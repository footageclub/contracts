// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {
    ERC721ContractMetadataUpgradeable,
    ISeaDropTokenContractMetadataUpgradeable
} from "../ERC721ContractMetadataUpgradeable.sol";


import {
    ReentrancyGuardUpgradeable
} from "../../lib-upgradeable/solmate/src/utils/ReentrancyGuardUpgradeable.sol";

import {
    IERC165Upgradeable
} from "../../lib/openzeppelin-contracts-upgradeable/contracts/utils/introspection/IERC165Upgradeable.sol";

import {
    DefaultOperatorFiltererUpgradeable
} from "../../lib-upgradeable/operator-filter-registry/src/upgradeable/DefaultOperatorFiltererUpgradeable.sol";

import { ERC721ClubStorage } from "./ERC721ClubStorage.sol";

import {
    ERC721ContractMetadataStorage
} from "../ERC721ContractMetadataStorage.sol";

import { IERC721AUpgradeable } from "../../lib/ERC721A-Upgradeable/contracts/IERC721AUpgradeable.sol";
import { ECDSAUpgradeable }  from "../../lib/openzeppelin-contracts-upgradeable/contracts/utils/cryptography/ECDSAUpgradeable.sol";
import {
    SeaDropErrorsAndEventsUpgradeable
} from "../lib/SeaDropErrorsAndEventsUpgradeable.sol";

