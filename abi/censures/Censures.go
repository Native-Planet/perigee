// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package censures

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

// CensuresMetaData contains all meta data concerning the Censures contract.
var CensuresMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"uint16\"}],\"name\":\"getCensuredBy\",\"outputs\":[{\"name\":\"cens\",\"type\":\"uint16[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint16\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"censuring\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_whose\",\"type\":\"uint16\"}],\"name\":\"getCensuringCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"censuredByIndexes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_as\",\"type\":\"uint16\"},{\"name\":\"_who\",\"type\":\"uint32\"}],\"name\":\"censure\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_whose\",\"type\":\"uint16\"}],\"name\":\"getCensuring\",\"outputs\":[{\"name\":\"cens\",\"type\":\"uint32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"azimuth\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_as\",\"type\":\"uint16\"},{\"name\":\"_who\",\"type\":\"uint32\"}],\"name\":\"forgive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint16\"},{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"censuringIndexes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"censuredBy\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"uint16\"}],\"name\":\"getCensuredByCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_azimuth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"by\",\"type\":\"uint16\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"uint32\"}],\"name\":\"Censured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"by\",\"type\":\"uint16\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"uint32\"}],\"name\":\"Forgiven\",\"type\":\"event\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516020806113948339810180604052810190808051906020019092919050505080806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505061130f806100856000396000f3006080604052600436106100af576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630fb7a5e6146100b457806320c2637e1461013a578063658cc376146101955780638618da88146101da5780638f7a47381461022f578063b6bcf35414610270578063d40ffacb146102f6578063d487758a1461034d578063e83113a81461038e578063e8a7dc05146103e3578063f98139be1461043c575b600080fd5b3480156100c057600080fd5b506100e3600480360381019080803561ffff169060200190929190505050610481565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561012657808201518184015260208101905061010b565b505050509050019250505060405180910390f35b34801561014657600080fd5b50610173600480360381019080803561ffff1690602001909291908035906020019092919050505061051e565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b3480156101a157600080fd5b506101c4600480360381019080803561ffff169060200190929190505050610566565b6040518082815260200191505060405180910390f35b3480156101e657600080fd5b50610219600480360381019080803563ffffffff169060200190929190803561ffff16906020019092919050505061058e565b6040518082815260200191505060405180910390f35b34801561023b57600080fd5b5061026e600480360381019080803561ffff169060200190929190803563ffffffff1690602001909291905050506105b3565b005b34801561027c57600080fd5b5061029f600480360381019080803561ffff169060200190929190505050610baf565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156102e25780820151818401526020810190506102c7565b505050509050019250505060405180910390f35b34801561030257600080fd5b5061030b610c4e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561035957600080fd5b5061038c600480360381019080803561ffff169060200190929190803563ffffffff169060200190929190505050610c73565b005b34801561039a57600080fd5b506103cd600480360381019080803561ffff169060200190929190803563ffffffff1690602001909291905050506111b5565b6040518082815260200191505060405180910390f35b3480156103ef57600080fd5b5061041e600480360381019080803563ffffffff169060200190929190803590602001909291905050506111da565b604051808261ffff1661ffff16815260200191505060405180910390f35b34801561044857600080fd5b5061046b600480360381019080803561ffff169060200190929190505050611220565b6040518082815260200191505060405180910390f35b6060600260008361ffff1663ffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561051257602002820191906000526020600020906000905b82829054906101000a900461ffff1661ffff16815260200190600201906020826001010492830192600103820291508084116104d95790505b50505050509050919050565b60016020528160005260406000208181548110151561053957fe5b9060005260206000209060089182820401919006600402915091509054906101000a900463ffffffff1681565b6000600160008361ffff1661ffff168152602001908152602001600020805490509050919050565b6004602052816000526040600020602052806000526040600020600091509150505481565b6000808361ffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639137fe0a82336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561068b57600080fd5b505af115801561069f573d6000803e3d6000fd5b505050506040513d60208110156106b557600080fd5b810190808051906020019092919050505080156107a557506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561076957600080fd5b505af115801561077d573d6000803e3d6000fd5b505050506040513d602081101561079357600080fd5b81019080805190602001909291905050505b15156107b057600080fd5b8363ffffffff168561ffff161415801561080257506000600360008761ffff1661ffff16815260200190815260200160002060008663ffffffff1663ffffffff16815260200190815260200160002054145b151561080d57600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166393976405866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808261ffff1663ffffffff168152602001915050602060405180830381600087803b1580156108a757600080fd5b505af11580156108bb573d6000803e3d6000fd5b505050506040513d60208110156108d157600080fd5b810190808051906020019092919050505092506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166393976405856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561098057600080fd5b505af1158015610994573d6000803e3d6000fd5b505050506040513d60208110156109aa57600080fd5b810190808051906020019092919050505091508260028111156109c957fe5b8260028111156109d557fe5b101515156109e257600080fd5b600160008661ffff1661ffff16815260200190815260200160002084908060018154018082558091505090600182039060005260206000209060089182820401919006600402909192909190916101000a81548163ffffffff021916908363ffffffff16021790555050600160008661ffff1661ffff16815260200190815260200160002080549050600360008761ffff1661ffff16815260200190815260200160002060008663ffffffff1663ffffffff16815260200190815260200160002081905550600260008563ffffffff1663ffffffff16815260200190815260200160002085908060018154018082558091505090600182039060005260206000209060109182820401919006600202909192909190916101000a81548161ffff021916908361ffff16021790555050600260008563ffffffff1663ffffffff16815260200190815260200160002080549050600460008663ffffffff1663ffffffff16815260200190815260200160002060008761ffff1661ffff168152602001908152602001600020819055508363ffffffff168561ffff167f672de430319eb045be436e23c1e2dde54bab11643b876afbcf00c6b8f290238a60405160405180910390a35050505050565b6060600160008361ffff1661ffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015610c4257602002820191906000526020600020906000905b82829054906101000a900463ffffffff1663ffffffff1681526020019060040190602082600301049283019260010382029150808411610c055790505b50505050509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000806000806000808961ffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639137fe0a82336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015610d5457600080fd5b505af1158015610d68573d6000803e3d6000fd5b505050506040513d6020811015610d7e57600080fd5b81019080805190602001909291905050508015610e6e57506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015610e3257600080fd5b505af1158015610e46573d6000803e3d6000fd5b505050506040513d6020811015610e5c57600080fd5b81019080805190602001909291905050505b1515610e7957600080fd5b600360008c61ffff1661ffff16815260200190815260200160002060008b63ffffffff1663ffffffff168152602001908152602001600020549850600460008b63ffffffff1663ffffffff16815260200190815260200160002060008c61ffff1661ffff168152602001908152602001600020549750600089118015610eff5750600088115b1515610f0a57600080fd5b888060019003995050878060019003985050600160008c61ffff1661ffff1681526020019081526020016000209650600260008b63ffffffff1663ffffffff168152602001908152602001600020955060018780549050039450600186805490500393508685815481101515610f7c57fe5b90600052602060002090600891828204019190066004029054906101000a900463ffffffff1692508584815481101515610fb257fe5b90600052602060002090601091828204019190066002029054906101000a900461ffff16915082878a815481101515610fe757fe5b90600052602060002090600891828204019190066004026101000a81548163ffffffff021916908363ffffffff16021790555081868981548110151561102957fe5b90600052602060002090601091828204019190066002026101000a81548161ffff021916908361ffff16021790555060018901600360008d61ffff1661ffff16815260200190815260200160002060008563ffffffff1663ffffffff1681526020019081526020016000208190555060018801600460008c63ffffffff1663ffffffff16815260200190815260200160002060008461ffff1661ffff168152602001908152602001600020819055508487816110e5919061124a565b508386816110f39190611284565b506000600360008d61ffff1661ffff16815260200190815260200160002060008c63ffffffff1663ffffffff168152602001908152602001600020819055506000600460008c63ffffffff1663ffffffff16815260200190815260200160002060008d61ffff1661ffff168152602001908152602001600020819055508963ffffffff168b61ffff167f3117bd64ae3afde1b1fbaf78e7e84647e30969d8804ce132b4278e96d71446ce60405160405180910390a35050505050505050505050565b6003602052816000526040600020602052806000526040600020600091509150505481565b6002602052816000526040600020818154811015156111f557fe5b9060005260206000209060109182820401919006600202915091509054906101000a900461ffff1681565b6000600260008361ffff1663ffffffff168152602001908152602001600020805490509050919050565b81548183558181111561127f57600701600890048160070160089004836000526020600020918201910161127e91906112be565b5b505050565b8154818355818111156112b957600f016010900481600f016010900483600052602060002091820191016112b891906112be565b5b505050565b6112e091905b808211156112dc5760008160009055506001016112c4565b5090565b905600a165627a7a723058207d08efc33e2d2cc83bb8edb2cc8d0ef58ff85e8eecafe7030f6bb130fc509aef0029",
}

