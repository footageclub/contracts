// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package clublock

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
)

// ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo is an auto generated low-level Go binding around an user-defined struct.
type ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo struct {
	RoyaltyAddress common.Address
	RoyaltyBps     *big.Int
}

// MintParams is an auto generated low-level Go binding around an user-defined struct.
type MintParams struct {
	MintPrice                *big.Int
	MaxTotalMintableByWallet *big.Int
	StartTime                *big.Int
	EndTime                  *big.Int
	DropStageIndex           *big.Int
	MaxTokenSupplyForStage   *big.Int
	FeeBps                   *big.Int
	RestrictFeeRecipients    bool
}

// PublicDrop is an auto generated low-level Go binding around an user-defined struct.
type PublicDrop struct {
	MintPrice                *big.Int
	StartTime                *big.Int
	EndTime                  *big.Int
	MaxTotalMintableByWallet uint16
	FeeBps                   uint16
	RestrictFeeRecipients    bool
}

// SignedMintValidationParams is an auto generated low-level Go binding around an user-defined struct.
type SignedMintValidationParams struct {
	MinMintPrice                *big.Int
	MaxMaxTotalMintableByWallet *big.Int
	MinStartTime                *big.Int
	MaxEndTime                  *big.Int
	MaxMaxTokenSupplyForStage   *big.Int
	MinFeeBps                   uint16
	MaxFeeBps                   uint16
}

// TokenGatedDropStage is an auto generated low-level Go binding around an user-defined struct.
type TokenGatedDropStage struct {
	MintPrice                *big.Int
	MaxTotalMintableByWallet uint16
	StartTime                *big.Int
	EndTime                  *big.Int
	DropStageIndex           uint8
	MaxTokenSupplyForStage   uint32
	FeeBps                   uint16
	RestrictFeeRecipients    bool
}

