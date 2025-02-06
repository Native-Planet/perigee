// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package delegatedsending

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

// DelegatedsendingMetaData contains all meta data concerning the Delegatedsending contract.
var DelegatedsendingMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"getInviters\",\"outputs\":[{\"name\":\"invs\",\"type\":\"uint32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"isInviter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"poolStarsRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"invitedBy\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"pools\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_as\",\"type\":\"uint32\"},{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"sendPoint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"fromPool\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_as\",\"type\":\"uint32\"},{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"canSend\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inviters\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolStars\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"canReceive\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"invited\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"getPool\",\"outputs\":[{\"name\":\"pool\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"uint32\"}],\"name\":\"getPoolStars\",\"outputs\":[{\"name\":\"stars\",\"type\":\"uint16[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"azimuth\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"uint32\"}],\"name\":\"getInvited\",\"outputs\":[{\"name\":\"invd\",\"type\":\"uint32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_as\",\"type\":\"uint16\"},{\"name\":\"_for\",\"type\":\"uint32\"},{\"name\":\"_size\",\"type\":\"uint16\"}],\"name\":\"setPoolSize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_azimuth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"prefix\",\"type\":\"uint16\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"size\",\"type\":\"uint16\"}],\"name\":\"Pool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"prefix\",\"type\":\"uint16\"},{\"indexed\":true,\"name\":\"fromPool\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"by\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"point\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Sent\",\"type\":\"event\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051602080611fab8339810180604052810190808051906020019092919050505080806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050611f26806100856000396000f3006080604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631a6a57f4146100f65780632cb2b1fa1461016257806340603cc4146101ad5780634d6b4499146102065780634f4573ee1461025957806366008294146102b65780636dcadcdd1461031957806370646de91461036c5780637a84c366146103c75780637d6e36d31461041457806390d370ba1461046d5780639677c7eb146104c85780639da46ee314610525578063b7fc5a4814610578578063d40ffacb14610600578063e00fb7b714610657578063e78b8a37146106df575b600080fd5b34801561010257600080fd5b5061010b61072e565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561014e578082015181840152602081019050610133565b505050509050019250505060405180910390f35b34801561016e57600080fd5b50610193600480360381019080803563ffffffff1690602001909291905050506107b2565b604051808215151515815260200191505060405180910390f35b3480156101b957600080fd5b506101ec600480360381019080803563ffffffff169060200190929190803561ffff1690602001909291905050506107d2565b604051808215151515815260200191505060405180910390f35b34801561021257600080fd5b50610237600480360381019080803563ffffffff169060200190929190505050610801565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b34801561026557600080fd5b50610298600480360381019080803563ffffffff169060200190929190803561ffff169060200190929190505050610824565b604051808261ffff1661ffff16815260200191505060405180910390f35b3480156102c257600080fd5b50610317600480360381019080803563ffffffff169060200190929190803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610854565b005b34801561032557600080fd5b5061034a600480360381019080803563ffffffff169060200190929190505050610f10565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b34801561037857600080fd5b506103ad600480360381019080803563ffffffff169060200190929190803563ffffffff169060200190929190505050610f33565b604051808215151515815260200191505060405180910390f35b3480156103d357600080fd5b506103f2600480360381019080803590602001909291905050506115c5565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b34801561042057600080fd5b5061044f600480360381019080803563ffffffff169060200190929190803590602001909291905050506115fe565b604051808261ffff1661ffff16815260200191505060405180910390f35b34801561047957600080fd5b506104ae600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611644565b604051808215151515815260200191505060405180910390f35b3480156104d457600080fd5b50610503600480360381019080803563ffffffff16906020019092919080359060200190929190505050611849565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b34801561053157600080fd5b50610556600480360381019080803563ffffffff169060200190929190505050611891565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b34801561058457600080fd5b506105a9600480360381019080803563ffffffff1690602001909291905050506118e5565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156105ec5780820151818401526020810190506105d1565b505050509050019250505060405180910390f35b34801561060c57600080fd5b50610615611984565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561066357600080fd5b50610688600480360381019080803563ffffffff1690602001909291905050506119a9565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156106cb5780820151818401526020810190506106b0565b505050509050019250505060405180910390f35b3480156106eb57600080fd5b5061072c600480360381019080803561ffff169060200190929190803563ffffffff169060200190929190803561ffff169060200190929190505050611a4c565b005b606060058054806020026020016040519081016040528092919081815260200182805480156107a857602002820191906000526020600020906000905b82829054906101000a900463ffffffff1663ffffffff168152602001906004019060208260030104928301926001038202915080841161076b5790505b5050505050905090565b60066020528060005260406000206000915054906101000a900460ff1681565b60046020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b60086020528060005260406000206000915054906101000a900463ffffffff1681565b60016020528160005260406000206020528060005260406000206000915091509054906101000a900461ffff1681565b600080846000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663caf590f982336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561092857600080fd5b505af115801561093c573d6000803e3d6000fd5b505050506040513d602081101561095257600080fd5b81019080805190602001909291905050508015610a4257506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015610a0657600080fd5b505af1158015610a1a573d6000803e3d6000fd5b505050506040513d6020811015610a3057600080fd5b81019080805190602001909291905050505b1515610a4d57600080fd5b610a578686610f33565b1515610a6257600080fd5b8373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151515610a9d57600080fd5b610aa684611644565b1515610ab157600080fd5b610aba86611891565b92506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e4a358d7866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015610b5857600080fd5b505af1158015610b6c573d6000803e3d6000fd5b505050506040513d6020811015610b8257600080fd5b81019080805190602001909291905050509150600160008463ffffffff1663ffffffff16815260200190815260200160002060008361ffff1661ffff168152602001908152602001600020600081819054906101000a900461ffff16809291906001900391906101000a81548161ffff021916908361ffff1602179055505082600260008763ffffffff1663ffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff160217905550600760008763ffffffff1663ffffffff16815260200190815260200160002085908060018154018082558091505090600182039060005260206000209060089182820401919006600402909192909190916101000a81548163ffffffff021916908363ffffffff1602179055505085600860008763ffffffff1663ffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610d7057600080fd5b505af1158015610d84573d6000803e3d6000fd5b505050506040513d6020811015610d9a57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1663a0d3253f86866040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b158015610e5957600080fd5b505af1158015610e6d573d6000803e3d6000fd5b505050508263ffffffff168261ffff167f47638e3cddee220481e4c3f9183d639c0efea7f05fcd2df4188855729f715419888888604051808463ffffffff1663ffffffff1681526020018363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001935050505060405180910390a3505050505050565b60026020528060005260406000206000915054906101000a900463ffffffff1681565b60008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e4a358d7856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015610fd457600080fd5b505af1158015610fe8573d6000803e3d6000fd5b505050506040513d6020811015610ffe57600080fd5b8101908080519060200190929190505050915061101a85611891565b9050600160008263ffffffff1663ffffffff16815260200190815260200160002060008361ffff1661ffff16815260200190815260200160002060009054906101000a900461ffff1661ffff16600010801561116857506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663caf590f98560006040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561112c57600080fd5b505af1158015611140573d6000803e3d6000fd5b505050506040513d602081101561115657600080fd5b81019080805190602001909291905050505b801561127957506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324ba1a4683306040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808361ffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561123d57600080fd5b505af1158015611251573d6000803e3d6000fd5b505050506040513d602081101561126757600080fd5b81019080805190602001909291905050505b801561135657506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636d09887b836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808261ffff1663ffffffff168152602001915050602060405180830381600087803b15801561131a57600080fd5b505af115801561132e573d6000803e3d6000fd5b505050506040513d602081101561134457600080fd5b81019080805190602001909291905050505b80156115bb57506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156113e257600080fd5b505af11580156113f6573d6000803e3d6000fd5b505050506040513d602081101561140c57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1663ef20bff883426040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808361ffff1663ffffffff16815260200182815260200192505050602060405180830381600087803b15801561149d57600080fd5b505af11580156114b1573d6000803e3d6000fd5b505050506040513d60208110156114c757600080fd5b810190808051906020019092919050505063ffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663293a9169846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808261ffff1663ffffffff168152602001915050602060405180830381600087803b15801561157857600080fd5b505af115801561158c573d6000803e3d6000fd5b505050506040513d60208110156115a257600080fd5b810190808051906020019092919050505063ffffffff16105b9250505092915050565b6005818154811015156115d457fe5b9060005260206000209060089182820401919006600402915054906101000a900463ffffffff1681565b60036020528160005260406000208181548110151561161957fe5b9060005260206000209060109182820401919006600202915091509054906101000a900461ffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166312afbc78836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561170257600080fd5b505af1158015611716573d6000803e3d6000fd5b505050506040513d602081101561172c57600080fd5b8101908080519060200190929190505050600014801561184257506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e01cff84836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561180357600080fd5b505af1158015611817573d6000803e3d6000fd5b505050506040513d602081101561182d57600080fd5b81019080805190602001909291905050506000145b9050919050565b60076020528160005260406000208181548110151561186457fe5b9060005260206000209060089182820401919006600402915091509054906101000a900463ffffffff1681565b6000600260008363ffffffff1663ffffffff16815260200190815260200160002060009054906101000a900463ffffffff1690508063ffffffff16600014156118dc578190506118e0565b8090505b919050565b6060600360008363ffffffff1663ffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561197857602002820191906000526020600020906000905b82829054906101000a900461ffff1661ffff168152602001906002019060208260010104928301926001038202915080841161193f5790505b50505050509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060600760008363ffffffff1663ffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015611a4057602002820191906000526020600020906000905b82829054906101000a900463ffffffff1663ffffffff1681526020019060040190602082600301049283019260010382029150808411611a035790505b50505050509050919050565b8261ffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663caf590f982336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015611b2157600080fd5b505af1158015611b35573d6000803e3d6000fd5b505050506040513d6020811015611b4b57600080fd5b81019080805190602001909291905050508015611c3b57506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015611bff57600080fd5b505af1158015611c13573d6000803e3d6000fd5b505050506040513d6020811015611c2957600080fd5b81019080805190602001909291905050505b1515611c4657600080fd5b6000600260008563ffffffff1663ffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff16021790555081600160008563ffffffff1663ffffffff16815260200190815260200160002060008661ffff1661ffff16815260200190815260200160002060006101000a81548161ffff021916908361ffff160217905550600460008463ffffffff1663ffffffff16815260200190815260200160002060008561ffff1661ffff16815260200190815260200160002060009054906101000a900460ff161515600015151415611de657600360008463ffffffff1663ffffffff16815260200190815260200160002084908060018154018082558091505090600182039060005260206000209060109182820401919006600202909192909190916101000a81548161ffff021916908361ffff160217905550506001600460008563ffffffff1663ffffffff16815260200190815260200160002060008661ffff1661ffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b600660008463ffffffff1663ffffffff16815260200190815260200160002060009054906101000a900460ff161515600015151415611ea9576001600660008563ffffffff1663ffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600583908060018154018082558091505090600182039060005260206000209060089182820401919006600402909192909190916101000a81548163ffffffff021916908363ffffffff160217905550505b8263ffffffff168461ffff167fb05d354b3ecb8f9593b9298bcdeea36401f0e199bc0617e9a731a45cecb343ec84604051808261ffff1661ffff16815260200191505060405180910390a3505050505600a165627a7a72305820fcfbc064562659fc1048f5aba22e26a3f8ebe0ed301d5b5b5526b0af440167750029",
}

