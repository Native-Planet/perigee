// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ecliptic

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

// EclipticMetaData contains all meta data concerning the Ecliptic contract.
var EclipticMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_galaxy\",\"type\":\"uint8\"},{\"name\":\"_proposal\",\"type\":\"bytes32\"}],\"name\":\"startDocumentPoll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"detach\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"approved\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_approved\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposal\",\"type\":\"bytes32\"}],\"name\":\"updateDocumentPoll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"onUpgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"InterfaceId_ERC165\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_reset\",\"type\":\"bool\"}],\"name\":\"transferPoint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_galaxy\",\"type\":\"uint8\"},{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"createGalaxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depositAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_transferProxy\",\"type\":\"address\"}],\"name\":\"setTransferProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryUpgradeHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_encryptionKey\",\"type\":\"bytes32\"},{\"name\":\"_authenticationKey\",\"type\":\"bytes32\"},{\"name\":\"_cryptoSuiteVersion\",\"type\":\"uint32\"},{\"name\":\"_discontinuous\",\"type\":\"bool\"}],\"name\":\"configureKeys\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_sponsor\",\"type\":\"uint32\"}],\"name\":\"canEscapeTo\",\"outputs\":[{\"name\":\"canEscape\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_treasuryImpl\",\"type\":\"address\"}],\"name\":\"upgradeTreasury\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"doesExist\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_galaxy\",\"type\":\"uint8\"},{\"name\":\"_proposal\",\"type\":\"address\"},{\"name\":\"_vote\",\"type\":\"bool\"}],\"name\":\"castUpgradeVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposal\",\"type\":\"address\"}],\"name\":\"updateUpgradePoll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_galaxy\",\"type\":\"uint8\"},{\"name\":\"_proposal\",\"type\":\"bytes32\"},{\"name\":\"_vote\",\"type\":\"bool\"}],\"name\":\"castDocumentVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"setManagementProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryUpgraded\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_galaxy\",\"type\":\"uint8\"},{\"name\":\"_proposal\",\"type\":\"address\"}],\"name\":\"startUpgradePoll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"spawn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_galaxy\",\"type\":\"uint8\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"setVotingProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_prefix\",\"type\":\"uint16\"},{\"name\":\"_spawnProxy\",\"type\":\"address\"}],\"name\":\"setSpawnProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_primary\",\"type\":\"string\"},{\"name\":\"_secondary\",\"type\":\"string\"},{\"name\":\"_tertiary\",\"type\":\"string\"}],\"name\":\"setDnsDomains\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"reject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_sponsor\",\"type\":\"uint32\"}],\"name\":\"escape\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"adopt\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"cancelEscape\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"_tokenURI\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"azimuth\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"claims\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"polls\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousEcliptic\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"getSpawnLimit\",\"outputs\":[{\"name\":\"limit\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_previous\",\"type\":\"address\"},{\"name\":\"_azimuth\",\"type\":\"address\"},{\"name\":\"_polls\",\"type\":\"address\"},{\"name\":\"_claims\",\"type\":\"address\"},{\"name\":\"_treasuryProxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
	Bin: "0x60806040526000600560146101000a81548160ff0219169083151502179055503480156200002c57600080fd5b5060405160a08062009e67833981018060405281019080805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919050505084848481336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505082600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050620001c06301ffc9a77c01000000000000000000000000000000000000000000000000000000000262000304640100000000026401000000009004565b81600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506200027f6380ac58cd7c01000000000000000000000000000000000000000000000000000000000262000304640100000000026401000000009004565b620002bc635b5e139f7c01000000000000000000000000000000000000000000000000000000000262000304640100000000026401000000009004565b620002f9637f5828d07c01000000000000000000000000000000000000000000000000000000000262000304640100000000026401000000009004565b5050505050620003c3565b63ffffffff7c010000000000000000000000000000000000000000000000000000000002817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916141515156200035657600080fd5b600160046000837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b619a9480620003d36000396000f30060806040526004361061025c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806301ffc9a71461026157806305bbf5db146102c557806306fdde0314610303578063073a780414610393578063081812fc146103c6578063095ea7b314610433578063172a735c146104805780631824a46b146104b157806319fa8f50146104c85780631e79a85b1461053157806323b872dd1461059057806326295b52146105fd57806328f833b71461064d5780632c7ba564146106a457806339d42b15146106f757806342842e0e1461072a5780634447e48c146107975780634a013296146108025780634e7e92991461085d5780634f558e79146108a057806351de5541146108e55780635623715b146109415780635d978826146109845780636352211e146109db5780636eb5822414610a4857806370a0823114610a92578063715018a614610ae95780638866bb2c14610b005780638da5cb5b14610b5357806395d89b4114610baa5780639e6e402814610c3a5780639e988d1314610c69578063a0d3253f14610cb9578063a22cb46514610d0c578063a60e8bd614610d5b578063ae32622114610dab578063b88d4fde14610dfc578063bac55edd14610eaf578063bbe21ca514610f1a578063bf5772b914610f4d578063c1b9d98b14610f90578063c6d761d414610fc3578063c87b56dd14610ff6578063d40ffacb1461109c578063dcc59b6f146110f3578063e64853c41461114a578063e985e9c5146111a1578063ed969f681461121c578063ef20bff814611273578063f2fde38b146112d0575b600080fd5b34801561026d57600080fd5b506102ab60048036038101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190505050611313565b604051808215151515815260200191505060405180910390f35b3480156102d157600080fd5b50610301600480360381019080803560ff169060200190929190803560001916906020019092919050505061137b565b005b34801561030f57600080fd5b5061031861162c565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561035857808201518184015260208101905061033d565b50505050905090810190601f1680156103855780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561039f57600080fd5b506103c4600480360381019080803563ffffffff169060200190929190505050611669565b005b3480156103d257600080fd5b506103f160048036038101908080359060200190929190505050611b1b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561043f57600080fd5b5061047e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611cf4565b005b34801561048c57600080fd5b506104af6004803603810190808035600019169060200190929190505050611d17565b005b3480156104bd57600080fd5b506104c6611def565b005b3480156104d457600080fd5b506104dd61203b565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b34801561053d57600080fd5b5061058e600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803515159060200190929190505050612062565b005b34801561059c57600080fd5b506105fb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506136f5565b005b34801561060957600080fd5b5061064b600480360381019080803560ff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613838565b005b34801561065957600080fd5b50610662613ab2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156106b057600080fd5b506106f5600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613aca565b005b34801561070357600080fd5b5061070c613f8a565b60405180826000191660001916815260200191505060405180910390f35b34801561073657600080fd5b50610795600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050613fae565b005b3480156107a357600080fd5b50610800600480360381019080803563ffffffff16906020019092919080356000191690602001909291908035600019169060200190929190803563ffffffff169060200190929190803515159060200190929190505050613fcf565b005b34801561080e57600080fd5b50610843600480360381019080803563ffffffff169060200190929190803563ffffffff16906020019092919050505061449c565b604051808215151515815260200191505060405180910390f35b34801561086957600080fd5b5061089e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061486a565b005b3480156108ac57600080fd5b506108cb60048036038101908080359060200190929190505050614a23565b604051808215151515815260200191505060405180910390f35b3480156108f157600080fd5b5061093f600480360381019080803560ff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803515159060200190929190505050614b14565b005b34801561094d57600080fd5b50610982600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050614e3c565b005b34801561099057600080fd5b50610999614f4c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156109e757600080fd5b50610a0660048036038101908080359060200190929190505050614f72565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610a5457600080fd5b50610a90600480360381019080803560ff1690602001909291908035600019169060200190929190803515159060200190929190505050615150565b005b348015610a9e57600080fd5b50610ad3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050615440565b6040518082815260200191505060405180910390f35b348015610af557600080fd5b50610afe615567565b005b348015610b0c57600080fd5b50610b51600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050615669565b005b348015610b5f57600080fd5b50610b68615a77565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610bb657600080fd5b50610bbf615a9c565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610bff578082015181840152602081019050610be4565b50505050905090810190601f168015610c2c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610c4657600080fd5b50610c4f615ad9565b604051808215151515815260200191505060405180910390f35b348015610c7557600080fd5b50610cb7600480360381019080803560ff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050615aec565b005b348015610cc557600080fd5b50610d0a600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050615e99565b005b348015610d1857600080fd5b50610d59600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035151590602001909291905050506168b5565b005b348015610d6757600080fd5b50610da9600480360381019080803560ff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050616a5d565b005b348015610db757600080fd5b50610dfa600480360381019080803561ffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050616d43565b005b348015610e0857600080fd5b50610ead600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061727b565b005b348015610ebb57600080fd5b50610f186004803603810190808035906020019082018035906020019190919293919293908035906020019082018035906020019190919293919293908035906020019082018035906020019190919293919293905050506174a4565b005b348015610f2657600080fd5b50610f4b600480360381019080803563ffffffff16906020019092919050505061760e565b005b348015610f5957600080fd5b50610f8e600480360381019080803563ffffffff169060200190929190803563ffffffff169060200190929190505050617ac0565b005b348015610f9c57600080fd5b50610fc1600480360381019080803563ffffffff169060200190929190505050617fe9565b005b348015610fcf57600080fd5b50610ff4600480360381019080803563ffffffff1690602001909291905050506185c3565b005b34801561100257600080fd5b5061102160048036038101908080359060200190929190505050618874565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015611061578082015181840152602081019050611046565b50505050905090810190601f16801561108e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156110a857600080fd5b506110b1618de3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156110ff57600080fd5b50611108618e09565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561115657600080fd5b5061115f618e2f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156111ad57600080fd5b50611202600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050618e55565b604051808215151515815260200191505060405180910390f35b34801561122857600080fd5b50611231618f8b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561127f57600080fd5b506112ae600480360381019080803563ffffffff16906020019092919080359060200190929190505050618fb1565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b3480156112dc57600080fd5b50611311600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050619122565b005b600060046000837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060009054906101000a900460ff169050919050565b8160ff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b65afedd82336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561145057600080fd5b505af1158015611464573d6000803e3d6000fd5b505050506040513d602081101561147a57600080fd5b8101908080519060200190929190505050801561156b5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561152f57600080fd5b505af1158015611543573d6000803e3d6000fd5b505050506040513d602081101561155957600080fd5b81019080805190602001909291905050505b151561157657600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639a33aff9836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050600060405180830381600087803b15801561160f57600080fd5b505af1158015611623573d6000803e3d6000fd5b50505050505050565b60606040805190810160405280600e81526020017f417a696d75746820506f696e7473000000000000000000000000000000000000815250905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663439f7d3c836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561170857600080fd5b505af115801561171c573d6000803e3d6000fd5b505050506040513d602081101561173257600080fd5b81019080805190602001909291905050509050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166377eb4c50836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b1580156117e257600080fd5b505af11580156117f6573d6000803e3d6000fd5b505050506040513d602081101561180c57600080fd5b810190808051906020019092919050505080156119315750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639137fe0a82336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b1580156118f557600080fd5b505af1158015611909573d6000803e3d6000fd5b505050506040513d602081101561191f57600080fd5b81019080805190602001909291905050505b151561193c57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663621b23e2826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b1580156119d957600080fd5b505af11580156119ed573d6000803e3d6000fd5b505050506040513d6020811015611a0357600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1673111111111111111111111111111111111111111173ffffffffffffffffffffffffffffffffffffffff1614151515611a6257600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bc562b9e836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050600060405180830381600087803b158015611aff57600080fd5b505af1158015611b13573d6000803e3d6000fd5b505050505050565b60008164010000000081101515611b3157600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015611bce57600080fd5b505af1158015611be2573d6000803e3d6000fd5b505050506040513d6020811015611bf857600080fd5b81019080805190602001909291905050501515611c1457600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324541f78846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015611cb157600080fd5b505af1158015611cc5573d6000803e3d6000fd5b505050506040513d6020811015611cdb57600080fd5b8101908080519060200190929190505050915050919050565b8064010000000081101515611d0857600080fd5b611d128284613aca565b505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663172a735c826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b158015611db057600080fd5b505af1158015611dc4573d6000803e3d6000fd5b505050506040513d6020811015611dda57600080fd5b81019080805190602001909291905050505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148015611f375750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015611ecd57600080fd5b505af1158015611ee1573d6000803e3d6000fd5b505050506040513d6020811015611ef757600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16145b801561202e5750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015611fc457600080fd5b505af1158015611fd8573d6000803e3d6000fd5b505050506040513d6020811015611fee57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16145b151561203957600080fd5b565b6301ffc9a77c01000000000000000000000000000000000000000000000000000000000281565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663728aa85785336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561213557600080fd5b505af1158015612149573d6000803e3d6000fd5b505050506040513d602081101561215f57600080fd5b8101908080519060200190929190505050151561217b57600080fd5b8273ffffffffffffffffffffffffffffffffffffffff1673111111111111111111111111111111111111111173ffffffffffffffffffffffffffffffffffffffff161415806123b75750600060028111156121d257fe5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166393976405866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561226f57600080fd5b505af1158015612283573d6000803e3d6000fd5b505050506040513d602081101561229957600080fd5b810190808051906020019092919050505060028111156122b557fe5b141580156123b657506123b4600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663621b23e2866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561235e57600080fd5b505af1158015612372573d6000803e3d6000fd5b505050506040513d602081101561238857600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff16619189565b155b5b15156123c257600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561245f57600080fd5b505af1158015612473573d6000803e3d6000fd5b505050506040513d602081101561248957600080fd5b8101908080519060200190929190505050151561255657600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637bc702a1856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050600060405180830381600087803b15801561253d57600080fd5b505af1158015612551573d6000803e3d6000fd5b505050505b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663caf590f985856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561262757600080fd5b505af115801561263b573d6000803e3d6000fd5b505050506040513d602081101561265157600080fd5b8101908080519060200190929190505050151561296157600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663621b23e2856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561270557600080fd5b505af1158015612719573d6000803e3d6000fd5b505050506040513d602081101561272f57600080fd5b81019080805190602001909291905050509050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ddc3595085856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b15801561281357600080fd5b505af1158015612827573d6000803e3d6000fd5b50505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c7ba5648560006040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b1580156128e757600080fd5b505af11580156128fb573d6000803e3d6000fd5b505050508363ffffffff168373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a45b8273ffffffffffffffffffffffffffffffffffffffff1673111111111111111111111111111111111111111173ffffffffffffffffffffffffffffffffffffffff161415612f4b57600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f9b87d408560008060006040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808563ffffffff1663ffffffff16815260200184600102600019168152602001836001026000191681526020018263ffffffff168152602001945050505050600060405180830381600087803b158015612a7457600080fd5b505af1158015612a88573d6000803e3d6000fd5b50505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638866bb2c8560006040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b158015612b4857600080fd5b505af1158015612b5c573d6000803e3d6000fd5b50505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a297c1d88560006040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b158015612c1c57600080fd5b505af1158015612c30573d6000803e3d6000fd5b50505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c7ba5648560006040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b158015612cf057600080fd5b505af1158015612d04573d6000803e3d6000fd5b50505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632a19642c8560006040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b158015612dc457600080fd5b505af1158015612dd8573d6000803e3d6000fd5b50505050600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663eaae46e5856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050600060405180830381600087803b158015612e7957600080fd5b505af1158015612e8d573d6000803e3d6000fd5b50505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c6d761d4856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050600060405180830381600087803b158015612f2e57600080fd5b505af1158015612f42573d6000803e3d6000fd5b505050506136ef565b81156136ee57600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636d09887b856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015612fee57600080fd5b505af1158015613002573d6000803e3d6000fd5b505050506040513d602081101561301857600080fd5b8101908080519060200190929190505050156131c757600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f4914919856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050600060405180830381600087803b1580156130cb57600080fd5b505af11580156130df573d6000803e3d6000fd5b50505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f9b87d408560008060006040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808563ffffffff1663ffffffff16815260200184600102600019168152602001836001026000191681526020018263ffffffff168152602001945050505050600060405180830381600087803b1580156131ae57600080fd5b505af11580156131c2573d6000803e3d6000fd5b505050505b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638866bb2c8560006040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b15801561328357600080fd5b505af1158015613297573d6000803e3d6000fd5b50505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a297c1d88560006040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b15801561335757600080fd5b505af115801561336b573d6000803e3d6000fd5b50505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c7ba5648560006040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b15801561342b57600080fd5b505af115801561343f573d6000803e3d6000fd5b50505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f4e3be2d856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b1580156134e057600080fd5b505af11580156134f4573d6000803e3d6000fd5b505050506040513d602081101561350a57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1673111111111111111111111111111111111111111173ffffffffffffffffffffffffffffffffffffffff1614151561363857600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632a19642c8560006040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b15801561361f57600080fd5b505af1158015613633573d6000803e3d6000fd5b505050505b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663eaae46e5856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050600060405180830381600087803b1580156136d557600080fd5b505af11580156136e9573d6000803e3d6000fd5b505050505b5b50505050565b6000816401000000008110151561370b57600080fd5b829150600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663caf590f983876040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b1580156137df57600080fd5b505af11580156137f3573d6000803e3d6000fd5b505050506040513d602081101561380957600080fd5b8101908080519060200190929190505050151561382557600080fd5b61383182856001612062565b5050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561389357600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663caf590f98360006040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808360ff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561394c57600080fd5b505af1158015613960573d6000803e3d6000fd5b505050506040513d602081101561397657600080fd5b810190808051906020019092919050505080156139aa57508073ffffffffffffffffffffffffffffffffffffffff16600014155b15156139b557600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663246d41a96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b158015613a3b57600080fd5b505af1158015613a4f573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415613a9d57613a988260ff16826001600061919c565b613aae565b613aad8260ff168260003361919c565b5b5050565b73111111111111111111111111111111111111111181565b600082600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663621b23e2826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015613b6a57600080fd5b505af1158015613b7e573d6000803e3d6000fd5b505050506040513d6020811015613b9457600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1673111111111111111111111111111111111111111173ffffffffffffffffffffffffffffffffffffffff1614151515613bf357600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663621b23e2856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015613c9057600080fd5b505af1158015613ca4573d6000803e3d6000fd5b505050506040513d6020811015613cba57600080fd5b810190808051906020019092919050505091503373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161480613e2f5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b6363cf283336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015613df357600080fd5b505af1158015613e07573d6000803e3d6000fd5b505050506040513d6020811015613e1d57600080fd5b81019080805190602001909291905050505b1515613e3a57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c7ba56485856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b158015613f0b57600080fd5b505af1158015613f1f573d6000803e3d6000fd5b505050508363ffffffff168373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a450505050565b7f26f3eae628fa1a4d23e34b91a4d412526a47620ced37c80928906f9fa07c077481565b613fca838383602060405190810160405280600081525061727b565b505050565b84600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639137fe0a82336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b1580156140a157600080fd5b505af11580156140b5573d6000803e3d6000fd5b505050506040513d60208110156140cb57600080fd5b810190808051906020019092919050505080156141bc5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561418057600080fd5b505af1158015614194573d6000803e3d6000fd5b505050506040513d60208110156141aa57600080fd5b81019080805190602001909291905050505b15156141c757600080fd5b85600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663621b23e2826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561426557600080fd5b505af1158015614279573d6000803e3d6000fd5b505050506040513d602081101561428f57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1673111111111111111111111111111111111111111173ffffffffffffffffffffffffffffffffffffffff16141515156142ee57600080fd5b82156143aa57600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f4914919886040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050600060405180830381600087803b15801561439157600080fd5b505af11580156143a5573d6000803e3d6000fd5b505050505b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f9b87d40888888886040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808563ffffffff1663ffffffff168152602001846000191660001916815260200183600019166000191681526020018263ffffffff1663ffffffff168152602001945050505050600060405180830381600087803b15801561447b57600080fd5b505af115801561448f573d6000803e3d6000fd5b5050505050505050505050565b6000806000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636d09887b856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561453e57600080fd5b505af1158015614552573d6000803e3d6000fd5b505050506040513d602081101561456857600080fd5b810190808051906020019092919050505015156145885760009250614862565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166393976405866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561462557600080fd5b505af1158015614639573d6000803e3d6000fd5b505050506040513d602081101561464f57600080fd5b81019080805190602001909291905050509150600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166393976405856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b1580156146ff57600080fd5b505af1158015614713573d6000803e3d6000fd5b505050506040513d602081101561472957600080fd5b8101908080519060200190929190505050905081600281111561474857fe5b60ff16600182600281111561475957fe5b0160ff16148061485f575081600281111561477057fe5b81600281111561477c57fe5b14801561485e5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636d09887b866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561482157600080fd5b505af1158015614835573d6000803e3d6000fd5b505050506040513d602081101561484b57600080fd5b8101908080519060200190929190505050155b5b92505b505092915050565b600560149054906101000a900460ff1615151561488657600080fd5b7f26f3eae628fa1a4d23e34b91a4d412526a47620ced37c80928906f9fa07c07746000191681604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c0100000000000000000000000002815260140191505060405180910390206000191614151561490c57600080fd5b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633659cfe6826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156149c957600080fd5b505af11580156149dd573d6000803e3d6000fd5b505050506040513d60208110156149f357600080fd5b8101908080519060200190929190505050506001600560146101000a81548160ff02191690831515021790555050565b600064010000000082108015614b0d5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015614ad157600080fd5b505af1158015614ae5573d6000803e3d6000fd5b505050506040513d6020811015614afb57600080fd5b81019080805190602001909291905050505b9050919050565b60008360ff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b65afedd82336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015614beb57600080fd5b505af1158015614bff573d6000803e3d6000fd5b505050506040513d6020811015614c1557600080fd5b81019080805190602001909291905050508015614d065750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015614cca57600080fd5b505af1158015614cde573d6000803e3d6000fd5b505050506040513d6020811015614cf457600080fd5b81019080805190602001909291905050505b1515614d1157600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166351de55418686866040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808460ff1660ff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001821515151581526020019350505050602060405180830381600087803b158015614de857600080fd5b505af1158015614dfc573d6000803e3d6000fd5b505050506040513d6020811015614e1257600080fd5b810190808051906020019092919050505091508115614e3557614e34846196cc565b5b5050505050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635623715b836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015614efb57600080fd5b505af1158015614f0f573d6000803e3d6000fd5b505050506040513d6020811015614f2557600080fd5b810190808051906020019092919050505090508015614f4857614f47826196cc565b5b5050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000808264010000000081101515614f8957600080fd5b839150600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561502957600080fd5b505af115801561503d573d6000803e3d6000fd5b505050506040513d602081101561505357600080fd5b8101908080519060200190929190505050151561506f57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663621b23e2836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561510c57600080fd5b505af1158015615120573d6000803e3d6000fd5b505050506040513d602081101561513657600080fd5b810190808051906020019092919050505092505050919050565b8260ff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b65afedd82336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561522557600080fd5b505af1158015615239573d6000803e3d6000fd5b505050506040513d602081101561524f57600080fd5b810190808051906020019092919050505080156153405750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561530457600080fd5b505af1158015615318573d6000803e3d6000fd5b505050506040513d602081101561532e57600080fd5b81019080805190602001909291905050505b151561534b57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636eb582248585856040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808460ff1660ff1681526020018360001916600019168152602001821515151581526020019350505050602060405180830381600087803b1580156153fe57600080fd5b505af1158015615412573d6000803e3d6000fd5b505050506040513d602081101561542857600080fd5b81019080805190602001909291905050505050505050565b60008173ffffffffffffffffffffffffffffffffffffffff1660001415151561546857600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166312afbc78836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561552557600080fd5b505af1158015615539573d6000803e3d6000fd5b505050506040513d602081101561554f57600080fd5b81019080805190602001909291905050509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156155c257600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b81600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639137fe0a82336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561573b57600080fd5b505af115801561574f573d6000803e3d6000fd5b505050506040513d602081101561576557600080fd5b810190808051906020019092919050505080156158565750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561581a57600080fd5b505af115801561582e573d6000803e3d6000fd5b505050506040513d602081101561584457600080fd5b81019080805190602001909291905050505b151561586157600080fd5b82600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663621b23e2826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b1580156158ff57600080fd5b505af1158015615913573d6000803e3d6000fd5b505050506040513d602081101561592957600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1673111111111111111111111111111111111111111173ffffffffffffffffffffffffffffffffffffffff161415151561598857600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638866bb2c85856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b158015615a5957600080fd5b505af1158015615a6d573d6000803e3d6000fd5b5050505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606040805190810160405280600381526020017f415a500000000000000000000000000000000000000000000000000000000000815250905090565b600560149054906101000a900460ff1681565b8160ff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b65afedd82336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015615bc157600080fd5b505af1158015615bd5573d6000803e3d6000fd5b505050506040513d6020811015615beb57600080fd5b81019080805190602001909291905050508015615cdc5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015615ca057600080fd5b505af1158015615cb4573d6000803e3d6000fd5b505050506040513d6020811015615cca57600080fd5b81019080805190602001909291905050505b1515615ce757600080fd5b3073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1663ed969f686040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015615d6257600080fd5b505af1158015615d76573d6000803e3d6000fd5b505050506040513d6020811015615d8c57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff16141515615dbf57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e31f3e0c836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b158015615e7c57600080fd5b505af1158015615e90573d6000803e3d6000fd5b50505050505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663caf590f98460006040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015615f5757600080fd5b505af1158015615f6b573d6000803e3d6000fd5b505050506040513d6020811015615f8157600080fd5b81019080805190602001909291905050501515615f9d57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e4a358d7846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561603a57600080fd5b505af115801561604e573d6000803e3d6000fd5b505050506040513d602081101561606457600080fd5b81019080805190602001909291905050509050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663621b23e2826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808261ffff1663ffffffff168152602001915050602060405180830381600087803b15801561611257600080fd5b505af1158015616126573d6000803e3d6000fd5b505050506040513d602081101561613c57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1673111111111111111111111111111111111111111173ffffffffffffffffffffffffffffffffffffffff161415151561619b57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f4e3be2d826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808261ffff1663ffffffff168152602001915050602060405180830381600087803b15801561623657600080fd5b505af115801561624a573d6000803e3d6000fd5b505050506040513d602081101561626057600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1673111111111111111111111111111111111111111173ffffffffffffffffffffffffffffffffffffffff16141515156162bf57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166393976405846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561635c57600080fd5b505af1158015616370573d6000803e3d6000fd5b505050506040513d602081101561638657600080fd5b810190808051906020019092919050505060028111156163a257fe5b60ff1660018060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166393976405846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808261ffff1663ffffffff168152602001915050602060405180830381600087803b15801561644157600080fd5b505af1158015616455573d6000803e3d6000fd5b505050506040513d602081101561646b57600080fd5b8101908080519060200190929190505050600281111561648757fe5b0160ff1614151561649757600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636d09887b826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808261ffff1663ffffffff168152602001915050602060405180830381600087803b15801561653257600080fd5b505af1158015616546573d6000803e3d6000fd5b505050506040513d602081101561655c57600080fd5b8101908080519060200190929190505050801561666657506165828161ffff1642618fb1565b63ffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663293a9169836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808261ffff1663ffffffff168152602001915050602060405180830381600087803b15801561662357600080fd5b505af1158015616637573d6000803e3d6000fd5b505050506040513d602081101561664d57600080fd5b810190808051906020019092919050505063ffffffff16105b151561667157600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638a27bf5982336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808361ffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561674057600080fd5b505af1158015616754573d6000803e3d6000fd5b505050506040513d602081101561676a57600080fd5b8101908080519060200190929190505050151561678657600080fd5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156167cd576167c883836001600061919c565b6168b0565b6168af83836000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663621b23e2866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808261ffff1663ffffffff168152602001915050602060405180830381600087803b15801561686f57600080fd5b505af1158015616883573d6000803e3d6000fd5b505050506040513d602081101561689957600080fd5b810190808051906020019092919050505061919c565b5b505050565b8173ffffffffffffffffffffffffffffffffffffffff166000141515156168db57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bc735d903384846040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001821515151581526020019350505050600060405180830381600087803b1580156169d857600080fd5b505af11580156169ec573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051808215151515815260200191505060405180910390a35050565b8160ff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b65afedd82336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015616b3257600080fd5b505af1158015616b46573d6000803e3d6000fd5b505050506040513d6020811015616b5c57600080fd5b81019080805190602001909291905050508015616c4d5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015616c1157600080fd5b505af1158015616c25573d6000803e3d6000fd5b505050506040513d6020811015616c3b57600080fd5b81019080805190602001909291905050505b1515616c5857600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a297c1d884846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808360ff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b158015616d2657600080fd5b505af1158015616d3a573d6000803e3d6000fd5b50505050505050565b8161ffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638a27bf5982336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015616e1957600080fd5b505af1158015616e2d573d6000803e3d6000fd5b505050506040513d6020811015616e4357600080fd5b81019080805190602001909291905050508015616f345750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015616ef857600080fd5b505af1158015616f0c573d6000803e3d6000fd5b505050506040513d6020811015616f2257600080fd5b81019080805190602001909291905050505b1515616f3f57600080fd5b8261ffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663621b23e2826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015616fe157600080fd5b505af1158015616ff5573d6000803e3d6000fd5b505050506040513d602081101561700b57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1673111111111111111111111111111111111111111173ffffffffffffffffffffffffffffffffffffffff161415151561706a57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f4e3be2d856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808261ffff1663ffffffff168152602001915050602060405180830381600087803b15801561710557600080fd5b505af1158015617119573d6000803e3d6000fd5b505050506040513d602081101561712f57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1673111111111111111111111111111111111111111173ffffffffffffffffffffffffffffffffffffffff161415151561718e57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632a19642c85856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808361ffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b15801561725d57600080fd5b505af1158015617271573d6000803e3d6000fd5b5050505050505050565b60006172888585856136f5565b6172a78473ffffffffffffffffffffffffffffffffffffffff16619189565b1561749d578373ffffffffffffffffffffffffffffffffffffffff1663150b7a02338786866040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156173a1578082015181840152602081019050617386565b50505050905090810190601f1680156173ce5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b1580156173f057600080fd5b505af1158015617404573d6000803e3d6000fd5b505050506040513d602081101561741a57600080fd5b8101908080519060200190929190505050905063150b7a027c0100000000000000000000000000000000000000000000000000000000027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614151561749c57600080fd5b5b5050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156174ff57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bac55edd8787878787876040518763ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001806020018060200184810384528a8a82818152602001925080828437820191505084810383528888828181526020019250808284378201915050848103825286868281815260200192508082843782019150509950505050505050505050600060405180830381600087803b1580156175ee57600080fd5b505af1158015617602573d6000803e3d6000fd5b50505050505050505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f5af6621836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b1580156176ad57600080fd5b505af11580156176c1573d6000803e3d6000fd5b505050506040513d60208110156176d757600080fd5b81019080805190602001909291905050509050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639b350e12836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561778757600080fd5b505af115801561779b573d6000803e3d6000fd5b505050506040513d60208110156177b157600080fd5b810190808051906020019092919050505080156178d65750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639137fe0a82336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561789a57600080fd5b505af11580156178ae573d6000803e3d6000fd5b505050506040513d60208110156178c457600080fd5b81019080805190602001909291905050505b15156178e157600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663621b23e2826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561797e57600080fd5b505af1158015617992573d6000803e3d6000fd5b505050506040513d60208110156179a857600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1673111111111111111111111111111111111111111173ffffffffffffffffffffffffffffffffffffffff1614151515617a0757600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c6d761d4836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050600060405180830381600087803b158015617aa457600080fd5b505af1158015617ab8573d6000803e3d6000fd5b505050505050565b81600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639137fe0a82336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015617b9257600080fd5b505af1158015617ba6573d6000803e3d6000fd5b505050506040513d6020811015617bbc57600080fd5b81019080805190602001909291905050508015617cad5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015617c7157600080fd5b505af1158015617c85573d6000803e3d6000fd5b505050506040513d6020811015617c9b57600080fd5b81019080805190602001909291905050505b1515617cb857600080fd5b82600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663621b23e2826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015617d5657600080fd5b505af1158015617d6a573d6000803e3d6000fd5b505050506040513d6020811015617d8057600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1673111111111111111111111111111111111111111173ffffffffffffffffffffffffffffffffffffffff1614151515617ddf57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663621b23e2846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015617e7c57600080fd5b505af1158015617e90573d6000803e3d6000fd5b505050506040513d6020811015617ea657600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1673111111111111111111111111111111111111111173ffffffffffffffffffffffffffffffffffffffff1614151515617f0557600080fd5b617f0f848461449c565b1515617f1a57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a634585985856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018263ffffffff1663ffffffff16815260200192505050600060405180830381600087803b158015617fcb57600080fd5b505af1158015617fdf573d6000803e3d6000fd5b5050505050505050565b600081600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663621b23e2826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561808957600080fd5b505af115801561809d573d6000803e3d6000fd5b505050506040513d60208110156180b357600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1673111111111111111111111111111111111111111173ffffffffffffffffffffffffffffffffffffffff161415151561811257600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f5af6621846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b1580156181af57600080fd5b505af11580156181c3573d6000803e3d6000fd5b505050506040513d60208110156181d957600080fd5b81019080805190602001909291905050509150600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639b350e12846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561828957600080fd5b505af115801561829d573d6000803e3d6000fd5b505050506040513d60208110156182b357600080fd5b810190808051906020019092919050505080156183d85750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639137fe0a83336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561839c57600080fd5b505af11580156183b0573d6000803e3d6000fd5b505050506040513d60208110156183c657600080fd5b81019080805190602001909291905050505b15156183e357600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663621b23e2836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561848057600080fd5b505af1158015618494573d6000803e3d6000fd5b505050506040513d60208110156184aa57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1673111111111111111111111111111111111111111173ffffffffffffffffffffffffffffffffffffffff161415151561850957600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166313063180846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050600060405180830381600087803b1580156185a657600080fd5b505af11580156185ba573d6000803e3d6000fd5b50505050505050565b80600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639137fe0a82336040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561869557600080fd5b505af11580156186a9573d6000803e3d6000fd5b505050506040513d60208110156186bf57600080fd5b810190808051906020019092919050505080156187b05750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e19b305826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561877457600080fd5b505af1158015618788573d6000803e3d6000fd5b505050506040513d602081101561879e57600080fd5b81019080805190602001909291905050505b15156187bb57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c6d761d4836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050600060405180830381600087803b15801561885857600080fd5b505af115801561886c573d6000803e3d6000fd5b505050505050565b606080826401000000008110151561888b57600080fd5b606060405190810160405280602e81526020017f68747470733a2f2f617a696d7574682e6e6574776f726b2f6572633732312f3081526020017f3030303030303030302e6a736f6e0000000000000000000000000000000000008152509250829150600a633b9aca00858115156188fe57fe5b0481151561890857fe5b066030017f01000000000000000000000000000000000000000000000000000000000000000282601f81518110151561893d57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a6305f5e1008581151561897e57fe5b0481151561898857fe5b066030017f0100000000000000000000000000000000000000000000000000000000000000028260208151811015156189bd57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a62989680858115156189fd57fe5b04811515618a0757fe5b066030017f010000000000000000000000000000000000000000000000000000000000000002826021815181101515618a3c57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a620f424085811515618a7c57fe5b04811515618a8657fe5b066030017f010000000000000000000000000000000000000000000000000000000000000002826022815181101515618abb57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a620186a085811515618afb57fe5b04811515618b0557fe5b066030017f010000000000000000000000000000000000000000000000000000000000000002826023815181101515618b3a57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a61271085811515618b7957fe5b04811515618b8357fe5b066030017f010000000000000000000000000000000000000000000000000000000000000002826024815181101515618bb857fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a6103e885811515618bf757fe5b04811515618c0157fe5b066030017f010000000000000000000000000000000000000000000000000000000000000002826025815181101515618c3657fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a606485811515618c7457fe5b04811515618c7e57fe5b066030017f010000000000000000000000000000000000000000000000000000000000000002826026815181101515618cb357fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a8085811515618cf057fe5b04811515618cfa57fe5b066030017f010000000000000000000000000000000000000000000000000000000000000002826027815181101515618d2f57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a600185811515618d6d57fe5b04811515618d7757fe5b066030017f010000000000000000000000000000000000000000000000000000000000000002826028815181101515618dac57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b6363cf284846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015618f4857600080fd5b505af1158015618f5c573d6000803e3d6000fd5b505050506040513d6020811015618f7257600080fd5b8101908080519060200190929190505050905092915050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166393976405866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561905357600080fd5b505af1158015619067573d6000803e3d6000fd5b505050506040513d602081101561907d57600080fd5b810190808051906020019092919050505091506000600281111561909d57fe5b8260028111156190a957fe5b14156190b85760ff925061911a565b600160028111156190c557fe5b8260028111156190d157fe5b1415619115576301e13380635c2aad8085038115156190ec57fe5b0490506006811015619107578060020a61040002925061910d565b61ffff92505b82925061911a565b600092505b505092915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561917d57600080fd5b6191868161996e565b50565b600080823b905060008111915050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c17ad9d8856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050600060405180830381600087803b15801561923957600080fd5b505af115801561924d573d6000803e3d6000fd5b50505050811561944657600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637bc702a1856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808263ffffffff1663ffffffff168152602001915050600060405180830381600087803b1580156192f457600080fd5b505af1158015619308573d6000803e3d6000fd5b50505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ddc3595085856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b1580156193dd57600080fd5b505af11580156193f1573d6000803e3d6000fd5b505050508363ffffffff168373ffffffffffffffffffffffffffffffffffffffff1660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a46196c6565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ddc3595085836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b15801561951757600080fd5b505af115801561952b573d6000803e3d6000fd5b50505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c7ba56485856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808363ffffffff1663ffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b15801561960057600080fd5b505af1158015619614573d6000803e3d6000fd5b505050508363ffffffff168173ffffffffffffffffffffffffffffffffffffffff1660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a48363ffffffff168373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b50505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f2fde38b826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561978957600080fd5b505af115801561979d573d6000803e3d6000fd5b50505050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f2fde38b826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561985e57600080fd5b505af1158015619872573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff16631824a46b6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b1580156198da57600080fd5b505af11580156198ee573d6000803e3d6000fd5b505050507fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b81604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a18073ffffffffffffffffffffffffffffffffffffffff16ff5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156199aa57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a723058206ea1bfcaf77e9b1ba4f4c3df26e526c5464b9846b21fa221d5aaef3641534a270029",
}

