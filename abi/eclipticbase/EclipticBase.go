// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eclipticbase

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
	_ = abi.ConvertType
)

// EclipticbaseMetaData contains all meta data concerning the Eclipticbase contract.
var EclipticbaseMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"onUpgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"azimuth\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"polls\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousEcliptic\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_previous\",\"type\":\"address\"},{\"name\":\"_azimuth\",\"type\":\"address\"},{\"name\":\"_polls\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
}

// EclipticbaseABI is the input ABI used to generate the binding from.
// Deprecated: Use EclipticbaseMetaData.ABI instead.
var EclipticbaseABI = EclipticbaseMetaData.ABI

// Eclipticbase is an auto generated Go binding around an Ethereum contract.
type Eclipticbase struct {
	EclipticbaseCaller     // Read-only binding to the contract
	EclipticbaseTransactor // Write-only binding to the contract
	EclipticbaseFilterer   // Log filterer for contract events
}

// EclipticbaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type EclipticbaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EclipticbaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EclipticbaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EclipticbaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EclipticbaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EclipticbaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EclipticbaseSession struct {
	Contract     *Eclipticbase     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EclipticbaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EclipticbaseCallerSession struct {
	Contract *EclipticbaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EclipticbaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EclipticbaseTransactorSession struct {
	Contract     *EclipticbaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EclipticbaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type EclipticbaseRaw struct {
	Contract *Eclipticbase // Generic contract binding to access the raw methods on
}

// EclipticbaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EclipticbaseCallerRaw struct {
	Contract *EclipticbaseCaller // Generic read-only contract binding to access the raw methods on
}

// EclipticbaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EclipticbaseTransactorRaw struct {
	Contract *EclipticbaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEclipticbase creates a new instance of Eclipticbase, bound to a specific deployed contract.
func NewEclipticbase(address common.Address, backend bind.ContractBackend) (*Eclipticbase, error) {
	contract, err := bindEclipticbase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eclipticbase{EclipticbaseCaller: EclipticbaseCaller{contract: contract}, EclipticbaseTransactor: EclipticbaseTransactor{contract: contract}, EclipticbaseFilterer: EclipticbaseFilterer{contract: contract}}, nil
}

// NewEclipticbaseCaller creates a new read-only instance of Eclipticbase, bound to a specific deployed contract.
func NewEclipticbaseCaller(address common.Address, caller bind.ContractCaller) (*EclipticbaseCaller, error) {
	contract, err := bindEclipticbase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EclipticbaseCaller{contract: contract}, nil
}

// NewEclipticbaseTransactor creates a new write-only instance of Eclipticbase, bound to a specific deployed contract.
func NewEclipticbaseTransactor(address common.Address, transactor bind.ContractTransactor) (*EclipticbaseTransactor, error) {
	contract, err := bindEclipticbase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EclipticbaseTransactor{contract: contract}, nil
}

// NewEclipticbaseFilterer creates a new log filterer instance of Eclipticbase, bound to a specific deployed contract.
func NewEclipticbaseFilterer(address common.Address, filterer bind.ContractFilterer) (*EclipticbaseFilterer, error) {
	contract, err := bindEclipticbase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EclipticbaseFilterer{contract: contract}, nil
}

// bindEclipticbase binds a generic wrapper to an already deployed contract.
func bindEclipticbase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EclipticbaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eclipticbase *EclipticbaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eclipticbase.Contract.EclipticbaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eclipticbase *EclipticbaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eclipticbase.Contract.EclipticbaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eclipticbase *EclipticbaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eclipticbase.Contract.EclipticbaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eclipticbase *EclipticbaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eclipticbase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eclipticbase *EclipticbaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eclipticbase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eclipticbase *EclipticbaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eclipticbase.Contract.contract.Transact(opts, method, params...)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Eclipticbase *EclipticbaseCaller) Azimuth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Eclipticbase.contract.Call(opts, &out, "azimuth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Eclipticbase *EclipticbaseSession) Azimuth() (common.Address, error) {
	return _Eclipticbase.Contract.Azimuth(&_Eclipticbase.CallOpts)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Eclipticbase *EclipticbaseCallerSession) Azimuth() (common.Address, error) {
	return _Eclipticbase.Contract.Azimuth(&_Eclipticbase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Eclipticbase *EclipticbaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Eclipticbase.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Eclipticbase *EclipticbaseSession) Owner() (common.Address, error) {
	return _Eclipticbase.Contract.Owner(&_Eclipticbase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Eclipticbase *EclipticbaseCallerSession) Owner() (common.Address, error) {
	return _Eclipticbase.Contract.Owner(&_Eclipticbase.CallOpts)
}

// Polls is a free data retrieval call binding the contract method 0xe64853c4.
//
// Solidity: function polls() view returns(address)
func (_Eclipticbase *EclipticbaseCaller) Polls(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Eclipticbase.contract.Call(opts, &out, "polls")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Polls is a free data retrieval call binding the contract method 0xe64853c4.
//
// Solidity: function polls() view returns(address)
func (_Eclipticbase *EclipticbaseSession) Polls() (common.Address, error) {
	return _Eclipticbase.Contract.Polls(&_Eclipticbase.CallOpts)
}

// Polls is a free data retrieval call binding the contract method 0xe64853c4.
//
// Solidity: function polls() view returns(address)
func (_Eclipticbase *EclipticbaseCallerSession) Polls() (common.Address, error) {
	return _Eclipticbase.Contract.Polls(&_Eclipticbase.CallOpts)
}

// PreviousEcliptic is a free data retrieval call binding the contract method 0xed969f68.
//
// Solidity: function previousEcliptic() view returns(address)
func (_Eclipticbase *EclipticbaseCaller) PreviousEcliptic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Eclipticbase.contract.Call(opts, &out, "previousEcliptic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreviousEcliptic is a free data retrieval call binding the contract method 0xed969f68.
//
// Solidity: function previousEcliptic() view returns(address)
func (_Eclipticbase *EclipticbaseSession) PreviousEcliptic() (common.Address, error) {
	return _Eclipticbase.Contract.PreviousEcliptic(&_Eclipticbase.CallOpts)
}

// PreviousEcliptic is a free data retrieval call binding the contract method 0xed969f68.
//
// Solidity: function previousEcliptic() view returns(address)
func (_Eclipticbase *EclipticbaseCallerSession) PreviousEcliptic() (common.Address, error) {
	return _Eclipticbase.Contract.PreviousEcliptic(&_Eclipticbase.CallOpts)
}

// OnUpgrade is a paid mutator transaction binding the contract method 0x1824a46b.
//
// Solidity: function onUpgrade() returns()
func (_Eclipticbase *EclipticbaseTransactor) OnUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eclipticbase.contract.Transact(opts, "onUpgrade")
}

// OnUpgrade is a paid mutator transaction binding the contract method 0x1824a46b.
//
// Solidity: function onUpgrade() returns()
func (_Eclipticbase *EclipticbaseSession) OnUpgrade() (*types.Transaction, error) {
	return _Eclipticbase.Contract.OnUpgrade(&_Eclipticbase.TransactOpts)
}

// OnUpgrade is a paid mutator transaction binding the contract method 0x1824a46b.
//
// Solidity: function onUpgrade() returns()
func (_Eclipticbase *EclipticbaseTransactorSession) OnUpgrade() (*types.Transaction, error) {
	return _Eclipticbase.Contract.OnUpgrade(&_Eclipticbase.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Eclipticbase *EclipticbaseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eclipticbase.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Eclipticbase *EclipticbaseSession) RenounceOwnership() (*types.Transaction, error) {
	return _Eclipticbase.Contract.RenounceOwnership(&_Eclipticbase.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Eclipticbase *EclipticbaseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Eclipticbase.Contract.RenounceOwnership(&_Eclipticbase.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Eclipticbase *EclipticbaseTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Eclipticbase.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Eclipticbase *EclipticbaseSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Eclipticbase.Contract.TransferOwnership(&_Eclipticbase.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Eclipticbase *EclipticbaseTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Eclipticbase.Contract.TransferOwnership(&_Eclipticbase.TransactOpts, _newOwner)
}

// EclipticbaseOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Eclipticbase contract.
type EclipticbaseOwnershipRenouncedIterator struct {
	Event *EclipticbaseOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *EclipticbaseOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EclipticbaseOwnershipRenounced)
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
		it.Event = new(EclipticbaseOwnershipRenounced)
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
func (it *EclipticbaseOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EclipticbaseOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EclipticbaseOwnershipRenounced represents a OwnershipRenounced event raised by the Eclipticbase contract.
type EclipticbaseOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Eclipticbase *EclipticbaseFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*EclipticbaseOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Eclipticbase.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EclipticbaseOwnershipRenouncedIterator{contract: _Eclipticbase.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Eclipticbase *EclipticbaseFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *EclipticbaseOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Eclipticbase.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EclipticbaseOwnershipRenounced)
				if err := _Eclipticbase.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Eclipticbase *EclipticbaseFilterer) ParseOwnershipRenounced(log types.Log) (*EclipticbaseOwnershipRenounced, error) {
	event := new(EclipticbaseOwnershipRenounced)
	if err := _Eclipticbase.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EclipticbaseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Eclipticbase contract.
type EclipticbaseOwnershipTransferredIterator struct {
	Event *EclipticbaseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EclipticbaseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EclipticbaseOwnershipTransferred)
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
		it.Event = new(EclipticbaseOwnershipTransferred)
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
func (it *EclipticbaseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EclipticbaseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EclipticbaseOwnershipTransferred represents a OwnershipTransferred event raised by the Eclipticbase contract.
type EclipticbaseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Eclipticbase *EclipticbaseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EclipticbaseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Eclipticbase.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EclipticbaseOwnershipTransferredIterator{contract: _Eclipticbase.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Eclipticbase *EclipticbaseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EclipticbaseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Eclipticbase.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EclipticbaseOwnershipTransferred)
				if err := _Eclipticbase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Eclipticbase *EclipticbaseFilterer) ParseOwnershipTransferred(log types.Log) (*EclipticbaseOwnershipTransferred, error) {
	event := new(EclipticbaseOwnershipTransferred)
	if err := _Eclipticbase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EclipticbaseUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Eclipticbase contract.
type EclipticbaseUpgradedIterator struct {
	Event *EclipticbaseUpgraded // Event containing the contract specifics and raw log

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
func (it *EclipticbaseUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EclipticbaseUpgraded)
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
		it.Event = new(EclipticbaseUpgraded)
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
func (it *EclipticbaseUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EclipticbaseUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EclipticbaseUpgraded represents a Upgraded event raised by the Eclipticbase contract.
type EclipticbaseUpgraded struct {
	To  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address to)
func (_Eclipticbase *EclipticbaseFilterer) FilterUpgraded(opts *bind.FilterOpts) (*EclipticbaseUpgradedIterator, error) {

	logs, sub, err := _Eclipticbase.contract.FilterLogs(opts, "Upgraded")
	if err != nil {
		return nil, err
	}
	return &EclipticbaseUpgradedIterator{contract: _Eclipticbase.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address to)
func (_Eclipticbase *EclipticbaseFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EclipticbaseUpgraded) (event.Subscription, error) {

	logs, sub, err := _Eclipticbase.contract.WatchLogs(opts, "Upgraded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EclipticbaseUpgraded)
				if err := _Eclipticbase.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address to)
func (_Eclipticbase *EclipticbaseFilterer) ParseUpgraded(log types.Log) (*EclipticbaseUpgraded, error) {
	event := new(EclipticbaseUpgraded)
	if err := _Eclipticbase.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
