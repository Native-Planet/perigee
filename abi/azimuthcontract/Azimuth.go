// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package azimuth

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

// AzimuthMetaData contains all meta data concerning the Azimuth contract.
var AzimuthMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"escapeRequestsIndexes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_whose\",\"type\":\"address\"}],\"name\":\"getOwnedPoints\",\"outputs\":[{\"name\":\"ownedPoints\",\"type\":\"uint32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"votingFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"rights\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"managementProxy\",\"type\":\"address\"},{\"name\":\"spawnProxy\",\"type\":\"address\"},{\"name\":\"votingProxy\",\"type\":\"address\"},{\"name\":\"transferProxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferringFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_sponsor\",\"type\":\"uint32\"}],\"name\":\"isSponsor\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"getManagementProxy\",\"outputs\":[{\"name\":\"manager\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"getContinuityNumber\",\"outputs\":[{\"name\":\"continuityNumber\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sponsoring\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_whose\",\"type\":\"address\"}],\"name\":\"getOwnedPointCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"doEscape\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_whose\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getOwnedPointAtIndex\",\"outputs\":[{\"name\":\"point\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"getTransferProxy\",\"outputs\":[{\"name\":\"transferProxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"isSpawnProxy\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pointsOwnedBy\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"getSpawnCount\",\"outputs\":[{\"name\":\"spawnCount\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"setSpawnProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"setTransferProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"sponsoringIndexes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"isTransferProxy\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"getVotingProxy\",\"outputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"isManagementProxy\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"votingForIndexes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"isLive\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sponsor\",\"type\":\"uint32\"}],\"name\":\"getEscapeRequests\",\"outputs\":[{\"name\":\"requests\",\"type\":\"uint32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"getSponsor\",\"outputs\":[{\"name\":\"sponsor\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sponsor\",\"type\":\"uint32\"}],\"name\":\"getEscapeRequestsCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"pointOwnerIndexes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getManagerFor\",\"outputs\":[{\"name\":\"mfor\",\"type\":\"uint32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"managerForIndexes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"managerFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"isActive\",\"outputs\":[{\"name\":\"equals\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"points\",\"outputs\":[{\"name\":\"encryptionKey\",\"type\":\"bytes32\"},{\"name\":\"authenticationKey\",\"type\":\"bytes32\"},{\"name\":\"hasSponsor\",\"type\":\"bool\"},{\"name\":\"active\",\"type\":\"bool\"},{\"name\":\"escapeRequested\",\"type\":\"bool\"},{\"name\":\"sponsor\",\"type\":\"uint32\"},{\"name\":\"escapeRequestedTo\",\"type\":\"uint32\"},{\"name\":\"cryptoSuiteVersion\",\"type\":\"uint32\"},{\"name\":\"keyRevisionNumber\",\"type\":\"uint32\"},{\"name\":\"continuityNumber\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getSpawningForCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"hasBeenLinked\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"canTransfer\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"spawningForIndexes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"hasSponsor\",\"outputs\":[{\"name\":\"has\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"activatePoint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"getSpawned\",\"outputs\":[{\"name\":\"spawned\",\"type\":\"uint32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"setManagementProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"canSpawnAs\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"getKeyRevisionNumber\",\"outputs\":[{\"name\":\"revision\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"canManage\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getTransferringFor\",\"outputs\":[{\"name\":\"tfor\",\"type\":\"uint32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"getPointSize\",\"outputs\":[{\"name\":\"_size\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sponsor\",\"type\":\"uint32\"}],\"name\":\"getSponsoringCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getSpawningFor\",\"outputs\":[{\"name\":\"sfor\",\"type\":\"uint32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"isEscaping\",\"outputs\":[{\"name\":\"escaping\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"escapeRequests\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"setVotingProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sponsor\",\"type\":\"uint32\"}],\"name\":\"getSponsoring\",\"outputs\":[{\"name\":\"sponsees\",\"type\":\"uint32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_sponsor\",\"type\":\"uint32\"}],\"name\":\"setEscapeRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"canVoteAs\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_primary\",\"type\":\"string\"},{\"name\":\"_secondary\",\"type\":\"string\"},{\"name\":\"_tertiary\",\"type\":\"string\"}],\"name\":\"setDnsDomains\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"loseSponsor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"registerSpawned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"getKeys\",\"outputs\":[{\"name\":\"crypt\",\"type\":\"bytes32\"},{\"name\":\"auth\",\"type\":\"bytes32\"},{\"name\":\"suite\",\"type\":\"uint32\"},{\"name\":\"revision\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"cancelEscape\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getManagerForCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"transferringForIndexes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getVotingFor\",\"outputs\":[{\"name\":\"vfor\",\"type\":\"uint32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spawningFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getTransferringForCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"isVotingProxy\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"getPrefix\",\"outputs\":[{\"name\":\"prefix\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dnsDomains\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"incrementContinuityNumber\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"getSpawnProxy\",\"outputs\":[{\"name\":\"spawnProxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"}],\"name\":\"getEscapeRequest\",\"outputs\":[{\"name\":\"escape\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_encryptionKey\",\"type\":\"bytes32\"},{\"name\":\"_authenticationKey\",\"type\":\"bytes32\"},{\"name\":\"_cryptoSuiteVersion\",\"type\":\"uint32\"}],\"name\":\"setKeys\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getVotingForCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_point\",\"type\":\"uint32\"},{\"name\":\"_sponsor\",\"type\":\"uint32\"}],\"name\":\"isRequestingEscapeTo\",\"outputs\":[{\"name\":\"equals\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"point\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"point\",\"type\":\"uint32\"}],\"name\":\"Activated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"prefix\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"child\",\"type\":\"uint32\"}],\"name\":\"Spawned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"point\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"sponsor\",\"type\":\"uint32\"}],\"name\":\"EscapeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"point\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"sponsor\",\"type\":\"uint32\"}],\"name\":\"EscapeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"point\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"sponsor\",\"type\":\"uint32\"}],\"name\":\"EscapeAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"point\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"sponsor\",\"type\":\"uint32\"}],\"name\":\"LostSponsor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"point\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"encryptionKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"authenticationKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"cryptoSuiteVersion\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"keyRevisionNumber\",\"type\":\"uint32\"}],\"name\":\"ChangedKeys\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"point\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"number\",\"type\":\"uint32\"}],\"name\":\"BrokeContinuity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"point\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"spawnProxy\",\"type\":\"address\"}],\"name\":\"ChangedSpawnProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"point\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"transferProxy\",\"type\":\"address\"}],\"name\":\"ChangedTransferProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"point\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"managementProxy\",\"type\":\"address\"}],\"name\":\"ChangedManagementProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"point\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"votingProxy\",\"type\":\"address\"}],\"name\":\"ChangedVotingProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"primary\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"secondary\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"tertiary\",\"type\":\"string\"}],\"name\":\"ChangedDns\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506200010d6040805190810160405280600b81526020017f6578616d706c652e636f6d0000000000000000000000000000000000000000008152506040805190810160405280600b81526020017f6578616d706c652e636f6d0000000000000000000000000000000000000000008152506040805190810160405280600b81526020017f6578616d706c652e636f6d00000000000000000000000000000000000000000081525062000113640100000000026401000000009004565b6200041d565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156200016f57600080fd5b82600460006003811015156200018157fe5b019080519060200190620001979291906200036e565b508160046001600381101515620001aa57fe5b019080519060200190620001c09291906200036e565b508060046002600381101515620001d357fe5b019080519060200190620001e99291906200036e565b507ffafd04ade1daae2e1fdb0fc1cc6a899fd424063ed5c92120e67e073053b9489883838360405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b83811015620002585780820151818401526020810190506200023b565b50505050905090810190601f168015620002865780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b83811015620002c1578082015181840152602081019050620002a4565b50505050905090810190601f168015620002ef5780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b838110156200032a5780820151818401526020810190506200030d565b50505050905090810190601f168015620003585780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a1505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003b157805160ff1916838001178555620003e2565b82800160010185558215620003e2579182015b82811115620003e1578251825591602001919060010190620003c4565b5b509050620003f19190620003f5565b5090565b6200041a91905b8082111562000416576000816000905550600101620003fc565b5090565b90565b616fe3806200042d6000396000f3006080604052600436106103bc576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806301d8a97f146103c15780630333d4991461041857806304bbb61d146104b057806304bc94e51461051d578063073167da1461065c57806308083039146106c95780630f934e6614610724578063108d995b1461079757806312ace2e6146107ea57806312afbc7814610847578063130631801461089e57806322a7906f146108d157806324541f781461093e57806324ba1a46146109b157806324cf61b514610a1c5780632637588e14610a89578063293a916914610b045780632a19642c14610b575780632c7ba56414610baa5780632cce0e2914610bfd5780633262207c14610c5457806333f3372414610cbf578063375e24c114610d3257806337de45c114610d9d5780633b17b4ea14610e045780633c01a9fa14610e4f578063439f7d3c14610ed75780634448f5fb14610f2a578063474dbae714610f71578063483a15ed14610fd857806352c0683b146110705780635a83df54146110d75780635e19b30514611144578063621b23e21461118f57806363fa9a87146112025780636436b1b9146112e05780636d09887b14611337578063715018a614611382578063728aa8571461139957806376944d5b1461140457806377eb4c501461146b5780637bc702a1146114b6578063846f652b146114e95780638866bb2c146115715780638a27bf59146115c45780638d2d3f411461162f5780638da5cb5b146116825780639137fe0a146116d957806391398f251461174457806393976405146117dc578063942324c3146118315780639b0d1b2c146118785780639b350e12146119105780639dd901901461195b578063a297c1d8146119b8578063a46afef114611a0b578063a634585914611a93578063b6363cf214611ad6578063b65afedd14611b51578063bac55edd14611bbc578063bc562b9e14611cb1578063bc735d9014611ce4578063c17ad9d814611d53578063c210c62314611d86578063c6d761d414611e0a578063caf590f914611e3d578063d154126414611ea8578063d2588e2014611eff578063d2a9b6ed14611f66578063d5c5164f14611ffe578063ddc359501461206b578063e01cff84146120be578063e213c84014612115578063e4a358d714612180578063eccc8ff1146121cf578063f2fde38b14612275578063f4914919146122b8578063f4e3be2d146122eb578063f5af66211461235e578063f9b87d40146123b1578063fdb357be14612410578063fde9073314612467575b600080fd5b3480156103cd57600080fd5b50610402600480360381019080803563ffffffff169060200190929190803563ffffffff1690602001909291905050506124c2565b6040518082815260200191505060405180910390f35b34801561042457600080fd5b50610459600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506124e7565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561049c578082015181840152602081019050610481565b505050509050019250505060405180910390f35b3480156104bc57600080fd5b506104fb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506125aa565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b34801561052957600080fd5b5061054e600480360381019080803563ffffffff1690602001909291905050506125f2565b604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019550505050505060405180910390f35b34801561066857600080fd5b506106a7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506126c8565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b3480156106d557600080fd5b5061070a600480360381019080803563ffffffff169060200190929190803563ffffffff169060200190929190505050612710565b604051808215151515815260200191505060405180910390f35b34801561073057600080fd5b50610755600480360381019080803563ffffffff16906020019092919050505061277a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156107a357600080fd5b506107c8600480360381019080803563ffffffff1690602001909291905050506127c6565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b3480156107f657600080fd5b50610825600480360381019080803563ffffffff16906020019092919080359060200190929190505050612802565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b34801561085357600080fd5b50610888600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061284a565b6040518082815260200191505060405180910390f35b3480156108aa57600080fd5b506108cf600480360381019080803563ffffffff169060200190929190505050612896565b005b3480156108dd57600080fd5b5061091c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506129af565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b34801561094a57600080fd5b5061096f600480360381019080803563ffffffff169060200190929190505050612a42565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156109bd57600080fd5b50610a02600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612a8e565b604051808215151515815260200191505060405180910390f35b348015610a2857600080fd5b50610a67600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050612b09565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b348015610a9557600080fd5b50610aea600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612b51565b604051808215151515815260200191505060405180910390f35b348015610b1057600080fd5b50610b35600480360381019080803563ffffffff169060200190929190505050612b80565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b348015610b6357600080fd5b50610ba8600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612bc4565b005b348015610bb657600080fd5b50610bfb600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506130bd565b005b348015610c0957600080fd5b50610c3e600480360381019080803563ffffffff169060200190929190803563ffffffff1690602001909291905050506135b6565b6040518082815260200191505060405180910390f35b348015610c6057600080fd5b50610ca5600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506135db565b604051808215151515815260200191505060405180910390f35b348015610ccb57600080fd5b50610cf0600480360381019080803563ffffffff169060200190929190505050613656565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610d3e57600080fd5b50610d83600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506136a2565b604051808215151515815260200191505060405180910390f35b348015610da957600080fd5b50610dee600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803563ffffffff16906020019092919050505061371d565b6040518082815260200191505060405180910390f35b348015610e1057600080fd5b50610e35600480360381019080803563ffffffff169060200190929190505050613742565b604051808215151515815260200191505060405180910390f35b348015610e5b57600080fd5b50610e80600480360381019080803563ffffffff1690602001909291905050506137bd565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610ec3578082015181840152602081019050610ea8565b505050509050019250505060405180910390f35b348015610ee357600080fd5b50610f08600480360381019080803563ffffffff169060200190929190505050613860565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b348015610f3657600080fd5b50610f5b600480360381019080803563ffffffff16906020019092919050505061389c565b6040518082815260200191505060405180910390f35b348015610f7d57600080fd5b50610fc2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803563ffffffff1690602001909291905050506138c8565b6040518082815260200191505060405180910390f35b348015610fe457600080fd5b50611019600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506138ed565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561105c578082015181840152602081019050611041565b505050509050019250505060405180910390f35b34801561107c57600080fd5b506110c1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803563ffffffff1690602001909291905050506139b0565b6040518082815260200191505060405180910390f35b3480156110e357600080fd5b50611122600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506139d5565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b34801561115057600080fd5b50611175600480360381019080803563ffffffff169060200190929190505050613a1d565b604051808215151515815260200191505060405180910390f35b34801561119b57600080fd5b506111c0600480360381019080803563ffffffff169060200190929190505050613a56565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561120e57600080fd5b50611233600480360381019080803563ffffffff169060200190929190505050613aa2565b604051808b600019166000191681526020018a600019166000191681526020018915151515815260200188151515158152602001871515151581526020018663ffffffff1663ffffffff1681526020018563ffffffff1663ffffffff1681526020018463ffffffff1663ffffffff1681526020018363ffffffff1663ffffffff1681526020018263ffffffff1663ffffffff1681526020019a505050505050505050505060405180910390f35b3480156112ec57600080fd5b50611321600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613b6d565b6040518082815260200191505060405180910390f35b34801561134357600080fd5b50611368600480360381019080803563ffffffff169060200190929190505050613bb9565b604051808215151515815260200191505060405180910390f35b34801561138e57600080fd5b50611397613bfd565b005b3480156113a557600080fd5b506113ea600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613cff565b604051808215151515815260200191505060405180910390f35b34801561141057600080fd5b50611455600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803563ffffffff169060200190929190505050613eb1565b6040518082815260200191505060405180910390f35b34801561147757600080fd5b5061149c600480360381019080803563ffffffff169060200190929190505050613ed6565b604051808215151515815260200191505060405180910390f35b3480156114c257600080fd5b506114e7600480360381019080803563ffffffff169060200190929190505050613f0f565b005b3480156114f557600080fd5b5061151a600480360381019080803563ffffffff169060200190929190505050614017565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561155d578082015181840152602081019050611542565b505050509050019250505060405180910390f35b34801561157d57600080fd5b506115c2600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506140bd565b005b3480156115d057600080fd5b50611615600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506145b6565b604051808215151515815260200191505060405180910390f35b34801561163b57600080fd5b50611660600480360381019080803563ffffffff1690602001909291905050506146b3565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b34801561168e57600080fd5b506116976146ef565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156116e557600080fd5b5061172a600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050614714565b604051808215151515815260200191505060405180910390f35b34801561175057600080fd5b50611785600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050614811565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156117c85780820151818401526020810190506117ad565b505050509050019250505060405180910390f35b3480156117e857600080fd5b5061180d600480360381019080803563ffffffff1690602001909291905050506148d4565b6040518082600281111561181d57fe5b60ff16815260200191505060405180910390f35b34801561183d57600080fd5b50611862600480360381019080803563ffffffff169060200190929190505050614913565b6040518082815260200191505060405180910390f35b34801561188457600080fd5b506118b9600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061493f565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156118fc5780820151818401526020810190506118e1565b505050509050019250505060405180910390f35b34801561191c57600080fd5b50611941600480360381019080803563ffffffff169060200190929190505050614a02565b604051808215151515815260200191505060405180910390f35b34801561196757600080fd5b50611996600480360381019080803563ffffffff16906020019092919080359060200190929190505050614a3b565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b3480156119c457600080fd5b50611a09600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050614a83565b005b348015611a1757600080fd5b50611a3c600480360381019080803563ffffffff169060200190929190505050614f7c565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015611a7f578082015181840152602081019050611a64565b505050509050019250505060405180910390f35b348015611a9f57600080fd5b50611ad4600480360381019080803563ffffffff169060200190929190803563ffffffff16906020019092919050505061501f565b005b348015611ae257600080fd5b50611b37600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506150d9565b604051808215151515815260200191505060405180910390f35b348015611b5d57600080fd5b50611ba2600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061516d565b604051808215151515815260200191505060405180910390f35b348015611bc857600080fd5b50611caf600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061526a565b005b348015611cbd57600080fd5b50611ce2600480360381019080803563ffffffff1690602001909291905050506154b2565b005b348015611cf057600080fd5b50611d51600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035151590602001909291905050506155c0565b005b348015611d5f57600080fd5b50611d84600480360381019080803563ffffffff1690602001909291905050506156b4565b005b348015611d9257600080fd5b50611db7600480360381019080803563ffffffff1690602001909291905050506157e9565b60405180856000191660001916815260200184600019166000191681526020018363ffffffff1663ffffffff1681526020018263ffffffff1663ffffffff16815260200194505050505060405180910390f35b348015611e1657600080fd5b50611e3b600480360381019080803563ffffffff169060200190929190505050615856565b005b348015611e4957600080fd5b50611e8e600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050615955565b604051808215151515815260200191505060405180910390f35b348015611eb457600080fd5b50611ee9600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506159d0565b6040518082815260200191505060405180910390f35b348015611f0b57600080fd5b50611f50600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803563ffffffff169060200190929190505050615a1c565b6040518082815260200191505060405180910390f35b348015611f7257600080fd5b50611fa7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050615a41565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015611fea578082015181840152602081019050611fcf565b505050509050019250505060405180910390f35b34801561200a57600080fd5b50612049600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050615b04565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b34801561207757600080fd5b506120bc600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050615b4c565b005b3480156120ca57600080fd5b506120ff600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061609b565b6040518082815260200191505060405180910390f35b34801561212157600080fd5b50612166600480360381019080803563ffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506160e7565b604051808215151515815260200191505060405180910390f35b34801561218c57600080fd5b506121b1600480360381019080803563ffffffff169060200190929190505050616162565b604051808261ffff1661ffff16815260200191505060405180910390f35b3480156121db57600080fd5b506121fa600480360381019080803590602001909291905050506161ad565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561223a57808201518184015260208101905061221f565b50505050905090810190601f1680156122675780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561228157600080fd5b506122b6600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061625f565b005b3480156122c457600080fd5b506122e9600480360381019080803563ffffffff1690602001909291905050506162c6565b005b3480156122f757600080fd5b5061231c600480360381019080803563ffffffff1690602001909291905050506163e3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561236a57600080fd5b5061238f600480360381019080803563ffffffff16906020019092919050505061642f565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b3480156123bd57600080fd5b5061240e600480360381019080803563ffffffff16906020019092919080356000191690602001909291908035600019169060200190929190803563ffffffff16906020019092919050505061646b565b005b34801561241c57600080fd5b50612451600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050616658565b6040518082815260200191505060405180910390f35b34801561247357600080fd5b506124a8600480360381019080803563ffffffff169060200190929190803563ffffffff1690602001909291905050506166a4565b604051808215151515815260200191505060405180910390f35b600a602052816000526040600020602052806000526040600020600091509150505481565b6060600b60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561259e57602002820191906000526020600020906000905b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116125615790505b50505050509050919050565b6011602052816000526040600020818154811015156125c557fe5b9060005260206000209060089182820401919006600402915091509054906101000a900463ffffffff1681565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905085565b6013602052816000526040600020818154811015156126e357fe5b9060005260206000209060089182820401919006600402915091509054906101000a900463ffffffff1681565b600080600160008563ffffffff1663ffffffff16815260200190815260200160002090508060030160009054906101000a900460ff16801561277157508263ffffffff168160030160039054906101000a900463ffffffff1663ffffffff16145b91505092915050565b6000600260008363ffffffff1663ffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6000600160008363ffffffff1663ffffffff16815260200190815260200160002060030160139054906101000a900463ffffffff169050919050565b60076020528160005260406000208181548110151561281d57fe5b9060005260206000209060089182820401919006600402915091509054906101000a900463ffffffff1681565b6000600b60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490509050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156128f357600080fd5b600160008363ffffffff1663ffffffff16815260200190815260200160002090508060030160029054906101000a900460ff16151561293157600080fd5b6129518260018360030160079054906101000a900463ffffffff1661670e565b61295d82600080616a76565b8060030160039054906101000a900463ffffffff1663ffffffff168263ffffffff167f7e447c9b1bda4b174b0796e100bf7f34ebf36dbb7fe665490b1bfce6246a9da560405160405180910390a35050565b600080600b60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050808054905083101515612a0557600080fd5b8083815481101515612a1357fe5b90600052602060002090600891828204019190066004029054906101000a900463ffffffff1691505092915050565b6000600260008363ffffffff1663ffffffff16815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60008173ffffffffffffffffffffffffffffffffffffffff16600260008563ffffffff1663ffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614905092915050565b600b60205281600052604060002081815481101515612b2457fe5b9060005260206000209060089182820401919006600402915091509054906101000a900463ffffffff1681565b60036020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b600080600160008463ffffffff1663ffffffff16815260200190815260200160002060020180549050905064010000000081101515612bbb57fe5b80915050919050565b60008060008060008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515612c2a57600080fd5b600260008a63ffffffff1663ffffffff16815260200190815260200160002096508660020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1695508773ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161415612cab576130b2565b8573ffffffffffffffffffffffffffffffffffffffff166000141515612f0d57601060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a63ffffffff1663ffffffff168152602001908152602001600020549450600085111515612d3657fe5b848060019003955050600f60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209350600184805490500392508383815481101515612d9857fe5b90600052602060002090600891828204019190066004029054906101000a900463ffffffff169150818486815481101515612dcf57fe5b90600052602060002090600891828204019190066004026101000a81548163ffffffff021916908363ffffffff16021790555060018501601060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008463ffffffff1663ffffffff168152602001908152602001600020819055508383815481101515612e7457fe5b90600052602060002090600891828204019190066004026101000a81549063ffffffff0219169055828481612ea99190616ed8565b506000601060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008b63ffffffff1663ffffffff168152602001908152602001600020819055505b8773ffffffffffffffffffffffffffffffffffffffff16600014151561302457600f60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508089908060018154018082558091505090600182039060005260206000209060089182820401919006600402909192909190916101000a81548163ffffffff021916908363ffffffff160217905550508080549050601060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008b63ffffffff1663ffffffff168152602001908152602001600020819055505b878760020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508773ffffffffffffffffffffffffffffffffffffffff168963ffffffff167f902736af7b3cefe10d9e840aed0d687e35c84095122b25051a20ead8866f006d60405160405180910390a35b505050505050505050565b60008060008060008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561312357600080fd5b600260008a63ffffffff1663ffffffff16815260200190815260200160002096508660040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1695508773ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614156131a4576135ab565b8573ffffffffffffffffffffffffffffffffffffffff16600014151561340657601460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a63ffffffff1663ffffffff16815260200190815260200160002054945060008511151561322f57fe5b848060019003955050601360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020935060018480549050039250838381548110151561329157fe5b90600052602060002090600891828204019190066004029054906101000a900463ffffffff1691508184868154811015156132c857fe5b90600052602060002090600891828204019190066004026101000a81548163ffffffff021916908363ffffffff16021790555060018501601460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008463ffffffff1663ffffffff16815260200190815260200160002081905550838381548110151561336d57fe5b90600052602060002090600891828204019190066004026101000a81549063ffffffff02191690558284816133a29190616ed8565b506000601460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008b63ffffffff1663ffffffff168152602001908152602001600020819055505b8773ffffffffffffffffffffffffffffffffffffffff16600014151561351d57601360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508089908060018154018082558091505090600182039060005260206000209060089182820401919006600402909192909190916101000a81548163ffffffff021916908363ffffffff160217905550508080549050601460008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008b63ffffffff1663ffffffff168152602001908152602001600020819055505b878760040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508773ffffffffffffffffffffffffffffffffffffffff168963ffffffff167fcfe369b7197e7f0cf06793ae2472a9b13583fecbed2f78dfa14d1f10796b847c60405160405180910390a35b505050505050505050565b6008602052816000526040600020602052806000526040600020600091509150505481565b60008173ffffffffffffffffffffffffffffffffffffffff16600260008563ffffffff1663ffffffff16815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614905092915050565b6000600260008363ffffffff1663ffffffff16815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60008173ffffffffffffffffffffffffffffffffffffffff16600260008563ffffffff1663ffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614905092915050565b6012602052816000526040600020602052806000526040600020600091509150505481565b600080600160008463ffffffff1663ffffffff168152602001908152602001600020905060006001028160000154600019161415801561378e5750600060010281600101546000191614155b80156137b55750600081600301600b9054906101000a900463ffffffff1663ffffffff1614155b915050919050565b6060600960008363ffffffff1663ffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561385457602002820191906000526020600020906000905b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116138175790505b50505050509050919050565b6000600160008363ffffffff1663ffffffff16815260200190815260200160002060030160039054906101000a900463ffffffff169050919050565b6000600960008363ffffffff1663ffffffff168152602001908152602001600020805490509050919050565b600c602052816000526040600020602052806000526040600020600091509150505481565b6060600d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156139a457602002820191906000526020600020906000905b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116139675790505b50505050509050919050565b600e602052816000526040600020602052806000526040600020600091509150505481565b600d602052816000526040600020818154811015156139f057fe5b9060005260206000209060089182820401919006600402915091509054906101000a900463ffffffff1681565b6000600160008363ffffffff1663ffffffff16815260200190815260200160002060030160019054906101000a900460ff169050919050565b6000600260008363ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60016020528060005260406000206000915090508060000154908060010154908060030160009054906101000a900460ff16908060030160019054906101000a900460ff16908060030160029054906101000a900460ff16908060030160039054906101000a900463ffffffff16908060030160079054906101000a900463ffffffff169080600301600b9054906101000a900463ffffffff169080600301600f9054906101000a900463ffffffff16908060030160139054906101000a900463ffffffff1690508a565b6000600f60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490509050919050565b600080600160008463ffffffff1663ffffffff168152602001908152602001600020600301600f9054906101000a900463ffffffff1663ffffffff16119050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515613c5857600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600080600260008563ffffffff1663ffffffff16815260200190815260200160002090508273ffffffffffffffffffffffffffffffffffffffff16600014158015613ea857508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161480613df257508060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b80613ea75750600360008260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff165b5b91505092915050565b6010602052816000526040600020602052806000526040600020600091509150505481565b6000600160008363ffffffff1663ffffffff16815260200190815260200160002060030160009054906101000a900460ff169050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515613f6c57600080fd5b600160008363ffffffff1663ffffffff16815260200190815260200160002090508060030160019054906101000a900460ff16151515613fab57600080fd5b60018160030160016101000a81548160ff021916908315150217905550613fe0826001613fd785616162565b61ffff1661670e565b8163ffffffff167fe74c03809d0769e1b1f706cc8414258cd1f3b6fe020cd15d0165c210ba503a0f60405160405180910390a25050565b6060600160008363ffffffff1663ffffffff1681526020019081526020016000206002018054806020026020016040519081016040528092919081815260200182805480156140b157602002820191906000526020600020906000905b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116140745790505b50505050509050919050565b60008060008060008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561412357600080fd5b600260008a63ffffffff1663ffffffff16815260200190815260200160002096508660010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1695508773ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614156141a4576145ab565b8573ffffffffffffffffffffffffffffffffffffffff16600014151561440657600e60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a63ffffffff1663ffffffff16815260200190815260200160002054945060008511151561422f57fe5b848060019003955050600d60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020935060018480549050039250838381548110151561429157fe5b90600052602060002090600891828204019190066004029054906101000a900463ffffffff1691508184868154811015156142c857fe5b90600052602060002090600891828204019190066004026101000a81548163ffffffff021916908363ffffffff16021790555060018501600e60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008463ffffffff1663ffffffff16815260200190815260200160002081905550838381548110151561436d57fe5b90600052602060002090600891828204019190066004026101000a81549063ffffffff02191690558284816143a29190616ed8565b506000600e60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008b63ffffffff1663ffffffff168152602001908152602001600020819055505b8773ffffffffffffffffffffffffffffffffffffffff16600014151561451d57600d60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508089908060018154018082558091505090600182039060005260206000209060089182820401919006600402909192909190916101000a81548163ffffffff021916908363ffffffff160217905550508080549050600e60008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008b63ffffffff1663ffffffff168152602001908152602001600020819055505b878760010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508773ffffffffffffffffffffffffffffffffffffffff168963ffffffff167fab9c9327cffd2acc168fafedbe06139f5f55cb84c761df05e0511c251e2ee9bf60405160405180910390a35b505050505050505050565b600080600260008563ffffffff1663ffffffff16815260200190815260200160002090508273ffffffffffffffffffffffffffffffffffffffff166000141580156146aa57508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614806146a957508060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b5b91505092915050565b6000600160008363ffffffff1663ffffffff168152602001908152602001600020600301600f9054906101000a900463ffffffff169050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600260008563ffffffff1663ffffffff16815260200190815260200160002090508273ffffffffffffffffffffffffffffffffffffffff1660001415801561480857508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16148061480757508060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b5b91505092915050565b6060601360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156148c857602002820191906000526020600020906000905b82829054906101000a900463ffffffff1663ffffffff168152602001906004019060208260030104928301926001038202915080841161488b5790505b50505050509050919050565b60006101008263ffffffff1610156148ef576000905061490e565b620100008263ffffffff161015614909576001905061490e565b600290505b919050565b6000600760008363ffffffff1663ffffffff168152602001908152602001600020805490509050919050565b6060600f60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156149f657602002820191906000526020600020906000905b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116149b95790505b50505050509050919050565b6000600160008363ffffffff1663ffffffff16815260200190815260200160002060030160029054906101000a900460ff169050919050565b600960205281600052604060002081815481101515614a5657fe5b9060005260206000209060089182820401919006600402915091509054906101000a900463ffffffff1681565b60008060008060008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515614ae957600080fd5b600260008a63ffffffff1663ffffffff16815260200190815260200160002096508660030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1695508773ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161415614b6a57614f71565b8573ffffffffffffffffffffffffffffffffffffffff166000141515614dcc57601260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a63ffffffff1663ffffffff168152602001908152602001600020549450600085111515614bf557fe5b848060019003955050601160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209350600184805490500392508383815481101515614c5757fe5b90600052602060002090600891828204019190066004029054906101000a900463ffffffff169150818486815481101515614c8e57fe5b90600052602060002090600891828204019190066004026101000a81548163ffffffff021916908363ffffffff16021790555060018501601260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008463ffffffff1663ffffffff168152602001908152602001600020819055508383815481101515614d3357fe5b90600052602060002090600891828204019190066004026101000a81549063ffffffff0219169055828481614d689190616ed8565b506000601260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008b63ffffffff1663ffffffff168152602001908152602001600020819055505b8773ffffffffffffffffffffffffffffffffffffffff166000141515614ee357601160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508089908060018154018082558091505090600182039060005260206000209060089182820401919006600402909192909190916101000a81548163ffffffff021916908363ffffffff160217905550508080549050601260008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008b63ffffffff1663ffffffff168152602001908152602001600020819055505b878760030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508773ffffffffffffffffffffffffffffffffffffffff168963ffffffff167fcbd6269ec71457f2c7b1a22774f246f6c5a2eae3795ed7300db517680c61c80560405160405180910390a35b505050505050505050565b6060600760008363ffffffff1663ffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561501357602002820191906000526020600020906000905b82829054906101000a900463ffffffff1663ffffffff1681526020019060040190602082600301049283019260010382029150808411614fd65790505b50505050509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561507a57600080fd5b61508482826166a4565b1561508e576150d5565b61509a82600183616a76565b8063ffffffff168263ffffffff167fb4d4850b8f218218141c5665cba379e53e9bb015b51e8d934be70210aead874a60405160405180910390a35b5050565b6000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600080600260008563ffffffff1663ffffffff16815260200190815260200160002090508273ffffffffffffffffffffffffffffffffffffffff1660001415801561526157508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16148061526057508060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b5b91505092915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156152c557600080fd5b82600460006003811015156152d657fe5b0190805190602001906152ea929190616f12565b5081600460016003811015156152fc57fe5b019080519060200190615310929190616f12565b50806004600260038110151561532257fe5b019080519060200190615336929190616f12565b507ffafd04ade1daae2e1fdb0fc1cc6a899fd424063ed5c92120e67e073053b9489883838360405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b838110156153a3578082015181840152602081019050615388565b50505050905090810190601f1680156153d05780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b838110156154095780820151818401526020810190506153ee565b50505050905090810190601f1680156154365780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b8381101561546f578082015181840152602081019050615454565b50505050905090810190601f16801561549c5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a1505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561550f57600080fd5b600160008363ffffffff1663ffffffff16815260200190815260200160002090508060030160009054906101000a900460ff16151561554d576155bc565b61556d8260008360030160039054906101000a900463ffffffff1661670e565b8060030160039054906101000a900463ffffffff1663ffffffff168263ffffffff167fd7704f9a25193dbd0b0cb4a809feffffa7f19d1aae8817a71346c194448210d560405160405180910390a35b5050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561561b57600080fd5b80600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561571157600080fd5b61571a82616162565b61ffff1690508163ffffffff168163ffffffff161415615739576157e5565b600160008263ffffffff1663ffffffff16815260200190815260200160002060020182908060018154018082558091505090600182039060005260206000209060089182820401919006600402909192909190916101000a81548163ffffffff021916908363ffffffff160217905550508163ffffffff168163ffffffff167fb2d3a6e7a339f5c8ff96265e2f03a010a8541070f3744a24709096441508154660405160405180910390a35b5050565b6000806000806000600160008763ffffffff1663ffffffff16815260200190815260200160002090508060000154816001015482600301600b9054906101000a900463ffffffff1683600301600f9054906101000a900463ffffffff169450945094509450509193509193565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156158b457600080fd5b600160008463ffffffff1663ffffffff16815260200190815260200160002091508160030160029054906101000a900460ff1615156158f257615950565b8160030160079054906101000a900463ffffffff16905061591583600080616a76565b8063ffffffff168363ffffffff167fd653bb0e0bb7ce8393e624d98fbf17cda5902c8328ed0cd09988f36890d9932a60405160405180910390a35b505050565b60008173ffffffffffffffffffffffffffffffffffffffff16600260008563ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614905092915050565b6000600d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490509050919050565b6014602052816000526040600020602052806000526040600020600091509150505481565b6060601160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015615af857602002820191906000526020600020906000905b82829054906101000a900463ffffffff1663ffffffff1681526020019060040190602082600301049283019260010382029150808411615abb5790505b50505050509050919050565b600f60205281600052604060002081815481101515615b1f57fe5b9060005260206000209060089182820401919006600402915091509054906101000a900463ffffffff1681565b60008060008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515615baf57600080fd5b8573ffffffffffffffffffffffffffffffffffffffff16600014151515615bd557600080fd5b600260008863ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1694508573ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161415615c5357616092565b8473ffffffffffffffffffffffffffffffffffffffff166000141515615eb557600c60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008863ffffffff1663ffffffff168152602001908152602001600020549350600084111515615cde57fe5b838060019003945050600b60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250600183805490500391508282815481101515615d4057fe5b90600052602060002090600891828204019190066004029054906101000a900463ffffffff169050808385815481101515615d7757fe5b90600052602060002090600891828204019190066004026101000a81548163ffffffff021916908363ffffffff16021790555060018401600c60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008363ffffffff1663ffffffff168152602001908152602001600020819055508282815481101515615e1c57fe5b90600052602060002090600891828204019190066004026101000a81549063ffffffff0219169055818381615e519190616ed8565b506000600c60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008963ffffffff1663ffffffff168152602001908152602001600020819055505b85600260008963ffffffff1663ffffffff16815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600b60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002087908060018154018082558091505090600182039060005260206000209060089182820401919006600402909192909190916101000a81548163ffffffff021916908363ffffffff16021790555050600b60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080549050600c60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008963ffffffff1663ffffffff168152602001908152602001600020819055508573ffffffffffffffffffffffffffffffffffffffff168763ffffffff167f16d0f539d49c6cad822b767a9445bfb1cf7ea6f2a6c2b120a7ea4cc7660d8fda60405160405180910390a35b50505050505050565b6000601360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490509050919050565b60008173ffffffffffffffffffffffffffffffffffffffff16600260008563ffffffff1663ffffffff16815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614905092915050565b6000620100008263ffffffff161015616190576101008263ffffffff1681151561618857fe5b0690506161a8565b620100008263ffffffff168115156161a457fe5b0690505b919050565b6004816003811015156161bc57fe5b016000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156162575780601f1061622c57610100808354040283529160200191616257565b820191906000526020600020905b81548152906001019060200180831161623a57829003601f168201915b505050505081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156162ba57600080fd5b6162c381616dde565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561632357600080fd5b600160008363ffffffff1663ffffffff168152602001908152602001600020905080600301601381819054906101000a900463ffffffff168092919060010191906101000a81548163ffffffff021916908363ffffffff160217905550508163ffffffff167f29294799f1c21a37ef838e15f79dd91bcee2df99d63cd1c18ac968b129514e6e8260030160139054906101000a900463ffffffff16604051808263ffffffff1663ffffffff16815260200191505060405180910390a25050565b6000600260008363ffffffff1663ffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6000600160008363ffffffff1663ffffffff16815260200190815260200160002060030160079054906101000a900463ffffffff169050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156164c857600080fd5b600160008663ffffffff1663ffffffff1681526020019081526020016000209050836000191681600001546000191614801561650f57508260001916816001015460001916145b801561653a57508163ffffffff1681600301600b9054906101000a900463ffffffff1663ffffffff16145b1561654457616651565b83816000018160001916905550828160010181600019169055508181600301600b6101000a81548163ffffffff021916908363ffffffff16021790555080600301600f81819054906101000a900463ffffffff168092919060010191906101000a81548163ffffffff021916908363ffffffff160217905550508463ffffffff167faa10e7a0117d4323f1d99d630ec169bebb3a988e895770e351987e01ff5423d585858585600301600f9054906101000a900463ffffffff1660405180856000191660001916815260200184600019166000191681526020018363ffffffff1663ffffffff1681526020018263ffffffff1663ffffffff16815260200194505050505060405180910390a25b5050505050565b6000601160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490509050919050565b600080600160008563ffffffff1663ffffffff16815260200190815260200160002090508060030160029054906101000a900460ff16801561670557508263ffffffff168160030160079054906101000a900463ffffffff1663ffffffff16145b91505092915050565b600080600080600080600080600160008c63ffffffff1663ffffffff16815260200190815260200160002097508760030160009054906101000a900460ff1696508760030160039054906101000a900463ffffffff16955086158015616772575089155b8061679a57508680156167825750895b801561679957508863ffffffff168663ffffffff16145b5b156167a457616a69565b861561696c57600860008763ffffffff1663ffffffff16815260200190815260200160002060008c63ffffffff1663ffffffff1681526020019081526020016000205494506000851115156167f557fe5b848060019003955050600760008763ffffffff1663ffffffff168152602001908152602001600020935060018480549050039250838381548110151561683757fe5b90600052602060002090600891828204019190066004029054906101000a900463ffffffff16915081848681548110151561686e57fe5b90600052602060002090600891828204019190066004026101000a81548163ffffffff021916908363ffffffff16021790555060018501600860008863ffffffff1663ffffffff16815260200190815260200160002060008463ffffffff1663ffffffff1681526020019081526020016000208190555083838154811015156168f357fe5b90600052602060002090600891828204019190066004026101000a81549063ffffffff02191690558284816169289190616ed8565b506000600860008863ffffffff1663ffffffff16815260200190815260200160002060008d63ffffffff1663ffffffff168152602001908152602001600020819055505b8915616a2957600760008a63ffffffff1663ffffffff1681526020019081526020016000209050808b908060018154018082558091505090600182039060005260206000209060089182820401919006600402909192909190916101000a81548163ffffffff021916908363ffffffff160217905550508080549050600860008b63ffffffff1663ffffffff16815260200190815260200160002060008d63ffffffff1663ffffffff168152602001908152602001600020819055505b888860030160036101000a81548163ffffffff021916908363ffffffff160217905550898860030160006101000a81548160ff0219169083151502179055505b5050505050505050505050565b600080600080600080600080600160008c63ffffffff1663ffffffff16815260200190815260200160002097508760030160029054906101000a900460ff1696508760030160079054906101000a900463ffffffff16955086158015616ada575089155b80616b025750868015616aea5750895b8015616b0157508863ffffffff168663ffffffff16145b5b15616b0c57616dd1565b8615616cd457600a60008763ffffffff1663ffffffff16815260200190815260200160002060008c63ffffffff1663ffffffff168152602001908152602001600020549450600085111515616b5d57fe5b848060019003955050600960008763ffffffff1663ffffffff1681526020019081526020016000209350600184805490500392508383815481101515616b9f57fe5b90600052602060002090600891828204019190066004029054906101000a900463ffffffff169150818486815481101515616bd657fe5b90600052602060002090600891828204019190066004026101000a81548163ffffffff021916908363ffffffff16021790555060018501600a60008863ffffffff1663ffffffff16815260200190815260200160002060008463ffffffff1663ffffffff168152602001908152602001600020819055508383815481101515616c5b57fe5b90600052602060002090600891828204019190066004026101000a81549063ffffffff0219169055828481616c909190616ed8565b506000600a60008863ffffffff1663ffffffff16815260200190815260200160002060008d63ffffffff1663ffffffff168152602001908152602001600020819055505b8915616d9157600960008a63ffffffff1663ffffffff1681526020019081526020016000209050808b908060018154018082558091505090600182039060005260206000209060089182820401919006600402909192909190916101000a81548163ffffffff021916908363ffffffff160217905550508080549050600a60008b63ffffffff1663ffffffff16815260200190815260200160002060008d63ffffffff1663ffffffff168152602001908152602001600020819055505b888860030160076101000a81548163ffffffff021916908363ffffffff160217905550898860030160026101000a81548160ff0219169083151502179055505b5050505050505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515616e1a57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b815481835581811115616f0d576007016008900481600701600890048360005260206000209182019101616f0c9190616f92565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10616f5357805160ff1916838001178555616f81565b82800160010185558215616f81579182015b82811115616f80578251825591602001919060010190616f65565b5b509050616f8e9190616f92565b5090565b616fb491905b80821115616fb0576000816000905550600101616f98565b5090565b905600a165627a7a723058205ffe833e9de79062286683ecfac75f4897fa7b12ded63f0fc5c4c566487b39200029",
}

