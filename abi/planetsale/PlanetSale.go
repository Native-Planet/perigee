// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package planetsale

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

// PlanetsaleMetaData contains all meta data concerning the Planetsale contract.
var PlanetsaleMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"close\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_planet\",\"type\":\"uint32\"}],\"name\":\"purchase\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"azimuth\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_planet\",\"type\":\"uint32\"}],\"name\":\"available\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_azimuth\",\"type\":\"address\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"prefix\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"planet\",\"type\":\"uint32\"}],\"name\":\"PlanetSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405160408061109f8339810180604052810190808051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600010151561008c57600080fd5b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506100e5816100ec640100000000026401000000009004565b5050610160565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561014757600080fd5b80600010151561015657600080fd5b8060028190555050565b610f308061016f6000396000f3006080604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806351cff8d9146100a9578063715018a6146100ec5780638da5cb5b1461010357806391b7f5ed1461015a578063a035b1fe14610187578063c74073a1146101b2578063c7f04e65146101f5578063d40ffacb1461021b578063d57b4fca14610272578063f2fde38b146102bd575b600080fd5b3480156100b557600080fd5b506100ea600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610300565b005b3480156100f857600080fd5b506101016103e2565b005b34801561010f57600080fd5b506101186104e4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561016657600080fd5b5061018560048036038101908080359060200190929190505050610509565b005b34801561019357600080fd5b5061019c61057d565b6040518082815260200191505060405180910390f35b3480156101be57600080fd5b506101f3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610583565b005b610219600480360381019080803563ffffffff16906020019092919050505061061d565b005b34801561022757600080fd5b506102306109b1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561027e57600080fd5b506102a3600480360381019080803563ffffffff1690602001909291905050506109d7565b604051808215151515815260200191505060405180910390f35b3480156102c957600080fd5b506102fe600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610da3565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561035b57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff1660001415151561038157600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166108fc3073ffffffffffffffffffffffffffffffffffffffff16319081150290604051600060405180830381858888f193505050501580156103de573d6000803e3d6000fd5b5050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561043d57600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561056457600080fd5b80600010151561057357600080fd5b8060028190555050565b60025481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156105de57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff1660001415151561060457600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16ff5b6000600254341480156106355750610634826109d7565b5b151561064057600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156106c657600080fd5b505af11580156106da573d6000803e3d6000fd5b505050506040513d60208110156106f057600080fd5b810190808051906020019092919050505090508073ffffffffffffffffffffffffffffffffffffffff1663a0d3253f83306040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b1580156107b257600080fd5b505af11580156107c6573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff16631e79a85b833360006040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808463ffffffff1663ffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001821515151581526020019350505050600060405180830381600087803b15801561088657600080fd5b505af115801561089a573d6000803e3d6000fd5b505050508163ffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e4a358d7846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561094257600080fd5b505af1158015610956573d6000803e3d6000fd5b505050506040513d602081101561096c57600080fd5b810190808051906020019092919050505061ffff167f74bf1310a7a2b0bdd412eec8b0b11538ad4a803d561686380b4ccd083afc02c260405160405180910390a35050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e4a358d7846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015610a7757600080fd5b505af1158015610a8b573d6000803e3d6000fd5b505050506040513d6020811015610aa157600080fd5b81019080805190602001909291905050509050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663caf590f98460006040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015610b7057600080fd5b505af1158015610b84573d6000803e3d6000fd5b505050506040513d6020811015610b9a57600080fd5b81019080805190602001909291905050508015610cbd5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324ba1a4682306040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808361ffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015610c8157600080fd5b505af1158015610c95573d6000803e3d6000fd5b505050506040513d6020811015610cab57600080fd5b81019080805190602001909291905050505b8015610d9b5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636d09887b826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808261ffff1663ffffffff168152602001915050602060405180830381600087803b158015610d5f57600080fd5b505af1158015610d73573d6000803e3d6000fd5b505050506040513d6020811015610d8957600080fd5b81019080805190602001909291905050505b915050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610dfe57600080fd5b610e0781610e0a565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610e4657600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a72305820fe4309ca0beb80dca563ebe07c264c6aa63ed5c86f8cc51d3963a7e2f39928040029",
}

// PlanetsaleABI is the input ABI used to generate the binding from.
// Deprecated: Use PlanetsaleMetaData.ABI instead.
var PlanetsaleABI = PlanetsaleMetaData.ABI

// PlanetsaleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PlanetsaleMetaData.Bin instead.
var PlanetsaleBin = PlanetsaleMetaData.Bin

