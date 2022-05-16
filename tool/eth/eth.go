package eth

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/houyanzu/eth-product/config"
	"github.com/houyanzu/eth-product/lib/contract/standardcoin"
	"github.com/houyanzu/eth-product/lib/contract/unipair"
	"github.com/shopspring/decimal"
	"math/big"
	"strings"
)

func BalanceOf[T any](conf config.Config[T], token, wallet string) (balance decimal.Decimal) {
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

func GetTxStatus[T any](conf config.Config[T], hash string) (status uint64, err error) {
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

func GetUniPrice[T any](conf config.Config[T], pair, token string, amount *big.Int) (price *big.Int, err error) {
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

	if tokenStr == token0Str {
		return quote(amount, reserves.Reserve0, reserves.Reserve1), nil
	} else if tokenStr == token1Str {
		return quote(amount, reserves.Reserve1, reserves.Reserve0), nil
	}

	return nil, errors.New("wrong token")
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
