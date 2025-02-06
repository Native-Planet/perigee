// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package publicresolver

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

// PublicresolverMetaData contains all meta data concerning the Publicresolver contract.
var PublicresolverMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setText\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"contentTypes\",\"type\":\"uint256\"}],\"name\":\"ABI\",\"outputs\":[{\"name\":\"contentType\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"x\",\"type\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"setPubkey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"content\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"contentType\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setABI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"setMultihash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"setContent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"pubkey\",\"outputs\":[{\"name\":\"x\",\"type\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"multihash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"ensAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ContentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"contentType\",\"type\":\"uint256\"}],\"name\":\"ABIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"x\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"PubkeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"key\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"MultihashChanged\",\"type\":\"event\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051602080611f8083398101806040528101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050611efd806100836000396000f3006080604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806301ffc9a7146100e057806310f13a8c146101445780632203ab561461020157806329cd62ea146102bc5780632dff6941146103095780633b3b57de1461035657806359d1d43c146103c7578063623195b0146104b7578063691f34311461053857806377372213146105e2578063aa4cb54714610659578063c3d014d6146106d0578063c86902331461070f578063d5fa2b001461076b578063e89401a1146107bc575b600080fd5b3480156100ec57600080fd5b5061012a60048036038101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190505050610866565b604051808215151515815260200191505060405180910390f35b34801561015057600080fd5b506101ff6004803603810190808035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610bc0565b005b34801561020d57600080fd5b5061023a600480360381019080803560001916906020019092919080359060200190929190505050610e7d565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610280578082015181840152602081019050610265565b50505050905090810190601f1680156102ad5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b3480156102c857600080fd5b50610307600480360381019080803560001916906020019092919080356000191690602001909291908035600019169060200190929190505050610fc0565b005b34801561031557600080fd5b506103386004803603810190808035600019169060200190929190505050611185565b60405180826000191660001916815260200191505060405180910390f35b34801561036257600080fd5b5061038560048036038101908080356000191690602001909291905050506111ad565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103d357600080fd5b5061043c6004803603810190808035600019169060200190929190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506111f5565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561047c578082015181840152602081019050610461565b50505050905090810190601f1680156104a95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156104c357600080fd5b50610536600480360381019080803560001916906020019092919080359060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061131f565b005b34801561054457600080fd5b5061056760048036038101908080356000191690602001909291905050506114bc565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105a757808201518184015260208101905061058c565b50505050905090810190601f1680156105d45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156105ee57600080fd5b506106576004803603810190808035600019169060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061157c565b005b34801561066557600080fd5b506106ce6004803603810190808035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611762565b005b3480156106dc57600080fd5b5061070d60048036038101908080356000191690602001909291908035600019169060200190929190505050611948565b005b34801561071b57600080fd5b5061073e6004803603810190808035600019169060200190929190505050611ac5565b60405180836000191660001916815260200182600019166000191681526020019250505060405180910390f35b34801561077757600080fd5b506107ba6004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611b15565b005b3480156107c857600080fd5b506107eb6004803603810190808035600019169060200190929190505050611cec565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561082b578082015181840152602081019050610810565b50505050905090810190601f1680156108585780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6000633b3b57de7c0100000000000000000000000000000000000000000000000000000000027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610937575063d8389dc57c0100000000000000000000000000000000000000000000000000000000027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806109a2575063691f34317c0100000000000000000000000000000000000000000000000000000000027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610a0d5750632203ab567c0100000000000000000000000000000000000000000000000000000000027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610a78575063c86902337c0100000000000000000000000000000000000000000000000000000000027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610ae357506359d1d43c7c0100000000000000000000000000000000000000000000000000000000027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610b4e575063e89401a17c0100000000000000000000000000000000000000000000000000000000027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610bb957506301ffc9a77c0100000000000000000000000000000000000000000000000000000000027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b9050919050565b823373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166302571be3836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b158015610c7057600080fd5b505af1158015610c84573d6000803e3d6000fd5b505050506040513d6020811015610c9a57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff16141515610ccd57600080fd5b81600160008660001916600019168152602001908152602001600020600501846040518082805190602001908083835b602083101515610d225780518252602082019150602081019050602083039250610cfd565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390209080519060200190610d68929190611dac565b5083600019167fd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a75508485604051808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610dd5578082015181840152602081019050610dba565b50505050905090810190601f168015610e025780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b83811015610e3b578082015181840152602081019050610e20565b50505050905090810190601f168015610e685780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a250505050565b6000606060006001600086600019166000191681526020019081526020016000209050600192505b8383111515610fb357600084841614158015610eea57506000816006016000858152602001908152602001600020805460018160011615610100020316600290049050115b15610fa4578060060160008481526020019081526020016000208054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f985780601f10610f6d57610100808354040283529160200191610f98565b820191906000526020600020905b815481529060010190602001808311610f7b57829003601f168201915b50505050509150610fb8565b6001839060020a029250610ea5565b600092505b509250929050565b823373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166302571be3836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b15801561107057600080fd5b505af1158015611084573d6000803e3d6000fd5b505050506040513d602081101561109a57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff161415156110cd57600080fd5b6040805190810160405280846000191681526020018360001916815250600160008660001916600019168152602001908152602001600020600301600082015181600001906000191690556020820151816001019060001916905590505083600019167f1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46848460405180836000191660001916815260200182600019166000191681526020019250505060405180910390a250505050565b6000600160008360001916600019168152602001908152602001600020600101549050919050565b600060016000836000191660001916815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6060600160008460001916600019168152602001908152602001600020600501826040518082805190602001908083835b60208310151561124b5780518252602082019150602081019050602083039250611226565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113125780601f106112e757610100808354040283529160200191611312565b820191906000526020600020905b8154815290600101906020018083116112f557829003601f168201915b5050505050905092915050565b823373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166302571be3836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b1580156113cf57600080fd5b505af11580156113e3573d6000803e3d6000fd5b505050506040513d60208110156113f957600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1614151561142c57600080fd5b600083600185031614151561144057600080fd5b8160016000866000191660001916815260200190815260200160002060060160008581526020019081526020016000209080519060200190611483929190611e2c565b508284600019167faa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe360405160405180910390a350505050565b60606001600083600019166000191681526020019081526020016000206002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156115705780601f1061154557610100808354040283529160200191611570565b820191906000526020600020905b81548152906001019060200180831161155357829003601f168201915b50505050509050919050565b813373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166302571be3836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b15801561162c57600080fd5b505af1158015611640573d6000803e3d6000fd5b505050506040513d602081101561165657600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1614151561168957600080fd5b8160016000856000191660001916815260200190815260200160002060020190805190602001906116bb929190611dac565b5082600019167fb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7836040518080602001828103825283818151815260200191508051906020019080838360005b83811015611723578082015181840152602081019050611708565b50505050905090810190601f1680156117505780820380516001836020036101000a031916815260200191505b509250505060405180910390a2505050565b813373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166302571be3836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b15801561181257600080fd5b505af1158015611826573d6000803e3d6000fd5b505050506040513d602081101561183c57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1614151561186f57600080fd5b8160016000856000191660001916815260200190815260200160002060070190805190602001906118a1929190611e2c565b5082600019167fc0b0fc07269fc2749adada3221c095a1d2187b2d075b51c915857b520f3a5021836040518080602001828103825283818151815260200191508051906020019080838360005b838110156119095780820151818401526020810190506118ee565b50505050905090810190601f1680156119365780820380516001836020036101000a031916815260200191505b509250505060405180910390a2505050565b813373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166302571be3836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b1580156119f857600080fd5b505af1158015611a0c573d6000803e3d6000fd5b505050506040513d6020811015611a2257600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff16141515611a5557600080fd5b81600160008560001916600019168152602001908152602001600020600101816000191690555082600019167f0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc8360405180826000191660001916815260200191505060405180910390a2505050565b600080600160008460001916600019168152602001908152602001600020600301600001546001600085600019166000191681526020019081526020016000206003016001015491509150915091565b813373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166302571be3836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b158015611bc557600080fd5b505af1158015611bd9573d6000803e3d6000fd5b505050506040513d6020811015611bef57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff16141515611c2257600080fd5b8160016000856000191660001916815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600019167f52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd283604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a2505050565b60606001600083600019166000191681526020019081526020016000206007018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611da05780601f10611d7557610100808354040283529160200191611da0565b820191906000526020600020905b815481529060010190602001808311611d8357829003601f168201915b50505050509050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611ded57805160ff1916838001178555611e1b565b82800160010185558215611e1b579182015b82811115611e1a578251825591602001919060010190611dff565b5b509050611e289190611eac565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611e6d57805160ff1916838001178555611e9b565b82800160010185558215611e9b579182015b82811115611e9a578251825591602001919060010190611e7f565b5b509050611ea89190611eac565b5090565b611ece91905b80821115611eca576000816000905550600101611eb2565b5090565b905600a165627a7a7230582035eee055a6798b403c5f9f5be6cf2c48d5f813b0eb2aebe916470ae616d24be50029",
}

// PublicresolverABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicresolverMetaData.ABI instead.
var PublicresolverABI = PublicresolverMetaData.ABI

// PublicresolverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicresolverMetaData.Bin instead.
var PublicresolverBin = PublicresolverMetaData.Bin

// DeployPublicresolver deploys a new Ethereum contract, binding an instance of Publicresolver to it.
func DeployPublicresolver(auth *bind.TransactOpts, backend bind.ContractBackend, ensAddr common.Address) (common.Address, *types.Transaction, *Publicresolver, error) {
	parsed, err := PublicresolverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicresolverBin), backend, ensAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Publicresolver{PublicresolverCaller: PublicresolverCaller{contract: contract}, PublicresolverTransactor: PublicresolverTransactor{contract: contract}, PublicresolverFilterer: PublicresolverFilterer{contract: contract}}, nil
}

// Publicresolver is an auto generated Go binding around an Ethereum contract.
type Publicresolver struct {
	PublicresolverCaller     // Read-only binding to the contract
	PublicresolverTransactor // Write-only binding to the contract
	PublicresolverFilterer   // Log filterer for contract events
}

// PublicresolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicresolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicresolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicresolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicresolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicresolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicresolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicresolverSession struct {
	Contract     *Publicresolver   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PublicresolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicresolverCallerSession struct {
	Contract *PublicresolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PublicresolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicresolverTransactorSession struct {
	Contract     *PublicresolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PublicresolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicresolverRaw struct {
	Contract *Publicresolver // Generic contract binding to access the raw methods on
}

// PublicresolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicresolverCallerRaw struct {
	Contract *PublicresolverCaller // Generic read-only contract binding to access the raw methods on
}

// PublicresolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicresolverTransactorRaw struct {
	Contract *PublicresolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicresolver creates a new instance of Publicresolver, bound to a specific deployed contract.
func NewPublicresolver(address common.Address, backend bind.ContractBackend) (*Publicresolver, error) {
	contract, err := bindPublicresolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Publicresolver{PublicresolverCaller: PublicresolverCaller{contract: contract}, PublicresolverTransactor: PublicresolverTransactor{contract: contract}, PublicresolverFilterer: PublicresolverFilterer{contract: contract}}, nil
}

// NewPublicresolverCaller creates a new read-only instance of Publicresolver, bound to a specific deployed contract.
func NewPublicresolverCaller(address common.Address, caller bind.ContractCaller) (*PublicresolverCaller, error) {
	contract, err := bindPublicresolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicresolverCaller{contract: contract}, nil
}

// NewPublicresolverTransactor creates a new write-only instance of Publicresolver, bound to a specific deployed contract.
func NewPublicresolverTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicresolverTransactor, error) {
	contract, err := bindPublicresolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicresolverTransactor{contract: contract}, nil
}

