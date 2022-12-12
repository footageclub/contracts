// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
/// comment UUPS?
/// comment 由于该合约为动态部署合约,根据需求无须使用可升级合约。另外,动态合约需要动态部署，如何使用UUPS将造成动态部署困难。
/// @custom:security-contact info@footage.club
contract ClubERC20 is ERC20, ERC20Burnable, Pausable, Ownable {
    using SafeMath for uint256;
    uint256 public Price;

    event Received(address indexed sender, uint indexed value);

    // 支取事件
    event Withdraw(address indexed address_, uint indexed value_);

    constructor(string memory _name, string memory _symbol, uint256 _price) ERC20(_name, _symbol) {
        Price = _price;
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    receive() external payable {
        emit Received(_msgSender(), msg.value);
    }

    /// comment 这里的 `Price` 和 mint 的 `amount` 没有关系？只要发送 `Print` 个 ETH，就能 mint 任意数量的 `ClubERC20` 吗？
    /// comment 已修改
    function mint(address to, uint256 amount) external payable {
        uint paying = Price.mul(amount);
        require(msg.value == paying, "Ether value sent is not correct");
        _mint(to, amount);
    }

    function _beforeTokenTransfer(address from, address to, uint256 amount)
        internal
        whenNotPaused
        override
    {
        super._beforeTokenTransfer(from, to, amount);
    }

    /**
     * @dev 取出合约中的余额
     */
    function withdraw() external onlyOwner {
        uint value =  address(this).balance;
        (bool success,) = payable(_msgSender()).call{value: value}("");
        require(success, "Failed to send Enter");
        emit Withdraw(_msgSender(), value);
    }
}