// EclipticABI is the input ABI used to generate the binding from.
// Deprecated: Use EclipticMetaData.ABI instead.
var EclipticABI = EclipticMetaData.ABI

// EclipticBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EclipticMetaData.Bin instead.
var EclipticBin = EclipticMetaData.Bin

// DeployEcliptic deploys a new Ethereum contract, binding an instance of Ecliptic to it.
func DeployEcliptic(auth *bind.TransactOpts, backend bind.ContractBackend, _previous common.Address, _azimuth common.Address, _polls common.Address, _claims common.Address, _treasuryProxy common.Address) (common.Address, *types.Transaction, *Ecliptic, error) {
	parsed, err := EclipticMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EclipticBin), backend, _previous, _azimuth, _polls, _claims, _treasuryProxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ecliptic{EclipticCaller: EclipticCaller{contract: contract}, EclipticTransactor: EclipticTransactor{contract: contract}, EclipticFilterer: EclipticFilterer{contract: contract}}, nil
}

// Ecliptic is an auto generated Go binding around an Ethereum contract.
type Ecliptic struct {
	EclipticCaller     // Read-only binding to the contract
	EclipticTransactor // Write-only binding to the contract
	EclipticFilterer   // Log filterer for contract events
}

// EclipticCaller is an auto generated read-only Go binding around an Ethereum contract.
type EclipticCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EclipticTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EclipticTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EclipticFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EclipticFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EclipticSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EclipticSession struct {
	Contract     *Ecliptic         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EclipticCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EclipticCallerSession struct {
	Contract *EclipticCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EclipticTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EclipticTransactorSession struct {
	Contract     *EclipticTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EclipticRaw is an auto generated low-level Go binding around an Ethereum contract.
type EclipticRaw struct {
	Contract *Ecliptic // Generic contract binding to access the raw methods on
}

// EclipticCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EclipticCallerRaw struct {
	Contract *EclipticCaller // Generic read-only contract binding to access the raw methods on
}

// EclipticTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EclipticTransactorRaw struct {
	Contract *EclipticTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEcliptic creates a new instance of Ecliptic, bound to a specific deployed contract.
func NewEcliptic(address common.Address, backend bind.ContractBackend) (*Ecliptic, error) {
	contract, err := bindEcliptic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ecliptic{EclipticCaller: EclipticCaller{contract: contract}, EclipticTransactor: EclipticTransactor{contract: contract}, EclipticFilterer: EclipticFilterer{contract: contract}}, nil
}

// NewEclipticCaller creates a new read-only instance of Ecliptic, bound to a specific deployed contract.
func NewEclipticCaller(address common.Address, caller bind.ContractCaller) (*EclipticCaller, error) {
	contract, err := bindEcliptic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EclipticCaller{contract: contract}, nil
}

// NewEclipticTransactor creates a new write-only instance of Ecliptic, bound to a specific deployed contract.
func NewEclipticTransactor(address common.Address, transactor bind.ContractTransactor) (*EclipticTransactor, error) {
	contract, err := bindEcliptic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EclipticTransactor{contract: contract}, nil
}

// NewEclipticFilterer creates a new log filterer instance of Ecliptic, bound to a specific deployed contract.
func NewEclipticFilterer(address common.Address, filterer bind.ContractFilterer) (*EclipticFilterer, error) {
	contract, err := bindEcliptic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EclipticFilterer{contract: contract}, nil
}

// bindEcliptic binds a generic wrapper to an already deployed contract.
func bindEcliptic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EclipticMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ecliptic *EclipticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ecliptic.Contract.EclipticCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ecliptic *EclipticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ecliptic.Contract.EclipticTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ecliptic *EclipticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ecliptic.Contract.EclipticTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ecliptic *EclipticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ecliptic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ecliptic *EclipticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ecliptic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ecliptic *EclipticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ecliptic.Contract.contract.Transact(opts, method, params...)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() view returns(bytes4)
func (_Ecliptic *EclipticCaller) InterfaceIdERC165(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "InterfaceId_ERC165")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() view returns(bytes4)
func (_Ecliptic *EclipticSession) InterfaceIdERC165() ([4]byte, error) {
	return _Ecliptic.Contract.InterfaceIdERC165(&_Ecliptic.CallOpts)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() view returns(bytes4)
func (_Ecliptic *EclipticCallerSession) InterfaceIdERC165() ([4]byte, error) {
	return _Ecliptic.Contract.InterfaceIdERC165(&_Ecliptic.CallOpts)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Ecliptic *EclipticCaller) Azimuth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "azimuth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Ecliptic *EclipticSession) Azimuth() (common.Address, error) {
	return _Ecliptic.Contract.Azimuth(&_Ecliptic.CallOpts)
}

// Azimuth is a free data retrieval call binding the contract method 0xd40ffacb.
//
// Solidity: function azimuth() view returns(address)
func (_Ecliptic *EclipticCallerSession) Azimuth() (common.Address, error) {
	return _Ecliptic.Contract.Azimuth(&_Ecliptic.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Ecliptic *EclipticCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Ecliptic *EclipticSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Ecliptic.Contract.BalanceOf(&_Ecliptic.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Ecliptic *EclipticCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Ecliptic.Contract.BalanceOf(&_Ecliptic.CallOpts, _owner)
}

// CanEscapeTo is a free data retrieval call binding the contract method 0x4a013296.
//
// Solidity: function canEscapeTo(uint32 _point, uint32 _sponsor) view returns(bool canEscape)
func (_Ecliptic *EclipticCaller) CanEscapeTo(opts *bind.CallOpts, _point uint32, _sponsor uint32) (bool, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "canEscapeTo", _point, _sponsor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanEscapeTo is a free data retrieval call binding the contract method 0x4a013296.
//
// Solidity: function canEscapeTo(uint32 _point, uint32 _sponsor) view returns(bool canEscape)
func (_Ecliptic *EclipticSession) CanEscapeTo(_point uint32, _sponsor uint32) (bool, error) {
	return _Ecliptic.Contract.CanEscapeTo(&_Ecliptic.CallOpts, _point, _sponsor)
}

// CanEscapeTo is a free data retrieval call binding the contract method 0x4a013296.
//
// Solidity: function canEscapeTo(uint32 _point, uint32 _sponsor) view returns(bool canEscape)
func (_Ecliptic *EclipticCallerSession) CanEscapeTo(_point uint32, _sponsor uint32) (bool, error) {
	return _Ecliptic.Contract.CanEscapeTo(&_Ecliptic.CallOpts, _point, _sponsor)
}

// Claims is a free data retrieval call binding the contract method 0xdcc59b6f.
//
// Solidity: function claims() view returns(address)
func (_Ecliptic *EclipticCaller) Claims(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "claims")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Claims is a free data retrieval call binding the contract method 0xdcc59b6f.
//
// Solidity: function claims() view returns(address)
func (_Ecliptic *EclipticSession) Claims() (common.Address, error) {
	return _Ecliptic.Contract.Claims(&_Ecliptic.CallOpts)
}

// Claims is a free data retrieval call binding the contract method 0xdcc59b6f.
//
// Solidity: function claims() view returns(address)
func (_Ecliptic *EclipticCallerSession) Claims() (common.Address, error) {
	return _Ecliptic.Contract.Claims(&_Ecliptic.CallOpts)
}

// DepositAddress is a free data retrieval call binding the contract method 0x28f833b7.
//
// Solidity: function depositAddress() view returns(address)
func (_Ecliptic *EclipticCaller) DepositAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "depositAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositAddress is a free data retrieval call binding the contract method 0x28f833b7.
//
// Solidity: function depositAddress() view returns(address)
func (_Ecliptic *EclipticSession) DepositAddress() (common.Address, error) {
	return _Ecliptic.Contract.DepositAddress(&_Ecliptic.CallOpts)
}

// DepositAddress is a free data retrieval call binding the contract method 0x28f833b7.
//
// Solidity: function depositAddress() view returns(address)
func (_Ecliptic *EclipticCallerSession) DepositAddress() (common.Address, error) {
	return _Ecliptic.Contract.DepositAddress(&_Ecliptic.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) view returns(bool doesExist)
func (_Ecliptic *EclipticCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "exists", _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) view returns(bool doesExist)
func (_Ecliptic *EclipticSession) Exists(_tokenId *big.Int) (bool, error) {
	return _Ecliptic.Contract.Exists(&_Ecliptic.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) view returns(bool doesExist)
func (_Ecliptic *EclipticCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _Ecliptic.Contract.Exists(&_Ecliptic.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address approved)
func (_Ecliptic *EclipticCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "getApproved", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address approved)
func (_Ecliptic *EclipticSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Ecliptic.Contract.GetApproved(&_Ecliptic.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address approved)
func (_Ecliptic *EclipticCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Ecliptic.Contract.GetApproved(&_Ecliptic.CallOpts, _tokenId)
}

// GetSpawnLimit is a free data retrieval call binding the contract method 0xef20bff8.
//
// Solidity: function getSpawnLimit(uint32 _point, uint256 _time) view returns(uint32 limit)
func (_Ecliptic *EclipticCaller) GetSpawnLimit(opts *bind.CallOpts, _point uint32, _time *big.Int) (uint32, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "getSpawnLimit", _point, _time)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetSpawnLimit is a free data retrieval call binding the contract method 0xef20bff8.
//
// Solidity: function getSpawnLimit(uint32 _point, uint256 _time) view returns(uint32 limit)
func (_Ecliptic *EclipticSession) GetSpawnLimit(_point uint32, _time *big.Int) (uint32, error) {
	return _Ecliptic.Contract.GetSpawnLimit(&_Ecliptic.CallOpts, _point, _time)
}

// GetSpawnLimit is a free data retrieval call binding the contract method 0xef20bff8.
//
// Solidity: function getSpawnLimit(uint32 _point, uint256 _time) view returns(uint32 limit)
func (_Ecliptic *EclipticCallerSession) GetSpawnLimit(_point uint32, _time *big.Int) (uint32, error) {
	return _Ecliptic.Contract.GetSpawnLimit(&_Ecliptic.CallOpts, _point, _time)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool result)
func (_Ecliptic *EclipticCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "isApprovedForAll", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool result)
func (_Ecliptic *EclipticSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Ecliptic.Contract.IsApprovedForAll(&_Ecliptic.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool result)
func (_Ecliptic *EclipticCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Ecliptic.Contract.IsApprovedForAll(&_Ecliptic.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ecliptic *EclipticCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ecliptic *EclipticSession) Name() (string, error) {
	return _Ecliptic.Contract.Name(&_Ecliptic.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ecliptic *EclipticCallerSession) Name() (string, error) {
	return _Ecliptic.Contract.Name(&_Ecliptic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ecliptic *EclipticCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ecliptic *EclipticSession) Owner() (common.Address, error) {
	return _Ecliptic.Contract.Owner(&_Ecliptic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ecliptic *EclipticCallerSession) Owner() (common.Address, error) {
	return _Ecliptic.Contract.Owner(&_Ecliptic.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_Ecliptic *EclipticCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "ownerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_Ecliptic *EclipticSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Ecliptic.Contract.OwnerOf(&_Ecliptic.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_Ecliptic *EclipticCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Ecliptic.Contract.OwnerOf(&_Ecliptic.CallOpts, _tokenId)
}

// Polls is a free data retrieval call binding the contract method 0xe64853c4.
//
// Solidity: function polls() view returns(address)
func (_Ecliptic *EclipticCaller) Polls(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "polls")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Polls is a free data retrieval call binding the contract method 0xe64853c4.
//
// Solidity: function polls() view returns(address)
func (_Ecliptic *EclipticSession) Polls() (common.Address, error) {
	return _Ecliptic.Contract.Polls(&_Ecliptic.CallOpts)
}

// Polls is a free data retrieval call binding the contract method 0xe64853c4.
//
// Solidity: function polls() view returns(address)
func (_Ecliptic *EclipticCallerSession) Polls() (common.Address, error) {
	return _Ecliptic.Contract.Polls(&_Ecliptic.CallOpts)
}

// PreviousEcliptic is a free data retrieval call binding the contract method 0xed969f68.
//
// Solidity: function previousEcliptic() view returns(address)
func (_Ecliptic *EclipticCaller) PreviousEcliptic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "previousEcliptic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreviousEcliptic is a free data retrieval call binding the contract method 0xed969f68.
//
// Solidity: function previousEcliptic() view returns(address)
func (_Ecliptic *EclipticSession) PreviousEcliptic() (common.Address, error) {
	return _Ecliptic.Contract.PreviousEcliptic(&_Ecliptic.CallOpts)
}

// PreviousEcliptic is a free data retrieval call binding the contract method 0xed969f68.
//
// Solidity: function previousEcliptic() view returns(address)
func (_Ecliptic *EclipticCallerSession) PreviousEcliptic() (common.Address, error) {
	return _Ecliptic.Contract.PreviousEcliptic(&_Ecliptic.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Ecliptic *EclipticCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Ecliptic *EclipticSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Ecliptic.Contract.SupportsInterface(&_Ecliptic.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Ecliptic *EclipticCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Ecliptic.Contract.SupportsInterface(&_Ecliptic.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ecliptic *EclipticCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ecliptic *EclipticSession) Symbol() (string, error) {
	return _Ecliptic.Contract.Symbol(&_Ecliptic.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ecliptic *EclipticCallerSession) Symbol() (string, error) {
	return _Ecliptic.Contract.Symbol(&_Ecliptic.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string _tokenURI)
func (_Ecliptic *EclipticCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string _tokenURI)
func (_Ecliptic *EclipticSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Ecliptic.Contract.TokenURI(&_Ecliptic.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string _tokenURI)
func (_Ecliptic *EclipticCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Ecliptic.Contract.TokenURI(&_Ecliptic.CallOpts, _tokenId)
}

// TreasuryProxy is a free data retrieval call binding the contract method 0x5d978826.
//
// Solidity: function treasuryProxy() view returns(address)
func (_Ecliptic *EclipticCaller) TreasuryProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "treasuryProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryProxy is a free data retrieval call binding the contract method 0x5d978826.
//
// Solidity: function treasuryProxy() view returns(address)
func (_Ecliptic *EclipticSession) TreasuryProxy() (common.Address, error) {
	return _Ecliptic.Contract.TreasuryProxy(&_Ecliptic.CallOpts)
}

// TreasuryProxy is a free data retrieval call binding the contract method 0x5d978826.
//
// Solidity: function treasuryProxy() view returns(address)
func (_Ecliptic *EclipticCallerSession) TreasuryProxy() (common.Address, error) {
	return _Ecliptic.Contract.TreasuryProxy(&_Ecliptic.CallOpts)
}

// TreasuryUpgradeHash is a free data retrieval call binding the contract method 0x39d42b15.
//
// Solidity: function treasuryUpgradeHash() view returns(bytes32)
func (_Ecliptic *EclipticCaller) TreasuryUpgradeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "treasuryUpgradeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TreasuryUpgradeHash is a free data retrieval call binding the contract method 0x39d42b15.
//
// Solidity: function treasuryUpgradeHash() view returns(bytes32)
func (_Ecliptic *EclipticSession) TreasuryUpgradeHash() ([32]byte, error) {
	return _Ecliptic.Contract.TreasuryUpgradeHash(&_Ecliptic.CallOpts)
}

// TreasuryUpgradeHash is a free data retrieval call binding the contract method 0x39d42b15.
//
// Solidity: function treasuryUpgradeHash() view returns(bytes32)
func (_Ecliptic *EclipticCallerSession) TreasuryUpgradeHash() ([32]byte, error) {
	return _Ecliptic.Contract.TreasuryUpgradeHash(&_Ecliptic.CallOpts)
}

// TreasuryUpgraded is a free data retrieval call binding the contract method 0x9e6e4028.
//
// Solidity: function treasuryUpgraded() view returns(bool)
func (_Ecliptic *EclipticCaller) TreasuryUpgraded(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ecliptic.contract.Call(opts, &out, "treasuryUpgraded")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TreasuryUpgraded is a free data retrieval call binding the contract method 0x9e6e4028.
//
// Solidity: function treasuryUpgraded() view returns(bool)
func (_Ecliptic *EclipticSession) TreasuryUpgraded() (bool, error) {
	return _Ecliptic.Contract.TreasuryUpgraded(&_Ecliptic.CallOpts)
}

// TreasuryUpgraded is a free data retrieval call binding the contract method 0x9e6e4028.
//
// Solidity: function treasuryUpgraded() view returns(bool)
func (_Ecliptic *EclipticCallerSession) TreasuryUpgraded() (bool, error) {
	return _Ecliptic.Contract.TreasuryUpgraded(&_Ecliptic.CallOpts)
}

// Adopt is a paid mutator transaction binding the contract method 0xc1b9d98b.
//
// Solidity: function adopt(uint32 _point) returns()
func (_Ecliptic *EclipticTransactor) Adopt(opts *bind.TransactOpts, _point uint32) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "adopt", _point)
}

// Adopt is a paid mutator transaction binding the contract method 0xc1b9d98b.
//
// Solidity: function adopt(uint32 _point) returns()
func (_Ecliptic *EclipticSession) Adopt(_point uint32) (*types.Transaction, error) {
	return _Ecliptic.Contract.Adopt(&_Ecliptic.TransactOpts, _point)
}

// Adopt is a paid mutator transaction binding the contract method 0xc1b9d98b.
//
// Solidity: function adopt(uint32 _point) returns()
func (_Ecliptic *EclipticTransactorSession) Adopt(_point uint32) (*types.Transaction, error) {
	return _Ecliptic.Contract.Adopt(&_Ecliptic.TransactOpts, _point)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) returns()
func (_Ecliptic *EclipticTransactor) Approve(opts *bind.TransactOpts, _approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "approve", _approved, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) returns()
func (_Ecliptic *EclipticSession) Approve(_approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ecliptic.Contract.Approve(&_Ecliptic.TransactOpts, _approved, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) returns()
func (_Ecliptic *EclipticTransactorSession) Approve(_approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ecliptic.Contract.Approve(&_Ecliptic.TransactOpts, _approved, _tokenId)
}

// CancelEscape is a paid mutator transaction binding the contract method 0xc6d761d4.
//
// Solidity: function cancelEscape(uint32 _point) returns()
func (_Ecliptic *EclipticTransactor) CancelEscape(opts *bind.TransactOpts, _point uint32) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "cancelEscape", _point)
}

// CancelEscape is a paid mutator transaction binding the contract method 0xc6d761d4.
//
// Solidity: function cancelEscape(uint32 _point) returns()
func (_Ecliptic *EclipticSession) CancelEscape(_point uint32) (*types.Transaction, error) {
	return _Ecliptic.Contract.CancelEscape(&_Ecliptic.TransactOpts, _point)
}

// CancelEscape is a paid mutator transaction binding the contract method 0xc6d761d4.
//
// Solidity: function cancelEscape(uint32 _point) returns()
func (_Ecliptic *EclipticTransactorSession) CancelEscape(_point uint32) (*types.Transaction, error) {
	return _Ecliptic.Contract.CancelEscape(&_Ecliptic.TransactOpts, _point)
}

// CastDocumentVote is a paid mutator transaction binding the contract method 0x6eb58224.
//
// Solidity: function castDocumentVote(uint8 _galaxy, bytes32 _proposal, bool _vote) returns()
func (_Ecliptic *EclipticTransactor) CastDocumentVote(opts *bind.TransactOpts, _galaxy uint8, _proposal [32]byte, _vote bool) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "castDocumentVote", _galaxy, _proposal, _vote)
}

// CastDocumentVote is a paid mutator transaction binding the contract method 0x6eb58224.
//
// Solidity: function castDocumentVote(uint8 _galaxy, bytes32 _proposal, bool _vote) returns()
func (_Ecliptic *EclipticSession) CastDocumentVote(_galaxy uint8, _proposal [32]byte, _vote bool) (*types.Transaction, error) {
	return _Ecliptic.Contract.CastDocumentVote(&_Ecliptic.TransactOpts, _galaxy, _proposal, _vote)
}

// CastDocumentVote is a paid mutator transaction binding the contract method 0x6eb58224.
//
// Solidity: function castDocumentVote(uint8 _galaxy, bytes32 _proposal, bool _vote) returns()
func (_Ecliptic *EclipticTransactorSession) CastDocumentVote(_galaxy uint8, _proposal [32]byte, _vote bool) (*types.Transaction, error) {
	return _Ecliptic.Contract.CastDocumentVote(&_Ecliptic.TransactOpts, _galaxy, _proposal, _vote)
}

// CastUpgradeVote is a paid mutator transaction binding the contract method 0x51de5541.
//
// Solidity: function castUpgradeVote(uint8 _galaxy, address _proposal, bool _vote) returns()
func (_Ecliptic *EclipticTransactor) CastUpgradeVote(opts *bind.TransactOpts, _galaxy uint8, _proposal common.Address, _vote bool) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "castUpgradeVote", _galaxy, _proposal, _vote)
}

// CastUpgradeVote is a paid mutator transaction binding the contract method 0x51de5541.
//
// Solidity: function castUpgradeVote(uint8 _galaxy, address _proposal, bool _vote) returns()
func (_Ecliptic *EclipticSession) CastUpgradeVote(_galaxy uint8, _proposal common.Address, _vote bool) (*types.Transaction, error) {
	return _Ecliptic.Contract.CastUpgradeVote(&_Ecliptic.TransactOpts, _galaxy, _proposal, _vote)
}

// CastUpgradeVote is a paid mutator transaction binding the contract method 0x51de5541.
//
// Solidity: function castUpgradeVote(uint8 _galaxy, address _proposal, bool _vote) returns()
func (_Ecliptic *EclipticTransactorSession) CastUpgradeVote(_galaxy uint8, _proposal common.Address, _vote bool) (*types.Transaction, error) {
	return _Ecliptic.Contract.CastUpgradeVote(&_Ecliptic.TransactOpts, _galaxy, _proposal, _vote)
}

// ConfigureKeys is a paid mutator transaction binding the contract method 0x4447e48c.
//
// Solidity: function configureKeys(uint32 _point, bytes32 _encryptionKey, bytes32 _authenticationKey, uint32 _cryptoSuiteVersion, bool _discontinuous) returns()
func (_Ecliptic *EclipticTransactor) ConfigureKeys(opts *bind.TransactOpts, _point uint32, _encryptionKey [32]byte, _authenticationKey [32]byte, _cryptoSuiteVersion uint32, _discontinuous bool) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "configureKeys", _point, _encryptionKey, _authenticationKey, _cryptoSuiteVersion, _discontinuous)
}

// ConfigureKeys is a paid mutator transaction binding the contract method 0x4447e48c.
//
// Solidity: function configureKeys(uint32 _point, bytes32 _encryptionKey, bytes32 _authenticationKey, uint32 _cryptoSuiteVersion, bool _discontinuous) returns()
func (_Ecliptic *EclipticSession) ConfigureKeys(_point uint32, _encryptionKey [32]byte, _authenticationKey [32]byte, _cryptoSuiteVersion uint32, _discontinuous bool) (*types.Transaction, error) {
	return _Ecliptic.Contract.ConfigureKeys(&_Ecliptic.TransactOpts, _point, _encryptionKey, _authenticationKey, _cryptoSuiteVersion, _discontinuous)
}

// ConfigureKeys is a paid mutator transaction binding the contract method 0x4447e48c.
//
// Solidity: function configureKeys(uint32 _point, bytes32 _encryptionKey, bytes32 _authenticationKey, uint32 _cryptoSuiteVersion, bool _discontinuous) returns()
func (_Ecliptic *EclipticTransactorSession) ConfigureKeys(_point uint32, _encryptionKey [32]byte, _authenticationKey [32]byte, _cryptoSuiteVersion uint32, _discontinuous bool) (*types.Transaction, error) {
	return _Ecliptic.Contract.ConfigureKeys(&_Ecliptic.TransactOpts, _point, _encryptionKey, _authenticationKey, _cryptoSuiteVersion, _discontinuous)
}

// CreateGalaxy is a paid mutator transaction binding the contract method 0x26295b52.
//
// Solidity: function createGalaxy(uint8 _galaxy, address _target) returns()
func (_Ecliptic *EclipticTransactor) CreateGalaxy(opts *bind.TransactOpts, _galaxy uint8, _target common.Address) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "createGalaxy", _galaxy, _target)
}

// CreateGalaxy is a paid mutator transaction binding the contract method 0x26295b52.
//
// Solidity: function createGalaxy(uint8 _galaxy, address _target) returns()
func (_Ecliptic *EclipticSession) CreateGalaxy(_galaxy uint8, _target common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.CreateGalaxy(&_Ecliptic.TransactOpts, _galaxy, _target)
}

// CreateGalaxy is a paid mutator transaction binding the contract method 0x26295b52.
//
// Solidity: function createGalaxy(uint8 _galaxy, address _target) returns()
func (_Ecliptic *EclipticTransactorSession) CreateGalaxy(_galaxy uint8, _target common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.CreateGalaxy(&_Ecliptic.TransactOpts, _galaxy, _target)
}

// Detach is a paid mutator transaction binding the contract method 0x073a7804.
//
// Solidity: function detach(uint32 _point) returns()
func (_Ecliptic *EclipticTransactor) Detach(opts *bind.TransactOpts, _point uint32) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "detach", _point)
}

// Detach is a paid mutator transaction binding the contract method 0x073a7804.
//
// Solidity: function detach(uint32 _point) returns()
func (_Ecliptic *EclipticSession) Detach(_point uint32) (*types.Transaction, error) {
	return _Ecliptic.Contract.Detach(&_Ecliptic.TransactOpts, _point)
}

// Detach is a paid mutator transaction binding the contract method 0x073a7804.
//
// Solidity: function detach(uint32 _point) returns()
func (_Ecliptic *EclipticTransactorSession) Detach(_point uint32) (*types.Transaction, error) {
	return _Ecliptic.Contract.Detach(&_Ecliptic.TransactOpts, _point)
}

// Escape is a paid mutator transaction binding the contract method 0xbf5772b9.
//
// Solidity: function escape(uint32 _point, uint32 _sponsor) returns()
func (_Ecliptic *EclipticTransactor) Escape(opts *bind.TransactOpts, _point uint32, _sponsor uint32) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "escape", _point, _sponsor)
}

// Escape is a paid mutator transaction binding the contract method 0xbf5772b9.
//
// Solidity: function escape(uint32 _point, uint32 _sponsor) returns()
func (_Ecliptic *EclipticSession) Escape(_point uint32, _sponsor uint32) (*types.Transaction, error) {
	return _Ecliptic.Contract.Escape(&_Ecliptic.TransactOpts, _point, _sponsor)
}

// Escape is a paid mutator transaction binding the contract method 0xbf5772b9.
//
// Solidity: function escape(uint32 _point, uint32 _sponsor) returns()
func (_Ecliptic *EclipticTransactorSession) Escape(_point uint32, _sponsor uint32) (*types.Transaction, error) {
	return _Ecliptic.Contract.Escape(&_Ecliptic.TransactOpts, _point, _sponsor)
}

// OnUpgrade is a paid mutator transaction binding the contract method 0x1824a46b.
//
// Solidity: function onUpgrade() returns()
func (_Ecliptic *EclipticTransactor) OnUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "onUpgrade")
}

// OnUpgrade is a paid mutator transaction binding the contract method 0x1824a46b.
//
// Solidity: function onUpgrade() returns()
func (_Ecliptic *EclipticSession) OnUpgrade() (*types.Transaction, error) {
	return _Ecliptic.Contract.OnUpgrade(&_Ecliptic.TransactOpts)
}

// OnUpgrade is a paid mutator transaction binding the contract method 0x1824a46b.
//
// Solidity: function onUpgrade() returns()
func (_Ecliptic *EclipticTransactorSession) OnUpgrade() (*types.Transaction, error) {
	return _Ecliptic.Contract.OnUpgrade(&_Ecliptic.TransactOpts)
}

// Reject is a paid mutator transaction binding the contract method 0xbbe21ca5.
//
// Solidity: function reject(uint32 _point) returns()
func (_Ecliptic *EclipticTransactor) Reject(opts *bind.TransactOpts, _point uint32) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "reject", _point)
}

// Reject is a paid mutator transaction binding the contract method 0xbbe21ca5.
//
// Solidity: function reject(uint32 _point) returns()
func (_Ecliptic *EclipticSession) Reject(_point uint32) (*types.Transaction, error) {
	return _Ecliptic.Contract.Reject(&_Ecliptic.TransactOpts, _point)
}

// Reject is a paid mutator transaction binding the contract method 0xbbe21ca5.
//
// Solidity: function reject(uint32 _point) returns()
func (_Ecliptic *EclipticTransactorSession) Reject(_point uint32) (*types.Transaction, error) {
	return _Ecliptic.Contract.Reject(&_Ecliptic.TransactOpts, _point)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ecliptic *EclipticTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ecliptic *EclipticSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ecliptic.Contract.RenounceOwnership(&_Ecliptic.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ecliptic *EclipticTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ecliptic.Contract.RenounceOwnership(&_Ecliptic.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Ecliptic *EclipticTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Ecliptic *EclipticSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ecliptic.Contract.SafeTransferFrom(&_Ecliptic.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Ecliptic *EclipticTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ecliptic.Contract.SafeTransferFrom(&_Ecliptic.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Ecliptic *EclipticTransactor) SafeTransferFrom0(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "safeTransferFrom0", _from, _to, _tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Ecliptic *EclipticSession) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ecliptic.Contract.SafeTransferFrom0(&_Ecliptic.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Ecliptic *EclipticTransactorSession) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ecliptic.Contract.SafeTransferFrom0(&_Ecliptic.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Ecliptic *EclipticTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Ecliptic *EclipticSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Ecliptic.Contract.SetApprovalForAll(&_Ecliptic.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Ecliptic *EclipticTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Ecliptic.Contract.SetApprovalForAll(&_Ecliptic.TransactOpts, _operator, _approved)
}

// SetDnsDomains is a paid mutator transaction binding the contract method 0xbac55edd.
//
// Solidity: function setDnsDomains(string _primary, string _secondary, string _tertiary) returns()
func (_Ecliptic *EclipticTransactor) SetDnsDomains(opts *bind.TransactOpts, _primary string, _secondary string, _tertiary string) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "setDnsDomains", _primary, _secondary, _tertiary)
}

// SetDnsDomains is a paid mutator transaction binding the contract method 0xbac55edd.
//
// Solidity: function setDnsDomains(string _primary, string _secondary, string _tertiary) returns()
func (_Ecliptic *EclipticSession) SetDnsDomains(_primary string, _secondary string, _tertiary string) (*types.Transaction, error) {
	return _Ecliptic.Contract.SetDnsDomains(&_Ecliptic.TransactOpts, _primary, _secondary, _tertiary)
}

// SetDnsDomains is a paid mutator transaction binding the contract method 0xbac55edd.
//
// Solidity: function setDnsDomains(string _primary, string _secondary, string _tertiary) returns()
func (_Ecliptic *EclipticTransactorSession) SetDnsDomains(_primary string, _secondary string, _tertiary string) (*types.Transaction, error) {
	return _Ecliptic.Contract.SetDnsDomains(&_Ecliptic.TransactOpts, _primary, _secondary, _tertiary)
}

// SetManagementProxy is a paid mutator transaction binding the contract method 0x8866bb2c.
//
// Solidity: function setManagementProxy(uint32 _point, address _manager) returns()
func (_Ecliptic *EclipticTransactor) SetManagementProxy(opts *bind.TransactOpts, _point uint32, _manager common.Address) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "setManagementProxy", _point, _manager)
}

// SetManagementProxy is a paid mutator transaction binding the contract method 0x8866bb2c.
//
// Solidity: function setManagementProxy(uint32 _point, address _manager) returns()
func (_Ecliptic *EclipticSession) SetManagementProxy(_point uint32, _manager common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.SetManagementProxy(&_Ecliptic.TransactOpts, _point, _manager)
}

// SetManagementProxy is a paid mutator transaction binding the contract method 0x8866bb2c.
//
// Solidity: function setManagementProxy(uint32 _point, address _manager) returns()
func (_Ecliptic *EclipticTransactorSession) SetManagementProxy(_point uint32, _manager common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.SetManagementProxy(&_Ecliptic.TransactOpts, _point, _manager)
}

// SetSpawnProxy is a paid mutator transaction binding the contract method 0xae326221.
//
// Solidity: function setSpawnProxy(uint16 _prefix, address _spawnProxy) returns()
func (_Ecliptic *EclipticTransactor) SetSpawnProxy(opts *bind.TransactOpts, _prefix uint16, _spawnProxy common.Address) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "setSpawnProxy", _prefix, _spawnProxy)
}

// SetSpawnProxy is a paid mutator transaction binding the contract method 0xae326221.
//
// Solidity: function setSpawnProxy(uint16 _prefix, address _spawnProxy) returns()
func (_Ecliptic *EclipticSession) SetSpawnProxy(_prefix uint16, _spawnProxy common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.SetSpawnProxy(&_Ecliptic.TransactOpts, _prefix, _spawnProxy)
}

// SetSpawnProxy is a paid mutator transaction binding the contract method 0xae326221.
//
// Solidity: function setSpawnProxy(uint16 _prefix, address _spawnProxy) returns()
func (_Ecliptic *EclipticTransactorSession) SetSpawnProxy(_prefix uint16, _spawnProxy common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.SetSpawnProxy(&_Ecliptic.TransactOpts, _prefix, _spawnProxy)
}

// SetTransferProxy is a paid mutator transaction binding the contract method 0x2c7ba564.
//
// Solidity: function setTransferProxy(uint32 _point, address _transferProxy) returns()
func (_Ecliptic *EclipticTransactor) SetTransferProxy(opts *bind.TransactOpts, _point uint32, _transferProxy common.Address) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "setTransferProxy", _point, _transferProxy)
}

