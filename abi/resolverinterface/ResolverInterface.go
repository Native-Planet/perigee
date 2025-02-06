// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package resolverinterface

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

// ResolverinterfaceMetaData contains all meta data concerning the Resolverinterface contract.
var ResolverinterfaceMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ResolverinterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use ResolverinterfaceMetaData.ABI instead.
var ResolverinterfaceABI = ResolverinterfaceMetaData.ABI

// Resolverinterface is an auto generated Go binding around an Ethereum contract.
type Resolverinterface struct {
	ResolverinterfaceCaller     // Read-only binding to the contract
	ResolverinterfaceTransactor // Write-only binding to the contract
	ResolverinterfaceFilterer   // Log filterer for contract events
}

// ResolverinterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResolverinterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverinterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResolverinterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverinterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResolverinterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverinterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResolverinterfaceSession struct {
	Contract     *Resolverinterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ResolverinterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResolverinterfaceCallerSession struct {
	Contract *ResolverinterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ResolverinterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResolverinterfaceTransactorSession struct {
	Contract     *ResolverinterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ResolverinterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResolverinterfaceRaw struct {
	Contract *Resolverinterface // Generic contract binding to access the raw methods on
}

// ResolverinterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResolverinterfaceCallerRaw struct {
	Contract *ResolverinterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ResolverinterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResolverinterfaceTransactorRaw struct {
	Contract *ResolverinterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResolverinterface creates a new instance of Resolverinterface, bound to a specific deployed contract.
func NewResolverinterface(address common.Address, backend bind.ContractBackend) (*Resolverinterface, error) {
	contract, err := bindResolverinterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Resolverinterface{ResolverinterfaceCaller: ResolverinterfaceCaller{contract: contract}, ResolverinterfaceTransactor: ResolverinterfaceTransactor{contract: contract}, ResolverinterfaceFilterer: ResolverinterfaceFilterer{contract: contract}}, nil
}

// NewResolverinterfaceCaller creates a new read-only instance of Resolverinterface, bound to a specific deployed contract.
func NewResolverinterfaceCaller(address common.Address, caller bind.ContractCaller) (*ResolverinterfaceCaller, error) {
	contract, err := bindResolverinterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverinterfaceCaller{contract: contract}, nil
}

// NewResolverinterfaceTransactor creates a new write-only instance of Resolverinterface, bound to a specific deployed contract.
func NewResolverinterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ResolverinterfaceTransactor, error) {
	contract, err := bindResolverinterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverinterfaceTransactor{contract: contract}, nil
}

// NewResolverinterfaceFilterer creates a new log filterer instance of Resolverinterface, bound to a specific deployed contract.
func NewResolverinterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ResolverinterfaceFilterer, error) {
	contract, err := bindResolverinterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResolverinterfaceFilterer{contract: contract}, nil
}

// bindResolverinterface binds a generic wrapper to an already deployed contract.
func bindResolverinterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ResolverinterfaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Resolverinterface *ResolverinterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Resolverinterface.Contract.ResolverinterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Resolverinterface *ResolverinterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Resolverinterface.Contract.ResolverinterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Resolverinterface *ResolverinterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Resolverinterface.Contract.ResolverinterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Resolverinterface *ResolverinterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Resolverinterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Resolverinterface *ResolverinterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Resolverinterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Resolverinterface *ResolverinterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Resolverinterface.Contract.contract.Transact(opts, method, params...)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Resolverinterface *ResolverinterfaceCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Resolverinterface.contract.Call(opts, &out, "addr", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Resolverinterface *ResolverinterfaceSession) Addr(node [32]byte) (common.Address, error) {
	return _Resolverinterface.Contract.Addr(&_Resolverinterface.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Resolverinterface *ResolverinterfaceCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _Resolverinterface.Contract.Addr(&_Resolverinterface.CallOpts, node)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Resolverinterface *ResolverinterfaceCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _Resolverinterface.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Resolverinterface *ResolverinterfaceSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Resolverinterface.Contract.SupportsInterface(&_Resolverinterface.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Resolverinterface *ResolverinterfaceCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Resolverinterface.Contract.SupportsInterface(&_Resolverinterface.CallOpts, interfaceID)
}
