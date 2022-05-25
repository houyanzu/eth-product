package monitor

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/houyanzu/eth-product/config"
	"github.com/houyanzu/eth-product/database/chainrecord"
	"math/big"
	"strings"
)

type EthLog struct {
	logs        []types.Log
	netLastNum  uint64
	endBlockNum uint64
	contract    string
}

var initBlock = make(map[string]uint64)

func InitBlockNum(contract string, blockNum uint64) {
	contract = strings.ToLower(contract)
	initBlock[contract] = blockNum
}

func Monitor(contract string, blockDiff uint64) (res EthLog, err error) {
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

	contractAddress := common.HexToAddress(contract)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(lastBlockNum + 1)),
		ToBlock:   big.NewInt(int64(endBlockNum)),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return
	}
	res.logs = logs
	res.netLastNum = netLastNum
	res.endBlockNum = endBlockNum
	res.contract = contract
	return
}

func (e EthLog) Foreach(f func(index int, log types.Log)) {
	for k, v := range e.logs {
		blockNum := v.BlockNumber
		hash := v.TxHash.Hex()
		record := chainrecord.New(nil)
		record.Data.Contract = e.contract
		record.Data.BlockNum = blockNum
		record.Data.EventId = v.Topics[0].Hex()
		record.Data.Hash = hash
		record.Add()
		f(k, v)
	}
	if e.endBlockNum < e.netLastNum {
		record := chainrecord.New(nil)
		record.Data.Contract = e.contract
		record.Data.BlockNum = e.endBlockNum
		record.Add()
	}
}