// SetTransferProxy is a paid mutator transaction binding the contract method 0x2c7ba564.
//
// Solidity: function setTransferProxy(uint32 _point, address _transferProxy) returns()
func (_Ecliptic *EclipticSession) SetTransferProxy(_point uint32, _transferProxy common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.SetTransferProxy(&_Ecliptic.TransactOpts, _point, _transferProxy)
}

// SetTransferProxy is a paid mutator transaction binding the contract method 0x2c7ba564.
//
// Solidity: function setTransferProxy(uint32 _point, address _transferProxy) returns()
func (_Ecliptic *EclipticTransactorSession) SetTransferProxy(_point uint32, _transferProxy common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.SetTransferProxy(&_Ecliptic.TransactOpts, _point, _transferProxy)
}

// SetVotingProxy is a paid mutator transaction binding the contract method 0xa60e8bd6.
//
// Solidity: function setVotingProxy(uint8 _galaxy, address _voter) returns()
func (_Ecliptic *EclipticTransactor) SetVotingProxy(opts *bind.TransactOpts, _galaxy uint8, _voter common.Address) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "setVotingProxy", _galaxy, _voter)
}

// SetVotingProxy is a paid mutator transaction binding the contract method 0xa60e8bd6.
//
// Solidity: function setVotingProxy(uint8 _galaxy, address _voter) returns()
func (_Ecliptic *EclipticSession) SetVotingProxy(_galaxy uint8, _voter common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.SetVotingProxy(&_Ecliptic.TransactOpts, _galaxy, _voter)
}

