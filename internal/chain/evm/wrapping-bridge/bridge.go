// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappingBridge

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// WrappingBridgeMetaData contains all meta data concerning the WrappingBridge contract.
var WrappingBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_oracles\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"destChain\",\"type\":\"int256\"}],\"name\":\"Lock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destAddress\",\"type\":\"string\"}],\"name\":\"LockERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"CoinsUnlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"UnlockERC20\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_destAddress\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_destChain\",\"type\":\"int256\"}],\"name\":\"lock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_destAddress\",\"type\":\"string\"}],\"name\":\"lockERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unlockERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WrappingBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use WrappingBridgeMetaData.ABI instead.
var WrappingBridgeABI = WrappingBridgeMetaData.ABI

// WrappingBridge is an auto generated Go binding around an Ethereum contract.
type WrappingBridge struct {
	WrappingBridgeCaller     // Read-only binding to the contract
	WrappingBridgeTransactor // Write-only binding to the contract
	WrappingBridgeFilterer   // Log filterer for contract events
}

// WrappingBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrappingBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappingBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrappingBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappingBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WrappingBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappingBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrappingBridgeSession struct {
	Contract     *WrappingBridge   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WrappingBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrappingBridgeCallerSession struct {
	Contract *WrappingBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// WrappingBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrappingBridgeTransactorSession struct {
	Contract     *WrappingBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// WrappingBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrappingBridgeRaw struct {
	Contract *WrappingBridge // Generic contract binding to access the raw methods on
}

// WrappingBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrappingBridgeCallerRaw struct {
	Contract *WrappingBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// WrappingBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrappingBridgeTransactorRaw struct {
	Contract *WrappingBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrappingBridge creates a new instance of WrappingBridge, bound to a specific deployed contract.
func NewWrappingBridge(address common.Address, backend bind.ContractBackend) (*WrappingBridge, error) {
	contract, err := bindWrappingBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrappingBridge{WrappingBridgeCaller: WrappingBridgeCaller{contract: contract}, WrappingBridgeTransactor: WrappingBridgeTransactor{contract: contract}, WrappingBridgeFilterer: WrappingBridgeFilterer{contract: contract}}, nil
}

// NewWrappingBridgeCaller creates a new read-only instance of WrappingBridge, bound to a specific deployed contract.
func NewWrappingBridgeCaller(address common.Address, caller bind.ContractCaller) (*WrappingBridgeCaller, error) {
	contract, err := bindWrappingBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrappingBridgeCaller{contract: contract}, nil
}

// NewWrappingBridgeTransactor creates a new write-only instance of WrappingBridge, bound to a specific deployed contract.
func NewWrappingBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*WrappingBridgeTransactor, error) {
	contract, err := bindWrappingBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrappingBridgeTransactor{contract: contract}, nil
}

// NewWrappingBridgeFilterer creates a new log filterer instance of WrappingBridge, bound to a specific deployed contract.
func NewWrappingBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*WrappingBridgeFilterer, error) {
	contract, err := bindWrappingBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrappingBridgeFilterer{contract: contract}, nil
}

