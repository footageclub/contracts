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

// PartnerMetaData contains all meta data concerning the Partner contract.
var PartnerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AdministratorMustInitializeWithFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BurnIncorrectSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"CannotExceedMaxSupplyOfUint64\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"basisPoints\",\"type\":\"uint256\"}],\"name\":\"InvalidRoyaltyBasisPoints\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAdministratorIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotNextAdministrator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotNextOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdministrator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedSeaDrop\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrAdministrator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProvenanceHashCannotBeSetAfterMintStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoyaltyAddressCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignersMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousAdministrator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdministrator\",\"type\":\"address\"}],\"name\":\"AdministratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowedSeaDrop\",\"type\":\"address[]\"}],\"name\":\"AllowedSeaDropUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"}],\"name\":\"ContractURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPotentialAdministrator\",\"type\":\"address\"}],\"name\":\"PotentialAdministratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPotentialAdministrator\",\"type\":\"address\"}],\"name\":\"PotentialOwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newHash\",\"type\":\"bytes32\"}],\"name\":\"ProvenanceHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"name\":\"RoyaltyInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SeaDropTokenDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptAdministration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"administrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelAdministrationTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"}],\"name\":\"emitBatchMetadataUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBurnAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"getMintStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minterNumMinted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"administrator\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"allowedSeaDrop\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"mintSeaDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contractURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"dropURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"publicKeyURIs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"internalType\":\"structAllowListData\",\"name\":\"allowListData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"creatorPayoutAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"provenanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"allowedFeeRecipients\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedFeeRecipients\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowedPayers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedPayers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenGatedAllowedNftTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structTokenGatedDropStage[]\",\"name\":\"tokenGatedDropStages\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedTokenGatedAllowedNftTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"internalType\":\"structSignedMintValidationParams[]\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedSigners\",\"type\":\"address[]\"}],\"internalType\":\"structERC721SeaDropStructsErrorsAndEventsUpgradeable.MultiConfigureStruct\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"multiConfigure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"potentialAdministrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provenanceHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdministration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBurnAddress\",\"type\":\"address\"}],\"name\":\"setBurnAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newProvenanceHash\",\"type\":\"bytes32\"}],\"name\":\"setProvenanceHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"royaltyAddress\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"royaltyBps\",\"type\":\"uint96\"}],\"internalType\":\"structISeaDropTokenContractMetadataUpgradeable.RoyaltyInfo\",\"name\":\"newInfo\",\"type\":\"tuple\"}],\"name\":\"setRoyaltyInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdministrator\",\"type\":\"address\"}],\"name\":\"transferAdministration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPotentialOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"publicKeyURIs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"internalType\":\"structAllowListData\",\"name\":\"allowListData\",\"type\":\"tuple\"}],\"name\":\"updateAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateAllowedFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"allowedSeaDrop\",\"type\":\"address[]\"}],\"name\":\"updateAllowedSeaDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payoutAddress\",\"type\":\"address\"}],\"name\":\"updateCreatorPayoutAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"dropURI\",\"type\":\"string\"}],\"name\":\"updateDropURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updatePayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"}],\"name\":\"updatePublicDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"internalType\":\"structSignedMintValidationParams\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple\"}],\"name\":\"updateSignedMintValidationParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structTokenGatedDropStage\",\"name\":\"dropStage\",\"type\":\"tuple\"}],\"name\":\"updateTokenGatedDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615566806100206000396000f3fe608060405234801561001057600080fd5b50600436106103425760003560e01c8063715018a6116101b8578063ad2f852a11610104578063d5e7feb8116100a2578063f0025d961161007c578063f0025d9614610718578063f2fde38b14610720578063f53d0a8e14610733578063fbc46ada1461073b57600080fd5b8063d5e7feb8146106f5578063e8a3d485146106fd578063e985e9c51461070557600080fd5b8063c780b63d116100de578063c780b63d146106bf578063c87b56dd146106c7578063cb743ba8146106da578063d5abeb01146106ed57600080fd5b8063ad2f852a14610667578063b88d4fde14610685578063c6ab67a31461069857600080fd5b8063911f456b1161017157806395d89b411161014b57806395d89b41146106315780639794ed4014610639578063a22cb46514610641578063a48301141461065457600080fd5b8063911f456b146105f8578063913ee93d1461060b578063938e3d7b1461061e57600080fd5b8063715018a61461058c57806379ba5097146105945780637a05bc821461059c5780637bc2be76146105af578063840e15d4146105c25780638da5cb5b146105f057600080fd5b806342842e0e1161029257806360c308b61161023057806366251b691161020a57806366251b691461054b5780636c0360eb1461055e5780636f8b44b01461056657806370a082311461057957600080fd5b806360c308b6146105125780636352211e1461052557806364869dad1461053857600080fd5b806348a4c1011161026c57806348a4c101146104c65780634b0e7216146104d9578063511aa644146104ec57806355f804b3146104ff57600080fd5b806342842e0e1461048d57806342966c68146104a057806344dae42c146104b357600080fd5b80631b73593c116102ff5780632a55205a116102d95780632a55205a1461041b5780633680620d1461044d57806338b39d291461046057806342260b5d1461046857600080fd5b80631b73593c146103ed57806323452b9c1461040057806323b872dd1461040857600080fd5b806301ffc9a71461034757806306fdde031461036f578063081812fc14610384578063095ea7b3146103af578063099b6bfa146103c457806318160ddd146103d7575b600080fd5b61035a610355366004613d05565b61074e565b60405190151581526020015b60405180910390f35b610377610794565b6040516103669190613d72565b610397610392366004613d85565b61082f565b6040516001600160a01b039091168152602001610366565b6103c26103bd366004613db3565b61087c565b005b6103c26103d2366004613d85565b610895565b6103df61092b565b604051908152602001610366565b6103c26103fb366004613ddf565b61094b565b6103c2610b6b565b6103c2610416366004613e20565b610bc7565b61042e610429366004613e61565b610bf2565b604080516001600160a01b039093168352602083019190915201610366565b6103c261045b366004613e83565b610c45565b610397610d01565b6000805160206154f183398151915254600160a01b90046001600160601b03166103df565b6103c261049b366004613e20565b610d34565b6103c26104ae366004613d85565b610d59565b6103c26104c1366004613ed9565b610daf565b6103c26104d4366004613f0f565b610ed2565b6103c26104e7366004613f5a565b610f78565b6103c26104fa366004614089565b610fc1565b6103c261050d3660046141a8565b6111a3565b6103c26105203660046141e9565b611234565b610397610533366004613d85565b61128f565b6103c2610546366004613db3565b61129a565b6103c261055936600461425d565b61130b565b61037761134a565b6103c2610574366004613d85565b611359565b6103df610587366004613f5a565b6113cd565b6103c2611435565b6103c2611449565b6103c26105aa36600461428b565b6114dc565b6103c26105bd3660046142df565b611564565b6105d56105d0366004613f5a565b61170b565b60408051938452602084019290925290820152606001610366565b61039761173d565b6103c2610606366004614334565b611765565b6103c2610619366004613f5a565b612382565b6103c261062c3660046141a8565b612435565b61037761249b565b6103c26124b3565b6103c261064f36600461436f565b612531565b6103c2610662366004613e61565b612545565b6000805160206154f1833981519152546001600160a01b0316610397565b6103c26106933660046143f4565b612583565b7fb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031c546103df565b6103c26125b0565b6103776106d5366004613d85565b612629565b6103c26106e8366004613f0f565b6126ad565b6103df61273d565b610397612755565b610377612771565b61035a61071336600461425d565b612787565b6103c26127c4565b6103c261072e366004613f5a565b612822565b6103976128ac565b6103c2610749366004614493565b6128b6565b60006001600160e01b03198216630c487f4760e11b148061077f57506001600160e01b03198216639c15441560e01b145b8061078e575061078e82612ad5565b92915050565b606061079e612b15565b60020180546107ac9061459c565b80601f01602080910402602001604051908101604052809291908181526020018280546107d89061459c565b80156108255780601f106107fa57610100808354040283529160200191610825565b820191906000526020600020905b81548152906001019060200180831161080857829003601f168201915b5050505050905090565b600061083a82612b39565b610857576040516333d1c03960e21b815260040160405180910390fd5b61085f612b15565b60009283526006016020525060409020546001600160a01b031690565b8161088681612b82565b6108908383612c3b565b505050565b61089d612ce9565b60006108a7612d2f565b11156108c65760405163e03264af60e01b815260040160405180910390fd5b7fb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031c80549082905560408051828152602081018490527f7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c91015b60405180910390a15050565b60006001610937612b15565b60010154610943612b15565b540303919050565b61095361173d565b6001600160a01b0316336001600160a01b03161461099c57610973612d42565b546001600160a01b0316331461099c576040516359d9793760e01b815260040160405180910390fd5b6109a582612d66565b604051632f1a98a760e21b81523060048201526000906001600160a01b0384169063bc6a629c9060240160c060405180830381865afa1580156109ec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a109190614605565b90506000610a233684900384018461469d565b9050610a2d612d42565b546001600160a01b03163314610a8057816060015161ffff16600003610a6657604051634f4ca83d60e11b815260040160405180910390fd5b60808083015161ffff1690820152600160a0820152610aba565b606082015161ffff8116610a95576001610a97565b805b61ffff9081166060850152608092830151169183019190915250600160a0820152805b604080516301308e6560e01b815282516001600160501b03166004820152602083015165ffffffffffff9081166024830152918301519091166044820152606082015161ffff9081166064830152608083015116608482015260a0820151151560a48201526001600160a01b038516906301308e659060c401600060405180830381600087803b158015610b4d57600080fd5b505af1158015610b61573d6000803e3d6000fd5b5050505050505050565b610b73612db1565b60008051602061547183398151915280546001600160a01b0319169055604051600081527f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da906020015b60405180910390a1565b826001600160a01b0381163314610be157610be133612b82565b610bec848484612dfb565b50505050565b6000805160206154f18339815191528054600091829161271090610c2690600160a01b90046001600160601b031686614734565b610c30919061474b565b90546001600160a01b03169590945092505050565b610c4d61173d565b6001600160a01b0316336001600160a01b031614610c9657610c6d612d42565b546001600160a01b03163314610c96576040516359d9793760e01b815260040160405180910390fd5b610c9f82612d66565b60405163ebb4a55f60e01b81526001600160a01b0383169063ebb4a55f90610ccb9084906004016148ae565b600060405180830381600087803b158015610ce557600080fd5b505af1158015610cf9573d6000803e3d6000fd5b505050505050565b60007f756747167942d99688c60edc6b1b1d2b892b7afd96275ad2b396f42e61e7f0e05b546001600160a01b0316919050565b826001600160a01b0381163314610d4e57610d4e33612b82565b610bec848484612fdc565b7f756747167942d99688c60edc6b1b1d2b892b7afd96275ad2b396f42e61e7f0e0546001600160a01b03163314610da35760405163ea885d8360e01b815260040160405180910390fd5b610dac81612ff7565b50565b610db7612ce9565b6000610dc66020830183613f5a565b6001600160a01b031603610ded57604051631cc0baef60e01b815260040160405180910390fd5b612710610e0060408301602084016148d6565b6001600160601b03161115610e4957610e1f60408201602083016148d6565b604051633cadbafb60e01b81526001600160601b0390911660048201526024015b60405180910390fd5b806000805160206154f1833981519152610e6382826148f3565b507ff21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d9050610e946020830183613f5a565b610ea460408401602085016148d6565b604080516001600160a01b0390931683526001600160601b039091166020830152015b60405180910390a150565b610eda612d42565b546001600160a01b03163314610f06576040516001620aed3360e41b0319815260040160405180910390fd5b610f0f83612d66565b604051638e7d1e4360e01b81526001600160a01b0383811660048301528215156024830152841690638e7d1e43906044015b600060405180830381600087803b158015610f5b57600080fd5b505af1158015610f6f573d6000803e3d6000fd5b50505050505050565b610f80612db1565b7f756747167942d99688c60edc6b1b1d2b892b7afd96275ad2b396f42e61e7f0e080546001600160a01b0319166001600160a01b0392909216919091179055565b610fc961173d565b6001600160a01b0316336001600160a01b03161461101257610fe9612d42565b546001600160a01b03163314611012576040516359d9793760e01b815260040160405180910390fd5b61101b83612d66565b6040516381bf9af360e01b81523060048201526001600160a01b038381166024830152600091908516906381bf9af39060440160e060405180830381865afa15801561106b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061108f9190614940565b90508161109a612d42565b546001600160a01b031633146110f957816020015162ffffff166000036110d457604051634f4ca83d60e11b815260040160405180910390fd5b60a08083015161ffff9081169183019190915260c0808401519091169082015261113c565b602082015162ffffff811661110f576001611111565b805b62ffffff1660208401525060a08082015161ffff9081169184019190915260c0918201511690820152805b6040516309a7002f60e31b81526001600160a01b03861690634d3801789061116a9087908590600401614a62565b600060405180830381600087803b15801561118457600080fd5b505af1158015611198573d6000803e3d6000fd5b505050505050505050565b6111ab612ce9565b7fb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031a6111d7828483614ac6565b506111e061092b565b15611230577f6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c600180611211613002565b61121b9190614b85565b6040805192835260208301919091520161091f565b5050565b61123c61173d565b6001600160a01b0316336001600160a01b0316146112855761125c612d42565b546001600160a01b03163314611285576040516359d9793760e01b815260040160405180910390fd5b611230828261300c565b600061078e8261318e565b6112a333612d66565b6112ab61273d565b816112b4612d2f565b6112be9190614b98565b111561130157806112cd612d2f565b6112d79190614b98565b6112df61273d565b60405163384b48c560e21b815260048101929092526024820152604401610e40565b6112308282613222565b611313612db1565b61131c82612d66565b60405163024e71b760e31b81526001600160a01b0382811660048301528316906312738db890602401610ccb565b6060611354613339565b905090565b611361612ce9565b6001600160401b0381111561138c5760405163b43e913760e01b815260048101829052602401610e40565b806000805160206154d1833981519152556040518181527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c90602001610ec7565b60006001600160a01b0382166113f6576040516323d3ad8160e21b815260040160405180910390fd5b6001600160401b03611406612b15565b6005016000846001600160a01b03166001600160a01b0316815260200190815260200160002054169050919050565b61143d612db1565b6114476000613358565b565b600080516020615471833981519152546001600160a01b031633811461148257604051636b7584e760e11b815260040160405180910390fd5b60008051602061547183398151915280546001600160a01b0319169055604080516000815290517f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da9181900360200190a1610dac81613358565b6114e461173d565b6001600160a01b0316336001600160a01b03161461152d57611504612d42565b546001600160a01b0316331461152d576040516359d9793760e01b815260040160405180910390fd5b61153683612d66565b60405163b957d0cb60e01b81526001600160a01b0384169063b957d0cb90610f419085908590600401614bab565b61156c61173d565b6001600160a01b0316336001600160a01b0316146115b55761158c612d42565b546001600160a01b031633146115b5576040516359d9793760e01b815260040160405180910390fd5b6115be83612d66565b604051630587453760e11b81523060048201526001600160a01b03838116602483015260009190851690630b0e8a6e9060440161010060405180830381865afa15801561160f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116339190614bf6565b9050600061164636849003840184614cb4565b9050611650612d42565b546001600160a01b031633146116a357816020015161ffff1660000361168957604051634f4ca83d60e11b815260040160405180910390fd5b60c08083015161ffff1690820152600160e08201526116dd565b602082015161ffff81166116b85760016116ba565b805b61ffff908116602085015260c092830151169183019190915250600160e0820152805b604051637ecd591560e11b81526001600160a01b0386169063fd9ab22a9061116a9087908590600401614ddb565b6000806000611719846133c9565b9250611723612d2f565b91506000805160206154d183398151915254929491935050565b60007ff73863e3917595a7e67829ed60f0c51bf14c7a0d0de47c7b02a00abd48488e4a610d25565b61176d612db1565b8035156117c6576040516306f8b44b60e41b8152813560048201523090636f8b44b090602401600060405180830381600087803b1580156117ad57600080fd5b505af11580156117c1573d6000803e3d6000fd5b505050505b6117d36020820182614df9565b15905061183d57306355f804b36117ed6020840184614df9565b6040518363ffffffff1660e01b815260040161180a929190614bab565b600060405180830381600087803b15801561182457600080fd5b505af1158015611838573d6000803e3d6000fd5b505050505b61184a6040820182614df9565b1590506118b4573063938e3d7b6118646040840184614df9565b6040518363ffffffff1660e01b8152600401611881929190614bab565b600060405180830381600087803b15801561189b57600080fd5b505af11580156118af573d6000803e3d6000fd5b505050505b6118d46118c760e0830160c08401614e3f565b65ffffffffffff16151590565b6118e76118c760c0840160a08501614e3f565b176001036119595730631b73593c6119056080840160608501613f5a565b836080016040518363ffffffff1660e01b8152600401611926929190614e5c565b600060405180830381600087803b15801561194057600080fd5b505af1158015611954573d6000803e3d6000fd5b505050505b611967610140820182614df9565b1590506119e35730637a05bc826119846080840160608501613f5a565b611992610140850185614df9565b6040518463ffffffff1660e01b81526004016119b093929190614f0a565b600060405180830381600087803b1580156119ca57600080fd5b505af11580156119de573d6000803e3d6000fd5b505050505b60006119f3610160830183614f38565b3514611a6d5730633680620d611a0f6080840160608501613f5a565b611a1d610160850185614f38565b6040518363ffffffff1660e01b8152600401611a3a929190614f58565b600060405180830381600087803b158015611a5457600080fd5b505af1158015611a68573d6000803e3d6000fd5b505050505b6000611a816101a083016101808401613f5a565b6001600160a01b031614611b1857306366251b69611aa56080840160608501613f5a565b611ab76101a085016101808601613f5a565b6040516001600160e01b031960e085901b1681526001600160a01b03928316600482015291166024820152604401600060405180830381600087803b158015611aff57600080fd5b505af1158015611b13573d6000803e3d6000fd5b505050505b6101a081013515611b79576040516304cdb5fd60e11b81526101a08201356004820152309063099b6bfa90602401600060405180830381600087803b158015611b6057600080fd5b505af1158015611b74573d6000803e3d6000fd5b505050505b6000611b896101c0830183614f7c565b90501115611c535760005b611ba26101c0830183614f7c565b9050811015611c5157306348a4c101611bc16080850160608601613f5a565b611bcf6101c0860186614f7c565b85818110611bdf57611bdf614fc5565b9050602002016020810190611bf49190613f5a565b60016040518463ffffffff1660e01b8152600401611c1493929190614fdb565b600060405180830381600087803b158015611c2e57600080fd5b505af1158015611c42573d6000803e3d6000fd5b50505050806001019050611b94565b505b6000611c636101e0830183614f7c565b90501115611d2d5760005b611c7c6101e0830183614f7c565b9050811015611d2b57306348a4c101611c9b6080850160608601613f5a565b611ca96101e0860186614f7c565b85818110611cb957611cb9614fc5565b9050602002016020810190611cce9190613f5a565b60006040518463ffffffff1660e01b8152600401611cee93929190614fdb565b600060405180830381600087803b158015611d0857600080fd5b505af1158015611d1c573d6000803e3d6000fd5b50505050806001019050611c6e565b505b6000611d3d610200830183614f7c565b90501115611e075760005b611d56610200830183614f7c565b9050811015611e05573063cb743ba8611d756080850160608601613f5a565b611d83610200860186614f7c565b85818110611d9357611d93614fc5565b9050602002016020810190611da89190613f5a565b60016040518463ffffffff1660e01b8152600401611dc893929190614fdb565b600060405180830381600087803b158015611de257600080fd5b505af1158015611df6573d6000803e3d6000fd5b50505050806001019050611d48565b505b6000611e17610220830183614f7c565b90501115611ee15760005b611e30610220830183614f7c565b9050811015611edf573063cb743ba8611e4f6080850160608601613f5a565b611e5d610220860186614f7c565b85818110611e6d57611e6d614fc5565b9050602002016020810190611e829190613f5a565b60006040518463ffffffff1660e01b8152600401611ea293929190614fdb565b600060405180830381600087803b158015611ebc57600080fd5b505af1158015611ed0573d6000803e3d6000fd5b50505050806001019050611e22565b505b6000611ef1610260830183614fff565b9050111561201c57611f07610240820182614f7c565b9050611f17610260830183614fff565b905014611f375760405163b81aa63960e01b815260040160405180910390fd5b60005b611f48610260830183614fff565b905081101561201a5730637bc2be76611f676080850160608601613f5a565b611f75610240860186614f7c565b85818110611f8557611f85614fc5565b9050602002016020810190611f9a9190613f5a565b611fa8610260870187614fff565b86818110611fb857611fb8614fc5565b905061010002016040518463ffffffff1660e01b8152600401611fdd93929190615048565b600060405180830381600087803b158015611ff757600080fd5b505af115801561200b573d6000803e3d6000fd5b50505050806001019050611f3a565b505b600061202c610280830183614f7c565b905011156121375760005b612045610280830183614f7c565b9050811015612135576040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915230637bc2be766120a56080860160608701613f5a565b6120b3610280870187614f7c565b868181106120c3576120c3614fc5565b90506020020160208101906120d89190613f5a565b846040518463ffffffff1660e01b81526004016120f793929190615123565b600060405180830381600087803b15801561211157600080fd5b505af1158015612125573d6000803e3d6000fd5b5050505081600101915050612037565b505b60006121476102c0830183615149565b905011156122715761215d6102a0820182614f7c565b905061216d6102c0830183615149565b90501461218d576040516374ef6df760e01b815260040160405180910390fd5b60005b61219e6102c0830183615149565b905081101561226f573063511aa6446121bd6080850160608601613f5a565b6121cb6102a0860186614f7c565b858181106121db576121db614fc5565b90506020020160208101906121f09190613f5a565b6121fe6102c0870187615149565b8681811061220e5761220e614fc5565b905060e002016040518463ffffffff1660e01b815260040161223293929190615191565b600060405180830381600087803b15801561224c57600080fd5b505af1158015612260573d6000803e3d6000fd5b50505050806001019050612190565b505b60006122816102e0830183614f7c565b90501115610dac5760005b61229a6102e0830183614f7c565b9050811015611230576040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091523063511aa6446122f26080860160608701613f5a565b6123006102e0870187614f7c565b8681811061231057612310614fc5565b90506020020160208101906123259190613f5a565b846040518463ffffffff1660e01b815260040161234493929190615256565b600060405180830381600087803b15801561235e57600080fd5b505af1158015612372573d6000803e3d6000fd5b505050508160010191505061228c565b61238a612d42565b546001600160a01b031633146123b6576040516001620aed3360e41b0319815260040160405180910390fd5b6001600160a01b0381166123dd57604051633536be7f60e21b815260040160405180910390fd5b806123e6612d42565b60010180546001600160a01b0319166001600160a01b0392831617905560405190821681527fffa60f32d5278b35b1a3350ca92518fb5fe53e54ad07ac6355a17f54c5296b1f90602001610ec7565b61243d612ce9565b7fb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031b612469828483614ac6565b507f905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac37378828260405161091f929190614bab565b60606124a5612b15565b60030180546107ac9061459c565b6124bb612d42565b546001600160a01b031633146124e7576040516001620aed3360e41b0319815260040160405180910390fd5b6124ef612d42565b60010180546001600160a01b0319169055604051600081527fffa60f32d5278b35b1a3350ca92518fb5fe53e54ad07ac6355a17f54c5296b1f90602001610bbd565b8161253b81612b82565b610890838361340e565b61254d612ce9565b60408051838152602081018390527f6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c910161091f565b836001600160a01b038116331461259d5761259d33612b82565b6125a98585858561348b565b5050505050565b6125b8612d42565b546001600160a01b031633146125e4576040516001620aed3360e41b0319815260040160405180910390fd5b6125ec612d42565b80546001600160a01b031916905560405160009033907f222c8e95a03c7aa95fc5d110485e0d38e767f07ab1ec878a6eac644ef1d83122908390a3565b606061263482612b39565b61265157604051630a14c4b560e41b815260040160405180910390fd5b600061265b613339565b9050805160000361267b57604051806020016040528060008152506126a6565b80612685846134cf565b60405160200161269692919061527c565b6040516020818303038152906040525b9392505050565b6126b561173d565b6001600160a01b0316336001600160a01b0316146126fe576126d5612d42565b546001600160a01b031633146126fe576040516359d9793760e01b815260040160405180910390fd5b61270783612d66565b604051633f952e6560e11b81526001600160a01b0383811660048301528215156024830152841690637f2a5cca90604401610f41565b60006000805160206154d18339815191525b54919050565b600061275f612d42565b600101546001600160a01b0316919050565b60606000805160206154d183398151915261079e565b6000612791612b15565b6001600160a01b039384166000908152600791909101602090815260408083209490951682529290925250205460ff1690565b60006127ce612d42565b600101546001600160a01b031690503381146127fd576040516353bb059b60e01b815260040160405180910390fd5b61280681613513565b61280e612d42565b60010180546001600160a01b031916905550565b61282a612db1565b6001600160a01b03811661285157604051633a247dd760e11b815260040160405180910390fd5b60008051602061547183398151915280546001600160a01b0319166001600160a01b0383169081179091556040519081527f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da90602001610ec7565b6000610d25612d42565b600054610100900460ff16158080156128d65750600054600160ff909116105b806128f05750303b1580156128f0575060005460ff166001145b6129535760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610e40565b6000805460ff191660011790558015612976576000805461ff0019166101001790555b60008051602061551183398151915254610100900460ff166129ab576000805160206155118339815191525460ff16156129af565b303b155b612a215760405162461bcd60e51b815260206004820152603760248201527f455243373231415f5f496e697469616c697a61626c653a20636f6e747261637460448201527f20697320616c726561647920696e697469616c697a65640000000000000000006064820152608401610e40565b60008051602061551183398151915254610100900460ff16158015612a5d57600080516020615511833981519152805461ffff19166101011790555b612a6986868686613568565b8015612a8857600080516020615511833981519152805461ff00191690555b5080156125a9576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050505050565b60006001600160e01b0319821663152a902d60e11b1480612b065750632483248360e11b6001600160e01b03198316145b8061078e575061078e826135db565b7f2569078dfb4b0305704d3008e7403993ae9601b85f7ae5e742de3de8f8011c4090565b600081600111158015612b535750612b4f612b15565b5482105b801561078e5750600160e01b612b67612b15565b60008481526004919091016020526040902054161592915050565b6daaeb6d7670e522a718067333cd4e3b15610dac57604051633185c44d60e21b81523060048201526001600160a01b03821660248201526daaeb6d7670e522a718067333cd4e9063c617113490604401602060405180830381865afa158015612bef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c1391906152ab565b610dac57604051633b79c77360e21b81526001600160a01b0382166004820152602401610e40565b6000612c468261128f565b9050336001600160a01b03821614612c7f57612c628133612787565b612c7f576040516367d9dca160e11b815260040160405180910390fd5b82612c88612b15565b6000848152600691909101602052604080822080546001600160a01b0319166001600160a01b0394851617905551849286811692908516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259190a4505050565b303314612d0e612cf761173d565b6001600160a01b0316336001600160a01b03161490565b1760000361144757604051635fc483c560e01b815260040160405180910390fd5b60006001612d3b612b15565b5403919050565b7fb3623c06c2ed11908644eb46053665cc2b67a32ab7b445094be7495b2530a9d390565b6001600160a01b0381166000908152600080516020615491833981519152602052604090205460ff161515600114610dac576040516315e26ff360e01b815260040160405180910390fd5b337ff73863e3917595a7e67829ed60f0c51bf14c7a0d0de47c7b02a00abd48488e4a546001600160a01b03161461144757604051635fc483c560e01b815260040160405180910390fd5b6000612e068261318e565b9050836001600160a01b0316816001600160a01b031614612e395760405162a1148160e81b815260040160405180910390fd5b600080612e4584613629565b91509150612e6a8187612e553390565b6001600160a01b039081169116811491141790565b612e9557612e788633612787565b612e9557604051632ce44b5f60e11b815260040160405180910390fd5b6001600160a01b038516612ebc57604051633a954ecd60e21b815260040160405180910390fd5b8015612ec757600082555b612ecf612b15565b6001600160a01b0387166000908152600591909101602052604090208054600019019055612efb612b15565b6001600160a01b03861660008181526005929092016020526040909120805460010190554260a01b17600160e11b17612f32612b15565b60008681526004919091016020526040812091909155600160e11b84169003612fa85760018401612f61612b15565b600082815260049190910160205260408120549003612fa657612f82612b15565b548114612fa65783612f92612b15565b600083815260049190910160205260409020555b505b83856001600160a01b0316876001600160a01b03166000805160206154b183398151915260405160405180910390a4610cf9565b61089083838360405180602001604052806000815250612583565b610dac816000613651565b600061274f612b15565b7ff268be8736a07172c20cb8afb46ffa17fa1131bf48395e58d9c0ce565c5047f4548160005b828110156130ac5760006000805160206154918339815191526000600080516020615491833981519152600101848154811061307057613070614fc5565b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff1916911515919091179055600101613032565b5060005b8181101561312257600160008051602061549183398151915260008787858181106130dd576130dd614fc5565b90506020020160208101906130f29190613f5a565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790556001016130b0565b5061314e7ff268be8736a07172c20cb8afb46ffa17fa1131bf48395e58d9c0ce565c5047f48585613c22565b507fbbd3b69c138de4d317d0bc4290282c4e1cbd1e58b579a5b4f114b598c237454d84846040516131809291906152c8565b60405180910390a150505050565b60008180600111613209576131a1612b15565b548110156132095760006131b3612b15565b600083815260049190910160205260408120549150600160e01b82169003613207575b806000036126a6576131e6612b15565b600019909201600081815260049390930160205260409092205490506131d6565b505b604051636f96cda160e11b815260040160405180910390fd5b600061322c612b15565b54905060008290036132515760405163b562e8dd60e01b815260040160405180910390fd5b680100000000000000018202613265612b15565b6001600160a01b038516600081815260059290920160205260409091208054929092019091554260a01b6001841460e11b17176132a0612b15565b600083815260049190910160205260408120919091556001600160a01b0384169083830190839083906000805160206154b18339815191528180a4600183015b81811461330657808360006000805160206154b1833981519152600080a46001016132e0565b508160000361332757604051622e076360e81b815260040160405180910390fd5b80613330612b15565b55506108909050565b60606000805160206154d183398151915260010180546107ac9061459c565b7ff73863e3917595a7e67829ed60f0c51bf14c7a0d0de47c7b02a00abd48488e4a80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60006001600160401b0360406133dd612b15565b6005016000856001600160a01b03166001600160a01b0316815260200190815260200160002054901c169050919050565b80613417612b15565b336000818152600792909201602090815260408084206001600160a01b03881680865290835293819020805460ff19169515159590951790945592518415158152919290917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b613496848484610bc7565b6001600160a01b0383163b15610bec576134b2848484846137c0565b610bec576040516368d2bf6b60e11b815260040160405180910390fd5b606060a06040510180604052602081039150506000815280825b600183039250600a81066030018353600a9004806134e95750819003601f19909101908152919050565b8061351c612d42565b80546001600160a01b0319166001600160a01b039283161790556040519082169033907f222c8e95a03c7aa95fc5d110485e0d38e767f07ab1ec878a6eac644ef1d8312290600090a350565b600054610100900460ff1661358f5760405162461bcd60e51b8152600401610e4090615316565b61359984846138ac565b6135a161396c565b6135a9613993565b6135b384846139c2565b6135bb6139e9565b6135c6848483613a36565b6135cf82613b2f565b610bec84848484613b5f565b60006301ffc9a760e01b6001600160e01b03198316148061360c57506380ac58cd60e01b6001600160e01b03198316145b8061078e5750506001600160e01b031916635b5e139f60e01b1490565b6000806000613636612b15565b60009485526006016020525050604090912080549092909150565b600061365c8361318e565b90508060008061366b86613629565b9150915084156136ab57613680818433612e55565b6136ab5761368e8333612787565b6136ab57604051632ce44b5f60e11b815260040160405180910390fd5b80156136b657600082555b6fffffffffffffffffffffffffffffffff6136cf612b15565b6001600160a01b038516600081815260059290920160205260409091208054929092019091554260a01b17600360e01b17613708612b15565b60008881526004919091016020526040812091909155600160e11b8516900361377e5760018601613737612b15565b60008281526004919091016020526040812054900361377c57613758612b15565b54811461377c5784613768612b15565b600083815260049190910160205260409020555b505b60405186906000906001600160a01b038616906000805160206154b1833981519152908390a46137ac612b15565b600190810180549091019055505050505050565b604051630a85bd0160e11b81526000906001600160a01b0385169063150b7a02906137f5903390899088908890600401615361565b6020604051808303816000875af1925050508015613830575060408051601f3d908101601f1916820190925261382d91810190615394565b60015b61388e573d80801561385e576040519150601f19603f3d011682016040523d82523d6000602084013e613863565b606091505b508051600003613886576040516368d2bf6b60e11b815260040160405180910390fd5b805181602001fd5b6001600160e01b031916630a85bd0160e11b1490505b949350505050565b60008051602061551183398151915254610100900460ff1661392d5760405162461bcd60e51b815260206004820152603460248201527f455243373231415f5f496e697469616c697a61626c653a20636f6e7472616374604482015273206973206e6f7420696e697469616c697a696e6760601b6064820152608401610e40565b81613936612b15565b6002019061394490826153b1565b508061394e612b15565b6003019061395c90826153b1565b506001613967612b15565b555050565b600054610100900460ff166114475760405162461bcd60e51b8152600401610e4090615316565b600054610100900460ff166139ba5760405162461bcd60e51b8152600401610e4090615316565b611447613b86565b600054610100900460ff166112305760405162461bcd60e51b8152600401610e4090615316565b600054610100900460ff16613a105760405162461bcd60e51b8152600401610e4090615316565b60017fd59f8a8c0d1463371c77782499276e5cbe466fd192ada543ceaea0a36604c1f255565b600054610100900460ff16613a5d5760405162461bcd60e51b8152600401610e4090615316565b805160005b81811015613acc5760016000805160206154918339815191526000016000858481518110613a9257613a92614fc5565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055600101613a62565b508151613aff907ff268be8736a07172c20cb8afb46ffa17fa1131bf48395e58d9c0ce565c5047f4906020850190613c85565b506040517fd7aca75208b9be5ffc04c6a01922020ffd62b55e68e502e317f5344960279af890600090a150505050565b600054610100900460ff16613b565760405162461bcd60e51b8152600401610e4090615316565b610dac81613bae565b600054610100900460ff16610bec5760405162461bcd60e51b8152600401610e4090615316565b303b15613ba55760405162dc149f60e41b815260040160405180910390fd5b61144733613358565b303b15613bcd5760405162dc149f60e41b815260040160405180910390fd5b80613bd6612d42565b80546001600160a01b0319166001600160a01b03928316179055604051908216906000907f222c8e95a03c7aa95fc5d110485e0d38e767f07ab1ec878a6eac644ef1d83122908290a350565b828054828255906000526020600020908101928215613c75579160200282015b82811115613c755781546001600160a01b0319166001600160a01b03843516178255602090920191600190910190613c42565b50613c81929150613cda565b5090565b828054828255906000526020600020908101928215613c75579160200282015b82811115613c7557825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613ca5565b5b80821115613c815760008155600101613cdb565b6001600160e01b031981168114610dac57600080fd5b600060208284031215613d1757600080fd5b81356126a681613cef565b60005b83811015613d3d578181015183820152602001613d25565b50506000910152565b60008151808452613d5e816020860160208601613d22565b601f01601f19169290920160200192915050565b6020815260006126a66020830184613d46565b600060208284031215613d9757600080fd5b5035919050565b6001600160a01b0381168114610dac57600080fd5b60008060408385031215613dc657600080fd5b8235613dd181613d9e565b946020939093013593505050565b60008082840360e0811215613df357600080fd5b8335613dfe81613d9e565b925060c0601f1982011215613e1257600080fd5b506020830190509250929050565b600080600060608486031215613e3557600080fd5b8335613e4081613d9e565b92506020840135613e5081613d9e565b929592945050506040919091013590565b60008060408385031215613e7457600080fd5b50508035926020909101359150565b60008060408385031215613e9657600080fd5b8235613ea181613d9e565b915060208301356001600160401b03811115613ebc57600080fd5b830160608186031215613ece57600080fd5b809150509250929050565b600060408284031215613eeb57600080fd5b50919050565b8015158114610dac57600080fd5b8035613f0a81613ef1565b919050565b600080600060608486031215613f2457600080fd5b8335613f2f81613d9e565b92506020840135613f3f81613d9e565b91506040840135613f4f81613ef1565b809150509250925092565b600060208284031215613f6c57600080fd5b81356126a681613d9e565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b0381118282101715613faf57613faf613f77565b60405290565b60405160c081016001600160401b0381118282101715613faf57613faf613f77565b60405161010081016001600160401b0381118282101715613faf57613faf613f77565b604051601f8201601f191681016001600160401b038111828210171561402257614022613f77565b604052919050565b6001600160501b0381168114610dac57600080fd5b62ffffff81168114610dac57600080fd5b64ffffffffff81168114610dac57600080fd5b8035613f0a81614050565b61ffff81168114610dac57600080fd5b8035613f0a8161406e565b60008060008385036101208112156140a057600080fd5b84356140ab81613d9e565b935060208501356140bb81613d9e565b925060e0603f19820112156140cf57600080fd5b506140d8613f8d565b60408501356140e68161402a565b815260608501356140f68161403f565b6020820152608085013561410981614050565b604082015260a085013561411c81614050565b606082015261412d60c08601614063565b608082015261413e60e0860161407e565b60a0820152614150610100860161407e565b60c0820152809150509250925092565b60008083601f84011261417257600080fd5b5081356001600160401b0381111561418957600080fd5b6020830191508360208285010111156141a157600080fd5b9250929050565b600080602083850312156141bb57600080fd5b82356001600160401b038111156141d157600080fd5b6141dd85828601614160565b90969095509350505050565b600080602083850312156141fc57600080fd5b82356001600160401b038082111561421357600080fd5b818501915085601f83011261422757600080fd5b81358181111561423657600080fd5b8660208260051b850101111561424b57600080fd5b60209290920196919550909350505050565b6000806040838503121561427057600080fd5b823561427b81613d9e565b91506020830135613ece81613d9e565b6000806000604084860312156142a057600080fd5b83356142ab81613d9e565b925060208401356001600160401b038111156142c657600080fd5b6142d286828701614160565b9497909650939450505050565b60008060008385036101408112156142f657600080fd5b843561430181613d9e565b9350602085013561431181613d9e565b9250610100603f198201121561432657600080fd5b506040840190509250925092565b60006020828403121561434657600080fd5b81356001600160401b0381111561435c57600080fd5b820161030081850312156126a657600080fd5b6000806040838503121561438257600080fd5b823561438d81613d9e565b91506020830135613ece81613ef1565b60006001600160401b038311156143b6576143b6613f77565b6143c9601f8401601f1916602001613ffa565b90508281528383830111156143dd57600080fd5b828260208301376000602084830101529392505050565b6000806000806080858703121561440a57600080fd5b843561441581613d9e565b9350602085013561442581613d9e565b92506040850135915060608501356001600160401b0381111561444757600080fd5b8501601f8101871361445857600080fd5b6144678782356020840161439d565b91505092959194509250565b600082601f83011261448457600080fd5b6126a68383356020850161439d565b600080600080608085870312156144a957600080fd5b84356001600160401b03808211156144c057600080fd5b6144cc88838901614473565b95506020915081870135818111156144e357600080fd5b6144ef89828a01614473565b955050604087013561450081613d9e565b935060608701358181111561451457600080fd5b8701601f8101891361452557600080fd5b80358281111561453757614537613f77565b8060051b9250614548848401613ffa565b818152928201840192848101908b85111561456257600080fd5b928501925b8484101561458c578335925061457c83613d9e565b8282529285019290850190614567565b989b979a50959850505050505050565b600181811c908216806145b057607f821691505b602082108103613eeb57634e487b7160e01b600052602260045260246000fd5b65ffffffffffff81168114610dac57600080fd5b8051613f0a816145d0565b8051613f0a8161406e565b8051613f0a81613ef1565b600060c0828403121561461757600080fd5b61461f613fb5565b825161462a8161402a565b8152602083015161463a816145d0565b6020820152604083015161464d816145d0565b604082015260608301516146608161406e565b606082015260808301516146738161406e565b608082015260a083015161468681613ef1565b60a08201529392505050565b8035613f0a816145d0565b600060c082840312156146af57600080fd5b6146b7613fb5565b82356146c28161402a565b815260208301356146d2816145d0565b602082015260408301356146e5816145d0565b604082015260608301356146f88161406e565b6060820152608083013561470b8161406e565b608082015260a083013561468681613ef1565b634e487b7160e01b600052601160045260246000fd5b808202811582820484141761078e5761078e61471e565b60008261476857634e487b7160e01b600052601260045260246000fd5b500490565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000808335601e198436030181126147ad57600080fd5b83016020810192503590506001600160401b038111156147cc57600080fd5b8036038213156141a157600080fd5b60006060830182358452602080840135601e198536030181126147fd57600080fd5b840181810190356001600160401b0381111561481857600080fd5b8060051b80360383131561482b57600080fd5b6060848901529381905260809387018401938290880160005b8381101561487e57898703607f1901825261485f8386614796565b61486a89828461476d565b985050509185019190850190600101614844565b5050505050506148916040840184614796565b85830360408701526148a483828461476d565b9695505050505050565b6020815260006126a660208301846147db565b6001600160601b0381168114610dac57600080fd5b6000602082840312156148e857600080fd5b81356126a6816148c1565b81356148fe81613d9e565b81546001600160a01b03199081166001600160a01b03929092169182178355602084013561492b816148c1565b60a01b1617905550565b8051613f0a81614050565b600060e0828403121561495257600080fd5b60405160e081018181106001600160401b038211171561497457614974613f77565b60405282516149828161402a565b815260208301516149928161403f565b602082015260408301516149a581614050565b604082015260608301516149b881614050565b60608201526149c960808401614935565b60808201526149da60a084016145ef565b60a08201526149eb60c084016145ef565b60c08201529392505050565b6001600160501b03815116825262ffffff6020820151166020830152604081015164ffffffffff8082166040850152806060840151166060850152806080840151166080850152505060a081015161ffff80821660a08501528060c08401511660c085015250505050565b6001600160a01b038316815261010081016126a660208301846149f7565b601f82111561089057600081815260208120601f850160051c81016020861015614aa75750805b601f850160051c820191505b81811015610cf957828155600101614ab3565b6001600160401b03831115614add57614add613f77565b614af183614aeb835461459c565b83614a80565b6000601f841160018114614b255760008515614b0d5750838201355b600019600387901b1c1916600186901b1783556125a9565b600083815260209020601f19861690835b82811015614b565786850135825560209485019460019092019101614b36565b5086821015614b735760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b8181038181111561078e5761078e61471e565b8082018082111561078e5761078e61471e565b6020815260006138a460208301848661476d565b60ff81168114610dac57600080fd5b8051613f0a81614bbf565b63ffffffff81168114610dac57600080fd5b8051613f0a81614bd9565b60006101008284031215614c0957600080fd5b614c11613fd7565b8251614c1c8161402a565b81526020830151614c2c8161406e565b6020820152614c3d604084016145e4565b6040820152614c4e606084016145e4565b6060820152614c5f60808401614bce565b6080820152614c7060a08401614beb565b60a0820152614c8160c084016145ef565b60c0820152614c9260e084016145fa565b60e08201529392505050565b8035613f0a81614bbf565b8035613f0a81614bd9565b60006101008284031215614cc757600080fd5b614ccf613fd7565b8235614cda8161402a565b81526020830135614cea8161406e565b6020820152614cfb60408401614692565b6040820152614d0c60608401614692565b6060820152614d1d60808401614c9e565b6080820152614d2e60a08401614ca9565b60a0820152614d3f60c0840161407e565b60c0820152614c9260e08401613eff565b6001600160501b03815116825261ffff6020820151166020830152604081015165ffffffffffff8082166040850152806060840151166060850152505060ff608082015116608083015260a0810151614db160a084018263ffffffff169052565b5060c0810151614dc760c084018261ffff169052565b5060e081015161089060e084018215159052565b6001600160a01b038316815261012081016126a66020830184614d50565b6000808335601e19843603018112614e1057600080fd5b8301803591506001600160401b03821115614e2a57600080fd5b6020019150368190038213156141a157600080fd5b600060208284031215614e5157600080fd5b81356126a6816145d0565b6001600160a01b038316815260e081018235614e778161402a565b6001600160501b0381166020840152506020830135614e95816145d0565b65ffffffffffff808216604085015260408501359150614eb4826145d0565b16606083810191909152830135614eca8161406e565b61ffff81166080840152506080830135614ee38161406e565b61ffff811660a084015250614efa60a08401613eff565b80151560c0840152509392505050565b6001600160a01b0384168152604060208201819052600090614f2f908301848661476d565b95945050505050565b60008235605e19833603018112614f4e57600080fd5b9190910192915050565b6001600160a01b03831681526040602082018190526000906138a4908301846147db565b6000808335601e19843603018112614f9357600080fd5b8301803591506001600160401b03821115614fad57600080fd5b6020019150600581901b36038213156141a157600080fd5b634e487b7160e01b600052603260045260246000fd5b6001600160a01b039384168152919092166020820152901515604082015260600190565b6000808335601e1984360301811261501657600080fd5b8301803591506001600160401b0382111561503057600080fd5b6020019150600881901b36038213156141a157600080fd5b6001600160a01b038481168252831660208201526101408101823561506c8161402a565b6001600160501b0316604083015260208301356150888161406e565b61ffff16606083015261509d60408401614692565b65ffffffffffff1660808301526150b660608401614692565b65ffffffffffff1660a08301526150cf60808401614c9e565b60ff1660c08301526150e360a08401614ca9565b63ffffffff1660e08301526150fa60c0840161407e565b61ffff1661010083015261511060e08401613eff565b8015156101208401525b50949350505050565b6001600160a01b0384811682528316602082015261014081016138a46040830184614d50565b6000808335601e1984360301811261516057600080fd5b8301803591506001600160401b0382111561517a57600080fd5b602001915060e0810236038213156141a157600080fd5b6001600160a01b03848116825283166020820152610120810182356151b58161402a565b6001600160501b0316604083015260208301356151d18161403f565b62ffffff16606083015260408301356151e981614050565b64ffffffffff16608083015261520160608401614063565b64ffffffffff1660a083015261521960808401614063565b64ffffffffff1660c083015261523160a0840161407e565b61ffff1660e083015261524660c0840161407e565b61ffff811661010084015261511a565b6001600160a01b0384811682528316602082015261012081016138a460408301846149f7565b6000835161528e818460208801613d22565b8351908301906152a2818360208801613d22565b01949350505050565b6000602082840312156152bd57600080fd5b81516126a681613ef1565b60208082528181018390526000908460408401835b8681101561530b5782356152f081613d9e565b6001600160a01b0316825291830191908301906001016152dd565b509695505050505050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906148a490830184613d46565b6000602082840312156153a657600080fd5b81516126a681613cef565b81516001600160401b038111156153ca576153ca613f77565b6153de816153d8845461459c565b84614a80565b602080601f83116001811461541357600084156153fb5750858301515b600019600386901b1c1916600185901b178555610cf9565b600085815260208120601f198616915b8281101561544257888601518255948401946001909101908401615423565b50858210156154605787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fef73863e3917595a7e67829ed60f0c51bf14c7a0d0de47c7b02a00abd48488e4bf268be8736a07172c20cb8afb46ffa17fa1131bf48395e58d9c0ce565c5047f3ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efb847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a40319b847c145db4703dd7415562b7f4e3aaa5b8cbca80b6c5135ed89cd14e7a4031dee151c8401928dc223602bb187aff91b9a56c7cae5476ef1b3287b085a16c85fa2646970667358221220a05c0abc4c3e38e5029251abdccc07a04e42536e1b12dea069d09b6fed87730164736f6c63430008110033",
}

