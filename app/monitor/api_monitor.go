package monitor

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/houyanzu/eth-product/config"
	"github.com/houyanzu/eth-product/database/chainrecord"
	"github.com/houyanzu/eth-product/lib/httptool"
	"github.com/houyanzu/eth-product/tool/eth"
	"strings"
	"time"
)

type apiLogRes struct {
	Result []types.Log
}

func ApiMonitor(contract string, blockDiff uint64) (res EventLog, err error) {
	contract = strings.ToLower(contract)
	conf := config.GetConfig()

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
	netLastNum, err := eth.GetApiLastBlockNum()
	if err != nil {
		return
	}
	if netLastNum == 0 {
		err = errors.New("zero")
		return
	}
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

	str := strings.ReplaceAll(string(resp), `"logIndex":"0x"`, `"logIndex":"0x0"`)
	str = strings.ReplaceAll(str, `"transactionIndex":"0x"`, `"transactionIndex":"0x0"`)
	var logRes apiLogRes
	err = json.Unmarshal([]byte(str), &logRes)
	if err != nil {
		fmt.Println(url)
		return
	}
	res.logs = logRes.Result
	res.netLastNum = netLastNum
	res.endBlockNum = endBlockNum
	res.contract = contract
	return
}