// SetVotingProxy is a paid mutator transaction binding the contract method 0xa60e8bd6.
//
// Solidity: function setVotingProxy(uint8 _galaxy, address _voter) returns()
func (_Ecliptic *EclipticTransactorSession) SetVotingProxy(_galaxy uint8, _voter common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.SetVotingProxy(&_Ecliptic.TransactOpts, _galaxy, _voter)
}

// Spawn is a paid mutator transaction binding the contract method 0xa0d3253f.
//
// Solidity: function spawn(uint32 _point, address _target) returns()
func (_Ecliptic *EclipticTransactor) Spawn(opts *bind.TransactOpts, _point uint32, _target common.Address) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "spawn", _point, _target)
}

// Spawn is a paid mutator transaction binding the contract method 0xa0d3253f.
//
// Solidity: function spawn(uint32 _point, address _target) returns()
func (_Ecliptic *EclipticSession) Spawn(_point uint32, _target common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.Spawn(&_Ecliptic.TransactOpts, _point, _target)
}

// Spawn is a paid mutator transaction binding the contract method 0xa0d3253f.
//
// Solidity: function spawn(uint32 _point, address _target) returns()
func (_Ecliptic *EclipticTransactorSession) Spawn(_point uint32, _target common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.Spawn(&_Ecliptic.TransactOpts, _point, _target)
}

