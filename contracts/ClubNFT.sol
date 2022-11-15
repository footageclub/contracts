// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts/token/ERC1155/extensions/ERC1155Supply.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

contract ClubNFT is Ownable, ERC1155Supply {
    using SafeMath for uint256;

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


    constructor(address recipientAddress_, address rRecipientAddress_, uint rPercebtage_, uint platformFee_, string memory uri_) ERC1155(uri_){
        RecipientAddress = recipientAddress_;
        RRecipientAddress = rRecipientAddress_;
        RPercebtage = rPercebtage_;
        PlatformFee = platformFee_;
    }

    function setRecipientAddress(address recipientAddress_) external onlyOwner() {
        RecipientAddress = recipientAddress_;
    }

    function setRRecipientAddress(address rRecipientAddress_)  external onlyOwner() {
        RRecipientAddress = rRecipientAddress_;
    }

    function setRPercebtage(uint rPercebtage_) external onlyOwner() {
         RPercebtage = rPercebtage_;
    }

    function setClubNFT(uint256[] memory tokenids, uint[] memory totalSupplys, uint[] memory peerSupplys, uint256[] memory prices) external onlyOwner() {
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

     function withdraw() external onlyOwner {
       uint amount = address(this).balance;
       payable(RecipientAddress).transfer(amount);
    }

    function mint(uint256 tokenId, uint256 amount) external payable {
        uint paying = _salePrices[tokenId].mul(amount);
        require(paying > 0 && msg.value >= paying, "msg.value is incorrect");
        require(_totalSupplyLimit[tokenId] > 0 && totalSupply(tokenId) + amount <= _totalSupplyLimit[tokenId], "total limit");
        

        uint256 userAmount = balanceOf(msg.sender, tokenId);
        require(userAmount + amount <=  _peerSupplyLimit[tokenId], "peer total limit");

        if (_mintDateLimit[tokenId].length > 0 ) {
            require(block.timestamp > _mintDateLimit[tokenId][0], "start at");

            if (_mintDateLimit[tokenId].length > 1 ) {
                require(block.timestamp < _mintDateLimit[tokenId][1], "endt at");
            }
        } 

        _mint(msg.sender, tokenId, amount, "");
    }
}