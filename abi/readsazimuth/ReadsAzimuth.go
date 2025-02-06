// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package readsazimuth

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

// ReadsazimuthMetaData contains all meta data concerning the Readsazimuth contract.
var ReadsazimuthMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"azimuth\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_azimuth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405160208061016b83398101806040528101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505060e9806100826000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063d40ffacb146044575b600080fd5b348015604f57600080fd5b5060566098565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820e2a00cab990cea14429737e2351dc5f0a42f1636be3490a71a07b027080922130029",
}

// ReadsazimuthABI is the input ABI used to generate the binding from.
// Deprecated: Use ReadsazimuthMetaData.ABI instead.
var ReadsazimuthABI = ReadsazimuthMetaData.ABI

// ReadsazimuthBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReadsazimuthMetaData.Bin instead.
var ReadsazimuthBin = ReadsazimuthMetaData.Bin

// DeployReadsazimuth deploys a new Ethereum contract, binding an instance of Readsazimuth to it.
func DeployReadsazimuth(auth *bind.TransactOpts, backend bind.ContractBackend, _azimuth common.Address) (common.Address, *types.Transaction, *Readsazimuth, error) {
	parsed, err := ReadsazimuthMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReadsazimuthBin), backend, _azimuth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Readsazimuth{ReadsazimuthCaller: ReadsazimuthCaller{contract: contract}, ReadsazimuthTransactor: ReadsazimuthTransactor{contract: contract}, ReadsazimuthFilterer: ReadsazimuthFilterer{contract: contract}}, nil
}

// Readsazimuth is an auto generated Go binding around an Ethereum contract.
type Readsazimuth struct {
	ReadsazimuthCaller     // Read-only binding to the contract
	ReadsazimuthTransactor // Write-only binding to the contract
	ReadsazimuthFilterer   // Log filterer for contract events
}

// ReadsazimuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReadsazimuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadsazimuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReadsazimuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadsazimuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReadsazimuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadsazimuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReadsazimuthSession struct {
	Contract     *Readsazimuth     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReadsazimuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReadsazimuthCallerSession struct {
	Contract *ReadsazimuthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ReadsazimuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReadsazimuthTransactorSession struct {
	Contract     *ReadsazimuthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ReadsazimuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReadsazimuthRaw struct {
	Contract *Readsazimuth // Generic contract binding to access the raw methods on
}

// ReadsazimuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReadsazimuthCallerRaw struct {
	Contract *ReadsazimuthCaller // Generic read-only contract binding to access the raw methods on
}

// ReadsazimuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReadsazimuthTransactorRaw struct {
	Contract *ReadsazimuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReadsazimuth creates a new instance of Readsazimuth, bound to a specific deployed contract.
func NewReadsazimuth(address common.Address, backend bind.ContractBackend) (*Readsazimuth, error) {
	contract, err := bindReadsazimuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Readsazimuth{ReadsazimuthCaller: ReadsazimuthCaller{contract: contract}, ReadsazimuthTransactor: ReadsazimuthTransactor{contract: contract}, ReadsazimuthFilterer: ReadsazimuthFilterer{contract: contract}}, nil
}

// NewReadsazimuthCaller creates a new read-only instance of Readsazimuth, bound to a specific deployed contract.
func NewReadsazimuthCaller(address common.Address, caller bind.ContractCaller) (*ReadsazimuthCaller, error) {
	contract, err := bindReadsazimuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReadsazimuthCaller{contract: contract}, nil
}

// NewReadsazimuthTransactor creates a new write-only instance of Readsazimuth, bound to a specific deployed contract.
func NewReadsazimuthTransactor(address common.Address, transactor bind.ContractTransactor) (*ReadsazimuthTransactor, error) {
	contract, err := bindReadsazimuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReadsazimuthTransactor{contract: contract}, nil
}

// NewReadsazimuthFilterer creates a new log filterer instance of Readsazimuth, bound to a specific deployed contract.
func NewReadsazimuthFilterer(address common.Address, filterer bind.ContractFilterer) (*ReadsazimuthFilterer, error) {
	contract, err := bindReadsazimuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReadsazimuthFilterer{contract: contract}, nil
}

// bindReadsazimuth binds a generic wrapper to an already deployed contract.
func bindReadsazimuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReadsazimuthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Readsazimuth *ReadsazimuthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Readsazimuth.Contract.ReadsazimuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Readsazimuth *ReadsazimuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Readsazimuth.Contract.ReadsazimuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Readsazimuth *ReadsazimuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Readsazimuth.Contract.ReadsazimuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Readsazimuth *ReadsazimuthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Readsazimuth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Readsazimuth *ReadsazimuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Readsazimuth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Readsazimuth *ReadsazimuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Readsazimuth.Contract.contract.Transact(opts, method, params...)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Readsazimuth *ReadsazimuthCaller) Azimuth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Readsazimuth.contract.Call(opts, &out, "azimuth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Readsazimuth *ReadsazimuthSession) Azimuth() (common.Address, error) {
	return _Readsazimuth.Contract.Azimuth(&_Readsazimuth.CallOpts)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Readsazimuth *ReadsazimuthCallerSession) Azimuth() (common.Address, error) {
	return _Readsazimuth.Contract.Azimuth(&_Readsazimuth.CallOpts)
}
