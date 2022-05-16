package multitransfer

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MultitransferABI is the input ABI used to generate the binding from.
const MultitransferABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_dsts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"multiTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_dsts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"multiTransferToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Multitransfer is an auto generated Go binding around an Ethereum contract.
type Multitransfer struct {
	MultitransferCaller     // Read-only binding to the contract
	MultitransferTransactor // Write-only binding to the contract
	MultitransferFilterer   // Log filterer for contract events
}

// MultitransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultitransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultitransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultitransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultitransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultitransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultitransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultitransferSession struct {
	Contract     *Multitransfer    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultitransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultitransferCallerSession struct {
	Contract *MultitransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MultitransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultitransferTransactorSession struct {
	Contract     *MultitransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MultitransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultitransferRaw struct {
	Contract *Multitransfer // Generic contract binding to access the raw methods on
}

// MultitransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultitransferCallerRaw struct {
	Contract *MultitransferCaller // Generic read-only contract binding to access the raw methods on
}

// MultitransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultitransferTransactorRaw struct {
	Contract *MultitransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultitransfer creates a new instance of Multitransfer, bound to a specific deployed contract.
func NewMultitransfer(address common.Address, backend bind.ContractBackend) (*Multitransfer, error) {
	contract, err := bindMultitransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multitransfer{MultitransferCaller: MultitransferCaller{contract: contract}, MultitransferTransactor: MultitransferTransactor{contract: contract}, MultitransferFilterer: MultitransferFilterer{contract: contract}}, nil
}

// NewMultitransferCaller creates a new read-only instance of Multitransfer, bound to a specific deployed contract.
func NewMultitransferCaller(address common.Address, caller bind.ContractCaller) (*MultitransferCaller, error) {
	contract, err := bindMultitransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultitransferCaller{contract: contract}, nil
}

// NewMultitransferTransactor creates a new write-only instance of Multitransfer, bound to a specific deployed contract.
func NewMultitransferTransactor(address common.Address, transactor bind.ContractTransactor) (*MultitransferTransactor, error) {
	contract, err := bindMultitransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultitransferTransactor{contract: contract}, nil
}

// NewMultitransferFilterer creates a new log filterer instance of Multitransfer, bound to a specific deployed contract.
func NewMultitransferFilterer(address common.Address, filterer bind.ContractFilterer) (*MultitransferFilterer, error) {
	contract, err := bindMultitransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultitransferFilterer{contract: contract}, nil
}

// bindMultitransfer binds a generic wrapper to an already deployed contract.
func bindMultitransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultitransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multitransfer *MultitransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multitransfer.Contract.MultitransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multitransfer *MultitransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multitransfer.Contract.MultitransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multitransfer *MultitransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multitransfer.Contract.MultitransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multitransfer *MultitransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multitransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multitransfer *MultitransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multitransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multitransfer *MultitransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multitransfer.Contract.contract.Transact(opts, method, params...)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0x1e89d545.
//
// Solidity: function multiTransfer(address[] _dsts, uint256[] _values) payable returns()
func (_Multitransfer *MultitransferTransactor) MultiTransfer(opts *bind.TransactOpts, _dsts []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Multitransfer.contract.Transact(opts, "multiTransfer", _dsts, _values)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0x1e89d545.
//
// Solidity: function multiTransfer(address[] _dsts, uint256[] _values) payable returns()
func (_Multitransfer *MultitransferSession) MultiTransfer(_dsts []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Multitransfer.Contract.MultiTransfer(&_Multitransfer.TransactOpts, _dsts, _values)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0x1e89d545.
//
// Solidity: function multiTransfer(address[] _dsts, uint256[] _values) payable returns()
func (_Multitransfer *MultitransferTransactorSession) MultiTransfer(_dsts []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Multitransfer.Contract.MultiTransfer(&_Multitransfer.TransactOpts, _dsts, _values)
}

// MultiTransferToken is a paid mutator transaction binding the contract method 0xbba4b0ac.
//
// Solidity: function multiTransferToken(address[] _tokens, address[] _dsts, uint256[] _values) returns()
func (_Multitransfer *MultitransferTransactor) MultiTransferToken(opts *bind.TransactOpts, _tokens []common.Address, _dsts []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Multitransfer.contract.Transact(opts, "multiTransferToken", _tokens, _dsts, _values)
}

// MultiTransferToken is a paid mutator transaction binding the contract method 0xbba4b0ac.
//
// Solidity: function multiTransferToken(address[] _tokens, address[] _dsts, uint256[] _values) returns()
func (_Multitransfer *MultitransferSession) MultiTransferToken(_tokens []common.Address, _dsts []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Multitransfer.Contract.MultiTransferToken(&_Multitransfer.TransactOpts, _tokens, _dsts, _values)
}

// MultiTransferToken is a paid mutator transaction binding the contract method 0xbba4b0ac.
//
// Solidity: function multiTransferToken(address[] _tokens, address[] _dsts, uint256[] _values) returns()
func (_Multitransfer *MultitransferTransactorSession) MultiTransferToken(_tokens []common.Address, _dsts []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Multitransfer.Contract.MultiTransferToken(&_Multitransfer.TransactOpts, _tokens, _dsts, _values)
}
