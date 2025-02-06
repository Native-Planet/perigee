// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polls

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

// PollsMetaData contains all meta data concerning the Polls contract.
var PollsMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"pollDuration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposal\",\"type\":\"bytes32\"}],\"name\":\"updateDocumentPoll\",\"outputs\":[{\"name\":\"majority\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUpgradeProposals\",\"outputs\":[{\"name\":\"proposals\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"incrementTotalVoters\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDocumentMajorities\",\"outputs\":[{\"name\":\"majorities\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_as\",\"type\":\"uint8\"},{\"name\":\"_proposal\",\"type\":\"address\"},{\"name\":\"_vote\",\"type\":\"bool\"}],\"name\":\"castUpgradeVote\",\"outputs\":[{\"name\":\"majority\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposal\",\"type\":\"address\"}],\"name\":\"updateUpgradePoll\",\"outputs\":[{\"name\":\"majority\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_galaxy\",\"type\":\"uint8\"},{\"name\":\"_proposal\",\"type\":\"bytes32\"}],\"name\":\"hasVotedOnDocumentPoll\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_as\",\"type\":\"uint8\"},{\"name\":\"_proposal\",\"type\":\"bytes32\"},{\"name\":\"_vote\",\"type\":\"bool\"}],\"name\":\"castDocumentVote\",\"outputs\":[{\"name\":\"majority\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"documentHasAchievedMajority\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposal\",\"type\":\"bytes32\"}],\"name\":\"startDocumentPoll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"upgradeProposals\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"documentPolls\",\"outputs\":[{\"name\":\"start\",\"type\":\"uint256\"},{\"name\":\"yesVotes\",\"type\":\"uint16\"},{\"name\":\"noVotes\",\"type\":\"uint16\"},{\"name\":\"duration\",\"type\":\"uint256\"},{\"name\":\"cooldown\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"documentProposals\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUpgradeProposalCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDocumentProposalCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDocumentProposals\",\"outputs\":[{\"name\":\"proposals\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pollCooldown\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollDuration\",\"type\":\"uint256\"},{\"name\":\"_pollCooldown\",\"type\":\"uint256\"}],\"name\":\"reconfigure\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"upgradePolls\",\"outputs\":[{\"name\":\"start\",\"type\":\"uint256\"},{\"name\":\"yesVotes\",\"type\":\"uint16\"},{\"name\":\"noVotes\",\"type\":\"uint16\"},{\"name\":\"duration\",\"type\":\"uint256\"},{\"name\":\"cooldown\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"upgradeHasAchievedMajority\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposal\",\"type\":\"address\"}],\"name\":\"startUpgradePoll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"documentMajorities\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_galaxy\",\"type\":\"uint8\"},{\"name\":\"_proposal\",\"type\":\"address\"}],\"name\":\"hasVotedOnUpgradePoll\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_pollDuration\",\"type\":\"uint256\"},{\"name\":\"_pollCooldown\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"UpgradePollStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"proposal\",\"type\":\"bytes32\"}],\"name\":\"DocumentPollStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"UpgradeMajority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"proposal\",\"type\":\"bytes32\"}],\"name\":\"DocumentMajority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405160408062001f508339810180604052810190808051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506200009a8282620000a2640100000000026401000000009004565b505062000153565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515620000fe57600080fd5b8162069780111580156200011557506276a7008211155b8015620001255750806206978011155b80156200013557506276a7008111155b15156200014157600080fd5b81600181905550806002819055505050565b611ded80620001636000396000f30060806040526004361061016a576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806303a6930b1461016f578063172a735c1461019a5780631d075862146101e3578063246d41a91461024f57806334fe00b11461026657806351de5541146102d25780635623715b146103465780635f5300ff146103a157806369422924146103d45780636eb582241461042a578063715018a61461048c578063834e3822146104a35780638da5cb5b146104ec5780639a33aff9146105435780639fa4276914610574578063a2ebd2d5146105e1578063a46d41cf14610652578063a62459741461069b578063b953a9ac146106c6578063bc634abb146106f1578063d22687e41461075d578063da3d7d7f14610788578063dbbd96d2146107bf578063dc2b054e14610842578063e31f3e0c1461089d578063f2fde38b146108e0578063f740d10014610923578063f8abb93d1461096c575b600080fd5b34801561017b57600080fd5b506101846109d4565b6040518082815260200191505060405180910390f35b3480156101a657600080fd5b506101c960048036038101908080356000191690602001909291905050506109da565b604051808215151515815260200191505060405180910390f35b3480156101ef57600080fd5b506101f8610bc5565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561023b578082015181840152602081019050610220565b505050509050019250505060405180910390f35b34801561025b57600080fd5b50610264610c53565b005b34801561027257600080fd5b5061027b610d18565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156102be5780820151818401526020810190506102a3565b505050509050019250505060405180910390f35b3480156102de57600080fd5b5061032c600480360381019080803560ff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803515159060200190929190505050610d74565b604051808215151515815260200191505060405180910390f35b34801561035257600080fd5b50610387600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e31565b604051808215151515815260200191505060405180910390f35b3480156103ad57600080fd5b506103b66110d5565b604051808261ffff1661ffff16815260200191505060405180910390f35b3480156103e057600080fd5b50610410600480360381019080803560ff16906020019092919080356000191690602001909291905050506110e9565b604051808215151515815260200191505060405180910390f35b34801561043657600080fd5b50610472600480360381019080803560ff1690602001909291908035600019169060200190929190803515159060200190929190505050611138565b604051808215151515815260200191505060405180910390f35b34801561049857600080fd5b506104a16111d1565b005b3480156104af57600080fd5b506104d260048036038101908080356000191690602001909291905050506112d3565b604051808215151515815260200191505060405180910390f35b3480156104f857600080fd5b506105016112f3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561054f57600080fd5b506105726004803603810190808035600019169060200190929190505050611318565b005b34801561058057600080fd5b5061059f60048036038101908080359060200190929190505050611453565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156105ed57600080fd5b506106106004803603810190808035600019169060200190929190505050611491565b604051808681526020018561ffff1661ffff1681526020018461ffff1661ffff1681526020018381526020018281526020019550505050505060405180910390f35b34801561065e57600080fd5b5061067d600480360381019080803590602001909291905050506114e3565b60405180826000191660001916815260200191505060405180910390f35b3480156106a757600080fd5b506106b0611506565b6040518082815260200191505060405180910390f35b3480156106d257600080fd5b506106db611513565b6040518082815260200191505060405180910390f35b3480156106fd57600080fd5b50610706611520565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561074957808201518184015260208101905061072e565b505050509050019250505060405180910390f35b34801561076957600080fd5b5061077261157c565b6040518082815260200191505060405180910390f35b34801561079457600080fd5b506107bd6004803603810190808035906020019092919080359060200190929190505050611582565b005b3480156107cb57600080fd5b50610800600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061162e565b604051808681526020018561ffff1661ffff1681526020018461ffff1661ffff1681526020018381526020018281526020019550505050505060405180910390f35b34801561084e57600080fd5b50610883600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611680565b604051808215151515815260200191505060405180910390f35b3480156108a957600080fd5b506108de600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506116a0565b005b3480156108ec57600080fd5b50610921600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061187b565b005b34801561092f57600080fd5b5061094e600480360381019080803590602001909291905050506118e2565b60405180826000191660001916815260200191505060405180910390f35b34801561097857600080fd5b506109ba600480360381019080803560ff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611905565b604051808215151515815260200191505060405180910390f35b60015481565b60008060096000846000191660001916815260200190815260200160002060009054906101000a900460ff16151515610a1257600080fd5b6008600084600019166000191681526020019081526020016000209050610b0e8160c0604051908101604052908160008201548152602001600182016101008060200260405190810160405280929190826101008015610aac576020028201916000905b82829054906101000a900460ff16151581526020019060010190602082600001049283019260010382029150808411610a765790505b505050505081526020016009820160009054906101000a900461ffff1661ffff1661ffff1681526020016009820160029054906101000a900461ffff1661ffff1661ffff168152602001600a8201548152602001600b82015481525050611978565b91508115610bbc57600160096000856000191660001916815260200190815260200160002060006101000a81548160ff021916908315150217905550600a8390806001815401808255809150509060018203906000526020600020016000909192909190915090600019169055507fb11c3a231f918730636de6d815e478d63aead4aaaac95266c95b0278ff7183198360405180826000191660001916815260200191505060405180910390a15b81915050919050565b60606004805480602002602001604051908101604052809291908181526020018280548015610c4957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610bff575b5050505050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610cae57600080fd5b610100600360009054906101000a900461ffff1661ffff16101515610cd257600080fd5b610cfa6001600360009054906101000a900461ffff1661ffff16611a3190919063ffffffff16565b600360006101000a81548161ffff021916908361ffff160217905550565b6060600a805480602002602001604051908101604052809291908181526020018280548015610d6a57602002820191906000526020600020905b81546000191681526020019060010190808311610d52575b5050505050905090565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610dd257600080fd5b600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050610e1e818685611a57565b610e2784610e31565b9150509392505050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610e8f57600080fd5b600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151515610ee857600080fd5b600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506110088160c0604051908101604052908160008201548152602001600182016101008060200260405190810160405280929190826101008015610fa6576020028201916000905b82829054906101000a900460ff16151581526020019060010190602082600001049283019260010382029150808411610f705790505b505050505081526020016009820160009054906101000a900461ffff1661ffff1661ffff1681526020016009820160029054906101000a900461ffff1661ffff1661ffff168152602001600a8201548152602001600b82015481525050611978565b915081156110cc576001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f1ea969d04c4e4e370253bdeb7fb9638e50e75c206b4c42be36eb9844d37f444183604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a15b81915050919050565b600360009054906101000a900461ffff1681565b60006008600083600019166000191681526020019081526020016000206001018360ff166101008110151561111a57fe5b602091828204019190069054906101000a900460ff16905092915050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561119657600080fd5b60086000856000191660001916815260200190815260200160002090506111be818685611a57565b6111c7846109da565b9150509392505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561122c57600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60096020528060005260406000206000915054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561137557600080fd5b60096000836000191660001916815260200190815260200160002060009054906101000a900460ff161515156113aa57600080fd5b60086000836000191660001916815260200190815260200160002090508060000154600014156114075760078290806001815401808255809150509060018203906000526020600020016000909192909190915090600019169055505b61141081611b9a565b7fac68c1f0dc8aca8ee1bf445a687669d7d404b12dd0886f747e696115f00485278260405180826000191660001916815260200191505060405180910390a15050565b60048181548110151561146257fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60086020528060005260406000206000915090508060000154908060090160009054906101000a900461ffff16908060090160029054906101000a900461ffff169080600a01549080600b0154905085565b6007818154811015156114f257fe5b906000526020600020016000915090505481565b6000600480549050905090565b6000600780549050905090565b6060600780548060200260200160405190810160405280929190818152602001828054801561157257602002820191906000526020600020905b8154600019168152602001906001019080831161155a575b5050505050905090565b60025481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156115dd57600080fd5b8162069780111580156115f357506276a7008211155b80156116025750806206978011155b801561161157506276a7008111155b151561161c57600080fd5b81600181905550806002819055505050565b60056020528060005260406000206000915090508060000154908060090160009054906101000a900461ffff16908060090160029054906101000a900461ffff169080600a01549080600b0154905085565b60066020528060005260406000206000915054906101000a900460ff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156116fd57600080fd5b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151561175657600080fd5b600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905080600001546000141561180b5760048290806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b61181481611b9a565b7f6ffd05e058c57083a3d6fea058fbdbec172b55ce532b0f778f16e929a0ff516b82604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a15050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156118d657600080fd5b6118df81611c4a565b50565b600a818154811015156118f157fe5b906000526020600020016000915090505481565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018360ff166101008110151561195a57fe5b602091828204019190069054906101000a900460ff16905092915050565b60006004600360009054906101000a900461ffff1661ffff1681151561199a57fe5b0461ffff16826040015161ffff16101580156119c55750816060015161ffff16826040015161ffff16115b8015611a2a57506119e782608001518360000151611d4490919063ffffffff16565b421180611a295750611a1a8260400151600360009054906101000a900461ffff1661ffff16611d6090919063ffffffff16565b61ffff16826040015161ffff16115b5b9050919050565b60008082840190508361ffff168161ffff1610151515611a4d57fe5b8091505092915050565b82600001544210151515611a6757fe5b826001018260ff1661010081101515611a7c57fe5b602091828204019190069054906101000a900460ff16158015611ab85750611ab583600a01548460000154611d4490919063ffffffff16565b42105b1515611ac357600080fd5b6001836001018360ff1661010081101515611ada57fe5b602091828204019190066101000a81548160ff0219169083151502179055508015611b4c57611b2960018460090160009054906101000a900461ffff1661ffff16611a3190919063ffffffff16565b8360090160006101000a81548161ffff021916908361ffff160217905550611b95565b611b7660018460090160029054906101000a900461ffff1661ffff16611a3190919063ffffffff16565b8360090160026101000a81548161ffff021916908361ffff1602179055505b505050565b611bcb611bb882600b015483600a0154611d4490919063ffffffff16565b8260000154611d4490919063ffffffff16565b42111515611bd857600080fd5b428160000181905550806001016000611bf19190611d81565b60008160090160006101000a81548161ffff021916908361ffff16021790555060008160090160026101000a81548161ffff021916908361ffff16021790555060015481600a018190555060025481600b018190555050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515611c8657600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008183019050828110151515611d5757fe5b80905092915050565b60008261ffff168261ffff1611151515611d7657fe5b818303905092915050565b5080610100601f01602090040190611d999190611d9c565b50565b611dbe91905b80821115611dba576000816000905550600101611da2565b5090565b905600a165627a7a7230582068bd58209af89cb15753035fa222cb6b96fd5f7e3c95b00ecf2e0093ce2925bc0029",
}

