// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc721receiver

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

// Erc721receiverMetaData contains all meta data concerning the Erc721receiver contract.
var Erc721receiverMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Erc721receiverABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc721receiverMetaData.ABI instead.
var Erc721receiverABI = Erc721receiverMetaData.ABI

// Erc721receiver is an auto generated Go binding around an Ethereum contract.
type Erc721receiver struct {
	Erc721receiverCaller     // Read-only binding to the contract
	Erc721receiverTransactor // Write-only binding to the contract
	Erc721receiverFilterer   // Log filterer for contract events
}

// Erc721receiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc721receiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721receiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc721receiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721receiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc721receiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721receiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc721receiverSession struct {
	Contract     *Erc721receiver   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc721receiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc721receiverCallerSession struct {
	Contract *Erc721receiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Erc721receiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc721receiverTransactorSession struct {
	Contract     *Erc721receiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Erc721receiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc721receiverRaw struct {
	Contract *Erc721receiver // Generic contract binding to access the raw methods on
}

// Erc721receiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc721receiverCallerRaw struct {
	Contract *Erc721receiverCaller // Generic read-only contract binding to access the raw methods on
}

// Erc721receiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc721receiverTransactorRaw struct {
	Contract *Erc721receiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc721receiver creates a new instance of Erc721receiver, bound to a specific deployed contract.
func NewErc721receiver(address common.Address, backend bind.ContractBackend) (*Erc721receiver, error) {
	contract, err := bindErc721receiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc721receiver{Erc721receiverCaller: Erc721receiverCaller{contract: contract}, Erc721receiverTransactor: Erc721receiverTransactor{contract: contract}, Erc721receiverFilterer: Erc721receiverFilterer{contract: contract}}, nil
}

// NewErc721receiverCaller creates a new read-only instance of Erc721receiver, bound to a specific deployed contract.
func NewErc721receiverCaller(address common.Address, caller bind.ContractCaller) (*Erc721receiverCaller, error) {
	contract, err := bindErc721receiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721receiverCaller{contract: contract}, nil
}

// NewErc721receiverTransactor creates a new write-only instance of Erc721receiver, bound to a specific deployed contract.
func NewErc721receiverTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc721receiverTransactor, error) {
	contract, err := bindErc721receiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721receiverTransactor{contract: contract}, nil
}

// NewErc721receiverFilterer creates a new log filterer instance of Erc721receiver, bound to a specific deployed contract.
func NewErc721receiverFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc721receiverFilterer, error) {
	contract, err := bindErc721receiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc721receiverFilterer{contract: contract}, nil
}

// bindErc721receiver binds a generic wrapper to an already deployed contract.
func bindErc721receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc721receiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721receiver *Erc721receiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721receiver.Contract.Erc721receiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721receiver *Erc721receiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721receiver.Contract.Erc721receiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721receiver *Erc721receiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721receiver.Contract.Erc721receiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721receiver *Erc721receiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721receiver *Erc721receiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721receiver *Erc721receiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721receiver.Contract.contract.Transact(opts, method, params...)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_Erc721receiver *Erc721receiverTransactor) OnERC721Received(opts *bind.TransactOpts, _operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Erc721receiver.contract.Transact(opts, "onERC721Received", _operator, _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_Erc721receiver *Erc721receiverSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Erc721receiver.Contract.OnERC721Received(&_Erc721receiver.TransactOpts, _operator, _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_Erc721receiver *Erc721receiverTransactorSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Erc721receiver.Contract.OnERC721Received(&_Erc721receiver.TransactOpts, _operator, _from, _tokenId, _data)
}
