package eth

import (
	"fmt"
	"github.com/houyanzu/eth-product/config"
	"github.com/shopspring/decimal"
	"testing"
)

func TestGetGasFeeByHash(t *testing.T) {
	err := config.ParseConfigByFile("D:\\work\\gowork\\eth-product\\config\\config.json")
	if err != nil {
		panic(err)
	}
	gas, err := GetGasFeeByHash("0xb9e7ec0d2a44552a3e99a37f2a04ced59dc12903003eac137bdcc3b2567d0309")
	if err != nil {
		panic(err)
	}
	fmt.Println(gas.Div(decimal.New(1, 18)).String())
}
