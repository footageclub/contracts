// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";


contract ClubERC20 is Ownable, ERC20 {

    uint256 public price;

    constructor(address _owner, string memory _name, string memory _symbol, uint256 _price) ERC20(_name, _symbol) {
        price = _price;
        _transferOwnership(_owner);
    }

    /**
     * @dev 取出合约中的余额
     */
    function withdraw() external onlyOwner {
       uint amount = address(this).balance;

       payable(msg.sender).transfer(amount);
    }

    /**
     * @dev 查看合约中的余额
     */
    function balance() external view onlyOwner returns (uint256) {
        return address(this).balance;
    }

    function mint() external payable {
        require(msg.value == price, "Ether value sent is not correct");
        _mint(msg.sender, 1);
    }
}