// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package claims

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

// ClaimsMetaData contains all meta data concerning the Claims contract.
var ClaimsMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claims\",\"outputs\":[{\"name\":\"protocol\",\"type\":\"string\"},{\"name\":\"claim\",\"type\":\"string\"},{\"name\":\"dossier\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_whose\",\"type\":\"uint32\"},{\"name\":\"_protocol\",\"type\":\"string\"},{\"name\":\"_claim\",\"type\":\"string\"}],\"name\":\"findClaim\",\"outputs\":[{\"name\":\"index\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_protocol\",\"type\":\"string\"},{\"name\":\"_claim\",\"type\":\"string\"}],\"name\":\"removeClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"azimuth\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"clearClaims\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_protocol\",\"type\":\"string\"},{\"name\":\"_claim\",\"type\":\"string\"},{\"name\":\"_dossier\",\"type\":\"bytes\"}],\"name\":\"addClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_azimuth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"by\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"_protocol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_claim\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_dossier\",\"type\":\"bytes\"}],\"name\":\"ClaimAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"by\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"_protocol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_claim\",\"type\":\"string\"}],\"name\":\"ClaimRemoved\",\"type\":\"event\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516020806119378339810180604052810190808051906020019092919050505080806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506118b2806100856000396000f300608060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632906247e1461007d5780632945a57d1461020b578063296e3661146102e4578063d40ffacb14610347578063eaae46e51461039e578063fb1d8201146103d1575b600080fd5b34801561008957600080fd5b506100b8600480360381019080803563ffffffff1690602001909291908035906020019092919050505061044c565b60405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b838110156101005780820151818401526020810190506100e5565b50505050905090810190601f16801561012d5780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b8381101561016657808201518184015260208101905061014b565b50505050905090810190601f1680156101935780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b838110156101cc5780820151818401526020810190506101b1565b50505050905090810190601f1680156101f95780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b34801561021757600080fd5b506102c8600480360381019080803563ffffffff169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610650565b604051808260ff1660ff16815260200191505060405180910390f35b3480156102f057600080fd5b50610345600480360381019080803563ffffffff169060200190929190803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390505050610891565b005b34801561035357600080fd5b5061035c610bfb565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103aa57600080fd5b506103cf600480360381019080803563ffffffff169060200190929190505050610c20565b005b3480156103dd57600080fd5b5061044a600480360381019080803563ffffffff16906020019092919080359060200190820180359060200191909192939192939080359060200190820180359060200191909192939192939080359060200190820180359060200191909192939192939050505061106b565b005b60016020528160005260406000208160108110151561046757fe5b6003020160009150915050806000018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561050a5780601f106104df5761010080835404028352916020019161050a565b820191906000526020600020905b8154815290600101906020018083116104ed57829003601f168201915b505050505090806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105a85780601f1061057d576101008083540402835291602001916105a8565b820191906000526020600020905b81548152906001019060200180831161058b57829003601f168201915b505050505090806002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106465780601f1061061b57610100808354040283529160200191610646565b820191906000526020600020905b81548152906001019060200180831161062957829003601f168201915b5050505050905083565b600080600080600080876040518082805190602001908083835b60208310151561068f578051825260208201915060208101905060208303925061066a565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209450866040518082805190602001908083835b6020831015156106f457805182526020820191506020810190506020830392506106cf565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209350600160008a63ffffffff1663ffffffff1681526020019081526020016000209250600091505b601060ff168260ff16101561088057828260ff1660108110151561076957fe5b6003020190508060000160405180828054600181600116156101000203166002900480156107ce5780601f106107ac5761010080835404028352918201916107ce565b820191906000526020600020905b8154815290600101906020018083116107ba575b50509150506040518091039020600019168560001916148015610863575080600101604051808280546001816001161561010002031660029004801561084b5780601f1061082957610100808354040283529182019161084b565b820191906000526020600020905b815481529060010190602001808311610837575b50509150506040518091039020600019168460001916145b1561087357600182019550610885565b8180600101925050610749565b600095505b50505050509392505050565b6000856000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639137fe0a82336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561096457600080fd5b505af1158015610978573d6000803e3d6000fd5b505050506040513d602081101561098e57600080fd5b81019080805190602001909291905050508015610a7e57506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015610a4257600080fd5b505af1158015610a56573d6000803e3d6000fd5b505050506040513d6020811015610a6c57600080fd5b81019080805190602001909291905050505b1515610a8957600080fd5b610af88787878080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505086868080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050610650565b60ff169150600082111515610b0c57600080fd5b818060019003925050600160008863ffffffff1663ffffffff16815260200190815260200160002082601081101515610b4157fe5b6003020160008082016000610b5691906116d1565b600182016000610b6691906116d1565b600282016000610b769190611719565b50508663ffffffff167fdd924d662463d64f1eaae95a37d26f5bbbf9bbe2443adb897397e8b57c0b0513878787876040518080602001806020018381038352878782818152602001925080828437820191505083810382528585828181526020019250808284378201915050965050505050505060405180910390a250505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639137fe0a84336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015610cf357600080fd5b505af1158015610d07573d6000803e3d6000fd5b505050506040513d6020811015610d1d57600080fd5b810190808051906020019092919050505080610e2357506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610db957600080fd5b505af1158015610dcd573d6000803e3d6000fd5b505050506040513d6020811015610de357600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515610e2e57600080fd5b600160008463ffffffff1663ffffffff1681526020019081526020016000209150600090505b601060ff168160ff16101561106657818160ff16601081101515610e7457fe5b6003020160010180546001816001161561010002031660029004905060001015611011578263ffffffff167fdd924d662463d64f1eaae95a37d26f5bbbf9bbe2443adb897397e8b57c0b0513838360ff16601081101515610ed157fe5b60030201600001848460ff16601081101515610ee957fe5b60030201600101604051808060200180602001838103835285818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015610f7d5780601f10610f5257610100808354040283529160200191610f7d565b820191906000526020600020905b815481529060010190602001808311610f6057829003601f168201915b50508381038252848181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156110005780601f10610fd557610100808354040283529160200191611000565b820191906000526020600020905b815481529060010190602001808311610fe357829003601f168201915b505094505050505060405180910390a25b818160ff1660108110151561102257fe5b600302016000808201600061103791906116d1565b60018201600061104791906116d1565b6002820160006110579190611719565b50508080600101915050610e54565b505050565b600080886000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639137fe0a82336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561113f57600080fd5b505af1158015611153573d6000803e3d6000fd5b505050506040513d602081101561116957600080fd5b8101908080519060200190929190505050801561125957506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561121d57600080fd5b505af1158015611231573d6000803e3d6000fd5b505050506040513d602081101561124757600080fd5b81019080805190602001909291905050505b151561126457600080fd5b88889050600010801561127a5750868690506000105b151561128557600080fd5b6112f48a8a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505089898080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050610650565b925060008360ff1614156114515761130b8a61163c565b91506060604051908101604052808a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815260200188888080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815260200186868080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815250600160008c63ffffffff1663ffffffff1681526020019081526020016000208360ff166010811015156113ee57fe5b60030201600082015181600001908051906020019061140e929190611761565b50602082015181600101908051906020019061142b929190611761565b5060408201518160020190805190602001906114489291906117e1565b50905050611594565b6060604051908101604052808a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815260200188888080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815260200186868080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815250600160008c63ffffffff1663ffffffff1681526020019081526020016000206001850360ff1660108110151561153557fe5b600302016000820151816000019080519060200190611555929190611761565b506020820151816001019080519060200190611572929190611761565b50604082015181600201908051906020019061158f9291906117e1565b509050505b8963ffffffff167fef768946812a98aaa648b07282fa428f69903e34f6a38d8a9b208bd8ee53bb538a8a8a8a8a8a6040518080602001806020018060200184810384528a8a8281815260200192508082843782019150508481038352888882818152602001925080828437820191505084810382528686828181526020019250808284378201915050995050505050505050505060405180910390a250505050505050505050565b600080600080600160008663ffffffff1663ffffffff1681526020019081526020016000209250600091505b601060ff168260ff1610156116c457828260ff1660108110151561168857fe5b60030201905080600101805460018160011615610100020316600290049050600014156116b7578193506116c9565b8180600101925050611668565b600080fd5b505050919050565b50805460018160011615610100020316600290046000825580601f106116f75750611716565b601f0160209004906000526020600020908101906117159190611861565b5b50565b50805460018160011615610100020316600290046000825580601f1061173f575061175e565b601f01602090049060005260206000209081019061175d9190611861565b5b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106117a257805160ff19168380011785556117d0565b828001600101855582156117d0579182015b828111156117cf5782518255916020019190600101906117b4565b5b5090506117dd9190611861565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061182257805160ff1916838001178555611850565b82800160010185558215611850579182015b8281111561184f578251825591602001919060010190611834565b5b50905061185d9190611861565b5090565b61188391905b8082111561187f576000816000905550600101611867565b5090565b905600a165627a7a72305820916f0b03133ec7fbf76882deb3b5aa2a1fcbc161fb1b0c8b71a64ac703d3fb2a0029",
}

