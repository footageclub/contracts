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

// TokenMetaData contains all meta data concerning the Token contract.
var TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"allowedSeaDrop\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"CannotExceedMaxSupplyOfUint64\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"basisPoints\",\"type\":\"uint256\"}],\"name\":\"InvalidRoyaltyBasisPoints\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotNextOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedSeaDrop\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProvenanceHashCannotBeSetAfterMintStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoyaltyAddressCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignersMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowedSeaDrop\",\"type\":\"address[]\"}],\"name\":\"AllowedSeaDropUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"}],\"name\":\"ContractURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPotentialAdministrator\",\"type\":\"address\"}],\"name\":\"PotentialOwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newHash\",\"type\":\"bytes32\"}],\"name\":\"ProvenanceHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"name\":\"RoyaltyInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SeaDropTokenDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OPERATOR_FILTER_REGISTRY\",\"outputs\":[{\"internalType\":\"contractIOperatorFilterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"}],\"name\":\"emitBatchMetadataUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"getMintStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minterNumMinted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"mintSeaDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contractURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"dropURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"publicKeyURIs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"internalType\":\"structAllowListData\",\"name\":\"allowListData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"creatorPayoutAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"provenanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"allowedFeeRecipients\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedFeeRecipients\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowedPayers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedPayers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenGatedAllowedNftTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structTokenGatedDropStage[]\",\"name\":\"tokenGatedDropStages\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedTokenGatedAllowedNftTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"internalType\":\"structSignedMintValidationParams[]\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"disallowedSigners\",\"type\":\"address[]\"}],\"internalType\":\"structERC721SeaDropStructsErrorsAndEvents.MultiConfigureStruct\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"multiConfigure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provenanceHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newProvenanceHash\",\"type\":\"bytes32\"}],\"name\":\"setProvenanceHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"royaltyAddress\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"royaltyBps\",\"type\":\"uint96\"}],\"internalType\":\"structISeaDropTokenContractMetadata.RoyaltyInfo\",\"name\":\"newInfo\",\"type\":\"tuple\"}],\"name\":\"setRoyaltyInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPotentialOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"publicKeyURIs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"internalType\":\"structAllowListData\",\"name\":\"allowListData\",\"type\":\"tuple\"}],\"name\":\"updateAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateAllowedFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"allowedSeaDrop\",\"type\":\"address[]\"}],\"name\":\"updateAllowedSeaDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payoutAddress\",\"type\":\"address\"}],\"name\":\"updateCreatorPayoutAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"dropURI\",\"type\":\"string\"}],\"name\":\"updateDropURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updatePayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"}],\"name\":\"updatePublicDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"internalType\":\"structSignedMintValidationParams\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple\"}],\"name\":\"updateSignedMintValidationParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seaDropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structTokenGatedDropStage\",\"name\":\"dropStage\",\"type\":\"tuple\"}],\"name\":\"updateTokenGatedDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526001600f553480156200001657600080fd5b5060405162003fea38038062003fea833981016040819052620000399162000450565b733cc6cdda760b79bafa08df41ecfa224f810dceb66001848481816002620000628382620005f6565b506003620000718282620005f6565b5050600160005550620000836200027a565b50506daaeb6d7670e522a718067333cd4e3b15620001ca5780156200011857604051633e9f1edf60e11b81523060048201526001600160a01b03831660248201526daaeb6d7670e522a718067333cd4e90637d3e3dbe906044015b600060405180830381600087803b158015620000f957600080fd5b505af11580156200010e573d6000803e3d6000fd5b50505050620001ca565b6001600160a01b03821615620001695760405163a0af290360e01b81523060048201526001600160a01b03831660248201526daaeb6d7670e522a718067333cd4e9063a0af290390604401620000de565b604051632210724360e11b81523060048201526daaeb6d7670e522a718067333cd4e90634420e48690602401600060405180830381600087803b158015620001b057600080fd5b505af1158015620001c5573d6000803e3d6000fd5b505050505b5050805160005b818110156200023057600160106000858481518110620001f557620001f5620006c2565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055600101620001d1565b50815162000246906011906020850190620002f9565b506040517fd7aca75208b9be5ffc04c6a01922020ffd62b55e68e502e317f5344960279af890600090a150505050620006d8565b303b156200029a5760405162dc149f60e41b815260040160405180910390fd5b620002a533620002a7565b565b600880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b82805482825590600052602060002090810192821562000351579160200282015b828111156200035157825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906200031a565b506200035f92915062000363565b5090565b5b808211156200035f576000815560010162000364565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620003bb57620003bb6200037a565b604052919050565b600082601f830112620003d557600080fd5b81516001600160401b03811115620003f157620003f16200037a565b602062000407601f8301601f1916820162000390565b82815285828487010111156200041c57600080fd5b60005b838110156200043c5785810183015182820184015282016200041f565b506000928101909101919091529392505050565b6000806000606084860312156200046657600080fd5b83516001600160401b03808211156200047e57600080fd5b6200048c87838801620003c3565b9450602091508186015181811115620004a457600080fd5b620004b288828901620003c3565b945050604086015181811115620004c857600080fd5b8601601f81018813620004da57600080fd5b805182811115620004ef57620004ef6200037a565b8060051b92506200050284840162000390565b818152928201840192848101908a8511156200051d57600080fd5b928501925b848410156200055757835192506001600160a01b0383168314620005465760008081fd5b828252928501929085019062000522565b8096505050505050509250925092565b600181811c908216806200057c57607f821691505b6020821081036200059d57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620005f157600081815260208120601f850160051c81016020861015620005cc5750805b601f850160051c820191505b81811015620005ed57828155600101620005d8565b5050505b505050565b81516001600160401b038111156200061257620006126200037a565b6200062a8162000623845462000567565b84620005a3565b602080601f831160018114620006625760008415620006495750858301515b600019600386901b1c1916600185901b178555620005ed565b600085815260208120601f198616915b82811015620006935788860151825594840194600190910190840162000672565b5085821015620006b25787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b61390280620006e86000396000f3fe608060405234801561001057600080fd5b506004361061027f5760003560e01c80636c0360eb1161015c57806395d89b41116100ce578063c87b56dd11610087578063c87b56dd1461059c578063cb743ba8146105af578063d5abeb01146105c2578063e8a3d485146105ca578063e985e9c5146105d2578063f2fde38b146105e557600080fd5b806395d89b4114610542578063a22cb4651461054a578063a48301141461055d578063ad2f852a14610570578063b88d4fde14610581578063c6ab67a31461059457600080fd5b80637a05bc82116101205780637a05bc82146104b75780637bc2be76146104ca578063840e15d4146104dd5780638da5cb5b1461050b578063911f456b1461051c578063938e3d7b1461052f57600080fd5b80636c0360eb146104795780636f8b44b01461048157806370a0823114610494578063715018a6146104a757806379ba5097146104af57600080fd5b806341f43434116101f5578063511aa644116101b9578063511aa6441461040757806355f804b31461041a57806360c308b61461042d5780636352211e1461044057806364869dad1461045357806366251b691461046657600080fd5b806341f43434146103a157806342260b5d146103b657806342842e0e146103ce57806344dae42c146103e157806348a4c101146103f457600080fd5b806318160ddd1161024757806318160ddd146103145780631b73593c1461032e57806323452b9c1461034157806323b872dd146103495780632a55205a1461035c5780633680620d1461038e57600080fd5b806301ffc9a71461028457806306fdde03146102ac578063081812fc146102c1578063095ea7b3146102ec578063099b6bfa14610301575b600080fd5b610297610292366004612779565b6105f8565b60405190151581526020015b60405180910390f35b6102b461063e565b6040516102a391906127e6565b6102d46102cf3660046127f9565b6106d0565b6040516001600160a01b0390911681526020016102a3565b6102ff6102fa366004612827565b610714565b005b6102ff61030f3660046127f9565b61072d565b60015460005403600019015b6040519081526020016102a3565b6102ff61033c366004612853565b6107a0565b6102ff610813565b6102ff610357366004612894565b610861565b61036f61036a3660046128d5565b61088c565b604080516001600160a01b0390931683526020830191909152016102a3565b6102ff61039c3660046128f7565b6108d2565b6102d46daaeb6d7670e522a718067333cd4e81565b600e54600160a01b90046001600160601b0316610320565b6102ff6103dc366004612894565b61090f565b6102ff6103ef36600461294d565b610934565b6102ff610402366004612983565b610a4a565b6102ff610415366004612a8d565b610ac4565b6102ff610428366004612ba4565b610b03565b6102ff61043b366004612be5565b610b79565b6102d461044e3660046127f9565b610b8b565b6102ff610461366004612827565b610b96565b6102ff610474366004612c59565b610c54565b6102b4610c93565b6102ff61048f3660046127f9565b610ca2565b6103206104a2366004612c87565b610d0a565b6102ff610d58565b6102ff610d6c565b6102ff6104c5366004612ca4565b610de8565b6102ff6104d8366004612cf8565b610e27565b6104f06104eb366004612c87565b610e66565b604080519384526020840192909252908201526060016102a3565b6008546001600160a01b03166102d4565b6102ff61052a366004612d4d565b610ea7565b6102ff61053d366004612ba4565b611ac4565b6102b4611b0b565b6102ff610558366004612d88565b611b1a565b6102ff61056b3660046128d5565b611b2e565b600e546001600160a01b03166102d4565b6102ff61058f366004612db6565b611b6c565b600d54610320565b6102b46105aa3660046127f9565b611b99565b6102ff6105bd366004612983565b611c1d565b600a54610320565b6102b4611c64565b6102976105e0366004612c59565b611c73565b6102ff6105f3366004612c87565b611ca1565b60006001600160e01b03198216630c487f4760e11b148061062957506001600160e01b03198216639c15441560e01b145b80610638575061063882611d1e565b92915050565b60606002805461064d90612e79565b80601f016020809104026020016040519081016040528092919081815260200182805461067990612e79565b80156106c65780601f1061069b576101008083540402835291602001916106c6565b820191906000526020600020905b8154815290600101906020018083116106a957829003601f168201915b5050505050905090565b60006106db82611d5e565b6106f8576040516333d1c03960e21b815260040160405180910390fd5b506000908152600660205260409020546001600160a01b031690565b8161071e81611d93565b6107288383611e4c565b505050565b610735611eec565b600054600019011561075a5760405163e03264af60e01b815260040160405180910390fd5b600d80549082905560408051828152602081018490527f7c22004198bf87da0f0dab623c72e66ca1200f4454aa3b9ca30f436275428b7c91015b60405180910390a15050565b6107a8611eec565b6107b182611f3c565b6040516301308e6560e01b81526001600160a01b038316906301308e65906107dd908490600401612f4e565b600060405180830381600087803b1580156107f757600080fd5b505af115801561080b573d6000803e3d6000fd5b505050505050565b61081b611f7a565b600980546001600160a01b0319169055604051600081527f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da9060200160405180910390a1565b826001600160a01b038116331461087b5761087b33611d93565b610886848484611fa5565b50505050565b600e80546000918291612710906108b390600160a01b90046001600160601b031686612f72565b6108bd9190612f89565b90546001600160a01b03169590945092505050565b6108da611eec565b6108e382611f3c565b60405163ebb4a55f60e01b81526001600160a01b0383169063ebb4a55f906107dd9084906004016130ec565b826001600160a01b03811633146109295761092933611d93565b61088684848461213a565b61093c611eec565b600061094b6020830183612c87565b6001600160a01b03160361097257604051631cc0baef60e01b815260040160405180910390fd5b6127106109856040830160208401613114565b6001600160601b031611156109ce576109a46040820160208301613114565b604051633cadbafb60e01b81526001600160601b0390911660048201526024015b60405180910390fd5b80600e6109db8282613131565b507ff21fccf4d64d86d532c4e4eb86c007b6ad57a460c27d724188625e755ec6cf6d9050610a0c6020830183612c87565b610a1c6040840160208501613114565b604080516001600160a01b0390931683526001600160601b039091166020830152015b60405180910390a150565b610a52611eec565b610a5b83611f3c565b604051638e7d1e4360e01b81526001600160a01b0383811660048301528215156024830152841690638e7d1e43906044015b600060405180830381600087803b158015610aa757600080fd5b505af1158015610abb573d6000803e3d6000fd5b50505050505050565b610acc611eec565b610ad583611f3c565b6040516309a7002f60e31b81526001600160a01b03841690634d38017890610a8d90859085906004016131de565b610b0b611eec565b600b610b18828483613242565b50600154600054036000190115610b75577f6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c600180610b5660005490565b610b609190613301565b60408051928352602083019190915201610794565b5050565b610b81611f7a565b610b758282612155565b60006106388261226f565b600f54600114610bd55760405162461bcd60e51b815260206004820152600a6024820152695245454e5452414e435960b01b60448201526064016109c5565b6002600f55610be333611f3c565b600a5481610bf46000546000190190565b610bfe9190613314565b1115610c415780610c126000546000190190565b610c1c9190613314565b600a5460405163384b48c560e21b8152600481019290925260248201526044016109c5565b610c4b82826122de565b50506001600f55565b610c5c611eec565b610c6582611f3c565b60405163024e71b760e31b81526001600160a01b0382811660048301528316906312738db8906024016107dd565b6060610c9d6122f8565b905090565b610caa611eec565b6001600160401b03811115610cd55760405163b43e913760e01b8152600481018290526024016109c5565b600a8190556040518181527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c90602001610a3f565b60006001600160a01b038216610d33576040516323d3ad8160e21b815260040160405180910390fd5b506001600160a01b03166000908152600560205260409020546001600160401b031690565b610d60611f7a565b610d6a6000612307565b565b6009546001600160a01b0316338114610d9857604051636b7584e760e11b815260040160405180910390fd5b600980546001600160a01b0319169055604051600081527f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da9060200160405180910390a1610de581612307565b50565b610df0611eec565b610df983611f3c565b60405163b957d0cb60e01b81526001600160a01b0384169063b957d0cb90610a8d9085908590600401613327565b610e2f611eec565b610e3883611f3c565b604051637ecd591560e11b81526001600160a01b0384169063fd9ab22a90610a8d9085908590600401613409565b6001600160a01b03811660009081526005602052604080822054901c6001600160401b03169080610e9a6000546000190190565b600a549395909450915050565b610eaf611f7a565b803515610f08576040516306f8b44b60e41b8152813560048201523090636f8b44b090602401600060405180830381600087803b158015610eef57600080fd5b505af1158015610f03573d6000803e3d6000fd5b505050505b610f156020820182613427565b159050610f7f57306355f804b3610f2f6020840184613427565b6040518363ffffffff1660e01b8152600401610f4c929190613327565b600060405180830381600087803b158015610f6657600080fd5b505af1158015610f7a573d6000803e3d6000fd5b505050505b610f8c6040820182613427565b159050610ff6573063938e3d7b610fa66040840184613427565b6040518363ffffffff1660e01b8152600401610fc3929190613327565b600060405180830381600087803b158015610fdd57600080fd5b505af1158015610ff1573d6000803e3d6000fd5b505050505b61101661100960e0830160c0840161346d565b65ffffffffffff16151590565b61102961100960c0840160a0850161346d565b1760010361109b5730631b73593c6110476080840160608501612c87565b836080016040518363ffffffff1660e01b8152600401611068929190613488565b600060405180830381600087803b15801561108257600080fd5b505af1158015611096573d6000803e3d6000fd5b505050505b6110a9610140820182613427565b1590506111255730637a05bc826110c66080840160608501612c87565b6110d4610140850185613427565b6040518463ffffffff1660e01b81526004016110f2939291906134a5565b600060405180830381600087803b15801561110c57600080fd5b505af1158015611120573d6000803e3d6000fd5b505050505b60006111356101608301836134d3565b35146111af5730633680620d6111516080840160608501612c87565b61115f6101608501856134d3565b6040518363ffffffff1660e01b815260040161117c9291906134f3565b600060405180830381600087803b15801561119657600080fd5b505af11580156111aa573d6000803e3d6000fd5b505050505b60006111c36101a083016101808401612c87565b6001600160a01b03161461125a57306366251b696111e76080840160608501612c87565b6111f96101a085016101808601612c87565b6040516001600160e01b031960e085901b1681526001600160a01b03928316600482015291166024820152604401600060405180830381600087803b15801561124157600080fd5b505af1158015611255573d6000803e3d6000fd5b505050505b6101a0810135156112bb576040516304cdb5fd60e11b81526101a08201356004820152309063099b6bfa90602401600060405180830381600087803b1580156112a257600080fd5b505af11580156112b6573d6000803e3d6000fd5b505050505b60006112cb6101c0830183613517565b905011156113955760005b6112e46101c0830183613517565b905081101561139357306348a4c1016113036080850160608601612c87565b6113116101c0860186613517565b8581811061132157611321613560565b90506020020160208101906113369190612c87565b60016040518463ffffffff1660e01b815260040161135693929190613576565b600060405180830381600087803b15801561137057600080fd5b505af1158015611384573d6000803e3d6000fd5b505050508060010190506112d6565b505b60006113a56101e0830183613517565b9050111561146f5760005b6113be6101e0830183613517565b905081101561146d57306348a4c1016113dd6080850160608601612c87565b6113eb6101e0860186613517565b858181106113fb576113fb613560565b90506020020160208101906114109190612c87565b60006040518463ffffffff1660e01b815260040161143093929190613576565b600060405180830381600087803b15801561144a57600080fd5b505af115801561145e573d6000803e3d6000fd5b505050508060010190506113b0565b505b600061147f610200830183613517565b905011156115495760005b611498610200830183613517565b9050811015611547573063cb743ba86114b76080850160608601612c87565b6114c5610200860186613517565b858181106114d5576114d5613560565b90506020020160208101906114ea9190612c87565b60016040518463ffffffff1660e01b815260040161150a93929190613576565b600060405180830381600087803b15801561152457600080fd5b505af1158015611538573d6000803e3d6000fd5b5050505080600101905061148a565b505b6000611559610220830183613517565b905011156116235760005b611572610220830183613517565b9050811015611621573063cb743ba86115916080850160608601612c87565b61159f610220860186613517565b858181106115af576115af613560565b90506020020160208101906115c49190612c87565b60006040518463ffffffff1660e01b81526004016115e493929190613576565b600060405180830381600087803b1580156115fe57600080fd5b505af1158015611612573d6000803e3d6000fd5b50505050806001019050611564565b505b600061163361026083018361359a565b9050111561175e57611649610240820182613517565b905061165961026083018361359a565b9050146116795760405163b81aa63960e01b815260040160405180910390fd5b60005b61168a61026083018361359a565b905081101561175c5730637bc2be766116a96080850160608601612c87565b6116b7610240860186613517565b858181106116c7576116c7613560565b90506020020160208101906116dc9190612c87565b6116ea61026087018761359a565b868181106116fa576116fa613560565b905061010002016040518463ffffffff1660e01b815260040161171f939291906135e3565b600060405180830381600087803b15801561173957600080fd5b505af115801561174d573d6000803e3d6000fd5b5050505080600101905061167c565b505b600061176e610280830183613517565b905011156118795760005b611787610280830183613517565b9050811015611877576040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915230637bc2be766117e76080860160608701612c87565b6117f5610280870187613517565b8681811061180557611805613560565b905060200201602081019061181a9190612c87565b846040518463ffffffff1660e01b815260040161183993929190613609565b600060405180830381600087803b15801561185357600080fd5b505af1158015611867573d6000803e3d6000fd5b5050505081600101915050611779565b505b60006118896102c08301836136bb565b905011156119b35761189f6102a0820182613517565b90506118af6102c08301836136bb565b9050146118cf576040516374ef6df760e01b815260040160405180910390fd5b60005b6118e06102c08301836136bb565b90508110156119b1573063511aa6446118ff6080850160608601612c87565b61190d6102a0860186613517565b8581811061191d5761191d613560565b90506020020160208101906119329190612c87565b6119406102c08701876136bb565b8681811061195057611950613560565b905060e002016040518463ffffffff1660e01b815260040161197493929190613703565b600060405180830381600087803b15801561198e57600080fd5b505af11580156119a2573d6000803e3d6000fd5b505050508060010190506118d2565b505b60006119c36102e0830183613517565b90501115610de55760005b6119dc6102e0830183613517565b9050811015610b75576040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091523063511aa644611a346080860160608701612c87565b611a426102e0870187613517565b86818110611a5257611a52613560565b9050602002016020810190611a679190612c87565b846040518463ffffffff1660e01b8152600401611a86939291906137bc565b600060405180830381600087803b158015611aa057600080fd5b505af1158015611ab4573d6000803e3d6000fd5b50505050816001019150506119ce565b611acc611eec565b600c611ad9828483613242565b507f905d981207a7d0b6c62cc46ab0be2a076d0298e4a86d0ab79882dbd01ac373788282604051610794929190613327565b60606003805461064d90612e79565b81611b2481611d93565b6107288383612359565b611b36611eec565b60408051838152602081018390527f6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c9101610794565b836001600160a01b0381163314611b8657611b8633611d93565b611b92858585856123c5565b5050505050565b6060611ba482611d5e565b611bc157604051630a14c4b560e41b815260040160405180910390fd5b6000611bcb6122f8565b90508051600003611beb5760405180602001604052806000815250611c16565b80611bf584612409565b604051602001611c069291906137e2565b6040516020818303038152906040525b9392505050565b611c25611eec565b611c2e83611f3c565b604051633f952e6560e11b81526001600160a01b0383811660048301528215156024830152841690637f2a5cca90604401610a8d565b6060600c805461064d90612e79565b6001600160a01b03918216600090815260076020908152604080832093909416825291909152205460ff1690565b611ca9611f7a565b6001600160a01b038116611cd057604051633a247dd760e11b815260040160405180910390fd5b600980546001600160a01b0319166001600160a01b0383169081179091556040519081527f11a3cf439fb225bfe74225716b6774765670ec1060e3796802e62139d69974da90602001610a3f565b60006001600160e01b0319821663152a902d60e11b1480611d4f5750632483248360e11b6001600160e01b03198316145b8061063857506106388261244d565b600081600111158015611d72575060005482105b8015610638575050600090815260046020526040902054600160e01b161590565b6daaeb6d7670e522a718067333cd4e3b15610de557604051633185c44d60e21b81523060048201526001600160a01b03821660248201526daaeb6d7670e522a718067333cd4e9063c617113490604401602060405180830381865afa158015611e00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e249190613811565b610de557604051633b79c77360e21b81526001600160a01b03821660048201526024016109c5565b6000611e5782610b8b565b9050336001600160a01b03821614611e9057611e738133611c73565b611e90576040516367d9dca160e11b815260040160405180910390fd5b60008281526006602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b303314611f1b611f046008546001600160a01b031690565b6001600160a01b0316336001600160a01b03161490565b17600003610d6a57604051635fc483c560e01b815260040160405180910390fd5b6001600160a01b03811660009081526010602052604090205460ff161515600114610de5576040516315e26ff360e01b815260040160405180910390fd5b6008546001600160a01b03163314610d6a57604051635fc483c560e01b815260040160405180910390fd5b6000611fb08261226f565b9050836001600160a01b0316816001600160a01b031614611fe35760405162a1148160e81b815260040160405180910390fd5b60008281526006602052604090208054338082146001600160a01b03881690911417612030576120138633611c73565b61203057604051632ce44b5f60e11b815260040160405180910390fd5b6001600160a01b03851661205757604051633a954ecd60e21b815260040160405180910390fd5b801561206257600082555b6001600160a01b038681166000908152600560205260408082208054600019019055918716808252919020805460010190554260a01b17600160e11b17600085815260046020526040812091909155600160e11b841690036120f4576001840160008181526004602052604081205490036120f25760005481146120f25760008181526004602052604090208490555b505b83856001600160a01b0316876001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a461080b565b61072883838360405180602001604052806000815250611b6c565b6011548160005b828110156121b9576000601060006011848154811061217d5761217d613560565b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff191691151591909117905560010161215c565b5060005b81811015612222576001601060008787858181106121dd576121dd613560565b90506020020160208101906121f29190612c87565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790556001016121bd565b5061222f601185856126eb565b507fbbd3b69c138de4d317d0bc4290282c4e1cbd1e58b579a5b4f114b598c237454d848460405161226192919061382e565b60405180910390a150505050565b600081806001116122c5576000548110156122c55760008181526004602052604081205490600160e01b821690036122c3575b80600003611c165750600019016000818152600460205260409020546122a2565b505b604051636f96cda160e11b815260040160405180910390fd5b610b7582826040518060200160405280600081525061249b565b6060600b805461064d90612e79565b600880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b3360008181526007602090815260408083206001600160a01b03871680855290835292819020805460ff191686151590811790915590519081529192917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b6123d0848484610861565b6001600160a01b0383163b15610886576123ec84848484612501565b610886576040516368d2bf6b60e11b815260040160405180910390fd5b606060a06040510180604052602081039150506000815280825b600183039250600a81066030018353600a9004806124235750819003601f19909101908152919050565b60006301ffc9a760e01b6001600160e01b03198316148061247e57506380ac58cd60e01b6001600160e01b03198316145b806106385750506001600160e01b031916635b5e139f60e01b1490565b6124a583836125ed565b6001600160a01b0383163b15610728576000548281035b6124cf6000868380600101945086612501565b6124ec576040516368d2bf6b60e11b815260040160405180910390fd5b8181106124bc578160005414611b9257600080fd5b604051630a85bd0160e11b81526000906001600160a01b0385169063150b7a029061253690339089908890889060040161387c565b6020604051808303816000875af1925050508015612571575060408051601f3d908101601f1916820190925261256e918101906138af565b60015b6125cf573d80801561259f576040519150601f19603f3d011682016040523d82523d6000602084013e6125a4565b606091505b5080516000036125c7576040516368d2bf6b60e11b815260040160405180910390fd5b805181602001fd5b6001600160e01b031916630a85bd0160e11b1490505b949350505050565b60008054908290036126125760405163b562e8dd60e01b815260040160405180910390fd5b6001600160a01b03831660008181526005602090815260408083208054680100000000000000018802019055848352600490915281206001851460e11b4260a01b178317905582840190839083907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8180a4600183015b8181146126c157808360007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef600080a4600101612689565b50816000036126e257604051622e076360e81b815260040160405180910390fd5b60005550505050565b82805482825590600052602060002090810192821561273e579160200282015b8281111561273e5781546001600160a01b0319166001600160a01b0384351617825560209092019160019091019061270b565b5061274a92915061274e565b5090565b5b8082111561274a576000815560010161274f565b6001600160e01b031981168114610de557600080fd5b60006020828403121561278b57600080fd5b8135611c1681612763565b60005b838110156127b1578181015183820152602001612799565b50506000910152565b600081518084526127d2816020860160208601612796565b601f01601f19169290920160200192915050565b602081526000611c1660208301846127ba565b60006020828403121561280b57600080fd5b5035919050565b6001600160a01b0381168114610de557600080fd5b6000806040838503121561283a57600080fd5b823561284581612812565b946020939093013593505050565b60008082840360e081121561286757600080fd5b833561287281612812565b925060c0601f198201121561288657600080fd5b506020830190509250929050565b6000806000606084860312156128a957600080fd5b83356128b481612812565b925060208401356128c481612812565b929592945050506040919091013590565b600080604083850312156128e857600080fd5b50508035926020909101359150565b6000806040838503121561290a57600080fd5b823561291581612812565b915060208301356001600160401b0381111561293057600080fd5b83016060818603121561294257600080fd5b809150509250929050565b60006040828403121561295f57600080fd5b50919050565b8015158114610de557600080fd5b803561297e81612965565b919050565b60008060006060848603121561299857600080fd5b83356129a381612812565b925060208401356129b381612812565b915060408401356129c381612965565b809150509250925092565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b0381118282101715612a0657612a066129ce565b60405290565b604051601f8201601f191681016001600160401b0381118282101715612a3457612a346129ce565b604052919050565b80356001600160501b038116811461297e57600080fd5b803562ffffff8116811461297e57600080fd5b803564ffffffffff8116811461297e57600080fd5b803561ffff8116811461297e57600080fd5b6000806000838503610120811215612aa457600080fd5b8435612aaf81612812565b93506020850135612abf81612812565b925060e0603f1982011215612ad357600080fd5b50612adc6129e4565b612ae860408601612a3c565b8152612af660608601612a53565b6020820152612b0760808601612a66565b6040820152612b1860a08601612a66565b6060820152612b2960c08601612a66565b6080820152612b3a60e08601612a7b565b60a0820152612b4c6101008601612a7b565b60c0820152809150509250925092565b60008083601f840112612b6e57600080fd5b5081356001600160401b03811115612b8557600080fd5b602083019150836020828501011115612b9d57600080fd5b9250929050565b60008060208385031215612bb757600080fd5b82356001600160401b03811115612bcd57600080fd5b612bd985828601612b5c565b90969095509350505050565b60008060208385031215612bf857600080fd5b82356001600160401b0380821115612c0f57600080fd5b818501915085601f830112612c2357600080fd5b813581811115612c3257600080fd5b8660208260051b8501011115612c4757600080fd5b60209290920196919550909350505050565b60008060408385031215612c6c57600080fd5b8235612c7781612812565b9150602083013561294281612812565b600060208284031215612c9957600080fd5b8135611c1681612812565b600080600060408486031215612cb957600080fd5b8335612cc481612812565b925060208401356001600160401b03811115612cdf57600080fd5b612ceb86828701612b5c565b9497909650939450505050565b6000806000838503610140811215612d0f57600080fd5b8435612d1a81612812565b93506020850135612d2a81612812565b9250610100603f1982011215612d3f57600080fd5b506040840190509250925092565b600060208284031215612d5f57600080fd5b81356001600160401b03811115612d7557600080fd5b82016103008185031215611c1657600080fd5b60008060408385031215612d9b57600080fd5b8235612da681612812565b9150602083013561294281612965565b60008060008060808587031215612dcc57600080fd5b8435612dd781612812565b9350602085810135612de881612812565b93506040860135925060608601356001600160401b0380821115612e0b57600080fd5b818801915088601f830112612e1f57600080fd5b813581811115612e3157612e316129ce565b612e43601f8201601f19168501612a0c565b91508082528984828501011115612e5957600080fd5b808484018584013760008482840101525080935050505092959194509250565b600181811c90821680612e8d57607f821691505b60208210810361295f57634e487b7160e01b600052602260045260246000fd5b803565ffffffffffff8116811461297e57600080fd5b6001600160501b03612ed482612a3c565b168252612ee360208201612ead565b65ffffffffffff808216602085015280612eff60408501612ead565b1660408501525050612f1360608201612a7b565b61ffff808216606085015280612f2b60808501612a7b565b166080850152505060a0810135612f4181612965565b80151560a0840152505050565b60c081016106388284612ec3565b634e487b7160e01b600052601160045260246000fd5b808202811582820484141761063857610638612f5c565b600082612fa657634e487b7160e01b600052601260045260246000fd5b500490565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000808335601e19843603018112612feb57600080fd5b83016020810192503590506001600160401b0381111561300a57600080fd5b803603821315612b9d57600080fd5b60006060830182358452602080840135601e1985360301811261303b57600080fd5b840181810190356001600160401b0381111561305657600080fd5b8060051b80360383131561306957600080fd5b6060848901529381905260809387018401938290880160005b838110156130bc57898703607f1901825261309d8386612fd4565b6130a8898284612fab565b985050509185019190850190600101613082565b5050505050506130cf6040840184612fd4565b85830360408701526130e2838284612fab565b9695505050505050565b602081526000611c166020830184613019565b6001600160601b0381168114610de557600080fd5b60006020828403121561312657600080fd5b8135611c16816130ff565b813561313c81612812565b81546001600160a01b03199081166001600160a01b039290921691821783556020840135613169816130ff565b60a01b1617905550565b6001600160501b03815116825262ffffff6020820151166020830152604081015164ffffffffff8082166040850152806060840151166060850152806080840151166080850152505060a081015161ffff80821660a08501528060c08401511660c085015250505050565b6001600160a01b03831681526101008101611c166020830184613173565b601f82111561072857600081815260208120601f850160051c810160208610156132235750805b601f850160051c820191505b8181101561080b5782815560010161322f565b6001600160401b03831115613259576132596129ce565b61326d836132678354612e79565b836131fc565b6000601f8411600181146132a157600085156132895750838201355b600019600387901b1c1916600186901b178355611b92565b600083815260209020601f19861690835b828110156132d257868501358255602094850194600190920191016132b2565b50868210156132ef5760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b8181038181111561063857610638612f5c565b8082018082111561063857610638612f5c565b6020815260006125e5602083018486612fab565b803563ffffffff8116811461297e57600080fd5b6001600160501b0361336082612a3c565b16825261ffff61337260208301612a7b565b16602083015261338460408201612ead565b65ffffffffffff8082166040850152806133a060608501612ead565b1660608501525050608081013560ff811681146133bc57600080fd5b60ff1660808301526133d060a0820161333b565b63ffffffff1660a08301526133e760c08201612a7b565b61ffff1660c08301526133fc60e08201612973565b80151560e0840152505050565b6001600160a01b03831681526101208101611c16602083018461334f565b6000808335601e1984360301811261343e57600080fd5b8301803591506001600160401b0382111561345857600080fd5b602001915036819003821315612b9d57600080fd5b60006020828403121561347f57600080fd5b611c1682612ead565b6001600160a01b038316815260e08101611c166020830184612ec3565b6001600160a01b03841681526040602082018190526000906134ca9083018486612fab565b95945050505050565b60008235605e198336030181126134e957600080fd5b9190910192915050565b6001600160a01b03831681526040602082018190526000906125e590830184613019565b6000808335601e1984360301811261352e57600080fd5b8301803591506001600160401b0382111561354857600080fd5b6020019150600581901b3603821315612b9d57600080fd5b634e487b7160e01b600052603260045260246000fd5b6001600160a01b039384168152919092166020820152901515604082015260600190565b6000808335601e198436030181126135b157600080fd5b8301803591506001600160401b038211156135cb57600080fd5b6020019150600881901b3603821315612b9d57600080fd5b6001600160a01b0384811682528316602082015261014081016125e5604083018461334f565b60006101408201905060018060a01b0380861683528085166020840152506001600160501b03835116604083015261ffff602084015116606083015265ffffffffffff6040840151166080830152606083015161367060a084018265ffffffffffff169052565b50608083015160ff811660c08401525060a083015163ffffffff811660e08401525060c083015161ffff81166101008401525060e08301518015156101208401525b50949350505050565b6000808335601e198436030181126136d257600080fd5b8301803591506001600160401b038211156136ec57600080fd5b602001915060e081023603821315612b9d57600080fd5b6001600160a01b0384811682528316602082015261012081016001600160501b0361372d84612a3c565b16604083015262ffffff61374360208501612a53565b16606083015261375560408401612a66565b64ffffffffff80821660808501528061377060608701612a66565b1660a08501528061378360808701612a66565b1660c0850152505061379760a08401612a7b565b61ffff1660e08301526137ac60c08401612a7b565b61ffff81166101008401526136b2565b6001600160a01b0384811682528316602082015261012081016125e56040830184613173565b600083516137f4818460208801612796565b835190830190613808818360208801612796565b01949350505050565b60006020828403121561382357600080fd5b8151611c1681612965565b60208082528181018390526000908460408401835b8681101561387157823561385681612812565b6001600160a01b031682529183019190830190600101613843565b509695505050505050565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906130e2908301846127ba565b6000602082840312156138c157600080fd5b8151611c168161276356fea264697066735822122015d264760cb55ebe196c194e56a9105fa2100b6993888d7e97ca4031d9c4183f64736f6c63430008110033",
}

// TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMetaData.ABI instead.
var TokenABI = TokenMetaData.ABI

// TokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenMetaData.Bin instead.
var TokenBin = TokenMetaData.Bin

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, allowedSeaDrop []common.Address) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := TokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenBin), backend, name, symbol, allowedSeaDrop)
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

// OPERATORFILTERREGISTRY is a free data retrieval call binding the contract method 0x41f43434.
//
// Solidity: function OPERATOR_FILTER_REGISTRY() view returns(address)
func (_Token *TokenCaller) OPERATORFILTERREGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "OPERATOR_FILTER_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OPERATORFILTERREGISTRY is a free data retrieval call binding the contract method 0x41f43434.
//
// Solidity: function OPERATOR_FILTER_REGISTRY() view returns(address)
func (_Token *TokenSession) OPERATORFILTERREGISTRY() (common.Address, error) {
	return _Token.Contract.OPERATORFILTERREGISTRY(&_Token.CallOpts)
}

// OPERATORFILTERREGISTRY is a free data retrieval call binding the contract method 0x41f43434.
//
// Solidity: function OPERATOR_FILTER_REGISTRY() view returns(address)
func (_Token *TokenCallerSession) OPERATORFILTERREGISTRY() (common.Address, error) {
	return _Token.Contract.OPERATORFILTERREGISTRY(&_Token.CallOpts)
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
func (_Token *TokenTransactor) MultiConfigure(opts *bind.TransactOpts, config ERC721SeaDropStructsErrorsAndEventsMultiConfigureStruct) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "multiConfigure", config)
}

