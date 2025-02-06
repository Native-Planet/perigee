// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package safemath8

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

// Safemath8MetaData contains all meta data concerning the Safemath8 contract.
var Safemath8MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820f082c1309e604d64845fcd12935b85cd945cad58a250ef165b77aee0ba65f5360029",
}

// Safemath8ABI is the input ABI used to generate the binding from.
// Deprecated: Use Safemath8MetaData.ABI instead.
var Safemath8ABI = Safemath8MetaData.ABI

// Safemath8Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Safemath8MetaData.Bin instead.
var Safemath8Bin = Safemath8MetaData.Bin

// DeploySafemath8 deploys a new Ethereum contract, binding an instance of Safemath8 to it.
func DeploySafemath8(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Safemath8, error) {
	parsed, err := Safemath8MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Safemath8Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Safemath8{Safemath8Caller: Safemath8Caller{contract: contract}, Safemath8Transactor: Safemath8Transactor{contract: contract}, Safemath8Filterer: Safemath8Filterer{contract: contract}}, nil
}

// Safemath8 is an auto generated Go binding around an Ethereum contract.
type Safemath8 struct {
	Safemath8Caller     // Read-only binding to the contract
	Safemath8Transactor // Write-only binding to the contract
	Safemath8Filterer   // Log filterer for contract events
}

// Safemath8Caller is an auto generated read-only Go binding around an Ethereum contract.
type Safemath8Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Safemath8Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Safemath8Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Safemath8Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Safemath8Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Safemath8Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Safemath8Session struct {
	Contract     *Safemath8        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Safemath8CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Safemath8CallerSession struct {
	Contract *Safemath8Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Safemath8TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Safemath8TransactorSession struct {
	Contract     *Safemath8Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Safemath8Raw is an auto generated low-level Go binding around an Ethereum contract.
type Safemath8Raw struct {
	Contract *Safemath8 // Generic contract binding to access the raw methods on
}

// Safemath8CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Safemath8CallerRaw struct {
	Contract *Safemath8Caller // Generic read-only contract binding to access the raw methods on
}

// Safemath8TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Safemath8TransactorRaw struct {
	Contract *Safemath8Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafemath8 creates a new instance of Safemath8, bound to a specific deployed contract.
func NewSafemath8(address common.Address, backend bind.ContractBackend) (*Safemath8, error) {
	contract, err := bindSafemath8(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Safemath8{Safemath8Caller: Safemath8Caller{contract: contract}, Safemath8Transactor: Safemath8Transactor{contract: contract}, Safemath8Filterer: Safemath8Filterer{contract: contract}}, nil
}

// NewSafemath8Caller creates a new read-only instance of Safemath8, bound to a specific deployed contract.
func NewSafemath8Caller(address common.Address, caller bind.ContractCaller) (*Safemath8Caller, error) {
	contract, err := bindSafemath8(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Safemath8Caller{contract: contract}, nil
}

// NewSafemath8Transactor creates a new write-only instance of Safemath8, bound to a specific deployed contract.
func NewSafemath8Transactor(address common.Address, transactor bind.ContractTransactor) (*Safemath8Transactor, error) {
	contract, err := bindSafemath8(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Safemath8Transactor{contract: contract}, nil
}

// NewSafemath8Filterer creates a new log filterer instance of Safemath8, bound to a specific deployed contract.
func NewSafemath8Filterer(address common.Address, filterer bind.ContractFilterer) (*Safemath8Filterer, error) {
	contract, err := bindSafemath8(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Safemath8Filterer{contract: contract}, nil
}

// bindSafemath8 binds a generic wrapper to an already deployed contract.
func bindSafemath8(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Safemath8MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Safemath8 *Safemath8Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Safemath8.Contract.Safemath8Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Safemath8 *Safemath8Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Safemath8.Contract.Safemath8Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Safemath8 *Safemath8Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Safemath8.Contract.Safemath8Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Safemath8 *Safemath8CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Safemath8.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Safemath8 *Safemath8TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Safemath8.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Safemath8 *Safemath8TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Safemath8.Contract.contract.Transact(opts, method, params...)
}