// PartnerABI is the input ABI used to generate the binding from.
// Deprecated: Use PartnerMetaData.ABI instead.
var PartnerABI = PartnerMetaData.ABI

// PartnerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PartnerMetaData.Bin instead.
var PartnerBin = PartnerMetaData.Bin

// DeployPartner deploys a new Ethereum contract, binding an instance of Partner to it.
func DeployPartner(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Partner, error) {
	parsed, err := PartnerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PartnerBin), backend)
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

// GetBurnAddress is a free data retrieval call binding the contract method 0x38b39d29.
//
// Solidity: function getBurnAddress() view returns(address)
func (_Partner *PartnerCaller) GetBurnAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Partner.contract.Call(opts, &out, "getBurnAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBurnAddress is a free data retrieval call binding the contract method 0x38b39d29.
//
// Solidity: function getBurnAddress() view returns(address)
func (_Partner *PartnerSession) GetBurnAddress() (common.Address, error) {
	return _Partner.Contract.GetBurnAddress(&_Partner.CallOpts)
}

// GetBurnAddress is a free data retrieval call binding the contract method 0x38b39d29.
//
// Solidity: function getBurnAddress() view returns(address)
func (_Partner *PartnerCallerSession) GetBurnAddress() (common.Address, error) {
	return _Partner.Contract.GetBurnAddress(&_Partner.CallOpts)
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

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Partner *PartnerTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Partner *PartnerSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Partner.Contract.Burn(&_Partner.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Partner *PartnerTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Partner.Contract.Burn(&_Partner.TransactOpts, tokenId)
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

// Initialize is a paid mutator transaction binding the contract method 0xfbc46ada.
//
// Solidity: function initialize(string name, string symbol, address administrator, address[] allowedSeaDrop) returns()
func (_Partner *PartnerTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, administrator common.Address, allowedSeaDrop []common.Address) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "initialize", name, symbol, administrator, allowedSeaDrop)
}

// Initialize is a paid mutator transaction binding the contract method 0xfbc46ada.
//
// Solidity: function initialize(string name, string symbol, address administrator, address[] allowedSeaDrop) returns()
func (_Partner *PartnerSession) Initialize(name string, symbol string, administrator common.Address, allowedSeaDrop []common.Address) (*types.Transaction, error) {
	return _Partner.Contract.Initialize(&_Partner.TransactOpts, name, symbol, administrator, allowedSeaDrop)
}

// Initialize is a paid mutator transaction binding the contract method 0xfbc46ada.
//
// Solidity: function initialize(string name, string symbol, address administrator, address[] allowedSeaDrop) returns()
func (_Partner *PartnerTransactorSession) Initialize(name string, symbol string, administrator common.Address, allowedSeaDrop []common.Address) (*types.Transaction, error) {
	return _Partner.Contract.Initialize(&_Partner.TransactOpts, name, symbol, administrator, allowedSeaDrop)
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
func (_Partner *PartnerTransactor) MultiConfigure(opts *bind.TransactOpts, config ERC721SeaDropStructsErrorsAndEventsUpgradeableMultiConfigureStruct) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "multiConfigure", config)
}

