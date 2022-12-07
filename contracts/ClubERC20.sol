// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/// @custom:security-contact info@footage.club
/// @comment UUPS?
contract ClubERC20 is ERC20, ERC20Burnable, Pausable, Ownable {
    uint256 public Price;

    constructor(string memory _name, string memory _symbol, uint256 _price) ERC20(_name, _symbol) {
        Price = _price;
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    /// @comment 这里的 `Price` 和 mint 的 `amount` 没有关系？只要发送 `Print` 个 ETH，就能 mint 任意数量的 `ClubERC20` 吗？
    function mint(address to, uint256 amount) external payable {
        require(msg.value >= Price, "Ether value sent is not correct");
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
        (bool success,) = payable(msg.sender).call{value: address(this).balance}("");
        require(success, "Failed to send Enter");
    }
}