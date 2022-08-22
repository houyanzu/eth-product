package monitor

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/houyanzu/eth-product/config"
	"github.com/houyanzu/eth-product/lib/contract/standardcoin"
	"github.com/houyanzu/eth-product/lib/httptool"
	"strings"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	fmt.Println(len(logTransferSigHash.Hex()))
}

func TestApiMonitor(t *testing.T) {
	err := config.ParseConfigByFile("D:\\work\\gowork\\eth-product\\config\\config.json")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()

	contract := strings.ToLower("0xe561479bebee0e606c19bb1973fc4761613e3c42")
	conf := config.GetConfig()
	lastBlockNum := uint64(4993830)
	endBlockNum := uint64(4993832)

	url := conf.Eth.ApiHost +
		"?module=logs&action=getLogs" +
		"&fromBlock=" + fmt.Sprintf("%d", lastBlockNum+1) +
		"&toBlock=" + fmt.Sprintf("%d", endBlockNum) +
		"&address=" + contract +
		"&apikey=" + conf.Eth.ApiKey
	resp, code, err := httptool.Get(url, 20*time.Second)
	if err != nil {
		return
	}
	if code != 200 {
		err = errors.New(string(resp))
		return
	}

	var logRes apiLogRes
	err = json.Unmarshal(resp, &logRes)
	if err != nil {
		return
	}
	con, err := standardcoin.NewStandardcoin(common.HexToAddress(""), nil)
	if err != nil {
		return
	}
	for _, v := range logRes.Result {
		fmt.Println(v.TxHash.Hex())
		trans, err := con.ParseTransfer(v)
		if err != nil {
			return
		}
		fmt.Println(trans.From.Hex())
		fmt.Println(trans.To.Hex())
		fmt.Println(trans.Value.String())
		break
	}

}
