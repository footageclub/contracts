// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/token/ERC1155/ERC1155Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC1155/extensions/ERC1155BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC1155/extensions/ERC1155SupplyUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/math/SafeMathUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol";

/// @custom:security-contact info@footage.club
contract ClubNFTPublic is Initializable, ERC1155Upgradeable, AccessControlUpgradeable, PausableUpgradeable, ERC1155BurnableUpgradeable, ERC1155SupplyUpgradeable, UUPSUpgradeable {
    using SafeMathUpgradeable for uint256;
    using ECDSAUpgradeable for bytes32;

    bytes32 public constant URI_SETTER_ROLE = keccak256("URI_SETTER_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant SIGNER_ROLE = keccak256("SIGNER_ROLE");
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");
    bytes32 public constant WITHDRAW_ROLE = keccak256("WITHDRAW_ROLE");

    // 设置事件
    event ClubNFTSetting(uint256 indexed clubId, uint256[] indexed tokenids, uint[] indexed totalSupplys, uint[] peerSupplys, uint256[] prices);

    // 收益人地址 clubid -> address
    mapping(uint256 => address) public Recipients;
    mapping(uint256 => uint256) public RecipientAmount;

    // 版权收益人地址 clubid -> address
    mapping(uint256 => address) public RRecipients;

    // 版权比例 clubid -> Percentage
    mapping(uint256 => uint) public RPercentages;

    // 平台费用
    uint public PlatformFee;
    
    // token_id - > clubid 
    mapping(uint256 => uint256) private _clubTokens;

    // 限制tokenid铸造数量 tokenId -> clubId -> limit
    mapping(uint256 => mapping(uint256 => uint)) private _totalSupplyLimit;

    // 限制单个用户tokenid铸造数量 tokenId -> clubId -> limit
    mapping(uint256 => mapping(uint256 => uint)) private _peerSupplyLimit;

    // nft的价格设置
    mapping(uint256 => uint256) private _salePrices;

    // 铸造开始/结束时间
    mapping(uint256 => mapping(uint256 => uint256[])) private _mintDateLimit; 

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(string memory uri_) initializer public {
        __ERC1155_init(uri_);
        __AccessControl_init();
        __Pausable_init();
        __ERC1155Burnable_init();
        __ERC1155Supply_init();
        __UUPSUpgradeable_init();

        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(URI_SETTER_ROLE, msg.sender);
        _grantRole(PAUSER_ROLE, msg.sender);
        _grantRole(SIGNER_ROLE, msg.sender);
        _grantRole(UPGRADER_ROLE, msg.sender);
        _grantRole(WITHDRAW_ROLE, msg.sender);

        PlatformFee = 250;
    }

    function setURI(string memory newuri) public onlyRole(URI_SETTER_ROLE) {
        _setURI(newuri);
    }

    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    function mint(uint256 clubId, uint256 tokenId, uint256 amount, bytes calldata sign) external payable {
        require(hasRole(SIGNER_ROLE, keccak256(abi.encodePacked(this, msg.sender,clubId, tokenId, amount)).toEthSignedMessageHash().recover(sign)), "sign err");
        _requireRegister(clubId, tokenId);
        uint256 price = _salePrices[tokenId];
        if (price > 0) {
             uint paying = price.mul(amount);
             require(paying > 0 && msg.value >= paying, "msg.value is incorrect");
        }
        
        uint totalSupplyToken = _totalSupplyLimit[tokenId][clubId];
        uint pSupplyToken = _peerSupplyLimit[tokenId][clubId];
        require(totalSupplyToken == 0 || totalSupply(tokenId) + amount <= totalSupplyToken, "total limit");
        
        uint256 userAmount = balanceOf(msg.sender, tokenId);
        require(userAmount + amount <=  pSupplyToken, "peer total limit");

        if (_mintDateLimit[tokenId][clubId].length > 0 ) {
            require(block.timestamp > _mintDateLimit[tokenId][clubId][0], "start at");

            if (_mintDateLimit[tokenId][clubId].length > 1 ) {
                require(block.timestamp < _mintDateLimit[tokenId][clubId][1], "endt at");
            }
        } 

        _mint(msg.sender, tokenId, amount, "");
        RecipientAmount[clubId]+=msg.value;
    }

    function _beforeTokenTransfer(address operator, address from, address to, uint256[] memory ids, uint256[] memory amounts, bytes memory data)
        internal
        whenNotPaused
        override(ERC1155Upgradeable, ERC1155SupplyUpgradeable)
    {
        super._beforeTokenTransfer(operator, from, to, ids, amounts, data);
    }

    function _authorizeUpgrade(address newImplementation)
        internal
        onlyRole(UPGRADER_ROLE)
        override
    {}

    // The following functions are overrides required by Solidity.
    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC1155Upgradeable, AccessControlUpgradeable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }

    function setRecipientAddress(uint256 clubId, address recipientAddress_) external onlyRole(DEFAULT_ADMIN_ROLE) {
        Recipients[clubId] = recipientAddress_;
    }

    function setRRecipientAddress(uint256 clubId, address rRecipientAddress_)  external onlyRole(DEFAULT_ADMIN_ROLE) {
        RRecipients[clubId] = rRecipientAddress_;
    }

    function setRPercentage(uint256 clubId, uint rPercentage_) external onlyRole(DEFAULT_ADMIN_ROLE) {
         RPercentages[clubId] = rPercentage_;
    }

    function setClubNFT(uint256 clubId, uint256[] memory tokenids, uint[] memory totalSupplys, uint[] memory peerSupplys, uint256[] memory prices, bytes calldata sign) external {
        require(hasRole(SIGNER_ROLE, keccak256(abi.encodePacked(this, msg.sender, clubId, tokenids, totalSupplys, peerSupplys, prices)).toEthSignedMessageHash().recover(sign)), "sign err");
        require(tokenids.length == totalSupplys.length, "data err");
        require(tokenids.length == peerSupplys.length, "data err");
        require(tokenids.length == prices.length, "data err");
        for(uint i = 0; i <  tokenids.length;i ++) {
            if(_clubTokens[tokenids[i]] != 0 && _clubTokens[tokenids[i]] != clubId) {
                revert("bad token");
            }
        }

        for(uint i = 0; i <  tokenids.length;i ++) {
            _totalSupplyLimit[tokenids[i]][clubId] = totalSupplys[i];
            _peerSupplyLimit[tokenids[i]][clubId] = peerSupplys[i];
            _salePrices[tokenids[i]] = prices[i];
            _clubTokens[tokenids[i]] = clubId;
        }

        emit ClubNFTSetting(clubId, tokenids, totalSupplys, peerSupplys, prices);
    }

    function withdraw(uint256 clubId, bytes calldata sign) external {
       require(hasRole(WITHDRAW_ROLE, keccak256(abi.encodePacked(this, msg.sender, clubId, "withdraw")).toEthSignedMessageHash().recover(sign)), "sign err");
       require(Recipients[clubId] == msg.sender && Recipients[clubId] != address(0), "bad sender");

       uint256 amount = RecipientAmount[clubId];
       RecipientAmount[clubId] = 0;
       require(amount > 0, "zero amount");

       (bool success,) = payable(msg.sender).call{value: amount}("");
       require(success, "Failed to send Enter");
    }

    function _requireRegister(uint256 clubId, uint256 tokenId) internal view virtual {
        require(_clubTokens[tokenId] == clubId && _clubTokens[tokenId] > 0, "ClubNFTPublic: invalid token ID");
    }
}