// StartDocumentPoll is a paid mutator transaction binding the contract method 0x05bbf5db.
//
// Solidity: function startDocumentPoll(uint8 _galaxy, bytes32 _proposal) returns()
func (_Ecliptic *EclipticTransactor) StartDocumentPoll(opts *bind.TransactOpts, _galaxy uint8, _proposal [32]byte) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "startDocumentPoll", _galaxy, _proposal)
}

// StartDocumentPoll is a paid mutator transaction binding the contract method 0x05bbf5db.
//
// Solidity: function startDocumentPoll(uint8 _galaxy, bytes32 _proposal) returns()
func (_Ecliptic *EclipticSession) StartDocumentPoll(_galaxy uint8, _proposal [32]byte) (*types.Transaction, error) {
	return _Ecliptic.Contract.StartDocumentPoll(&_Ecliptic.TransactOpts, _galaxy, _proposal)
}

// StartDocumentPoll is a paid mutator transaction binding the contract method 0x05bbf5db.
//
// Solidity: function startDocumentPoll(uint8 _galaxy, bytes32 _proposal) returns()
func (_Ecliptic *EclipticTransactorSession) StartDocumentPoll(_galaxy uint8, _proposal [32]byte) (*types.Transaction, error) {
	return _Ecliptic.Contract.StartDocumentPoll(&_Ecliptic.TransactOpts, _galaxy, _proposal)
}

