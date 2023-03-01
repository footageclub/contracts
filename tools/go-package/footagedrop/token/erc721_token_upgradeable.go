// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// AllowListData is an auto generated low-level Go binding around an user-defined struct.
type AllowListData struct {
	MerkleRoot    [32]byte
	PublicKeyURIs []string
	AllowListURI  string
}

// ERC721SeaDropStructsErrorsAndEventsUpgradeableMultiConfigureStruct is an auto generated low-level Go binding around an user-defined struct.
type ERC721SeaDropStructsErrorsAndEventsUpgradeableMultiConfigureStruct struct {
	MaxSupply                            *big.Int
	BaseURI                              string
	ContractURI                          string
	SeaDropImpl                          common.Address
	PublicDrop                           PublicDrop
	DropURI                              string
	AllowListData                        AllowListData
	CreatorPayoutAddress                 common.Address
	ProvenanceHash                       [32]byte
	AllowedFeeRecipients                 []common.Address
	DisallowedFeeRecipients              []common.Address
	AllowedPayers                        []common.Address
	DisallowedPayers                     []common.Address
	TokenGatedAllowedNftTokens           []common.Address
	TokenGatedDropStages                 []TokenGatedDropStage
	DisallowedTokenGatedAllowedNftTokens []common.Address
	Signers                              []common.Address
	SignedMintValidationParams           []SignedMintValidationParams
	DisallowedSigners                    []common.Address
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

// TokenMetaData contains all meta data concerning the Token contract.
var TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BurnIncorrectSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"CannotExceedMaxSupplyOfUint64\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"basisPoints\",\"type\":\"uint256\"}],\"name\":\"InvalidRoyaltyBasisPoints\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotNextOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedSeaDrop\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProvenanceHashCannotBeSetAfterMintStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoyaltyAddressCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignersMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowedSeaDrop\",\"type\":\"address[]\"}],\"name\":\"AllowedSeaDropUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"}],\"name\":\"ContractURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPotentialAdministrator\",\"type\":\"address\"}],\"name\":\"PotentialOwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newHash\",\"type\":\"bytes32\"}],\"name\":\"ProvenanceHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"name\":\"RoyaltyInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SeaDropTokenDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"}],\"name\":\"emitBatchMetadataUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBurnAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"getMintStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minterNumMinted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"allowedSeaDrop\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"mintSeaDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contractURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"dropURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"publicKeyURIs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"internalType\":\"structAllowListData\",\"name\":\"allowListData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"creatorPayoutAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"provenanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"allowedFeeRecipients\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedFeeRecipients\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowedPayers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedPayers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenGatedAllowedNftTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structTokenGatedDropStage[]\",\"name\":\"tokenGatedDropStages\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedTokenGatedAllowedNftTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"internalType\":\"structSignedMintValidationParams[]\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedSigners\",\"type\":\"address[]\"}],\"internalType\":\"structERC721SeaDropStructsErrorsAndEventsUpgradeable.MultiConfigureStruct\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"multiConfigure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provenanceHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBurnAddress\",\"type\":\"address\"}],\"name\":\"setBurnAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newProvenanceHash\",\"type\":\"bytes32\"}],\"name\":\"setProvenanceHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"royaltyAddress\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"royaltyBps\",\"type\":\"uint96\"}],\"internalType\":\"structISeaDropTokenContractMetadataUpgradeable.RoyaltyInfo\",\"name\":\"newInfo\",\"type\":\"tuple\"}],\"name\":\"setRoyaltyInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPotentialOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"publicKeyURIs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"internalType\":\"structAllowListData\",\"name\":\"allowListData\",\"type\":\"tuple\"}],\"name\":\"updateAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateAllowedFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"allowedSeaDrop\",\"type\":\"address[]\"}],\"name\":\"updateAllowedSeaDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payoutAddress\",\"type\":\"address\"}],\"name\":\"updateCreatorPayoutAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"dropURI\",\"type\":\"string\"}],\"name\":\"updateDropURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updatePayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"}],\"name\":\"updatePublicDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"internalType\":\"structSignedMintValidationParams\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple\"}],\"name\":\"updateSignedMintValidationParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structTokenGatedDropStage\",\"name\":\"dropStage\",\"type\":\"tuple\"}],\"name\":\"updateTokenGatedDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506149a6806100206000396000f3fe608060405234801561001057600080fd5b50600436106102a05760003560e01c80636c0360eb11610167578063a22cb465116100ce578063ca27593111610087578063ca27593114610615578063cb743ba814610628578063d5abeb011461063b578063e8a3d48514610643578063e985e9c51461064b578063f2fde38b1461065e57600080fd5b8063a22cb46514610584578063a483011414610597578063ad2f852a146105aa578063b88d4fde146105c8578063c6ab67a3146105db578063c87b56dd1461060257600080fd5b80637bc2be76116101205780637bc2be761461050d578063840e15d4146105205780638da5cb5b1461054e578063911f456b14610556578063938e3d7b1461056957806395d89b411461057c57600080fd5b80636c0360eb146104bc5780636f8b44b0146104c457806370a08231146104d7578063715018a6146104ea57806379ba5097146104f25780637a05bc82146104fa57600080fd5b806342260b5d1161020b578063511aa644116101c4578063511aa6441461044a57806355f804b31461045d57806360c308b6146104705780636352211e1461048357806364869dad1461049657806366251b69146104a957600080fd5b806342260b5d146103c657806342842e0e146103eb57806342966c68146103fe57806344dae42c1461041157806348a4c101146104245780634b0e72161461043757600080fd5b80631b73593c1161025d5780631b73593c1461034b57806323452b9c1461035e57806323b872dd146103665780632a55205a146103795780633680620d146103ab57806338b39d29146103be57600080fd5b806301ffc9a7146102a557806306fdde03146102cd578063081812fc146102e2578063095ea7b31461030d578063099b6bfa1461032257806318160ddd14610335575b600080fd5b6102b86102b3366004613508565b610671565b60405190151581526020015b60405180910390f35b6102d56106b7565b6040516102c49190613575565b6102f56102f0366004613588565b610752565b6040516001600160a01b0390911681526020016102c4565b61032061031b3660046135b6565b61079f565b005b610320610330366004613588565b6107b8565b61033d61084e565b6040519081526020016102c4565b6103206103593660046135e2565b61086e565b6103206108e1565b610320610374366004613623565b61093c565b61038c610387366004613664565b610967565b604080516001600160a01b0390931683526020830191909152016102c4565b6103206103b9366004613686565b6109ba565b6102f56109f7565b60008051602061493183398151915254600160a01b90046001600160601b031661033d565b6103206103f9366004613623565b610a2a565b61032061040c366004613588565b610a4f565b61032061041f3660046136dc565b610aa5565b610320610432366004613712565b610bc8565b61032061044536600461375d565b610c42565b610320610458366004613839565b610c8b565b61032061046b366004613950565b610cca565b61032061047e366004613991565b610d5b565b6102f5610491366004613588565b610d6d565b6103206104a43660046135b6565b610d78565b6103206104b7366004613a05565b610e62565b6102d5610ea1565b6103206104d2366004613588565b610eb0565b61033d6104e536600461375d565b610f24565b610320610f8c565b610320610fa0565b610320610508366004613a33565b611033565b61032061051b366004613a87565b611072565b61053361052e36600461375d565b6110b1565b604080519384526020840192909252908201526060016102c4565b6102f56110e3565b610320610564366004613adc565b61110b565b610320610577366004613950565b611d28565b6102d5611d8e565b610320610592366004613b17565b611da6565b6103206105a5366004613664565b611dba565b600080516020614931833981519152546001600160a01b03166102f5565b6103206105d6366004613b9c565b611df8565b7fb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031c5461033d565b6102d5610610366004613588565b611e25565b610320610623366004613c3b565b611ea9565b610320610636366004613712565b6120c7565b61033d61210e565b6102d5612126565b6102b8610659366004613a05565b61213c565b61032061066c36600461375d565b612179565b60006001600160e01b03198216630c487f4760e11b14806106a257506001600160e01b03198216639c15441560e01b145b806106b157506106b182612203565b92915050565b60606106c1612243565b60020180546106cf90613d33565b80601f01602080910402602001604051908101604052809291908181526020018280546106fb90613d33565b80156107485780601f1061071d57610100808354040283529160200191610748565b820191906000526020600020905b81548152906001019060200180831161072b57829003601f168201915b5050505050905090565b600061075d82612267565b61077a576040516333d1c03960e21b815260040160405180910390fd5b610782612243565b60009283526006016020525060409020546001600160a01b031690565b816107a9816122b0565b6107b38383612369565b505050565b6107c0612417565b60006107ca61245d565b11156107e95760405163e03264af60e01b815260040160405180910390fd5b7fb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031c80549082905560408051828152602081018490527f7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c91015b60405180910390a15050565b6000600161085a612243565b60010154610866612243565b540303919050565b610876612470565b61087f826124ba565b6040516301308e6560e01b81526001600160a01b038316906301308e65906108ab908490600401613e08565b600060405180830381600087803b1580156108c557600080fd5b505af11580156108d9573d6000803e3d6000fd5b505050505050565b6108e9612470565b60008051602061489183398151915280546001600160a01b0319169055604080516000815290517f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da9181900360200190a1565b826001600160a01b038116331461095657610956336122b0565b610961848484612505565b50505050565b600080516020614931833981519152805460009182916127109061099b90600160a01b90046001600160601b031686613e2c565b6109a59190613e43565b90546001600160a01b03169590945092505050565b6109c2612470565b6109cb826124ba565b60405163ebb4a55f60e01b81526001600160a01b0383169063ebb4a55f906108ab908490600401613fa6565b60007f12b889d8cf32baab74978bbc0a53ed4e7550971861e47aa7865383dea5acf0935b546001600160a01b0316919050565b826001600160a01b0381163314610a4457610a44336122b0565b6109618484846126e6565b7f12b889d8cf32baab74978bbc0a53ed4e7550971861e47aa7865383dea5acf093546001600160a01b03163314610a995760405163ea885d8360e01b815260040160405180910390fd5b610aa281612701565b50565b610aad612417565b6000610abc602083018361375d565b6001600160a01b031603610ae357604051631cc0baef60e01b815260040160405180910390fd5b612710610af66040830160208401613fce565b6001600160601b03161115610b3f57610b156040820160208301613fce565b604051633cadbafb60e01b81526001600160601b0390911660048201526024015b60405180910390fd5b80600080516020614931833981519152610b598282613feb565b507ff21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d9050610b8a602083018361375d565b610b9a6040840160208501613fce565b604080516001600160a01b0390931683526001600160601b039091166020830152015b60405180910390a150565b610bd0612470565b610bd9836124ba565b604051638e7d1e4360e01b81526001600160a01b0383811660048301528215156024830152841690638e7d1e43906044015b600060405180830381600087803b158015610c2557600080fd5b505af1158015610c39573d6000803e3d6000fd5b50505050505050565b610c4a612470565b7f12b889d8cf32baab74978bbc0a53ed4e7550971861e47aa7865383dea5acf09380546001600160a01b0319166001600160a01b0392909216919091179055565b610c93612470565b610c9c836124ba565b6040516309a7002f60e31b81526001600160a01b03841690634d38017890610c0b9085908590600401614098565b610cd2612417565b7fb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031a610cfe8284836140fc565b50610d0761084e565b15610d57577f6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c600180610d3861270c565b610d4291906141bb565b60408051928352602083019190915201610842565b5050565b610d63612470565b610d578282612716565b60006106b18261288a565b6000805160206148b183398151915254600114610dc45760405162461bcd60e51b815260206004820152600a6024820152695245454e5452414e435960b01b6044820152606401610b36565b60026000805160206148b183398151915255610ddf336124ba565b610de761210e565b81610df061245d565b610dfa91906141ce565b1115610e3d5780610e0961245d565b610e1391906141ce565b610e1b61210e565b60405163384b48c560e21b815260048101929092526024820152604401610b36565b610e47828261291e565b60016000805160206148b1833981519152555050565b555050565b610e6a612470565b610e73826124ba565b60405163024e71b760e31b81526001600160a01b0382811660048301528316906312738db8906024016108ab565b6060610eab612938565b905090565b610eb8612417565b6001600160401b03811115610ee35760405163b43e913760e01b815260048101829052602401610b36565b80600080516020614911833981519152556040518181527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c90602001610bbd565b60006001600160a01b038216610f4d576040516323d3ad8160e21b815260040160405180910390fd5b6001600160401b03610f5d612243565b6005016000846001600160a01b03166001600160a01b0316815260200190815260200160002054169050919050565b610f94612470565b610f9e6000612957565b565b600080516020614891833981519152546001600160a01b0316338114610fd957604051636b7584e760e11b815260040160405180910390fd5b60008051602061489183398151915280546001600160a01b0319169055604080516000815290517f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da9181900360200190a1610aa281612957565b61103b612470565b611044836124ba565b60405163b957d0cb60e01b81526001600160a01b0384169063b957d0cb90610c0b90859085906004016141e1565b61107a612470565b611083836124ba565b604051637ecd591560e11b81526001600160a01b0384169063fd9ab22a90610c0b90859085906004016142c3565b60008060006110bf846129c8565b92506110c961245d565b915060008051602061491183398151915254929491935050565b60007ff73863e3917595a7e67829ed60f0c51bf14c7a0d0de47c7b02a00abd48488e4a610a1b565b611113612470565b80351561116c576040516306f8b44b60e41b8152813560048201523090636f8b44b090602401600060405180830381600087803b15801561115357600080fd5b505af1158015611167573d6000803e3d6000fd5b505050505b61117960208201826142e1565b1590506111e357306355f804b361119360208401846142e1565b6040518363ffffffff1660e01b81526004016111b09291906141e1565b600060405180830381600087803b1580156111ca57600080fd5b505af11580156111de573d6000803e3d6000fd5b505050505b6111f060408201826142e1565b15905061125a573063938e3d7b61120a60408401846142e1565b6040518363ffffffff1660e01b81526004016112279291906141e1565b600060405180830381600087803b15801561124157600080fd5b505af1158015611255573d6000803e3d6000fd5b505050505b61127a61126d60e0830160c08401614327565b65ffffffffffff16151590565b61128d61126d60c0840160a08501614327565b176001036112ff5730631b73593c6112ab608084016060850161375d565b836080016040518363ffffffff1660e01b81526004016112cc929190614342565b600060405180830381600087803b1580156112e657600080fd5b505af11580156112fa573d6000803e3d6000fd5b505050505b61130d6101408201826142e1565b1590506113895730637a05bc8261132a608084016060850161375d565b6113386101408501856142e1565b6040518463ffffffff1660e01b81526004016113569392919061435f565b600060405180830381600087803b15801561137057600080fd5b505af1158015611384573d6000803e3d6000fd5b505050505b600061139961016083018361438d565b35146114135730633680620d6113b5608084016060850161375d565b6113c361016085018561438d565b6040518363ffffffff1660e01b81526004016113e09291906143ad565b600060405180830381600087803b1580156113fa57600080fd5b505af115801561140e573d6000803e3d6000fd5b505050505b60006114276101a08301610180840161375d565b6001600160a01b0316146114be57306366251b6961144b608084016060850161375d565b61145d6101a08501610180860161375d565b6040516001600160e01b031960e085901b1681526001600160a01b03928316600482015291166024820152604401600060405180830381600087803b1580156114a557600080fd5b505af11580156114b9573d6000803e3d6000fd5b505050505b6101a08101351561151f576040516304cdb5fd60e11b81526101a08201356004820152309063099b6bfa90602401600060405180830381600087803b15801561150657600080fd5b505af115801561151a573d6000803e3d6000fd5b505050505b600061152f6101c08301836143d1565b905011156115f95760005b6115486101c08301836143d1565b90508110156115f757306348a4c101611567608085016060860161375d565b6115756101c08601866143d1565b858181106115855761158561441a565b905060200201602081019061159a919061375d565b60016040518463ffffffff1660e01b81526004016115ba93929190614430565b600060405180830381600087803b1580156115d457600080fd5b505af11580156115e8573d6000803e3d6000fd5b5050505080600101905061153a565b505b60006116096101e08301836143d1565b905011156116d35760005b6116226101e08301836143d1565b90508110156116d157306348a4c101611641608085016060860161375d565b61164f6101e08601866143d1565b8581811061165f5761165f61441a565b9050602002016020810190611674919061375d565b60006040518463ffffffff1660e01b815260040161169493929190614430565b600060405180830381600087803b1580156116ae57600080fd5b505af11580156116c2573d6000803e3d6000fd5b50505050806001019050611614565b505b60006116e36102008301836143d1565b905011156117ad5760005b6116fc6102008301836143d1565b90508110156117ab573063cb743ba861171b608085016060860161375d565b6117296102008601866143d1565b858181106117395761173961441a565b905060200201602081019061174e919061375d565b60016040518463ffffffff1660e01b815260040161176e93929190614430565b600060405180830381600087803b15801561178857600080fd5b505af115801561179c573d6000803e3d6000fd5b505050508060010190506116ee565b505b60006117bd6102208301836143d1565b905011156118875760005b6117d66102208301836143d1565b9050811015611885573063cb743ba86117f5608085016060860161375d565b6118036102208601866143d1565b858181106118135761181361441a565b9050602002016020810190611828919061375d565b60006040518463ffffffff1660e01b815260040161184893929190614430565b600060405180830381600087803b15801561186257600080fd5b505af1158015611876573d6000803e3d6000fd5b505050508060010190506117c8565b505b6000611897610260830183614454565b905011156119c2576118ad6102408201826143d1565b90506118bd610260830183614454565b9050146118dd5760405163b81aa63960e01b815260040160405180910390fd5b60005b6118ee610260830183614454565b90508110156119c05730637bc2be7661190d608085016060860161375d565b61191b6102408601866143d1565b8581811061192b5761192b61441a565b9050602002016020810190611940919061375d565b61194e610260870187614454565b8681811061195e5761195e61441a565b905061010002016040518463ffffffff1660e01b81526004016119839392919061449d565b600060405180830381600087803b15801561199d57600080fd5b505af11580156119b1573d6000803e3d6000fd5b505050508060010190506118e0565b505b60006119d26102808301836143d1565b90501115611add5760005b6119eb6102808301836143d1565b9050811015611adb576040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915230637bc2be76611a4b608086016060870161375d565b611a596102808701876143d1565b86818110611a6957611a6961441a565b9050602002016020810190611a7e919061375d565b846040518463ffffffff1660e01b8152600401611a9d939291906144c3565b600060405180830381600087803b158015611ab757600080fd5b505af1158015611acb573d6000803e3d6000fd5b50505050816001019150506119dd565b505b6000611aed6102c0830183614575565b90501115611c1757611b036102a08201826143d1565b9050611b136102c0830183614575565b905014611b33576040516374ef6df760e01b815260040160405180910390fd5b60005b611b446102c0830183614575565b9050811015611c15573063511aa644611b63608085016060860161375d565b611b716102a08601866143d1565b85818110611b8157611b8161441a565b9050602002016020810190611b96919061375d565b611ba46102c0870187614575565b86818110611bb457611bb461441a565b905060e002016040518463ffffffff1660e01b8152600401611bd8939291906145bd565b600060405180830381600087803b158015611bf257600080fd5b505af1158015611c06573d6000803e3d6000fd5b50505050806001019050611b36565b505b6000611c276102e08301836143d1565b90501115610aa25760005b611c406102e08301836143d1565b9050811015610d57576040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091523063511aa644611c98608086016060870161375d565b611ca66102e08701876143d1565b86818110611cb657611cb661441a565b9050602002016020810190611ccb919061375d565b846040518463ffffffff1660e01b8152600401611cea93929190614676565b600060405180830381600087803b158015611d0457600080fd5b505af1158015611d18573d6000803e3d6000fd5b5050505081600101915050611c32565b611d30612417565b7fb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031b611d5c8284836140fc565b507f905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac3737882826040516108429291906141e1565b6060611d98612243565b60030180546106cf90613d33565b81611db0816122b0565b6107b38383612a0d565b611dc2612417565b60408051838152602081018390527f6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c9101610842565b836001600160a01b0381163314611e1257611e12336122b0565b611e1e85858585612a8a565b5050505050565b6060611e3082612267565b611e4d57604051630a14c4b560e41b815260040160405180910390fd5b6000611e57612938565b90508051600003611e775760405180602001604052806000815250611ea2565b80611e8184612ace565b604051602001611e9292919061469c565b6040516020818303038152906040525b9392505050565b600054610100900460ff1615808015611ec95750600054600160ff909116105b80611ee35750303b158015611ee3575060005460ff166001145b611f465760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610b36565b6000805460ff191660011790558015611f69576000805461ff0019166101001790555b60008051602061495183398151915254610100900460ff16611f9e576000805160206149518339815191525460ff1615611fa2565b303b155b6120145760405162461bcd60e51b815260206004820152603760248201527f455243373231415f5f496e697469616c697a61626c653a20636f6e747261637460448201527f20697320616c726561647920696e697469616c697a65640000000000000000006064820152608401610b36565b60008051602061495183398151915254610100900460ff1615801561205057600080516020614951833981519152805461ffff19166101011790555b61205b858585612b12565b801561207a57600080516020614951833981519152805461ff00191690555b508015610961576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a150505050565b6120cf612470565b6120d8836124ba565b604051633f952e6560e11b81526001600160a01b0383811660048301528215156024830152841690637f2a5cca90604401610c0b565b60006000805160206149118339815191525b54919050565b60606000805160206149118339815191526106c1565b6000612146612243565b6001600160a01b039384166000908152600791909101602090815260408083209490951682529290925250205460ff1690565b612181612470565b6001600160a01b0381166121a857604051633a247dd760e11b815260040160405180910390fd5b60008051602061489183398151915280546001600160a01b0319166001600160a01b0383169081179091556040519081527f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da90602001610bbd565b60006001600160e01b0319821663152a902d60e11b14806122345750632483248360e11b6001600160e01b03198316145b806106b157506106b182612b78565b7f2569078dfb4b0305704d3008e7403993ae9601b85f7ae5e742de3de8f8011c4090565b600081600111158015612281575061227d612243565b5482105b80156106b15750600160e01b612295612243565b60008481526004919091016020526040902054161592915050565b6daaeb6d7670e522a718067333cd4e3b15610aa257604051633185c44d60e21b81523060048201526001600160a01b03821660248201526daaeb6d7670e522a718067333cd4e9063c617113490604401602060405180830381865afa15801561231d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061234191906146cb565b610aa257604051633b79c77360e21b81526001600160a01b0382166004820152602401610b36565b600061237482610d6d565b9050336001600160a01b038216146123ad57612390813361213c565b6123ad576040516367d9dca160e11b815260040160405180910390fd5b826123b6612243565b6000848152600691909101602052604080822080546001600160a01b0319166001600160a01b0394851617905551849286811692908516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259190a4505050565b30331461243c6124256110e3565b6001600160a01b0316336001600160a01b03161490565b17600003610f9e57604051635fc483c560e01b815260040160405180910390fd5b60006001612469612243565b5403919050565b337ff73863e3917595a7e67829ed60f0c51bf14c7a0d0de47c7b02a00abd48488e4a546001600160a01b031614610f9e57604051635fc483c560e01b815260040160405180910390fd5b6001600160a01b03811660009081526000805160206148d1833981519152602052604090205460ff161515600114610aa2576040516315e26ff360e01b815260040160405180910390fd5b60006125108261288a565b9050836001600160a01b0316816001600160a01b0316146125435760405162a1148160e81b815260040160405180910390fd5b60008061254f84612bc6565b91509150612574818761255f3390565b6001600160a01b039081169116811491141790565b61259f57612582863361213c565b61259f57604051632ce44b5f60e11b815260040160405180910390fd5b6001600160a01b0385166125c657604051633a954ecd60e21b815260040160405180910390fd5b80156125d157600082555b6125d9612243565b6001600160a01b0387166000908152600591909101602052604090208054600019019055612605612243565b6001600160a01b03861660008181526005929092016020526040909120805460010190554260a01b17600160e11b1761263c612243565b60008681526004919091016020526040812091909155600160e11b841690036126b2576001840161266b612243565b6000828152600491909101602052604081205490036126b05761268c612243565b5481146126b0578361269c612243565b600083815260049190910160205260409020555b505b83856001600160a01b0316876001600160a01b03166000805160206148f183398151915260405160405180910390a46108d9565b6107b383838360405180602001604052806000815250611df8565b610aa2816000612bee565b6000612120612243565b7ff268be8736a07172c20cb8afb46ffa17fa1131bf48395e58d9c0ce565c5047f4548160005b828110156127b65760006000805160206148d183398151915260006000805160206148d1833981519152600101848154811061277a5761277a61441a565b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff191691151591909117905560010161273c565b5060005b8181101561282c5760016000805160206148d183398151915260008787858181106127e7576127e761441a565b90506020020160208101906127fc919061375d565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790556001016127ba565b506128587ff268be8736a07172c20cb8afb46ffa17fa1131bf48395e58d9c0ce565c5047f48585613425565b507fbbd3b69c138de4d317d0bc4290282c4e1cbd1e58b579a5b4f114b598c237454d84846040516120b99291906146e8565b600081806001116129055761289d612243565b548110156129055760006128af612243565b600083815260049190910160205260408120549150600160e01b82169003612903575b80600003611ea2576128e2612243565b600019909201600081815260049390930160205260409092205490506128d2565b505b604051636f96cda160e11b815260040160405180910390fd5b610d57828260405180602001604052806000815250612d5d565b606060008051602061491183398151915260010180546106cf90613d33565b7ff73863e3917595a7e67829ed60f0c51bf14c7a0d0de47c7b02a00abd48488e4a80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60006001600160401b0360406129dc612243565b6005016000856001600160a01b03166001600160a01b0316815260200190815260200160002054901c169050919050565b80612a16612243565b336000818152600792909201602090815260408084206001600160a01b03881680865290835293819020805460ff19169515159590951790945592518415158152919290917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b612a9584848461093c565b6001600160a01b0383163b1561096157612ab184848484612dd3565b610961576040516368d2bf6b60e11b815260040160405180910390fd5b606060a06040510180604052602081039150506000815280825b600183039250600a81066030018353600a900480612ae85750819003601f19909101908152919050565b600054610100900460ff16612b395760405162461bcd60e51b8152600401610b3690614736565b612b438383612ebf565b612b4b612f7a565b612b53612fa1565b612b5d8383612fd0565b612b65612ff7565b612b6d613032565b6107b3838383613078565b60006301ffc9a760e01b6001600160e01b031983161480612ba957506380ac58cd60e01b6001600160e01b03198316145b806106b15750506001600160e01b031916635b5e139f60e01b1490565b6000806000612bd3612243565b60009485526006016020525050604090912080549092909150565b6000612bf98361288a565b905080600080612c0886612bc6565b915091508415612c4857612c1d81843361255f565b612c4857612c2b833361213c565b612c4857604051632ce44b5f60e11b815260040160405180910390fd5b8015612c5357600082555b6fffffffffffffffffffffffffffffffff612c6c612243565b6001600160a01b038516600081815260059290920160205260409091208054929092019091554260a01b17600360e01b17612ca5612243565b60008881526004919091016020526040812091909155600160e11b85169003612d1b5760018601612cd4612243565b600082815260049190910160205260408120549003612d1957612cf5612243565b548114612d195784612d05612243565b600083815260049190910160205260409020555b505b60405186906000906001600160a01b038616906000805160206148f1833981519152908390a4612d49612243565b600190810180549091019055505050505050565b612d678383613171565b6001600160a01b0383163b156107b3576000612d81612243565b5490508281035b612d9b6000868380600101945086612dd3565b612db8576040516368d2bf6b60e11b815260040160405180910390fd5b818110612d885781612dc8612243565b5414611e1e57600080fd5b604051630a85bd0160e11b81526000906001600160a01b0385169063150b7a0290612e08903390899088908890600401614781565b6020604051808303816000875af1925050508015612e43575060408051601f3d908101601f19168201909252612e40918101906147b4565b60015b612ea1573d808015612e71576040519150601f19603f3d011682016040523d82523d6000602084013e612e76565b606091505b508051600003612e99576040516368d2bf6b60e11b815260040160405180910390fd5b805181602001fd5b6001600160e01b031916630a85bd0160e11b1490505b949350505050565b60008051602061495183398151915254610100900460ff16612f405760405162461bcd60e51b815260206004820152603460248201527f455243373231415f5f496e697469616c697a61626c653a20636f6e7472616374604482015273206973206e6f7420696e697469616c697a696e6760601b6064820152608401610b36565b81612f49612243565b60020190612f5790826147d1565b5080612f61612243565b60030190612f6f90826147d1565b506001610e5d612243565b600054610100900460ff16610f9e5760405162461bcd60e51b8152600401610b3690614736565b600054610100900460ff16612fc85760405162461bcd60e51b8152600401610b3690614736565b610f9e613288565b600054610100900460ff16610d575760405162461bcd60e51b8152600401610b3690614736565b600054610100900460ff1661301e5760405162461bcd60e51b8152600401610b3690614736565b60016000805160206148b183398151915255565b600054610100900460ff166130595760405162461bcd60e51b8152600401610b3690614736565b610f9e733cc6cdda760b79bafa08df41ecfa224f810dceb660016132b0565b600054610100900460ff1661309f5760405162461bcd60e51b8152600401610b3690614736565b805160005b8181101561310e5760016000805160206148d183398151915260000160008584815181106130d4576130d461441a565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff19169115159190911790556001016130a4565b508151613141907ff268be8736a07172c20cb8afb46ffa17fa1131bf48395e58d9c0ce565c5047f4906020850190613488565b506040517fd7aca75208b9be5ffc04c6a01922020ffd62b55e68e502e317f5344960279af890600090a150505050565b600061317b612243565b54905060008290036131a05760405163b562e8dd60e01b815260040160405180910390fd5b6801000000000000000182026131b4612243565b6001600160a01b038516600081815260059290920160205260409091208054929092019091554260a01b6001841460e11b17176131ef612243565b600083815260049190910160205260408120919091556001600160a01b0384169083830190839083906000805160206148f18339815191528180a4600183015b81811461325557808360006000805160206148f1833981519152600080a460010161322f565b508160000361327657604051622e076360e81b815260040160405180910390fd5b8061327f612243565b55506107b39050565b303b156132a75760405162dc149f60e41b815260040160405180910390fd5b610f9e33612957565b600054610100900460ff166132d75760405162461bcd60e51b8152600401610b3690614736565b6daaeb6d7670e522a718067333cd4e3b15610d575760405163c3c5a54760e01b81523060048201526daaeb6d7670e522a718067333cd4e9063c3c5a547906024016020604051808303816000875af1158015613337573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061335b91906146cb565b610d575780156133a557604051633e9f1edf60e11b81523060048201526001600160a01b03831660248201526daaeb6d7670e522a718067333cd4e90637d3e3dbe906044016108ab565b6001600160a01b038216156133f45760405163a0af290360e01b81523060048201526001600160a01b03831660248201526daaeb6d7670e522a718067333cd4e9063a0af2903906044016108ab565b604051632210724360e11b81523060048201526daaeb6d7670e522a718067333cd4e90634420e486906024016108ab565b828054828255906000526020600020908101928215613478579160200282015b828111156134785781546001600160a01b0319166001600160a01b03843516178255602090920191600190910190613445565b506134849291506134dd565b5090565b828054828255906000526020600020908101928215613478579160200282015b8281111561347857825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906134a8565b5b8082111561348457600081556001016134de565b6001600160e01b031981168114610aa257600080fd5b60006020828403121561351a57600080fd5b8135611ea2816134f2565b60005b83811015613540578181015183820152602001613528565b50506000910152565b60008151808452613561816020860160208601613525565b601f01601f19169290920160200192915050565b602081526000611ea26020830184613549565b60006020828403121561359a57600080fd5b5035919050565b6001600160a01b0381168114610aa257600080fd5b600080604083850312156135c957600080fd5b82356135d4816135a1565b946020939093013593505050565b60008082840360e08112156135f657600080fd5b8335613601816135a1565b925060c0601f198201121561361557600080fd5b506020830190509250929050565b60008060006060848603121561363857600080fd5b8335613643816135a1565b92506020840135613653816135a1565b929592945050506040919091013590565b6000806040838503121561367757600080fd5b50508035926020909101359150565b6000806040838503121561369957600080fd5b82356136a4816135a1565b915060208301356001600160401b038111156136bf57600080fd5b8301606081860312156136d157600080fd5b809150509250929050565b6000604082840312156136ee57600080fd5b50919050565b8015158114610aa257600080fd5b803561370d816136f4565b919050565b60008060006060848603121561372757600080fd5b8335613732816135a1565b92506020840135613742816135a1565b91506040840135613752816136f4565b809150509250925092565b60006020828403121561376f57600080fd5b8135611ea2816135a1565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b03811182821017156137b2576137b261377a565b60405290565b604051601f8201601f191681016001600160401b03811182821017156137e0576137e061377a565b604052919050565b80356001600160501b038116811461370d57600080fd5b803562ffffff8116811461370d57600080fd5b803564ffffffffff8116811461370d57600080fd5b803561ffff8116811461370d57600080fd5b600080600083850361012081121561385057600080fd5b843561385b816135a1565b9350602085013561386b816135a1565b925060e0603f198201121561387f57600080fd5b50613888613790565b613894604086016137e8565b81526138a2606086016137ff565b60208201526138b360808601613812565b60408201526138c460a08601613812565b60608201526138d560c08601613812565b60808201526138e660e08601613827565b60a08201526138f86101008601613827565b60c0820152809150509250925092565b60008083601f84011261391a57600080fd5b5081356001600160401b0381111561393157600080fd5b60208301915083602082850101111561394957600080fd5b9250929050565b6000806020838503121561396357600080fd5b82356001600160401b0381111561397957600080fd5b61398585828601613908565b90969095509350505050565b600080602083850312156139a457600080fd5b82356001600160401b03808211156139bb57600080fd5b818501915085601f8301126139cf57600080fd5b8135818111156139de57600080fd5b8660208260051b85010111156139f357600080fd5b60209290920196919550909350505050565b60008060408385031215613a1857600080fd5b8235613a23816135a1565b915060208301356136d1816135a1565b600080600060408486031215613a4857600080fd5b8335613a53816135a1565b925060208401356001600160401b03811115613a6e57600080fd5b613a7a86828701613908565b9497909650939450505050565b6000806000838503610140811215613a9e57600080fd5b8435613aa9816135a1565b93506020850135613ab9816135a1565b9250610100603f1982011215613ace57600080fd5b506040840190509250925092565b600060208284031215613aee57600080fd5b81356001600160401b03811115613b0457600080fd5b82016103008185031215611ea257600080fd5b60008060408385031215613b2a57600080fd5b8235613b35816135a1565b915060208301356136d1816136f4565b60006001600160401b03831115613b5e57613b5e61377a565b613b71601f8401601f19166020016137b8565b9050828152838383011115613b8557600080fd5b828260208301376000602084830101529392505050565b60008060008060808587031215613bb257600080fd5b8435613bbd816135a1565b93506020850135613bcd816135a1565b92506040850135915060608501356001600160401b03811115613bef57600080fd5b8501601f81018713613c0057600080fd5b613c0f87823560208401613b45565b91505092959194509250565b600082601f830112613c2c57600080fd5b611ea283833560208501613b45565b600080600060608486031215613c5057600080fd5b83356001600160401b0380821115613c6757600080fd5b613c7387838801613c1b565b9450602091508186013581811115613c8a57600080fd5b613c9688828901613c1b565b945050604086013581811115613cab57600080fd5b8601601f81018813613cbc57600080fd5b803582811115613cce57613cce61377a565b8060051b9250613cdf8484016137b8565b818152928201840192848101908a851115613cf957600080fd5b928501925b84841015613d235783359250613d13836135a1565b8282529285019290850190613cfe565b8096505050505050509250925092565b600181811c90821680613d4757607f821691505b6020821081036136ee57634e487b7160e01b600052602260045260246000fd5b803565ffffffffffff8116811461370d57600080fd5b6001600160501b03613d8e826137e8565b168252613d9d60208201613d67565b65ffffffffffff808216602085015280613db960408501613d67565b1660408501525050613dcd60608201613827565b61ffff808216606085015280613de560808501613827565b166080850152505060a0810135613dfb816136f4565b80151560a0840152505050565b60c081016106b18284613d7d565b634e487b7160e01b600052601160045260246000fd5b80820281158282048414176106b1576106b1613e16565b600082613e6057634e487b7160e01b600052601260045260246000fd5b500490565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000808335601e19843603018112613ea557600080fd5b83016020810192503590506001600160401b03811115613ec457600080fd5b80360382131561394957600080fd5b60006060830182358452602080840135601e19853603018112613ef557600080fd5b840181810190356001600160401b03811115613f1057600080fd5b8060051b803603831315613f2357600080fd5b6060848901529381905260809387018401938290880160005b83811015613f7657898703607f19018252613f578386613e8e565b613f62898284613e65565b985050509185019190850190600101613f3c565b505050505050613f896040840184613e8e565b8583036040870152613f9c838284613e65565b9695505050505050565b602081526000611ea26020830184613ed3565b6001600160601b0381168114610aa257600080fd5b600060208284031215613fe057600080fd5b8135611ea281613fb9565b8135613ff6816135a1565b81546001600160a01b03199081166001600160a01b03929092169182178355602084013561402381613fb9565b60a01b1617905550565b6001600160501b03815116825262ffffff6020820151166020830152604081015164ffffffffff8082166040850152806060840151166060850152806080840151166080850152505060a081015161ffff80821660a08501528060c08401511660c085015250505050565b6001600160a01b03831681526101008101611ea2602083018461402d565b601f8211156107b357600081815260208120601f850160051c810160208610156140dd5750805b601f850160051c820191505b818110156108d9578281556001016140e9565b6001600160401b038311156141135761411361377a565b614127836141218354613d33565b836140b6565b6000601f84116001811461415b57600085156141435750838201355b600019600387901b1c1916600186901b178355611e1e565b600083815260209020601f19861690835b8281101561418c578685013582556020948501946001909201910161416c565b50868210156141a95760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b818103818111156106b1576106b1613e16565b808201808211156106b1576106b1613e16565b602081526000612eb7602083018486613e65565b803563ffffffff8116811461370d57600080fd5b6001600160501b0361421a826137e8565b16825261ffff61422c60208301613827565b16602083015261423e60408201613d67565b65ffffffffffff80821660408501528061425a60608501613d67565b1660608501525050608081013560ff8116811461427657600080fd5b60ff16608083015261428a60a082016141f5565b63ffffffff1660a08301526142a160c08201613827565b61ffff1660c08301526142b660e08201613702565b80151560e0840152505050565b6001600160a01b03831681526101208101611ea26020830184614209565b6000808335601e198436030181126142f857600080fd5b8301803591506001600160401b0382111561431257600080fd5b60200191503681900382131561394957600080fd5b60006020828403121561433957600080fd5b611ea282613d67565b6001600160a01b038316815260e08101611ea26020830184613d7d565b6001600160a01b03841681526040602082018190526000906143849083018486613e65565b95945050505050565b60008235605e198336030181126143a357600080fd5b9190910192915050565b6001600160a01b0383168152604060208201819052600090612eb790830184613ed3565b6000808335601e198436030181126143e857600080fd5b8301803591506001600160401b0382111561440257600080fd5b6020019150600581901b360382131561394957600080fd5b634e487b7160e01b600052603260045260246000fd5b6001600160a01b039384168152919092166020820152901515604082015260600190565b6000808335601e1984360301811261446b57600080fd5b8301803591506001600160401b0382111561448557600080fd5b6020019150600881901b360382131561394957600080fd5b6001600160a01b038481168252831660208201526101408101612eb76040830184614209565b60006101408201905060018060a01b0380861683528085166020840152506001600160501b03835116604083015261ffff602084015116606083015265ffffffffffff6040840151166080830152606083015161452a60a084018265ffffffffffff169052565b50608083015160ff811660c08401525060a083015163ffffffff811660e08401525060c083015161ffff81166101008401525060e08301518015156101208401525b50949350505050565b6000808335601e1984360301811261458c57600080fd5b8301803591506001600160401b038211156145a657600080fd5b602001915060e08102360382131561394957600080fd5b6001600160a01b0384811682528316602082015261012081016001600160501b036145e7846137e8565b16604083015262ffffff6145fd602085016137ff565b16606083015261460f60408401613812565b64ffffffffff80821660808501528061462a60608701613812565b1660a08501528061463d60808701613812565b1660c0850152505061465160a08401613827565b61ffff1660e083015261466660c08401613827565b61ffff811661010084015261456c565b6001600160a01b038481168252831660208201526101208101612eb7604083018461402d565b600083516146ae818460208801613525565b8351908301906146c2818360208801613525565b01949350505050565b6000602082840312156146dd57600080fd5b8151611ea2816136f4565b60208082528181018390526000908460408401835b8681101561472b578235614710816135a1565b6001600160a01b0316825291830191908301906001016146fd565b509695505050505050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090613f9c90830184613549565b6000602082840312156147c657600080fd5b8151611ea2816134f2565b81516001600160401b038111156147ea576147ea61377a565b6147fe816147f88454613d33565b846140b6565b602080601f831160018114614833576000841561481b5750858301515b600019600386901b1c1916600185901b1785556108d9565b600085815260208120601f198616915b8281101561486257888601518255948401946001909101908401614843565b50858210156148805787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fef73863e3917595a7e67829ed60f0c51bf14c7a0d0de47c7b02a00abd48488e4bd59f8a8c0d1463371c77782499276e5cbe466fd192ada543ceaea0a36604c1f2f268be8736a07172c20cb8afb46ffa17fa1131bf48395e58d9c0ce565c5047f3ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a40319b847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031dee151c8401928dc223602bb187aff91b9a56c7cae5476ef1b3287b085a16c85fa26469706673582212205f559756fe15d7a42f1c60021abb337aabde2ad42189e59580f91e7bbb6babb864736f6c63430008110033",
}

// TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMetaData.ABI instead.
var TokenABI = TokenMetaData.ABI

// TokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenMetaData.Bin instead.
var TokenBin = TokenMetaData.Bin

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := TokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Token *TokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Token *TokenCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Token *TokenSession) BaseURI() (string, error) {
	return _Token.Contract.BaseURI(&_Token.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Token *TokenCallerSession) BaseURI() (string, error) {
	return _Token.Contract.BaseURI(&_Token.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Token *TokenCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Token *TokenSession) ContractURI() (string, error) {
	return _Token.Contract.ContractURI(&_Token.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Token *TokenCallerSession) ContractURI() (string, error) {
	return _Token.Contract.ContractURI(&_Token.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Token *TokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Token *TokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Token.Contract.GetApproved(&_Token.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Token *TokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Token.Contract.GetApproved(&_Token.CallOpts, tokenId)
}

// GetBurnAddress is a free data retrieval call binding the contract method 0x38b39d29.
//
// Solidity: function getBurnAddress() view returns(address)
func (_Token *TokenCaller) GetBurnAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getBurnAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBurnAddress is a free data retrieval call binding the contract method 0x38b39d29.
//
// Solidity: function getBurnAddress() view returns(address)
func (_Token *TokenSession) GetBurnAddress() (common.Address, error) {
	return _Token.Contract.GetBurnAddress(&_Token.CallOpts)
}

// GetBurnAddress is a free data retrieval call binding the contract method 0x38b39d29.
//
// Solidity: function getBurnAddress() view returns(address)
func (_Token *TokenCallerSession) GetBurnAddress() (common.Address, error) {
	return _Token.Contract.GetBurnAddress(&_Token.CallOpts)
}

// GetMintStats is a free data retrieval call binding the contract method 0x840e15d4.
//
// Solidity: function getMintStats(address minter) view returns(uint256 minterNumMinted, uint256 currentTotalSupply, uint256 maxSupply)
func (_Token *TokenCaller) GetMintStats(opts *bind.CallOpts, minter common.Address) (struct {
	MinterNumMinted    *big.Int
	CurrentTotalSupply *big.Int
	MaxSupply          *big.Int
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getMintStats", minter)

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
func (_Token *TokenSession) GetMintStats(minter common.Address) (struct {
	MinterNumMinted    *big.Int
	CurrentTotalSupply *big.Int
	MaxSupply          *big.Int
}, error) {
	return _Token.Contract.GetMintStats(&_Token.CallOpts, minter)
}

// GetMintStats is a free data retrieval call binding the contract method 0x840e15d4.
//
// Solidity: function getMintStats(address minter) view returns(uint256 minterNumMinted, uint256 currentTotalSupply, uint256 maxSupply)
func (_Token *TokenCallerSession) GetMintStats(minter common.Address) (struct {
	MinterNumMinted    *big.Int
	CurrentTotalSupply *big.Int
	MaxSupply          *big.Int
}, error) {
	return _Token.Contract.GetMintStats(&_Token.CallOpts, minter)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Token *TokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Token *TokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Token.Contract.IsApprovedForAll(&_Token.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Token *TokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Token.Contract.IsApprovedForAll(&_Token.CallOpts, owner, operator)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Token *TokenCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Token *TokenSession) MaxSupply() (*big.Int, error) {
	return _Token.Contract.MaxSupply(&_Token.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Token *TokenCallerSession) MaxSupply() (*big.Int, error) {
	return _Token.Contract.MaxSupply(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCallerSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Token *TokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Token *TokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Token.Contract.OwnerOf(&_Token.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Token *TokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Token.Contract.OwnerOf(&_Token.CallOpts, tokenId)
}

// ProvenanceHash is a free data retrieval call binding the contract method 0xc6ab67a3.
//
// Solidity: function provenanceHash() view returns(bytes32)
func (_Token *TokenCaller) ProvenanceHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "provenanceHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvenanceHash is a free data retrieval call binding the contract method 0xc6ab67a3.
//
// Solidity: function provenanceHash() view returns(bytes32)
func (_Token *TokenSession) ProvenanceHash() ([32]byte, error) {
	return _Token.Contract.ProvenanceHash(&_Token.CallOpts)
}

// ProvenanceHash is a free data retrieval call binding the contract method 0xc6ab67a3.
//
// Solidity: function provenanceHash() view returns(bytes32)
func (_Token *TokenCallerSession) ProvenanceHash() ([32]byte, error) {
	return _Token.Contract.ProvenanceHash(&_Token.CallOpts)
}

// RoyaltyAddress is a free data retrieval call binding the contract method 0xad2f852a.
//
// Solidity: function royaltyAddress() view returns(address)
func (_Token *TokenCaller) RoyaltyAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "royaltyAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyAddress is a free data retrieval call binding the contract method 0xad2f852a.
//
// Solidity: function royaltyAddress() view returns(address)
func (_Token *TokenSession) RoyaltyAddress() (common.Address, error) {
	return _Token.Contract.RoyaltyAddress(&_Token.CallOpts)
}

// RoyaltyAddress is a free data retrieval call binding the contract method 0xad2f852a.
//
// Solidity: function royaltyAddress() view returns(address)
func (_Token *TokenCallerSession) RoyaltyAddress() (common.Address, error) {
	return _Token.Contract.RoyaltyAddress(&_Token.CallOpts)
}

// RoyaltyBasisPoints is a free data retrieval call binding the contract method 0x42260b5d.
//
// Solidity: function royaltyBasisPoints() view returns(uint256)
func (_Token *TokenCaller) RoyaltyBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "royaltyBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoyaltyBasisPoints is a free data retrieval call binding the contract method 0x42260b5d.
//
// Solidity: function royaltyBasisPoints() view returns(uint256)
func (_Token *TokenSession) RoyaltyBasisPoints() (*big.Int, error) {
	return _Token.Contract.RoyaltyBasisPoints(&_Token.CallOpts)
}

// RoyaltyBasisPoints is a free data retrieval call binding the contract method 0x42260b5d.
//
// Solidity: function royaltyBasisPoints() view returns(uint256)
func (_Token *TokenCallerSession) RoyaltyBasisPoints() (*big.Int, error) {
	return _Token.Contract.RoyaltyBasisPoints(&_Token.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Token *TokenCaller) RoyaltyInfo(opts *bind.CallOpts, arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "royaltyInfo", arg0, _salePrice)

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
func (_Token *TokenSession) RoyaltyInfo(arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Token.Contract.RoyaltyInfo(&_Token.CallOpts, arg0, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Token *TokenCallerSession) RoyaltyInfo(arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Token.Contract.RoyaltyInfo(&_Token.CallOpts, arg0, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Token *TokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Token *TokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Token.Contract.SupportsInterface(&_Token.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Token *TokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Token.Contract.SupportsInterface(&_Token.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Token *TokenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Token *TokenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Token.Contract.TokenURI(&_Token.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Token *TokenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Token.Contract.TokenURI(&_Token.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Token *TokenTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Token *TokenSession) AcceptOwnership() (*types.Transaction, error) {
	return _Token.Contract.AcceptOwnership(&_Token.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Token *TokenTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Token.Contract.AcceptOwnership(&_Token.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Token *TokenSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Token *TokenTransactorSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, operator, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Token *TokenTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Token *TokenSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Token *TokenTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, tokenId)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_Token *TokenTransactor) CancelOwnershipTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "cancelOwnershipTransfer")
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_Token *TokenSession) CancelOwnershipTransfer() (*types.Transaction, error) {
	return _Token.Contract.CancelOwnershipTransfer(&_Token.TransactOpts)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_Token *TokenTransactorSession) CancelOwnershipTransfer() (*types.Transaction, error) {
	return _Token.Contract.CancelOwnershipTransfer(&_Token.TransactOpts)
}

// EmitBatchMetadataUpdate is a paid mutator transaction binding the contract method 0xa4830114.
//
// Solidity: function emitBatchMetadataUpdate(uint256 fromTokenId, uint256 toTokenId) returns()
func (_Token *TokenTransactor) EmitBatchMetadataUpdate(opts *bind.TransactOpts, fromTokenId *big.Int, toTokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "emitBatchMetadataUpdate", fromTokenId, toTokenId)
}

// EmitBatchMetadataUpdate is a paid mutator transaction binding the contract method 0xa4830114.
//
// Solidity: function emitBatchMetadataUpdate(uint256 fromTokenId, uint256 toTokenId) returns()
func (_Token *TokenSession) EmitBatchMetadataUpdate(fromTokenId *big.Int, toTokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.EmitBatchMetadataUpdate(&_Token.TransactOpts, fromTokenId, toTokenId)
}

// EmitBatchMetadataUpdate is a paid mutator transaction binding the contract method 0xa4830114.
//
// Solidity: function emitBatchMetadataUpdate(uint256 fromTokenId, uint256 toTokenId) returns()
func (_Token *TokenTransactorSession) EmitBatchMetadataUpdate(fromTokenId *big.Int, toTokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.EmitBatchMetadataUpdate(&_Token.TransactOpts, fromTokenId, toTokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0xca275931.
//
// Solidity: function initialize(string name, string symbol, address[] allowedSeaDrop) returns()
func (_Token *TokenTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, allowedSeaDrop []common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "initialize", name, symbol, allowedSeaDrop)
}

// Initialize is a paid mutator transaction binding the contract method 0xca275931.
//
// Solidity: function initialize(string name, string symbol, address[] allowedSeaDrop) returns()
func (_Token *TokenSession) Initialize(name string, symbol string, allowedSeaDrop []common.Address) (*types.Transaction, error) {
	return _Token.Contract.Initialize(&_Token.TransactOpts, name, symbol, allowedSeaDrop)
}

// Initialize is a paid mutator transaction binding the contract method 0xca275931.
//
// Solidity: function initialize(string name, string symbol, address[] allowedSeaDrop) returns()
func (_Token *TokenTransactorSession) Initialize(name string, symbol string, allowedSeaDrop []common.Address) (*types.Transaction, error) {
	return _Token.Contract.Initialize(&_Token.TransactOpts, name, symbol, allowedSeaDrop)
}

// MintSeaDrop is a paid mutator transaction binding the contract method 0x64869dad.
//
// Solidity: function mintSeaDrop(address minter, uint256 quantity) returns()
func (_Token *TokenTransactor) MintSeaDrop(opts *bind.TransactOpts, minter common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mintSeaDrop", minter, quantity)
}

// MintSeaDrop is a paid mutator transaction binding the contract method 0x64869dad.
//
// Solidity: function mintSeaDrop(address minter, uint256 quantity) returns()
func (_Token *TokenSession) MintSeaDrop(minter common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Token.Contract.MintSeaDrop(&_Token.TransactOpts, minter, quantity)
}

// MintSeaDrop is a paid mutator transaction binding the contract method 0x64869dad.
//
// Solidity: function mintSeaDrop(address minter, uint256 quantity) returns()
func (_Token *TokenTransactorSession) MintSeaDrop(minter common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Token.Contract.MintSeaDrop(&_Token.TransactOpts, minter, quantity)
}

// MultiConfigure is a paid mutator transaction binding the contract method 0x911f456b.
//
// Solidity: function multiConfigure((uint256,string,string,address,(uint80,uint48,uint48,uint16,uint16,bool),string,(bytes32,string[],string),address,bytes32,address[],address[],address[],address[],address[],(uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool)[],address[],address[],(uint80,uint24,uint40,uint40,uint40,uint16,uint16)[],address[]) config) returns()
func (_Token *TokenTransactor) MultiConfigure(opts *bind.TransactOpts, config ERC721SeaDropStructsErrorsAndEventsUpgradeableMultiConfigureStruct) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "multiConfigure", config)
}

// MultiConfigure is a paid mutator transaction binding the contract method 0x911f456b.
//
// Solidity: function multiConfigure((uint256,string,string,address,(uint80,uint48,uint48,uint16,uint16,bool),string,(bytes32,string[],string),address,bytes32,address[],address[],address[],address[],address[],(uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool)[],address[],address[],(uint80,uint24,uint40,uint40,uint40,uint16,uint16)[],address[]) config) returns()
func (_Token *TokenSession) MultiConfigure(config ERC721SeaDropStructsErrorsAndEventsUpgradeableMultiConfigureStruct) (*types.Transaction, error) {
	return _Token.Contract.MultiConfigure(&_Token.TransactOpts, config)
}

// MultiConfigure is a paid mutator transaction binding the contract method 0x911f456b.
//
// Solidity: function multiConfigure((uint256,string,string,address,(uint80,uint48,uint48,uint16,uint16,bool),string,(bytes32,string[],string),address,bytes32,address[],address[],address[],address[],address[],(uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool)[],address[],address[],(uint80,uint24,uint40,uint40,uint40,uint16,uint16)[],address[]) config) returns()
func (_Token *TokenTransactorSession) MultiConfigure(config ERC721SeaDropStructsErrorsAndEventsUpgradeableMultiConfigureStruct) (*types.Transaction, error) {
	return _Token.Contract.MultiConfigure(&_Token.TransactOpts, config)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token.Contract.RenounceOwnership(&_Token.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token.Contract.RenounceOwnership(&_Token.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SafeTransferFrom(&_Token.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SafeTransferFrom(&_Token.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Token *TokenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Token *TokenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.Contract.SafeTransferFrom0(&_Token.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Token *TokenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.Contract.SafeTransferFrom0(&_Token.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token *TokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token *TokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token.Contract.SetApprovalForAll(&_Token.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token *TokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token.Contract.SetApprovalForAll(&_Token.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Token *TokenTransactor) SetBaseURI(opts *bind.TransactOpts, newBaseURI string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setBaseURI", newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Token *TokenSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Token.Contract.SetBaseURI(&_Token.TransactOpts, newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Token *TokenTransactorSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Token.Contract.SetBaseURI(&_Token.TransactOpts, newBaseURI)
}

// SetBurnAddress is a paid mutator transaction binding the contract method 0x4b0e7216.
//
// Solidity: function setBurnAddress(address newBurnAddress) returns()
func (_Token *TokenTransactor) SetBurnAddress(opts *bind.TransactOpts, newBurnAddress common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setBurnAddress", newBurnAddress)
}

// SetBurnAddress is a paid mutator transaction binding the contract method 0x4b0e7216.
//
// Solidity: function setBurnAddress(address newBurnAddress) returns()
func (_Token *TokenSession) SetBurnAddress(newBurnAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetBurnAddress(&_Token.TransactOpts, newBurnAddress)
}

// SetBurnAddress is a paid mutator transaction binding the contract method 0x4b0e7216.
//
// Solidity: function setBurnAddress(address newBurnAddress) returns()
func (_Token *TokenTransactorSession) SetBurnAddress(newBurnAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetBurnAddress(&_Token.TransactOpts, newBurnAddress)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_Token *TokenTransactor) SetContractURI(opts *bind.TransactOpts, newContractURI string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setContractURI", newContractURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_Token *TokenSession) SetContractURI(newContractURI string) (*types.Transaction, error) {
	return _Token.Contract.SetContractURI(&_Token.TransactOpts, newContractURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_Token *TokenTransactorSession) SetContractURI(newContractURI string) (*types.Transaction, error) {
	return _Token.Contract.SetContractURI(&_Token.TransactOpts, newContractURI)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_Token *TokenTransactor) SetMaxSupply(opts *bind.TransactOpts, newMaxSupply *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setMaxSupply", newMaxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_Token *TokenSession) SetMaxSupply(newMaxSupply *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetMaxSupply(&_Token.TransactOpts, newMaxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_Token *TokenTransactorSession) SetMaxSupply(newMaxSupply *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetMaxSupply(&_Token.TransactOpts, newMaxSupply)
}

// SetProvenanceHash is a paid mutator transaction binding the contract method 0x099b6bfa.
//
// Solidity: function setProvenanceHash(bytes32 newProvenanceHash) returns()
func (_Token *TokenTransactor) SetProvenanceHash(opts *bind.TransactOpts, newProvenanceHash [32]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setProvenanceHash", newProvenanceHash)
}

// SetProvenanceHash is a paid mutator transaction binding the contract method 0x099b6bfa.
//
// Solidity: function setProvenanceHash(bytes32 newProvenanceHash) returns()
func (_Token *TokenSession) SetProvenanceHash(newProvenanceHash [32]byte) (*types.Transaction, error) {
	return _Token.Contract.SetProvenanceHash(&_Token.TransactOpts, newProvenanceHash)
}

// SetProvenanceHash is a paid mutator transaction binding the contract method 0x099b6bfa.
//
// Solidity: function setProvenanceHash(bytes32 newProvenanceHash) returns()
func (_Token *TokenTransactorSession) SetProvenanceHash(newProvenanceHash [32]byte) (*types.Transaction, error) {
	return _Token.Contract.SetProvenanceHash(&_Token.TransactOpts, newProvenanceHash)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_Token *TokenTransactor) SetRoyaltyInfo(opts *bind.TransactOpts, newInfo ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setRoyaltyInfo", newInfo)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_Token *TokenSession) SetRoyaltyInfo(newInfo ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo) (*types.Transaction, error) {
	return _Token.Contract.SetRoyaltyInfo(&_Token.TransactOpts, newInfo)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_Token *TokenTransactorSession) SetRoyaltyInfo(newInfo ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo) (*types.Transaction, error) {
	return _Token.Contract.SetRoyaltyInfo(&_Token.TransactOpts, newInfo)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newPotentialOwner) returns()
func (_Token *TokenTransactor) TransferOwnership(opts *bind.TransactOpts, newPotentialOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferOwnership", newPotentialOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newPotentialOwner) returns()
func (_Token *TokenSession) TransferOwnership(newPotentialOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newPotentialOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newPotentialOwner) returns()
func (_Token *TokenTransactorSession) TransferOwnership(newPotentialOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newPotentialOwner)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0x3680620d.
//
// Solidity: function updateAllowList(address seaDropImpl, (bytes32,string[],string) allowListData) returns()
func (_Token *TokenTransactor) UpdateAllowList(opts *bind.TransactOpts, seaDropImpl common.Address, allowListData AllowListData) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateAllowList", seaDropImpl, allowListData)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0x3680620d.
//
// Solidity: function updateAllowList(address seaDropImpl, (bytes32,string[],string) allowListData) returns()
func (_Token *TokenSession) UpdateAllowList(seaDropImpl common.Address, allowListData AllowListData) (*types.Transaction, error) {
	return _Token.Contract.UpdateAllowList(&_Token.TransactOpts, seaDropImpl, allowListData)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0x3680620d.
//
// Solidity: function updateAllowList(address seaDropImpl, (bytes32,string[],string) allowListData) returns()
func (_Token *TokenTransactorSession) UpdateAllowList(seaDropImpl common.Address, allowListData AllowListData) (*types.Transaction, error) {
	return _Token.Contract.UpdateAllowList(&_Token.TransactOpts, seaDropImpl, allowListData)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x48a4c101.
//
// Solidity: function updateAllowedFeeRecipient(address seaDropImpl, address feeRecipient, bool allowed) returns()
func (_Token *TokenTransactor) UpdateAllowedFeeRecipient(opts *bind.TransactOpts, seaDropImpl common.Address, feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateAllowedFeeRecipient", seaDropImpl, feeRecipient, allowed)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x48a4c101.
//
// Solidity: function updateAllowedFeeRecipient(address seaDropImpl, address feeRecipient, bool allowed) returns()
func (_Token *TokenSession) UpdateAllowedFeeRecipient(seaDropImpl common.Address, feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _Token.Contract.UpdateAllowedFeeRecipient(&_Token.TransactOpts, seaDropImpl, feeRecipient, allowed)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x48a4c101.
//
// Solidity: function updateAllowedFeeRecipient(address seaDropImpl, address feeRecipient, bool allowed) returns()
func (_Token *TokenTransactorSession) UpdateAllowedFeeRecipient(seaDropImpl common.Address, feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _Token.Contract.UpdateAllowedFeeRecipient(&_Token.TransactOpts, seaDropImpl, feeRecipient, allowed)
}

// UpdateAllowedSeaDrop is a paid mutator transaction binding the contract method 0x60c308b6.
//
// Solidity: function updateAllowedSeaDrop(address[] allowedSeaDrop) returns()
func (_Token *TokenTransactor) UpdateAllowedSeaDrop(opts *bind.TransactOpts, allowedSeaDrop []common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateAllowedSeaDrop", allowedSeaDrop)
}

// UpdateAllowedSeaDrop is a paid mutator transaction binding the contract method 0x60c308b6.
//
// Solidity: function updateAllowedSeaDrop(address[] allowedSeaDrop) returns()
func (_Token *TokenSession) UpdateAllowedSeaDrop(allowedSeaDrop []common.Address) (*types.Transaction, error) {
	return _Token.Contract.UpdateAllowedSeaDrop(&_Token.TransactOpts, allowedSeaDrop)
}

// UpdateAllowedSeaDrop is a paid mutator transaction binding the contract method 0x60c308b6.
//
// Solidity: function updateAllowedSeaDrop(address[] allowedSeaDrop) returns()
func (_Token *TokenTransactorSession) UpdateAllowedSeaDrop(allowedSeaDrop []common.Address) (*types.Transaction, error) {
	return _Token.Contract.UpdateAllowedSeaDrop(&_Token.TransactOpts, allowedSeaDrop)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x66251b69.
//
// Solidity: function updateCreatorPayoutAddress(address seaDropImpl, address payoutAddress) returns()
func (_Token *TokenTransactor) UpdateCreatorPayoutAddress(opts *bind.TransactOpts, seaDropImpl common.Address, payoutAddress common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateCreatorPayoutAddress", seaDropImpl, payoutAddress)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x66251b69.
//
// Solidity: function updateCreatorPayoutAddress(address seaDropImpl, address payoutAddress) returns()
func (_Token *TokenSession) UpdateCreatorPayoutAddress(seaDropImpl common.Address, payoutAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.UpdateCreatorPayoutAddress(&_Token.TransactOpts, seaDropImpl, payoutAddress)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x66251b69.
//
// Solidity: function updateCreatorPayoutAddress(address seaDropImpl, address payoutAddress) returns()
func (_Token *TokenTransactorSession) UpdateCreatorPayoutAddress(seaDropImpl common.Address, payoutAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.UpdateCreatorPayoutAddress(&_Token.TransactOpts, seaDropImpl, payoutAddress)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0x7a05bc82.
//
// Solidity: function updateDropURI(address seaDropImpl, string dropURI) returns()
func (_Token *TokenTransactor) UpdateDropURI(opts *bind.TransactOpts, seaDropImpl common.Address, dropURI string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateDropURI", seaDropImpl, dropURI)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0x7a05bc82.
//
// Solidity: function updateDropURI(address seaDropImpl, string dropURI) returns()
func (_Token *TokenSession) UpdateDropURI(seaDropImpl common.Address, dropURI string) (*types.Transaction, error) {
	return _Token.Contract.UpdateDropURI(&_Token.TransactOpts, seaDropImpl, dropURI)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0x7a05bc82.
//
// Solidity: function updateDropURI(address seaDropImpl, string dropURI) returns()
func (_Token *TokenTransactorSession) UpdateDropURI(seaDropImpl common.Address, dropURI string) (*types.Transaction, error) {
	return _Token.Contract.UpdateDropURI(&_Token.TransactOpts, seaDropImpl, dropURI)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0xcb743ba8.
//
// Solidity: function updatePayer(address seaDropImpl, address payer, bool allowed) returns()
func (_Token *TokenTransactor) UpdatePayer(opts *bind.TransactOpts, seaDropImpl common.Address, payer common.Address, allowed bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updatePayer", seaDropImpl, payer, allowed)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0xcb743ba8.
//
// Solidity: function updatePayer(address seaDropImpl, address payer, bool allowed) returns()
func (_Token *TokenSession) UpdatePayer(seaDropImpl common.Address, payer common.Address, allowed bool) (*types.Transaction, error) {
	return _Token.Contract.UpdatePayer(&_Token.TransactOpts, seaDropImpl, payer, allowed)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0xcb743ba8.
//
// Solidity: function updatePayer(address seaDropImpl, address payer, bool allowed) returns()
func (_Token *TokenTransactorSession) UpdatePayer(seaDropImpl common.Address, payer common.Address, allowed bool) (*types.Transaction, error) {
	return _Token.Contract.UpdatePayer(&_Token.TransactOpts, seaDropImpl, payer, allowed)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x1b73593c.
//
// Solidity: function updatePublicDrop(address seaDropImpl, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_Token *TokenTransactor) UpdatePublicDrop(opts *bind.TransactOpts, seaDropImpl common.Address, publicDrop PublicDrop) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updatePublicDrop", seaDropImpl, publicDrop)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x1b73593c.
//
// Solidity: function updatePublicDrop(address seaDropImpl, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_Token *TokenSession) UpdatePublicDrop(seaDropImpl common.Address, publicDrop PublicDrop) (*types.Transaction, error) {
	return _Token.Contract.UpdatePublicDrop(&_Token.TransactOpts, seaDropImpl, publicDrop)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x1b73593c.
//
// Solidity: function updatePublicDrop(address seaDropImpl, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_Token *TokenTransactorSession) UpdatePublicDrop(seaDropImpl common.Address, publicDrop PublicDrop) (*types.Transaction, error) {
	return _Token.Contract.UpdatePublicDrop(&_Token.TransactOpts, seaDropImpl, publicDrop)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x511aa644.
//
// Solidity: function updateSignedMintValidationParams(address seaDropImpl, address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_Token *TokenTransactor) UpdateSignedMintValidationParams(opts *bind.TransactOpts, seaDropImpl common.Address, signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateSignedMintValidationParams", seaDropImpl, signer, signedMintValidationParams)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x511aa644.
//
// Solidity: function updateSignedMintValidationParams(address seaDropImpl, address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_Token *TokenSession) UpdateSignedMintValidationParams(seaDropImpl common.Address, signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _Token.Contract.UpdateSignedMintValidationParams(&_Token.TransactOpts, seaDropImpl, signer, signedMintValidationParams)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x511aa644.
//
// Solidity: function updateSignedMintValidationParams(address seaDropImpl, address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_Token *TokenTransactorSession) UpdateSignedMintValidationParams(seaDropImpl common.Address, signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _Token.Contract.UpdateSignedMintValidationParams(&_Token.TransactOpts, seaDropImpl, signer, signedMintValidationParams)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0x7bc2be76.
//
// Solidity: function updateTokenGatedDrop(address seaDropImpl, address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_Token *TokenTransactor) UpdateTokenGatedDrop(opts *bind.TransactOpts, seaDropImpl common.Address, allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateTokenGatedDrop", seaDropImpl, allowedNftToken, dropStage)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0x7bc2be76.
//
// Solidity: function updateTokenGatedDrop(address seaDropImpl, address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_Token *TokenSession) UpdateTokenGatedDrop(seaDropImpl common.Address, allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _Token.Contract.UpdateTokenGatedDrop(&_Token.TransactOpts, seaDropImpl, allowedNftToken, dropStage)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0x7bc2be76.
//
// Solidity: function updateTokenGatedDrop(address seaDropImpl, address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_Token *TokenTransactorSession) UpdateTokenGatedDrop(seaDropImpl common.Address, allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _Token.Contract.UpdateTokenGatedDrop(&_Token.TransactOpts, seaDropImpl, allowedNftToken, dropStage)
}

// TokenAllowedSeaDropUpdatedIterator is returned from FilterAllowedSeaDropUpdated and is used to iterate over the raw logs and unpacked data for AllowedSeaDropUpdated events raised by the Token contract.
type TokenAllowedSeaDropUpdatedIterator struct {
	Event *TokenAllowedSeaDropUpdated // Event containing the contract specifics and raw log

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
func (it *TokenAllowedSeaDropUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAllowedSeaDropUpdated)
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
		it.Event = new(TokenAllowedSeaDropUpdated)
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
func (it *TokenAllowedSeaDropUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAllowedSeaDropUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAllowedSeaDropUpdated represents a AllowedSeaDropUpdated event raised by the Token contract.
type TokenAllowedSeaDropUpdated struct {
	AllowedSeaDrop []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAllowedSeaDropUpdated is a free log retrieval operation binding the contract event 0xbbd3b69c138de4d317d0bc4290282c4e1cbd1e58b579a5b4f114b598c237454d.
//
// Solidity: event AllowedSeaDropUpdated(address[] allowedSeaDrop)
func (_Token *TokenFilterer) FilterAllowedSeaDropUpdated(opts *bind.FilterOpts) (*TokenAllowedSeaDropUpdatedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "AllowedSeaDropUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenAllowedSeaDropUpdatedIterator{contract: _Token.contract, event: "AllowedSeaDropUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedSeaDropUpdated is a free log subscription operation binding the contract event 0xbbd3b69c138de4d317d0bc4290282c4e1cbd1e58b579a5b4f114b598c237454d.
//
// Solidity: event AllowedSeaDropUpdated(address[] allowedSeaDrop)
func (_Token *TokenFilterer) WatchAllowedSeaDropUpdated(opts *bind.WatchOpts, sink chan<- *TokenAllowedSeaDropUpdated) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "AllowedSeaDropUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAllowedSeaDropUpdated)
				if err := _Token.contract.UnpackLog(event, "AllowedSeaDropUpdated", log); err != nil {
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

// ParseAllowedSeaDropUpdated is a log parse operation binding the contract event 0xbbd3b69c138de4d317d0bc4290282c4e1cbd1e58b579a5b4f114b598c237454d.
//
// Solidity: event AllowedSeaDropUpdated(address[] allowedSeaDrop)
func (_Token *TokenFilterer) ParseAllowedSeaDropUpdated(log types.Log) (*TokenAllowedSeaDropUpdated, error) {
	event := new(TokenAllowedSeaDropUpdated)
	if err := _Token.contract.UnpackLog(event, "AllowedSeaDropUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*TokenApprovalIterator, error) {

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

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Token *TokenFilterer) ParseApproval(log types.Log) (*TokenApproval, error) {
	event := new(TokenApproval)
	if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Token contract.
type TokenApprovalForAllIterator struct {
	Event *TokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApprovalForAll)
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
		it.Event = new(TokenApprovalForAll)
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
func (it *TokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApprovalForAll represents a ApprovalForAll event raised by the Token contract.
type TokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Token *TokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalForAllIterator{contract: _Token.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Token *TokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApprovalForAll)
				if err := _Token.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Token *TokenFilterer) ParseApprovalForAll(log types.Log) (*TokenApprovalForAll, error) {
	event := new(TokenApprovalForAll)
	if err := _Token.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Token contract.
type TokenBatchMetadataUpdateIterator struct {
	Event *TokenBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *TokenBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBatchMetadataUpdate)
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
		it.Event = new(TokenBatchMetadataUpdate)
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
func (it *TokenBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Token contract.
type TokenBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Token *TokenFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*TokenBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &TokenBatchMetadataUpdateIterator{contract: _Token.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Token *TokenFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *TokenBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBatchMetadataUpdate)
				if err := _Token.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_Token *TokenFilterer) ParseBatchMetadataUpdate(log types.Log) (*TokenBatchMetadataUpdate, error) {
	event := new(TokenBatchMetadataUpdate)
	if err := _Token.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Token contract.
type TokenConsecutiveTransferIterator struct {
	Event *TokenConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *TokenConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenConsecutiveTransfer)
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
		it.Event = new(TokenConsecutiveTransfer)
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
func (it *TokenConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Token contract.
type TokenConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Token *TokenFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*TokenConsecutiveTransferIterator, error) {

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

	logs, sub, err := _Token.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenConsecutiveTransferIterator{contract: _Token.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Token *TokenFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *TokenConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Token.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenConsecutiveTransfer)
				if err := _Token.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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
func (_Token *TokenFilterer) ParseConsecutiveTransfer(log types.Log) (*TokenConsecutiveTransfer, error) {
	event := new(TokenConsecutiveTransfer)
	if err := _Token.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenContractURIUpdatedIterator is returned from FilterContractURIUpdated and is used to iterate over the raw logs and unpacked data for ContractURIUpdated events raised by the Token contract.
type TokenContractURIUpdatedIterator struct {
	Event *TokenContractURIUpdated // Event containing the contract specifics and raw log

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
func (it *TokenContractURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenContractURIUpdated)
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
		it.Event = new(TokenContractURIUpdated)
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
func (it *TokenContractURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenContractURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenContractURIUpdated represents a ContractURIUpdated event raised by the Token contract.
type TokenContractURIUpdated struct {
	NewContractURI string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContractURIUpdated is a free log retrieval operation binding the contract event 0x905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac37378.
//
// Solidity: event ContractURIUpdated(string newContractURI)
func (_Token *TokenFilterer) FilterContractURIUpdated(opts *bind.FilterOpts) (*TokenContractURIUpdatedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenContractURIUpdatedIterator{contract: _Token.contract, event: "ContractURIUpdated", logs: logs, sub: sub}, nil
}

// WatchContractURIUpdated is a free log subscription operation binding the contract event 0x905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac37378.
//
// Solidity: event ContractURIUpdated(string newContractURI)
func (_Token *TokenFilterer) WatchContractURIUpdated(opts *bind.WatchOpts, sink chan<- *TokenContractURIUpdated) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenContractURIUpdated)
				if err := _Token.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
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
func (_Token *TokenFilterer) ParseContractURIUpdated(log types.Log) (*TokenContractURIUpdated, error) {
	event := new(TokenContractURIUpdated)
	if err := _Token.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Token contract.
type TokenInitializedIterator struct {
	Event *TokenInitialized // Event containing the contract specifics and raw log

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
func (it *TokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenInitialized)
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
		it.Event = new(TokenInitialized)
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
func (it *TokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenInitialized represents a Initialized event raised by the Token contract.
type TokenInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Token *TokenFilterer) FilterInitialized(opts *bind.FilterOpts) (*TokenInitializedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenInitializedIterator{contract: _Token.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Token *TokenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenInitialized) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenInitialized)
				if err := _Token.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Token *TokenFilterer) ParseInitialized(log types.Log) (*TokenInitialized, error) {
	event := new(TokenInitialized)
	if err := _Token.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenMaxSupplyUpdatedIterator is returned from FilterMaxSupplyUpdated and is used to iterate over the raw logs and unpacked data for MaxSupplyUpdated events raised by the Token contract.
type TokenMaxSupplyUpdatedIterator struct {
	Event *TokenMaxSupplyUpdated // Event containing the contract specifics and raw log

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
func (it *TokenMaxSupplyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMaxSupplyUpdated)
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
		it.Event = new(TokenMaxSupplyUpdated)
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
func (it *TokenMaxSupplyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMaxSupplyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMaxSupplyUpdated represents a MaxSupplyUpdated event raised by the Token contract.
type TokenMaxSupplyUpdated struct {
	NewMaxSupply *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMaxSupplyUpdated is a free log retrieval operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 newMaxSupply)
func (_Token *TokenFilterer) FilterMaxSupplyUpdated(opts *bind.FilterOpts) (*TokenMaxSupplyUpdatedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenMaxSupplyUpdatedIterator{contract: _Token.contract, event: "MaxSupplyUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxSupplyUpdated is a free log subscription operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 newMaxSupply)
func (_Token *TokenFilterer) WatchMaxSupplyUpdated(opts *bind.WatchOpts, sink chan<- *TokenMaxSupplyUpdated) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMaxSupplyUpdated)
				if err := _Token.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
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
func (_Token *TokenFilterer) ParseMaxSupplyUpdated(log types.Log) (*TokenMaxSupplyUpdated, error) {
	event := new(TokenMaxSupplyUpdated)
	if err := _Token.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Token contract.
type TokenOwnershipTransferredIterator struct {
	Event *TokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenOwnershipTransferred)
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
		it.Event = new(TokenOwnershipTransferred)
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
func (it *TokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenOwnershipTransferred represents a OwnershipTransferred event raised by the Token contract.
type TokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenOwnershipTransferredIterator{contract: _Token.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenOwnershipTransferred)
				if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Token *TokenFilterer) ParseOwnershipTransferred(log types.Log) (*TokenOwnershipTransferred, error) {
	event := new(TokenOwnershipTransferred)
	if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenPotentialOwnerUpdatedIterator is returned from FilterPotentialOwnerUpdated and is used to iterate over the raw logs and unpacked data for PotentialOwnerUpdated events raised by the Token contract.
type TokenPotentialOwnerUpdatedIterator struct {
	Event *TokenPotentialOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *TokenPotentialOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPotentialOwnerUpdated)
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
		it.Event = new(TokenPotentialOwnerUpdated)
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
func (it *TokenPotentialOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenPotentialOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenPotentialOwnerUpdated represents a PotentialOwnerUpdated event raised by the Token contract.
type TokenPotentialOwnerUpdated struct {
	NewPotentialAdministrator common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterPotentialOwnerUpdated is a free log retrieval operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address newPotentialAdministrator)
func (_Token *TokenFilterer) FilterPotentialOwnerUpdated(opts *bind.FilterOpts) (*TokenPotentialOwnerUpdatedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "PotentialOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenPotentialOwnerUpdatedIterator{contract: _Token.contract, event: "PotentialOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchPotentialOwnerUpdated is a free log subscription operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address newPotentialAdministrator)
func (_Token *TokenFilterer) WatchPotentialOwnerUpdated(opts *bind.WatchOpts, sink chan<- *TokenPotentialOwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "PotentialOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenPotentialOwnerUpdated)
				if err := _Token.contract.UnpackLog(event, "PotentialOwnerUpdated", log); err != nil {
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
func (_Token *TokenFilterer) ParsePotentialOwnerUpdated(log types.Log) (*TokenPotentialOwnerUpdated, error) {
	event := new(TokenPotentialOwnerUpdated)
	if err := _Token.contract.UnpackLog(event, "PotentialOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenProvenanceHashUpdatedIterator is returned from FilterProvenanceHashUpdated and is used to iterate over the raw logs and unpacked data for ProvenanceHashUpdated events raised by the Token contract.
type TokenProvenanceHashUpdatedIterator struct {
	Event *TokenProvenanceHashUpdated // Event containing the contract specifics and raw log

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
func (it *TokenProvenanceHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenProvenanceHashUpdated)
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
		it.Event = new(TokenProvenanceHashUpdated)
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
func (it *TokenProvenanceHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenProvenanceHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenProvenanceHashUpdated represents a ProvenanceHashUpdated event raised by the Token contract.
type TokenProvenanceHashUpdated struct {
	PreviousHash [32]byte
	NewHash      [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProvenanceHashUpdated is a free log retrieval operation binding the contract event 0x7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c.
//
// Solidity: event ProvenanceHashUpdated(bytes32 previousHash, bytes32 newHash)
func (_Token *TokenFilterer) FilterProvenanceHashUpdated(opts *bind.FilterOpts) (*TokenProvenanceHashUpdatedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "ProvenanceHashUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenProvenanceHashUpdatedIterator{contract: _Token.contract, event: "ProvenanceHashUpdated", logs: logs, sub: sub}, nil
}

// WatchProvenanceHashUpdated is a free log subscription operation binding the contract event 0x7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c.
//
// Solidity: event ProvenanceHashUpdated(bytes32 previousHash, bytes32 newHash)
func (_Token *TokenFilterer) WatchProvenanceHashUpdated(opts *bind.WatchOpts, sink chan<- *TokenProvenanceHashUpdated) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "ProvenanceHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenProvenanceHashUpdated)
				if err := _Token.contract.UnpackLog(event, "ProvenanceHashUpdated", log); err != nil {
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
func (_Token *TokenFilterer) ParseProvenanceHashUpdated(log types.Log) (*TokenProvenanceHashUpdated, error) {
	event := new(TokenProvenanceHashUpdated)
	if err := _Token.contract.UnpackLog(event, "ProvenanceHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRoyaltyInfoUpdatedIterator is returned from FilterRoyaltyInfoUpdated and is used to iterate over the raw logs and unpacked data for RoyaltyInfoUpdated events raised by the Token contract.
type TokenRoyaltyInfoUpdatedIterator struct {
	Event *TokenRoyaltyInfoUpdated // Event containing the contract specifics and raw log

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
func (it *TokenRoyaltyInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRoyaltyInfoUpdated)
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
		it.Event = new(TokenRoyaltyInfoUpdated)
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
func (it *TokenRoyaltyInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRoyaltyInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRoyaltyInfoUpdated represents a RoyaltyInfoUpdated event raised by the Token contract.
type TokenRoyaltyInfoUpdated struct {
	Receiver common.Address
	Bps      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyInfoUpdated is a free log retrieval operation binding the contract event 0xf21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d.
//
// Solidity: event RoyaltyInfoUpdated(address receiver, uint256 bps)
func (_Token *TokenFilterer) FilterRoyaltyInfoUpdated(opts *bind.FilterOpts) (*TokenRoyaltyInfoUpdatedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "RoyaltyInfoUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenRoyaltyInfoUpdatedIterator{contract: _Token.contract, event: "RoyaltyInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchRoyaltyInfoUpdated is a free log subscription operation binding the contract event 0xf21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d.
//
// Solidity: event RoyaltyInfoUpdated(address receiver, uint256 bps)
func (_Token *TokenFilterer) WatchRoyaltyInfoUpdated(opts *bind.WatchOpts, sink chan<- *TokenRoyaltyInfoUpdated) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "RoyaltyInfoUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRoyaltyInfoUpdated)
				if err := _Token.contract.UnpackLog(event, "RoyaltyInfoUpdated", log); err != nil {
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
func (_Token *TokenFilterer) ParseRoyaltyInfoUpdated(log types.Log) (*TokenRoyaltyInfoUpdated, error) {
	event := new(TokenRoyaltyInfoUpdated)
	if err := _Token.contract.UnpackLog(event, "RoyaltyInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSeaDropTokenDeployedIterator is returned from FilterSeaDropTokenDeployed and is used to iterate over the raw logs and unpacked data for SeaDropTokenDeployed events raised by the Token contract.
type TokenSeaDropTokenDeployedIterator struct {
	Event *TokenSeaDropTokenDeployed // Event containing the contract specifics and raw log

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
func (it *TokenSeaDropTokenDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSeaDropTokenDeployed)
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
		it.Event = new(TokenSeaDropTokenDeployed)
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
func (it *TokenSeaDropTokenDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSeaDropTokenDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSeaDropTokenDeployed represents a SeaDropTokenDeployed event raised by the Token contract.
type TokenSeaDropTokenDeployed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSeaDropTokenDeployed is a free log retrieval operation binding the contract event 0xd7aca75208b9be5ffc04c6a01922020ffd62b55e68e502e317f5344960279af8.
//
// Solidity: event SeaDropTokenDeployed()
func (_Token *TokenFilterer) FilterSeaDropTokenDeployed(opts *bind.FilterOpts) (*TokenSeaDropTokenDeployedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "SeaDropTokenDeployed")
	if err != nil {
		return nil, err
	}
	return &TokenSeaDropTokenDeployedIterator{contract: _Token.contract, event: "SeaDropTokenDeployed", logs: logs, sub: sub}, nil
}

// WatchSeaDropTokenDeployed is a free log subscription operation binding the contract event 0xd7aca75208b9be5ffc04c6a01922020ffd62b55e68e502e317f5344960279af8.
//
// Solidity: event SeaDropTokenDeployed()
func (_Token *TokenFilterer) WatchSeaDropTokenDeployed(opts *bind.WatchOpts, sink chan<- *TokenSeaDropTokenDeployed) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "SeaDropTokenDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSeaDropTokenDeployed)
				if err := _Token.contract.UnpackLog(event, "SeaDropTokenDeployed", log); err != nil {
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

// ParseSeaDropTokenDeployed is a log parse operation binding the contract event 0xd7aca75208b9be5ffc04c6a01922020ffd62b55e68e502e317f5344960279af8.
//
// Solidity: event SeaDropTokenDeployed()
func (_Token *TokenFilterer) ParseSeaDropTokenDeployed(log types.Log) (*TokenSeaDropTokenDeployed, error) {
	event := new(TokenSeaDropTokenDeployed)
	if err := _Token.contract.UnpackLog(event, "SeaDropTokenDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*TokenTransferIterator, error) {

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

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Token *TokenFilterer) ParseTransfer(log types.Log) (*TokenTransfer, error) {
	event := new(TokenTransfer)
	if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
