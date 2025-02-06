// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package linearstarrelease

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

// LinearstarreleaseMetaData contains all meta data concerning the Linearstarrelease contract.
var LinearstarreleaseMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"startReleasing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_participant\",\"type\":\"address\"}],\"name\":\"getRemainingStars\",\"outputs\":[{\"name\":\"stars\",\"type\":\"uint16[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_participant\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawOverdue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_participant\",\"type\":\"address\"}],\"name\":\"verifyBalance\",\"outputs\":[{\"name\":\"correct\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"start\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"transferBatch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_participant\",\"type\":\"address\"},{\"name\":\"_windup\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint16\"},{\"name\":\"_rate\",\"type\":\"uint16\"},{\"name\":\"_rateUnit\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"azimuth\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"batches\",\"outputs\":[{\"name\":\"windup\",\"type\":\"uint256\"},{\"name\":\"rateUnit\",\"type\":\"uint256\"},{\"name\":\"withdrawn\",\"type\":\"uint16\"},{\"name\":\"rate\",\"type\":\"uint16\"},{\"name\":\"amount\",\"type\":\"uint16\"},{\"name\":\"approvedTransferTo\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"approveBatchTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_participant\",\"type\":\"address\"},{\"name\":\"_star\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_participant\",\"type\":\"address\"}],\"name\":\"withdrawLimit\",\"outputs\":[{\"name\":\"limit\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_azimuth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051602080612872833981018060405281019080805190602001909291905050508080336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506127aa806100c86000396000f3006080604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063049efcc7146100f65780633ccfd60b1461010d57806351cff8d914610124578063715018a61461016757806377bffc9a1461017e5780638da5cb5b146102165780639882e15e1461026d578063a00ddad1146102d0578063be9a65551461032b578063bf54789414610356578063bfca1ead14610399578063d40ffacb1461040c578063e2ea265814610463578063e596d81114610521578063e6deefa914610564578063f2fde38b146105b5578063fce33f01146105f8575b600080fd5b34801561010257600080fd5b5061010b610657565b005b34801561011957600080fd5b506101226106cc565b005b34801561013057600080fd5b50610165600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106d7565b005b34801561017357600080fd5b5061017c61076d565b005b34801561018a57600080fd5b506101bf600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061086f565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156102025780820151818401526020810190506101e7565b505050509050019250505060405180910390f35b34801561022257600080fd5b5061022b610931565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561027957600080fd5b506102ce600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610956565b005b3480156102dc57600080fd5b50610311600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a33565b604051808215151515815260200191505060405180910390f35b34801561033757600080fd5b50610340610ac7565b6040518082815260200191505060405180910390f35b34801561036257600080fd5b50610397600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610acd565b005b3480156103a557600080fd5b5061040a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803561ffff169060200190929190803561ffff16906020019092919080359060200190929190505050610f2a565b005b34801561041857600080fd5b50610421611074565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561046f57600080fd5b506104a4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061109a565b604051808781526020018681526020018561ffff1661ffff1681526020018461ffff1661ffff1681526020018361ffff1661ffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001965050505050505060405180910390f35b34801561052d57600080fd5b50610562600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611120565b005b34801561057057600080fd5b506105b3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803561ffff169060200190929190505050611268565b005b3480156105c157600080fd5b506105f6600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506113d7565b005b34801561060457600080fd5b50610639600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061143e565b604051808261ffff1661ffff16815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156106b257600080fd5b60025460001415156106c357600080fd5b42600281905550565b6106d5336106d7565b565b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816000018054905011801561075257506107353361143e565b61ffff168160030160009054906101000a900461ffff1661ffff16105b151561075d57600080fd5b61076981836001611563565b5050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156107c857600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6060600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000180548060200260200160405190810160405280929190818152602001828054801561092557602002820191906000526020600020906000905b82829054906101000a900461ffff1661ffff16815260200190600201906020826001010492830192600103820291508084116108ec5790505b50505050509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156109b157600080fd5b60025460001080156109da57506109d76312cc030060025461164990919063ffffffff16565b42115b15156109e557600080fd5b610a2f600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020826000611563565b5050565b600080600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060000180549050610aba8260030160009054906101000a900461ffff168360030160049054906101000a900461ffff1661ffff1661166590919063ffffffff16565b61ffff1614915050919050565b60025481565b60003373ffffffffffffffffffffffffffffffffffffffff16600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160069054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610b6b57600080fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160049054906101000a900461ffff1661ffff166000141515610bce57600080fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905080600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820181600001908054610c659291906125e4565b5060018201548160010155600282015481600201556003820160009054906101000a900461ffff168160030160006101000a81548161ffff021916908361ffff1602179055506003820160029054906101000a900461ffff168160030160026101000a81548161ffff021916908361ffff1602179055506003820160049054906101000a900461ffff168160030160046101000a81548161ffff021916908361ffff1602179055506003820160069054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160030160066101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555090505060e0604051908101604052806000604051908082528060200260200182016040528015610db35781602001602082028038833980820191505090505b5081526020016000815260200160008152602001600061ffff168152602001600061ffff168152602001600061ffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815250600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000019080519060200190610e5e929190612644565b50602082015181600101556040820151816002015560608201518160030160006101000a81548161ffff021916908361ffff16021790555060808201518160030160026101000a81548161ffff021916908361ffff16021790555060a08201518160030160046101000a81548161ffff021916908361ffff16021790555060c08201518160030160066101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610f8757600080fd5b600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060030160049054906101000a900461ffff1661ffff166000141515610fed57600080fd5b60008361ffff161180156110015750600082115b8015611011575060008461ffff16115b151561101c57600080fd5b848160010181905550838160030160046101000a81548161ffff021916908361ffff160217905550828160030160026101000a81548161ffff021916908361ffff160217905550818160020181905550505050505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60036020528060005260406000206000915090508060010154908060020154908060030160009054906101000a900461ffff16908060030160029054906101000a900461ffff16908060030160049054906101000a900461ffff16908060030160069054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905086565b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160049054906101000a900461ffff1661ffff166000141580156111d95750600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160049054906101000a900461ffff1661ffff166000145b15156111e457600080fd5b80600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160066101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156112c557600080fd5b600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060ff8261ffff1611801561135e57506113508160030160009054906101000a900461ffff168260030160049054906101000a900461ffff1661ffff1661166590919063ffffffff16565b61ffff168160000180549050105b151561136957600080fd5b6113788261ffff166001611686565b151561138357600080fd5b8060000182908060018154018082558091505090600182039060005260206000209060109182820401919006600202909192909190916101000a81548161ffff021916908361ffff16021790555050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561143257600080fd5b61143b816121b8565b50565b60008060008060025460001415611458576000935061155b565b60009250600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002091506114b6826001015460025461164990919063ffffffff16565b90508042111561150d5761150a82600201546114db83426122b290919063ffffffff16565b8115156114e457fe5b048360030160029054906101000a900461ffff1661ffff166122cb90919063ffffffff16565b92505b600183101561151f576001935061155b565b8160030160049054906101000a900461ffff1661ffff16831115611557578160030160049054906101000a900461ffff16935061155b565b8293505b505050919050565b600083600001611584600186600001805490506122b290919063ffffffff16565b81548110151561159057fe5b90600052602060002090601091828204019190066002029054906101000a900461ffff1690506115d1600185600001805490506122b290919063ffffffff16565b84600001816115e091906126ee565b5061160b60018560030160009054906101000a900461ffff1661ffff1661230390919063ffffffff16565b8460030160006101000a81548161ffff021916908361ffff1602179055506116388161ffff168484612329565b151561164357600080fd5b50505050565b6000818301905082811015151561165c57fe5b80905092915050565b60008261ffff168261ffff161115151561167b57fe5b818303905092915050565b6000806000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561171157600080fd5b505af1158015611725573d6000803e3d6000fd5b505050506040513d602081101561173b57600080fd5b81019080805190602001909291905050509150600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e4a358d7866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b1580156117eb57600080fd5b505af11580156117ff573d6000803e3d6000fd5b505050506040513d602081101561181557600080fd5b81019080805190602001909291905050509050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663caf590f98660006040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b1580156118e457600080fd5b505af11580156118f8573d6000803e3d6000fd5b505050506040513d602081101561190e57600080fd5b81019080805190602001909291905050508015611a315750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663caf590f982336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808361ffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b1580156119f557600080fd5b505af1158015611a09573d6000803e3d6000fd5b505050506040513d6020811015611a1f57600080fd5b81019080805190602001909291905050505b8015611b435750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324ba1a4682306040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808361ffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015611b0757600080fd5b505af1158015611b1b573d6000803e3d6000fd5b505050506040513d6020811015611b3157600080fd5b81019080805190602001909291905050505b8015611cea5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663293a9169826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808261ffff1663ffffffff168152602001915050602060405180830381600087803b158015611be557600080fd5b505af1158015611bf9573d6000803e3d6000fd5b505050506040513d6020811015611c0f57600080fd5b810190808051906020019092919050505063ffffffff168273ffffffffffffffffffffffffffffffffffffffff1663ef20bff883426040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808361ffff1663ffffffff16815260200182815260200192505050602060405180830381600087803b158015611ca757600080fd5b505af1158015611cbb573d6000803e3d6000fd5b505050506040513d6020811015611cd157600080fd5b810190808051906020019092919050505063ffffffff16115b15611dbf578173ffffffffffffffffffffffffffffffffffffffff1663a0d3253f86306040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b158015611d9e57600080fd5b505af1158015611db2573d6000803e3d6000fd5b50505050600192506121b0565b831580611ea15750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636d09887b866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015611e6457600080fd5b505af1158015611e78573d6000803e3d6000fd5b505050506040513d6020811015611e8e57600080fd5b8101908080519060200190929190505050155b8015611fb55750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663caf590f986336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015611f7957600080fd5b505af1158015611f8d573d6000803e3d6000fd5b505050506040513d6020811015611fa357600080fd5b81019080805190602001909291905050505b80156120c95750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663728aa85786306040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561208d57600080fd5b505af11580156120a1573d6000803e3d6000fd5b505050506040513d60208110156120b757600080fd5b81019080805190602001909291905050505b156121ab578173ffffffffffffffffffffffffffffffffffffffff16631e79a85b863060016040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808463ffffffff1663ffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001821515151581526020019350505050600060405180830381600087803b15801561218a57600080fd5b505af115801561219e573d6000803e3d6000fd5b50505050600192506121b0565b600092505b505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156121f457600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008282111515156122c057fe5b818303905092915050565b6000808314156122de57600090506122fd565b81830290508183828115156122ef57fe5b041415156122f957fe5b8090505b92915050565b60008082840190508361ffff168161ffff161015151561231f57fe5b8091505092915050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663caf590f985306040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b1580156123fc57600080fd5b505af1158015612410573d6000803e3d6000fd5b505050506040513d602081101561242657600080fd5b8101908080519060200190929190505050156125d857600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156124c257600080fd5b505af11580156124d6573d6000803e3d6000fd5b505050506040513d60208110156124ec57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff16631e79a85b8585856040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808463ffffffff1663ffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001821515151581526020019350505050600060405180830381600087803b1580156125b757600080fd5b505af11580156125cb573d6000803e3d6000fd5b50505050600190506125dd565b600090505b9392505050565b82805482825590600052602060002090600f0160109004810192821561263357600052602060002091600f016010900482015b82811115612632578254825591600101919060010190612617565b5b5090506126409190612728565b5090565b82805482825590600052602060002090600f016010900481019282156126dd5791602002820160005b838211156126ad57835183826101000a81548161ffff021916908361ffff160217905550926020019260020160208160010104928301926001030261266d565b80156126db5782816101000a81549061ffff02191690556002016020816001010492830192600103026126ad565b505b5090506126ea9190612728565b5090565b81548183558181111561272357600f016010900481600f016010900483600052602060002091820191016127229190612759565b5b505050565b61275691905b8082111561275257600081816101000a81549061ffff02191690555060010161272e565b5090565b90565b61277b91905b8082111561277757600081600090555060010161275f565b5090565b905600a165627a7a72305820fc8714256c24f7055dec48f06168d27b996fb06a783b741a2fed8fbdef2976a60029",
}

