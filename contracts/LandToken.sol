// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/utils/ERC721HolderUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/StringsUpgradeable.sol";
import "./libs/HasSecondarySaleFeesUpgradeable.sol";

contract LandToken is Initializable, AccessControlUpgradeable, ERC721EnumerableUpgradeable, ERC721HolderUpgradeable, HasSecondarySaleFeesUpgradeable {
  using ECDSAUpgradeable for bytes32;
  using StringsUpgradeable for uint256;

  bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

  // Optional mapping for token URIs
  mapping(uint256 => string) private _tokenURIs;

  // 定义NFT元数据地址前缀
  string public tokenURIPrefix;

  // 定义合约地址
  string public contractURI;

  // 定义二级市场版税收益账户和费用
  address payable public rcAddress;

  // 费用计算: 0.25 * 10000 (百分比 * 10000)
  uint256 public rcValue;

  /// @custom:oz-upgrades-unsafe-allow constructor
  constructor() {
      _disableInitializers();
  }

  function initialize(address _owner, address payable _recipient, uint256 _fee, string memory _name, string memory _symbol, string memory _tokenURIPrefix, string memory _contractURI) public initializer {
      __Context_init_unchained();
      __ERC165_init_unchained();
      __AccessControl_init_unchained();
      __ERC721Enumerable_init_unchained();
      __ERC721Holder_init_unchained();
      __ERC721_init_unchained(_name, _symbol);
      __HasSecondarySaleFees_init_unchained();

      // 权限设置
      _setupRole(DEFAULT_ADMIN_ROLE, _owner);
      _setupRole(MINTER_ROLE, _owner);

      tokenURIPrefix = _tokenURIPrefix;
      contractURI = _contractURI;
      rcAddress = _recipient;
      rcValue = _fee;
  }

  /**
   * @dev See {IERC165-supportsInterface}.
   */
  function supportsInterface(bytes4 interfaceId) 
      public 
      view 
      virtual 
      override(AccessControlUpgradeable, ERC165StorageUpgradeable, ERC721EnumerableUpgradeable)
      returns (bool) 
      {
        return super.supportsInterface(interfaceId);
  }

  /**
    * @dev See {IERC721Metadata-tokenURI}.
    */
  function tokenURI(uint256 tokenId) public view virtual override returns (string memory) {
      _requireMinted(tokenId);

      string memory _tokenURI = _tokenURIs[tokenId];
      string memory base = _baseURI();

      // If there is no base URI, return the token URI.
      if (bytes(base).length == 0) {
          return _tokenURI;
      }
      // If both are set, concatenate the baseURI and tokenURI (via abi.encodePacked).
      if (bytes(_tokenURI).length > 0) {
          return string(abi.encodePacked(base, _tokenURI));
      }

      return super.tokenURI(tokenId);
  }

  /**
   * @dev 获取费用收益地址 
   */
  function getFeeRecipients(uint256 id) public view override returns (address payable[] memory) {
      require(id > 0, "id 0");
      address payable[] memory result = new address payable[](1);
      result[0] = rcAddress;
      return result;
  }

  /**
   * @dev 获取费用收益数值
   */
  function getFeeBps(uint256 id) public view override returns (uint[] memory) {
      require(id > 0, "id 0");
      uint[] memory result = new uint[](1);
      result[0] = rcValue;
      return result;
  }

  /**
   * @dev 设置收益账户与数值 
   */
  function setFeeRecipient(address payable _recipient, uint256 _fee) public onlyRole(DEFAULT_ADMIN_ROLE) {
    require(_recipient != address(0), "Recipient be present");
    require(_fee > 0, "Fee value be positive");

    rcAddress = _recipient;
    rcValue = _fee;
    address[] memory recipients = new address[](1);
    recipients[0] = _recipient;

    uint[] memory bps = new uint[](1);
    bps[0] = _fee;

    emit SecondarySaleFees(0, recipients, bps);
  }

   /**
   * 获取用户所有的NFT
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
  function withdraw(address payable to) external onlyRole(DEFAULT_ADMIN_ROLE) {
      require(to != address(0), "address zero");
      to.transfer(address(this).balance);
  }

  /**
   * @dev 查看合约中的余额
   */
  function balance() external view onlyRole(DEFAULT_ADMIN_ROLE) returns (uint256) {
    return address(this).balance;
  }

  /**
   * @dev 创建
   */
  function mint(uint amount, uint256 tokenId, string memory tokenUrl, bytes calldata sign) external payable {
     require(hasRole(MINTER_ROLE, keccak256(abi.encodePacked(this, msg.sender, amount, tokenId, tokenUrl)).toEthSignedMessageHash().recover(sign)), "sign err");
     require(msg.value == amount, "Ether value sent is not correct");
     _safeMint(msg.sender, tokenId);
     _setTokenURI(tokenId, tokenUrl);
  }


  /**
   * @dev Internal function to set the contract URI
   * @param _contractURI string URI prefix to assign
   */
  function setContractURI(string memory _contractURI) public onlyRole(DEFAULT_ADMIN_ROLE) {
        contractURI = _contractURI;
  }

  /**
   * @dev 设置NFT前缀
   */
  function setBaseURI(string memory _tokenURIPrefix) public onlyRole(DEFAULT_ADMIN_ROLE) {
        tokenURIPrefix = _tokenURIPrefix;
  }

  /**
   * @dev Base URI for computing {tokenURI}. If set, the resulting URI for each
   * token will be the concatenation of the `baseURI` and the `tokenId`. Empty
   * by default, can be overridden in child contracts.
   */
  function _baseURI() internal view virtual override returns (string memory) {
      return tokenURIPrefix;
  }

  /**
   * @dev Sets `_tokenURI` as the tokenURI of `tokenId`.
   *
   * Requirements:
   *
   * - `tokenId` must exist.
   */
  function _setTokenURI(uint256 tokenId, string memory _tokenURI) internal virtual {
      require(_exists(tokenId), "ERC721URIStorage: URI set of nonexistent token");
      _tokenURIs[tokenId] = _tokenURI;
  }
}
