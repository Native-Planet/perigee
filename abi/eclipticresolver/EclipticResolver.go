// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eclipticresolver

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

// EclipticresolverMetaData contains all meta data concerning the Eclipticresolver contract.
var EclipticresolverMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_azimuth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405160208061034583398101806040528101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506102c2806100836000396000f30060806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806301ffc9a71461005e5780633b3b57de146100c2575b34801561005857600080fd5b50600080fd5b34801561006a57600080fd5b506100a860048036038101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190505050610133565b604051808215151515815260200191505060405180910390f35b3480156100ce57600080fd5b506100f160048036038101908080356000191690602001909291905050506101cd565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000633b3b57de7c010000000000000000000000000000000000000000000000000000000002827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806101c657506301ffc9a77c010000000000000000000000000000000000000000000000000000000002827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b9050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561025457600080fd5b505af1158015610268573d6000803e3d6000fd5b505050506040513d602081101561027e57600080fd5b810190808051906020019092919050505090509190505600a165627a7a7230582029cb4c821406e6c1c8f2e5609bbd3f752d990da94ccd6ef718556126758e9c300029",
}

// EclipticresolverABI is the input ABI used to generate the binding from.
// Deprecated: Use EclipticresolverMetaData.ABI instead.
var EclipticresolverABI = EclipticresolverMetaData.ABI

// EclipticresolverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EclipticresolverMetaData.Bin instead.
var EclipticresolverBin = EclipticresolverMetaData.Bin

// DeployEclipticresolver deploys a new Ethereum contract, binding an instance of Eclipticresolver to it.
func DeployEclipticresolver(auth *bind.TransactOpts, backend bind.ContractBackend, _azimuth common.Address) (common.Address, *types.Transaction, *Eclipticresolver, error) {
	parsed, err := EclipticresolverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EclipticresolverBin), backend, _azimuth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Eclipticresolver{EclipticresolverCaller: EclipticresolverCaller{contract: contract}, EclipticresolverTransactor: EclipticresolverTransactor{contract: contract}, EclipticresolverFilterer: EclipticresolverFilterer{contract: contract}}, nil
}

// Eclipticresolver is an auto generated Go binding around an Ethereum contract.
type Eclipticresolver struct {
	EclipticresolverCaller     // Read-only binding to the contract
	EclipticresolverTransactor // Write-only binding to the contract
	EclipticresolverFilterer   // Log filterer for contract events
}

// EclipticresolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type EclipticresolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EclipticresolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EclipticresolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EclipticresolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EclipticresolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EclipticresolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EclipticresolverSession struct {
	Contract     *Eclipticresolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EclipticresolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EclipticresolverCallerSession struct {
	Contract *EclipticresolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// EclipticresolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EclipticresolverTransactorSession struct {
	Contract     *EclipticresolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// EclipticresolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type EclipticresolverRaw struct {
	Contract *Eclipticresolver // Generic contract binding to access the raw methods on
}

// EclipticresolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EclipticresolverCallerRaw struct {
	Contract *EclipticresolverCaller // Generic read-only contract binding to access the raw methods on
}

// EclipticresolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EclipticresolverTransactorRaw struct {
	Contract *EclipticresolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEclipticresolver creates a new instance of Eclipticresolver, bound to a specific deployed contract.
func NewEclipticresolver(address common.Address, backend bind.ContractBackend) (*Eclipticresolver, error) {
	contract, err := bindEclipticresolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eclipticresolver{EclipticresolverCaller: EclipticresolverCaller{contract: contract}, EclipticresolverTransactor: EclipticresolverTransactor{contract: contract}, EclipticresolverFilterer: EclipticresolverFilterer{contract: contract}}, nil
}

// NewEclipticresolverCaller creates a new read-only instance of Eclipticresolver, bound to a specific deployed contract.
func NewEclipticresolverCaller(address common.Address, caller bind.ContractCaller) (*EclipticresolverCaller, error) {
	contract, err := bindEclipticresolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EclipticresolverCaller{contract: contract}, nil
}

// NewEclipticresolverTransactor creates a new write-only instance of Eclipticresolver, bound to a specific deployed contract.
func NewEclipticresolverTransactor(address common.Address, transactor bind.ContractTransactor) (*EclipticresolverTransactor, error) {
	contract, err := bindEclipticresolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EclipticresolverTransactor{contract: contract}, nil
}

// NewEclipticresolverFilterer creates a new log filterer instance of Eclipticresolver, bound to a specific deployed contract.
func NewEclipticresolverFilterer(address common.Address, filterer bind.ContractFilterer) (*EclipticresolverFilterer, error) {
	contract, err := bindEclipticresolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EclipticresolverFilterer{contract: contract}, nil
}

// bindEclipticresolver binds a generic wrapper to an already deployed contract.
func bindEclipticresolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EclipticresolverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eclipticresolver *EclipticresolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eclipticresolver.Contract.EclipticresolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eclipticresolver *EclipticresolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eclipticresolver.Contract.EclipticresolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eclipticresolver *EclipticresolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eclipticresolver.Contract.EclipticresolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eclipticresolver *EclipticresolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eclipticresolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eclipticresolver *EclipticresolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eclipticresolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eclipticresolver *EclipticresolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eclipticresolver.Contract.contract.Transact(opts, method, params...)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Eclipticresolver *EclipticresolverCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Eclipticresolver.contract.Call(opts, &out, "addr", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Eclipticresolver *EclipticresolverSession) Addr(node [32]byte) (common.Address, error) {
	return _Eclipticresolver.Contract.Addr(&_Eclipticresolver.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Eclipticresolver *EclipticresolverCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _Eclipticresolver.Contract.Addr(&_Eclipticresolver.CallOpts, node)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Eclipticresolver *EclipticresolverCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _Eclipticresolver.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Eclipticresolver *EclipticresolverSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Eclipticresolver.Contract.SupportsInterface(&_Eclipticresolver.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Eclipticresolver *EclipticresolverCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Eclipticresolver.Contract.SupportsInterface(&_Eclipticresolver.CallOpts, interfaceID)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Eclipticresolver *EclipticresolverTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Eclipticresolver.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Eclipticresolver *EclipticresolverSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Eclipticresolver.Contract.Fallback(&_Eclipticresolver.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Eclipticresolver *EclipticresolverTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Eclipticresolver.Contract.Fallback(&_Eclipticresolver.TransactOpts, calldata)
}
