// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package partner

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

// ERC721SeaDropStructsErrorsAndEventsMultiConfigureStruct is an auto generated low-level Go binding around an user-defined struct.
type ERC721SeaDropStructsErrorsAndEventsMultiConfigureStruct struct {
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

// ISeaDropTokenContractMetadataRoyaltyInfo is an auto generated low-level Go binding around an user-defined struct.
type ISeaDropTokenContractMetadataRoyaltyInfo struct {
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

// PartnerMetaData contains all meta data concerning the Partner contract.
var PartnerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"allowedSeaDrop\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AdministratorMustInitializeWithFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"CannotExceedMaxSupplyOfUint64\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"basisPoints\",\"type\":\"uint256\"}],\"name\":\"InvalidRoyaltyBasisPoints\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAdministratorIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotNextAdministrator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotNextOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdministrator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedSeaDrop\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrAdministrator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProvenanceHashCannotBeSetAfterMintStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoyaltyAddressCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignersMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousAdministrator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdministrator\",\"type\":\"address\"}],\"name\":\"AdministratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowedSeaDrop\",\"type\":\"address[]\"}],\"name\":\"AllowedSeaDropUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"}],\"name\":\"ContractURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPotentialAdministrator\",\"type\":\"address\"}],\"name\":\"PotentialAdministratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPotentialAdministrator\",\"type\":\"address\"}],\"name\":\"PotentialOwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newHash\",\"type\":\"bytes32\"}],\"name\":\"ProvenanceHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"name\":\"RoyaltyInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SeaDropTokenDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OPERATOR_FILTER_REGISTRY\",\"outputs\":[{\"internalType\":\"contractIOperatorFilterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdministration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"administrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelAdministrationTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"}],\"name\":\"emitBatchMetadataUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"getMintStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minterNumMinted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"mintSeaDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contractURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"dropURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"publicKeyURIs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"internalType\":\"structAllowListData\",\"name\":\"allowListData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"creatorPayoutAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"provenanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"allowedFeeRecipients\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedFeeRecipients\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowedPayers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedPayers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenGatedAllowedNftTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structTokenGatedDropStage[]\",\"name\":\"tokenGatedDropStages\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedTokenGatedAllowedNftTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"internalType\":\"structSignedMintValidationParams[]\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedSigners\",\"type\":\"address[]\"}],\"internalType\":\"structERC721SeaDropStructsErrorsAndEvents.MultiConfigureStruct\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"multiConfigure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"potentialAdministrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provenanceHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdministration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newProvenanceHash\",\"type\":\"bytes32\"}],\"name\":\"setProvenanceHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"royaltyAddress\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"royaltyBps\",\"type\":\"uint96\"}],\"internalType\":\"structISeaDropTokenContractMetadata.RoyaltyInfo\",\"name\":\"newInfo\",\"type\":\"tuple\"}],\"name\":\"setRoyaltyInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdministrator\",\"type\":\"address\"}],\"name\":\"transferAdministration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPotentialOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"publicKeyURIs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"internalType\":\"structAllowListData\",\"name\":\"allowListData\",\"type\":\"tuple\"}],\"name\":\"updateAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateAllowedFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"allowedSeaDrop\",\"type\":\"address[]\"}],\"name\":\"updateAllowedSeaDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payoutAddress\",\"type\":\"address\"}],\"name\":\"updateCreatorPayoutAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"dropURI\",\"type\":\"string\"}],\"name\":\"updateDropURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updatePayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"}],\"name\":\"updatePublicDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"internalType\":\"structSignedMintValidationParams\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple\"}],\"name\":\"updateSignedMintValidationParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structTokenGatedDropStage\",\"name\":\"dropStage\",\"type\":\"tuple\"}],\"name\":\"updateTokenGatedDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526001600f553480156200001657600080fd5b5060405162004c0738038062004c078339810160408190526200003991620004f3565b81848483733cc6cdda760b79bafa08df41ecfa224f810dceb6600184848181600262000066838262000699565b50600362000075828262000699565b50506001600055506200008762000294565b50506daaeb6d7670e522a718067333cd4e3b15620001ce5780156200011c57604051633e9f1edf60e11b81523060048201526001600160a01b03831660248201526daaeb6d7670e522a718067333cd4e90637d3e3dbe906044015b600060405180830381600087803b158015620000fd57600080fd5b505af115801562000112573d6000803e3d6000fd5b50505050620001ce565b6001600160a01b038216156200016d5760405163a0af290360e01b81523060048201526001600160a01b03831660248201526daaeb6d7670e522a718067333cd4e9063a0af290390604401620000e2565b604051632210724360e11b81523060048201526daaeb6d7670e522a718067333cd4e90634420e48690602401600060405180830381600087803b158015620001b457600080fd5b505af1158015620001c9573d6000803e3d6000fd5b505050505b5050805160005b818110156200023457600160106000858481518110620001f957620001f962000765565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055600101620001d5565b5081516200024a9060119060208501906200037f565b506040517fd7aca75208b9be5ffc04c6a01922020ffd62b55e68e502e317f5344960279af890600090a1505050506200028981620002c160201b60201c565b50505050506200077b565b303b15620002b45760405162dc149f60e41b815260040160405180910390fd5b620002bf336200032d565b565b303b15620002e15760405162dc149f60e41b815260040160405180910390fd5b601280546001600160a01b0319166001600160a01b0383169081179091556040516000907f222c8e95a03c7aa95fc5d110485e0d38e767f07ab1ec878a6eac644ef1d83122908290a350565b600880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b828054828255906000526020600020908101928215620003d7579160200282015b82811115620003d757825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620003a0565b50620003e5929150620003e9565b5090565b5b80821115620003e55760008155600101620003ea565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171562000441576200044162000400565b604052919050565b600082601f8301126200045b57600080fd5b81516001600160401b0381111562000477576200047762000400565b60206200048d601f8301601f1916820162000416565b8281528582848701011115620004a257600080fd5b60005b83811015620004c2578581018301518282018401528201620004a5565b506000928101909101919091529392505050565b80516001600160a01b0381168114620004ee57600080fd5b919050565b600080600080608085870312156200050a57600080fd5b84516001600160401b03808211156200052257600080fd5b620005308883890162000449565b95506020915081870151818111156200054857600080fd5b6200055689828a0162000449565b9550506200056760408801620004d6565b93506060870151818111156200057c57600080fd5b8701601f810189136200058e57600080fd5b805182811115620005a357620005a362000400565b8060051b9250620005b684840162000416565b818152928201840192848101908b851115620005d157600080fd5b928501925b84841015620005fa57620005ea84620004d6565b82529285019290850190620005d6565b989b979a50959850505050505050565b600181811c908216806200061f57607f821691505b6020821081036200064057634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200069457600081815260208120601f850160051c810160208610156200066f5750805b601f850160051c820191505b8181101562000690578281556001016200067b565b5050505b505050565b81516001600160401b03811115620006b557620006b562000400565b620006cd81620006c684546200060a565b8462000646565b602080601f831160018114620007055760008415620006ec5750858301515b600019600386901b1c1916600185901b17855562000690565b600085815260208120601f198616915b82811015620007365788860151825594840194600190910190840162000715565b5085821015620007555787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b61447c806200078b6000396000f3fe608060405234801561001057600080fd5b50600436106102f15760003560e01c8063715018a61161019d578063ad2f852a116100e9578063d5abeb01116100a2578063e985e9c51161007c578063e985e9c51461067a578063f0025d961461068d578063f2fde38b14610695578063f53d0a8e146106a857600080fd5b8063d5abeb0114610657578063d5e7feb81461065f578063e8a3d4851461067257600080fd5b8063ad2f852a146105fd578063b88d4fde1461060e578063c6ab67a314610621578063c780b63d14610629578063c87b56dd14610631578063cb743ba81461064457600080fd5b8063911f456b1161015657806395d89b411161013057806395d89b41146105c75780639794ed40146105cf578063a22cb465146105d7578063a4830114146105ea57600080fd5b8063911f456b1461058e578063913ee93d146105a1578063938e3d7b146105b457600080fd5b8063715018a61461051957806379ba5097146105215780637a05bc82146105295780637bc2be761461053c578063840e15d41461054f5780638da5cb5b1461057d57600080fd5b806342260b5d1161025c57806360c308b61161021557806366251b69116101ef57806366251b69146104d85780636c0360eb146104eb5780636f8b44b0146104f357806370a082311461050657600080fd5b806360c308b61461049f5780636352211e146104b257806364869dad146104c557600080fd5b806342260b5d1461042857806342842e0e1461044057806344dae42c1461045357806348a4c10114610466578063511aa6441461047957806355f804b31461048c57600080fd5b80631b73593c116102ae5780631b73593c146103a057806323452b9c146103b357806323b872dd146103bb5780632a55205a146103ce5780633680620d1461040057806341f434341461041357600080fd5b806301ffc9a7146102f657806306fdde031461031e578063081812fc14610333578063095ea7b31461035e578063099b6bfa1461037357806318160ddd14610386575b600080fd5b610309610304366004612f21565b6106bb565b60405190151581526020015b60405180910390f35b610326610701565b6040516103159190612f8e565b610346610341366004612fa1565b610793565b6040516001600160a01b039091168152602001610315565b61037161036c366004612fcf565b6107d7565b005b610371610381366004612fa1565b6107f0565b60015460005403600019015b604051908152602001610315565b6103716103ae366004612ffb565b610863565b610371610a6a565b6103716103c936600461303c565b610ab9565b6103e16103dc36600461307d565b610ae4565b604080516001600160a01b039093168352602083019190915201610315565b61037161040e36600461309f565b610b2a565b6103466daaeb6d7670e522a718067333cd4e81565b600e54600160a01b90046001600160601b0316610392565b61037161044e36600461303c565b610bd2565b6103716104613660046130f5565b610bf7565b61037161047436600461312b565b610d0d565b610371610487366004613288565b610dad565b61037161049a3660046133a7565b610f77565b6103716104ad3660046133e8565b610fed565b6103466104c0366004612fa1565b611034565b6103716104d3366004612fcf565b61103f565b6103716104e636600461345c565b6110b0565b6103266110ef565b610371610501366004612fa1565b6110fe565b61039261051436600461348a565b611166565b6103716111b4565b6103716111c8565b6103716105373660046134a7565b611244565b61037161054a3660046134fb565b6112b8565b61056261055d36600461348a565b611446565b60408051938452602084019290925290820152606001610315565b6008546001600160a01b0316610346565b61037161059c366004613550565b611487565b6103716105af36600461348a565b6120a4565b6103716105c23660046133a7565b612147565b61032661218e565b61037161219d565b6103716105e536600461358b565b61220c565b6103716105f836600461307d565b612220565b600e546001600160a01b0316610346565b61037161061c3660046135b9565b61225e565b600d54610392565b61037161228b565b61032661063f366004612fa1565b6122f8565b61037161065236600461312b565b61237c565b600a54610392565b601354610346906001600160a01b031681565b6103266123f8565b61030961068836600461345c565b612407565b610371612435565b6103716106a336600461348a565b61247d565b601254610346906001600160a01b031681565b60006001600160e01b03198216630c487f4760e11b14806106ec57506001600160e01b03198216639c15441560e01b145b806106fb57506106fb826124fa565b92915050565b6060600280546107109061367c565b80601f016020809104026020016040519081016040528092919081815260200182805461073c9061367c565b80156107895780601f1061075e57610100808354040283529160200191610789565b820191906000526020600020905b81548152906001019060200180831161076c57829003601f168201915b5050505050905090565b600061079e8261253a565b6107bb576040516333d1c03960e21b815260040160405180910390fd5b506000908152600660205260409020546001600160a01b031690565b816107e18161256f565b6107eb8383612628565b505050565b6107f86126c8565b600054600019011561081d5760405163e03264af60e01b815260040160405180910390fd5b600d80549082905560408051828152602081018490527f7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c91015b60405180910390a15050565b6008546001600160a01b031633146108a0576012546001600160a01b031633146108a0576040516359d9793760e01b815260040160405180910390fd5b6108a982612718565b604051632f1a98a760e21b81523060048201526000906001600160a01b0384169063bc6a629c9060240160c060405180830381865afa1580156108f0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061091491906136e5565b905060006109273684900384018461377d565b6012549091506001600160a01b0316331461097f57816060015161ffff1660000361096557604051634f4ca83d60e11b815260040160405180910390fd5b60808083015161ffff1690820152600160a08201526109b9565b606082015161ffff8116610994576001610996565b805b61ffff9081166060850152608092830151169183019190915250600160a0820152805b604080516301308e6560e01b815282516001600160501b03166004820152602083015165ffffffffffff9081166024830152918301519091166044820152606082015161ffff9081166064830152608083015116608482015260a0820151151560a48201526001600160a01b038516906301308e659060c401600060405180830381600087803b158015610a4c57600080fd5b505af1158015610a60573d6000803e3d6000fd5b5050505050505050565b610a72612756565b600980546001600160a01b0319169055604051600081527f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da906020015b60405180910390a1565b826001600160a01b0381163314610ad357610ad33361256f565b610ade848484612781565b50505050565b600e8054600091829161271090610b0b90600160a01b90046001600160601b031686613814565b610b15919061382b565b90546001600160a01b03169590945092505050565b6008546001600160a01b03163314610b67576012546001600160a01b03163314610b67576040516359d9793760e01b815260040160405180910390fd5b610b7082612718565b60405163ebb4a55f60e01b81526001600160a01b0383169063ebb4a55f90610b9c90849060040161398e565b600060405180830381600087803b158015610bb657600080fd5b505af1158015610bca573d6000803e3d6000fd5b505050505050565b826001600160a01b0381163314610bec57610bec3361256f565b610ade848484612916565b610bff6126c8565b6000610c0e602083018361348a565b6001600160a01b031603610c3557604051631cc0baef60e01b815260040160405180910390fd5b612710610c4860408301602084016139b6565b6001600160601b03161115610c9157610c6760408201602083016139b6565b604051633cadbafb60e01b81526001600160601b0390911660048201526024015b60405180910390fd5b80600e610c9e82826139d3565b507ff21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d9050610ccf602083018361348a565b610cdf60408401602085016139b6565b604080516001600160a01b0390931683526001600160601b039091166020830152015b60405180910390a150565b6012546001600160a01b03163314610d3b576040516001620aed3360e41b0319815260040160405180910390fd5b610d4483612718565b604051638e7d1e4360e01b81526001600160a01b0383811660048301528215156024830152841690638e7d1e43906044015b600060405180830381600087803b158015610d9057600080fd5b505af1158015610da4573d6000803e3d6000fd5b50505050505050565b6008546001600160a01b03163314610dea576012546001600160a01b03163314610dea576040516359d9793760e01b815260040160405180910390fd5b610df383612718565b6040516381bf9af360e01b81523060048201526001600160a01b038381166024830152600091908516906381bf9af39060440160e060405180830381865afa158015610e43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e679190613a20565b60125490915082906001600160a01b03163314610ecd57816020015162ffffff16600003610ea857604051634f4ca83d60e11b815260040160405180910390fd5b60a08083015161ffff9081169183019190915260c08084015190911690820152610f10565b602082015162ffffff8116610ee3576001610ee5565b805b62ffffff1660208401525060a08082015161ffff9081169184019190915260c0918201511690820152805b6040516309a7002f60e31b81526001600160a01b03861690634d38017890610f3e9087908590600401613b42565b600060405180830381600087803b158015610f5857600080fd5b505af1158015610f6c573d6000803e3d6000fd5b505050505050505050565b610f7f6126c8565b600b610f8c828483613ba6565b50600154600054036000190115610fe9577f6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c600180610fca60005490565b610fd49190613c65565b60408051928352602083019190915201610857565b5050565b6008546001600160a01b0316331461102a576012546001600160a01b0316331461102a576040516359d9793760e01b815260040160405180910390fd5b610fe98282612931565b60006106fb82612a4b565b61104833612718565b600a54816110596000546000190190565b6110639190613c78565b11156110a657806110776000546000190190565b6110819190613c78565b600a5460405163384b48c560e21b815260048101929092526024820152604401610c88565b610fe98282612aba565b6110b86126c8565b6110c182612718565b60405163024e71b760e31b81526001600160a01b0382811660048301528316906312738db890602401610b9c565b60606110f9612bb8565b905090565b6111066126c8565b6001600160401b038111156111315760405163b43e913760e01b815260048101829052602401610c88565b600a8190556040518181527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c90602001610d02565b60006001600160a01b03821661118f576040516323d3ad8160e21b815260040160405180910390fd5b506001600160a01b03166000908152600560205260409020546001600160401b031690565b6111bc612756565b6111c66000612bc7565b565b6009546001600160a01b03163381146111f457604051636b7584e760e11b815260040160405180910390fd5b600980546001600160a01b0319169055604051600081527f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da9060200160405180910390a161124181612bc7565b50565b6008546001600160a01b03163314611281576012546001600160a01b03163314611281576040516359d9793760e01b815260040160405180910390fd5b61128a83612718565b60405163b957d0cb60e01b81526001600160a01b0384169063b957d0cb90610d769085908590600401613c8b565b6008546001600160a01b031633146112f5576012546001600160a01b031633146112f5576040516359d9793760e01b815260040160405180910390fd5b6112fe83612718565b604051630587453760e11b81523060048201526001600160a01b03838116602483015260009190851690630b0e8a6e9060440161010060405180830381865afa15801561134f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113739190613cd6565b9050600061138636849003840184613d94565b6012549091506001600160a01b031633146113de57816020015161ffff166000036113c457604051634f4ca83d60e11b815260040160405180910390fd5b60c08083015161ffff1690820152600160e0820152611418565b602082015161ffff81166113f35760016113f5565b805b61ffff908116602085015260c092830151169183019190915250600160e0820152805b604051637ecd591560e11b81526001600160a01b0386169063fd9ab22a90610f3e9087908590600401613ebb565b6001600160a01b03811660009081526005602052604080822054901c6001600160401b0316908061147a6000546000190190565b600a549395909450915050565b61148f612756565b8035156114e8576040516306f8b44b60e41b8152813560048201523090636f8b44b090602401600060405180830381600087803b1580156114cf57600080fd5b505af11580156114e3573d6000803e3d6000fd5b505050505b6114f56020820182613ed9565b15905061155f57306355f804b361150f6020840184613ed9565b6040518363ffffffff1660e01b815260040161152c929190613c8b565b600060405180830381600087803b15801561154657600080fd5b505af115801561155a573d6000803e3d6000fd5b505050505b61156c6040820182613ed9565b1590506115d6573063938e3d7b6115866040840184613ed9565b6040518363ffffffff1660e01b81526004016115a3929190613c8b565b600060405180830381600087803b1580156115bd57600080fd5b505af11580156115d1573d6000803e3d6000fd5b505050505b6115f66115e960e0830160c08401613f1f565b65ffffffffffff16151590565b6116096115e960c0840160a08501613f1f565b1760010361167b5730631b73593c611627608084016060850161348a565b836080016040518363ffffffff1660e01b8152600401611648929190613f3c565b600060405180830381600087803b15801561166257600080fd5b505af1158015611676573d6000803e3d6000fd5b505050505b611689610140820182613ed9565b1590506117055730637a05bc826116a6608084016060850161348a565b6116b4610140850185613ed9565b6040518463ffffffff1660e01b81526004016116d293929190613fea565b600060405180830381600087803b1580156116ec57600080fd5b505af1158015611700573d6000803e3d6000fd5b505050505b6000611715610160830183614018565b351461178f5730633680620d611731608084016060850161348a565b61173f610160850185614018565b6040518363ffffffff1660e01b815260040161175c929190614038565b600060405180830381600087803b15801561177657600080fd5b505af115801561178a573d6000803e3d6000fd5b505050505b60006117a36101a08301610180840161348a565b6001600160a01b03161461183a57306366251b696117c7608084016060850161348a565b6117d96101a08501610180860161348a565b6040516001600160e01b031960e085901b1681526001600160a01b03928316600482015291166024820152604401600060405180830381600087803b15801561182157600080fd5b505af1158015611835573d6000803e3d6000fd5b505050505b6101a08101351561189b576040516304cdb5fd60e11b81526101a08201356004820152309063099b6bfa90602401600060405180830381600087803b15801561188257600080fd5b505af1158015611896573d6000803e3d6000fd5b505050505b60006118ab6101c083018361405c565b905011156119755760005b6118c46101c083018361405c565b905081101561197357306348a4c1016118e3608085016060860161348a565b6118f16101c086018661405c565b85818110611901576119016140a5565b9050602002016020810190611916919061348a565b60016040518463ffffffff1660e01b8152600401611936939291906140bb565b600060405180830381600087803b15801561195057600080fd5b505af1158015611964573d6000803e3d6000fd5b505050508060010190506118b6565b505b60006119856101e083018361405c565b90501115611a4f5760005b61199e6101e083018361405c565b9050811015611a4d57306348a4c1016119bd608085016060860161348a565b6119cb6101e086018661405c565b858181106119db576119db6140a5565b90506020020160208101906119f0919061348a565b60006040518463ffffffff1660e01b8152600401611a10939291906140bb565b600060405180830381600087803b158015611a2a57600080fd5b505af1158015611a3e573d6000803e3d6000fd5b50505050806001019050611990565b505b6000611a5f61020083018361405c565b90501115611b295760005b611a7861020083018361405c565b9050811015611b27573063cb743ba8611a97608085016060860161348a565b611aa561020086018661405c565b85818110611ab557611ab56140a5565b9050602002016020810190611aca919061348a565b60016040518463ffffffff1660e01b8152600401611aea939291906140bb565b600060405180830381600087803b158015611b0457600080fd5b505af1158015611b18573d6000803e3d6000fd5b50505050806001019050611a6a565b505b6000611b3961022083018361405c565b90501115611c035760005b611b5261022083018361405c565b9050811015611c01573063cb743ba8611b71608085016060860161348a565b611b7f61022086018661405c565b85818110611b8f57611b8f6140a5565b9050602002016020810190611ba4919061348a565b60006040518463ffffffff1660e01b8152600401611bc4939291906140bb565b600060405180830381600087803b158015611bde57600080fd5b505af1158015611bf2573d6000803e3d6000fd5b50505050806001019050611b44565b505b6000611c136102608301836140df565b90501115611d3e57611c2961024082018261405c565b9050611c396102608301836140df565b905014611c595760405163b81aa63960e01b815260040160405180910390fd5b60005b611c6a6102608301836140df565b9050811015611d3c5730637bc2be76611c89608085016060860161348a565b611c9761024086018661405c565b85818110611ca757611ca76140a5565b9050602002016020810190611cbc919061348a565b611cca6102608701876140df565b86818110611cda57611cda6140a5565b905061010002016040518463ffffffff1660e01b8152600401611cff93929190614128565b600060405180830381600087803b158015611d1957600080fd5b505af1158015611d2d573d6000803e3d6000fd5b50505050806001019050611c5c565b505b6000611d4e61028083018361405c565b90501115611e595760005b611d6761028083018361405c565b9050811015611e57576040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915230637bc2be76611dc7608086016060870161348a565b611dd561028087018761405c565b86818110611de557611de56140a5565b9050602002016020810190611dfa919061348a565b846040518463ffffffff1660e01b8152600401611e1993929190614203565b600060405180830381600087803b158015611e3357600080fd5b505af1158015611e47573d6000803e3d6000fd5b5050505081600101915050611d59565b505b6000611e696102c0830183614229565b90501115611f9357611e7f6102a082018261405c565b9050611e8f6102c0830183614229565b905014611eaf576040516374ef6df760e01b815260040160405180910390fd5b60005b611ec06102c0830183614229565b9050811015611f91573063511aa644611edf608085016060860161348a565b611eed6102a086018661405c565b85818110611efd57611efd6140a5565b9050602002016020810190611f12919061348a565b611f206102c0870187614229565b86818110611f3057611f306140a5565b905060e002016040518463ffffffff1660e01b8152600401611f5493929190614271565b600060405180830381600087803b158015611f6e57600080fd5b505af1158015611f82573d6000803e3d6000fd5b50505050806001019050611eb2565b505b6000611fa36102e083018361405c565b905011156112415760005b611fbc6102e083018361405c565b9050811015610fe9576040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091523063511aa644612014608086016060870161348a565b6120226102e087018761405c565b86818110612032576120326140a5565b9050602002016020810190612047919061348a565b846040518463ffffffff1660e01b815260040161206693929190614336565b600060405180830381600087803b15801561208057600080fd5b505af1158015612094573d6000803e3d6000fd5b5050505081600101915050611fae565b6012546001600160a01b031633146120d2576040516001620aed3360e41b0319815260040160405180910390fd5b6001600160a01b0381166120f957604051633536be7f60e21b815260040160405180910390fd5b601380546001600160a01b0319166001600160a01b0383169081179091556040519081527fffa60f32d5278b35b1a3350ca92518fb5fe53e54ad07ac6355a17f54c5296b1f90602001610d02565b61214f6126c8565b600c61215c828483613ba6565b507f905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac373788282604051610857929190613c8b565b6060600380546107109061367c565b6012546001600160a01b031633146121cb576040516001620aed3360e41b0319815260040160405180910390fd5b601380546001600160a01b0319169055604051600081527fffa60f32d5278b35b1a3350ca92518fb5fe53e54ad07ac6355a17f54c5296b1f90602001610aaf565b816122168161256f565b6107eb8383612c19565b6122286126c8565b60408051838152602081018390527f6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c9101610857565b836001600160a01b0381163314612278576122783361256f565b61228485858585612c85565b5050505050565b6012546001600160a01b031633146122b9576040516001620aed3360e41b0319815260040160405180910390fd5b601280546001600160a01b031916905560405160009033907f222c8e95a03c7aa95fc5d110485e0d38e767f07ab1ec878a6eac644ef1d83122908390a3565b60606123038261253a565b61232057604051630a14c4b560e41b815260040160405180910390fd5b600061232a612bb8565b9050805160000361234a5760405180602001604052806000815250612375565b8061235484612cc9565b60405160200161236592919061435c565b6040516020818303038152906040525b9392505050565b6008546001600160a01b031633146123b9576012546001600160a01b031633146123b9576040516359d9793760e01b815260040160405180910390fd5b6123c283612718565b604051633f952e6560e11b81526001600160a01b0383811660048301528215156024830152841690637f2a5cca90604401610d76565b6060600c80546107109061367c565b6001600160a01b03918216600090815260076020908152604080832093909416825291909152205460ff1690565b6013546001600160a01b0316338114612461576040516353bb059b60e01b815260040160405180910390fd5b61246a81612d0d565b50601380546001600160a01b0319169055565b612485612756565b6001600160a01b0381166124ac57604051633a247dd760e11b815260040160405180910390fd5b600980546001600160a01b0319166001600160a01b0383169081179091556040519081527f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da90602001610d02565b60006001600160e01b0319821663152a902d60e11b148061252b5750632483248360e11b6001600160e01b03198316145b806106fb57506106fb82612d59565b60008160011115801561254e575060005482105b80156106fb575050600090815260046020526040902054600160e01b161590565b6daaeb6d7670e522a718067333cd4e3b1561124157604051633185c44d60e21b81523060048201526001600160a01b03821660248201526daaeb6d7670e522a718067333cd4e9063c617113490604401602060405180830381865afa1580156125dc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612600919061438b565b61124157604051633b79c77360e21b81526001600160a01b0382166004820152602401610c88565b600061263382611034565b9050336001600160a01b0382161461266c5761264f8133612407565b61266c576040516367d9dca160e11b815260040160405180910390fd5b60008281526006602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b3033146126f76126e06008546001600160a01b031690565b6001600160a01b0316336001600160a01b03161490565b176000036111c657604051635fc483c560e01b815260040160405180910390fd5b6001600160a01b03811660009081526010602052604090205460ff161515600114611241576040516315e26ff360e01b815260040160405180910390fd5b6008546001600160a01b031633146111c657604051635fc483c560e01b815260040160405180910390fd5b600061278c82612a4b565b9050836001600160a01b0316816001600160a01b0316146127bf5760405162a1148160e81b815260040160405180910390fd5b60008281526006602052604090208054338082146001600160a01b0388169091141761280c576127ef8633612407565b61280c57604051632ce44b5f60e11b815260040160405180910390fd5b6001600160a01b03851661283357604051633a954ecd60e21b815260040160405180910390fd5b801561283e57600082555b6001600160a01b038681166000908152600560205260408082208054600019019055918716808252919020805460010190554260a01b17600160e11b17600085815260046020526040812091909155600160e11b841690036128d0576001840160008181526004602052604081205490036128ce5760005481146128ce5760008181526004602052604090208490555b505b83856001600160a01b0316876001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4610bca565b6107eb8383836040518060200160405280600081525061225e565b6011548160005b828110156129955760006010600060118481548110612959576129596140a5565b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff1916911515919091179055600101612938565b5060005b818110156129fe576001601060008787858181106129b9576129b96140a5565b90506020020160208101906129ce919061348a565b6001600160a01b031681526020810191909152604001600020805460ff1916911515919091179055600101612999565b50612a0b60118585612e93565b507fbbd3b69c138de4d317d0bc4290282c4e1cbd1e58b579a5b4f114b598c237454d8484604051612a3d9291906143a8565b60405180910390a150505050565b60008180600111612aa157600054811015612aa15760008181526004602052604081205490600160e01b82169003612a9f575b80600003612375575060001901600081815260046020526040902054612a7e565b505b604051636f96cda160e11b815260040160405180910390fd5b6000805490829003612adf5760405163b562e8dd60e01b815260040160405180910390fd5b6001600160a01b03831660008181526005602090815260408083208054680100000000000000018802019055848352600490915281206001851460e11b4260a01b178317905582840190839083907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8180a4600183015b818114612b8e57808360007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef600080a4600101612b56565b5081600003612baf57604051622e076360e81b815260040160405180910390fd5b60005550505050565b6060600b80546107109061367c565b600880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b3360008181526007602090815260408083206001600160a01b03871680855290835292819020805460ff191686151590811790915590519081529192917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b612c90848484610ab9565b6001600160a01b0383163b15610ade57612cac84848484612da7565b610ade576040516368d2bf6b60e11b815260040160405180910390fd5b606060a06040510180604052602081039150506000815280825b600183039250600a81066030018353600a900480612ce35750819003601f19909101908152919050565b601280546001600160a01b0319166001600160a01b03831690811790915560405133907f222c8e95a03c7aa95fc5d110485e0d38e767f07ab1ec878a6eac644ef1d8312290600090a350565b60006301ffc9a760e01b6001600160e01b031983161480612d8a57506380ac58cd60e01b6001600160e01b03198316145b806106fb5750506001600160e01b031916635b5e139f60e01b1490565b604051630a85bd0160e11b81526000906001600160a01b0385169063150b7a0290612ddc9033908990889088906004016143f6565b6020604051808303816000875af1925050508015612e17575060408051601f3d908101601f19168201909252612e1491810190614429565b60015b612e75573d808015612e45576040519150601f19603f3d011682016040523d82523d6000602084013e612e4a565b606091505b508051600003612e6d576040516368d2bf6b60e11b815260040160405180910390fd5b805181602001fd5b6001600160e01b031916630a85bd0160e11b1490505b949350505050565b828054828255906000526020600020908101928215612ee6579160200282015b82811115612ee65781546001600160a01b0319166001600160a01b03843516178255602090920191600190910190612eb3565b50612ef2929150612ef6565b5090565b5b80821115612ef25760008155600101612ef7565b6001600160e01b03198116811461124157600080fd5b600060208284031215612f3357600080fd5b813561237581612f0b565b60005b83811015612f59578181015183820152602001612f41565b50506000910152565b60008151808452612f7a816020860160208601612f3e565b601f01601f19169290920160200192915050565b6020815260006123756020830184612f62565b600060208284031215612fb357600080fd5b5035919050565b6001600160a01b038116811461124157600080fd5b60008060408385031215612fe257600080fd5b8235612fed81612fba565b946020939093013593505050565b60008082840360e081121561300f57600080fd5b833561301a81612fba565b925060c0601f198201121561302e57600080fd5b506020830190509250929050565b60008060006060848603121561305157600080fd5b833561305c81612fba565b9250602084013561306c81612fba565b929592945050506040919091013590565b6000806040838503121561309057600080fd5b50508035926020909101359150565b600080604083850312156130b257600080fd5b82356130bd81612fba565b915060208301356001600160401b038111156130d857600080fd5b8301606081860312156130ea57600080fd5b809150509250929050565b60006040828403121561310757600080fd5b50919050565b801515811461124157600080fd5b80356131268161310d565b919050565b60008060006060848603121561314057600080fd5b833561314b81612fba565b9250602084013561315b81612fba565b9150604084013561316b8161310d565b809150509250925092565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b03811182821017156131ae576131ae613176565b60405290565b60405160c081016001600160401b03811182821017156131ae576131ae613176565b60405161010081016001600160401b03811182821017156131ae576131ae613176565b604051601f8201601f191681016001600160401b038111828210171561322157613221613176565b604052919050565b6001600160501b038116811461124157600080fd5b62ffffff8116811461124157600080fd5b64ffffffffff8116811461124157600080fd5b80356131268161324f565b61ffff8116811461124157600080fd5b80356131268161326d565b600080600083850361012081121561329f57600080fd5b84356132aa81612fba565b935060208501356132ba81612fba565b925060e0603f19820112156132ce57600080fd5b506132d761318c565b60408501356132e581613229565b815260608501356132f58161323e565b602082015260808501356133088161324f565b604082015260a085013561331b8161324f565b606082015261332c60c08601613262565b608082015261333d60e0860161327d565b60a082015261334f610100860161327d565b60c0820152809150509250925092565b60008083601f84011261337157600080fd5b5081356001600160401b0381111561338857600080fd5b6020830191508360208285010111156133a057600080fd5b9250929050565b600080602083850312156133ba57600080fd5b82356001600160401b038111156133d057600080fd5b6133dc8582860161335f565b90969095509350505050565b600080602083850312156133fb57600080fd5b82356001600160401b038082111561341257600080fd5b818501915085601f83011261342657600080fd5b81358181111561343557600080fd5b8660208260051b850101111561344a57600080fd5b60209290920196919550909350505050565b6000806040838503121561346f57600080fd5b823561347a81612fba565b915060208301356130ea81612fba565b60006020828403121561349c57600080fd5b813561237581612fba565b6000806000604084860312156134bc57600080fd5b83356134c781612fba565b925060208401356001600160401b038111156134e257600080fd5b6134ee8682870161335f565b9497909650939450505050565b600080600083850361014081121561351257600080fd5b843561351d81612fba565b9350602085013561352d81612fba565b9250610100603f198201121561354257600080fd5b506040840190509250925092565b60006020828403121561356257600080fd5b81356001600160401b0381111561357857600080fd5b8201610300818503121561237557600080fd5b6000806040838503121561359e57600080fd5b82356135a981612fba565b915060208301356130ea8161310d565b600080600080608085870312156135cf57600080fd5b84356135da81612fba565b93506020858101356135eb81612fba565b93506040860135925060608601356001600160401b038082111561360e57600080fd5b818801915088601f83011261362257600080fd5b81358181111561363457613634613176565b613646601f8201601f191685016131f9565b9150808252898482850101111561365c57600080fd5b808484018584013760008482840101525080935050505092959194509250565b600181811c9082168061369057607f821691505b60208210810361310757634e487b7160e01b600052602260045260246000fd5b65ffffffffffff8116811461124157600080fd5b8051613126816136b0565b80516131268161326d565b80516131268161310d565b600060c082840312156136f757600080fd5b6136ff6131b4565b825161370a81613229565b8152602083015161371a816136b0565b6020820152604083015161372d816136b0565b604082015260608301516137408161326d565b606082015260808301516137538161326d565b608082015260a08301516137668161310d565b60a08201529392505050565b8035613126816136b0565b600060c0828403121561378f57600080fd5b6137976131b4565b82356137a281613229565b815260208301356137b2816136b0565b602082015260408301356137c5816136b0565b604082015260608301356137d88161326d565b606082015260808301356137eb8161326d565b608082015260a08301356137668161310d565b634e487b7160e01b600052601160045260246000fd5b80820281158282048414176106fb576106fb6137fe565b60008261384857634e487b7160e01b600052601260045260246000fd5b500490565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000808335601e1984360301811261388d57600080fd5b83016020810192503590506001600160401b038111156138ac57600080fd5b8036038213156133a057600080fd5b60006060830182358452602080840135601e198536030181126138dd57600080fd5b840181810190356001600160401b038111156138f857600080fd5b8060051b80360383131561390b57600080fd5b6060848901529381905260809387018401938290880160005b8381101561395e57898703607f1901825261393f8386613876565b61394a89828461384d565b985050509185019190850190600101613924565b5050505050506139716040840184613876565b858303604087015261398483828461384d565b9695505050505050565b60208152600061237560208301846138bb565b6001600160601b038116811461124157600080fd5b6000602082840312156139c857600080fd5b8135612375816139a1565b81356139de81612fba565b81546001600160a01b03199081166001600160a01b039290921691821783556020840135613a0b816139a1565b60a01b1617905550565b80516131268161324f565b600060e08284031215613a3257600080fd5b60405160e081018181106001600160401b0382111715613a5457613a54613176565b6040528251613a6281613229565b81526020830151613a728161323e565b60208201526040830151613a858161324f565b60408201526060830151613a988161324f565b6060820152613aa960808401613a15565b6080820152613aba60a084016136cf565b60a0820152613acb60c084016136cf565b60c08201529392505050565b6001600160501b03815116825262ffffff6020820151166020830152604081015164ffffffffff8082166040850152806060840151166060850152806080840151166080850152505060a081015161ffff80821660a08501528060c08401511660c085015250505050565b6001600160a01b038316815261010081016123756020830184613ad7565b601f8211156107eb57600081815260208120601f850160051c81016020861015613b875750805b601f850160051c820191505b81811015610bca57828155600101613b93565b6001600160401b03831115613bbd57613bbd613176565b613bd183613bcb835461367c565b83613b60565b6000601f841160018114613c055760008515613bed5750838201355b600019600387901b1c1916600186901b178355612284565b600083815260209020601f19861690835b82811015613c365786850135825560209485019460019092019101613c16565b5086821015613c535760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b818103818111156106fb576106fb6137fe565b808201808211156106fb576106fb6137fe565b602081526000612e8b60208301848661384d565b60ff8116811461124157600080fd5b805161312681613c9f565b63ffffffff8116811461124157600080fd5b805161312681613cb9565b60006101008284031215613ce957600080fd5b613cf16131d6565b8251613cfc81613229565b81526020830151613d0c8161326d565b6020820152613d1d604084016136c4565b6040820152613d2e606084016136c4565b6060820152613d3f60808401613cae565b6080820152613d5060a08401613ccb565b60a0820152613d6160c084016136cf565b60c0820152613d7260e084016136da565b60e08201529392505050565b803561312681613c9f565b803561312681613cb9565b60006101008284031215613da757600080fd5b613daf6131d6565b8235613dba81613229565b81526020830135613dca8161326d565b6020820152613ddb60408401613772565b6040820152613dec60608401613772565b6060820152613dfd60808401613d7e565b6080820152613e0e60a08401613d89565b60a0820152613e1f60c0840161327d565b60c0820152613d7260e0840161311b565b6001600160501b03815116825261ffff6020820151166020830152604081015165ffffffffffff8082166040850152806060840151166060850152505060ff608082015116608083015260a0810151613e9160a084018263ffffffff169052565b5060c0810151613ea760c084018261ffff169052565b5060e08101516107eb60e084018215159052565b6001600160a01b038316815261012081016123756020830184613e30565b6000808335601e19843603018112613ef057600080fd5b8301803591506001600160401b03821115613f0a57600080fd5b6020019150368190038213156133a057600080fd5b600060208284031215613f3157600080fd5b8135612375816136b0565b6001600160a01b038316815260e081018235613f5781613229565b6001600160501b0381166020840152506020830135613f75816136b0565b65ffffffffffff808216604085015260408501359150613f94826136b0565b16606083810191909152830135613faa8161326d565b61ffff81166080840152506080830135613fc38161326d565b61ffff811660a084015250613fda60a0840161311b565b80151560c0840152509392505050565b6001600160a01b038416815260406020820181905260009061400f908301848661384d565b95945050505050565b60008235605e1983360301811261402e57600080fd5b9190910192915050565b6001600160a01b0383168152604060208201819052600090612e8b908301846138bb565b6000808335601e1984360301811261407357600080fd5b8301803591506001600160401b0382111561408d57600080fd5b6020019150600581901b36038213156133a057600080fd5b634e487b7160e01b600052603260045260246000fd5b6001600160a01b039384168152919092166020820152901515604082015260600190565b6000808335601e198436030181126140f657600080fd5b8301803591506001600160401b0382111561411057600080fd5b6020019150600881901b36038213156133a057600080fd5b6001600160a01b038481168252831660208201526101408101823561414c81613229565b6001600160501b0316604083015260208301356141688161326d565b61ffff16606083015261417d60408401613772565b65ffffffffffff16608083015261419660608401613772565b65ffffffffffff1660a08301526141af60808401613d7e565b60ff1660c08301526141c360a08401613d89565b63ffffffff1660e08301526141da60c0840161327d565b61ffff166101008301526141f060e0840161311b565b8015156101208401525b50949350505050565b6001600160a01b038481168252831660208201526101408101612e8b6040830184613e30565b6000808335601e1984360301811261424057600080fd5b8301803591506001600160401b0382111561425a57600080fd5b602001915060e0810236038213156133a057600080fd5b6001600160a01b038481168252831660208201526101208101823561429581613229565b6001600160501b0316604083015260208301356142b18161323e565b62ffffff16606083015260408301356142c98161324f565b64ffffffffff1660808301526142e160608401613262565b64ffffffffff1660a08301526142f960808401613262565b64ffffffffff1660c083015261431160a0840161327d565b61ffff1660e083015261432660c0840161327d565b61ffff81166101008401526141fa565b6001600160a01b038481168252831660208201526101208101612e8b6040830184613ad7565b6000835161436e818460208801612f3e565b835190830190614382818360208801612f3e565b01949350505050565b60006020828403121561439d57600080fd5b81516123758161310d565b60208082528181018390526000908460408401835b868110156143eb5782356143d081612fba565b6001600160a01b0316825291830191908301906001016143bd565b509695505050505050565b6001600160a01b038581168252841660208201526040810183905260806060820181905260009061398490830184612f62565b60006020828403121561443b57600080fd5b815161237581612f0b56fea2646970667358221220e993d2818d799b6999b27493829e07dd3199ac69bbf1014c6af0740f7c20ed8d64736f6c63430008110033",
}

