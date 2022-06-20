// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package locktransfer

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

// LocktransferABI is the input ABI used to generate the binding from.
const LocktransferABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transferAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseStartTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseCycle\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseTimes\",\"type\":\"uint256\"}],\"name\":\"LockTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOperator\",\"type\":\"address\"}],\"name\":\"OperatorAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOperator\",\"type\":\"address\"}],\"name\":\"OperatorDel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"delOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getLockAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releaseStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releaseCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releaseTimes\",\"type\":\"uint256\"}],\"name\":\"lockTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Locktransfer is an auto generated Go binding around an Ethereum contract.
type Locktransfer struct {
	LocktransferCaller     // Read-only binding to the contract
	LocktransferTransactor // Write-only binding to the contract
	LocktransferFilterer   // Log filterer for contract events
}

// LocktransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type LocktransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LocktransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LocktransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LocktransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LocktransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LocktransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LocktransferSession struct {
	Contract     *Locktransfer     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LocktransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LocktransferCallerSession struct {
	Contract *LocktransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LocktransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LocktransferTransactorSession struct {
	Contract     *LocktransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LocktransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type LocktransferRaw struct {
	Contract *Locktransfer // Generic contract binding to access the raw methods on
}

// LocktransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LocktransferCallerRaw struct {
	Contract *LocktransferCaller // Generic read-only contract binding to access the raw methods on
}

// LocktransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LocktransferTransactorRaw struct {
	Contract *LocktransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLocktransfer creates a new instance of Locktransfer, bound to a specific deployed contract.
func NewLocktransfer(address common.Address, backend bind.ContractBackend) (*Locktransfer, error) {
	contract, err := bindLocktransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Locktransfer{LocktransferCaller: LocktransferCaller{contract: contract}, LocktransferTransactor: LocktransferTransactor{contract: contract}, LocktransferFilterer: LocktransferFilterer{contract: contract}}, nil
}

// NewLocktransferCaller creates a new read-only instance of Locktransfer, bound to a specific deployed contract.
func NewLocktransferCaller(address common.Address, caller bind.ContractCaller) (*LocktransferCaller, error) {
	contract, err := bindLocktransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LocktransferCaller{contract: contract}, nil
}

// NewLocktransferTransactor creates a new write-only instance of Locktransfer, bound to a specific deployed contract.
func NewLocktransferTransactor(address common.Address, transactor bind.ContractTransactor) (*LocktransferTransactor, error) {
	contract, err := bindLocktransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LocktransferTransactor{contract: contract}, nil
}

// NewLocktransferFilterer creates a new log filterer instance of Locktransfer, bound to a specific deployed contract.
func NewLocktransferFilterer(address common.Address, filterer bind.ContractFilterer) (*LocktransferFilterer, error) {
	contract, err := bindLocktransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LocktransferFilterer{contract: contract}, nil
}

// bindLocktransfer binds a generic wrapper to an already deployed contract.
func bindLocktransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LocktransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Locktransfer *LocktransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Locktransfer.Contract.LocktransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Locktransfer *LocktransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Locktransfer.Contract.LocktransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Locktransfer *LocktransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Locktransfer.Contract.LocktransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Locktransfer *LocktransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Locktransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Locktransfer *LocktransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Locktransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Locktransfer *LocktransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Locktransfer.Contract.contract.Transact(opts, method, params...)
}