// LinearstarreleaseABI is the input ABI used to generate the binding from.
// Deprecated: Use LinearstarreleaseMetaData.ABI instead.
var LinearstarreleaseABI = LinearstarreleaseMetaData.ABI

// LinearstarreleaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LinearstarreleaseMetaData.Bin instead.
var LinearstarreleaseBin = LinearstarreleaseMetaData.Bin

// DeployLinearstarrelease deploys a new Ethereum contract, binding an instance of Linearstarrelease to it.
func DeployLinearstarrelease(auth *bind.TransactOpts, backend bind.ContractBackend, _azimuth common.Address) (common.Address, *types.Transaction, *Linearstarrelease, error) {
	parsed, err := LinearstarreleaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LinearstarreleaseBin), backend, _azimuth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Linearstarrelease{LinearstarreleaseCaller: LinearstarreleaseCaller{contract: contract}, LinearstarreleaseTransactor: LinearstarreleaseTransactor{contract: contract}, LinearstarreleaseFilterer: LinearstarreleaseFilterer{contract: contract}}, nil
}

// Linearstarrelease is an auto generated Go binding around an Ethereum contract.
type Linearstarrelease struct {
	LinearstarreleaseCaller     // Read-only binding to the contract
	LinearstarreleaseTransactor // Write-only binding to the contract
	LinearstarreleaseFilterer   // Log filterer for contract events
}

// LinearstarreleaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type LinearstarreleaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinearstarreleaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LinearstarreleaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinearstarreleaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LinearstarreleaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinearstarreleaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LinearstarreleaseSession struct {
	Contract     *Linearstarrelease // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LinearstarreleaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LinearstarreleaseCallerSession struct {
	Contract *LinearstarreleaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LinearstarreleaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LinearstarreleaseTransactorSession struct {
	Contract     *LinearstarreleaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LinearstarreleaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type LinearstarreleaseRaw struct {
	Contract *Linearstarrelease // Generic contract binding to access the raw methods on
}

// LinearstarreleaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LinearstarreleaseCallerRaw struct {
	Contract *LinearstarreleaseCaller // Generic read-only contract binding to access the raw methods on
}

// LinearstarreleaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LinearstarreleaseTransactorRaw struct {
	Contract *LinearstarreleaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLinearstarrelease creates a new instance of Linearstarrelease, bound to a specific deployed contract.
func NewLinearstarrelease(address common.Address, backend bind.ContractBackend) (*Linearstarrelease, error) {
	contract, err := bindLinearstarrelease(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Linearstarrelease{LinearstarreleaseCaller: LinearstarreleaseCaller{contract: contract}, LinearstarreleaseTransactor: LinearstarreleaseTransactor{contract: contract}, LinearstarreleaseFilterer: LinearstarreleaseFilterer{contract: contract}}, nil
}

// NewLinearstarreleaseCaller creates a new read-only instance of Linearstarrelease, bound to a specific deployed contract.
func NewLinearstarreleaseCaller(address common.Address, caller bind.ContractCaller) (*LinearstarreleaseCaller, error) {
	contract, err := bindLinearstarrelease(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LinearstarreleaseCaller{contract: contract}, nil
}

// NewLinearstarreleaseTransactor creates a new write-only instance of Linearstarrelease, bound to a specific deployed contract.
func NewLinearstarreleaseTransactor(address common.Address, transactor bind.ContractTransactor) (*LinearstarreleaseTransactor, error) {
	contract, err := bindLinearstarrelease(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LinearstarreleaseTransactor{contract: contract}, nil
}

// NewLinearstarreleaseFilterer creates a new log filterer instance of Linearstarrelease, bound to a specific deployed contract.
func NewLinearstarreleaseFilterer(address common.Address, filterer bind.ContractFilterer) (*LinearstarreleaseFilterer, error) {
	contract, err := bindLinearstarrelease(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LinearstarreleaseFilterer{contract: contract}, nil
}

// bindLinearstarrelease binds a generic wrapper to an already deployed contract.
func bindLinearstarrelease(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LinearstarreleaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Linearstarrelease *LinearstarreleaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Linearstarrelease.Contract.LinearstarreleaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Linearstarrelease *LinearstarreleaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.LinearstarreleaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Linearstarrelease *LinearstarreleaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.LinearstarreleaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Linearstarrelease *LinearstarreleaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Linearstarrelease.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Linearstarrelease *LinearstarreleaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Linearstarrelease *LinearstarreleaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.contract.Transact(opts, method, params...)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Linearstarrelease *LinearstarreleaseCaller) Azimuth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Linearstarrelease.contract.Call(opts, &out, "azimuth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Linearstarrelease *LinearstarreleaseSession) Azimuth() (common.Address, error) {
	return _Linearstarrelease.Contract.Azimuth(&_Linearstarrelease.CallOpts)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Linearstarrelease *LinearstarreleaseCallerSession) Azimuth() (common.Address, error) {
	return _Linearstarrelease.Contract.Azimuth(&_Linearstarrelease.CallOpts)
}

// Batches is a free data retrieval call binding the contract method 0xe2ea2658.
//
// Solidity: function batches(address ) view returns(uint256 windup, uint256 rateUnit, uint16 withdrawn, uint16 rate, uint16 amount, address approvedTransferTo)
func (_Linearstarrelease *LinearstarreleaseCaller) Batches(opts *bind.CallOpts, arg0 common.Address) (struct {
	Windup             *big.Int
	RateUnit           *big.Int
	Withdrawn          uint16
	Rate               uint16
	Amount             uint16
	ApprovedTransferTo common.Address
}, error) {
	var out []interface{}
	err := _Linearstarrelease.contract.Call(opts, &out, "batches", arg0)

	outstruct := new(struct {
		Windup             *big.Int
		RateUnit           *big.Int
		Withdrawn          uint16
		Rate               uint16
		Amount             uint16
		ApprovedTransferTo common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Windup = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RateUnit = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Withdrawn = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.Rate = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.Amount = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.ApprovedTransferTo = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Batches is a free data retrieval call binding the contract method 0xe2ea2658.
//
// Solidity: function batches(address ) view returns(uint256 windup, uint256 rateUnit, uint16 withdrawn, uint16 rate, uint16 amount, address approvedTransferTo)
func (_Linearstarrelease *LinearstarreleaseSession) Batches(arg0 common.Address) (struct {
	Windup             *big.Int
	RateUnit           *big.Int
	Withdrawn          uint16
	Rate               uint16
	Amount             uint16
	ApprovedTransferTo common.Address
}, error) {
	return _Linearstarrelease.Contract.Batches(&_Linearstarrelease.CallOpts, arg0)
}

// Batches is a free data retrieval call binding the contract method 0xe2ea2658.
//
// Solidity: function batches(address ) view returns(uint256 windup, uint256 rateUnit, uint16 withdrawn, uint16 rate, uint16 amount, address approvedTransferTo)
func (_Linearstarrelease *LinearstarreleaseCallerSession) Batches(arg0 common.Address) (struct {
	Windup             *big.Int
	RateUnit           *big.Int
	Withdrawn          uint16
	Rate               uint16
	Amount             uint16
	ApprovedTransferTo common.Address
}, error) {
	return _Linearstarrelease.Contract.Batches(&_Linearstarrelease.CallOpts, arg0)
}

// GetRemainingStars is a free data retrieval call binding the contract method 0x77bffc9a.
//
// Solidity: function getRemainingStars(address _participant) view returns(uint16[] stars)
func (_Linearstarrelease *LinearstarreleaseCaller) GetRemainingStars(opts *bind.CallOpts, _participant common.Address) ([]uint16, error) {
	var out []interface{}
	err := _Linearstarrelease.contract.Call(opts, &out, "getRemainingStars", _participant)

	if err != nil {
		return *new([]uint16), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint16)).(*[]uint16)

	return out0, err

}

// GetRemainingStars is a free data retrieval call binding the contract method 0x77bffc9a.
//
// Solidity: function getRemainingStars(address _participant) view returns(uint16[] stars)
func (_Linearstarrelease *LinearstarreleaseSession) GetRemainingStars(_participant common.Address) ([]uint16, error) {
	return _Linearstarrelease.Contract.GetRemainingStars(&_Linearstarrelease.CallOpts, _participant)
}

// GetRemainingStars is a free data retrieval call binding the contract method 0x77bffc9a.
//
// Solidity: function getRemainingStars(address _participant) view returns(uint16[] stars)
func (_Linearstarrelease *LinearstarreleaseCallerSession) GetRemainingStars(_participant common.Address) ([]uint16, error) {
	return _Linearstarrelease.Contract.GetRemainingStars(&_Linearstarrelease.CallOpts, _participant)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Linearstarrelease *LinearstarreleaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Linearstarrelease.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Linearstarrelease *LinearstarreleaseSession) Owner() (common.Address, error) {
	return _Linearstarrelease.Contract.Owner(&_Linearstarrelease.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Linearstarrelease *LinearstarreleaseCallerSession) Owner() (common.Address, error) {
	return _Linearstarrelease.Contract.Owner(&_Linearstarrelease.CallOpts)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() view returns(uint256)
func (_Linearstarrelease *LinearstarreleaseCaller) Start(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Linearstarrelease.contract.Call(opts, &out, "start")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() view returns(uint256)
func (_Linearstarrelease *LinearstarreleaseSession) Start() (*big.Int, error) {
	return _Linearstarrelease.Contract.Start(&_Linearstarrelease.CallOpts)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() view returns(uint256)
func (_Linearstarrelease *LinearstarreleaseCallerSession) Start() (*big.Int, error) {
	return _Linearstarrelease.Contract.Start(&_Linearstarrelease.CallOpts)
}

// VerifyBalance is a free data retrieval call binding the contract method 0xa00ddad1.
//
// Solidity: function verifyBalance(address _participant) view returns(bool correct)
func (_Linearstarrelease *LinearstarreleaseCaller) VerifyBalance(opts *bind.CallOpts, _participant common.Address) (bool, error) {
	var out []interface{}
	err := _Linearstarrelease.contract.Call(opts, &out, "verifyBalance", _participant)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyBalance is a free data retrieval call binding the contract method 0xa00ddad1.
//
// Solidity: function verifyBalance(address _participant) view returns(bool correct)
func (_Linearstarrelease *LinearstarreleaseSession) VerifyBalance(_participant common.Address) (bool, error) {
	return _Linearstarrelease.Contract.VerifyBalance(&_Linearstarrelease.CallOpts, _participant)
}

// VerifyBalance is a free data retrieval call binding the contract method 0xa00ddad1.
//
// Solidity: function verifyBalance(address _participant) view returns(bool correct)
func (_Linearstarrelease *LinearstarreleaseCallerSession) VerifyBalance(_participant common.Address) (bool, error) {
	return _Linearstarrelease.Contract.VerifyBalance(&_Linearstarrelease.CallOpts, _participant)
}

// WithdrawLimit is a free data retrieval call binding the contract method 0xfce33f01.
//
// Solidity: function withdrawLimit(address _participant) view returns(uint16 limit)
func (_Linearstarrelease *LinearstarreleaseCaller) WithdrawLimit(opts *bind.CallOpts, _participant common.Address) (uint16, error) {
	var out []interface{}
	err := _Linearstarrelease.contract.Call(opts, &out, "withdrawLimit", _participant)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// WithdrawLimit is a free data retrieval call binding the contract method 0xfce33f01.
//
// Solidity: function withdrawLimit(address _participant) view returns(uint16 limit)
func (_Linearstarrelease *LinearstarreleaseSession) WithdrawLimit(_participant common.Address) (uint16, error) {
	return _Linearstarrelease.Contract.WithdrawLimit(&_Linearstarrelease.CallOpts, _participant)
}

// WithdrawLimit is a free data retrieval call binding the contract method 0xfce33f01.
//
// Solidity: function withdrawLimit(address _participant) view returns(uint16 limit)
func (_Linearstarrelease *LinearstarreleaseCallerSession) WithdrawLimit(_participant common.Address) (uint16, error) {
	return _Linearstarrelease.Contract.WithdrawLimit(&_Linearstarrelease.CallOpts, _participant)
}

// ApproveBatchTransfer is a paid mutator transaction binding the contract method 0xe596d811.
//
// Solidity: function approveBatchTransfer(address _to) returns()
func (_Linearstarrelease *LinearstarreleaseTransactor) ApproveBatchTransfer(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Linearstarrelease.contract.Transact(opts, "approveBatchTransfer", _to)
}

// ApproveBatchTransfer is a paid mutator transaction binding the contract method 0xe596d811.
//
// Solidity: function approveBatchTransfer(address _to) returns()
func (_Linearstarrelease *LinearstarreleaseSession) ApproveBatchTransfer(_to common.Address) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.ApproveBatchTransfer(&_Linearstarrelease.TransactOpts, _to)
}

// ApproveBatchTransfer is a paid mutator transaction binding the contract method 0xe596d811.
//
// Solidity: function approveBatchTransfer(address _to) returns()
func (_Linearstarrelease *LinearstarreleaseTransactorSession) ApproveBatchTransfer(_to common.Address) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.ApproveBatchTransfer(&_Linearstarrelease.TransactOpts, _to)
}

// Deposit is a paid mutator transaction binding the contract method 0xe6deefa9.
//
// Solidity: function deposit(address _participant, uint16 _star) returns()
func (_Linearstarrelease *LinearstarreleaseTransactor) Deposit(opts *bind.TransactOpts, _participant common.Address, _star uint16) (*types.Transaction, error) {
	return _Linearstarrelease.contract.Transact(opts, "deposit", _participant, _star)
}

// Deposit is a paid mutator transaction binding the contract method 0xe6deefa9.
//
// Solidity: function deposit(address _participant, uint16 _star) returns()
func (_Linearstarrelease *LinearstarreleaseSession) Deposit(_participant common.Address, _star uint16) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.Deposit(&_Linearstarrelease.TransactOpts, _participant, _star)
}

// Deposit is a paid mutator transaction binding the contract method 0xe6deefa9.
//
// Solidity: function deposit(address _participant, uint16 _star) returns()
func (_Linearstarrelease *LinearstarreleaseTransactorSession) Deposit(_participant common.Address, _star uint16) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.Deposit(&_Linearstarrelease.TransactOpts, _participant, _star)
}

// Register is a paid mutator transaction binding the contract method 0xbfca1ead.
//
// Solidity: function register(address _participant, uint256 _windup, uint16 _amount, uint16 _rate, uint256 _rateUnit) returns()
func (_Linearstarrelease *LinearstarreleaseTransactor) Register(opts *bind.TransactOpts, _participant common.Address, _windup *big.Int, _amount uint16, _rate uint16, _rateUnit *big.Int) (*types.Transaction, error) {
	return _Linearstarrelease.contract.Transact(opts, "register", _participant, _windup, _amount, _rate, _rateUnit)
}

// Register is a paid mutator transaction binding the contract method 0xbfca1ead.
//
// Solidity: function register(address _participant, uint256 _windup, uint16 _amount, uint16 _rate, uint256 _rateUnit) returns()
func (_Linearstarrelease *LinearstarreleaseSession) Register(_participant common.Address, _windup *big.Int, _amount uint16, _rate uint16, _rateUnit *big.Int) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.Register(&_Linearstarrelease.TransactOpts, _participant, _windup, _amount, _rate, _rateUnit)
}

// Register is a paid mutator transaction binding the contract method 0xbfca1ead.
//
// Solidity: function register(address _participant, uint256 _windup, uint16 _amount, uint16 _rate, uint256 _rateUnit) returns()
func (_Linearstarrelease *LinearstarreleaseTransactorSession) Register(_participant common.Address, _windup *big.Int, _amount uint16, _rate uint16, _rateUnit *big.Int) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.Register(&_Linearstarrelease.TransactOpts, _participant, _windup, _amount, _rate, _rateUnit)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Linearstarrelease *LinearstarreleaseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Linearstarrelease.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Linearstarrelease *LinearstarreleaseSession) RenounceOwnership() (*types.Transaction, error) {
	return _Linearstarrelease.Contract.RenounceOwnership(&_Linearstarrelease.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Linearstarrelease *LinearstarreleaseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Linearstarrelease.Contract.RenounceOwnership(&_Linearstarrelease.TransactOpts)
}

// StartReleasing is a paid mutator transaction binding the contract method 0x049efcc7.
//
// Solidity: function startReleasing() returns()
func (_Linearstarrelease *LinearstarreleaseTransactor) StartReleasing(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Linearstarrelease.contract.Transact(opts, "startReleasing")
}

// StartReleasing is a paid mutator transaction binding the contract method 0x049efcc7.
//
// Solidity: function startReleasing() returns()
func (_Linearstarrelease *LinearstarreleaseSession) StartReleasing() (*types.Transaction, error) {
	return _Linearstarrelease.Contract.StartReleasing(&_Linearstarrelease.TransactOpts)
}

// StartReleasing is a paid mutator transaction binding the contract method 0x049efcc7.
//
// Solidity: function startReleasing() returns()
func (_Linearstarrelease *LinearstarreleaseTransactorSession) StartReleasing() (*types.Transaction, error) {
	return _Linearstarrelease.Contract.StartReleasing(&_Linearstarrelease.TransactOpts)
}

// TransferBatch is a paid mutator transaction binding the contract method 0xbf547894.
//
// Solidity: function transferBatch(address _from) returns()
func (_Linearstarrelease *LinearstarreleaseTransactor) TransferBatch(opts *bind.TransactOpts, _from common.Address) (*types.Transaction, error) {
	return _Linearstarrelease.contract.Transact(opts, "transferBatch", _from)
}

// TransferBatch is a paid mutator transaction binding the contract method 0xbf547894.
//
// Solidity: function transferBatch(address _from) returns()
func (_Linearstarrelease *LinearstarreleaseSession) TransferBatch(_from common.Address) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.TransferBatch(&_Linearstarrelease.TransactOpts, _from)
}

// TransferBatch is a paid mutator transaction binding the contract method 0xbf547894.
//
// Solidity: function transferBatch(address _from) returns()
func (_Linearstarrelease *LinearstarreleaseTransactorSession) TransferBatch(_from common.Address) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.TransferBatch(&_Linearstarrelease.TransactOpts, _from)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Linearstarrelease *LinearstarreleaseTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Linearstarrelease.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Linearstarrelease *LinearstarreleaseSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.TransferOwnership(&_Linearstarrelease.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Linearstarrelease *LinearstarreleaseTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.TransferOwnership(&_Linearstarrelease.TransactOpts, _newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Linearstarrelease *LinearstarreleaseTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Linearstarrelease.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Linearstarrelease *LinearstarreleaseSession) Withdraw() (*types.Transaction, error) {
	return _Linearstarrelease.Contract.Withdraw(&_Linearstarrelease.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Linearstarrelease *LinearstarreleaseTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Linearstarrelease.Contract.Withdraw(&_Linearstarrelease.TransactOpts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _to) returns()
func (_Linearstarrelease *LinearstarreleaseTransactor) Withdraw0(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Linearstarrelease.contract.Transact(opts, "withdraw0", _to)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _to) returns()
func (_Linearstarrelease *LinearstarreleaseSession) Withdraw0(_to common.Address) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.Withdraw0(&_Linearstarrelease.TransactOpts, _to)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _to) returns()
func (_Linearstarrelease *LinearstarreleaseTransactorSession) Withdraw0(_to common.Address) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.Withdraw0(&_Linearstarrelease.TransactOpts, _to)
}

// WithdrawOverdue is a paid mutator transaction binding the contract method 0x9882e15e.
//
// Solidity: function withdrawOverdue(address _participant, address _to) returns()
func (_Linearstarrelease *LinearstarreleaseTransactor) WithdrawOverdue(opts *bind.TransactOpts, _participant common.Address, _to common.Address) (*types.Transaction, error) {
	return _Linearstarrelease.contract.Transact(opts, "withdrawOverdue", _participant, _to)
}

// WithdrawOverdue is a paid mutator transaction binding the contract method 0x9882e15e.
//
// Solidity: function withdrawOverdue(address _participant, address _to) returns()
func (_Linearstarrelease *LinearstarreleaseSession) WithdrawOverdue(_participant common.Address, _to common.Address) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.WithdrawOverdue(&_Linearstarrelease.TransactOpts, _participant, _to)
}

// WithdrawOverdue is a paid mutator transaction binding the contract method 0x9882e15e.
//
// Solidity: function withdrawOverdue(address _participant, address _to) returns()
func (_Linearstarrelease *LinearstarreleaseTransactorSession) WithdrawOverdue(_participant common.Address, _to common.Address) (*types.Transaction, error) {
	return _Linearstarrelease.Contract.WithdrawOverdue(&_Linearstarrelease.TransactOpts, _participant, _to)
}

// LinearstarreleaseOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Linearstarrelease contract.
type LinearstarreleaseOwnershipRenouncedIterator struct {
	Event *LinearstarreleaseOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *LinearstarreleaseOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinearstarreleaseOwnershipRenounced)
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
		it.Event = new(LinearstarreleaseOwnershipRenounced)
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
func (it *LinearstarreleaseOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LinearstarreleaseOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LinearstarreleaseOwnershipRenounced represents a OwnershipRenounced event raised by the Linearstarrelease contract.
type LinearstarreleaseOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Linearstarrelease *LinearstarreleaseFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*LinearstarreleaseOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Linearstarrelease.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LinearstarreleaseOwnershipRenouncedIterator{contract: _Linearstarrelease.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Linearstarrelease *LinearstarreleaseFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *LinearstarreleaseOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Linearstarrelease.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LinearstarreleaseOwnershipRenounced)
				if err := _Linearstarrelease.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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
func (_Linearstarrelease *LinearstarreleaseFilterer) ParseOwnershipRenounced(log types.Log) (*LinearstarreleaseOwnershipRenounced, error) {
	event := new(LinearstarreleaseOwnershipRenounced)
	if err := _Linearstarrelease.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LinearstarreleaseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Linearstarrelease contract.
type LinearstarreleaseOwnershipTransferredIterator struct {
	Event *LinearstarreleaseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LinearstarreleaseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinearstarreleaseOwnershipTransferred)
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
		it.Event = new(LinearstarreleaseOwnershipTransferred)
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
func (it *LinearstarreleaseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LinearstarreleaseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LinearstarreleaseOwnershipTransferred represents a OwnershipTransferred event raised by the Linearstarrelease contract.
type LinearstarreleaseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Linearstarrelease *LinearstarreleaseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LinearstarreleaseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Linearstarrelease.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LinearstarreleaseOwnershipTransferredIterator{contract: _Linearstarrelease.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Linearstarrelease *LinearstarreleaseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LinearstarreleaseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Linearstarrelease.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LinearstarreleaseOwnershipTransferred)
				if err := _Linearstarrelease.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Linearstarrelease *LinearstarreleaseFilterer) ParseOwnershipTransferred(log types.Log) (*LinearstarreleaseOwnershipTransferred, error) {
	event := new(LinearstarreleaseOwnershipTransferred)
	if err := _Linearstarrelease.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
