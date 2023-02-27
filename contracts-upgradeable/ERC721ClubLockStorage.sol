// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { MintParams } from '../lib/SeaDropStructsUpgradeable.sol';

library ERC721ClubLockStorage {
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

    bytes32 internal constant _SIGNED_MINT_TYPEHASH =
    keccak256(
        "MintClublock("
            "address minter,"
            "address to,"
            "bytes32 assets,"
            "uint256 salt"
        ")"
    );


    struct Layout {
        /// @notice The only address that can burn tokens on this contract.
        address burnAddress;

        /// @notice 地块资产存储
        mapping(string => uint256) _clublockByAssetId;
        mapping(uint256 => string) _clublockByTokenId;

        /// @notice 签名者
        address _signer;

        /// @notice 掉落配置
        MintParams _mintParams;

        /// @notice 用于存储已使用过的签名
        mapping(bytes32 => bool) _usedDigests;

        /// @notice 记录链ID
        uint256 _CHAIN_ID;

        /// @notice domain hash
        bytes32 _DOMAIN_SEPARATOR;

        /// @notice 创建者地址
        address _creatorPayoutAddress;
    }

    bytes32 internal constant STORAGE_SLOT = keccak256("footage.contracts.storage.ERC721ClubLock");
    function layout() internal pure returns (Layout storage l) {
        bytes32 slot = STORAGE_SLOT;
        assembly {
            l.slot := slot
        }
    }
}