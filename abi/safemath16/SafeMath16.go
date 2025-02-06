// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package safemath16

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

// Safemath16MetaData contains all meta data concerning the Safemath16 contract.
var Safemath16MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820aa8f708dd85a6bc129840e6463175dd48755cb31a64c5b67326ca77b9a1408070029",
}

// Safemath16ABI is the input ABI used to generate the binding from.
// Deprecated: Use Safemath16MetaData.ABI instead.
var Safemath16ABI = Safemath16MetaData.ABI

// Safemath16Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Safemath16MetaData.Bin instead.
var Safemath16Bin = Safemath16MetaData.Bin

// DeploySafemath16 deploys a new Ethereum contract, binding an instance of Safemath16 to it.
func DeploySafemath16(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Safemath16, error) {
	parsed, err := Safemath16MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Safemath16Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Safemath16{Safemath16Caller: Safemath16Caller{contract: contract}, Safemath16Transactor: Safemath16Transactor{contract: contract}, Safemath16Filterer: Safemath16Filterer{contract: contract}}, nil
}

// Safemath16 is an auto generated Go binding around an Ethereum contract.
type Safemath16 struct {
	Safemath16Caller     // Read-only binding to the contract
	Safemath16Transactor // Write-only binding to the contract
	Safemath16Filterer   // Log filterer for contract events
}

// Safemath16Caller is an auto generated read-only Go binding around an Ethereum contract.
type Safemath16Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Safemath16Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Safemath16Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Safemath16Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Safemath16Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Safemath16Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Safemath16Session struct {
	Contract     *Safemath16       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Safemath16CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Safemath16CallerSession struct {
	Contract *Safemath16Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Safemath16TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Safemath16TransactorSession struct {
	Contract     *Safemath16Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Safemath16Raw is an auto generated low-level Go binding around an Ethereum contract.
type Safemath16Raw struct {
	Contract *Safemath16 // Generic contract binding to access the raw methods on
}

// Safemath16CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Safemath16CallerRaw struct {
	Contract *Safemath16Caller // Generic read-only contract binding to access the raw methods on
}

// Safemath16TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Safemath16TransactorRaw struct {
	Contract *Safemath16Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafemath16 creates a new instance of Safemath16, bound to a specific deployed contract.
func NewSafemath16(address common.Address, backend bind.ContractBackend) (*Safemath16, error) {
	contract, err := bindSafemath16(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Safemath16{Safemath16Caller: Safemath16Caller{contract: contract}, Safemath16Transactor: Safemath16Transactor{contract: contract}, Safemath16Filterer: Safemath16Filterer{contract: contract}}, nil
}

// NewSafemath16Caller creates a new read-only instance of Safemath16, bound to a specific deployed contract.
func NewSafemath16Caller(address common.Address, caller bind.ContractCaller) (*Safemath16Caller, error) {
	contract, err := bindSafemath16(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Safemath16Caller{contract: contract}, nil
}

// NewSafemath16Transactor creates a new write-only instance of Safemath16, bound to a specific deployed contract.
func NewSafemath16Transactor(address common.Address, transactor bind.ContractTransactor) (*Safemath16Transactor, error) {
	contract, err := bindSafemath16(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Safemath16Transactor{contract: contract}, nil
}

// NewSafemath16Filterer creates a new log filterer instance of Safemath16, bound to a specific deployed contract.
func NewSafemath16Filterer(address common.Address, filterer bind.ContractFilterer) (*Safemath16Filterer, error) {
	contract, err := bindSafemath16(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Safemath16Filterer{contract: contract}, nil
}

// bindSafemath16 binds a generic wrapper to an already deployed contract.
func bindSafemath16(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Safemath16MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Safemath16 *Safemath16Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Safemath16.Contract.Safemath16Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Safemath16 *Safemath16Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Safemath16.Contract.Safemath16Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Safemath16 *Safemath16Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Safemath16.Contract.Safemath16Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Safemath16 *Safemath16CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Safemath16.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Safemath16 *Safemath16TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Safemath16.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Safemath16 *Safemath16TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Safemath16.Contract.contract.Transact(opts, method, params...)
}