// MultiConfigure is a paid mutator transaction binding the contract method 0x911f456b.
//
// Solidity: function multiConfigure((uint256,string,string,address,(uint80,uint48,uint48,uint16,uint16,bool),string,(bytes32,string[],string),address,bytes32,address[],address[],address[],address[],address[],(uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool)[],address[],address[],(uint80,uint24,uint40,uint40,uint40,uint16,uint16)[],address[]) config) returns()
func (_Partner *PartnerSession) MultiConfigure(config ERC721SeaDropStructsErrorsAndEventsUpgradeableMultiConfigureStruct) (*types.Transaction, error) {
	return _Partner.Contract.MultiConfigure(&_Partner.TransactOpts, config)
}

// MultiConfigure is a paid mutator transaction binding the contract method 0x911f456b.
//
// Solidity: function multiConfigure((uint256,string,string,address,(uint80,uint48,uint48,uint16,uint16,bool),string,(bytes32,string[],string),address,bytes32,address[],address[],address[],address[],address[],(uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool)[],address[],address[],(uint80,uint24,uint40,uint40,uint40,uint16,uint16)[],address[]) config) returns()
func (_Partner *PartnerTransactorSession) MultiConfigure(config ERC721SeaDropStructsErrorsAndEventsUpgradeableMultiConfigureStruct) (*types.Transaction, error) {
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

// SetBurnAddress is a paid mutator transaction binding the contract method 0x4b0e7216.
//
// Solidity: function setBurnAddress(address newBurnAddress) returns()
func (_Partner *PartnerTransactor) SetBurnAddress(opts *bind.TransactOpts, newBurnAddress common.Address) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "setBurnAddress", newBurnAddress)
}

