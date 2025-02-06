// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package takespoints

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

// TakespointsMetaData contains all meta data concerning the Takespoints contract.
var TakespointsMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"azimuth\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_azimuth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405160208061016d8339810180604052810190808051906020019092919050505080806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505060e9806100846000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063d40ffacb146044575b600080fd5b348015604f57600080fd5b5060566098565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a723058207a888111ea13cf348ff926ba848ceb6234d0c25d088973e8e753de85c719e1ad0029",
}

// TakespointsABI is the input ABI used to generate the binding from.
// Deprecated: Use TakespointsMetaData.ABI instead.
var TakespointsABI = TakespointsMetaData.ABI

// TakespointsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TakespointsMetaData.Bin instead.
var TakespointsBin = TakespointsMetaData.Bin

// DeployTakespoints deploys a new Ethereum contract, binding an instance of Takespoints to it.
func DeployTakespoints(auth *bind.TransactOpts, backend bind.ContractBackend, _azimuth common.Address) (common.Address, *types.Transaction, *Takespoints, error) {
	parsed, err := TakespointsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TakespointsBin), backend, _azimuth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Takespoints{TakespointsCaller: TakespointsCaller{contract: contract}, TakespointsTransactor: TakespointsTransactor{contract: contract}, TakespointsFilterer: TakespointsFilterer{contract: contract}}, nil
}

// Takespoints is an auto generated Go binding around an Ethereum contract.
type Takespoints struct {
	TakespointsCaller     // Read-only binding to the contract
	TakespointsTransactor // Write-only binding to the contract
	TakespointsFilterer   // Log filterer for contract events
}

// TakespointsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TakespointsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TakespointsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TakespointsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TakespointsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TakespointsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TakespointsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TakespointsSession struct {
	Contract     *Takespoints      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TakespointsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TakespointsCallerSession struct {
	Contract *TakespointsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TakespointsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TakespointsTransactorSession struct {
	Contract     *TakespointsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TakespointsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TakespointsRaw struct {
	Contract *Takespoints // Generic contract binding to access the raw methods on
}

// TakespointsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TakespointsCallerRaw struct {
	Contract *TakespointsCaller // Generic read-only contract binding to access the raw methods on
}

// TakespointsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TakespointsTransactorRaw struct {
	Contract *TakespointsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTakespoints creates a new instance of Takespoints, bound to a specific deployed contract.
func NewTakespoints(address common.Address, backend bind.ContractBackend) (*Takespoints, error) {
	contract, err := bindTakespoints(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Takespoints{TakespointsCaller: TakespointsCaller{contract: contract}, TakespointsTransactor: TakespointsTransactor{contract: contract}, TakespointsFilterer: TakespointsFilterer{contract: contract}}, nil
}

// NewTakespointsCaller creates a new read-only instance of Takespoints, bound to a specific deployed contract.
func NewTakespointsCaller(address common.Address, caller bind.ContractCaller) (*TakespointsCaller, error) {
	contract, err := bindTakespoints(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TakespointsCaller{contract: contract}, nil
}

// NewTakespointsTransactor creates a new write-only instance of Takespoints, bound to a specific deployed contract.
func NewTakespointsTransactor(address common.Address, transactor bind.ContractTransactor) (*TakespointsTransactor, error) {
	contract, err := bindTakespoints(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TakespointsTransactor{contract: contract}, nil
}

// NewTakespointsFilterer creates a new log filterer instance of Takespoints, bound to a specific deployed contract.
func NewTakespointsFilterer(address common.Address, filterer bind.ContractFilterer) (*TakespointsFilterer, error) {
	contract, err := bindTakespoints(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TakespointsFilterer{contract: contract}, nil
}

// bindTakespoints binds a generic wrapper to an already deployed contract.
func bindTakespoints(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TakespointsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Takespoints *TakespointsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Takespoints.Contract.TakespointsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Takespoints *TakespointsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Takespoints.Contract.TakespointsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Takespoints *TakespointsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Takespoints.Contract.TakespointsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Takespoints *TakespointsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Takespoints.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Takespoints *TakespointsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Takespoints.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Takespoints *TakespointsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Takespoints.Contract.contract.Transact(opts, method, params...)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Takespoints *TakespointsCaller) Azimuth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Takespoints.contract.Call(opts, &out, "azimuth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Takespoints *TakespointsSession) Azimuth() (common.Address, error) {
	return _Takespoints.Contract.Azimuth(&_Takespoints.CallOpts)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Takespoints *TakespointsCallerSession) Azimuth() (common.Address, error) {
	return _Takespoints.Contract.Azimuth(&_Takespoints.CallOpts)
}