// DeployPlanetsale deploys a new Ethereum contract, binding an instance of Planetsale to it.
func DeployPlanetsale(auth *bind.TransactOpts, backend bind.ContractBackend, _azimuth common.Address, _price *big.Int) (common.Address, *types.Transaction, *Planetsale, error) {
	parsed, err := PlanetsaleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PlanetsaleBin), backend, _azimuth, _price)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Planetsale{PlanetsaleCaller: PlanetsaleCaller{contract: contract}, PlanetsaleTransactor: PlanetsaleTransactor{contract: contract}, PlanetsaleFilterer: PlanetsaleFilterer{contract: contract}}, nil
}

// Planetsale is an auto generated Go binding around an Ethereum contract.
type Planetsale struct {
	PlanetsaleCaller     // Read-only binding to the contract
	PlanetsaleTransactor // Write-only binding to the contract
	PlanetsaleFilterer   // Log filterer for contract events
}

// PlanetsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlanetsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlanetsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlanetsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlanetsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlanetsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlanetsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlanetsaleSession struct {
	Contract     *Planetsale       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlanetsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlanetsaleCallerSession struct {
	Contract *PlanetsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PlanetsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlanetsaleTransactorSession struct {
	Contract     *PlanetsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PlanetsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlanetsaleRaw struct {
	Contract *Planetsale // Generic contract binding to access the raw methods on
}

// PlanetsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlanetsaleCallerRaw struct {
	Contract *PlanetsaleCaller // Generic read-only contract binding to access the raw methods on
}

// PlanetsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlanetsaleTransactorRaw struct {
	Contract *PlanetsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlanetsale creates a new instance of Planetsale, bound to a specific deployed contract.
func NewPlanetsale(address common.Address, backend bind.ContractBackend) (*Planetsale, error) {
	contract, err := bindPlanetsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Planetsale{PlanetsaleCaller: PlanetsaleCaller{contract: contract}, PlanetsaleTransactor: PlanetsaleTransactor{contract: contract}, PlanetsaleFilterer: PlanetsaleFilterer{contract: contract}}, nil
}

// NewPlanetsaleCaller creates a new read-only instance of Planetsale, bound to a specific deployed contract.
func NewPlanetsaleCaller(address common.Address, caller bind.ContractCaller) (*PlanetsaleCaller, error) {
	contract, err := bindPlanetsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlanetsaleCaller{contract: contract}, nil
}

// NewPlanetsaleTransactor creates a new write-only instance of Planetsale, bound to a specific deployed contract.
func NewPlanetsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*PlanetsaleTransactor, error) {
	contract, err := bindPlanetsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlanetsaleTransactor{contract: contract}, nil
}

// NewPlanetsaleFilterer creates a new log filterer instance of Planetsale, bound to a specific deployed contract.
func NewPlanetsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*PlanetsaleFilterer, error) {
	contract, err := bindPlanetsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlanetsaleFilterer{contract: contract}, nil
}

// bindPlanetsale binds a generic wrapper to an already deployed contract.
func bindPlanetsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PlanetsaleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Planetsale *PlanetsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Planetsale.Contract.PlanetsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Planetsale *PlanetsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Planetsale.Contract.PlanetsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Planetsale *PlanetsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Planetsale.Contract.PlanetsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Planetsale *PlanetsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Planetsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Planetsale *PlanetsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Planetsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Planetsale *PlanetsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Planetsale.Contract.contract.Transact(opts, method, params...)
}

