package monitor

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestName(t *testing.T) {
	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	fmt.Println(len(logTransferSigHash.Hex()))
}
