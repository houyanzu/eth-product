package transfer

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/houyanzu/eth-product/config"
	"github.com/houyanzu/eth-product/database/pwdwt"
	"github.com/houyanzu/eth-product/database/transferdetails"
	"github.com/houyanzu/eth-product/database/transferrecords"
	"github.com/houyanzu/eth-product/lib/contract/multitransfer"
	"github.com/houyanzu/eth-product/lib/crypto/aes"
	"github.com/houyanzu/eth-product/tool/eth"
	"math/big"
)

var privateKeyStr string

func InitTrans(priKeyCt aes.Decoder, password []byte) (e error) {
	defer func() {
		err := recover()
		if err != nil {
			pwdwt.New(nil).Wrong()
			e = errors.New("wrong password")
			return
		}
	}()
	times := pwdwt.New(nil).GetTimes()
	if times >= 5 {
		e = errors.New("locked")
		return
	}
	privateKeyByte := priKeyCt.Decode(password)
	privateKeyStr = privateKeyByte.ToString()
	pwdwt.New(nil).ResetTimes()
	return
}

func Transfer(limit int) (err error) {
	conf := config.GetConfig()
	pending := transferrecords.New(nil).InitPending()
	if pending.Data.ID > 0 {
		var status uint64
		status, err = eth.GetTxStatus(pending.Data.Hash)
		if err != nil {
			//TODO:覆盖操作
			return
		}
		if status == 1 {
			pending.SetSuccess()
			transferdetails.New(nil).SetSuccess(pending.Data.ID)
		} else if status == 0 {
			pending.SetFail()
			transferdetails.New(nil).SetFail(pending.Data.ID)
		}
	}

	waitingList := transferdetails.New(nil).InitWaitingList(limit)
	length := len(waitingList.List)
	if length == 0 {
		return
	}

	tokens, tos := make([]common.Address, length), make([]common.Address, length)
	ids := make([]uint, length)
	amounts := make([]*big.Int, length)
	totalValue := big.NewInt(0)
	waitingList.Foreach(func(index int, m *transferdetails.Model) {
		tokens[index] = common.HexToAddress(m.Data.Token)
		tos[index] = common.HexToAddress(m.Data.To)
		amount := m.Data.Amount.BigInt()
		amounts[index] = amount
		ids[index] = m.Data.ID
		if m.Data.Token == "0x0000000000000000000000000000000000000000" {
			totalValue.Add(totalValue, amount)
		}
	})
	client, err := ethclient.Dial(conf.Eth.Host)
	if err != nil {
		return
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.NonceAt(context.Background(), fromAddress, nil)
	if err != nil {
		return
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(conf.Eth.ChainId))
	if err != nil {
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = totalValue         // in wei
	auth.GasLimit = uint64(2000000) // in units
	auth.GasPrice = gasPrice

	multiCon := common.HexToAddress(conf.Eth.MultiTransferContract)
	multiTransferInstance, err := multitransfer.NewMultitransfer(multiCon, client)
	if err != nil {
		return
	}
	tx, err := multiTransferInstance.MultiTransferToken(auth, tokens, tos, amounts)
	if err != nil {
		return
	}
	hash := tx.Hash().Hex()

	tr := transferrecords.New(nil)
	tr.Data.Hash = hash
	tr.Data.Nonce = nonce
	tr.Add()

	transferdetails.New(nil).SetExec(ids, tr.Data.ID)
	return
}
