// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package supportsinterfacewithlookup

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

// SupportsinterfacewithlookupMetaData contains all meta data concerning the Supportsinterfacewithlookup contract.
var SupportsinterfacewithlookupMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"InterfaceId_ERC165\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061004c6301ffc9a77c010000000000000000000000000000000000000000000000000000000002610051640100000000026401000000009004565b61010e565b63ffffffff7c010000000000000000000000000000000000000000000000000000000002817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916141515156100a257600080fd5b6001600080837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6101d88061011d6000396000f30060806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806301ffc9a71461005157806319fa8f50146100b5575b600080fd5b34801561005d57600080fd5b5061009b60048036038101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916906020019092919050505061011e565b604051808215151515815260200191505060405180910390f35b3480156100c157600080fd5b506100ca610185565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b6000806000837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060009054906101000a900460ff169050919050565b6301ffc9a77c010000000000000000000000000000000000000000000000000000000002815600a165627a7a723058208960b6becfd3c1282bb9e2503bc961c3b6f240c0d98f17a9ee1b69bc902950050029",
}

// SupportsinterfacewithlookupABI is the input ABI used to generate the binding from.
// Deprecated: Use SupportsinterfacewithlookupMetaData.ABI instead.
var SupportsinterfacewithlookupABI = SupportsinterfacewithlookupMetaData.ABI

// SupportsinterfacewithlookupBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SupportsinterfacewithlookupMetaData.Bin instead.
var SupportsinterfacewithlookupBin = SupportsinterfacewithlookupMetaData.Bin

// DeploySupportsinterfacewithlookup deploys a new Ethereum contract, binding an instance of Supportsinterfacewithlookup to it.
func DeploySupportsinterfacewithlookup(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Supportsinterfacewithlookup, error) {
	parsed, err := SupportsinterfacewithlookupMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SupportsinterfacewithlookupBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Supportsinterfacewithlookup{SupportsinterfacewithlookupCaller: SupportsinterfacewithlookupCaller{contract: contract}, SupportsinterfacewithlookupTransactor: SupportsinterfacewithlookupTransactor{contract: contract}, SupportsinterfacewithlookupFilterer: SupportsinterfacewithlookupFilterer{contract: contract}}, nil
}

// Supportsinterfacewithlookup is an auto generated Go binding around an Ethereum contract.
type Supportsinterfacewithlookup struct {
	SupportsinterfacewithlookupCaller     // Read-only binding to the contract
	SupportsinterfacewithlookupTransactor // Write-only binding to the contract
	SupportsinterfacewithlookupFilterer   // Log filterer for contract events
}

// SupportsinterfacewithlookupCaller is an auto generated read-only Go binding around an Ethereum contract.
type SupportsinterfacewithlookupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupportsinterfacewithlookupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SupportsinterfacewithlookupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupportsinterfacewithlookupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SupportsinterfacewithlookupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupportsinterfacewithlookupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SupportsinterfacewithlookupSession struct {
	Contract     *Supportsinterfacewithlookup // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SupportsinterfacewithlookupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SupportsinterfacewithlookupCallerSession struct {
	Contract *SupportsinterfacewithlookupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// SupportsinterfacewithlookupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SupportsinterfacewithlookupTransactorSession struct {
	Contract     *SupportsinterfacewithlookupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// SupportsinterfacewithlookupRaw is an auto generated low-level Go binding around an Ethereum contract.
type SupportsinterfacewithlookupRaw struct {
	Contract *Supportsinterfacewithlookup // Generic contract binding to access the raw methods on
}

// SupportsinterfacewithlookupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SupportsinterfacewithlookupCallerRaw struct {
	Contract *SupportsinterfacewithlookupCaller // Generic read-only contract binding to access the raw methods on
}

// SupportsinterfacewithlookupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SupportsinterfacewithlookupTransactorRaw struct {
	Contract *SupportsinterfacewithlookupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSupportsinterfacewithlookup creates a new instance of Supportsinterfacewithlookup, bound to a specific deployed contract.
