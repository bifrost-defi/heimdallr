// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package locker

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

// LockerMetaData contains all meta data concerning the Locker contract.
var LockerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destination\",\"type\":\"string\"}],\"name\":\"AVAXLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AVAXUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destination\",\"type\":\"string\"}],\"name\":\"USDCLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"USDCUnlocked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usdcContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_destination\",\"type\":\"string\"}],\"name\":\"lockAVAX\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_destination\",\"type\":\"string\"}],\"name\":\"lockUSDC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unlockAVAX\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unlockUSDC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LockerABI is the input ABI used to generate the binding from.
// Deprecated: Use LockerMetaData.ABI instead.
var LockerABI = LockerMetaData.ABI

// Locker is an auto generated Go binding around an Ethereum contract.
type Locker struct {
	LockerCaller     // Read-only binding to the contract
	LockerTransactor // Write-only binding to the contract
	LockerFilterer   // Log filterer for contract events
}

// LockerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockerSession struct {
	Contract     *Locker           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockerCallerSession struct {
	Contract *LockerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LockerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockerTransactorSession struct {
	Contract     *LockerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockerRaw struct {
	Contract *Locker // Generic contract binding to access the raw methods on
}

// LockerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockerCallerRaw struct {
	Contract *LockerCaller // Generic read-only contract binding to access the raw methods on
}

// LockerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockerTransactorRaw struct {
	Contract *LockerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLocker creates a new instance of Locker, bound to a specific deployed contract.
func NewLocker(address common.Address, backend bind.ContractBackend) (*Locker, error) {
	contract, err := bindLocker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Locker{LockerCaller: LockerCaller{contract: contract}, LockerTransactor: LockerTransactor{contract: contract}, LockerFilterer: LockerFilterer{contract: contract}}, nil
}

// NewLockerCaller creates a new read-only instance of Locker, bound to a specific deployed contract.
func NewLockerCaller(address common.Address, caller bind.ContractCaller) (*LockerCaller, error) {
	contract, err := bindLocker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockerCaller{contract: contract}, nil
}

// NewLockerTransactor creates a new write-only instance of Locker, bound to a specific deployed contract.
func NewLockerTransactor(address common.Address, transactor bind.ContractTransactor) (*LockerTransactor, error) {
	contract, err := bindLocker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockerTransactor{contract: contract}, nil
}

// NewLockerFilterer creates a new log filterer instance of Locker, bound to a specific deployed contract.
func NewLockerFilterer(address common.Address, filterer bind.ContractFilterer) (*LockerFilterer, error) {
	contract, err := bindLocker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockerFilterer{contract: contract}, nil
}

// bindLocker binds a generic wrapper to an already deployed contract.
func bindLocker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Locker *LockerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Locker.Contract.LockerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Locker *LockerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Locker.Contract.LockerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Locker *LockerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Locker.Contract.LockerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Locker *LockerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Locker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Locker *LockerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Locker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Locker *LockerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Locker.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Locker *LockerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Locker.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Locker *LockerSession) Owner() (common.Address, error) {
	return _Locker.Contract.Owner(&_Locker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Locker *LockerCallerSession) Owner() (common.Address, error) {
	return _Locker.Contract.Owner(&_Locker.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address usdcContract) returns()
func (_Locker *LockerTransactor) Initialize(opts *bind.TransactOpts, usdcContract common.Address) (*types.Transaction, error) {
	return _Locker.contract.Transact(opts, "initialize", usdcContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address usdcContract) returns()
func (_Locker *LockerSession) Initialize(usdcContract common.Address) (*types.Transaction, error) {
	return _Locker.Contract.Initialize(&_Locker.TransactOpts, usdcContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address usdcContract) returns()
func (_Locker *LockerTransactorSession) Initialize(usdcContract common.Address) (*types.Transaction, error) {
	return _Locker.Contract.Initialize(&_Locker.TransactOpts, usdcContract)
}

// LockAVAX is a paid mutator transaction binding the contract method 0x1df4b98d.
//
// Solidity: function lockAVAX(string _destination) payable returns(bool success)
func (_Locker *LockerTransactor) LockAVAX(opts *bind.TransactOpts, _destination string) (*types.Transaction, error) {
	return _Locker.contract.Transact(opts, "lockAVAX", _destination)
}

// LockAVAX is a paid mutator transaction binding the contract method 0x1df4b98d.
//
// Solidity: function lockAVAX(string _destination) payable returns(bool success)
func (_Locker *LockerSession) LockAVAX(_destination string) (*types.Transaction, error) {
	return _Locker.Contract.LockAVAX(&_Locker.TransactOpts, _destination)
}

// LockAVAX is a paid mutator transaction binding the contract method 0x1df4b98d.
//
// Solidity: function lockAVAX(string _destination) payable returns(bool success)
func (_Locker *LockerTransactorSession) LockAVAX(_destination string) (*types.Transaction, error) {
	return _Locker.Contract.LockAVAX(&_Locker.TransactOpts, _destination)
}

// LockUSDC is a paid mutator transaction binding the contract method 0xce747b3b.
//
// Solidity: function lockUSDC(uint256 _amount, string _destination) returns(bool success)
func (_Locker *LockerTransactor) LockUSDC(opts *bind.TransactOpts, _amount *big.Int, _destination string) (*types.Transaction, error) {
	return _Locker.contract.Transact(opts, "lockUSDC", _amount, _destination)
}

// LockUSDC is a paid mutator transaction binding the contract method 0xce747b3b.
//
// Solidity: function lockUSDC(uint256 _amount, string _destination) returns(bool success)
func (_Locker *LockerSession) LockUSDC(_amount *big.Int, _destination string) (*types.Transaction, error) {
	return _Locker.Contract.LockUSDC(&_Locker.TransactOpts, _amount, _destination)
}

// LockUSDC is a paid mutator transaction binding the contract method 0xce747b3b.
//
// Solidity: function lockUSDC(uint256 _amount, string _destination) returns(bool success)
func (_Locker *LockerTransactorSession) LockUSDC(_amount *big.Int, _destination string) (*types.Transaction, error) {
	return _Locker.Contract.LockUSDC(&_Locker.TransactOpts, _amount, _destination)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Locker *LockerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Locker.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Locker *LockerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Locker.Contract.RenounceOwnership(&_Locker.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Locker *LockerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Locker.Contract.RenounceOwnership(&_Locker.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Locker *LockerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Locker.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Locker *LockerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Locker.Contract.TransferOwnership(&_Locker.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Locker *LockerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Locker.Contract.TransferOwnership(&_Locker.TransactOpts, newOwner)
}

// UnlockAVAX is a paid mutator transaction binding the contract method 0x25971dd1.
//
// Solidity: function unlockAVAX(address _user, uint256 _amount) payable returns(bool success)
func (_Locker *LockerTransactor) UnlockAVAX(opts *bind.TransactOpts, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Locker.contract.Transact(opts, "unlockAVAX", _user, _amount)
}

// UnlockAVAX is a paid mutator transaction binding the contract method 0x25971dd1.
//
// Solidity: function unlockAVAX(address _user, uint256 _amount) payable returns(bool success)
func (_Locker *LockerSession) UnlockAVAX(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Locker.Contract.UnlockAVAX(&_Locker.TransactOpts, _user, _amount)
}

// UnlockAVAX is a paid mutator transaction binding the contract method 0x25971dd1.
//
// Solidity: function unlockAVAX(address _user, uint256 _amount) payable returns(bool success)
func (_Locker *LockerTransactorSession) UnlockAVAX(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Locker.Contract.UnlockAVAX(&_Locker.TransactOpts, _user, _amount)
}

// UnlockUSDC is a paid mutator transaction binding the contract method 0xa7b49fcb.
//
// Solidity: function unlockUSDC(address _user, uint256 _amount) returns(bool success)
func (_Locker *LockerTransactor) UnlockUSDC(opts *bind.TransactOpts, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Locker.contract.Transact(opts, "unlockUSDC", _user, _amount)
}

// UnlockUSDC is a paid mutator transaction binding the contract method 0xa7b49fcb.
//
// Solidity: function unlockUSDC(address _user, uint256 _amount) returns(bool success)
func (_Locker *LockerSession) UnlockUSDC(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Locker.Contract.UnlockUSDC(&_Locker.TransactOpts, _user, _amount)
}

// UnlockUSDC is a paid mutator transaction binding the contract method 0xa7b49fcb.
//
// Solidity: function unlockUSDC(address _user, uint256 _amount) returns(bool success)
func (_Locker *LockerTransactorSession) UnlockUSDC(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Locker.Contract.UnlockUSDC(&_Locker.TransactOpts, _user, _amount)
}

// LockerAVAXLockedIterator is returned from FilterAVAXLocked and is used to iterate over the raw logs and unpacked data for AVAXLocked events raised by the Locker contract.
type LockerAVAXLockedIterator struct {
	Event *LockerAVAXLocked // Event containing the contract specifics and raw log

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
func (it *LockerAVAXLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockerAVAXLocked)
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
		it.Event = new(LockerAVAXLocked)
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
func (it *LockerAVAXLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockerAVAXLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockerAVAXLocked represents a AVAXLocked event raised by the Locker contract.
type LockerAVAXLocked struct {
	User        common.Address
	Amount      *big.Int
	Destination string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAVAXLocked is a free log retrieval operation binding the contract event 0x379592cb51967f509c548d68f5f4ab55be6fecaa7d68487804decf8f7fcc374a.
//
// Solidity: event AVAXLocked(address user, uint256 amount, string destination)
func (_Locker *LockerFilterer) FilterAVAXLocked(opts *bind.FilterOpts) (*LockerAVAXLockedIterator, error) {

	logs, sub, err := _Locker.contract.FilterLogs(opts, "AVAXLocked")
	if err != nil {
		return nil, err
	}
	return &LockerAVAXLockedIterator{contract: _Locker.contract, event: "AVAXLocked", logs: logs, sub: sub}, nil
}

// WatchAVAXLocked is a free log subscription operation binding the contract event 0x379592cb51967f509c548d68f5f4ab55be6fecaa7d68487804decf8f7fcc374a.
//
// Solidity: event AVAXLocked(address user, uint256 amount, string destination)
func (_Locker *LockerFilterer) WatchAVAXLocked(opts *bind.WatchOpts, sink chan<- *LockerAVAXLocked) (event.Subscription, error) {

	logs, sub, err := _Locker.contract.WatchLogs(opts, "AVAXLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockerAVAXLocked)
				if err := _Locker.contract.UnpackLog(event, "AVAXLocked", log); err != nil {
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

// ParseAVAXLocked is a log parse operation binding the contract event 0x379592cb51967f509c548d68f5f4ab55be6fecaa7d68487804decf8f7fcc374a.
//
// Solidity: event AVAXLocked(address user, uint256 amount, string destination)
func (_Locker *LockerFilterer) ParseAVAXLocked(log types.Log) (*LockerAVAXLocked, error) {
	event := new(LockerAVAXLocked)
	if err := _Locker.contract.UnpackLog(event, "AVAXLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockerAVAXUnlockedIterator is returned from FilterAVAXUnlocked and is used to iterate over the raw logs and unpacked data for AVAXUnlocked events raised by the Locker contract.
type LockerAVAXUnlockedIterator struct {
	Event *LockerAVAXUnlocked // Event containing the contract specifics and raw log

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
func (it *LockerAVAXUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockerAVAXUnlocked)
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
		it.Event = new(LockerAVAXUnlocked)
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
func (it *LockerAVAXUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockerAVAXUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockerAVAXUnlocked represents a AVAXUnlocked event raised by the Locker contract.
type LockerAVAXUnlocked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAVAXUnlocked is a free log retrieval operation binding the contract event 0xba7ea2ed5478abb1ad27fd86093575771c2814ae60137c492411b0be84caaa76.
//
// Solidity: event AVAXUnlocked(address user, uint256 amount)
func (_Locker *LockerFilterer) FilterAVAXUnlocked(opts *bind.FilterOpts) (*LockerAVAXUnlockedIterator, error) {

	logs, sub, err := _Locker.contract.FilterLogs(opts, "AVAXUnlocked")
	if err != nil {
		return nil, err
	}
	return &LockerAVAXUnlockedIterator{contract: _Locker.contract, event: "AVAXUnlocked", logs: logs, sub: sub}, nil
}

// WatchAVAXUnlocked is a free log subscription operation binding the contract event 0xba7ea2ed5478abb1ad27fd86093575771c2814ae60137c492411b0be84caaa76.
//
// Solidity: event AVAXUnlocked(address user, uint256 amount)
func (_Locker *LockerFilterer) WatchAVAXUnlocked(opts *bind.WatchOpts, sink chan<- *LockerAVAXUnlocked) (event.Subscription, error) {

	logs, sub, err := _Locker.contract.WatchLogs(opts, "AVAXUnlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockerAVAXUnlocked)
				if err := _Locker.contract.UnpackLog(event, "AVAXUnlocked", log); err != nil {
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

// ParseAVAXUnlocked is a log parse operation binding the contract event 0xba7ea2ed5478abb1ad27fd86093575771c2814ae60137c492411b0be84caaa76.
//
// Solidity: event AVAXUnlocked(address user, uint256 amount)
func (_Locker *LockerFilterer) ParseAVAXUnlocked(log types.Log) (*LockerAVAXUnlocked, error) {
	event := new(LockerAVAXUnlocked)
	if err := _Locker.contract.UnpackLog(event, "AVAXUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Locker contract.
type LockerOwnershipTransferredIterator struct {
	Event *LockerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LockerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockerOwnershipTransferred)
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
		it.Event = new(LockerOwnershipTransferred)
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
func (it *LockerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockerOwnershipTransferred represents a OwnershipTransferred event raised by the Locker contract.
type LockerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Locker *LockerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LockerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Locker.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LockerOwnershipTransferredIterator{contract: _Locker.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Locker *LockerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LockerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Locker.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockerOwnershipTransferred)
				if err := _Locker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Locker *LockerFilterer) ParseOwnershipTransferred(log types.Log) (*LockerOwnershipTransferred, error) {
	event := new(LockerOwnershipTransferred)
	if err := _Locker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockerUSDCLockedIterator is returned from FilterUSDCLocked and is used to iterate over the raw logs and unpacked data for USDCLocked events raised by the Locker contract.
type LockerUSDCLockedIterator struct {
	Event *LockerUSDCLocked // Event containing the contract specifics and raw log

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
func (it *LockerUSDCLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockerUSDCLocked)
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
		it.Event = new(LockerUSDCLocked)
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
func (it *LockerUSDCLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockerUSDCLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockerUSDCLocked represents a USDCLocked event raised by the Locker contract.
type LockerUSDCLocked struct {
	User        common.Address
	Amount      *big.Int
	Destination string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUSDCLocked is a free log retrieval operation binding the contract event 0xd01dddd7bfcb80343deac08bd170d7705dcb766dc86bf16ece39c9405f03aa95.
//
// Solidity: event USDCLocked(address user, uint256 amount, string destination)
func (_Locker *LockerFilterer) FilterUSDCLocked(opts *bind.FilterOpts) (*LockerUSDCLockedIterator, error) {

	logs, sub, err := _Locker.contract.FilterLogs(opts, "USDCLocked")
	if err != nil {
		return nil, err
	}
	return &LockerUSDCLockedIterator{contract: _Locker.contract, event: "USDCLocked", logs: logs, sub: sub}, nil
}

// WatchUSDCLocked is a free log subscription operation binding the contract event 0xd01dddd7bfcb80343deac08bd170d7705dcb766dc86bf16ece39c9405f03aa95.
//
// Solidity: event USDCLocked(address user, uint256 amount, string destination)
func (_Locker *LockerFilterer) WatchUSDCLocked(opts *bind.WatchOpts, sink chan<- *LockerUSDCLocked) (event.Subscription, error) {

	logs, sub, err := _Locker.contract.WatchLogs(opts, "USDCLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockerUSDCLocked)
				if err := _Locker.contract.UnpackLog(event, "USDCLocked", log); err != nil {
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

// ParseUSDCLocked is a log parse operation binding the contract event 0xd01dddd7bfcb80343deac08bd170d7705dcb766dc86bf16ece39c9405f03aa95.
//
// Solidity: event USDCLocked(address user, uint256 amount, string destination)
func (_Locker *LockerFilterer) ParseUSDCLocked(log types.Log) (*LockerUSDCLocked, error) {
	event := new(LockerUSDCLocked)
	if err := _Locker.contract.UnpackLog(event, "USDCLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockerUSDCUnlockedIterator is returned from FilterUSDCUnlocked and is used to iterate over the raw logs and unpacked data for USDCUnlocked events raised by the Locker contract.
type LockerUSDCUnlockedIterator struct {
	Event *LockerUSDCUnlocked // Event containing the contract specifics and raw log

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
func (it *LockerUSDCUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockerUSDCUnlocked)
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
		it.Event = new(LockerUSDCUnlocked)
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
func (it *LockerUSDCUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockerUSDCUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockerUSDCUnlocked represents a USDCUnlocked event raised by the Locker contract.
type LockerUSDCUnlocked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUSDCUnlocked is a free log retrieval operation binding the contract event 0x09f7642ed9ed28d798b13e7c738ffcdee4ef07563135359c48125ab831093f08.
//
// Solidity: event USDCUnlocked(address user, uint256 amount)
func (_Locker *LockerFilterer) FilterUSDCUnlocked(opts *bind.FilterOpts) (*LockerUSDCUnlockedIterator, error) {

	logs, sub, err := _Locker.contract.FilterLogs(opts, "USDCUnlocked")
	if err != nil {
		return nil, err
	}
	return &LockerUSDCUnlockedIterator{contract: _Locker.contract, event: "USDCUnlocked", logs: logs, sub: sub}, nil
}

// WatchUSDCUnlocked is a free log subscription operation binding the contract event 0x09f7642ed9ed28d798b13e7c738ffcdee4ef07563135359c48125ab831093f08.
//
// Solidity: event USDCUnlocked(address user, uint256 amount)
func (_Locker *LockerFilterer) WatchUSDCUnlocked(opts *bind.WatchOpts, sink chan<- *LockerUSDCUnlocked) (event.Subscription, error) {

	logs, sub, err := _Locker.contract.WatchLogs(opts, "USDCUnlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockerUSDCUnlocked)
				if err := _Locker.contract.UnpackLog(event, "USDCUnlocked", log); err != nil {
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

// ParseUSDCUnlocked is a log parse operation binding the contract event 0x09f7642ed9ed28d798b13e7c738ffcdee4ef07563135359c48125ab831093f08.
//
// Solidity: event USDCUnlocked(address user, uint256 amount)
func (_Locker *LockerFilterer) ParseUSDCUnlocked(log types.Log) (*LockerUSDCUnlocked, error) {
	event := new(LockerUSDCUnlocked)
	if err := _Locker.contract.UnpackLog(event, "USDCUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