// PartnerABI is the input ABI used to generate the binding from.
// Deprecated: Use PartnerMetaData.ABI instead.
var PartnerABI = PartnerMetaData.ABI

// PartnerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PartnerMetaData.Bin instead.
var PartnerBin = PartnerMetaData.Bin

// DeployPartner deploys a new Ethereum contract, binding an instance of Partner to it.
func DeployPartner(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, administrator common.Address, allowedSeaDrop []common.Address) (common.Address, *types.Transaction, *Partner, error) {
	parsed, err := PartnerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PartnerBin), backend, name, symbol, administrator, allowedSeaDrop)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Partner{PartnerCaller: PartnerCaller{contract: contract}, PartnerTransactor: PartnerTransactor{contract: contract}, PartnerFilterer: PartnerFilterer{contract: contract}}, nil
}

// Partner is an auto generated Go binding around an Ethereum contract.
type Partner struct {
	PartnerCaller     // Read-only binding to the contract
	PartnerTransactor // Write-only binding to the contract
	PartnerFilterer   // Log filterer for contract events
}

// PartnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PartnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PartnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PartnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PartnerSession struct {
	Contract     *Partner          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PartnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PartnerCallerSession struct {
	Contract *PartnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PartnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PartnerTransactorSession struct {
	Contract     *PartnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PartnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PartnerRaw struct {
	Contract *Partner // Generic contract binding to access the raw methods on
}

// PartnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PartnerCallerRaw struct {
	Contract *PartnerCaller // Generic read-only contract binding to access the raw methods on
}

// PartnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PartnerTransactorRaw struct {
	Contract *PartnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPartner creates a new instance of Partner, bound to a specific deployed contract.
func NewPartner(address common.Address, backend bind.ContractBackend) (*Partner, error) {
	contract, err := bindPartner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Partner{PartnerCaller: PartnerCaller{contract: contract}, PartnerTransactor: PartnerTransactor{contract: contract}, PartnerFilterer: PartnerFilterer{contract: contract}}, nil
}

// NewPartnerCaller creates a new read-only instance of Partner, bound to a specific deployed contract.
func NewPartnerCaller(address common.Address, caller bind.ContractCaller) (*PartnerCaller, error) {
	contract, err := bindPartner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PartnerCaller{contract: contract}, nil
}

// NewPartnerTransactor creates a new write-only instance of Partner, bound to a specific deployed contract.
func NewPartnerTransactor(address common.Address, transactor bind.ContractTransactor) (*PartnerTransactor, error) {
	contract, err := bindPartner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PartnerTransactor{contract: contract}, nil
}

// NewPartnerFilterer creates a new log filterer instance of Partner, bound to a specific deployed contract.
func NewPartnerFilterer(address common.Address, filterer bind.ContractFilterer) (*PartnerFilterer, error) {
	contract, err := bindPartner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PartnerFilterer{contract: contract}, nil
}

// bindPartner binds a generic wrapper to an already deployed contract.
func bindPartner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PartnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Partner *PartnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Partner.Contract.PartnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Partner *PartnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Partner.Contract.PartnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Partner *PartnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Partner.Contract.PartnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Partner *PartnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Partner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Partner *PartnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Partner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Partner *PartnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Partner.Contract.contract.Transact(opts, method, params...)
}

