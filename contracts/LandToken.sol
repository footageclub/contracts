// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/StringsUpgradeable.sol";
import "./libs/HasSecondarySaleFeesUpgradeable.sol";

/// @custom:security-contact info@footage.club
contract LandToken is Initializable, ERC721Upgradeable, ERC721EnumerableUpgradeable, ERC721URIStorageUpgradeable, PausableUpgradeable, AccessControlUpgradeable, ERC721BurnableUpgradeable, UUPSUpgradeable, HasSecondarySaleFeesUpgradeable {
    using ECDSAUpgradeable for bytes32;
    using StringsUpgradeable for uint256;

    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");
    bytes32 public constant WITHDRAW_ROLE = keccak256("WITHDRAW_ROLE");

    event Received(address indexed sender, uint indexed value);

     // 支取事件
    event Withdraw(address indexed address_, uint indexed value_);

    // 定义二级市场版税收益账户和费用
    address payable public RRecipientAddress;

    // 费用计算: 0.25 * 10000 (百分比 * 10000)
    uint256 public RPercebtage;

    // 设置基础uri
    string public baseURI;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(string memory _name, string memory _symbol, string memory _uri) initializer public {
        __ERC721_init(_name, _symbol);
        __ERC721Enumerable_init();
        __ERC721URIStorage_init();
        __Pausable_init();
        __AccessControl_init();
        __ERC721Burnable_init();
        __UUPSUpgradeable_init();

        /// comment `msg.sender` 改为 `_msgSender()` 更安全
        /// comment 已修改
        _grantRole(DEFAULT_ADMIN_ROLE, _msgSender());
        _grantRole(PAUSER_ROLE, _msgSender());
        _grantRole(MINTER_ROLE, _msgSender());
        _grantRole(UPGRADER_ROLE, _msgSender());
        _grantRole(WITHDRAW_ROLE, _msgSender());

        RRecipientAddress = payable(_msgSender());
        RPercebtage = 250;
        baseURI = _uri;
    }

    function _baseURI() internal view override returns (string memory) {
        return baseURI;
    }

    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    receive() external payable {
        emit Received(_msgSender(), msg.value);
    }

    /// comment 接收 ETH 转帐的合约一般需要定义一个 `receive()` 函数 (https://docs.soliditylang.org/en/latest/contracts.html#receive-ether-function)
    /// comment 已修改
    function mint(uint amount, uint256 tokenId, string memory uri, bytes calldata sign) external payable{
        require(hasRole(MINTER_ROLE, keccak256(abi.encodePacked(this, _msgSender(), msg.value, tokenId, uri)).toEthSignedMessageHash().recover(sign)), "sign err");
        /// comment 可以严格要求 msg.value == amount，防止用户超付
        require(msg.value >= amount, "Ether value sent is not correct");
        /// comment 最好检查一下 `tokenId` 是否已经被占用，不然相当于把之前的 NFT 给覆盖了。
        /// 另外，不清楚具体需求，不过一般 tokenId 都是在 mint 时使用一个 `Counter` 计数器来顺序生成，而不是由外部指定
        _safeMint(_msgSender(), tokenId);
        _setTokenURI(tokenId, uri);
    }

    function _beforeTokenTransfer(address from, address to, uint256 tokenId, uint256 batchSize)
        internal
        whenNotPaused
        override(ERC721Upgradeable, ERC721EnumerableUpgradeable)
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
        override(ERC721Upgradeable, ERC721EnumerableUpgradeable, AccessControlUpgradeable, ERC165StorageUpgradeable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }

    /**
     * @dev 获取费用收益地址 
     */
    /// comment 直接返回 `address` 就好，为何要包装成一个 `address[]` ？
    /// comment 这是对二级市场的定义，会在二级市场需求确认下来后，进行修改
    function getFeeRecipients(uint256 id) public view override returns (address payable[] memory) {
        require(id > 0, "id 0");
        address payable[] memory result = new address payable[](1);
        result[0] = RRecipientAddress;
        return result;
    }

    /**
     * @dev 获取费用收益数值
     */
    /// comment  同上，直接返回 unit 就好
    /// comment 这是对二级市场的定义，会在二级市场需求确认下来后，进行修改
    function getFeeBps(uint256 id) public view override returns (uint[] memory) {
        require(id > 0, "id 0");
        uint[] memory result = new uint[](1);
        result[0] = RPercebtage;
        return result;
    }

    /**
     * @dev 设置收益账户与数值 
     */
    function setFeeRecipient(address payable _recipient, uint256 _fee) public onlyRole(DEFAULT_ADMIN_ROLE) {
        require(_recipient != address(0), "Recipient be present");
        /// comment 多些检查； _fee 不能大于 100
        /// comment 已修改
        require(_fee > 0 && _fee <= 10000, "Fee value be positive");

        RRecipientAddress = _recipient;
        RPercebtage = _fee;
        address[] memory recipients = new address[](1);
        recipients[0] = _recipient;

        uint[] memory bps = new uint[](1);
        bps[0] = _fee;

        /// comment 同上，没必要包装成 `address[]` 和 `uint[]`
        /// comment 这是对二级市场的定义，会在二级市场需求确认下来后，进行修改
        emit SecondarySaleFees(0, recipients, bps);
    }

    /**
     * @dev 获取用户所有的NFT
     */
    function balanceByOwner(address owner) external view returns (uint256[] memory) {
        uint256 amount = balanceOf(owner);
        uint256[] memory balances = new uint256[](amount);

        for (uint256 i = 0; i < amount; i++) {
          balances[i] = tokenOfOwnerByIndex(owner, i);
        } 
        return balances;
    }

    /**
     * @dev 取出合约中的余额
     */
    /// comment 取款时可以 emit 一个事件
    /// comment 已修改
    function withdraw(address payable to) external onlyRole(WITHDRAW_ROLE) {
        require(to != address(0), "address zero");
        uint value = address(this).balance;
        (bool success, ) = to.call{value: value}("");
        require(success, "Failed to send Enter");

        emit Withdraw(to, value);
    }

    /**
     * @dev 改变基础url
     */
    function setBaseURI(string memory uri) external onlyRole(DEFAULT_ADMIN_ROLE) {
        baseURI = uri;
    }
}
