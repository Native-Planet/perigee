// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package itreasuryproxy

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

// ItreasuryproxyMetaData contains all meta data concerning the Itreasuryproxy contract.
var ItreasuryproxyMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_impl\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"freeze\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ItreasuryproxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ItreasuryproxyMetaData.ABI instead.
var ItreasuryproxyABI = ItreasuryproxyMetaData.ABI

// Itreasuryproxy is an auto generated Go binding around an Ethereum contract.
type Itreasuryproxy struct {
	ItreasuryproxyCaller     // Read-only binding to the contract
	ItreasuryproxyTransactor // Write-only binding to the contract
	ItreasuryproxyFilterer   // Log filterer for contract events
}

// ItreasuryproxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ItreasuryproxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItreasuryproxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ItreasuryproxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItreasuryproxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ItreasuryproxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItreasuryproxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ItreasuryproxySession struct {
	Contract     *Itreasuryproxy   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ItreasuryproxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ItreasuryproxyCallerSession struct {
	Contract *ItreasuryproxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ItreasuryproxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ItreasuryproxyTransactorSession struct {
	Contract     *ItreasuryproxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ItreasuryproxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ItreasuryproxyRaw struct {
	Contract *Itreasuryproxy // Generic contract binding to access the raw methods on
}

// ItreasuryproxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ItreasuryproxyCallerRaw struct {
	Contract *ItreasuryproxyCaller // Generic read-only contract binding to access the raw methods on
}

// ItreasuryproxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ItreasuryproxyTransactorRaw struct {
	Contract *ItreasuryproxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewItreasuryproxy creates a new instance of Itreasuryproxy, bound to a specific deployed contract.
func NewItreasuryproxy(address common.Address, backend bind.ContractBackend) (*Itreasuryproxy, error) {
	contract, err := bindItreasuryproxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Itreasuryproxy{ItreasuryproxyCaller: ItreasuryproxyCaller{contract: contract}, ItreasuryproxyTransactor: ItreasuryproxyTransactor{contract: contract}, ItreasuryproxyFilterer: ItreasuryproxyFilterer{contract: contract}}, nil
}

// NewItreasuryproxyCaller creates a new read-only instance of Itreasuryproxy, bound to a specific deployed contract.
func NewItreasuryproxyCaller(address common.Address, caller bind.ContractCaller) (*ItreasuryproxyCaller, error) {
	contract, err := bindItreasuryproxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ItreasuryproxyCaller{contract: contract}, nil
}

// NewItreasuryproxyTransactor creates a new write-only instance of Itreasuryproxy, bound to a specific deployed contract.
func NewItreasuryproxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ItreasuryproxyTransactor, error) {
	contract, err := bindItreasuryproxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ItreasuryproxyTransactor{contract: contract}, nil
}

// NewItreasuryproxyFilterer creates a new log filterer instance of Itreasuryproxy, bound to a specific deployed contract.
func NewItreasuryproxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ItreasuryproxyFilterer, error) {
	contract, err := bindItreasuryproxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ItreasuryproxyFilterer{contract: contract}, nil
}

// bindItreasuryproxy binds a generic wrapper to an already deployed contract.
func bindItreasuryproxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ItreasuryproxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Itreasuryproxy *ItreasuryproxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Itreasuryproxy.Contract.ItreasuryproxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Itreasuryproxy *ItreasuryproxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Itreasuryproxy.Contract.ItreasuryproxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Itreasuryproxy *ItreasuryproxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Itreasuryproxy.Contract.ItreasuryproxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Itreasuryproxy *ItreasuryproxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Itreasuryproxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Itreasuryproxy *ItreasuryproxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Itreasuryproxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Itreasuryproxy *ItreasuryproxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Itreasuryproxy.Contract.contract.Transact(opts, method, params...)
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns(bool)
func (_Itreasuryproxy *ItreasuryproxyTransactor) Freeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Itreasuryproxy.contract.Transact(opts, "freeze")
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns(bool)
func (_Itreasuryproxy *ItreasuryproxySession) Freeze() (*types.Transaction, error) {
	return _Itreasuryproxy.Contract.Freeze(&_Itreasuryproxy.TransactOpts)
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns(bool)
func (_Itreasuryproxy *ItreasuryproxyTransactorSession) Freeze() (*types.Transaction, error) {
	return _Itreasuryproxy.Contract.Freeze(&_Itreasuryproxy.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address _impl) returns(bool)
func (_Itreasuryproxy *ItreasuryproxyTransactor) UpgradeTo(opts *bind.TransactOpts, _impl common.Address) (*types.Transaction, error) {
	return _Itreasuryproxy.contract.Transact(opts, "upgradeTo", _impl)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address _impl) returns(bool)
func (_Itreasuryproxy *ItreasuryproxySession) UpgradeTo(_impl common.Address) (*types.Transaction, error) {
	return _Itreasuryproxy.Contract.UpgradeTo(&_Itreasuryproxy.TransactOpts, _impl)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address _impl) returns(bool)
func (_Itreasuryproxy *ItreasuryproxyTransactorSession) UpgradeTo(_impl common.Address) (*types.Transaction, error) {
	return _Itreasuryproxy.Contract.UpgradeTo(&_Itreasuryproxy.TransactOpts, _impl)
}