// AzimuthABI is the input ABI used to generate the binding from.
// Deprecated: Use AzimuthMetaData.ABI instead.
var AzimuthABI = AzimuthMetaData.ABI

// AzimuthBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AzimuthMetaData.Bin instead.
var AzimuthBin = AzimuthMetaData.Bin

// DeployAzimuth deploys a new Ethereum contract, binding an instance of Azimuth to it.
func DeployAzimuth(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Azimuth, error) {
	parsed, err := AzimuthMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AzimuthBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Azimuth{AzimuthCaller: AzimuthCaller{contract: contract}, AzimuthTransactor: AzimuthTransactor{contract: contract}, AzimuthFilterer: AzimuthFilterer{contract: contract}}, nil
}

// Azimuth is an auto generated Go binding around an Ethereum contract.
type Azimuth struct {
	AzimuthCaller     // Read-only binding to the contract
	AzimuthTransactor // Write-only binding to the contract
	AzimuthFilterer   // Log filterer for contract events
}

// AzimuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type AzimuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AzimuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AzimuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AzimuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AzimuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AzimuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AzimuthSession struct {
	Contract     *Azimuth          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AzimuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AzimuthCallerSession struct {
	Contract *AzimuthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AzimuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AzimuthTransactorSession struct {
	Contract     *AzimuthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AzimuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type AzimuthRaw struct {
	Contract *Azimuth // Generic contract binding to access the raw methods on
}

// AzimuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AzimuthCallerRaw struct {
	Contract *AzimuthCaller // Generic read-only contract binding to access the raw methods on
}

// AzimuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AzimuthTransactorRaw struct {
	Contract *AzimuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAzimuth creates a new instance of Azimuth, bound to a specific deployed contract.
func NewAzimuth(address common.Address, backend bind.ContractBackend) (*Azimuth, error) {
	contract, err := bindAzimuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Azimuth{AzimuthCaller: AzimuthCaller{contract: contract}, AzimuthTransactor: AzimuthTransactor{contract: contract}, AzimuthFilterer: AzimuthFilterer{contract: contract}}, nil
}

// NewAzimuthCaller creates a new read-only instance of Azimuth, bound to a specific deployed contract.
func NewAzimuthCaller(address common.Address, caller bind.ContractCaller) (*AzimuthCaller, error) {
	contract, err := bindAzimuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AzimuthCaller{contract: contract}, nil
}

// NewAzimuthTransactor creates a new write-only instance of Azimuth, bound to a specific deployed contract.
func NewAzimuthTransactor(address common.Address, transactor bind.ContractTransactor) (*AzimuthTransactor, error) {
	contract, err := bindAzimuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AzimuthTransactor{contract: contract}, nil
}

// NewAzimuthFilterer creates a new log filterer instance of Azimuth, bound to a specific deployed contract.
func NewAzimuthFilterer(address common.Address, filterer bind.ContractFilterer) (*AzimuthFilterer, error) {
	contract, err := bindAzimuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AzimuthFilterer{contract: contract}, nil
}

// bindAzimuth binds a generic wrapper to an already deployed contract.
func bindAzimuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AzimuthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Azimuth *AzimuthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Azimuth.Contract.AzimuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Azimuth *AzimuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Azimuth.Contract.AzimuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Azimuth *AzimuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Azimuth.Contract.AzimuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Azimuth *AzimuthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Azimuth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Azimuth *AzimuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Azimuth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Azimuth *AzimuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Azimuth.Contract.contract.Transact(opts, method, params...)
}