// StartUpgradePoll is a paid mutator transaction binding the contract method 0x9e988d13.
//
// Solidity: function startUpgradePoll(uint8 _galaxy, address _proposal) returns()
func (_Ecliptic *EclipticTransactor) StartUpgradePoll(opts *bind.TransactOpts, _galaxy uint8, _proposal common.Address) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "startUpgradePoll", _galaxy, _proposal)
}

// StartUpgradePoll is a paid mutator transaction binding the contract method 0x9e988d13.
//
// Solidity: function startUpgradePoll(uint8 _galaxy, address _proposal) returns()
func (_Ecliptic *EclipticSession) StartUpgradePoll(_galaxy uint8, _proposal common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.StartUpgradePoll(&_Ecliptic.TransactOpts, _galaxy, _proposal)
}

// StartUpgradePoll is a paid mutator transaction binding the contract method 0x9e988d13.
//
// Solidity: function startUpgradePoll(uint8 _galaxy, address _proposal) returns()
func (_Ecliptic *EclipticTransactorSession) StartUpgradePoll(_galaxy uint8, _proposal common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.StartUpgradePoll(&_Ecliptic.TransactOpts, _galaxy, _proposal)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Ecliptic *EclipticTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Ecliptic *EclipticSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ecliptic.Contract.TransferFrom(&_Ecliptic.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Ecliptic *EclipticTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Ecliptic.Contract.TransferFrom(&_Ecliptic.TransactOpts, _from, _to, _tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Ecliptic *EclipticTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Ecliptic *EclipticSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.TransferOwnership(&_Ecliptic.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Ecliptic *EclipticTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.TransferOwnership(&_Ecliptic.TransactOpts, _newOwner)
}

// TransferPoint is a paid mutator transaction binding the contract method 0x1e79a85b.
//
// Solidity: function transferPoint(uint32 _point, address _target, bool _reset) returns()
func (_Ecliptic *EclipticTransactor) TransferPoint(opts *bind.TransactOpts, _point uint32, _target common.Address, _reset bool) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "transferPoint", _point, _target, _reset)
}

// TransferPoint is a paid mutator transaction binding the contract method 0x1e79a85b.
//
// Solidity: function transferPoint(uint32 _point, address _target, bool _reset) returns()
func (_Ecliptic *EclipticSession) TransferPoint(_point uint32, _target common.Address, _reset bool) (*types.Transaction, error) {
	return _Ecliptic.Contract.TransferPoint(&_Ecliptic.TransactOpts, _point, _target, _reset)
}

// TransferPoint is a paid mutator transaction binding the contract method 0x1e79a85b.
//
// Solidity: function transferPoint(uint32 _point, address _target, bool _reset) returns()
func (_Ecliptic *EclipticTransactorSession) TransferPoint(_point uint32, _target common.Address, _reset bool) (*types.Transaction, error) {
	return _Ecliptic.Contract.TransferPoint(&_Ecliptic.TransactOpts, _point, _target, _reset)
}

// UpdateDocumentPoll is a paid mutator transaction binding the contract method 0x172a735c.
//
// Solidity: function updateDocumentPoll(bytes32 _proposal) returns()
func (_Ecliptic *EclipticTransactor) UpdateDocumentPoll(opts *bind.TransactOpts, _proposal [32]byte) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "updateDocumentPoll", _proposal)
}

