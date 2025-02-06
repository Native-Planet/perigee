// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package addressutils

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

// AddressutilsMetaData contains all meta data concerning the Addressutils contract.
var AddressutilsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820835d8bada385b022daca962f6c5e5737adeecf2ba6666df4be63f5e53f151d8c0029",
}

// AddressutilsABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressutilsMetaData.ABI instead.
var AddressutilsABI = AddressutilsMetaData.ABI

// AddressutilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressutilsMetaData.Bin instead.
var AddressutilsBin = AddressutilsMetaData.Bin

// DeployAddressutils deploys a new Ethereum contract, binding an instance of Addressutils to it.
func DeployAddressutils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Addressutils, error) {
	parsed, err := AddressutilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressutilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Addressutils{AddressutilsCaller: AddressutilsCaller{contract: contract}, AddressutilsTransactor: AddressutilsTransactor{contract: contract}, AddressutilsFilterer: AddressutilsFilterer{contract: contract}}, nil
}

// Addressutils is an auto generated Go binding around an Ethereum contract.
type Addressutils struct {
	AddressutilsCaller     // Read-only binding to the contract
	AddressutilsTransactor // Write-only binding to the contract
	AddressutilsFilterer   // Log filterer for contract events
}

// AddressutilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressutilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressutilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressutilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressutilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressutilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressutilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressutilsSession struct {
	Contract     *Addressutils     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressutilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressutilsCallerSession struct {
	Contract *AddressutilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AddressutilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressutilsTransactorSession struct {
	Contract     *AddressutilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AddressutilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressutilsRaw struct {
	Contract *Addressutils // Generic contract binding to access the raw methods on
}

// AddressutilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressutilsCallerRaw struct {
	Contract *AddressutilsCaller // Generic read-only contract binding to access the raw methods on
}

// AddressutilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressutilsTransactorRaw struct {
	Contract *AddressutilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressutils creates a new instance of Addressutils, bound to a specific deployed contract.
func NewAddressutils(address common.Address, backend bind.ContractBackend) (*Addressutils, error) {
	contract, err := bindAddressutils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Addressutils{AddressutilsCaller: AddressutilsCaller{contract: contract}, AddressutilsTransactor: AddressutilsTransactor{contract: contract}, AddressutilsFilterer: AddressutilsFilterer{contract: contract}}, nil
}

// NewAddressutilsCaller creates a new read-only instance of Addressutils, bound to a specific deployed contract.
func NewAddressutilsCaller(address common.Address, caller bind.ContractCaller) (*AddressutilsCaller, error) {
	contract, err := bindAddressutils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressutilsCaller{contract: contract}, nil
}

// NewAddressutilsTransactor creates a new write-only instance of Addressutils, bound to a specific deployed contract.
func NewAddressutilsTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressutilsTransactor, error) {
	contract, err := bindAddressutils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressutilsTransactor{contract: contract}, nil
}

// NewAddressutilsFilterer creates a new log filterer instance of Addressutils, bound to a specific deployed contract.
func NewAddressutilsFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressutilsFilterer, error) {
	contract, err := bindAddressutils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressutilsFilterer{contract: contract}, nil
}

// bindAddressutils binds a generic wrapper to an already deployed contract.
func bindAddressutils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddressutilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Addressutils *AddressutilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Addressutils.Contract.AddressutilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Addressutils *AddressutilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Addressutils.Contract.AddressutilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Addressutils *AddressutilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Addressutils.Contract.AddressutilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Addressutils *AddressutilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Addressutils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Addressutils *AddressutilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Addressutils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Addressutils *AddressutilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Addressutils.Contract.contract.Transact(opts, method, params...)
}
