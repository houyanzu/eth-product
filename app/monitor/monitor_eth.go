package monitor

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/houyanzu/eth-product/config"
	"github.com/houyanzu/eth-product/database/chainrecord"
	"math/big"
)

type EthLog struct {
	logs        []transferLog
	netLastNum  uint64
	endBlockNum uint64
	contract    string
}

type transferLog struct {
	TxHash      common.Hash
	BlockNumber uint64
	From        common.Address
	To          common.Address
	Amount      *big.Int
}

func MonitorEth(blockDiff uint64) (res EthLog, err error) {
	contract := "0x0000000000000000000000000000000000000000"
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
	startBlockNum := lastBlockNum + 1
	endBlockNum := lastBlockNum + blockDiff

	logs := make([]transferLog, 0)
	for i := startBlockNum; i <= endBlockNum; i++ {
		var block *types.Block
		block, err = client.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		if err != nil {
			return
		}
		for _, tx := range block.Transactions() {
			msg, _ := tx.AsMessage(types.NewEIP155Signer(big.NewInt(conf.Eth.ChainId)), nil)
			logs = append(logs, transferLog{tx.Hash(), i, msg.From(), *tx.To(), tx.Value()})
		}
	}

	res.logs = logs
	res.netLastNum = netLastNum
	res.endBlockNum = endBlockNum
	res.contract = contract
	return
}

func (e EthLog) Foreach(f func(index int, log transferLog)) {
	for k, v := range e.logs {
		blockNum := v.BlockNumber
		hash := v.TxHash.Hex()
		record := chainrecord.New(nil)
		record.Data.Contract = e.contract
		record.Data.BlockNum = blockNum
		record.Data.EventId = ""
		record.Data.Hash = hash
		record.Add()
		f(k, v)
	}
	if e.endBlockNum <= e.netLastNum {
		record := chainrecord.New(nil)
		record.Data.Contract = e.contract
		record.Data.BlockNum = e.endBlockNum
		record.Add()
	}
}
