// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package club

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

// ERC721ClubUpgradeableClubParams is an auto generated low-level Go binding around an user-defined struct.
type ERC721ClubUpgradeableClubParams struct {
	TokenId            *big.Int
	Token              []*big.Int
	CurrentSpace       *big.Int
	Space              *big.Int
	CurrentVideoLength *big.Int
	Videolength        *big.Int
}

// ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo is an auto generated low-level Go binding around an user-defined struct.
type ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo struct {
	RoyaltyAddress common.Address
	RoyaltyBps     *big.Int
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

// ClubMetaData contains all meta data concerning the Club contract.
var ClubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BurnIncorrectSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"CannotExceedMaxSupplyOfUint64\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CreatorPayoutAddressCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateAllowedPlegedContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateFeeRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicatePayer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientNotPresent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"IncorrectPayment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeBps\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"basisPoints\",\"type\":\"uint256\"}],\"name\":\"InvalidRoyaltyBasisPoints\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recoveredSigner\",\"type\":\"address\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedEndTime\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumOrMaximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedFeeBps\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedMaxTokenSupplyForStage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedMaxTotalMintableByWallet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedMintPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintQuantityCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allowed\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxMintedPerWallet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxSupply\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxTokenSupplyForStage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"}],\"name\":\"NotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotNextOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedPlegeContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OnlyClubOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"OnlyINonFungibleSeaDropToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyIsApprovedForAll\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPleged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerNotPresent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProvenanceHashCannotBeSetAfterMintStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoyaltyAddressCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignedMintsMustRestrictFeeRecipients\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerNotPresent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedDropAllowedNftTokenCannotBeDropToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedDropAllowedNftTokenCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedDropStageNotPresent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowedNftTokenId\",\"type\":\"uint256\"}],\"name\":\"TokenGatedNotTokenOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowedNftTokenId\",\"type\":\"uint256\"}],\"name\":\"TokenGatedTokenIdAlreadyRedeemed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousMerkleRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newMerkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"publicKeyURI\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"name\":\"AllowListUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AllowedFeeRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AllowedPlegedContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ClubTokenDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"}],\"name\":\"ContractURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPayoutAddress\",\"type\":\"address\"}],\"name\":\"CreatorPayoutAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newDropURI\",\"type\":\"string\"}],\"name\":\"DropURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"PayerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"clubId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"token\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"currentSpace\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVideoLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Videolength\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structERC721ClubUpgradeable.ClubParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"minted\",\"type\":\"bool\"}],\"name\":\"PlegedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPotentialAdministrator\",\"type\":\"address\"}],\"name\":\"PotentialOwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newHash\",\"type\":\"bytes32\"}],\"name\":\"ProvenanceHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"}],\"name\":\"PublicDropUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"name\":\"RoyaltyInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantityMinted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unitMintPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dropStageIndex\",\"type\":\"uint256\"}],\"name\":\"SeaDropMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structSignedMintValidationParams\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple\"}],\"name\":\"SignedMintValidationParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"SpaceSizeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTokenGatedDropStage\",\"name\":\"dropStage\",\"type\":\"tuple\"}],\"name\":\"TokenGatedDropStageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"clubId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"token\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"currentSpace\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVideoLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Videolength\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structERC721ClubUpgradeable.ClubParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"UnplegedUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"}],\"name\":\"emitBatchMetadataUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBurnAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getClubId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getClubSpace\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"clubId\",\"type\":\"uint256\"}],\"name\":\"getClubTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"getMintStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minterNumMinted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"plegeSigner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"defaultSpace\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultVideoLength\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"clubId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"plegeToken\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"plege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provenanceHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBurnAddress\",\"type\":\"address\"}],\"name\":\"setBurnAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newProvenanceHash\",\"type\":\"bytes32\"}],\"name\":\"setProvenanceHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"royaltyAddress\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"royaltyBps\",\"type\":\"uint96\"}],\"internalType\":\"structISeaDropTokenContractMetadataUpgradeable.RoyaltyInfo\",\"name\":\"newInfo\",\"type\":\"tuple\"}],\"name\":\"setRoyaltyInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPotentialOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"clubId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"plegedToken\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"unplege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateAllowedPledgeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"updateSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"updateSpaceSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50614045806100206000396000f3fe608060405234801561001057600080fd5b506004361061027f5760003560e01c80636f8b44b01161015c578063a7ecd37e116100ce578063c87b56dd11610087578063c87b56dd146105d0578063d5abeb01146105e3578063e8a3d485146105eb578063e985e9c5146105f3578063f2fde38b14610606578063fc46607a1461061957600080fd5b8063a7ecd37e1461053f578063ad2f852a14610552578063b88d4fde14610570578063b8b52cc114610583578063c6938bc214610596578063c6ab67a3146105a957600080fd5b8063840e15d411610120578063840e15d4146104c85780638da5cb5b146104f6578063938e3d7b146104fe57806395d89b4114610511578063a22cb46514610519578063a48301141461052c57600080fd5b80636f8b44b01461047f57806370a0823114610492578063715018a6146104a557806374823132146104ad57806379ba5097146104c057600080fd5b806342260b5d116101f55780634c182abc116101b95780634c182abc14610403578063530c3de71461041657806355f804b3146104295780635c4056811461043c5780636352211e146104645780636c0360eb1461047757600080fd5b806342260b5d1461039257806342842e0e146103b757806342966c68146103ca57806344dae42c146103dd5780634b0e7216146103f057600080fd5b80630dfa40b7116102475780630dfa40b71461031457806318160ddd1461033557806323452b9c1461033d57806323b872dd146103455780632a55205a1461035857806338b39d291461038a57600080fd5b806301ffc9a71461028457806306fdde03146102ac578063081812fc146102c1578063095ea7b3146102ec578063099b6bfa14610301575b600080fd5b610297610292366004613542565b61062c565b60405190151581526020015b60405180910390f35b6102b4610657565b6040516102a391906135af565b6102d46102cf3660046135c2565b6106f2565b6040516001600160a01b0390911681526020016102a3565b6102ff6102fa3660046135f0565b61073f565b005b6102ff61030f3660046135c2565b610758565b6103276103223660046135c2565b6107ee565b6040519081526020016102a3565b61032761080c565b6102ff61082c565b6102ff61035336600461361c565b610887565b61036b61036636600461365d565b6108b2565b604080516001600160a01b0390931683526020830191909152016102a3565b6102d4610905565b600080516020613fd083398151915254600160a01b90046001600160601b0316610327565b6102ff6103c536600461361c565b61091e565b6102ff6103d83660046135c2565b610943565b6102ff6103eb36600461367f565b610980565b6102ff6103fe366004613697565b610aa3565b6102ff6104113660046136f6565b610ad5565b6103276104243660046135c2565b610f6a565b6102ff6104373660046137c6565b610f88565b61044f61044a3660046135c2565b611019565b604080519283526020830191909152016102a3565b6102d46104723660046135c2565b611057565b6102b4611062565b6102ff61048d3660046135c2565b611071565b6103276104a0366004613697565b6110e6565b6102ff61114f565b6102ff6104bb3660046138b4565b611163565b6102ff6113b7565b6104db6104d6366004613697565b61144a565b604080519384526020840192909252908201526060016102a3565b6102d461147c565b6102ff61050c3660046137c6565b6114a4565b6102b461150a565b6102ff61052736600461394a565b611522565b6102ff61053a36600461365d565b611536565b6102ff61054d366004613697565b611574565b600080516020613fd0833981519152546001600160a01b03166102d4565b6102ff61057e366004613983565b6115d4565b6102ff61059136600461394a565b611601565b6102ff6105a43660046136f6565b61179d565b7fb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031c54610327565b6102b46105de3660046135c2565b611c7b565b610327611cff565b6102b4611d17565b610297610601366004613a03565b611d2d565b6102ff610614366004613697565b611d6a565b6102ff61062736600461365d565b611df4565b60006001600160e01b03198216639c15441560e01b1480610651575061065182611e98565b92915050565b6060610661611ed8565b600201805461066f90613a31565b80601f016020809104026020016040519081016040528092919081815260200182805461069b90613a31565b80156106e85780601f106106bd576101008083540402835291602001916106e8565b820191906000526020600020905b8154815290600101906020018083116106cb57829003601f168201915b5050505050905090565b60006106fd82611efc565b61071a576040516333d1c03960e21b815260040160405180910390fd5b610722611ed8565b60009283526006016020525060409020546001600160a01b031690565b8161074981611f45565b6107538383611ffe565b505050565b6107606120ac565b600061076a6120f2565b11156107895760405163e03264af60e01b815260040160405180910390fd5b7fb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031c80549082905560408051828152602081018490527f7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c91015b60405180910390a15050565b60006107f8612105565b600092835260010160205250604090205490565b60006001610818611ed8565b60010154610824611ed8565b540303919050565b610834612129565b600080516020613f5083398151915280546001600160a01b0319169055604080516000815290517f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da9181900360200190a1565b826001600160a01b03811633146108a1576108a133611f45565b6108ac848484612173565b50505050565b600080516020613fd083398151915280546000918291612710906108e690600160a01b90046001600160601b031686613a7b565b6108f09190613a92565b90546001600160a01b03169590945092505050565b600061090f612105565b546001600160a01b0316919050565b826001600160a01b03811633146109385761093833611f45565b6108ac848484612354565b61094b612105565b546001600160a01b031633146109745760405163ea885d8360e01b815260040160405180910390fd5b61097d8161236f565b50565b6109886120ac565b60006109976020830183613697565b6001600160a01b0316036109be57604051631cc0baef60e01b815260040160405180910390fd5b6127106109d16040830160208401613ac9565b6001600160601b03161115610a1a576109f06040820160208301613ac9565b604051633cadbafb60e01b81526001600160601b0390911660048201526024015b60405180910390fd5b80600080516020613fd0833981519152610a348282613ae6565b507ff21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d9050610a656020830183613697565b610a756040840160208501613ac9565b604080516001600160a01b0390931683526001600160601b039091166020830152015b60405180910390a150565b610aab612129565b80610ab4612105565b80546001600160a01b0319166001600160a01b039290921691909117905550565b600080516020613f7083398151915254600114610b215760405162461bcd60e51b815260206004820152600a6024820152695245454e5452414e435960b01b6044820152606401610a11565b6002600080516020613f7083398151915255610b3c8661237a565b610b4688886123c4565b610b7788888888888888887fecd3fbbcdd4013be93b976ea0feb9b776e977e17ef78c34a0efb7b315e2a6f76612423565b6000610b81612105565b60008981526004919091016020526040812091505b85811015610d6e576001600160a01b038816600090815260208390526040812090888884818110610bc957610bc9613b28565b602090810292909201358352508101919091526040016000205460ff16610c035760405163cdbff65760e01b815260040160405180910390fd5b6000610c0d612105565b60008b8152600491909101602090815260408083206001600160a01b038d168452909152812090898985818110610c4657610c46613b28565b90506020020135815260200190815260200160002060006101000a81548160ff021916908315150217905550876001600160a01b03166323b872dd30338a8a86818110610c9557610c95613b28565b6040516001600160e01b031960e088901b1681526001600160a01b03958616600482015294909316602485015250602090910201356044820152606401600060405180830381600087803b158015610cec57600080fd5b505af1158015610d00573d6000803e3d6000fd5b50505050610d0c612105565b60008a8152600491909101602090815260408083206001600160a01b038c168452909152812090888884818110610d4557610d45613b28565b60209081029290920135835250810191909152604001600020805460ff19169055600101610b96565b50600085610d7a612105565b60070154610d889190613a7b565b9050600086610d95612105565b60080154610da39190613a7b565b905060008083610db1612105565b60008e815260059190910160205260409020541115610df35783610dd3612105565b60008e81526005919091016020526040902054610df09190613b3e565b91505b82610dfc612105565b60008e815260069190910160205260409020541115610e3e5782610e1e612105565b60008e81526006919091016020526040902054610e3b9190613b3e565b90505b81610e47612105565b60008e8152600591909101602052604090205580610e63612105565b60060160008e81526020019081526020016000208190555060006040518060c001604052808e81526020018c8c8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050509082525060208101859052604001610ed6612105565b600701548152602001838152602001610eed612105565b600801548152509050336001600160a01b03168c6001600160a01b03168f7f8e8823d9b5c6f5ebda354c473fe0a4b30300c54dc7ca47de51363766a2a1eb1784604051610f3a9190613bd3565b60405180910390a45050505050506001610f5f600080516020613f7083398151915290565b555050505050505050565b6000610f74612105565b600092835260020160205250604090205490565b610f906120ac565b7fb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031a610fbc828483613c2c565b50610fc561080c565b15611015577f6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c600180610ff6612565565b6110009190613b3e565b604080519283526020830191909152016107e2565b5050565b600080611024612105565b6000848152600591909101602052604090205461103f612105565b60009485526006016020526040909320549293915050565b60006106518261256f565b606061106c612603565b905090565b6110796120ac565b67ffffffffffffffff8111156110a55760405163b43e913760e01b815260048101829052602401610a11565b80600080516020613fb0833981519152556040518181527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c90602001610a98565b60006001600160a01b03821661110f576040516323d3ad8160e21b815260040160405180910390fd5b67ffffffffffffffff611120611ed8565b6005016000846001600160a01b03166001600160a01b0316815260200190815260200160002054169050919050565b611157612129565b6111616000612622565b565b600054610100900460ff16158080156111835750600054600160ff909116105b8061119d5750303b15801561119d575060005460ff166001145b6112005760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610a11565b6000805460ff191660011790558015611223576000805461ff0019166101001790555b600080516020613ff083398151915254610100900460ff1661125857600080516020613ff08339815191525460ff161561125c565b303b155b6112ce5760405162461bcd60e51b815260206004820152603760248201527f455243373231415f5f496e697469616c697a61626c653a20636f6e747261637460448201527f20697320616c726561647920696e697469616c697a65640000000000000000006064820152608401610a11565b600080516020613ff083398151915254610100900460ff1615801561130a57600080516020613ff0833981519152805461ffff19166101011790555b6113148787612693565b61131c612753565b61132461277a565b61132e87876127a9565b6113366127d0565b61133e61280b565b611349858585612851565b801561136857600080516020613ff0833981519152805461ff00191690555b5080156113af576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b600080516020613f50833981519152546001600160a01b03163381146113f057604051636b7584e760e11b815260040160405180910390fd5b600080516020613f5083398151915280546001600160a01b0319169055604080516000815290517f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da9181900360200190a161097d81612622565b60008060006114588461290b565b92506114626120f2565b9150600080516020613fb083398151915254929491935050565b60007ff73863e3917595a7e67829ed60f0c51bf14c7a0d0de47c7b02a00abd48488e4a61090f565b6114ac6120ac565b7fb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031b6114d8828483613c2c565b507f905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac3737882826040516107e2929190613cec565b6060611514611ed8565b600301805461066f90613a31565b8161152c81611f45565b6107538383612951565b61153e6120ac565b60408051838152602081018390527f6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c91016107e2565b61157c612129565b80611585612105565b60090180546001600160a01b0319166001600160a01b039283161790556040519082169033907f2d025324f0a785e8c12d0a0d91a9caa49df4ef20ff87e0df7213a1d4f3157beb90600090a350565b836001600160a01b03811633146115ee576115ee33611f45565b6115fa858585856129ce565b5050505050565b611609612129565b80156116be57611617612105565b6001600160a01b0383166000908152600391909101602052604090205460ff16156116845760405162461bcd60e51b815260206004820152601e60248201527f4475706c6963617465416c6c6f776564506c65676564436f6e747261637400006044820152606401610a11565b600161168e612105565b6001600160a01b038416600090815260039190910160205260409020805460ff191691151591909117905561175f565b6116c6612105565b6001600160a01b0383166000908152600391909101602052604090205460ff166117325760405162461bcd60e51b815260206004820152601f60248201527f416c6c6f776564506c65676564436f6e74726163744e6f7450726573656e74006044820152606401610a11565b61173a612105565b6001600160a01b038316600090815260039190910160205260409020805460ff191690555b604051811515906001600160a01b0384169033907febdafa528849c3f6349852f55f3238b5f0f0ec565571bc32a6de61f4a27b6e1890600090a45050565b600080516020613f70833981519152546001146117e95760405162461bcd60e51b815260206004820152600a6024820152695245454e5452414e435960b01b6044820152606401610a11565b6002600080516020613f70833981519152556118048661237a565b61180e8888612a12565b60405163e985e9c560e01b81523360048201523060248201526001600160a01b0387169063e985e9c590604401602060405180830381865afa158015611858573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061187c9190613d1b565b6118995760405163bc31f0f160e01b815260040160405180910390fd5b6118ca88888888888888887f421b8352eded66e05e1120b999a305150aef8a4e6083e0477935a7732b74feac612423565b60006118d4612105565b60008981526004919091016020526040812091505b85811015611a7d576001600160a01b03881660009081526020839052604081209088888481811061191c5761191c613b28565b602090810292909201358352508101919091526040016000205460ff16156119785760405162461bcd60e51b815260206004820152600f60248201526e111d5c1b1a58d85d19541b1959d959608a1b6044820152606401610a11565b876001600160a01b03166323b872dd33308a8a8681811061199b5761199b613b28565b6040516001600160e01b031960e088901b1681526001600160a01b03958616600482015294909316602485015250602090910201356044820152606401600060405180830381600087803b1580156119f257600080fd5b505af1158015611a06573d6000803e3d6000fd5b505050506001611a14612105565b60008b8152600491909101602090815260408083206001600160a01b038d168452909152812090898985818110611a4d57611a4d613b28565b60209081029290920135835250810191909152604001600020805460ff19169115159190911790556001016118e9565b50600088600003611ad957611a90612565565b9850611a9d336001612a88565b50600188611aa9612105565b60008c8152600191909101602052604090205589611ac5612105565b60008b815260029190910160205260409020555b600086611ae4612105565b60070154611af29190613a7b565b611afa612105565b60008c81526005919091016020526040902054611b179190613d38565b9050600087611b24612105565b60080154611b329190613a7b565b611b3a612105565b60008d81526006919091016020526040902054611b579190613d38565b905081611b62612105565b60008d8152600591909101602052604090205580611b7e612105565b60060160008d81526020019081526020016000208190555060006040518060c001604052808d81526020018b8b8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050509082525060208101859052604001611bf1612105565b600701548152602001838152602001611c08612105565b600801548152509050336001600160a01b03168b6001600160a01b03168e7ffc7b82318a55b2522e1143d198ffa980407c894361a928b287d4ea366b14b90d8488604051611c57929190613d4b565b60405180910390a450505050506001610f5f600080516020613f7083398151915290565b6060611c8682611efc565b611ca357604051630a14c4b560e41b815260040160405180910390fd5b6000611cad612603565b90508051600003611ccd5760405180602001604052806000815250611cf8565b80611cd784612aa2565b604051602001611ce8929190613d6f565b6040516020818303038152906040525b9392505050565b6000600080516020613fb08339815191525b54919050565b6060600080516020613fb0833981519152610661565b6000611d37611ed8565b6001600160a01b039384166000908152600791909101602090815260408083209490951682529290925250205460ff1690565b611d72612129565b6001600160a01b038116611d9957604051633a247dd760e11b815260040160405180910390fd5b600080516020613f5083398151915280546001600160a01b0319166001600160a01b0383169081179091556040519081527f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da90602001610a98565b611dfc612129565b811580611e07575080155b15611e485760405162461bcd60e51b815260206004820152601160248201527053706163654f724c656e6774685a65726f60781b6044820152606401610a11565b81611e51612105565b6007015580611e5e612105565b6008015560408051838152602081018390527f47208e794e45f5c352c883a84f2d7136a6144e7516979bf97a13266f807c418591016107e2565b60006001600160e01b0319821663152a902d60e11b1480611ec95750632483248360e11b6001600160e01b03198316145b80610651575061065182612ae6565b7f2569078dfb4b0305704d3008e7403993ae9601b85f7ae5e742de3de8f8011c4090565b600081600111158015611f165750611f12611ed8565b5482105b80156106515750600160e01b611f2a611ed8565b60008481526004919091016020526040902054161592915050565b6daaeb6d7670e522a718067333cd4e3b1561097d57604051633185c44d60e21b81523060048201526001600160a01b03821660248201526daaeb6d7670e522a718067333cd4e9063c617113490604401602060405180830381865afa158015611fb2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fd69190613d1b565b61097d57604051633b79c77360e21b81526001600160a01b0382166004820152602401610a11565b600061200982611057565b9050336001600160a01b03821614612042576120258133611d2d565b612042576040516367d9dca160e11b815260040160405180910390fd5b8261204b611ed8565b6000848152600691909101602052604080822080546001600160a01b0319166001600160a01b0394851617905551849286811692908516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259190a4505050565b3033146120d16120ba61147c565b6001600160a01b0316336001600160a01b03161490565b1760000361116157604051635fc483c560e01b815260040160405180910390fd5b600060016120fe611ed8565b5403919050565b7f98f8583596240e040053e538e63672db2a3cad6e6b237f934724be549c06619d90565b337ff73863e3917595a7e67829ed60f0c51bf14c7a0d0de47c7b02a00abd48488e4a546001600160a01b03161461116157604051635fc483c560e01b815260040160405180910390fd5b600061217e8261256f565b9050836001600160a01b0316816001600160a01b0316146121b15760405162a1148160e81b815260040160405180910390fd5b6000806121bd84612b34565b915091506121e281876121cd3390565b6001600160a01b039081169116811491141790565b61220d576121f08633611d2d565b61220d57604051632ce44b5f60e11b815260040160405180910390fd5b6001600160a01b03851661223457604051633a954ecd60e21b815260040160405180910390fd5b801561223f57600082555b612247611ed8565b6001600160a01b0387166000908152600591909101602052604090208054600019019055612273611ed8565b6001600160a01b03861660008181526005929092016020526040909120805460010190554260a01b17600160e11b176122aa611ed8565b60008681526004919091016020526040812091909155600160e11b8416900361232057600184016122d9611ed8565b60008281526004919091016020526040812054900361231e576122fa611ed8565b54811461231e578361230a611ed8565b600083815260049190910160205260409020555b505b83856001600160a01b0316876001600160a01b0316600080516020613f9083398151915260405160405180910390a46113af565b610753838383604051806020016040528060008152506115d4565b61097d816000612b5c565b612382612105565b6001600160a01b0382166000908152600391909101602052604090205460ff16151560011461097d5760405163287915b160e11b815260040160405180910390fd5b60006123ce612105565b60008481526001919091016020526040902054905080831415612402336123f485611057565b6001600160a01b0316141590565b1760000361075357604051635fc483c560e01b815260040160405180910390fd5b600061246988338c8c8b8b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508d92508a9150612ccb9050565b9050612473612105565b6000828152600a91909101602052604090205460ff16156124a75760405163900bb2c960e01b815260040160405180910390fd5b60016124b1612105565b6000838152600a9190910160209081526040808320805460ff1916941515949094179093558251601f870182900482028101820190935285835290916125149187908790819084018382808284376000920191909152508693925050612da49050565b905061251e612105565b600901546001600160a01b0382811691161461255857604051633615713d60e21b81526001600160a01b0382166004820152602401610a11565b5050505050505050505050565b6000611d11611ed8565b600081806001116125ea57612582611ed8565b548110156125ea576000612594611ed8565b600083815260049190910160205260408120549150600160e01b821690036125e8575b80600003611cf8576125c7611ed8565b600019909201600081815260049390930160205260409092205490506125b7565b505b604051636f96cda160e11b815260040160405180910390fd5b6060600080516020613fb0833981519152600101805461066f90613a31565b7ff73863e3917595a7e67829ed60f0c51bf14c7a0d0de47c7b02a00abd48488e4a80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600080516020613ff083398151915254610100900460ff166127145760405162461bcd60e51b815260206004820152603460248201527f455243373231415f5f496e697469616c697a61626c653a20636f6e7472616374604482015273206973206e6f7420696e697469616c697a696e6760601b6064820152608401610a11565b8161271d611ed8565b6002019061272b9082613d9e565b5080612735611ed8565b600301906127439082613d9e565b50600161274e611ed8565b555050565b600054610100900460ff166111615760405162461bcd60e51b8152600401610a1190613e5e565b600054610100900460ff166127a15760405162461bcd60e51b8152600401610a1190613e5e565b611161612dc8565b600054610100900460ff166110155760405162461bcd60e51b8152600401610a1190613e5e565b600054610100900460ff166127f75760405162461bcd60e51b8152600401610a1190613e5e565b6001600080516020613f7083398151915255565b600054610100900460ff166128325760405162461bcd60e51b8152600401610a1190613e5e565b611161733cc6cdda760b79bafa08df41ecfa224f810dceb66001612df0565b600054610100900460ff166128785760405162461bcd60e51b8152600401610a1190613e5e565b82612881612105565b60090180546001600160a01b0319166001600160a01b0392909216919091179055816128ab612105565b60070155806128b8612105565b60080155466128c5612105565b600b01556128d1612f8f565b6128d9612105565b600c01556040517fad4c8d8275325862df1444d2f29d1d82d632a8ea3377c392a29f0dbc6d635a2a90600090a1505050565b600067ffffffffffffffff6040612920611ed8565b6005016000856001600160a01b03166001600160a01b0316815260200190815260200160002054901c169050919050565b8061295a611ed8565b336000818152600792909201602090815260408084206001600160a01b03881680865290835293819020805460ff19169515159590951790945592518415158152919290917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b6129d9848484610887565b6001600160a01b0383163b156108ac576129f584848484613034565b6108ac576040516368d2bf6b60e11b815260040160405180910390fd5b6000612a1c612105565b600084815260019190910160205260409020549050818114612a5157604051635fc483c560e01b815260040160405180910390fd5b81156107535733612a6183611057565b6001600160a01b03161461075357604051635fc483c560e01b815260040160405180910390fd5b61101582826040518060200160405280600081525061311f565b606060a06040510180604052602081039150506000815280825b600183039250600a81066030018353600a900480612abc5750819003601f19909101908152919050565b60006301ffc9a760e01b6001600160e01b031983161480612b1757506380ac58cd60e01b6001600160e01b03198316145b806106515750506001600160e01b031916635b5e139f60e01b1490565b6000806000612b41611ed8565b60009485526006016020525050604090912080549092909150565b6000612b678361256f565b905080600080612b7686612b34565b915091508415612bb657612b8b8184336121cd565b612bb657612b998333611d2d565b612bb657604051632ce44b5f60e11b815260040160405180910390fd5b8015612bc157600082555b6fffffffffffffffffffffffffffffffff612bda611ed8565b6001600160a01b038516600081815260059290920160205260409091208054929092019091554260a01b17600360e01b17612c13611ed8565b60008881526004919091016020526040812091909155600160e11b85169003612c895760018601612c42611ed8565b600082815260049190910160205260408120549003612c8757612c63611ed8565b548114612c875784612c73611ed8565b600083815260049190910160205260409020555b505b60405186906000906001600160a01b03861690600080516020613f90833981519152908390a4612cb7611ed8565b600190810180549091019055505050505050565b600061190160f01b612cdb613195565b838a8a8a8a8a604051602001612cf19190613ea9565b60408051601f198184030181528282528051602091820120908301979097526001600160a01b0395861690820152939092166060840152608083015260a082015260c081019190915260e081018690526101000160408051601f198184030181529082905280516020918201206001600160f01b03199094169082015260228101919091526042810191909152606201604051602081830303815290604052805190602001209050979650505050505050565b6000806000612db385856131c2565b91509150612dc081613207565b509392505050565b303b15612de75760405162dc149f60e41b815260040160405180910390fd5b61116133612622565b600054610100900460ff16612e175760405162461bcd60e51b8152600401610a1190613e5e565b6daaeb6d7670e522a718067333cd4e3b156110155760405163c3c5a54760e01b81523060048201526daaeb6d7670e522a718067333cd4e9063c3c5a547906024016020604051808303816000875af1158015612e77573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e9b9190613d1b565b611015578015612f0f57604051633e9f1edf60e11b81523060048201526001600160a01b03831660248201526daaeb6d7670e522a718067333cd4e90637d3e3dbe906044015b600060405180830381600087803b158015612efb57600080fd5b505af11580156113af573d6000803e3d6000fd5b6001600160a01b03821615612f5e5760405163a0af290360e01b81523060048201526001600160a01b03831660248201526daaeb6d7670e522a718067333cd4e9063a0af290390604401612ee1565b604051632210724360e11b81523060048201526daaeb6d7670e522a718067333cd4e90634420e48690602401612ee1565b604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527fc09b5f8e6ecd4ed1ee5281e0121f0f52e8c1ea7842745e539b385012ed62e444918101919091527fe6bbd6277e1bf288eed5e8d1780f9a50b239e86b153736bceebccf4ea79d90b360608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b604051630a85bd0160e11b81526000906001600160a01b0385169063150b7a0290613069903390899088908890600401613edf565b6020604051808303816000875af19250505080156130a4575060408051601f3d908101601f191682019092526130a191810190613f1c565b60015b613102573d8080156130d2576040519150601f19603f3d011682016040523d82523d6000602084013e6130d7565b606091505b5080516000036130fa576040516368d2bf6b60e11b815260040160405180910390fd5b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050949350505050565b6131298383613351565b6001600160a01b0383163b15610753576000613143611ed8565b5490508281035b61315d6000868380600101945086613034565b61317a576040516368d2bf6b60e11b815260040160405180910390fd5b81811061314a578161318a611ed8565b54146115fa57600080fd5b600061319f612105565b600b015446146131b15761106c612f8f565b6131b9612105565b600c0154905090565b60008082516041036131f85760208301516040840151606085015160001a6131ec87828585613468565b94509450505050613200565b506000905060025b9250929050565b600081600481111561321b5761321b613f39565b036132235750565b600181600481111561323757613237613f39565b036132845760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610a11565b600281600481111561329857613298613f39565b036132e55760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610a11565b60038160048111156132f9576132f9613f39565b0361097d5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610a11565b600061335b611ed8565b54905060008290036133805760405163b562e8dd60e01b815260040160405180910390fd5b680100000000000000018202613394611ed8565b6001600160a01b038516600081815260059290920160205260409091208054929092019091554260a01b6001841460e11b17176133cf611ed8565b600083815260049190910160205260408120919091556001600160a01b038416908383019083908390600080516020613f908339815191528180a4600183015b8181146134355780836000600080516020613f90833981519152600080a460010161340f565b508160000361345657604051622e076360e81b815260040160405180910390fd5b8061345f611ed8565b55506107539050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561349f5750600090506003613523565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156134f3573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661351c57600060019250925050613523565b9150600090505b94509492505050565b6001600160e01b03198116811461097d57600080fd5b60006020828403121561355457600080fd5b8135611cf88161352c565b60005b8381101561357a578181015183820152602001613562565b50506000910152565b6000815180845261359b81602086016020860161355f565b601f01601f19169290920160200192915050565b602081526000611cf86020830184613583565b6000602082840312156135d457600080fd5b5035919050565b6001600160a01b038116811461097d57600080fd5b6000806040838503121561360357600080fd5b823561360e816135db565b946020939093013593505050565b60008060006060848603121561363157600080fd5b833561363c816135db565b9250602084013561364c816135db565b929592945050506040919091013590565b6000806040838503121561367057600080fd5b50508035926020909101359150565b60006040828403121561369157600080fd5b50919050565b6000602082840312156136a957600080fd5b8135611cf8816135db565b60008083601f8401126136c657600080fd5b50813567ffffffffffffffff8111156136de57600080fd5b60208301915083602082850101111561320057600080fd5b60008060008060008060008060c0898b03121561371257600080fd5b8835975060208901359650604089013561372b816135db565b9550606089013567ffffffffffffffff8082111561374857600080fd5b818b0191508b601f83011261375c57600080fd5b81358181111561376b57600080fd5b8c60208260051b850101111561378057600080fd5b6020830197508096505060808b0135945060a08b01359150808211156137a557600080fd5b506137b28b828c016136b4565b999c989b5096995094979396929594505050565b600080602083850312156137d957600080fd5b823567ffffffffffffffff8111156137f057600080fd5b6137fc858286016136b4565b90969095509350505050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff8084111561383957613839613808565b604051601f8501601f19908116603f0116810190828211818310171561386157613861613808565b8160405280935085815286868601111561387a57600080fd5b858560208301376000602087830101525050509392505050565b600082601f8301126138a557600080fd5b611cf88383356020850161381e565b600080600080600060a086880312156138cc57600080fd5b853567ffffffffffffffff808211156138e457600080fd5b6138f089838a01613894565b9650602088013591508082111561390657600080fd5b5061391388828901613894565b9450506040860135613924816135db565b94979396509394606081013594506080013592915050565b801515811461097d57600080fd5b6000806040838503121561395d57600080fd5b8235613968816135db565b915060208301356139788161393c565b809150509250929050565b6000806000806080858703121561399957600080fd5b84356139a4816135db565b935060208501356139b4816135db565b925060408501359150606085013567ffffffffffffffff8111156139d757600080fd5b8501601f810187136139e857600080fd5b6139f78782356020840161381e565b91505092959194509250565b60008060408385031215613a1657600080fd5b8235613a21816135db565b91506020830135613978816135db565b600181811c90821680613a4557607f821691505b60208210810361369157634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b808202811582820484141761065157610651613a65565b600082613aaf57634e487b7160e01b600052601260045260246000fd5b500490565b6001600160601b038116811461097d57600080fd5b600060208284031215613adb57600080fd5b8135611cf881613ab4565b8135613af1816135db565b81546001600160a01b03199081166001600160a01b039290921691821783556020840135613b1e81613ab4565b60a01b1617905550565b634e487b7160e01b600052603260045260246000fd5b8181038181111561065157610651613a65565b600060c083018251845260208084015160c08287015282815180855260e0880191508383019450600092505b80831015613b9d5784518252938301936001929092019190830190613b7d565b5060408601516040880152606086015160608801526080860151608088015260a086015160a08801528094505050505092915050565b602081526000611cf86020830184613b51565b601f82111561075357600081815260208120601f850160051c81016020861015613c0d5750805b601f850160051c820191505b818110156113af57828155600101613c19565b67ffffffffffffffff831115613c4457613c44613808565b613c5883613c528354613a31565b83613be6565b6000601f841160018114613c8c5760008515613c745750838201355b600019600387901b1c1916600186901b1783556115fa565b600083815260209020601f19861690835b82811015613cbd5786850135825560209485019460019092019101613c9d565b5086821015613cda5760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b600060208284031215613d2d57600080fd5b8151611cf88161393c565b8082018082111561065157610651613a65565b604081526000613d5e6040830185613b51565b905082151560208301529392505050565b60008351613d8181846020880161355f565b835190830190613d9581836020880161355f565b01949350505050565b815167ffffffffffffffff811115613db857613db8613808565b613dcc81613dc68454613a31565b84613be6565b602080601f831160018114613e015760008415613de95750858301515b600019600386901b1c1916600185901b1785556113af565b600085815260208120601f198616915b82811015613e3057888601518255948401946001909101908401613e11565b5085821015613e4e5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b815160009082906020808601845b83811015613ed357815185529382019390820190600101613eb7565b50929695505050505050565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090613f1290830184613583565b9695505050505050565b600060208284031215613f2e57600080fd5b8151611cf88161352c565b634e487b7160e01b600052602160045260246000fdfef73863e3917595a7e67829ed60f0c51bf14c7a0d0de47c7b02a00abd48488e4bd59f8a8c0d1463371c77782499276e5cbe466fd192ada543ceaea0a36604c1f2ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a40319b847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031dee151c8401928dc223602bb187aff91b9a56c7cae5476ef1b3287b085a16c85fa264697066735822122023f6d4e5541490a42e1da5b4454ce594bde35081121e2e14dcb5c6c7a5ba63fa64736f6c63430008110033",
}