// bindWrappingBridge binds a generic wrapper to an already deployed contract.
func bindWrappingBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WrappingBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappingBridge *WrappingBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappingBridge.Contract.WrappingBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappingBridge *WrappingBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappingBridge.Contract.WrappingBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappingBridge *WrappingBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappingBridge.Contract.WrappingBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappingBridge *WrappingBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappingBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappingBridge *WrappingBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappingBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappingBridge *WrappingBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappingBridge.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WrappingBridge *WrappingBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrappingBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WrappingBridge *WrappingBridgeSession) Owner() (common.Address, error) {
	return _WrappingBridge.Contract.Owner(&_WrappingBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WrappingBridge *WrappingBridgeCallerSession) Owner() (common.Address, error) {
	return _WrappingBridge.Contract.Owner(&_WrappingBridge.CallOpts)
}

// Lock is a paid mutator transaction binding the contract method 0x42d09351.
//
// Solidity: function lock(string _destAddress, int256 _destChain) payable returns(bool success)
func (_WrappingBridge *WrappingBridgeTransactor) Lock(opts *bind.TransactOpts, _destAddress string, _destChain *big.Int) (*types.Transaction, error) {
	return _WrappingBridge.contract.Transact(opts, "lock", _destAddress, _destChain)
}

// Lock is a paid mutator transaction binding the contract method 0x42d09351.
//
// Solidity: function lock(string _destAddress, int256 _destChain) payable returns(bool success)
func (_WrappingBridge *WrappingBridgeSession) Lock(_destAddress string, _destChain *big.Int) (*types.Transaction, error) {
	return _WrappingBridge.Contract.Lock(&_WrappingBridge.TransactOpts, _destAddress, _destChain)
}

// Lock is a paid mutator transaction binding the contract method 0x42d09351.
//
// Solidity: function lock(string _destAddress, int256 _destChain) payable returns(bool success)
func (_WrappingBridge *WrappingBridgeTransactorSession) Lock(_destAddress string, _destChain *big.Int) (*types.Transaction, error) {
	return _WrappingBridge.Contract.Lock(&_WrappingBridge.TransactOpts, _destAddress, _destChain)
}

// LockERC20 is a paid mutator transaction binding the contract method 0xabb18184.
//
// Solidity: function lockERC20(address _token, uint256 _amount, string _destAddress) returns(bool success)
func (_WrappingBridge *WrappingBridgeTransactor) LockERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _destAddress string) (*types.Transaction, error) {
	return _WrappingBridge.contract.Transact(opts, "lockERC20", _token, _amount, _destAddress)
}

// LockERC20 is a paid mutator transaction binding the contract method 0xabb18184.
//
// Solidity: function lockERC20(address _token, uint256 _amount, string _destAddress) returns(bool success)
func (_WrappingBridge *WrappingBridgeSession) LockERC20(_token common.Address, _amount *big.Int, _destAddress string) (*types.Transaction, error) {
	return _WrappingBridge.Contract.LockERC20(&_WrappingBridge.TransactOpts, _token, _amount, _destAddress)
}