// CanManage is a free data retrieval call binding the contract method 0x9137fe0a.
//
// Solidity: function canManage(uint32 _point, address _who) view returns(bool result)
func (_Azimuth *AzimuthCaller) CanManage(opts *bind.CallOpts, _point uint32, _who common.Address) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "canManage", _point, _who)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanManage is a free data retrieval call binding the contract method 0x9137fe0a.
//
// Solidity: function canManage(uint32 _point, address _who) view returns(bool result)
func (_Azimuth *AzimuthSession) CanManage(_point uint32, _who common.Address) (bool, error) {
	return _Azimuth.Contract.CanManage(&_Azimuth.CallOpts, _point, _who)
}

// CanManage is a free data retrieval call binding the contract method 0x9137fe0a.
//
// Solidity: function canManage(uint32 _point, address _who) view returns(bool result)
func (_Azimuth *AzimuthCallerSession) CanManage(_point uint32, _who common.Address) (bool, error) {
	return _Azimuth.Contract.CanManage(&_Azimuth.CallOpts, _point, _who)
}

// CanSpawnAs is a free data retrieval call binding the contract method 0x8a27bf59.
//
// Solidity: function canSpawnAs(uint32 _point, address _who) view returns(bool result)
func (_Azimuth *AzimuthCaller) CanSpawnAs(opts *bind.CallOpts, _point uint32, _who common.Address) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "canSpawnAs", _point, _who)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanSpawnAs is a free data retrieval call binding the contract method 0x8a27bf59.
//
// Solidity: function canSpawnAs(uint32 _point, address _who) view returns(bool result)
func (_Azimuth *AzimuthSession) CanSpawnAs(_point uint32, _who common.Address) (bool, error) {
	return _Azimuth.Contract.CanSpawnAs(&_Azimuth.CallOpts, _point, _who)
}

// CanSpawnAs is a free data retrieval call binding the contract method 0x8a27bf59.
//
// Solidity: function canSpawnAs(uint32 _point, address _who) view returns(bool result)
func (_Azimuth *AzimuthCallerSession) CanSpawnAs(_point uint32, _who common.Address) (bool, error) {
	return _Azimuth.Contract.CanSpawnAs(&_Azimuth.CallOpts, _point, _who)
}

// CanTransfer is a free data retrieval call binding the contract method 0x728aa857.
//
// Solidity: function canTransfer(uint32 _point, address _who) view returns(bool result)
func (_Azimuth *AzimuthCaller) CanTransfer(opts *bind.CallOpts, _point uint32, _who common.Address) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "canTransfer", _point, _who)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanTransfer is a free data retrieval call binding the contract method 0x728aa857.
//
// Solidity: function canTransfer(uint32 _point, address _who) view returns(bool result)
func (_Azimuth *AzimuthSession) CanTransfer(_point uint32, _who common.Address) (bool, error) {
	return _Azimuth.Contract.CanTransfer(&_Azimuth.CallOpts, _point, _who)
}

// CanTransfer is a free data retrieval call binding the contract method 0x728aa857.
//
// Solidity: function canTransfer(uint32 _point, address _who) view returns(bool result)
func (_Azimuth *AzimuthCallerSession) CanTransfer(_point uint32, _who common.Address) (bool, error) {
	return _Azimuth.Contract.CanTransfer(&_Azimuth.CallOpts, _point, _who)
}

// CanVoteAs is a free data retrieval call binding the contract method 0xb65afedd.
//
// Solidity: function canVoteAs(uint32 _point, address _who) view returns(bool result)
func (_Azimuth *AzimuthCaller) CanVoteAs(opts *bind.CallOpts, _point uint32, _who common.Address) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "canVoteAs", _point, _who)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanVoteAs is a free data retrieval call binding the contract method 0xb65afedd.
//
// Solidity: function canVoteAs(uint32 _point, address _who) view returns(bool result)
func (_Azimuth *AzimuthSession) CanVoteAs(_point uint32, _who common.Address) (bool, error) {
	return _Azimuth.Contract.CanVoteAs(&_Azimuth.CallOpts, _point, _who)
}

// CanVoteAs is a free data retrieval call binding the contract method 0xb65afedd.
//
// Solidity: function canVoteAs(uint32 _point, address _who) view returns(bool result)
func (_Azimuth *AzimuthCallerSession) CanVoteAs(_point uint32, _who common.Address) (bool, error) {
	return _Azimuth.Contract.CanVoteAs(&_Azimuth.CallOpts, _point, _who)
}

// DnsDomains is a free data retrieval call binding the contract method 0xeccc8ff1.
//
// Solidity: function dnsDomains(uint256 ) view returns(string)
func (_Azimuth *AzimuthCaller) DnsDomains(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "dnsDomains", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DnsDomains is a free data retrieval call binding the contract method 0xeccc8ff1.
//
// Solidity: function dnsDomains(uint256 ) view returns(string)
func (_Azimuth *AzimuthSession) DnsDomains(arg0 *big.Int) (string, error) {
	return _Azimuth.Contract.DnsDomains(&_Azimuth.CallOpts, arg0)
}

// DnsDomains is a free data retrieval call binding the contract method 0xeccc8ff1.
//
// Solidity: function dnsDomains(uint256 ) view returns(string)
func (_Azimuth *AzimuthCallerSession) DnsDomains(arg0 *big.Int) (string, error) {
	return _Azimuth.Contract.DnsDomains(&_Azimuth.CallOpts, arg0)
}

// EscapeRequests is a free data retrieval call binding the contract method 0x9dd90190.
//
// Solidity: function escapeRequests(uint32 , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthCaller) EscapeRequests(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "escapeRequests", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// EscapeRequests is a free data retrieval call binding the contract method 0x9dd90190.
//
// Solidity: function escapeRequests(uint32 , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthSession) EscapeRequests(arg0 uint32, arg1 *big.Int) (uint32, error) {
	return _Azimuth.Contract.EscapeRequests(&_Azimuth.CallOpts, arg0, arg1)
}

// EscapeRequests is a free data retrieval call binding the contract method 0x9dd90190.
//
// Solidity: function escapeRequests(uint32 , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthCallerSession) EscapeRequests(arg0 uint32, arg1 *big.Int) (uint32, error) {
	return _Azimuth.Contract.EscapeRequests(&_Azimuth.CallOpts, arg0, arg1)
}

// EscapeRequestsIndexes is a free data retrieval call binding the contract method 0x01d8a97f.
//
// Solidity: function escapeRequestsIndexes(uint32 , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthCaller) EscapeRequestsIndexes(opts *bind.CallOpts, arg0 uint32, arg1 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "escapeRequestsIndexes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EscapeRequestsIndexes is a free data retrieval call binding the contract method 0x01d8a97f.
//
// Solidity: function escapeRequestsIndexes(uint32 , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthSession) EscapeRequestsIndexes(arg0 uint32, arg1 uint32) (*big.Int, error) {
	return _Azimuth.Contract.EscapeRequestsIndexes(&_Azimuth.CallOpts, arg0, arg1)
}

// EscapeRequestsIndexes is a free data retrieval call binding the contract method 0x01d8a97f.
//
// Solidity: function escapeRequestsIndexes(uint32 , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthCallerSession) EscapeRequestsIndexes(arg0 uint32, arg1 uint32) (*big.Int, error) {
	return _Azimuth.Contract.EscapeRequestsIndexes(&_Azimuth.CallOpts, arg0, arg1)
}