// UpdateDocumentPoll is a paid mutator transaction binding the contract method 0x172a735c.
//
// Solidity: function updateDocumentPoll(bytes32 _proposal) returns()
func (_Ecliptic *EclipticSession) UpdateDocumentPoll(_proposal [32]byte) (*types.Transaction, error) {
	return _Ecliptic.Contract.UpdateDocumentPoll(&_Ecliptic.TransactOpts, _proposal)
}

// UpdateDocumentPoll is a paid mutator transaction binding the contract method 0x172a735c.
//
// Solidity: function updateDocumentPoll(bytes32 _proposal) returns()
func (_Ecliptic *EclipticTransactorSession) UpdateDocumentPoll(_proposal [32]byte) (*types.Transaction, error) {
	return _Ecliptic.Contract.UpdateDocumentPoll(&_Ecliptic.TransactOpts, _proposal)
}

// UpdateUpgradePoll is a paid mutator transaction binding the contract method 0x5623715b.
//
// Solidity: function updateUpgradePoll(address _proposal) returns()
func (_Ecliptic *EclipticTransactor) UpdateUpgradePoll(opts *bind.TransactOpts, _proposal common.Address) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "updateUpgradePoll", _proposal)
}

// UpdateUpgradePoll is a paid mutator transaction binding the contract method 0x5623715b.
//
// Solidity: function updateUpgradePoll(address _proposal) returns()
func (_Ecliptic *EclipticSession) UpdateUpgradePoll(_proposal common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.UpdateUpgradePoll(&_Ecliptic.TransactOpts, _proposal)
}