// PollsABI is the input ABI used to generate the binding from.
// Deprecated: Use PollsMetaData.ABI instead.
var PollsABI = PollsMetaData.ABI

// PollsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PollsMetaData.Bin instead.
var PollsBin = PollsMetaData.Bin

// DeployPolls deploys a new Ethereum contract, binding an instance of Polls to it.
func DeployPolls(auth *bind.TransactOpts, backend bind.ContractBackend, _pollDuration *big.Int, _pollCooldown *big.Int) (common.Address, *types.Transaction, *Polls, error) {
	parsed, err := PollsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PollsBin), backend, _pollDuration, _pollCooldown)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polls{PollsCaller: PollsCaller{contract: contract}, PollsTransactor: PollsTransactor{contract: contract}, PollsFilterer: PollsFilterer{contract: contract}}, nil
}

// Polls is an auto generated Go binding around an Ethereum contract.
type Polls struct {
	PollsCaller     // Read-only binding to the contract
	PollsTransactor // Write-only binding to the contract
	PollsFilterer   // Log filterer for contract events
}

// PollsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PollsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PollsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PollsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PollsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PollsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PollsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PollsSession struct {
	Contract     *Polls            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PollsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PollsCallerSession struct {
	Contract *PollsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PollsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PollsTransactorSession struct {
	Contract     *PollsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PollsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PollsRaw struct {
	Contract *Polls // Generic contract binding to access the raw methods on
}

// PollsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PollsCallerRaw struct {
	Contract *PollsCaller // Generic read-only contract binding to access the raw methods on
}

// PollsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PollsTransactorRaw struct {
	Contract *PollsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolls creates a new instance of Polls, bound to a specific deployed contract.
func NewPolls(address common.Address, backend bind.ContractBackend) (*Polls, error) {
	contract, err := bindPolls(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polls{PollsCaller: PollsCaller{contract: contract}, PollsTransactor: PollsTransactor{contract: contract}, PollsFilterer: PollsFilterer{contract: contract}}, nil
}

// NewPollsCaller creates a new read-only instance of Polls, bound to a specific deployed contract.
func NewPollsCaller(address common.Address, caller bind.ContractCaller) (*PollsCaller, error) {
	contract, err := bindPolls(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PollsCaller{contract: contract}, nil
}

// NewPollsTransactor creates a new write-only instance of Polls, bound to a specific deployed contract.
func NewPollsTransactor(address common.Address, transactor bind.ContractTransactor) (*PollsTransactor, error) {
	contract, err := bindPolls(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PollsTransactor{contract: contract}, nil
}

// NewPollsFilterer creates a new log filterer instance of Polls, bound to a specific deployed contract.
func NewPollsFilterer(address common.Address, filterer bind.ContractFilterer) (*PollsFilterer, error) {
	contract, err := bindPolls(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PollsFilterer{contract: contract}, nil
}

// bindPolls binds a generic wrapper to an already deployed contract.
func bindPolls(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PollsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polls *PollsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polls.Contract.PollsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polls *PollsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polls.Contract.PollsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polls *PollsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polls.Contract.PollsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polls *PollsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polls.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polls *PollsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polls.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polls *PollsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polls.Contract.contract.Transact(opts, method, params...)
}

// DocumentHasAchievedMajority is a free data retrieval call binding the contract method 0x834e3822.
//
// Solidity: function documentHasAchievedMajority(bytes32 ) view returns(bool)
func (_Polls *PollsCaller) DocumentHasAchievedMajority(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "documentHasAchievedMajority", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DocumentHasAchievedMajority is a free data retrieval call binding the contract method 0x834e3822.
//
// Solidity: function documentHasAchievedMajority(bytes32 ) view returns(bool)
func (_Polls *PollsSession) DocumentHasAchievedMajority(arg0 [32]byte) (bool, error) {
	return _Polls.Contract.DocumentHasAchievedMajority(&_Polls.CallOpts, arg0)
}

// DocumentHasAchievedMajority is a free data retrieval call binding the contract method 0x834e3822.
//
// Solidity: function documentHasAchievedMajority(bytes32 ) view returns(bool)
func (_Polls *PollsCallerSession) DocumentHasAchievedMajority(arg0 [32]byte) (bool, error) {
	return _Polls.Contract.DocumentHasAchievedMajority(&_Polls.CallOpts, arg0)
}

// DocumentMajorities is a free data retrieval call binding the contract method 0xf740d100.
//
// Solidity: function documentMajorities(uint256 ) view returns(bytes32)
func (_Polls *PollsCaller) DocumentMajorities(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "documentMajorities", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DocumentMajorities is a free data retrieval call binding the contract method 0xf740d100.
//
// Solidity: function documentMajorities(uint256 ) view returns(bytes32)
func (_Polls *PollsSession) DocumentMajorities(arg0 *big.Int) ([32]byte, error) {
	return _Polls.Contract.DocumentMajorities(&_Polls.CallOpts, arg0)
}

// DocumentMajorities is a free data retrieval call binding the contract method 0xf740d100.
//
// Solidity: function documentMajorities(uint256 ) view returns(bytes32)
func (_Polls *PollsCallerSession) DocumentMajorities(arg0 *big.Int) ([32]byte, error) {
	return _Polls.Contract.DocumentMajorities(&_Polls.CallOpts, arg0)
}

// DocumentPolls is a free data retrieval call binding the contract method 0xa2ebd2d5.
//
// Solidity: function documentPolls(bytes32 ) view returns(uint256 start, uint16 yesVotes, uint16 noVotes, uint256 duration, uint256 cooldown)
func (_Polls *PollsCaller) DocumentPolls(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Start    *big.Int
	YesVotes uint16
	NoVotes  uint16
	Duration *big.Int
	Cooldown *big.Int
}, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "documentPolls", arg0)

	outstruct := new(struct {
		Start    *big.Int
		YesVotes uint16
		NoVotes  uint16
		Duration *big.Int
		Cooldown *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Start = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.YesVotes = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.NoVotes = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.Duration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Cooldown = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DocumentPolls is a free data retrieval call binding the contract method 0xa2ebd2d5.
//
// Solidity: function documentPolls(bytes32 ) view returns(uint256 start, uint16 yesVotes, uint16 noVotes, uint256 duration, uint256 cooldown)
func (_Polls *PollsSession) DocumentPolls(arg0 [32]byte) (struct {
	Start    *big.Int
	YesVotes uint16
	NoVotes  uint16
	Duration *big.Int
	Cooldown *big.Int
}, error) {
	return _Polls.Contract.DocumentPolls(&_Polls.CallOpts, arg0)
}

// DocumentPolls is a free data retrieval call binding the contract method 0xa2ebd2d5.
//
// Solidity: function documentPolls(bytes32 ) view returns(uint256 start, uint16 yesVotes, uint16 noVotes, uint256 duration, uint256 cooldown)
func (_Polls *PollsCallerSession) DocumentPolls(arg0 [32]byte) (struct {
	Start    *big.Int
	YesVotes uint16
	NoVotes  uint16
	Duration *big.Int
	Cooldown *big.Int
}, error) {
	return _Polls.Contract.DocumentPolls(&_Polls.CallOpts, arg0)
}

// DocumentProposals is a free data retrieval call binding the contract method 0xa46d41cf.
//
// Solidity: function documentProposals(uint256 ) view returns(bytes32)
func (_Polls *PollsCaller) DocumentProposals(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "documentProposals", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DocumentProposals is a free data retrieval call binding the contract method 0xa46d41cf.
//
// Solidity: function documentProposals(uint256 ) view returns(bytes32)
func (_Polls *PollsSession) DocumentProposals(arg0 *big.Int) ([32]byte, error) {
	return _Polls.Contract.DocumentProposals(&_Polls.CallOpts, arg0)
}

// DocumentProposals is a free data retrieval call binding the contract method 0xa46d41cf.
//
// Solidity: function documentProposals(uint256 ) view returns(bytes32)
func (_Polls *PollsCallerSession) DocumentProposals(arg0 *big.Int) ([32]byte, error) {
	return _Polls.Contract.DocumentProposals(&_Polls.CallOpts, arg0)
}

// GetDocumentMajorities is a free data retrieval call binding the contract method 0x34fe00b1.
//
// Solidity: function getDocumentMajorities() view returns(bytes32[] majorities)
func (_Polls *PollsCaller) GetDocumentMajorities(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "getDocumentMajorities")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetDocumentMajorities is a free data retrieval call binding the contract method 0x34fe00b1.
//
// Solidity: function getDocumentMajorities() view returns(bytes32[] majorities)
func (_Polls *PollsSession) GetDocumentMajorities() ([][32]byte, error) {
	return _Polls.Contract.GetDocumentMajorities(&_Polls.CallOpts)
}

// GetDocumentMajorities is a free data retrieval call binding the contract method 0x34fe00b1.
//
// Solidity: function getDocumentMajorities() view returns(bytes32[] majorities)
func (_Polls *PollsCallerSession) GetDocumentMajorities() ([][32]byte, error) {
	return _Polls.Contract.GetDocumentMajorities(&_Polls.CallOpts)
}

// GetDocumentProposalCount is a free data retrieval call binding the contract method 0xb953a9ac.
//
// Solidity: function getDocumentProposalCount() view returns(uint256 count)
func (_Polls *PollsCaller) GetDocumentProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "getDocumentProposalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDocumentProposalCount is a free data retrieval call binding the contract method 0xb953a9ac.
//
// Solidity: function getDocumentProposalCount() view returns(uint256 count)
func (_Polls *PollsSession) GetDocumentProposalCount() (*big.Int, error) {
	return _Polls.Contract.GetDocumentProposalCount(&_Polls.CallOpts)
}

// GetDocumentProposalCount is a free data retrieval call binding the contract method 0xb953a9ac.
//
// Solidity: function getDocumentProposalCount() view returns(uint256 count)
func (_Polls *PollsCallerSession) GetDocumentProposalCount() (*big.Int, error) {
	return _Polls.Contract.GetDocumentProposalCount(&_Polls.CallOpts)
}

// GetDocumentProposals is a free data retrieval call binding the contract method 0xbc634abb.
//
// Solidity: function getDocumentProposals() view returns(bytes32[] proposals)
func (_Polls *PollsCaller) GetDocumentProposals(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "getDocumentProposals")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetDocumentProposals is a free data retrieval call binding the contract method 0xbc634abb.
//
// Solidity: function getDocumentProposals() view returns(bytes32[] proposals)
func (_Polls *PollsSession) GetDocumentProposals() ([][32]byte, error) {
	return _Polls.Contract.GetDocumentProposals(&_Polls.CallOpts)
}

// GetDocumentProposals is a free data retrieval call binding the contract method 0xbc634abb.
//
// Solidity: function getDocumentProposals() view returns(bytes32[] proposals)
func (_Polls *PollsCallerSession) GetDocumentProposals() ([][32]byte, error) {
	return _Polls.Contract.GetDocumentProposals(&_Polls.CallOpts)
}

// GetUpgradeProposalCount is a free data retrieval call binding the contract method 0xa6245974.
//
// Solidity: function getUpgradeProposalCount() view returns(uint256 count)
func (_Polls *PollsCaller) GetUpgradeProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "getUpgradeProposalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUpgradeProposalCount is a free data retrieval call binding the contract method 0xa6245974.
//
// Solidity: function getUpgradeProposalCount() view returns(uint256 count)
func (_Polls *PollsSession) GetUpgradeProposalCount() (*big.Int, error) {
	return _Polls.Contract.GetUpgradeProposalCount(&_Polls.CallOpts)
}

// GetUpgradeProposalCount is a free data retrieval call binding the contract method 0xa6245974.
//
// Solidity: function getUpgradeProposalCount() view returns(uint256 count)
func (_Polls *PollsCallerSession) GetUpgradeProposalCount() (*big.Int, error) {
	return _Polls.Contract.GetUpgradeProposalCount(&_Polls.CallOpts)
}

// GetUpgradeProposals is a free data retrieval call binding the contract method 0x1d075862.
//
// Solidity: function getUpgradeProposals() view returns(address[] proposals)
func (_Polls *PollsCaller) GetUpgradeProposals(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "getUpgradeProposals")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUpgradeProposals is a free data retrieval call binding the contract method 0x1d075862.
//
// Solidity: function getUpgradeProposals() view returns(address[] proposals)
func (_Polls *PollsSession) GetUpgradeProposals() ([]common.Address, error) {
	return _Polls.Contract.GetUpgradeProposals(&_Polls.CallOpts)
}

// GetUpgradeProposals is a free data retrieval call binding the contract method 0x1d075862.
//
// Solidity: function getUpgradeProposals() view returns(address[] proposals)
func (_Polls *PollsCallerSession) GetUpgradeProposals() ([]common.Address, error) {
	return _Polls.Contract.GetUpgradeProposals(&_Polls.CallOpts)
}

// HasVotedOnDocumentPoll is a free data retrieval call binding the contract method 0x69422924.
//
// Solidity: function hasVotedOnDocumentPoll(uint8 _galaxy, bytes32 _proposal) view returns(bool result)
func (_Polls *PollsCaller) HasVotedOnDocumentPoll(opts *bind.CallOpts, _galaxy uint8, _proposal [32]byte) (bool, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "hasVotedOnDocumentPoll", _galaxy, _proposal)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVotedOnDocumentPoll is a free data retrieval call binding the contract method 0x69422924.
//
// Solidity: function hasVotedOnDocumentPoll(uint8 _galaxy, bytes32 _proposal) view returns(bool result)
func (_Polls *PollsSession) HasVotedOnDocumentPoll(_galaxy uint8, _proposal [32]byte) (bool, error) {
	return _Polls.Contract.HasVotedOnDocumentPoll(&_Polls.CallOpts, _galaxy, _proposal)
}

// HasVotedOnDocumentPoll is a free data retrieval call binding the contract method 0x69422924.
//
// Solidity: function hasVotedOnDocumentPoll(uint8 _galaxy, bytes32 _proposal) view returns(bool result)
func (_Polls *PollsCallerSession) HasVotedOnDocumentPoll(_galaxy uint8, _proposal [32]byte) (bool, error) {
	return _Polls.Contract.HasVotedOnDocumentPoll(&_Polls.CallOpts, _galaxy, _proposal)
}

// HasVotedOnUpgradePoll is a free data retrieval call binding the contract method 0xf8abb93d.
//
// Solidity: function hasVotedOnUpgradePoll(uint8 _galaxy, address _proposal) view returns(bool result)
func (_Polls *PollsCaller) HasVotedOnUpgradePoll(opts *bind.CallOpts, _galaxy uint8, _proposal common.Address) (bool, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "hasVotedOnUpgradePoll", _galaxy, _proposal)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVotedOnUpgradePoll is a free data retrieval call binding the contract method 0xf8abb93d.
//
// Solidity: function hasVotedOnUpgradePoll(uint8 _galaxy, address _proposal) view returns(bool result)
func (_Polls *PollsSession) HasVotedOnUpgradePoll(_galaxy uint8, _proposal common.Address) (bool, error) {
	return _Polls.Contract.HasVotedOnUpgradePoll(&_Polls.CallOpts, _galaxy, _proposal)
}

// HasVotedOnUpgradePoll is a free data retrieval call binding the contract method 0xf8abb93d.
//
// Solidity: function hasVotedOnUpgradePoll(uint8 _galaxy, address _proposal) view returns(bool result)
func (_Polls *PollsCallerSession) HasVotedOnUpgradePoll(_galaxy uint8, _proposal common.Address) (bool, error) {
	return _Polls.Contract.HasVotedOnUpgradePoll(&_Polls.CallOpts, _galaxy, _proposal)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Polls *PollsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Polls *PollsSession) Owner() (common.Address, error) {
	return _Polls.Contract.Owner(&_Polls.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Polls *PollsCallerSession) Owner() (common.Address, error) {
	return _Polls.Contract.Owner(&_Polls.CallOpts)
}

// PollCooldown is a free data retrieval call binding the contract method 0xd22687e4.
//
// Solidity: function pollCooldown() view returns(uint256)
func (_Polls *PollsCaller) PollCooldown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "pollCooldown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PollCooldown is a free data retrieval call binding the contract method 0xd22687e4.
//
// Solidity: function pollCooldown() view returns(uint256)
func (_Polls *PollsSession) PollCooldown() (*big.Int, error) {
	return _Polls.Contract.PollCooldown(&_Polls.CallOpts)
}

// PollCooldown is a free data retrieval call binding the contract method 0xd22687e4.
//
// Solidity: function pollCooldown() view returns(uint256)
func (_Polls *PollsCallerSession) PollCooldown() (*big.Int, error) {
	return _Polls.Contract.PollCooldown(&_Polls.CallOpts)
}

// PollDuration is a free data retrieval call binding the contract method 0x03a6930b.
//
// Solidity: function pollDuration() view returns(uint256)
func (_Polls *PollsCaller) PollDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "pollDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PollDuration is a free data retrieval call binding the contract method 0x03a6930b.
//
// Solidity: function pollDuration() view returns(uint256)
func (_Polls *PollsSession) PollDuration() (*big.Int, error) {
	return _Polls.Contract.PollDuration(&_Polls.CallOpts)
}

// PollDuration is a free data retrieval call binding the contract method 0x03a6930b.
//
// Solidity: function pollDuration() view returns(uint256)
func (_Polls *PollsCallerSession) PollDuration() (*big.Int, error) {
	return _Polls.Contract.PollDuration(&_Polls.CallOpts)
}

// TotalVoters is a free data retrieval call binding the contract method 0x5f5300ff.
//
// Solidity: function totalVoters() view returns(uint16)
func (_Polls *PollsCaller) TotalVoters(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "totalVoters")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TotalVoters is a free data retrieval call binding the contract method 0x5f5300ff.
//
// Solidity: function totalVoters() view returns(uint16)
func (_Polls *PollsSession) TotalVoters() (uint16, error) {
	return _Polls.Contract.TotalVoters(&_Polls.CallOpts)
}

// TotalVoters is a free data retrieval call binding the contract method 0x5f5300ff.
//
// Solidity: function totalVoters() view returns(uint16)
func (_Polls *PollsCallerSession) TotalVoters() (uint16, error) {
	return _Polls.Contract.TotalVoters(&_Polls.CallOpts)
}

// UpgradeHasAchievedMajority is a free data retrieval call binding the contract method 0xdc2b054e.
//
// Solidity: function upgradeHasAchievedMajority(address ) view returns(bool)
func (_Polls *PollsCaller) UpgradeHasAchievedMajority(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "upgradeHasAchievedMajority", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UpgradeHasAchievedMajority is a free data retrieval call binding the contract method 0xdc2b054e.
//
// Solidity: function upgradeHasAchievedMajority(address ) view returns(bool)
func (_Polls *PollsSession) UpgradeHasAchievedMajority(arg0 common.Address) (bool, error) {
	return _Polls.Contract.UpgradeHasAchievedMajority(&_Polls.CallOpts, arg0)
}

// UpgradeHasAchievedMajority is a free data retrieval call binding the contract method 0xdc2b054e.
//
// Solidity: function upgradeHasAchievedMajority(address ) view returns(bool)
func (_Polls *PollsCallerSession) UpgradeHasAchievedMajority(arg0 common.Address) (bool, error) {
	return _Polls.Contract.UpgradeHasAchievedMajority(&_Polls.CallOpts, arg0)
}

// UpgradePolls is a free data retrieval call binding the contract method 0xdbbd96d2.
//
// Solidity: function upgradePolls(address ) view returns(uint256 start, uint16 yesVotes, uint16 noVotes, uint256 duration, uint256 cooldown)
func (_Polls *PollsCaller) UpgradePolls(opts *bind.CallOpts, arg0 common.Address) (struct {
	Start    *big.Int
	YesVotes uint16
	NoVotes  uint16
	Duration *big.Int
	Cooldown *big.Int
}, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "upgradePolls", arg0)

	outstruct := new(struct {
		Start    *big.Int
		YesVotes uint16
		NoVotes  uint16
		Duration *big.Int
		Cooldown *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Start = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.YesVotes = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.NoVotes = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.Duration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Cooldown = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UpgradePolls is a free data retrieval call binding the contract method 0xdbbd96d2.
//
// Solidity: function upgradePolls(address ) view returns(uint256 start, uint16 yesVotes, uint16 noVotes, uint256 duration, uint256 cooldown)
func (_Polls *PollsSession) UpgradePolls(arg0 common.Address) (struct {
	Start    *big.Int
	YesVotes uint16
	NoVotes  uint16
	Duration *big.Int
	Cooldown *big.Int
}, error) {
	return _Polls.Contract.UpgradePolls(&_Polls.CallOpts, arg0)
}

// UpgradePolls is a free data retrieval call binding the contract method 0xdbbd96d2.
//
// Solidity: function upgradePolls(address ) view returns(uint256 start, uint16 yesVotes, uint16 noVotes, uint256 duration, uint256 cooldown)
func (_Polls *PollsCallerSession) UpgradePolls(arg0 common.Address) (struct {
	Start    *big.Int
	YesVotes uint16
	NoVotes  uint16
	Duration *big.Int
	Cooldown *big.Int
}, error) {
	return _Polls.Contract.UpgradePolls(&_Polls.CallOpts, arg0)
}

// UpgradeProposals is a free data retrieval call binding the contract method 0x9fa42769.
//
// Solidity: function upgradeProposals(uint256 ) view returns(address)
func (_Polls *PollsCaller) UpgradeProposals(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Polls.contract.Call(opts, &out, "upgradeProposals", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradeProposals is a free data retrieval call binding the contract method 0x9fa42769.
//
// Solidity: function upgradeProposals(uint256 ) view returns(address)
func (_Polls *PollsSession) UpgradeProposals(arg0 *big.Int) (common.Address, error) {
	return _Polls.Contract.UpgradeProposals(&_Polls.CallOpts, arg0)
}

// UpgradeProposals is a free data retrieval call binding the contract method 0x9fa42769.
//
// Solidity: function upgradeProposals(uint256 ) view returns(address)
func (_Polls *PollsCallerSession) UpgradeProposals(arg0 *big.Int) (common.Address, error) {
	return _Polls.Contract.UpgradeProposals(&_Polls.CallOpts, arg0)
}

// CastDocumentVote is a paid mutator transaction binding the contract method 0x6eb58224.
//
// Solidity: function castDocumentVote(uint8 _as, bytes32 _proposal, bool _vote) returns(bool majority)
func (_Polls *PollsTransactor) CastDocumentVote(opts *bind.TransactOpts, _as uint8, _proposal [32]byte, _vote bool) (*types.Transaction, error) {
	return _Polls.contract.Transact(opts, "castDocumentVote", _as, _proposal, _vote)
}

// CastDocumentVote is a paid mutator transaction binding the contract method 0x6eb58224.
//
// Solidity: function castDocumentVote(uint8 _as, bytes32 _proposal, bool _vote) returns(bool majority)
func (_Polls *PollsSession) CastDocumentVote(_as uint8, _proposal [32]byte, _vote bool) (*types.Transaction, error) {
	return _Polls.Contract.CastDocumentVote(&_Polls.TransactOpts, _as, _proposal, _vote)
}

// CastDocumentVote is a paid mutator transaction binding the contract method 0x6eb58224.
//
// Solidity: function castDocumentVote(uint8 _as, bytes32 _proposal, bool _vote) returns(bool majority)
func (_Polls *PollsTransactorSession) CastDocumentVote(_as uint8, _proposal [32]byte, _vote bool) (*types.Transaction, error) {
	return _Polls.Contract.CastDocumentVote(&_Polls.TransactOpts, _as, _proposal, _vote)
}

// CastUpgradeVote is a paid mutator transaction binding the contract method 0x51de5541.
//
// Solidity: function castUpgradeVote(uint8 _as, address _proposal, bool _vote) returns(bool majority)
func (_Polls *PollsTransactor) CastUpgradeVote(opts *bind.TransactOpts, _as uint8, _proposal common.Address, _vote bool) (*types.Transaction, error) {
	return _Polls.contract.Transact(opts, "castUpgradeVote", _as, _proposal, _vote)
}

// CastUpgradeVote is a paid mutator transaction binding the contract method 0x51de5541.
//
// Solidity: function castUpgradeVote(uint8 _as, address _proposal, bool _vote) returns(bool majority)
func (_Polls *PollsSession) CastUpgradeVote(_as uint8, _proposal common.Address, _vote bool) (*types.Transaction, error) {
	return _Polls.Contract.CastUpgradeVote(&_Polls.TransactOpts, _as, _proposal, _vote)
}

// CastUpgradeVote is a paid mutator transaction binding the contract method 0x51de5541.
//
// Solidity: function castUpgradeVote(uint8 _as, address _proposal, bool _vote) returns(bool majority)
func (_Polls *PollsTransactorSession) CastUpgradeVote(_as uint8, _proposal common.Address, _vote bool) (*types.Transaction, error) {
	return _Polls.Contract.CastUpgradeVote(&_Polls.TransactOpts, _as, _proposal, _vote)
}

// IncrementTotalVoters is a paid mutator transaction binding the contract method 0x246d41a9.
//
// Solidity: function incrementTotalVoters() returns()
func (_Polls *PollsTransactor) IncrementTotalVoters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polls.contract.Transact(opts, "incrementTotalVoters")
}

// IncrementTotalVoters is a paid mutator transaction binding the contract method 0x246d41a9.
//
// Solidity: function incrementTotalVoters() returns()
func (_Polls *PollsSession) IncrementTotalVoters() (*types.Transaction, error) {
	return _Polls.Contract.IncrementTotalVoters(&_Polls.TransactOpts)
}

// IncrementTotalVoters is a paid mutator transaction binding the contract method 0x246d41a9.
//
// Solidity: function incrementTotalVoters() returns()
func (_Polls *PollsTransactorSession) IncrementTotalVoters() (*types.Transaction, error) {
	return _Polls.Contract.IncrementTotalVoters(&_Polls.TransactOpts)
}

// Reconfigure is a paid mutator transaction binding the contract method 0xda3d7d7f.
//
// Solidity: function reconfigure(uint256 _pollDuration, uint256 _pollCooldown) returns()
func (_Polls *PollsTransactor) Reconfigure(opts *bind.TransactOpts, _pollDuration *big.Int, _pollCooldown *big.Int) (*types.Transaction, error) {
	return _Polls.contract.Transact(opts, "reconfigure", _pollDuration, _pollCooldown)
}

// Reconfigure is a paid mutator transaction binding the contract method 0xda3d7d7f.
//
// Solidity: function reconfigure(uint256 _pollDuration, uint256 _pollCooldown) returns()
func (_Polls *PollsSession) Reconfigure(_pollDuration *big.Int, _pollCooldown *big.Int) (*types.Transaction, error) {
	return _Polls.Contract.Reconfigure(&_Polls.TransactOpts, _pollDuration, _pollCooldown)
}

// Reconfigure is a paid mutator transaction binding the contract method 0xda3d7d7f.
//
// Solidity: function reconfigure(uint256 _pollDuration, uint256 _pollCooldown) returns()
func (_Polls *PollsTransactorSession) Reconfigure(_pollDuration *big.Int, _pollCooldown *big.Int) (*types.Transaction, error) {
	return _Polls.Contract.Reconfigure(&_Polls.TransactOpts, _pollDuration, _pollCooldown)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Polls *PollsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polls.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Polls *PollsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Polls.Contract.RenounceOwnership(&_Polls.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Polls *PollsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Polls.Contract.RenounceOwnership(&_Polls.TransactOpts)
}

// StartDocumentPoll is a paid mutator transaction binding the contract method 0x9a33aff9.
//
// Solidity: function startDocumentPoll(bytes32 _proposal) returns()
func (_Polls *PollsTransactor) StartDocumentPoll(opts *bind.TransactOpts, _proposal [32]byte) (*types.Transaction, error) {
	return _Polls.contract.Transact(opts, "startDocumentPoll", _proposal)
}

// StartDocumentPoll is a paid mutator transaction binding the contract method 0x9a33aff9.
//
// Solidity: function startDocumentPoll(bytes32 _proposal) returns()
func (_Polls *PollsSession) StartDocumentPoll(_proposal [32]byte) (*types.Transaction, error) {
	return _Polls.Contract.StartDocumentPoll(&_Polls.TransactOpts, _proposal)
}

// StartDocumentPoll is a paid mutator transaction binding the contract method 0x9a33aff9.
//
// Solidity: function startDocumentPoll(bytes32 _proposal) returns()
func (_Polls *PollsTransactorSession) StartDocumentPoll(_proposal [32]byte) (*types.Transaction, error) {
	return _Polls.Contract.StartDocumentPoll(&_Polls.TransactOpts, _proposal)
}

// StartUpgradePoll is a paid mutator transaction binding the contract method 0xe31f3e0c.
//
// Solidity: function startUpgradePoll(address _proposal) returns()
func (_Polls *PollsTransactor) StartUpgradePoll(opts *bind.TransactOpts, _proposal common.Address) (*types.Transaction, error) {
	return _Polls.contract.Transact(opts, "startUpgradePoll", _proposal)
}

// StartUpgradePoll is a paid mutator transaction binding the contract method 0xe31f3e0c.
//
// Solidity: function startUpgradePoll(address _proposal) returns()
func (_Polls *PollsSession) StartUpgradePoll(_proposal common.Address) (*types.Transaction, error) {
	return _Polls.Contract.StartUpgradePoll(&_Polls.TransactOpts, _proposal)
}

// StartUpgradePoll is a paid mutator transaction binding the contract method 0xe31f3e0c.
//
// Solidity: function startUpgradePoll(address _proposal) returns()
func (_Polls *PollsTransactorSession) StartUpgradePoll(_proposal common.Address) (*types.Transaction, error) {
	return _Polls.Contract.StartUpgradePoll(&_Polls.TransactOpts, _proposal)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Polls *PollsTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Polls.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Polls *PollsSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Polls.Contract.TransferOwnership(&_Polls.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Polls *PollsTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Polls.Contract.TransferOwnership(&_Polls.TransactOpts, _newOwner)
}

// UpdateDocumentPoll is a paid mutator transaction binding the contract method 0x172a735c.
//
// Solidity: function updateDocumentPoll(bytes32 _proposal) returns(bool majority)
func (_Polls *PollsTransactor) UpdateDocumentPoll(opts *bind.TransactOpts, _proposal [32]byte) (*types.Transaction, error) {
	return _Polls.contract.Transact(opts, "updateDocumentPoll", _proposal)
}

// UpdateDocumentPoll is a paid mutator transaction binding the contract method 0x172a735c.
//
// Solidity: function updateDocumentPoll(bytes32 _proposal) returns(bool majority)
func (_Polls *PollsSession) UpdateDocumentPoll(_proposal [32]byte) (*types.Transaction, error) {
	return _Polls.Contract.UpdateDocumentPoll(&_Polls.TransactOpts, _proposal)
}

// UpdateDocumentPoll is a paid mutator transaction binding the contract method 0x172a735c.
//
// Solidity: function updateDocumentPoll(bytes32 _proposal) returns(bool majority)
func (_Polls *PollsTransactorSession) UpdateDocumentPoll(_proposal [32]byte) (*types.Transaction, error) {
	return _Polls.Contract.UpdateDocumentPoll(&_Polls.TransactOpts, _proposal)
}

// UpdateUpgradePoll is a paid mutator transaction binding the contract method 0x5623715b.
//
// Solidity: function updateUpgradePoll(address _proposal) returns(bool majority)
func (_Polls *PollsTransactor) UpdateUpgradePoll(opts *bind.TransactOpts, _proposal common.Address) (*types.Transaction, error) {
	return _Polls.contract.Transact(opts, "updateUpgradePoll", _proposal)
}

// UpdateUpgradePoll is a paid mutator transaction binding the contract method 0x5623715b.
//
// Solidity: function updateUpgradePoll(address _proposal) returns(bool majority)
func (_Polls *PollsSession) UpdateUpgradePoll(_proposal common.Address) (*types.Transaction, error) {
	return _Polls.Contract.UpdateUpgradePoll(&_Polls.TransactOpts, _proposal)
}

// UpdateUpgradePoll is a paid mutator transaction binding the contract method 0x5623715b.
//
// Solidity: function updateUpgradePoll(address _proposal) returns(bool majority)
func (_Polls *PollsTransactorSession) UpdateUpgradePoll(_proposal common.Address) (*types.Transaction, error) {
	return _Polls.Contract.UpdateUpgradePoll(&_Polls.TransactOpts, _proposal)
}

// PollsDocumentMajorityIterator is returned from FilterDocumentMajority and is used to iterate over the raw logs and unpacked data for DocumentMajority events raised by the Polls contract.
type PollsDocumentMajorityIterator struct {
	Event *PollsDocumentMajority // Event containing the contract specifics and raw log

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
func (it *PollsDocumentMajorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PollsDocumentMajority)
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
		it.Event = new(PollsDocumentMajority)
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
func (it *PollsDocumentMajorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PollsDocumentMajorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PollsDocumentMajority represents a DocumentMajority event raised by the Polls contract.
type PollsDocumentMajority struct {
	Proposal [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDocumentMajority is a free log retrieval operation binding the contract event 0xb11c3a231f918730636de6d815e478d63aead4aaaac95266c95b0278ff718319.
//
// Solidity: event DocumentMajority(bytes32 proposal)
func (_Polls *PollsFilterer) FilterDocumentMajority(opts *bind.FilterOpts) (*PollsDocumentMajorityIterator, error) {

	logs, sub, err := _Polls.contract.FilterLogs(opts, "DocumentMajority")
	if err != nil {
		return nil, err
	}
	return &PollsDocumentMajorityIterator{contract: _Polls.contract, event: "DocumentMajority", logs: logs, sub: sub}, nil
}

// WatchDocumentMajority is a free log subscription operation binding the contract event 0xb11c3a231f918730636de6d815e478d63aead4aaaac95266c95b0278ff718319.
//
// Solidity: event DocumentMajority(bytes32 proposal)
func (_Polls *PollsFilterer) WatchDocumentMajority(opts *bind.WatchOpts, sink chan<- *PollsDocumentMajority) (event.Subscription, error) {

	logs, sub, err := _Polls.contract.WatchLogs(opts, "DocumentMajority")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PollsDocumentMajority)
				if err := _Polls.contract.UnpackLog(event, "DocumentMajority", log); err != nil {
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

// ParseDocumentMajority is a log parse operation binding the contract event 0xb11c3a231f918730636de6d815e478d63aead4aaaac95266c95b0278ff718319.
//
// Solidity: event DocumentMajority(bytes32 proposal)
func (_Polls *PollsFilterer) ParseDocumentMajority(log types.Log) (*PollsDocumentMajority, error) {
	event := new(PollsDocumentMajority)
	if err := _Polls.contract.UnpackLog(event, "DocumentMajority", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PollsDocumentPollStartedIterator is returned from FilterDocumentPollStarted and is used to iterate over the raw logs and unpacked data for DocumentPollStarted events raised by the Polls contract.
type PollsDocumentPollStartedIterator struct {
	Event *PollsDocumentPollStarted // Event containing the contract specifics and raw log

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
func (it *PollsDocumentPollStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PollsDocumentPollStarted)
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
		it.Event = new(PollsDocumentPollStarted)
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
func (it *PollsDocumentPollStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PollsDocumentPollStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PollsDocumentPollStarted represents a DocumentPollStarted event raised by the Polls contract.
type PollsDocumentPollStarted struct {
	Proposal [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDocumentPollStarted is a free log retrieval operation binding the contract event 0xac68c1f0dc8aca8ee1bf445a687669d7d404b12dd0886f747e696115f0048527.
//
// Solidity: event DocumentPollStarted(bytes32 proposal)
func (_Polls *PollsFilterer) FilterDocumentPollStarted(opts *bind.FilterOpts) (*PollsDocumentPollStartedIterator, error) {

	logs, sub, err := _Polls.contract.FilterLogs(opts, "DocumentPollStarted")
	if err != nil {
		return nil, err
	}
	return &PollsDocumentPollStartedIterator{contract: _Polls.contract, event: "DocumentPollStarted", logs: logs, sub: sub}, nil
}

// WatchDocumentPollStarted is a free log subscription operation binding the contract event 0xac68c1f0dc8aca8ee1bf445a687669d7d404b12dd0886f747e696115f0048527.
//
// Solidity: event DocumentPollStarted(bytes32 proposal)
func (_Polls *PollsFilterer) WatchDocumentPollStarted(opts *bind.WatchOpts, sink chan<- *PollsDocumentPollStarted) (event.Subscription, error) {

	logs, sub, err := _Polls.contract.WatchLogs(opts, "DocumentPollStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PollsDocumentPollStarted)
				if err := _Polls.contract.UnpackLog(event, "DocumentPollStarted", log); err != nil {
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

// ParseDocumentPollStarted is a log parse operation binding the contract event 0xac68c1f0dc8aca8ee1bf445a687669d7d404b12dd0886f747e696115f0048527.
//
// Solidity: event DocumentPollStarted(bytes32 proposal)
func (_Polls *PollsFilterer) ParseDocumentPollStarted(log types.Log) (*PollsDocumentPollStarted, error) {
	event := new(PollsDocumentPollStarted)
	if err := _Polls.contract.UnpackLog(event, "DocumentPollStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PollsOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Polls contract.
type PollsOwnershipRenouncedIterator struct {
	Event *PollsOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PollsOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PollsOwnershipRenounced)
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
		it.Event = new(PollsOwnershipRenounced)
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
func (it *PollsOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PollsOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PollsOwnershipRenounced represents a OwnershipRenounced event raised by the Polls contract.
type PollsOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Polls *PollsFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PollsOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Polls.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PollsOwnershipRenouncedIterator{contract: _Polls.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Polls *PollsFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PollsOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Polls.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PollsOwnershipRenounced)
				if err := _Polls.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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
func (_Polls *PollsFilterer) ParseOwnershipRenounced(log types.Log) (*PollsOwnershipRenounced, error) {
	event := new(PollsOwnershipRenounced)
	if err := _Polls.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PollsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Polls contract.
type PollsOwnershipTransferredIterator struct {
	Event *PollsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PollsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PollsOwnershipTransferred)
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
		it.Event = new(PollsOwnershipTransferred)
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
func (it *PollsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PollsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PollsOwnershipTransferred represents a OwnershipTransferred event raised by the Polls contract.
type PollsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Polls *PollsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PollsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Polls.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PollsOwnershipTransferredIterator{contract: _Polls.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Polls *PollsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PollsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Polls.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PollsOwnershipTransferred)
				if err := _Polls.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Polls *PollsFilterer) ParseOwnershipTransferred(log types.Log) (*PollsOwnershipTransferred, error) {
	event := new(PollsOwnershipTransferred)
	if err := _Polls.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PollsUpgradeMajorityIterator is returned from FilterUpgradeMajority and is used to iterate over the raw logs and unpacked data for UpgradeMajority events raised by the Polls contract.
type PollsUpgradeMajorityIterator struct {
	Event *PollsUpgradeMajority // Event containing the contract specifics and raw log

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
func (it *PollsUpgradeMajorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PollsUpgradeMajority)
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
		it.Event = new(PollsUpgradeMajority)
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
func (it *PollsUpgradeMajorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PollsUpgradeMajorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PollsUpgradeMajority represents a UpgradeMajority event raised by the Polls contract.
type PollsUpgradeMajority struct {
	Proposal common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpgradeMajority is a free log retrieval operation binding the contract event 0x1ea969d04c4e4e370253bdeb7fb9638e50e75c206b4c42be36eb9844d37f4441.
//
// Solidity: event UpgradeMajority(address proposal)
func (_Polls *PollsFilterer) FilterUpgradeMajority(opts *bind.FilterOpts) (*PollsUpgradeMajorityIterator, error) {

	logs, sub, err := _Polls.contract.FilterLogs(opts, "UpgradeMajority")
	if err != nil {
		return nil, err
	}
	return &PollsUpgradeMajorityIterator{contract: _Polls.contract, event: "UpgradeMajority", logs: logs, sub: sub}, nil
}

// WatchUpgradeMajority is a free log subscription operation binding the contract event 0x1ea969d04c4e4e370253bdeb7fb9638e50e75c206b4c42be36eb9844d37f4441.
//
// Solidity: event UpgradeMajority(address proposal)
func (_Polls *PollsFilterer) WatchUpgradeMajority(opts *bind.WatchOpts, sink chan<- *PollsUpgradeMajority) (event.Subscription, error) {

	logs, sub, err := _Polls.contract.WatchLogs(opts, "UpgradeMajority")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PollsUpgradeMajority)
				if err := _Polls.contract.UnpackLog(event, "UpgradeMajority", log); err != nil {
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

// ParseUpgradeMajority is a log parse operation binding the contract event 0x1ea969d04c4e4e370253bdeb7fb9638e50e75c206b4c42be36eb9844d37f4441.
//
// Solidity: event UpgradeMajority(address proposal)
func (_Polls *PollsFilterer) ParseUpgradeMajority(log types.Log) (*PollsUpgradeMajority, error) {
	event := new(PollsUpgradeMajority)
	if err := _Polls.contract.UnpackLog(event, "UpgradeMajority", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PollsUpgradePollStartedIterator is returned from FilterUpgradePollStarted and is used to iterate over the raw logs and unpacked data for UpgradePollStarted events raised by the Polls contract.
type PollsUpgradePollStartedIterator struct {
	Event *PollsUpgradePollStarted // Event containing the contract specifics and raw log

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
func (it *PollsUpgradePollStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PollsUpgradePollStarted)
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
		it.Event = new(PollsUpgradePollStarted)
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
func (it *PollsUpgradePollStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PollsUpgradePollStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PollsUpgradePollStarted represents a UpgradePollStarted event raised by the Polls contract.
type PollsUpgradePollStarted struct {
	Proposal common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpgradePollStarted is a free log retrieval operation binding the contract event 0x6ffd05e058c57083a3d6fea058fbdbec172b55ce532b0f778f16e929a0ff516b.
//
// Solidity: event UpgradePollStarted(address proposal)
func (_Polls *PollsFilterer) FilterUpgradePollStarted(opts *bind.FilterOpts) (*PollsUpgradePollStartedIterator, error) {

	logs, sub, err := _Polls.contract.FilterLogs(opts, "UpgradePollStarted")
	if err != nil {
		return nil, err
	}
	return &PollsUpgradePollStartedIterator{contract: _Polls.contract, event: "UpgradePollStarted", logs: logs, sub: sub}, nil
}

// WatchUpgradePollStarted is a free log subscription operation binding the contract event 0x6ffd05e058c57083a3d6fea058fbdbec172b55ce532b0f778f16e929a0ff516b.
//
// Solidity: event UpgradePollStarted(address proposal)
func (_Polls *PollsFilterer) WatchUpgradePollStarted(opts *bind.WatchOpts, sink chan<- *PollsUpgradePollStarted) (event.Subscription, error) {

	logs, sub, err := _Polls.contract.WatchLogs(opts, "UpgradePollStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PollsUpgradePollStarted)
				if err := _Polls.contract.UnpackLog(event, "UpgradePollStarted", log); err != nil {
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

// ParseUpgradePollStarted is a log parse operation binding the contract event 0x6ffd05e058c57083a3d6fea058fbdbec172b55ce532b0f778f16e929a0ff516b.
//
// Solidity: event UpgradePollStarted(address proposal)
func (_Polls *PollsFilterer) ParseUpgradePollStarted(log types.Log) (*PollsUpgradePollStarted, error) {
	event := new(PollsUpgradePollStarted)
	if err := _Polls.contract.UnpackLog(event, "UpgradePollStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