// GetLockAmount is a free data retrieval call binding the contract method 0x339fd959.
//
// Solidity: function getLockAmount(address token, address account) view returns(uint256)
func (_Locktransfer *LocktransferCaller) GetLockAmount(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Locktransfer.contract.Call(opts, &out, "getLockAmount", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockAmount is a free data retrieval call binding the contract method 0x339fd959.
//
// Solidity: function getLockAmount(address token, address account) view returns(uint256)
func (_Locktransfer *LocktransferSession) GetLockAmount(token common.Address, account common.Address) (*big.Int, error) {
	return _Locktransfer.Contract.GetLockAmount(&_Locktransfer.CallOpts, token, account)
}

// GetLockAmount is a free data retrieval call binding the contract method 0x339fd959.
//
// Solidity: function getLockAmount(address token, address account) view returns(uint256)
func (_Locktransfer *LocktransferCallerSession) GetLockAmount(token common.Address, account common.Address) (*big.Int, error) {
	return _Locktransfer.Contract.GetLockAmount(&_Locktransfer.CallOpts, token, account)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_Locktransfer *LocktransferCaller) IsOperator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Locktransfer.contract.Call(opts, &out, "isOperator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_Locktransfer *LocktransferSession) IsOperator(account common.Address) (bool, error) {
	return _Locktransfer.Contract.IsOperator(&_Locktransfer.CallOpts, account)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_Locktransfer *LocktransferCallerSession) IsOperator(account common.Address) (bool, error) {
	return _Locktransfer.Contract.IsOperator(&_Locktransfer.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Locktransfer *LocktransferCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Locktransfer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Locktransfer *LocktransferSession) Owner() (common.Address, error) {
	return _Locktransfer.Contract.Owner(&_Locktransfer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Locktransfer *LocktransferCallerSession) Owner() (common.Address, error) {
	return _Locktransfer.Contract.Owner(&_Locktransfer.CallOpts)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_Locktransfer *LocktransferTransactor) AddOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _Locktransfer.contract.Transact(opts, "addOperator", newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_Locktransfer *LocktransferSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _Locktransfer.Contract.AddOperator(&_Locktransfer.TransactOpts, newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_Locktransfer *LocktransferTransactorSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _Locktransfer.Contract.AddOperator(&_Locktransfer.TransactOpts, newOperator)
}

// DelOperator is a paid mutator transaction binding the contract method 0x3e30838d.
//
// Solidity: function delOperator(address operator) returns()
func (_Locktransfer *LocktransferTransactor) DelOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Locktransfer.contract.Transact(opts, "delOperator", operator)
}

// DelOperator is a paid mutator transaction binding the contract method 0x3e30838d.
//
// Solidity: function delOperator(address operator) returns()
func (_Locktransfer *LocktransferSession) DelOperator(operator common.Address) (*types.Transaction, error) {
	return _Locktransfer.Contract.DelOperator(&_Locktransfer.TransactOpts, operator)
}

// DelOperator is a paid mutator transaction binding the contract method 0x3e30838d.
//
// Solidity: function delOperator(address operator) returns()
func (_Locktransfer *LocktransferTransactorSession) DelOperator(operator common.Address) (*types.Transaction, error) {
	return _Locktransfer.Contract.DelOperator(&_Locktransfer.TransactOpts, operator)
}

// LockTransfer is a paid mutator transaction binding the contract method 0xb3cf65df.
//
// Solidity: function lockTransfer(address token, address to, uint256 amount, uint256 releaseStartTime, uint256 releaseCycle, uint256 releaseTimes) returns(bool)
func (_Locktransfer *LocktransferTransactor) LockTransfer(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, releaseStartTime *big.Int, releaseCycle *big.Int, releaseTimes *big.Int) (*types.Transaction, error) {
	return _Locktransfer.contract.Transact(opts, "lockTransfer", token, to, amount, releaseStartTime, releaseCycle, releaseTimes)
}

// LockTransfer is a paid mutator transaction binding the contract method 0xb3cf65df.
//
// Solidity: function lockTransfer(address token, address to, uint256 amount, uint256 releaseStartTime, uint256 releaseCycle, uint256 releaseTimes) returns(bool)
func (_Locktransfer *LocktransferSession) LockTransfer(token common.Address, to common.Address, amount *big.Int, releaseStartTime *big.Int, releaseCycle *big.Int, releaseTimes *big.Int) (*types.Transaction, error) {
	return _Locktransfer.Contract.LockTransfer(&_Locktransfer.TransactOpts, token, to, amount, releaseStartTime, releaseCycle, releaseTimes)
}

// LockTransfer is a paid mutator transaction binding the contract method 0xb3cf65df.
//
// Solidity: function lockTransfer(address token, address to, uint256 amount, uint256 releaseStartTime, uint256 releaseCycle, uint256 releaseTimes) returns(bool)
func (_Locktransfer *LocktransferTransactorSession) LockTransfer(token common.Address, to common.Address, amount *big.Int, releaseStartTime *big.Int, releaseCycle *big.Int, releaseTimes *big.Int) (*types.Transaction, error) {
	return _Locktransfer.Contract.LockTransfer(&_Locktransfer.TransactOpts, token, to, amount, releaseStartTime, releaseCycle, releaseTimes)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Locktransfer *LocktransferTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Locktransfer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Locktransfer *LocktransferSession) RenounceOwnership() (*types.Transaction, error) {
	return _Locktransfer.Contract.RenounceOwnership(&_Locktransfer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Locktransfer *LocktransferTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Locktransfer.Contract.RenounceOwnership(&_Locktransfer.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Locktransfer *LocktransferTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Locktransfer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Locktransfer *LocktransferSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Locktransfer.Contract.TransferOwnership(&_Locktransfer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Locktransfer *LocktransferTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Locktransfer.Contract.TransferOwnership(&_Locktransfer.TransactOpts, newOwner)
}

// LocktransferLockTransferIterator is returned from FilterLockTransfer and is used to iterate over the raw logs and unpacked data for LockTransfer events raised by the Locktransfer contract.
type LocktransferLockTransferIterator struct {
	Event *LocktransferLockTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LocktransferLockTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LocktransferLockTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LocktransferLockTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LocktransferLockTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LocktransferLockTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LocktransferLockTransfer represents a LockTransfer event raised by the Locktransfer contract.
type LocktransferLockTransfer struct {
	Token            common.Address
	To               common.Address
	Amount           *big.Int
	TransferAmount   *big.Int
	ReleaseStartTime *big.Int
	ReleaseCycle     *big.Int
	ReleaseTimes     *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLockTransfer is a free log retrieval operation binding the contract event 0x2586d1dca3391bd657d4003132d5a4a3b8e2662be5d6973104f496bbe794722e.
//
// Solidity: event LockTransfer(address indexed token, address indexed to, uint256 amount, uint256 transferAmount, uint256 releaseStartTime, uint256 releaseCycle, uint256 releaseTimes)
func (_Locktransfer *LocktransferFilterer) FilterLockTransfer(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*LocktransferLockTransferIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Locktransfer.contract.FilterLogs(opts, "LockTransfer", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LocktransferLockTransferIterator{contract: _Locktransfer.contract, event: "LockTransfer", logs: logs, sub: sub}, nil
}

// WatchLockTransfer is a free log subscription operation binding the contract event 0x2586d1dca3391bd657d4003132d5a4a3b8e2662be5d6973104f496bbe794722e.
//
// Solidity: event LockTransfer(address indexed token, address indexed to, uint256 amount, uint256 transferAmount, uint256 releaseStartTime, uint256 releaseCycle, uint256 releaseTimes)
func (_Locktransfer *LocktransferFilterer) WatchLockTransfer(opts *bind.WatchOpts, sink chan<- *LocktransferLockTransfer, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Locktransfer.contract.WatchLogs(opts, "LockTransfer", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LocktransferLockTransfer)
				if err := _Locktransfer.contract.UnpackLog(event, "LockTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLockTransfer is a log parse operation binding the contract event 0x2586d1dca3391bd657d4003132d5a4a3b8e2662be5d6973104f496bbe794722e.
//
// Solidity: event LockTransfer(address indexed token, address indexed to, uint256 amount, uint256 transferAmount, uint256 releaseStartTime, uint256 releaseCycle, uint256 releaseTimes)
func (_Locktransfer *LocktransferFilterer) ParseLockTransfer(log types.Log) (*LocktransferLockTransfer, error) {
	event := new(LocktransferLockTransfer)
	if err := _Locktransfer.contract.UnpackLog(event, "LockTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LocktransferOperatorAddIterator is returned from FilterOperatorAdd and is used to iterate over the raw logs and unpacked data for OperatorAdd events raised by the Locktransfer contract.
type LocktransferOperatorAddIterator struct {
	Event *LocktransferOperatorAdd // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LocktransferOperatorAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LocktransferOperatorAdd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LocktransferOperatorAdd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LocktransferOperatorAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LocktransferOperatorAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LocktransferOperatorAdd represents a OperatorAdd event raised by the Locktransfer contract.
type LocktransferOperatorAdd struct {
	OldOperator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdd is a free log retrieval operation binding the contract event 0xb8a01e9b4715127016e813d6838a2752d018a19920f431467f17cc3c31df0686.
//
// Solidity: event OperatorAdd(address indexed oldOperator)
func (_Locktransfer *LocktransferFilterer) FilterOperatorAdd(opts *bind.FilterOpts, oldOperator []common.Address) (*LocktransferOperatorAddIterator, error) {

	var oldOperatorRule []interface{}
	for _, oldOperatorItem := range oldOperator {
		oldOperatorRule = append(oldOperatorRule, oldOperatorItem)
	}

	logs, sub, err := _Locktransfer.contract.FilterLogs(opts, "OperatorAdd", oldOperatorRule)
	if err != nil {
		return nil, err
	}
	return &LocktransferOperatorAddIterator{contract: _Locktransfer.contract, event: "OperatorAdd", logs: logs, sub: sub}, nil
}

// WatchOperatorAdd is a free log subscription operation binding the contract event 0xb8a01e9b4715127016e813d6838a2752d018a19920f431467f17cc3c31df0686.
//
// Solidity: event OperatorAdd(address indexed oldOperator)
func (_Locktransfer *LocktransferFilterer) WatchOperatorAdd(opts *bind.WatchOpts, sink chan<- *LocktransferOperatorAdd, oldOperator []common.Address) (event.Subscription, error) {

	var oldOperatorRule []interface{}
	for _, oldOperatorItem := range oldOperator {
		oldOperatorRule = append(oldOperatorRule, oldOperatorItem)
	}

	logs, sub, err := _Locktransfer.contract.WatchLogs(opts, "OperatorAdd", oldOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LocktransferOperatorAdd)
				if err := _Locktransfer.contract.UnpackLog(event, "OperatorAdd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorAdd is a log parse operation binding the contract event 0xb8a01e9b4715127016e813d6838a2752d018a19920f431467f17cc3c31df0686.
//
// Solidity: event OperatorAdd(address indexed oldOperator)
func (_Locktransfer *LocktransferFilterer) ParseOperatorAdd(log types.Log) (*LocktransferOperatorAdd, error) {
	event := new(LocktransferOperatorAdd)
	if err := _Locktransfer.contract.UnpackLog(event, "OperatorAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LocktransferOperatorDelIterator is returned from FilterOperatorDel and is used to iterate over the raw logs and unpacked data for OperatorDel events raised by the Locktransfer contract.
type LocktransferOperatorDelIterator struct {
	Event *LocktransferOperatorDel // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LocktransferOperatorDelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LocktransferOperatorDel)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LocktransferOperatorDel)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LocktransferOperatorDelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LocktransferOperatorDelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LocktransferOperatorDel represents a OperatorDel event raised by the Locktransfer contract.
type LocktransferOperatorDel struct {
	OldOperator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorDel is a free log retrieval operation binding the contract event 0x26d80e044a097a0a4a866f16db1a525159d516fd5e0f157e94acb4f0f5d86c95.
//
// Solidity: event OperatorDel(address indexed oldOperator)
func (_Locktransfer *LocktransferFilterer) FilterOperatorDel(opts *bind.FilterOpts, oldOperator []common.Address) (*LocktransferOperatorDelIterator, error) {

	var oldOperatorRule []interface{}
	for _, oldOperatorItem := range oldOperator {
		oldOperatorRule = append(oldOperatorRule, oldOperatorItem)
	}

	logs, sub, err := _Locktransfer.contract.FilterLogs(opts, "OperatorDel", oldOperatorRule)
	if err != nil {
		return nil, err
	}
	return &LocktransferOperatorDelIterator{contract: _Locktransfer.contract, event: "OperatorDel", logs: logs, sub: sub}, nil
}

// WatchOperatorDel is a free log subscription operation binding the contract event 0x26d80e044a097a0a4a866f16db1a525159d516fd5e0f157e94acb4f0f5d86c95.
//
// Solidity: event OperatorDel(address indexed oldOperator)
func (_Locktransfer *LocktransferFilterer) WatchOperatorDel(opts *bind.WatchOpts, sink chan<- *LocktransferOperatorDel, oldOperator []common.Address) (event.Subscription, error) {

	var oldOperatorRule []interface{}
	for _, oldOperatorItem := range oldOperator {
		oldOperatorRule = append(oldOperatorRule, oldOperatorItem)
	}

	logs, sub, err := _Locktransfer.contract.WatchLogs(opts, "OperatorDel", oldOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LocktransferOperatorDel)
				if err := _Locktransfer.contract.UnpackLog(event, "OperatorDel", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorDel is a log parse operation binding the contract event 0x26d80e044a097a0a4a866f16db1a525159d516fd5e0f157e94acb4f0f5d86c95.
//
// Solidity: event OperatorDel(address indexed oldOperator)
func (_Locktransfer *LocktransferFilterer) ParseOperatorDel(log types.Log) (*LocktransferOperatorDel, error) {
	event := new(LocktransferOperatorDel)
	if err := _Locktransfer.contract.UnpackLog(event, "OperatorDel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LocktransferOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Locktransfer contract.
type LocktransferOwnershipTransferredIterator struct {
	Event *LocktransferOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LocktransferOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LocktransferOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LocktransferOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LocktransferOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LocktransferOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LocktransferOwnershipTransferred represents a OwnershipTransferred event raised by the Locktransfer contract.
type LocktransferOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Locktransfer *LocktransferFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LocktransferOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Locktransfer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LocktransferOwnershipTransferredIterator{contract: _Locktransfer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Locktransfer *LocktransferFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LocktransferOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Locktransfer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LocktransferOwnershipTransferred)
				if err := _Locktransfer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Locktransfer *LocktransferFilterer) ParseOwnershipTransferred(log types.Log) (*LocktransferOwnershipTransferred, error) {
	event := new(LocktransferOwnershipTransferred)
	if err := _Locktransfer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
