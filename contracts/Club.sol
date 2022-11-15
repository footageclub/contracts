// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "./LandToken.sol";


contract Club is AccessControl, ERC721, ERC721Holder {
    using ECDSAUpgradeable for bytes32;

    bytes32 public constant SIGNER_ROLE = keccak256("SIGNER_ROLE");

    // 质押事件
    event PledgeTransfer(address indexed to, uint256 indexed tokenId, uint256 indexed landTokenId);

    // 解除质押
    event UnpledgeTransfer(address indexed to, uint256 indexed tokenId, uint256 indexed landTokenId);

    // 记录质押地块 LandId => CludId 
    mapping(uint256 => uint256) private pledged;

    // CludId => LandId
    mapping(uint256 => uint256[]) private clubPledged;
    
    // CludId => LandId => index
    mapping(uint256 => mapping(uint256 => uint)) private clubPledgedIndex;

    LandToken public land; 

    constructor(address _owner, string memory _name, string memory _symbol, LandToken _land) ERC721(_name, _symbol) {
        _setupRole(SIGNER_ROLE, _owner);
        _setupRole(DEFAULT_ADMIN_ROLE, _owner);
        land = _land;
    }

    /**
     * @dev See {IERC165-supportsInterface}.
     */
    function supportsInterface(bytes4 interfaceId) 
      public 
      view 
      virtual 
      override(AccessControl, ERC721)
      returns (bool) 
      {
        return super.supportsInterface(interfaceId);
    }

    /**
     * @dev 实现地块质押
     */
    function pledge(uint256 tokenId, uint256 landTokenId, bytes calldata sign) external {
        require(hasRole(SIGNER_ROLE, keccak256(abi.encodePacked(this, msg.sender, tokenId)).toEthSignedMessageHash().recover(sign)), "sign err");
        require(!_requireLedged(landTokenId), "Ledged land");
        address landOwner = land.ownerOf(landTokenId);
        require(landOwner == msg.sender, "address not a valid land owner");
        require(tokenId > 0 && (land.getApproved(landTokenId) ==  address(this) || land.isApprovedForAll(msg.sender, address(this))), "Invalid token ID");

        if (!_exists(tokenId)) {
            _safeMint(msg.sender, tokenId);
        } else {
            address owner = ownerOf(tokenId);
            require(owner == msg.sender, "address not a valid owner");
        }

        land.safeTransferFrom(msg.sender, address(this), landTokenId);
        pledged[landTokenId] = tokenId;
        clubPledged[tokenId].push(landTokenId);
        clubPledgedIndex[tokenId][landTokenId] = clubPledged[tokenId].length - 1;

        emit PledgeTransfer(msg.sender, tokenId, landTokenId);
    }

    /**
     * @dev 实现地块赎回
     */
    function unpledge(uint256 tokenId, uint256 landTokenId, bytes calldata sign) external {
        require(hasRole(SIGNER_ROLE, keccak256(abi.encodePacked(this, msg.sender, tokenId)).toEthSignedMessageHash().recover(sign)), "sign err");
        _requireMinted(tokenId);
        address owner = ownerOf(tokenId);
        require(owner == msg.sender, "address not a valid owner");
        require(pledged[landTokenId] == tokenId, "Invalid Owner");
        require(land.ownerOf(landTokenId) == address(this), "Invalid token ID");

        land.safeTransferFrom(address(this), owner, landTokenId);
        pledged[landTokenId] = 0;
        clubPledged[tokenId][clubPledgedIndex[tokenId][landTokenId]]  =  clubPledged[tokenId][clubPledged[tokenId].length -1];
        clubPledged[tokenId].pop();

        delete pledged[landTokenId];
        delete clubPledgedIndex[tokenId][landTokenId];

        emit UnpledgeTransfer(msg.sender, tokenId, landTokenId);
    }
 
    function countClubLand(uint256 tokenId) external view returns(uint) {
        return clubPledged[tokenId].length;
    }

    function clubLands(uint256 tokenId) external view returns(uint256[] memory) {
        uint256[] memory lands = new uint256[](clubPledged[tokenId].length);
        for(uint i=0; i < clubPledged[tokenId].length;i++) {
            lands[i] =  clubPledged[tokenId][i];
        }
        return lands;
    }

    function tokenPledged(uint256 tokenId) external view returns (uint256) {
        return pledged[tokenId];
    }

    function _requireLedged(uint256 landTokenId) internal view virtual returns (bool) {
        return pledged[landTokenId] > 0;
    }
}