// ClublockMetaData contains all meta data concerning the Clublock contract.
var ClublockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BurnIncorrectSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"CannotExceedMaxSupplyOfUint64\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CreatorPayoutAddressCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateFeeRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicatePayer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientNotPresent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"IncorrectPayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAssetId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeBps\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"basisPoints\",\"type\":\"uint256\"}],\"name\":\"InvalidRoyaltyBasisPoints\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recoveredSigner\",\"type\":\"address\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedEndTime\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumOrMaximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedFeeBps\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedMaxTokenSupplyForStage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedMaxTotalMintableByWallet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedMintPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintQuantityCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allowed\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxMintedPerWallet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxSupply\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxTokenSupplyForStage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"}],\"name\":\"NotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotNextOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"OnlyINonFungibleSeaDropToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerNotPresent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProvenanceHashCannotBeSetAfterMintStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoyaltyAddressCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignedMintsMustRestrictFeeRecipients\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerNotPresent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedDropAllowedNftTokenCannotBeDropToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedDropAllowedNftTokenCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedDropStageNotPresent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowedNftTokenId\",\"type\":\"uint256\"}],\"name\":\"TokenGatedNotTokenOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowedNftTokenId\",\"type\":\"uint256\"}],\"name\":\"TokenGatedTokenIdAlreadyRedeemed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"assets\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"trace\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousMerkleRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newMerkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"publicKeyURI\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"name\":\"AllowListUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AllowedFeeRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ClubLockDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"}],\"name\":\"ContractURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPayoutAddress\",\"type\":\"address\"}],\"name\":\"CreatorPayoutAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newDropURI\",\"type\":\"string\"}],\"name\":\"DropURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dropStageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structMintParams\",\"name\":\"mintParams\",\"type\":\"tuple\"}],\"name\":\"MintParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"PayerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPotentialAdministrator\",\"type\":\"address\"}],\"name\":\"PotentialOwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newHash\",\"type\":\"bytes32\"}],\"name\":\"ProvenanceHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"}],\"name\":\"PublicDropUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"name\":\"RoyaltyInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantityMinted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unitMintPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dropStageIndex\",\"type\":\"uint256\"}],\"name\":\"SeaDropMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structSignedMintValidationParams\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple\"}],\"name\":\"SignedMintValidationParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTokenGatedDropStage\",\"name\":\"dropStage\",\"type\":\"tuple\"}],\"name\":\"TokenGatedDropStageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"assets\",\"type\":\"string[]\"}],\"name\":\"mintClublockUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payoutAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dropStageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structMintParams\",\"name\":\"mintParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"royaltyAddress\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"royaltyBps\",\"type\":\"uint96\"}],\"internalType\":\"structISeaDropTokenContractMetadataUpgradeable.RoyaltyInfo\",\"name\":\"royalty\",\"type\":\"tuple\"}],\"name\":\"config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"}],\"name\":\"emitBatchMetadataUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBurnAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMintParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dropStageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structMintParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"getMintStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minterNumMinted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"assets\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mintClublock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provenanceHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBurnAddress\",\"type\":\"address\"}],\"name\":\"setBurnAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newProvenanceHash\",\"type\":\"bytes32\"}],\"name\":\"setProvenanceHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"royaltyAddress\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"royaltyBps\",\"type\":\"uint96\"}],\"internalType\":\"structISeaDropTokenContractMetadataUpgradeable.RoyaltyInfo\",\"name\":\"newInfo\",\"type\":\"tuple\"}],\"name\":\"setRoyaltyInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPotentialOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payoutAddress\",\"type\":\"address\"}],\"name\":\"updateCreatorPayoutAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dropStageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structMintParams\",\"name\":\"mintParams\",\"type\":\"tuple\"}],\"name\":\"updateMintParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"updateSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50613fe2806100206000396000f3fe60806040526004361061025c5760003560e01c806370a0823111610144578063a9b34c3b116100b6578063c87b56dd1161007a578063c87b56dd14610746578063d1e7f79c14610766578063d5abeb01146107de578063e8a3d485146107f3578063e985e9c514610808578063f2fde38b1461082857600080fd5b8063a9b34c3b14610694578063ad2f852a146106a7578063b7c5e064146106d2578063b88d4fde146106f2578063c6ab67a31461071257600080fd5b80638ff27867116101085780638ff27867146105df578063938e3d7b146105ff57806395d89b411461061f578063a22cb46514610634578063a483011414610654578063a7ecd37e1461067457600080fd5b806370a0823114610545578063715018a61461056557806379ba50971461057a578063840e15d41461058f5780638da5cb5b146105ca57600080fd5b80632a55205a116101dd57806344dae42c116101a157806344dae42c146104905780634b0e7216146104b057806355f804b3146104d05780636352211e146104f05780636c0360eb146105105780636f8b44b01461052557600080fd5b80632a55205a146103ca57806338b39d291461040957806342260b5d1461041e57806342842e0e1461045057806342966c681461047057600080fd5b8063099b6bfa11610224578063099b6bfa1461033257806312738db81461035257806318160ddd1461037257806323452b9c1461039557806323b872dd146103aa57600080fd5b806301ffc9a71461026157806306fdde0314610296578063077f224a146102b8578063081812fc146102da578063095ea7b314610312575b600080fd5b34801561026d57600080fd5b5061028161027c36600461327f565b610848565b60405190151581526020015b60405180910390f35b3480156102a257600080fd5b506102ab610873565b60405161028d91906132ec565b3480156102c457600080fd5b506102d86102d33660046133d1565b61090e565b005b3480156102e657600080fd5b506102fa6102f5366004613448565b610b63565b6040516001600160a01b03909116815260200161028d565b34801561031e57600080fd5b506102d861032d366004613461565b610bb0565b34801561033e57600080fd5b506102d861034d366004613448565b610bc9565b34801561035e57600080fd5b506102d861036d36600461348d565b610c5f565b34801561037e57600080fd5b50610387610ce6565b60405190815260200161028d565b3480156103a157600080fd5b506102d8610d06565b3480156103b657600080fd5b506102d86103c53660046134aa565b610d61565b3480156103d657600080fd5b506103ea6103e53660046134eb565b610d86565b604080516001600160a01b03909316835260208301919091520161028d565b34801561041557600080fd5b506102fa610dd9565b34801561042a57600080fd5b50600080516020613f6d83398151915254600160a01b90046001600160601b0316610387565b34801561045c57600080fd5b506102d861046b3660046134aa565b610df2565b34801561047c57600080fd5b506102d861048b366004613448565b610e17565b34801561049c57600080fd5b506102d86104ab36600461351f565b610e54565b3480156104bc57600080fd5b506102d86104cb36600461348d565b610f72565b3480156104dc57600080fd5b506102d86104eb36600461357c565b610fa4565b3480156104fc57600080fd5b506102fa61050b366004613448565b611035565b34801561051c57600080fd5b506102ab611040565b34801561053157600080fd5b506102d8610540366004613448565b61104f565b34801561055157600080fd5b5061038761056036600461348d565b6110c3565b34801561057157600080fd5b506102d861112b565b34801561058657600080fd5b506102d861113f565b34801561059b57600080fd5b506105af6105aa36600461348d565b6111d2565b6040805193845260208401929092529082015260600161028d565b3480156105d657600080fd5b506102fa611204565b3480156105eb57600080fd5b506102d86105fa3660046135d0565b61122c565b34801561060b57600080fd5b506102d861061a36600461357c565b6112b6565b34801561062b57600080fd5b506102ab61131c565b34801561064057600080fd5b506102d861064f3660046135fb565b611334565b34801561066057600080fd5b506102d861066f3660046134eb565b611348565b34801561068057600080fd5b506102d861068f36600461348d565b611386565b6102d86106a2366004613634565b6113e6565b3480156106b357600080fd5b50600080516020613f6d833981519152546001600160a01b03166102fa565b3480156106de57600080fd5b506102d86106ed3660046136f0565b61170a565b3480156106fe57600080fd5b506102d861070d366004613743565b6118e3565b34801561071e57600080fd5b507fb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031c54610387565b34801561075257600080fd5b506102ab610761366004613448565b611910565b34801561077257600080fd5b5061077b611994565b60405161028d9190600061010082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015260c083015160c083015260e0830151151560e083015292915050565b3480156107ea57600080fd5b50610387611a4b565b3480156107ff57600080fd5b506102ab611a63565b34801561081457600080fd5b506102816108233660046137c2565b611a79565b34801561083457600080fd5b506102d861084336600461348d565b611ab6565b60006001600160e01b03198216639c15441560e01b148061086d575061086d82611b40565b92915050565b606061087d611b80565b600201805461088b906137f0565b80601f01602080910402602001604051908101604052809291908181526020018280546108b7906137f0565b80156109045780601f106108d957610100808354040283529160200191610904565b820191906000526020600020905b8154815290600101906020018083116108e757829003601f168201915b5050505050905090565b600054610100900460ff161580801561092e5750600054600160ff909116105b806109485750303b158015610948575060005460ff166001145b6109b05760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff1916600117905580156109d3576000805461ff0019166101001790555b600080516020613f8d83398151915254610100900460ff16610a0857600080516020613f8d8339815191525460ff1615610a0c565b303b155b610a7e5760405162461bcd60e51b815260206004820152603760248201527f455243373231415f5f496e697469616c697a61626c653a20636f6e747261637460448201527f20697320616c726561647920696e697469616c697a656400000000000000000060648201526084016109a7565b600080516020613f8d83398151915254610100900460ff16158015610aba57600080516020613f8d833981519152805461ffff19166101011790555b610ac48585611ba4565b610acc611c64565b610ad4611c8b565b610ade8585611cba565b610ae6611ce1565b610aee611d1c565b610af783611d62565b8015610b1657600080516020613f8d833981519152805461ff00191690555b508015610b5d576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b6000610b6e82611e00565b610b8b576040516333d1c03960e21b815260040160405180910390fd5b610b93611b80565b60009283526006016020525060409020546001600160a01b031690565b81610bba81611e49565b610bc48383611f02565b505050565b610bd1611fb0565b6000610bdb611ff6565b1115610bfa5760405163e03264af60e01b815260040160405180910390fd5b7fb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031c80549082905560408051828152602081018490527f7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c91015b60405180910390a15050565b610c67612009565b6001600160a01b038116610c8e57604051633f00976960e01b815260040160405180910390fd5b80610c97612053565b600f0180546001600160a01b0319166001600160a01b039283161790556040519082169033907f0c69f21751e800ea5960436c9a94370c7adbf54c733a20a025293fbbe8f1625290600090a350565b60006001610cf2611b80565b60010154610cfe611b80565b540303919050565b610d0e612009565b600080516020613eed83398151915280546001600160a01b0319169055604080516000815290517f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da9181900360200190a1565b826001600160a01b0381163314610d7b57610d7b33611e49565b610b5d848484612077565b600080516020613f6d8339815191528054600091829161271090610dba90600160a01b90046001600160601b03168661383a565b610dc49190613851565b90546001600160a01b03169590945092505050565b6000610de3612053565b546001600160a01b0316919050565b826001600160a01b0381163314610e0c57610e0c33611e49565b610b5d84848461225c565b610e1f612053565b546001600160a01b03163314610e485760405163ea885d8360e01b815260040160405180910390fd5b610e5181612277565b50565b610e5c611fb0565b6000610e6b602083018361348d565b6001600160a01b031603610e9257604051631cc0baef60e01b815260040160405180910390fd5b612710610ea56040830160208401613888565b6001600160601b03161115610ee957610ec46040820160208301613888565b604051633cadbafb60e01b81526001600160601b0390911660048201526024016109a7565b80600080516020613f6d833981519152610f0382826138a5565b507ff21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d9050610f34602083018361348d565b610f446040840160208501613888565b604080516001600160a01b0390931683526001600160601b039091166020830152015b60405180910390a150565b610f7a612009565b80610f83612053565b80546001600160a01b0319166001600160a01b039290921691909117905550565b610fac611fb0565b7fb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031a610fd882848361392d565b50610fe1610ce6565b15611031577f6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c600180611012612282565b61101c91906139ec565b60408051928352602083019190915201610c53565b5050565b600061086d8261228c565b606061104a612320565b905090565b611057611fb0565b6001600160401b038111156110825760405163b43e913760e01b8152600481018290526024016109a7565b80600080516020613f4d833981519152556040518181527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c90602001610f67565b60006001600160a01b0382166110ec576040516323d3ad8160e21b815260040160405180910390fd5b6001600160401b036110fc611b80565b6005016000846001600160a01b03166001600160a01b0316815260200190815260200160002054169050919050565b611133612009565b61113d600061233f565b565b600080516020613eed833981519152546001600160a01b031633811461117857604051636b7584e760e11b815260040160405180910390fd5b600080516020613eed83398151915280546001600160a01b0319169055604080516000815290517f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da9181900360200190a1610e518161233f565b60008060006111e0846123b0565b92506111ea611ff6565b9150600080516020613f4d83398151915254929491935050565b60007ff73863e3917595a7e67829ed60f0c51bf14c7a0d0de47c7b02a00abd48488e4a610de3565b611234611fb0565b6127108160c00135111561126157604051631994fc9960e11b815260c082013560048201526024016109a7565b8061126a612053565b60040161127782826139ff565b505060405133907fe2b0a093e383d45f177604d69d4d2172f9e459de75908c02c9f9dd0ba8591bea906112ab908490613a67565b60405180910390a250565b6112be611fb0565b7fb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031b6112ea82848361392d565b507f905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac373788282604051610c53929190613af6565b6060611326611b80565b600301805461088b906137f0565b8161133e81611e49565b610bc483836123f5565b611350611fb0565b60408051838152602081018390527f6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c9101610c53565b61138e612009565b80611397612053565b60030180546001600160a01b0319166001600160a01b039283161790556040519082169033907f2d025324f0a785e8c12d0a0d91a9caa49df4ef20ff87e0df7213a1d4f3157beb90600090a350565b600080516020613f0d833981519152546001146114325760405162461bcd60e51b815260206004820152600a6024820152695245454e5452414e435960b01b60448201526064016109a7565b6002600080516020613f0d8339815191525560006001600160a01b03871661145a573361145c565b865b9050846000611469612053565b60408051610100810182526004830154815260058301546020820152600683015491810182905260078301546060820181905260088401546080830152600984015460a0830152600a84015460c0830152600b9093015460ff16151560e082015292506114d591612472565b6114e9838383602001518460a001516124b5565b6114f78282600001516125c6565b600061150e848b6115088b8d613b0a565b8a612603565b9050611518612053565b6000828152600c91909101602052604090205460ff161561154c5760405163900bb2c960e01b815260040160405180910390fd5b6001611556612053565b6000838152600c9190910160209081526040808320805460ff1916941515949094179093558251601f890182900482028101820190935287835290916115b991899089908190840183828082843760009201919091525086939250506126df9050565b90506115c3612053565b600301546001600160a01b038281169116146115fd57604051633615713d60e21b81526001600160a01b03821660048201526024016109a7565b505060005b8281101561167e57611612612053565b60010189898381811061162757611627613b8d565b90506020028101906116399190613ba3565b604051611647929190613be9565b908152602001604051809103902054600014611676576040516307d7e52d60e51b815260040160405180910390fd5b600101611602565b506116898383612703565b61169460008061271d565b886001600160a01b0316836001600160a01b03167fbeeb7c756bb80c977a1eeb91b3bdb1af5d9107550f8c929feee514b674722e9483600001518b8b6040516116df93929190613bf9565b60405180910390a35050506001611701600080516020613f0d83398151915290565b55505050505050565b611712612009565b6001600160a01b0384161561177b5760405163024e71b760e31b81526001600160a01b038516600482015230906312738db890602401600060405180830381600087803b15801561176257600080fd5b505af1158015611776573d6000803e3d6000fd5b505050505b82156117d3576040516306f8b44b60e41b8152600481018490523090636f8b44b090602401600060405180830381600087803b1580156117ba57600080fd5b505af11580156117ce573d6000803e3d6000fd5b505050505b60608201351515604083013515151760010361183f57604051638ff2786760e01b81523090638ff278679061180c908590600401613a67565b600060405180830381600087803b15801561182657600080fd5b505af115801561183a573d6000803e3d6000fd5b505050505b6118606118526040830160208401613888565b6001600160601b0316151590565b6118806000611872602085018561348d565b6001600160a01b0316141590565b17600103610b5d57604051631136b90b60e21b815230906344dae42c906118ab908490600401613ca0565b600060405180830381600087803b1580156118c557600080fd5b505af11580156118d9573d6000803e3d6000fd5b5050505050505050565b836001600160a01b03811633146118fd576118fd33611e49565b611909858585856127c5565b5050505050565b606061191b82611e00565b61193857604051630a14c4b560e41b815260040160405180910390fd5b6000611942612320565b90508051600003611962576040518060200160405280600081525061198d565b8061196c84612809565b60405160200161197d929190613cde565b6040516020818303038152906040525b9392505050565b6119de604051806101000160405280600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000151581525090565b6119e6612053565b604080516101008101825260048301548152600583015460208201526006830154918101919091526007820154606082015260088201546080820152600982015460a0820152600a82015460c0820152600b9091015460ff16151560e0820152919050565b6000600080516020613f4d8339815191525b54919050565b6060600080516020613f4d83398151915261087d565b6000611a83611b80565b6001600160a01b039384166000908152600791909101602090815260408083209490951682529290925250205460ff1690565b611abe612009565b6001600160a01b038116611ae557604051633a247dd760e11b815260040160405180910390fd5b600080516020613eed83398151915280546001600160a01b0319166001600160a01b0383169081179091556040519081527f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da90602001610f67565b60006001600160e01b0319821663152a902d60e11b1480611b715750632483248360e11b6001600160e01b03198316145b8061086d575061086d8261284d565b7f2569078dfb4b0305704d3008e7403993ae9601b85f7ae5e742de3de8f8011c4090565b600080516020613f8d83398151915254610100900460ff16611c255760405162461bcd60e51b815260206004820152603460248201527f455243373231415f5f496e697469616c697a61626c653a20636f6e7472616374604482015273206973206e6f7420696e697469616c697a696e6760601b60648201526084016109a7565b81611c2e611b80565b60020190611c3c9082613d0d565b5080611c46611b80565b60030190611c549082613d0d565b506001611c5f611b80565b555050565b600054610100900460ff1661113d5760405162461bcd60e51b81526004016109a790613dcc565b600054610100900460ff16611cb25760405162461bcd60e51b81526004016109a790613dcc565b61113d61289b565b600054610100900460ff166110315760405162461bcd60e51b81526004016109a790613dcc565b600054610100900460ff16611d085760405162461bcd60e51b81526004016109a790613dcc565b6001600080516020613f0d83398151915255565b600054610100900460ff16611d435760405162461bcd60e51b81526004016109a790613dcc565b61113d733cc6cdda760b79bafa08df41ecfa224f810dceb660016128c3565b600054610100900460ff16611d895760405162461bcd60e51b81526004016109a790613dcc565b46611d92612053565b600d0155611d9e612a62565b611da6612053565b600e015580611db3612053565b60030180546001600160a01b0319166001600160a01b03929092169190911790556040517f6d9aef9de83ac46d2f116c7876c158dd6c92761f13ce1fe88c97d3c3a3f1347d90600090a150565b600081600111158015611e1a5750611e16611b80565b5482105b801561086d5750600160e01b611e2e611b80565b60008481526004919091016020526040902054161592915050565b6daaeb6d7670e522a718067333cd4e3b15610e5157604051633185c44d60e21b81523060048201526001600160a01b03821660248201526daaeb6d7670e522a718067333cd4e9063c617113490604401602060405180830381865afa158015611eb6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611eda9190613e17565b610e5157604051633b79c77360e21b81526001600160a01b03821660048201526024016109a7565b6000611f0d82611035565b9050336001600160a01b03821614611f4657611f298133611a79565b611f46576040516367d9dca160e11b815260040160405180910390fd5b82611f4f611b80565b6000848152600691909101602052604080822080546001600160a01b0319166001600160a01b0394851617905551849286811692908516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259190a4505050565b303314611fd5611fbe611204565b6001600160a01b0316336001600160a01b03161490565b1760000361113d57604051635fc483c560e01b815260040160405180910390fd5b60006001612002611b80565b5403919050565b337ff73863e3917595a7e67829ed60f0c51bf14c7a0d0de47c7b02a00abd48488e4a546001600160a01b03161461113d57604051635fc483c560e01b815260040160405180910390fd5b7fd08209350903ede075247a6e3c0e53083833967b311cefaa6031e4037ad52c3490565b60006120828261228c565b9050836001600160a01b0316816001600160a01b0316146120b55760405162a1148160e81b815260040160405180910390fd5b6000806120c184612b07565b915091506120e681876120d13390565b6001600160a01b039081169116811491141790565b612111576120f48633611a79565b61211157604051632ce44b5f60e11b815260040160405180910390fd5b6001600160a01b03851661213857604051633a954ecd60e21b815260040160405180910390fd5b801561214357600082555b61214b611b80565b6001600160a01b0387166000908152600591909101602052604090208054600019019055612177611b80565b6001600160a01b03861660008181526005929092016020526040909120805460010190554260a01b17600160e11b176121ae611b80565b60008681526004919091016020526040812091909155600160e11b8416900361222457600184016121dd611b80565b600082815260049190910160205260408120549003612222576121fe611b80565b548114612222578361220e611b80565b600083815260049190910160205260409020555b505b83856001600160a01b0316876001600160a01b0316600080516020613f2d83398151915260405160405180910390a45b505050505050565b610bc4838383604051806020016040528060008152506118e3565b610e51816000612b2f565b6000611a5d611b80565b600081806001116123075761229f611b80565b548110156123075760006122b1611b80565b600083815260049190910160205260408120549150600160e01b82169003612305575b8060000361198d576122e4611b80565b600019909201600081815260049390930160205260409092205490506122d4565b505b604051636f96cda160e11b815260040160405180910390fd5b6060600080516020613f4d833981519152600101805461088b906137f0565b7ff73863e3917595a7e67829ed60f0c51bf14c7a0d0de47c7b02a00abd48488e4a80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60006001600160401b0360406123c4611b80565b6005016000856001600160a01b03166001600160a01b0316815260200190815260200160002054901c169050919050565b806123fe611b80565b336000818152600792909201602090815260408084206001600160a01b03881680865290835293819020805460ff19169515159590951790945592518415158152919290917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b61247b81421190565b61248483421090565b17600103611031576040516309ed117960e11b815242600482015260248101839052604481018290526064016109a7565b826000036124d65760405163198441cb60e01b815260040160405180910390fd5b60006124e1856123b0565b905060006124ed611ff6565b600080516020613f4d833981519152549091508461250b8488613e34565b111561253f5761251b8387613e34565b60405163edc0127360e01b81526004810191909152602481018690526044016109a7565b8061254a8388613e34565b111561257e5761255a8287613e34565b60405163384b48c560e21b81526004810191909152602481018290526044016109a7565b836125898388613e34565b11156125bd576125998287613e34565b604051635cc6d5f560e11b81526004810191909152602481018590526044016109a7565b50505050505050565b6125d0818361383a565b341461103157346125e1828461383a565b604051630d35e92160e01b8152600481019290925260248201526044016109a7565b600061190160f01b612613612c9e565b7f7e105447652a752be7ce7f1a14062d1d57344f91fb0eeebbeec189681e32a329878761263f88612ccb565b8051602091820120604051612683959493928a91019485526001600160a01b0393841660208601529190921660408401526060830191909152608082015260a00190565b60408051601f198184030181529082905280516020918201206001600160f01b03199094169082015260228101919091526042810191909152606201604051602081830303815290604052805190602001209050949350505050565b60008060006126ee8585612d4c565b915091506126fb81612d91565b509392505050565b611031828260405180602001604052806000815250612edb565b61271081111561274357604051631994fc9960e11b8152600481018290526024016109a7565b600061274d612053565b600f01546001600160a01b031690508061277a57604051633f00976960e01b815260040160405180910390fd5b8160000361278c57610bc48134612f51565b600061271061279b843461383a565b6127a59190613851565b90503481900381156127bb576127bb8583612f51565b6119098382612f51565b6127d0848484610d61565b6001600160a01b0383163b15610b5d576127ec84848484612fa2565b610b5d576040516368d2bf6b60e11b815260040160405180910390fd5b606060a06040510180604052602081039150506000815280825b600183039250600a81066030018353600a9004806128235750819003601f19909101908152919050565b60006301ffc9a760e01b6001600160e01b03198316148061287e57506380ac58cd60e01b6001600160e01b03198316145b8061086d5750506001600160e01b031916635b5e139f60e01b1490565b303b156128ba5760405162dc149f60e41b815260040160405180910390fd5b61113d3361233f565b600054610100900460ff166128ea5760405162461bcd60e51b81526004016109a790613dcc565b6daaeb6d7670e522a718067333cd4e3b156110315760405163c3c5a54760e01b81523060048201526daaeb6d7670e522a718067333cd4e9063c3c5a547906024016020604051808303816000875af115801561294a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061296e9190613e17565b6110315780156129e257604051633e9f1edf60e11b81523060048201526001600160a01b03831660248201526daaeb6d7670e522a718067333cd4e90637d3e3dbe906044015b600060405180830381600087803b1580156129ce57600080fd5b505af1158015612254573d6000803e3d6000fd5b6001600160a01b03821615612a315760405163a0af290360e01b81523060048201526001600160a01b03831660248201526daaeb6d7670e522a718067333cd4e9063a0af2903906044016129b4565b604051632210724360e11b81523060048201526daaeb6d7670e522a718067333cd4e90634420e486906024016129b4565b604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527fc09b5f8e6ecd4ed1ee5281e0121f0f52e8c1ea7842745e539b385012ed62e444918101919091527fe6bbd6277e1bf288eed5e8d1780f9a50b239e86b153736bceebccf4ea79d90b360608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b6000806000612b14611b80565b60009485526006016020525050604090912080549092909150565b6000612b3a8361228c565b905080600080612b4986612b07565b915091508415612b8957612b5e8184336120d1565b612b8957612b6c8333611a79565b612b8957604051632ce44b5f60e11b815260040160405180910390fd5b8015612b9457600082555b6fffffffffffffffffffffffffffffffff612bad611b80565b6001600160a01b038516600081815260059290920160205260409091208054929092019091554260a01b17600360e01b17612be6611b80565b60008881526004919091016020526040812091909155600160e11b85169003612c5c5760018601612c15611b80565b600082815260049190910160205260408120549003612c5a57612c36611b80565b548114612c5a5784612c46611b80565b600083815260049190910160205260409020555b505b60405186906000906001600160a01b03861690600080516020613f2d833981519152908390a4612c8a611b80565b600190810180549091019055505050505050565b6000612ca8612053565b600d01544614612cba5761104a612a62565b612cc2612053565b600e0154905090565b606060005b8251811015612d465781838281518110612cec57612cec613b8d565b6020026020010151604051602001612d049190613e47565b60408051601f1981840301815290829052612d229291602001613cde565b60405160208183030381529060405291508080612d3e90613e63565b915050612cd0565b50919050565b6000808251604103612d825760208301516040840151606085015160001a612d768782858561308e565b94509450505050612d8a565b506000905060025b9250929050565b6000816004811115612da557612da5613e7c565b03612dad5750565b6001816004811115612dc157612dc1613e7c565b03612e0e5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016109a7565b6002816004811115612e2257612e22613e7c565b03612e6f5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016109a7565b6003816004811115612e8357612e83613e7c565b03610e515760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016109a7565b612ee58383613152565b6001600160a01b0383163b15610bc4576000612eff611b80565b5490508281035b612f196000868380600101945086612fa2565b612f36576040516368d2bf6b60e11b815260040160405180910390fd5b818110612f065781612f46611b80565b541461190957600080fd5b600080600080600085875af1905080610bc45760405162461bcd60e51b815260206004820152601360248201527211551217d514905394d1915497d19052531151606a1b60448201526064016109a7565b604051630a85bd0160e11b81526000906001600160a01b0385169063150b7a0290612fd7903390899088908890600401613e92565b6020604051808303816000875af1925050508015613012575060408051601f3d908101601f1916820190925261300f91810190613ecf565b60015b613070573d808015613040576040519150601f19603f3d011682016040523d82523d6000602084013e613045565b606091505b508051600003613068576040516368d2bf6b60e11b815260040160405180910390fd5b805181602001fd5b6001600160e01b031916630a85bd0160e11b1490505b949350505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156130c55750600090506003613149565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613119573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661314257600060019250925050613149565b9150600090505b94509492505050565b600061315c611b80565b54905060008290036131815760405163b562e8dd60e01b815260040160405180910390fd5b680100000000000000018202613195611b80565b6001600160a01b038516600081815260059290920160205260409091208054929092019091554260a01b6001841460e11b17176131d0611b80565b600083815260049190910160205260408120919091556001600160a01b038416908383019083908390600080516020613f2d8339815191528180a4600183015b8181146132365780836000600080516020613f2d833981519152600080a4600101613210565b508160000361325757604051622e076360e81b815260040160405180910390fd5b80613260611b80565b5550610bc49050565b6001600160e01b031981168114610e5157600080fd5b60006020828403121561329157600080fd5b813561198d81613269565b60005b838110156132b757818101518382015260200161329f565b50506000910152565b600081518084526132d881602086016020860161329c565b601f01601f19169290920160200192915050565b60208152600061198d60208301846132c0565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171561333d5761333d6132ff565b604052919050565b60006001600160401b0383111561335e5761335e6132ff565b613371601f8401601f1916602001613315565b905082815283838301111561338557600080fd5b828260208301376000602084830101529392505050565b600082601f8301126133ad57600080fd5b61198d83833560208501613345565b6001600160a01b0381168114610e5157600080fd5b6000806000606084860312156133e657600080fd5b83356001600160401b03808211156133fd57600080fd5b6134098783880161339c565b9450602086013591508082111561341f57600080fd5b5061342c8682870161339c565b925050604084013561343d816133bc565b809150509250925092565b60006020828403121561345a57600080fd5b5035919050565b6000806040838503121561347457600080fd5b823561347f816133bc565b946020939093013593505050565b60006020828403121561349f57600080fd5b813561198d816133bc565b6000806000606084860312156134bf57600080fd5b83356134ca816133bc565b925060208401356134da816133bc565b929592945050506040919091013590565b600080604083850312156134fe57600080fd5b50508035926020909101359150565b600060408284031215612d4657600080fd5b60006040828403121561353157600080fd5b61198d838361350d565b60008083601f84011261354d57600080fd5b5081356001600160401b0381111561356457600080fd5b602083019150836020828501011115612d8a57600080fd5b6000806020838503121561358f57600080fd5b82356001600160401b038111156135a557600080fd5b6135b18582860161353b565b90969095509350505050565b60006101008284031215612d4657600080fd5b600061010082840312156135e357600080fd5b61198d83836135bd565b8015158114610e5157600080fd5b6000806040838503121561360e57600080fd5b8235613619816133bc565b91506020830135613629816135ed565b809150509250929050565b6000806000806000806080878903121561364d57600080fd5b8635613658816133bc565b955060208701356001600160401b038082111561367457600080fd5b818901915089601f83011261368857600080fd5b81358181111561369757600080fd5b8a60208260051b85010111156136ac57600080fd5b602083019750809650506040890135945060608901359150808211156136d157600080fd5b506136de89828a0161353b565b979a9699509497509295939492505050565b600080600080610180858703121561370757600080fd5b8435613712816133bc565b93506020850135925061372886604087016135bd565b915061373886610140870161350d565b905092959194509250565b6000806000806080858703121561375957600080fd5b8435613764816133bc565b93506020850135613774816133bc565b92506040850135915060608501356001600160401b0381111561379657600080fd5b8501601f810187136137a757600080fd5b6137b687823560208401613345565b91505092959194509250565b600080604083850312156137d557600080fd5b82356137e0816133bc565b91506020830135613629816133bc565b600181811c9082168061380457607f821691505b602082108103612d4657634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b808202811582820484141761086d5761086d613824565b60008261386e57634e487b7160e01b600052601260045260246000fd5b500490565b6001600160601b0381168114610e5157600080fd5b60006020828403121561389a57600080fd5b813561198d81613873565b81356138b0816133bc565b81546001600160a01b03199081166001600160a01b0392909216918217835560208401356138dd81613873565b60a01b1617905550565b601f821115610bc457600081815260208120601f850160051c8101602086101561390e5750805b601f850160051c820191505b818110156122545782815560010161391a565b6001600160401b03831115613944576139446132ff565b6139588361395283546137f0565b836138e7565b6000601f84116001811461398c57600085156139745750838201355b600019600387901b1c1916600186901b178355611909565b600083815260209020601f19861690835b828110156139bd578685013582556020948501946001909201910161399d565b50868210156139da5760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b8181038181111561086d5761086d613824565b813581556020820135600182015560408201356002820155606082013560038201556080820135600482015560a0820135600582015560c082013560068201556007810160e0830135613a51816135ed565b815490151560ff1660ff19919091161790555050565b600061010082019050823582526020830135602083015260408301356040830152606083013560608301526080830135608083015260a083013560a083015260c083013560c083015260e0830135613abe816135ed565b80151560e08401525092915050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b602081526000613086602083018486613acd565b60006001600160401b0380841115613b2457613b246132ff565b8360051b6020613b35818301613315565b868152918501918181019036841115613b4d57600080fd5b865b84811015613b8157803586811115613b675760008081fd5b613b7336828b0161339c565b845250918301918301613b4f565b50979650505050505050565b634e487b7160e01b600052603260045260246000fd5b6000808335601e19843603018112613bba57600080fd5b8301803591506001600160401b03821115613bd457600080fd5b602001915036819003821315612d8a57600080fd5b8183823760009101908152919050565b60006040820185835260206040818501528185835260608501905060608660051b86010192508660005b87811015613c9257868503605f190183528135368a9003601e19018112613c4957600080fd5b890184810190356001600160401b03811115613c6457600080fd5b803603821315613c7357600080fd5b613c7e878284613acd565b965050509183019190830190600101613c23565b509298975050505050505050565b604081018235613caf816133bc565b6001600160a01b031682526020830135613cc881613873565b6001600160601b03811660208401525092915050565b60008351613cf081846020880161329c565b835190830190613d0481836020880161329c565b01949350505050565b81516001600160401b03811115613d2657613d266132ff565b613d3a81613d3484546137f0565b846138e7565b602080601f831160018114613d6f5760008415613d575750858301515b600019600386901b1c1916600185901b178555612254565b600085815260208120601f198616915b82811015613d9e57888601518255948401946001909101908401613d7f565b5085821015613dbc5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b600060208284031215613e2957600080fd5b815161198d816135ed565b8082018082111561086d5761086d613824565b60008251613e5981846020870161329c565b9190910192915050565b600060018201613e7557613e75613824565b5060010190565b634e487b7160e01b600052602160045260246000fd5b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090613ec5908301846132c0565b9695505050505050565b600060208284031215613ee157600080fd5b815161198d8161326956fef73863e3917595a7e67829ed60f0c51bf14c7a0d0de47c7b02a00abd48488e4bd59f8a8c0d1463371c77782499276e5cbe466fd192ada543ceaea0a36604c1f2ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a40319b847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031dee151c8401928dc223602bb187aff91b9a56c7cae5476ef1b3287b085a16c85fa2646970667358221220656a79602e8f427ded9538ad35634cd79794eba46fffc61e121150b9df29a52c64736f6c63430008110033",
}

