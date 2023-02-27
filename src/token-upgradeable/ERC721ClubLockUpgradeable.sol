// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;


import {
    ISeaDropTokenContractMetadataUpgradeable
} from "seadrop-upgradeable/interfaces/ISeaDropTokenContractMetadataUpgradeable.sol";

import { 
    ERC721ContractMetadataUpgradeable
} from "seadrop-upgradeable/ERC721ContractMetadataUpgradeable.sol";

import {
    ReentrancyGuardUpgradeable
} from "../../lib-upgradeable/solmate/src/utils/ReentrancyGuardUpgradeable.sol";

import {
    IERC165Upgradeable
} from "../../lib-upgradeable/openzeppelin-contracts-upgradeable/contracts/utils/introspection/IERC165Upgradeable.sol";

import {
    DefaultOperatorFiltererUpgradeable
} from "../../lib-upgradeable/operator-filter-registry/src/upgradeable/DefaultOperatorFiltererUpgradeable.sol";

import {
    ERC721ContractMetadataStorage
} from "seadrop-upgradeable/ERC721ContractMetadataStorage.sol";

import { ERC721SeaDropUpgradeable } from "seadrop-upgradeable/ERC721SeaDropUpgradeable.sol";
import { ERC721ClubLockStorage } from "./ERC721ClubLockStorage.sol";
import { MintParams } from "seadrop-upgradeable/lib/SeaDropStructsUpgradeable.sol";
import { 
    ECDSAUpgradeable 
}  from "../../lib-upgradeable/openzeppelin-contracts-upgradeable/contracts/utils/cryptography/ECDSAUpgradeable.sol";

import {
    SeaDropErrorsAndEventsUpgradeable
} from "seadrop-upgradeable/lib/SeaDropErrorsAndEventsUpgradeable.sol";