// DelegatedsendingABI is the input ABI used to generate the binding from.
// Deprecated: Use DelegatedsendingMetaData.ABI instead.
var DelegatedsendingABI = DelegatedsendingMetaData.ABI

// DelegatedsendingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DelegatedsendingMetaData.Bin instead.
var DelegatedsendingBin = DelegatedsendingMetaData.Bin

// DeployDelegatedsending deploys a new Ethereum contract, binding an instance of Delegatedsending to it.
func DeployDelegatedsending(auth *bind.TransactOpts, backend bind.ContractBackend, _azimuth common.Address) (common.Address, *types.Transaction, *Delegatedsending, error) {
	parsed, err := DelegatedsendingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DelegatedsendingBin), backend, _azimuth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Delegatedsending{DelegatedsendingCaller: DelegatedsendingCaller{contract: contract}, DelegatedsendingTransactor: DelegatedsendingTransactor{contract: contract}, DelegatedsendingFilterer: DelegatedsendingFilterer{contract: contract}}, nil
}

// Delegatedsending is an auto generated Go binding around an Ethereum contract.
type Delegatedsending struct {
	DelegatedsendingCaller     // Read-only binding to the contract
	DelegatedsendingTransactor // Write-only binding to the contract
	DelegatedsendingFilterer   // Log filterer for contract events
}

// DelegatedsendingCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelegatedsendingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatedsendingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegatedsendingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatedsendingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelegatedsendingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatedsendingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegatedsendingSession struct {
	Contract     *Delegatedsending // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelegatedsendingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegatedsendingCallerSession struct {
	Contract *DelegatedsendingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DelegatedsendingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegatedsendingTransactorSession struct {
	Contract     *DelegatedsendingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DelegatedsendingRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelegatedsendingRaw struct {
	Contract *Delegatedsending // Generic contract binding to access the raw methods on
}

// DelegatedsendingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegatedsendingCallerRaw struct {
	Contract *DelegatedsendingCaller // Generic read-only contract binding to access the raw methods on
}

// DelegatedsendingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegatedsendingTransactorRaw struct {
	Contract *DelegatedsendingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegatedsending creates a new instance of Delegatedsending, bound to a specific deployed contract.
func NewDelegatedsending(address common.Address, backend bind.ContractBackend) (*Delegatedsending, error) {
	contract, err := bindDelegatedsending(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Delegatedsending{DelegatedsendingCaller: DelegatedsendingCaller{contract: contract}, DelegatedsendingTransactor: DelegatedsendingTransactor{contract: contract}, DelegatedsendingFilterer: DelegatedsendingFilterer{contract: contract}}, nil
}

// NewDelegatedsendingCaller creates a new read-only instance of Delegatedsending, bound to a specific deployed contract.
func NewDelegatedsendingCaller(address common.Address, caller bind.ContractCaller) (*DelegatedsendingCaller, error) {
	contract, err := bindDelegatedsending(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelegatedsendingCaller{contract: contract}, nil
}

// NewDelegatedsendingTransactor creates a new write-only instance of Delegatedsending, bound to a specific deployed contract.
func NewDelegatedsendingTransactor(address common.Address, transactor bind.ContractTransactor) (*DelegatedsendingTransactor, error) {
	contract, err := bindDelegatedsending(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelegatedsendingTransactor{contract: contract}, nil
}

// NewDelegatedsendingFilterer creates a new log filterer instance of Delegatedsending, bound to a specific deployed contract.
func NewDelegatedsendingFilterer(address common.Address, filterer bind.ContractFilterer) (*DelegatedsendingFilterer, error) {
	contract, err := bindDelegatedsending(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelegatedsendingFilterer{contract: contract}, nil
}

// bindDelegatedsending binds a generic wrapper to an already deployed contract.
func bindDelegatedsending(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DelegatedsendingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delegatedsending *DelegatedsendingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Delegatedsending.Contract.DelegatedsendingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delegatedsending *DelegatedsendingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegatedsending.Contract.DelegatedsendingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delegatedsending *DelegatedsendingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delegatedsending.Contract.DelegatedsendingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delegatedsending *DelegatedsendingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Delegatedsending.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delegatedsending *DelegatedsendingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegatedsending.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delegatedsending *DelegatedsendingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delegatedsending.Contract.contract.Transact(opts, method, params...)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Delegatedsending *DelegatedsendingCaller) Azimuth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Delegatedsending.contract.Call(opts, &out, "azimuth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Delegatedsending *DelegatedsendingSession) Azimuth() (common.Address, error) {
	return _Delegatedsending.Contract.Azimuth(&_Delegatedsending.CallOpts)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Delegatedsending *DelegatedsendingCallerSession) Azimuth() (common.Address, error) {
	return _Delegatedsending.Contract.Azimuth(&_Delegatedsending.CallOpts)
}

// CanReceive is a free data retrieval call binding the contract method 0x90d370ba.
//
// Solidity: function canReceive(address _recipient) view returns(bool result)
func (_Delegatedsending *DelegatedsendingCaller) CanReceive(opts *bind.CallOpts, _recipient common.Address) (bool, error) {
	var out []interface{}
	err := _Delegatedsending.contract.Call(opts, &out, "canReceive", _recipient)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanReceive is a free data retrieval call binding the contract method 0x90d370ba.
//
// Solidity: function canReceive(address _recipient) view returns(bool result)
func (_Delegatedsending *DelegatedsendingSession) CanReceive(_recipient common.Address) (bool, error) {
	return _Delegatedsending.Contract.CanReceive(&_Delegatedsending.CallOpts, _recipient)
}

// CanReceive is a free data retrieval call binding the contract method 0x90d370ba.
//
// Solidity: function canReceive(address _recipient) view returns(bool result)
func (_Delegatedsending *DelegatedsendingCallerSession) CanReceive(_recipient common.Address) (bool, error) {
	return _Delegatedsending.Contract.CanReceive(&_Delegatedsending.CallOpts, _recipient)
}

// CanSend is a free data retrieval call binding the contract method 0x70646de9.
//
// Solidity: function canSend(uint32 _as, uint32 _point) view returns(bool result)
func (_Delegatedsending *DelegatedsendingCaller) CanSend(opts *bind.CallOpts, _as uint32, _point uint32) (bool, error) {
	var out []interface{}
	err := _Delegatedsending.contract.Call(opts, &out, "canSend", _as, _point)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanSend is a free data retrieval call binding the contract method 0x70646de9.
//
// Solidity: function canSend(uint32 _as, uint32 _point) view returns(bool result)
func (_Delegatedsending *DelegatedsendingSession) CanSend(_as uint32, _point uint32) (bool, error) {
	return _Delegatedsending.Contract.CanSend(&_Delegatedsending.CallOpts, _as, _point)
}

// CanSend is a free data retrieval call binding the contract method 0x70646de9.
//
// Solidity: function canSend(uint32 _as, uint32 _point) view returns(bool result)
func (_Delegatedsending *DelegatedsendingCallerSession) CanSend(_as uint32, _point uint32) (bool, error) {
	return _Delegatedsending.Contract.CanSend(&_Delegatedsending.CallOpts, _as, _point)
}

// FromPool is a free data retrieval call binding the contract method 0x6dcadcdd.
//
// Solidity: function fromPool(uint32 ) view returns(uint32)
func (_Delegatedsending *DelegatedsendingCaller) FromPool(opts *bind.CallOpts, arg0 uint32) (uint32, error) {
	var out []interface{}
	err := _Delegatedsending.contract.Call(opts, &out, "fromPool", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FromPool is a free data retrieval call binding the contract method 0x6dcadcdd.
//
// Solidity: function fromPool(uint32 ) view returns(uint32)
func (_Delegatedsending *DelegatedsendingSession) FromPool(arg0 uint32) (uint32, error) {
	return _Delegatedsending.Contract.FromPool(&_Delegatedsending.CallOpts, arg0)
}

// FromPool is a free data retrieval call binding the contract method 0x6dcadcdd.
//
// Solidity: function fromPool(uint32 ) view returns(uint32)
func (_Delegatedsending *DelegatedsendingCallerSession) FromPool(arg0 uint32) (uint32, error) {
	return _Delegatedsending.Contract.FromPool(&_Delegatedsending.CallOpts, arg0)
}

// GetInvited is a free data retrieval call binding the contract method 0xe00fb7b7.
//
// Solidity: function getInvited(uint32 _who) view returns(uint32[] invd)
func (_Delegatedsending *DelegatedsendingCaller) GetInvited(opts *bind.CallOpts, _who uint32) ([]uint32, error) {
	var out []interface{}
	err := _Delegatedsending.contract.Call(opts, &out, "getInvited", _who)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetInvited is a free data retrieval call binding the contract method 0xe00fb7b7.
//
// Solidity: function getInvited(uint32 _who) view returns(uint32[] invd)
func (_Delegatedsending *DelegatedsendingSession) GetInvited(_who uint32) ([]uint32, error) {
	return _Delegatedsending.Contract.GetInvited(&_Delegatedsending.CallOpts, _who)
}

// GetInvited is a free data retrieval call binding the contract method 0xe00fb7b7.
//
// Solidity: function getInvited(uint32 _who) view returns(uint32[] invd)
func (_Delegatedsending *DelegatedsendingCallerSession) GetInvited(_who uint32) ([]uint32, error) {
	return _Delegatedsending.Contract.GetInvited(&_Delegatedsending.CallOpts, _who)
}

// GetInviters is a free data retrieval call binding the contract method 0x1a6a57f4.
//
// Solidity: function getInviters() view returns(uint32[] invs)
func (_Delegatedsending *DelegatedsendingCaller) GetInviters(opts *bind.CallOpts) ([]uint32, error) {
	var out []interface{}
	err := _Delegatedsending.contract.Call(opts, &out, "getInviters")

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetInviters is a free data retrieval call binding the contract method 0x1a6a57f4.
//
// Solidity: function getInviters() view returns(uint32[] invs)
func (_Delegatedsending *DelegatedsendingSession) GetInviters() ([]uint32, error) {
	return _Delegatedsending.Contract.GetInviters(&_Delegatedsending.CallOpts)
}

// GetInviters is a free data retrieval call binding the contract method 0x1a6a57f4.
//
// Solidity: function getInviters() view returns(uint32[] invs)
func (_Delegatedsending *DelegatedsendingCallerSession) GetInviters() ([]uint32, error) {
	return _Delegatedsending.Contract.GetInviters(&_Delegatedsending.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0x9da46ee3.
//
// Solidity: function getPool(uint32 _point) view returns(uint32 pool)
func (_Delegatedsending *DelegatedsendingCaller) GetPool(opts *bind.CallOpts, _point uint32) (uint32, error) {
	var out []interface{}
	err := _Delegatedsending.contract.Call(opts, &out, "getPool", _point)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x9da46ee3.
//
// Solidity: function getPool(uint32 _point) view returns(uint32 pool)
func (_Delegatedsending *DelegatedsendingSession) GetPool(_point uint32) (uint32, error) {
	return _Delegatedsending.Contract.GetPool(&_Delegatedsending.CallOpts, _point)
}

// GetPool is a free data retrieval call binding the contract method 0x9da46ee3.
//
// Solidity: function getPool(uint32 _point) view returns(uint32 pool)
func (_Delegatedsending *DelegatedsendingCallerSession) GetPool(_point uint32) (uint32, error) {
	return _Delegatedsending.Contract.GetPool(&_Delegatedsending.CallOpts, _point)
}

// GetPoolStars is a free data retrieval call binding the contract method 0xb7fc5a48.
//
// Solidity: function getPoolStars(uint32 _who) view returns(uint16[] stars)
func (_Delegatedsending *DelegatedsendingCaller) GetPoolStars(opts *bind.CallOpts, _who uint32) ([]uint16, error) {
	var out []interface{}
	err := _Delegatedsending.contract.Call(opts, &out, "getPoolStars", _who)

	if err != nil {
		return *new([]uint16), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint16)).(*[]uint16)

	return out0, err

}

// GetPoolStars is a free data retrieval call binding the contract method 0xb7fc5a48.
//
// Solidity: function getPoolStars(uint32 _who) view returns(uint16[] stars)
func (_Delegatedsending *DelegatedsendingSession) GetPoolStars(_who uint32) ([]uint16, error) {
	return _Delegatedsending.Contract.GetPoolStars(&_Delegatedsending.CallOpts, _who)
}

// GetPoolStars is a free data retrieval call binding the contract method 0xb7fc5a48.
//
// Solidity: function getPoolStars(uint32 _who) view returns(uint16[] stars)
func (_Delegatedsending *DelegatedsendingCallerSession) GetPoolStars(_who uint32) ([]uint16, error) {
	return _Delegatedsending.Contract.GetPoolStars(&_Delegatedsending.CallOpts, _who)
}

// Invited is a free data retrieval call binding the contract method 0x9677c7eb.
//
// Solidity: function invited(uint32 , uint256 ) view returns(uint32)
func (_Delegatedsending *DelegatedsendingCaller) Invited(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Delegatedsending.contract.Call(opts, &out, "invited", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Invited is a free data retrieval call binding the contract method 0x9677c7eb.
//
// Solidity: function invited(uint32 , uint256 ) view returns(uint32)
func (_Delegatedsending *DelegatedsendingSession) Invited(arg0 uint32, arg1 *big.Int) (uint32, error) {
	return _Delegatedsending.Contract.Invited(&_Delegatedsending.CallOpts, arg0, arg1)
}

// Invited is a free data retrieval call binding the contract method 0x9677c7eb.
//
// Solidity: function invited(uint32 , uint256 ) view returns(uint32)
func (_Delegatedsending *DelegatedsendingCallerSession) Invited(arg0 uint32, arg1 *big.Int) (uint32, error) {
	return _Delegatedsending.Contract.Invited(&_Delegatedsending.CallOpts, arg0, arg1)
}

// InvitedBy is a free data retrieval call binding the contract method 0x4d6b4499.
//
// Solidity: function invitedBy(uint32 ) view returns(uint32)
func (_Delegatedsending *DelegatedsendingCaller) InvitedBy(opts *bind.CallOpts, arg0 uint32) (uint32, error) {
	var out []interface{}
	err := _Delegatedsending.contract.Call(opts, &out, "invitedBy", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// InvitedBy is a free data retrieval call binding the contract method 0x4d6b4499.
//
// Solidity: function invitedBy(uint32 ) view returns(uint32)
func (_Delegatedsending *DelegatedsendingSession) InvitedBy(arg0 uint32) (uint32, error) {
	return _Delegatedsending.Contract.InvitedBy(&_Delegatedsending.CallOpts, arg0)
}

// InvitedBy is a free data retrieval call binding the contract method 0x4d6b4499.
//
// Solidity: function invitedBy(uint32 ) view returns(uint32)
func (_Delegatedsending *DelegatedsendingCallerSession) InvitedBy(arg0 uint32) (uint32, error) {
	return _Delegatedsending.Contract.InvitedBy(&_Delegatedsending.CallOpts, arg0)
}

// Inviters is a free data retrieval call binding the contract method 0x7a84c366.
//
// Solidity: function inviters(uint256 ) view returns(uint32)
func (_Delegatedsending *DelegatedsendingCaller) Inviters(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Delegatedsending.contract.Call(opts, &out, "inviters", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Inviters is a free data retrieval call binding the contract method 0x7a84c366.
//
// Solidity: function inviters(uint256 ) view returns(uint32)
func (_Delegatedsending *DelegatedsendingSession) Inviters(arg0 *big.Int) (uint32, error) {
	return _Delegatedsending.Contract.Inviters(&_Delegatedsending.CallOpts, arg0)
}

// Inviters is a free data retrieval call binding the contract method 0x7a84c366.
//
// Solidity: function inviters(uint256 ) view returns(uint32)
func (_Delegatedsending *DelegatedsendingCallerSession) Inviters(arg0 *big.Int) (uint32, error) {
	return _Delegatedsending.Contract.Inviters(&_Delegatedsending.CallOpts, arg0)
}

// IsInviter is a free data retrieval call binding the contract method 0x2cb2b1fa.
//
// Solidity: function isInviter(uint32 ) view returns(bool)
func (_Delegatedsending *DelegatedsendingCaller) IsInviter(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _Delegatedsending.contract.Call(opts, &out, "isInviter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInviter is a free data retrieval call binding the contract method 0x2cb2b1fa.
//
// Solidity: function isInviter(uint32 ) view returns(bool)
func (_Delegatedsending *DelegatedsendingSession) IsInviter(arg0 uint32) (bool, error) {
	return _Delegatedsending.Contract.IsInviter(&_Delegatedsending.CallOpts, arg0)
}

// IsInviter is a free data retrieval call binding the contract method 0x2cb2b1fa.
//
// Solidity: function isInviter(uint32 ) view returns(bool)
func (_Delegatedsending *DelegatedsendingCallerSession) IsInviter(arg0 uint32) (bool, error) {
	return _Delegatedsending.Contract.IsInviter(&_Delegatedsending.CallOpts, arg0)
}

// PoolStars is a free data retrieval call binding the contract method 0x7d6e36d3.
//
// Solidity: function poolStars(uint32 , uint256 ) view returns(uint16)
func (_Delegatedsending *DelegatedsendingCaller) PoolStars(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (uint16, error) {
	var out []interface{}
	err := _Delegatedsending.contract.Call(opts, &out, "poolStars", arg0, arg1)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PoolStars is a free data retrieval call binding the contract method 0x7d6e36d3.
//
// Solidity: function poolStars(uint32 , uint256 ) view returns(uint16)
func (_Delegatedsending *DelegatedsendingSession) PoolStars(arg0 uint32, arg1 *big.Int) (uint16, error) {
	return _Delegatedsending.Contract.PoolStars(&_Delegatedsending.CallOpts, arg0, arg1)
}

// PoolStars is a free data retrieval call binding the contract method 0x7d6e36d3.
//
// Solidity: function poolStars(uint32 , uint256 ) view returns(uint16)
func (_Delegatedsending *DelegatedsendingCallerSession) PoolStars(arg0 uint32, arg1 *big.Int) (uint16, error) {
	return _Delegatedsending.Contract.PoolStars(&_Delegatedsending.CallOpts, arg0, arg1)
}

// PoolStarsRegistered is a free data retrieval call binding the contract method 0x40603cc4.
//
// Solidity: function poolStarsRegistered(uint32 , uint16 ) view returns(bool)
func (_Delegatedsending *DelegatedsendingCaller) PoolStarsRegistered(opts *bind.CallOpts, arg0 uint32, arg1 uint16) (bool, error) {
	var out []interface{}
	err := _Delegatedsending.contract.Call(opts, &out, "poolStarsRegistered", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PoolStarsRegistered is a free data retrieval call binding the contract method 0x40603cc4.
//
// Solidity: function poolStarsRegistered(uint32 , uint16 ) view returns(bool)
func (_Delegatedsending *DelegatedsendingSession) PoolStarsRegistered(arg0 uint32, arg1 uint16) (bool, error) {
	return _Delegatedsending.Contract.PoolStarsRegistered(&_Delegatedsending.CallOpts, arg0, arg1)
}

// PoolStarsRegistered is a free data retrieval call binding the contract method 0x40603cc4.
//
// Solidity: function poolStarsRegistered(uint32 , uint16 ) view returns(bool)
func (_Delegatedsending *DelegatedsendingCallerSession) PoolStarsRegistered(arg0 uint32, arg1 uint16) (bool, error) {
	return _Delegatedsending.Contract.PoolStarsRegistered(&_Delegatedsending.CallOpts, arg0, arg1)
}

// Pools is a free data retrieval call binding the contract method 0x4f4573ee.
//
// Solidity: function pools(uint32 , uint16 ) view returns(uint16)
func (_Delegatedsending *DelegatedsendingCaller) Pools(opts *bind.CallOpts, arg0 uint32, arg1 uint16) (uint16, error) {
	var out []interface{}
	err := _Delegatedsending.contract.Call(opts, &out, "pools", arg0, arg1)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Pools is a free data retrieval call binding the contract method 0x4f4573ee.
//
// Solidity: function pools(uint32 , uint16 ) view returns(uint16)
func (_Delegatedsending *DelegatedsendingSession) Pools(arg0 uint32, arg1 uint16) (uint16, error) {
	return _Delegatedsending.Contract.Pools(&_Delegatedsending.CallOpts, arg0, arg1)
}

// Pools is a free data retrieval call binding the contract method 0x4f4573ee.
//
// Solidity: function pools(uint32 , uint16 ) view returns(uint16)
func (_Delegatedsending *DelegatedsendingCallerSession) Pools(arg0 uint32, arg1 uint16) (uint16, error) {
	return _Delegatedsending.Contract.Pools(&_Delegatedsending.CallOpts, arg0, arg1)
}

// SendPoint is a paid mutator transaction binding the contract method 0x66008294.
//
// Solidity: function sendPoint(uint32 _as, uint32 _point, address _to) returns()
func (_Delegatedsending *DelegatedsendingTransactor) SendPoint(opts *bind.TransactOpts, _as uint32, _point uint32, _to common.Address) (*types.Transaction, error) {
	return _Delegatedsending.contract.Transact(opts, "sendPoint", _as, _point, _to)
}

// SendPoint is a paid mutator transaction binding the contract method 0x66008294.
//
// Solidity: function sendPoint(uint32 _as, uint32 _point, address _to) returns()
func (_Delegatedsending *DelegatedsendingSession) SendPoint(_as uint32, _point uint32, _to common.Address) (*types.Transaction, error) {
	return _Delegatedsending.Contract.SendPoint(&_Delegatedsending.TransactOpts, _as, _point, _to)
}

// SendPoint is a paid mutator transaction binding the contract method 0x66008294.
//
// Solidity: function sendPoint(uint32 _as, uint32 _point, address _to) returns()
func (_Delegatedsending *DelegatedsendingTransactorSession) SendPoint(_as uint32, _point uint32, _to common.Address) (*types.Transaction, error) {
	return _Delegatedsending.Contract.SendPoint(&_Delegatedsending.TransactOpts, _as, _point, _to)
}

// SetPoolSize is a paid mutator transaction binding the contract method 0xe78b8a37.
//
// Solidity: function setPoolSize(uint16 _as, uint32 _for, uint16 _size) returns()
func (_Delegatedsending *DelegatedsendingTransactor) SetPoolSize(opts *bind.TransactOpts, _as uint16, _for uint32, _size uint16) (*types.Transaction, error) {
	return _Delegatedsending.contract.Transact(opts, "setPoolSize", _as, _for, _size)
}

// SetPoolSize is a paid mutator transaction binding the contract method 0xe78b8a37.
//
// Solidity: function setPoolSize(uint16 _as, uint32 _for, uint16 _size) returns()
func (_Delegatedsending *DelegatedsendingSession) SetPoolSize(_as uint16, _for uint32, _size uint16) (*types.Transaction, error) {
	return _Delegatedsending.Contract.SetPoolSize(&_Delegatedsending.TransactOpts, _as, _for, _size)
}

// SetPoolSize is a paid mutator transaction binding the contract method 0xe78b8a37.
//
// Solidity: function setPoolSize(uint16 _as, uint32 _for, uint16 _size) returns()
func (_Delegatedsending *DelegatedsendingTransactorSession) SetPoolSize(_as uint16, _for uint32, _size uint16) (*types.Transaction, error) {
	return _Delegatedsending.Contract.SetPoolSize(&_Delegatedsending.TransactOpts, _as, _for, _size)
}

// DelegatedsendingPoolIterator is returned from FilterPool and is used to iterate over the raw logs and unpacked data for Pool events raised by the Delegatedsending contract.
type DelegatedsendingPoolIterator struct {
	Event *DelegatedsendingPool // Event containing the contract specifics and raw log

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
func (it *DelegatedsendingPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatedsendingPool)
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
		it.Event = new(DelegatedsendingPool)
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
func (it *DelegatedsendingPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatedsendingPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatedsendingPool represents a Pool event raised by the Delegatedsending contract.
type DelegatedsendingPool struct {
	Prefix uint16
	Who    uint32
	Size   uint16
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPool is a free log retrieval operation binding the contract event 0xb05d354b3ecb8f9593b9298bcdeea36401f0e199bc0617e9a731a45cecb343ec.
//
// Solidity: event Pool(uint16 indexed prefix, uint32 indexed who, uint16 size)
func (_Delegatedsending *DelegatedsendingFilterer) FilterPool(opts *bind.FilterOpts, prefix []uint16, who []uint32) (*DelegatedsendingPoolIterator, error) {

	var prefixRule []interface{}
	for _, prefixItem := range prefix {
		prefixRule = append(prefixRule, prefixItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Delegatedsending.contract.FilterLogs(opts, "Pool", prefixRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DelegatedsendingPoolIterator{contract: _Delegatedsending.contract, event: "Pool", logs: logs, sub: sub}, nil
}

// WatchPool is a free log subscription operation binding the contract event 0xb05d354b3ecb8f9593b9298bcdeea36401f0e199bc0617e9a731a45cecb343ec.
//
// Solidity: event Pool(uint16 indexed prefix, uint32 indexed who, uint16 size)
func (_Delegatedsending *DelegatedsendingFilterer) WatchPool(opts *bind.WatchOpts, sink chan<- *DelegatedsendingPool, prefix []uint16, who []uint32) (event.Subscription, error) {

	var prefixRule []interface{}
	for _, prefixItem := range prefix {
		prefixRule = append(prefixRule, prefixItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Delegatedsending.contract.WatchLogs(opts, "Pool", prefixRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatedsendingPool)
				if err := _Delegatedsending.contract.UnpackLog(event, "Pool", log); err != nil {
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

// ParsePool is a log parse operation binding the contract event 0xb05d354b3ecb8f9593b9298bcdeea36401f0e199bc0617e9a731a45cecb343ec.
//
// Solidity: event Pool(uint16 indexed prefix, uint32 indexed who, uint16 size)
func (_Delegatedsending *DelegatedsendingFilterer) ParsePool(log types.Log) (*DelegatedsendingPool, error) {
	event := new(DelegatedsendingPool)
	if err := _Delegatedsending.contract.UnpackLog(event, "Pool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatedsendingSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the Delegatedsending contract.
type DelegatedsendingSentIterator struct {
	Event *DelegatedsendingSent // Event containing the contract specifics and raw log

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
func (it *DelegatedsendingSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatedsendingSent)
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
		it.Event = new(DelegatedsendingSent)
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
func (it *DelegatedsendingSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatedsendingSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatedsendingSent represents a Sent event raised by the Delegatedsending contract.
type DelegatedsendingSent struct {
	Prefix   uint16
	FromPool uint32
	By       uint32
	Point    uint32
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0x47638e3cddee220481e4c3f9183d639c0efea7f05fcd2df4188855729f715419.
//
// Solidity: event Sent(uint16 indexed prefix, uint32 indexed fromPool, uint32 by, uint32 point, address to)
func (_Delegatedsending *DelegatedsendingFilterer) FilterSent(opts *bind.FilterOpts, prefix []uint16, fromPool []uint32) (*DelegatedsendingSentIterator, error) {

	var prefixRule []interface{}
	for _, prefixItem := range prefix {
		prefixRule = append(prefixRule, prefixItem)
	}
	var fromPoolRule []interface{}
	for _, fromPoolItem := range fromPool {
		fromPoolRule = append(fromPoolRule, fromPoolItem)
	}

	logs, sub, err := _Delegatedsending.contract.FilterLogs(opts, "Sent", prefixRule, fromPoolRule)
	if err != nil {
		return nil, err
	}
	return &DelegatedsendingSentIterator{contract: _Delegatedsending.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0x47638e3cddee220481e4c3f9183d639c0efea7f05fcd2df4188855729f715419.
//
// Solidity: event Sent(uint16 indexed prefix, uint32 indexed fromPool, uint32 by, uint32 point, address to)
func (_Delegatedsending *DelegatedsendingFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *DelegatedsendingSent, prefix []uint16, fromPool []uint32) (event.Subscription, error) {

	var prefixRule []interface{}
	for _, prefixItem := range prefix {
		prefixRule = append(prefixRule, prefixItem)
	}
	var fromPoolRule []interface{}
	for _, fromPoolItem := range fromPool {
		fromPoolRule = append(fromPoolRule, fromPoolItem)
	}

	logs, sub, err := _Delegatedsending.contract.WatchLogs(opts, "Sent", prefixRule, fromPoolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatedsendingSent)
				if err := _Delegatedsending.contract.UnpackLog(event, "Sent", log); err != nil {
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

// ParseSent is a log parse operation binding the contract event 0x47638e3cddee220481e4c3f9183d639c0efea7f05fcd2df4188855729f715419.
//
// Solidity: event Sent(uint16 indexed prefix, uint32 indexed fromPool, uint32 by, uint32 point, address to)
func (_Delegatedsending *DelegatedsendingFilterer) ParseSent(log types.Log) (*DelegatedsendingSent, error) {
	event := new(DelegatedsendingSent)
	if err := _Delegatedsending.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
