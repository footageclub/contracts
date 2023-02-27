// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

library ERC721ClubStorage {
    bytes32 internal constant _NAME_HASH = keccak256("Footage");
    bytes32 internal constant _VERSION_HASH = keccak256("1.0");
    bytes32 internal constant _EIP_712_DOMAIN_TYPEHASH =
        // prettier-ignore
        keccak256(
            "EIP712Domain("
                "string name,"
                "string version,"
                "uint256 chainId,"
                "address verifyingContract"
            ")"
        );

    bytes32 internal constant _SIGNED_PLEGE_TYPEHASH =
    keccak256(
        "Plege("
            "address nftContract,"
            "address sender,"
            "uint256 clubId,"
            "uint256 tokenId,"
            "uint256[] plegedToken,"
            "uint256 salt"
        ")"
    );

    bytes32 internal constant _SIGNED_UNPLEGE_TYPEHASH =
    keccak256(
        "Unplege("
            "address nftContract,"
            "address sender,"
            "uint256 clubId,"
            "uint256 tokenId,"
            "uint256[] plegedToken,"
            "uint256 salt"
        ")"
    );

    struct Layout {
        /// @notice The only address that can burn tokens on this contract.
        address burnAddress;

        /// @notice 俱乐部ID与TokenID关系映射 clubId => tokenId
        mapping(uint256 => uint256) _club;

        // @notice 俱乐部ID与TokenID关系映射  tokenId => clubId
        mapping(uint256 => uint256) _clubByToken;

        ///  @notice 允许质押的NFT合约
        mapping(address => bool) _allowedPledgedContracts;

        /// @notice 记录质押地块 clubId => contract address => token = > boo;
        mapping(uint256 => mapping(address => mapping (uint256 => bool))) _pledged;

        /// @notice 设置club质押NFT增长存储空间大小 (kb)
        mapping(uint256 => uint256) _storageFileSpace;

        /// @notice 设置club质押NFT增长的视频存储时长大小 (ms)
         mapping(uint256 => uint256) _storageVideoSpace;
        
        /// @notice 设置默认存储空间大小
        uint256 _fileSpace;

         /// @notice 设置默认视频存储时长大小 (ms)
        uint256 _videoSpace;
        
        /// @notice 质押签名者
        address _signer; 

        mapping(bytes32 => bool) _usedDigests;

        uint256 _CHAIN_ID;

        bytes32 _DOMAIN_SEPARATOR;
    }

    bytes32 internal constant STORAGE_SLOT = keccak256("footage.contracts.storage.ERC721Club");
    function layout() internal pure returns (Layout storage l) {
        bytes32 slot = STORAGE_SLOT;
        assembly {
            l.slot := slot
        }
    }
}