contract ERC721ClubLockUpgradeable is ERC721ContractMetadataUpgradeable,
    ReentrancyGuardUpgradeable,
    DefaultOperatorFiltererUpgradeable,
    SeaDropErrorsAndEventsUpgradeable {
    using ERC721ClubLockStorage for ERC721ClubLockStorage.Layout;
    using ERC721ContractMetadataStorage for ERC721ContractMetadataStorage.Layout;
    using ECDSAUpgradeable for bytes32;

    /**
     * @notice A token can only be burned by the set burn address.
     */
    error BurnIncorrectSender();

    /**
     * @notice 错误的用户资产ID
     */
    error InvalidAssetId();

    /**
     * @notice 合约部署成功
     */
    event ClubLockDeployed();

    /**
     * @notice 铸造成功事件
     */
     event mintClublockUpdated(
        address indexed minter, 
        address indexed to,
        uint256 mintPrice,
        string[] assets
    );

    error trace( address minter,
        address to,
        string[] assets,
        uint256 salt);

    /**
     * @notice 更新签名地址事件
     */
    event SignerUpdated(address indexed sender, address indexed signer);

    /**
     * @notice 铸造信息已经更新
     */
    event MintParamsUpdated(address indexed sender, MintParams mintParams);

    /**
     * @notice Initialize the token contract with its name, symbol,
     *         and allowed SeaDrop addresses.
     */
    function initialize(
        string memory name,
        string memory symbol,
        address signer
    ) external initializer initializerERC721A {
        __ERC721A_init_unchained(name, symbol);
        __ConstructorInitializable_init_unchained();
        __TwoStepOwnable_init_unchained();
        __ERC721ContractMetadata_init_unchained(name, symbol);
        __ReentrancyGuard_init_unchained();
        __DefaultOperatorFilterer_init();
        __ERC721ClubLock_init_unchained(signer);
    }

    function __ERC721ClubLock_init_unchained(address signer) internal onlyInitializing {
        ERC721ClubLockStorage.layout()._CHAIN_ID = block.chainid;
        ERC721ClubLockStorage.layout()._DOMAIN_SEPARATOR = _deriveDomainSeparator();
        ERC721ClubLockStorage.layout()._signer = signer;
        // Emit an event noting the contract deployment.
        emit ClubLockDeployed();
    }

    function mintClublock(address to, string[] calldata assets, uint256 salt, bytes calldata signature) external payable nonReentrant {
         // Get the minter address.
        address minter = to != address(0)
            ? to
            : msg.sender;

        uint256 quantity = assets.length;

        MintParams memory mintParams =  ERC721ClubLockStorage.layout()._mintParams;

        // Validate that the dropStage is active.
        _checkActive(mintParams.startTime, mintParams.endTime);

          // Check that the minter is allowed to mint the desired quantity.
        _checkMintQuantity(minter, quantity, mintParams.maxTotalMintableByWallet, mintParams.maxTokenSupplyForStage);

           // Validate payment is correct for number minted.
        _checkCorrectPayment(quantity, mintParams.mintPrice);

        // Validate the signature in a block scope to avoid "stack too deep".
        {
            // Get the digest to verify the EIP-712 signature.
            bytes32 digest = _getDigest(
                minter,
                to,
                assets,
                salt
            );

            // Ensure the digest has not already been used.
            if (ERC721ClubLockStorage.layout()._usedDigests[digest]) {
                revert SignatureAlreadyUsed();
            }

            // Mark the digest as used.
            ERC721ClubLockStorage.layout()._usedDigests[digest] = true;

            // Use the recover method to see what address was used to create
            // the signature on this data.
            // Note that if the digest doesn't exactly match what was signed we'll
            // get a random recovered address.
            address recoveredAddress = digest.recover(signature);
            if (recoveredAddress != ERC721ClubLockStorage.layout()._signer ){
                revert InvalidSignature(recoveredAddress);
            }
        }

        for(uint256 i = 0; i < quantity;) {
            if(ERC721ClubLockStorage.layout()._clublockByAssetId[assets[i]] != 0) {
                revert InvalidAssetId(); 
            }

            unchecked {
                ++i;
            }
        }


        // Mint the quantity of tokens to the minter.
        _safeMint(minter, quantity);

        // eth to creator
        _splitPayout(address(0), 0);

        // emit event
        emit mintClublockUpdated(minter, to, mintParams.mintPrice, assets);
    }

    function updateSigner(address signer) external onlyOwner {
        ERC721ClubLockStorage.layout()._signer = signer;

        emit SignerUpdated(msg.sender, signer);
    }

    function updateMintParams(MintParams calldata mintParams) external {
        // Only owner or self call
        _onlyOwnerOrSelf();

        if(mintParams.feeBps > 10_000) {
            revert InvalidFeeBps(mintParams.feeBps);
        }

        ERC721ClubLockStorage.layout()._mintParams = mintParams;
        emit MintParamsUpdated(msg.sender, mintParams);
    }

    /**
     * @notice Update the creator payout address for this nft contract on SeaDrop.
     *         Only the owner can set the creator payout address.
     *
     * @param payoutAddress The new payout address.
     */
    function updateCreatorPayoutAddress( address payoutAddress ) external onlyOwner {

        // Update the creator payout address.
        if (payoutAddress == address(0)) {
            revert CreatorPayoutAddressCannotBeZeroAddress();
        }

        // Set the creator payout address.
        ERC721ClubLockStorage.layout()._creatorPayoutAddress  = payoutAddress;

        // Emit an event with the update.
        emit CreatorPayoutAddressUpdated(msg.sender, payoutAddress);
       
    }

    function config(address payoutAddress, uint256 maxSupply, MintParams calldata mintParams, RoyaltyInfo calldata royalty) external onlyOwner {
        if(payoutAddress != address(0)) {
            this.updateCreatorPayoutAddress(payoutAddress);
        }
        
        if ( maxSupply != 0 ) {
            this.setMaxSupply(maxSupply);
        }

        if (
            _cast(mintParams.startTime != 0) |
            _cast(mintParams.endTime != 0) == 1
        ) {
            this.updateMintParams(mintParams);
        }

        if(
            _cast(royalty.royaltyAddress != address(0)) |
            _cast(royalty.royaltyBps != 0) == 1
        ) {
            this.setRoyaltyInfo(royalty);
        }
    }

    function getMintParams() external view returns(MintParams memory) {
        return ERC721ClubLockStorage.layout()._mintParams;
    }

    /**
     * @notice Returns whether the interface is supported.
     *
     * @param interfaceId The interface id to check against.
     */
    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(ERC721ContractMetadataUpgradeable)
        returns (bool)
    {
        return interfaceId ==
            type(ISeaDropTokenContractMetadataUpgradeable).interfaceId ||
            // ERC721ContractMetadata returns supportsInterface true for
            //     EIP-2981
            // ERC721A returns supportsInterface true for
            //     ERC165, ERC721, ERC721Metadata
            super.supportsInterface(interfaceId);
    }

    /**
     * @dev Approve or remove `operator` as an operator for the caller.
     * Operators can call {transferFrom} or {safeTransferFrom}
     * for any token owned by the caller.
     *
     * Requirements:
     *
     * - The `operator` cannot be the caller.
     * - The `operator` must be allowed.
     *
     * Emits an {ApprovalForAll} event.
     */
    function setApprovalForAll(address operator, bool approved)
        public
        override
        onlyAllowedOperatorApproval(operator)
    {
        super.setApprovalForAll(operator, approved);
    }

    /**
     * @dev Gives permission to `to` to transfer `tokenId` token to another account.
     * The approval is cleared when the token is transferred.
     *
     * Only a single account can be approved at a time, so approving the
     * zero address clears previous approvals.
     *
     * Requirements:
     *
     * - The caller must own the token or be an approved operator.
     * - `tokenId` must exist.
     * - The `operator` mut be allowed.
     *
     * Emits an {Approval} event.
     */
    function approve(address operator, uint256 tokenId)
        public
        override
        onlyAllowedOperatorApproval(operator)
    {
        super.approve(operator, tokenId);
    }

    /**
     * @dev Transfers `tokenId` from `from` to `to`.
     *
     * Requirements:
     *
     * - `from` cannot be the zero address.
     * - `to` cannot be the zero address.
     * - `tokenId` token must be owned by `from`.
     * - If the caller is not `from`, it must be approved to move this token
     * by either {approve} or {setApprovalForAll}.
     * - The operator must be allowed.
     *
     * Emits a {Transfer} event.
     */
    function transferFrom(
        address from,
        address to,
        uint256 tokenId
    ) public override onlyAllowedOperator(from) {
        super.transferFrom(from, to, tokenId);
    }

    /**
     * @dev Equivalent to `safeTransferFrom(from, to, tokenId, '')`.
     */
    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId
    ) public override onlyAllowedOperator(from) {
        super.safeTransferFrom(from, to, tokenId);
    }

    /**
     * @dev Safely transfers `tokenId` token from `from` to `to`.
     *
     * Requirements:
     *
     * - `from` cannot be the zero address.
     * - `to` cannot be the zero address.
     * - `tokenId` token must exist and be owned by `from`.
     * - If the caller is not `from`, it must be approved to move this token
     * by either {approve} or {setApprovalForAll}.
     * - If `to` refers to a smart contract, it must implement
     * {IERC721Receiver-onERC721Received}, which is called upon a safe transfer.
     * - The operator must be allowed.
     *
     * Emits a {Transfer} event.
     */
    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId,
        bytes memory data
    ) public override onlyAllowedOperator(from) {
        super.safeTransferFrom(from, to, tokenId, data);
    }

    function setBurnAddress(address newBurnAddress) external onlyOwner {
        ERC721ClubLockStorage.layout().burnAddress = newBurnAddress;
    }

    function getBurnAddress() public view returns (address) {
        return ERC721ClubLockStorage.layout().burnAddress;
    }

    /**
     * @notice Destroys `tokenId`, only callable by the set burn address.
     *
     * @param tokenId The token id to burn.
     */
    function burn(uint256 tokenId) external {
        if (msg.sender != ERC721ClubLockStorage.layout().burnAddress) {
            revert BurnIncorrectSender();
        }
        _burn(tokenId);
    }

    /**
     * @notice Revert if the payment is not the quantity times the mint price.
     *
     * @param quantity  The number of tokens to mint.
     * @param mintPrice The mint price per token.
     */
    function _checkCorrectPayment(uint256 quantity, uint256 mintPrice)
        internal
        view
    {
        // Revert if the tx's value doesn't match the total cost.
        if (msg.value != quantity * mintPrice) {
            revert IncorrectPayment(msg.value, quantity * mintPrice);
        }
    }

     /**
     * @notice Check that the drop stage is active.
     *
     * @param startTime The drop stage start time.
     * @param endTime   The drop stage end time.
     */
    function _checkActive(uint256 startTime, uint256 endTime) internal view {
        if (
            _cast(block.timestamp < startTime) |
                _cast(block.timestamp > endTime) ==
            1
        ) {
            // Revert if the drop stage is not active.
            revert NotActive(block.timestamp, startTime, endTime);
        }
    }

    /**
     * @notice Check that the wallet is allowed to mint the desired quantity.
     *
     * @param minter                   The mint recipient.
     * @param quantity                 The number of tokens to mint.
     * @param maxTotalMintableByWallet The max allowed mints per wallet.
     * @param maxTokenSupplyForStage   The max token supply for the drop stage.
     */
    function _checkMintQuantity(
        address minter,
        uint256 quantity,
        uint256 maxTotalMintableByWallet,
        uint256 maxTokenSupplyForStage
    ) internal view {
        // Mint quantity of zero is not valid.
        if (quantity == 0) {
            revert MintQuantityCannotBeZero();
        }

        // Get the mint stats.
        uint256 minterNumMinted =  _numberMinted(minter);
        uint256 currentTotalSupply = _totalMinted();
        uint256 maxSupply = ERC721ContractMetadataStorage.layout()._maxSupply;
 

        // Ensure mint quantity doesn't exceed maxTotalMintableByWallet.
        if (quantity + minterNumMinted > maxTotalMintableByWallet) {
            revert MintQuantityExceedsMaxMintedPerWallet(
                quantity + minterNumMinted,
                maxTotalMintableByWallet
            );
        }

        // Ensure mint quantity doesn't exceed maxSupply.
        if (quantity + currentTotalSupply > maxSupply) {
            revert MintQuantityExceedsMaxSupply(
                quantity + currentTotalSupply,
                maxSupply
            );
        }

        // Ensure mint quantity doesn't exceed maxTokenSupplyForStage.
        if (quantity + currentTotalSupply > maxTokenSupplyForStage) {
            revert MintQuantityExceedsMaxTokenSupplyForStage(
                quantity + currentTotalSupply,
                maxTokenSupplyForStage
            );
        }
    }

    /**
     * @notice Returns a set of mint stats for the address.
     *         This assists SeaDrop in enforcing maxSupply,
     *         maxTotalMintableByWallet, and maxTokenSupplyForStage checks.
     *
     * @dev    NOTE: Implementing contracts should always update these numbers
     *         before transferring any tokens with _safeMint() to mitigate
     *         consequences of malicious onERC721Received() hooks.
     *
     * @param minter The minter address.
     */
    function getMintStats(address minter)
        external
        view
        returns (
            uint256 minterNumMinted,
            uint256 currentTotalSupply,
            uint256 maxSupply
        )
    {
        minterNumMinted = _numberMinted(minter);
        currentTotalSupply = _totalMinted();
        maxSupply = ERC721ContractMetadataStorage.layout()._maxSupply;
    }

    /**
     * @dev Internal view function to get the EIP-712 domain separator. If the
     *      chainId matches the chainId set on deployment, the cached domain
     *      separator will be returned; otherwise, it will be derived from
     *      scratch.
     *
     * @return The domain separator.
     */
    function _domainSeparator() internal view returns (bytes32) {
        // prettier-ignore
        return block.chainid == ERC721ClubLockStorage.layout()._CHAIN_ID
            ? ERC721ClubLockStorage.layout()._DOMAIN_SEPARATOR
            : _deriveDomainSeparator();
    }

    /**
     * @dev Internal view function to derive the EIP-712 domain separator.
     *
     * @return The derived domain separator.
     */
    function _deriveDomainSeparator() internal view returns (bytes32) {
        // prettier-ignore
        return keccak256(
            abi.encode(
                ERC721ClubLockStorage._EIP_712_DOMAIN_TYPEHASH,
                ERC721ClubLockStorage._NAME_HASH,
                ERC721ClubLockStorage._VERSION_HASH,
                block.chainid,
                address(this)
            )
        );
    }

    /**
     * @notice Verify an EIP-712 signature by recreating the data structure
     *         that we signed on the client side, and then using that to recover
     *         the address that signed the signature for this data.
     *
     * @param minter       The mint address.
     * @param to           The to address.
     * @param assets       The clublock bind id
     * @param salt         The salt for the signed mint.
     */
    function _getDigest(
        address minter,
        address to,
        string[] memory assets,
        uint256 salt
    ) internal view returns (bytes32 digest) {
        digest = keccak256(
            bytes.concat(
                bytes2(0x1901),
                _domainSeparator(),
                keccak256(
                    abi.encode(
                        ERC721ClubLockStorage._SIGNED_MINT_TYPEHASH,
                        minter,
                        to,
                        keccak256(encodeTightlyPacked(assets)),
                        salt
                    )
                )
            )
        );
    }

    function encodeTightlyPacked(string[] memory path) internal pure returns(bytes memory encoded) {
        for (uint i = 0; i < path.length; i++) {
            encoded = bytes.concat(
                encoded, abi.encodePacked(path[i])
            );
        }
    }

    /**
     * @notice Split the payment payout for the creator and fee recipient.
     *
     * @param feeRecipient The fee recipient.
     * @param feeBps       The fee basis points.
     */
    function _splitPayout(
        address feeRecipient,
        uint256 feeBps
    ) internal {
        // Revert if the fee basis points is greater than 10_000.
        if (feeBps > 10_000) {
            revert InvalidFeeBps(feeBps);
        }

        // Get the creator payout address.
        address creatorPayoutAddress =  ERC721ClubLockStorage.layout()._creatorPayoutAddress;

        // Ensure the creator payout address is not the zero address.
        if (creatorPayoutAddress == address(0)) {
            revert CreatorPayoutAddressCannotBeZeroAddress();
        }

        // msg.value has already been validated by this point, so can use it directly.

        // If the fee is zero, just transfer to the creator and return.
        if (feeBps == 0) {
            _safeTransferETH(creatorPayoutAddress, msg.value);
            return;
        }

        // Get the fee amount.
        // Note that the fee amount is rounded down in favor of the creator.
        uint256 feeAmount = (msg.value * feeBps) / 10_000;

        // Get the creator payout amount. Fee amount is <= msg.value per above.
        uint256 payoutAmount;
        unchecked {
            payoutAmount = msg.value - feeAmount;
        }

        // Transfer the fee amount to the fee recipient.
        if (feeAmount > 0) {
            _safeTransferETH(feeRecipient, feeAmount);
        }

        // Transfer the creator payout amount to the creator.
       _safeTransferETH(creatorPayoutAddress, payoutAmount);
    }

    function _safeTransferETH(address to, uint256 amount) internal {
        bool success;

        assembly {
            // Transfer the ETH and store if it succeeded or not.
            success := call(gas(), to, amount, 0, 0, 0, 0)
        }

        require(success, "ETH_TRANSFER_FAILED");
    }

    /**
     * @dev Overrides the `_startTokenId` function from ERC721A
     *      to start at token id `1`.
     *
     *      This is to avoid future possible problems since `0` is usually
     *      used to signal values that have not been set or have been removed.
     */
    function _startTokenId() internal view virtual override returns (uint256) {
        return 1;
    }

}