// CensuresABI is the input ABI used to generate the binding from.
// Deprecated: Use CensuresMetaData.ABI instead.
var CensuresABI = CensuresMetaData.ABI

// CensuresBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CensuresMetaData.Bin instead.
var CensuresBin = CensuresMetaData.Bin

// DeployCensures deploys a new Ethereum contract, binding an instance of Censures to it.
func DeployCensures(auth *bind.TransactOpts, backend bind.ContractBackend, _azimuth common.Address) (common.Address, *types.Transaction, *Censures, error) {
	parsed, err := CensuresMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CensuresBin), backend, _azimuth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Censures{CensuresCaller: CensuresCaller{contract: contract}, CensuresTransactor: CensuresTransactor{contract: contract}, CensuresFilterer: CensuresFilterer{contract: contract}}, nil
}

// Censures is an auto generated Go binding around an Ethereum contract.
type Censures struct {
	CensuresCaller     // Read-only binding to the contract
	CensuresTransactor // Write-only binding to the contract
	CensuresFilterer   // Log filterer for contract events
}

// CensuresCaller is an auto generated read-only Go binding around an Ethereum contract.
type CensuresCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CensuresTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CensuresTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CensuresFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CensuresFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CensuresSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CensuresSession struct {
	Contract     *Censures         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CensuresCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CensuresCallerSession struct {
	Contract *CensuresCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CensuresTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CensuresTransactorSession struct {
	Contract     *CensuresTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CensuresRaw is an auto generated low-level Go binding around an Ethereum contract.
type CensuresRaw struct {
	Contract *Censures // Generic contract binding to access the raw methods on
}

// CensuresCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CensuresCallerRaw struct {
	Contract *CensuresCaller // Generic read-only contract binding to access the raw methods on
}

// CensuresTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CensuresTransactorRaw struct {
	Contract *CensuresTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCensures creates a new instance of Censures, bound to a specific deployed contract.
func NewCensures(address common.Address, backend bind.ContractBackend) (*Censures, error) {
	contract, err := bindCensures(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Censures{CensuresCaller: CensuresCaller{contract: contract}, CensuresTransactor: CensuresTransactor{contract: contract}, CensuresFilterer: CensuresFilterer{contract: contract}}, nil
}

// NewCensuresCaller creates a new read-only instance of Censures, bound to a specific deployed contract.
func NewCensuresCaller(address common.Address, caller bind.ContractCaller) (*CensuresCaller, error) {
	contract, err := bindCensures(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CensuresCaller{contract: contract}, nil
}

// NewCensuresTransactor creates a new write-only instance of Censures, bound to a specific deployed contract.
func NewCensuresTransactor(address common.Address, transactor bind.ContractTransactor) (*CensuresTransactor, error) {
	contract, err := bindCensures(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CensuresTransactor{contract: contract}, nil
}

// NewCensuresFilterer creates a new log filterer instance of Censures, bound to a specific deployed contract.
func NewCensuresFilterer(address common.Address, filterer bind.ContractFilterer) (*CensuresFilterer, error) {
	contract, err := bindCensures(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CensuresFilterer{contract: contract}, nil
}

// bindCensures binds a generic wrapper to an already deployed contract.
func bindCensures(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CensuresMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Censures *CensuresRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Censures.Contract.CensuresCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Censures *CensuresRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Censures.Contract.CensuresTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Censures *CensuresRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Censures.Contract.CensuresTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Censures *CensuresCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Censures.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Censures *CensuresTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Censures.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Censures *CensuresTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Censures.Contract.contract.Transact(opts, method, params...)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Censures *CensuresCaller) Azimuth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Censures.contract.Call(opts, &out, "azimuth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Censures *CensuresSession) Azimuth() (common.Address, error) {
	return _Censures.Contract.Azimuth(&_Censures.CallOpts)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Censures *CensuresCallerSession) Azimuth() (common.Address, error) {
	return _Censures.Contract.Azimuth(&_Censures.CallOpts)
}

// CensuredBy is a free data retrieval call binding the contract method 0xe8a7dc05.
//
// Solidity: function censuredBy(uint32 , uint256 ) view returns(uint16)
func (_Censures *CensuresCaller) CensuredBy(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (uint16, error) {
	var out []interface{}
	err := _Censures.contract.Call(opts, &out, "censuredBy", arg0, arg1)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// CensuredBy is a free data retrieval call binding the contract method 0xe8a7dc05.
//
// Solidity: function censuredBy(uint32 , uint256 ) view returns(uint16)
func (_Censures *CensuresSession) CensuredBy(arg0 uint32, arg1 *big.Int) (uint16, error) {
	return _Censures.Contract.CensuredBy(&_Censures.CallOpts, arg0, arg1)
}

// CensuredBy is a free data retrieval call binding the contract method 0xe8a7dc05.
//
// Solidity: function censuredBy(uint32 , uint256 ) view returns(uint16)
func (_Censures *CensuresCallerSession) CensuredBy(arg0 uint32, arg1 *big.Int) (uint16, error) {
	return _Censures.Contract.CensuredBy(&_Censures.CallOpts, arg0, arg1)
}

// CensuredByIndexes is a free data retrieval call binding the contract method 0x8618da88.
//
// Solidity: function censuredByIndexes(uint32 , uint16 ) view returns(uint256)
func (_Censures *CensuresCaller) CensuredByIndexes(opts *bind.CallOpts, arg0 uint32, arg1 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Censures.contract.Call(opts, &out, "censuredByIndexes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CensuredByIndexes is a free data retrieval call binding the contract method 0x8618da88.
//
// Solidity: function censuredByIndexes(uint32 , uint16 ) view returns(uint256)
func (_Censures *CensuresSession) CensuredByIndexes(arg0 uint32, arg1 uint16) (*big.Int, error) {
	return _Censures.Contract.CensuredByIndexes(&_Censures.CallOpts, arg0, arg1)
}

// CensuredByIndexes is a free data retrieval call binding the contract method 0x8618da88.
//
// Solidity: function censuredByIndexes(uint32 , uint16 ) view returns(uint256)
func (_Censures *CensuresCallerSession) CensuredByIndexes(arg0 uint32, arg1 uint16) (*big.Int, error) {
	return _Censures.Contract.CensuredByIndexes(&_Censures.CallOpts, arg0, arg1)
}

// Censuring is a free data retrieval call binding the contract method 0x20c2637e.
//
// Solidity: function censuring(uint16 , uint256 ) view returns(uint32)
func (_Censures *CensuresCaller) Censuring(opts *bind.CallOpts, arg0 uint16, arg1 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Censures.contract.Call(opts, &out, "censuring", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Censuring is a free data retrieval call binding the contract method 0x20c2637e.
//
// Solidity: function censuring(uint16 , uint256 ) view returns(uint32)
func (_Censures *CensuresSession) Censuring(arg0 uint16, arg1 *big.Int) (uint32, error) {
	return _Censures.Contract.Censuring(&_Censures.CallOpts, arg0, arg1)
}

// Censuring is a free data retrieval call binding the contract method 0x20c2637e.
//
// Solidity: function censuring(uint16 , uint256 ) view returns(uint32)
func (_Censures *CensuresCallerSession) Censuring(arg0 uint16, arg1 *big.Int) (uint32, error) {
	return _Censures.Contract.Censuring(&_Censures.CallOpts, arg0, arg1)
}

// CensuringIndexes is a free data retrieval call binding the contract method 0xe83113a8.
//
// Solidity: function censuringIndexes(uint16 , uint32 ) view returns(uint256)
func (_Censures *CensuresCaller) CensuringIndexes(opts *bind.CallOpts, arg0 uint16, arg1 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Censures.contract.Call(opts, &out, "censuringIndexes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CensuringIndexes is a free data retrieval call binding the contract method 0xe83113a8.
//
// Solidity: function censuringIndexes(uint16 , uint32 ) view returns(uint256)
func (_Censures *CensuresSession) CensuringIndexes(arg0 uint16, arg1 uint32) (*big.Int, error) {
	return _Censures.Contract.CensuringIndexes(&_Censures.CallOpts, arg0, arg1)
}

// CensuringIndexes is a free data retrieval call binding the contract method 0xe83113a8.
//
// Solidity: function censuringIndexes(uint16 , uint32 ) view returns(uint256)
func (_Censures *CensuresCallerSession) CensuringIndexes(arg0 uint16, arg1 uint32) (*big.Int, error) {
	return _Censures.Contract.CensuringIndexes(&_Censures.CallOpts, arg0, arg1)
}

// GetCensuredBy is a free data retrieval call binding the contract method 0x0fb7a5e6.
//
// Solidity: function getCensuredBy(uint16 _who) view returns(uint16[] cens)
func (_Censures *CensuresCaller) GetCensuredBy(opts *bind.CallOpts, _who uint16) ([]uint16, error) {
	var out []interface{}
	err := _Censures.contract.Call(opts, &out, "getCensuredBy", _who)

	if err != nil {
		return *new([]uint16), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint16)).(*[]uint16)

	return out0, err

}

// GetCensuredBy is a free data retrieval call binding the contract method 0x0fb7a5e6.
//
// Solidity: function getCensuredBy(uint16 _who) view returns(uint16[] cens)
func (_Censures *CensuresSession) GetCensuredBy(_who uint16) ([]uint16, error) {
	return _Censures.Contract.GetCensuredBy(&_Censures.CallOpts, _who)
}

// GetCensuredBy is a free data retrieval call binding the contract method 0x0fb7a5e6.
//
// Solidity: function getCensuredBy(uint16 _who) view returns(uint16[] cens)
func (_Censures *CensuresCallerSession) GetCensuredBy(_who uint16) ([]uint16, error) {
	return _Censures.Contract.GetCensuredBy(&_Censures.CallOpts, _who)
}

// GetCensuredByCount is a free data retrieval call binding the contract method 0xf98139be.
//
// Solidity: function getCensuredByCount(uint16 _who) view returns(uint256 count)
func (_Censures *CensuresCaller) GetCensuredByCount(opts *bind.CallOpts, _who uint16) (*big.Int, error) {
	var out []interface{}
	err := _Censures.contract.Call(opts, &out, "getCensuredByCount", _who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCensuredByCount is a free data retrieval call binding the contract method 0xf98139be.
//
// Solidity: function getCensuredByCount(uint16 _who) view returns(uint256 count)
func (_Censures *CensuresSession) GetCensuredByCount(_who uint16) (*big.Int, error) {
	return _Censures.Contract.GetCensuredByCount(&_Censures.CallOpts, _who)
}

// GetCensuredByCount is a free data retrieval call binding the contract method 0xf98139be.
//
// Solidity: function getCensuredByCount(uint16 _who) view returns(uint256 count)
func (_Censures *CensuresCallerSession) GetCensuredByCount(_who uint16) (*big.Int, error) {
	return _Censures.Contract.GetCensuredByCount(&_Censures.CallOpts, _who)
}

// GetCensuring is a free data retrieval call binding the contract method 0xb6bcf354.
//
// Solidity: function getCensuring(uint16 _whose) view returns(uint32[] cens)
func (_Censures *CensuresCaller) GetCensuring(opts *bind.CallOpts, _whose uint16) ([]uint32, error) {
	var out []interface{}
	err := _Censures.contract.Call(opts, &out, "getCensuring", _whose)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetCensuring is a free data retrieval call binding the contract method 0xb6bcf354.
//
// Solidity: function getCensuring(uint16 _whose) view returns(uint32[] cens)
func (_Censures *CensuresSession) GetCensuring(_whose uint16) ([]uint32, error) {
	return _Censures.Contract.GetCensuring(&_Censures.CallOpts, _whose)
}

// GetCensuring is a free data retrieval call binding the contract method 0xb6bcf354.
//
// Solidity: function getCensuring(uint16 _whose) view returns(uint32[] cens)
func (_Censures *CensuresCallerSession) GetCensuring(_whose uint16) ([]uint32, error) {
	return _Censures.Contract.GetCensuring(&_Censures.CallOpts, _whose)
}

// GetCensuringCount is a free data retrieval call binding the contract method 0x658cc376.
//
// Solidity: function getCensuringCount(uint16 _whose) view returns(uint256 count)
func (_Censures *CensuresCaller) GetCensuringCount(opts *bind.CallOpts, _whose uint16) (*big.Int, error) {
	var out []interface{}
	err := _Censures.contract.Call(opts, &out, "getCensuringCount", _whose)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCensuringCount is a free data retrieval call binding the contract method 0x658cc376.
//
// Solidity: function getCensuringCount(uint16 _whose) view returns(uint256 count)
func (_Censures *CensuresSession) GetCensuringCount(_whose uint16) (*big.Int, error) {
	return _Censures.Contract.GetCensuringCount(&_Censures.CallOpts, _whose)
}

// GetCensuringCount is a free data retrieval call binding the contract method 0x658cc376.
//
// Solidity: function getCensuringCount(uint16 _whose) view returns(uint256 count)
func (_Censures *CensuresCallerSession) GetCensuringCount(_whose uint16) (*big.Int, error) {
	return _Censures.Contract.GetCensuringCount(&_Censures.CallOpts, _whose)
}

// Censure is a paid mutator transaction binding the contract method 0x8f7a4738.
//
// Solidity: function censure(uint16 _as, uint32 _who) returns()
func (_Censures *CensuresTransactor) Censure(opts *bind.TransactOpts, _as uint16, _who uint32) (*types.Transaction, error) {
	return _Censures.contract.Transact(opts, "censure", _as, _who)
}

// Censure is a paid mutator transaction binding the contract method 0x8f7a4738.
//
// Solidity: function censure(uint16 _as, uint32 _who) returns()
func (_Censures *CensuresSession) Censure(_as uint16, _who uint32) (*types.Transaction, error) {
	return _Censures.Contract.Censure(&_Censures.TransactOpts, _as, _who)
}

// Censure is a paid mutator transaction binding the contract method 0x8f7a4738.
//
// Solidity: function censure(uint16 _as, uint32 _who) returns()
func (_Censures *CensuresTransactorSession) Censure(_as uint16, _who uint32) (*types.Transaction, error) {
	return _Censures.Contract.Censure(&_Censures.TransactOpts, _as, _who)
}

// Forgive is a paid mutator transaction binding the contract method 0xd487758a.
//
// Solidity: function forgive(uint16 _as, uint32 _who) returns()
func (_Censures *CensuresTransactor) Forgive(opts *bind.TransactOpts, _as uint16, _who uint32) (*types.Transaction, error) {
	return _Censures.contract.Transact(opts, "forgive", _as, _who)
}

// Forgive is a paid mutator transaction binding the contract method 0xd487758a.
//
// Solidity: function forgive(uint16 _as, uint32 _who) returns()
func (_Censures *CensuresSession) Forgive(_as uint16, _who uint32) (*types.Transaction, error) {
	return _Censures.Contract.Forgive(&_Censures.TransactOpts, _as, _who)
}

// Forgive is a paid mutator transaction binding the contract method 0xd487758a.
//
// Solidity: function forgive(uint16 _as, uint32 _who) returns()
func (_Censures *CensuresTransactorSession) Forgive(_as uint16, _who uint32) (*types.Transaction, error) {
	return _Censures.Contract.Forgive(&_Censures.TransactOpts, _as, _who)
}

// CensuresCensuredIterator is returned from FilterCensured and is used to iterate over the raw logs and unpacked data for Censured events raised by the Censures contract.
type CensuresCensuredIterator struct {
	Event *CensuresCensured // Event containing the contract specifics and raw log

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
func (it *CensuresCensuredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CensuresCensured)
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
		it.Event = new(CensuresCensured)
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
func (it *CensuresCensuredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CensuresCensuredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CensuresCensured represents a Censured event raised by the Censures contract.
type CensuresCensured struct {
	By  uint16
	Who uint32
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCensured is a free log retrieval operation binding the contract event 0x672de430319eb045be436e23c1e2dde54bab11643b876afbcf00c6b8f290238a.
//
// Solidity: event Censured(uint16 indexed by, uint32 indexed who)
func (_Censures *CensuresFilterer) FilterCensured(opts *bind.FilterOpts, by []uint16, who []uint32) (*CensuresCensuredIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Censures.contract.FilterLogs(opts, "Censured", byRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &CensuresCensuredIterator{contract: _Censures.contract, event: "Censured", logs: logs, sub: sub}, nil
}

// WatchCensured is a free log subscription operation binding the contract event 0x672de430319eb045be436e23c1e2dde54bab11643b876afbcf00c6b8f290238a.
//
// Solidity: event Censured(uint16 indexed by, uint32 indexed who)
func (_Censures *CensuresFilterer) WatchCensured(opts *bind.WatchOpts, sink chan<- *CensuresCensured, by []uint16, who []uint32) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Censures.contract.WatchLogs(opts, "Censured", byRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CensuresCensured)
				if err := _Censures.contract.UnpackLog(event, "Censured", log); err != nil {
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

// ParseCensured is a log parse operation binding the contract event 0x672de430319eb045be436e23c1e2dde54bab11643b876afbcf00c6b8f290238a.
//
// Solidity: event Censured(uint16 indexed by, uint32 indexed who)
func (_Censures *CensuresFilterer) ParseCensured(log types.Log) (*CensuresCensured, error) {
	event := new(CensuresCensured)
	if err := _Censures.contract.UnpackLog(event, "Censured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CensuresForgivenIterator is returned from FilterForgiven and is used to iterate over the raw logs and unpacked data for Forgiven events raised by the Censures contract.
type CensuresForgivenIterator struct {
	Event *CensuresForgiven // Event containing the contract specifics and raw log

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
func (it *CensuresForgivenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CensuresForgiven)
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
		it.Event = new(CensuresForgiven)
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
func (it *CensuresForgivenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CensuresForgivenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CensuresForgiven represents a Forgiven event raised by the Censures contract.
type CensuresForgiven struct {
	By  uint16
	Who uint32
	Raw types.Log // Blockchain specific contextual infos
}

// FilterForgiven is a free log retrieval operation binding the contract event 0x3117bd64ae3afde1b1fbaf78e7e84647e30969d8804ce132b4278e96d71446ce.
//
// Solidity: event Forgiven(uint16 indexed by, uint32 indexed who)
func (_Censures *CensuresFilterer) FilterForgiven(opts *bind.FilterOpts, by []uint16, who []uint32) (*CensuresForgivenIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Censures.contract.FilterLogs(opts, "Forgiven", byRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &CensuresForgivenIterator{contract: _Censures.contract, event: "Forgiven", logs: logs, sub: sub}, nil
}

// WatchForgiven is a free log subscription operation binding the contract event 0x3117bd64ae3afde1b1fbaf78e7e84647e30969d8804ce132b4278e96d71446ce.
//
// Solidity: event Forgiven(uint16 indexed by, uint32 indexed who)
func (_Censures *CensuresFilterer) WatchForgiven(opts *bind.WatchOpts, sink chan<- *CensuresForgiven, by []uint16, who []uint32) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Censures.contract.WatchLogs(opts, "Forgiven", byRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CensuresForgiven)
				if err := _Censures.contract.UnpackLog(event, "Forgiven", log); err != nil {
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

// ParseForgiven is a log parse operation binding the contract event 0x3117bd64ae3afde1b1fbaf78e7e84647e30969d8804ce132b4278e96d71446ce.
//
// Solidity: event Forgiven(uint16 indexed by, uint32 indexed who)
func (_Censures *CensuresFilterer) ParseForgiven(log types.Log) (*CensuresForgiven, error) {
	event := new(CensuresForgiven)
	if err := _Censures.contract.UnpackLog(event, "Forgiven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