// MultiConfigure is a paid mutator transaction binding the contract method 0x911f456b.
//
// Solidity: function multiConfigure((uint256,string,string,address,(uint80,uint48,uint48,uint16,uint16,bool),string,(bytes32,string[],string),address,bytes32,address[],address[],address[],address[],address[],(uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool)[],address[],address[],(uint80,uint24,uint40,uint40,uint40,uint16,uint16)[],address[]) config) returns()
func (_Token *TokenSession) MultiConfigure(config ERC721SeaDropStructsErrorsAndEventsMultiConfigureStruct) (*types.Transaction, error) {
	return _Token.Contract.MultiConfigure(&_Token.TransactOpts, config)
}

// MultiConfigure is a paid mutator transaction binding the contract method 0x911f456b.
//
// Solidity: function multiConfigure((uint256,string,string,address,(uint80,uint48,uint48,uint16,uint16,bool),string,(bytes32,string[],string),address,bytes32,address[],address[],address[],address[],address[],(uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool)[],address[],address[],(uint80,uint24,uint40,uint40,uint40,uint16,uint16)[],address[]) config) returns()
func (_Token *TokenTransactorSession) MultiConfigure(config ERC721SeaDropStructsErrorsAndEventsMultiConfigureStruct) (*types.Transaction, error) {
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
func (_Token *TokenTransactor) SetRoyaltyInfo(opts *bind.TransactOpts, newInfo ISeaDropTokenContractMetadataRoyaltyInfo) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setRoyaltyInfo", newInfo)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_Token *TokenSession) SetRoyaltyInfo(newInfo ISeaDropTokenContractMetadataRoyaltyInfo) (*types.Transaction, error) {
	return _Token.Contract.SetRoyaltyInfo(&_Token.TransactOpts, newInfo)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x44dae42c.
//
// Solidity: function setRoyaltyInfo((address,uint96) newInfo) returns()
func (_Token *TokenTransactorSession) SetRoyaltyInfo(newInfo ISeaDropTokenContractMetadataRoyaltyInfo) (*types.Transaction, error) {
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
