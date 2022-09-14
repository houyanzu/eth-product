package eth

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/houyanzu/eth-product/config"
	"github.com/houyanzu/eth-product/lib/contract/standardcoin"
	"github.com/houyanzu/eth-product/lib/contract/unipair"
	"github.com/houyanzu/eth-product/lib/httptool"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GetClientAndAuth(priKey string, gasLimit uint64, value *big.Int) (client *ethclient.Client, auth *bind.TransactOpts, err error) {
	conf := config.GetConfig()
	client, err = ethclient.Dial(conf.Eth.Host)
	if err != nil {
		return
	}

	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		return
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	nonce, err := client.NonceAt(context.Background(), common.HexToAddress(fromAddress), nil)
	if err != nil {
		return
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(conf.Eth.ChainId))
	if err != nil {
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value       // in wei
	auth.GasLimit = gasLimit // in units
	auth.GasPrice = gasPrice
	return
}

func BalanceOf(token, wallet string) (balance decimal.Decimal, err error) {
	conf := config.GetConfig()
	balance = decimal.Zero
	client, err := ethclient.Dial(conf.Eth.Host)
	if err != nil {
		return
	}

	coin, err := standardcoin.NewStandardcoin(common.HexToAddress(token), client)
	if err != nil {
		return
	}

	ba, err := coin.BalanceOf(nil, common.HexToAddress(wallet))
	if err != nil {
		return
	}
	balance = decimal.NewFromBigInt(ba, 0)
	return
}

func TokenSymbol(token string) (res string, err error) {
	conf := config.GetConfig()
	client, err := ethclient.Dial(conf.Eth.Host)
	if err != nil {
		return
	}

	coin, err := standardcoin.NewStandardcoin(common.HexToAddress(token), client)
	if err != nil {
		return
	}

	res, err = coin.Symbol(nil)
	if err != nil {
		return
	}
	return
}

func TokenName(token string) (res string, err error) {
	conf := config.GetConfig()
	client, err := ethclient.Dial(conf.Eth.Host)
	if err != nil {
		return
	}

	coin, err := standardcoin.NewStandardcoin(common.HexToAddress(token), client)
	if err != nil {
		return
	}

	res, err = coin.Name(nil)
	if err != nil {
		return
	}
	return
}

func TokenTotalSupply(token string) (res decimal.Decimal, err error) {
	conf := config.GetConfig()
	client, err := ethclient.Dial(conf.Eth.Host)
	if err != nil {
		return
	}

	coin, err := standardcoin.NewStandardcoin(common.HexToAddress(token), client)
	if err != nil {
		return
	}

	resBig, err := coin.TotalSupply(nil)
	if err != nil {
		return
	}
	res = decimal.NewFromBigInt(resBig, 0)
	return
}

func TokenDecimals(token string) (res uint8, err error) {
	conf := config.GetConfig()
	client, err := ethclient.Dial(conf.Eth.Host)
	if err != nil {
		return
	}

	coin, err := standardcoin.NewStandardcoin(common.HexToAddress(token), client)
	if err != nil {
		return
	}

	res, err = coin.Decimals(nil)
	if err != nil {
		return
	}
	return
}

func BalanceAt(addr string) (balance decimal.Decimal, err error) {
	balance = decimal.Zero
	conf := config.GetConfig()
	client, err := ethclient.Dial(conf.Eth.Host)
	if err != nil {
		return
	}

	account := common.HexToAddress(addr)
	ba, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return
	}

	balance = decimal.NewFromBigInt(ba, 0)
	return
}

func GetTxStatus(hash string) (status uint64, err error) {
	conf := config.GetConfig()
	client, err := ethclient.Dial(conf.Eth.Host)
	if err != nil {
		return
	}
	txHash := common.HexToHash(hash)
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return
	}
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return
	}
	status = receipt.Status
	return
}

func GetUniPrice(pair, token string, amount decimal.Decimal) (price decimal.Decimal, err error) {
	conf := config.GetConfig()
	client, err := ethclient.Dial(conf.Eth.Host)
	if err != nil {
		return
	}

	uniPair, err := unipair.NewUnipair(common.HexToAddress(pair), client)
	if err != nil {
		return
	}

	token0, err := uniPair.Token0(nil)
	if err != nil {
		return
	}
	token1, err := uniPair.Token1(nil)
	if err != nil {
		return
	}
	token0Str := strings.ToLower(token0.Hex())
	token1Str := strings.ToLower(token1.Hex())
	tokenStr := strings.ToLower(token)

	reserves, err := uniPair.GetReserves(nil)

	amountBig := amount.BigInt()

	if tokenStr == token0Str {
		res := quote(amountBig, reserves.Reserve0, reserves.Reserve1)
		return decimal.NewFromBigInt(res, 0), nil
	} else if tokenStr == token1Str {
		res := quote(amountBig, reserves.Reserve1, reserves.Reserve0)
		return decimal.NewFromBigInt(res, 0), nil
	}

	return decimal.Zero, errors.New("wrong token")
}

func IsAddress(addr string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(addr)
}

func IsContract(addr string) (res bool, err error) {
	if !IsAddress(addr) {
		return false, nil
	}
	conf := config.GetConfig()
	client, err := ethclient.Dial(conf.Eth.Host)
	if err != nil {
		return
	}
	address := common.HexToAddress(addr)
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	res = len(bytecode) > 0
	return
}

func GetGasFeeByHash(hash string) (decimal.Decimal, error) {
	conf := config.GetConfig()

	client, err := ethclient.Dial(conf.Eth.Host)
	if err != nil {
		return decimal.Zero, err
	}
	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		return decimal.Decimal{}, err
	}
	gasPrice := decimal.NewFromBigInt(tx.GasPrice(), 0)

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err != nil {
		return decimal.Decimal{}, err
	}
	gasUsedInt := receipt.GasUsed
	gasUsed, _ := decimal.NewFromString(fmt.Sprintf("%d", gasUsedInt))
	gasFee := gasPrice.Mul(gasUsed)
	return gasFee, nil
}

func GetApiLastBlockNum() (num uint64, err error) {
	var res struct {
		Status string `json:"status"`
		Result string `json:"result"`
	}
	conf := config.GetConfig()
	now := time.Now().Unix() - 5
	url := conf.Eth.ApiHost +
		"?module=block&action=getblocknobytime" +
		"&timestamp=" + fmt.Sprintf("%d", now) +
		"&closest=before" +
		"&apikey=" + conf.Eth.ApiKey
	resp, code, err := httptool.Get(url, 20*time.Second)
	if err != nil {
		return
	}
	if code != 200 {
		err = errors.New(string(resp))
		return
	}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	return strconv.ParseUint(res.Result, 10, 64)
}

func CreateAddress() (address, privateKeyString string, err error) {
	privateKey, _ := crypto.GenerateKey()
	if err != nil {
		return
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyString = hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}

	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return
}

func quote(amountA, reserveA, reserveB *big.Int) *big.Int {
	if amountA.Cmp(big.NewInt(0)) != 1 {
		return big.NewInt(0)
	}
	if reserveA.Cmp(big.NewInt(0)) != 1 || reserveB.Cmp(big.NewInt(0)) != 1 {
		return big.NewInt(0)
	}

	result := big.NewInt(0)
	result.Mul(amountA, reserveB).Div(result, reserveA)
	return result
}
