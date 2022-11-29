// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/utils/introspection/ERC165StorageUpgradeable.sol";

abstract contract HasSecondarySaleFeesUpgradeable is ERC165StorageUpgradeable {
    event SecondarySaleFees(uint256 tokenId,address[] recipients, uint[] bps);
    /*
     * bytes4(keccak256('getFeeBps(uint256)')) == 0x0ebd4c7f
     * bytes4(keccak256('getFeeRecipients(uint256)')) == 0xb9c4d9fb
     * bytes4(keccak256('_SECONDARY_SALE_INTERFACE_ID')) == 0x957d5e37
     *
     * => 0x0ebd4c7f ^ 0xb9c4d9fb ^ 0x957d5e37 == 0x2204cbb3
     */
    bytes4 private constant _SECONDARY_SALE_INTERFACE_ID = 0x2204cbb3;

    function __HasSecondarySaleFees_init() internal onlyInitializing {
        __HasSecondarySaleFees_init_unchained();
    }

    function __HasSecondarySaleFees_init_unchained() internal onlyInitializing {
         _registerInterface(_SECONDARY_SALE_INTERFACE_ID);
    }

    function getFeeRecipients(uint256 id) public view virtual returns (address payable[] memory);
    function getFeeBps(uint256 id) public view virtual returns (uint[] memory);

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}
