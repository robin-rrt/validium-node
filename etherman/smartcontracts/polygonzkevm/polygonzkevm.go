// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonzkevm

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

// PolygonZkEVMBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonZkEVMBatchData struct {
	BatchHash          [32]byte
	GlobalExitRoot     [32]byte
	Timestamp          uint64
	MinForcedTimestamp uint64
}

// PolygonZkEVMDAData is an auto generated low-level Go binding around an user-defined struct.
type PolygonZkEVMDAData struct {
	BlockNumber uint32
	Proof       [][32]byte
	Width       *big.Int
	Index       *big.Int
}

// PolygonZkEVMForcedBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonZkEVMForcedBatchData struct {
	Transactions       []byte
	GlobalExitRoot     [32]byte
	MinForcedTimestamp uint64
}

// PolygonZkEVMInitializePackedParameters is an auto generated low-level Go binding around an user-defined struct.
type PolygonZkEVMInitializePackedParameters struct {
	Admin                    common.Address
	TrustedSequencer         common.Address
	PendingStateTimeout      uint64
	TrustedAggregator        common.Address
	TrustedAggregatorTimeout uint64
}

// PolygonzkevmMetaData contains all meta data concerning the Polygonzkevm contract.
var PolygonzkevmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIDataAvailabilityRouter\",\"name\":\"_dataAvailabilityRouter\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_matic\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"_rollupVerifier\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_forkID\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyPosted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DataRootNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDAProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ActivateForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"ConsolidatePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"OverridePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storedStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"provedStateRoot\",\"type\":\"bytes32\"}],\"name\":\"ProveNonDeterministicPendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"SetMultiplierBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"SetPendingStateTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"SetTrustedAggregatorTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"SetVerifyBatchTimeTarget\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"UpdateZkEVMVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequencedBatchNum\",\"type\":\"uint64\"}],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"batchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newStateRoot\",\"type\":\"uint256\"}],\"name\":\"checkStateRootInsidePrime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"consolidatePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maticAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forkID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trustedSequencer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateTimeout\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"trustedAggregatorTimeout\",\"type\":\"uint64\"}],\"internalType\":\"structPolygonZkEVM.InitializePackedParameters\",\"name\":\"initializePackedParameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"genesisRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_trustedSequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isBatchPosted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isForcedBatchDisallowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"isPendingStateConsolidable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPendingState\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPendingStateConsolidated\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matic\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplierBatchFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"overridePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingStateTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingStateTransitions\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"proveNonDeterministicPendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupVerifier\",\"outputs\":[{\"internalType\":\"contractIVerifierRollup\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractIDataAvailabilityRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPolygonZkEVM.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"width\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structPolygonZkEVM.DAData[]\",\"name\":\"daData\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPolygonZkEVM.ForcedBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"sequencedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"setMultiplierBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"setPendingStateTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDataAvailabilityRouter\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"setTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"setTrustedAggregatorTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"setVerifyBatchTimeTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregatorTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyBatchTimeTarget\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x61014060405234801562000011575f80fd5b5060405162005dd238038062005dd28339810160408190526200003491620000b6565b606680546001600160a01b0319166001600160a01b0397881617905595851660c05292841660805290831660a05290911660e0526001600160401b0390811661010052166101205262000154565b6001600160a01b038116811462000097575f80fd5b50565b80516001600160401b0381168114620000b1575f80fd5b919050565b5f805f805f805f60e0888a031215620000cd575f80fd5b8751620000da8162000082565b6020890151909750620000ed8162000082565b6040890151909650620001008162000082565b6060890151909550620001138162000082565b6080890151909450620001268162000082565b92506200013660a089016200009a565b91506200014660c089016200009a565b905092959891949750929550565b60805160a05160c05160e0516101005161012051615bb86200021a5f395f818161069701528181610ded015261285501525f81816107f60152610dc301525f81816107bc01528181612ebb01528181613a40015261497801525f818161096701528181610f5201528181611103015281816118e701528181613046015281816136a6015261455901525f8181610a1401528181613edc015261431001525f81816108b101528181611d710152818161301b01528181613a0e0152613fbb0152615bb85ff3fe608060405234801561000f575f80fd5b50600436106103d9575f3560e01c80638c3d73011161020f578063c89e42df1161012e578063e8bf92ed116100c3578063f14916d611610093578063f851a44011610079578063f851a44014610aaf578063f887ea4014610ac2578063f8b823e414610ad5575f80fd5b8063f14916d614610a89578063f2fde38b14610a9c575f80fd5b8063e8bf92ed14610a0f578063eaeb077b14610a36578063ed6b010414610a49578063efde529014610a76575f80fd5b8063d8d1091b116100fe578063d8d1091b1461099c578063d939b315146109af578063dbc16976146109d7578063e7a7ed02146109df575f80fd5b8063c89e42df1461092d578063cfa8ed4714610940578063d02103ca14610962578063d2e129f914610989575f80fd5b8063adc879e9116101a4578063ba58ae3911610174578063ba58ae39146108d3578063c0d78655146108e6578063c0ed84e0146108f9578063c754c7ed14610901575f80fd5b8063adc879e9146107f1578063afd23cbe14610818578063b4d63f5814610846578063b6b0b097146108ac575f80fd5b80639c9f3dfe116101df5780639c9f3dfe14610791578063a066215c146107a4578063a3c573eb146107b7578063ada8f919146107de575f80fd5b80638c3d73011461075d5780638da5cb5b1461076557806399f5634e146107765780639aa972a31461077e575f80fd5b80634a1a89a7116102fb5780636b8616ce116102905780637fcb3653116102605780637fcb36531461067e578063831c7ead14610692578063837a4738146106b9578063841b24d71461072d575f80fd5b80636b8616ce146106315780636ff512cc14610650578063715018a6146106635780637215541a1461066b575f80fd5b8063542028d5116102cb578063542028d5146106065780635ec919581461060e5780636046916914610616578063621dd4111461061e575f80fd5b80634a1a89a7146105a15780634a910e6a146105c15780634e487706146105d45780635392c5e0146105e7575f80fd5b80632b0006fa1161037157806340b10f1c1161034157806340b10f1c14610518578063423fa856146105455780634560526714610565578063458c04771461058d575f80fd5b80632b0006fa146104cc5780632c1f816a146104df578063383b3be8146104f2578063394218e914610505575f80fd5b806319d8ac61116103ac57806319d8ac611461045b578063220d78991461046f578063267822471461048257806329878983146104ad575f80fd5b80630a0d9fbe146103dd578063107bf28c1461041457806315064c96146104295780631816b7e514610446575b5f80fd5b6071546103f690610100900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020015b60405180910390f35b61041c610ade565b60405161040b9190614ff2565b6071546104369060ff1681565b604051901515815260200161040b565b61045961045436600461500b565b610b6a565b005b6075546103f69067ffffffffffffffff1681565b61041c61047d366004615043565b610c75565b607d54610495906001600160a01b031681565b6040516001600160a01b03909116815260200161040b565b607654610495906801000000000000000090046001600160a01b031681565b6104596104da3660046150a3565b610e4b565b6104596104ed366004615107565b610ffb565b61043661050036600461517c565b6111e9565b61045961051336600461517c565b61123e565b610537610526366004615195565b60656020525f908152604090205481565b60405190815260200161040b565b6075546103f69068010000000000000000900467ffffffffffffffff1681565b6075546103f690700100000000000000000000000000000000900467ffffffffffffffff1681565b607b546103f69067ffffffffffffffff1681565b607b546103f69068010000000000000000900467ffffffffffffffff1681565b6104596105cf36600461517c565b6113b5565b6104596105e236600461517c565b61145b565b6105376105f536600461517c565b60776020525f908152604090205481565b61041c6115d2565b6104596115df565b6105376116d1565b61045961062c3660046150a3565b6116e6565b61053761063f36600461517c565b60736020525f908152604090205481565b61045961065e3660046151cb565b611a57565b610459611b12565b61045961067936600461517c565b611b25565b6076546103f69067ffffffffffffffff1681565b6103f67f000000000000000000000000000000000000000000000000000000000000000081565b6107016106c7366004615195565b607a6020525f908152604090208054600182015460029092015467ffffffffffffffff808316936801000000000000000090930416919084565b6040805167ffffffffffffffff958616815294909316602085015291830152606082015260800161040b565b607b546103f6907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b610459611c85565b6033546001600160a01b0316610495565b610537611d37565b61045961078c366004615107565b611e7d565b61045961079f36600461517c565b611f2d565b6104596107b236600461517c565b61209c565b6104957f000000000000000000000000000000000000000000000000000000000000000081565b6104596107ec3660046151cb565b612195565b6103f67f000000000000000000000000000000000000000000000000000000000000000081565b607154610833906901000000000000000000900461ffff1681565b60405161ffff909116815260200161040b565b61088661085436600461517c565b60746020525f90815260409020805460019091015467ffffffffffffffff808216916801000000000000000090041683565b6040805193845267ffffffffffffffff928316602085015291169082015260600161040b565b6104957f000000000000000000000000000000000000000000000000000000000000000081565b6104366108e1366004615195565b61223f565b6104596108f43660046151cb565b6122c7565b6103f6612309565b607d546103f69074010000000000000000000000000000000000000000900467ffffffffffffffff1681565b61045961093b3660046152a4565b61235c565b607154610495906b01000000000000000000000090046001600160a01b031681565b6104957f000000000000000000000000000000000000000000000000000000000000000081565b61045961099736600461531b565b6123dc565b6104596109aa366004615407565b6128f8565b607b546103f690700100000000000000000000000000000000900467ffffffffffffffff1681565b610459612e75565b6075546103f6907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b6104957f000000000000000000000000000000000000000000000000000000000000000081565b610459610a44366004615446565b612f2f565b607d54610436907c0100000000000000000000000000000000000000000000000000000000900460ff1681565b610459610a8436600461548e565b6132ab565b610459610a973660046151cb565b613aee565b610459610aaa3660046151cb565b613ba6565b607c54610495906001600160a01b031681565b606654610495906001600160a01b031681565b61053760725481565b60798054610aeb90615534565b80601f0160208091040260200160405190810160405280929190818152602001828054610b1790615534565b8015610b625780601f10610b3957610100808354040283529160200191610b62565b820191905f5260205f20905b815481529060010190602001808311610b4557829003601f168201915b505050505081565b607c546001600160a01b03163314610bae576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88161ffff161080610bc757506103ff8161ffff16115b15610bfe576040517f4c2533c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607180547fffffffffffffffffffffffffffffffffffffffffff0000ffffffffffffffffff16690100000000000000000061ffff8416908102919091179091556040519081527f7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5906020015b60405180910390a150565b67ffffffffffffffff8086165f818152607460205260408082205493881682529020546060929115801590610ca8575081155b15610cdf576040517f6818c29e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80610d16576040517f66385b5100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d1f8461223f565b610d55576040517f176b913c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080517fffffffffffffffffffffffffffffffffffffffff0000000000000000000000003360601b166020820152603481019690965260548601929092527fffffffffffffffff00000000000000000000000000000000000000000000000060c098891b811660748701527f0000000000000000000000000000000000000000000000000000000000000000891b8116607c8701527f0000000000000000000000000000000000000000000000000000000000000000891b81166084870152608c86019490945260ac85015260cc840194909452509290931b90911660ec830152805180830360d401815260f4909201905290565b6076546801000000000000000090046001600160a01b03163314610e9b576040517fbbcbbc0500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ea9868686868686613c4d565b607680547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8681169182179092555f908152607760205260409020839055607b541615610f2357607b80547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018490527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906333d6247d906024015f604051808303815f87803b158015610f9b575f80fd5b505af1158015610fad573d5f803e3d5ffd5b505060405184815233925067ffffffffffffffff871691507fcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe906020015b60405180910390a3505050505050565b6076546801000000000000000090046001600160a01b0316331461104b576040517fbbcbbc0500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61105a87878787878787613fee565b607680547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8681169182179092555f908152607760205260409020839055607b5416156110d457607b80547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018490527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906333d6247d906024015f604051808303815f87803b15801561114c575f80fd5b505af115801561115e573d5f803e3d5ffd5b5050607b805477ffffffffffffffffffffffffffffffffffffffffffffffff167a093a800000000000000000000000000000000000000000000000001790555050604051828152339067ffffffffffffffff8616907fcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf729060200160405180910390a350505050505050565b607b5467ffffffffffffffff8281165f908152607a60205260408120549092429261122c92700100000000000000000000000000000000909204811691166155b2565b67ffffffffffffffff16111592915050565b607c546001600160a01b03163314611282576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff821611156112c9576040517f1d06e87900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60715460ff1661133857607b5467ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811690821610611338576040517f401636df00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527f1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a190602001610c6a565b6076546801000000000000000090046001600160a01b0316331461144f5760715460ff1615611410576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611419816111e9565b61144f576040517f0ce9e4a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61145881614410565b50565b607c546001600160a01b0316331461149f576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff821611156114e6576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60715460ff1661155157607d5467ffffffffffffffff74010000000000000000000000000000000000000000909104811690821610611551576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607d80547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b90602001610c6a565b60788054610aeb90615534565b607c546001600160a01b03163314611623576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607d547c0100000000000000000000000000000000000000000000000000000000900460ff1661167f576040517ff6ba91a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607d80547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690556040517f854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f905f90a1565b5f60725460646116e191906155da565b905090565b60715460ff1615611723576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b5467ffffffffffffffff8581165f90815260746020526040902060010154429261176f927801000000000000000000000000000000000000000000000000909104811691166155b2565b67ffffffffffffffff1611156117b1576040517f8a0704d300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e86117be86866155f1565b67ffffffffffffffff161115611800576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61180e868686868686613c4d565b61181784614610565b607b54700100000000000000000000000000000000900467ffffffffffffffff165f0361194b57607680547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8681169182179092555f908152607760205260409020839055607b5416156118b857607b80547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018490527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906333d6247d906024015f604051808303815f87803b158015611930575f80fd5b505af1158015611942573d5f803e3d5ffd5b50505050611a19565b6119536147ea565b607b805467ffffffffffffffff16905f61196c83615612565b825467ffffffffffffffff9182166101009390930a928302928202191691909117909155604080516080810182524283168152878316602080830191825282840189815260608401898152607b5487165f908152607a90935294909120925183549251861668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009093169516949094171781559151600183015551600290910155505b604051828152339067ffffffffffffffff8616907f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f596690602001610feb565b607c546001600160a01b03163314611a9b576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607180547fff0000000000000000000000000000000000000000ffffffffffffffffffffff166b0100000000000000000000006001600160a01b038416908102919091179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610c6a565b611b1a614899565b611b235f61490d565b565b6033546001600160a01b03163314611c7d575f611b40612309565b90508067ffffffffffffffff168267ffffffffffffffff1611611b8f576040517f812a372d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60755467ffffffffffffffff6801000000000000000090910481169083161180611bd4575067ffffffffffffffff8083165f9081526074602052604090206001015416155b15611c0b576040517f98c5c01400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8083165f908152607460205260409020600101544291611c399162093a8091166155b2565b67ffffffffffffffff161115611c7b576040517fd257555a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b611458614976565b607d546001600160a01b03163314611cc9576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607d54607c80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0390921691821790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9081906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015611db6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611dda9190615638565b90505f611de5612309565b60755467ffffffffffffffff680100000000000000008204811691611e3d91700100000000000000000000000000000000820481169178010000000000000000000000000000000000000000000000009004166155f1565b611e4791906155b2565b611e5191906155f1565b67ffffffffffffffff169050805f03611e6c575f9250505090565b611e76818361567c565b9250505090565b60715460ff1615611eba576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611ec987878787878787613fee565b67ffffffffffffffff84165f908152607760209081526040918290205482519081529081018490527f1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010910160405180910390a1611f24614976565b50505050505050565b607c546001600160a01b03163314611f71576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115611fb8576040517fcc96507000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60715460ff1661201f57607b5467ffffffffffffffff70010000000000000000000000000000000090910481169082161061201f576040517f48a05a9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b80547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c7590602001610c6a565b607c546001600160a01b031633146120e0576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620151808167ffffffffffffffff161115612127576040517fe067dfe800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607180547fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff1661010067ffffffffffffffff8416908102919091179091556040519081527f1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c2890602001610c6a565b607c546001600160a01b031633146121d9576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607d80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610c6a565b5f67ffffffff0000000167ffffffffffffffff8316108015612276575067ffffffff00000001604083901c67ffffffffffffffff16105b8015612297575067ffffffff00000001608083901c67ffffffffffffffff16105b80156122ae575067ffffffff0000000160c083901c105b156122bb57506001919050565b505f919050565b919050565b6122cf614899565b606680547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b607b545f9067ffffffffffffffff161561234b5750607b5467ffffffffffffffff9081165f908152607a60205260409020546801000000000000000090041690565b5060765467ffffffffffffffff1690565b607c546001600160a01b031633146123a0576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60786123ac82826156d4565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610c6a9190614ff2565b5f54610100900460ff16158080156123fa57505f54600160ff909116105b806124135750303b15801561241357505f5460ff166001145b6124a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015612500575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61250d60208801886151cb565b607c80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039290921691909117905561255560408801602089016151cb565b607180546001600160a01b03929092166b010000000000000000000000027fff0000000000000000000000000000000000000000ffffffffffffffffffffff9092169190911790556125ad60808801606089016151cb565b607680546001600160a01b039290921668010000000000000000027fffffffff0000000000000000000000000000000000000000ffffffffffffffff9092169190911790555f805260776020527f5bc4dfdf00d266d79007b974fdf03467b7a6429c65a950cd7d967e8c7af65da6869055607861262a86826156d4565b50607961263785826156d4565b5062093a8061264c6060890160408a0161517c565b67ffffffffffffffff16111561268e576040517fcc96507000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61269e606088016040890161517c565b607b805467ffffffffffffffff92909216700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff90921691909117905562093a8061270060a0890160808a0161517c565b67ffffffffffffffff161115612742576040517f1d06e87900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61275260a088016080890161517c565b607b805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff939093169290920291909117905567016345785d8a0000607255607180547fffffffffffffffffffffffffffffffffffffffffff00000000000000000000ff166a03ea000000000000070800179055607d80547fffffff000000000000000000ffffffffffffffffffffffffffffffffffffffff167c01000000000006978000000000000000000000000000000000000000001790556128316149ec565b7fed7be53c9f1a96a481223b15568a5b1a475e01a74b347d6ca187c8bf0c078cd65f7f0000000000000000000000000000000000000000000000000000000000000000858560405161288694939291906157f7565b60405180910390a18015611f24575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050505050565b607d547c0100000000000000000000000000000000000000000000000000000000900460ff1615612955576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60715460ff1615612992576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f8190036129cd576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115612a09576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60755467ffffffffffffffff78010000000000000000000000000000000000000000000000008204811691612a5491849170010000000000000000000000000000000090041661582e565b1115612a8c576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60755467ffffffffffffffff68010000000000000000820481165f8181526074602052604081205491937001000000000000000000000000000000009004909216915b84811015612d13575f878783818110612aea57612aea615841565b9050602002810190612afc919061586e565b612b05906158aa565b905083612b1181615612565b825180516020918201208185015160408087015190519499509194505f93612b729386939101928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b60408051601f19818403018152918152815160209283012067ffffffffffffffff89165f90815260739093529120549091508114612bdc576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff86165f90815260736020526040812055612c00600189615936565b8403612c6f5742607d60149054906101000a900467ffffffffffffffff168460400151612c2d91906155b2565b67ffffffffffffffff161115612c6f576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020838101516040805192830188905282018490526060808301919091524260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016608083015233901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088820152609c016040516020818303038152906040528051906020012094505050508080612d0b90615949565b915050612acf565b50612d1e84846155b2565b6075805467ffffffffffffffff4281167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009092168217808455604080516060810182528781526020808201958652680100000000000000009384900485168284019081528589165f818152607490935284832093518455965160019390930180549151871686027fffffffffffffffffffffffffffffffff0000000000000000000000000000000090921693871693909317179091558554938916700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff938602939093167fffffffffffffffff00000000000000000000000000000000ffffffffffffffff90941693909317919091179093559151929550917f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a49190a2505050505050565b607c546001600160a01b03163314612eb9576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663dbc169766040518163ffffffff1660e01b81526004015f604051808303815f87803b158015612f11575f80fd5b505af1158015612f23573d5f803e3d5ffd5b50505050611b23614a8b565b607d547c0100000000000000000000000000000000000000000000000000000000900460ff1615612f8c576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60715460ff1615612fc9576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f612fd26116d1565b90508181111561300e576040517f4732fdb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6130436001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016333084614b19565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156130a0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130c49190615638565b60758054919250780100000000000000000000000000000000000000000000000090910467ffffffffffffffff169060186130fe83615612565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550508484604051613135929190615980565b60408051918290038220602083015281018290527fffffffffffffffff0000000000000000000000000000000000000000000000004260c01b16606082015260680160408051808303601f1901815291815281516020928301206075547801000000000000000000000000000000000000000000000000900467ffffffffffffffff165f9081526073909352912055323303613245576075546040805183815233602082015260609181018290525f91810191909152780100000000000000000000000000000000000000000000000090910467ffffffffffffffff16907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319060800160405180910390a26132a4565b607560189054906101000a900467ffffffffffffffff1667ffffffffffffffff167ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9318233888860405161329b949392919061598f565b60405180910390a25b5050505050565b60715460ff16156132e8576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6071546b01000000000000000000000090046001600160a01b0316331461333b576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b835f819003613376576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88111156133b2576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b818110156134975761348f8585838181106133d1576133d1615841565b90506020028101906133e391906159b7565b6133f19060208101906159e9565b86868481811061340357613403615841565b905060200281019061341591906159b7565b613423906020810190615a0c565b88888681811061343557613435615841565b905060200281019061344791906159b7565b6040013589898781811061345d5761345d615841565b905060200281019061346f91906159b7565b606001358c8c8881811061348557613485615841565b5050505050505050565b6001016133b4565b5060755467ffffffffffffffff68010000000000000000820481165f81815260746020526040812054838516949293700100000000000000000000000000000000909304909216919082905b86811015613869575f8c8c838181106134fe576134fe615841565b9050608002018036038101906135149190615a70565b805160608201519192509067ffffffffffffffff1615613663578561353881615612565b9650505f818360200151846060015160405160200161358f93929190928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b60408051601f19818403018152918152815160209283012067ffffffffffffffff8a165f908152607390935291205490915081146135f9576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8088165f90815260736020526040808220919091556060850151908501519082169116101561365d576040517f7f7ab87200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50613751565b60208201511580159061371a575060208201516040517f257b363200000000000000000000000000000000000000000000000000000000815260048101919091527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063257b3632906024016020604051808303815f875af11580156136f4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906137189190615638565b155b15613751576040517f73bd668d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8767ffffffffffffffff16826040015167ffffffffffffffff161080613784575042826040015167ffffffffffffffff16115b156137bb576040517fea82791600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602082810151604080850151815193840189905290830184905260608084019290925260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528b901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088820152609c01604051602081830303815290604052805190602001209450816040015197505050808061386190615949565b9150506134e3565b5061387486856155b2565b60755490945067ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811690841611156138dd576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6138e882856155f1565b6138fc9067ffffffffffffffff1688615936565b604080516060810182528581524267ffffffffffffffff908116602080840191825260758054680100000000000000009081900485168688019081528d86165f8181526074909552979093209551865592516001909501805492519585167fffffffffffffffffffffffffffffffff000000000000000000000000000000009384161795851684029590951790945583548c84169116179302929092179055909150828116908516146139f157607580547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8716021790555b613a36333083607254613a0491906155da565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016929190614b19565b613a3e6147ea565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b158015613a96575f80fd5b505af1158015613aa8573d5f803e3d5ffd5b505060405167ffffffffffffffff881692507f303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce91505f90a2505050505050505050505050565b607c546001600160a01b03163314613b32576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607680547fffffffff0000000000000000000000000000000000000000ffffffffffffffff16680100000000000000006001600160a01b038416908102919091179091556040519081527f61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca90602001610c6a565b613bae614899565b6001600160a01b038116613c44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161249b565b6114588161490d565b5f80613c57612309565b905067ffffffffffffffff881615613d2657607b5467ffffffffffffffff9081169089161115613cb3576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8089165f908152607a60205260409020600281015481549094509091898116680100000000000000009092041614613d20576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50613dc6565b67ffffffffffffffff87165f90815260776020526040902054915081613d78576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168767ffffffffffffffff161115613dc6576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168667ffffffffffffffff1611613e13576040517fb9b18f5700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f613e218888888689610c75565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001600283604051613e559190615adf565b602060405180830381855afa158015613e70573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190613e939190615638565b613e9d9190615af0565b6040805160208101825282815290517f9121da8a0000000000000000000000000000000000000000000000000000000081529192506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691639121da8a91613f1291899190600401615b03565b602060405180830381865afa158015613f2d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613f519190615b3d565b613f87576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613fe233613f95858b6155f1565b67ffffffffffffffff16613fa7611d37565b613fb191906155da565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169190614bd8565b50505050505050505050565b5f67ffffffffffffffff8816156140ba57607b5467ffffffffffffffff9081169089161115614049576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5067ffffffffffffffff8088165f908152607a602052604090206002810154815490928881166801000000000000000090920416146140b4576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50614155565b5067ffffffffffffffff85165f908152607760205260409020548061410b576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60765467ffffffffffffffff9081169087161115614155576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b5467ffffffffffffffff908116908816118061418757508767ffffffffffffffff168767ffffffffffffffff1611155b806141ae5750607b5467ffffffffffffffff68010000000000000000909104811690881611155b156141e5576040517fbfa7079f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8781165f908152607a6020526040902054680100000000000000009004811690861614614247576040517f32a2a77f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6142558787878588610c75565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016002836040516142899190615adf565b602060405180830381855afa1580156142a4573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906142c79190615638565b6142d19190615af0565b6040805160208101825282815290517f9121da8a0000000000000000000000000000000000000000000000000000000081529192506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691639121da8a9161434691889190600401615b03565b602060405180830381865afa158015614361573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906143859190615b3d565b6143bb576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff89165f908152607a6020526040902060020154859003613fe2576040517fa47276bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b5467ffffffffffffffff68010000000000000000909104811690821611158061444a5750607b5467ffffffffffffffff908116908216115b15614481576040517fd086b70b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8181165f818152607a6020908152604080832080546076805468010000000000000000928390049098167fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000090981688179055600282015487865260779094529382902092909255607b80547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff169390940292909217909255600182015490517f33d6247d00000000000000000000000000000000000000000000000000000000815260048101919091529091907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906333d6247d906024015f604051808303815f87803b1580156145a2575f80fd5b505af11580156145b4573d5f803e3d5ffd5b505050508267ffffffffffffffff168167ffffffffffffffff167f328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e846002015460405161460391815260200190565b60405180910390a3505050565b5f614619612309565b9050815f8061462884846155f1565b60715467ffffffffffffffff91821692505f9161464b9161010090041642615936565b90505b8467ffffffffffffffff168467ffffffffffffffff16146146d55767ffffffffffffffff8085165f90815260746020526040902060018101549091168210156146b357600181015468010000000000000000900467ffffffffffffffff1694506146cf565b6146bd86866155f1565b67ffffffffffffffff169350506146d5565b5061464e565b5f6146e08484615936565b90508381101561473757808403600c81116146fb57806146fe565b600c5b9050806103e80a81607160099054906101000a900461ffff1661ffff160a607254028161472d5761472d61564f565b04607255506147a6565b838103600c8111614748578061474b565b600c5b90505f816103e80a82607160099054906101000a900461ffff1661ffff160a670de0b6b3a764000002816147815761478161564f565b04905080607254670de0b6b3a7640000028161479f5761479f61564f565b0460725550505b683635c9adc5dea0000060725411156147cb57683635c9adc5dea00000607255611f24565b633b9aca006072541015611f2457633b9aca0060725550505050505050565b607b5467ffffffffffffffff680100000000000000008204811691161115611b2357607b545f906148329068010000000000000000900467ffffffffffffffff1660016155b2565b905061483d816111e9565b1561145857607b545f9060029061485f90849067ffffffffffffffff166155f1565b6148699190615b5c565b61487390836155b2565b905061487e816111e9565b156148905761488c81614410565b5050565b61488c82614410565b6033546001600160a01b03163314611b23576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161249b565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632072f6c56040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156149ce575f80fd5b505af11580156149e0573d5f803e3d5ffd5b50505050611b23614c26565b5f54610100900460ff16614a82576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161249b565b611b233361490d565b60715460ff16614ac7576040517f5386698100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b6040516001600160a01b0380851660248301528316604482015260648101829052614bca9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152614cb8565b50505050565b505050505050565b6040516001600160a01b038316602482015260448101829052614c219084907fa9059cbb0000000000000000000000000000000000000000000000000000000090606401614b66565b505050565b60715460ff1615614c63576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b5f614d0c826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316614db69092919063ffffffff16565b805190915015614c215780806020019051810190614d2a9190615b3d565b614c21576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161249b565b6060614dc484845f85614dcc565b949350505050565b606082471015614e5e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161249b565b5f80866001600160a01b03168587604051614e799190615adf565b5f6040518083038185875af1925050503d805f8114614eb3576040519150601f19603f3d011682016040523d82523d5f602084013e614eb8565b606091505b5091509150614ec987838387614ed4565b979650505050505050565b60608315614f5c5782515f03614f55576001600160a01b0385163b614f55576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161249b565b5081614dc4565b614dc48383815115614f715781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161249b9190614ff2565b5f5b83811015614fbf578181015183820152602001614fa7565b50505f910152565b5f8151808452614fde816020860160208601614fa5565b601f01601f19169290920160200192915050565b602081525f6150046020830184614fc7565b9392505050565b5f6020828403121561501b575f80fd5b813561ffff81168114615004575f80fd5b803567ffffffffffffffff811681146122c2575f80fd5b5f805f805f60a08688031215615057575f80fd5b6150608661502c565b945061506e6020870161502c565b94979496505050506040830135926060810135926080909101359150565b80610300810183101561509d575f80fd5b92915050565b5f805f805f806103a087890312156150b9575f80fd5b6150c28761502c565b95506150d06020880161502c565b94506150de6040880161502c565b935060608701359250608087013591506150fb8860a0890161508c565b90509295509295509295565b5f805f805f805f6103c0888a03121561511e575f80fd5b6151278861502c565b96506151356020890161502c565b95506151436040890161502c565b94506151516060890161502c565b93506080880135925060a0880135915061516e8960c08a0161508c565b905092959891949750929550565b5f6020828403121561518c575f80fd5b6150048261502c565b5f602082840312156151a5575f80fd5b5035919050565b6001600160a01b0381168114611458575f80fd5b80356122c2816151ac565b5f602082840312156151db575f80fd5b8135615004816151ac565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f67ffffffffffffffff8084111561522d5761522d6151e6565b604051601f8501601f19908116603f01168101908282118183101715615255576152556151e6565b8160405280935085815286868601111561526d575f80fd5b858560208301375f602087830101525050509392505050565b5f82601f830112615295575f80fd5b61500483833560208501615213565b5f602082840312156152b4575f80fd5b813567ffffffffffffffff8111156152ca575f80fd5b614dc484828501615286565b5f8083601f8401126152e6575f80fd5b50813567ffffffffffffffff8111156152fd575f80fd5b602083019150836020828501011115615314575f80fd5b9250929050565b5f805f805f80868803610120811215615332575f80fd5b60a081121561533f575f80fd5b5086955060a0870135945060c087013567ffffffffffffffff80821115615364575f80fd5b6153708a838b01615286565b955060e0890135915080821115615385575f80fd5b6153918a838b01615286565b94506101008901359150808211156153a7575f80fd5b506153b489828a016152d6565b979a9699509497509295939492505050565b5f8083601f8401126153d6575f80fd5b50813567ffffffffffffffff8111156153ed575f80fd5b6020830191508360208260051b8501011115615314575f80fd5b5f8060208385031215615418575f80fd5b823567ffffffffffffffff81111561542e575f80fd5b61543a858286016153c6565b90969095509350505050565b5f805f60408486031215615458575f80fd5b833567ffffffffffffffff81111561546e575f80fd5b61547a868287016152d6565b909790965060209590950135949350505050565b5f805f805f606086880312156154a2575f80fd5b853567ffffffffffffffff808211156154b9575f80fd5b818801915088601f8301126154cc575f80fd5b8135818111156154da575f80fd5b8960208260071b85010111156154ee575f80fd5b602092830197509550908701359080821115615508575f80fd5b50615515888289016153c6565b90945092506155289050604087016151c0565b90509295509295909350565b600181811c9082168061554857607f821691505b60208210810361557f577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b67ffffffffffffffff8181168382160190808211156155d3576155d3615585565b5092915050565b808202811582820484141761509d5761509d615585565b67ffffffffffffffff8281168282160390808211156155d3576155d3615585565b5f67ffffffffffffffff80831681810361562e5761562e615585565b6001019392505050565b5f60208284031215615648575f80fd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f8261568a5761568a61564f565b500490565b601f821115614c21575f81815260208120601f850160051c810160208610156156b55750805b601f850160051c820191505b81811015614bd0578281556001016156c1565b815167ffffffffffffffff8111156156ee576156ee6151e6565b615702816156fc8454615534565b8461568f565b602080601f831160018114615754575f841561571e5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555614bd0565b5f85815260208120601f198616915b8281101561578257888601518255948401946001909101908401615763565b50858210156157be57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b81835281816020850137505f602082840101525f6020601f19601f840116840101905092915050565b5f67ffffffffffffffff8087168352808616602084015250606060408301526158246060830184866157ce565b9695505050505050565b8082018082111561509d5761509d615585565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa18336030181126158a0575f80fd5b9190910192915050565b5f606082360312156158ba575f80fd5b6040516060810167ffffffffffffffff82821081831117156158de576158de6151e6565b8160405284359150808211156158f2575f80fd5b50830136601f820112615903575f80fd5b61591236823560208401615213565b8252506020830135602082015261592b6040840161502c565b604082015292915050565b8181038181111561509d5761509d615585565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361597957615979615585565b5060010190565b818382375f9101908152919050565b8481526001600160a01b0384166020820152606060408201525f6158246060830184866157ce565b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818336030181126158a0575f80fd5b5f602082840312156159f9575f80fd5b813563ffffffff81168114615004575f80fd5b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112615a3f575f80fd5b83018035915067ffffffffffffffff821115615a59575f80fd5b6020019150600581901b3603821315615314575f80fd5b5f60808284031215615a80575f80fd5b6040516080810181811067ffffffffffffffff82111715615aa357615aa36151e6565b80604052508235815260208301356020820152615ac26040840161502c565b6040820152615ad36060840161502c565b60608201529392505050565b5f82516158a0818460208701614fa5565b5f82615afe57615afe61564f565b500690565b6103208101610300808584378201835f5b6001811015615b33578151835260209283019290910190600101615b14565b5050509392505050565b5f60208284031215615b4d575f80fd5b81518015158114615004575f80fd5b5f67ffffffffffffffff80841680615b7657615b7661564f565b9216919091049291505056fea2646970667358221220855074d2996c7ccb0d681345a450fb738458fab74812841930af4df69c784b0764736f6c63430008140033",
}

// PolygonzkevmABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonzkevmMetaData.ABI instead.
var PolygonzkevmABI = PolygonzkevmMetaData.ABI

// PolygonzkevmBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonzkevmMetaData.Bin instead.
var PolygonzkevmBin = PolygonzkevmMetaData.Bin

// DeployPolygonzkevm deploys a new Ethereum contract, binding an instance of Polygonzkevm to it.
func DeployPolygonzkevm(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _dataAvailabilityRouter common.Address, _matic common.Address, _rollupVerifier common.Address, _bridgeAddress common.Address, _chainID uint64, _forkID uint64) (common.Address, *types.Transaction, *Polygonzkevm, error) {
	parsed, err := PolygonzkevmMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonzkevmBin), backend, _globalExitRootManager, _dataAvailabilityRouter, _matic, _rollupVerifier, _bridgeAddress, _chainID, _forkID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonzkevm{PolygonzkevmCaller: PolygonzkevmCaller{contract: contract}, PolygonzkevmTransactor: PolygonzkevmTransactor{contract: contract}, PolygonzkevmFilterer: PolygonzkevmFilterer{contract: contract}}, nil
}

// Polygonzkevm is an auto generated Go binding around an Ethereum contract.
type Polygonzkevm struct {
	PolygonzkevmCaller     // Read-only binding to the contract
	PolygonzkevmTransactor // Write-only binding to the contract
	PolygonzkevmFilterer   // Log filterer for contract events
}

// PolygonzkevmCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonzkevmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonzkevmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonzkevmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonzkevmSession struct {
	Contract     *Polygonzkevm     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolygonzkevmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonzkevmCallerSession struct {
	Contract *PolygonzkevmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PolygonzkevmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonzkevmTransactorSession struct {
	Contract     *PolygonzkevmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PolygonzkevmRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonzkevmRaw struct {
	Contract *Polygonzkevm // Generic contract binding to access the raw methods on
}

// PolygonzkevmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonzkevmCallerRaw struct {
	Contract *PolygonzkevmCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonzkevmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonzkevmTransactorRaw struct {
	Contract *PolygonzkevmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonzkevm creates a new instance of Polygonzkevm, bound to a specific deployed contract.
func NewPolygonzkevm(address common.Address, backend bind.ContractBackend) (*Polygonzkevm, error) {
	contract, err := bindPolygonzkevm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevm{PolygonzkevmCaller: PolygonzkevmCaller{contract: contract}, PolygonzkevmTransactor: PolygonzkevmTransactor{contract: contract}, PolygonzkevmFilterer: PolygonzkevmFilterer{contract: contract}}, nil
}

// NewPolygonzkevmCaller creates a new read-only instance of Polygonzkevm, bound to a specific deployed contract.
func NewPolygonzkevmCaller(address common.Address, caller bind.ContractCaller) (*PolygonzkevmCaller, error) {
	contract, err := bindPolygonzkevm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmCaller{contract: contract}, nil
}

// NewPolygonzkevmTransactor creates a new write-only instance of Polygonzkevm, bound to a specific deployed contract.
func NewPolygonzkevmTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonzkevmTransactor, error) {
	contract, err := bindPolygonzkevm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmTransactor{contract: contract}, nil
}

// NewPolygonzkevmFilterer creates a new log filterer instance of Polygonzkevm, bound to a specific deployed contract.
func NewPolygonzkevmFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonzkevmFilterer, error) {
	contract, err := bindPolygonzkevm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmFilterer{contract: contract}, nil
}

// bindPolygonzkevm binds a generic wrapper to an already deployed contract.
func bindPolygonzkevm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonzkevmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevm *PolygonzkevmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevm.Contract.PolygonzkevmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevm *PolygonzkevmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.PolygonzkevmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevm *PolygonzkevmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.PolygonzkevmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevm *PolygonzkevmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevm *PolygonzkevmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevm *PolygonzkevmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) Admin() (common.Address, error) {
	return _Polygonzkevm.Contract.Admin(&_Polygonzkevm.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) Admin() (common.Address, error) {
	return _Polygonzkevm.Contract.Admin(&_Polygonzkevm.CallOpts)
}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmCaller) BatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "batchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmSession) BatchFee() (*big.Int, error) {
	return _Polygonzkevm.Contract.BatchFee(&_Polygonzkevm.CallOpts)
}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmCallerSession) BatchFee() (*big.Int, error) {
	return _Polygonzkevm.Contract.BatchFee(&_Polygonzkevm.CallOpts)
}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmCaller) BatchNumToStateRoot(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "batchNumToStateRoot", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmSession) BatchNumToStateRoot(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevm.Contract.BatchNumToStateRoot(&_Polygonzkevm.CallOpts, arg0)
}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmCallerSession) BatchNumToStateRoot(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevm.Contract.BatchNumToStateRoot(&_Polygonzkevm.CallOpts, arg0)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevm.Contract.BridgeAddress(&_Polygonzkevm.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevm.Contract.BridgeAddress(&_Polygonzkevm.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmCaller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonzkevm.Contract.CalculateRewardPerBatch(&_Polygonzkevm.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmCallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonzkevm.Contract.CalculateRewardPerBatch(&_Polygonzkevm.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCaller) ChainID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmSession) ChainID() (uint64, error) {
	return _Polygonzkevm.Contract.ChainID(&_Polygonzkevm.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCallerSession) ChainID() (uint64, error) {
	return _Polygonzkevm.Contract.ChainID(&_Polygonzkevm.CallOpts)
}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Polygonzkevm *PolygonzkevmCaller) CheckStateRootInsidePrime(opts *bind.CallOpts, newStateRoot *big.Int) (bool, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "checkStateRootInsidePrime", newStateRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Polygonzkevm *PolygonzkevmSession) CheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Polygonzkevm.Contract.CheckStateRootInsidePrime(&_Polygonzkevm.CallOpts, newStateRoot)
}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Polygonzkevm *PolygonzkevmCallerSession) CheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Polygonzkevm.Contract.CheckStateRootInsidePrime(&_Polygonzkevm.CallOpts, newStateRoot)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonzkevm.Contract.ForceBatchTimeout(&_Polygonzkevm.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonzkevm.Contract.ForceBatchTimeout(&_Polygonzkevm.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevm.Contract.ForcedBatches(&_Polygonzkevm.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevm *PolygonzkevmCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevm.Contract.ForcedBatches(&_Polygonzkevm.CallOpts, arg0)
}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCaller) ForkID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "forkID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmSession) ForkID() (uint64, error) {
	return _Polygonzkevm.Contract.ForkID(&_Polygonzkevm.CallOpts)
}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCallerSession) ForkID() (uint64, error) {
	return _Polygonzkevm.Contract.ForkID(&_Polygonzkevm.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmCaller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmSession) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonzkevm.Contract.GetForcedBatchFee(&_Polygonzkevm.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonzkevm *PolygonzkevmCallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonzkevm.Contract.GetForcedBatchFee(&_Polygonzkevm.CallOpts)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonzkevm *PolygonzkevmCaller) GetInputSnarkBytes(opts *bind.CallOpts, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "getInputSnarkBytes", initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonzkevm *PolygonzkevmSession) GetInputSnarkBytes(initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonzkevm.Contract.GetInputSnarkBytes(&_Polygonzkevm.CallOpts, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonzkevm *PolygonzkevmCallerSession) GetInputSnarkBytes(initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonzkevm.Contract.GetInputSnarkBytes(&_Polygonzkevm.CallOpts, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCaller) GetLastVerifiedBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "getLastVerifiedBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmSession) GetLastVerifiedBatch() (uint64, error) {
	return _Polygonzkevm.Contract.GetLastVerifiedBatch(&_Polygonzkevm.CallOpts)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCallerSession) GetLastVerifiedBatch() (uint64, error) {
	return _Polygonzkevm.Contract.GetLastVerifiedBatch(&_Polygonzkevm.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevm.Contract.GlobalExitRootManager(&_Polygonzkevm.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevm.Contract.GlobalExitRootManager(&_Polygonzkevm.CallOpts)
}

// IsBatchPosted is a free data retrieval call binding the contract method 0x40b10f1c.
//
// Solidity: function isBatchPosted(bytes32 ) view returns(uint256)
func (_Polygonzkevm *PolygonzkevmCaller) IsBatchPosted(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "isBatchPosted", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsBatchPosted is a free data retrieval call binding the contract method 0x40b10f1c.
//
// Solidity: function isBatchPosted(bytes32 ) view returns(uint256)
func (_Polygonzkevm *PolygonzkevmSession) IsBatchPosted(arg0 [32]byte) (*big.Int, error) {
	return _Polygonzkevm.Contract.IsBatchPosted(&_Polygonzkevm.CallOpts, arg0)
}

// IsBatchPosted is a free data retrieval call binding the contract method 0x40b10f1c.
//
// Solidity: function isBatchPosted(bytes32 ) view returns(uint256)
func (_Polygonzkevm *PolygonzkevmCallerSession) IsBatchPosted(arg0 [32]byte) (*big.Int, error) {
	return _Polygonzkevm.Contract.IsBatchPosted(&_Polygonzkevm.CallOpts, arg0)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonzkevm *PolygonzkevmCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonzkevm *PolygonzkevmSession) IsEmergencyState() (bool, error) {
	return _Polygonzkevm.Contract.IsEmergencyState(&_Polygonzkevm.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonzkevm *PolygonzkevmCallerSession) IsEmergencyState() (bool, error) {
	return _Polygonzkevm.Contract.IsEmergencyState(&_Polygonzkevm.CallOpts)
}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Polygonzkevm *PolygonzkevmCaller) IsForcedBatchDisallowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "isForcedBatchDisallowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Polygonzkevm *PolygonzkevmSession) IsForcedBatchDisallowed() (bool, error) {
	return _Polygonzkevm.Contract.IsForcedBatchDisallowed(&_Polygonzkevm.CallOpts)
}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Polygonzkevm *PolygonzkevmCallerSession) IsForcedBatchDisallowed() (bool, error) {
	return _Polygonzkevm.Contract.IsForcedBatchDisallowed(&_Polygonzkevm.CallOpts)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Polygonzkevm *PolygonzkevmCaller) IsPendingStateConsolidable(opts *bind.CallOpts, pendingStateNum uint64) (bool, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "isPendingStateConsolidable", pendingStateNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Polygonzkevm *PolygonzkevmSession) IsPendingStateConsolidable(pendingStateNum uint64) (bool, error) {
	return _Polygonzkevm.Contract.IsPendingStateConsolidable(&_Polygonzkevm.CallOpts, pendingStateNum)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Polygonzkevm *PolygonzkevmCallerSession) IsPendingStateConsolidable(pendingStateNum uint64) (bool, error) {
	return _Polygonzkevm.Contract.IsPendingStateConsolidable(&_Polygonzkevm.CallOpts, pendingStateNum)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCaller) LastBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "lastBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmSession) LastBatchSequenced() (uint64, error) {
	return _Polygonzkevm.Contract.LastBatchSequenced(&_Polygonzkevm.CallOpts)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCallerSession) LastBatchSequenced() (uint64, error) {
	return _Polygonzkevm.Contract.LastBatchSequenced(&_Polygonzkevm.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmSession) LastForceBatch() (uint64, error) {
	return _Polygonzkevm.Contract.LastForceBatch(&_Polygonzkevm.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCallerSession) LastForceBatch() (uint64, error) {
	return _Polygonzkevm.Contract.LastForceBatch(&_Polygonzkevm.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonzkevm.Contract.LastForceBatchSequenced(&_Polygonzkevm.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonzkevm.Contract.LastForceBatchSequenced(&_Polygonzkevm.CallOpts)
}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCaller) LastPendingState(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "lastPendingState")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmSession) LastPendingState() (uint64, error) {
	return _Polygonzkevm.Contract.LastPendingState(&_Polygonzkevm.CallOpts)
}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCallerSession) LastPendingState() (uint64, error) {
	return _Polygonzkevm.Contract.LastPendingState(&_Polygonzkevm.CallOpts)
}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCaller) LastPendingStateConsolidated(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "lastPendingStateConsolidated")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmSession) LastPendingStateConsolidated() (uint64, error) {
	return _Polygonzkevm.Contract.LastPendingStateConsolidated(&_Polygonzkevm.CallOpts)
}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCallerSession) LastPendingStateConsolidated() (uint64, error) {
	return _Polygonzkevm.Contract.LastPendingStateConsolidated(&_Polygonzkevm.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCaller) LastTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "lastTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmSession) LastTimestamp() (uint64, error) {
	return _Polygonzkevm.Contract.LastTimestamp(&_Polygonzkevm.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCallerSession) LastTimestamp() (uint64, error) {
	return _Polygonzkevm.Contract.LastTimestamp(&_Polygonzkevm.CallOpts)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCaller) LastVerifiedBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "lastVerifiedBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmSession) LastVerifiedBatch() (uint64, error) {
	return _Polygonzkevm.Contract.LastVerifiedBatch(&_Polygonzkevm.CallOpts)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCallerSession) LastVerifiedBatch() (uint64, error) {
	return _Polygonzkevm.Contract.LastVerifiedBatch(&_Polygonzkevm.CallOpts)
}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) Matic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "matic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) Matic() (common.Address, error) {
	return _Polygonzkevm.Contract.Matic(&_Polygonzkevm.CallOpts)
}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) Matic() (common.Address, error) {
	return _Polygonzkevm.Contract.Matic(&_Polygonzkevm.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonzkevm *PolygonzkevmCaller) MultiplierBatchFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "multiplierBatchFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonzkevm *PolygonzkevmSession) MultiplierBatchFee() (uint16, error) {
	return _Polygonzkevm.Contract.MultiplierBatchFee(&_Polygonzkevm.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonzkevm *PolygonzkevmCallerSession) MultiplierBatchFee() (uint16, error) {
	return _Polygonzkevm.Contract.MultiplierBatchFee(&_Polygonzkevm.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevm *PolygonzkevmCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevm *PolygonzkevmSession) NetworkName() (string, error) {
	return _Polygonzkevm.Contract.NetworkName(&_Polygonzkevm.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevm *PolygonzkevmCallerSession) NetworkName() (string, error) {
	return _Polygonzkevm.Contract.NetworkName(&_Polygonzkevm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) Owner() (common.Address, error) {
	return _Polygonzkevm.Contract.Owner(&_Polygonzkevm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) Owner() (common.Address, error) {
	return _Polygonzkevm.Contract.Owner(&_Polygonzkevm.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) PendingAdmin() (common.Address, error) {
	return _Polygonzkevm.Contract.PendingAdmin(&_Polygonzkevm.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) PendingAdmin() (common.Address, error) {
	return _Polygonzkevm.Contract.PendingAdmin(&_Polygonzkevm.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCaller) PendingStateTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "pendingStateTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmSession) PendingStateTimeout() (uint64, error) {
	return _Polygonzkevm.Contract.PendingStateTimeout(&_Polygonzkevm.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCallerSession) PendingStateTimeout() (uint64, error) {
	return _Polygonzkevm.Contract.PendingStateTimeout(&_Polygonzkevm.CallOpts)
}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns(uint64 timestamp, uint64 lastVerifiedBatch, bytes32 exitRoot, bytes32 stateRoot)
func (_Polygonzkevm *PolygonzkevmCaller) PendingStateTransitions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "pendingStateTransitions", arg0)

	outstruct := new(struct {
		Timestamp         uint64
		LastVerifiedBatch uint64
		ExitRoot          [32]byte
		StateRoot         [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.LastVerifiedBatch = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ExitRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.StateRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns(uint64 timestamp, uint64 lastVerifiedBatch, bytes32 exitRoot, bytes32 stateRoot)
func (_Polygonzkevm *PolygonzkevmSession) PendingStateTransitions(arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	return _Polygonzkevm.Contract.PendingStateTransitions(&_Polygonzkevm.CallOpts, arg0)
}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns(uint64 timestamp, uint64 lastVerifiedBatch, bytes32 exitRoot, bytes32 stateRoot)
func (_Polygonzkevm *PolygonzkevmCallerSession) PendingStateTransitions(arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	return _Polygonzkevm.Contract.PendingStateTransitions(&_Polygonzkevm.CallOpts, arg0)
}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) RollupVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "rollupVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) RollupVerifier() (common.Address, error) {
	return _Polygonzkevm.Contract.RollupVerifier(&_Polygonzkevm.CallOpts)
}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) RollupVerifier() (common.Address, error) {
	return _Polygonzkevm.Contract.RollupVerifier(&_Polygonzkevm.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) Router() (common.Address, error) {
	return _Polygonzkevm.Contract.Router(&_Polygonzkevm.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) Router() (common.Address, error) {
	return _Polygonzkevm.Contract.Router(&_Polygonzkevm.CallOpts)
}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns(bytes32 accInputHash, uint64 sequencedTimestamp, uint64 previousLastBatchSequenced)
func (_Polygonzkevm *PolygonzkevmCaller) SequencedBatches(opts *bind.CallOpts, arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "sequencedBatches", arg0)

	outstruct := new(struct {
		AccInputHash               [32]byte
		SequencedTimestamp         uint64
		PreviousLastBatchSequenced uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AccInputHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.SequencedTimestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.PreviousLastBatchSequenced = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns(bytes32 accInputHash, uint64 sequencedTimestamp, uint64 previousLastBatchSequenced)
func (_Polygonzkevm *PolygonzkevmSession) SequencedBatches(arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	return _Polygonzkevm.Contract.SequencedBatches(&_Polygonzkevm.CallOpts, arg0)
}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns(bytes32 accInputHash, uint64 sequencedTimestamp, uint64 previousLastBatchSequenced)
func (_Polygonzkevm *PolygonzkevmCallerSession) SequencedBatches(arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	return _Polygonzkevm.Contract.SequencedBatches(&_Polygonzkevm.CallOpts, arg0)
}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) TrustedAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "trustedAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) TrustedAggregator() (common.Address, error) {
	return _Polygonzkevm.Contract.TrustedAggregator(&_Polygonzkevm.CallOpts)
}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) TrustedAggregator() (common.Address, error) {
	return _Polygonzkevm.Contract.TrustedAggregator(&_Polygonzkevm.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCaller) TrustedAggregatorTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "trustedAggregatorTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Polygonzkevm.Contract.TrustedAggregatorTimeout(&_Polygonzkevm.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCallerSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Polygonzkevm.Contract.TrustedAggregatorTimeout(&_Polygonzkevm.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevm *PolygonzkevmCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevm *PolygonzkevmSession) TrustedSequencer() (common.Address, error) {
	return _Polygonzkevm.Contract.TrustedSequencer(&_Polygonzkevm.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevm *PolygonzkevmCallerSession) TrustedSequencer() (common.Address, error) {
	return _Polygonzkevm.Contract.TrustedSequencer(&_Polygonzkevm.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevm *PolygonzkevmCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevm *PolygonzkevmSession) TrustedSequencerURL() (string, error) {
	return _Polygonzkevm.Contract.TrustedSequencerURL(&_Polygonzkevm.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevm *PolygonzkevmCallerSession) TrustedSequencerURL() (string, error) {
	return _Polygonzkevm.Contract.TrustedSequencerURL(&_Polygonzkevm.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCaller) VerifyBatchTimeTarget(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevm.contract.Call(opts, &out, "verifyBatchTimeTarget")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Polygonzkevm.Contract.VerifyBatchTimeTarget(&_Polygonzkevm.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonzkevm *PolygonzkevmCallerSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Polygonzkevm.Contract.VerifyBatchTimeTarget(&_Polygonzkevm.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevm *PolygonzkevmTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevm *PolygonzkevmSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonzkevm.Contract.AcceptAdminRole(&_Polygonzkevm.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonzkevm.Contract.AcceptAdminRole(&_Polygonzkevm.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) ActivateEmergencyState(opts *bind.TransactOpts, sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "activateEmergencyState", sequencedBatchNum)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Polygonzkevm *PolygonzkevmSession) ActivateEmergencyState(sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.ActivateEmergencyState(&_Polygonzkevm.TransactOpts, sequencedBatchNum)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) ActivateEmergencyState(sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.ActivateEmergencyState(&_Polygonzkevm.TransactOpts, sequencedBatchNum)
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Polygonzkevm *PolygonzkevmTransactor) ActivateForceBatches(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "activateForceBatches")
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Polygonzkevm *PolygonzkevmSession) ActivateForceBatches() (*types.Transaction, error) {
	return _Polygonzkevm.Contract.ActivateForceBatches(&_Polygonzkevm.TransactOpts)
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) ActivateForceBatches() (*types.Transaction, error) {
	return _Polygonzkevm.Contract.ActivateForceBatches(&_Polygonzkevm.TransactOpts)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) ConsolidatePendingState(opts *bind.TransactOpts, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "consolidatePendingState", pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Polygonzkevm *PolygonzkevmSession) ConsolidatePendingState(pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.ConsolidatePendingState(&_Polygonzkevm.TransactOpts, pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) ConsolidatePendingState(pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.ConsolidatePendingState(&_Polygonzkevm.TransactOpts, pendingStateNum)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonzkevm *PolygonzkevmTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonzkevm *PolygonzkevmSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonzkevm.Contract.DeactivateEmergencyState(&_Polygonzkevm.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonzkevm.Contract.DeactivateEmergencyState(&_Polygonzkevm.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "forceBatch", transactions, maticAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Polygonzkevm *PolygonzkevmSession) ForceBatch(transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.ForceBatch(&_Polygonzkevm.TransactOpts, transactions, maticAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) ForceBatch(transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.ForceBatch(&_Polygonzkevm.TransactOpts, transactions, maticAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) Initialize(opts *bind.TransactOpts, initializePackedParameters PolygonZkEVMInitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "initialize", initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Polygonzkevm *PolygonzkevmSession) Initialize(initializePackedParameters PolygonZkEVMInitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.Initialize(&_Polygonzkevm.TransactOpts, initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) Initialize(initializePackedParameters PolygonZkEVMInitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.Initialize(&_Polygonzkevm.TransactOpts, initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) OverridePendingState(opts *bind.TransactOpts, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "overridePendingState", initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevm *PolygonzkevmSession) OverridePendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.OverridePendingState(&_Polygonzkevm.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) OverridePendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.OverridePendingState(&_Polygonzkevm.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) ProveNonDeterministicPendingState(opts *bind.TransactOpts, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "proveNonDeterministicPendingState", initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevm *PolygonzkevmSession) ProveNonDeterministicPendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.ProveNonDeterministicPendingState(&_Polygonzkevm.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) ProveNonDeterministicPendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.ProveNonDeterministicPendingState(&_Polygonzkevm.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Polygonzkevm *PolygonzkevmTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Polygonzkevm *PolygonzkevmSession) RenounceOwnership() (*types.Transaction, error) {
	return _Polygonzkevm.Contract.RenounceOwnership(&_Polygonzkevm.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Polygonzkevm.Contract.RenounceOwnership(&_Polygonzkevm.TransactOpts)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xefde5290.
//
// Solidity: function sequenceBatches((bytes32,bytes32,uint64,uint64)[] batches, (uint32,bytes32[],uint256,uint256)[] daData, address l2Coinbase) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SequenceBatches(opts *bind.TransactOpts, batches []PolygonZkEVMBatchData, daData []PolygonZkEVMDAData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "sequenceBatches", batches, daData, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xefde5290.
//
// Solidity: function sequenceBatches((bytes32,bytes32,uint64,uint64)[] batches, (uint32,bytes32[],uint256,uint256)[] daData, address l2Coinbase) returns()
func (_Polygonzkevm *PolygonzkevmSession) SequenceBatches(batches []PolygonZkEVMBatchData, daData []PolygonZkEVMDAData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SequenceBatches(&_Polygonzkevm.TransactOpts, batches, daData, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xefde5290.
//
// Solidity: function sequenceBatches((bytes32,bytes32,uint64,uint64)[] batches, (uint32,bytes32[],uint256,uint256)[] daData, address l2Coinbase) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SequenceBatches(batches []PolygonZkEVMBatchData, daData []PolygonZkEVMDAData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SequenceBatches(&_Polygonzkevm.TransactOpts, batches, daData, l2Coinbase)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []PolygonZkEVMForcedBatchData) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Polygonzkevm *PolygonzkevmSession) SequenceForceBatches(batches []PolygonZkEVMForcedBatchData) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SequenceForceBatches(&_Polygonzkevm.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SequenceForceBatches(batches []PolygonZkEVMForcedBatchData) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SequenceForceBatches(&_Polygonzkevm.TransactOpts, batches)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevm *PolygonzkevmSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetForceBatchTimeout(&_Polygonzkevm.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetForceBatchTimeout(&_Polygonzkevm.TransactOpts, newforceBatchTimeout)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SetMultiplierBatchFee(opts *bind.TransactOpts, newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "setMultiplierBatchFee", newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonzkevm *PolygonzkevmSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetMultiplierBatchFee(&_Polygonzkevm.TransactOpts, newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetMultiplierBatchFee(&_Polygonzkevm.TransactOpts, newMultiplierBatchFee)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SetPendingStateTimeout(opts *bind.TransactOpts, newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "setPendingStateTimeout", newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonzkevm *PolygonzkevmSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetPendingStateTimeout(&_Polygonzkevm.TransactOpts, newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetPendingStateTimeout(&_Polygonzkevm.TransactOpts, newPendingStateTimeout)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SetRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "setRouter", _router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_Polygonzkevm *PolygonzkevmSession) SetRouter(_router common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetRouter(&_Polygonzkevm.TransactOpts, _router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SetRouter(_router common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetRouter(&_Polygonzkevm.TransactOpts, _router)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SetTrustedAggregator(opts *bind.TransactOpts, newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "setTrustedAggregator", newTrustedAggregator)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Polygonzkevm *PolygonzkevmSession) SetTrustedAggregator(newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetTrustedAggregator(&_Polygonzkevm.TransactOpts, newTrustedAggregator)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SetTrustedAggregator(newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetTrustedAggregator(&_Polygonzkevm.TransactOpts, newTrustedAggregator)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SetTrustedAggregatorTimeout(opts *bind.TransactOpts, newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "setTrustedAggregatorTimeout", newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonzkevm *PolygonzkevmSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetTrustedAggregatorTimeout(&_Polygonzkevm.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetTrustedAggregatorTimeout(&_Polygonzkevm.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevm *PolygonzkevmSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetTrustedSequencer(&_Polygonzkevm.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetTrustedSequencer(&_Polygonzkevm.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevm *PolygonzkevmSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetTrustedSequencerURL(&_Polygonzkevm.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetTrustedSequencerURL(&_Polygonzkevm.TransactOpts, newTrustedSequencerURL)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) SetVerifyBatchTimeTarget(opts *bind.TransactOpts, newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "setVerifyBatchTimeTarget", newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonzkevm *PolygonzkevmSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetVerifyBatchTimeTarget(&_Polygonzkevm.TransactOpts, newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.SetVerifyBatchTimeTarget(&_Polygonzkevm.TransactOpts, newVerifyBatchTimeTarget)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevm *PolygonzkevmSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.TransferAdminRole(&_Polygonzkevm.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.TransferAdminRole(&_Polygonzkevm.TransactOpts, newPendingAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Polygonzkevm *PolygonzkevmSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.TransferOwnership(&_Polygonzkevm.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.TransferOwnership(&_Polygonzkevm.TransactOpts, newOwner)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) VerifyBatches(opts *bind.TransactOpts, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "verifyBatches", pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevm *PolygonzkevmSession) VerifyBatches(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.VerifyBatches(&_Polygonzkevm.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) VerifyBatches(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.VerifyBatches(&_Polygonzkevm.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevm *PolygonzkevmTransactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevm.contract.Transact(opts, "verifyBatchesTrustedAggregator", pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevm *PolygonzkevmSession) VerifyBatchesTrustedAggregator(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.VerifyBatchesTrustedAggregator(&_Polygonzkevm.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevm *PolygonzkevmTransactorSession) VerifyBatchesTrustedAggregator(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevm.Contract.VerifyBatchesTrustedAggregator(&_Polygonzkevm.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// PolygonzkevmAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Polygonzkevm contract.
type PolygonzkevmAcceptAdminRoleIterator struct {
	Event *PolygonzkevmAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmAcceptAdminRole)
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
		it.Event = new(PolygonzkevmAcceptAdminRole)
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
func (it *PolygonzkevmAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmAcceptAdminRole represents a AcceptAdminRole event raised by the Polygonzkevm contract.
type PolygonzkevmAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*PolygonzkevmAcceptAdminRoleIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmAcceptAdminRoleIterator{contract: _Polygonzkevm.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonzkevmAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmAcceptAdminRole)
				if err := _Polygonzkevm.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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

// ParseAcceptAdminRole is a log parse operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseAcceptAdminRole(log types.Log) (*PolygonzkevmAcceptAdminRole, error) {
	event := new(PolygonzkevmAcceptAdminRole)
	if err := _Polygonzkevm.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmActivateForceBatchesIterator is returned from FilterActivateForceBatches and is used to iterate over the raw logs and unpacked data for ActivateForceBatches events raised by the Polygonzkevm contract.
type PolygonzkevmActivateForceBatchesIterator struct {
	Event *PolygonzkevmActivateForceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmActivateForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmActivateForceBatches)
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
		it.Event = new(PolygonzkevmActivateForceBatches)
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
func (it *PolygonzkevmActivateForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmActivateForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmActivateForceBatches represents a ActivateForceBatches event raised by the Polygonzkevm contract.
type PolygonzkevmActivateForceBatches struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterActivateForceBatches is a free log retrieval operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Polygonzkevm *PolygonzkevmFilterer) FilterActivateForceBatches(opts *bind.FilterOpts) (*PolygonzkevmActivateForceBatchesIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "ActivateForceBatches")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmActivateForceBatchesIterator{contract: _Polygonzkevm.contract, event: "ActivateForceBatches", logs: logs, sub: sub}, nil
}

// WatchActivateForceBatches is a free log subscription operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Polygonzkevm *PolygonzkevmFilterer) WatchActivateForceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmActivateForceBatches) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "ActivateForceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmActivateForceBatches)
				if err := _Polygonzkevm.contract.UnpackLog(event, "ActivateForceBatches", log); err != nil {
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

// ParseActivateForceBatches is a log parse operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Polygonzkevm *PolygonzkevmFilterer) ParseActivateForceBatches(log types.Log) (*PolygonzkevmActivateForceBatches, error) {
	event := new(PolygonzkevmActivateForceBatches)
	if err := _Polygonzkevm.contract.UnpackLog(event, "ActivateForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmConsolidatePendingStateIterator is returned from FilterConsolidatePendingState and is used to iterate over the raw logs and unpacked data for ConsolidatePendingState events raised by the Polygonzkevm contract.
type PolygonzkevmConsolidatePendingStateIterator struct {
	Event *PolygonzkevmConsolidatePendingState // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmConsolidatePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmConsolidatePendingState)
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
		it.Event = new(PolygonzkevmConsolidatePendingState)
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
func (it *PolygonzkevmConsolidatePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmConsolidatePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmConsolidatePendingState represents a ConsolidatePendingState event raised by the Polygonzkevm contract.
type PolygonzkevmConsolidatePendingState struct {
	NumBatch        uint64
	StateRoot       [32]byte
	PendingStateNum uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterConsolidatePendingState is a free log retrieval operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterConsolidatePendingState(opts *bind.FilterOpts, numBatch []uint64, pendingStateNum []uint64) (*PolygonzkevmConsolidatePendingStateIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var pendingStateNumRule []interface{}
	for _, pendingStateNumItem := range pendingStateNum {
		pendingStateNumRule = append(pendingStateNumRule, pendingStateNumItem)
	}

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "ConsolidatePendingState", numBatchRule, pendingStateNumRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmConsolidatePendingStateIterator{contract: _Polygonzkevm.contract, event: "ConsolidatePendingState", logs: logs, sub: sub}, nil
}

// WatchConsolidatePendingState is a free log subscription operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchConsolidatePendingState(opts *bind.WatchOpts, sink chan<- *PolygonzkevmConsolidatePendingState, numBatch []uint64, pendingStateNum []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var pendingStateNumRule []interface{}
	for _, pendingStateNumItem := range pendingStateNum {
		pendingStateNumRule = append(pendingStateNumRule, pendingStateNumItem)
	}

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "ConsolidatePendingState", numBatchRule, pendingStateNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmConsolidatePendingState)
				if err := _Polygonzkevm.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
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

// ParseConsolidatePendingState is a log parse operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseConsolidatePendingState(log types.Log) (*PolygonzkevmConsolidatePendingState, error) {
	event := new(PolygonzkevmConsolidatePendingState)
	if err := _Polygonzkevm.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Polygonzkevm contract.
type PolygonzkevmEmergencyStateActivatedIterator struct {
	Event *PolygonzkevmEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmEmergencyStateActivated)
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
		it.Event = new(PolygonzkevmEmergencyStateActivated)
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
func (it *PolygonzkevmEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmEmergencyStateActivated represents a EmergencyStateActivated event raised by the Polygonzkevm contract.
type PolygonzkevmEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonzkevm *PolygonzkevmFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*PolygonzkevmEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmEmergencyStateActivatedIterator{contract: _Polygonzkevm.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonzkevm *PolygonzkevmFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *PolygonzkevmEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmEmergencyStateActivated)
				if err := _Polygonzkevm.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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

// ParseEmergencyStateActivated is a log parse operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonzkevm *PolygonzkevmFilterer) ParseEmergencyStateActivated(log types.Log) (*PolygonzkevmEmergencyStateActivated, error) {
	event := new(PolygonzkevmEmergencyStateActivated)
	if err := _Polygonzkevm.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Polygonzkevm contract.
type PolygonzkevmEmergencyStateDeactivatedIterator struct {
	Event *PolygonzkevmEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmEmergencyStateDeactivated)
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
		it.Event = new(PolygonzkevmEmergencyStateDeactivated)
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
func (it *PolygonzkevmEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Polygonzkevm contract.
type PolygonzkevmEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonzkevm *PolygonzkevmFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*PolygonzkevmEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmEmergencyStateDeactivatedIterator{contract: _Polygonzkevm.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonzkevm *PolygonzkevmFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *PolygonzkevmEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmEmergencyStateDeactivated)
				if err := _Polygonzkevm.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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

// ParseEmergencyStateDeactivated is a log parse operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonzkevm *PolygonzkevmFilterer) ParseEmergencyStateDeactivated(log types.Log) (*PolygonzkevmEmergencyStateDeactivated, error) {
	event := new(PolygonzkevmEmergencyStateDeactivated)
	if err := _Polygonzkevm.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Polygonzkevm contract.
type PolygonzkevmForceBatchIterator struct {
	Event *PolygonzkevmForceBatch // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmForceBatch)
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
		it.Event = new(PolygonzkevmForceBatch)
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
func (it *PolygonzkevmForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmForceBatch represents a ForceBatch event raised by the Polygonzkevm contract.
type PolygonzkevmForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*PolygonzkevmForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmForceBatchIterator{contract: _Polygonzkevm.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *PolygonzkevmForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmForceBatch)
				if err := _Polygonzkevm.contract.UnpackLog(event, "ForceBatch", log); err != nil {
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

// ParseForceBatch is a log parse operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseForceBatch(log types.Log) (*PolygonzkevmForceBatch, error) {
	event := new(PolygonzkevmForceBatch)
	if err := _Polygonzkevm.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonzkevm contract.
type PolygonzkevmInitializedIterator struct {
	Event *PolygonzkevmInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmInitialized)
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
		it.Event = new(PolygonzkevmInitialized)
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
func (it *PolygonzkevmInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmInitialized represents a Initialized event raised by the Polygonzkevm contract.
type PolygonzkevmInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonzkevmInitializedIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmInitializedIterator{contract: _Polygonzkevm.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonzkevmInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmInitialized)
				if err := _Polygonzkevm.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseInitialized(log types.Log) (*PolygonzkevmInitialized, error) {
	event := new(PolygonzkevmInitialized)
	if err := _Polygonzkevm.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmOverridePendingStateIterator is returned from FilterOverridePendingState and is used to iterate over the raw logs and unpacked data for OverridePendingState events raised by the Polygonzkevm contract.
type PolygonzkevmOverridePendingStateIterator struct {
	Event *PolygonzkevmOverridePendingState // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmOverridePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmOverridePendingState)
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
		it.Event = new(PolygonzkevmOverridePendingState)
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
func (it *PolygonzkevmOverridePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmOverridePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmOverridePendingState represents a OverridePendingState event raised by the Polygonzkevm contract.
type PolygonzkevmOverridePendingState struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOverridePendingState is a free log retrieval operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterOverridePendingState(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonzkevmOverridePendingStateIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "OverridePendingState", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmOverridePendingStateIterator{contract: _Polygonzkevm.contract, event: "OverridePendingState", logs: logs, sub: sub}, nil
}

// WatchOverridePendingState is a free log subscription operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchOverridePendingState(opts *bind.WatchOpts, sink chan<- *PolygonzkevmOverridePendingState, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "OverridePendingState", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmOverridePendingState)
				if err := _Polygonzkevm.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
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

// ParseOverridePendingState is a log parse operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseOverridePendingState(log types.Log) (*PolygonzkevmOverridePendingState, error) {
	event := new(PolygonzkevmOverridePendingState)
	if err := _Polygonzkevm.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Polygonzkevm contract.
type PolygonzkevmOwnershipTransferredIterator struct {
	Event *PolygonzkevmOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmOwnershipTransferred)
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
		it.Event = new(PolygonzkevmOwnershipTransferred)
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
func (it *PolygonzkevmOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmOwnershipTransferred represents a OwnershipTransferred event raised by the Polygonzkevm contract.
type PolygonzkevmOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PolygonzkevmOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmOwnershipTransferredIterator{contract: _Polygonzkevm.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PolygonzkevmOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmOwnershipTransferred)
				if err := _Polygonzkevm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Polygonzkevm *PolygonzkevmFilterer) ParseOwnershipTransferred(log types.Log) (*PolygonzkevmOwnershipTransferred, error) {
	event := new(PolygonzkevmOwnershipTransferred)
	if err := _Polygonzkevm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmProveNonDeterministicPendingStateIterator is returned from FilterProveNonDeterministicPendingState and is used to iterate over the raw logs and unpacked data for ProveNonDeterministicPendingState events raised by the Polygonzkevm contract.
type PolygonzkevmProveNonDeterministicPendingStateIterator struct {
	Event *PolygonzkevmProveNonDeterministicPendingState // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmProveNonDeterministicPendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmProveNonDeterministicPendingState)
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
		it.Event = new(PolygonzkevmProveNonDeterministicPendingState)
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
func (it *PolygonzkevmProveNonDeterministicPendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmProveNonDeterministicPendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmProveNonDeterministicPendingState represents a ProveNonDeterministicPendingState event raised by the Polygonzkevm contract.
type PolygonzkevmProveNonDeterministicPendingState struct {
	StoredStateRoot [32]byte
	ProvedStateRoot [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProveNonDeterministicPendingState is a free log retrieval operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterProveNonDeterministicPendingState(opts *bind.FilterOpts) (*PolygonzkevmProveNonDeterministicPendingStateIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmProveNonDeterministicPendingStateIterator{contract: _Polygonzkevm.contract, event: "ProveNonDeterministicPendingState", logs: logs, sub: sub}, nil
}

// WatchProveNonDeterministicPendingState is a free log subscription operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchProveNonDeterministicPendingState(opts *bind.WatchOpts, sink chan<- *PolygonzkevmProveNonDeterministicPendingState) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmProveNonDeterministicPendingState)
				if err := _Polygonzkevm.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
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

// ParseProveNonDeterministicPendingState is a log parse operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseProveNonDeterministicPendingState(log types.Log) (*PolygonzkevmProveNonDeterministicPendingState, error) {
	event := new(PolygonzkevmProveNonDeterministicPendingState)
	if err := _Polygonzkevm.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Polygonzkevm contract.
type PolygonzkevmSequenceBatchesIterator struct {
	Event *PolygonzkevmSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSequenceBatches)
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
		it.Event = new(PolygonzkevmSequenceBatches)
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
func (it *PolygonzkevmSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSequenceBatches represents a SequenceBatches event raised by the Polygonzkevm contract.
type PolygonzkevmSequenceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonzkevmSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSequenceBatchesIterator{contract: _Polygonzkevm.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSequenceBatches)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
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

// ParseSequenceBatches is a log parse operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSequenceBatches(log types.Log) (*PolygonzkevmSequenceBatches, error) {
	event := new(PolygonzkevmSequenceBatches)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Polygonzkevm contract.
type PolygonzkevmSequenceForceBatchesIterator struct {
	Event *PolygonzkevmSequenceForceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSequenceForceBatches)
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
		it.Event = new(PolygonzkevmSequenceForceBatches)
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
func (it *PolygonzkevmSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSequenceForceBatches represents a SequenceForceBatches event raised by the Polygonzkevm contract.
type PolygonzkevmSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonzkevmSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSequenceForceBatchesIterator{contract: _Polygonzkevm.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSequenceForceBatches)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
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

// ParseSequenceForceBatches is a log parse operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSequenceForceBatches(log types.Log) (*PolygonzkevmSequenceForceBatches, error) {
	event := new(PolygonzkevmSequenceForceBatches)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Polygonzkevm contract.
type PolygonzkevmSetForceBatchTimeoutIterator struct {
	Event *PolygonzkevmSetForceBatchTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSetForceBatchTimeout)
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
		it.Event = new(PolygonzkevmSetForceBatchTimeout)
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
func (it *PolygonzkevmSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Polygonzkevm contract.
type PolygonzkevmSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*PolygonzkevmSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSetForceBatchTimeoutIterator{contract: _Polygonzkevm.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSetForceBatchTimeout)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
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

// ParseSetForceBatchTimeout is a log parse operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSetForceBatchTimeout(log types.Log) (*PolygonzkevmSetForceBatchTimeout, error) {
	event := new(PolygonzkevmSetForceBatchTimeout)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSetMultiplierBatchFeeIterator is returned from FilterSetMultiplierBatchFee and is used to iterate over the raw logs and unpacked data for SetMultiplierBatchFee events raised by the Polygonzkevm contract.
type PolygonzkevmSetMultiplierBatchFeeIterator struct {
	Event *PolygonzkevmSetMultiplierBatchFee // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmSetMultiplierBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSetMultiplierBatchFee)
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
		it.Event = new(PolygonzkevmSetMultiplierBatchFee)
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
func (it *PolygonzkevmSetMultiplierBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSetMultiplierBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSetMultiplierBatchFee represents a SetMultiplierBatchFee event raised by the Polygonzkevm contract.
type PolygonzkevmSetMultiplierBatchFee struct {
	NewMultiplierBatchFee uint16
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetMultiplierBatchFee is a free log retrieval operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSetMultiplierBatchFee(opts *bind.FilterOpts) (*PolygonzkevmSetMultiplierBatchFeeIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSetMultiplierBatchFeeIterator{contract: _Polygonzkevm.contract, event: "SetMultiplierBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetMultiplierBatchFee is a free log subscription operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSetMultiplierBatchFee(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSetMultiplierBatchFee) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSetMultiplierBatchFee)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
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

// ParseSetMultiplierBatchFee is a log parse operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSetMultiplierBatchFee(log types.Log) (*PolygonzkevmSetMultiplierBatchFee, error) {
	event := new(PolygonzkevmSetMultiplierBatchFee)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSetPendingStateTimeoutIterator is returned from FilterSetPendingStateTimeout and is used to iterate over the raw logs and unpacked data for SetPendingStateTimeout events raised by the Polygonzkevm contract.
type PolygonzkevmSetPendingStateTimeoutIterator struct {
	Event *PolygonzkevmSetPendingStateTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmSetPendingStateTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSetPendingStateTimeout)
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
		it.Event = new(PolygonzkevmSetPendingStateTimeout)
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
func (it *PolygonzkevmSetPendingStateTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSetPendingStateTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSetPendingStateTimeout represents a SetPendingStateTimeout event raised by the Polygonzkevm contract.
type PolygonzkevmSetPendingStateTimeout struct {
	NewPendingStateTimeout uint64
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetPendingStateTimeout is a free log retrieval operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSetPendingStateTimeout(opts *bind.FilterOpts) (*PolygonzkevmSetPendingStateTimeoutIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSetPendingStateTimeoutIterator{contract: _Polygonzkevm.contract, event: "SetPendingStateTimeout", logs: logs, sub: sub}, nil
}

// WatchSetPendingStateTimeout is a free log subscription operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSetPendingStateTimeout(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSetPendingStateTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSetPendingStateTimeout)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
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

// ParseSetPendingStateTimeout is a log parse operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSetPendingStateTimeout(log types.Log) (*PolygonzkevmSetPendingStateTimeout, error) {
	event := new(PolygonzkevmSetPendingStateTimeout)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSetTrustedAggregatorIterator is returned from FilterSetTrustedAggregator and is used to iterate over the raw logs and unpacked data for SetTrustedAggregator events raised by the Polygonzkevm contract.
type PolygonzkevmSetTrustedAggregatorIterator struct {
	Event *PolygonzkevmSetTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmSetTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSetTrustedAggregator)
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
		it.Event = new(PolygonzkevmSetTrustedAggregator)
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
func (it *PolygonzkevmSetTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSetTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSetTrustedAggregator represents a SetTrustedAggregator event raised by the Polygonzkevm contract.
type PolygonzkevmSetTrustedAggregator struct {
	NewTrustedAggregator common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregator is a free log retrieval operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSetTrustedAggregator(opts *bind.FilterOpts) (*PolygonzkevmSetTrustedAggregatorIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSetTrustedAggregatorIterator{contract: _Polygonzkevm.contract, event: "SetTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregator is a free log subscription operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSetTrustedAggregator(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSetTrustedAggregator) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSetTrustedAggregator)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
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

// ParseSetTrustedAggregator is a log parse operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSetTrustedAggregator(log types.Log) (*PolygonzkevmSetTrustedAggregator, error) {
	event := new(PolygonzkevmSetTrustedAggregator)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSetTrustedAggregatorTimeoutIterator is returned from FilterSetTrustedAggregatorTimeout and is used to iterate over the raw logs and unpacked data for SetTrustedAggregatorTimeout events raised by the Polygonzkevm contract.
type PolygonzkevmSetTrustedAggregatorTimeoutIterator struct {
	Event *PolygonzkevmSetTrustedAggregatorTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmSetTrustedAggregatorTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSetTrustedAggregatorTimeout)
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
		it.Event = new(PolygonzkevmSetTrustedAggregatorTimeout)
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
func (it *PolygonzkevmSetTrustedAggregatorTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSetTrustedAggregatorTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSetTrustedAggregatorTimeout represents a SetTrustedAggregatorTimeout event raised by the Polygonzkevm contract.
type PolygonzkevmSetTrustedAggregatorTimeout struct {
	NewTrustedAggregatorTimeout uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregatorTimeout is a free log retrieval operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSetTrustedAggregatorTimeout(opts *bind.FilterOpts) (*PolygonzkevmSetTrustedAggregatorTimeoutIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSetTrustedAggregatorTimeoutIterator{contract: _Polygonzkevm.contract, event: "SetTrustedAggregatorTimeout", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregatorTimeout is a free log subscription operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSetTrustedAggregatorTimeout(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSetTrustedAggregatorTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSetTrustedAggregatorTimeout)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
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

// ParseSetTrustedAggregatorTimeout is a log parse operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSetTrustedAggregatorTimeout(log types.Log) (*PolygonzkevmSetTrustedAggregatorTimeout, error) {
	event := new(PolygonzkevmSetTrustedAggregatorTimeout)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Polygonzkevm contract.
type PolygonzkevmSetTrustedSequencerIterator struct {
	Event *PolygonzkevmSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSetTrustedSequencer)
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
		it.Event = new(PolygonzkevmSetTrustedSequencer)
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
func (it *PolygonzkevmSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSetTrustedSequencer represents a SetTrustedSequencer event raised by the Polygonzkevm contract.
type PolygonzkevmSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*PolygonzkevmSetTrustedSequencerIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSetTrustedSequencerIterator{contract: _Polygonzkevm.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSetTrustedSequencer)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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

// ParseSetTrustedSequencer is a log parse operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSetTrustedSequencer(log types.Log) (*PolygonzkevmSetTrustedSequencer, error) {
	event := new(PolygonzkevmSetTrustedSequencer)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Polygonzkevm contract.
type PolygonzkevmSetTrustedSequencerURLIterator struct {
	Event *PolygonzkevmSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSetTrustedSequencerURL)
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
		it.Event = new(PolygonzkevmSetTrustedSequencerURL)
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
func (it *PolygonzkevmSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Polygonzkevm contract.
type PolygonzkevmSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*PolygonzkevmSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSetTrustedSequencerURLIterator{contract: _Polygonzkevm.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSetTrustedSequencerURL)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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

// ParseSetTrustedSequencerURL is a log parse operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSetTrustedSequencerURL(log types.Log) (*PolygonzkevmSetTrustedSequencerURL, error) {
	event := new(PolygonzkevmSetTrustedSequencerURL)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmSetVerifyBatchTimeTargetIterator is returned from FilterSetVerifyBatchTimeTarget and is used to iterate over the raw logs and unpacked data for SetVerifyBatchTimeTarget events raised by the Polygonzkevm contract.
type PolygonzkevmSetVerifyBatchTimeTargetIterator struct {
	Event *PolygonzkevmSetVerifyBatchTimeTarget // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmSetVerifyBatchTimeTargetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmSetVerifyBatchTimeTarget)
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
		it.Event = new(PolygonzkevmSetVerifyBatchTimeTarget)
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
func (it *PolygonzkevmSetVerifyBatchTimeTargetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmSetVerifyBatchTimeTargetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmSetVerifyBatchTimeTarget represents a SetVerifyBatchTimeTarget event raised by the Polygonzkevm contract.
type PolygonzkevmSetVerifyBatchTimeTarget struct {
	NewVerifyBatchTimeTarget uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetVerifyBatchTimeTarget is a free log retrieval operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterSetVerifyBatchTimeTarget(opts *bind.FilterOpts) (*PolygonzkevmSetVerifyBatchTimeTargetIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmSetVerifyBatchTimeTargetIterator{contract: _Polygonzkevm.contract, event: "SetVerifyBatchTimeTarget", logs: logs, sub: sub}, nil
}

// WatchSetVerifyBatchTimeTarget is a free log subscription operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchSetVerifyBatchTimeTarget(opts *bind.WatchOpts, sink chan<- *PolygonzkevmSetVerifyBatchTimeTarget) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmSetVerifyBatchTimeTarget)
				if err := _Polygonzkevm.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
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

// ParseSetVerifyBatchTimeTarget is a log parse operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseSetVerifyBatchTimeTarget(log types.Log) (*PolygonzkevmSetVerifyBatchTimeTarget, error) {
	event := new(PolygonzkevmSetVerifyBatchTimeTarget)
	if err := _Polygonzkevm.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Polygonzkevm contract.
type PolygonzkevmTransferAdminRoleIterator struct {
	Event *PolygonzkevmTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmTransferAdminRole)
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
		it.Event = new(PolygonzkevmTransferAdminRole)
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
func (it *PolygonzkevmTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmTransferAdminRole represents a TransferAdminRole event raised by the Polygonzkevm contract.
type PolygonzkevmTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*PolygonzkevmTransferAdminRoleIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmTransferAdminRoleIterator{contract: _Polygonzkevm.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonzkevmTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmTransferAdminRole)
				if err := _Polygonzkevm.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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

// ParseTransferAdminRole is a log parse operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseTransferAdminRole(log types.Log) (*PolygonzkevmTransferAdminRole, error) {
	event := new(PolygonzkevmTransferAdminRole)
	if err := _Polygonzkevm.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmUpdateZkEVMVersionIterator is returned from FilterUpdateZkEVMVersion and is used to iterate over the raw logs and unpacked data for UpdateZkEVMVersion events raised by the Polygonzkevm contract.
type PolygonzkevmUpdateZkEVMVersionIterator struct {
	Event *PolygonzkevmUpdateZkEVMVersion // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmUpdateZkEVMVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmUpdateZkEVMVersion)
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
		it.Event = new(PolygonzkevmUpdateZkEVMVersion)
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
func (it *PolygonzkevmUpdateZkEVMVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmUpdateZkEVMVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmUpdateZkEVMVersion represents a UpdateZkEVMVersion event raised by the Polygonzkevm contract.
type PolygonzkevmUpdateZkEVMVersion struct {
	NumBatch uint64
	ForkID   uint64
	Version  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateZkEVMVersion is a free log retrieval operation binding the contract event 0xed7be53c9f1a96a481223b15568a5b1a475e01a74b347d6ca187c8bf0c078cd6.
//
// Solidity: event UpdateZkEVMVersion(uint64 numBatch, uint64 forkID, string version)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterUpdateZkEVMVersion(opts *bind.FilterOpts) (*PolygonzkevmUpdateZkEVMVersionIterator, error) {

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "UpdateZkEVMVersion")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmUpdateZkEVMVersionIterator{contract: _Polygonzkevm.contract, event: "UpdateZkEVMVersion", logs: logs, sub: sub}, nil
}

// WatchUpdateZkEVMVersion is a free log subscription operation binding the contract event 0xed7be53c9f1a96a481223b15568a5b1a475e01a74b347d6ca187c8bf0c078cd6.
//
// Solidity: event UpdateZkEVMVersion(uint64 numBatch, uint64 forkID, string version)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchUpdateZkEVMVersion(opts *bind.WatchOpts, sink chan<- *PolygonzkevmUpdateZkEVMVersion) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "UpdateZkEVMVersion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmUpdateZkEVMVersion)
				if err := _Polygonzkevm.contract.UnpackLog(event, "UpdateZkEVMVersion", log); err != nil {
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

// ParseUpdateZkEVMVersion is a log parse operation binding the contract event 0xed7be53c9f1a96a481223b15568a5b1a475e01a74b347d6ca187c8bf0c078cd6.
//
// Solidity: event UpdateZkEVMVersion(uint64 numBatch, uint64 forkID, string version)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseUpdateZkEVMVersion(log types.Log) (*PolygonzkevmUpdateZkEVMVersion, error) {
	event := new(PolygonzkevmUpdateZkEVMVersion)
	if err := _Polygonzkevm.contract.UnpackLog(event, "UpdateZkEVMVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonzkevm contract.
type PolygonzkevmVerifyBatchesIterator struct {
	Event *PolygonzkevmVerifyBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmVerifyBatches)
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
		it.Event = new(PolygonzkevmVerifyBatches)
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
func (it *PolygonzkevmVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmVerifyBatches represents a VerifyBatches event raised by the Polygonzkevm contract.
type PolygonzkevmVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonzkevmVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmVerifyBatchesIterator{contract: _Polygonzkevm.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmVerifyBatches)
				if err := _Polygonzkevm.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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

// ParseVerifyBatches is a log parse operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseVerifyBatches(log types.Log) (*PolygonzkevmVerifyBatches, error) {
	event := new(PolygonzkevmVerifyBatches)
	if err := _Polygonzkevm.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmVerifyBatchesTrustedAggregatorIterator is returned from FilterVerifyBatchesTrustedAggregator and is used to iterate over the raw logs and unpacked data for VerifyBatchesTrustedAggregator events raised by the Polygonzkevm contract.
type PolygonzkevmVerifyBatchesTrustedAggregatorIterator struct {
	Event *PolygonzkevmVerifyBatchesTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmVerifyBatchesTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmVerifyBatchesTrustedAggregator)
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
		it.Event = new(PolygonzkevmVerifyBatchesTrustedAggregator)
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
func (it *PolygonzkevmVerifyBatchesTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmVerifyBatchesTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmVerifyBatchesTrustedAggregator represents a VerifyBatchesTrustedAggregator event raised by the Polygonzkevm contract.
type PolygonzkevmVerifyBatchesTrustedAggregator struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatchesTrustedAggregator is a free log retrieval operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevm *PolygonzkevmFilterer) FilterVerifyBatchesTrustedAggregator(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonzkevmVerifyBatchesTrustedAggregatorIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevm.contract.FilterLogs(opts, "VerifyBatchesTrustedAggregator", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmVerifyBatchesTrustedAggregatorIterator{contract: _Polygonzkevm.contract, event: "VerifyBatchesTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchVerifyBatchesTrustedAggregator is a free log subscription operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevm *PolygonzkevmFilterer) WatchVerifyBatchesTrustedAggregator(opts *bind.WatchOpts, sink chan<- *PolygonzkevmVerifyBatchesTrustedAggregator, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevm.contract.WatchLogs(opts, "VerifyBatchesTrustedAggregator", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmVerifyBatchesTrustedAggregator)
				if err := _Polygonzkevm.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
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

// ParseVerifyBatchesTrustedAggregator is a log parse operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevm *PolygonzkevmFilterer) ParseVerifyBatchesTrustedAggregator(log types.Log) (*PolygonzkevmVerifyBatchesTrustedAggregator, error) {
	event := new(PolygonzkevmVerifyBatchesTrustedAggregator)
	if err := _Polygonzkevm.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
