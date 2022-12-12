// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/utils/ERC721HolderUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol";
import "./LandToken.sol";

/// @custom:security-contact test@qq.com
contract Club is Initializable, ERC721Upgradeable, ERC721URIStorageUpgradeable, PausableUpgradeable, AccessControlUpgradeable, ERC721BurnableUpgradeable, UUPSUpgradeable, ERC721HolderUpgradeable {
    using ECDSAUpgradeable for bytes32;

    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant SIGNER_ROLE = keccak256("SIGNER_ROLE");
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");
    bytes32 public constant URI_SETTER_ROLE = keccak256("URI_SETTER_ROLE");

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

    LandToken public Land; 

    string public BaseURI; 

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(string memory _name, string memory _symbol, LandToken _land, string memory _uri) initializer public {
        __ERC721_init(_name, _symbol);
        __ERC721URIStorage_init();
        __Pausable_init();
        __AccessControl_init();
        __ERC721Burnable_init();
        __UUPSUpgradeable_init();

        _grantRole(DEFAULT_ADMIN_ROLE, _msgSender());
        _grantRole(PAUSER_ROLE, _msgSender());
        _grantRole(SIGNER_ROLE, _msgSender());
        _grantRole(UPGRADER_ROLE, _msgSender());
        _grantRole(URI_SETTER_ROLE, _msgSender());

        Land = _land;
        BaseURI = _uri;
    }

    function _baseURI() internal view override returns (string memory) {
        return BaseURI;
    }

    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    function _beforeTokenTransfer(address from, address to, uint256 tokenId, uint256 batchSize)
        internal
        whenNotPaused
        override
    {
        super._beforeTokenTransfer(from, to, tokenId, batchSize);
    }

    function _authorizeUpgrade(address newImplementation)
        internal
        onlyRole(UPGRADER_ROLE)
        override
    {}

    // The following functions are overrides required by Solidity.

    function _burn(uint256 tokenId)
        internal
        override(ERC721Upgradeable, ERC721URIStorageUpgradeable)
    {
        super._burn(tokenId);
    }

    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721Upgradeable, ERC721URIStorageUpgradeable)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721Upgradeable, AccessControlUpgradeable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }

    /**
     * @dev 改变基础url
     */
    function setBaseURI(string memory uri) external onlyRole(URI_SETTER_ROLE) {
        BaseURI = uri;
    }

     /**
     * @dev 实现地块质押
     */
    function pledge(uint256 tokenId, uint256 landTokenId, string memory _tokenURI, bytes calldata sign) external {
        require(hasRole(SIGNER_ROLE, keccak256(abi.encodePacked(this, _msgSender(), tokenId, landTokenId)).toEthSignedMessageHash().recover(sign)), "sign err");
        require(!_requireLedged(landTokenId), "Ledged land");
        address landOwner = Land.ownerOf(landTokenId);
        require(landOwner == _msgSender(), "address not a valid land owner");
        require(tokenId > 0 && (Land.getApproved(landTokenId) ==  address(this) || Land.isApprovedForAll(_msgSender(), address(this))), "Invalid token ID");

        if (!_exists(tokenId)) {
            _safeMint(_msgSender(), tokenId);
            _setTokenURI(tokenId, _tokenURI);
        } else {
            address owner = ownerOf(tokenId);
            require(owner == _msgSender(), "address not a valid owner");
        }

        Land.safeTransferFrom(_msgSender(), address(this), landTokenId);
        pledged[landTokenId] = tokenId;
        clubPledged[tokenId].push(landTokenId);
        clubPledgedIndex[tokenId][landTokenId] = clubPledged[tokenId].length - 1;

        emit PledgeTransfer(_msgSender(), tokenId, landTokenId);
    }

    /**
     * @dev 实现地块赎回
     */
    function unpledge(uint256 tokenId, uint256 landTokenId, bytes calldata sign) external {
        require(hasRole(SIGNER_ROLE, keccak256(abi.encodePacked(this, _msgSender(), tokenId, landTokenId)).toEthSignedMessageHash().recover(sign)), "sign err");
        _requireMinted(tokenId);
        address owner = ownerOf(tokenId);
        require(owner == _msgSender(), "address not a valid owner");
        require(pledged[landTokenId] == tokenId, "Invalid Owner");
        require(Land.ownerOf(landTokenId) == address(this), "Invalid token ID");

        Land.safeTransferFrom(address(this), owner, landTokenId);
        pledged[landTokenId] = 0;
        clubPledged[tokenId][clubPledgedIndex[tokenId][landTokenId]]  =  clubPledged[tokenId][clubPledged[tokenId].length -1];
        clubPledged[tokenId].pop();

        delete pledged[landTokenId];
        delete clubPledgedIndex[tokenId][landTokenId];

        emit UnpledgeTransfer(_msgSender(), tokenId, landTokenId);
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