// UpdateUpgradePoll is a paid mutator transaction binding the contract method 0x5623715b.
//
// Solidity: function updateUpgradePoll(address _proposal) returns()
func (_Ecliptic *EclipticTransactorSession) UpdateUpgradePoll(_proposal common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.UpdateUpgradePoll(&_Ecliptic.TransactOpts, _proposal)
}

// UpgradeTreasury is a paid mutator transaction binding the contract method 0x4e7e9299.
//
// Solidity: function upgradeTreasury(address _treasuryImpl) returns()
func (_Ecliptic *EclipticTransactor) UpgradeTreasury(opts *bind.TransactOpts, _treasuryImpl common.Address) (*types.Transaction, error) {
	return _Ecliptic.contract.Transact(opts, "upgradeTreasury", _treasuryImpl)
}

// UpgradeTreasury is a paid mutator transaction binding the contract method 0x4e7e9299.
//
// Solidity: function upgradeTreasury(address _treasuryImpl) returns()
func (_Ecliptic *EclipticSession) UpgradeTreasury(_treasuryImpl common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.UpgradeTreasury(&_Ecliptic.TransactOpts, _treasuryImpl)
}

// UpgradeTreasury is a paid mutator transaction binding the contract method 0x4e7e9299.
//
// Solidity: function upgradeTreasury(address _treasuryImpl) returns()
func (_Ecliptic *EclipticTransactorSession) UpgradeTreasury(_treasuryImpl common.Address) (*types.Transaction, error) {
	return _Ecliptic.Contract.UpgradeTreasury(&_Ecliptic.TransactOpts, _treasuryImpl)
}

// EclipticApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ecliptic contract.
type EclipticApprovalIterator struct {
	Event *EclipticApproval // Event containing the contract specifics and raw log

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
func (it *EclipticApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EclipticApproval)
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
		it.Event = new(EclipticApproval)
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
func (it *EclipticApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EclipticApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EclipticApproval represents a Approval event raised by the Ecliptic contract.
type EclipticApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Ecliptic *EclipticFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*EclipticApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Ecliptic.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &EclipticApprovalIterator{contract: _Ecliptic.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Ecliptic *EclipticFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EclipticApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Ecliptic.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EclipticApproval)
				if err := _Ecliptic.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Ecliptic *EclipticFilterer) ParseApproval(log types.Log) (*EclipticApproval, error) {
	event := new(EclipticApproval)
	if err := _Ecliptic.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EclipticApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Ecliptic contract.
type EclipticApprovalForAllIterator struct {
	Event *EclipticApprovalForAll // Event containing the contract specifics and raw log

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
func (it *EclipticApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EclipticApprovalForAll)
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
		it.Event = new(EclipticApprovalForAll)
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
func (it *EclipticApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EclipticApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EclipticApprovalForAll represents a ApprovalForAll event raised by the Ecliptic contract.
type EclipticApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Ecliptic *EclipticFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*EclipticApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Ecliptic.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &EclipticApprovalForAllIterator{contract: _Ecliptic.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Ecliptic *EclipticFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *EclipticApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Ecliptic.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EclipticApprovalForAll)
				if err := _Ecliptic.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Ecliptic *EclipticFilterer) ParseApprovalForAll(log types.Log) (*EclipticApprovalForAll, error) {
	event := new(EclipticApprovalForAll)
	if err := _Ecliptic.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EclipticOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Ecliptic contract.
type EclipticOwnershipRenouncedIterator struct {
	Event *EclipticOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *EclipticOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EclipticOwnershipRenounced)
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
		it.Event = new(EclipticOwnershipRenounced)
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
func (it *EclipticOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EclipticOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EclipticOwnershipRenounced represents a OwnershipRenounced event raised by the Ecliptic contract.
type EclipticOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Ecliptic *EclipticFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*EclipticOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ecliptic.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EclipticOwnershipRenouncedIterator{contract: _Ecliptic.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Ecliptic *EclipticFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *EclipticOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ecliptic.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EclipticOwnershipRenounced)
				if err := _Ecliptic.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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
func (_Ecliptic *EclipticFilterer) ParseOwnershipRenounced(log types.Log) (*EclipticOwnershipRenounced, error) {
	event := new(EclipticOwnershipRenounced)
	if err := _Ecliptic.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EclipticOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ecliptic contract.
type EclipticOwnershipTransferredIterator struct {
	Event *EclipticOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EclipticOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EclipticOwnershipTransferred)
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
		it.Event = new(EclipticOwnershipTransferred)
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
func (it *EclipticOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EclipticOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EclipticOwnershipTransferred represents a OwnershipTransferred event raised by the Ecliptic contract.
type EclipticOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ecliptic *EclipticFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EclipticOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ecliptic.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EclipticOwnershipTransferredIterator{contract: _Ecliptic.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ecliptic *EclipticFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EclipticOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ecliptic.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EclipticOwnershipTransferred)
				if err := _Ecliptic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ecliptic *EclipticFilterer) ParseOwnershipTransferred(log types.Log) (*EclipticOwnershipTransferred, error) {
	event := new(EclipticOwnershipTransferred)
	if err := _Ecliptic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EclipticTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ecliptic contract.
type EclipticTransferIterator struct {
	Event *EclipticTransfer // Event containing the contract specifics and raw log

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
func (it *EclipticTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EclipticTransfer)
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
		it.Event = new(EclipticTransfer)
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
func (it *EclipticTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EclipticTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EclipticTransfer represents a Transfer event raised by the Ecliptic contract.
type EclipticTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Ecliptic *EclipticFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*EclipticTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Ecliptic.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &EclipticTransferIterator{contract: _Ecliptic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Ecliptic *EclipticFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EclipticTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Ecliptic.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EclipticTransfer)
				if err := _Ecliptic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Ecliptic *EclipticFilterer) ParseTransfer(log types.Log) (*EclipticTransfer, error) {
	event := new(EclipticTransfer)
	if err := _Ecliptic.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EclipticUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Ecliptic contract.
type EclipticUpgradedIterator struct {
	Event *EclipticUpgraded // Event containing the contract specifics and raw log

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
func (it *EclipticUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EclipticUpgraded)
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
		it.Event = new(EclipticUpgraded)
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
func (it *EclipticUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EclipticUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EclipticUpgraded represents a Upgraded event raised by the Ecliptic contract.
type EclipticUpgraded struct {
	To  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address to)
func (_Ecliptic *EclipticFilterer) FilterUpgraded(opts *bind.FilterOpts) (*EclipticUpgradedIterator, error) {

	logs, sub, err := _Ecliptic.contract.FilterLogs(opts, "Upgraded")
	if err != nil {
		return nil, err
	}
	return &EclipticUpgradedIterator{contract: _Ecliptic.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address to)
func (_Ecliptic *EclipticFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EclipticUpgraded) (event.Subscription, error) {

	logs, sub, err := _Ecliptic.contract.WatchLogs(opts, "Upgraded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EclipticUpgraded)
				if err := _Ecliptic.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address to)
func (_Ecliptic *EclipticFilterer) ParseUpgraded(log types.Log) (*EclipticUpgraded, error) {
	event := new(EclipticUpgraded)
	if err := _Ecliptic.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
