// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nftokenreceivertestmock

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

// NftokenreceivertestmockMetaData contains all meta data concerning the Nftokenreceivertestmock contract.
var NftokenreceivertestmockMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061017a806100206000396000f300608060405260043610610041576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063150b7a0214610046575b600080fd5b34801561005257600080fd5b506100c9600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190820180359060200191909192939192939050505061011d565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b600063150b7a027c0100000000000000000000000000000000000000000000000000000000029050959450505050505600a165627a7a723058209cb039039b7f336a72927073c57b7515f540acd4087afe7341fc4363f91c40890029",
}

// NftokenreceivertestmockABI is the input ABI used to generate the binding from.
// Deprecated: Use NftokenreceivertestmockMetaData.ABI instead.
var NftokenreceivertestmockABI = NftokenreceivertestmockMetaData.ABI

// NftokenreceivertestmockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NftokenreceivertestmockMetaData.Bin instead.
var NftokenreceivertestmockBin = NftokenreceivertestmockMetaData.Bin

// DeployNftokenreceivertestmock deploys a new Ethereum contract, binding an instance of Nftokenreceivertestmock to it.
func DeployNftokenreceivertestmock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Nftokenreceivertestmock, error) {
	parsed, err := NftokenreceivertestmockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NftokenreceivertestmockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Nftokenreceivertestmock{NftokenreceivertestmockCaller: NftokenreceivertestmockCaller{contract: contract}, NftokenreceivertestmockTransactor: NftokenreceivertestmockTransactor{contract: contract}, NftokenreceivertestmockFilterer: NftokenreceivertestmockFilterer{contract: contract}}, nil
}

// Nftokenreceivertestmock is an auto generated Go binding around an Ethereum contract.
type Nftokenreceivertestmock struct {
	NftokenreceivertestmockCaller     // Read-only binding to the contract
	NftokenreceivertestmockTransactor // Write-only binding to the contract
	NftokenreceivertestmockFilterer   // Log filterer for contract events
}

// NftokenreceivertestmockCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftokenreceivertestmockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftokenreceivertestmockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftokenreceivertestmockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftokenreceivertestmockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftokenreceivertestmockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftokenreceivertestmockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftokenreceivertestmockSession struct {
	Contract     *Nftokenreceivertestmock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// NftokenreceivertestmockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftokenreceivertestmockCallerSession struct {
	Contract *NftokenreceivertestmockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// NftokenreceivertestmockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftokenreceivertestmockTransactorSession struct {
	Contract     *NftokenreceivertestmockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// NftokenreceivertestmockRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftokenreceivertestmockRaw struct {
	Contract *Nftokenreceivertestmock // Generic contract binding to access the raw methods on
}

// NftokenreceivertestmockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftokenreceivertestmockCallerRaw struct {
	Contract *NftokenreceivertestmockCaller // Generic read-only contract binding to access the raw methods on
}

// NftokenreceivertestmockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftokenreceivertestmockTransactorRaw struct {
	Contract *NftokenreceivertestmockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNftokenreceivertestmock creates a new instance of Nftokenreceivertestmock, bound to a specific deployed contract.
func NewNftokenreceivertestmock(address common.Address, backend bind.ContractBackend) (*Nftokenreceivertestmock, error) {
	contract, err := bindNftokenreceivertestmock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nftokenreceivertestmock{NftokenreceivertestmockCaller: NftokenreceivertestmockCaller{contract: contract}, NftokenreceivertestmockTransactor: NftokenreceivertestmockTransactor{contract: contract}, NftokenreceivertestmockFilterer: NftokenreceivertestmockFilterer{contract: contract}}, nil
}

// NewNftokenreceivertestmockCaller creates a new read-only instance of Nftokenreceivertestmock, bound to a specific deployed contract.
func NewNftokenreceivertestmockCaller(address common.Address, caller bind.ContractCaller) (*NftokenreceivertestmockCaller, error) {
	contract, err := bindNftokenreceivertestmock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NftokenreceivertestmockCaller{contract: contract}, nil
}

// NewNftokenreceivertestmockTransactor creates a new write-only instance of Nftokenreceivertestmock, bound to a specific deployed contract.
func NewNftokenreceivertestmockTransactor(address common.Address, transactor bind.ContractTransactor) (*NftokenreceivertestmockTransactor, error) {
	contract, err := bindNftokenreceivertestmock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NftokenreceivertestmockTransactor{contract: contract}, nil
}

// NewNftokenreceivertestmockFilterer creates a new log filterer instance of Nftokenreceivertestmock, bound to a specific deployed contract.
func NewNftokenreceivertestmockFilterer(address common.Address, filterer bind.ContractFilterer) (*NftokenreceivertestmockFilterer, error) {
	contract, err := bindNftokenreceivertestmock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NftokenreceivertestmockFilterer{contract: contract}, nil
}

// bindNftokenreceivertestmock binds a generic wrapper to an already deployed contract.
func bindNftokenreceivertestmock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NftokenreceivertestmockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nftokenreceivertestmock *NftokenreceivertestmockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nftokenreceivertestmock.Contract.NftokenreceivertestmockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nftokenreceivertestmock *NftokenreceivertestmockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nftokenreceivertestmock.Contract.NftokenreceivertestmockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nftokenreceivertestmock *NftokenreceivertestmockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nftokenreceivertestmock.Contract.NftokenreceivertestmockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nftokenreceivertestmock *NftokenreceivertestmockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nftokenreceivertestmock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nftokenreceivertestmock *NftokenreceivertestmockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nftokenreceivertestmock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nftokenreceivertestmock *NftokenreceivertestmockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nftokenreceivertestmock.Contract.contract.Transact(opts, method, params...)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_Nftokenreceivertestmock *NftokenreceivertestmockTransactor) OnERC721Received(opts *bind.TransactOpts, _operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Nftokenreceivertestmock.contract.Transact(opts, "onERC721Received", _operator, _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_Nftokenreceivertestmock *NftokenreceivertestmockSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Nftokenreceivertestmock.Contract.OnERC721Received(&_Nftokenreceivertestmock.TransactOpts, _operator, _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_Nftokenreceivertestmock *NftokenreceivertestmockTransactorSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Nftokenreceivertestmock.Contract.OnERC721Received(&_Nftokenreceivertestmock.TransactOpts, _operator, _from, _tokenId, _data)
}