// ClubABI is the input ABI used to generate the binding from.
// Deprecated: Use ClubMetaData.ABI instead.
var ClubABI = ClubMetaData.ABI

// ClubBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ClubMetaData.Bin instead.
var ClubBin = ClubMetaData.Bin

// DeployClub deploys a new Ethereum contract, binding an instance of Club to it.
func DeployClub(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Club, error) {
	parsed, err := ClubMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ClubBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Club{ClubCaller: ClubCaller{contract: contract}, ClubTransactor: ClubTransactor{contract: contract}, ClubFilterer: ClubFilterer{contract: contract}}, nil
}

// Club is an auto generated Go binding around an Ethereum contract.
type Club struct {
	ClubCaller     // Read-only binding to the contract
	ClubTransactor // Write-only binding to the contract
	ClubFilterer   // Log filterer for contract events
}

// ClubCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClubSession struct {
	Contract     *Club             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClubCallerSession struct {
	Contract *ClubCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ClubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClubTransactorSession struct {
	Contract     *ClubTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClubRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClubRaw struct {
	Contract *Club // Generic contract binding to access the raw methods on
}

// ClubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClubCallerRaw struct {
	Contract *ClubCaller // Generic read-only contract binding to access the raw methods on
}

// ClubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClubTransactorRaw struct {
	Contract *ClubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClub creates a new instance of Club, bound to a specific deployed contract.
func NewClub(address common.Address, backend bind.ContractBackend) (*Club, error) {
	contract, err := bindClub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Club{ClubCaller: ClubCaller{contract: contract}, ClubTransactor: ClubTransactor{contract: contract}, ClubFilterer: ClubFilterer{contract: contract}}, nil
}

// NewClubCaller creates a new read-only instance of Club, bound to a specific deployed contract.
func NewClubCaller(address common.Address, caller bind.ContractCaller) (*ClubCaller, error) {
	contract, err := bindClub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClubCaller{contract: contract}, nil
}

// NewClubTransactor creates a new write-only instance of Club, bound to a specific deployed contract.
func NewClubTransactor(address common.Address, transactor bind.ContractTransactor) (*ClubTransactor, error) {
	contract, err := bindClub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClubTransactor{contract: contract}, nil
}

// NewClubFilterer creates a new log filterer instance of Club, bound to a specific deployed contract.
func NewClubFilterer(address common.Address, filterer bind.ContractFilterer) (*ClubFilterer, error) {
	contract, err := bindClub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClubFilterer{contract: contract}, nil
}

// bindClub binds a generic wrapper to an already deployed contract.
func bindClub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Club *ClubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Club.Contract.ClubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Club *ClubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Club.Contract.ClubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Club *ClubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Club.Contract.ClubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Club *ClubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Club.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Club *ClubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Club.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Club *ClubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Club.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Club *ClubCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Club *ClubSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Club.Contract.BalanceOf(&_Club.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Club *ClubCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Club.Contract.BalanceOf(&_Club.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Club *ClubCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Club *ClubSession) BaseURI() (string, error) {
	return _Club.Contract.BaseURI(&_Club.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Club *ClubCallerSession) BaseURI() (string, error) {
	return _Club.Contract.BaseURI(&_Club.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Club *ClubCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Club *ClubSession) ContractURI() (string, error) {
	return _Club.Contract.ContractURI(&_Club.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Club *ClubCallerSession) ContractURI() (string, error) {
	return _Club.Contract.ContractURI(&_Club.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Club *ClubCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Club *ClubSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Club.Contract.GetApproved(&_Club.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Club *ClubCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Club.Contract.GetApproved(&_Club.CallOpts, tokenId)
}

// GetBurnAddress is a free data retrieval call binding the contract method 0x38b39d29.
//
// Solidity: function getBurnAddress() view returns(address)
func (_Club *ClubCaller) GetBurnAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "getBurnAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBurnAddress is a free data retrieval call binding the contract method 0x38b39d29.
//
// Solidity: function getBurnAddress() view returns(address)
func (_Club *ClubSession) GetBurnAddress() (common.Address, error) {
	return _Club.Contract.GetBurnAddress(&_Club.CallOpts)
}

// GetBurnAddress is a free data retrieval call binding the contract method 0x38b39d29.
//
// Solidity: function getBurnAddress() view returns(address)
func (_Club *ClubCallerSession) GetBurnAddress() (common.Address, error) {
	return _Club.Contract.GetBurnAddress(&_Club.CallOpts)
}

// GetClubId is a free data retrieval call binding the contract method 0x530c3de7.
//
// Solidity: function getClubId(uint256 tokenId) view returns(uint256)
func (_Club *ClubCaller) GetClubId(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "getClubId", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClubId is a free data retrieval call binding the contract method 0x530c3de7.
//
// Solidity: function getClubId(uint256 tokenId) view returns(uint256)
func (_Club *ClubSession) GetClubId(tokenId *big.Int) (*big.Int, error) {
	return _Club.Contract.GetClubId(&_Club.CallOpts, tokenId)
}

// GetClubId is a free data retrieval call binding the contract method 0x530c3de7.
//
// Solidity: function getClubId(uint256 tokenId) view returns(uint256)
func (_Club *ClubCallerSession) GetClubId(tokenId *big.Int) (*big.Int, error) {
	return _Club.Contract.GetClubId(&_Club.CallOpts, tokenId)
}

// GetClubTokenId is a free data retrieval call binding the contract method 0x0dfa40b7.
//
// Solidity: function getClubTokenId(uint256 clubId) view returns(uint256)
func (_Club *ClubCaller) GetClubTokenId(opts *bind.CallOpts, clubId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "getClubTokenId", clubId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClubTokenId is a free data retrieval call binding the contract method 0x0dfa40b7.
//
// Solidity: function getClubTokenId(uint256 clubId) view returns(uint256)
func (_Club *ClubSession) GetClubTokenId(clubId *big.Int) (*big.Int, error) {
	return _Club.Contract.GetClubTokenId(&_Club.CallOpts, clubId)
}

// GetClubTokenId is a free data retrieval call binding the contract method 0x0dfa40b7.
//
// Solidity: function getClubTokenId(uint256 clubId) view returns(uint256)
func (_Club *ClubCallerSession) GetClubTokenId(clubId *big.Int) (*big.Int, error) {
	return _Club.Contract.GetClubTokenId(&_Club.CallOpts, clubId)
}

// GetMintStats is a free data retrieval call binding the contract method 0x840e15d4.
//
// Solidity: function getMintStats(address minter) view returns(uint256 minterNumMinted, uint256 currentTotalSupply, uint256 maxSupply)
func (_Club *ClubCaller) GetMintStats(opts *bind.CallOpts, minter common.Address) (struct {
	MinterNumMinted    *big.Int
	CurrentTotalSupply *big.Int
	MaxSupply          *big.Int
}, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "getMintStats", minter)

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
func (_Club *ClubSession) GetMintStats(minter common.Address) (struct {
	MinterNumMinted    *big.Int
	CurrentTotalSupply *big.Int
	MaxSupply          *big.Int
}, error) {
	return _Club.Contract.GetMintStats(&_Club.CallOpts, minter)
}

// GetMintStats is a free data retrieval call binding the contract method 0x840e15d4.
//
// Solidity: function getMintStats(address minter) view returns(uint256 minterNumMinted, uint256 currentTotalSupply, uint256 maxSupply)
func (_Club *ClubCallerSession) GetMintStats(minter common.Address) (struct {
	MinterNumMinted    *big.Int
	CurrentTotalSupply *big.Int
	MaxSupply          *big.Int
}, error) {
	return _Club.Contract.GetMintStats(&_Club.CallOpts, minter)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Club *ClubCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Club *ClubSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Club.Contract.IsApprovedForAll(&_Club.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Club *ClubCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Club.Contract.IsApprovedForAll(&_Club.CallOpts, owner, operator)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Club *ClubCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Club *ClubSession) MaxSupply() (*big.Int, error) {
	return _Club.Contract.MaxSupply(&_Club.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Club *ClubCallerSession) MaxSupply() (*big.Int, error) {
	return _Club.Contract.MaxSupply(&_Club.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Club *ClubCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Club *ClubSession) Name() (string, error) {
	return _Club.Contract.Name(&_Club.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Club *ClubCallerSession) Name() (string, error) {
	return _Club.Contract.Name(&_Club.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Club *ClubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Club *ClubSession) Owner() (common.Address, error) {
	return _Club.Contract.Owner(&_Club.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Club *ClubCallerSession) Owner() (common.Address, error) {
	return _Club.Contract.Owner(&_Club.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Club *ClubCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Club *ClubSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Club.Contract.OwnerOf(&_Club.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Club *ClubCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Club.Contract.OwnerOf(&_Club.CallOpts, tokenId)
}

// ProvenanceHash is a free data retrieval call binding the contract method 0xc6ab67a3.
//
// Solidity: function provenanceHash() view returns(bytes32)
func (_Club *ClubCaller) ProvenanceHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "provenanceHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvenanceHash is a free data retrieval call binding the contract method 0xc6ab67a3.
//
// Solidity: function provenanceHash() view returns(bytes32)
func (_Club *ClubSession) ProvenanceHash() ([32]byte, error) {
	return _Club.Contract.ProvenanceHash(&_Club.CallOpts)
}

// ProvenanceHash is a free data retrieval call binding the contract method 0xc6ab67a3.
//
// Solidity: function provenanceHash() view returns(bytes32)
func (_Club *ClubCallerSession) ProvenanceHash() ([32]byte, error) {
	return _Club.Contract.ProvenanceHash(&_Club.CallOpts)
}

// RoyaltyAddress is a free data retrieval call binding the contract method 0xad2f852a.
//
// Solidity: function royaltyAddress() view returns(address)
func (_Club *ClubCaller) RoyaltyAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "royaltyAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyAddress is a free data retrieval call binding the contract method 0xad2f852a.
//
// Solidity: function royaltyAddress() view returns(address)
func (_Club *ClubSession) RoyaltyAddress() (common.Address, error) {
	return _Club.Contract.RoyaltyAddress(&_Club.CallOpts)
}

// RoyaltyAddress is a free data retrieval call binding the contract method 0xad2f852a.
//
// Solidity: function royaltyAddress() view returns(address)
func (_Club *ClubCallerSession) RoyaltyAddress() (common.Address, error) {
	return _Club.Contract.RoyaltyAddress(&_Club.CallOpts)
}

// RoyaltyBasisPoints is a free data retrieval call binding the contract method 0x42260b5d.
//
// Solidity: function royaltyBasisPoints() view returns(uint256)
func (_Club *ClubCaller) RoyaltyBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "royaltyBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoyaltyBasisPoints is a free data retrieval call binding the contract method 0x42260b5d.
//
// Solidity: function royaltyBasisPoints() view returns(uint256)
func (_Club *ClubSession) RoyaltyBasisPoints() (*big.Int, error) {
	return _Club.Contract.RoyaltyBasisPoints(&_Club.CallOpts)
}

// RoyaltyBasisPoints is a free data retrieval call binding the contract method 0x42260b5d.
//
// Solidity: function royaltyBasisPoints() view returns(uint256)
func (_Club *ClubCallerSession) RoyaltyBasisPoints() (*big.Int, error) {
	return _Club.Contract.RoyaltyBasisPoints(&_Club.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Club *ClubCaller) RoyaltyInfo(opts *bind.CallOpts, arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "royaltyInfo", arg0, _salePrice)

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
func (_Club *ClubSession) RoyaltyInfo(arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Club.Contract.RoyaltyInfo(&_Club.CallOpts, arg0, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Club *ClubCallerSession) RoyaltyInfo(arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Club.Contract.RoyaltyInfo(&_Club.CallOpts, arg0, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Club *ClubCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Club *ClubSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Club.Contract.SupportsInterface(&_Club.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Club *ClubCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Club.Contract.SupportsInterface(&_Club.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Club *ClubCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Club *ClubSession) Symbol() (string, error) {
	return _Club.Contract.Symbol(&_Club.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Club *ClubCallerSession) Symbol() (string, error) {
	return _Club.Contract.Symbol(&_Club.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Club *ClubCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Club *ClubSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Club.Contract.TokenURI(&_Club.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Club *ClubCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Club.Contract.TokenURI(&_Club.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Club *ClubCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Club.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Club *ClubSession) TotalSupply() (*big.Int, error) {
	return _Club.Contract.TotalSupply(&_Club.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Club *ClubCallerSession) TotalSupply() (*big.Int, error) {
	return _Club.Contract.TotalSupply(&_Club.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Club *ClubTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Club *ClubSession) AcceptOwnership() (*types.Transaction, error) {
	return _Club.Contract.AcceptOwnership(&_Club.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Club *ClubTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Club.Contract.AcceptOwnership(&_Club.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Club *ClubTransactor) Approve(opts *bind.TransactOpts, operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "approve", operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Club *ClubSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Club.Contract.Approve(&_Club.TransactOpts, operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Club *ClubTransactorSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Club.Contract.Approve(&_Club.TransactOpts, operator, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Club *ClubTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Club *ClubSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Club.Contract.Burn(&_Club.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Club *ClubTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Club.Contract.Burn(&_Club.TransactOpts, tokenId)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_Club *ClubTransactor) CancelOwnershipTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "cancelOwnershipTransfer")
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_Club *ClubSession) CancelOwnershipTransfer() (*types.Transaction, error) {
	return _Club.Contract.CancelOwnershipTransfer(&_Club.TransactOpts)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_Club *ClubTransactorSession) CancelOwnershipTransfer() (*types.Transaction, error) {
	return _Club.Contract.CancelOwnershipTransfer(&_Club.TransactOpts)
}

// EmitBatchMetadataUpdate is a paid mutator transaction binding the contract method 0xa4830114.
//
// Solidity: function emitBatchMetadataUpdate(uint256 fromTokenId, uint256 toTokenId) returns()
func (_Club *ClubTransactor) EmitBatchMetadataUpdate(opts *bind.TransactOpts, fromTokenId *big.Int, toTokenId *big.Int) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "emitBatchMetadataUpdate", fromTokenId, toTokenId)
}

// EmitBatchMetadataUpdate is a paid mutator transaction binding the contract method 0xa4830114.
//
// Solidity: function emitBatchMetadataUpdate(uint256 fromTokenId, uint256 toTokenId) returns()
func (_Club *ClubSession) EmitBatchMetadataUpdate(fromTokenId *big.Int, toTokenId *big.Int) (*types.Transaction, error) {
	return _Club.Contract.EmitBatchMetadataUpdate(&_Club.TransactOpts, fromTokenId, toTokenId)
}

// EmitBatchMetadataUpdate is a paid mutator transaction binding the contract method 0xa4830114.
//
// Solidity: function emitBatchMetadataUpdate(uint256 fromTokenId, uint256 toTokenId) returns()
func (_Club *ClubTransactorSession) EmitBatchMetadataUpdate(fromTokenId *big.Int, toTokenId *big.Int) (*types.Transaction, error) {
	return _Club.Contract.EmitBatchMetadataUpdate(&_Club.TransactOpts, fromTokenId, toTokenId)
}

// GetClubSpace is a paid mutator transaction binding the contract method 0x5c405681.
//
// Solidity: function getClubSpace(uint256 tokenId) returns(uint256, uint256)
func (_Club *ClubTransactor) GetClubSpace(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "getClubSpace", tokenId)
}

// GetClubSpace is a paid mutator transaction binding the contract method 0x5c405681.
//
// Solidity: function getClubSpace(uint256 tokenId) returns(uint256, uint256)
func (_Club *ClubSession) GetClubSpace(tokenId *big.Int) (*types.Transaction, error) {
	return _Club.Contract.GetClubSpace(&_Club.TransactOpts, tokenId)
}

// GetClubSpace is a paid mutator transaction binding the contract method 0x5c405681.
//
// Solidity: function getClubSpace(uint256 tokenId) returns(uint256, uint256)
func (_Club *ClubTransactorSession) GetClubSpace(tokenId *big.Int) (*types.Transaction, error) {
	return _Club.Contract.GetClubSpace(&_Club.TransactOpts, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x74823132.
//
// Solidity: function initialize(string name, string symbol, address plegeSigner, uint256 defaultSpace, uint256 defaultVideoLength) returns()
func (_Club *ClubTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, plegeSigner common.Address, defaultSpace *big.Int, defaultVideoLength *big.Int) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "initialize", name, symbol, plegeSigner, defaultSpace, defaultVideoLength)
}

// Initialize is a paid mutator transaction binding the contract method 0x74823132.
//
// Solidity: function initialize(string name, string symbol, address plegeSigner, uint256 defaultSpace, uint256 defaultVideoLength) returns()
func (_Club *ClubSession) Initialize(name string, symbol string, plegeSigner common.Address, defaultSpace *big.Int, defaultVideoLength *big.Int) (*types.Transaction, error) {
	return _Club.Contract.Initialize(&_Club.TransactOpts, name, symbol, plegeSigner, defaultSpace, defaultVideoLength)
}

// Initialize is a paid mutator transaction binding the contract method 0x74823132.
//
// Solidity: function initialize(string name, string symbol, address plegeSigner, uint256 defaultSpace, uint256 defaultVideoLength) returns()
func (_Club *ClubTransactorSession) Initialize(name string, symbol string, plegeSigner common.Address, defaultSpace *big.Int, defaultVideoLength *big.Int) (*types.Transaction, error) {
	return _Club.Contract.Initialize(&_Club.TransactOpts, name, symbol, plegeSigner, defaultSpace, defaultVideoLength)
}

// Plege is a paid mutator transaction binding the contract method 0xc6938bc2.
//
// Solidity: function plege(uint256 clubId, uint256 tokenId, address nftContract, uint256[] plegeToken, uint256 salt, bytes signature) returns()
func (_Club *ClubTransactor) Plege(opts *bind.TransactOpts, clubId *big.Int, tokenId *big.Int, nftContract common.Address, plegeToken []*big.Int, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "plege", clubId, tokenId, nftContract, plegeToken, salt, signature)
}

// Plege is a paid mutator transaction binding the contract method 0xc6938bc2.
//
// Solidity: function plege(uint256 clubId, uint256 tokenId, address nftContract, uint256[] plegeToken, uint256 salt, bytes signature) returns()
func (_Club *ClubSession) Plege(clubId *big.Int, tokenId *big.Int, nftContract common.Address, plegeToken []*big.Int, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Club.Contract.Plege(&_Club.TransactOpts, clubId, tokenId, nftContract, plegeToken, salt, signature)
}

// Plege is a paid mutator transaction binding the contract method 0xc6938bc2.
//
// Solidity: function plege(uint256 clubId, uint256 tokenId, address nftContract, uint256[] plegeToken, uint256 salt, bytes signature) returns()
func (_Club *ClubTransactorSession) Plege(clubId *big.Int, tokenId *big.Int, nftContract common.Address, plegeToken []*big.Int, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Club.Contract.Plege(&_Club.TransactOpts, clubId, tokenId, nftContract, plegeToken, salt, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Club *ClubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Club *ClubSession) RenounceOwnership() (*types.Transaction, error) {
	return _Club.Contract.RenounceOwnership(&_Club.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Club *ClubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Club.Contract.RenounceOwnership(&_Club.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Club *ClubTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Club *ClubSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Club.Contract.SafeTransferFrom(&_Club.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Club *ClubTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Club.Contract.SafeTransferFrom(&_Club.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Club *ClubTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Club *ClubSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Club.Contract.SafeTransferFrom0(&_Club.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Club *ClubTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Club.Contract.SafeTransferFrom0(&_Club.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Club *ClubTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Club *ClubSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Club.Contract.SetApprovalForAll(&_Club.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Club *ClubTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Club.Contract.SetApprovalForAll(&_Club.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Club *ClubTransactor) SetBaseURI(opts *bind.TransactOpts, newBaseURI string) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "setBaseURI", newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Club *ClubSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Club.Contract.SetBaseURI(&_Club.TransactOpts, newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Club *ClubTransactorSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Club.Contract.SetBaseURI(&_Club.TransactOpts, newBaseURI)
}

// SetBurnAddress is a paid mutator transaction binding the contract method 0x4b0e7216.
//
// Solidity: function setBurnAddress(address newBurnAddress) returns()
func (_Club *ClubTransactor) SetBurnAddress(opts *bind.TransactOpts, newBurnAddress common.Address) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "setBurnAddress", newBurnAddress)
}

// SetBurnAddress is a paid mutator transaction binding the contract method 0x4b0e7216.
//
// Solidity: function setBurnAddress(address newBurnAddress) returns()
func (_Club *ClubSession) SetBurnAddress(newBurnAddress common.Address) (*types.Transaction, error) {
	return _Club.Contract.SetBurnAddress(&_Club.TransactOpts, newBurnAddress)
}

// SetBurnAddress is a paid mutator transaction binding the contract method 0x4b0e7216.
//
// Solidity: function setBurnAddress(address newBurnAddress) returns()
func (_Club *ClubTransactorSession) SetBurnAddress(newBurnAddress common.Address) (*types.Transaction, error) {
	return _Club.Contract.SetBurnAddress(&_Club.TransactOpts, newBurnAddress)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_Club *ClubTransactor) SetContractURI(opts *bind.TransactOpts, newContractURI string) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "setContractURI", newContractURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_Club *ClubSession) SetContractURI(newContractURI string) (*types.Transaction, error) {
	return _Club.Contract.SetContractURI(&_Club.TransactOpts, newContractURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_Club *ClubTransactorSession) SetContractURI(newContractURI string) (*types.Transaction, error) {
	return _Club.Contract.SetContractURI(&_Club.TransactOpts, newContractURI)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_Club *ClubTransactor) SetMaxSupply(opts *bind.TransactOpts, newMaxSupply *big.Int) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "setMaxSupply", newMaxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_Club *ClubSession) SetMaxSupply(newMaxSupply *big.Int) (*types.Transaction, error) {
	return _Club.Contract.SetMaxSupply(&_Club.TransactOpts, newMaxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_Club *ClubTransactorSession) SetMaxSupply(newMaxSupply *big.Int) (*types.Transaction, error) {
	return _Club.Contract.SetMaxSupply(&_Club.TransactOpts, newMaxSupply)
}

// SetProvenanceHash is a paid mutator transaction binding the contract method 0x099b6bfa.
//
// Solidity: function setProvenanceHash(bytes32 newProvenanceHash) returns()
func (_Club *ClubTransactor) SetProvenanceHash(opts *bind.TransactOpts, newProvenanceHash [32]byte) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "setProvenanceHash", newProvenanceHash)
}

// SetProvenanceHash is a paid mutator transaction binding the contract method 0x099b6bfa.
//
// Solidity: function setProvenanceHash(bytes32 newProvenanceHash) returns()
func (_Club *ClubSession) SetProvenanceHash(newProvenanceHash [32]byte) (*types.Transaction, error) {
	return _Club.Contract.SetProvenanceHash(&_Club.TransactOpts, newProvenanceHash)
}

// SetProvenanceHash is a paid mutator transaction binding the contract method 0x099b6bfa.
//
// Solidity: function setProvenanceHash(bytes32 newProvenanceHash) returns()
func (_Club *ClubTransactorSession) SetProvenanceHash(newProvenanceHash [32]byte) (*types.Transaction, error) {
	return _Club.Contract.SetProvenanceHash(&_Club.TransactOpts, newProvenanceHash)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_Club *ClubTransactor) SetRoyaltyInfo(opts *bind.TransactOpts, newInfo ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "setRoyaltyInfo", newInfo)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_Club *ClubSession) SetRoyaltyInfo(newInfo ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo) (*types.Transaction, error) {
	return _Club.Contract.SetRoyaltyInfo(&_Club.TransactOpts, newInfo)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_Club *ClubTransactorSession) SetRoyaltyInfo(newInfo ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo) (*types.Transaction, error) {
	return _Club.Contract.SetRoyaltyInfo(&_Club.TransactOpts, newInfo)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Club *ClubTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Club *ClubSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Club.Contract.TransferFrom(&_Club.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Club *ClubTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Club.Contract.TransferFrom(&_Club.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newPotentialOwner) returns()
func (_Club *ClubTransactor) TransferOwnership(opts *bind.TransactOpts, newPotentialOwner common.Address) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "transferOwnership", newPotentialOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newPotentialOwner) returns()
func (_Club *ClubSession) TransferOwnership(newPotentialOwner common.Address) (*types.Transaction, error) {
	return _Club.Contract.TransferOwnership(&_Club.TransactOpts, newPotentialOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newPotentialOwner) returns()
func (_Club *ClubTransactorSession) TransferOwnership(newPotentialOwner common.Address) (*types.Transaction, error) {
	return _Club.Contract.TransferOwnership(&_Club.TransactOpts, newPotentialOwner)
}

// Unplege is a paid mutator transaction binding the contract method 0x4c182abc.
//
// Solidity: function unplege(uint256 clubId, uint256 tokenId, address nftContract, uint256[] plegedToken, uint256 salt, bytes signature) returns()
func (_Club *ClubTransactor) Unplege(opts *bind.TransactOpts, clubId *big.Int, tokenId *big.Int, nftContract common.Address, plegedToken []*big.Int, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "unplege", clubId, tokenId, nftContract, plegedToken, salt, signature)
}

// Unplege is a paid mutator transaction binding the contract method 0x4c182abc.
//
// Solidity: function unplege(uint256 clubId, uint256 tokenId, address nftContract, uint256[] plegedToken, uint256 salt, bytes signature) returns()
func (_Club *ClubSession) Unplege(clubId *big.Int, tokenId *big.Int, nftContract common.Address, plegedToken []*big.Int, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Club.Contract.Unplege(&_Club.TransactOpts, clubId, tokenId, nftContract, plegedToken, salt, signature)
}

// Unplege is a paid mutator transaction binding the contract method 0x4c182abc.
//
// Solidity: function unplege(uint256 clubId, uint256 tokenId, address nftContract, uint256[] plegedToken, uint256 salt, bytes signature) returns()
func (_Club *ClubTransactorSession) Unplege(clubId *big.Int, tokenId *big.Int, nftContract common.Address, plegedToken []*big.Int, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Club.Contract.Unplege(&_Club.TransactOpts, clubId, tokenId, nftContract, plegedToken, salt, signature)
}

// UpdateAllowedPledgeToken is a paid mutator transaction binding the contract method 0xb8b52cc1.
//
// Solidity: function updateAllowedPledgeToken(address nftContract, bool allowed) returns()
func (_Club *ClubTransactor) UpdateAllowedPledgeToken(opts *bind.TransactOpts, nftContract common.Address, allowed bool) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "updateAllowedPledgeToken", nftContract, allowed)
}

// UpdateAllowedPledgeToken is a paid mutator transaction binding the contract method 0xb8b52cc1.
//
// Solidity: function updateAllowedPledgeToken(address nftContract, bool allowed) returns()
func (_Club *ClubSession) UpdateAllowedPledgeToken(nftContract common.Address, allowed bool) (*types.Transaction, error) {
	return _Club.Contract.UpdateAllowedPledgeToken(&_Club.TransactOpts, nftContract, allowed)
}

// UpdateAllowedPledgeToken is a paid mutator transaction binding the contract method 0xb8b52cc1.
//
// Solidity: function updateAllowedPledgeToken(address nftContract, bool allowed) returns()
func (_Club *ClubTransactorSession) UpdateAllowedPledgeToken(nftContract common.Address, allowed bool) (*types.Transaction, error) {
	return _Club.Contract.UpdateAllowedPledgeToken(&_Club.TransactOpts, nftContract, allowed)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address signer) returns()
func (_Club *ClubTransactor) UpdateSigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "updateSigner", signer)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address signer) returns()
func (_Club *ClubSession) UpdateSigner(signer common.Address) (*types.Transaction, error) {
	return _Club.Contract.UpdateSigner(&_Club.TransactOpts, signer)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address signer) returns()
func (_Club *ClubTransactorSession) UpdateSigner(signer common.Address) (*types.Transaction, error) {
	return _Club.Contract.UpdateSigner(&_Club.TransactOpts, signer)
}

// UpdateSpaceSize is a paid mutator transaction binding the contract method 0xfc46607a.
//
// Solidity: function updateSpaceSize(uint256 size, uint256 length) returns()
func (_Club *ClubTransactor) UpdateSpaceSize(opts *bind.TransactOpts, size *big.Int, length *big.Int) (*types.Transaction, error) {
	return _Club.contract.Transact(opts, "updateSpaceSize", size, length)
}

// UpdateSpaceSize is a paid mutator transaction binding the contract method 0xfc46607a.
//
// Solidity: function updateSpaceSize(uint256 size, uint256 length) returns()
func (_Club *ClubSession) UpdateSpaceSize(size *big.Int, length *big.Int) (*types.Transaction, error) {
	return _Club.Contract.UpdateSpaceSize(&_Club.TransactOpts, size, length)
}

// UpdateSpaceSize is a paid mutator transaction binding the contract method 0xfc46607a.
//
// Solidity: function updateSpaceSize(uint256 size, uint256 length) returns()
func (_Club *ClubTransactorSession) UpdateSpaceSize(size *big.Int, length *big.Int) (*types.Transaction, error) {
	return _Club.Contract.UpdateSpaceSize(&_Club.TransactOpts, size, length)
}

// ClubAllowListUpdatedIterator is returned from FilterAllowListUpdated and is used to iterate over the raw logs and unpacked data for AllowListUpdated events raised by the Club contract.
type ClubAllowListUpdatedIterator struct {
	Event *ClubAllowListUpdated // Event containing the contract specifics and raw log

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
func (it *ClubAllowListUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubAllowListUpdated)
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
		it.Event = new(ClubAllowListUpdated)
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
func (it *ClubAllowListUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubAllowListUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubAllowListUpdated represents a AllowListUpdated event raised by the Club contract.
type ClubAllowListUpdated struct {
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
func (_Club *ClubFilterer) FilterAllowListUpdated(opts *bind.FilterOpts, nftContract []common.Address, previousMerkleRoot [][32]byte, newMerkleRoot [][32]byte) (*ClubAllowListUpdatedIterator, error) {

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

	logs, sub, err := _Club.contract.FilterLogs(opts, "AllowListUpdated", nftContractRule, previousMerkleRootRule, newMerkleRootRule)
	if err != nil {
		return nil, err
	}
	return &ClubAllowListUpdatedIterator{contract: _Club.contract, event: "AllowListUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowListUpdated is a free log subscription operation binding the contract event 0xefcd7e019bc8b47d27881fd59e2619280ca5894f285950f10ab049870652efa5.
//
// Solidity: event AllowListUpdated(address indexed nftContract, bytes32 indexed previousMerkleRoot, bytes32 indexed newMerkleRoot, string[] publicKeyURI, string allowListURI)
func (_Club *ClubFilterer) WatchAllowListUpdated(opts *bind.WatchOpts, sink chan<- *ClubAllowListUpdated, nftContract []common.Address, previousMerkleRoot [][32]byte, newMerkleRoot [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Club.contract.WatchLogs(opts, "AllowListUpdated", nftContractRule, previousMerkleRootRule, newMerkleRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubAllowListUpdated)
				if err := _Club.contract.UnpackLog(event, "AllowListUpdated", log); err != nil {
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
func (_Club *ClubFilterer) ParseAllowListUpdated(log types.Log) (*ClubAllowListUpdated, error) {
	event := new(ClubAllowListUpdated)
	if err := _Club.contract.UnpackLog(event, "AllowListUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubAllowedFeeRecipientUpdatedIterator is returned from FilterAllowedFeeRecipientUpdated and is used to iterate over the raw logs and unpacked data for AllowedFeeRecipientUpdated events raised by the Club contract.
type ClubAllowedFeeRecipientUpdatedIterator struct {
	Event *ClubAllowedFeeRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *ClubAllowedFeeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubAllowedFeeRecipientUpdated)
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
		it.Event = new(ClubAllowedFeeRecipientUpdated)
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
func (it *ClubAllowedFeeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubAllowedFeeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubAllowedFeeRecipientUpdated represents a AllowedFeeRecipientUpdated event raised by the Club contract.
type ClubAllowedFeeRecipientUpdated struct {
	NftContract  common.Address
	FeeRecipient common.Address
	Allowed      bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAllowedFeeRecipientUpdated is a free log retrieval operation binding the contract event 0x6486c31f9d664e241acf94ec2541d328f6b9e97257ae16a1d887f296f879719f.
//
// Solidity: event AllowedFeeRecipientUpdated(address indexed nftContract, address indexed feeRecipient, bool indexed allowed)
func (_Club *ClubFilterer) FilterAllowedFeeRecipientUpdated(opts *bind.FilterOpts, nftContract []common.Address, feeRecipient []common.Address, allowed []bool) (*ClubAllowedFeeRecipientUpdatedIterator, error) {

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

	logs, sub, err := _Club.contract.FilterLogs(opts, "AllowedFeeRecipientUpdated", nftContractRule, feeRecipientRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &ClubAllowedFeeRecipientUpdatedIterator{contract: _Club.contract, event: "AllowedFeeRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedFeeRecipientUpdated is a free log subscription operation binding the contract event 0x6486c31f9d664e241acf94ec2541d328f6b9e97257ae16a1d887f296f879719f.
//
// Solidity: event AllowedFeeRecipientUpdated(address indexed nftContract, address indexed feeRecipient, bool indexed allowed)
func (_Club *ClubFilterer) WatchAllowedFeeRecipientUpdated(opts *bind.WatchOpts, sink chan<- *ClubAllowedFeeRecipientUpdated, nftContract []common.Address, feeRecipient []common.Address, allowed []bool) (event.Subscription, error) {

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

	logs, sub, err := _Club.contract.WatchLogs(opts, "AllowedFeeRecipientUpdated", nftContractRule, feeRecipientRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubAllowedFeeRecipientUpdated)
				if err := _Club.contract.UnpackLog(event, "AllowedFeeRecipientUpdated", log); err != nil {
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
func (_Club *ClubFilterer) ParseAllowedFeeRecipientUpdated(log types.Log) (*ClubAllowedFeeRecipientUpdated, error) {
	event := new(ClubAllowedFeeRecipientUpdated)
	if err := _Club.contract.UnpackLog(event, "AllowedFeeRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubAllowedPlegedContractUpdatedIterator is returned from FilterAllowedPlegedContractUpdated and is used to iterate over the raw logs and unpacked data for AllowedPlegedContractUpdated events raised by the Club contract.
type ClubAllowedPlegedContractUpdatedIterator struct {
	Event *ClubAllowedPlegedContractUpdated // Event containing the contract specifics and raw log

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
func (it *ClubAllowedPlegedContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubAllowedPlegedContractUpdated)
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
		it.Event = new(ClubAllowedPlegedContractUpdated)
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
func (it *ClubAllowedPlegedContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubAllowedPlegedContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubAllowedPlegedContractUpdated represents a AllowedPlegedContractUpdated event raised by the Club contract.
type ClubAllowedPlegedContractUpdated struct {
	Sender   common.Address
	NftToken common.Address
	Allowed  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAllowedPlegedContractUpdated is a free log retrieval operation binding the contract event 0xebdafa528849c3f6349852f55f3238b5f0f0ec565571bc32a6de61f4a27b6e18.
//
// Solidity: event AllowedPlegedContractUpdated(address indexed sender, address indexed nftToken, bool indexed allowed)
func (_Club *ClubFilterer) FilterAllowedPlegedContractUpdated(opts *bind.FilterOpts, sender []common.Address, nftToken []common.Address, allowed []bool) (*ClubAllowedPlegedContractUpdatedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var nftTokenRule []interface{}
	for _, nftTokenItem := range nftToken {
		nftTokenRule = append(nftTokenRule, nftTokenItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Club.contract.FilterLogs(opts, "AllowedPlegedContractUpdated", senderRule, nftTokenRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &ClubAllowedPlegedContractUpdatedIterator{contract: _Club.contract, event: "AllowedPlegedContractUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedPlegedContractUpdated is a free log subscription operation binding the contract event 0xebdafa528849c3f6349852f55f3238b5f0f0ec565571bc32a6de61f4a27b6e18.
//
// Solidity: event AllowedPlegedContractUpdated(address indexed sender, address indexed nftToken, bool indexed allowed)
func (_Club *ClubFilterer) WatchAllowedPlegedContractUpdated(opts *bind.WatchOpts, sink chan<- *ClubAllowedPlegedContractUpdated, sender []common.Address, nftToken []common.Address, allowed []bool) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var nftTokenRule []interface{}
	for _, nftTokenItem := range nftToken {
		nftTokenRule = append(nftTokenRule, nftTokenItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Club.contract.WatchLogs(opts, "AllowedPlegedContractUpdated", senderRule, nftTokenRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubAllowedPlegedContractUpdated)
				if err := _Club.contract.UnpackLog(event, "AllowedPlegedContractUpdated", log); err != nil {
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

// ParseAllowedPlegedContractUpdated is a log parse operation binding the contract event 0xebdafa528849c3f6349852f55f3238b5f0f0ec565571bc32a6de61f4a27b6e18.
//
// Solidity: event AllowedPlegedContractUpdated(address indexed sender, address indexed nftToken, bool indexed allowed)
func (_Club *ClubFilterer) ParseAllowedPlegedContractUpdated(log types.Log) (*ClubAllowedPlegedContractUpdated, error) {
	event := new(ClubAllowedPlegedContractUpdated)
	if err := _Club.contract.UnpackLog(event, "AllowedPlegedContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Club contract.
type ClubApprovalIterator struct {
	Event *ClubApproval // Event containing the contract specifics and raw log

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
func (it *ClubApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubApproval)
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
		it.Event = new(ClubApproval)
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
func (it *ClubApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubApproval represents a Approval event raised by the Club contract.
type ClubApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Club *ClubFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ClubApprovalIterator, error) {

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

	logs, sub, err := _Club.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ClubApprovalIterator{contract: _Club.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Club *ClubFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ClubApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Club.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubApproval)
				if err := _Club.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Club *ClubFilterer) ParseApproval(log types.Log) (*ClubApproval, error) {
	event := new(ClubApproval)
	if err := _Club.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Club contract.
type ClubApprovalForAllIterator struct {
	Event *ClubApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ClubApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubApprovalForAll)
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
		it.Event = new(ClubApprovalForAll)
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
func (it *ClubApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubApprovalForAll represents a ApprovalForAll event raised by the Club contract.
type ClubApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Club *ClubFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ClubApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Club.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ClubApprovalForAllIterator{contract: _Club.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Club *ClubFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ClubApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Club.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubApprovalForAll)
				if err := _Club.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Club *ClubFilterer) ParseApprovalForAll(log types.Log) (*ClubApprovalForAll, error) {
	event := new(ClubApprovalForAll)
	if err := _Club.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Club contract.
type ClubBatchMetadataUpdateIterator struct {
	Event *ClubBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ClubBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubBatchMetadataUpdate)
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
		it.Event = new(ClubBatchMetadataUpdate)
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
func (it *ClubBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Club contract.
type ClubBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Club *ClubFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*ClubBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Club.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ClubBatchMetadataUpdateIterator{contract: _Club.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Club *ClubFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ClubBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Club.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubBatchMetadataUpdate)
				if err := _Club.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_Club *ClubFilterer) ParseBatchMetadataUpdate(log types.Log) (*ClubBatchMetadataUpdate, error) {
	event := new(ClubBatchMetadataUpdate)
	if err := _Club.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubClubTokenDeployedIterator is returned from FilterClubTokenDeployed and is used to iterate over the raw logs and unpacked data for ClubTokenDeployed events raised by the Club contract.
type ClubClubTokenDeployedIterator struct {
	Event *ClubClubTokenDeployed // Event containing the contract specifics and raw log

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
func (it *ClubClubTokenDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubClubTokenDeployed)
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
		it.Event = new(ClubClubTokenDeployed)
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
func (it *ClubClubTokenDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubClubTokenDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubClubTokenDeployed represents a ClubTokenDeployed event raised by the Club contract.
type ClubClubTokenDeployed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClubTokenDeployed is a free log retrieval operation binding the contract event 0xad4c8d8275325862df1444d2f29d1d82d632a8ea3377c392a29f0dbc6d635a2a.
//
// Solidity: event ClubTokenDeployed()
func (_Club *ClubFilterer) FilterClubTokenDeployed(opts *bind.FilterOpts) (*ClubClubTokenDeployedIterator, error) {

	logs, sub, err := _Club.contract.FilterLogs(opts, "ClubTokenDeployed")
	if err != nil {
		return nil, err
	}
	return &ClubClubTokenDeployedIterator{contract: _Club.contract, event: "ClubTokenDeployed", logs: logs, sub: sub}, nil
}

// WatchClubTokenDeployed is a free log subscription operation binding the contract event 0xad4c8d8275325862df1444d2f29d1d82d632a8ea3377c392a29f0dbc6d635a2a.
//
// Solidity: event ClubTokenDeployed()
func (_Club *ClubFilterer) WatchClubTokenDeployed(opts *bind.WatchOpts, sink chan<- *ClubClubTokenDeployed) (event.Subscription, error) {

	logs, sub, err := _Club.contract.WatchLogs(opts, "ClubTokenDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubClubTokenDeployed)
				if err := _Club.contract.UnpackLog(event, "ClubTokenDeployed", log); err != nil {
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

// ParseClubTokenDeployed is a log parse operation binding the contract event 0xad4c8d8275325862df1444d2f29d1d82d632a8ea3377c392a29f0dbc6d635a2a.
//
// Solidity: event ClubTokenDeployed()
func (_Club *ClubFilterer) ParseClubTokenDeployed(log types.Log) (*ClubClubTokenDeployed, error) {
	event := new(ClubClubTokenDeployed)
	if err := _Club.contract.UnpackLog(event, "ClubTokenDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Club contract.
type ClubConsecutiveTransferIterator struct {
	Event *ClubConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *ClubConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubConsecutiveTransfer)
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
		it.Event = new(ClubConsecutiveTransfer)
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
func (it *ClubConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Club contract.
type ClubConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Club *ClubFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*ClubConsecutiveTransferIterator, error) {

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

	logs, sub, err := _Club.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ClubConsecutiveTransferIterator{contract: _Club.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Club *ClubFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *ClubConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Club.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubConsecutiveTransfer)
				if err := _Club.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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
func (_Club *ClubFilterer) ParseConsecutiveTransfer(log types.Log) (*ClubConsecutiveTransfer, error) {
	event := new(ClubConsecutiveTransfer)
	if err := _Club.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubContractURIUpdatedIterator is returned from FilterContractURIUpdated and is used to iterate over the raw logs and unpacked data for ContractURIUpdated events raised by the Club contract.
type ClubContractURIUpdatedIterator struct {
	Event *ClubContractURIUpdated // Event containing the contract specifics and raw log

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
func (it *ClubContractURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubContractURIUpdated)
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
		it.Event = new(ClubContractURIUpdated)
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
func (it *ClubContractURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubContractURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubContractURIUpdated represents a ContractURIUpdated event raised by the Club contract.
type ClubContractURIUpdated struct {
	NewContractURI string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContractURIUpdated is a free log retrieval operation binding the contract event 0x905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac37378.
//
// Solidity: event ContractURIUpdated(string newContractURI)
func (_Club *ClubFilterer) FilterContractURIUpdated(opts *bind.FilterOpts) (*ClubContractURIUpdatedIterator, error) {

	logs, sub, err := _Club.contract.FilterLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return &ClubContractURIUpdatedIterator{contract: _Club.contract, event: "ContractURIUpdated", logs: logs, sub: sub}, nil
}

// WatchContractURIUpdated is a free log subscription operation binding the contract event 0x905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac37378.
//
// Solidity: event ContractURIUpdated(string newContractURI)
func (_Club *ClubFilterer) WatchContractURIUpdated(opts *bind.WatchOpts, sink chan<- *ClubContractURIUpdated) (event.Subscription, error) {

	logs, sub, err := _Club.contract.WatchLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubContractURIUpdated)
				if err := _Club.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
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
func (_Club *ClubFilterer) ParseContractURIUpdated(log types.Log) (*ClubContractURIUpdated, error) {
	event := new(ClubContractURIUpdated)
	if err := _Club.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubCreatorPayoutAddressUpdatedIterator is returned from FilterCreatorPayoutAddressUpdated and is used to iterate over the raw logs and unpacked data for CreatorPayoutAddressUpdated events raised by the Club contract.
type ClubCreatorPayoutAddressUpdatedIterator struct {
	Event *ClubCreatorPayoutAddressUpdated // Event containing the contract specifics and raw log

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
func (it *ClubCreatorPayoutAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubCreatorPayoutAddressUpdated)
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
		it.Event = new(ClubCreatorPayoutAddressUpdated)
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
func (it *ClubCreatorPayoutAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubCreatorPayoutAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubCreatorPayoutAddressUpdated represents a CreatorPayoutAddressUpdated event raised by the Club contract.
type ClubCreatorPayoutAddressUpdated struct {
	NftContract      common.Address
	NewPayoutAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCreatorPayoutAddressUpdated is a free log retrieval operation binding the contract event 0x0c69f21751e800ea5960436c9a94370c7adbf54c733a20a025293fbbe8f16252.
//
// Solidity: event CreatorPayoutAddressUpdated(address indexed nftContract, address indexed newPayoutAddress)
func (_Club *ClubFilterer) FilterCreatorPayoutAddressUpdated(opts *bind.FilterOpts, nftContract []common.Address, newPayoutAddress []common.Address) (*ClubCreatorPayoutAddressUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var newPayoutAddressRule []interface{}
	for _, newPayoutAddressItem := range newPayoutAddress {
		newPayoutAddressRule = append(newPayoutAddressRule, newPayoutAddressItem)
	}

	logs, sub, err := _Club.contract.FilterLogs(opts, "CreatorPayoutAddressUpdated", nftContractRule, newPayoutAddressRule)
	if err != nil {
		return nil, err
	}
	return &ClubCreatorPayoutAddressUpdatedIterator{contract: _Club.contract, event: "CreatorPayoutAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchCreatorPayoutAddressUpdated is a free log subscription operation binding the contract event 0x0c69f21751e800ea5960436c9a94370c7adbf54c733a20a025293fbbe8f16252.
//
// Solidity: event CreatorPayoutAddressUpdated(address indexed nftContract, address indexed newPayoutAddress)
func (_Club *ClubFilterer) WatchCreatorPayoutAddressUpdated(opts *bind.WatchOpts, sink chan<- *ClubCreatorPayoutAddressUpdated, nftContract []common.Address, newPayoutAddress []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var newPayoutAddressRule []interface{}
	for _, newPayoutAddressItem := range newPayoutAddress {
		newPayoutAddressRule = append(newPayoutAddressRule, newPayoutAddressItem)
	}

	logs, sub, err := _Club.contract.WatchLogs(opts, "CreatorPayoutAddressUpdated", nftContractRule, newPayoutAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubCreatorPayoutAddressUpdated)
				if err := _Club.contract.UnpackLog(event, "CreatorPayoutAddressUpdated", log); err != nil {
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
func (_Club *ClubFilterer) ParseCreatorPayoutAddressUpdated(log types.Log) (*ClubCreatorPayoutAddressUpdated, error) {
	event := new(ClubCreatorPayoutAddressUpdated)
	if err := _Club.contract.UnpackLog(event, "CreatorPayoutAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubDropURIUpdatedIterator is returned from FilterDropURIUpdated and is used to iterate over the raw logs and unpacked data for DropURIUpdated events raised by the Club contract.
type ClubDropURIUpdatedIterator struct {
	Event *ClubDropURIUpdated // Event containing the contract specifics and raw log

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
func (it *ClubDropURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubDropURIUpdated)
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
		it.Event = new(ClubDropURIUpdated)
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
func (it *ClubDropURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubDropURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubDropURIUpdated represents a DropURIUpdated event raised by the Club contract.
type ClubDropURIUpdated struct {
	NftContract common.Address
	NewDropURI  string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDropURIUpdated is a free log retrieval operation binding the contract event 0xa0295608d25b3033c2e2c41cbac8746c2d08767bcfde6d47fae1ed7ba1d32150.
//
// Solidity: event DropURIUpdated(address indexed nftContract, string newDropURI)
func (_Club *ClubFilterer) FilterDropURIUpdated(opts *bind.FilterOpts, nftContract []common.Address) (*ClubDropURIUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Club.contract.FilterLogs(opts, "DropURIUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return &ClubDropURIUpdatedIterator{contract: _Club.contract, event: "DropURIUpdated", logs: logs, sub: sub}, nil
}

// WatchDropURIUpdated is a free log subscription operation binding the contract event 0xa0295608d25b3033c2e2c41cbac8746c2d08767bcfde6d47fae1ed7ba1d32150.
//
// Solidity: event DropURIUpdated(address indexed nftContract, string newDropURI)
func (_Club *ClubFilterer) WatchDropURIUpdated(opts *bind.WatchOpts, sink chan<- *ClubDropURIUpdated, nftContract []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Club.contract.WatchLogs(opts, "DropURIUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubDropURIUpdated)
				if err := _Club.contract.UnpackLog(event, "DropURIUpdated", log); err != nil {
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
func (_Club *ClubFilterer) ParseDropURIUpdated(log types.Log) (*ClubDropURIUpdated, error) {
	event := new(ClubDropURIUpdated)
	if err := _Club.contract.UnpackLog(event, "DropURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Club contract.
type ClubInitializedIterator struct {
	Event *ClubInitialized // Event containing the contract specifics and raw log

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
func (it *ClubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubInitialized)
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
		it.Event = new(ClubInitialized)
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
func (it *ClubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubInitialized represents a Initialized event raised by the Club contract.
type ClubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Club *ClubFilterer) FilterInitialized(opts *bind.FilterOpts) (*ClubInitializedIterator, error) {

	logs, sub, err := _Club.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ClubInitializedIterator{contract: _Club.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Club *ClubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ClubInitialized) (event.Subscription, error) {

	logs, sub, err := _Club.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubInitialized)
				if err := _Club.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Club *ClubFilterer) ParseInitialized(log types.Log) (*ClubInitialized, error) {
	event := new(ClubInitialized)
	if err := _Club.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubMaxSupplyUpdatedIterator is returned from FilterMaxSupplyUpdated and is used to iterate over the raw logs and unpacked data for MaxSupplyUpdated events raised by the Club contract.
type ClubMaxSupplyUpdatedIterator struct {
	Event *ClubMaxSupplyUpdated // Event containing the contract specifics and raw log

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
func (it *ClubMaxSupplyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubMaxSupplyUpdated)
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
		it.Event = new(ClubMaxSupplyUpdated)
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
func (it *ClubMaxSupplyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubMaxSupplyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubMaxSupplyUpdated represents a MaxSupplyUpdated event raised by the Club contract.
type ClubMaxSupplyUpdated struct {
	NewMaxSupply *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMaxSupplyUpdated is a free log retrieval operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 newMaxSupply)
func (_Club *ClubFilterer) FilterMaxSupplyUpdated(opts *bind.FilterOpts) (*ClubMaxSupplyUpdatedIterator, error) {

	logs, sub, err := _Club.contract.FilterLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return &ClubMaxSupplyUpdatedIterator{contract: _Club.contract, event: "MaxSupplyUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxSupplyUpdated is a free log subscription operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 newMaxSupply)
func (_Club *ClubFilterer) WatchMaxSupplyUpdated(opts *bind.WatchOpts, sink chan<- *ClubMaxSupplyUpdated) (event.Subscription, error) {

	logs, sub, err := _Club.contract.WatchLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubMaxSupplyUpdated)
				if err := _Club.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
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
func (_Club *ClubFilterer) ParseMaxSupplyUpdated(log types.Log) (*ClubMaxSupplyUpdated, error) {
	event := new(ClubMaxSupplyUpdated)
	if err := _Club.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Club contract.
type ClubOwnershipTransferredIterator struct {
	Event *ClubOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ClubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubOwnershipTransferred)
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
		it.Event = new(ClubOwnershipTransferred)
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
func (it *ClubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubOwnershipTransferred represents a OwnershipTransferred event raised by the Club contract.
type ClubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Club *ClubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ClubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Club.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ClubOwnershipTransferredIterator{contract: _Club.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Club *ClubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ClubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Club.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubOwnershipTransferred)
				if err := _Club.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Club *ClubFilterer) ParseOwnershipTransferred(log types.Log) (*ClubOwnershipTransferred, error) {
	event := new(ClubOwnershipTransferred)
	if err := _Club.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubPayerUpdatedIterator is returned from FilterPayerUpdated and is used to iterate over the raw logs and unpacked data for PayerUpdated events raised by the Club contract.
type ClubPayerUpdatedIterator struct {
	Event *ClubPayerUpdated // Event containing the contract specifics and raw log

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
func (it *ClubPayerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubPayerUpdated)
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
		it.Event = new(ClubPayerUpdated)
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
func (it *ClubPayerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubPayerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubPayerUpdated represents a PayerUpdated event raised by the Club contract.
type ClubPayerUpdated struct {
	NftContract common.Address
	Payer       common.Address
	Allowed     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayerUpdated is a free log retrieval operation binding the contract event 0x55a5cfa4bc68ffb9d833b75bf93f6d9c9aadc558dbfa587a9b5bb0ea7d5c38a3.
//
// Solidity: event PayerUpdated(address indexed nftContract, address indexed payer, bool indexed allowed)
func (_Club *ClubFilterer) FilterPayerUpdated(opts *bind.FilterOpts, nftContract []common.Address, payer []common.Address, allowed []bool) (*ClubPayerUpdatedIterator, error) {

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

	logs, sub, err := _Club.contract.FilterLogs(opts, "PayerUpdated", nftContractRule, payerRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &ClubPayerUpdatedIterator{contract: _Club.contract, event: "PayerUpdated", logs: logs, sub: sub}, nil
}

// WatchPayerUpdated is a free log subscription operation binding the contract event 0x55a5cfa4bc68ffb9d833b75bf93f6d9c9aadc558dbfa587a9b5bb0ea7d5c38a3.
//
// Solidity: event PayerUpdated(address indexed nftContract, address indexed payer, bool indexed allowed)
func (_Club *ClubFilterer) WatchPayerUpdated(opts *bind.WatchOpts, sink chan<- *ClubPayerUpdated, nftContract []common.Address, payer []common.Address, allowed []bool) (event.Subscription, error) {

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

	logs, sub, err := _Club.contract.WatchLogs(opts, "PayerUpdated", nftContractRule, payerRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubPayerUpdated)
				if err := _Club.contract.UnpackLog(event, "PayerUpdated", log); err != nil {
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
func (_Club *ClubFilterer) ParsePayerUpdated(log types.Log) (*ClubPayerUpdated, error) {
	event := new(ClubPayerUpdated)
	if err := _Club.contract.UnpackLog(event, "PayerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubPlegedUpdatedIterator is returned from FilterPlegedUpdated and is used to iterate over the raw logs and unpacked data for PlegedUpdated events raised by the Club contract.
type ClubPlegedUpdatedIterator struct {
	Event *ClubPlegedUpdated // Event containing the contract specifics and raw log

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
func (it *ClubPlegedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubPlegedUpdated)
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
		it.Event = new(ClubPlegedUpdated)
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
func (it *ClubPlegedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubPlegedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubPlegedUpdated represents a PlegedUpdated event raised by the Club contract.
type ClubPlegedUpdated struct {
	ClubId      *big.Int
	NftContract common.Address
	Sender      common.Address
	Params      ERC721ClubUpgradeableClubParams
	Minted      bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPlegedUpdated is a free log retrieval operation binding the contract event 0xfc7b82318a55b2522e1143d198ffa980407c894361a928b287d4ea366b14b90d.
//
// Solidity: event PlegedUpdated(uint256 indexed clubId, address indexed nftContract, address indexed sender, (uint256,uint256[],uint256,uint256,uint256,uint256) params, bool minted)
func (_Club *ClubFilterer) FilterPlegedUpdated(opts *bind.FilterOpts, clubId []*big.Int, nftContract []common.Address, sender []common.Address) (*ClubPlegedUpdatedIterator, error) {

	var clubIdRule []interface{}
	for _, clubIdItem := range clubId {
		clubIdRule = append(clubIdRule, clubIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Club.contract.FilterLogs(opts, "PlegedUpdated", clubIdRule, nftContractRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ClubPlegedUpdatedIterator{contract: _Club.contract, event: "PlegedUpdated", logs: logs, sub: sub}, nil
}

// WatchPlegedUpdated is a free log subscription operation binding the contract event 0xfc7b82318a55b2522e1143d198ffa980407c894361a928b287d4ea366b14b90d.
//
// Solidity: event PlegedUpdated(uint256 indexed clubId, address indexed nftContract, address indexed sender, (uint256,uint256[],uint256,uint256,uint256,uint256) params, bool minted)
func (_Club *ClubFilterer) WatchPlegedUpdated(opts *bind.WatchOpts, sink chan<- *ClubPlegedUpdated, clubId []*big.Int, nftContract []common.Address, sender []common.Address) (event.Subscription, error) {

	var clubIdRule []interface{}
	for _, clubIdItem := range clubId {
		clubIdRule = append(clubIdRule, clubIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Club.contract.WatchLogs(opts, "PlegedUpdated", clubIdRule, nftContractRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubPlegedUpdated)
				if err := _Club.contract.UnpackLog(event, "PlegedUpdated", log); err != nil {
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

// ParsePlegedUpdated is a log parse operation binding the contract event 0xfc7b82318a55b2522e1143d198ffa980407c894361a928b287d4ea366b14b90d.
//
// Solidity: event PlegedUpdated(uint256 indexed clubId, address indexed nftContract, address indexed sender, (uint256,uint256[],uint256,uint256,uint256,uint256) params, bool minted)
func (_Club *ClubFilterer) ParsePlegedUpdated(log types.Log) (*ClubPlegedUpdated, error) {
	event := new(ClubPlegedUpdated)
	if err := _Club.contract.UnpackLog(event, "PlegedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubPotentialOwnerUpdatedIterator is returned from FilterPotentialOwnerUpdated and is used to iterate over the raw logs and unpacked data for PotentialOwnerUpdated events raised by the Club contract.
type ClubPotentialOwnerUpdatedIterator struct {
	Event *ClubPotentialOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *ClubPotentialOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubPotentialOwnerUpdated)
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
		it.Event = new(ClubPotentialOwnerUpdated)
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
func (it *ClubPotentialOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubPotentialOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubPotentialOwnerUpdated represents a PotentialOwnerUpdated event raised by the Club contract.
type ClubPotentialOwnerUpdated struct {
	NewPotentialAdministrator common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterPotentialOwnerUpdated is a free log retrieval operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address newPotentialAdministrator)
func (_Club *ClubFilterer) FilterPotentialOwnerUpdated(opts *bind.FilterOpts) (*ClubPotentialOwnerUpdatedIterator, error) {

	logs, sub, err := _Club.contract.FilterLogs(opts, "PotentialOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &ClubPotentialOwnerUpdatedIterator{contract: _Club.contract, event: "PotentialOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchPotentialOwnerUpdated is a free log subscription operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address newPotentialAdministrator)
func (_Club *ClubFilterer) WatchPotentialOwnerUpdated(opts *bind.WatchOpts, sink chan<- *ClubPotentialOwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _Club.contract.WatchLogs(opts, "PotentialOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubPotentialOwnerUpdated)
				if err := _Club.contract.UnpackLog(event, "PotentialOwnerUpdated", log); err != nil {
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
func (_Club *ClubFilterer) ParsePotentialOwnerUpdated(log types.Log) (*ClubPotentialOwnerUpdated, error) {
	event := new(ClubPotentialOwnerUpdated)
	if err := _Club.contract.UnpackLog(event, "PotentialOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubProvenanceHashUpdatedIterator is returned from FilterProvenanceHashUpdated and is used to iterate over the raw logs and unpacked data for ProvenanceHashUpdated events raised by the Club contract.
type ClubProvenanceHashUpdatedIterator struct {
	Event *ClubProvenanceHashUpdated // Event containing the contract specifics and raw log

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
func (it *ClubProvenanceHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubProvenanceHashUpdated)
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
		it.Event = new(ClubProvenanceHashUpdated)
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
func (it *ClubProvenanceHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubProvenanceHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubProvenanceHashUpdated represents a ProvenanceHashUpdated event raised by the Club contract.
type ClubProvenanceHashUpdated struct {
	PreviousHash [32]byte
	NewHash      [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProvenanceHashUpdated is a free log retrieval operation binding the contract event 0x7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c.
//
// Solidity: event ProvenanceHashUpdated(bytes32 previousHash, bytes32 newHash)
func (_Club *ClubFilterer) FilterProvenanceHashUpdated(opts *bind.FilterOpts) (*ClubProvenanceHashUpdatedIterator, error) {

	logs, sub, err := _Club.contract.FilterLogs(opts, "ProvenanceHashUpdated")
	if err != nil {
		return nil, err
	}
	return &ClubProvenanceHashUpdatedIterator{contract: _Club.contract, event: "ProvenanceHashUpdated", logs: logs, sub: sub}, nil
}

// WatchProvenanceHashUpdated is a free log subscription operation binding the contract event 0x7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c.
//
// Solidity: event ProvenanceHashUpdated(bytes32 previousHash, bytes32 newHash)
func (_Club *ClubFilterer) WatchProvenanceHashUpdated(opts *bind.WatchOpts, sink chan<- *ClubProvenanceHashUpdated) (event.Subscription, error) {

	logs, sub, err := _Club.contract.WatchLogs(opts, "ProvenanceHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubProvenanceHashUpdated)
				if err := _Club.contract.UnpackLog(event, "ProvenanceHashUpdated", log); err != nil {
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
func (_Club *ClubFilterer) ParseProvenanceHashUpdated(log types.Log) (*ClubProvenanceHashUpdated, error) {
	event := new(ClubProvenanceHashUpdated)
	if err := _Club.contract.UnpackLog(event, "ProvenanceHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubPublicDropUpdatedIterator is returned from FilterPublicDropUpdated and is used to iterate over the raw logs and unpacked data for PublicDropUpdated events raised by the Club contract.
type ClubPublicDropUpdatedIterator struct {
	Event *ClubPublicDropUpdated // Event containing the contract specifics and raw log

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
func (it *ClubPublicDropUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubPublicDropUpdated)
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
		it.Event = new(ClubPublicDropUpdated)
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
func (it *ClubPublicDropUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubPublicDropUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubPublicDropUpdated represents a PublicDropUpdated event raised by the Club contract.
type ClubPublicDropUpdated struct {
	NftContract common.Address
	PublicDrop  PublicDrop
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPublicDropUpdated is a free log retrieval operation binding the contract event 0x3e30d8e1f739ea4795c481b21c23f905e938b80339305f3508e43c558e5dead3.
//
// Solidity: event PublicDropUpdated(address indexed nftContract, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop)
func (_Club *ClubFilterer) FilterPublicDropUpdated(opts *bind.FilterOpts, nftContract []common.Address) (*ClubPublicDropUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Club.contract.FilterLogs(opts, "PublicDropUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return &ClubPublicDropUpdatedIterator{contract: _Club.contract, event: "PublicDropUpdated", logs: logs, sub: sub}, nil
}

// WatchPublicDropUpdated is a free log subscription operation binding the contract event 0x3e30d8e1f739ea4795c481b21c23f905e938b80339305f3508e43c558e5dead3.
//
// Solidity: event PublicDropUpdated(address indexed nftContract, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop)
func (_Club *ClubFilterer) WatchPublicDropUpdated(opts *bind.WatchOpts, sink chan<- *ClubPublicDropUpdated, nftContract []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Club.contract.WatchLogs(opts, "PublicDropUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubPublicDropUpdated)
				if err := _Club.contract.UnpackLog(event, "PublicDropUpdated", log); err != nil {
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
func (_Club *ClubFilterer) ParsePublicDropUpdated(log types.Log) (*ClubPublicDropUpdated, error) {
	event := new(ClubPublicDropUpdated)
	if err := _Club.contract.UnpackLog(event, "PublicDropUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubRoyaltyInfoUpdatedIterator is returned from FilterRoyaltyInfoUpdated and is used to iterate over the raw logs and unpacked data for RoyaltyInfoUpdated events raised by the Club contract.
type ClubRoyaltyInfoUpdatedIterator struct {
	Event *ClubRoyaltyInfoUpdated // Event containing the contract specifics and raw log

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
func (it *ClubRoyaltyInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubRoyaltyInfoUpdated)
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
		it.Event = new(ClubRoyaltyInfoUpdated)
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
func (it *ClubRoyaltyInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubRoyaltyInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubRoyaltyInfoUpdated represents a RoyaltyInfoUpdated event raised by the Club contract.
type ClubRoyaltyInfoUpdated struct {
	Receiver common.Address
	Bps      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyInfoUpdated is a free log retrieval operation binding the contract event 0xf21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d.
//
// Solidity: event RoyaltyInfoUpdated(address receiver, uint256 bps)
func (_Club *ClubFilterer) FilterRoyaltyInfoUpdated(opts *bind.FilterOpts) (*ClubRoyaltyInfoUpdatedIterator, error) {

	logs, sub, err := _Club.contract.FilterLogs(opts, "RoyaltyInfoUpdated")
	if err != nil {
		return nil, err
	}
	return &ClubRoyaltyInfoUpdatedIterator{contract: _Club.contract, event: "RoyaltyInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchRoyaltyInfoUpdated is a free log subscription operation binding the contract event 0xf21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d.
//
// Solidity: event RoyaltyInfoUpdated(address receiver, uint256 bps)
func (_Club *ClubFilterer) WatchRoyaltyInfoUpdated(opts *bind.WatchOpts, sink chan<- *ClubRoyaltyInfoUpdated) (event.Subscription, error) {

	logs, sub, err := _Club.contract.WatchLogs(opts, "RoyaltyInfoUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubRoyaltyInfoUpdated)
				if err := _Club.contract.UnpackLog(event, "RoyaltyInfoUpdated", log); err != nil {
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
func (_Club *ClubFilterer) ParseRoyaltyInfoUpdated(log types.Log) (*ClubRoyaltyInfoUpdated, error) {
	event := new(ClubRoyaltyInfoUpdated)
	if err := _Club.contract.UnpackLog(event, "RoyaltyInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubSeaDropMintIterator is returned from FilterSeaDropMint and is used to iterate over the raw logs and unpacked data for SeaDropMint events raised by the Club contract.
type ClubSeaDropMintIterator struct {
	Event *ClubSeaDropMint // Event containing the contract specifics and raw log

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
func (it *ClubSeaDropMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubSeaDropMint)
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
		it.Event = new(ClubSeaDropMint)
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
func (it *ClubSeaDropMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubSeaDropMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubSeaDropMint represents a SeaDropMint event raised by the Club contract.
type ClubSeaDropMint struct {
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
func (_Club *ClubFilterer) FilterSeaDropMint(opts *bind.FilterOpts, nftContract []common.Address, minter []common.Address, feeRecipient []common.Address) (*ClubSeaDropMintIterator, error) {

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

	logs, sub, err := _Club.contract.FilterLogs(opts, "SeaDropMint", nftContractRule, minterRule, feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &ClubSeaDropMintIterator{contract: _Club.contract, event: "SeaDropMint", logs: logs, sub: sub}, nil
}

// WatchSeaDropMint is a free log subscription operation binding the contract event 0xe90cf9cc0a552cf52ea6ff74ece0f1c8ae8cc9ad630d3181f55ac43ca076b7d6.
//
// Solidity: event SeaDropMint(address indexed nftContract, address indexed minter, address indexed feeRecipient, address payer, uint256 quantityMinted, uint256 unitMintPrice, uint256 feeBps, uint256 dropStageIndex)
func (_Club *ClubFilterer) WatchSeaDropMint(opts *bind.WatchOpts, sink chan<- *ClubSeaDropMint, nftContract []common.Address, minter []common.Address, feeRecipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Club.contract.WatchLogs(opts, "SeaDropMint", nftContractRule, minterRule, feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubSeaDropMint)
				if err := _Club.contract.UnpackLog(event, "SeaDropMint", log); err != nil {
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
func (_Club *ClubFilterer) ParseSeaDropMint(log types.Log) (*ClubSeaDropMint, error) {
	event := new(ClubSeaDropMint)
	if err := _Club.contract.UnpackLog(event, "SeaDropMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubSignedMintValidationParamsUpdatedIterator is returned from FilterSignedMintValidationParamsUpdated and is used to iterate over the raw logs and unpacked data for SignedMintValidationParamsUpdated events raised by the Club contract.
type ClubSignedMintValidationParamsUpdatedIterator struct {
	Event *ClubSignedMintValidationParamsUpdated // Event containing the contract specifics and raw log

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
func (it *ClubSignedMintValidationParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubSignedMintValidationParamsUpdated)
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
		it.Event = new(ClubSignedMintValidationParamsUpdated)
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
func (it *ClubSignedMintValidationParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubSignedMintValidationParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubSignedMintValidationParamsUpdated represents a SignedMintValidationParamsUpdated event raised by the Club contract.
type ClubSignedMintValidationParamsUpdated struct {
	NftContract                common.Address
	Signer                     common.Address
	SignedMintValidationParams SignedMintValidationParams
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterSignedMintValidationParamsUpdated is a free log retrieval operation binding the contract event 0xcaeb4009c05208df426d15ff50b608287b05d21dee1f790552ea451a540a7be0.
//
// Solidity: event SignedMintValidationParamsUpdated(address indexed nftContract, address indexed signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams)
func (_Club *ClubFilterer) FilterSignedMintValidationParamsUpdated(opts *bind.FilterOpts, nftContract []common.Address, signer []common.Address) (*ClubSignedMintValidationParamsUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Club.contract.FilterLogs(opts, "SignedMintValidationParamsUpdated", nftContractRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &ClubSignedMintValidationParamsUpdatedIterator{contract: _Club.contract, event: "SignedMintValidationParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchSignedMintValidationParamsUpdated is a free log subscription operation binding the contract event 0xcaeb4009c05208df426d15ff50b608287b05d21dee1f790552ea451a540a7be0.
//
// Solidity: event SignedMintValidationParamsUpdated(address indexed nftContract, address indexed signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams)
func (_Club *ClubFilterer) WatchSignedMintValidationParamsUpdated(opts *bind.WatchOpts, sink chan<- *ClubSignedMintValidationParamsUpdated, nftContract []common.Address, signer []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Club.contract.WatchLogs(opts, "SignedMintValidationParamsUpdated", nftContractRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubSignedMintValidationParamsUpdated)
				if err := _Club.contract.UnpackLog(event, "SignedMintValidationParamsUpdated", log); err != nil {
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
func (_Club *ClubFilterer) ParseSignedMintValidationParamsUpdated(log types.Log) (*ClubSignedMintValidationParamsUpdated, error) {
	event := new(ClubSignedMintValidationParamsUpdated)
	if err := _Club.contract.UnpackLog(event, "SignedMintValidationParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubSignerUpdatedIterator is returned from FilterSignerUpdated and is used to iterate over the raw logs and unpacked data for SignerUpdated events raised by the Club contract.
type ClubSignerUpdatedIterator struct {
	Event *ClubSignerUpdated // Event containing the contract specifics and raw log

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
func (it *ClubSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubSignerUpdated)
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
		it.Event = new(ClubSignerUpdated)
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
func (it *ClubSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubSignerUpdated represents a SignerUpdated event raised by the Club contract.
type ClubSignerUpdated struct {
	Sender common.Address
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerUpdated is a free log retrieval operation binding the contract event 0x2d025324f0a785e8c12d0a0d91a9caa49df4ef20ff87e0df7213a1d4f3157beb.
//
// Solidity: event SignerUpdated(address indexed sender, address indexed signer)
func (_Club *ClubFilterer) FilterSignerUpdated(opts *bind.FilterOpts, sender []common.Address, signer []common.Address) (*ClubSignerUpdatedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Club.contract.FilterLogs(opts, "SignerUpdated", senderRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &ClubSignerUpdatedIterator{contract: _Club.contract, event: "SignerUpdated", logs: logs, sub: sub}, nil
}

// WatchSignerUpdated is a free log subscription operation binding the contract event 0x2d025324f0a785e8c12d0a0d91a9caa49df4ef20ff87e0df7213a1d4f3157beb.
//
// Solidity: event SignerUpdated(address indexed sender, address indexed signer)
func (_Club *ClubFilterer) WatchSignerUpdated(opts *bind.WatchOpts, sink chan<- *ClubSignerUpdated, sender []common.Address, signer []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Club.contract.WatchLogs(opts, "SignerUpdated", senderRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubSignerUpdated)
				if err := _Club.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
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
func (_Club *ClubFilterer) ParseSignerUpdated(log types.Log) (*ClubSignerUpdated, error) {
	event := new(ClubSignerUpdated)
	if err := _Club.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubSpaceSizeUpdatedIterator is returned from FilterSpaceSizeUpdated and is used to iterate over the raw logs and unpacked data for SpaceSizeUpdated events raised by the Club contract.
type ClubSpaceSizeUpdatedIterator struct {
	Event *ClubSpaceSizeUpdated // Event containing the contract specifics and raw log

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
func (it *ClubSpaceSizeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubSpaceSizeUpdated)
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
		it.Event = new(ClubSpaceSizeUpdated)
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
func (it *ClubSpaceSizeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubSpaceSizeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubSpaceSizeUpdated represents a SpaceSizeUpdated event raised by the Club contract.
type ClubSpaceSizeUpdated struct {
	Size   *big.Int
	Length *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSpaceSizeUpdated is a free log retrieval operation binding the contract event 0x47208e794e45f5c352c883a84f2d7136a6144e7516979bf97a13266f807c4185.
//
// Solidity: event SpaceSizeUpdated(uint256 size, uint256 length)
func (_Club *ClubFilterer) FilterSpaceSizeUpdated(opts *bind.FilterOpts) (*ClubSpaceSizeUpdatedIterator, error) {

	logs, sub, err := _Club.contract.FilterLogs(opts, "SpaceSizeUpdated")
	if err != nil {
		return nil, err
	}
	return &ClubSpaceSizeUpdatedIterator{contract: _Club.contract, event: "SpaceSizeUpdated", logs: logs, sub: sub}, nil
}

// WatchSpaceSizeUpdated is a free log subscription operation binding the contract event 0x47208e794e45f5c352c883a84f2d7136a6144e7516979bf97a13266f807c4185.
//
// Solidity: event SpaceSizeUpdated(uint256 size, uint256 length)
func (_Club *ClubFilterer) WatchSpaceSizeUpdated(opts *bind.WatchOpts, sink chan<- *ClubSpaceSizeUpdated) (event.Subscription, error) {

	logs, sub, err := _Club.contract.WatchLogs(opts, "SpaceSizeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubSpaceSizeUpdated)
				if err := _Club.contract.UnpackLog(event, "SpaceSizeUpdated", log); err != nil {
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

// ParseSpaceSizeUpdated is a log parse operation binding the contract event 0x47208e794e45f5c352c883a84f2d7136a6144e7516979bf97a13266f807c4185.
//
// Solidity: event SpaceSizeUpdated(uint256 size, uint256 length)
func (_Club *ClubFilterer) ParseSpaceSizeUpdated(log types.Log) (*ClubSpaceSizeUpdated, error) {
	event := new(ClubSpaceSizeUpdated)
	if err := _Club.contract.UnpackLog(event, "SpaceSizeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubTokenGatedDropStageUpdatedIterator is returned from FilterTokenGatedDropStageUpdated and is used to iterate over the raw logs and unpacked data for TokenGatedDropStageUpdated events raised by the Club contract.
type ClubTokenGatedDropStageUpdatedIterator struct {
	Event *ClubTokenGatedDropStageUpdated // Event containing the contract specifics and raw log

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
func (it *ClubTokenGatedDropStageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubTokenGatedDropStageUpdated)
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
		it.Event = new(ClubTokenGatedDropStageUpdated)
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
func (it *ClubTokenGatedDropStageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubTokenGatedDropStageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubTokenGatedDropStageUpdated represents a TokenGatedDropStageUpdated event raised by the Club contract.
type ClubTokenGatedDropStageUpdated struct {
	NftContract     common.Address
	AllowedNftToken common.Address
	DropStage       TokenGatedDropStage
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenGatedDropStageUpdated is a free log retrieval operation binding the contract event 0xc695f93ae16034280e4fc93181b6afca9af23027ac1f1842a2287ba25cdc4476.
//
// Solidity: event TokenGatedDropStageUpdated(address indexed nftContract, address indexed allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage)
func (_Club *ClubFilterer) FilterTokenGatedDropStageUpdated(opts *bind.FilterOpts, nftContract []common.Address, allowedNftToken []common.Address) (*ClubTokenGatedDropStageUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var allowedNftTokenRule []interface{}
	for _, allowedNftTokenItem := range allowedNftToken {
		allowedNftTokenRule = append(allowedNftTokenRule, allowedNftTokenItem)
	}

	logs, sub, err := _Club.contract.FilterLogs(opts, "TokenGatedDropStageUpdated", nftContractRule, allowedNftTokenRule)
	if err != nil {
		return nil, err
	}
	return &ClubTokenGatedDropStageUpdatedIterator{contract: _Club.contract, event: "TokenGatedDropStageUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenGatedDropStageUpdated is a free log subscription operation binding the contract event 0xc695f93ae16034280e4fc93181b6afca9af23027ac1f1842a2287ba25cdc4476.
//
// Solidity: event TokenGatedDropStageUpdated(address indexed nftContract, address indexed allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage)
func (_Club *ClubFilterer) WatchTokenGatedDropStageUpdated(opts *bind.WatchOpts, sink chan<- *ClubTokenGatedDropStageUpdated, nftContract []common.Address, allowedNftToken []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var allowedNftTokenRule []interface{}
	for _, allowedNftTokenItem := range allowedNftToken {
		allowedNftTokenRule = append(allowedNftTokenRule, allowedNftTokenItem)
	}

	logs, sub, err := _Club.contract.WatchLogs(opts, "TokenGatedDropStageUpdated", nftContractRule, allowedNftTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubTokenGatedDropStageUpdated)
				if err := _Club.contract.UnpackLog(event, "TokenGatedDropStageUpdated", log); err != nil {
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
func (_Club *ClubFilterer) ParseTokenGatedDropStageUpdated(log types.Log) (*ClubTokenGatedDropStageUpdated, error) {
	event := new(ClubTokenGatedDropStageUpdated)
	if err := _Club.contract.UnpackLog(event, "TokenGatedDropStageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Club contract.
type ClubTransferIterator struct {
	Event *ClubTransfer // Event containing the contract specifics and raw log

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
func (it *ClubTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubTransfer)
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
		it.Event = new(ClubTransfer)
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
func (it *ClubTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubTransfer represents a Transfer event raised by the Club contract.
type ClubTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Club *ClubFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ClubTransferIterator, error) {

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

	logs, sub, err := _Club.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ClubTransferIterator{contract: _Club.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Club *ClubFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ClubTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Club.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubTransfer)
				if err := _Club.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Club *ClubFilterer) ParseTransfer(log types.Log) (*ClubTransfer, error) {
	event := new(ClubTransfer)
	if err := _Club.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClubUnplegedUpdatedIterator is returned from FilterUnplegedUpdated and is used to iterate over the raw logs and unpacked data for UnplegedUpdated events raised by the Club contract.
type ClubUnplegedUpdatedIterator struct {
	Event *ClubUnplegedUpdated // Event containing the contract specifics and raw log

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
func (it *ClubUnplegedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClubUnplegedUpdated)
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
		it.Event = new(ClubUnplegedUpdated)
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
func (it *ClubUnplegedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClubUnplegedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClubUnplegedUpdated represents a UnplegedUpdated event raised by the Club contract.
type ClubUnplegedUpdated struct {
	ClubId      *big.Int
	NftContract common.Address
	Sender      common.Address
	Params      ERC721ClubUpgradeableClubParams
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnplegedUpdated is a free log retrieval operation binding the contract event 0x8e8823d9b5c6f5ebda354c473fe0a4b30300c54dc7ca47de51363766a2a1eb17.
//
// Solidity: event UnplegedUpdated(uint256 indexed clubId, address indexed nftContract, address indexed sender, (uint256,uint256[],uint256,uint256,uint256,uint256) params)
func (_Club *ClubFilterer) FilterUnplegedUpdated(opts *bind.FilterOpts, clubId []*big.Int, nftContract []common.Address, sender []common.Address) (*ClubUnplegedUpdatedIterator, error) {

	var clubIdRule []interface{}
	for _, clubIdItem := range clubId {
		clubIdRule = append(clubIdRule, clubIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Club.contract.FilterLogs(opts, "UnplegedUpdated", clubIdRule, nftContractRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ClubUnplegedUpdatedIterator{contract: _Club.contract, event: "UnplegedUpdated", logs: logs, sub: sub}, nil
}

// WatchUnplegedUpdated is a free log subscription operation binding the contract event 0x8e8823d9b5c6f5ebda354c473fe0a4b30300c54dc7ca47de51363766a2a1eb17.
//
// Solidity: event UnplegedUpdated(uint256 indexed clubId, address indexed nftContract, address indexed sender, (uint256,uint256[],uint256,uint256,uint256,uint256) params)
func (_Club *ClubFilterer) WatchUnplegedUpdated(opts *bind.WatchOpts, sink chan<- *ClubUnplegedUpdated, clubId []*big.Int, nftContract []common.Address, sender []common.Address) (event.Subscription, error) {

	var clubIdRule []interface{}
	for _, clubIdItem := range clubId {
		clubIdRule = append(clubIdRule, clubIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Club.contract.WatchLogs(opts, "UnplegedUpdated", clubIdRule, nftContractRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClubUnplegedUpdated)
				if err := _Club.contract.UnpackLog(event, "UnplegedUpdated", log); err != nil {
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

// ParseUnplegedUpdated is a log parse operation binding the contract event 0x8e8823d9b5c6f5ebda354c473fe0a4b30300c54dc7ca47de51363766a2a1eb17.
//
// Solidity: event UnplegedUpdated(uint256 indexed clubId, address indexed nftContract, address indexed sender, (uint256,uint256[],uint256,uint256,uint256,uint256) params)
func (_Club *ClubFilterer) ParseUnplegedUpdated(log types.Log) (*ClubUnplegedUpdated, error) {
	event := new(ClubUnplegedUpdated)
	if err := _Club.contract.UnpackLog(event, "UnplegedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