func NewSupportsinterfacewithlookup(address common.Address, backend bind.ContractBackend) (*Supportsinterfacewithlookup, error) {
	contract, err := bindSupportsinterfacewithlookup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Supportsinterfacewithlookup{SupportsinterfacewithlookupCaller: SupportsinterfacewithlookupCaller{contract: contract}, SupportsinterfacewithlookupTransactor: SupportsinterfacewithlookupTransactor{contract: contract}, SupportsinterfacewithlookupFilterer: SupportsinterfacewithlookupFilterer{contract: contract}}, nil
}

// NewSupportsinterfacewithlookupCaller creates a new read-only instance of Supportsinterfacewithlookup, bound to a specific deployed contract.
func NewSupportsinterfacewithlookupCaller(address common.Address, caller bind.ContractCaller) (*SupportsinterfacewithlookupCaller, error) {
	contract, err := bindSupportsinterfacewithlookup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SupportsinterfacewithlookupCaller{contract: contract}, nil
}

// NewSupportsinterfacewithlookupTransactor creates a new write-only instance of Supportsinterfacewithlookup, bound to a specific deployed contract.
func NewSupportsinterfacewithlookupTransactor(address common.Address, transactor bind.ContractTransactor) (*SupportsinterfacewithlookupTransactor, error) {
	contract, err := bindSupportsinterfacewithlookup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SupportsinterfacewithlookupTransactor{contract: contract}, nil
}

// NewSupportsinterfacewithlookupFilterer creates a new log filterer instance of Supportsinterfacewithlookup, bound to a specific deployed contract.
func NewSupportsinterfacewithlookupFilterer(address common.Address, filterer bind.ContractFilterer) (*SupportsinterfacewithlookupFilterer, error) {
	contract, err := bindSupportsinterfacewithlookup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SupportsinterfacewithlookupFilterer{contract: contract}, nil
}

// bindSupportsinterfacewithlookup binds a generic wrapper to an already deployed contract.
func bindSupportsinterfacewithlookup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SupportsinterfacewithlookupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Supportsinterfacewithlookup *SupportsinterfacewithlookupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Supportsinterfacewithlookup.Contract.SupportsinterfacewithlookupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Supportsinterfacewithlookup *SupportsinterfacewithlookupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supportsinterfacewithlookup.Contract.SupportsinterfacewithlookupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Supportsinterfacewithlookup *SupportsinterfacewithlookupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Supportsinterfacewithlookup.Contract.SupportsinterfacewithlookupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Supportsinterfacewithlookup *SupportsinterfacewithlookupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Supportsinterfacewithlookup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Supportsinterfacewithlookup *SupportsinterfacewithlookupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supportsinterfacewithlookup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Supportsinterfacewithlookup *SupportsinterfacewithlookupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Supportsinterfacewithlookup.Contract.contract.Transact(opts, method, params...)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() view returns(bytes4)
func (_Supportsinterfacewithlookup *SupportsinterfacewithlookupCaller) InterfaceIdERC165(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Supportsinterfacewithlookup.contract.Call(opts, &out, "InterfaceId_ERC165")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() view returns(bytes4)
func (_Supportsinterfacewithlookup *SupportsinterfacewithlookupSession) InterfaceIdERC165() ([4]byte, error) {
	return _Supportsinterfacewithlookup.Contract.InterfaceIdERC165(&_Supportsinterfacewithlookup.CallOpts)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() view returns(bytes4)
func (_Supportsinterfacewithlookup *SupportsinterfacewithlookupCallerSession) InterfaceIdERC165() ([4]byte, error) {
	return _Supportsinterfacewithlookup.Contract.InterfaceIdERC165(&_Supportsinterfacewithlookup.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Supportsinterfacewithlookup *SupportsinterfacewithlookupCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Supportsinterfacewithlookup.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Supportsinterfacewithlookup *SupportsinterfacewithlookupSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Supportsinterfacewithlookup.Contract.SupportsInterface(&_Supportsinterfacewithlookup.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Supportsinterfacewithlookup *SupportsinterfacewithlookupCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Supportsinterfacewithlookup.Contract.SupportsInterface(&_Supportsinterfacewithlookup.CallOpts, _interfaceId)
}