// OPERATORFILTERREGISTRY is a free data retrieval call binding the contract method 0x41f43434.
//
// Solidity: function OPERATOR_FILTER_REGISTRY() view returns(address)
func (_Partner *PartnerCaller) OPERATORFILTERREGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "OPERATOR_FILTER_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OPERATORFILTERREGISTRY is a free data retrieval call binding the contract method 0x41f43434.
//
// Solidity: function OPERATOR_FILTER_REGISTRY() view returns(address)
func (_Partner *PartnerSession) OPERATORFILTERREGISTRY() (common.Address, error) {
	return _Partner.Contract.OPERATORFILTERREGISTRY(&_Partner.CallOpts)
}

// OPERATORFILTERREGISTRY is a free data retrieval call binding the contract method 0x41f43434.
//
// Solidity: function OPERATOR_FILTER_REGISTRY() view returns(address)
func (_Partner *PartnerCallerSession) OPERATORFILTERREGISTRY() (common.Address, error) {
	return _Partner.Contract.OPERATORFILTERREGISTRY(&_Partner.CallOpts)
}

// Administrator is a free data retrieval call binding the contract method 0xf53d0a8e.
//
// Solidity: function administrator() view returns(address)
func (_Partner *PartnerCaller) Administrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "administrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Administrator is a free data retrieval call binding the contract method 0xf53d0a8e.
//
// Solidity: function administrator() view returns(address)
func (_Partner *PartnerSession) Administrator() (common.Address, error) {
	return _Partner.Contract.Administrator(&_Partner.CallOpts)
}

