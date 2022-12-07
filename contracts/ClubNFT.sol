
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/token/ERC1155/extensions/ERC1155Burnable.sol";
import "@openzeppelin/contracts/token/ERC1155/extensions/ERC1155Supply.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

/// @custom:security-contact info@footage.club
/// @comments ERC1155 的合约也可以支持 UUPS 升级 (ERC1155Upgradeable)
contract ClubNFT is ERC1155, AccessControl, Pausable, ERC1155Burnable, ERC1155Supply {
    using SafeMath for uint256;

    bytes32 public constant URI_SETTER_ROLE = keccak256("URI_SETTER_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant WITHDRAW_ROLE = keccak256("WITHDRAW_ROLE");

    // 设置事件
    event ClubNFTSetting(uint256[] indexed tokenids, uint[] indexed totalSupplys, uint[] indexed peerSupplys, uint256[] prices);

    // 收益人地址
    address public RecipientAddress;

    // 版权收益人地址
    address public RRecipientAddress;

    // 版权比例
    uint public RPercebtage;

    // 平台费用
    uint public PlatformFee;

    // 限制tokenid铸造数量
    mapping(uint256 => uint) private _totalSupplyLimit;

    // 限制单个用户tokenid铸造数量
    mapping(uint256 => uint) private _peerSupplyLimit;

    // nft的价格设置
    mapping(uint256 => uint256) private _salePrices;

    // 铸造开始/结束时间
    mapping(uint256 => uint256[]) private _mintDateLimit; 

    constructor(address recipientAddress_, address rRecipientAddress_, uint rPercebtage_, uint platformFee_, string memory uri_) ERC1155(uri_) {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(URI_SETTER_ROLE, msg.sender);
        _grantRole(PAUSER_ROLE, msg.sender);
        _grantRole(WITHDRAW_ROLE, msg.sender);

        RecipientAddress = recipientAddress_;
        RRecipientAddress = rRecipientAddress_;
        RPercebtage = rPercebtage_;
        PlatformFee = platformFee_;
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

    function mint(uint256 tokenId, uint256 amount, bytes memory data) external payable {
        /// @comment 数字计算可以使用 OpenZeppline 的 SafeMathUpgradeable
        uint paying = _salePrices[tokenId].mul(amount);
        /// @comment 可以严格要求 msg.value == paying，防止用户超付
        require(paying > 0 && msg.value >= paying, "msg.value is incorrect");
        require(_totalSupplyLimit[tokenId] > 0 && totalSupply(tokenId) + amount <= _totalSupplyLimit[tokenId], "total limit");
        
        /// @comment `msg.sender` 和 `msg.data` 可以统一用 `_msgSender()` 和 `_msgData()` 来替代
        uint256 userAmount = balanceOf(msg.sender, tokenId);
        require(userAmount + amount <=  _peerSupplyLimit[tokenId], "peer total limit");

        if (_mintDateLimit[tokenId].length > 0 ) {
            require(block.timestamp > _mintDateLimit[tokenId][0], "start at");

            if (_mintDateLimit[tokenId].length > 1 ) {
                require(block.timestamp < _mintDateLimit[tokenId][1], "endt at");
            }
        } 

        _mint(msg.sender, tokenId, amount, data);
    }

    function _beforeTokenTransfer(address operator, address from, address to, uint256[] memory ids, uint256[] memory amounts, bytes memory data)
        internal
        whenNotPaused
        override(ERC1155, ERC1155Supply)
    {
        super._beforeTokenTransfer(operator, from, to, ids, amounts, data);
    }

    // The following functions are overrides required by Solidity.

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC1155, AccessControl)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }

    function setRecipientAddress(address recipientAddress_) external onlyRole(DEFAULT_ADMIN_ROLE) {
        RecipientAddress = recipientAddress_;
    }

    function setRRecipientAddress(address rRecipientAddress_)  external onlyRole(DEFAULT_ADMIN_ROLE){
        RRecipientAddress = rRecipientAddress_;
    }

    function setRPercentage(uint rPercebtage_) external onlyRole(DEFAULT_ADMIN_ROLE) {
        /// @admin 检查一下 `rPercebtage_` 的值是否合法
         RPercebtage = rPercebtage_;
    }

    function setClubNFT(uint256[] memory tokenids, uint[] memory totalSupplys, uint[] memory peerSupplys, uint256[] memory prices) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(tokenids.length == totalSupplys.length, "data err");
        require(tokenids.length == peerSupplys.length, "data err");
        require(tokenids.length == prices.length, "data err");

        for(uint i = 0; i <  tokenids.length;i ++) {
            _totalSupplyLimit[tokenids[i]] = totalSupplys[i];
            _peerSupplyLimit[tokenids[i]] = peerSupplys[i];
            _salePrices[tokenids[i]] = prices[i];
        }

        emit ClubNFTSetting(tokenids, totalSupplys, peerSupplys, prices);
    }

    /**
     * @dev 取出合约中的余额
     */
    /// @comment emit 一个 Withdraw 事件记录一下
    function withdraw() external onlyRole(WITHDRAW_ROLE) {
        (bool success,) = RecipientAddress.call{value: address(this).balance}("");
        require(success, "Failed to send Enter");
    }
}