// GetContinuityNumber is a free data retrieval call binding the contract method 0x108d995b.
//
// Solidity: function getContinuityNumber(uint32 _point) view returns(uint32 continuityNumber)
func (_Azimuth *AzimuthCaller) GetContinuityNumber(opts *bind.CallOpts, _point uint32) (uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getContinuityNumber", _point)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetContinuityNumber is a free data retrieval call binding the contract method 0x108d995b.
//
// Solidity: function getContinuityNumber(uint32 _point) view returns(uint32 continuityNumber)
func (_Azimuth *AzimuthSession) GetContinuityNumber(_point uint32) (uint32, error) {
	return _Azimuth.Contract.GetContinuityNumber(&_Azimuth.CallOpts, _point)
}

// GetContinuityNumber is a free data retrieval call binding the contract method 0x108d995b.
//
// Solidity: function getContinuityNumber(uint32 _point) view returns(uint32 continuityNumber)
func (_Azimuth *AzimuthCallerSession) GetContinuityNumber(_point uint32) (uint32, error) {
	return _Azimuth.Contract.GetContinuityNumber(&_Azimuth.CallOpts, _point)
}

// GetEscapeRequest is a free data retrieval call binding the contract method 0xf5af6621.
//
// Solidity: function getEscapeRequest(uint32 _point) view returns(uint32 escape)
func (_Azimuth *AzimuthCaller) GetEscapeRequest(opts *bind.CallOpts, _point uint32) (uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getEscapeRequest", _point)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetEscapeRequest is a free data retrieval call binding the contract method 0xf5af6621.
//
// Solidity: function getEscapeRequest(uint32 _point) view returns(uint32 escape)
func (_Azimuth *AzimuthSession) GetEscapeRequest(_point uint32) (uint32, error) {
	return _Azimuth.Contract.GetEscapeRequest(&_Azimuth.CallOpts, _point)
}

// GetEscapeRequest is a free data retrieval call binding the contract method 0xf5af6621.
//
// Solidity: function getEscapeRequest(uint32 _point) view returns(uint32 escape)
func (_Azimuth *AzimuthCallerSession) GetEscapeRequest(_point uint32) (uint32, error) {
	return _Azimuth.Contract.GetEscapeRequest(&_Azimuth.CallOpts, _point)
}

// GetEscapeRequests is a free data retrieval call binding the contract method 0x3c01a9fa.
//
// Solidity: function getEscapeRequests(uint32 _sponsor) view returns(uint32[] requests)
func (_Azimuth *AzimuthCaller) GetEscapeRequests(opts *bind.CallOpts, _sponsor uint32) ([]uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getEscapeRequests", _sponsor)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetEscapeRequests is a free data retrieval call binding the contract method 0x3c01a9fa.
//
// Solidity: function getEscapeRequests(uint32 _sponsor) view returns(uint32[] requests)
func (_Azimuth *AzimuthSession) GetEscapeRequests(_sponsor uint32) ([]uint32, error) {
	return _Azimuth.Contract.GetEscapeRequests(&_Azimuth.CallOpts, _sponsor)
}

// GetEscapeRequests is a free data retrieval call binding the contract method 0x3c01a9fa.
//
// Solidity: function getEscapeRequests(uint32 _sponsor) view returns(uint32[] requests)
func (_Azimuth *AzimuthCallerSession) GetEscapeRequests(_sponsor uint32) ([]uint32, error) {
	return _Azimuth.Contract.GetEscapeRequests(&_Azimuth.CallOpts, _sponsor)
}

// GetEscapeRequestsCount is a free data retrieval call binding the contract method 0x4448f5fb.
//
// Solidity: function getEscapeRequestsCount(uint32 _sponsor) view returns(uint256 count)
func (_Azimuth *AzimuthCaller) GetEscapeRequestsCount(opts *bind.CallOpts, _sponsor uint32) (*big.Int, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getEscapeRequestsCount", _sponsor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEscapeRequestsCount is a free data retrieval call binding the contract method 0x4448f5fb.
//
// Solidity: function getEscapeRequestsCount(uint32 _sponsor) view returns(uint256 count)
func (_Azimuth *AzimuthSession) GetEscapeRequestsCount(_sponsor uint32) (*big.Int, error) {
	return _Azimuth.Contract.GetEscapeRequestsCount(&_Azimuth.CallOpts, _sponsor)
}

// GetEscapeRequestsCount is a free data retrieval call binding the contract method 0x4448f5fb.
//
// Solidity: function getEscapeRequestsCount(uint32 _sponsor) view returns(uint256 count)
func (_Azimuth *AzimuthCallerSession) GetEscapeRequestsCount(_sponsor uint32) (*big.Int, error) {
	return _Azimuth.Contract.GetEscapeRequestsCount(&_Azimuth.CallOpts, _sponsor)
}

// GetKeyRevisionNumber is a free data retrieval call binding the contract method 0x8d2d3f41.
//
// Solidity: function getKeyRevisionNumber(uint32 _point) view returns(uint32 revision)
func (_Azimuth *AzimuthCaller) GetKeyRevisionNumber(opts *bind.CallOpts, _point uint32) (uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getKeyRevisionNumber", _point)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetKeyRevisionNumber is a free data retrieval call binding the contract method 0x8d2d3f41.
//
// Solidity: function getKeyRevisionNumber(uint32 _point) view returns(uint32 revision)
func (_Azimuth *AzimuthSession) GetKeyRevisionNumber(_point uint32) (uint32, error) {
	return _Azimuth.Contract.GetKeyRevisionNumber(&_Azimuth.CallOpts, _point)
}

// GetKeyRevisionNumber is a free data retrieval call binding the contract method 0x8d2d3f41.
//
// Solidity: function getKeyRevisionNumber(uint32 _point) view returns(uint32 revision)
func (_Azimuth *AzimuthCallerSession) GetKeyRevisionNumber(_point uint32) (uint32, error) {
	return _Azimuth.Contract.GetKeyRevisionNumber(&_Azimuth.CallOpts, _point)
}

// GetKeys is a free data retrieval call binding the contract method 0xc210c623.
//
// Solidity: function getKeys(uint32 _point) view returns(bytes32 crypt, bytes32 auth, uint32 suite, uint32 revision)
func (_Azimuth *AzimuthCaller) GetKeys(opts *bind.CallOpts, _point uint32) (struct {
	Crypt    [32]byte
	Auth     [32]byte
	Suite    uint32
	Revision uint32
}, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getKeys", _point)

	outstruct := new(struct {
		Crypt    [32]byte
		Auth     [32]byte
		Suite    uint32
		Revision uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Crypt = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Auth = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Suite = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Revision = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetKeys is a free data retrieval call binding the contract method 0xc210c623.
//
// Solidity: function getKeys(uint32 _point) view returns(bytes32 crypt, bytes32 auth, uint32 suite, uint32 revision)
func (_Azimuth *AzimuthSession) GetKeys(_point uint32) (struct {
	Crypt    [32]byte
	Auth     [32]byte
	Suite    uint32
	Revision uint32
}, error) {
	return _Azimuth.Contract.GetKeys(&_Azimuth.CallOpts, _point)
}

// GetKeys is a free data retrieval call binding the contract method 0xc210c623.
//
// Solidity: function getKeys(uint32 _point) view returns(bytes32 crypt, bytes32 auth, uint32 suite, uint32 revision)
func (_Azimuth *AzimuthCallerSession) GetKeys(_point uint32) (struct {
	Crypt    [32]byte
	Auth     [32]byte
	Suite    uint32
	Revision uint32
}, error) {
	return _Azimuth.Contract.GetKeys(&_Azimuth.CallOpts, _point)
}

// GetManagementProxy is a free data retrieval call binding the contract method 0x0f934e66.
//
// Solidity: function getManagementProxy(uint32 _point) view returns(address manager)
func (_Azimuth *AzimuthCaller) GetManagementProxy(opts *bind.CallOpts, _point uint32) (common.Address, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getManagementProxy", _point)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetManagementProxy is a free data retrieval call binding the contract method 0x0f934e66.
//
// Solidity: function getManagementProxy(uint32 _point) view returns(address manager)
func (_Azimuth *AzimuthSession) GetManagementProxy(_point uint32) (common.Address, error) {
	return _Azimuth.Contract.GetManagementProxy(&_Azimuth.CallOpts, _point)
}

// GetManagementProxy is a free data retrieval call binding the contract method 0x0f934e66.
//
// Solidity: function getManagementProxy(uint32 _point) view returns(address manager)
func (_Azimuth *AzimuthCallerSession) GetManagementProxy(_point uint32) (common.Address, error) {
	return _Azimuth.Contract.GetManagementProxy(&_Azimuth.CallOpts, _point)
}

// GetManagerFor is a free data retrieval call binding the contract method 0x483a15ed.
//
// Solidity: function getManagerFor(address _proxy) view returns(uint32[] mfor)
func (_Azimuth *AzimuthCaller) GetManagerFor(opts *bind.CallOpts, _proxy common.Address) ([]uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getManagerFor", _proxy)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetManagerFor is a free data retrieval call binding the contract method 0x483a15ed.
//
// Solidity: function getManagerFor(address _proxy) view returns(uint32[] mfor)
func (_Azimuth *AzimuthSession) GetManagerFor(_proxy common.Address) ([]uint32, error) {
	return _Azimuth.Contract.GetManagerFor(&_Azimuth.CallOpts, _proxy)
}

// GetManagerFor is a free data retrieval call binding the contract method 0x483a15ed.
//
// Solidity: function getManagerFor(address _proxy) view returns(uint32[] mfor)
func (_Azimuth *AzimuthCallerSession) GetManagerFor(_proxy common.Address) ([]uint32, error) {
	return _Azimuth.Contract.GetManagerFor(&_Azimuth.CallOpts, _proxy)
}

// GetManagerForCount is a free data retrieval call binding the contract method 0xd1541264.
//
// Solidity: function getManagerForCount(address _proxy) view returns(uint256 count)
func (_Azimuth *AzimuthCaller) GetManagerForCount(opts *bind.CallOpts, _proxy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getManagerForCount", _proxy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetManagerForCount is a free data retrieval call binding the contract method 0xd1541264.
//
// Solidity: function getManagerForCount(address _proxy) view returns(uint256 count)
func (_Azimuth *AzimuthSession) GetManagerForCount(_proxy common.Address) (*big.Int, error) {
	return _Azimuth.Contract.GetManagerForCount(&_Azimuth.CallOpts, _proxy)
}

// GetManagerForCount is a free data retrieval call binding the contract method 0xd1541264.
//
// Solidity: function getManagerForCount(address _proxy) view returns(uint256 count)
func (_Azimuth *AzimuthCallerSession) GetManagerForCount(_proxy common.Address) (*big.Int, error) {
	return _Azimuth.Contract.GetManagerForCount(&_Azimuth.CallOpts, _proxy)
}

// GetOwnedPointAtIndex is a free data retrieval call binding the contract method 0x22a7906f.
//
// Solidity: function getOwnedPointAtIndex(address _whose, uint256 _index) view returns(uint32 point)
func (_Azimuth *AzimuthCaller) GetOwnedPointAtIndex(opts *bind.CallOpts, _whose common.Address, _index *big.Int) (uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getOwnedPointAtIndex", _whose, _index)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetOwnedPointAtIndex is a free data retrieval call binding the contract method 0x22a7906f.
//
// Solidity: function getOwnedPointAtIndex(address _whose, uint256 _index) view returns(uint32 point)
func (_Azimuth *AzimuthSession) GetOwnedPointAtIndex(_whose common.Address, _index *big.Int) (uint32, error) {
	return _Azimuth.Contract.GetOwnedPointAtIndex(&_Azimuth.CallOpts, _whose, _index)
}

// GetOwnedPointAtIndex is a free data retrieval call binding the contract method 0x22a7906f.
//
// Solidity: function getOwnedPointAtIndex(address _whose, uint256 _index) view returns(uint32 point)
func (_Azimuth *AzimuthCallerSession) GetOwnedPointAtIndex(_whose common.Address, _index *big.Int) (uint32, error) {
	return _Azimuth.Contract.GetOwnedPointAtIndex(&_Azimuth.CallOpts, _whose, _index)
}

// GetOwnedPointCount is a free data retrieval call binding the contract method 0x12afbc78.
//
// Solidity: function getOwnedPointCount(address _whose) view returns(uint256 count)
func (_Azimuth *AzimuthCaller) GetOwnedPointCount(opts *bind.CallOpts, _whose common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getOwnedPointCount", _whose)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOwnedPointCount is a free data retrieval call binding the contract method 0x12afbc78.
//
// Solidity: function getOwnedPointCount(address _whose) view returns(uint256 count)
func (_Azimuth *AzimuthSession) GetOwnedPointCount(_whose common.Address) (*big.Int, error) {
	return _Azimuth.Contract.GetOwnedPointCount(&_Azimuth.CallOpts, _whose)
}

// GetOwnedPointCount is a free data retrieval call binding the contract method 0x12afbc78.
//
// Solidity: function getOwnedPointCount(address _whose) view returns(uint256 count)
func (_Azimuth *AzimuthCallerSession) GetOwnedPointCount(_whose common.Address) (*big.Int, error) {
	return _Azimuth.Contract.GetOwnedPointCount(&_Azimuth.CallOpts, _whose)
}

// GetOwnedPoints is a free data retrieval call binding the contract method 0x0333d499.
//
// Solidity: function getOwnedPoints(address _whose) view returns(uint32[] ownedPoints)
func (_Azimuth *AzimuthCaller) GetOwnedPoints(opts *bind.CallOpts, _whose common.Address) ([]uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getOwnedPoints", _whose)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetOwnedPoints is a free data retrieval call binding the contract method 0x0333d499.
//
// Solidity: function getOwnedPoints(address _whose) view returns(uint32[] ownedPoints)
func (_Azimuth *AzimuthSession) GetOwnedPoints(_whose common.Address) ([]uint32, error) {
	return _Azimuth.Contract.GetOwnedPoints(&_Azimuth.CallOpts, _whose)
}

// GetOwnedPoints is a free data retrieval call binding the contract method 0x0333d499.
//
// Solidity: function getOwnedPoints(address _whose) view returns(uint32[] ownedPoints)
func (_Azimuth *AzimuthCallerSession) GetOwnedPoints(_whose common.Address) ([]uint32, error) {
	return _Azimuth.Contract.GetOwnedPoints(&_Azimuth.CallOpts, _whose)
}

// GetOwner is a free data retrieval call binding the contract method 0x621b23e2.
//
// Solidity: function getOwner(uint32 _point) view returns(address owner)
func (_Azimuth *AzimuthCaller) GetOwner(opts *bind.CallOpts, _point uint32) (common.Address, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getOwner", _point)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x621b23e2.
//
// Solidity: function getOwner(uint32 _point) view returns(address owner)
func (_Azimuth *AzimuthSession) GetOwner(_point uint32) (common.Address, error) {
	return _Azimuth.Contract.GetOwner(&_Azimuth.CallOpts, _point)
}

// GetOwner is a free data retrieval call binding the contract method 0x621b23e2.
//
// Solidity: function getOwner(uint32 _point) view returns(address owner)
func (_Azimuth *AzimuthCallerSession) GetOwner(_point uint32) (common.Address, error) {
	return _Azimuth.Contract.GetOwner(&_Azimuth.CallOpts, _point)
}

// GetPointSize is a free data retrieval call binding the contract method 0x93976405.
//
// Solidity: function getPointSize(uint32 _point) pure returns(uint8 _size)
func (_Azimuth *AzimuthCaller) GetPointSize(opts *bind.CallOpts, _point uint32) (uint8, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getPointSize", _point)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPointSize is a free data retrieval call binding the contract method 0x93976405.
//
// Solidity: function getPointSize(uint32 _point) pure returns(uint8 _size)
func (_Azimuth *AzimuthSession) GetPointSize(_point uint32) (uint8, error) {
	return _Azimuth.Contract.GetPointSize(&_Azimuth.CallOpts, _point)
}

// GetPointSize is a free data retrieval call binding the contract method 0x93976405.
//
// Solidity: function getPointSize(uint32 _point) pure returns(uint8 _size)
func (_Azimuth *AzimuthCallerSession) GetPointSize(_point uint32) (uint8, error) {
	return _Azimuth.Contract.GetPointSize(&_Azimuth.CallOpts, _point)
}

// GetPrefix is a free data retrieval call binding the contract method 0xe4a358d7.
//
// Solidity: function getPrefix(uint32 _point) pure returns(uint16 prefix)
func (_Azimuth *AzimuthCaller) GetPrefix(opts *bind.CallOpts, _point uint32) (uint16, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getPrefix", _point)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetPrefix is a free data retrieval call binding the contract method 0xe4a358d7.
//
// Solidity: function getPrefix(uint32 _point) pure returns(uint16 prefix)
func (_Azimuth *AzimuthSession) GetPrefix(_point uint32) (uint16, error) {
	return _Azimuth.Contract.GetPrefix(&_Azimuth.CallOpts, _point)
}

// GetPrefix is a free data retrieval call binding the contract method 0xe4a358d7.
//
// Solidity: function getPrefix(uint32 _point) pure returns(uint16 prefix)
func (_Azimuth *AzimuthCallerSession) GetPrefix(_point uint32) (uint16, error) {
	return _Azimuth.Contract.GetPrefix(&_Azimuth.CallOpts, _point)
}

// GetSpawnCount is a free data retrieval call binding the contract method 0x293a9169.
//
// Solidity: function getSpawnCount(uint32 _point) view returns(uint32 spawnCount)
func (_Azimuth *AzimuthCaller) GetSpawnCount(opts *bind.CallOpts, _point uint32) (uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getSpawnCount", _point)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetSpawnCount is a free data retrieval call binding the contract method 0x293a9169.
//
// Solidity: function getSpawnCount(uint32 _point) view returns(uint32 spawnCount)
func (_Azimuth *AzimuthSession) GetSpawnCount(_point uint32) (uint32, error) {
	return _Azimuth.Contract.GetSpawnCount(&_Azimuth.CallOpts, _point)
}

// GetSpawnCount is a free data retrieval call binding the contract method 0x293a9169.
//
// Solidity: function getSpawnCount(uint32 _point) view returns(uint32 spawnCount)
func (_Azimuth *AzimuthCallerSession) GetSpawnCount(_point uint32) (uint32, error) {
	return _Azimuth.Contract.GetSpawnCount(&_Azimuth.CallOpts, _point)
}

// GetSpawnProxy is a free data retrieval call binding the contract method 0xf4e3be2d.
//
// Solidity: function getSpawnProxy(uint32 _point) view returns(address spawnProxy)
func (_Azimuth *AzimuthCaller) GetSpawnProxy(opts *bind.CallOpts, _point uint32) (common.Address, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getSpawnProxy", _point)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSpawnProxy is a free data retrieval call binding the contract method 0xf4e3be2d.
//
// Solidity: function getSpawnProxy(uint32 _point) view returns(address spawnProxy)
func (_Azimuth *AzimuthSession) GetSpawnProxy(_point uint32) (common.Address, error) {
	return _Azimuth.Contract.GetSpawnProxy(&_Azimuth.CallOpts, _point)
}

// GetSpawnProxy is a free data retrieval call binding the contract method 0xf4e3be2d.
//
// Solidity: function getSpawnProxy(uint32 _point) view returns(address spawnProxy)
func (_Azimuth *AzimuthCallerSession) GetSpawnProxy(_point uint32) (common.Address, error) {
	return _Azimuth.Contract.GetSpawnProxy(&_Azimuth.CallOpts, _point)
}

// GetSpawned is a free data retrieval call binding the contract method 0x846f652b.
//
// Solidity: function getSpawned(uint32 _point) view returns(uint32[] spawned)
func (_Azimuth *AzimuthCaller) GetSpawned(opts *bind.CallOpts, _point uint32) ([]uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getSpawned", _point)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetSpawned is a free data retrieval call binding the contract method 0x846f652b.
//
// Solidity: function getSpawned(uint32 _point) view returns(uint32[] spawned)
func (_Azimuth *AzimuthSession) GetSpawned(_point uint32) ([]uint32, error) {
	return _Azimuth.Contract.GetSpawned(&_Azimuth.CallOpts, _point)
}

// GetSpawned is a free data retrieval call binding the contract method 0x846f652b.
//
// Solidity: function getSpawned(uint32 _point) view returns(uint32[] spawned)
func (_Azimuth *AzimuthCallerSession) GetSpawned(_point uint32) ([]uint32, error) {
	return _Azimuth.Contract.GetSpawned(&_Azimuth.CallOpts, _point)
}

// GetSpawningFor is a free data retrieval call binding the contract method 0x9b0d1b2c.
//
// Solidity: function getSpawningFor(address _proxy) view returns(uint32[] sfor)
func (_Azimuth *AzimuthCaller) GetSpawningFor(opts *bind.CallOpts, _proxy common.Address) ([]uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getSpawningFor", _proxy)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetSpawningFor is a free data retrieval call binding the contract method 0x9b0d1b2c.
//
// Solidity: function getSpawningFor(address _proxy) view returns(uint32[] sfor)
func (_Azimuth *AzimuthSession) GetSpawningFor(_proxy common.Address) ([]uint32, error) {
	return _Azimuth.Contract.GetSpawningFor(&_Azimuth.CallOpts, _proxy)
}

// GetSpawningFor is a free data retrieval call binding the contract method 0x9b0d1b2c.
//
// Solidity: function getSpawningFor(address _proxy) view returns(uint32[] sfor)
func (_Azimuth *AzimuthCallerSession) GetSpawningFor(_proxy common.Address) ([]uint32, error) {
	return _Azimuth.Contract.GetSpawningFor(&_Azimuth.CallOpts, _proxy)
}

// GetSpawningForCount is a free data retrieval call binding the contract method 0x6436b1b9.
//
// Solidity: function getSpawningForCount(address _proxy) view returns(uint256 count)
func (_Azimuth *AzimuthCaller) GetSpawningForCount(opts *bind.CallOpts, _proxy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getSpawningForCount", _proxy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpawningForCount is a free data retrieval call binding the contract method 0x6436b1b9.
//
// Solidity: function getSpawningForCount(address _proxy) view returns(uint256 count)
func (_Azimuth *AzimuthSession) GetSpawningForCount(_proxy common.Address) (*big.Int, error) {
	return _Azimuth.Contract.GetSpawningForCount(&_Azimuth.CallOpts, _proxy)
}

// GetSpawningForCount is a free data retrieval call binding the contract method 0x6436b1b9.
//
// Solidity: function getSpawningForCount(address _proxy) view returns(uint256 count)
func (_Azimuth *AzimuthCallerSession) GetSpawningForCount(_proxy common.Address) (*big.Int, error) {
	return _Azimuth.Contract.GetSpawningForCount(&_Azimuth.CallOpts, _proxy)
}

// GetSponsor is a free data retrieval call binding the contract method 0x439f7d3c.
//
// Solidity: function getSponsor(uint32 _point) view returns(uint32 sponsor)
func (_Azimuth *AzimuthCaller) GetSponsor(opts *bind.CallOpts, _point uint32) (uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getSponsor", _point)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetSponsor is a free data retrieval call binding the contract method 0x439f7d3c.
//
// Solidity: function getSponsor(uint32 _point) view returns(uint32 sponsor)
func (_Azimuth *AzimuthSession) GetSponsor(_point uint32) (uint32, error) {
	return _Azimuth.Contract.GetSponsor(&_Azimuth.CallOpts, _point)
}

// GetSponsor is a free data retrieval call binding the contract method 0x439f7d3c.
//
// Solidity: function getSponsor(uint32 _point) view returns(uint32 sponsor)
func (_Azimuth *AzimuthCallerSession) GetSponsor(_point uint32) (uint32, error) {
	return _Azimuth.Contract.GetSponsor(&_Azimuth.CallOpts, _point)
}

// GetSponsoring is a free data retrieval call binding the contract method 0xa46afef1.
//
// Solidity: function getSponsoring(uint32 _sponsor) view returns(uint32[] sponsees)
func (_Azimuth *AzimuthCaller) GetSponsoring(opts *bind.CallOpts, _sponsor uint32) ([]uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getSponsoring", _sponsor)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetSponsoring is a free data retrieval call binding the contract method 0xa46afef1.
//
// Solidity: function getSponsoring(uint32 _sponsor) view returns(uint32[] sponsees)
func (_Azimuth *AzimuthSession) GetSponsoring(_sponsor uint32) ([]uint32, error) {
	return _Azimuth.Contract.GetSponsoring(&_Azimuth.CallOpts, _sponsor)
}

// GetSponsoring is a free data retrieval call binding the contract method 0xa46afef1.
//
// Solidity: function getSponsoring(uint32 _sponsor) view returns(uint32[] sponsees)
func (_Azimuth *AzimuthCallerSession) GetSponsoring(_sponsor uint32) ([]uint32, error) {
	return _Azimuth.Contract.GetSponsoring(&_Azimuth.CallOpts, _sponsor)
}

// GetSponsoringCount is a free data retrieval call binding the contract method 0x942324c3.
//
// Solidity: function getSponsoringCount(uint32 _sponsor) view returns(uint256 count)
func (_Azimuth *AzimuthCaller) GetSponsoringCount(opts *bind.CallOpts, _sponsor uint32) (*big.Int, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getSponsoringCount", _sponsor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSponsoringCount is a free data retrieval call binding the contract method 0x942324c3.
//
// Solidity: function getSponsoringCount(uint32 _sponsor) view returns(uint256 count)
func (_Azimuth *AzimuthSession) GetSponsoringCount(_sponsor uint32) (*big.Int, error) {
	return _Azimuth.Contract.GetSponsoringCount(&_Azimuth.CallOpts, _sponsor)
}

// GetSponsoringCount is a free data retrieval call binding the contract method 0x942324c3.
//
// Solidity: function getSponsoringCount(uint32 _sponsor) view returns(uint256 count)
func (_Azimuth *AzimuthCallerSession) GetSponsoringCount(_sponsor uint32) (*big.Int, error) {
	return _Azimuth.Contract.GetSponsoringCount(&_Azimuth.CallOpts, _sponsor)
}

// GetTransferProxy is a free data retrieval call binding the contract method 0x24541f78.
//
// Solidity: function getTransferProxy(uint32 _point) view returns(address transferProxy)
func (_Azimuth *AzimuthCaller) GetTransferProxy(opts *bind.CallOpts, _point uint32) (common.Address, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getTransferProxy", _point)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTransferProxy is a free data retrieval call binding the contract method 0x24541f78.
//
// Solidity: function getTransferProxy(uint32 _point) view returns(address transferProxy)
func (_Azimuth *AzimuthSession) GetTransferProxy(_point uint32) (common.Address, error) {
	return _Azimuth.Contract.GetTransferProxy(&_Azimuth.CallOpts, _point)
}

// GetTransferProxy is a free data retrieval call binding the contract method 0x24541f78.
//
// Solidity: function getTransferProxy(uint32 _point) view returns(address transferProxy)
func (_Azimuth *AzimuthCallerSession) GetTransferProxy(_point uint32) (common.Address, error) {
	return _Azimuth.Contract.GetTransferProxy(&_Azimuth.CallOpts, _point)
}

// GetTransferringFor is a free data retrieval call binding the contract method 0x91398f25.
//
// Solidity: function getTransferringFor(address _proxy) view returns(uint32[] tfor)
func (_Azimuth *AzimuthCaller) GetTransferringFor(opts *bind.CallOpts, _proxy common.Address) ([]uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getTransferringFor", _proxy)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetTransferringFor is a free data retrieval call binding the contract method 0x91398f25.
//
// Solidity: function getTransferringFor(address _proxy) view returns(uint32[] tfor)
func (_Azimuth *AzimuthSession) GetTransferringFor(_proxy common.Address) ([]uint32, error) {
	return _Azimuth.Contract.GetTransferringFor(&_Azimuth.CallOpts, _proxy)
}

// GetTransferringFor is a free data retrieval call binding the contract method 0x91398f25.
//
// Solidity: function getTransferringFor(address _proxy) view returns(uint32[] tfor)
func (_Azimuth *AzimuthCallerSession) GetTransferringFor(_proxy common.Address) ([]uint32, error) {
	return _Azimuth.Contract.GetTransferringFor(&_Azimuth.CallOpts, _proxy)
}

// GetTransferringForCount is a free data retrieval call binding the contract method 0xe01cff84.
//
// Solidity: function getTransferringForCount(address _proxy) view returns(uint256 count)
func (_Azimuth *AzimuthCaller) GetTransferringForCount(opts *bind.CallOpts, _proxy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getTransferringForCount", _proxy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransferringForCount is a free data retrieval call binding the contract method 0xe01cff84.
//
// Solidity: function getTransferringForCount(address _proxy) view returns(uint256 count)
func (_Azimuth *AzimuthSession) GetTransferringForCount(_proxy common.Address) (*big.Int, error) {
	return _Azimuth.Contract.GetTransferringForCount(&_Azimuth.CallOpts, _proxy)
}

// GetTransferringForCount is a free data retrieval call binding the contract method 0xe01cff84.
//
// Solidity: function getTransferringForCount(address _proxy) view returns(uint256 count)
func (_Azimuth *AzimuthCallerSession) GetTransferringForCount(_proxy common.Address) (*big.Int, error) {
	return _Azimuth.Contract.GetTransferringForCount(&_Azimuth.CallOpts, _proxy)
}

// GetVotingFor is a free data retrieval call binding the contract method 0xd2a9b6ed.
//
// Solidity: function getVotingFor(address _proxy) view returns(uint32[] vfor)
func (_Azimuth *AzimuthCaller) GetVotingFor(opts *bind.CallOpts, _proxy common.Address) ([]uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getVotingFor", _proxy)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetVotingFor is a free data retrieval call binding the contract method 0xd2a9b6ed.
//
// Solidity: function getVotingFor(address _proxy) view returns(uint32[] vfor)
func (_Azimuth *AzimuthSession) GetVotingFor(_proxy common.Address) ([]uint32, error) {
	return _Azimuth.Contract.GetVotingFor(&_Azimuth.CallOpts, _proxy)
}

// GetVotingFor is a free data retrieval call binding the contract method 0xd2a9b6ed.
//
// Solidity: function getVotingFor(address _proxy) view returns(uint32[] vfor)
func (_Azimuth *AzimuthCallerSession) GetVotingFor(_proxy common.Address) ([]uint32, error) {
	return _Azimuth.Contract.GetVotingFor(&_Azimuth.CallOpts, _proxy)
}

// GetVotingForCount is a free data retrieval call binding the contract method 0xfdb357be.
//
// Solidity: function getVotingForCount(address _proxy) view returns(uint256 count)
func (_Azimuth *AzimuthCaller) GetVotingForCount(opts *bind.CallOpts, _proxy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getVotingForCount", _proxy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotingForCount is a free data retrieval call binding the contract method 0xfdb357be.
//
// Solidity: function getVotingForCount(address _proxy) view returns(uint256 count)
func (_Azimuth *AzimuthSession) GetVotingForCount(_proxy common.Address) (*big.Int, error) {
	return _Azimuth.Contract.GetVotingForCount(&_Azimuth.CallOpts, _proxy)
}

// GetVotingForCount is a free data retrieval call binding the contract method 0xfdb357be.
//
// Solidity: function getVotingForCount(address _proxy) view returns(uint256 count)
func (_Azimuth *AzimuthCallerSession) GetVotingForCount(_proxy common.Address) (*big.Int, error) {
	return _Azimuth.Contract.GetVotingForCount(&_Azimuth.CallOpts, _proxy)
}

// GetVotingProxy is a free data retrieval call binding the contract method 0x33f33724.
//
// Solidity: function getVotingProxy(uint32 _point) view returns(address voter)
func (_Azimuth *AzimuthCaller) GetVotingProxy(opts *bind.CallOpts, _point uint32) (common.Address, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "getVotingProxy", _point)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVotingProxy is a free data retrieval call binding the contract method 0x33f33724.
//
// Solidity: function getVotingProxy(uint32 _point) view returns(address voter)
func (_Azimuth *AzimuthSession) GetVotingProxy(_point uint32) (common.Address, error) {
	return _Azimuth.Contract.GetVotingProxy(&_Azimuth.CallOpts, _point)
}

// GetVotingProxy is a free data retrieval call binding the contract method 0x33f33724.
//
// Solidity: function getVotingProxy(uint32 _point) view returns(address voter)
func (_Azimuth *AzimuthCallerSession) GetVotingProxy(_point uint32) (common.Address, error) {
	return _Azimuth.Contract.GetVotingProxy(&_Azimuth.CallOpts, _point)
}

// HasBeenLinked is a free data retrieval call binding the contract method 0x6d09887b.
//
// Solidity: function hasBeenLinked(uint32 _point) view returns(bool result)
func (_Azimuth *AzimuthCaller) HasBeenLinked(opts *bind.CallOpts, _point uint32) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "hasBeenLinked", _point)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasBeenLinked is a free data retrieval call binding the contract method 0x6d09887b.
//
// Solidity: function hasBeenLinked(uint32 _point) view returns(bool result)
func (_Azimuth *AzimuthSession) HasBeenLinked(_point uint32) (bool, error) {
	return _Azimuth.Contract.HasBeenLinked(&_Azimuth.CallOpts, _point)
}

// HasBeenLinked is a free data retrieval call binding the contract method 0x6d09887b.
//
// Solidity: function hasBeenLinked(uint32 _point) view returns(bool result)
func (_Azimuth *AzimuthCallerSession) HasBeenLinked(_point uint32) (bool, error) {
	return _Azimuth.Contract.HasBeenLinked(&_Azimuth.CallOpts, _point)
}

// HasSponsor is a free data retrieval call binding the contract method 0x77eb4c50.
//
// Solidity: function hasSponsor(uint32 _point) view returns(bool has)
func (_Azimuth *AzimuthCaller) HasSponsor(opts *bind.CallOpts, _point uint32) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "hasSponsor", _point)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasSponsor is a free data retrieval call binding the contract method 0x77eb4c50.
//
// Solidity: function hasSponsor(uint32 _point) view returns(bool has)
func (_Azimuth *AzimuthSession) HasSponsor(_point uint32) (bool, error) {
	return _Azimuth.Contract.HasSponsor(&_Azimuth.CallOpts, _point)
}

// HasSponsor is a free data retrieval call binding the contract method 0x77eb4c50.
//
// Solidity: function hasSponsor(uint32 _point) view returns(bool has)
func (_Azimuth *AzimuthCallerSession) HasSponsor(_point uint32) (bool, error) {
	return _Azimuth.Contract.HasSponsor(&_Azimuth.CallOpts, _point)
}

// IsActive is a free data retrieval call binding the contract method 0x5e19b305.
//
// Solidity: function isActive(uint32 _point) view returns(bool equals)
func (_Azimuth *AzimuthCaller) IsActive(opts *bind.CallOpts, _point uint32) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "isActive", _point)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x5e19b305.
//
// Solidity: function isActive(uint32 _point) view returns(bool equals)
func (_Azimuth *AzimuthSession) IsActive(_point uint32) (bool, error) {
	return _Azimuth.Contract.IsActive(&_Azimuth.CallOpts, _point)
}

// IsActive is a free data retrieval call binding the contract method 0x5e19b305.
//
// Solidity: function isActive(uint32 _point) view returns(bool equals)
func (_Azimuth *AzimuthCallerSession) IsActive(_point uint32) (bool, error) {
	return _Azimuth.Contract.IsActive(&_Azimuth.CallOpts, _point)
}

// IsEscaping is a free data retrieval call binding the contract method 0x9b350e12.
//
// Solidity: function isEscaping(uint32 _point) view returns(bool escaping)
func (_Azimuth *AzimuthCaller) IsEscaping(opts *bind.CallOpts, _point uint32) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "isEscaping", _point)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEscaping is a free data retrieval call binding the contract method 0x9b350e12.
//
// Solidity: function isEscaping(uint32 _point) view returns(bool escaping)
func (_Azimuth *AzimuthSession) IsEscaping(_point uint32) (bool, error) {
	return _Azimuth.Contract.IsEscaping(&_Azimuth.CallOpts, _point)
}

// IsEscaping is a free data retrieval call binding the contract method 0x9b350e12.
//
// Solidity: function isEscaping(uint32 _point) view returns(bool escaping)
func (_Azimuth *AzimuthCallerSession) IsEscaping(_point uint32) (bool, error) {
	return _Azimuth.Contract.IsEscaping(&_Azimuth.CallOpts, _point)
}

// IsLive is a free data retrieval call binding the contract method 0x3b17b4ea.
//
// Solidity: function isLive(uint32 _point) view returns(bool result)
func (_Azimuth *AzimuthCaller) IsLive(opts *bind.CallOpts, _point uint32) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "isLive", _point)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLive is a free data retrieval call binding the contract method 0x3b17b4ea.
//
// Solidity: function isLive(uint32 _point) view returns(bool result)
func (_Azimuth *AzimuthSession) IsLive(_point uint32) (bool, error) {
	return _Azimuth.Contract.IsLive(&_Azimuth.CallOpts, _point)
}

// IsLive is a free data retrieval call binding the contract method 0x3b17b4ea.
//
// Solidity: function isLive(uint32 _point) view returns(bool result)
func (_Azimuth *AzimuthCallerSession) IsLive(_point uint32) (bool, error) {
	return _Azimuth.Contract.IsLive(&_Azimuth.CallOpts, _point)
}

// IsManagementProxy is a free data retrieval call binding the contract method 0x375e24c1.
//
// Solidity: function isManagementProxy(uint32 _point, address _proxy) view returns(bool result)
func (_Azimuth *AzimuthCaller) IsManagementProxy(opts *bind.CallOpts, _point uint32, _proxy common.Address) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "isManagementProxy", _point, _proxy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManagementProxy is a free data retrieval call binding the contract method 0x375e24c1.
//
// Solidity: function isManagementProxy(uint32 _point, address _proxy) view returns(bool result)
func (_Azimuth *AzimuthSession) IsManagementProxy(_point uint32, _proxy common.Address) (bool, error) {
	return _Azimuth.Contract.IsManagementProxy(&_Azimuth.CallOpts, _point, _proxy)
}

// IsManagementProxy is a free data retrieval call binding the contract method 0x375e24c1.
//
// Solidity: function isManagementProxy(uint32 _point, address _proxy) view returns(bool result)
func (_Azimuth *AzimuthCallerSession) IsManagementProxy(_point uint32, _proxy common.Address) (bool, error) {
	return _Azimuth.Contract.IsManagementProxy(&_Azimuth.CallOpts, _point, _proxy)
}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address _owner, address _operator) view returns(bool result)
func (_Azimuth *AzimuthCaller) IsOperator(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "isOperator", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address _owner, address _operator) view returns(bool result)
func (_Azimuth *AzimuthSession) IsOperator(_owner common.Address, _operator common.Address) (bool, error) {
	return _Azimuth.Contract.IsOperator(&_Azimuth.CallOpts, _owner, _operator)
}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address _owner, address _operator) view returns(bool result)
func (_Azimuth *AzimuthCallerSession) IsOperator(_owner common.Address, _operator common.Address) (bool, error) {
	return _Azimuth.Contract.IsOperator(&_Azimuth.CallOpts, _owner, _operator)
}

// IsOwner is a free data retrieval call binding the contract method 0xcaf590f9.
//
// Solidity: function isOwner(uint32 _point, address _address) view returns(bool result)
func (_Azimuth *AzimuthCaller) IsOwner(opts *bind.CallOpts, _point uint32, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "isOwner", _point, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0xcaf590f9.
//
// Solidity: function isOwner(uint32 _point, address _address) view returns(bool result)
func (_Azimuth *AzimuthSession) IsOwner(_point uint32, _address common.Address) (bool, error) {
	return _Azimuth.Contract.IsOwner(&_Azimuth.CallOpts, _point, _address)
}

// IsOwner is a free data retrieval call binding the contract method 0xcaf590f9.
//
// Solidity: function isOwner(uint32 _point, address _address) view returns(bool result)
func (_Azimuth *AzimuthCallerSession) IsOwner(_point uint32, _address common.Address) (bool, error) {
	return _Azimuth.Contract.IsOwner(&_Azimuth.CallOpts, _point, _address)
}

// IsRequestingEscapeTo is a free data retrieval call binding the contract method 0xfde90733.
//
// Solidity: function isRequestingEscapeTo(uint32 _point, uint32 _sponsor) view returns(bool equals)
func (_Azimuth *AzimuthCaller) IsRequestingEscapeTo(opts *bind.CallOpts, _point uint32, _sponsor uint32) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "isRequestingEscapeTo", _point, _sponsor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRequestingEscapeTo is a free data retrieval call binding the contract method 0xfde90733.
//
// Solidity: function isRequestingEscapeTo(uint32 _point, uint32 _sponsor) view returns(bool equals)
func (_Azimuth *AzimuthSession) IsRequestingEscapeTo(_point uint32, _sponsor uint32) (bool, error) {
	return _Azimuth.Contract.IsRequestingEscapeTo(&_Azimuth.CallOpts, _point, _sponsor)
}

// IsRequestingEscapeTo is a free data retrieval call binding the contract method 0xfde90733.
//
// Solidity: function isRequestingEscapeTo(uint32 _point, uint32 _sponsor) view returns(bool equals)
func (_Azimuth *AzimuthCallerSession) IsRequestingEscapeTo(_point uint32, _sponsor uint32) (bool, error) {
	return _Azimuth.Contract.IsRequestingEscapeTo(&_Azimuth.CallOpts, _point, _sponsor)
}

// IsSpawnProxy is a free data retrieval call binding the contract method 0x24ba1a46.
//
// Solidity: function isSpawnProxy(uint32 _point, address _proxy) view returns(bool result)
func (_Azimuth *AzimuthCaller) IsSpawnProxy(opts *bind.CallOpts, _point uint32, _proxy common.Address) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "isSpawnProxy", _point, _proxy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSpawnProxy is a free data retrieval call binding the contract method 0x24ba1a46.
//
// Solidity: function isSpawnProxy(uint32 _point, address _proxy) view returns(bool result)
func (_Azimuth *AzimuthSession) IsSpawnProxy(_point uint32, _proxy common.Address) (bool, error) {
	return _Azimuth.Contract.IsSpawnProxy(&_Azimuth.CallOpts, _point, _proxy)
}

// IsSpawnProxy is a free data retrieval call binding the contract method 0x24ba1a46.
//
// Solidity: function isSpawnProxy(uint32 _point, address _proxy) view returns(bool result)
func (_Azimuth *AzimuthCallerSession) IsSpawnProxy(_point uint32, _proxy common.Address) (bool, error) {
	return _Azimuth.Contract.IsSpawnProxy(&_Azimuth.CallOpts, _point, _proxy)
}

// IsSponsor is a free data retrieval call binding the contract method 0x08083039.
//
// Solidity: function isSponsor(uint32 _point, uint32 _sponsor) view returns(bool result)
func (_Azimuth *AzimuthCaller) IsSponsor(opts *bind.CallOpts, _point uint32, _sponsor uint32) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "isSponsor", _point, _sponsor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSponsor is a free data retrieval call binding the contract method 0x08083039.
//
// Solidity: function isSponsor(uint32 _point, uint32 _sponsor) view returns(bool result)
func (_Azimuth *AzimuthSession) IsSponsor(_point uint32, _sponsor uint32) (bool, error) {
	return _Azimuth.Contract.IsSponsor(&_Azimuth.CallOpts, _point, _sponsor)
}

// IsSponsor is a free data retrieval call binding the contract method 0x08083039.
//
// Solidity: function isSponsor(uint32 _point, uint32 _sponsor) view returns(bool result)
func (_Azimuth *AzimuthCallerSession) IsSponsor(_point uint32, _sponsor uint32) (bool, error) {
	return _Azimuth.Contract.IsSponsor(&_Azimuth.CallOpts, _point, _sponsor)
}

// IsTransferProxy is a free data retrieval call binding the contract method 0x3262207c.
//
// Solidity: function isTransferProxy(uint32 _point, address _proxy) view returns(bool result)
func (_Azimuth *AzimuthCaller) IsTransferProxy(opts *bind.CallOpts, _point uint32, _proxy common.Address) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "isTransferProxy", _point, _proxy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransferProxy is a free data retrieval call binding the contract method 0x3262207c.
//
// Solidity: function isTransferProxy(uint32 _point, address _proxy) view returns(bool result)
func (_Azimuth *AzimuthSession) IsTransferProxy(_point uint32, _proxy common.Address) (bool, error) {
	return _Azimuth.Contract.IsTransferProxy(&_Azimuth.CallOpts, _point, _proxy)
}

// IsTransferProxy is a free data retrieval call binding the contract method 0x3262207c.
//
// Solidity: function isTransferProxy(uint32 _point, address _proxy) view returns(bool result)
func (_Azimuth *AzimuthCallerSession) IsTransferProxy(_point uint32, _proxy common.Address) (bool, error) {
	return _Azimuth.Contract.IsTransferProxy(&_Azimuth.CallOpts, _point, _proxy)
}

// IsVotingProxy is a free data retrieval call binding the contract method 0xe213c840.
//
// Solidity: function isVotingProxy(uint32 _point, address _proxy) view returns(bool result)
func (_Azimuth *AzimuthCaller) IsVotingProxy(opts *bind.CallOpts, _point uint32, _proxy common.Address) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "isVotingProxy", _point, _proxy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVotingProxy is a free data retrieval call binding the contract method 0xe213c840.
//
// Solidity: function isVotingProxy(uint32 _point, address _proxy) view returns(bool result)
func (_Azimuth *AzimuthSession) IsVotingProxy(_point uint32, _proxy common.Address) (bool, error) {
	return _Azimuth.Contract.IsVotingProxy(&_Azimuth.CallOpts, _point, _proxy)
}

// IsVotingProxy is a free data retrieval call binding the contract method 0xe213c840.
//
// Solidity: function isVotingProxy(uint32 _point, address _proxy) view returns(bool result)
func (_Azimuth *AzimuthCallerSession) IsVotingProxy(_point uint32, _proxy common.Address) (bool, error) {
	return _Azimuth.Contract.IsVotingProxy(&_Azimuth.CallOpts, _point, _proxy)
}

// ManagerFor is a free data retrieval call binding the contract method 0x5a83df54.
//
// Solidity: function managerFor(address , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthCaller) ManagerFor(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "managerFor", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ManagerFor is a free data retrieval call binding the contract method 0x5a83df54.
//
// Solidity: function managerFor(address , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthSession) ManagerFor(arg0 common.Address, arg1 *big.Int) (uint32, error) {
	return _Azimuth.Contract.ManagerFor(&_Azimuth.CallOpts, arg0, arg1)
}

// ManagerFor is a free data retrieval call binding the contract method 0x5a83df54.
//
// Solidity: function managerFor(address , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthCallerSession) ManagerFor(arg0 common.Address, arg1 *big.Int) (uint32, error) {
	return _Azimuth.Contract.ManagerFor(&_Azimuth.CallOpts, arg0, arg1)
}

// ManagerForIndexes is a free data retrieval call binding the contract method 0x52c0683b.
//
// Solidity: function managerForIndexes(address , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthCaller) ManagerForIndexes(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "managerForIndexes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerForIndexes is a free data retrieval call binding the contract method 0x52c0683b.
//
// Solidity: function managerForIndexes(address , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthSession) ManagerForIndexes(arg0 common.Address, arg1 uint32) (*big.Int, error) {
	return _Azimuth.Contract.ManagerForIndexes(&_Azimuth.CallOpts, arg0, arg1)
}

// ManagerForIndexes is a free data retrieval call binding the contract method 0x52c0683b.
//
// Solidity: function managerForIndexes(address , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthCallerSession) ManagerForIndexes(arg0 common.Address, arg1 uint32) (*big.Int, error) {
	return _Azimuth.Contract.ManagerForIndexes(&_Azimuth.CallOpts, arg0, arg1)
}

// Operators is a free data retrieval call binding the contract method 0x2637588e.
//
// Solidity: function operators(address , address ) view returns(bool)
func (_Azimuth *AzimuthCaller) Operators(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "operators", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0x2637588e.
//
// Solidity: function operators(address , address ) view returns(bool)
func (_Azimuth *AzimuthSession) Operators(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Azimuth.Contract.Operators(&_Azimuth.CallOpts, arg0, arg1)
}

// Operators is a free data retrieval call binding the contract method 0x2637588e.
//
// Solidity: function operators(address , address ) view returns(bool)
func (_Azimuth *AzimuthCallerSession) Operators(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Azimuth.Contract.Operators(&_Azimuth.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Azimuth *AzimuthCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Azimuth *AzimuthSession) Owner() (common.Address, error) {
	return _Azimuth.Contract.Owner(&_Azimuth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Azimuth *AzimuthCallerSession) Owner() (common.Address, error) {
	return _Azimuth.Contract.Owner(&_Azimuth.CallOpts)
}

// PointOwnerIndexes is a free data retrieval call binding the contract method 0x474dbae7.
//
// Solidity: function pointOwnerIndexes(address , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthCaller) PointOwnerIndexes(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "pointOwnerIndexes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PointOwnerIndexes is a free data retrieval call binding the contract method 0x474dbae7.
//
// Solidity: function pointOwnerIndexes(address , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthSession) PointOwnerIndexes(arg0 common.Address, arg1 uint32) (*big.Int, error) {
	return _Azimuth.Contract.PointOwnerIndexes(&_Azimuth.CallOpts, arg0, arg1)
}

// PointOwnerIndexes is a free data retrieval call binding the contract method 0x474dbae7.
//
// Solidity: function pointOwnerIndexes(address , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthCallerSession) PointOwnerIndexes(arg0 common.Address, arg1 uint32) (*big.Int, error) {
	return _Azimuth.Contract.PointOwnerIndexes(&_Azimuth.CallOpts, arg0, arg1)
}

// Points is a free data retrieval call binding the contract method 0x63fa9a87.
//
// Solidity: function points(uint32 ) view returns(bytes32 encryptionKey, bytes32 authenticationKey, bool hasSponsor, bool active, bool escapeRequested, uint32 sponsor, uint32 escapeRequestedTo, uint32 cryptoSuiteVersion, uint32 keyRevisionNumber, uint32 continuityNumber)
func (_Azimuth *AzimuthCaller) Points(opts *bind.CallOpts, arg0 uint32) (struct {
	EncryptionKey      [32]byte
	AuthenticationKey  [32]byte
	HasSponsor         bool
	Active             bool
	EscapeRequested    bool
	Sponsor            uint32
	EscapeRequestedTo  uint32
	CryptoSuiteVersion uint32
	KeyRevisionNumber  uint32
	ContinuityNumber   uint32
}, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "points", arg0)

	outstruct := new(struct {
		EncryptionKey      [32]byte
		AuthenticationKey  [32]byte
		HasSponsor         bool
		Active             bool
		EscapeRequested    bool
		Sponsor            uint32
		EscapeRequestedTo  uint32
		CryptoSuiteVersion uint32
		KeyRevisionNumber  uint32
		ContinuityNumber   uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EncryptionKey = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.AuthenticationKey = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.HasSponsor = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Active = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.EscapeRequested = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Sponsor = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.EscapeRequestedTo = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.CryptoSuiteVersion = *abi.ConvertType(out[7], new(uint32)).(*uint32)
	outstruct.KeyRevisionNumber = *abi.ConvertType(out[8], new(uint32)).(*uint32)
	outstruct.ContinuityNumber = *abi.ConvertType(out[9], new(uint32)).(*uint32)

	return *outstruct, err

}

// Points is a free data retrieval call binding the contract method 0x63fa9a87.
//
// Solidity: function points(uint32 ) view returns(bytes32 encryptionKey, bytes32 authenticationKey, bool hasSponsor, bool active, bool escapeRequested, uint32 sponsor, uint32 escapeRequestedTo, uint32 cryptoSuiteVersion, uint32 keyRevisionNumber, uint32 continuityNumber)
func (_Azimuth *AzimuthSession) Points(arg0 uint32) (struct {
	EncryptionKey      [32]byte
	AuthenticationKey  [32]byte
	HasSponsor         bool
	Active             bool
	EscapeRequested    bool
	Sponsor            uint32
	EscapeRequestedTo  uint32
	CryptoSuiteVersion uint32
	KeyRevisionNumber  uint32
	ContinuityNumber   uint32
}, error) {
	return _Azimuth.Contract.Points(&_Azimuth.CallOpts, arg0)
}

// Points is a free data retrieval call binding the contract method 0x63fa9a87.
//
// Solidity: function points(uint32 ) view returns(bytes32 encryptionKey, bytes32 authenticationKey, bool hasSponsor, bool active, bool escapeRequested, uint32 sponsor, uint32 escapeRequestedTo, uint32 cryptoSuiteVersion, uint32 keyRevisionNumber, uint32 continuityNumber)
func (_Azimuth *AzimuthCallerSession) Points(arg0 uint32) (struct {
	EncryptionKey      [32]byte
	AuthenticationKey  [32]byte
	HasSponsor         bool
	Active             bool
	EscapeRequested    bool
	Sponsor            uint32
	EscapeRequestedTo  uint32
	CryptoSuiteVersion uint32
	KeyRevisionNumber  uint32
	ContinuityNumber   uint32
}, error) {
	return _Azimuth.Contract.Points(&_Azimuth.CallOpts, arg0)
}

// PointsOwnedBy is a free data retrieval call binding the contract method 0x24cf61b5.
//
// Solidity: function pointsOwnedBy(address , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthCaller) PointsOwnedBy(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "pointsOwnedBy", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PointsOwnedBy is a free data retrieval call binding the contract method 0x24cf61b5.
//
// Solidity: function pointsOwnedBy(address , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthSession) PointsOwnedBy(arg0 common.Address, arg1 *big.Int) (uint32, error) {
	return _Azimuth.Contract.PointsOwnedBy(&_Azimuth.CallOpts, arg0, arg1)
}

// PointsOwnedBy is a free data retrieval call binding the contract method 0x24cf61b5.
//
// Solidity: function pointsOwnedBy(address , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthCallerSession) PointsOwnedBy(arg0 common.Address, arg1 *big.Int) (uint32, error) {
	return _Azimuth.Contract.PointsOwnedBy(&_Azimuth.CallOpts, arg0, arg1)
}

// Rights is a free data retrieval call binding the contract method 0x04bc94e5.
//
// Solidity: function rights(uint32 ) view returns(address owner, address managementProxy, address spawnProxy, address votingProxy, address transferProxy)
func (_Azimuth *AzimuthCaller) Rights(opts *bind.CallOpts, arg0 uint32) (struct {
	Owner           common.Address
	ManagementProxy common.Address
	SpawnProxy      common.Address
	VotingProxy     common.Address
	TransferProxy   common.Address
}, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "rights", arg0)

	outstruct := new(struct {
		Owner           common.Address
		ManagementProxy common.Address
		SpawnProxy      common.Address
		VotingProxy     common.Address
		TransferProxy   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ManagementProxy = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.SpawnProxy = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.VotingProxy = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.TransferProxy = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Rights is a free data retrieval call binding the contract method 0x04bc94e5.
//
// Solidity: function rights(uint32 ) view returns(address owner, address managementProxy, address spawnProxy, address votingProxy, address transferProxy)
func (_Azimuth *AzimuthSession) Rights(arg0 uint32) (struct {
	Owner           common.Address
	ManagementProxy common.Address
	SpawnProxy      common.Address
	VotingProxy     common.Address
	TransferProxy   common.Address
}, error) {
	return _Azimuth.Contract.Rights(&_Azimuth.CallOpts, arg0)
}

// Rights is a free data retrieval call binding the contract method 0x04bc94e5.
//
// Solidity: function rights(uint32 ) view returns(address owner, address managementProxy, address spawnProxy, address votingProxy, address transferProxy)
func (_Azimuth *AzimuthCallerSession) Rights(arg0 uint32) (struct {
	Owner           common.Address
	ManagementProxy common.Address
	SpawnProxy      common.Address
	VotingProxy     common.Address
	TransferProxy   common.Address
}, error) {
	return _Azimuth.Contract.Rights(&_Azimuth.CallOpts, arg0)
}

// SpawningFor is a free data retrieval call binding the contract method 0xd5c5164f.
//
// Solidity: function spawningFor(address , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthCaller) SpawningFor(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "spawningFor", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SpawningFor is a free data retrieval call binding the contract method 0xd5c5164f.
//
// Solidity: function spawningFor(address , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthSession) SpawningFor(arg0 common.Address, arg1 *big.Int) (uint32, error) {
	return _Azimuth.Contract.SpawningFor(&_Azimuth.CallOpts, arg0, arg1)
}

// SpawningFor is a free data retrieval call binding the contract method 0xd5c5164f.
//
// Solidity: function spawningFor(address , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthCallerSession) SpawningFor(arg0 common.Address, arg1 *big.Int) (uint32, error) {
	return _Azimuth.Contract.SpawningFor(&_Azimuth.CallOpts, arg0, arg1)
}

// SpawningForIndexes is a free data retrieval call binding the contract method 0x76944d5b.
//
// Solidity: function spawningForIndexes(address , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthCaller) SpawningForIndexes(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "spawningForIndexes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpawningForIndexes is a free data retrieval call binding the contract method 0x76944d5b.
//
// Solidity: function spawningForIndexes(address , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthSession) SpawningForIndexes(arg0 common.Address, arg1 uint32) (*big.Int, error) {
	return _Azimuth.Contract.SpawningForIndexes(&_Azimuth.CallOpts, arg0, arg1)
}

// SpawningForIndexes is a free data retrieval call binding the contract method 0x76944d5b.
//
// Solidity: function spawningForIndexes(address , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthCallerSession) SpawningForIndexes(arg0 common.Address, arg1 uint32) (*big.Int, error) {
	return _Azimuth.Contract.SpawningForIndexes(&_Azimuth.CallOpts, arg0, arg1)
}

// Sponsoring is a free data retrieval call binding the contract method 0x12ace2e6.
//
// Solidity: function sponsoring(uint32 , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthCaller) Sponsoring(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "sponsoring", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Sponsoring is a free data retrieval call binding the contract method 0x12ace2e6.
//
// Solidity: function sponsoring(uint32 , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthSession) Sponsoring(arg0 uint32, arg1 *big.Int) (uint32, error) {
	return _Azimuth.Contract.Sponsoring(&_Azimuth.CallOpts, arg0, arg1)
}

// Sponsoring is a free data retrieval call binding the contract method 0x12ace2e6.
//
// Solidity: function sponsoring(uint32 , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthCallerSession) Sponsoring(arg0 uint32, arg1 *big.Int) (uint32, error) {
	return _Azimuth.Contract.Sponsoring(&_Azimuth.CallOpts, arg0, arg1)
}

// SponsoringIndexes is a free data retrieval call binding the contract method 0x2cce0e29.
//
// Solidity: function sponsoringIndexes(uint32 , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthCaller) SponsoringIndexes(opts *bind.CallOpts, arg0 uint32, arg1 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "sponsoringIndexes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SponsoringIndexes is a free data retrieval call binding the contract method 0x2cce0e29.
//
// Solidity: function sponsoringIndexes(uint32 , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthSession) SponsoringIndexes(arg0 uint32, arg1 uint32) (*big.Int, error) {
	return _Azimuth.Contract.SponsoringIndexes(&_Azimuth.CallOpts, arg0, arg1)
}

// SponsoringIndexes is a free data retrieval call binding the contract method 0x2cce0e29.
//
// Solidity: function sponsoringIndexes(uint32 , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthCallerSession) SponsoringIndexes(arg0 uint32, arg1 uint32) (*big.Int, error) {
	return _Azimuth.Contract.SponsoringIndexes(&_Azimuth.CallOpts, arg0, arg1)
}

// TransferringFor is a free data retrieval call binding the contract method 0x073167da.
//
// Solidity: function transferringFor(address , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthCaller) TransferringFor(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "transferringFor", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TransferringFor is a free data retrieval call binding the contract method 0x073167da.
//
// Solidity: function transferringFor(address , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthSession) TransferringFor(arg0 common.Address, arg1 *big.Int) (uint32, error) {
	return _Azimuth.Contract.TransferringFor(&_Azimuth.CallOpts, arg0, arg1)
}

// TransferringFor is a free data retrieval call binding the contract method 0x073167da.
//
// Solidity: function transferringFor(address , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthCallerSession) TransferringFor(arg0 common.Address, arg1 *big.Int) (uint32, error) {
	return _Azimuth.Contract.TransferringFor(&_Azimuth.CallOpts, arg0, arg1)
}

// TransferringForIndexes is a free data retrieval call binding the contract method 0xd2588e20.
//
// Solidity: function transferringForIndexes(address , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthCaller) TransferringForIndexes(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "transferringForIndexes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferringForIndexes is a free data retrieval call binding the contract method 0xd2588e20.
//
// Solidity: function transferringForIndexes(address , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthSession) TransferringForIndexes(arg0 common.Address, arg1 uint32) (*big.Int, error) {
	return _Azimuth.Contract.TransferringForIndexes(&_Azimuth.CallOpts, arg0, arg1)
}

// TransferringForIndexes is a free data retrieval call binding the contract method 0xd2588e20.
//
// Solidity: function transferringForIndexes(address , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthCallerSession) TransferringForIndexes(arg0 common.Address, arg1 uint32) (*big.Int, error) {
	return _Azimuth.Contract.TransferringForIndexes(&_Azimuth.CallOpts, arg0, arg1)
}

// VotingFor is a free data retrieval call binding the contract method 0x04bbb61d.
//
// Solidity: function votingFor(address , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthCaller) VotingFor(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "votingFor", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// VotingFor is a free data retrieval call binding the contract method 0x04bbb61d.
//
// Solidity: function votingFor(address , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthSession) VotingFor(arg0 common.Address, arg1 *big.Int) (uint32, error) {
	return _Azimuth.Contract.VotingFor(&_Azimuth.CallOpts, arg0, arg1)
}

// VotingFor is a free data retrieval call binding the contract method 0x04bbb61d.
//
// Solidity: function votingFor(address , uint256 ) view returns(uint32)
func (_Azimuth *AzimuthCallerSession) VotingFor(arg0 common.Address, arg1 *big.Int) (uint32, error) {
	return _Azimuth.Contract.VotingFor(&_Azimuth.CallOpts, arg0, arg1)
}

// VotingForIndexes is a free data retrieval call binding the contract method 0x37de45c1.
//
// Solidity: function votingForIndexes(address , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthCaller) VotingForIndexes(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Azimuth.contract.Call(opts, &out, "votingForIndexes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingForIndexes is a free data retrieval call binding the contract method 0x37de45c1.
//
// Solidity: function votingForIndexes(address , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthSession) VotingForIndexes(arg0 common.Address, arg1 uint32) (*big.Int, error) {
	return _Azimuth.Contract.VotingForIndexes(&_Azimuth.CallOpts, arg0, arg1)
}

// VotingForIndexes is a free data retrieval call binding the contract method 0x37de45c1.
//
// Solidity: function votingForIndexes(address , uint32 ) view returns(uint256)
func (_Azimuth *AzimuthCallerSession) VotingForIndexes(arg0 common.Address, arg1 uint32) (*big.Int, error) {
	return _Azimuth.Contract.VotingForIndexes(&_Azimuth.CallOpts, arg0, arg1)
}

// ActivatePoint is a paid mutator transaction binding the contract method 0x7bc702a1.
//
// Solidity: function activatePoint(uint32 _point) returns()
func (_Azimuth *AzimuthTransactor) ActivatePoint(opts *bind.TransactOpts, _point uint32) (*types.Transaction, error) {
	return _Azimuth.contract.Transact(opts, "activatePoint", _point)
}

// ActivatePoint is a paid mutator transaction binding the contract method 0x7bc702a1.
//
// Solidity: function activatePoint(uint32 _point) returns()
func (_Azimuth *AzimuthSession) ActivatePoint(_point uint32) (*types.Transaction, error) {
	return _Azimuth.Contract.ActivatePoint(&_Azimuth.TransactOpts, _point)
}

// ActivatePoint is a paid mutator transaction binding the contract method 0x7bc702a1.
//
// Solidity: function activatePoint(uint32 _point) returns()
func (_Azimuth *AzimuthTransactorSession) ActivatePoint(_point uint32) (*types.Transaction, error) {
	return _Azimuth.Contract.ActivatePoint(&_Azimuth.TransactOpts, _point)
}

// CancelEscape is a paid mutator transaction binding the contract method 0xc6d761d4.
//
// Solidity: function cancelEscape(uint32 _point) returns()
func (_Azimuth *AzimuthTransactor) CancelEscape(opts *bind.TransactOpts, _point uint32) (*types.Transaction, error) {
	return _Azimuth.contract.Transact(opts, "cancelEscape", _point)
}

// CancelEscape is a paid mutator transaction binding the contract method 0xc6d761d4.
//
// Solidity: function cancelEscape(uint32 _point) returns()
func (_Azimuth *AzimuthSession) CancelEscape(_point uint32) (*types.Transaction, error) {
	return _Azimuth.Contract.CancelEscape(&_Azimuth.TransactOpts, _point)
}

// CancelEscape is a paid mutator transaction binding the contract method 0xc6d761d4.
//
// Solidity: function cancelEscape(uint32 _point) returns()
func (_Azimuth *AzimuthTransactorSession) CancelEscape(_point uint32) (*types.Transaction, error) {
	return _Azimuth.Contract.CancelEscape(&_Azimuth.TransactOpts, _point)
}

// DoEscape is a paid mutator transaction binding the contract method 0x13063180.
//
// Solidity: function doEscape(uint32 _point) returns()
func (_Azimuth *AzimuthTransactor) DoEscape(opts *bind.TransactOpts, _point uint32) (*types.Transaction, error) {
	return _Azimuth.contract.Transact(opts, "doEscape", _point)
}

// DoEscape is a paid mutator transaction binding the contract method 0x13063180.
//
// Solidity: function doEscape(uint32 _point) returns()
func (_Azimuth *AzimuthSession) DoEscape(_point uint32) (*types.Transaction, error) {
	return _Azimuth.Contract.DoEscape(&_Azimuth.TransactOpts, _point)
}

// DoEscape is a paid mutator transaction binding the contract method 0x13063180.
//
// Solidity: function doEscape(uint32 _point) returns()
func (_Azimuth *AzimuthTransactorSession) DoEscape(_point uint32) (*types.Transaction, error) {
	return _Azimuth.Contract.DoEscape(&_Azimuth.TransactOpts, _point)
}

// IncrementContinuityNumber is a paid mutator transaction binding the contract method 0xf4914919.
//
// Solidity: function incrementContinuityNumber(uint32 _point) returns()
func (_Azimuth *AzimuthTransactor) IncrementContinuityNumber(opts *bind.TransactOpts, _point uint32) (*types.Transaction, error) {
	return _Azimuth.contract.Transact(opts, "incrementContinuityNumber", _point)
}

// IncrementContinuityNumber is a paid mutator transaction binding the contract method 0xf4914919.
//
// Solidity: function incrementContinuityNumber(uint32 _point) returns()
func (_Azimuth *AzimuthSession) IncrementContinuityNumber(_point uint32) (*types.Transaction, error) {
	return _Azimuth.Contract.IncrementContinuityNumber(&_Azimuth.TransactOpts, _point)
}

// IncrementContinuityNumber is a paid mutator transaction binding the contract method 0xf4914919.
//
// Solidity: function incrementContinuityNumber(uint32 _point) returns()
func (_Azimuth *AzimuthTransactorSession) IncrementContinuityNumber(_point uint32) (*types.Transaction, error) {
	return _Azimuth.Contract.IncrementContinuityNumber(&_Azimuth.TransactOpts, _point)
}

// LoseSponsor is a paid mutator transaction binding the contract method 0xbc562b9e.
//
// Solidity: function loseSponsor(uint32 _point) returns()
func (_Azimuth *AzimuthTransactor) LoseSponsor(opts *bind.TransactOpts, _point uint32) (*types.Transaction, error) {
	return _Azimuth.contract.Transact(opts, "loseSponsor", _point)
}

// LoseSponsor is a paid mutator transaction binding the contract method 0xbc562b9e.
//
// Solidity: function loseSponsor(uint32 _point) returns()
func (_Azimuth *AzimuthSession) LoseSponsor(_point uint32) (*types.Transaction, error) {
	return _Azimuth.Contract.LoseSponsor(&_Azimuth.TransactOpts, _point)
}

// LoseSponsor is a paid mutator transaction binding the contract method 0xbc562b9e.
//
// Solidity: function loseSponsor(uint32 _point) returns()
func (_Azimuth *AzimuthTransactorSession) LoseSponsor(_point uint32) (*types.Transaction, error) {
	return _Azimuth.Contract.LoseSponsor(&_Azimuth.TransactOpts, _point)
}

// RegisterSpawned is a paid mutator transaction binding the contract method 0xc17ad9d8.
//
// Solidity: function registerSpawned(uint32 _point) returns()
func (_Azimuth *AzimuthTransactor) RegisterSpawned(opts *bind.TransactOpts, _point uint32) (*types.Transaction, error) {
	return _Azimuth.contract.Transact(opts, "registerSpawned", _point)
}

// RegisterSpawned is a paid mutator transaction binding the contract method 0xc17ad9d8.
//
// Solidity: function registerSpawned(uint32 _point) returns()
func (_Azimuth *AzimuthSession) RegisterSpawned(_point uint32) (*types.Transaction, error) {
	return _Azimuth.Contract.RegisterSpawned(&_Azimuth.TransactOpts, _point)
}

// RegisterSpawned is a paid mutator transaction binding the contract method 0xc17ad9d8.
//
// Solidity: function registerSpawned(uint32 _point) returns()
func (_Azimuth *AzimuthTransactorSession) RegisterSpawned(_point uint32) (*types.Transaction, error) {
	return _Azimuth.Contract.RegisterSpawned(&_Azimuth.TransactOpts, _point)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Azimuth *AzimuthTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Azimuth.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Azimuth *AzimuthSession) RenounceOwnership() (*types.Transaction, error) {
	return _Azimuth.Contract.RenounceOwnership(&_Azimuth.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Azimuth *AzimuthTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Azimuth.Contract.RenounceOwnership(&_Azimuth.TransactOpts)
}

// SetDnsDomains is a paid mutator transaction binding the contract method 0xbac55edd.
//
// Solidity: function setDnsDomains(string _primary, string _secondary, string _tertiary) returns()
func (_Azimuth *AzimuthTransactor) SetDnsDomains(opts *bind.TransactOpts, _primary string, _secondary string, _tertiary string) (*types.Transaction, error) {
	return _Azimuth.contract.Transact(opts, "setDnsDomains", _primary, _secondary, _tertiary)
}

// SetDnsDomains is a paid mutator transaction binding the contract method 0xbac55edd.
//
// Solidity: function setDnsDomains(string _primary, string _secondary, string _tertiary) returns()
func (_Azimuth *AzimuthSession) SetDnsDomains(_primary string, _secondary string, _tertiary string) (*types.Transaction, error) {
	return _Azimuth.Contract.SetDnsDomains(&_Azimuth.TransactOpts, _primary, _secondary, _tertiary)
}

// SetDnsDomains is a paid mutator transaction binding the contract method 0xbac55edd.
//
// Solidity: function setDnsDomains(string _primary, string _secondary, string _tertiary) returns()
func (_Azimuth *AzimuthTransactorSession) SetDnsDomains(_primary string, _secondary string, _tertiary string) (*types.Transaction, error) {
	return _Azimuth.Contract.SetDnsDomains(&_Azimuth.TransactOpts, _primary, _secondary, _tertiary)
}

// SetEscapeRequest is a paid mutator transaction binding the contract method 0xa6345859.
//
// Solidity: function setEscapeRequest(uint32 _point, uint32 _sponsor) returns()
func (_Azimuth *AzimuthTransactor) SetEscapeRequest(opts *bind.TransactOpts, _point uint32, _sponsor uint32) (*types.Transaction, error) {
	return _Azimuth.contract.Transact(opts, "setEscapeRequest", _point, _sponsor)
}

// SetEscapeRequest is a paid mutator transaction binding the contract method 0xa6345859.
//
// Solidity: function setEscapeRequest(uint32 _point, uint32 _sponsor) returns()
func (_Azimuth *AzimuthSession) SetEscapeRequest(_point uint32, _sponsor uint32) (*types.Transaction, error) {
	return _Azimuth.Contract.SetEscapeRequest(&_Azimuth.TransactOpts, _point, _sponsor)
}

// SetEscapeRequest is a paid mutator transaction binding the contract method 0xa6345859.
//
// Solidity: function setEscapeRequest(uint32 _point, uint32 _sponsor) returns()
func (_Azimuth *AzimuthTransactorSession) SetEscapeRequest(_point uint32, _sponsor uint32) (*types.Transaction, error) {
	return _Azimuth.Contract.SetEscapeRequest(&_Azimuth.TransactOpts, _point, _sponsor)
}

// SetKeys is a paid mutator transaction binding the contract method 0xf9b87d40.
//
// Solidity: function setKeys(uint32 _point, bytes32 _encryptionKey, bytes32 _authenticationKey, uint32 _cryptoSuiteVersion) returns()
func (_Azimuth *AzimuthTransactor) SetKeys(opts *bind.TransactOpts, _point uint32, _encryptionKey [32]byte, _authenticationKey [32]byte, _cryptoSuiteVersion uint32) (*types.Transaction, error) {
	return _Azimuth.contract.Transact(opts, "setKeys", _point, _encryptionKey, _authenticationKey, _cryptoSuiteVersion)
}

// SetKeys is a paid mutator transaction binding the contract method 0xf9b87d40.
//
// Solidity: function setKeys(uint32 _point, bytes32 _encryptionKey, bytes32 _authenticationKey, uint32 _cryptoSuiteVersion) returns()
func (_Azimuth *AzimuthSession) SetKeys(_point uint32, _encryptionKey [32]byte, _authenticationKey [32]byte, _cryptoSuiteVersion uint32) (*types.Transaction, error) {
	return _Azimuth.Contract.SetKeys(&_Azimuth.TransactOpts, _point, _encryptionKey, _authenticationKey, _cryptoSuiteVersion)
}

// SetKeys is a paid mutator transaction binding the contract method 0xf9b87d40.
//
// Solidity: function setKeys(uint32 _point, bytes32 _encryptionKey, bytes32 _authenticationKey, uint32 _cryptoSuiteVersion) returns()
func (_Azimuth *AzimuthTransactorSession) SetKeys(_point uint32, _encryptionKey [32]byte, _authenticationKey [32]byte, _cryptoSuiteVersion uint32) (*types.Transaction, error) {
	return _Azimuth.Contract.SetKeys(&_Azimuth.TransactOpts, _point, _encryptionKey, _authenticationKey, _cryptoSuiteVersion)
}

// SetManagementProxy is a paid mutator transaction binding the contract method 0x8866bb2c.
//
// Solidity: function setManagementProxy(uint32 _point, address _proxy) returns()
func (_Azimuth *AzimuthTransactor) SetManagementProxy(opts *bind.TransactOpts, _point uint32, _proxy common.Address) (*types.Transaction, error) {
	return _Azimuth.contract.Transact(opts, "setManagementProxy", _point, _proxy)
}

// SetManagementProxy is a paid mutator transaction binding the contract method 0x8866bb2c.
//
// Solidity: function setManagementProxy(uint32 _point, address _proxy) returns()
func (_Azimuth *AzimuthSession) SetManagementProxy(_point uint32, _proxy common.Address) (*types.Transaction, error) {
	return _Azimuth.Contract.SetManagementProxy(&_Azimuth.TransactOpts, _point, _proxy)
}

// SetManagementProxy is a paid mutator transaction binding the contract method 0x8866bb2c.
//
// Solidity: function setManagementProxy(uint32 _point, address _proxy) returns()
func (_Azimuth *AzimuthTransactorSession) SetManagementProxy(_point uint32, _proxy common.Address) (*types.Transaction, error) {
	return _Azimuth.Contract.SetManagementProxy(&_Azimuth.TransactOpts, _point, _proxy)
}

// SetOperator is a paid mutator transaction binding the contract method 0xbc735d90.
//
// Solidity: function setOperator(address _owner, address _operator, bool _approved) returns()
func (_Azimuth *AzimuthTransactor) SetOperator(opts *bind.TransactOpts, _owner common.Address, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Azimuth.contract.Transact(opts, "setOperator", _owner, _operator, _approved)
}

// SetOperator is a paid mutator transaction binding the contract method 0xbc735d90.
//
// Solidity: function setOperator(address _owner, address _operator, bool _approved) returns()
func (_Azimuth *AzimuthSession) SetOperator(_owner common.Address, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Azimuth.Contract.SetOperator(&_Azimuth.TransactOpts, _owner, _operator, _approved)
}

// SetOperator is a paid mutator transaction binding the contract method 0xbc735d90.
//
// Solidity: function setOperator(address _owner, address _operator, bool _approved) returns()
func (_Azimuth *AzimuthTransactorSession) SetOperator(_owner common.Address, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Azimuth.Contract.SetOperator(&_Azimuth.TransactOpts, _owner, _operator, _approved)
}

// SetOwner is a paid mutator transaction binding the contract method 0xddc35950.
//
// Solidity: function setOwner(uint32 _point, address _owner) returns()
func (_Azimuth *AzimuthTransactor) SetOwner(opts *bind.TransactOpts, _point uint32, _owner common.Address) (*types.Transaction, error) {
	return _Azimuth.contract.Transact(opts, "setOwner", _point, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0xddc35950.
//
// Solidity: function setOwner(uint32 _point, address _owner) returns()
func (_Azimuth *AzimuthSession) SetOwner(_point uint32, _owner common.Address) (*types.Transaction, error) {
	return _Azimuth.Contract.SetOwner(&_Azimuth.TransactOpts, _point, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0xddc35950.
//
// Solidity: function setOwner(uint32 _point, address _owner) returns()
func (_Azimuth *AzimuthTransactorSession) SetOwner(_point uint32, _owner common.Address) (*types.Transaction, error) {
	return _Azimuth.Contract.SetOwner(&_Azimuth.TransactOpts, _point, _owner)
}

// SetSpawnProxy is a paid mutator transaction binding the contract method 0x2a19642c.
//
// Solidity: function setSpawnProxy(uint32 _point, address _proxy) returns()
func (_Azimuth *AzimuthTransactor) SetSpawnProxy(opts *bind.TransactOpts, _point uint32, _proxy common.Address) (*types.Transaction, error) {
	return _Azimuth.contract.Transact(opts, "setSpawnProxy", _point, _proxy)
}

// SetSpawnProxy is a paid mutator transaction binding the contract method 0x2a19642c.
//
// Solidity: function setSpawnProxy(uint32 _point, address _proxy) returns()
func (_Azimuth *AzimuthSession) SetSpawnProxy(_point uint32, _proxy common.Address) (*types.Transaction, error) {
	return _Azimuth.Contract.SetSpawnProxy(&_Azimuth.TransactOpts, _point, _proxy)
}

// SetSpawnProxy is a paid mutator transaction binding the contract method 0x2a19642c.
//
// Solidity: function setSpawnProxy(uint32 _point, address _proxy) returns()
func (_Azimuth *AzimuthTransactorSession) SetSpawnProxy(_point uint32, _proxy common.Address) (*types.Transaction, error) {
	return _Azimuth.Contract.SetSpawnProxy(&_Azimuth.TransactOpts, _point, _proxy)
}

// SetTransferProxy is a paid mutator transaction binding the contract method 0x2c7ba564.
//
// Solidity: function setTransferProxy(uint32 _point, address _proxy) returns()
func (_Azimuth *AzimuthTransactor) SetTransferProxy(opts *bind.TransactOpts, _point uint32, _proxy common.Address) (*types.Transaction, error) {
	return _Azimuth.contract.Transact(opts, "setTransferProxy", _point, _proxy)
}

// SetTransferProxy is a paid mutator transaction binding the contract method 0x2c7ba564.
//
// Solidity: function setTransferProxy(uint32 _point, address _proxy) returns()
func (_Azimuth *AzimuthSession) SetTransferProxy(_point uint32, _proxy common.Address) (*types.Transaction, error) {
	return _Azimuth.Contract.SetTransferProxy(&_Azimuth.TransactOpts, _point, _proxy)
}

// SetTransferProxy is a paid mutator transaction binding the contract method 0x2c7ba564.
//
// Solidity: function setTransferProxy(uint32 _point, address _proxy) returns()
func (_Azimuth *AzimuthTransactorSession) SetTransferProxy(_point uint32, _proxy common.Address) (*types.Transaction, error) {
	return _Azimuth.Contract.SetTransferProxy(&_Azimuth.TransactOpts, _point, _proxy)
}

// SetVotingProxy is a paid mutator transaction binding the contract method 0xa297c1d8.
//
// Solidity: function setVotingProxy(uint32 _point, address _proxy) returns()
func (_Azimuth *AzimuthTransactor) SetVotingProxy(opts *bind.TransactOpts, _point uint32, _proxy common.Address) (*types.Transaction, error) {
	return _Azimuth.contract.Transact(opts, "setVotingProxy", _point, _proxy)
}

// SetVotingProxy is a paid mutator transaction binding the contract method 0xa297c1d8.
//
// Solidity: function setVotingProxy(uint32 _point, address _proxy) returns()
func (_Azimuth *AzimuthSession) SetVotingProxy(_point uint32, _proxy common.Address) (*types.Transaction, error) {
	return _Azimuth.Contract.SetVotingProxy(&_Azimuth.TransactOpts, _point, _proxy)
}

// SetVotingProxy is a paid mutator transaction binding the contract method 0xa297c1d8.
//
// Solidity: function setVotingProxy(uint32 _point, address _proxy) returns()
func (_Azimuth *AzimuthTransactorSession) SetVotingProxy(_point uint32, _proxy common.Address) (*types.Transaction, error) {
	return _Azimuth.Contract.SetVotingProxy(&_Azimuth.TransactOpts, _point, _proxy)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Azimuth *AzimuthTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Azimuth.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Azimuth *AzimuthSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Azimuth.Contract.TransferOwnership(&_Azimuth.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Azimuth *AzimuthTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Azimuth.Contract.TransferOwnership(&_Azimuth.TransactOpts, _newOwner)
}

// AzimuthActivatedIterator is returned from FilterActivated and is used to iterate over the raw logs and unpacked data for Activated events raised by the Azimuth contract.
type AzimuthActivatedIterator struct {
	Event *AzimuthActivated // Event containing the contract specifics and raw log

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
func (it *AzimuthActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AzimuthActivated)
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
		it.Event = new(AzimuthActivated)
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
func (it *AzimuthActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AzimuthActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AzimuthActivated represents a Activated event raised by the Azimuth contract.
type AzimuthActivated struct {
	Point uint32
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterActivated is a free log retrieval operation binding the contract event 0xe74c03809d0769e1b1f706cc8414258cd1f3b6fe020cd15d0165c210ba503a0f.
//
// Solidity: event Activated(uint32 indexed point)
func (_Azimuth *AzimuthFilterer) FilterActivated(opts *bind.FilterOpts, point []uint32) (*AzimuthActivatedIterator, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _Azimuth.contract.FilterLogs(opts, "Activated", pointRule)
	if err != nil {
		return nil, err
	}
	return &AzimuthActivatedIterator{contract: _Azimuth.contract, event: "Activated", logs: logs, sub: sub}, nil
}

// WatchActivated is a free log subscription operation binding the contract event 0xe74c03809d0769e1b1f706cc8414258cd1f3b6fe020cd15d0165c210ba503a0f.
//
// Solidity: event Activated(uint32 indexed point)
func (_Azimuth *AzimuthFilterer) WatchActivated(opts *bind.WatchOpts, sink chan<- *AzimuthActivated, point []uint32) (event.Subscription, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _Azimuth.contract.WatchLogs(opts, "Activated", pointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AzimuthActivated)
				if err := _Azimuth.contract.UnpackLog(event, "Activated", log); err != nil {
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

// ParseActivated is a log parse operation binding the contract event 0xe74c03809d0769e1b1f706cc8414258cd1f3b6fe020cd15d0165c210ba503a0f.
//
// Solidity: event Activated(uint32 indexed point)
func (_Azimuth *AzimuthFilterer) ParseActivated(log types.Log) (*AzimuthActivated, error) {
	event := new(AzimuthActivated)
	if err := _Azimuth.contract.UnpackLog(event, "Activated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AzimuthBrokeContinuityIterator is returned from FilterBrokeContinuity and is used to iterate over the raw logs and unpacked data for BrokeContinuity events raised by the Azimuth contract.
type AzimuthBrokeContinuityIterator struct {
	Event *AzimuthBrokeContinuity // Event containing the contract specifics and raw log

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
func (it *AzimuthBrokeContinuityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AzimuthBrokeContinuity)
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
		it.Event = new(AzimuthBrokeContinuity)
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
func (it *AzimuthBrokeContinuityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AzimuthBrokeContinuityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AzimuthBrokeContinuity represents a BrokeContinuity event raised by the Azimuth contract.
type AzimuthBrokeContinuity struct {
	Point  uint32
	Number uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBrokeContinuity is a free log retrieval operation binding the contract event 0x29294799f1c21a37ef838e15f79dd91bcee2df99d63cd1c18ac968b129514e6e.
//
// Solidity: event BrokeContinuity(uint32 indexed point, uint32 number)
func (_Azimuth *AzimuthFilterer) FilterBrokeContinuity(opts *bind.FilterOpts, point []uint32) (*AzimuthBrokeContinuityIterator, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _Azimuth.contract.FilterLogs(opts, "BrokeContinuity", pointRule)
	if err != nil {
		return nil, err
	}
	return &AzimuthBrokeContinuityIterator{contract: _Azimuth.contract, event: "BrokeContinuity", logs: logs, sub: sub}, nil
}

// WatchBrokeContinuity is a free log subscription operation binding the contract event 0x29294799f1c21a37ef838e15f79dd91bcee2df99d63cd1c18ac968b129514e6e.
//
// Solidity: event BrokeContinuity(uint32 indexed point, uint32 number)
func (_Azimuth *AzimuthFilterer) WatchBrokeContinuity(opts *bind.WatchOpts, sink chan<- *AzimuthBrokeContinuity, point []uint32) (event.Subscription, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _Azimuth.contract.WatchLogs(opts, "BrokeContinuity", pointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AzimuthBrokeContinuity)
				if err := _Azimuth.contract.UnpackLog(event, "BrokeContinuity", log); err != nil {
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

// ParseBrokeContinuity is a log parse operation binding the contract event 0x29294799f1c21a37ef838e15f79dd91bcee2df99d63cd1c18ac968b129514e6e.
//
// Solidity: event BrokeContinuity(uint32 indexed point, uint32 number)
func (_Azimuth *AzimuthFilterer) ParseBrokeContinuity(log types.Log) (*AzimuthBrokeContinuity, error) {
	event := new(AzimuthBrokeContinuity)
	if err := _Azimuth.contract.UnpackLog(event, "BrokeContinuity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AzimuthChangedDnsIterator is returned from FilterChangedDns and is used to iterate over the raw logs and unpacked data for ChangedDns events raised by the Azimuth contract.
type AzimuthChangedDnsIterator struct {
	Event *AzimuthChangedDns // Event containing the contract specifics and raw log

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
func (it *AzimuthChangedDnsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AzimuthChangedDns)
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
		it.Event = new(AzimuthChangedDns)
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
func (it *AzimuthChangedDnsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AzimuthChangedDnsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AzimuthChangedDns represents a ChangedDns event raised by the Azimuth contract.
type AzimuthChangedDns struct {
	Primary   string
	Secondary string
	Tertiary  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangedDns is a free log retrieval operation binding the contract event 0xfafd04ade1daae2e1fdb0fc1cc6a899fd424063ed5c92120e67e073053b94898.
//
// Solidity: event ChangedDns(string primary, string secondary, string tertiary)
func (_Azimuth *AzimuthFilterer) FilterChangedDns(opts *bind.FilterOpts) (*AzimuthChangedDnsIterator, error) {

	logs, sub, err := _Azimuth.contract.FilterLogs(opts, "ChangedDns")
	if err != nil {
		return nil, err
	}
	return &AzimuthChangedDnsIterator{contract: _Azimuth.contract, event: "ChangedDns", logs: logs, sub: sub}, nil
}

// WatchChangedDns is a free log subscription operation binding the contract event 0xfafd04ade1daae2e1fdb0fc1cc6a899fd424063ed5c92120e67e073053b94898.
//
// Solidity: event ChangedDns(string primary, string secondary, string tertiary)
func (_Azimuth *AzimuthFilterer) WatchChangedDns(opts *bind.WatchOpts, sink chan<- *AzimuthChangedDns) (event.Subscription, error) {

	logs, sub, err := _Azimuth.contract.WatchLogs(opts, "ChangedDns")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AzimuthChangedDns)
				if err := _Azimuth.contract.UnpackLog(event, "ChangedDns", log); err != nil {
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

// ParseChangedDns is a log parse operation binding the contract event 0xfafd04ade1daae2e1fdb0fc1cc6a899fd424063ed5c92120e67e073053b94898.
//
// Solidity: event ChangedDns(string primary, string secondary, string tertiary)
func (_Azimuth *AzimuthFilterer) ParseChangedDns(log types.Log) (*AzimuthChangedDns, error) {
	event := new(AzimuthChangedDns)
	if err := _Azimuth.contract.UnpackLog(event, "ChangedDns", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AzimuthChangedKeysIterator is returned from FilterChangedKeys and is used to iterate over the raw logs and unpacked data for ChangedKeys events raised by the Azimuth contract.
type AzimuthChangedKeysIterator struct {
	Event *AzimuthChangedKeys // Event containing the contract specifics and raw log

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
func (it *AzimuthChangedKeysIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AzimuthChangedKeys)
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
		it.Event = new(AzimuthChangedKeys)
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
func (it *AzimuthChangedKeysIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AzimuthChangedKeysIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AzimuthChangedKeys represents a ChangedKeys event raised by the Azimuth contract.
type AzimuthChangedKeys struct {
	Point              uint32
	EncryptionKey      [32]byte
	AuthenticationKey  [32]byte
	CryptoSuiteVersion uint32
	KeyRevisionNumber  uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterChangedKeys is a free log retrieval operation binding the contract event 0xaa10e7a0117d4323f1d99d630ec169bebb3a988e895770e351987e01ff5423d5.
//
// Solidity: event ChangedKeys(uint32 indexed point, bytes32 encryptionKey, bytes32 authenticationKey, uint32 cryptoSuiteVersion, uint32 keyRevisionNumber)
func (_Azimuth *AzimuthFilterer) FilterChangedKeys(opts *bind.FilterOpts, point []uint32) (*AzimuthChangedKeysIterator, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _Azimuth.contract.FilterLogs(opts, "ChangedKeys", pointRule)
	if err != nil {
		return nil, err
	}
	return &AzimuthChangedKeysIterator{contract: _Azimuth.contract, event: "ChangedKeys", logs: logs, sub: sub}, nil
}

// WatchChangedKeys is a free log subscription operation binding the contract event 0xaa10e7a0117d4323f1d99d630ec169bebb3a988e895770e351987e01ff5423d5.
//
// Solidity: event ChangedKeys(uint32 indexed point, bytes32 encryptionKey, bytes32 authenticationKey, uint32 cryptoSuiteVersion, uint32 keyRevisionNumber)
func (_Azimuth *AzimuthFilterer) WatchChangedKeys(opts *bind.WatchOpts, sink chan<- *AzimuthChangedKeys, point []uint32) (event.Subscription, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _Azimuth.contract.WatchLogs(opts, "ChangedKeys", pointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AzimuthChangedKeys)
				if err := _Azimuth.contract.UnpackLog(event, "ChangedKeys", log); err != nil {
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

// ParseChangedKeys is a log parse operation binding the contract event 0xaa10e7a0117d4323f1d99d630ec169bebb3a988e895770e351987e01ff5423d5.
//
// Solidity: event ChangedKeys(uint32 indexed point, bytes32 encryptionKey, bytes32 authenticationKey, uint32 cryptoSuiteVersion, uint32 keyRevisionNumber)
func (_Azimuth *AzimuthFilterer) ParseChangedKeys(log types.Log) (*AzimuthChangedKeys, error) {
	event := new(AzimuthChangedKeys)
	if err := _Azimuth.contract.UnpackLog(event, "ChangedKeys", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AzimuthChangedManagementProxyIterator is returned from FilterChangedManagementProxy and is used to iterate over the raw logs and unpacked data for ChangedManagementProxy events raised by the Azimuth contract.
type AzimuthChangedManagementProxyIterator struct {
	Event *AzimuthChangedManagementProxy // Event containing the contract specifics and raw log

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
func (it *AzimuthChangedManagementProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AzimuthChangedManagementProxy)
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
		it.Event = new(AzimuthChangedManagementProxy)
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
func (it *AzimuthChangedManagementProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AzimuthChangedManagementProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AzimuthChangedManagementProxy represents a ChangedManagementProxy event raised by the Azimuth contract.
type AzimuthChangedManagementProxy struct {
	Point           uint32
	ManagementProxy common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterChangedManagementProxy is a free log retrieval operation binding the contract event 0xab9c9327cffd2acc168fafedbe06139f5f55cb84c761df05e0511c251e2ee9bf.
//
// Solidity: event ChangedManagementProxy(uint32 indexed point, address indexed managementProxy)
func (_Azimuth *AzimuthFilterer) FilterChangedManagementProxy(opts *bind.FilterOpts, point []uint32, managementProxy []common.Address) (*AzimuthChangedManagementProxyIterator, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var managementProxyRule []interface{}
	for _, managementProxyItem := range managementProxy {
		managementProxyRule = append(managementProxyRule, managementProxyItem)
	}

	logs, sub, err := _Azimuth.contract.FilterLogs(opts, "ChangedManagementProxy", pointRule, managementProxyRule)
	if err != nil {
		return nil, err
	}
	return &AzimuthChangedManagementProxyIterator{contract: _Azimuth.contract, event: "ChangedManagementProxy", logs: logs, sub: sub}, nil
}

// WatchChangedManagementProxy is a free log subscription operation binding the contract event 0xab9c9327cffd2acc168fafedbe06139f5f55cb84c761df05e0511c251e2ee9bf.
//
// Solidity: event ChangedManagementProxy(uint32 indexed point, address indexed managementProxy)
func (_Azimuth *AzimuthFilterer) WatchChangedManagementProxy(opts *bind.WatchOpts, sink chan<- *AzimuthChangedManagementProxy, point []uint32, managementProxy []common.Address) (event.Subscription, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var managementProxyRule []interface{}
	for _, managementProxyItem := range managementProxy {
		managementProxyRule = append(managementProxyRule, managementProxyItem)
	}

	logs, sub, err := _Azimuth.contract.WatchLogs(opts, "ChangedManagementProxy", pointRule, managementProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AzimuthChangedManagementProxy)
				if err := _Azimuth.contract.UnpackLog(event, "ChangedManagementProxy", log); err != nil {
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

// ParseChangedManagementProxy is a log parse operation binding the contract event 0xab9c9327cffd2acc168fafedbe06139f5f55cb84c761df05e0511c251e2ee9bf.
//
// Solidity: event ChangedManagementProxy(uint32 indexed point, address indexed managementProxy)
func (_Azimuth *AzimuthFilterer) ParseChangedManagementProxy(log types.Log) (*AzimuthChangedManagementProxy, error) {
	event := new(AzimuthChangedManagementProxy)
	if err := _Azimuth.contract.UnpackLog(event, "ChangedManagementProxy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AzimuthChangedSpawnProxyIterator is returned from FilterChangedSpawnProxy and is used to iterate over the raw logs and unpacked data for ChangedSpawnProxy events raised by the Azimuth contract.
type AzimuthChangedSpawnProxyIterator struct {
	Event *AzimuthChangedSpawnProxy // Event containing the contract specifics and raw log

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
func (it *AzimuthChangedSpawnProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AzimuthChangedSpawnProxy)
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
		it.Event = new(AzimuthChangedSpawnProxy)
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
func (it *AzimuthChangedSpawnProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AzimuthChangedSpawnProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AzimuthChangedSpawnProxy represents a ChangedSpawnProxy event raised by the Azimuth contract.
type AzimuthChangedSpawnProxy struct {
	Point      uint32
	SpawnProxy common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChangedSpawnProxy is a free log retrieval operation binding the contract event 0x902736af7b3cefe10d9e840aed0d687e35c84095122b25051a20ead8866f006d.
//
// Solidity: event ChangedSpawnProxy(uint32 indexed point, address indexed spawnProxy)
func (_Azimuth *AzimuthFilterer) FilterChangedSpawnProxy(opts *bind.FilterOpts, point []uint32, spawnProxy []common.Address) (*AzimuthChangedSpawnProxyIterator, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var spawnProxyRule []interface{}
	for _, spawnProxyItem := range spawnProxy {
		spawnProxyRule = append(spawnProxyRule, spawnProxyItem)
	}

	logs, sub, err := _Azimuth.contract.FilterLogs(opts, "ChangedSpawnProxy", pointRule, spawnProxyRule)
	if err != nil {
		return nil, err
	}
	return &AzimuthChangedSpawnProxyIterator{contract: _Azimuth.contract, event: "ChangedSpawnProxy", logs: logs, sub: sub}, nil
}

// WatchChangedSpawnProxy is a free log subscription operation binding the contract event 0x902736af7b3cefe10d9e840aed0d687e35c84095122b25051a20ead8866f006d.
//
// Solidity: event ChangedSpawnProxy(uint32 indexed point, address indexed spawnProxy)
func (_Azimuth *AzimuthFilterer) WatchChangedSpawnProxy(opts *bind.WatchOpts, sink chan<- *AzimuthChangedSpawnProxy, point []uint32, spawnProxy []common.Address) (event.Subscription, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var spawnProxyRule []interface{}
	for _, spawnProxyItem := range spawnProxy {
		spawnProxyRule = append(spawnProxyRule, spawnProxyItem)
	}

	logs, sub, err := _Azimuth.contract.WatchLogs(opts, "ChangedSpawnProxy", pointRule, spawnProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AzimuthChangedSpawnProxy)
				if err := _Azimuth.contract.UnpackLog(event, "ChangedSpawnProxy", log); err != nil {
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

// ParseChangedSpawnProxy is a log parse operation binding the contract event 0x902736af7b3cefe10d9e840aed0d687e35c84095122b25051a20ead8866f006d.
//
// Solidity: event ChangedSpawnProxy(uint32 indexed point, address indexed spawnProxy)
func (_Azimuth *AzimuthFilterer) ParseChangedSpawnProxy(log types.Log) (*AzimuthChangedSpawnProxy, error) {
	event := new(AzimuthChangedSpawnProxy)
	if err := _Azimuth.contract.UnpackLog(event, "ChangedSpawnProxy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AzimuthChangedTransferProxyIterator is returned from FilterChangedTransferProxy and is used to iterate over the raw logs and unpacked data for ChangedTransferProxy events raised by the Azimuth contract.
type AzimuthChangedTransferProxyIterator struct {
	Event *AzimuthChangedTransferProxy // Event containing the contract specifics and raw log

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
func (it *AzimuthChangedTransferProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AzimuthChangedTransferProxy)
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
		it.Event = new(AzimuthChangedTransferProxy)
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
func (it *AzimuthChangedTransferProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AzimuthChangedTransferProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AzimuthChangedTransferProxy represents a ChangedTransferProxy event raised by the Azimuth contract.
type AzimuthChangedTransferProxy struct {
	Point         uint32
	TransferProxy common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChangedTransferProxy is a free log retrieval operation binding the contract event 0xcfe369b7197e7f0cf06793ae2472a9b13583fecbed2f78dfa14d1f10796b847c.
//
// Solidity: event ChangedTransferProxy(uint32 indexed point, address indexed transferProxy)
func (_Azimuth *AzimuthFilterer) FilterChangedTransferProxy(opts *bind.FilterOpts, point []uint32, transferProxy []common.Address) (*AzimuthChangedTransferProxyIterator, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var transferProxyRule []interface{}
	for _, transferProxyItem := range transferProxy {
		transferProxyRule = append(transferProxyRule, transferProxyItem)
	}

	logs, sub, err := _Azimuth.contract.FilterLogs(opts, "ChangedTransferProxy", pointRule, transferProxyRule)
	if err != nil {
		return nil, err
	}
	return &AzimuthChangedTransferProxyIterator{contract: _Azimuth.contract, event: "ChangedTransferProxy", logs: logs, sub: sub}, nil
}

// WatchChangedTransferProxy is a free log subscription operation binding the contract event 0xcfe369b7197e7f0cf06793ae2472a9b13583fecbed2f78dfa14d1f10796b847c.
//
// Solidity: event ChangedTransferProxy(uint32 indexed point, address indexed transferProxy)
func (_Azimuth *AzimuthFilterer) WatchChangedTransferProxy(opts *bind.WatchOpts, sink chan<- *AzimuthChangedTransferProxy, point []uint32, transferProxy []common.Address) (event.Subscription, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var transferProxyRule []interface{}
	for _, transferProxyItem := range transferProxy {
		transferProxyRule = append(transferProxyRule, transferProxyItem)
	}

	logs, sub, err := _Azimuth.contract.WatchLogs(opts, "ChangedTransferProxy", pointRule, transferProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AzimuthChangedTransferProxy)
				if err := _Azimuth.contract.UnpackLog(event, "ChangedTransferProxy", log); err != nil {
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

// ParseChangedTransferProxy is a log parse operation binding the contract event 0xcfe369b7197e7f0cf06793ae2472a9b13583fecbed2f78dfa14d1f10796b847c.
//
// Solidity: event ChangedTransferProxy(uint32 indexed point, address indexed transferProxy)
func (_Azimuth *AzimuthFilterer) ParseChangedTransferProxy(log types.Log) (*AzimuthChangedTransferProxy, error) {
	event := new(AzimuthChangedTransferProxy)
	if err := _Azimuth.contract.UnpackLog(event, "ChangedTransferProxy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AzimuthChangedVotingProxyIterator is returned from FilterChangedVotingProxy and is used to iterate over the raw logs and unpacked data for ChangedVotingProxy events raised by the Azimuth contract.
type AzimuthChangedVotingProxyIterator struct {
	Event *AzimuthChangedVotingProxy // Event containing the contract specifics and raw log

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
func (it *AzimuthChangedVotingProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AzimuthChangedVotingProxy)
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
		it.Event = new(AzimuthChangedVotingProxy)
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
func (it *AzimuthChangedVotingProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AzimuthChangedVotingProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AzimuthChangedVotingProxy represents a ChangedVotingProxy event raised by the Azimuth contract.
type AzimuthChangedVotingProxy struct {
	Point       uint32
	VotingProxy common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChangedVotingProxy is a free log retrieval operation binding the contract event 0xcbd6269ec71457f2c7b1a22774f246f6c5a2eae3795ed7300db517680c61c805.
//
// Solidity: event ChangedVotingProxy(uint32 indexed point, address indexed votingProxy)
func (_Azimuth *AzimuthFilterer) FilterChangedVotingProxy(opts *bind.FilterOpts, point []uint32, votingProxy []common.Address) (*AzimuthChangedVotingProxyIterator, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var votingProxyRule []interface{}
	for _, votingProxyItem := range votingProxy {
		votingProxyRule = append(votingProxyRule, votingProxyItem)
	}

	logs, sub, err := _Azimuth.contract.FilterLogs(opts, "ChangedVotingProxy", pointRule, votingProxyRule)
	if err != nil {
		return nil, err
	}
	return &AzimuthChangedVotingProxyIterator{contract: _Azimuth.contract, event: "ChangedVotingProxy", logs: logs, sub: sub}, nil
}

// WatchChangedVotingProxy is a free log subscription operation binding the contract event 0xcbd6269ec71457f2c7b1a22774f246f6c5a2eae3795ed7300db517680c61c805.
//
// Solidity: event ChangedVotingProxy(uint32 indexed point, address indexed votingProxy)
func (_Azimuth *AzimuthFilterer) WatchChangedVotingProxy(opts *bind.WatchOpts, sink chan<- *AzimuthChangedVotingProxy, point []uint32, votingProxy []common.Address) (event.Subscription, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var votingProxyRule []interface{}
	for _, votingProxyItem := range votingProxy {
		votingProxyRule = append(votingProxyRule, votingProxyItem)
	}

	logs, sub, err := _Azimuth.contract.WatchLogs(opts, "ChangedVotingProxy", pointRule, votingProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AzimuthChangedVotingProxy)
				if err := _Azimuth.contract.UnpackLog(event, "ChangedVotingProxy", log); err != nil {
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

// ParseChangedVotingProxy is a log parse operation binding the contract event 0xcbd6269ec71457f2c7b1a22774f246f6c5a2eae3795ed7300db517680c61c805.
//
// Solidity: event ChangedVotingProxy(uint32 indexed point, address indexed votingProxy)
func (_Azimuth *AzimuthFilterer) ParseChangedVotingProxy(log types.Log) (*AzimuthChangedVotingProxy, error) {
	event := new(AzimuthChangedVotingProxy)
	if err := _Azimuth.contract.UnpackLog(event, "ChangedVotingProxy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AzimuthEscapeAcceptedIterator is returned from FilterEscapeAccepted and is used to iterate over the raw logs and unpacked data for EscapeAccepted events raised by the Azimuth contract.
type AzimuthEscapeAcceptedIterator struct {
	Event *AzimuthEscapeAccepted // Event containing the contract specifics and raw log

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
func (it *AzimuthEscapeAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AzimuthEscapeAccepted)
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
		it.Event = new(AzimuthEscapeAccepted)
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
func (it *AzimuthEscapeAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AzimuthEscapeAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AzimuthEscapeAccepted represents a EscapeAccepted event raised by the Azimuth contract.
type AzimuthEscapeAccepted struct {
	Point   uint32
	Sponsor uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEscapeAccepted is a free log retrieval operation binding the contract event 0x7e447c9b1bda4b174b0796e100bf7f34ebf36dbb7fe665490b1bfce6246a9da5.
//
// Solidity: event EscapeAccepted(uint32 indexed point, uint32 indexed sponsor)
func (_Azimuth *AzimuthFilterer) FilterEscapeAccepted(opts *bind.FilterOpts, point []uint32, sponsor []uint32) (*AzimuthEscapeAcceptedIterator, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _Azimuth.contract.FilterLogs(opts, "EscapeAccepted", pointRule, sponsorRule)
	if err != nil {
		return nil, err
	}
	return &AzimuthEscapeAcceptedIterator{contract: _Azimuth.contract, event: "EscapeAccepted", logs: logs, sub: sub}, nil
}

// WatchEscapeAccepted is a free log subscription operation binding the contract event 0x7e447c9b1bda4b174b0796e100bf7f34ebf36dbb7fe665490b1bfce6246a9da5.
//
// Solidity: event EscapeAccepted(uint32 indexed point, uint32 indexed sponsor)
func (_Azimuth *AzimuthFilterer) WatchEscapeAccepted(opts *bind.WatchOpts, sink chan<- *AzimuthEscapeAccepted, point []uint32, sponsor []uint32) (event.Subscription, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _Azimuth.contract.WatchLogs(opts, "EscapeAccepted", pointRule, sponsorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AzimuthEscapeAccepted)
				if err := _Azimuth.contract.UnpackLog(event, "EscapeAccepted", log); err != nil {
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

// ParseEscapeAccepted is a log parse operation binding the contract event 0x7e447c9b1bda4b174b0796e100bf7f34ebf36dbb7fe665490b1bfce6246a9da5.
//
// Solidity: event EscapeAccepted(uint32 indexed point, uint32 indexed sponsor)
func (_Azimuth *AzimuthFilterer) ParseEscapeAccepted(log types.Log) (*AzimuthEscapeAccepted, error) {
	event := new(AzimuthEscapeAccepted)
	if err := _Azimuth.contract.UnpackLog(event, "EscapeAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AzimuthEscapeCanceledIterator is returned from FilterEscapeCanceled and is used to iterate over the raw logs and unpacked data for EscapeCanceled events raised by the Azimuth contract.
type AzimuthEscapeCanceledIterator struct {
	Event *AzimuthEscapeCanceled // Event containing the contract specifics and raw log

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
func (it *AzimuthEscapeCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AzimuthEscapeCanceled)
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
		it.Event = new(AzimuthEscapeCanceled)
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
func (it *AzimuthEscapeCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AzimuthEscapeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AzimuthEscapeCanceled represents a EscapeCanceled event raised by the Azimuth contract.
type AzimuthEscapeCanceled struct {
	Point   uint32
	Sponsor uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEscapeCanceled is a free log retrieval operation binding the contract event 0xd653bb0e0bb7ce8393e624d98fbf17cda5902c8328ed0cd09988f36890d9932a.
//
// Solidity: event EscapeCanceled(uint32 indexed point, uint32 indexed sponsor)
func (_Azimuth *AzimuthFilterer) FilterEscapeCanceled(opts *bind.FilterOpts, point []uint32, sponsor []uint32) (*AzimuthEscapeCanceledIterator, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _Azimuth.contract.FilterLogs(opts, "EscapeCanceled", pointRule, sponsorRule)
	if err != nil {
		return nil, err
	}
	return &AzimuthEscapeCanceledIterator{contract: _Azimuth.contract, event: "EscapeCanceled", logs: logs, sub: sub}, nil
}

// WatchEscapeCanceled is a free log subscription operation binding the contract event 0xd653bb0e0bb7ce8393e624d98fbf17cda5902c8328ed0cd09988f36890d9932a.
//
// Solidity: event EscapeCanceled(uint32 indexed point, uint32 indexed sponsor)
func (_Azimuth *AzimuthFilterer) WatchEscapeCanceled(opts *bind.WatchOpts, sink chan<- *AzimuthEscapeCanceled, point []uint32, sponsor []uint32) (event.Subscription, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _Azimuth.contract.WatchLogs(opts, "EscapeCanceled", pointRule, sponsorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AzimuthEscapeCanceled)
				if err := _Azimuth.contract.UnpackLog(event, "EscapeCanceled", log); err != nil {
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

// ParseEscapeCanceled is a log parse operation binding the contract event 0xd653bb0e0bb7ce8393e624d98fbf17cda5902c8328ed0cd09988f36890d9932a.
//
// Solidity: event EscapeCanceled(uint32 indexed point, uint32 indexed sponsor)
func (_Azimuth *AzimuthFilterer) ParseEscapeCanceled(log types.Log) (*AzimuthEscapeCanceled, error) {
	event := new(AzimuthEscapeCanceled)
	if err := _Azimuth.contract.UnpackLog(event, "EscapeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AzimuthEscapeRequestedIterator is returned from FilterEscapeRequested and is used to iterate over the raw logs and unpacked data for EscapeRequested events raised by the Azimuth contract.
type AzimuthEscapeRequestedIterator struct {
	Event *AzimuthEscapeRequested // Event containing the contract specifics and raw log

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
func (it *AzimuthEscapeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AzimuthEscapeRequested)
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
		it.Event = new(AzimuthEscapeRequested)
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
func (it *AzimuthEscapeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AzimuthEscapeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AzimuthEscapeRequested represents a EscapeRequested event raised by the Azimuth contract.
type AzimuthEscapeRequested struct {
	Point   uint32
	Sponsor uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEscapeRequested is a free log retrieval operation binding the contract event 0xb4d4850b8f218218141c5665cba379e53e9bb015b51e8d934be70210aead874a.
//
// Solidity: event EscapeRequested(uint32 indexed point, uint32 indexed sponsor)
func (_Azimuth *AzimuthFilterer) FilterEscapeRequested(opts *bind.FilterOpts, point []uint32, sponsor []uint32) (*AzimuthEscapeRequestedIterator, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _Azimuth.contract.FilterLogs(opts, "EscapeRequested", pointRule, sponsorRule)
	if err != nil {
		return nil, err
	}
	return &AzimuthEscapeRequestedIterator{contract: _Azimuth.contract, event: "EscapeRequested", logs: logs, sub: sub}, nil
}

// WatchEscapeRequested is a free log subscription operation binding the contract event 0xb4d4850b8f218218141c5665cba379e53e9bb015b51e8d934be70210aead874a.
//
// Solidity: event EscapeRequested(uint32 indexed point, uint32 indexed sponsor)
func (_Azimuth *AzimuthFilterer) WatchEscapeRequested(opts *bind.WatchOpts, sink chan<- *AzimuthEscapeRequested, point []uint32, sponsor []uint32) (event.Subscription, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _Azimuth.contract.WatchLogs(opts, "EscapeRequested", pointRule, sponsorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AzimuthEscapeRequested)
				if err := _Azimuth.contract.UnpackLog(event, "EscapeRequested", log); err != nil {
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

// ParseEscapeRequested is a log parse operation binding the contract event 0xb4d4850b8f218218141c5665cba379e53e9bb015b51e8d934be70210aead874a.
//
// Solidity: event EscapeRequested(uint32 indexed point, uint32 indexed sponsor)
func (_Azimuth *AzimuthFilterer) ParseEscapeRequested(log types.Log) (*AzimuthEscapeRequested, error) {
	event := new(AzimuthEscapeRequested)
	if err := _Azimuth.contract.UnpackLog(event, "EscapeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AzimuthLostSponsorIterator is returned from FilterLostSponsor and is used to iterate over the raw logs and unpacked data for LostSponsor events raised by the Azimuth contract.
type AzimuthLostSponsorIterator struct {
	Event *AzimuthLostSponsor // Event containing the contract specifics and raw log

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
func (it *AzimuthLostSponsorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AzimuthLostSponsor)
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
		it.Event = new(AzimuthLostSponsor)
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
func (it *AzimuthLostSponsorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AzimuthLostSponsorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AzimuthLostSponsor represents a LostSponsor event raised by the Azimuth contract.
type AzimuthLostSponsor struct {
	Point   uint32
	Sponsor uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLostSponsor is a free log retrieval operation binding the contract event 0xd7704f9a25193dbd0b0cb4a809feffffa7f19d1aae8817a71346c194448210d5.
//
// Solidity: event LostSponsor(uint32 indexed point, uint32 indexed sponsor)
func (_Azimuth *AzimuthFilterer) FilterLostSponsor(opts *bind.FilterOpts, point []uint32, sponsor []uint32) (*AzimuthLostSponsorIterator, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _Azimuth.contract.FilterLogs(opts, "LostSponsor", pointRule, sponsorRule)
	if err != nil {
		return nil, err
	}
	return &AzimuthLostSponsorIterator{contract: _Azimuth.contract, event: "LostSponsor", logs: logs, sub: sub}, nil
}

// WatchLostSponsor is a free log subscription operation binding the contract event 0xd7704f9a25193dbd0b0cb4a809feffffa7f19d1aae8817a71346c194448210d5.
//
// Solidity: event LostSponsor(uint32 indexed point, uint32 indexed sponsor)
func (_Azimuth *AzimuthFilterer) WatchLostSponsor(opts *bind.WatchOpts, sink chan<- *AzimuthLostSponsor, point []uint32, sponsor []uint32) (event.Subscription, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _Azimuth.contract.WatchLogs(opts, "LostSponsor", pointRule, sponsorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AzimuthLostSponsor)
				if err := _Azimuth.contract.UnpackLog(event, "LostSponsor", log); err != nil {
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

// ParseLostSponsor is a log parse operation binding the contract event 0xd7704f9a25193dbd0b0cb4a809feffffa7f19d1aae8817a71346c194448210d5.
//
// Solidity: event LostSponsor(uint32 indexed point, uint32 indexed sponsor)
func (_Azimuth *AzimuthFilterer) ParseLostSponsor(log types.Log) (*AzimuthLostSponsor, error) {
	event := new(AzimuthLostSponsor)
	if err := _Azimuth.contract.UnpackLog(event, "LostSponsor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AzimuthOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Azimuth contract.
type AzimuthOwnerChangedIterator struct {
	Event *AzimuthOwnerChanged // Event containing the contract specifics and raw log

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
func (it *AzimuthOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AzimuthOwnerChanged)
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
		it.Event = new(AzimuthOwnerChanged)
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
func (it *AzimuthOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AzimuthOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AzimuthOwnerChanged represents a OwnerChanged event raised by the Azimuth contract.
type AzimuthOwnerChanged struct {
	Point uint32
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0x16d0f539d49c6cad822b767a9445bfb1cf7ea6f2a6c2b120a7ea4cc7660d8fda.
//
// Solidity: event OwnerChanged(uint32 indexed point, address indexed owner)
func (_Azimuth *AzimuthFilterer) FilterOwnerChanged(opts *bind.FilterOpts, point []uint32, owner []common.Address) (*AzimuthOwnerChangedIterator, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Azimuth.contract.FilterLogs(opts, "OwnerChanged", pointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &AzimuthOwnerChangedIterator{contract: _Azimuth.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0x16d0f539d49c6cad822b767a9445bfb1cf7ea6f2a6c2b120a7ea4cc7660d8fda.
//
// Solidity: event OwnerChanged(uint32 indexed point, address indexed owner)
func (_Azimuth *AzimuthFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *AzimuthOwnerChanged, point []uint32, owner []common.Address) (event.Subscription, error) {

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Azimuth.contract.WatchLogs(opts, "OwnerChanged", pointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AzimuthOwnerChanged)
				if err := _Azimuth.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0x16d0f539d49c6cad822b767a9445bfb1cf7ea6f2a6c2b120a7ea4cc7660d8fda.
//
// Solidity: event OwnerChanged(uint32 indexed point, address indexed owner)
func (_Azimuth *AzimuthFilterer) ParseOwnerChanged(log types.Log) (*AzimuthOwnerChanged, error) {
	event := new(AzimuthOwnerChanged)
	if err := _Azimuth.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AzimuthOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Azimuth contract.
type AzimuthOwnershipRenouncedIterator struct {
	Event *AzimuthOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *AzimuthOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AzimuthOwnershipRenounced)
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
		it.Event = new(AzimuthOwnershipRenounced)
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
func (it *AzimuthOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AzimuthOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AzimuthOwnershipRenounced represents a OwnershipRenounced event raised by the Azimuth contract.
type AzimuthOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Azimuth *AzimuthFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*AzimuthOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Azimuth.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AzimuthOwnershipRenouncedIterator{contract: _Azimuth.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Azimuth *AzimuthFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *AzimuthOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Azimuth.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AzimuthOwnershipRenounced)
				if err := _Azimuth.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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
func (_Azimuth *AzimuthFilterer) ParseOwnershipRenounced(log types.Log) (*AzimuthOwnershipRenounced, error) {
	event := new(AzimuthOwnershipRenounced)
	if err := _Azimuth.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AzimuthOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Azimuth contract.
type AzimuthOwnershipTransferredIterator struct {
	Event *AzimuthOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AzimuthOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AzimuthOwnershipTransferred)
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
		it.Event = new(AzimuthOwnershipTransferred)
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
func (it *AzimuthOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AzimuthOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AzimuthOwnershipTransferred represents a OwnershipTransferred event raised by the Azimuth contract.
type AzimuthOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Azimuth *AzimuthFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AzimuthOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Azimuth.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AzimuthOwnershipTransferredIterator{contract: _Azimuth.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Azimuth *AzimuthFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AzimuthOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Azimuth.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AzimuthOwnershipTransferred)
				if err := _Azimuth.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Azimuth *AzimuthFilterer) ParseOwnershipTransferred(log types.Log) (*AzimuthOwnershipTransferred, error) {
	event := new(AzimuthOwnershipTransferred)
	if err := _Azimuth.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AzimuthSpawnedIterator is returned from FilterSpawned and is used to iterate over the raw logs and unpacked data for Spawned events raised by the Azimuth contract.
type AzimuthSpawnedIterator struct {
	Event *AzimuthSpawned // Event containing the contract specifics and raw log

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
func (it *AzimuthSpawnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AzimuthSpawned)
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
		it.Event = new(AzimuthSpawned)
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
func (it *AzimuthSpawnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AzimuthSpawnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AzimuthSpawned represents a Spawned event raised by the Azimuth contract.
type AzimuthSpawned struct {
	Prefix uint32
	Child  uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSpawned is a free log retrieval operation binding the contract event 0xb2d3a6e7a339f5c8ff96265e2f03a010a8541070f3744a247090964415081546.
//
// Solidity: event Spawned(uint32 indexed prefix, uint32 indexed child)
func (_Azimuth *AzimuthFilterer) FilterSpawned(opts *bind.FilterOpts, prefix []uint32, child []uint32) (*AzimuthSpawnedIterator, error) {

	var prefixRule []interface{}
	for _, prefixItem := range prefix {
		prefixRule = append(prefixRule, prefixItem)
	}
	var childRule []interface{}
	for _, childItem := range child {
		childRule = append(childRule, childItem)
	}

	logs, sub, err := _Azimuth.contract.FilterLogs(opts, "Spawned", prefixRule, childRule)
	if err != nil {
		return nil, err
	}
	return &AzimuthSpawnedIterator{contract: _Azimuth.contract, event: "Spawned", logs: logs, sub: sub}, nil
}

// WatchSpawned is a free log subscription operation binding the contract event 0xb2d3a6e7a339f5c8ff96265e2f03a010a8541070f3744a247090964415081546.
//
// Solidity: event Spawned(uint32 indexed prefix, uint32 indexed child)
func (_Azimuth *AzimuthFilterer) WatchSpawned(opts *bind.WatchOpts, sink chan<- *AzimuthSpawned, prefix []uint32, child []uint32) (event.Subscription, error) {

	var prefixRule []interface{}
	for _, prefixItem := range prefix {
		prefixRule = append(prefixRule, prefixItem)
	}
	var childRule []interface{}
	for _, childItem := range child {
		childRule = append(childRule, childItem)
	}

	logs, sub, err := _Azimuth.contract.WatchLogs(opts, "Spawned", prefixRule, childRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AzimuthSpawned)
				if err := _Azimuth.contract.UnpackLog(event, "Spawned", log); err != nil {
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

// ParseSpawned is a log parse operation binding the contract event 0xb2d3a6e7a339f5c8ff96265e2f03a010a8541070f3744a247090964415081546.
//
// Solidity: event Spawned(uint32 indexed prefix, uint32 indexed child)
func (_Azimuth *AzimuthFilterer) ParseSpawned(log types.Log) (*AzimuthSpawned, error) {
	event := new(AzimuthSpawned)
	if err := _Azimuth.contract.UnpackLog(event, "Spawned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