// ClaimsABI is the input ABI used to generate the binding from.
// Deprecated: Use ClaimsMetaData.ABI instead.
var ClaimsABI = ClaimsMetaData.ABI

// ClaimsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ClaimsMetaData.Bin instead.
var ClaimsBin = ClaimsMetaData.Bin

// DeployClaims deploys a new Ethereum contract, binding an instance of Claims to it.
func DeployClaims(auth *bind.TransactOpts, backend bind.ContractBackend, _azimuth common.Address) (common.Address, *types.Transaction, *Claims, error) {
	parsed, err := ClaimsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ClaimsBin), backend, _azimuth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Claims{ClaimsCaller: ClaimsCaller{contract: contract}, ClaimsTransactor: ClaimsTransactor{contract: contract}, ClaimsFilterer: ClaimsFilterer{contract: contract}}, nil
}

// Claims is an auto generated Go binding around an Ethereum contract.
type Claims struct {
	ClaimsCaller     // Read-only binding to the contract
	ClaimsTransactor // Write-only binding to the contract
	ClaimsFilterer   // Log filterer for contract events
}

// ClaimsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimsSession struct {
	Contract     *Claims           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimsCallerSession struct {
	Contract *ClaimsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ClaimsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimsTransactorSession struct {
	Contract     *ClaimsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimsRaw struct {
	Contract *Claims // Generic contract binding to access the raw methods on
}

// ClaimsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimsCallerRaw struct {
	Contract *ClaimsCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimsTransactorRaw struct {
	Contract *ClaimsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaims creates a new instance of Claims, bound to a specific deployed contract.
func NewClaims(address common.Address, backend bind.ContractBackend) (*Claims, error) {
	contract, err := bindClaims(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Claims{ClaimsCaller: ClaimsCaller{contract: contract}, ClaimsTransactor: ClaimsTransactor{contract: contract}, ClaimsFilterer: ClaimsFilterer{contract: contract}}, nil
}

// NewClaimsCaller creates a new read-only instance of Claims, bound to a specific deployed contract.
func NewClaimsCaller(address common.Address, caller bind.ContractCaller) (*ClaimsCaller, error) {
	contract, err := bindClaims(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimsCaller{contract: contract}, nil
}

// NewClaimsTransactor creates a new write-only instance of Claims, bound to a specific deployed contract.
func NewClaimsTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimsTransactor, error) {
	contract, err := bindClaims(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimsTransactor{contract: contract}, nil
}

// NewClaimsFilterer creates a new log filterer instance of Claims, bound to a specific deployed contract.
func NewClaimsFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimsFilterer, error) {
	contract, err := bindClaims(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimsFilterer{contract: contract}, nil
}

// bindClaims binds a generic wrapper to an already deployed contract.
func bindClaims(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClaimsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Claims *ClaimsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Claims.Contract.ClaimsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Claims *ClaimsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Claims.Contract.ClaimsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Claims *ClaimsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Claims.Contract.ClaimsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Claims *ClaimsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Claims.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Claims *ClaimsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Claims.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Claims *ClaimsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Claims.Contract.contract.Transact(opts, method, params...)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Claims *ClaimsCaller) Azimuth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "azimuth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Claims *ClaimsSession) Azimuth() (common.Address, error) {
	return _Claims.Contract.Azimuth(&_Claims.CallOpts)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Claims *ClaimsCallerSession) Azimuth() (common.Address, error) {
	return _Claims.Contract.Azimuth(&_Claims.CallOpts)
}

// Claims is a free data retrieval call binding the contract method 0x2906247e.
//
// Solidity: function claims(uint32 , uint256 ) view returns(string protocol, string claim, bytes dossier)
func (_Claims *ClaimsCaller) Claims(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (struct {
	Protocol string
	Claim    string
	Dossier  []byte
}, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "claims", arg0, arg1)

	outstruct := new(struct {
		Protocol string
		Claim    string
		Dossier  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Protocol = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Claim = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Dossier = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Claims is a free data retrieval call binding the contract method 0x2906247e.
//
// Solidity: function claims(uint32 , uint256 ) view returns(string protocol, string claim, bytes dossier)
func (_Claims *ClaimsSession) Claims(arg0 uint32, arg1 *big.Int) (struct {
	Protocol string
	Claim    string
	Dossier  []byte
}, error) {
	return _Claims.Contract.Claims(&_Claims.CallOpts, arg0, arg1)
}

// Claims is a free data retrieval call binding the contract method 0x2906247e.
//
// Solidity: function claims(uint32 , uint256 ) view returns(string protocol, string claim, bytes dossier)
func (_Claims *ClaimsCallerSession) Claims(arg0 uint32, arg1 *big.Int) (struct {
	Protocol string
	Claim    string
	Dossier  []byte
}, error) {
	return _Claims.Contract.Claims(&_Claims.CallOpts, arg0, arg1)
}

// FindClaim is a free data retrieval call binding the contract method 0x2945a57d.
//
// Solidity: function findClaim(uint32 _whose, string _protocol, string _claim) view returns(uint8 index)
func (_Claims *ClaimsCaller) FindClaim(opts *bind.CallOpts, _whose uint32, _protocol string, _claim string) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "findClaim", _whose, _protocol, _claim)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// FindClaim is a free data retrieval call binding the contract method 0x2945a57d.
//
// Solidity: function findClaim(uint32 _whose, string _protocol, string _claim) view returns(uint8 index)
func (_Claims *ClaimsSession) FindClaim(_whose uint32, _protocol string, _claim string) (uint8, error) {
	return _Claims.Contract.FindClaim(&_Claims.CallOpts, _whose, _protocol, _claim)
}

// FindClaim is a free data retrieval call binding the contract method 0x2945a57d.
//
// Solidity: function findClaim(uint32 _whose, string _protocol, string _claim) view returns(uint8 index)
func (_Claims *ClaimsCallerSession) FindClaim(_whose uint32, _protocol string, _claim string) (uint8, error) {
	return _Claims.Contract.FindClaim(&_Claims.CallOpts, _whose, _protocol, _claim)
}

// AddClaim is a paid mutator transaction binding the contract method 0xfb1d8201.
//
// Solidity: function addClaim(uint32 _point, string _protocol, string _claim, bytes _dossier) returns()
func (_Claims *ClaimsTransactor) AddClaim(opts *bind.TransactOpts, _point uint32, _protocol string, _claim string, _dossier []byte) (*types.Transaction, error) {
	return _Claims.contract.Transact(opts, "addClaim", _point, _protocol, _claim, _dossier)
}

// AddClaim is a paid mutator transaction binding the contract method 0xfb1d8201.
//
// Solidity: function addClaim(uint32 _point, string _protocol, string _claim, bytes _dossier) returns()
func (_Claims *ClaimsSession) AddClaim(_point uint32, _protocol string, _claim string, _dossier []byte) (*types.Transaction, error) {
	return _Claims.Contract.AddClaim(&_Claims.TransactOpts, _point, _protocol, _claim, _dossier)
}

// AddClaim is a paid mutator transaction binding the contract method 0xfb1d8201.
//
// Solidity: function addClaim(uint32 _point, string _protocol, string _claim, bytes _dossier) returns()
func (_Claims *ClaimsTransactorSession) AddClaim(_point uint32, _protocol string, _claim string, _dossier []byte) (*types.Transaction, error) {
	return _Claims.Contract.AddClaim(&_Claims.TransactOpts, _point, _protocol, _claim, _dossier)
}

// ClearClaims is a paid mutator transaction binding the contract method 0xeaae46e5.
//
// Solidity: function clearClaims(uint32 _point) returns()
func (_Claims *ClaimsTransactor) ClearClaims(opts *bind.TransactOpts, _point uint32) (*types.Transaction, error) {
	return _Claims.contract.Transact(opts, "clearClaims", _point)
}

// ClearClaims is a paid mutator transaction binding the contract method 0xeaae46e5.
//
// Solidity: function clearClaims(uint32 _point) returns()
func (_Claims *ClaimsSession) ClearClaims(_point uint32) (*types.Transaction, error) {
	return _Claims.Contract.ClearClaims(&_Claims.TransactOpts, _point)
}

// ClearClaims is a paid mutator transaction binding the contract method 0xeaae46e5.
//
// Solidity: function clearClaims(uint32 _point) returns()
func (_Claims *ClaimsTransactorSession) ClearClaims(_point uint32) (*types.Transaction, error) {
	return _Claims.Contract.ClearClaims(&_Claims.TransactOpts, _point)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x296e3661.
//
// Solidity: function removeClaim(uint32 _point, string _protocol, string _claim) returns()
func (_Claims *ClaimsTransactor) RemoveClaim(opts *bind.TransactOpts, _point uint32, _protocol string, _claim string) (*types.Transaction, error) {
	return _Claims.contract.Transact(opts, "removeClaim", _point, _protocol, _claim)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x296e3661.
//
// Solidity: function removeClaim(uint32 _point, string _protocol, string _claim) returns()
func (_Claims *ClaimsSession) RemoveClaim(_point uint32, _protocol string, _claim string) (*types.Transaction, error) {
	return _Claims.Contract.RemoveClaim(&_Claims.TransactOpts, _point, _protocol, _claim)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x296e3661.
//
// Solidity: function removeClaim(uint32 _point, string _protocol, string _claim) returns()
func (_Claims *ClaimsTransactorSession) RemoveClaim(_point uint32, _protocol string, _claim string) (*types.Transaction, error) {
	return _Claims.Contract.RemoveClaim(&_Claims.TransactOpts, _point, _protocol, _claim)
}

// ClaimsClaimAddedIterator is returned from FilterClaimAdded and is used to iterate over the raw logs and unpacked data for ClaimAdded events raised by the Claims contract.
type ClaimsClaimAddedIterator struct {
	Event *ClaimsClaimAdded // Event containing the contract specifics and raw log

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
func (it *ClaimsClaimAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsClaimAdded)
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
		it.Event = new(ClaimsClaimAdded)
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
func (it *ClaimsClaimAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsClaimAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsClaimAdded represents a ClaimAdded event raised by the Claims contract.
type ClaimsClaimAdded struct {
	By       uint32
	Protocol string
	Claim    string
	Dossier  []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimAdded is a free log retrieval operation binding the contract event 0xef768946812a98aaa648b07282fa428f69903e34f6a38d8a9b208bd8ee53bb53.
//
// Solidity: event ClaimAdded(uint32 indexed by, string _protocol, string _claim, bytes _dossier)
func (_Claims *ClaimsFilterer) FilterClaimAdded(opts *bind.FilterOpts, by []uint32) (*ClaimsClaimAddedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Claims.contract.FilterLogs(opts, "ClaimAdded", byRule)
	if err != nil {
		return nil, err
	}
	return &ClaimsClaimAddedIterator{contract: _Claims.contract, event: "ClaimAdded", logs: logs, sub: sub}, nil
}

// WatchClaimAdded is a free log subscription operation binding the contract event 0xef768946812a98aaa648b07282fa428f69903e34f6a38d8a9b208bd8ee53bb53.
//
// Solidity: event ClaimAdded(uint32 indexed by, string _protocol, string _claim, bytes _dossier)
func (_Claims *ClaimsFilterer) WatchClaimAdded(opts *bind.WatchOpts, sink chan<- *ClaimsClaimAdded, by []uint32) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Claims.contract.WatchLogs(opts, "ClaimAdded", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsClaimAdded)
				if err := _Claims.contract.UnpackLog(event, "ClaimAdded", log); err != nil {
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

// ParseClaimAdded is a log parse operation binding the contract event 0xef768946812a98aaa648b07282fa428f69903e34f6a38d8a9b208bd8ee53bb53.
//
// Solidity: event ClaimAdded(uint32 indexed by, string _protocol, string _claim, bytes _dossier)
func (_Claims *ClaimsFilterer) ParseClaimAdded(log types.Log) (*ClaimsClaimAdded, error) {
	event := new(ClaimsClaimAdded)
	if err := _Claims.contract.UnpackLog(event, "ClaimAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsClaimRemovedIterator is returned from FilterClaimRemoved and is used to iterate over the raw logs and unpacked data for ClaimRemoved events raised by the Claims contract.
type ClaimsClaimRemovedIterator struct {
	Event *ClaimsClaimRemoved // Event containing the contract specifics and raw log

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
func (it *ClaimsClaimRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsClaimRemoved)
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
		it.Event = new(ClaimsClaimRemoved)
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
func (it *ClaimsClaimRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsClaimRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsClaimRemoved represents a ClaimRemoved event raised by the Claims contract.
type ClaimsClaimRemoved struct {
	By       uint32
	Protocol string
	Claim    string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimRemoved is a free log retrieval operation binding the contract event 0xdd924d662463d64f1eaae95a37d26f5bbbf9bbe2443adb897397e8b57c0b0513.
//
// Solidity: event ClaimRemoved(uint32 indexed by, string _protocol, string _claim)
func (_Claims *ClaimsFilterer) FilterClaimRemoved(opts *bind.FilterOpts, by []uint32) (*ClaimsClaimRemovedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Claims.contract.FilterLogs(opts, "ClaimRemoved", byRule)
	if err != nil {
		return nil, err
	}
	return &ClaimsClaimRemovedIterator{contract: _Claims.contract, event: "ClaimRemoved", logs: logs, sub: sub}, nil
}

// WatchClaimRemoved is a free log subscription operation binding the contract event 0xdd924d662463d64f1eaae95a37d26f5bbbf9bbe2443adb897397e8b57c0b0513.
//
// Solidity: event ClaimRemoved(uint32 indexed by, string _protocol, string _claim)
func (_Claims *ClaimsFilterer) WatchClaimRemoved(opts *bind.WatchOpts, sink chan<- *ClaimsClaimRemoved, by []uint32) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Claims.contract.WatchLogs(opts, "ClaimRemoved", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsClaimRemoved)
				if err := _Claims.contract.UnpackLog(event, "ClaimRemoved", log); err != nil {
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

// ParseClaimRemoved is a log parse operation binding the contract event 0xdd924d662463d64f1eaae95a37d26f5bbbf9bbe2443adb897397e8b57c0b0513.
//
// Solidity: event ClaimRemoved(uint32 indexed by, string _protocol, string _claim)
func (_Claims *ClaimsFilterer) ParseClaimRemoved(log types.Log) (*ClaimsClaimRemoved, error) {
	event := new(ClaimsClaimRemoved)
	if err := _Claims.contract.UnpackLog(event, "ClaimRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