// LockERC20 is a paid mutator transaction binding the contract method 0xabb18184.
//
// Solidity: function lockERC20(address _token, uint256 _amount, string _destAddress) returns(bool success)
func (_WrappingBridge *WrappingBridgeTransactorSession) LockERC20(_token common.Address, _amount *big.Int, _destAddress string) (*types.Transaction, error) {
	return _WrappingBridge.Contract.LockERC20(&_WrappingBridge.TransactOpts, _token, _amount, _destAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WrappingBridge *WrappingBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappingBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WrappingBridge *WrappingBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _WrappingBridge.Contract.RenounceOwnership(&_WrappingBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WrappingBridge *WrappingBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WrappingBridge.Contract.RenounceOwnership(&_WrappingBridge.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WrappingBridge *WrappingBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WrappingBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WrappingBridge *WrappingBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WrappingBridge.Contract.TransferOwnership(&_WrappingBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WrappingBridge *WrappingBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WrappingBridge.Contract.TransferOwnership(&_WrappingBridge.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address _to, uint256 _amount) returns(bool success)
func (_WrappingBridge *WrappingBridgeTransactor) Unlock(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WrappingBridge.contract.Transact(opts, "unlock", _to, _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address _to, uint256 _amount) returns(bool success)
func (_WrappingBridge *WrappingBridgeSession) Unlock(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WrappingBridge.Contract.Unlock(&_WrappingBridge.TransactOpts, _to, _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address _to, uint256 _amount) returns(bool success)
func (_WrappingBridge *WrappingBridgeTransactorSession) Unlock(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WrappingBridge.Contract.Unlock(&_WrappingBridge.TransactOpts, _to, _amount)
}

// UnlockERC20 is a paid mutator transaction binding the contract method 0x3c298e78.
//
// Solidity: function unlockERC20(address _token, address _to, uint256 _amount) returns(bool success)
func (_WrappingBridge *WrappingBridgeTransactor) UnlockERC20(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WrappingBridge.contract.Transact(opts, "unlockERC20", _token, _to, _amount)
}

// UnlockERC20 is a paid mutator transaction binding the contract method 0x3c298e78.
//
// Solidity: function unlockERC20(address _token, address _to, uint256 _amount) returns(bool success)
func (_WrappingBridge *WrappingBridgeSession) UnlockERC20(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WrappingBridge.Contract.UnlockERC20(&_WrappingBridge.TransactOpts, _token, _to, _amount)
}

// UnlockERC20 is a paid mutator transaction binding the contract method 0x3c298e78.
//
// Solidity: function unlockERC20(address _token, address _to, uint256 _amount) returns(bool success)
func (_WrappingBridge *WrappingBridgeTransactorSession) UnlockERC20(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WrappingBridge.Contract.UnlockERC20(&_WrappingBridge.TransactOpts, _token, _to, _amount)
}

// WrappingBridgeLockIterator is returned from FilterLock and is used to iterate over the raw logs and unpacked data for Lock events raised by the WrappingBridge contract.
type WrappingBridgeLockIterator struct {
	Event *WrappingBridgeLock // Event containing the contract specifics and raw log

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
func (it *WrappingBridgeLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappingBridgeLock)
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
		it.Event = new(WrappingBridgeLock)
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
func (it *WrappingBridgeLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappingBridgeLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappingBridgeLock represents a Lock event raised by the WrappingBridge contract.
type WrappingBridgeLock struct {
	From        common.Address
	Value       *big.Int
	DestAddress string
	DestChain   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLock is a free log retrieval operation binding the contract event 0x5e3839fd07d7a08dd22a9de0306651d18bf7dc0bd1cfce5507a989f08773c231.
//
// Solidity: event Lock(address indexed from, uint256 value, string destAddress, int256 destChain)
func (_WrappingBridge *WrappingBridgeFilterer) FilterLock(opts *bind.FilterOpts, from []common.Address) (*WrappingBridgeLockIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _WrappingBridge.contract.FilterLogs(opts, "Lock", fromRule)
	if err != nil {
		return nil, err
	}
	return &WrappingBridgeLockIterator{contract: _WrappingBridge.contract, event: "Lock", logs: logs, sub: sub}, nil
}

// WatchLock is a free log subscription operation binding the contract event 0x5e3839fd07d7a08dd22a9de0306651d18bf7dc0bd1cfce5507a989f08773c231.
//
// Solidity: event Lock(address indexed from, uint256 value, string destAddress, int256 destChain)
func (_WrappingBridge *WrappingBridgeFilterer) WatchLock(opts *bind.WatchOpts, sink chan<- *WrappingBridgeLock, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _WrappingBridge.contract.WatchLogs(opts, "Lock", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappingBridgeLock)
				if err := _WrappingBridge.contract.UnpackLog(event, "Lock", log); err != nil {
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

// ParseLock is a log parse operation binding the contract event 0x5e3839fd07d7a08dd22a9de0306651d18bf7dc0bd1cfce5507a989f08773c231.
//
// Solidity: event Lock(address indexed from, uint256 value, string destAddress, int256 destChain)
func (_WrappingBridge *WrappingBridgeFilterer) ParseLock(log types.Log) (*WrappingBridgeLock, error) {
	event := new(WrappingBridgeLock)
	if err := _WrappingBridge.contract.UnpackLog(event, "Lock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappingBridgeLockERC20Iterator is returned from FilterLockERC20 and is used to iterate over the raw logs and unpacked data for LockERC20 events raised by the WrappingBridge contract.
type WrappingBridgeLockERC20Iterator struct {
	Event *WrappingBridgeLockERC20 // Event containing the contract specifics and raw log

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
func (it *WrappingBridgeLockERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappingBridgeLockERC20)
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
		it.Event = new(WrappingBridgeLockERC20)
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
func (it *WrappingBridgeLockERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappingBridgeLockERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappingBridgeLockERC20 represents a LockERC20 event raised by the WrappingBridge contract.
type WrappingBridgeLockERC20 struct {
	Token       common.Address
	From        common.Address
	Value       *big.Int
	DestAddress string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLockERC20 is a free log retrieval operation binding the contract event 0x0e070005038f90e684f34b1affc54f54170ec8eda8baad9fa7ebcfae119181d4.
//
// Solidity: event LockERC20(address indexed token, address indexed from, uint256 value, string destAddress)
func (_WrappingBridge *WrappingBridgeFilterer) FilterLockERC20(opts *bind.FilterOpts, token []common.Address, from []common.Address) (*WrappingBridgeLockERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _WrappingBridge.contract.FilterLogs(opts, "LockERC20", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &WrappingBridgeLockERC20Iterator{contract: _WrappingBridge.contract, event: "LockERC20", logs: logs, sub: sub}, nil
}

// WatchLockERC20 is a free log subscription operation binding the contract event 0x0e070005038f90e684f34b1affc54f54170ec8eda8baad9fa7ebcfae119181d4.
//
// Solidity: event LockERC20(address indexed token, address indexed from, uint256 value, string destAddress)
func (_WrappingBridge *WrappingBridgeFilterer) WatchLockERC20(opts *bind.WatchOpts, sink chan<- *WrappingBridgeLockERC20, token []common.Address, from []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _WrappingBridge.contract.WatchLogs(opts, "LockERC20", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappingBridgeLockERC20)
				if err := _WrappingBridge.contract.UnpackLog(event, "LockERC20", log); err != nil {
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

// ParseLockERC20 is a log parse operation binding the contract event 0x0e070005038f90e684f34b1affc54f54170ec8eda8baad9fa7ebcfae119181d4.
//
// Solidity: event LockERC20(address indexed token, address indexed from, uint256 value, string destAddress)
func (_WrappingBridge *WrappingBridgeFilterer) ParseLockERC20(log types.Log) (*WrappingBridgeLockERC20, error) {
	event := new(WrappingBridgeLockERC20)
	if err := _WrappingBridge.contract.UnpackLog(event, "LockERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappingBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WrappingBridge contract.
type WrappingBridgeOwnershipTransferredIterator struct {
	Event *WrappingBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WrappingBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappingBridgeOwnershipTransferred)
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
		it.Event = new(WrappingBridgeOwnershipTransferred)
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
func (it *WrappingBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappingBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappingBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the WrappingBridge contract.
type WrappingBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WrappingBridge *WrappingBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WrappingBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WrappingBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WrappingBridgeOwnershipTransferredIterator{contract: _WrappingBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WrappingBridge *WrappingBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WrappingBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WrappingBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappingBridgeOwnershipTransferred)
				if err := _WrappingBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WrappingBridge *WrappingBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*WrappingBridgeOwnershipTransferred, error) {
	event := new(WrappingBridgeOwnershipTransferred)
	if err := _WrappingBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappingBridgeUnlockIterator is returned from FilterUnlock and is used to iterate over the raw logs and unpacked data for CoinsUnlock events raised by the WrappingBridge contract.
type WrappingBridgeUnlockIterator struct {
	Event *WrappingBridgeUnlock // Event containing the contract specifics and raw log

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
func (it *WrappingBridgeUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappingBridgeUnlock)
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
		it.Event = new(WrappingBridgeUnlock)
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
func (it *WrappingBridgeUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappingBridgeUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappingBridgeUnlock represents a CoinsUnlock event raised by the WrappingBridge contract.
type WrappingBridgeUnlock struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnlock is a free log retrieval operation binding the contract event 0x6381d9813cabeb57471b5a7e05078e64845ccdb563146a6911d536f24ce960f1.
//
// Solidity: event CoinsUnlock(address indexed to, uint256 value)
func (_WrappingBridge *WrappingBridgeFilterer) FilterUnlock(opts *bind.FilterOpts, to []common.Address) (*WrappingBridgeUnlockIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappingBridge.contract.FilterLogs(opts, "CoinsUnlock", toRule)
	if err != nil {
		return nil, err
	}
	return &WrappingBridgeUnlockIterator{contract: _WrappingBridge.contract, event: "CoinsUnlock", logs: logs, sub: sub}, nil
}

// WatchUnlock is a free log subscription operation binding the contract event 0x6381d9813cabeb57471b5a7e05078e64845ccdb563146a6911d536f24ce960f1.
//
// Solidity: event CoinsUnlock(address indexed to, uint256 value)
func (_WrappingBridge *WrappingBridgeFilterer) WatchUnlock(opts *bind.WatchOpts, sink chan<- *WrappingBridgeUnlock, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappingBridge.contract.WatchLogs(opts, "CoinsUnlock", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappingBridgeUnlock)
				if err := _WrappingBridge.contract.UnpackLog(event, "CoinsUnlock", log); err != nil {
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

// ParseUnlock is a log parse operation binding the contract event 0x6381d9813cabeb57471b5a7e05078e64845ccdb563146a6911d536f24ce960f1.
//
// Solidity: event CoinsUnlock(address indexed to, uint256 value)
func (_WrappingBridge *WrappingBridgeFilterer) ParseUnlock(log types.Log) (*WrappingBridgeUnlock, error) {
	event := new(WrappingBridgeUnlock)
	if err := _WrappingBridge.contract.UnpackLog(event, "CoinsUnlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappingBridgeUnlockERC20Iterator is returned from FilterUnlockERC20 and is used to iterate over the raw logs and unpacked data for UnlockERC20 events raised by the WrappingBridge contract.
type WrappingBridgeUnlockERC20Iterator struct {
	Event *WrappingBridgeUnlockERC20 // Event containing the contract specifics and raw log

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
func (it *WrappingBridgeUnlockERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappingBridgeUnlockERC20)
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
		it.Event = new(WrappingBridgeUnlockERC20)
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
func (it *WrappingBridgeUnlockERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappingBridgeUnlockERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappingBridgeUnlockERC20 represents a UnlockERC20 event raised by the WrappingBridge contract.
type WrappingBridgeUnlockERC20 struct {
	Token common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnlockERC20 is a free log retrieval operation binding the contract event 0xdf5187ef1bf8cac4aeb8ad0bd18066f1e6847312f80ac80ad4916bd3758a7bf2.
//
// Solidity: event UnlockERC20(address indexed token, address indexed to, uint256 value)
func (_WrappingBridge *WrappingBridgeFilterer) FilterUnlockERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*WrappingBridgeUnlockERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappingBridge.contract.FilterLogs(opts, "UnlockERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WrappingBridgeUnlockERC20Iterator{contract: _WrappingBridge.contract, event: "UnlockERC20", logs: logs, sub: sub}, nil
}

// WatchUnlockERC20 is a free log subscription operation binding the contract event 0xdf5187ef1bf8cac4aeb8ad0bd18066f1e6847312f80ac80ad4916bd3758a7bf2.
//
// Solidity: event UnlockERC20(address indexed token, address indexed to, uint256 value)
func (_WrappingBridge *WrappingBridgeFilterer) WatchUnlockERC20(opts *bind.WatchOpts, sink chan<- *WrappingBridgeUnlockERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappingBridge.contract.WatchLogs(opts, "UnlockERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappingBridgeUnlockERC20)
				if err := _WrappingBridge.contract.UnpackLog(event, "UnlockERC20", log); err != nil {
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

// ParseUnlockERC20 is a log parse operation binding the contract event 0xdf5187ef1bf8cac4aeb8ad0bd18066f1e6847312f80ac80ad4916bd3758a7bf2.
//
// Solidity: event UnlockERC20(address indexed token, address indexed to, uint256 value)
func (_WrappingBridge *WrappingBridgeFilterer) ParseUnlockERC20(log types.Log) (*WrappingBridgeUnlockERC20, error) {
	event := new(WrappingBridgeUnlockERC20)
	if err := _WrappingBridge.contract.UnpackLog(event, "UnlockERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