// ClublockABI is the input ABI used to generate the binding from.
// Deprecated: Use ClublockMetaData.ABI instead.
var ClublockABI = ClublockMetaData.ABI

// ClublockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ClublockMetaData.Bin instead.
var ClublockBin = ClublockMetaData.Bin

// DeployClublock deploys a new Ethereum contract, binding an instance of Clublock to it.
func DeployClublock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Clublock, error) {
	parsed, err := ClublockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ClublockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Clublock{ClublockCaller: ClublockCaller{contract: contract}, ClublockTransactor: ClublockTransactor{contract: contract}, ClublockFilterer: ClublockFilterer{contract: contract}}, nil
}

// Clublock is an auto generated Go binding around an Ethereum contract.
type Clublock struct {
	ClublockCaller     // Read-only binding to the contract
	ClublockTransactor // Write-only binding to the contract
	ClublockFilterer   // Log filterer for contract events
}

// ClublockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClublockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClublockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClublockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClublockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClublockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClublockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClublockSession struct {
	Contract     *Clublock         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClublockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClublockCallerSession struct {
	Contract *ClublockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ClublockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClublockTransactorSession struct {
	Contract     *ClublockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ClublockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClublockRaw struct {
	Contract *Clublock // Generic contract binding to access the raw methods on
}

// ClublockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClublockCallerRaw struct {
	Contract *ClublockCaller // Generic read-only contract binding to access the raw methods on
}

// ClublockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClublockTransactorRaw struct {
	Contract *ClublockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClublock creates a new instance of Clublock, bound to a specific deployed contract.
func NewClublock(address common.Address, backend bind.ContractBackend) (*Clublock, error) {
	contract, err := bindClublock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Clublock{ClublockCaller: ClublockCaller{contract: contract}, ClublockTransactor: ClublockTransactor{contract: contract}, ClublockFilterer: ClublockFilterer{contract: contract}}, nil
}

// NewClublockCaller creates a new read-only instance of Clublock, bound to a specific deployed contract.
func NewClublockCaller(address common.Address, caller bind.ContractCaller) (*ClublockCaller, error) {
	contract, err := bindClublock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClublockCaller{contract: contract}, nil
}

// NewClublockTransactor creates a new write-only instance of Clublock, bound to a specific deployed contract.
func NewClublockTransactor(address common.Address, transactor bind.ContractTransactor) (*ClublockTransactor, error) {
	contract, err := bindClublock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClublockTransactor{contract: contract}, nil
}

// NewClublockFilterer creates a new log filterer instance of Clublock, bound to a specific deployed contract.
func NewClublockFilterer(address common.Address, filterer bind.ContractFilterer) (*ClublockFilterer, error) {
	contract, err := bindClublock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClublockFilterer{contract: contract}, nil
}

// bindClublock binds a generic wrapper to an already deployed contract.
func bindClublock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClublockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clublock *ClublockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clublock.Contract.ClublockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clublock *ClublockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clublock.Contract.ClublockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clublock *ClublockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clublock.Contract.ClublockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clublock *ClublockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clublock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clublock *ClublockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clublock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clublock *ClublockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clublock.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Clublock *ClublockCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Clublock *ClublockSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Clublock.Contract.BalanceOf(&_Clublock.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Clublock *ClublockCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Clublock.Contract.BalanceOf(&_Clublock.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Clublock *ClublockCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Clublock *ClublockSession) BaseURI() (string, error) {
	return _Clublock.Contract.BaseURI(&_Clublock.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Clublock *ClublockCallerSession) BaseURI() (string, error) {
	return _Clublock.Contract.BaseURI(&_Clublock.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Clublock *ClublockCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Clublock *ClublockSession) ContractURI() (string, error) {
	return _Clublock.Contract.ContractURI(&_Clublock.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Clublock *ClublockCallerSession) ContractURI() (string, error) {
	return _Clublock.Contract.ContractURI(&_Clublock.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Clublock *ClublockCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Clublock *ClublockSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Clublock.Contract.GetApproved(&_Clublock.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Clublock *ClublockCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Clublock.Contract.GetApproved(&_Clublock.CallOpts, tokenId)
}

// GetBurnAddress is a free data retrieval call binding the contract method 0x38b39d29.
//
// Solidity: function getBurnAddress() view returns(address)
func (_Clublock *ClublockCaller) GetBurnAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "getBurnAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBurnAddress is a free data retrieval call binding the contract method 0x38b39d29.
//
// Solidity: function getBurnAddress() view returns(address)
func (_Clublock *ClublockSession) GetBurnAddress() (common.Address, error) {
	return _Clublock.Contract.GetBurnAddress(&_Clublock.CallOpts)
}

// GetBurnAddress is a free data retrieval call binding the contract method 0x38b39d29.
//
// Solidity: function getBurnAddress() view returns(address)
func (_Clublock *ClublockCallerSession) GetBurnAddress() (common.Address, error) {
	return _Clublock.Contract.GetBurnAddress(&_Clublock.CallOpts)
}

// GetMintParams is a free data retrieval call binding the contract method 0xd1e7f79c.
//
// Solidity: function getMintParams() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool))
func (_Clublock *ClublockCaller) GetMintParams(opts *bind.CallOpts) (MintParams, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "getMintParams")

	if err != nil {
		return *new(MintParams), err
	}

	out0 := *abi.ConvertType(out[0], new(MintParams)).(*MintParams)

	return out0, err

}

// GetMintParams is a free data retrieval call binding the contract method 0xd1e7f79c.
//
// Solidity: function getMintParams() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool))
func (_Clublock *ClublockSession) GetMintParams() (MintParams, error) {
	return _Clublock.Contract.GetMintParams(&_Clublock.CallOpts)
}

// GetMintParams is a free data retrieval call binding the contract method 0xd1e7f79c.
//
// Solidity: function getMintParams() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool))
func (_Clublock *ClublockCallerSession) GetMintParams() (MintParams, error) {
	return _Clublock.Contract.GetMintParams(&_Clublock.CallOpts)
}

// GetMintStats is a free data retrieval call binding the contract method 0x840e15d4.
//
// Solidity: function getMintStats(address minter) view returns(uint256 minterNumMinted, uint256 currentTotalSupply, uint256 maxSupply)
func (_Clublock *ClublockCaller) GetMintStats(opts *bind.CallOpts, minter common.Address) (struct {
	MinterNumMinted    *big.Int
	CurrentTotalSupply *big.Int
	MaxSupply          *big.Int
}, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "getMintStats", minter)

	outstruct := new(struct {
		MinterNumMinted    *big.Int
		CurrentTotalSupply *big.Int
		MaxSupply          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinterNumMinted = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentTotalSupply = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaxSupply = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetMintStats is a free data retrieval call binding the contract method 0x840e15d4.
//
// Solidity: function getMintStats(address minter) view returns(uint256 minterNumMinted, uint256 currentTotalSupply, uint256 maxSupply)
func (_Clublock *ClublockSession) GetMintStats(minter common.Address) (struct {
	MinterNumMinted    *big.Int
	CurrentTotalSupply *big.Int
	MaxSupply          *big.Int
}, error) {
	return _Clublock.Contract.GetMintStats(&_Clublock.CallOpts, minter)
}

// GetMintStats is a free data retrieval call binding the contract method 0x840e15d4.
//
// Solidity: function getMintStats(address minter) view returns(uint256 minterNumMinted, uint256 currentTotalSupply, uint256 maxSupply)
func (_Clublock *ClublockCallerSession) GetMintStats(minter common.Address) (struct {
	MinterNumMinted    *big.Int
	CurrentTotalSupply *big.Int
	MaxSupply          *big.Int
}, error) {
	return _Clublock.Contract.GetMintStats(&_Clublock.CallOpts, minter)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Clublock *ClublockCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Clublock *ClublockSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Clublock.Contract.IsApprovedForAll(&_Clublock.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Clublock *ClublockCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Clublock.Contract.IsApprovedForAll(&_Clublock.CallOpts, owner, operator)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Clublock *ClublockCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Clublock *ClublockSession) MaxSupply() (*big.Int, error) {
	return _Clublock.Contract.MaxSupply(&_Clublock.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Clublock *ClublockCallerSession) MaxSupply() (*big.Int, error) {
	return _Clublock.Contract.MaxSupply(&_Clublock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Clublock *ClublockCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Clublock *ClublockSession) Name() (string, error) {
	return _Clublock.Contract.Name(&_Clublock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Clublock *ClublockCallerSession) Name() (string, error) {
	return _Clublock.Contract.Name(&_Clublock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Clublock *ClublockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Clublock *ClublockSession) Owner() (common.Address, error) {
	return _Clublock.Contract.Owner(&_Clublock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Clublock *ClublockCallerSession) Owner() (common.Address, error) {
	return _Clublock.Contract.Owner(&_Clublock.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Clublock *ClublockCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Clublock *ClublockSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Clublock.Contract.OwnerOf(&_Clublock.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Clublock *ClublockCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Clublock.Contract.OwnerOf(&_Clublock.CallOpts, tokenId)
}

// ProvenanceHash is a free data retrieval call binding the contract method 0xc6ab67a3.
//
// Solidity: function provenanceHash() view returns(bytes32)
func (_Clublock *ClublockCaller) ProvenanceHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "provenanceHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvenanceHash is a free data retrieval call binding the contract method 0xc6ab67a3.
//
// Solidity: function provenanceHash() view returns(bytes32)
func (_Clublock *ClublockSession) ProvenanceHash() ([32]byte, error) {
	return _Clublock.Contract.ProvenanceHash(&_Clublock.CallOpts)
}

// ProvenanceHash is a free data retrieval call binding the contract method 0xc6ab67a3.
//
// Solidity: function provenanceHash() view returns(bytes32)
func (_Clublock *ClublockCallerSession) ProvenanceHash() ([32]byte, error) {
	return _Clublock.Contract.ProvenanceHash(&_Clublock.CallOpts)
}

// RoyaltyAddress is a free data retrieval call binding the contract method 0xad2f852a.
//
// Solidity: function royaltyAddress() view returns(address)
func (_Clublock *ClublockCaller) RoyaltyAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "royaltyAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyAddress is a free data retrieval call binding the contract method 0xad2f852a.
//
// Solidity: function royaltyAddress() view returns(address)
func (_Clublock *ClublockSession) RoyaltyAddress() (common.Address, error) {
	return _Clublock.Contract.RoyaltyAddress(&_Clublock.CallOpts)
}

// RoyaltyAddress is a free data retrieval call binding the contract method 0xad2f852a.
//
// Solidity: function royaltyAddress() view returns(address)
func (_Clublock *ClublockCallerSession) RoyaltyAddress() (common.Address, error) {
	return _Clublock.Contract.RoyaltyAddress(&_Clublock.CallOpts)
}

// RoyaltyBasisPoints is a free data retrieval call binding the contract method 0x42260b5d.
//
// Solidity: function royaltyBasisPoints() view returns(uint256)
func (_Clublock *ClublockCaller) RoyaltyBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "royaltyBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoyaltyBasisPoints is a free data retrieval call binding the contract method 0x42260b5d.
//
// Solidity: function royaltyBasisPoints() view returns(uint256)
func (_Clublock *ClublockSession) RoyaltyBasisPoints() (*big.Int, error) {
	return _Clublock.Contract.RoyaltyBasisPoints(&_Clublock.CallOpts)
}

// RoyaltyBasisPoints is a free data retrieval call binding the contract method 0x42260b5d.
//
// Solidity: function royaltyBasisPoints() view returns(uint256)
func (_Clublock *ClublockCallerSession) RoyaltyBasisPoints() (*big.Int, error) {
	return _Clublock.Contract.RoyaltyBasisPoints(&_Clublock.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Clublock *ClublockCaller) RoyaltyInfo(opts *bind.CallOpts, arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "royaltyInfo", arg0, _salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Clublock *ClublockSession) RoyaltyInfo(arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Clublock.Contract.RoyaltyInfo(&_Clublock.CallOpts, arg0, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Clublock *ClublockCallerSession) RoyaltyInfo(arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Clublock.Contract.RoyaltyInfo(&_Clublock.CallOpts, arg0, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Clublock *ClublockCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Clublock *ClublockSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Clublock.Contract.SupportsInterface(&_Clublock.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Clublock *ClublockCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Clublock.Contract.SupportsInterface(&_Clublock.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Clublock *ClublockCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Clublock *ClublockSession) Symbol() (string, error) {
	return _Clublock.Contract.Symbol(&_Clublock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Clublock *ClublockCallerSession) Symbol() (string, error) {
	return _Clublock.Contract.Symbol(&_Clublock.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Clublock *ClublockCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Clublock *ClublockSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Clublock.Contract.TokenURI(&_Clublock.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Clublock *ClublockCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Clublock.Contract.TokenURI(&_Clublock.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Clublock *ClublockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clublock.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Clublock *ClublockSession) TotalSupply() (*big.Int, error) {
	return _Clublock.Contract.TotalSupply(&_Clublock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Clublock *ClublockCallerSession) TotalSupply() (*big.Int, error) {
	return _Clublock.Contract.TotalSupply(&_Clublock.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Clublock *ClublockTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Clublock *ClublockSession) AcceptOwnership() (*types.Transaction, error) {
	return _Clublock.Contract.AcceptOwnership(&_Clublock.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Clublock *ClublockTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Clublock.Contract.AcceptOwnership(&_Clublock.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Clublock *ClublockTransactor) Approve(opts *bind.TransactOpts, operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "approve", operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Clublock *ClublockSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clublock.Contract.Approve(&_Clublock.TransactOpts, operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Clublock *ClublockTransactorSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clublock.Contract.Approve(&_Clublock.TransactOpts, operator, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Clublock *ClublockTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Clublock *ClublockSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Clublock.Contract.Burn(&_Clublock.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Clublock *ClublockTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Clublock.Contract.Burn(&_Clublock.TransactOpts, tokenId)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_Clublock *ClublockTransactor) CancelOwnershipTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "cancelOwnershipTransfer")
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_Clublock *ClublockSession) CancelOwnershipTransfer() (*types.Transaction, error) {
	return _Clublock.Contract.CancelOwnershipTransfer(&_Clublock.TransactOpts)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_Clublock *ClublockTransactorSession) CancelOwnershipTransfer() (*types.Transaction, error) {
	return _Clublock.Contract.CancelOwnershipTransfer(&_Clublock.TransactOpts)
}

// Config is a paid mutator transaction binding the contract method 0xb7c5e064.
//
// Solidity: function config(address payoutAddress, uint256 maxSupply, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, (address,uint96) royalty) returns()
func (_Clublock *ClublockTransactor) Config(opts *bind.TransactOpts, payoutAddress common.Address, maxSupply *big.Int, mintParams MintParams, royalty ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "config", payoutAddress, maxSupply, mintParams, royalty)
}

// Config is a paid mutator transaction binding the contract method 0xb7c5e064.
//
// Solidity: function config(address payoutAddress, uint256 maxSupply, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, (address,uint96) royalty) returns()
func (_Clublock *ClublockSession) Config(payoutAddress common.Address, maxSupply *big.Int, mintParams MintParams, royalty ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo) (*types.Transaction, error) {
	return _Clublock.Contract.Config(&_Clublock.TransactOpts, payoutAddress, maxSupply, mintParams, royalty)
}

// Config is a paid mutator transaction binding the contract method 0xb7c5e064.
//
// Solidity: function config(address payoutAddress, uint256 maxSupply, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, (address,uint96) royalty) returns()
func (_Clublock *ClublockTransactorSession) Config(payoutAddress common.Address, maxSupply *big.Int, mintParams MintParams, royalty ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo) (*types.Transaction, error) {
	return _Clublock.Contract.Config(&_Clublock.TransactOpts, payoutAddress, maxSupply, mintParams, royalty)
}

// EmitBatchMetadataUpdate is a paid mutator transaction binding the contract method 0xa4830114.
//
// Solidity: function emitBatchMetadataUpdate(uint256 fromTokenId, uint256 toTokenId) returns()
func (_Clublock *ClublockTransactor) EmitBatchMetadataUpdate(opts *bind.TransactOpts, fromTokenId *big.Int, toTokenId *big.Int) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "emitBatchMetadataUpdate", fromTokenId, toTokenId)
}

// EmitBatchMetadataUpdate is a paid mutator transaction binding the contract method 0xa4830114.
//
// Solidity: function emitBatchMetadataUpdate(uint256 fromTokenId, uint256 toTokenId) returns()
func (_Clublock *ClublockSession) EmitBatchMetadataUpdate(fromTokenId *big.Int, toTokenId *big.Int) (*types.Transaction, error) {
	return _Clublock.Contract.EmitBatchMetadataUpdate(&_Clublock.TransactOpts, fromTokenId, toTokenId)
}

// EmitBatchMetadataUpdate is a paid mutator transaction binding the contract method 0xa4830114.
//
// Solidity: function emitBatchMetadataUpdate(uint256 fromTokenId, uint256 toTokenId) returns()
func (_Clublock *ClublockTransactorSession) EmitBatchMetadataUpdate(fromTokenId *big.Int, toTokenId *big.Int) (*types.Transaction, error) {
	return _Clublock.Contract.EmitBatchMetadataUpdate(&_Clublock.TransactOpts, fromTokenId, toTokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address signer) returns()
func (_Clublock *ClublockTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, signer common.Address) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "initialize", name, symbol, signer)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address signer) returns()
func (_Clublock *ClublockSession) Initialize(name string, symbol string, signer common.Address) (*types.Transaction, error) {
	return _Clublock.Contract.Initialize(&_Clublock.TransactOpts, name, symbol, signer)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address signer) returns()
func (_Clublock *ClublockTransactorSession) Initialize(name string, symbol string, signer common.Address) (*types.Transaction, error) {
	return _Clublock.Contract.Initialize(&_Clublock.TransactOpts, name, symbol, signer)
}

// MintClublock is a paid mutator transaction binding the contract method 0xa9b34c3b.
//
// Solidity: function mintClublock(address to, string[] assets, uint256 salt, bytes signature) payable returns()
func (_Clublock *ClublockTransactor) MintClublock(opts *bind.TransactOpts, to common.Address, assets []string, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "mintClublock", to, assets, salt, signature)
}

// MintClublock is a paid mutator transaction binding the contract method 0xa9b34c3b.
//
// Solidity: function mintClublock(address to, string[] assets, uint256 salt, bytes signature) payable returns()
func (_Clublock *ClublockSession) MintClublock(to common.Address, assets []string, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Clublock.Contract.MintClublock(&_Clublock.TransactOpts, to, assets, salt, signature)
}

// MintClublock is a paid mutator transaction binding the contract method 0xa9b34c3b.
//
// Solidity: function mintClublock(address to, string[] assets, uint256 salt, bytes signature) payable returns()
func (_Clublock *ClublockTransactorSession) MintClublock(to common.Address, assets []string, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Clublock.Contract.MintClublock(&_Clublock.TransactOpts, to, assets, salt, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Clublock *ClublockTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Clublock *ClublockSession) RenounceOwnership() (*types.Transaction, error) {
	return _Clublock.Contract.RenounceOwnership(&_Clublock.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Clublock *ClublockTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Clublock.Contract.RenounceOwnership(&_Clublock.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Clublock *ClublockTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Clublock *ClublockSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clublock.Contract.SafeTransferFrom(&_Clublock.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Clublock *ClublockTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clublock.Contract.SafeTransferFrom(&_Clublock.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Clublock *ClublockTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Clublock *ClublockSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Clublock.Contract.SafeTransferFrom0(&_Clublock.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Clublock *ClublockTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Clublock.Contract.SafeTransferFrom0(&_Clublock.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Clublock *ClublockTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Clublock *ClublockSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Clublock.Contract.SetApprovalForAll(&_Clublock.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Clublock *ClublockTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Clublock.Contract.SetApprovalForAll(&_Clublock.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Clublock *ClublockTransactor) SetBaseURI(opts *bind.TransactOpts, newBaseURI string) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "setBaseURI", newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Clublock *ClublockSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Clublock.Contract.SetBaseURI(&_Clublock.TransactOpts, newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Clublock *ClublockTransactorSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Clublock.Contract.SetBaseURI(&_Clublock.TransactOpts, newBaseURI)
}

// SetBurnAddress is a paid mutator transaction binding the contract method 0x4b0e7216.
//
// Solidity: function setBurnAddress(address newBurnAddress) returns()
func (_Clublock *ClublockTransactor) SetBurnAddress(opts *bind.TransactOpts, newBurnAddress common.Address) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "setBurnAddress", newBurnAddress)
}

// SetBurnAddress is a paid mutator transaction binding the contract method 0x4b0e7216.
//
// Solidity: function setBurnAddress(address newBurnAddress) returns()
func (_Clublock *ClublockSession) SetBurnAddress(newBurnAddress common.Address) (*types.Transaction, error) {
	return _Clublock.Contract.SetBurnAddress(&_Clublock.TransactOpts, newBurnAddress)
}

// SetBurnAddress is a paid mutator transaction binding the contract method 0x4b0e7216.
//
// Solidity: function setBurnAddress(address newBurnAddress) returns()
func (_Clublock *ClublockTransactorSession) SetBurnAddress(newBurnAddress common.Address) (*types.Transaction, error) {
	return _Clublock.Contract.SetBurnAddress(&_Clublock.TransactOpts, newBurnAddress)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_Clublock *ClublockTransactor) SetContractURI(opts *bind.TransactOpts, newContractURI string) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "setContractURI", newContractURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_Clublock *ClublockSession) SetContractURI(newContractURI string) (*types.Transaction, error) {
	return _Clublock.Contract.SetContractURI(&_Clublock.TransactOpts, newContractURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_Clublock *ClublockTransactorSession) SetContractURI(newContractURI string) (*types.Transaction, error) {
	return _Clublock.Contract.SetContractURI(&_Clublock.TransactOpts, newContractURI)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_Clublock *ClublockTransactor) SetMaxSupply(opts *bind.TransactOpts, newMaxSupply *big.Int) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "setMaxSupply", newMaxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_Clublock *ClublockSession) SetMaxSupply(newMaxSupply *big.Int) (*types.Transaction, error) {
	return _Clublock.Contract.SetMaxSupply(&_Clublock.TransactOpts, newMaxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_Clublock *ClublockTransactorSession) SetMaxSupply(newMaxSupply *big.Int) (*types.Transaction, error) {
	return _Clublock.Contract.SetMaxSupply(&_Clublock.TransactOpts, newMaxSupply)
}

// SetProvenanceHash is a paid mutator transaction binding the contract method 0x099b6bfa.
//
// Solidity: function setProvenanceHash(bytes32 newProvenanceHash) returns()
func (_Clublock *ClublockTransactor) SetProvenanceHash(opts *bind.TransactOpts, newProvenanceHash [32]byte) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "setProvenanceHash", newProvenanceHash)
}

// SetProvenanceHash is a paid mutator transaction binding the contract method 0x099b6bfa.
//
// Solidity: function setProvenanceHash(bytes32 newProvenanceHash) returns()
func (_Clublock *ClublockSession) SetProvenanceHash(newProvenanceHash [32]byte) (*types.Transaction, error) {
	return _Clublock.Contract.SetProvenanceHash(&_Clublock.TransactOpts, newProvenanceHash)
}

// SetProvenanceHash is a paid mutator transaction binding the contract method 0x099b6bfa.
//
// Solidity: function setProvenanceHash(bytes32 newProvenanceHash) returns()
func (_Clublock *ClublockTransactorSession) SetProvenanceHash(newProvenanceHash [32]byte) (*types.Transaction, error) {
	return _Clublock.Contract.SetProvenanceHash(&_Clublock.TransactOpts, newProvenanceHash)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_Clublock *ClublockTransactor) SetRoyaltyInfo(opts *bind.TransactOpts, newInfo ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "setRoyaltyInfo", newInfo)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_Clublock *ClublockSession) SetRoyaltyInfo(newInfo ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo) (*types.Transaction, error) {
	return _Clublock.Contract.SetRoyaltyInfo(&_Clublock.TransactOpts, newInfo)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_Clublock *ClublockTransactorSession) SetRoyaltyInfo(newInfo ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo) (*types.Transaction, error) {
	return _Clublock.Contract.SetRoyaltyInfo(&_Clublock.TransactOpts, newInfo)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Clublock *ClublockTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Clublock *ClublockSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clublock.Contract.TransferFrom(&_Clublock.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Clublock *ClublockTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clublock.Contract.TransferFrom(&_Clublock.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newPotentialOwner) returns()
func (_Clublock *ClublockTransactor) TransferOwnership(opts *bind.TransactOpts, newPotentialOwner common.Address) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "transferOwnership", newPotentialOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newPotentialOwner) returns()
func (_Clublock *ClublockSession) TransferOwnership(newPotentialOwner common.Address) (*types.Transaction, error) {
	return _Clublock.Contract.TransferOwnership(&_Clublock.TransactOpts, newPotentialOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newPotentialOwner) returns()
func (_Clublock *ClublockTransactorSession) TransferOwnership(newPotentialOwner common.Address) (*types.Transaction, error) {
	return _Clublock.Contract.TransferOwnership(&_Clublock.TransactOpts, newPotentialOwner)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x12738db8.
//
// Solidity: function updateCreatorPayoutAddress(address payoutAddress) returns()
func (_Clublock *ClublockTransactor) UpdateCreatorPayoutAddress(opts *bind.TransactOpts, payoutAddress common.Address) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "updateCreatorPayoutAddress", payoutAddress)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x12738db8.
//
// Solidity: function updateCreatorPayoutAddress(address payoutAddress) returns()
func (_Clublock *ClublockSession) UpdateCreatorPayoutAddress(payoutAddress common.Address) (*types.Transaction, error) {
	return _Clublock.Contract.UpdateCreatorPayoutAddress(&_Clublock.TransactOpts, payoutAddress)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x12738db8.
//
// Solidity: function updateCreatorPayoutAddress(address payoutAddress) returns()
func (_Clublock *ClublockTransactorSession) UpdateCreatorPayoutAddress(payoutAddress common.Address) (*types.Transaction, error) {
	return _Clublock.Contract.UpdateCreatorPayoutAddress(&_Clublock.TransactOpts, payoutAddress)
}

// UpdateMintParams is a paid mutator transaction binding the contract method 0x8ff27867.
//
// Solidity: function updateMintParams((uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams) returns()
func (_Clublock *ClublockTransactor) UpdateMintParams(opts *bind.TransactOpts, mintParams MintParams) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "updateMintParams", mintParams)
}

// UpdateMintParams is a paid mutator transaction binding the contract method 0x8ff27867.
//
// Solidity: function updateMintParams((uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams) returns()
func (_Clublock *ClublockSession) UpdateMintParams(mintParams MintParams) (*types.Transaction, error) {
	return _Clublock.Contract.UpdateMintParams(&_Clublock.TransactOpts, mintParams)
}

// UpdateMintParams is a paid mutator transaction binding the contract method 0x8ff27867.
//
// Solidity: function updateMintParams((uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams) returns()
func (_Clublock *ClublockTransactorSession) UpdateMintParams(mintParams MintParams) (*types.Transaction, error) {
	return _Clublock.Contract.UpdateMintParams(&_Clublock.TransactOpts, mintParams)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address signer) returns()
func (_Clublock *ClublockTransactor) UpdateSigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Clublock.contract.Transact(opts, "updateSigner", signer)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address signer) returns()
func (_Clublock *ClublockSession) UpdateSigner(signer common.Address) (*types.Transaction, error) {
	return _Clublock.Contract.UpdateSigner(&_Clublock.TransactOpts, signer)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address signer) returns()
func (_Clublock *ClublockTransactorSession) UpdateSigner(signer common.Address) (*types.Transaction, error) {
	return _Clublock.Contract.UpdateSigner(&_Clublock.TransactOpts, signer)
}

// ClublockAllowListUpdatedIterator is returned from FilterAllowListUpdated and is used to iterate over the raw logs and unpacked data for AllowListUpdated events raised by the Clublock contract.
type ClublockAllowListUpdatedIterator struct {
	Event *ClublockAllowListUpdated // Event containing the contract specifics and raw log

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
func (it *ClublockAllowListUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockAllowListUpdated)
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
		it.Event = new(ClublockAllowListUpdated)
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
func (it *ClublockAllowListUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockAllowListUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockAllowListUpdated represents a AllowListUpdated event raised by the Clublock contract.
type ClublockAllowListUpdated struct {
	NftContract        common.Address
	PreviousMerkleRoot [32]byte
	NewMerkleRoot      [32]byte
	PublicKeyURI       []string
	AllowListURI       string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAllowListUpdated is a free log retrieval operation binding the contract event 0xefcd7e019bc8b47d27881fd59e2619280ca5894f285950f10ab049870652efa5.
//
// Solidity: event AllowListUpdated(address indexed nftContract, bytes32 indexed previousMerkleRoot, bytes32 indexed newMerkleRoot, string[] publicKeyURI, string allowListURI)
func (_Clublock *ClublockFilterer) FilterAllowListUpdated(opts *bind.FilterOpts, nftContract []common.Address, previousMerkleRoot [][32]byte, newMerkleRoot [][32]byte) (*ClublockAllowListUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var previousMerkleRootRule []interface{}
	for _, previousMerkleRootItem := range previousMerkleRoot {
		previousMerkleRootRule = append(previousMerkleRootRule, previousMerkleRootItem)
	}
	var newMerkleRootRule []interface{}
	for _, newMerkleRootItem := range newMerkleRoot {
		newMerkleRootRule = append(newMerkleRootRule, newMerkleRootItem)
	}

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "AllowListUpdated", nftContractRule, previousMerkleRootRule, newMerkleRootRule)
	if err != nil {
		return nil, err
	}
	return &ClublockAllowListUpdatedIterator{contract: _Clublock.contract, event: "AllowListUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowListUpdated is a free log subscription operation binding the contract event 0xefcd7e019bc8b47d27881fd59e2619280ca5894f285950f10ab049870652efa5.
//
// Solidity: event AllowListUpdated(address indexed nftContract, bytes32 indexed previousMerkleRoot, bytes32 indexed newMerkleRoot, string[] publicKeyURI, string allowListURI)
func (_Clublock *ClublockFilterer) WatchAllowListUpdated(opts *bind.WatchOpts, sink chan<- *ClublockAllowListUpdated, nftContract []common.Address, previousMerkleRoot [][32]byte, newMerkleRoot [][32]byte) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var previousMerkleRootRule []interface{}
	for _, previousMerkleRootItem := range previousMerkleRoot {
		previousMerkleRootRule = append(previousMerkleRootRule, previousMerkleRootItem)
	}
	var newMerkleRootRule []interface{}
	for _, newMerkleRootItem := range newMerkleRoot {
		newMerkleRootRule = append(newMerkleRootRule, newMerkleRootItem)
	}

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "AllowListUpdated", nftContractRule, previousMerkleRootRule, newMerkleRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockAllowListUpdated)
				if err := _Clublock.contract.UnpackLog(event, "AllowListUpdated", log); err != nil {
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

// ParseAllowListUpdated is a log parse operation binding the contract event 0xefcd7e019bc8b47d27881fd59e2619280ca5894f285950f10ab049870652efa5.
//
// Solidity: event AllowListUpdated(address indexed nftContract, bytes32 indexed previousMerkleRoot, bytes32 indexed newMerkleRoot, string[] publicKeyURI, string allowListURI)
func (_Clublock *ClublockFilterer) ParseAllowListUpdated(log types.Log) (*ClublockAllowListUpdated, error) {
	event := new(ClublockAllowListUpdated)
	if err := _Clublock.contract.UnpackLog(event, "AllowListUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockAllowedFeeRecipientUpdatedIterator is returned from FilterAllowedFeeRecipientUpdated and is used to iterate over the raw logs and unpacked data for AllowedFeeRecipientUpdated events raised by the Clublock contract.
type ClublockAllowedFeeRecipientUpdatedIterator struct {
	Event *ClublockAllowedFeeRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *ClublockAllowedFeeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockAllowedFeeRecipientUpdated)
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
		it.Event = new(ClublockAllowedFeeRecipientUpdated)
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
func (it *ClublockAllowedFeeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockAllowedFeeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockAllowedFeeRecipientUpdated represents a AllowedFeeRecipientUpdated event raised by the Clublock contract.
type ClublockAllowedFeeRecipientUpdated struct {
	NftContract  common.Address
	FeeRecipient common.Address
	Allowed      bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAllowedFeeRecipientUpdated is a free log retrieval operation binding the contract event 0x6486c31f9d664e241acf94ec2541d328f6b9e97257ae16a1d887f296f879719f.
//
// Solidity: event AllowedFeeRecipientUpdated(address indexed nftContract, address indexed feeRecipient, bool indexed allowed)
func (_Clublock *ClublockFilterer) FilterAllowedFeeRecipientUpdated(opts *bind.FilterOpts, nftContract []common.Address, feeRecipient []common.Address, allowed []bool) (*ClublockAllowedFeeRecipientUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "AllowedFeeRecipientUpdated", nftContractRule, feeRecipientRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &ClublockAllowedFeeRecipientUpdatedIterator{contract: _Clublock.contract, event: "AllowedFeeRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedFeeRecipientUpdated is a free log subscription operation binding the contract event 0x6486c31f9d664e241acf94ec2541d328f6b9e97257ae16a1d887f296f879719f.
//
// Solidity: event AllowedFeeRecipientUpdated(address indexed nftContract, address indexed feeRecipient, bool indexed allowed)
func (_Clublock *ClublockFilterer) WatchAllowedFeeRecipientUpdated(opts *bind.WatchOpts, sink chan<- *ClublockAllowedFeeRecipientUpdated, nftContract []common.Address, feeRecipient []common.Address, allowed []bool) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "AllowedFeeRecipientUpdated", nftContractRule, feeRecipientRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockAllowedFeeRecipientUpdated)
				if err := _Clublock.contract.UnpackLog(event, "AllowedFeeRecipientUpdated", log); err != nil {
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

// ParseAllowedFeeRecipientUpdated is a log parse operation binding the contract event 0x6486c31f9d664e241acf94ec2541d328f6b9e97257ae16a1d887f296f879719f.
//
// Solidity: event AllowedFeeRecipientUpdated(address indexed nftContract, address indexed feeRecipient, bool indexed allowed)
func (_Clublock *ClublockFilterer) ParseAllowedFeeRecipientUpdated(log types.Log) (*ClublockAllowedFeeRecipientUpdated, error) {
	event := new(ClublockAllowedFeeRecipientUpdated)
	if err := _Clublock.contract.UnpackLog(event, "AllowedFeeRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Clublock contract.
type ClublockApprovalIterator struct {
	Event *ClublockApproval // Event containing the contract specifics and raw log

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
func (it *ClublockApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockApproval)
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
		it.Event = new(ClublockApproval)
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
func (it *ClublockApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockApproval represents a Approval event raised by the Clublock contract.
type ClublockApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Clublock *ClublockFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ClublockApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ClublockApprovalIterator{contract: _Clublock.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Clublock *ClublockFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ClublockApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockApproval)
				if err := _Clublock.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Clublock *ClublockFilterer) ParseApproval(log types.Log) (*ClublockApproval, error) {
	event := new(ClublockApproval)
	if err := _Clublock.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Clublock contract.
type ClublockApprovalForAllIterator struct {
	Event *ClublockApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ClublockApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockApprovalForAll)
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
		it.Event = new(ClublockApprovalForAll)
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
func (it *ClublockApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockApprovalForAll represents a ApprovalForAll event raised by the Clublock contract.
type ClublockApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Clublock *ClublockFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ClublockApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ClublockApprovalForAllIterator{contract: _Clublock.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Clublock *ClublockFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ClublockApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockApprovalForAll)
				if err := _Clublock.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Clublock *ClublockFilterer) ParseApprovalForAll(log types.Log) (*ClublockApprovalForAll, error) {
	event := new(ClublockApprovalForAll)
	if err := _Clublock.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Clublock contract.
type ClublockBatchMetadataUpdateIterator struct {
	Event *ClublockBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ClublockBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockBatchMetadataUpdate)
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
		it.Event = new(ClublockBatchMetadataUpdate)
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
func (it *ClublockBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Clublock contract.
type ClublockBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Clublock *ClublockFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*ClublockBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ClublockBatchMetadataUpdateIterator{contract: _Clublock.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Clublock *ClublockFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ClublockBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockBatchMetadataUpdate)
				if err := _Clublock.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Clublock *ClublockFilterer) ParseBatchMetadataUpdate(log types.Log) (*ClublockBatchMetadataUpdate, error) {
	event := new(ClublockBatchMetadataUpdate)
	if err := _Clublock.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockClubLockDeployedIterator is returned from FilterClubLockDeployed and is used to iterate over the raw logs and unpacked data for ClubLockDeployed events raised by the Clublock contract.
type ClublockClubLockDeployedIterator struct {
	Event *ClublockClubLockDeployed // Event containing the contract specifics and raw log

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
func (it *ClublockClubLockDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockClubLockDeployed)
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
		it.Event = new(ClublockClubLockDeployed)
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
func (it *ClublockClubLockDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockClubLockDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockClubLockDeployed represents a ClubLockDeployed event raised by the Clublock contract.
type ClublockClubLockDeployed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClubLockDeployed is a free log retrieval operation binding the contract event 0x6d9aef9de83ac46d2f116c7876c158dd6c92761f13ce1fe88c97d3c3a3f1347d.
//
// Solidity: event ClubLockDeployed()
func (_Clublock *ClublockFilterer) FilterClubLockDeployed(opts *bind.FilterOpts) (*ClublockClubLockDeployedIterator, error) {

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "ClubLockDeployed")
	if err != nil {
		return nil, err
	}
	return &ClublockClubLockDeployedIterator{contract: _Clublock.contract, event: "ClubLockDeployed", logs: logs, sub: sub}, nil
}

// WatchClubLockDeployed is a free log subscription operation binding the contract event 0x6d9aef9de83ac46d2f116c7876c158dd6c92761f13ce1fe88c97d3c3a3f1347d.
//
// Solidity: event ClubLockDeployed()
func (_Clublock *ClublockFilterer) WatchClubLockDeployed(opts *bind.WatchOpts, sink chan<- *ClublockClubLockDeployed) (event.Subscription, error) {

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "ClubLockDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockClubLockDeployed)
				if err := _Clublock.contract.UnpackLog(event, "ClubLockDeployed", log); err != nil {
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

// ParseClubLockDeployed is a log parse operation binding the contract event 0x6d9aef9de83ac46d2f116c7876c158dd6c92761f13ce1fe88c97d3c3a3f1347d.
//
// Solidity: event ClubLockDeployed()
func (_Clublock *ClublockFilterer) ParseClubLockDeployed(log types.Log) (*ClublockClubLockDeployed, error) {
	event := new(ClublockClubLockDeployed)
	if err := _Clublock.contract.UnpackLog(event, "ClubLockDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Clublock contract.
type ClublockConsecutiveTransferIterator struct {
	Event *ClublockConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *ClublockConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockConsecutiveTransfer)
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
		it.Event = new(ClublockConsecutiveTransfer)
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
func (it *ClublockConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Clublock contract.
type ClublockConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Clublock *ClublockFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*ClublockConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ClublockConsecutiveTransferIterator{contract: _Clublock.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Clublock *ClublockFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *ClublockConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockConsecutiveTransfer)
				if err := _Clublock.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Clublock *ClublockFilterer) ParseConsecutiveTransfer(log types.Log) (*ClublockConsecutiveTransfer, error) {
	event := new(ClublockConsecutiveTransfer)
	if err := _Clublock.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockContractURIUpdatedIterator is returned from FilterContractURIUpdated and is used to iterate over the raw logs and unpacked data for ContractURIUpdated events raised by the Clublock contract.
type ClublockContractURIUpdatedIterator struct {
	Event *ClublockContractURIUpdated // Event containing the contract specifics and raw log

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
func (it *ClublockContractURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockContractURIUpdated)
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
		it.Event = new(ClublockContractURIUpdated)
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
func (it *ClublockContractURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockContractURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockContractURIUpdated represents a ContractURIUpdated event raised by the Clublock contract.
type ClublockContractURIUpdated struct {
	NewContractURI string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContractURIUpdated is a free log retrieval operation binding the contract event 0x905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac37378.
//
// Solidity: event ContractURIUpdated(string newContractURI)
func (_Clublock *ClublockFilterer) FilterContractURIUpdated(opts *bind.FilterOpts) (*ClublockContractURIUpdatedIterator, error) {

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return &ClublockContractURIUpdatedIterator{contract: _Clublock.contract, event: "ContractURIUpdated", logs: logs, sub: sub}, nil
}

// WatchContractURIUpdated is a free log subscription operation binding the contract event 0x905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac37378.
//
// Solidity: event ContractURIUpdated(string newContractURI)
func (_Clublock *ClublockFilterer) WatchContractURIUpdated(opts *bind.WatchOpts, sink chan<- *ClublockContractURIUpdated) (event.Subscription, error) {

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockContractURIUpdated)
				if err := _Clublock.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
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

// ParseContractURIUpdated is a log parse operation binding the contract event 0x905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac37378.
//
// Solidity: event ContractURIUpdated(string newContractURI)
func (_Clublock *ClublockFilterer) ParseContractURIUpdated(log types.Log) (*ClublockContractURIUpdated, error) {
	event := new(ClublockContractURIUpdated)
	if err := _Clublock.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockCreatorPayoutAddressUpdatedIterator is returned from FilterCreatorPayoutAddressUpdated and is used to iterate over the raw logs and unpacked data for CreatorPayoutAddressUpdated events raised by the Clublock contract.
type ClublockCreatorPayoutAddressUpdatedIterator struct {
	Event *ClublockCreatorPayoutAddressUpdated // Event containing the contract specifics and raw log

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
func (it *ClublockCreatorPayoutAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockCreatorPayoutAddressUpdated)
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
		it.Event = new(ClublockCreatorPayoutAddressUpdated)
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
func (it *ClublockCreatorPayoutAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockCreatorPayoutAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockCreatorPayoutAddressUpdated represents a CreatorPayoutAddressUpdated event raised by the Clublock contract.
type ClublockCreatorPayoutAddressUpdated struct {
	NftContract      common.Address
	NewPayoutAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCreatorPayoutAddressUpdated is a free log retrieval operation binding the contract event 0x0c69f21751e800ea5960436c9a94370c7adbf54c733a20a025293fbbe8f16252.
//
// Solidity: event CreatorPayoutAddressUpdated(address indexed nftContract, address indexed newPayoutAddress)
func (_Clublock *ClublockFilterer) FilterCreatorPayoutAddressUpdated(opts *bind.FilterOpts, nftContract []common.Address, newPayoutAddress []common.Address) (*ClublockCreatorPayoutAddressUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var newPayoutAddressRule []interface{}
	for _, newPayoutAddressItem := range newPayoutAddress {
		newPayoutAddressRule = append(newPayoutAddressRule, newPayoutAddressItem)
	}

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "CreatorPayoutAddressUpdated", nftContractRule, newPayoutAddressRule)
	if err != nil {
		return nil, err
	}
	return &ClublockCreatorPayoutAddressUpdatedIterator{contract: _Clublock.contract, event: "CreatorPayoutAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchCreatorPayoutAddressUpdated is a free log subscription operation binding the contract event 0x0c69f21751e800ea5960436c9a94370c7adbf54c733a20a025293fbbe8f16252.
//
// Solidity: event CreatorPayoutAddressUpdated(address indexed nftContract, address indexed newPayoutAddress)
func (_Clublock *ClublockFilterer) WatchCreatorPayoutAddressUpdated(opts *bind.WatchOpts, sink chan<- *ClublockCreatorPayoutAddressUpdated, nftContract []common.Address, newPayoutAddress []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var newPayoutAddressRule []interface{}
	for _, newPayoutAddressItem := range newPayoutAddress {
		newPayoutAddressRule = append(newPayoutAddressRule, newPayoutAddressItem)
	}

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "CreatorPayoutAddressUpdated", nftContractRule, newPayoutAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockCreatorPayoutAddressUpdated)
				if err := _Clublock.contract.UnpackLog(event, "CreatorPayoutAddressUpdated", log); err != nil {
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

// ParseCreatorPayoutAddressUpdated is a log parse operation binding the contract event 0x0c69f21751e800ea5960436c9a94370c7adbf54c733a20a025293fbbe8f16252.
//
// Solidity: event CreatorPayoutAddressUpdated(address indexed nftContract, address indexed newPayoutAddress)
func (_Clublock *ClublockFilterer) ParseCreatorPayoutAddressUpdated(log types.Log) (*ClublockCreatorPayoutAddressUpdated, error) {
	event := new(ClublockCreatorPayoutAddressUpdated)
	if err := _Clublock.contract.UnpackLog(event, "CreatorPayoutAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockDropURIUpdatedIterator is returned from FilterDropURIUpdated and is used to iterate over the raw logs and unpacked data for DropURIUpdated events raised by the Clublock contract.
type ClublockDropURIUpdatedIterator struct {
	Event *ClublockDropURIUpdated // Event containing the contract specifics and raw log

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
func (it *ClublockDropURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockDropURIUpdated)
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
		it.Event = new(ClublockDropURIUpdated)
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
func (it *ClublockDropURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockDropURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockDropURIUpdated represents a DropURIUpdated event raised by the Clublock contract.
type ClublockDropURIUpdated struct {
	NftContract common.Address
	NewDropURI  string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDropURIUpdated is a free log retrieval operation binding the contract event 0xa0295608d25b3033c2e2c41cbac8746c2d08767bcfde6d47fae1ed7ba1d32150.
//
// Solidity: event DropURIUpdated(address indexed nftContract, string newDropURI)
func (_Clublock *ClublockFilterer) FilterDropURIUpdated(opts *bind.FilterOpts, nftContract []common.Address) (*ClublockDropURIUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "DropURIUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return &ClublockDropURIUpdatedIterator{contract: _Clublock.contract, event: "DropURIUpdated", logs: logs, sub: sub}, nil
}

// WatchDropURIUpdated is a free log subscription operation binding the contract event 0xa0295608d25b3033c2e2c41cbac8746c2d08767bcfde6d47fae1ed7ba1d32150.
//
// Solidity: event DropURIUpdated(address indexed nftContract, string newDropURI)
func (_Clublock *ClublockFilterer) WatchDropURIUpdated(opts *bind.WatchOpts, sink chan<- *ClublockDropURIUpdated, nftContract []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "DropURIUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockDropURIUpdated)
				if err := _Clublock.contract.UnpackLog(event, "DropURIUpdated", log); err != nil {
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

// ParseDropURIUpdated is a log parse operation binding the contract event 0xa0295608d25b3033c2e2c41cbac8746c2d08767bcfde6d47fae1ed7ba1d32150.
//
// Solidity: event DropURIUpdated(address indexed nftContract, string newDropURI)
func (_Clublock *ClublockFilterer) ParseDropURIUpdated(log types.Log) (*ClublockDropURIUpdated, error) {
	event := new(ClublockDropURIUpdated)
	if err := _Clublock.contract.UnpackLog(event, "DropURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Clublock contract.
type ClublockInitializedIterator struct {
	Event *ClublockInitialized // Event containing the contract specifics and raw log

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
func (it *ClublockInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockInitialized)
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
		it.Event = new(ClublockInitialized)
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
func (it *ClublockInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockInitialized represents a Initialized event raised by the Clublock contract.
type ClublockInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Clublock *ClublockFilterer) FilterInitialized(opts *bind.FilterOpts) (*ClublockInitializedIterator, error) {

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ClublockInitializedIterator{contract: _Clublock.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Clublock *ClublockFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ClublockInitialized) (event.Subscription, error) {

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockInitialized)
				if err := _Clublock.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Clublock *ClublockFilterer) ParseInitialized(log types.Log) (*ClublockInitialized, error) {
	event := new(ClublockInitialized)
	if err := _Clublock.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockMaxSupplyUpdatedIterator is returned from FilterMaxSupplyUpdated and is used to iterate over the raw logs and unpacked data for MaxSupplyUpdated events raised by the Clublock contract.
type ClublockMaxSupplyUpdatedIterator struct {
	Event *ClublockMaxSupplyUpdated // Event containing the contract specifics and raw log

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
func (it *ClublockMaxSupplyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockMaxSupplyUpdated)
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
		it.Event = new(ClublockMaxSupplyUpdated)
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
func (it *ClublockMaxSupplyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockMaxSupplyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockMaxSupplyUpdated represents a MaxSupplyUpdated event raised by the Clublock contract.
type ClublockMaxSupplyUpdated struct {
	NewMaxSupply *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMaxSupplyUpdated is a free log retrieval operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 newMaxSupply)
func (_Clublock *ClublockFilterer) FilterMaxSupplyUpdated(opts *bind.FilterOpts) (*ClublockMaxSupplyUpdatedIterator, error) {

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return &ClublockMaxSupplyUpdatedIterator{contract: _Clublock.contract, event: "MaxSupplyUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxSupplyUpdated is a free log subscription operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 newMaxSupply)
func (_Clublock *ClublockFilterer) WatchMaxSupplyUpdated(opts *bind.WatchOpts, sink chan<- *ClublockMaxSupplyUpdated) (event.Subscription, error) {

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockMaxSupplyUpdated)
				if err := _Clublock.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
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

// ParseMaxSupplyUpdated is a log parse operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 newMaxSupply)
func (_Clublock *ClublockFilterer) ParseMaxSupplyUpdated(log types.Log) (*ClublockMaxSupplyUpdated, error) {
	event := new(ClublockMaxSupplyUpdated)
	if err := _Clublock.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockMintParamsUpdatedIterator is returned from FilterMintParamsUpdated and is used to iterate over the raw logs and unpacked data for MintParamsUpdated events raised by the Clublock contract.
type ClublockMintParamsUpdatedIterator struct {
	Event *ClublockMintParamsUpdated // Event containing the contract specifics and raw log

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
func (it *ClublockMintParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockMintParamsUpdated)
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
		it.Event = new(ClublockMintParamsUpdated)
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
func (it *ClublockMintParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockMintParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockMintParamsUpdated represents a MintParamsUpdated event raised by the Clublock contract.
type ClublockMintParamsUpdated struct {
	Sender     common.Address
	MintParams MintParams
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMintParamsUpdated is a free log retrieval operation binding the contract event 0xe2b0a093e383d45f177604d69d4d2172f9e459de75908c02c9f9dd0ba8591bea.
//
// Solidity: event MintParamsUpdated(address indexed sender, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams)
func (_Clublock *ClublockFilterer) FilterMintParamsUpdated(opts *bind.FilterOpts, sender []common.Address) (*ClublockMintParamsUpdatedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "MintParamsUpdated", senderRule)
	if err != nil {
		return nil, err
	}
	return &ClublockMintParamsUpdatedIterator{contract: _Clublock.contract, event: "MintParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchMintParamsUpdated is a free log subscription operation binding the contract event 0xe2b0a093e383d45f177604d69d4d2172f9e459de75908c02c9f9dd0ba8591bea.
//
// Solidity: event MintParamsUpdated(address indexed sender, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams)
func (_Clublock *ClublockFilterer) WatchMintParamsUpdated(opts *bind.WatchOpts, sink chan<- *ClublockMintParamsUpdated, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "MintParamsUpdated", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockMintParamsUpdated)
				if err := _Clublock.contract.UnpackLog(event, "MintParamsUpdated", log); err != nil {
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

// ParseMintParamsUpdated is a log parse operation binding the contract event 0xe2b0a093e383d45f177604d69d4d2172f9e459de75908c02c9f9dd0ba8591bea.
//
// Solidity: event MintParamsUpdated(address indexed sender, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams)
func (_Clublock *ClublockFilterer) ParseMintParamsUpdated(log types.Log) (*ClublockMintParamsUpdated, error) {
	event := new(ClublockMintParamsUpdated)
	if err := _Clublock.contract.UnpackLog(event, "MintParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Clublock contract.
type ClublockOwnershipTransferredIterator struct {
	Event *ClublockOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ClublockOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockOwnershipTransferred)
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
		it.Event = new(ClublockOwnershipTransferred)
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
func (it *ClublockOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockOwnershipTransferred represents a OwnershipTransferred event raised by the Clublock contract.
type ClublockOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Clublock *ClublockFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ClublockOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ClublockOwnershipTransferredIterator{contract: _Clublock.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Clublock *ClublockFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ClublockOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockOwnershipTransferred)
				if err := _Clublock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Clublock *ClublockFilterer) ParseOwnershipTransferred(log types.Log) (*ClublockOwnershipTransferred, error) {
	event := new(ClublockOwnershipTransferred)
	if err := _Clublock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockPayerUpdatedIterator is returned from FilterPayerUpdated and is used to iterate over the raw logs and unpacked data for PayerUpdated events raised by the Clublock contract.
type ClublockPayerUpdatedIterator struct {
	Event *ClublockPayerUpdated // Event containing the contract specifics and raw log

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
func (it *ClublockPayerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockPayerUpdated)
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
		it.Event = new(ClublockPayerUpdated)
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
func (it *ClublockPayerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockPayerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockPayerUpdated represents a PayerUpdated event raised by the Clublock contract.
type ClublockPayerUpdated struct {
	NftContract common.Address
	Payer       common.Address
	Allowed     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayerUpdated is a free log retrieval operation binding the contract event 0x55a5cfa4bc68ffb9d833b75bf93f6d9c9aadc558dbfa587a9b5bb0ea7d5c38a3.
//
// Solidity: event PayerUpdated(address indexed nftContract, address indexed payer, bool indexed allowed)
func (_Clublock *ClublockFilterer) FilterPayerUpdated(opts *bind.FilterOpts, nftContract []common.Address, payer []common.Address, allowed []bool) (*ClublockPayerUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "PayerUpdated", nftContractRule, payerRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &ClublockPayerUpdatedIterator{contract: _Clublock.contract, event: "PayerUpdated", logs: logs, sub: sub}, nil
}

// WatchPayerUpdated is a free log subscription operation binding the contract event 0x55a5cfa4bc68ffb9d833b75bf93f6d9c9aadc558dbfa587a9b5bb0ea7d5c38a3.
//
// Solidity: event PayerUpdated(address indexed nftContract, address indexed payer, bool indexed allowed)
func (_Clublock *ClublockFilterer) WatchPayerUpdated(opts *bind.WatchOpts, sink chan<- *ClublockPayerUpdated, nftContract []common.Address, payer []common.Address, allowed []bool) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "PayerUpdated", nftContractRule, payerRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockPayerUpdated)
				if err := _Clublock.contract.UnpackLog(event, "PayerUpdated", log); err != nil {
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

// ParsePayerUpdated is a log parse operation binding the contract event 0x55a5cfa4bc68ffb9d833b75bf93f6d9c9aadc558dbfa587a9b5bb0ea7d5c38a3.
//
// Solidity: event PayerUpdated(address indexed nftContract, address indexed payer, bool indexed allowed)
func (_Clublock *ClublockFilterer) ParsePayerUpdated(log types.Log) (*ClublockPayerUpdated, error) {
	event := new(ClublockPayerUpdated)
	if err := _Clublock.contract.UnpackLog(event, "PayerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockPotentialOwnerUpdatedIterator is returned from FilterPotentialOwnerUpdated and is used to iterate over the raw logs and unpacked data for PotentialOwnerUpdated events raised by the Clublock contract.
type ClublockPotentialOwnerUpdatedIterator struct {
	Event *ClublockPotentialOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *ClublockPotentialOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockPotentialOwnerUpdated)
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
		it.Event = new(ClublockPotentialOwnerUpdated)
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
func (it *ClublockPotentialOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockPotentialOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockPotentialOwnerUpdated represents a PotentialOwnerUpdated event raised by the Clublock contract.
type ClublockPotentialOwnerUpdated struct {
	NewPotentialAdministrator common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterPotentialOwnerUpdated is a free log retrieval operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address newPotentialAdministrator)
func (_Clublock *ClublockFilterer) FilterPotentialOwnerUpdated(opts *bind.FilterOpts) (*ClublockPotentialOwnerUpdatedIterator, error) {

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "PotentialOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &ClublockPotentialOwnerUpdatedIterator{contract: _Clublock.contract, event: "PotentialOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchPotentialOwnerUpdated is a free log subscription operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address newPotentialAdministrator)
func (_Clublock *ClublockFilterer) WatchPotentialOwnerUpdated(opts *bind.WatchOpts, sink chan<- *ClublockPotentialOwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "PotentialOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockPotentialOwnerUpdated)
				if err := _Clublock.contract.UnpackLog(event, "PotentialOwnerUpdated", log); err != nil {
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

// ParsePotentialOwnerUpdated is a log parse operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address newPotentialAdministrator)
func (_Clublock *ClublockFilterer) ParsePotentialOwnerUpdated(log types.Log) (*ClublockPotentialOwnerUpdated, error) {
	event := new(ClublockPotentialOwnerUpdated)
	if err := _Clublock.contract.UnpackLog(event, "PotentialOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockProvenanceHashUpdatedIterator is returned from FilterProvenanceHashUpdated and is used to iterate over the raw logs and unpacked data for ProvenanceHashUpdated events raised by the Clublock contract.
type ClublockProvenanceHashUpdatedIterator struct {
	Event *ClublockProvenanceHashUpdated // Event containing the contract specifics and raw log

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
func (it *ClublockProvenanceHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockProvenanceHashUpdated)
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
		it.Event = new(ClublockProvenanceHashUpdated)
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
func (it *ClublockProvenanceHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockProvenanceHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockProvenanceHashUpdated represents a ProvenanceHashUpdated event raised by the Clublock contract.
type ClublockProvenanceHashUpdated struct {
	PreviousHash [32]byte
	NewHash      [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProvenanceHashUpdated is a free log retrieval operation binding the contract event 0x7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c.
//
// Solidity: event ProvenanceHashUpdated(bytes32 previousHash, bytes32 newHash)
func (_Clublock *ClublockFilterer) FilterProvenanceHashUpdated(opts *bind.FilterOpts) (*ClublockProvenanceHashUpdatedIterator, error) {

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "ProvenanceHashUpdated")
	if err != nil {
		return nil, err
	}
	return &ClublockProvenanceHashUpdatedIterator{contract: _Clublock.contract, event: "ProvenanceHashUpdated", logs: logs, sub: sub}, nil
}

// WatchProvenanceHashUpdated is a free log subscription operation binding the contract event 0x7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c.
//
// Solidity: event ProvenanceHashUpdated(bytes32 previousHash, bytes32 newHash)
func (_Clublock *ClublockFilterer) WatchProvenanceHashUpdated(opts *bind.WatchOpts, sink chan<- *ClublockProvenanceHashUpdated) (event.Subscription, error) {

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "ProvenanceHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockProvenanceHashUpdated)
				if err := _Clublock.contract.UnpackLog(event, "ProvenanceHashUpdated", log); err != nil {
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

// ParseProvenanceHashUpdated is a log parse operation binding the contract event 0x7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c.
//
// Solidity: event ProvenanceHashUpdated(bytes32 previousHash, bytes32 newHash)
func (_Clublock *ClublockFilterer) ParseProvenanceHashUpdated(log types.Log) (*ClublockProvenanceHashUpdated, error) {
	event := new(ClublockProvenanceHashUpdated)
	if err := _Clublock.contract.UnpackLog(event, "ProvenanceHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockPublicDropUpdatedIterator is returned from FilterPublicDropUpdated and is used to iterate over the raw logs and unpacked data for PublicDropUpdated events raised by the Clublock contract.
type ClublockPublicDropUpdatedIterator struct {
	Event *ClublockPublicDropUpdated // Event containing the contract specifics and raw log

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
func (it *ClublockPublicDropUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockPublicDropUpdated)
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
		it.Event = new(ClublockPublicDropUpdated)
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
func (it *ClublockPublicDropUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockPublicDropUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockPublicDropUpdated represents a PublicDropUpdated event raised by the Clublock contract.
type ClublockPublicDropUpdated struct {
	NftContract common.Address
	PublicDrop  PublicDrop
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPublicDropUpdated is a free log retrieval operation binding the contract event 0x3e30d8e1f739ea4795c481b21c23f905e938b80339305f3508e43c558e5dead3.
//
// Solidity: event PublicDropUpdated(address indexed nftContract, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop)
func (_Clublock *ClublockFilterer) FilterPublicDropUpdated(opts *bind.FilterOpts, nftContract []common.Address) (*ClublockPublicDropUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "PublicDropUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return &ClublockPublicDropUpdatedIterator{contract: _Clublock.contract, event: "PublicDropUpdated", logs: logs, sub: sub}, nil
}

// WatchPublicDropUpdated is a free log subscription operation binding the contract event 0x3e30d8e1f739ea4795c481b21c23f905e938b80339305f3508e43c558e5dead3.
//
// Solidity: event PublicDropUpdated(address indexed nftContract, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop)
func (_Clublock *ClublockFilterer) WatchPublicDropUpdated(opts *bind.WatchOpts, sink chan<- *ClublockPublicDropUpdated, nftContract []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "PublicDropUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockPublicDropUpdated)
				if err := _Clublock.contract.UnpackLog(event, "PublicDropUpdated", log); err != nil {
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

// ParsePublicDropUpdated is a log parse operation binding the contract event 0x3e30d8e1f739ea4795c481b21c23f905e938b80339305f3508e43c558e5dead3.
//
// Solidity: event PublicDropUpdated(address indexed nftContract, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop)
func (_Clublock *ClublockFilterer) ParsePublicDropUpdated(log types.Log) (*ClublockPublicDropUpdated, error) {
	event := new(ClublockPublicDropUpdated)
	if err := _Clublock.contract.UnpackLog(event, "PublicDropUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockRoyaltyInfoUpdatedIterator is returned from FilterRoyaltyInfoUpdated and is used to iterate over the raw logs and unpacked data for RoyaltyInfoUpdated events raised by the Clublock contract.
type ClublockRoyaltyInfoUpdatedIterator struct {
	Event *ClublockRoyaltyInfoUpdated // Event containing the contract specifics and raw log

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
func (it *ClublockRoyaltyInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockRoyaltyInfoUpdated)
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
		it.Event = new(ClublockRoyaltyInfoUpdated)
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
func (it *ClublockRoyaltyInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockRoyaltyInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockRoyaltyInfoUpdated represents a RoyaltyInfoUpdated event raised by the Clublock contract.
type ClublockRoyaltyInfoUpdated struct {
	Receiver common.Address
	Bps      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyInfoUpdated is a free log retrieval operation binding the contract event 0xf21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d.
//
// Solidity: event RoyaltyInfoUpdated(address receiver, uint256 bps)
func (_Clublock *ClublockFilterer) FilterRoyaltyInfoUpdated(opts *bind.FilterOpts) (*ClublockRoyaltyInfoUpdatedIterator, error) {

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "RoyaltyInfoUpdated")
	if err != nil {
		return nil, err
	}
	return &ClublockRoyaltyInfoUpdatedIterator{contract: _Clublock.contract, event: "RoyaltyInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchRoyaltyInfoUpdated is a free log subscription operation binding the contract event 0xf21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d.
//
// Solidity: event RoyaltyInfoUpdated(address receiver, uint256 bps)
func (_Clublock *ClublockFilterer) WatchRoyaltyInfoUpdated(opts *bind.WatchOpts, sink chan<- *ClublockRoyaltyInfoUpdated) (event.Subscription, error) {

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "RoyaltyInfoUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockRoyaltyInfoUpdated)
				if err := _Clublock.contract.UnpackLog(event, "RoyaltyInfoUpdated", log); err != nil {
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

// ParseRoyaltyInfoUpdated is a log parse operation binding the contract event 0xf21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d.
//
// Solidity: event RoyaltyInfoUpdated(address receiver, uint256 bps)
func (_Clublock *ClublockFilterer) ParseRoyaltyInfoUpdated(log types.Log) (*ClublockRoyaltyInfoUpdated, error) {
	event := new(ClublockRoyaltyInfoUpdated)
	if err := _Clublock.contract.UnpackLog(event, "RoyaltyInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockSeaDropMintIterator is returned from FilterSeaDropMint and is used to iterate over the raw logs and unpacked data for SeaDropMint events raised by the Clublock contract.
type ClublockSeaDropMintIterator struct {
	Event *ClublockSeaDropMint // Event containing the contract specifics and raw log

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
func (it *ClublockSeaDropMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockSeaDropMint)
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
		it.Event = new(ClublockSeaDropMint)
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
func (it *ClublockSeaDropMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockSeaDropMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockSeaDropMint represents a SeaDropMint event raised by the Clublock contract.
type ClublockSeaDropMint struct {
	NftContract    common.Address
	Minter         common.Address
	FeeRecipient   common.Address
	Payer          common.Address
	QuantityMinted *big.Int
	UnitMintPrice  *big.Int
	FeeBps         *big.Int
	DropStageIndex *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSeaDropMint is a free log retrieval operation binding the contract event 0xe90cf9cc0a552cf52ea6ff74ece0f1c8ae8cc9ad630d3181f55ac43ca076b7d6.
//
// Solidity: event SeaDropMint(address indexed nftContract, address indexed minter, address indexed feeRecipient, address payer, uint256 quantityMinted, uint256 unitMintPrice, uint256 feeBps, uint256 dropStageIndex)
func (_Clublock *ClublockFilterer) FilterSeaDropMint(opts *bind.FilterOpts, nftContract []common.Address, minter []common.Address, feeRecipient []common.Address) (*ClublockSeaDropMintIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "SeaDropMint", nftContractRule, minterRule, feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &ClublockSeaDropMintIterator{contract: _Clublock.contract, event: "SeaDropMint", logs: logs, sub: sub}, nil
}

// WatchSeaDropMint is a free log subscription operation binding the contract event 0xe90cf9cc0a552cf52ea6ff74ece0f1c8ae8cc9ad630d3181f55ac43ca076b7d6.
//
// Solidity: event SeaDropMint(address indexed nftContract, address indexed minter, address indexed feeRecipient, address payer, uint256 quantityMinted, uint256 unitMintPrice, uint256 feeBps, uint256 dropStageIndex)
func (_Clublock *ClublockFilterer) WatchSeaDropMint(opts *bind.WatchOpts, sink chan<- *ClublockSeaDropMint, nftContract []common.Address, minter []common.Address, feeRecipient []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "SeaDropMint", nftContractRule, minterRule, feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockSeaDropMint)
				if err := _Clublock.contract.UnpackLog(event, "SeaDropMint", log); err != nil {
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

// ParseSeaDropMint is a log parse operation binding the contract event 0xe90cf9cc0a552cf52ea6ff74ece0f1c8ae8cc9ad630d3181f55ac43ca076b7d6.
//
// Solidity: event SeaDropMint(address indexed nftContract, address indexed minter, address indexed feeRecipient, address payer, uint256 quantityMinted, uint256 unitMintPrice, uint256 feeBps, uint256 dropStageIndex)
func (_Clublock *ClublockFilterer) ParseSeaDropMint(log types.Log) (*ClublockSeaDropMint, error) {
	event := new(ClublockSeaDropMint)
	if err := _Clublock.contract.UnpackLog(event, "SeaDropMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockSignedMintValidationParamsUpdatedIterator is returned from FilterSignedMintValidationParamsUpdated and is used to iterate over the raw logs and unpacked data for SignedMintValidationParamsUpdated events raised by the Clublock contract.
type ClublockSignedMintValidationParamsUpdatedIterator struct {
	Event *ClublockSignedMintValidationParamsUpdated // Event containing the contract specifics and raw log

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
func (it *ClublockSignedMintValidationParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockSignedMintValidationParamsUpdated)
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
		it.Event = new(ClublockSignedMintValidationParamsUpdated)
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
func (it *ClublockSignedMintValidationParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockSignedMintValidationParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockSignedMintValidationParamsUpdated represents a SignedMintValidationParamsUpdated event raised by the Clublock contract.
type ClublockSignedMintValidationParamsUpdated struct {
	NftContract                common.Address
	Signer                     common.Address
	SignedMintValidationParams SignedMintValidationParams
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterSignedMintValidationParamsUpdated is a free log retrieval operation binding the contract event 0xcaeb4009c05208df426d15ff50b608287b05d21dee1f790552ea451a540a7be0.
//
// Solidity: event SignedMintValidationParamsUpdated(address indexed nftContract, address indexed signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams)
func (_Clublock *ClublockFilterer) FilterSignedMintValidationParamsUpdated(opts *bind.FilterOpts, nftContract []common.Address, signer []common.Address) (*ClublockSignedMintValidationParamsUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "SignedMintValidationParamsUpdated", nftContractRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &ClublockSignedMintValidationParamsUpdatedIterator{contract: _Clublock.contract, event: "SignedMintValidationParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchSignedMintValidationParamsUpdated is a free log subscription operation binding the contract event 0xcaeb4009c05208df426d15ff50b608287b05d21dee1f790552ea451a540a7be0.
//
// Solidity: event SignedMintValidationParamsUpdated(address indexed nftContract, address indexed signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams)
func (_Clublock *ClublockFilterer) WatchSignedMintValidationParamsUpdated(opts *bind.WatchOpts, sink chan<- *ClublockSignedMintValidationParamsUpdated, nftContract []common.Address, signer []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "SignedMintValidationParamsUpdated", nftContractRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockSignedMintValidationParamsUpdated)
				if err := _Clublock.contract.UnpackLog(event, "SignedMintValidationParamsUpdated", log); err != nil {
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

// ParseSignedMintValidationParamsUpdated is a log parse operation binding the contract event 0xcaeb4009c05208df426d15ff50b608287b05d21dee1f790552ea451a540a7be0.
//
// Solidity: event SignedMintValidationParamsUpdated(address indexed nftContract, address indexed signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams)
func (_Clublock *ClublockFilterer) ParseSignedMintValidationParamsUpdated(log types.Log) (*ClublockSignedMintValidationParamsUpdated, error) {
	event := new(ClublockSignedMintValidationParamsUpdated)
	if err := _Clublock.contract.UnpackLog(event, "SignedMintValidationParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockSignerUpdatedIterator is returned from FilterSignerUpdated and is used to iterate over the raw logs and unpacked data for SignerUpdated events raised by the Clublock contract.
type ClublockSignerUpdatedIterator struct {
	Event *ClublockSignerUpdated // Event containing the contract specifics and raw log

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
func (it *ClublockSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockSignerUpdated)
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
		it.Event = new(ClublockSignerUpdated)
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
func (it *ClublockSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockSignerUpdated represents a SignerUpdated event raised by the Clublock contract.
type ClublockSignerUpdated struct {
	Sender common.Address
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerUpdated is a free log retrieval operation binding the contract event 0x2d025324f0a785e8c12d0a0d91a9caa49df4ef20ff87e0df7213a1d4f3157beb.
//
// Solidity: event SignerUpdated(address indexed sender, address indexed signer)
func (_Clublock *ClublockFilterer) FilterSignerUpdated(opts *bind.FilterOpts, sender []common.Address, signer []common.Address) (*ClublockSignerUpdatedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "SignerUpdated", senderRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &ClublockSignerUpdatedIterator{contract: _Clublock.contract, event: "SignerUpdated", logs: logs, sub: sub}, nil
}

// WatchSignerUpdated is a free log subscription operation binding the contract event 0x2d025324f0a785e8c12d0a0d91a9caa49df4ef20ff87e0df7213a1d4f3157beb.
//
// Solidity: event SignerUpdated(address indexed sender, address indexed signer)
func (_Clublock *ClublockFilterer) WatchSignerUpdated(opts *bind.WatchOpts, sink chan<- *ClublockSignerUpdated, sender []common.Address, signer []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "SignerUpdated", senderRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockSignerUpdated)
				if err := _Clublock.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
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

// ParseSignerUpdated is a log parse operation binding the contract event 0x2d025324f0a785e8c12d0a0d91a9caa49df4ef20ff87e0df7213a1d4f3157beb.
//
// Solidity: event SignerUpdated(address indexed sender, address indexed signer)
func (_Clublock *ClublockFilterer) ParseSignerUpdated(log types.Log) (*ClublockSignerUpdated, error) {
	event := new(ClublockSignerUpdated)
	if err := _Clublock.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockTokenGatedDropStageUpdatedIterator is returned from FilterTokenGatedDropStageUpdated and is used to iterate over the raw logs and unpacked data for TokenGatedDropStageUpdated events raised by the Clublock contract.
type ClublockTokenGatedDropStageUpdatedIterator struct {
	Event *ClublockTokenGatedDropStageUpdated // Event containing the contract specifics and raw log

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
func (it *ClublockTokenGatedDropStageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockTokenGatedDropStageUpdated)
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
		it.Event = new(ClublockTokenGatedDropStageUpdated)
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
func (it *ClublockTokenGatedDropStageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockTokenGatedDropStageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockTokenGatedDropStageUpdated represents a TokenGatedDropStageUpdated event raised by the Clublock contract.
type ClublockTokenGatedDropStageUpdated struct {
	NftContract     common.Address
	AllowedNftToken common.Address
	DropStage       TokenGatedDropStage
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenGatedDropStageUpdated is a free log retrieval operation binding the contract event 0xc695f93ae16034280e4fc93181b6afca9af23027ac1f1842a2287ba25cdc4476.
//
// Solidity: event TokenGatedDropStageUpdated(address indexed nftContract, address indexed allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage)
func (_Clublock *ClublockFilterer) FilterTokenGatedDropStageUpdated(opts *bind.FilterOpts, nftContract []common.Address, allowedNftToken []common.Address) (*ClublockTokenGatedDropStageUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var allowedNftTokenRule []interface{}
	for _, allowedNftTokenItem := range allowedNftToken {
		allowedNftTokenRule = append(allowedNftTokenRule, allowedNftTokenItem)
	}

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "TokenGatedDropStageUpdated", nftContractRule, allowedNftTokenRule)
	if err != nil {
		return nil, err
	}
	return &ClublockTokenGatedDropStageUpdatedIterator{contract: _Clublock.contract, event: "TokenGatedDropStageUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenGatedDropStageUpdated is a free log subscription operation binding the contract event 0xc695f93ae16034280e4fc93181b6afca9af23027ac1f1842a2287ba25cdc4476.
//
// Solidity: event TokenGatedDropStageUpdated(address indexed nftContract, address indexed allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage)
func (_Clublock *ClublockFilterer) WatchTokenGatedDropStageUpdated(opts *bind.WatchOpts, sink chan<- *ClublockTokenGatedDropStageUpdated, nftContract []common.Address, allowedNftToken []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var allowedNftTokenRule []interface{}
	for _, allowedNftTokenItem := range allowedNftToken {
		allowedNftTokenRule = append(allowedNftTokenRule, allowedNftTokenItem)
	}

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "TokenGatedDropStageUpdated", nftContractRule, allowedNftTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockTokenGatedDropStageUpdated)
				if err := _Clublock.contract.UnpackLog(event, "TokenGatedDropStageUpdated", log); err != nil {
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

// ParseTokenGatedDropStageUpdated is a log parse operation binding the contract event 0xc695f93ae16034280e4fc93181b6afca9af23027ac1f1842a2287ba25cdc4476.
//
// Solidity: event TokenGatedDropStageUpdated(address indexed nftContract, address indexed allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage)
func (_Clublock *ClublockFilterer) ParseTokenGatedDropStageUpdated(log types.Log) (*ClublockTokenGatedDropStageUpdated, error) {
	event := new(ClublockTokenGatedDropStageUpdated)
	if err := _Clublock.contract.UnpackLog(event, "TokenGatedDropStageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Clublock contract.
type ClublockTransferIterator struct {
	Event *ClublockTransfer // Event containing the contract specifics and raw log

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
func (it *ClublockTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockTransfer)
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
		it.Event = new(ClublockTransfer)
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
func (it *ClublockTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockTransfer represents a Transfer event raised by the Clublock contract.
type ClublockTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Clublock *ClublockFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ClublockTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ClublockTransferIterator{contract: _Clublock.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Clublock *ClublockFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ClublockTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockTransfer)
				if err := _Clublock.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Clublock *ClublockFilterer) ParseTransfer(log types.Log) (*ClublockTransfer, error) {
	event := new(ClublockTransfer)
	if err := _Clublock.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClublockMintClublockUpdatedIterator is returned from FilterMintClublockUpdated and is used to iterate over the raw logs and unpacked data for MintClublockUpdated events raised by the Clublock contract.
type ClublockMintClublockUpdatedIterator struct {
	Event *ClublockMintClublockUpdated // Event containing the contract specifics and raw log

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
func (it *ClublockMintClublockUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClublockMintClublockUpdated)
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
		it.Event = new(ClublockMintClublockUpdated)
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
func (it *ClublockMintClublockUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClublockMintClublockUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClublockMintClublockUpdated represents a MintClublockUpdated event raised by the Clublock contract.
type ClublockMintClublockUpdated struct {
	Minter    common.Address
	To        common.Address
	MintPrice *big.Int
	Assets    []string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMintClublockUpdated is a free log retrieval operation binding the contract event 0xbeeb7c756bb80c977a1eeb91b3bdb1af5d9107550f8c929feee514b674722e94.
//
// Solidity: event mintClublockUpdated(address indexed minter, address indexed to, uint256 mintPrice, string[] assets)
func (_Clublock *ClublockFilterer) FilterMintClublockUpdated(opts *bind.FilterOpts, minter []common.Address, to []common.Address) (*ClublockMintClublockUpdatedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Clublock.contract.FilterLogs(opts, "mintClublockUpdated", minterRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ClublockMintClublockUpdatedIterator{contract: _Clublock.contract, event: "mintClublockUpdated", logs: logs, sub: sub}, nil
}

// WatchMintClublockUpdated is a free log subscription operation binding the contract event 0xbeeb7c756bb80c977a1eeb91b3bdb1af5d9107550f8c929feee514b674722e94.
//
// Solidity: event mintClublockUpdated(address indexed minter, address indexed to, uint256 mintPrice, string[] assets)
func (_Clublock *ClublockFilterer) WatchMintClublockUpdated(opts *bind.WatchOpts, sink chan<- *ClublockMintClublockUpdated, minter []common.Address, to []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Clublock.contract.WatchLogs(opts, "mintClublockUpdated", minterRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClublockMintClublockUpdated)
				if err := _Clublock.contract.UnpackLog(event, "mintClublockUpdated", log); err != nil {
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

// ParseMintClublockUpdated is a log parse operation binding the contract event 0xbeeb7c756bb80c977a1eeb91b3bdb1af5d9107550f8c929feee514b674722e94.
//
// Solidity: event mintClublockUpdated(address indexed minter, address indexed to, uint256 mintPrice, string[] assets)
func (_Clublock *ClublockFilterer) ParseMintClublockUpdated(log types.Log) (*ClublockMintClublockUpdated, error) {
	event := new(ClublockMintClublockUpdated)
	if err := _Clublock.contract.UnpackLog(event, "mintClublockUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
