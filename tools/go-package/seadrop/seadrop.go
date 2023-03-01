// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package seadrop

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

// TokenGatedMintParams is an auto generated low-level Go binding around an user-defined struct.
type TokenGatedMintParams struct {
	AllowedNftToken    common.Address
	AllowedNftTokenIds []*big.Int
}

// SeadropMetaData contains all meta data concerning the Seadrop contract.
var SeadropMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CreatorPayoutAddressCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateFeeRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicatePayer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientNotPresent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"IncorrectPayment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeBps\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recoveredSigner\",\"type\":\"address\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedEndTime\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumOrMaximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedFeeBps\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedMaxTokenSupplyForStage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedMaxTotalMintableByWallet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedMintPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintQuantityCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allowed\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxMintedPerWallet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxSupply\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxTokenSupplyForStage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"}],\"name\":\"NotActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"OnlyINonFungibleSeaDropToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerNotPresent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignedMintsMustRestrictFeeRecipients\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerNotPresent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedDropAllowedNftTokenCannotBeDropToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedDropAllowedNftTokenCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedDropStageNotPresent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowedNftTokenId\",\"type\":\"uint256\"}],\"name\":\"TokenGatedNotTokenOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowedNftTokenId\",\"type\":\"uint256\"}],\"name\":\"TokenGatedTokenIdAlreadyRedeemed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousMerkleRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newMerkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"publicKeyURI\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"name\":\"AllowListUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AllowedFeeRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPayoutAddress\",\"type\":\"address\"}],\"name\":\"CreatorPayoutAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newDropURI\",\"type\":\"string\"}],\"name\":\"DropURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"PayerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"}],\"name\":\"PublicDropUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantityMinted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unitMintPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dropStageIndex\",\"type\":\"uint256\"}],\"name\":\"SeaDropMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structSignedMintValidationParams\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple\"}],\"name\":\"SignedMintValidationParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTokenGatedDropStage\",\"name\":\"dropStage\",\"type\":\"tuple\"}],\"name\":\"TokenGatedDropStageUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getAllowListMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getAllowedFeeRecipients\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowedNftTokenId\",\"type\":\"uint256\"}],\"name\":\"getAllowedNftTokenIdIsRedeemed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getCreatorPayoutAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"getFeeRecipientIsAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"}],\"name\":\"getPayerIsAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getPayers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getPublicDrop\",\"outputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structPublicDrop\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getSignedMintValidationParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"internalType\":\"structSignedMintValidationParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getTokenGatedAllowedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"}],\"name\":\"getTokenGatedDrop\",\"outputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structTokenGatedDropStage\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minterIfNotPayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dropStageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structMintParams\",\"name\":\"mintParams\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"mintAllowList\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minterIfNotPayer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"allowedNftTokenIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structTokenGatedMintParams\",\"name\":\"mintParams\",\"type\":\"tuple\"}],\"name\":\"mintAllowedTokenHolder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minterIfNotPayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"mintPublic\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minterIfNotPayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dropStageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structMintParams\",\"name\":\"mintParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mintSigned\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"publicKeyURIs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"internalType\":\"structAllowListData\",\"name\":\"allowListData\",\"type\":\"tuple\"}],\"name\":\"updateAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateAllowedFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payoutAddress\",\"type\":\"address\"}],\"name\":\"updateCreatorPayoutAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dropURI\",\"type\":\"string\"}],\"name\":\"updateDropURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updatePayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"}],\"name\":\"updatePublicDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"internalType\":\"structSignedMintValidationParams\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple\"}],\"name\":\"updateSignedMintValidationParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structTokenGatedDropStage\",\"name\":\"dropStage\",\"type\":\"tuple\"}],\"name\":\"updateTokenGatedDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405260016000554660805234801561001957600080fd5b506100c2604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f752a02269614d51d9b7bd0a2f05cf03e553ce6be8b487650a6a2a4990208d804918101919091527fe6bbd6277e1bf288eed5e8d1780f9a50b239e86b153736bceebccf4ea79d90b360608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b60a05260805160a051613df16100ea6000396000612a020152600061292d0152613df16000f3fe60806040526004361061014b5760003560e01c80637c35b982116100b657806399eb900f1161006f57806399eb900f146105a6578063b957d0cb146105b9578063bc6a629c146105d9578063e583141d14610722578063ebb4a55f1461076b578063fd9ab22a1461078b57600080fd5b80637c35b9821461035c5780637e3ba6af1461037c5780637f2a5cca1461039c57806381bf9af3146103bc57806388aa3d37146105365780638e7d1e431461058657600080fd5b806332bf11f51161010857806332bf11f5146102615780634300a4e6146102a55780634b61cd6f146102b85780634d380178146102cb5780635cb3c4d3146102eb578063686322741461033c57600080fd5b806301308e65146101505780630b0e8a6e1461017257806312738db8146101a8578063161ac21f146101c85780632db526eb146101db578063322e75d114610208575b600080fd5b34801561015c57600080fd5b5061017061016b366004612cfd565b6107ab565b005b34801561017e57600080fd5b5061019261018d366004612d2a565b6108e1565b60405161019f9190612d63565b60405180910390f35b3480156101b457600080fd5b506101706101c3366004612e06565b6109e3565b6101706101d6366004612e23565b610aeb565b3480156101e757600080fd5b506101fb6101f6366004612e06565b610c46565b60405161019f9190612e74565b34801561021457600080fd5b50610251610223366004612d2a565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b604051901515815260200161019f565b34801561026d57600080fd5b5061029761027c366004612e06565b6001600160a01b031660009081526003602052604090205490565b60405190815260200161019f565b6101706102b3366004612ed4565b610cbc565b6101706102c6366004612fe1565b610e33565b3480156102d757600080fd5b506101706102e6366004613084565b610fc7565b3480156102f757600080fd5b50610324610306366004612e06565b6001600160a01b039081166000908152600260205260409020541690565b6040516001600160a01b03909116815260200161019f565b34801561034857600080fd5b506101fb610357366004612e06565b611233565b34801561036857600080fd5b506101fb610377366004612e06565b6112a7565b34801561038857600080fd5b506101fb610397366004612e06565b61131b565b3480156103a857600080fd5b506101706103b73660046130e4565b61138f565b3480156103c857600080fd5b506104b96103d7366004612d2a565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152506001600160a01b039182166000908152600660209081526040808320939094168252918252829020825160e08101845290546001600160501b0381168252600160501b810462ffffff1692820192909252600160681b820464ffffffffff90811693820193909352600160901b820483166060820152600160b81b82049092166080830152600160e01b810461ffff90811660a0840152600160f01b9091041660c082015290565b60405161019f9190600060e0820190506001600160501b03835116825262ffffff6020840151166020830152604083015164ffffffffff8082166040850152806060860151166060850152806080860151166080850152505060a083015161ffff80821660a08501528060c08601511660c0850152505092915050565b34801561054257600080fd5b50610251610551366004613112565b6001600160a01b039283166000908152600d602090815260408083209490951682529283528381209181529152205460ff1690565b34801561059257600080fd5b506101706105a13660046130e4565b611591565b6101706105b4366004613153565b611793565b3480156105c557600080fd5b506101706105d43660046131cd565b611adf565b3480156105e557600080fd5b506106b36105f4366004612e06565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506001600160a01b0316600090815260016020908152604091829020825160c08101845290546001600160501b038116825265ffffffffffff600160501b8204811693830193909352600160801b81049092169281019290925261ffff600160b01b820481166060840152600160c01b820416608083015260ff600160d01b90910416151560a082015290565b60405161019f9190600060c0820190506001600160501b038351168252602083015165ffffffffffff80821660208501528060408601511660408501525050606083015161ffff8082166060850152806080860151166080850152505060a0830151151560a083015292915050565b34801561072e57600080fd5b5061025161073d366004612d2a565b6001600160a01b03918216600090815260096020908152604080832093909416825291909152205460ff1690565b34801561077757600080fd5b5061017061078636600461320f565b611bb0565b34801561079757600080fd5b506101706107a636600461324a565b611cb1565b6040516301ffc9a760e01b815233906301ffc9a7906107d590630c487f4760e11b90600401613281565b602060405180830381865afa1580156107f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108169190613296565b61083a576040516332c5d8cf60e01b81523360048201526024015b60405180910390fd5b61271061084d60a08301608084016132ce565b61ffff1611156108875761086760a08201608083016132ce565b604051631994fc9960e11b815261ffff9091166004820152602401610831565b33600090815260016020526040902081906108a2828261333b565b505060405133907f3e30d8e1f739ea4795c481b21c23f905e938b80339305f3508e43c558e5dead3906108d6908490613471565b60405180910390a250565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810191909152506001600160a01b038083166000908152600b6020908152604080832093851683529281529082902082516101008101845290546001600160501b038116825261ffff600160501b820481169383019390935265ffffffffffff600160601b8204811694830194909452600160901b8104909316606082015260ff600160c01b84048116608083015263ffffffff600160c81b85041660a0830152600160e81b840490921660c0820152600160f81b90920416151560e08201525b92915050565b6040516301ffc9a760e01b815233906301ffc9a790610a0d90630c487f4760e11b90600401613281565b602060405180830381865afa158015610a2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a4e9190613296565b610a6d576040516332c5d8cf60e01b8152336004820152602401610831565b6001600160a01b038116610a9457604051633f00976960e01b815260040160405180910390fd5b3360008181526002602052604080822080546001600160a01b0319166001600160a01b03861690811790915590519092917f0c69f21751e800ea5960436c9a94370c7adbf54c733a20a025293fbbe8f1625291a350565b6001600160a01b038416600090815260016020908152604091829020825160c08101845290546001600160501b038116825265ffffffffffff600160501b82048116938301849052600160801b82041693820184905261ffff600160b01b820481166060840152600160c01b820416608083015260ff600160d01b90910416151560a082015291610b7c9190611f18565b80516001600160501b0316610b918382611f5f565b60006001600160a01b038516610ba75733610ba9565b845b90506001600160a01b0381163314610bff576001600160a01b038716600090815260096020908152604080832033845290915290205460ff16610bff576040516303fcfb4160e31b815260040160405180910390fd5b610c16878286866060015161ffff16600019611f9c565b610c2587878560a001516120f3565b610c3d878286856000886080015161ffff168c61216b565b50505050505050565b6001600160a01b0381166000908152600c6020908152604091829020805483518184028101840190945280845260609392830182828015610cb057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610c92575b50505050509050919050565b610cce83604001358460600135611f18565b8235610cda8582611f5f565b60006001600160a01b038716610cf05733610cf2565b865b90506001600160a01b0381163314610d48576001600160a01b038916600090815260096020908152604080832033845290915290205460ff16610d48576040516303fcfb4160e31b815260040160405180910390fd5b610d5d89828888602001358960a00135611f9c565b610d788989610d73610100890160e08a01613508565b6120f3565b610df484848080602002602001604051908101604052809392919081815260200183836020028082843760009201829052506001600160a01b038f16815260036020908152604091829020549151919450610dd993508792508b9101613525565b60405160208183030381529060405280519060200120612294565b610e11576040516309bde33960e01b815260040160405180910390fd5b610e288982888589608001358a60c001358e61216b565b505050505050505050565b610e4584604001358560600135611f18565b610e50858535611f5f565b60006001600160a01b038716610e665733610e68565b865b90506001600160a01b0381163314610ebe576001600160a01b038916600090815260096020908152604080832033845290915290205460ff16610ebe576040516303fcfb4160e31b815260040160405180910390fd5b610ed389828888602001358960a00135611f9c565b610ee98989610d73610100890160e08a01613508565b6000610f068a838b610f00368b90038b018b61359b565b896122aa565b60008181526008602052604090205490915060ff1615610f395760405163900bb2c960e01b815260040160405180910390fd5b6000818152600860209081526040808320805460ff191660011790558051601f8701839004830281018301909152858152610f91918790879081908401838280828437600092019190915250869392505061242b9050565b9050610fac8b610fa6368a90038a018a61359b565b8361244f565b50610e289050898288883560808a013560c08b01358e61216b565b6040516301ffc9a760e01b815233906301ffc9a790610ff190630c487f4760e11b90600401613281565b602060405180830381865afa15801561100e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110329190613296565b611051576040516332c5d8cf60e01b8152336004820152602401610831565b6001600160a01b038216611078576040516367db084560e11b815260040160405180910390fd5b61271061108b60c0830160a084016132ce565b61ffff1611156110a55761086760c0820160a083016132ce565b6127106110b860e0830160c084016132ce565b61ffff1611156110d25761086760e0820160c083016132ce565b336000908152600760209081526040808320600683528184206001600160a01b03871685528084528285208054929591949093921592829161111891890190890161364e565b62ffffff161190508015611183576001600160a01b0387166000908152602085905260409020869061114a828261368b565b5050811561117e5784546001810186556000868152602090200180546001600160a01b0319166001600160a01b0389161790555b6111df565b8254600160501b900462ffffff166000036111b157604051632d018df960e21b815260040160405180910390fd5b3360009081526006602090815260408083206001600160a01b038b1684529091528120556111df8786612736565b866001600160a01b0316336001600160a01b03167fcaeb4009c05208df426d15ff50b608287b05d21dee1f790552ea451a540a7be08860405161122291906137da565b60405180910390a350505050505050565b6001600160a01b038116600090815260056020908152604091829020805483518184028101840190945280845260609392830182828015610cb0576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610c925750505050509050919050565b6001600160a01b0381166000908152600a6020908152604091829020805483518184028101840190945280845260609392830182828015610cb0576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610c925750505050509050919050565b6001600160a01b038116600090815260076020908152604091829020805483518184028101840190945280845260609392830182828015610cb0576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610c925750505050509050919050565b6040516301ffc9a760e01b815233906301ffc9a7906113b990630c487f4760e11b90600401613281565b602060405180830381865afa1580156113d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113fa9190613296565b611419576040516332c5d8cf60e01b8152336004820152602401610831565b6001600160a01b038216611440576040516334d11a2f60e21b815260040160405180910390fd5b336000908152600a60209081526040808320600990925290912082156114e3576001600160a01b03841660009081526020829052604090205460ff161561149a57604051636a47e97360e11b815260040160405180910390fd5b6001600160a01b0384166000818152602083815260408220805460ff191660019081179091558554908101865585835291200180546001600160a01b0319169091179055611551565b6001600160a01b03841660009081526020829052604090205460ff1661151c57604051634cc1171360e01b815260040160405180910390fd5b3360009081526009602090815260408083206001600160a01b03881684529091529020805460ff191690556115518483612736565b604051831515906001600160a01b0386169033907f55a5cfa4bc68ffb9d833b75bf93f6d9c9aadc558dbfa587a9b5bb0ea7d5c38a390600090a450505050565b6040516301ffc9a760e01b815233906301ffc9a7906115bb90630c487f4760e11b90600401613281565b602060405180830381865afa1580156115d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115fc9190613296565b61161b576040516332c5d8cf60e01b8152336004820152602401610831565b6001600160a01b03821661164257604051635136e8d560e01b815260040160405180910390fd5b336000908152600560209081526040808320600490925290912082156116e5576001600160a01b03841660009081526020829052604090205460ff161561169c57604051631e61c06b60e21b815260040160405180910390fd5b6001600160a01b0384166000818152602083815260408220805460ff191660019081179091558554908101865585835291200180546001600160a01b0319169091179055611753565b6001600160a01b03841660009081526020829052604090205460ff1661171e57604051630998fbbd60e01b815260040160405180910390fd5b3360009081526004602090815260408083206001600160a01b03881684529091529020805460ff191690556117538483612736565b604051831515906001600160a01b0386169033907f6486c31f9d664e241acf94ec2541d328f6b9e97257ae16a1d887f296f879719f90600090a450505050565b60006001600160a01b0383166117a957336117ab565b825b90506001600160a01b0381163314611801576001600160a01b038516600090815260096020908152604080832033845290915290205460ff16611801576040516303fcfb4160e31b815260040160405180910390fd5b60006118106020840184612e06565b6001600160a01b038088166000908152600b6020908152604080832093851683529281529082902082516101008101845290546001600160501b038116825261ffff600160501b820481169383019390935265ffffffffffff600160601b82048116948301859052600160901b8204166060830181905260ff600160c01b83048116608085015263ffffffff600160c81b84041660a0850152600160e81b830490941660c0840152600160f81b909104909216151560e08201529293506118d79190611f18565b6118e687878360e001516120f3565b60006118f56020860186613887565b905090506119108183600001516001600160501b0316611f5f565b61192f888583856020015161ffff168660a0015163ffffffff16611f9c565b60005b81811015611aa95760006119496020880188613887565b83818110611959576119596138d1565b905060200201359050856001600160a01b0316856001600160a01b0316636352211e836040518263ffffffff1660e01b815260040161199a91815260200190565b602060405180830381865afa1580156119b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119db91906138e7565b6001600160a01b031614611a1c5760405163da8c7bc760e01b81526001600160a01b03808c1660048301528616602482015260448101829052606401610831565b6001600160a01b03808b166000908152600d602090815260408083209389168352928152828220848352908190529190205460ff1615611a895760405163a93f299b60e01b81526001600160a01b03808d1660048301528716602482015260448101839052606401610831565b600091825260205260409020805460ff1916600190811790915501611932565b50611ad588858385600001516001600160501b0316866080015160ff168760c0015161ffff168d61216b565b5050505050505050565b6040516301ffc9a760e01b815233906301ffc9a790611b0990630c487f4760e11b90600401613281565b602060405180830381865afa158015611b26573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b4a9190613296565b611b69576040516332c5d8cf60e01b8152336004820152602401610831565b336001600160a01b03167fa0295608d25b3033c2e2c41cbac8746c2d08767bcfde6d47fae1ed7ba1d321508383604051611ba492919061392d565b60405180910390a25050565b6040516301ffc9a760e01b815233906301ffc9a790611bda90630c487f4760e11b90600401613281565b602060405180830381865afa158015611bf7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c1b9190613296565b611c3a576040516332c5d8cf60e01b8152336004820152602401610831565b33600081815260036020908152604090912080548435918290559290918391907fefcd7e019bc8b47d27881fd59e2619280ca5894f285950f10ab049870652efa590611c8890870187613887565b611c956040890189613949565b604051611ca59493929190613990565b60405180910390a45050565b6040516301ffc9a760e01b815233906301ffc9a790611cdb90630c487f4760e11b90600401613281565b602060405180830381865afa158015611cf8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d1c9190613296565b611d3b576040516332c5d8cf60e01b8152336004820152602401610831565b6001600160a01b038216611d62576040516367156a2d60e01b815260040160405180910390fd5b336001600160a01b03831603611d8b576040516302f02fbb60e31b815260040160405180910390fd5b612710611d9e60e0830160c084016132ce565b61ffff161115611db85761086760e0820160c083016132ce565b6000611dca60408301602084016132ce565b336000818152600b602090815260408083206001600160a01b03891684528252808320938352600c9091529020815461ffff93909316158015945091929091901590611e7857336000908152600b602090815260408083206001600160a01b038a16845290915290208590611e3f8282613a7d565b50508015611e735781546001810183556000838152602090200180546001600160a01b0319166001600160a01b0388161790555b611ec5565b8015611e9757604051632b60a32f60e01b815260040160405180910390fd5b336000908152600b602090815260408083206001600160a01b038a168452909152812055611ec58683612736565b856001600160a01b0316336001600160a01b03167fc695f93ae16034280e4fc93181b6afca9af23027ac1f1842a2287ba25cdc447687604051611f089190613c14565b60405180910390a3505050505050565b611f2181421190565b611f2a83421090565b17600103611f5b576040516309ed117960e11b81524260048201526024810183905260448101829052606401610831565b5050565b611f698183613ce9565b3414611f5b5734611f7a8284613ce9565b604051630d35e92160e01b815260048101929092526024820152604401610831565b82600003611fbd5760405163198441cb60e01b815260040160405180910390fd5b604051632103857560e21b81526001600160a01b038581166004830152600091829182919089169063840e15d490602401606060405180830381865afa15801561200b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061202f9190613d00565b91945092509050846120418488613d2e565b1115612075576120518387613d2e565b60405163edc0127360e01b8152600481019190915260248101869052604401610831565b806120808388613d2e565b11156120b4576120908287613d2e565b60405163384b48c560e21b8152600481019190915260248101829052604401610831565b836120bf8388613d2e565b1115611ad5576120cf8287613d2e565b604051635cc6d5f560e11b8152600481019190915260248101859052604401610831565b6001600160a01b03821661211a57604051635136e8d560e01b815260040160405180910390fd5b8015612166576001600160a01b0380841660009081526004602090815260408083209386168352929052205460ff166121665760405163f477d26f60e01b815260040160405180910390fd5b505050565b6000546001146121aa5760405162461bcd60e51b815260206004820152600a6024820152695245454e5452414e435960b01b6044820152606401610831565b60026000556040516364869dad60e01b81526001600160a01b038781166004830152602482018790528816906364869dad90604401600060405180830381600087803b1580156121f957600080fd5b505af115801561220d573d6000803e3d6000fd5b505050508360001461222457612224878284612832565b604080513381526020810187905290810185905260608101839052608081018490526001600160a01b0380831691888216918a16907fe90cf9cc0a552cf52ea6ff74ece0f1c8ae8cc9ad630d3181f55ac43ca076b7d69060a00160405180910390a4505060016000555050505050565b6000826122a185846128e4565b14949350505050565b6000807f632d30b7600fe596b016656d5dc3bc1e2c318bf422c8844e42871665322c484b846000015185602001518660400151876060015188608001518960a001518a60c001518b60e0015160405160200161234a99989796959493929190988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e083015215156101008201526101200190565b60405160208183030381529060405280519060200120905061190160f01b612370612929565b604080517f8927086ec138a7aa6009c4966fc394504ae49ab0c06ae815b48e32e85c8279c760208201526001600160a01b03808c1692820192909252818a166060820152908816608082015260a0810184905260c0810186905260e00160408051601f198184030181529082905280516020918201206001600160f01b031990941690820152602281019190915260428101919091526062016040516020818303038152906040528051906020012091505095945050505050565b600080600061243a8585612a24565b9150915061244781612a69565b509392505050565b6001600160a01b0380841660009081526006602090815260408083209385168352928152828220835160e08101855290546001600160501b038116825262ffffff600160501b82041692820183905264ffffffffff600160681b8204811695830195909552600160901b810485166060830152600160b81b8104909416608082015261ffff600160e01b8504811660a0830152600160f01b90940490931660c0840152900361251c57604051633615713d60e21b81526001600160a01b0383166004820152602401610831565b805183516001600160501b0390911611156125605782518151604051635061f68560e11b815260048101929092526001600160501b03166024820152604401610831565b806020015162ffffff16836020015111156125a657602080840151908201516040516309e860af60e31b8152600481019290925262ffffff166024820152604401610831565b806040015164ffffffffff16836040015110156125f25782604001518160400151604051630333d33d60e41b815260040161083192919091825264ffffffffff16602082015260400190565b806060015164ffffffffff168360600151111561263c5760608084015190820151604051636e1d357d60e01b8152600481019290925264ffffffffff166024820152604401610831565b806080015164ffffffffff168360a0015111156126865760a083015160808201516040516306d029c560e41b8152600481019290925264ffffffffff166024820152604401610831565b8060c0015161ffff168360c0015111156126ca5760c080840151908201516040516379fc44ed60e01b8152600481019290925261ffff166024820152604401610831565b8060a0015161ffff168360c00151101561270e5760c083015160a08201516040516379fc44ed60e01b8152600481019290925261ffff166024820152604401610831565b8260e001516127305760405163db8b2fad60e01b815260040160405180910390fd5b50505050565b805460005b8181101561273057836001600160a01b031683828154811061275f5761275f6138d1565b6000918252602090912001546001600160a01b03160361282a5782612785600184613d41565b81548110612795576127956138d1565b9060005260206000200160009054906101000a90046001600160a01b03168382815481106127c5576127c56138d1565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508280548061280357612803613d54565b600082815260209020810160001990810180546001600160a01b0319169055019055612730565b60010161273b565b61271081111561285857604051631994fc9960e11b815260048101829052602401610831565b6001600160a01b03808416600090815260026020526040902054168061289157604051633f00976960e01b815260040160405180910390fd5b816000036128a3576127308134612bb6565b60006127106128b28434613ce9565b6128bc9190613d6a565b90503481900381156128d2576128d28583612bb6565b6128dc8382612bb6565b505050505050565b600081815b84518110156124475761291582868381518110612908576129086138d1565b6020026020010151612c07565b91508061292181613d8c565b9150506128e9565b60007f000000000000000000000000000000000000000000000000000000000000000046146129ff576129fa604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f752a02269614d51d9b7bd0a2f05cf03e553ce6be8b487650a6a2a4990208d804918101919091527fe6bbd6277e1bf288eed5e8d1780f9a50b239e86b153736bceebccf4ea79d90b360608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b905090565b507f000000000000000000000000000000000000000000000000000000000000000090565b6000808251604103612a5a5760208301516040840151606085015160001a612a4e87828585612c39565b94509450505050612a62565b506000905060025b9250929050565b6000816004811115612a7d57612a7d613da5565b03612a855750565b6001816004811115612a9957612a99613da5565b03612ae65760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610831565b6002816004811115612afa57612afa613da5565b03612b475760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610831565b6003816004811115612b5b57612b5b613da5565b03612bb35760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610831565b50565b600080600080600085875af19050806121665760405162461bcd60e51b815260206004820152601360248201527211551217d514905394d1915497d19052531151606a1b6044820152606401610831565b6000818310612c23576000828152602084905260409020612c32565b60008381526020839052604090205b9392505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115612c705750600090506003612cf4565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015612cc4573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116612ced57600060019250925050612cf4565b9150600090505b94509492505050565b600060c08284031215612d0f57600080fd5b50919050565b6001600160a01b0381168114612bb357600080fd5b60008060408385031215612d3d57600080fd5b8235612d4881612d15565b91506020830135612d5881612d15565b809150509250929050565b6000610100820190506001600160501b03835116825261ffff6020840151166020830152604083015165ffffffffffff808216604085015280606086015116606085015250506080830151612dbd608084018260ff169052565b5060a0830151612dd560a084018263ffffffff169052565b5060c0830151612deb60c084018261ffff169052565b5060e0830151612dff60e084018215159052565b5092915050565b600060208284031215612e1857600080fd5b8135612c3281612d15565b60008060008060808587031215612e3957600080fd5b8435612e4481612d15565b93506020850135612e5481612d15565b92506040850135612e6481612d15565b9396929550929360600135925050565b6020808252825182820181905260009190848201906040850190845b81811015612eb55783516001600160a01b031683529284019291840191600101612e90565b50909695505050505050565b60006101008284031215612d0f57600080fd5b60008060008060008060006101a0888a031215612ef057600080fd5b8735612efb81612d15565b96506020880135612f0b81612d15565b95506040880135612f1b81612d15565b945060608801359350612f318960808a01612ec1565b925061018088013567ffffffffffffffff80821115612f4f57600080fd5b818a0191508a601f830112612f6357600080fd5b813581811115612f7257600080fd5b8b60208260051b8501011115612f8757600080fd5b60208301945080935050505092959891949750929550565b60008083601f840112612fb157600080fd5b50813567ffffffffffffffff811115612fc957600080fd5b602083019150836020828501011115612a6257600080fd5b6000806000806000806000806101c0898b031215612ffe57600080fd5b883561300981612d15565b9750602089013561301981612d15565b9650604089013561302981612d15565b95506060890135945061303f8a60808b01612ec1565b935061018089013592506101a089013567ffffffffffffffff81111561306457600080fd5b6130708b828c01612f9f565b999c989b5096995094979396929594505050565b60008082840361010081121561309957600080fd5b83356130a481612d15565b925060e0601f19820112156130b857600080fd5b506020830190509250929050565b8015158114612bb357600080fd5b80356130df816130c6565b919050565b600080604083850312156130f757600080fd5b823561310281612d15565b91506020830135612d58816130c6565b60008060006060848603121561312757600080fd5b833561313281612d15565b9250602084013561314281612d15565b929592945050506040919091013590565b6000806000806080858703121561316957600080fd5b843561317481612d15565b9350602085013561318481612d15565b9250604085013561319481612d15565b9150606085013567ffffffffffffffff8111156131b057600080fd5b8501604081880312156131c257600080fd5b939692955090935050565b600080602083850312156131e057600080fd5b823567ffffffffffffffff8111156131f757600080fd5b61320385828601612f9f565b90969095509350505050565b60006020828403121561322157600080fd5b813567ffffffffffffffff81111561323857600080fd5b820160608185031215612c3257600080fd5b600080610120838503121561325e57600080fd5b823561326981612d15565b91506132788460208501612ec1565b90509250929050565b6001600160e01b031991909116815260200190565b6000602082840312156132a857600080fd5b8151612c32816130c6565b61ffff81168114612bb357600080fd5b80356130df816132b3565b6000602082840312156132e057600080fd5b8135612c32816132b3565b6001600160501b0381168114612bb357600080fd5b65ffffffffffff81168114612bb357600080fd5b600081356109dd81613300565b600081356109dd816132b3565b600081356109dd816130c6565b8135613346816132eb565b815469ffffffffffffffffffff19166001600160501b03821617825550602082013561337181613300565b815465ffffffffffff60501b19811660509290921b65ffffffffffff60501b16918217835560408401356133a481613300565b65ffffffffffff60801b60809190911b166bffffffffffffffffffffffff60501b19821683178117845560608501356133dc816132b3565b6dffffffffffffffffffffffffffff60501b19929092169092179190911760b09190911b61ffff60b01b1617815561343961341960808401613321565b82805461ffff60c01b191660c09290921b61ffff60c01b16919091179055565b611f5b61344860a0840161332e565b82805460ff60d01b191691151560d01b60ff60d01b16919091179055565b80356130df81613300565b60c081018235613480816132eb565b6001600160501b03168252602083013561349981613300565b65ffffffffffff90811660208401526040840135906134b782613300565b16604083015260608301356134cb816132b3565b61ffff90811660608401526080840135906134e5826132b3565b16608083015260a08301356134f9816130c6565b80151560a08401525092915050565b60006020828403121561351a57600080fd5b8135612c32816130c6565b60006101208201905060018060a01b038416825282356020830152602083013560408301526040830135606083015260608301356080830152608083013560a083015260a083013560c083015260c083013560e083015260e083013561358a816130c6565b801515610100840152509392505050565b60006101008083850312156135af57600080fd5b6040519081019067ffffffffffffffff821181831017156135e057634e487b7160e01b600052604160045260246000fd5b81604052833581526020840135602082015260408401356040820152606084013560608201526080840135608082015260a084013560a082015260c084013560c082015261363060e085016130d4565b60e0820152949350505050565b62ffffff81168114612bb357600080fd5b60006020828403121561366057600080fd5b8135612c328161363d565b64ffffffffff81168114612bb357600080fd5b600081356109dd8161366b565b8135613696816132eb565b815469ffffffffffffffffffff19166001600160501b0382161782555060208201356136c18161363d565b815462ffffff60501b19811660509290921b62ffffff60501b16918217835560408401356136ee8161366b565b67ffffffffffffffff60501b199190911690911760689190911b64ffffffffff60681b1617815560608201356137238161366b565b815464ffffffffff60901b1916609082901b64ffffffffff60901b16178255506137786137526080840161367e565b82805464ffffffffff60b81b191660b89290921b64ffffffffff60b81b16919091179055565b6137a761378760a08401613321565b82805461ffff60e01b191660e09290921b61ffff60e01b16919091179055565b611f5b6137b660c08401613321565b8280546001600160f01b031660f09290921b6001600160f01b031916919091179055565b60e0810182356137e9816132eb565b6001600160501b0316825260208301356138028161363d565b62ffffff166020830152604083013561381a8161366b565b64ffffffffff90811660408401526060840135906138378261366b565b166060830152608083013561384b8161366b565b64ffffffffff16608083015261386360a084016132c3565b61ffff1660a083015261387860c084016132c3565b61ffff811660c0840152612dff565b6000808335601e1984360301811261389e57600080fd5b83018035915067ffffffffffffffff8211156138b957600080fd5b6020019150600581901b3603821315612a6257600080fd5b634e487b7160e01b600052603260045260246000fd5b6000602082840312156138f957600080fd5b8151612c3281612d15565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b602081526000613941602083018486613904565b949350505050565b6000808335601e1984360301811261396057600080fd5b83018035915067ffffffffffffffff82111561397b57600080fd5b602001915036819003821315612a6257600080fd5b6040808252810184905260006060600586901b8301810190830187835b88811015613a2157858403605f190183528135368b9003601e190181126139d357600080fd5b8a01602081810191359067ffffffffffffffff8211156139f257600080fd5b813603831315613a0157600080fd5b613a0c878385613904565b965094850194939093019250506001016139ad565b5050508281036020840152613a37818587613904565b979650505050505050565b60ff81168114612bb357600080fd5b600081356109dd81613a42565b63ffffffff81168114612bb357600080fd5b600081356109dd81613a5e565b8135613a88816132eb565b815469ffffffffffffffffffff19166001600160501b038216178255506020820135613ab3816132b3565b815461ffff60501b19811660509290921b61ffff60501b1691821783556040840135613ade81613300565b67ffffffffffffffff60501b1991909116909117606091821b65ffffffffffff60601b16178255613b3c90613b14908401613314565b82805465ffffffffffff60901b191660909290921b65ffffffffffff60901b16919091179055565b613b69613b4b60808401613a51565b82805460ff60c01b191660c09290921b60ff60c01b16919091179055565b613b9c613b7860a08401613a70565b82805463ffffffff60c81b191660c89290921b63ffffffff60c81b16919091179055565b613bcb613bab60c08401613321565b82805461ffff60e81b191660e89290921b61ffff60e81b16919091179055565b611f5b613bda60e0840161332e565b8280546001600160f81b031691151560f81b6001600160f81b031916919091179055565b80356130df81613a42565b80356130df81613a5e565b61010081018235613c24816132eb565b6001600160501b031682526020830135613c3d816132b3565b61ffff1660208301526040830135613c5481613300565b65ffffffffffff166040830152613c6d60608401613466565b65ffffffffffff166060830152613c8660808401613bfe565b60ff166080830152613c9a60a08401613c09565b63ffffffff1660a0830152613cb160c084016132c3565b61ffff1660c0830152613cc660e084016130d4565b80151560e0840152612dff565b634e487b7160e01b600052601160045260246000fd5b80820281158282048414176109dd576109dd613cd3565b600080600060608486031215613d1557600080fd5b8351925060208401519150604084015190509250925092565b808201808211156109dd576109dd613cd3565b818103818111156109dd576109dd613cd3565b634e487b7160e01b600052603160045260246000fd5b600082613d8757634e487b7160e01b600052601260045260246000fd5b500490565b600060018201613d9e57613d9e613cd3565b5060010190565b634e487b7160e01b600052602160045260246000fdfea26469706673582212200905715ff1de569eb1e3329b7952fd89d5cc94888d882c778623bb39dfc1ff4964736f6c63430008110033",
}