// Administrator is a free data retrieval call binding the contract method 0xf53d0a8e.
//
// Solidity: function administrator() view returns(address)
func (_Partner *PartnerCallerSession) Administrator() (common.Address, error) {
	return _Partner.Contract.Administrator(&_Partner.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Partner *PartnerCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Partner *PartnerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Partner.Contract.BalanceOf(&_Partner.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Partner *PartnerCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Partner.Contract.BalanceOf(&_Partner.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Partner *PartnerCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Partner *PartnerSession) BaseURI() (string, error) {
	return _Partner.Contract.BaseURI(&_Partner.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Partner *PartnerCallerSession) BaseURI() (string, error) {
	return _Partner.Contract.BaseURI(&_Partner.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Partner *PartnerCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Partner *PartnerSession) ContractURI() (string, error) {
	return _Partner.Contract.ContractURI(&_Partner.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Partner *PartnerCallerSession) ContractURI() (string, error) {
	return _Partner.Contract.ContractURI(&_Partner.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Partner *PartnerCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Partner *PartnerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Partner.Contract.GetApproved(&_Partner.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Partner *PartnerCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Partner.Contract.GetApproved(&_Partner.CallOpts, tokenId)
}

// GetMintStats is a free data retrieval call binding the contract method 0x840e15d4.
//
// Solidity: function getMintStats(address minter) view returns(uint256 minterNumMinted, uint256 currentTotalSupply, uint256 maxSupply)
func (_Partner *PartnerCaller) GetMintStats(opts *bind.CallOpts, minter common.Address) (struct {
	MinterNumMinted    *big.Int
	CurrentTotalSupply *big.Int
	MaxSupply          *big.Int
}, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "getMintStats", minter)

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
func (_Partner *PartnerSession) GetMintStats(minter common.Address) (struct {
	MinterNumMinted    *big.Int
	CurrentTotalSupply *big.Int
	MaxSupply          *big.Int
}, error) {
	return _Partner.Contract.GetMintStats(&_Partner.CallOpts, minter)
}

// GetMintStats is a free data retrieval call binding the contract method 0x840e15d4.
//
// Solidity: function getMintStats(address minter) view returns(uint256 minterNumMinted, uint256 currentTotalSupply, uint256 maxSupply)
func (_Partner *PartnerCallerSession) GetMintStats(minter common.Address) (struct {
	MinterNumMinted    *big.Int
	CurrentTotalSupply *big.Int
	MaxSupply          *big.Int
}, error) {
	return _Partner.Contract.GetMintStats(&_Partner.CallOpts, minter)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Partner *PartnerCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Partner *PartnerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Partner.Contract.IsApprovedForAll(&_Partner.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Partner *PartnerCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Partner.Contract.IsApprovedForAll(&_Partner.CallOpts, owner, operator)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Partner *PartnerCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Partner *PartnerSession) MaxSupply() (*big.Int, error) {
	return _Partner.Contract.MaxSupply(&_Partner.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Partner *PartnerCallerSession) MaxSupply() (*big.Int, error) {
	return _Partner.Contract.MaxSupply(&_Partner.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Partner *PartnerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Partner *PartnerSession) Name() (string, error) {
	return _Partner.Contract.Name(&_Partner.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Partner *PartnerCallerSession) Name() (string, error) {
	return _Partner.Contract.Name(&_Partner.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Partner *PartnerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Partner *PartnerSession) Owner() (common.Address, error) {
	return _Partner.Contract.Owner(&_Partner.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Partner *PartnerCallerSession) Owner() (common.Address, error) {
	return _Partner.Contract.Owner(&_Partner.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Partner *PartnerCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Partner *PartnerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Partner.Contract.OwnerOf(&_Partner.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Partner *PartnerCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Partner.Contract.OwnerOf(&_Partner.CallOpts, tokenId)
}

// PotentialAdministrator is a free data retrieval call binding the contract method 0xd5e7feb8.
//
// Solidity: function potentialAdministrator() view returns(address)
func (_Partner *PartnerCaller) PotentialAdministrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "potentialAdministrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PotentialAdministrator is a free data retrieval call binding the contract method 0xd5e7feb8.
//
// Solidity: function potentialAdministrator() view returns(address)
func (_Partner *PartnerSession) PotentialAdministrator() (common.Address, error) {
	return _Partner.Contract.PotentialAdministrator(&_Partner.CallOpts)
}

// PotentialAdministrator is a free data retrieval call binding the contract method 0xd5e7feb8.
//
// Solidity: function potentialAdministrator() view returns(address)
func (_Partner *PartnerCallerSession) PotentialAdministrator() (common.Address, error) {
	return _Partner.Contract.PotentialAdministrator(&_Partner.CallOpts)
}

// ProvenanceHash is a free data retrieval call binding the contract method 0xc6ab67a3.
//
// Solidity: function provenanceHash() view returns(bytes32)
func (_Partner *PartnerCaller) ProvenanceHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "provenanceHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvenanceHash is a free data retrieval call binding the contract method 0xc6ab67a3.
//
// Solidity: function provenanceHash() view returns(bytes32)
func (_Partner *PartnerSession) ProvenanceHash() ([32]byte, error) {
	return _Partner.Contract.ProvenanceHash(&_Partner.CallOpts)
}

// ProvenanceHash is a free data retrieval call binding the contract method 0xc6ab67a3.
//
// Solidity: function provenanceHash() view returns(bytes32)
func (_Partner *PartnerCallerSession) ProvenanceHash() ([32]byte, error) {
	return _Partner.Contract.ProvenanceHash(&_Partner.CallOpts)
}

// RoyaltyAddress is a free data retrieval call binding the contract method 0xad2f852a.
//
// Solidity: function royaltyAddress() view returns(address)
func (_Partner *PartnerCaller) RoyaltyAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "royaltyAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyAddress is a free data retrieval call binding the contract method 0xad2f852a.
//
// Solidity: function royaltyAddress() view returns(address)
func (_Partner *PartnerSession) RoyaltyAddress() (common.Address, error) {
	return _Partner.Contract.RoyaltyAddress(&_Partner.CallOpts)
}

// RoyaltyAddress is a free data retrieval call binding the contract method 0xad2f852a.
//
// Solidity: function royaltyAddress() view returns(address)
func (_Partner *PartnerCallerSession) RoyaltyAddress() (common.Address, error) {
	return _Partner.Contract.RoyaltyAddress(&_Partner.CallOpts)
}

// RoyaltyBasisPoints is a free data retrieval call binding the contract method 0x42260b5d.
//
// Solidity: function royaltyBasisPoints() view returns(uint256)
func (_Partner *PartnerCaller) RoyaltyBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "royaltyBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoyaltyBasisPoints is a free data retrieval call binding the contract method 0x42260b5d.
//
// Solidity: function royaltyBasisPoints() view returns(uint256)
func (_Partner *PartnerSession) RoyaltyBasisPoints() (*big.Int, error) {
	return _Partner.Contract.RoyaltyBasisPoints(&_Partner.CallOpts)
}

// RoyaltyBasisPoints is a free data retrieval call binding the contract method 0x42260b5d.
//
// Solidity: function royaltyBasisPoints() view returns(uint256)
func (_Partner *PartnerCallerSession) RoyaltyBasisPoints() (*big.Int, error) {
	return _Partner.Contract.RoyaltyBasisPoints(&_Partner.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Partner *PartnerCaller) RoyaltyInfo(opts *bind.CallOpts, arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "royaltyInfo", arg0, _salePrice)

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
func (_Partner *PartnerSession) RoyaltyInfo(arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Partner.Contract.RoyaltyInfo(&_Partner.CallOpts, arg0, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Partner *PartnerCallerSession) RoyaltyInfo(arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Partner.Contract.RoyaltyInfo(&_Partner.CallOpts, arg0, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Partner *PartnerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Partner *PartnerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Partner.Contract.SupportsInterface(&_Partner.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Partner *PartnerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Partner.Contract.SupportsInterface(&_Partner.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Partner *PartnerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Partner *PartnerSession) Symbol() (string, error) {
	return _Partner.Contract.Symbol(&_Partner.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Partner *PartnerCallerSession) Symbol() (string, error) {
	return _Partner.Contract.Symbol(&_Partner.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Partner *PartnerCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Partner *PartnerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Partner.Contract.TokenURI(&_Partner.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Partner *PartnerCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Partner.Contract.TokenURI(&_Partner.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Partner *PartnerCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Partner *PartnerSession) TotalSupply() (*big.Int, error) {
	return _Partner.Contract.TotalSupply(&_Partner.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Partner *PartnerCallerSession) TotalSupply() (*big.Int, error) {
	return _Partner.Contract.TotalSupply(&_Partner.CallOpts)
}

// AcceptAdministration is a paid mutator transaction binding the contract method 0xf0025d96.
//
// Solidity: function acceptAdministration() returns()
func (_Partner *PartnerTransactor) AcceptAdministration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "acceptAdministration")
}

// AcceptAdministration is a paid mutator transaction binding the contract method 0xf0025d96.
//
// Solidity: function acceptAdministration() returns()
func (_Partner *PartnerSession) AcceptAdministration() (*types.Transaction, error) {
	return _Partner.Contract.AcceptAdministration(&_Partner.TransactOpts)
}

// AcceptAdministration is a paid mutator transaction binding the contract method 0xf0025d96.
//
// Solidity: function acceptAdministration() returns()
func (_Partner *PartnerTransactorSession) AcceptAdministration() (*types.Transaction, error) {
	return _Partner.Contract.AcceptAdministration(&_Partner.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Partner *PartnerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Partner *PartnerSession) AcceptOwnership() (*types.Transaction, error) {
	return _Partner.Contract.AcceptOwnership(&_Partner.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Partner *PartnerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Partner.Contract.AcceptOwnership(&_Partner.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Partner *PartnerTransactor) Approve(opts *bind.TransactOpts, operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "approve", operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Partner *PartnerSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Partner.Contract.Approve(&_Partner.TransactOpts, operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Partner *PartnerTransactorSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Partner.Contract.Approve(&_Partner.TransactOpts, operator, tokenId)
}

// CancelAdministrationTransfer is a paid mutator transaction binding the contract method 0x9794ed40.
//
// Solidity: function cancelAdministrationTransfer() returns()
func (_Partner *PartnerTransactor) CancelAdministrationTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "cancelAdministrationTransfer")
}

// CancelAdministrationTransfer is a paid mutator transaction binding the contract method 0x9794ed40.
//
// Solidity: function cancelAdministrationTransfer() returns()
func (_Partner *PartnerSession) CancelAdministrationTransfer() (*types.Transaction, error) {
	return _Partner.Contract.CancelAdministrationTransfer(&_Partner.TransactOpts)
}

// CancelAdministrationTransfer is a paid mutator transaction binding the contract method 0x9794ed40.
//
// Solidity: function cancelAdministrationTransfer() returns()
func (_Partner *PartnerTransactorSession) CancelAdministrationTransfer() (*types.Transaction, error) {
	return _Partner.Contract.CancelAdministrationTransfer(&_Partner.TransactOpts)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_Partner *PartnerTransactor) CancelOwnershipTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "cancelOwnershipTransfer")
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_Partner *PartnerSession) CancelOwnershipTransfer() (*types.Transaction, error) {
	return _Partner.Contract.CancelOwnershipTransfer(&_Partner.TransactOpts)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_Partner *PartnerTransactorSession) CancelOwnershipTransfer() (*types.Transaction, error) {
	return _Partner.Contract.CancelOwnershipTransfer(&_Partner.TransactOpts)
}

// EmitBatchMetadataUpdate is a paid mutator transaction binding the contract method 0xa4830114.
//
// Solidity: function emitBatchMetadataUpdate(uint256 fromTokenId, uint256 toTokenId) returns()
func (_Partner *PartnerTransactor) EmitBatchMetadataUpdate(opts *bind.TransactOpts, fromTokenId *big.Int, toTokenId *big.Int) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "emitBatchMetadataUpdate", fromTokenId, toTokenId)
}

// EmitBatchMetadataUpdate is a paid mutator transaction binding the contract method 0xa4830114.
//
// Solidity: function emitBatchMetadataUpdate(uint256 fromTokenId, uint256 toTokenId) returns()
func (_Partner *PartnerSession) EmitBatchMetadataUpdate(fromTokenId *big.Int, toTokenId *big.Int) (*types.Transaction, error) {
	return _Partner.Contract.EmitBatchMetadataUpdate(&_Partner.TransactOpts, fromTokenId, toTokenId)
}

// EmitBatchMetadataUpdate is a paid mutator transaction binding the contract method 0xa4830114.
//
// Solidity: function emitBatchMetadataUpdate(uint256 fromTokenId, uint256 toTokenId) returns()
func (_Partner *PartnerTransactorSession) EmitBatchMetadataUpdate(fromTokenId *big.Int, toTokenId *big.Int) (*types.Transaction, error) {
	return _Partner.Contract.EmitBatchMetadataUpdate(&_Partner.TransactOpts, fromTokenId, toTokenId)
}

// MintSeaDrop is a paid mutator transaction binding the contract method 0x64869dad.
//
// Solidity: function mintSeaDrop(address minter, uint256 quantity) returns()
func (_Partner *PartnerTransactor) MintSeaDrop(opts *bind.TransactOpts, minter common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "mintSeaDrop", minter, quantity)
}

// MintSeaDrop is a paid mutator transaction binding the contract method 0x64869dad.
//
// Solidity: function mintSeaDrop(address minter, uint256 quantity) returns()
func (_Partner *PartnerSession) MintSeaDrop(minter common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Partner.Contract.MintSeaDrop(&_Partner.TransactOpts, minter, quantity)
}

// MintSeaDrop is a paid mutator transaction binding the contract method 0x64869dad.
//
// Solidity: function mintSeaDrop(address minter, uint256 quantity) returns()
func (_Partner *PartnerTransactorSession) MintSeaDrop(minter common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Partner.Contract.MintSeaDrop(&_Partner.TransactOpts, minter, quantity)
}

// MultiConfigure is a paid mutator transaction binding the contract method 0x911f456b.
//
// Solidity: function multiConfigure((uint256,string,string,address,(uint80,uint48,uint48,uint16,uint16,bool),string,(bytes32,string[],string),address,bytes32,address[],address[],address[],address[],address[],(uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool)[],address[],address[],(uint80,uint24,uint40,uint40,uint40,uint16,uint16)[],address[]) config) returns()
func (_Partner *PartnerTransactor) MultiConfigure(opts *bind.TransactOpts, config ERC721SeaDropStructsErrorsAndEventsMultiConfigureStruct) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "multiConfigure", config)
}

// MultiConfigure is a paid mutator transaction binding the contract method 0x911f456b.
//
// Solidity: function multiConfigure((uint256,string,string,address,(uint80,uint48,uint48,uint16,uint16,bool),string,(bytes32,string[],string),address,bytes32,address[],address[],address[],address[],address[],(uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool)[],address[],address[],(uint80,uint24,uint40,uint40,uint40,uint16,uint16)[],address[]) config) returns()
func (_Partner *PartnerSession) MultiConfigure(config ERC721SeaDropStructsErrorsAndEventsMultiConfigureStruct) (*types.Transaction, error) {
	return _Partner.Contract.MultiConfigure(&_Partner.TransactOpts, config)
}

// MultiConfigure is a paid mutator transaction binding the contract method 0x911f456b.
//
// Solidity: function multiConfigure((uint256,string,string,address,(uint80,uint48,uint48,uint16,uint16,bool),string,(bytes32,string[],string),address,bytes32,address[],address[],address[],address[],address[],(uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool)[],address[],address[],(uint80,uint24,uint40,uint40,uint40,uint16,uint16)[],address[]) config) returns()
func (_Partner *PartnerTransactorSession) MultiConfigure(config ERC721SeaDropStructsErrorsAndEventsMultiConfigureStruct) (*types.Transaction, error) {
	return _Partner.Contract.MultiConfigure(&_Partner.TransactOpts, config)
}

// RenounceAdministration is a paid mutator transaction binding the contract method 0xc780b63d.
//
// Solidity: function renounceAdministration() returns()
func (_Partner *PartnerTransactor) RenounceAdministration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "renounceAdministration")
}

// RenounceAdministration is a paid mutator transaction binding the contract method 0xc780b63d.
//
// Solidity: function renounceAdministration() returns()
func (_Partner *PartnerSession) RenounceAdministration() (*types.Transaction, error) {
	return _Partner.Contract.RenounceAdministration(&_Partner.TransactOpts)
}

// RenounceAdministration is a paid mutator transaction binding the contract method 0xc780b63d.
//
// Solidity: function renounceAdministration() returns()
func (_Partner *PartnerTransactorSession) RenounceAdministration() (*types.Transaction, error) {
	return _Partner.Contract.RenounceAdministration(&_Partner.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Partner *PartnerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Partner *PartnerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Partner.Contract.RenounceOwnership(&_Partner.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Partner *PartnerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Partner.Contract.RenounceOwnership(&_Partner.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Partner *PartnerTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Partner *PartnerSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Partner.Contract.SafeTransferFrom(&_Partner.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Partner *PartnerTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Partner.Contract.SafeTransferFrom(&_Partner.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Partner *PartnerTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Partner *PartnerSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Partner.Contract.SafeTransferFrom0(&_Partner.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Partner *PartnerTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Partner.Contract.SafeTransferFrom0(&_Partner.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Partner *PartnerTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Partner *PartnerSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Partner.Contract.SetApprovalForAll(&_Partner.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Partner *PartnerTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Partner.Contract.SetApprovalForAll(&_Partner.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Partner *PartnerTransactor) SetBaseURI(opts *bind.TransactOpts, newBaseURI string) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "setBaseURI", newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Partner *PartnerSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Partner.Contract.SetBaseURI(&_Partner.TransactOpts, newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Partner *PartnerTransactorSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Partner.Contract.SetBaseURI(&_Partner.TransactOpts, newBaseURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_Partner *PartnerTransactor) SetContractURI(opts *bind.TransactOpts, newContractURI string) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "setContractURI", newContractURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_Partner *PartnerSession) SetContractURI(newContractURI string) (*types.Transaction, error) {
	return _Partner.Contract.SetContractURI(&_Partner.TransactOpts, newContractURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newContractURI) returns()
func (_Partner *PartnerTransactorSession) SetContractURI(newContractURI string) (*types.Transaction, error) {
	return _Partner.Contract.SetContractURI(&_Partner.TransactOpts, newContractURI)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_Partner *PartnerTransactor) SetMaxSupply(opts *bind.TransactOpts, newMaxSupply *big.Int) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "setMaxSupply", newMaxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_Partner *PartnerSession) SetMaxSupply(newMaxSupply *big.Int) (*types.Transaction, error) {
	return _Partner.Contract.SetMaxSupply(&_Partner.TransactOpts, newMaxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_Partner *PartnerTransactorSession) SetMaxSupply(newMaxSupply *big.Int) (*types.Transaction, error) {
	return _Partner.Contract.SetMaxSupply(&_Partner.TransactOpts, newMaxSupply)
}

// SetProvenanceHash is a paid mutator transaction binding the contract method 0x099b6bfa.
//
// Solidity: function setProvenanceHash(bytes32 newProvenanceHash) returns()
func (_Partner *PartnerTransactor) SetProvenanceHash(opts *bind.TransactOpts, newProvenanceHash [32]byte) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "setProvenanceHash", newProvenanceHash)
}

// SetProvenanceHash is a paid mutator transaction binding the contract method 0x099b6bfa.
//
// Solidity: function setProvenanceHash(bytes32 newProvenanceHash) returns()
func (_Partner *PartnerSession) SetProvenanceHash(newProvenanceHash [32]byte) (*types.Transaction, error) {
	return _Partner.Contract.SetProvenanceHash(&_Partner.TransactOpts, newProvenanceHash)
}

// SetProvenanceHash is a paid mutator transaction binding the contract method 0x099b6bfa.
//
// Solidity: function setProvenanceHash(bytes32 newProvenanceHash) returns()
func (_Partner *PartnerTransactorSession) SetProvenanceHash(newProvenanceHash [32]byte) (*types.Transaction, error) {
	return _Partner.Contract.SetProvenanceHash(&_Partner.TransactOpts, newProvenanceHash)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_Partner *PartnerTransactor) SetRoyaltyInfo(opts *bind.TransactOpts, newInfo ISeaDropTokenContractMetadataRoyaltyInfo) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "setRoyaltyInfo", newInfo)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_Partner *PartnerSession) SetRoyaltyInfo(newInfo ISeaDropTokenContractMetadataRoyaltyInfo) (*types.Transaction, error) {
	return _Partner.Contract.SetRoyaltyInfo(&_Partner.TransactOpts, newInfo)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_Partner *PartnerTransactorSession) SetRoyaltyInfo(newInfo ISeaDropTokenContractMetadataRoyaltyInfo) (*types.Transaction, error) {
	return _Partner.Contract.SetRoyaltyInfo(&_Partner.TransactOpts, newInfo)
}

// TransferAdministration is a paid mutator transaction binding the contract method 0x913ee93d.
//
// Solidity: function transferAdministration(address newAdministrator) returns()
func (_Partner *PartnerTransactor) TransferAdministration(opts *bind.TransactOpts, newAdministrator common.Address) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "transferAdministration", newAdministrator)
}

// TransferAdministration is a paid mutator transaction binding the contract method 0x913ee93d.
//
// Solidity: function transferAdministration(address newAdministrator) returns()
func (_Partner *PartnerSession) TransferAdministration(newAdministrator common.Address) (*types.Transaction, error) {
	return _Partner.Contract.TransferAdministration(&_Partner.TransactOpts, newAdministrator)
}

// TransferAdministration is a paid mutator transaction binding the contract method 0x913ee93d.
//
// Solidity: function transferAdministration(address newAdministrator) returns()
func (_Partner *PartnerTransactorSession) TransferAdministration(newAdministrator common.Address) (*types.Transaction, error) {
	return _Partner.Contract.TransferAdministration(&_Partner.TransactOpts, newAdministrator)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Partner *PartnerTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Partner *PartnerSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Partner.Contract.TransferFrom(&_Partner.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Partner *PartnerTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Partner.Contract.TransferFrom(&_Partner.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newPotentialOwner) returns()
func (_Partner *PartnerTransactor) TransferOwnership(opts *bind.TransactOpts, newPotentialOwner common.Address) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "transferOwnership", newPotentialOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newPotentialOwner) returns()
func (_Partner *PartnerSession) TransferOwnership(newPotentialOwner common.Address) (*types.Transaction, error) {
	return _Partner.Contract.TransferOwnership(&_Partner.TransactOpts, newPotentialOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newPotentialOwner) returns()
func (_Partner *PartnerTransactorSession) TransferOwnership(newPotentialOwner common.Address) (*types.Transaction, error) {
	return _Partner.Contract.TransferOwnership(&_Partner.TransactOpts, newPotentialOwner)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0x3680620d.
//
// Solidity: function updateAllowList(address seaDropImpl, (bytes32,string[],string) allowListData) returns()
func (_Partner *PartnerTransactor) UpdateAllowList(opts *bind.TransactOpts, seaDropImpl common.Address, allowListData AllowListData) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "updateAllowList", seaDropImpl, allowListData)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0x3680620d.
//
// Solidity: function updateAllowList(address seaDropImpl, (bytes32,string[],string) allowListData) returns()
func (_Partner *PartnerSession) UpdateAllowList(seaDropImpl common.Address, allowListData AllowListData) (*types.Transaction, error) {
	return _Partner.Contract.UpdateAllowList(&_Partner.TransactOpts, seaDropImpl, allowListData)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0x3680620d.
//
// Solidity: function updateAllowList(address seaDropImpl, (bytes32,string[],string) allowListData) returns()
func (_Partner *PartnerTransactorSession) UpdateAllowList(seaDropImpl common.Address, allowListData AllowListData) (*types.Transaction, error) {
	return _Partner.Contract.UpdateAllowList(&_Partner.TransactOpts, seaDropImpl, allowListData)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x48a4c101.
//
// Solidity: function updateAllowedFeeRecipient(address seaDropImpl, address feeRecipient, bool allowed) returns()
func (_Partner *PartnerTransactor) UpdateAllowedFeeRecipient(opts *bind.TransactOpts, seaDropImpl common.Address, feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "updateAllowedFeeRecipient", seaDropImpl, feeRecipient, allowed)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x48a4c101.
//
// Solidity: function updateAllowedFeeRecipient(address seaDropImpl, address feeRecipient, bool allowed) returns()
func (_Partner *PartnerSession) UpdateAllowedFeeRecipient(seaDropImpl common.Address, feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _Partner.Contract.UpdateAllowedFeeRecipient(&_Partner.TransactOpts, seaDropImpl, feeRecipient, allowed)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x48a4c101.
//
// Solidity: function updateAllowedFeeRecipient(address seaDropImpl, address feeRecipient, bool allowed) returns()
func (_Partner *PartnerTransactorSession) UpdateAllowedFeeRecipient(seaDropImpl common.Address, feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _Partner.Contract.UpdateAllowedFeeRecipient(&_Partner.TransactOpts, seaDropImpl, feeRecipient, allowed)
}

// UpdateAllowedSeaDrop is a paid mutator transaction binding the contract method 0x60c308b6.
//
// Solidity: function updateAllowedSeaDrop(address[] allowedSeaDrop) returns()
func (_Partner *PartnerTransactor) UpdateAllowedSeaDrop(opts *bind.TransactOpts, allowedSeaDrop []common.Address) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "updateAllowedSeaDrop", allowedSeaDrop)
}

// UpdateAllowedSeaDrop is a paid mutator transaction binding the contract method 0x60c308b6.
//
// Solidity: function updateAllowedSeaDrop(address[] allowedSeaDrop) returns()
func (_Partner *PartnerSession) UpdateAllowedSeaDrop(allowedSeaDrop []common.Address) (*types.Transaction, error) {
	return _Partner.Contract.UpdateAllowedSeaDrop(&_Partner.TransactOpts, allowedSeaDrop)
}

// UpdateAllowedSeaDrop is a paid mutator transaction binding the contract method 0x60c308b6.
//
// Solidity: function updateAllowedSeaDrop(address[] allowedSeaDrop) returns()
func (_Partner *PartnerTransactorSession) UpdateAllowedSeaDrop(allowedSeaDrop []common.Address) (*types.Transaction, error) {
	return _Partner.Contract.UpdateAllowedSeaDrop(&_Partner.TransactOpts, allowedSeaDrop)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x66251b69.
//
// Solidity: function updateCreatorPayoutAddress(address seaDropImpl, address payoutAddress) returns()
func (_Partner *PartnerTransactor) UpdateCreatorPayoutAddress(opts *bind.TransactOpts, seaDropImpl common.Address, payoutAddress common.Address) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "updateCreatorPayoutAddress", seaDropImpl, payoutAddress)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x66251b69.
//
// Solidity: function updateCreatorPayoutAddress(address seaDropImpl, address payoutAddress) returns()
func (_Partner *PartnerSession) UpdateCreatorPayoutAddress(seaDropImpl common.Address, payoutAddress common.Address) (*types.Transaction, error) {
	return _Partner.Contract.UpdateCreatorPayoutAddress(&_Partner.TransactOpts, seaDropImpl, payoutAddress)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x66251b69.
//
// Solidity: function updateCreatorPayoutAddress(address seaDropImpl, address payoutAddress) returns()
func (_Partner *PartnerTransactorSession) UpdateCreatorPayoutAddress(seaDropImpl common.Address, payoutAddress common.Address) (*types.Transaction, error) {
	return _Partner.Contract.UpdateCreatorPayoutAddress(&_Partner.TransactOpts, seaDropImpl, payoutAddress)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0x7a05bc82.
//
// Solidity: function updateDropURI(address seaDropImpl, string dropURI) returns()
func (_Partner *PartnerTransactor) UpdateDropURI(opts *bind.TransactOpts, seaDropImpl common.Address, dropURI string) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "updateDropURI", seaDropImpl, dropURI)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0x7a05bc82.
//
// Solidity: function updateDropURI(address seaDropImpl, string dropURI) returns()
func (_Partner *PartnerSession) UpdateDropURI(seaDropImpl common.Address, dropURI string) (*types.Transaction, error) {
	return _Partner.Contract.UpdateDropURI(&_Partner.TransactOpts, seaDropImpl, dropURI)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0x7a05bc82.
//
// Solidity: function updateDropURI(address seaDropImpl, string dropURI) returns()
func (_Partner *PartnerTransactorSession) UpdateDropURI(seaDropImpl common.Address, dropURI string) (*types.Transaction, error) {
	return _Partner.Contract.UpdateDropURI(&_Partner.TransactOpts, seaDropImpl, dropURI)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0xcb743ba8.
//
// Solidity: function updatePayer(address seaDropImpl, address payer, bool allowed) returns()
func (_Partner *PartnerTransactor) UpdatePayer(opts *bind.TransactOpts, seaDropImpl common.Address, payer common.Address, allowed bool) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "updatePayer", seaDropImpl, payer, allowed)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0xcb743ba8.
//
// Solidity: function updatePayer(address seaDropImpl, address payer, bool allowed) returns()
func (_Partner *PartnerSession) UpdatePayer(seaDropImpl common.Address, payer common.Address, allowed bool) (*types.Transaction, error) {
	return _Partner.Contract.UpdatePayer(&_Partner.TransactOpts, seaDropImpl, payer, allowed)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0xcb743ba8.
//
// Solidity: function updatePayer(address seaDropImpl, address payer, bool allowed) returns()
func (_Partner *PartnerTransactorSession) UpdatePayer(seaDropImpl common.Address, payer common.Address, allowed bool) (*types.Transaction, error) {
	return _Partner.Contract.UpdatePayer(&_Partner.TransactOpts, seaDropImpl, payer, allowed)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x1b73593c.
//
// Solidity: function updatePublicDrop(address seaDropImpl, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_Partner *PartnerTransactor) UpdatePublicDrop(opts *bind.TransactOpts, seaDropImpl common.Address, publicDrop PublicDrop) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "updatePublicDrop", seaDropImpl, publicDrop)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x1b73593c.
//
// Solidity: function updatePublicDrop(address seaDropImpl, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_Partner *PartnerSession) UpdatePublicDrop(seaDropImpl common.Address, publicDrop PublicDrop) (*types.Transaction, error) {
	return _Partner.Contract.UpdatePublicDrop(&_Partner.TransactOpts, seaDropImpl, publicDrop)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x1b73593c.
//
// Solidity: function updatePublicDrop(address seaDropImpl, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_Partner *PartnerTransactorSession) UpdatePublicDrop(seaDropImpl common.Address, publicDrop PublicDrop) (*types.Transaction, error) {
	return _Partner.Contract.UpdatePublicDrop(&_Partner.TransactOpts, seaDropImpl, publicDrop)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x511aa644.
//
// Solidity: function updateSignedMintValidationParams(address seaDropImpl, address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_Partner *PartnerTransactor) UpdateSignedMintValidationParams(opts *bind.TransactOpts, seaDropImpl common.Address, signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "updateSignedMintValidationParams", seaDropImpl, signer, signedMintValidationParams)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x511aa644.
//
// Solidity: function updateSignedMintValidationParams(address seaDropImpl, address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_Partner *PartnerSession) UpdateSignedMintValidationParams(seaDropImpl common.Address, signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _Partner.Contract.UpdateSignedMintValidationParams(&_Partner.TransactOpts, seaDropImpl, signer, signedMintValidationParams)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x511aa644.
//
// Solidity: function updateSignedMintValidationParams(address seaDropImpl, address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_Partner *PartnerTransactorSession) UpdateSignedMintValidationParams(seaDropImpl common.Address, signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _Partner.Contract.UpdateSignedMintValidationParams(&_Partner.TransactOpts, seaDropImpl, signer, signedMintValidationParams)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0x7bc2be76.
//
// Solidity: function updateTokenGatedDrop(address seaDropImpl, address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_Partner *PartnerTransactor) UpdateTokenGatedDrop(opts *bind.TransactOpts, seaDropImpl common.Address, allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "updateTokenGatedDrop", seaDropImpl, allowedNftToken, dropStage)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0x7bc2be76.
//
// Solidity: function updateTokenGatedDrop(address seaDropImpl, address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_Partner *PartnerSession) UpdateTokenGatedDrop(seaDropImpl common.Address, allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _Partner.Contract.UpdateTokenGatedDrop(&_Partner.TransactOpts, seaDropImpl, allowedNftToken, dropStage)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0x7bc2be76.
//
// Solidity: function updateTokenGatedDrop(address seaDropImpl, address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_Partner *PartnerTransactorSession) UpdateTokenGatedDrop(seaDropImpl common.Address, allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _Partner.Contract.UpdateTokenGatedDrop(&_Partner.TransactOpts, seaDropImpl, allowedNftToken, dropStage)
}

// PartnerAdministratorUpdatedIterator is returned from FilterAdministratorUpdated and is used to iterate over the raw logs and unpacked data for AdministratorUpdated events raised by the Partner contract.
type PartnerAdministratorUpdatedIterator struct {
	Event *PartnerAdministratorUpdated // Event containing the contract specifics and raw log

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
func (it *PartnerAdministratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartnerAdministratorUpdated)
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
		it.Event = new(PartnerAdministratorUpdated)
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
func (it *PartnerAdministratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartnerAdministratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartnerAdministratorUpdated represents a AdministratorUpdated event raised by the Partner contract.
type PartnerAdministratorUpdated struct {
	PreviousAdministrator common.Address
	NewAdministrator      common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterAdministratorUpdated is a free log retrieval operation binding the contract event 0x222c8e95a03c7aa95fc5d110485e0d38e767f07ab1ec878a6eac644ef1d83122.
//
// Solidity: event AdministratorUpdated(address indexed previousAdministrator, address indexed newAdministrator)
func (_Partner *PartnerFilterer) FilterAdministratorUpdated(opts *bind.FilterOpts, previousAdministrator []common.Address, newAdministrator []common.Address) (*PartnerAdministratorUpdatedIterator, error) {

	var previousAdministratorRule []interface{}
	for _, previousAdministratorItem := range previousAdministrator {
		previousAdministratorRule = append(previousAdministratorRule, previousAdministratorItem)
	}
	var newAdministratorRule []interface{}
	for _, newAdministratorItem := range newAdministrator {
		newAdministratorRule = append(newAdministratorRule, newAdministratorItem)
	}

	logs, sub, err := _Partner.contract.FilterLogs(opts, "AdministratorUpdated", previousAdministratorRule, newAdministratorRule)
	if err != nil {
		return nil, err
	}
	return &PartnerAdministratorUpdatedIterator{contract: _Partner.contract, event: "AdministratorUpdated", logs: logs, sub: sub}, nil
}

// WatchAdministratorUpdated is a free log subscription operation binding the contract event 0x222c8e95a03c7aa95fc5d110485e0d38e767f07ab1ec878a6eac644ef1d83122.
//
// Solidity: event AdministratorUpdated(address indexed previousAdministrator, address indexed newAdministrator)
func (_Partner *PartnerFilterer) WatchAdministratorUpdated(opts *bind.WatchOpts, sink chan<- *PartnerAdministratorUpdated, previousAdministrator []common.Address, newAdministrator []common.Address) (event.Subscription, error) {

	var previousAdministratorRule []interface{}
	for _, previousAdministratorItem := range previousAdministrator {
		previousAdministratorRule = append(previousAdministratorRule, previousAdministratorItem)
	}
	var newAdministratorRule []interface{}
	for _, newAdministratorItem := range newAdministrator {
		newAdministratorRule = append(newAdministratorRule, newAdministratorItem)
	}

	logs, sub, err := _Partner.contract.WatchLogs(opts, "AdministratorUpdated", previousAdministratorRule, newAdministratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartnerAdministratorUpdated)
				if err := _Partner.contract.UnpackLog(event, "AdministratorUpdated", log); err != nil {
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

// ParseAdministratorUpdated is a log parse operation binding the contract event 0x222c8e95a03c7aa95fc5d110485e0d38e767f07ab1ec878a6eac644ef1d83122.
//
// Solidity: event AdministratorUpdated(address indexed previousAdministrator, address indexed newAdministrator)
func (_Partner *PartnerFilterer) ParseAdministratorUpdated(log types.Log) (*PartnerAdministratorUpdated, error) {
	event := new(PartnerAdministratorUpdated)
	if err := _Partner.contract.UnpackLog(event, "AdministratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartnerAllowedSeaDropUpdatedIterator is returned from FilterAllowedSeaDropUpdated and is used to iterate over the raw logs and unpacked data for AllowedSeaDropUpdated events raised by the Partner contract.
type PartnerAllowedSeaDropUpdatedIterator struct {
	Event *PartnerAllowedSeaDropUpdated // Event containing the contract specifics and raw log

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
func (it *PartnerAllowedSeaDropUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartnerAllowedSeaDropUpdated)
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
		it.Event = new(PartnerAllowedSeaDropUpdated)
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
func (it *PartnerAllowedSeaDropUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartnerAllowedSeaDropUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartnerAllowedSeaDropUpdated represents a AllowedSeaDropUpdated event raised by the Partner contract.
type PartnerAllowedSeaDropUpdated struct {
	AllowedSeaDrop []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAllowedSeaDropUpdated is a free log retrieval operation binding the contract event 0xbbd3b69c138de4d317d0bc4290282c4e1cbd1e58b579a5b4f114b598c237454d.
//
// Solidity: event AllowedSeaDropUpdated(address[] allowedSeaDrop)
func (_Partner *PartnerFilterer) FilterAllowedSeaDropUpdated(opts *bind.FilterOpts) (*PartnerAllowedSeaDropUpdatedIterator, error) {

	logs, sub, err := _Partner.contract.FilterLogs(opts, "AllowedSeaDropUpdated")
	if err != nil {
		return nil, err
	}
	return &PartnerAllowedSeaDropUpdatedIterator{contract: _Partner.contract, event: "AllowedSeaDropUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedSeaDropUpdated is a free log subscription operation binding the contract event 0xbbd3b69c138de4d317d0bc4290282c4e1cbd1e58b579a5b4f114b598c237454d.
//
// Solidity: event AllowedSeaDropUpdated(address[] allowedSeaDrop)
func (_Partner *PartnerFilterer) WatchAllowedSeaDropUpdated(opts *bind.WatchOpts, sink chan<- *PartnerAllowedSeaDropUpdated) (event.Subscription, error) {

	logs, sub, err := _Partner.contract.WatchLogs(opts, "AllowedSeaDropUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartnerAllowedSeaDropUpdated)
				if err := _Partner.contract.UnpackLog(event, "AllowedSeaDropUpdated", log); err != nil {
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
func (_Partner *PartnerFilterer) ParseAllowedSeaDropUpdated(log types.Log) (*PartnerAllowedSeaDropUpdated, error) {
	event := new(PartnerAllowedSeaDropUpdated)
	if err := _Partner.contract.UnpackLog(event, "AllowedSeaDropUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartnerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Partner contract.
type PartnerApprovalIterator struct {
	Event *PartnerApproval // Event containing the contract specifics and raw log

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
func (it *PartnerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartnerApproval)
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
		it.Event = new(PartnerApproval)
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
func (it *PartnerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartnerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartnerApproval represents a Approval event raised by the Partner contract.
type PartnerApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Partner *PartnerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*PartnerApprovalIterator, error) {

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

	logs, sub, err := _Partner.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PartnerApprovalIterator{contract: _Partner.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Partner *PartnerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PartnerApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Partner.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartnerApproval)
				if err := _Partner.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Partner *PartnerFilterer) ParseApproval(log types.Log) (*PartnerApproval, error) {
	event := new(PartnerApproval)
	if err := _Partner.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartnerApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Partner contract.
type PartnerApprovalForAllIterator struct {
	Event *PartnerApprovalForAll // Event containing the contract specifics and raw log

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
func (it *PartnerApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartnerApprovalForAll)
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
		it.Event = new(PartnerApprovalForAll)
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
func (it *PartnerApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartnerApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartnerApprovalForAll represents a ApprovalForAll event raised by the Partner contract.
type PartnerApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Partner *PartnerFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*PartnerApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Partner.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PartnerApprovalForAllIterator{contract: _Partner.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Partner *PartnerFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PartnerApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Partner.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartnerApprovalForAll)
				if err := _Partner.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Partner *PartnerFilterer) ParseApprovalForAll(log types.Log) (*PartnerApprovalForAll, error) {
	event := new(PartnerApprovalForAll)
	if err := _Partner.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartnerBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Partner contract.
type PartnerBatchMetadataUpdateIterator struct {
	Event *PartnerBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *PartnerBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartnerBatchMetadataUpdate)
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
		it.Event = new(PartnerBatchMetadataUpdate)
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
func (it *PartnerBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartnerBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartnerBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Partner contract.
type PartnerBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Partner *PartnerFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*PartnerBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Partner.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &PartnerBatchMetadataUpdateIterator{contract: _Partner.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Partner *PartnerFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *PartnerBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Partner.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartnerBatchMetadataUpdate)
				if err := _Partner.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_Partner *PartnerFilterer) ParseBatchMetadataUpdate(log types.Log) (*PartnerBatchMetadataUpdate, error) {
	event := new(PartnerBatchMetadataUpdate)
	if err := _Partner.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartnerConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Partner contract.
type PartnerConsecutiveTransferIterator struct {
	Event *PartnerConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *PartnerConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartnerConsecutiveTransfer)
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
		it.Event = new(PartnerConsecutiveTransfer)
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
func (it *PartnerConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartnerConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartnerConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Partner contract.
type PartnerConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Partner *PartnerFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*PartnerConsecutiveTransferIterator, error) {

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

	logs, sub, err := _Partner.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PartnerConsecutiveTransferIterator{contract: _Partner.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Partner *PartnerFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *PartnerConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Partner.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartnerConsecutiveTransfer)
				if err := _Partner.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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
func (_Partner *PartnerFilterer) ParseConsecutiveTransfer(log types.Log) (*PartnerConsecutiveTransfer, error) {
	event := new(PartnerConsecutiveTransfer)
	if err := _Partner.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartnerContractURIUpdatedIterator is returned from FilterContractURIUpdated and is used to iterate over the raw logs and unpacked data for ContractURIUpdated events raised by the Partner contract.
type PartnerContractURIUpdatedIterator struct {
	Event *PartnerContractURIUpdated // Event containing the contract specifics and raw log

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
func (it *PartnerContractURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartnerContractURIUpdated)
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
		it.Event = new(PartnerContractURIUpdated)
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
func (it *PartnerContractURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartnerContractURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartnerContractURIUpdated represents a ContractURIUpdated event raised by the Partner contract.
type PartnerContractURIUpdated struct {
	NewContractURI string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContractURIUpdated is a free log retrieval operation binding the contract event 0x905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac37378.
//
// Solidity: event ContractURIUpdated(string newContractURI)
func (_Partner *PartnerFilterer) FilterContractURIUpdated(opts *bind.FilterOpts) (*PartnerContractURIUpdatedIterator, error) {

	logs, sub, err := _Partner.contract.FilterLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return &PartnerContractURIUpdatedIterator{contract: _Partner.contract, event: "ContractURIUpdated", logs: logs, sub: sub}, nil
}

// WatchContractURIUpdated is a free log subscription operation binding the contract event 0x905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac37378.
//
// Solidity: event ContractURIUpdated(string newContractURI)
func (_Partner *PartnerFilterer) WatchContractURIUpdated(opts *bind.WatchOpts, sink chan<- *PartnerContractURIUpdated) (event.Subscription, error) {

	logs, sub, err := _Partner.contract.WatchLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartnerContractURIUpdated)
				if err := _Partner.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
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
func (_Partner *PartnerFilterer) ParseContractURIUpdated(log types.Log) (*PartnerContractURIUpdated, error) {
	event := new(PartnerContractURIUpdated)
	if err := _Partner.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartnerMaxSupplyUpdatedIterator is returned from FilterMaxSupplyUpdated and is used to iterate over the raw logs and unpacked data for MaxSupplyUpdated events raised by the Partner contract.
type PartnerMaxSupplyUpdatedIterator struct {
	Event *PartnerMaxSupplyUpdated // Event containing the contract specifics and raw log

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
func (it *PartnerMaxSupplyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartnerMaxSupplyUpdated)
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
		it.Event = new(PartnerMaxSupplyUpdated)
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
func (it *PartnerMaxSupplyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartnerMaxSupplyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartnerMaxSupplyUpdated represents a MaxSupplyUpdated event raised by the Partner contract.
type PartnerMaxSupplyUpdated struct {
	NewMaxSupply *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMaxSupplyUpdated is a free log retrieval operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 newMaxSupply)
func (_Partner *PartnerFilterer) FilterMaxSupplyUpdated(opts *bind.FilterOpts) (*PartnerMaxSupplyUpdatedIterator, error) {

	logs, sub, err := _Partner.contract.FilterLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return &PartnerMaxSupplyUpdatedIterator{contract: _Partner.contract, event: "MaxSupplyUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxSupplyUpdated is a free log subscription operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 newMaxSupply)
func (_Partner *PartnerFilterer) WatchMaxSupplyUpdated(opts *bind.WatchOpts, sink chan<- *PartnerMaxSupplyUpdated) (event.Subscription, error) {

	logs, sub, err := _Partner.contract.WatchLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartnerMaxSupplyUpdated)
				if err := _Partner.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
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
func (_Partner *PartnerFilterer) ParseMaxSupplyUpdated(log types.Log) (*PartnerMaxSupplyUpdated, error) {
	event := new(PartnerMaxSupplyUpdated)
	if err := _Partner.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartnerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Partner contract.
type PartnerOwnershipTransferredIterator struct {
	Event *PartnerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PartnerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartnerOwnershipTransferred)
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
		it.Event = new(PartnerOwnershipTransferred)
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
func (it *PartnerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartnerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartnerOwnershipTransferred represents a OwnershipTransferred event raised by the Partner contract.
type PartnerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Partner *PartnerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PartnerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Partner.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PartnerOwnershipTransferredIterator{contract: _Partner.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Partner *PartnerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PartnerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Partner.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartnerOwnershipTransferred)
				if err := _Partner.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Partner *PartnerFilterer) ParseOwnershipTransferred(log types.Log) (*PartnerOwnershipTransferred, error) {
	event := new(PartnerOwnershipTransferred)
	if err := _Partner.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartnerPotentialAdministratorUpdatedIterator is returned from FilterPotentialAdministratorUpdated and is used to iterate over the raw logs and unpacked data for PotentialAdministratorUpdated events raised by the Partner contract.
type PartnerPotentialAdministratorUpdatedIterator struct {
	Event *PartnerPotentialAdministratorUpdated // Event containing the contract specifics and raw log

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
func (it *PartnerPotentialAdministratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartnerPotentialAdministratorUpdated)
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
		it.Event = new(PartnerPotentialAdministratorUpdated)
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
func (it *PartnerPotentialAdministratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartnerPotentialAdministratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartnerPotentialAdministratorUpdated represents a PotentialAdministratorUpdated event raised by the Partner contract.
type PartnerPotentialAdministratorUpdated struct {
	NewPotentialAdministrator common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterPotentialAdministratorUpdated is a free log retrieval operation binding the contract event 0xffa60f32d5278b35b1a3350ca92518fb5fe53e54ad07ac6355a17f54c5296b1f.
//
// Solidity: event PotentialAdministratorUpdated(address newPotentialAdministrator)
func (_Partner *PartnerFilterer) FilterPotentialAdministratorUpdated(opts *bind.FilterOpts) (*PartnerPotentialAdministratorUpdatedIterator, error) {

	logs, sub, err := _Partner.contract.FilterLogs(opts, "PotentialAdministratorUpdated")
	if err != nil {
		return nil, err
	}
	return &PartnerPotentialAdministratorUpdatedIterator{contract: _Partner.contract, event: "PotentialAdministratorUpdated", logs: logs, sub: sub}, nil
}

// WatchPotentialAdministratorUpdated is a free log subscription operation binding the contract event 0xffa60f32d5278b35b1a3350ca92518fb5fe53e54ad07ac6355a17f54c5296b1f.
//
// Solidity: event PotentialAdministratorUpdated(address newPotentialAdministrator)
func (_Partner *PartnerFilterer) WatchPotentialAdministratorUpdated(opts *bind.WatchOpts, sink chan<- *PartnerPotentialAdministratorUpdated) (event.Subscription, error) {

	logs, sub, err := _Partner.contract.WatchLogs(opts, "PotentialAdministratorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartnerPotentialAdministratorUpdated)
				if err := _Partner.contract.UnpackLog(event, "PotentialAdministratorUpdated", log); err != nil {
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

// ParsePotentialAdministratorUpdated is a log parse operation binding the contract event 0xffa60f32d5278b35b1a3350ca92518fb5fe53e54ad07ac6355a17f54c5296b1f.
//
// Solidity: event PotentialAdministratorUpdated(address newPotentialAdministrator)
func (_Partner *PartnerFilterer) ParsePotentialAdministratorUpdated(log types.Log) (*PartnerPotentialAdministratorUpdated, error) {
	event := new(PartnerPotentialAdministratorUpdated)
	if err := _Partner.contract.UnpackLog(event, "PotentialAdministratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartnerPotentialOwnerUpdatedIterator is returned from FilterPotentialOwnerUpdated and is used to iterate over the raw logs and unpacked data for PotentialOwnerUpdated events raised by the Partner contract.
type PartnerPotentialOwnerUpdatedIterator struct {
	Event *PartnerPotentialOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *PartnerPotentialOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartnerPotentialOwnerUpdated)
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
		it.Event = new(PartnerPotentialOwnerUpdated)
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
func (it *PartnerPotentialOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartnerPotentialOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartnerPotentialOwnerUpdated represents a PotentialOwnerUpdated event raised by the Partner contract.
type PartnerPotentialOwnerUpdated struct {
	NewPotentialAdministrator common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterPotentialOwnerUpdated is a free log retrieval operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address newPotentialAdministrator)
func (_Partner *PartnerFilterer) FilterPotentialOwnerUpdated(opts *bind.FilterOpts) (*PartnerPotentialOwnerUpdatedIterator, error) {

	logs, sub, err := _Partner.contract.FilterLogs(opts, "PotentialOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &PartnerPotentialOwnerUpdatedIterator{contract: _Partner.contract, event: "PotentialOwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchPotentialOwnerUpdated is a free log subscription operation binding the contract event 0x11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da.
//
// Solidity: event PotentialOwnerUpdated(address newPotentialAdministrator)
func (_Partner *PartnerFilterer) WatchPotentialOwnerUpdated(opts *bind.WatchOpts, sink chan<- *PartnerPotentialOwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _Partner.contract.WatchLogs(opts, "PotentialOwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartnerPotentialOwnerUpdated)
				if err := _Partner.contract.UnpackLog(event, "PotentialOwnerUpdated", log); err != nil {
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
func (_Partner *PartnerFilterer) ParsePotentialOwnerUpdated(log types.Log) (*PartnerPotentialOwnerUpdated, error) {
	event := new(PartnerPotentialOwnerUpdated)
	if err := _Partner.contract.UnpackLog(event, "PotentialOwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartnerProvenanceHashUpdatedIterator is returned from FilterProvenanceHashUpdated and is used to iterate over the raw logs and unpacked data for ProvenanceHashUpdated events raised by the Partner contract.
type PartnerProvenanceHashUpdatedIterator struct {
	Event *PartnerProvenanceHashUpdated // Event containing the contract specifics and raw log

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
func (it *PartnerProvenanceHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartnerProvenanceHashUpdated)
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
		it.Event = new(PartnerProvenanceHashUpdated)
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
func (it *PartnerProvenanceHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartnerProvenanceHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartnerProvenanceHashUpdated represents a ProvenanceHashUpdated event raised by the Partner contract.
type PartnerProvenanceHashUpdated struct {
	PreviousHash [32]byte
	NewHash      [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProvenanceHashUpdated is a free log retrieval operation binding the contract event 0x7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c.
//
// Solidity: event ProvenanceHashUpdated(bytes32 previousHash, bytes32 newHash)
func (_Partner *PartnerFilterer) FilterProvenanceHashUpdated(opts *bind.FilterOpts) (*PartnerProvenanceHashUpdatedIterator, error) {

	logs, sub, err := _Partner.contract.FilterLogs(opts, "ProvenanceHashUpdated")
	if err != nil {
		return nil, err
	}
	return &PartnerProvenanceHashUpdatedIterator{contract: _Partner.contract, event: "ProvenanceHashUpdated", logs: logs, sub: sub}, nil
}

// WatchProvenanceHashUpdated is a free log subscription operation binding the contract event 0x7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c.
//
// Solidity: event ProvenanceHashUpdated(bytes32 previousHash, bytes32 newHash)
func (_Partner *PartnerFilterer) WatchProvenanceHashUpdated(opts *bind.WatchOpts, sink chan<- *PartnerProvenanceHashUpdated) (event.Subscription, error) {

	logs, sub, err := _Partner.contract.WatchLogs(opts, "ProvenanceHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartnerProvenanceHashUpdated)
				if err := _Partner.contract.UnpackLog(event, "ProvenanceHashUpdated", log); err != nil {
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
func (_Partner *PartnerFilterer) ParseProvenanceHashUpdated(log types.Log) (*PartnerProvenanceHashUpdated, error) {
	event := new(PartnerProvenanceHashUpdated)
	if err := _Partner.contract.UnpackLog(event, "ProvenanceHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartnerRoyaltyInfoUpdatedIterator is returned from FilterRoyaltyInfoUpdated and is used to iterate over the raw logs and unpacked data for RoyaltyInfoUpdated events raised by the Partner contract.
type PartnerRoyaltyInfoUpdatedIterator struct {
	Event *PartnerRoyaltyInfoUpdated // Event containing the contract specifics and raw log

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
func (it *PartnerRoyaltyInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartnerRoyaltyInfoUpdated)
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
		it.Event = new(PartnerRoyaltyInfoUpdated)
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
func (it *PartnerRoyaltyInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartnerRoyaltyInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartnerRoyaltyInfoUpdated represents a RoyaltyInfoUpdated event raised by the Partner contract.
type PartnerRoyaltyInfoUpdated struct {
	Receiver common.Address
	Bps      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyInfoUpdated is a free log retrieval operation binding the contract event 0xf21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d.
//
// Solidity: event RoyaltyInfoUpdated(address receiver, uint256 bps)
func (_Partner *PartnerFilterer) FilterRoyaltyInfoUpdated(opts *bind.FilterOpts) (*PartnerRoyaltyInfoUpdatedIterator, error) {

	logs, sub, err := _Partner.contract.FilterLogs(opts, "RoyaltyInfoUpdated")
	if err != nil {
		return nil, err
	}
	return &PartnerRoyaltyInfoUpdatedIterator{contract: _Partner.contract, event: "RoyaltyInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchRoyaltyInfoUpdated is a free log subscription operation binding the contract event 0xf21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d.
//
// Solidity: event RoyaltyInfoUpdated(address receiver, uint256 bps)
func (_Partner *PartnerFilterer) WatchRoyaltyInfoUpdated(opts *bind.WatchOpts, sink chan<- *PartnerRoyaltyInfoUpdated) (event.Subscription, error) {

	logs, sub, err := _Partner.contract.WatchLogs(opts, "RoyaltyInfoUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartnerRoyaltyInfoUpdated)
				if err := _Partner.contract.UnpackLog(event, "RoyaltyInfoUpdated", log); err != nil {
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
func (_Partner *PartnerFilterer) ParseRoyaltyInfoUpdated(log types.Log) (*PartnerRoyaltyInfoUpdated, error) {
	event := new(PartnerRoyaltyInfoUpdated)
	if err := _Partner.contract.UnpackLog(event, "RoyaltyInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartnerSeaDropTokenDeployedIterator is returned from FilterSeaDropTokenDeployed and is used to iterate over the raw logs and unpacked data for SeaDropTokenDeployed events raised by the Partner contract.
type PartnerSeaDropTokenDeployedIterator struct {
	Event *PartnerSeaDropTokenDeployed // Event containing the contract specifics and raw log

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
func (it *PartnerSeaDropTokenDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartnerSeaDropTokenDeployed)
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
		it.Event = new(PartnerSeaDropTokenDeployed)
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
func (it *PartnerSeaDropTokenDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartnerSeaDropTokenDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartnerSeaDropTokenDeployed represents a SeaDropTokenDeployed event raised by the Partner contract.
type PartnerSeaDropTokenDeployed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSeaDropTokenDeployed is a free log retrieval operation binding the contract event 0xd7aca75208b9be5ffc04c6a01922020ffd62b55e68e502e317f5344960279af8.
//
// Solidity: event SeaDropTokenDeployed()
func (_Partner *PartnerFilterer) FilterSeaDropTokenDeployed(opts *bind.FilterOpts) (*PartnerSeaDropTokenDeployedIterator, error) {

	logs, sub, err := _Partner.contract.FilterLogs(opts, "SeaDropTokenDeployed")
	if err != nil {
		return nil, err
	}
	return &PartnerSeaDropTokenDeployedIterator{contract: _Partner.contract, event: "SeaDropTokenDeployed", logs: logs, sub: sub}, nil
}

// WatchSeaDropTokenDeployed is a free log subscription operation binding the contract event 0xd7aca75208b9be5ffc04c6a01922020ffd62b55e68e502e317f5344960279af8.
//
// Solidity: event SeaDropTokenDeployed()
func (_Partner *PartnerFilterer) WatchSeaDropTokenDeployed(opts *bind.WatchOpts, sink chan<- *PartnerSeaDropTokenDeployed) (event.Subscription, error) {

	logs, sub, err := _Partner.contract.WatchLogs(opts, "SeaDropTokenDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartnerSeaDropTokenDeployed)
				if err := _Partner.contract.UnpackLog(event, "SeaDropTokenDeployed", log); err != nil {
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
func (_Partner *PartnerFilterer) ParseSeaDropTokenDeployed(log types.Log) (*PartnerSeaDropTokenDeployed, error) {
	event := new(PartnerSeaDropTokenDeployed)
	if err := _Partner.contract.UnpackLog(event, "SeaDropTokenDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartnerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Partner contract.
type PartnerTransferIterator struct {
	Event *PartnerTransfer // Event containing the contract specifics and raw log

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
func (it *PartnerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartnerTransfer)
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
		it.Event = new(PartnerTransfer)
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
func (it *PartnerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartnerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartnerTransfer represents a Transfer event raised by the Partner contract.
type PartnerTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Partner *PartnerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*PartnerTransferIterator, error) {

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

	logs, sub, err := _Partner.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PartnerTransferIterator{contract: _Partner.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Partner *PartnerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PartnerTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Partner.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartnerTransfer)
				if err := _Partner.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Partner *PartnerFilterer) ParseTransfer(log types.Log) (*PartnerTransfer, error) {
	event := new(PartnerTransfer)
	if err := _Partner.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