// SetBurnAddress is a paid mutator transaction binding the contract method 0x4b0e7216.
//
// Solidity: function setBurnAddress(address newBurnAddress) returns()
func (_Partner *PartnerSession) SetBurnAddress(newBurnAddress common.Address) (*types.Transaction, error) {
	return _Partner.Contract.SetBurnAddress(&_Partner.TransactOpts, newBurnAddress)
}

// SetBurnAddress is a paid mutator transaction binding the contract method 0x4b0e7216.
//
// Solidity: function setBurnAddress(address newBurnAddress) returns()
func (_Partner *PartnerTransactorSession) SetBurnAddress(newBurnAddress common.Address) (*types.Transaction, error) {
	return _Partner.Contract.SetBurnAddress(&_Partner.TransactOpts, newBurnAddress)
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
func (_Partner *PartnerTransactor) SetRoyaltyInfo(opts *bind.TransactOpts, newInfo ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo) (*types.Transaction, error) {
	return _Partner.contract.Transact(opts, "setRoyaltyInfo", newInfo)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_Partner *PartnerSession) SetRoyaltyInfo(newInfo ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo) (*types.Transaction, error) {
	return _Partner.Contract.SetRoyaltyInfo(&_Partner.TransactOpts, newInfo)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_Partner *PartnerTransactorSession) SetRoyaltyInfo(newInfo ISeaDropTokenContractMetadataUpgradeableRoyaltyInfo) (*types.Transaction, error) {
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

// PartnerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Partner contract.
type PartnerInitializedIterator struct {
	Event *PartnerInitialized // Event containing the contract specifics and raw log

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
func (it *PartnerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartnerInitialized)
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
		it.Event = new(PartnerInitialized)
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
func (it *PartnerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartnerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartnerInitialized represents a Initialized event raised by the Partner contract.
type PartnerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Partner *PartnerFilterer) FilterInitialized(opts *bind.FilterOpts) (*PartnerInitializedIterator, error) {

	logs, sub, err := _Partner.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PartnerInitializedIterator{contract: _Partner.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Partner *PartnerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PartnerInitialized) (event.Subscription, error) {

	logs, sub, err := _Partner.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartnerInitialized)
				if err := _Partner.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Partner *PartnerFilterer) ParseInitialized(log types.Log) (*PartnerInitialized, error) {
	event := new(PartnerInitialized)
	if err := _Partner.contract.UnpackLog(event, "Initialized", log); err != nil {
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