// NewPublicresolverFilterer creates a new log filterer instance of Publicresolver, bound to a specific deployed contract.
func NewPublicresolverFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicresolverFilterer, error) {
	contract, err := bindPublicresolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicresolverFilterer{contract: contract}, nil
}

// bindPublicresolver binds a generic wrapper to an already deployed contract.
func bindPublicresolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicresolverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Publicresolver *PublicresolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Publicresolver.Contract.PublicresolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Publicresolver *PublicresolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Publicresolver.Contract.PublicresolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Publicresolver *PublicresolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Publicresolver.Contract.PublicresolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Publicresolver *PublicresolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Publicresolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Publicresolver *PublicresolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Publicresolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Publicresolver *PublicresolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Publicresolver.Contract.contract.Transact(opts, method, params...)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256 contentType, bytes data)
func (_Publicresolver *PublicresolverCaller) ABI(opts *bind.CallOpts, node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	var out []interface{}
	err := _Publicresolver.contract.Call(opts, &out, "ABI", node, contentTypes)

	outstruct := new(struct {
		ContentType *big.Int
		Data        []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ContentType = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256 contentType, bytes data)
func (_Publicresolver *PublicresolverSession) ABI(node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	return _Publicresolver.Contract.ABI(&_Publicresolver.CallOpts, node, contentTypes)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256 contentType, bytes data)
func (_Publicresolver *PublicresolverCallerSession) ABI(node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	return _Publicresolver.Contract.ABI(&_Publicresolver.CallOpts, node, contentTypes)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Publicresolver *PublicresolverCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Publicresolver.contract.Call(opts, &out, "addr", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Publicresolver *PublicresolverSession) Addr(node [32]byte) (common.Address, error) {
	return _Publicresolver.Contract.Addr(&_Publicresolver.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Publicresolver *PublicresolverCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _Publicresolver.Contract.Addr(&_Publicresolver.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(bytes32 node) view returns(bytes32)
func (_Publicresolver *PublicresolverCaller) Content(opts *bind.CallOpts, node [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Publicresolver.contract.Call(opts, &out, "content", node)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(bytes32 node) view returns(bytes32)
func (_Publicresolver *PublicresolverSession) Content(node [32]byte) ([32]byte, error) {
	return _Publicresolver.Contract.Content(&_Publicresolver.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(bytes32 node) view returns(bytes32)
func (_Publicresolver *PublicresolverCallerSession) Content(node [32]byte) ([32]byte, error) {
	return _Publicresolver.Contract.Content(&_Publicresolver.CallOpts, node)
}

// Multihash is a free data retrieval call binding the contract method 0xe89401a1.
//
// Solidity: function multihash(bytes32 node) view returns(bytes)
func (_Publicresolver *PublicresolverCaller) Multihash(opts *bind.CallOpts, node [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Publicresolver.contract.Call(opts, &out, "multihash", node)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Multihash is a free data retrieval call binding the contract method 0xe89401a1.
//
// Solidity: function multihash(bytes32 node) view returns(bytes)
func (_Publicresolver *PublicresolverSession) Multihash(node [32]byte) ([]byte, error) {
	return _Publicresolver.Contract.Multihash(&_Publicresolver.CallOpts, node)
}

// Multihash is a free data retrieval call binding the contract method 0xe89401a1.
//
// Solidity: function multihash(bytes32 node) view returns(bytes)
func (_Publicresolver *PublicresolverCallerSession) Multihash(node [32]byte) ([]byte, error) {
	return _Publicresolver.Contract.Multihash(&_Publicresolver.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_Publicresolver *PublicresolverCaller) Name(opts *bind.CallOpts, node [32]byte) (string, error) {
	var out []interface{}
	err := _Publicresolver.contract.Call(opts, &out, "name", node)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_Publicresolver *PublicresolverSession) Name(node [32]byte) (string, error) {
	return _Publicresolver.Contract.Name(&_Publicresolver.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_Publicresolver *PublicresolverCallerSession) Name(node [32]byte) (string, error) {
	return _Publicresolver.Contract.Name(&_Publicresolver.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_Publicresolver *PublicresolverCaller) Pubkey(opts *bind.CallOpts, node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	var out []interface{}
	err := _Publicresolver.contract.Call(opts, &out, "pubkey", node)

	outstruct := new(struct {
		X [32]byte
		Y [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Y = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_Publicresolver *PublicresolverSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _Publicresolver.Contract.Pubkey(&_Publicresolver.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_Publicresolver *PublicresolverCallerSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _Publicresolver.Contract.Pubkey(&_Publicresolver.CallOpts, node)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Publicresolver *PublicresolverCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _Publicresolver.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Publicresolver *PublicresolverSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Publicresolver.Contract.SupportsInterface(&_Publicresolver.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Publicresolver *PublicresolverCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Publicresolver.Contract.SupportsInterface(&_Publicresolver.CallOpts, interfaceID)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_Publicresolver *PublicresolverCaller) Text(opts *bind.CallOpts, node [32]byte, key string) (string, error) {
	var out []interface{}
	err := _Publicresolver.contract.Call(opts, &out, "text", node, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_Publicresolver *PublicresolverSession) Text(node [32]byte, key string) (string, error) {
	return _Publicresolver.Contract.Text(&_Publicresolver.CallOpts, node, key)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_Publicresolver *PublicresolverCallerSession) Text(node [32]byte, key string) (string, error) {
	return _Publicresolver.Contract.Text(&_Publicresolver.CallOpts, node, key)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_Publicresolver *PublicresolverTransactor) SetABI(opts *bind.TransactOpts, node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _Publicresolver.contract.Transact(opts, "setABI", node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_Publicresolver *PublicresolverSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _Publicresolver.Contract.SetABI(&_Publicresolver.TransactOpts, node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_Publicresolver *PublicresolverTransactorSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _Publicresolver.Contract.SetABI(&_Publicresolver.TransactOpts, node, contentType, data)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address addr) returns()
func (_Publicresolver *PublicresolverTransactor) SetAddr(opts *bind.TransactOpts, node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Publicresolver.contract.Transact(opts, "setAddr", node, addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address addr) returns()
func (_Publicresolver *PublicresolverSession) SetAddr(node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Publicresolver.Contract.SetAddr(&_Publicresolver.TransactOpts, node, addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address addr) returns()
func (_Publicresolver *PublicresolverTransactorSession) SetAddr(node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Publicresolver.Contract.SetAddr(&_Publicresolver.TransactOpts, node, addr)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(bytes32 node, bytes32 hash) returns()
func (_Publicresolver *PublicresolverTransactor) SetContent(opts *bind.TransactOpts, node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _Publicresolver.contract.Transact(opts, "setContent", node, hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(bytes32 node, bytes32 hash) returns()
func (_Publicresolver *PublicresolverSession) SetContent(node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _Publicresolver.Contract.SetContent(&_Publicresolver.TransactOpts, node, hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(bytes32 node, bytes32 hash) returns()
func (_Publicresolver *PublicresolverTransactorSession) SetContent(node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _Publicresolver.Contract.SetContent(&_Publicresolver.TransactOpts, node, hash)
}

// SetMultihash is a paid mutator transaction binding the contract method 0xaa4cb547.
//
// Solidity: function setMultihash(bytes32 node, bytes hash) returns()
func (_Publicresolver *PublicresolverTransactor) SetMultihash(opts *bind.TransactOpts, node [32]byte, hash []byte) (*types.Transaction, error) {
	return _Publicresolver.contract.Transact(opts, "setMultihash", node, hash)
}

// SetMultihash is a paid mutator transaction binding the contract method 0xaa4cb547.
//
// Solidity: function setMultihash(bytes32 node, bytes hash) returns()
func (_Publicresolver *PublicresolverSession) SetMultihash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _Publicresolver.Contract.SetMultihash(&_Publicresolver.TransactOpts, node, hash)
}

// SetMultihash is a paid mutator transaction binding the contract method 0xaa4cb547.
//
// Solidity: function setMultihash(bytes32 node, bytes hash) returns()
func (_Publicresolver *PublicresolverTransactorSession) SetMultihash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _Publicresolver.Contract.SetMultihash(&_Publicresolver.TransactOpts, node, hash)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string name) returns()
func (_Publicresolver *PublicresolverTransactor) SetName(opts *bind.TransactOpts, node [32]byte, name string) (*types.Transaction, error) {
	return _Publicresolver.contract.Transact(opts, "setName", node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string name) returns()
func (_Publicresolver *PublicresolverSession) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _Publicresolver.Contract.SetName(&_Publicresolver.TransactOpts, node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string name) returns()
func (_Publicresolver *PublicresolverTransactorSession) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _Publicresolver.Contract.SetName(&_Publicresolver.TransactOpts, node, name)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_Publicresolver *PublicresolverTransactor) SetPubkey(opts *bind.TransactOpts, node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _Publicresolver.contract.Transact(opts, "setPubkey", node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_Publicresolver *PublicresolverSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _Publicresolver.Contract.SetPubkey(&_Publicresolver.TransactOpts, node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_Publicresolver *PublicresolverTransactorSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _Publicresolver.Contract.SetPubkey(&_Publicresolver.TransactOpts, node, x, y)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_Publicresolver *PublicresolverTransactor) SetText(opts *bind.TransactOpts, node [32]byte, key string, value string) (*types.Transaction, error) {
	return _Publicresolver.contract.Transact(opts, "setText", node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_Publicresolver *PublicresolverSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _Publicresolver.Contract.SetText(&_Publicresolver.TransactOpts, node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_Publicresolver *PublicresolverTransactorSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _Publicresolver.Contract.SetText(&_Publicresolver.TransactOpts, node, key, value)
}

// PublicresolverABIChangedIterator is returned from FilterABIChanged and is used to iterate over the raw logs and unpacked data for ABIChanged events raised by the Publicresolver contract.
type PublicresolverABIChangedIterator struct {
	Event *PublicresolverABIChanged // Event containing the contract specifics and raw log

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
func (it *PublicresolverABIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicresolverABIChanged)
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
		it.Event = new(PublicresolverABIChanged)
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
func (it *PublicresolverABIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicresolverABIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicresolverABIChanged represents a ABIChanged event raised by the Publicresolver contract.
type PublicresolverABIChanged struct {
	Node        [32]byte
	ContentType *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterABIChanged is a free log retrieval operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_Publicresolver *PublicresolverFilterer) FilterABIChanged(opts *bind.FilterOpts, node [][32]byte, contentType []*big.Int) (*PublicresolverABIChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _Publicresolver.contract.FilterLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return &PublicresolverABIChangedIterator{contract: _Publicresolver.contract, event: "ABIChanged", logs: logs, sub: sub}, nil
}

// WatchABIChanged is a free log subscription operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_Publicresolver *PublicresolverFilterer) WatchABIChanged(opts *bind.WatchOpts, sink chan<- *PublicresolverABIChanged, node [][32]byte, contentType []*big.Int) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _Publicresolver.contract.WatchLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicresolverABIChanged)
				if err := _Publicresolver.contract.UnpackLog(event, "ABIChanged", log); err != nil {
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

// ParseABIChanged is a log parse operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_Publicresolver *PublicresolverFilterer) ParseABIChanged(log types.Log) (*PublicresolverABIChanged, error) {
	event := new(PublicresolverABIChanged)
	if err := _Publicresolver.contract.UnpackLog(event, "ABIChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicresolverAddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the Publicresolver contract.
type PublicresolverAddrChangedIterator struct {
	Event *PublicresolverAddrChanged // Event containing the contract specifics and raw log

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
func (it *PublicresolverAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicresolverAddrChanged)
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
		it.Event = new(PublicresolverAddrChanged)
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
func (it *PublicresolverAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicresolverAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicresolverAddrChanged represents a AddrChanged event raised by the Publicresolver contract.
type PublicresolverAddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_Publicresolver *PublicresolverFilterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*PublicresolverAddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Publicresolver.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &PublicresolverAddrChangedIterator{contract: _Publicresolver.contract, event: "AddrChanged", logs: logs, sub: sub}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_Publicresolver *PublicresolverFilterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *PublicresolverAddrChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Publicresolver.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicresolverAddrChanged)
				if err := _Publicresolver.contract.UnpackLog(event, "AddrChanged", log); err != nil {
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

// ParseAddrChanged is a log parse operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_Publicresolver *PublicresolverFilterer) ParseAddrChanged(log types.Log) (*PublicresolverAddrChanged, error) {
	event := new(PublicresolverAddrChanged)
	if err := _Publicresolver.contract.UnpackLog(event, "AddrChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicresolverContentChangedIterator is returned from FilterContentChanged and is used to iterate over the raw logs and unpacked data for ContentChanged events raised by the Publicresolver contract.
type PublicresolverContentChangedIterator struct {
	Event *PublicresolverContentChanged // Event containing the contract specifics and raw log

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
func (it *PublicresolverContentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicresolverContentChanged)
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
		it.Event = new(PublicresolverContentChanged)
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
func (it *PublicresolverContentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicresolverContentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicresolverContentChanged represents a ContentChanged event raised by the Publicresolver contract.
type PublicresolverContentChanged struct {
	Node [32]byte
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterContentChanged is a free log retrieval operation binding the contract event 0x0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc.
//
// Solidity: event ContentChanged(bytes32 indexed node, bytes32 hash)
func (_Publicresolver *PublicresolverFilterer) FilterContentChanged(opts *bind.FilterOpts, node [][32]byte) (*PublicresolverContentChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Publicresolver.contract.FilterLogs(opts, "ContentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &PublicresolverContentChangedIterator{contract: _Publicresolver.contract, event: "ContentChanged", logs: logs, sub: sub}, nil
}

// WatchContentChanged is a free log subscription operation binding the contract event 0x0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc.
//
// Solidity: event ContentChanged(bytes32 indexed node, bytes32 hash)
func (_Publicresolver *PublicresolverFilterer) WatchContentChanged(opts *bind.WatchOpts, sink chan<- *PublicresolverContentChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Publicresolver.contract.WatchLogs(opts, "ContentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicresolverContentChanged)
				if err := _Publicresolver.contract.UnpackLog(event, "ContentChanged", log); err != nil {
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

// ParseContentChanged is a log parse operation binding the contract event 0x0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc.
//
// Solidity: event ContentChanged(bytes32 indexed node, bytes32 hash)
func (_Publicresolver *PublicresolverFilterer) ParseContentChanged(log types.Log) (*PublicresolverContentChanged, error) {
	event := new(PublicresolverContentChanged)
	if err := _Publicresolver.contract.UnpackLog(event, "ContentChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicresolverMultihashChangedIterator is returned from FilterMultihashChanged and is used to iterate over the raw logs and unpacked data for MultihashChanged events raised by the Publicresolver contract.
type PublicresolverMultihashChangedIterator struct {
	Event *PublicresolverMultihashChanged // Event containing the contract specifics and raw log

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
func (it *PublicresolverMultihashChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicresolverMultihashChanged)
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
		it.Event = new(PublicresolverMultihashChanged)
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
func (it *PublicresolverMultihashChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicresolverMultihashChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicresolverMultihashChanged represents a MultihashChanged event raised by the Publicresolver contract.
type PublicresolverMultihashChanged struct {
	Node [32]byte
	Hash []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMultihashChanged is a free log retrieval operation binding the contract event 0xc0b0fc07269fc2749adada3221c095a1d2187b2d075b51c915857b520f3a5021.
//
// Solidity: event MultihashChanged(bytes32 indexed node, bytes hash)
func (_Publicresolver *PublicresolverFilterer) FilterMultihashChanged(opts *bind.FilterOpts, node [][32]byte) (*PublicresolverMultihashChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Publicresolver.contract.FilterLogs(opts, "MultihashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &PublicresolverMultihashChangedIterator{contract: _Publicresolver.contract, event: "MultihashChanged", logs: logs, sub: sub}, nil
}

// WatchMultihashChanged is a free log subscription operation binding the contract event 0xc0b0fc07269fc2749adada3221c095a1d2187b2d075b51c915857b520f3a5021.
//
// Solidity: event MultihashChanged(bytes32 indexed node, bytes hash)
func (_Publicresolver *PublicresolverFilterer) WatchMultihashChanged(opts *bind.WatchOpts, sink chan<- *PublicresolverMultihashChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Publicresolver.contract.WatchLogs(opts, "MultihashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicresolverMultihashChanged)
				if err := _Publicresolver.contract.UnpackLog(event, "MultihashChanged", log); err != nil {
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

// ParseMultihashChanged is a log parse operation binding the contract event 0xc0b0fc07269fc2749adada3221c095a1d2187b2d075b51c915857b520f3a5021.
//
// Solidity: event MultihashChanged(bytes32 indexed node, bytes hash)
func (_Publicresolver *PublicresolverFilterer) ParseMultihashChanged(log types.Log) (*PublicresolverMultihashChanged, error) {
	event := new(PublicresolverMultihashChanged)
	if err := _Publicresolver.contract.UnpackLog(event, "MultihashChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicresolverNameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the Publicresolver contract.
type PublicresolverNameChangedIterator struct {
	Event *PublicresolverNameChanged // Event containing the contract specifics and raw log

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
func (it *PublicresolverNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicresolverNameChanged)
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
		it.Event = new(PublicresolverNameChanged)
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
func (it *PublicresolverNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicresolverNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicresolverNameChanged represents a NameChanged event raised by the Publicresolver contract.
type PublicresolverNameChanged struct {
	Node [32]byte
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_Publicresolver *PublicresolverFilterer) FilterNameChanged(opts *bind.FilterOpts, node [][32]byte) (*PublicresolverNameChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Publicresolver.contract.FilterLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &PublicresolverNameChangedIterator{contract: _Publicresolver.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_Publicresolver *PublicresolverFilterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *PublicresolverNameChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Publicresolver.contract.WatchLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicresolverNameChanged)
				if err := _Publicresolver.contract.UnpackLog(event, "NameChanged", log); err != nil {
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

// ParseNameChanged is a log parse operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_Publicresolver *PublicresolverFilterer) ParseNameChanged(log types.Log) (*PublicresolverNameChanged, error) {
	event := new(PublicresolverNameChanged)
	if err := _Publicresolver.contract.UnpackLog(event, "NameChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicresolverPubkeyChangedIterator is returned from FilterPubkeyChanged and is used to iterate over the raw logs and unpacked data for PubkeyChanged events raised by the Publicresolver contract.
type PublicresolverPubkeyChangedIterator struct {
	Event *PublicresolverPubkeyChanged // Event containing the contract specifics and raw log

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
func (it *PublicresolverPubkeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicresolverPubkeyChanged)
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
		it.Event = new(PublicresolverPubkeyChanged)
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
func (it *PublicresolverPubkeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicresolverPubkeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicresolverPubkeyChanged represents a PubkeyChanged event raised by the Publicresolver contract.
type PublicresolverPubkeyChanged struct {
	Node [32]byte
	X    [32]byte
	Y    [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPubkeyChanged is a free log retrieval operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_Publicresolver *PublicresolverFilterer) FilterPubkeyChanged(opts *bind.FilterOpts, node [][32]byte) (*PublicresolverPubkeyChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Publicresolver.contract.FilterLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &PublicresolverPubkeyChangedIterator{contract: _Publicresolver.contract, event: "PubkeyChanged", logs: logs, sub: sub}, nil
}

// WatchPubkeyChanged is a free log subscription operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_Publicresolver *PublicresolverFilterer) WatchPubkeyChanged(opts *bind.WatchOpts, sink chan<- *PublicresolverPubkeyChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Publicresolver.contract.WatchLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicresolverPubkeyChanged)
				if err := _Publicresolver.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
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

// ParsePubkeyChanged is a log parse operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_Publicresolver *PublicresolverFilterer) ParsePubkeyChanged(log types.Log) (*PublicresolverPubkeyChanged, error) {
	event := new(PublicresolverPubkeyChanged)
	if err := _Publicresolver.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicresolverTextChangedIterator is returned from FilterTextChanged and is used to iterate over the raw logs and unpacked data for TextChanged events raised by the Publicresolver contract.
type PublicresolverTextChangedIterator struct {
	Event *PublicresolverTextChanged // Event containing the contract specifics and raw log

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
func (it *PublicresolverTextChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicresolverTextChanged)
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
		it.Event = new(PublicresolverTextChanged)
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
func (it *PublicresolverTextChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicresolverTextChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicresolverTextChanged represents a TextChanged event raised by the Publicresolver contract.
type PublicresolverTextChanged struct {
	Node       [32]byte
	IndexedKey string
	Key        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTextChanged is a free log retrieval operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexedKey, string key)
func (_Publicresolver *PublicresolverFilterer) FilterTextChanged(opts *bind.FilterOpts, node [][32]byte) (*PublicresolverTextChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Publicresolver.contract.FilterLogs(opts, "TextChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &PublicresolverTextChangedIterator{contract: _Publicresolver.contract, event: "TextChanged", logs: logs, sub: sub}, nil
}

// WatchTextChanged is a free log subscription operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexedKey, string key)
func (_Publicresolver *PublicresolverFilterer) WatchTextChanged(opts *bind.WatchOpts, sink chan<- *PublicresolverTextChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Publicresolver.contract.WatchLogs(opts, "TextChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicresolverTextChanged)
				if err := _Publicresolver.contract.UnpackLog(event, "TextChanged", log); err != nil {
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

// ParseTextChanged is a log parse operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexedKey, string key)
func (_Publicresolver *PublicresolverFilterer) ParseTextChanged(log types.Log) (*PublicresolverTextChanged, error) {
	event := new(PublicresolverTextChanged)
	if err := _Publicresolver.contract.UnpackLog(event, "TextChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
