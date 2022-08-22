package monitor

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/houyanzu/eth-product/config"
	"github.com/houyanzu/eth-product/database/chainrecord"
	"github.com/houyanzu/eth-product/lib/httptool"
	"strings"
	"time"
)

type apiLogRes struct {
	Result []types.Log
}

func ApiMonitor(contract string, blockDiff uint64) (res EventLog, err error) {
	contract = strings.ToLower(contract)
	conf := config.GetConfig()
	client, err := ethclient.Dial(conf.Eth.Host)
	if err != nil {
		return
	}

	lastBlockNum := chainrecord.GetLastBlockNum(contract)
	if lastBlockNum == 0 {
		var ok bool
		if lastBlockNum, ok = initBlock[contract]; ok {
			record := chainrecord.New(nil)
			record.Data.Contract = contract
			record.Data.BlockNum = lastBlockNum
			record.Data.EventId = ""
			record.Data.Hash = ""
			record.Add()
		} else {
			panic("未初始化")
		}
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return
	}
	netLastNum := header.Number.Uint64()
	endBlockNum := lastBlockNum + blockDiff

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
	res.logs = logRes.Result
	res.netLastNum = netLastNum
	res.endBlockNum = endBlockNum
	res.contract = contract
	return
}
