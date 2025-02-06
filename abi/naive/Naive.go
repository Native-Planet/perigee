// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package naive

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

// NaiveMetaData contains all meta data concerning the Naive contract.
var NaiveMetaData = &bind.MetaData{
	ABI: "[{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Batch\",\"type\":\"event\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50606a80601d6000396000f3006080604052348015600f57600080fd5b507fcca739c72762deed05941b38d4aa82f2718c74457d5e2d8c5b1d7642caf2219660405160405180910390a10000a165627a7a723058203a3a9f89e6b4bb6cb4162a6670aae73d8cc45324ef9d50670bff9728313ddc7b0029",
}

// NaiveABI is the input ABI used to generate the binding from.
// Deprecated: Use NaiveMetaData.ABI instead.
var NaiveABI = NaiveMetaData.ABI

// NaiveBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NaiveMetaData.Bin instead.
var NaiveBin = NaiveMetaData.Bin

// DeployNaive deploys a new Ethereum contract, binding an instance of Naive to it.
func DeployNaive(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Naive, error) {
	parsed, err := NaiveMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NaiveBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Naive{NaiveCaller: NaiveCaller{contract: contract}, NaiveTransactor: NaiveTransactor{contract: contract}, NaiveFilterer: NaiveFilterer{contract: contract}}, nil
}

// Naive is an auto generated Go binding around an Ethereum contract.
type Naive struct {
	NaiveCaller     // Read-only binding to the contract
	NaiveTransactor // Write-only binding to the contract
	NaiveFilterer   // Log filterer for contract events
}

// NaiveCaller is an auto generated read-only Go binding around an Ethereum contract.
type NaiveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NaiveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NaiveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NaiveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NaiveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NaiveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NaiveSession struct {
	Contract     *Naive            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NaiveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NaiveCallerSession struct {
	Contract *NaiveCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NaiveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NaiveTransactorSession struct {
	Contract     *NaiveTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NaiveRaw is an auto generated low-level Go binding around an Ethereum contract.
type NaiveRaw struct {
	Contract *Naive // Generic contract binding to access the raw methods on
}

// NaiveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NaiveCallerRaw struct {
	Contract *NaiveCaller // Generic read-only contract binding to access the raw methods on
}

// NaiveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NaiveTransactorRaw struct {
	Contract *NaiveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNaive creates a new instance of Naive, bound to a specific deployed contract.
func NewNaive(address common.Address, backend bind.ContractBackend) (*Naive, error) {
	contract, err := bindNaive(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Naive{NaiveCaller: NaiveCaller{contract: contract}, NaiveTransactor: NaiveTransactor{contract: contract}, NaiveFilterer: NaiveFilterer{contract: contract}}, nil
}

// NewNaiveCaller creates a new read-only instance of Naive, bound to a specific deployed contract.
func NewNaiveCaller(address common.Address, caller bind.ContractCaller) (*NaiveCaller, error) {
	contract, err := bindNaive(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NaiveCaller{contract: contract}, nil
}

// NewNaiveTransactor creates a new write-only instance of Naive, bound to a specific deployed contract.
func NewNaiveTransactor(address common.Address, transactor bind.ContractTransactor) (*NaiveTransactor, error) {
	contract, err := bindNaive(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NaiveTransactor{contract: contract}, nil
}

// NewNaiveFilterer creates a new log filterer instance of Naive, bound to a specific deployed contract.
func NewNaiveFilterer(address common.Address, filterer bind.ContractFilterer) (*NaiveFilterer, error) {
	contract, err := bindNaive(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NaiveFilterer{contract: contract}, nil
}

// bindNaive binds a generic wrapper to an already deployed contract.
func bindNaive(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NaiveMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Naive *NaiveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Naive.Contract.NaiveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Naive *NaiveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Naive.Contract.NaiveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Naive *NaiveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Naive.Contract.NaiveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Naive *NaiveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Naive.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Naive *NaiveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Naive.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Naive *NaiveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Naive.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Naive *NaiveTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Naive.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Naive *NaiveSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Naive.Contract.Fallback(&_Naive.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Naive *NaiveTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Naive.Contract.Fallback(&_Naive.TransactOpts, calldata)
}

// NaiveBatchIterator is returned from FilterBatch and is used to iterate over the raw logs and unpacked data for Batch events raised by the Naive contract.
type NaiveBatchIterator struct {
	Event *NaiveBatch // Event containing the contract specifics and raw log

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
func (it *NaiveBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NaiveBatch)
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
		it.Event = new(NaiveBatch)
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
func (it *NaiveBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NaiveBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NaiveBatch represents a Batch event raised by the Naive contract.
type NaiveBatch struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBatch is a free log retrieval operation binding the contract event 0xcca739c72762deed05941b38d4aa82f2718c74457d5e2d8c5b1d7642caf22196.
//
// Solidity: event Batch()
func (_Naive *NaiveFilterer) FilterBatch(opts *bind.FilterOpts) (*NaiveBatchIterator, error) {

	logs, sub, err := _Naive.contract.FilterLogs(opts, "Batch")
	if err != nil {
		return nil, err
	}
	return &NaiveBatchIterator{contract: _Naive.contract, event: "Batch", logs: logs, sub: sub}, nil
}

// WatchBatch is a free log subscription operation binding the contract event 0xcca739c72762deed05941b38d4aa82f2718c74457d5e2d8c5b1d7642caf22196.
//
// Solidity: event Batch()
func (_Naive *NaiveFilterer) WatchBatch(opts *bind.WatchOpts, sink chan<- *NaiveBatch) (event.Subscription, error) {

	logs, sub, err := _Naive.contract.WatchLogs(opts, "Batch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NaiveBatch)
				if err := _Naive.contract.UnpackLog(event, "Batch", log); err != nil {
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

// ParseBatch is a log parse operation binding the contract event 0xcca739c72762deed05941b38d4aa82f2718c74457d5e2d8c5b1d7642caf22196.
//
// Solidity: event Batch()
func (_Naive *NaiveFilterer) ParseBatch(log types.Log) (*NaiveBatch, error) {
	event := new(NaiveBatch)
	if err := _Naive.contract.UnpackLog(event, "Batch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