// SeadropABI is the input ABI used to generate the binding from.
// Deprecated: Use SeadropMetaData.ABI instead.
var SeadropABI = SeadropMetaData.ABI

// SeadropBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SeadropMetaData.Bin instead.
var SeadropBin = SeadropMetaData.Bin

// DeploySeadrop deploys a new Ethereum contract, binding an instance of Seadrop to it.
func DeploySeadrop(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Seadrop, error) {
	parsed, err := SeadropMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SeadropBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Seadrop{SeadropCaller: SeadropCaller{contract: contract}, SeadropTransactor: SeadropTransactor{contract: contract}, SeadropFilterer: SeadropFilterer{contract: contract}}, nil
}

// Seadrop is an auto generated Go binding around an Ethereum contract.
type Seadrop struct {
	SeadropCaller     // Read-only binding to the contract
	SeadropTransactor // Write-only binding to the contract
	SeadropFilterer   // Log filterer for contract events
}

// SeadropCaller is an auto generated read-only Go binding around an Ethereum contract.
type SeadropCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeadropTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SeadropTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeadropFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SeadropFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeadropSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SeadropSession struct {
	Contract     *Seadrop          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SeadropCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SeadropCallerSession struct {
	Contract *SeadropCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SeadropTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SeadropTransactorSession struct {
	Contract     *SeadropTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SeadropRaw is an auto generated low-level Go binding around an Ethereum contract.
type SeadropRaw struct {
	Contract *Seadrop // Generic contract binding to access the raw methods on
}

// SeadropCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SeadropCallerRaw struct {
	Contract *SeadropCaller // Generic read-only contract binding to access the raw methods on
}

// SeadropTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SeadropTransactorRaw struct {
	Contract *SeadropTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSeadrop creates a new instance of Seadrop, bound to a specific deployed contract.
func NewSeadrop(address common.Address, backend bind.ContractBackend) (*Seadrop, error) {
	contract, err := bindSeadrop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Seadrop{SeadropCaller: SeadropCaller{contract: contract}, SeadropTransactor: SeadropTransactor{contract: contract}, SeadropFilterer: SeadropFilterer{contract: contract}}, nil
}

// NewSeadropCaller creates a new read-only instance of Seadrop, bound to a specific deployed contract.
func NewSeadropCaller(address common.Address, caller bind.ContractCaller) (*SeadropCaller, error) {
	contract, err := bindSeadrop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SeadropCaller{contract: contract}, nil
}

// NewSeadropTransactor creates a new write-only instance of Seadrop, bound to a specific deployed contract.
func NewSeadropTransactor(address common.Address, transactor bind.ContractTransactor) (*SeadropTransactor, error) {
	contract, err := bindSeadrop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SeadropTransactor{contract: contract}, nil
}

// NewSeadropFilterer creates a new log filterer instance of Seadrop, bound to a specific deployed contract.
func NewSeadropFilterer(address common.Address, filterer bind.ContractFilterer) (*SeadropFilterer, error) {
	contract, err := bindSeadrop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SeadropFilterer{contract: contract}, nil
}

// bindSeadrop binds a generic wrapper to an already deployed contract.
func bindSeadrop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SeadropABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Seadrop *SeadropRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Seadrop.Contract.SeadropCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Seadrop *SeadropRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seadrop.Contract.SeadropTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Seadrop *SeadropRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Seadrop.Contract.SeadropTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Seadrop *SeadropCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Seadrop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Seadrop *SeadropTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seadrop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Seadrop *SeadropTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Seadrop.Contract.contract.Transact(opts, method, params...)
}

// GetAllowListMerkleRoot is a free data retrieval call binding the contract method 0x32bf11f5.
//
// Solidity: function getAllowListMerkleRoot(address nftContract) view returns(bytes32)
func (_Seadrop *SeadropCaller) GetAllowListMerkleRoot(opts *bind.CallOpts, nftContract common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Seadrop.contract.Call(opts, &out, "getAllowListMerkleRoot", nftContract)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAllowListMerkleRoot is a free data retrieval call binding the contract method 0x32bf11f5.
//
// Solidity: function getAllowListMerkleRoot(address nftContract) view returns(bytes32)
func (_Seadrop *SeadropSession) GetAllowListMerkleRoot(nftContract common.Address) ([32]byte, error) {
	return _Seadrop.Contract.GetAllowListMerkleRoot(&_Seadrop.CallOpts, nftContract)
}

// GetAllowListMerkleRoot is a free data retrieval call binding the contract method 0x32bf11f5.
//
// Solidity: function getAllowListMerkleRoot(address nftContract) view returns(bytes32)
func (_Seadrop *SeadropCallerSession) GetAllowListMerkleRoot(nftContract common.Address) ([32]byte, error) {
	return _Seadrop.Contract.GetAllowListMerkleRoot(&_Seadrop.CallOpts, nftContract)
}

// GetAllowedFeeRecipients is a free data retrieval call binding the contract method 0x68632274.
//
// Solidity: function getAllowedFeeRecipients(address nftContract) view returns(address[])
func (_Seadrop *SeadropCaller) GetAllowedFeeRecipients(opts *bind.CallOpts, nftContract common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Seadrop.contract.Call(opts, &out, "getAllowedFeeRecipients", nftContract)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllowedFeeRecipients is a free data retrieval call binding the contract method 0x68632274.
//
// Solidity: function getAllowedFeeRecipients(address nftContract) view returns(address[])
func (_Seadrop *SeadropSession) GetAllowedFeeRecipients(nftContract common.Address) ([]common.Address, error) {
	return _Seadrop.Contract.GetAllowedFeeRecipients(&_Seadrop.CallOpts, nftContract)
}

// GetAllowedFeeRecipients is a free data retrieval call binding the contract method 0x68632274.
//
// Solidity: function getAllowedFeeRecipients(address nftContract) view returns(address[])
func (_Seadrop *SeadropCallerSession) GetAllowedFeeRecipients(nftContract common.Address) ([]common.Address, error) {
	return _Seadrop.Contract.GetAllowedFeeRecipients(&_Seadrop.CallOpts, nftContract)
}

// GetAllowedNftTokenIdIsRedeemed is a free data retrieval call binding the contract method 0x88aa3d37.
//
// Solidity: function getAllowedNftTokenIdIsRedeemed(address nftContract, address allowedNftToken, uint256 allowedNftTokenId) view returns(bool)
func (_Seadrop *SeadropCaller) GetAllowedNftTokenIdIsRedeemed(opts *bind.CallOpts, nftContract common.Address, allowedNftToken common.Address, allowedNftTokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Seadrop.contract.Call(opts, &out, "getAllowedNftTokenIdIsRedeemed", nftContract, allowedNftToken, allowedNftTokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetAllowedNftTokenIdIsRedeemed is a free data retrieval call binding the contract method 0x88aa3d37.
//
// Solidity: function getAllowedNftTokenIdIsRedeemed(address nftContract, address allowedNftToken, uint256 allowedNftTokenId) view returns(bool)
func (_Seadrop *SeadropSession) GetAllowedNftTokenIdIsRedeemed(nftContract common.Address, allowedNftToken common.Address, allowedNftTokenId *big.Int) (bool, error) {
	return _Seadrop.Contract.GetAllowedNftTokenIdIsRedeemed(&_Seadrop.CallOpts, nftContract, allowedNftToken, allowedNftTokenId)
}

// GetAllowedNftTokenIdIsRedeemed is a free data retrieval call binding the contract method 0x88aa3d37.
//
// Solidity: function getAllowedNftTokenIdIsRedeemed(address nftContract, address allowedNftToken, uint256 allowedNftTokenId) view returns(bool)
func (_Seadrop *SeadropCallerSession) GetAllowedNftTokenIdIsRedeemed(nftContract common.Address, allowedNftToken common.Address, allowedNftTokenId *big.Int) (bool, error) {
	return _Seadrop.Contract.GetAllowedNftTokenIdIsRedeemed(&_Seadrop.CallOpts, nftContract, allowedNftToken, allowedNftTokenId)
}

// GetCreatorPayoutAddress is a free data retrieval call binding the contract method 0x5cb3c4d3.
//
// Solidity: function getCreatorPayoutAddress(address nftContract) view returns(address)
func (_Seadrop *SeadropCaller) GetCreatorPayoutAddress(opts *bind.CallOpts, nftContract common.Address) (common.Address, error) {
	var out []interface{}
	err := _Seadrop.contract.Call(opts, &out, "getCreatorPayoutAddress", nftContract)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCreatorPayoutAddress is a free data retrieval call binding the contract method 0x5cb3c4d3.
//
// Solidity: function getCreatorPayoutAddress(address nftContract) view returns(address)
func (_Seadrop *SeadropSession) GetCreatorPayoutAddress(nftContract common.Address) (common.Address, error) {
	return _Seadrop.Contract.GetCreatorPayoutAddress(&_Seadrop.CallOpts, nftContract)
}

// GetCreatorPayoutAddress is a free data retrieval call binding the contract method 0x5cb3c4d3.
//
// Solidity: function getCreatorPayoutAddress(address nftContract) view returns(address)
func (_Seadrop *SeadropCallerSession) GetCreatorPayoutAddress(nftContract common.Address) (common.Address, error) {
	return _Seadrop.Contract.GetCreatorPayoutAddress(&_Seadrop.CallOpts, nftContract)
}

// GetFeeRecipientIsAllowed is a free data retrieval call binding the contract method 0x322e75d1.
//
// Solidity: function getFeeRecipientIsAllowed(address nftContract, address feeRecipient) view returns(bool)
func (_Seadrop *SeadropCaller) GetFeeRecipientIsAllowed(opts *bind.CallOpts, nftContract common.Address, feeRecipient common.Address) (bool, error) {
	var out []interface{}
	err := _Seadrop.contract.Call(opts, &out, "getFeeRecipientIsAllowed", nftContract, feeRecipient)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeeRecipientIsAllowed is a free data retrieval call binding the contract method 0x322e75d1.
//
// Solidity: function getFeeRecipientIsAllowed(address nftContract, address feeRecipient) view returns(bool)
func (_Seadrop *SeadropSession) GetFeeRecipientIsAllowed(nftContract common.Address, feeRecipient common.Address) (bool, error) {
	return _Seadrop.Contract.GetFeeRecipientIsAllowed(&_Seadrop.CallOpts, nftContract, feeRecipient)
}

// GetFeeRecipientIsAllowed is a free data retrieval call binding the contract method 0x322e75d1.
//
// Solidity: function getFeeRecipientIsAllowed(address nftContract, address feeRecipient) view returns(bool)
func (_Seadrop *SeadropCallerSession) GetFeeRecipientIsAllowed(nftContract common.Address, feeRecipient common.Address) (bool, error) {
	return _Seadrop.Contract.GetFeeRecipientIsAllowed(&_Seadrop.CallOpts, nftContract, feeRecipient)
}

// GetPayerIsAllowed is a free data retrieval call binding the contract method 0xe583141d.
//
// Solidity: function getPayerIsAllowed(address nftContract, address payer) view returns(bool)
func (_Seadrop *SeadropCaller) GetPayerIsAllowed(opts *bind.CallOpts, nftContract common.Address, payer common.Address) (bool, error) {
	var out []interface{}
	err := _Seadrop.contract.Call(opts, &out, "getPayerIsAllowed", nftContract, payer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetPayerIsAllowed is a free data retrieval call binding the contract method 0xe583141d.
//
// Solidity: function getPayerIsAllowed(address nftContract, address payer) view returns(bool)
func (_Seadrop *SeadropSession) GetPayerIsAllowed(nftContract common.Address, payer common.Address) (bool, error) {
	return _Seadrop.Contract.GetPayerIsAllowed(&_Seadrop.CallOpts, nftContract, payer)
}

// GetPayerIsAllowed is a free data retrieval call binding the contract method 0xe583141d.
//
// Solidity: function getPayerIsAllowed(address nftContract, address payer) view returns(bool)
func (_Seadrop *SeadropCallerSession) GetPayerIsAllowed(nftContract common.Address, payer common.Address) (bool, error) {
	return _Seadrop.Contract.GetPayerIsAllowed(&_Seadrop.CallOpts, nftContract, payer)
}

// GetPayers is a free data retrieval call binding the contract method 0x7c35b982.
//
// Solidity: function getPayers(address nftContract) view returns(address[])
func (_Seadrop *SeadropCaller) GetPayers(opts *bind.CallOpts, nftContract common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Seadrop.contract.Call(opts, &out, "getPayers", nftContract)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPayers is a free data retrieval call binding the contract method 0x7c35b982.
//
// Solidity: function getPayers(address nftContract) view returns(address[])
func (_Seadrop *SeadropSession) GetPayers(nftContract common.Address) ([]common.Address, error) {
	return _Seadrop.Contract.GetPayers(&_Seadrop.CallOpts, nftContract)
}

// GetPayers is a free data retrieval call binding the contract method 0x7c35b982.
//
// Solidity: function getPayers(address nftContract) view returns(address[])
func (_Seadrop *SeadropCallerSession) GetPayers(nftContract common.Address) ([]common.Address, error) {
	return _Seadrop.Contract.GetPayers(&_Seadrop.CallOpts, nftContract)
}

// GetPublicDrop is a free data retrieval call binding the contract method 0xbc6a629c.
//
// Solidity: function getPublicDrop(address nftContract) view returns((uint80,uint48,uint48,uint16,uint16,bool))
func (_Seadrop *SeadropCaller) GetPublicDrop(opts *bind.CallOpts, nftContract common.Address) (PublicDrop, error) {
	var out []interface{}
	err := _Seadrop.contract.Call(opts, &out, "getPublicDrop", nftContract)

	if err != nil {
		return *new(PublicDrop), err
	}

	out0 := *abi.ConvertType(out[0], new(PublicDrop)).(*PublicDrop)

	return out0, err

}

// GetPublicDrop is a free data retrieval call binding the contract method 0xbc6a629c.
//
// Solidity: function getPublicDrop(address nftContract) view returns((uint80,uint48,uint48,uint16,uint16,bool))
func (_Seadrop *SeadropSession) GetPublicDrop(nftContract common.Address) (PublicDrop, error) {
	return _Seadrop.Contract.GetPublicDrop(&_Seadrop.CallOpts, nftContract)
}

// GetPublicDrop is a free data retrieval call binding the contract method 0xbc6a629c.
//
// Solidity: function getPublicDrop(address nftContract) view returns((uint80,uint48,uint48,uint16,uint16,bool))
func (_Seadrop *SeadropCallerSession) GetPublicDrop(nftContract common.Address) (PublicDrop, error) {
	return _Seadrop.Contract.GetPublicDrop(&_Seadrop.CallOpts, nftContract)
}

// GetSignedMintValidationParams is a free data retrieval call binding the contract method 0x81bf9af3.
//
// Solidity: function getSignedMintValidationParams(address nftContract, address signer) view returns((uint80,uint24,uint40,uint40,uint40,uint16,uint16))
func (_Seadrop *SeadropCaller) GetSignedMintValidationParams(opts *bind.CallOpts, nftContract common.Address, signer common.Address) (SignedMintValidationParams, error) {
	var out []interface{}
	err := _Seadrop.contract.Call(opts, &out, "getSignedMintValidationParams", nftContract, signer)

	if err != nil {
		return *new(SignedMintValidationParams), err
	}

	out0 := *abi.ConvertType(out[0], new(SignedMintValidationParams)).(*SignedMintValidationParams)

	return out0, err

}

// GetSignedMintValidationParams is a free data retrieval call binding the contract method 0x81bf9af3.
//
// Solidity: function getSignedMintValidationParams(address nftContract, address signer) view returns((uint80,uint24,uint40,uint40,uint40,uint16,uint16))
func (_Seadrop *SeadropSession) GetSignedMintValidationParams(nftContract common.Address, signer common.Address) (SignedMintValidationParams, error) {
	return _Seadrop.Contract.GetSignedMintValidationParams(&_Seadrop.CallOpts, nftContract, signer)
}

// GetSignedMintValidationParams is a free data retrieval call binding the contract method 0x81bf9af3.
//
// Solidity: function getSignedMintValidationParams(address nftContract, address signer) view returns((uint80,uint24,uint40,uint40,uint40,uint16,uint16))
func (_Seadrop *SeadropCallerSession) GetSignedMintValidationParams(nftContract common.Address, signer common.Address) (SignedMintValidationParams, error) {
	return _Seadrop.Contract.GetSignedMintValidationParams(&_Seadrop.CallOpts, nftContract, signer)
}

// GetSigners is a free data retrieval call binding the contract method 0x7e3ba6af.
//
// Solidity: function getSigners(address nftContract) view returns(address[])
func (_Seadrop *SeadropCaller) GetSigners(opts *bind.CallOpts, nftContract common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Seadrop.contract.Call(opts, &out, "getSigners", nftContract)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSigners is a free data retrieval call binding the contract method 0x7e3ba6af.
//
// Solidity: function getSigners(address nftContract) view returns(address[])
func (_Seadrop *SeadropSession) GetSigners(nftContract common.Address) ([]common.Address, error) {
	return _Seadrop.Contract.GetSigners(&_Seadrop.CallOpts, nftContract)
}

// GetSigners is a free data retrieval call binding the contract method 0x7e3ba6af.
//
// Solidity: function getSigners(address nftContract) view returns(address[])
func (_Seadrop *SeadropCallerSession) GetSigners(nftContract common.Address) ([]common.Address, error) {
	return _Seadrop.Contract.GetSigners(&_Seadrop.CallOpts, nftContract)
}

// GetTokenGatedAllowedTokens is a free data retrieval call binding the contract method 0x2db526eb.
//
// Solidity: function getTokenGatedAllowedTokens(address nftContract) view returns(address[])
func (_Seadrop *SeadropCaller) GetTokenGatedAllowedTokens(opts *bind.CallOpts, nftContract common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Seadrop.contract.Call(opts, &out, "getTokenGatedAllowedTokens", nftContract)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenGatedAllowedTokens is a free data retrieval call binding the contract method 0x2db526eb.
//
// Solidity: function getTokenGatedAllowedTokens(address nftContract) view returns(address[])
func (_Seadrop *SeadropSession) GetTokenGatedAllowedTokens(nftContract common.Address) ([]common.Address, error) {
	return _Seadrop.Contract.GetTokenGatedAllowedTokens(&_Seadrop.CallOpts, nftContract)
}

// GetTokenGatedAllowedTokens is a free data retrieval call binding the contract method 0x2db526eb.
//
// Solidity: function getTokenGatedAllowedTokens(address nftContract) view returns(address[])
func (_Seadrop *SeadropCallerSession) GetTokenGatedAllowedTokens(nftContract common.Address) ([]common.Address, error) {
	return _Seadrop.Contract.GetTokenGatedAllowedTokens(&_Seadrop.CallOpts, nftContract)
}

// GetTokenGatedDrop is a free data retrieval call binding the contract method 0x0b0e8a6e.
//
// Solidity: function getTokenGatedDrop(address nftContract, address allowedNftToken) view returns((uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool))
func (_Seadrop *SeadropCaller) GetTokenGatedDrop(opts *bind.CallOpts, nftContract common.Address, allowedNftToken common.Address) (TokenGatedDropStage, error) {
	var out []interface{}
	err := _Seadrop.contract.Call(opts, &out, "getTokenGatedDrop", nftContract, allowedNftToken)

	if err != nil {
		return *new(TokenGatedDropStage), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenGatedDropStage)).(*TokenGatedDropStage)

	return out0, err

}

// GetTokenGatedDrop is a free data retrieval call binding the contract method 0x0b0e8a6e.
//
// Solidity: function getTokenGatedDrop(address nftContract, address allowedNftToken) view returns((uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool))
func (_Seadrop *SeadropSession) GetTokenGatedDrop(nftContract common.Address, allowedNftToken common.Address) (TokenGatedDropStage, error) {
	return _Seadrop.Contract.GetTokenGatedDrop(&_Seadrop.CallOpts, nftContract, allowedNftToken)
}

// GetTokenGatedDrop is a free data retrieval call binding the contract method 0x0b0e8a6e.
//
// Solidity: function getTokenGatedDrop(address nftContract, address allowedNftToken) view returns((uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool))
func (_Seadrop *SeadropCallerSession) GetTokenGatedDrop(nftContract common.Address, allowedNftToken common.Address) (TokenGatedDropStage, error) {
	return _Seadrop.Contract.GetTokenGatedDrop(&_Seadrop.CallOpts, nftContract, allowedNftToken)
}

// MintAllowList is a paid mutator transaction binding the contract method 0x4300a4e6.
//
// Solidity: function mintAllowList(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, bytes32[] proof) payable returns()
func (_Seadrop *SeadropTransactor) MintAllowList(opts *bind.TransactOpts, nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, proof [][32]byte) (*types.Transaction, error) {
	return _Seadrop.contract.Transact(opts, "mintAllowList", nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, proof)
}

// MintAllowList is a paid mutator transaction binding the contract method 0x4300a4e6.
//
// Solidity: function mintAllowList(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, bytes32[] proof) payable returns()
func (_Seadrop *SeadropSession) MintAllowList(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, proof [][32]byte) (*types.Transaction, error) {
	return _Seadrop.Contract.MintAllowList(&_Seadrop.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, proof)
}

// MintAllowList is a paid mutator transaction binding the contract method 0x4300a4e6.
//
// Solidity: function mintAllowList(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, bytes32[] proof) payable returns()
func (_Seadrop *SeadropTransactorSession) MintAllowList(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, proof [][32]byte) (*types.Transaction, error) {
	return _Seadrop.Contract.MintAllowList(&_Seadrop.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, proof)
}

// MintAllowedTokenHolder is a paid mutator transaction binding the contract method 0x99eb900f.
//
// Solidity: function mintAllowedTokenHolder(address nftContract, address feeRecipient, address minterIfNotPayer, (address,uint256[]) mintParams) payable returns()
func (_Seadrop *SeadropTransactor) MintAllowedTokenHolder(opts *bind.TransactOpts, nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, mintParams TokenGatedMintParams) (*types.Transaction, error) {
	return _Seadrop.contract.Transact(opts, "mintAllowedTokenHolder", nftContract, feeRecipient, minterIfNotPayer, mintParams)
}

// MintAllowedTokenHolder is a paid mutator transaction binding the contract method 0x99eb900f.
//
// Solidity: function mintAllowedTokenHolder(address nftContract, address feeRecipient, address minterIfNotPayer, (address,uint256[]) mintParams) payable returns()
func (_Seadrop *SeadropSession) MintAllowedTokenHolder(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, mintParams TokenGatedMintParams) (*types.Transaction, error) {
	return _Seadrop.Contract.MintAllowedTokenHolder(&_Seadrop.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, mintParams)
}

// MintAllowedTokenHolder is a paid mutator transaction binding the contract method 0x99eb900f.
//
// Solidity: function mintAllowedTokenHolder(address nftContract, address feeRecipient, address minterIfNotPayer, (address,uint256[]) mintParams) payable returns()
func (_Seadrop *SeadropTransactorSession) MintAllowedTokenHolder(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, mintParams TokenGatedMintParams) (*types.Transaction, error) {
	return _Seadrop.Contract.MintAllowedTokenHolder(&_Seadrop.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, mintParams)
}

// MintPublic is a paid mutator transaction binding the contract method 0x161ac21f.
//
// Solidity: function mintPublic(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity) payable returns()
func (_Seadrop *SeadropTransactor) MintPublic(opts *bind.TransactOpts, nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Seadrop.contract.Transact(opts, "mintPublic", nftContract, feeRecipient, minterIfNotPayer, quantity)
}

// MintPublic is a paid mutator transaction binding the contract method 0x161ac21f.
//
// Solidity: function mintPublic(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity) payable returns()
func (_Seadrop *SeadropSession) MintPublic(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Seadrop.Contract.MintPublic(&_Seadrop.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity)
}

// MintPublic is a paid mutator transaction binding the contract method 0x161ac21f.
//
// Solidity: function mintPublic(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity) payable returns()
func (_Seadrop *SeadropTransactorSession) MintPublic(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Seadrop.Contract.MintPublic(&_Seadrop.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity)
}

// MintSigned is a paid mutator transaction binding the contract method 0x4b61cd6f.
//
// Solidity: function mintSigned(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, uint256 salt, bytes signature) payable returns()
func (_Seadrop *SeadropTransactor) MintSigned(opts *bind.TransactOpts, nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Seadrop.contract.Transact(opts, "mintSigned", nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, salt, signature)
}

// MintSigned is a paid mutator transaction binding the contract method 0x4b61cd6f.
//
// Solidity: function mintSigned(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, uint256 salt, bytes signature) payable returns()
func (_Seadrop *SeadropSession) MintSigned(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Seadrop.Contract.MintSigned(&_Seadrop.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, salt, signature)
}

// MintSigned is a paid mutator transaction binding the contract method 0x4b61cd6f.
//
// Solidity: function mintSigned(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, uint256 salt, bytes signature) payable returns()
func (_Seadrop *SeadropTransactorSession) MintSigned(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Seadrop.Contract.MintSigned(&_Seadrop.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, salt, signature)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0xebb4a55f.
//
// Solidity: function updateAllowList((bytes32,string[],string) allowListData) returns()
func (_Seadrop *SeadropTransactor) UpdateAllowList(opts *bind.TransactOpts, allowListData AllowListData) (*types.Transaction, error) {
	return _Seadrop.contract.Transact(opts, "updateAllowList", allowListData)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0xebb4a55f.
//
// Solidity: function updateAllowList((bytes32,string[],string) allowListData) returns()
func (_Seadrop *SeadropSession) UpdateAllowList(allowListData AllowListData) (*types.Transaction, error) {
	return _Seadrop.Contract.UpdateAllowList(&_Seadrop.TransactOpts, allowListData)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0xebb4a55f.
//
// Solidity: function updateAllowList((bytes32,string[],string) allowListData) returns()
func (_Seadrop *SeadropTransactorSession) UpdateAllowList(allowListData AllowListData) (*types.Transaction, error) {
	return _Seadrop.Contract.UpdateAllowList(&_Seadrop.TransactOpts, allowListData)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x8e7d1e43.
//
// Solidity: function updateAllowedFeeRecipient(address feeRecipient, bool allowed) returns()
func (_Seadrop *SeadropTransactor) UpdateAllowedFeeRecipient(opts *bind.TransactOpts, feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _Seadrop.contract.Transact(opts, "updateAllowedFeeRecipient", feeRecipient, allowed)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x8e7d1e43.
//
// Solidity: function updateAllowedFeeRecipient(address feeRecipient, bool allowed) returns()
func (_Seadrop *SeadropSession) UpdateAllowedFeeRecipient(feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _Seadrop.Contract.UpdateAllowedFeeRecipient(&_Seadrop.TransactOpts, feeRecipient, allowed)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x8e7d1e43.
//
// Solidity: function updateAllowedFeeRecipient(address feeRecipient, bool allowed) returns()
func (_Seadrop *SeadropTransactorSession) UpdateAllowedFeeRecipient(feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _Seadrop.Contract.UpdateAllowedFeeRecipient(&_Seadrop.TransactOpts, feeRecipient, allowed)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x12738db8.
//
// Solidity: function updateCreatorPayoutAddress(address payoutAddress) returns()
func (_Seadrop *SeadropTransactor) UpdateCreatorPayoutAddress(opts *bind.TransactOpts, payoutAddress common.Address) (*types.Transaction, error) {
	return _Seadrop.contract.Transact(opts, "updateCreatorPayoutAddress", payoutAddress)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x12738db8.
//
// Solidity: function updateCreatorPayoutAddress(address payoutAddress) returns()
func (_Seadrop *SeadropSession) UpdateCreatorPayoutAddress(payoutAddress common.Address) (*types.Transaction, error) {
	return _Seadrop.Contract.UpdateCreatorPayoutAddress(&_Seadrop.TransactOpts, payoutAddress)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x12738db8.
//
// Solidity: function updateCreatorPayoutAddress(address payoutAddress) returns()
func (_Seadrop *SeadropTransactorSession) UpdateCreatorPayoutAddress(payoutAddress common.Address) (*types.Transaction, error) {
	return _Seadrop.Contract.UpdateCreatorPayoutAddress(&_Seadrop.TransactOpts, payoutAddress)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0xb957d0cb.
//
// Solidity: function updateDropURI(string dropURI) returns()
func (_Seadrop *SeadropTransactor) UpdateDropURI(opts *bind.TransactOpts, dropURI string) (*types.Transaction, error) {
	return _Seadrop.contract.Transact(opts, "updateDropURI", dropURI)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0xb957d0cb.
//
// Solidity: function updateDropURI(string dropURI) returns()
func (_Seadrop *SeadropSession) UpdateDropURI(dropURI string) (*types.Transaction, error) {
	return _Seadrop.Contract.UpdateDropURI(&_Seadrop.TransactOpts, dropURI)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0xb957d0cb.
//
// Solidity: function updateDropURI(string dropURI) returns()
func (_Seadrop *SeadropTransactorSession) UpdateDropURI(dropURI string) (*types.Transaction, error) {
	return _Seadrop.Contract.UpdateDropURI(&_Seadrop.TransactOpts, dropURI)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0x7f2a5cca.
//
// Solidity: function updatePayer(address payer, bool allowed) returns()
func (_Seadrop *SeadropTransactor) UpdatePayer(opts *bind.TransactOpts, payer common.Address, allowed bool) (*types.Transaction, error) {
	return _Seadrop.contract.Transact(opts, "updatePayer", payer, allowed)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0x7f2a5cca.
//
// Solidity: function updatePayer(address payer, bool allowed) returns()
func (_Seadrop *SeadropSession) UpdatePayer(payer common.Address, allowed bool) (*types.Transaction, error) {
	return _Seadrop.Contract.UpdatePayer(&_Seadrop.TransactOpts, payer, allowed)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0x7f2a5cca.
//
// Solidity: function updatePayer(address payer, bool allowed) returns()
func (_Seadrop *SeadropTransactorSession) UpdatePayer(payer common.Address, allowed bool) (*types.Transaction, error) {
	return _Seadrop.Contract.UpdatePayer(&_Seadrop.TransactOpts, payer, allowed)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x01308e65.
//
// Solidity: function updatePublicDrop((uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_Seadrop *SeadropTransactor) UpdatePublicDrop(opts *bind.TransactOpts, publicDrop PublicDrop) (*types.Transaction, error) {
	return _Seadrop.contract.Transact(opts, "updatePublicDrop", publicDrop)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x01308e65.
//
// Solidity: function updatePublicDrop((uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_Seadrop *SeadropSession) UpdatePublicDrop(publicDrop PublicDrop) (*types.Transaction, error) {
	return _Seadrop.Contract.UpdatePublicDrop(&_Seadrop.TransactOpts, publicDrop)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x01308e65.
//
// Solidity: function updatePublicDrop((uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_Seadrop *SeadropTransactorSession) UpdatePublicDrop(publicDrop PublicDrop) (*types.Transaction, error) {
	return _Seadrop.Contract.UpdatePublicDrop(&_Seadrop.TransactOpts, publicDrop)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x4d380178.
//
// Solidity: function updateSignedMintValidationParams(address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_Seadrop *SeadropTransactor) UpdateSignedMintValidationParams(opts *bind.TransactOpts, signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _Seadrop.contract.Transact(opts, "updateSignedMintValidationParams", signer, signedMintValidationParams)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x4d380178.
//
// Solidity: function updateSignedMintValidationParams(address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_Seadrop *SeadropSession) UpdateSignedMintValidationParams(signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _Seadrop.Contract.UpdateSignedMintValidationParams(&_Seadrop.TransactOpts, signer, signedMintValidationParams)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x4d380178.
//
// Solidity: function updateSignedMintValidationParams(address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_Seadrop *SeadropTransactorSession) UpdateSignedMintValidationParams(signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _Seadrop.Contract.UpdateSignedMintValidationParams(&_Seadrop.TransactOpts, signer, signedMintValidationParams)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0xfd9ab22a.
//
// Solidity: function updateTokenGatedDrop(address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_Seadrop *SeadropTransactor) UpdateTokenGatedDrop(opts *bind.TransactOpts, allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _Seadrop.contract.Transact(opts, "updateTokenGatedDrop", allowedNftToken, dropStage)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0xfd9ab22a.
//
// Solidity: function updateTokenGatedDrop(address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_Seadrop *SeadropSession) UpdateTokenGatedDrop(allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _Seadrop.Contract.UpdateTokenGatedDrop(&_Seadrop.TransactOpts, allowedNftToken, dropStage)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0xfd9ab22a.
//
// Solidity: function updateTokenGatedDrop(address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_Seadrop *SeadropTransactorSession) UpdateTokenGatedDrop(allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _Seadrop.Contract.UpdateTokenGatedDrop(&_Seadrop.TransactOpts, allowedNftToken, dropStage)
}

// SeadropAllowListUpdatedIterator is returned from FilterAllowListUpdated and is used to iterate over the raw logs and unpacked data for AllowListUpdated events raised by the Seadrop contract.
type SeadropAllowListUpdatedIterator struct {
	Event *SeadropAllowListUpdated // Event containing the contract specifics and raw log

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
func (it *SeadropAllowListUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeadropAllowListUpdated)
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
		it.Event = new(SeadropAllowListUpdated)
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
func (it *SeadropAllowListUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeadropAllowListUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeadropAllowListUpdated represents a AllowListUpdated event raised by the Seadrop contract.
type SeadropAllowListUpdated struct {
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
func (_Seadrop *SeadropFilterer) FilterAllowListUpdated(opts *bind.FilterOpts, nftContract []common.Address, previousMerkleRoot [][32]byte, newMerkleRoot [][32]byte) (*SeadropAllowListUpdatedIterator, error) {

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

	logs, sub, err := _Seadrop.contract.FilterLogs(opts, "AllowListUpdated", nftContractRule, previousMerkleRootRule, newMerkleRootRule)
	if err != nil {
		return nil, err
	}
	return &SeadropAllowListUpdatedIterator{contract: _Seadrop.contract, event: "AllowListUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowListUpdated is a free log subscription operation binding the contract event 0xefcd7e019bc8b47d27881fd59e2619280ca5894f285950f10ab049870652efa5.
//
// Solidity: event AllowListUpdated(address indexed nftContract, bytes32 indexed previousMerkleRoot, bytes32 indexed newMerkleRoot, string[] publicKeyURI, string allowListURI)
func (_Seadrop *SeadropFilterer) WatchAllowListUpdated(opts *bind.WatchOpts, sink chan<- *SeadropAllowListUpdated, nftContract []common.Address, previousMerkleRoot [][32]byte, newMerkleRoot [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Seadrop.contract.WatchLogs(opts, "AllowListUpdated", nftContractRule, previousMerkleRootRule, newMerkleRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeadropAllowListUpdated)
				if err := _Seadrop.contract.UnpackLog(event, "AllowListUpdated", log); err != nil {
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
func (_Seadrop *SeadropFilterer) ParseAllowListUpdated(log types.Log) (*SeadropAllowListUpdated, error) {
	event := new(SeadropAllowListUpdated)
	if err := _Seadrop.contract.UnpackLog(event, "AllowListUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeadropAllowedFeeRecipientUpdatedIterator is returned from FilterAllowedFeeRecipientUpdated and is used to iterate over the raw logs and unpacked data for AllowedFeeRecipientUpdated events raised by the Seadrop contract.
type SeadropAllowedFeeRecipientUpdatedIterator struct {
	Event *SeadropAllowedFeeRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *SeadropAllowedFeeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeadropAllowedFeeRecipientUpdated)
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
		it.Event = new(SeadropAllowedFeeRecipientUpdated)
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
func (it *SeadropAllowedFeeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeadropAllowedFeeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeadropAllowedFeeRecipientUpdated represents a AllowedFeeRecipientUpdated event raised by the Seadrop contract.
type SeadropAllowedFeeRecipientUpdated struct {
	NftContract  common.Address
	FeeRecipient common.Address
	Allowed      bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAllowedFeeRecipientUpdated is a free log retrieval operation binding the contract event 0x6486c31f9d664e241acf94ec2541d328f6b9e97257ae16a1d887f296f879719f.
//
// Solidity: event AllowedFeeRecipientUpdated(address indexed nftContract, address indexed feeRecipient, bool indexed allowed)
func (_Seadrop *SeadropFilterer) FilterAllowedFeeRecipientUpdated(opts *bind.FilterOpts, nftContract []common.Address, feeRecipient []common.Address, allowed []bool) (*SeadropAllowedFeeRecipientUpdatedIterator, error) {

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

	logs, sub, err := _Seadrop.contract.FilterLogs(opts, "AllowedFeeRecipientUpdated", nftContractRule, feeRecipientRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &SeadropAllowedFeeRecipientUpdatedIterator{contract: _Seadrop.contract, event: "AllowedFeeRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedFeeRecipientUpdated is a free log subscription operation binding the contract event 0x6486c31f9d664e241acf94ec2541d328f6b9e97257ae16a1d887f296f879719f.
//
// Solidity: event AllowedFeeRecipientUpdated(address indexed nftContract, address indexed feeRecipient, bool indexed allowed)
func (_Seadrop *SeadropFilterer) WatchAllowedFeeRecipientUpdated(opts *bind.WatchOpts, sink chan<- *SeadropAllowedFeeRecipientUpdated, nftContract []common.Address, feeRecipient []common.Address, allowed []bool) (event.Subscription, error) {

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

	logs, sub, err := _Seadrop.contract.WatchLogs(opts, "AllowedFeeRecipientUpdated", nftContractRule, feeRecipientRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeadropAllowedFeeRecipientUpdated)
				if err := _Seadrop.contract.UnpackLog(event, "AllowedFeeRecipientUpdated", log); err != nil {
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
func (_Seadrop *SeadropFilterer) ParseAllowedFeeRecipientUpdated(log types.Log) (*SeadropAllowedFeeRecipientUpdated, error) {
	event := new(SeadropAllowedFeeRecipientUpdated)
	if err := _Seadrop.contract.UnpackLog(event, "AllowedFeeRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeadropCreatorPayoutAddressUpdatedIterator is returned from FilterCreatorPayoutAddressUpdated and is used to iterate over the raw logs and unpacked data for CreatorPayoutAddressUpdated events raised by the Seadrop contract.
type SeadropCreatorPayoutAddressUpdatedIterator struct {
	Event *SeadropCreatorPayoutAddressUpdated // Event containing the contract specifics and raw log

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
func (it *SeadropCreatorPayoutAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeadropCreatorPayoutAddressUpdated)
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
		it.Event = new(SeadropCreatorPayoutAddressUpdated)
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
func (it *SeadropCreatorPayoutAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeadropCreatorPayoutAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeadropCreatorPayoutAddressUpdated represents a CreatorPayoutAddressUpdated event raised by the Seadrop contract.
type SeadropCreatorPayoutAddressUpdated struct {
	NftContract      common.Address
	NewPayoutAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCreatorPayoutAddressUpdated is a free log retrieval operation binding the contract event 0x0c69f21751e800ea5960436c9a94370c7adbf54c733a20a025293fbbe8f16252.
//
// Solidity: event CreatorPayoutAddressUpdated(address indexed nftContract, address indexed newPayoutAddress)
func (_Seadrop *SeadropFilterer) FilterCreatorPayoutAddressUpdated(opts *bind.FilterOpts, nftContract []common.Address, newPayoutAddress []common.Address) (*SeadropCreatorPayoutAddressUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var newPayoutAddressRule []interface{}
	for _, newPayoutAddressItem := range newPayoutAddress {
		newPayoutAddressRule = append(newPayoutAddressRule, newPayoutAddressItem)
	}

	logs, sub, err := _Seadrop.contract.FilterLogs(opts, "CreatorPayoutAddressUpdated", nftContractRule, newPayoutAddressRule)
	if err != nil {
		return nil, err
	}
	return &SeadropCreatorPayoutAddressUpdatedIterator{contract: _Seadrop.contract, event: "CreatorPayoutAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchCreatorPayoutAddressUpdated is a free log subscription operation binding the contract event 0x0c69f21751e800ea5960436c9a94370c7adbf54c733a20a025293fbbe8f16252.
//
// Solidity: event CreatorPayoutAddressUpdated(address indexed nftContract, address indexed newPayoutAddress)
func (_Seadrop *SeadropFilterer) WatchCreatorPayoutAddressUpdated(opts *bind.WatchOpts, sink chan<- *SeadropCreatorPayoutAddressUpdated, nftContract []common.Address, newPayoutAddress []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var newPayoutAddressRule []interface{}
	for _, newPayoutAddressItem := range newPayoutAddress {
		newPayoutAddressRule = append(newPayoutAddressRule, newPayoutAddressItem)
	}

	logs, sub, err := _Seadrop.contract.WatchLogs(opts, "CreatorPayoutAddressUpdated", nftContractRule, newPayoutAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeadropCreatorPayoutAddressUpdated)
				if err := _Seadrop.contract.UnpackLog(event, "CreatorPayoutAddressUpdated", log); err != nil {
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
func (_Seadrop *SeadropFilterer) ParseCreatorPayoutAddressUpdated(log types.Log) (*SeadropCreatorPayoutAddressUpdated, error) {
	event := new(SeadropCreatorPayoutAddressUpdated)
	if err := _Seadrop.contract.UnpackLog(event, "CreatorPayoutAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeadropDropURIUpdatedIterator is returned from FilterDropURIUpdated and is used to iterate over the raw logs and unpacked data for DropURIUpdated events raised by the Seadrop contract.
type SeadropDropURIUpdatedIterator struct {
	Event *SeadropDropURIUpdated // Event containing the contract specifics and raw log

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
func (it *SeadropDropURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeadropDropURIUpdated)
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
		it.Event = new(SeadropDropURIUpdated)
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
func (it *SeadropDropURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeadropDropURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeadropDropURIUpdated represents a DropURIUpdated event raised by the Seadrop contract.
type SeadropDropURIUpdated struct {
	NftContract common.Address
	NewDropURI  string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDropURIUpdated is a free log retrieval operation binding the contract event 0xa0295608d25b3033c2e2c41cbac8746c2d08767bcfde6d47fae1ed7ba1d32150.
//
// Solidity: event DropURIUpdated(address indexed nftContract, string newDropURI)
func (_Seadrop *SeadropFilterer) FilterDropURIUpdated(opts *bind.FilterOpts, nftContract []common.Address) (*SeadropDropURIUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Seadrop.contract.FilterLogs(opts, "DropURIUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return &SeadropDropURIUpdatedIterator{contract: _Seadrop.contract, event: "DropURIUpdated", logs: logs, sub: sub}, nil
}

// WatchDropURIUpdated is a free log subscription operation binding the contract event 0xa0295608d25b3033c2e2c41cbac8746c2d08767bcfde6d47fae1ed7ba1d32150.
//
// Solidity: event DropURIUpdated(address indexed nftContract, string newDropURI)
func (_Seadrop *SeadropFilterer) WatchDropURIUpdated(opts *bind.WatchOpts, sink chan<- *SeadropDropURIUpdated, nftContract []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Seadrop.contract.WatchLogs(opts, "DropURIUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeadropDropURIUpdated)
				if err := _Seadrop.contract.UnpackLog(event, "DropURIUpdated", log); err != nil {
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
func (_Seadrop *SeadropFilterer) ParseDropURIUpdated(log types.Log) (*SeadropDropURIUpdated, error) {
	event := new(SeadropDropURIUpdated)
	if err := _Seadrop.contract.UnpackLog(event, "DropURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeadropPayerUpdatedIterator is returned from FilterPayerUpdated and is used to iterate over the raw logs and unpacked data for PayerUpdated events raised by the Seadrop contract.
type SeadropPayerUpdatedIterator struct {
	Event *SeadropPayerUpdated // Event containing the contract specifics and raw log

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
func (it *SeadropPayerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeadropPayerUpdated)
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
		it.Event = new(SeadropPayerUpdated)
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
func (it *SeadropPayerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeadropPayerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeadropPayerUpdated represents a PayerUpdated event raised by the Seadrop contract.
type SeadropPayerUpdated struct {
	NftContract common.Address
	Payer       common.Address
	Allowed     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayerUpdated is a free log retrieval operation binding the contract event 0x55a5cfa4bc68ffb9d833b75bf93f6d9c9aadc558dbfa587a9b5bb0ea7d5c38a3.
//
// Solidity: event PayerUpdated(address indexed nftContract, address indexed payer, bool indexed allowed)
func (_Seadrop *SeadropFilterer) FilterPayerUpdated(opts *bind.FilterOpts, nftContract []common.Address, payer []common.Address, allowed []bool) (*SeadropPayerUpdatedIterator, error) {

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

	logs, sub, err := _Seadrop.contract.FilterLogs(opts, "PayerUpdated", nftContractRule, payerRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &SeadropPayerUpdatedIterator{contract: _Seadrop.contract, event: "PayerUpdated", logs: logs, sub: sub}, nil
}

// WatchPayerUpdated is a free log subscription operation binding the contract event 0x55a5cfa4bc68ffb9d833b75bf93f6d9c9aadc558dbfa587a9b5bb0ea7d5c38a3.
//
// Solidity: event PayerUpdated(address indexed nftContract, address indexed payer, bool indexed allowed)
func (_Seadrop *SeadropFilterer) WatchPayerUpdated(opts *bind.WatchOpts, sink chan<- *SeadropPayerUpdated, nftContract []common.Address, payer []common.Address, allowed []bool) (event.Subscription, error) {

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

	logs, sub, err := _Seadrop.contract.WatchLogs(opts, "PayerUpdated", nftContractRule, payerRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeadropPayerUpdated)
				if err := _Seadrop.contract.UnpackLog(event, "PayerUpdated", log); err != nil {
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
func (_Seadrop *SeadropFilterer) ParsePayerUpdated(log types.Log) (*SeadropPayerUpdated, error) {
	event := new(SeadropPayerUpdated)
	if err := _Seadrop.contract.UnpackLog(event, "PayerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeadropPublicDropUpdatedIterator is returned from FilterPublicDropUpdated and is used to iterate over the raw logs and unpacked data for PublicDropUpdated events raised by the Seadrop contract.
type SeadropPublicDropUpdatedIterator struct {
	Event *SeadropPublicDropUpdated // Event containing the contract specifics and raw log

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
func (it *SeadropPublicDropUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeadropPublicDropUpdated)
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
		it.Event = new(SeadropPublicDropUpdated)
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
func (it *SeadropPublicDropUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeadropPublicDropUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeadropPublicDropUpdated represents a PublicDropUpdated event raised by the Seadrop contract.
type SeadropPublicDropUpdated struct {
	NftContract common.Address
	PublicDrop  PublicDrop
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPublicDropUpdated is a free log retrieval operation binding the contract event 0x3e30d8e1f739ea4795c481b21c23f905e938b80339305f3508e43c558e5dead3.
//
// Solidity: event PublicDropUpdated(address indexed nftContract, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop)
func (_Seadrop *SeadropFilterer) FilterPublicDropUpdated(opts *bind.FilterOpts, nftContract []common.Address) (*SeadropPublicDropUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Seadrop.contract.FilterLogs(opts, "PublicDropUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return &SeadropPublicDropUpdatedIterator{contract: _Seadrop.contract, event: "PublicDropUpdated", logs: logs, sub: sub}, nil
}

// WatchPublicDropUpdated is a free log subscription operation binding the contract event 0x3e30d8e1f739ea4795c481b21c23f905e938b80339305f3508e43c558e5dead3.
//
// Solidity: event PublicDropUpdated(address indexed nftContract, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop)
func (_Seadrop *SeadropFilterer) WatchPublicDropUpdated(opts *bind.WatchOpts, sink chan<- *SeadropPublicDropUpdated, nftContract []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Seadrop.contract.WatchLogs(opts, "PublicDropUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeadropPublicDropUpdated)
				if err := _Seadrop.contract.UnpackLog(event, "PublicDropUpdated", log); err != nil {
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
func (_Seadrop *SeadropFilterer) ParsePublicDropUpdated(log types.Log) (*SeadropPublicDropUpdated, error) {
	event := new(SeadropPublicDropUpdated)
	if err := _Seadrop.contract.UnpackLog(event, "PublicDropUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeadropSeaDropMintIterator is returned from FilterSeaDropMint and is used to iterate over the raw logs and unpacked data for SeaDropMint events raised by the Seadrop contract.
type SeadropSeaDropMintIterator struct {
	Event *SeadropSeaDropMint // Event containing the contract specifics and raw log

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
func (it *SeadropSeaDropMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeadropSeaDropMint)
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
		it.Event = new(SeadropSeaDropMint)
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
func (it *SeadropSeaDropMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeadropSeaDropMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeadropSeaDropMint represents a SeaDropMint event raised by the Seadrop contract.
type SeadropSeaDropMint struct {
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
func (_Seadrop *SeadropFilterer) FilterSeaDropMint(opts *bind.FilterOpts, nftContract []common.Address, minter []common.Address, feeRecipient []common.Address) (*SeadropSeaDropMintIterator, error) {

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

	logs, sub, err := _Seadrop.contract.FilterLogs(opts, "SeaDropMint", nftContractRule, minterRule, feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &SeadropSeaDropMintIterator{contract: _Seadrop.contract, event: "SeaDropMint", logs: logs, sub: sub}, nil
}

// WatchSeaDropMint is a free log subscription operation binding the contract event 0xe90cf9cc0a552cf52ea6ff74ece0f1c8ae8cc9ad630d3181f55ac43ca076b7d6.
//
// Solidity: event SeaDropMint(address indexed nftContract, address indexed minter, address indexed feeRecipient, address payer, uint256 quantityMinted, uint256 unitMintPrice, uint256 feeBps, uint256 dropStageIndex)
func (_Seadrop *SeadropFilterer) WatchSeaDropMint(opts *bind.WatchOpts, sink chan<- *SeadropSeaDropMint, nftContract []common.Address, minter []common.Address, feeRecipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Seadrop.contract.WatchLogs(opts, "SeaDropMint", nftContractRule, minterRule, feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeadropSeaDropMint)
				if err := _Seadrop.contract.UnpackLog(event, "SeaDropMint", log); err != nil {
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
func (_Seadrop *SeadropFilterer) ParseSeaDropMint(log types.Log) (*SeadropSeaDropMint, error) {
	event := new(SeadropSeaDropMint)
	if err := _Seadrop.contract.UnpackLog(event, "SeaDropMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeadropSignedMintValidationParamsUpdatedIterator is returned from FilterSignedMintValidationParamsUpdated and is used to iterate over the raw logs and unpacked data for SignedMintValidationParamsUpdated events raised by the Seadrop contract.
type SeadropSignedMintValidationParamsUpdatedIterator struct {
	Event *SeadropSignedMintValidationParamsUpdated // Event containing the contract specifics and raw log

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
func (it *SeadropSignedMintValidationParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeadropSignedMintValidationParamsUpdated)
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
		it.Event = new(SeadropSignedMintValidationParamsUpdated)
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
func (it *SeadropSignedMintValidationParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeadropSignedMintValidationParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeadropSignedMintValidationParamsUpdated represents a SignedMintValidationParamsUpdated event raised by the Seadrop contract.
type SeadropSignedMintValidationParamsUpdated struct {
	NftContract                common.Address
	Signer                     common.Address
	SignedMintValidationParams SignedMintValidationParams
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterSignedMintValidationParamsUpdated is a free log retrieval operation binding the contract event 0xcaeb4009c05208df426d15ff50b608287b05d21dee1f790552ea451a540a7be0.
//
// Solidity: event SignedMintValidationParamsUpdated(address indexed nftContract, address indexed signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams)
func (_Seadrop *SeadropFilterer) FilterSignedMintValidationParamsUpdated(opts *bind.FilterOpts, nftContract []common.Address, signer []common.Address) (*SeadropSignedMintValidationParamsUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Seadrop.contract.FilterLogs(opts, "SignedMintValidationParamsUpdated", nftContractRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &SeadropSignedMintValidationParamsUpdatedIterator{contract: _Seadrop.contract, event: "SignedMintValidationParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchSignedMintValidationParamsUpdated is a free log subscription operation binding the contract event 0xcaeb4009c05208df426d15ff50b608287b05d21dee1f790552ea451a540a7be0.
//
// Solidity: event SignedMintValidationParamsUpdated(address indexed nftContract, address indexed signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams)
func (_Seadrop *SeadropFilterer) WatchSignedMintValidationParamsUpdated(opts *bind.WatchOpts, sink chan<- *SeadropSignedMintValidationParamsUpdated, nftContract []common.Address, signer []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Seadrop.contract.WatchLogs(opts, "SignedMintValidationParamsUpdated", nftContractRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeadropSignedMintValidationParamsUpdated)
				if err := _Seadrop.contract.UnpackLog(event, "SignedMintValidationParamsUpdated", log); err != nil {
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
func (_Seadrop *SeadropFilterer) ParseSignedMintValidationParamsUpdated(log types.Log) (*SeadropSignedMintValidationParamsUpdated, error) {
	event := new(SeadropSignedMintValidationParamsUpdated)
	if err := _Seadrop.contract.UnpackLog(event, "SignedMintValidationParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeadropTokenGatedDropStageUpdatedIterator is returned from FilterTokenGatedDropStageUpdated and is used to iterate over the raw logs and unpacked data for TokenGatedDropStageUpdated events raised by the Seadrop contract.
type SeadropTokenGatedDropStageUpdatedIterator struct {
	Event *SeadropTokenGatedDropStageUpdated // Event containing the contract specifics and raw log

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
func (it *SeadropTokenGatedDropStageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeadropTokenGatedDropStageUpdated)
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
		it.Event = new(SeadropTokenGatedDropStageUpdated)
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
func (it *SeadropTokenGatedDropStageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeadropTokenGatedDropStageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeadropTokenGatedDropStageUpdated represents a TokenGatedDropStageUpdated event raised by the Seadrop contract.
type SeadropTokenGatedDropStageUpdated struct {
	NftContract     common.Address
	AllowedNftToken common.Address
	DropStage       TokenGatedDropStage
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenGatedDropStageUpdated is a free log retrieval operation binding the contract event 0xc695f93ae16034280e4fc93181b6afca9af23027ac1f1842a2287ba25cdc4476.
//
// Solidity: event TokenGatedDropStageUpdated(address indexed nftContract, address indexed allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage)
func (_Seadrop *SeadropFilterer) FilterTokenGatedDropStageUpdated(opts *bind.FilterOpts, nftContract []common.Address, allowedNftToken []common.Address) (*SeadropTokenGatedDropStageUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var allowedNftTokenRule []interface{}
	for _, allowedNftTokenItem := range allowedNftToken {
		allowedNftTokenRule = append(allowedNftTokenRule, allowedNftTokenItem)
	}

	logs, sub, err := _Seadrop.contract.FilterLogs(opts, "TokenGatedDropStageUpdated", nftContractRule, allowedNftTokenRule)
	if err != nil {
		return nil, err
	}
	return &SeadropTokenGatedDropStageUpdatedIterator{contract: _Seadrop.contract, event: "TokenGatedDropStageUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenGatedDropStageUpdated is a free log subscription operation binding the contract event 0xc695f93ae16034280e4fc93181b6afca9af23027ac1f1842a2287ba25cdc4476.
//
// Solidity: event TokenGatedDropStageUpdated(address indexed nftContract, address indexed allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage)
func (_Seadrop *SeadropFilterer) WatchTokenGatedDropStageUpdated(opts *bind.WatchOpts, sink chan<- *SeadropTokenGatedDropStageUpdated, nftContract []common.Address, allowedNftToken []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var allowedNftTokenRule []interface{}
	for _, allowedNftTokenItem := range allowedNftToken {
		allowedNftTokenRule = append(allowedNftTokenRule, allowedNftTokenItem)
	}

	logs, sub, err := _Seadrop.contract.WatchLogs(opts, "TokenGatedDropStageUpdated", nftContractRule, allowedNftTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeadropTokenGatedDropStageUpdated)
				if err := _Seadrop.contract.UnpackLog(event, "TokenGatedDropStageUpdated", log); err != nil {
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
func (_Seadrop *SeadropFilterer) ParseTokenGatedDropStageUpdated(log types.Log) (*SeadropTokenGatedDropStageUpdated, error) {
	event := new(SeadropTokenGatedDropStageUpdated)
	if err := _Seadrop.contract.UnpackLog(event, "TokenGatedDropStageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