contract ERC721ClubUpgradeable is ERC721ContractMetadataUpgradeable,
    ReentrancyGuardUpgradeable,
    DefaultOperatorFiltererUpgradeable,
    SeaDropErrorsAndEventsUpgradeable
{
    using ERC721ClubStorage for ERC721ClubStorage.Layout;
    using ERC721ContractMetadataStorage for ERC721ContractMetadataStorage.Layout;
    using ECDSAUpgradeable for bytes32;


    /**
     * @notice A token can only be burned by the set burn address.
     */
    error BurnIncorrectSender();

    /**
     * @notice 仅允许已经指定的合约质押
     */
    error OnlyAllowedPlegeContract();

    /**
     * @notice 不允许合约重复
     */
    error DuplicateAllowedPlegedContract();

    /**
     * @notice 仅club拥有者才能操作
     */
    error OnlyClubOwner(uint256 tokenId);

    /**
     * @notice 合约没有授权
     */
    error OnlyIsApprovedForAll();

    /**
     * @notice 必须质押
     */
    error OnlyPleged();

    /**
     * @notice 合约部署成功
     */
    event ClubTokenDeployed();

    /**
     * @notice 更新允许质押的nft合约成功
     */
    event AllowedPlegedContractUpdated(address indexed sender, address indexed nftToken, bool indexed allowed);

    /**
     * @notice 更新club每次质押增加的空间
     */
    event SpaceSizeUpdated(uint256 size, uint256 length);

    /**
     * @notice 更新签名地址事件
     */
    event SignerUpdated(address indexed sender, address indexed signer);

    struct ClubParams {
        uint256 tokenId;

        uint256[] token;

        uint256 currentSpace;

        uint256 Space;

        uint256 currentVideoLength;

        uint256 Videolength;
    }

    /**
     * @notice 质押成功
     */
    event PlegedUpdated(
        uint256 indexed clubId,
        address indexed nftContract, 
        address indexed sender, 
        ClubParams params,
        bool minted);
        

    /**
     * @notice 解除质押成功
     */
    event UnplegedUpdated(
        uint256 indexed clubId,
        address indexed nftContract, 
        address indexed sender, 
        ClubParams params);
    /**
     * @notice Initialize the token contract with its name, symbol,
     *         administrator, and allowed SeaDrop addresses.
     */
    function initialize(
        string memory name,
        string memory symbol,
        address plegeSigner,
        uint256 defaultSpace,
        uint256 defaultVideoLength
    ) external initializer initializerERC721A {
        __ERC721A_init_unchained(name, symbol);
        __ConstructorInitializable_init_unchained();
        __TwoStepOwnable_init_unchained();
        __ERC721ContractMetadata_init_unchained(name, symbol);
        __ReentrancyGuard_init_unchained();
        __DefaultOperatorFilterer_init();
        __ERC721Club_init_unchained(plegeSigner, defaultSpace, defaultVideoLength);
    }

    function __ERC721Club_init_unchained(
        address signer,
        uint256 defaultSpace,
        uint256 defaultVideoLength
    ) internal onlyInitializing {
       

        // Set the enumeration.
        ERC721ClubStorage.layout()._signer = signer;
        ERC721ClubStorage.layout()._fileSpace = defaultSpace;
        ERC721ClubStorage.layout()._videoSpace = defaultVideoLength;
        ERC721ClubStorage.layout()._CHAIN_ID = block.chainid;
        ERC721ClubStorage.layout()._DOMAIN_SEPARATOR = _deriveDomainSeparator();

        // Emit an event noting the contract deployment.
        emit ClubTokenDeployed();
    }


    function _onlyAllowedPlegeContract(address nftcontract) internal view {
        if (ERC721ClubStorage.layout()._allowedPledgedContracts[nftcontract] != true) {
            revert OnlyAllowedPlegeContract();
        }
    }

    function _onlyClubOwnerOrZero(uint256 clubId, uint256 tokenId) internal view {
        uint256 clubTokenId =  ERC721ClubStorage.layout()._club[clubId];
        if (tokenId != clubTokenId) {
            revert OnlyOwner();
        }

        if( tokenId != 0 ) {
            if (ownerOf(tokenId) != msg.sender) {
                 revert OnlyOwner();
            }
        }
    }

    function _onlyClubOwner(uint256 clubId, uint256 tokenId) internal view {
        uint256 clubTokenId =  ERC721ClubStorage.layout()._club[clubId];
        if ( _cast(ownerOf(tokenId) != msg.sender) | _cast(clubTokenId != clubId)  == 0) {
            revert OnlyOwner();
        }
    }

    /**
     * @dev 质押nft到club中来
     * 
     * @param clubId club id
     * @param tokenId club token ID
     * @param nftContract 质押的nft合约地址
     * @param plegeToken 质押的nft token
     */
    function plege(uint256 clubId, uint256 tokenId, address nftContract, uint256[] calldata plegeToken, uint256 salt, bytes calldata signature) external virtual nonReentrant {
        _onlyAllowedPlegeContract(nftContract);

        _onlyClubOwnerOrZero(clubId, tokenId);
        if(!IERC721AUpgradeable(nftContract).isApprovedForAll(msg.sender, address(this))) {
            revert OnlyIsApprovedForAll();
        }

       // 签名验证
       _verifySignature(clubId, tokenId, nftContract, plegeToken, salt, signature, ERC721ClubStorage._SIGNED_PLEGE_TYPEHASH);
        mapping(address => mapping (uint256 => bool)) storage pledged = ERC721ClubStorage.layout()._pledged[tokenId];
        for (uint256 i = 0; i < plegeToken.length; ) {
            if (pledged[nftContract][plegeToken[i]]) {
                revert("DuplicatePleged");
            }
            
            IERC721AUpgradeable(nftContract).transferFrom(msg.sender, address(this), plegeToken[i]);
            ERC721ClubStorage.layout()._pledged[tokenId][nftContract][plegeToken[i]] = true;
            unchecked {
                ++i;
            }
        }

        bool minted = false;
        if ( tokenId == 0 ) {
            tokenId = _nextTokenId();
            _safeMint(msg.sender, 1);
            minted = true;
            ERC721ClubStorage.layout()._club[clubId] = tokenId;
            ERC721ClubStorage.layout()._clubByToken[tokenId] = clubId;
        }

        uint256 currentSpace = ERC721ClubStorage.layout()._storageFileSpace[tokenId] + ERC721ClubStorage.layout()._fileSpace * plegeToken.length;
        uint256 currentLength = ERC721ClubStorage.layout()._storageVideoSpace[tokenId] + ERC721ClubStorage.layout()._videoSpace * plegeToken.length;
        ERC721ClubStorage.layout()._storageFileSpace[tokenId] = currentSpace;
        ERC721ClubStorage.layout()._storageVideoSpace[tokenId] = currentLength;
        ClubParams memory params =  ClubParams(
            tokenId, 
            plegeToken, 
            currentSpace,
            ERC721ClubStorage.layout()._fileSpace,
            currentLength,
            ERC721ClubStorage.layout()._videoSpace);
        emit PlegedUpdated(clubId, nftContract, msg.sender, params, minted);
    }

    /**
     * @dev 解除质押
     * 
     * @param tokenId club ID
     * @param nftContract 质押的nft合约地址
     * @param plegedToken 已质押的nft token
     */
    function unplege(uint256 clubId, uint256 tokenId, address nftContract, uint256[] calldata plegedToken, uint256 salt, bytes calldata signature) external virtual nonReentrant {
        _onlyAllowedPlegeContract(nftContract);

        _onlyClubOwner(clubId, tokenId);

       // 签名验证
       _verifySignature(clubId, tokenId, nftContract, plegedToken, salt, signature, ERC721ClubStorage._SIGNED_UNPLEGE_TYPEHASH);
        mapping(address => mapping (uint256 => bool)) storage pledged = ERC721ClubStorage.layout()._pledged[tokenId];
        for (uint256 i = 0; i < plegedToken.length;) {
            if (!pledged[nftContract][plegedToken[i]]) {
                revert OnlyPleged();
            }
            
            ERC721ClubStorage.layout()._pledged[tokenId][nftContract][plegedToken[i]] = false;
            IERC721AUpgradeable(nftContract).transferFrom(address(this),msg.sender, plegedToken[i]);
            delete(ERC721ClubStorage.layout()._pledged[tokenId][nftContract][plegedToken[i]]);
            unchecked {
                ++i;
            }
        }


        uint256 space = ERC721ClubStorage.layout()._fileSpace * plegedToken.length;
        uint256 length = ERC721ClubStorage.layout()._videoSpace * plegedToken.length;
        uint256 currentSpace = 0;
        uint256 currentLength = 0;
        if (ERC721ClubStorage.layout()._storageFileSpace[tokenId] > space ) {
            currentSpace = ERC721ClubStorage.layout()._storageFileSpace[tokenId] - space;
        }

        if (ERC721ClubStorage.layout()._storageVideoSpace[tokenId] > length) {
            currentLength = ERC721ClubStorage.layout()._storageVideoSpace[tokenId] - length;
        }

        ERC721ClubStorage.layout()._storageFileSpace[tokenId] = currentSpace;
        ERC721ClubStorage.layout()._storageVideoSpace[tokenId] = currentLength;
        ClubParams memory params =  ClubParams(
            tokenId, 
            plegedToken, 
            currentSpace,
            ERC721ClubStorage.layout()._fileSpace,
            currentLength,
            ERC721ClubStorage.layout()._videoSpace);
        emit UnplegedUpdated(clubId, nftContract, msg.sender, params);
    }

    function _verifySignature(uint256 clubId, uint256 tokenId, address nftContract, uint256[] calldata plegedToken, uint256 salt, bytes calldata signature, bytes32 signHash) internal virtual {
        bytes32 digest = _getDigest(
            nftContract, 
            msg.sender, 
            clubId,
            tokenId, 
            plegedToken, 
            salt, signHash);

        if (ERC721ClubStorage.layout()._usedDigests[digest]) {
                revert SignatureAlreadyUsed();
        }

        ERC721ClubStorage.layout()._usedDigests[digest] = true;
        address recoveredAddress = digest.recover(signature);
        if( recoveredAddress != ERC721ClubStorage.layout()._signer){
            revert InvalidSignature(recoveredAddress);
        }
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

    function getClubId( uint256 tokenId ) external view returns( uint256 ) {
        return ERC721ClubStorage.layout()._clubByToken[tokenId];
    }

    function getClubTokenId( uint256 clubId ) external view returns( uint256 ) {
        return  ERC721ClubStorage.layout()._club[clubId];
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

    /**
     * 更新允许质押的NFT
     */
    function updateAllowedPledgeToken(address nftContract, bool allowed) external virtual onlyOwner {
        if (allowed) {
            if (ERC721ClubStorage.layout()._allowedPledgedContracts[nftContract]) {
                revert("DuplicateAllowedPlegedContract");
            }

           ERC721ClubStorage.layout()._allowedPledgedContracts[nftContract] = true;
        } else {
            if (!ERC721ClubStorage.layout()._allowedPledgedContracts[nftContract]) {
                revert("AllowedPlegedContractNotPresent");
            }

            delete ERC721ClubStorage.layout()._allowedPledgedContracts[nftContract];
        }

        // Emit an event with the update.
        emit AllowedPlegedContractUpdated(msg.sender, nftContract, allowed);
    }

    function updateSigner(address signer) external onlyOwner {
        ERC721ClubStorage.layout()._signer = signer;

        emit SignerUpdated(msg.sender, signer);
    }

    function updateSpaceSize(uint256 size, uint256 length) external virtual onlyOwner {
        if (size == 0 || length == 0) {
            revert("SpaceOrLengthZero");
        }
        
        ERC721ClubStorage.layout()._fileSpace = size;
        ERC721ClubStorage.layout()._videoSpace = length;
        emit SpaceSizeUpdated(size, length);
    }

    function getClubSpace(uint256 tokenId) external virtual returns (uint256, uint256) {
        return (
            ERC721ClubStorage.layout()._storageFileSpace[tokenId],
            ERC721ClubStorage.layout()._storageVideoSpace[tokenId]);
    }

    function setBurnAddress(address newBurnAddress) external onlyOwner {
        ERC721ClubStorage.layout().burnAddress = newBurnAddress;
    }

    function getBurnAddress() public view returns (address) {
        return ERC721ClubStorage.layout().burnAddress;
    }

    /**
     * @notice Destroys `tokenId`, only callable by the set burn address.
     *
     * @param tokenId The token id to burn.
     */
    function burn(uint256 tokenId) external {
        if (msg.sender != ERC721ClubStorage.layout().burnAddress) {
            revert BurnIncorrectSender();
        }
        _burn(tokenId);
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
        return block.chainid == ERC721ClubStorage.layout()._CHAIN_ID
            ? ERC721ClubStorage.layout()._DOMAIN_SEPARATOR
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
                ERC721ClubStorage._EIP_712_DOMAIN_TYPEHASH,
                ERC721ClubStorage._NAME_HASH,
                ERC721ClubStorage._VERSION_HASH,
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
     * @param nftContract  The nft contract.
     * @param sender       The sender address.
     * @param clubId       The club id.
     * @param tokenId      The club token id.
     * @param plegedToken  The nft tokens.
     * @param salt         The salt for the signed mint.
     * @param typeHash     hash
     */
    function _getDigest(
        address nftContract,
        address sender,
        uint256 clubId,
        uint256 tokenId,
        uint256[] memory plegedToken,
        uint256 salt,
        bytes32 typeHash
    ) internal view returns (bytes32 digest) {
        digest = keccak256(
            bytes.concat(
                bytes2(0x1901),
                _domainSeparator(),
                keccak256(
                    abi.encode(
                        typeHash,
                        nftContract,
                        sender,
                        clubId,
                        tokenId,
                        keccak256(abi.encodePacked(plegedToken)),
                        salt
                    )
                )
            )
        );
    }
}