// Available is a free data retrieval call binding the contract method 0xd57b4fca.
//
// Solidity: function available(uint32 _planet) view returns(bool result)
func (_Planetsale *PlanetsaleCaller) Available(opts *bind.CallOpts, _planet uint32) (bool, error) {
	var out []interface{}
	err := _Planetsale.contract.Call(opts, &out, "available", _planet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0xd57b4fca.
//
// Solidity: function available(uint32 _planet) view returns(bool result)
func (_Planetsale *PlanetsaleSession) Available(_planet uint32) (bool, error) {
	return _Planetsale.Contract.Available(&_Planetsale.CallOpts, _planet)
}

// Available is a free data retrieval call binding the contract method 0xd57b4fca.
//
// Solidity: function available(uint32 _planet) view returns(bool result)
func (_Planetsale *PlanetsaleCallerSession) Available(_planet uint32) (bool, error) {
	return _Planetsale.Contract.Available(&_Planetsale.CallOpts, _planet)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Planetsale *PlanetsaleCaller) Azimuth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Planetsale.contract.Call(opts, &out, "azimuth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Planetsale *PlanetsaleSession) Azimuth() (common.Address, error) {
	return _Planetsale.Contract.Azimuth(&_Planetsale.CallOpts)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Planetsale *PlanetsaleCallerSession) Azimuth() (common.Address, error) {
	return _Planetsale.Contract.Azimuth(&_Planetsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Planetsale *PlanetsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Planetsale.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Planetsale *PlanetsaleSession) Owner() (common.Address, error) {
	return _Planetsale.Contract.Owner(&_Planetsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Planetsale *PlanetsaleCallerSession) Owner() (common.Address, error) {
	return _Planetsale.Contract.Owner(&_Planetsale.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Planetsale *PlanetsaleCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Planetsale.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Planetsale *PlanetsaleSession) Price() (*big.Int, error) {
	return _Planetsale.Contract.Price(&_Planetsale.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Planetsale *PlanetsaleCallerSession) Price() (*big.Int, error) {
	return _Planetsale.Contract.Price(&_Planetsale.CallOpts)
}

// Close is a paid mutator transaction binding the contract method 0xc74073a1.
//
// Solidity: function close(address _target) returns()
func (_Planetsale *PlanetsaleTransactor) Close(opts *bind.TransactOpts, _target common.Address) (*types.Transaction, error) {
	return _Planetsale.contract.Transact(opts, "close", _target)
}

// Close is a paid mutator transaction binding the contract method 0xc74073a1.
//
// Solidity: function close(address _target) returns()
func (_Planetsale *PlanetsaleSession) Close(_target common.Address) (*types.Transaction, error) {
	return _Planetsale.Contract.Close(&_Planetsale.TransactOpts, _target)
}

// Close is a paid mutator transaction binding the contract method 0xc74073a1.
//
// Solidity: function close(address _target) returns()
func (_Planetsale *PlanetsaleTransactorSession) Close(_target common.Address) (*types.Transaction, error) {
	return _Planetsale.Contract.Close(&_Planetsale.TransactOpts, _target)
}

// Purchase is a paid mutator transaction binding the contract method 0xc7f04e65.
//
// Solidity: function purchase(uint32 _planet) payable returns()
func (_Planetsale *PlanetsaleTransactor) Purchase(opts *bind.TransactOpts, _planet uint32) (*types.Transaction, error) {
	return _Planetsale.contract.Transact(opts, "purchase", _planet)
}

// Purchase is a paid mutator transaction binding the contract method 0xc7f04e65.
//
// Solidity: function purchase(uint32 _planet) payable returns()
func (_Planetsale *PlanetsaleSession) Purchase(_planet uint32) (*types.Transaction, error) {
	return _Planetsale.Contract.Purchase(&_Planetsale.TransactOpts, _planet)
}

// Purchase is a paid mutator transaction binding the contract method 0xc7f04e65.
//
// Solidity: function purchase(uint32 _planet) payable returns()
func (_Planetsale *PlanetsaleTransactorSession) Purchase(_planet uint32) (*types.Transaction, error) {
	return _Planetsale.Contract.Purchase(&_Planetsale.TransactOpts, _planet)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Planetsale *PlanetsaleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Planetsale.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Planetsale *PlanetsaleSession) RenounceOwnership() (*types.Transaction, error) {
	return _Planetsale.Contract.RenounceOwnership(&_Planetsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Planetsale *PlanetsaleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Planetsale.Contract.RenounceOwnership(&_Planetsale.TransactOpts)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_Planetsale *PlanetsaleTransactor) SetPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Planetsale.contract.Transact(opts, "setPrice", _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_Planetsale *PlanetsaleSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _Planetsale.Contract.SetPrice(&_Planetsale.TransactOpts, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_Planetsale *PlanetsaleTransactorSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _Planetsale.Contract.SetPrice(&_Planetsale.TransactOpts, _price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Planetsale *PlanetsaleTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Planetsale.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Planetsale *PlanetsaleSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Planetsale.Contract.TransferOwnership(&_Planetsale.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Planetsale *PlanetsaleTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Planetsale.Contract.TransferOwnership(&_Planetsale.TransactOpts, _newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _target) returns()
func (_Planetsale *PlanetsaleTransactor) Withdraw(opts *bind.TransactOpts, _target common.Address) (*types.Transaction, error) {
	return _Planetsale.contract.Transact(opts, "withdraw", _target)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _target) returns()
func (_Planetsale *PlanetsaleSession) Withdraw(_target common.Address) (*types.Transaction, error) {
	return _Planetsale.Contract.Withdraw(&_Planetsale.TransactOpts, _target)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _target) returns()
func (_Planetsale *PlanetsaleTransactorSession) Withdraw(_target common.Address) (*types.Transaction, error) {
	return _Planetsale.Contract.Withdraw(&_Planetsale.TransactOpts, _target)
}

// PlanetsaleOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Planetsale contract.
type PlanetsaleOwnershipRenouncedIterator struct {
	Event *PlanetsaleOwnershipRenounced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlanetsaleOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlanetsaleOwnershipRenounced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlanetsaleOwnershipRenounced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlanetsaleOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlanetsaleOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlanetsaleOwnershipRenounced represents a OwnershipRenounced event raised by the Planetsale contract.
type PlanetsaleOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Planetsale *PlanetsaleFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PlanetsaleOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Planetsale.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PlanetsaleOwnershipRenouncedIterator{contract: _Planetsale.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Planetsale *PlanetsaleFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PlanetsaleOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Planetsale.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlanetsaleOwnershipRenounced)
				if err := _Planetsale.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Planetsale *PlanetsaleFilterer) ParseOwnershipRenounced(log types.Log) (*PlanetsaleOwnershipRenounced, error) {
	event := new(PlanetsaleOwnershipRenounced)
	if err := _Planetsale.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlanetsaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Planetsale contract.
type PlanetsaleOwnershipTransferredIterator struct {
	Event *PlanetsaleOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlanetsaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlanetsaleOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlanetsaleOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlanetsaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlanetsaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlanetsaleOwnershipTransferred represents a OwnershipTransferred event raised by the Planetsale contract.
type PlanetsaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Planetsale *PlanetsaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PlanetsaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Planetsale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PlanetsaleOwnershipTransferredIterator{contract: _Planetsale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Planetsale *PlanetsaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PlanetsaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Planetsale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlanetsaleOwnershipTransferred)
				if err := _Planetsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Planetsale *PlanetsaleFilterer) ParseOwnershipTransferred(log types.Log) (*PlanetsaleOwnershipTransferred, error) {
	event := new(PlanetsaleOwnershipTransferred)
	if err := _Planetsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlanetsalePlanetSoldIterator is returned from FilterPlanetSold and is used to iterate over the raw logs and unpacked data for PlanetSold events raised by the Planetsale contract.
type PlanetsalePlanetSoldIterator struct {
	Event *PlanetsalePlanetSold // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlanetsalePlanetSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlanetsalePlanetSold)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlanetsalePlanetSold)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlanetsalePlanetSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlanetsalePlanetSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlanetsalePlanetSold represents a PlanetSold event raised by the Planetsale contract.
type PlanetsalePlanetSold struct {
	Prefix uint32
	Planet uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPlanetSold is a free log retrieval operation binding the contract event 0x74bf1310a7a2b0bdd412eec8b0b11538ad4a803d561686380b4ccd083afc02c2.
//
// Solidity: event PlanetSold(uint32 indexed prefix, uint32 indexed planet)
func (_Planetsale *PlanetsaleFilterer) FilterPlanetSold(opts *bind.FilterOpts, prefix []uint32, planet []uint32) (*PlanetsalePlanetSoldIterator, error) {

	var prefixRule []interface{}
	for _, prefixItem := range prefix {
		prefixRule = append(prefixRule, prefixItem)
	}
	var planetRule []interface{}
	for _, planetItem := range planet {
		planetRule = append(planetRule, planetItem)
	}

	logs, sub, err := _Planetsale.contract.FilterLogs(opts, "PlanetSold", prefixRule, planetRule)
	if err != nil {
		return nil, err
	}
	return &PlanetsalePlanetSoldIterator{contract: _Planetsale.contract, event: "PlanetSold", logs: logs, sub: sub}, nil
}

// WatchPlanetSold is a free log subscription operation binding the contract event 0x74bf1310a7a2b0bdd412eec8b0b11538ad4a803d561686380b4ccd083afc02c2.
//
// Solidity: event PlanetSold(uint32 indexed prefix, uint32 indexed planet)
func (_Planetsale *PlanetsaleFilterer) WatchPlanetSold(opts *bind.WatchOpts, sink chan<- *PlanetsalePlanetSold, prefix []uint32, planet []uint32) (event.Subscription, error) {

	var prefixRule []interface{}
	for _, prefixItem := range prefix {
		prefixRule = append(prefixRule, prefixItem)
	}
	var planetRule []interface{}
	for _, planetItem := range planet {
		planetRule = append(planetRule, planetItem)
	}

	logs, sub, err := _Planetsale.contract.WatchLogs(opts, "PlanetSold", prefixRule, planetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlanetsalePlanetSold)
				if err := _Planetsale.contract.UnpackLog(event, "PlanetSold", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePlanetSold is a log parse operation binding the contract event 0x74bf1310a7a2b0bdd412eec8b0b11538ad4a803d561686380b4ccd083afc02c2.
//
// Solidity: event PlanetSold(uint32 indexed prefix, uint32 indexed planet)
func (_Planetsale *PlanetsaleFilterer) ParsePlanetSold(log types.Log) (*PlanetsalePlanetSold, error) {
	event := new(PlanetsalePlanetSold)
	if err := _Planetsale.contract.UnpackLog(event, "PlanetSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
