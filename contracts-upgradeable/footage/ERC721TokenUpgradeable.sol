// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import { ERC721SeaDropUpgradeable } from "../ERC721SeaDropUpgradeable.sol";
import { ERC721TokenStorage } from "./ERC721TokenStorage.sol";

/*
 * @notice This contract uses ERC721PartnerSeaDrop,
 *         an ERC721A token contract that is compatible with SeaDrop.
 *         The set burn address is the only sender that can burn tokens.
 */
contract ERC721TokenUpgradeable is ERC721SeaDropUpgradeable {
    using ERC721TokenStorage for ERC721TokenStorage.Layout;

    /**
     * @notice A token can only be burned by the set burn address.
     */
    error BurnIncorrectSender();

    /**
     * @notice Initialize the token contract with its name, symbol,
     *         administrator, and allowed SeaDrop addresses.
     */
    function initialize(
        string memory name,
        string memory symbol,
        address[] memory allowedSeaDrop
    ) external initializer initializerERC721A {
        ERC721SeaDropUpgradeable.__ERC721PartnerSeaDrop_init(
            name,
            symbol,
            allowedSeaDrop
        );
    }

    function setBurnAddress(address newBurnAddress) external onlyOwner {
        ERC721TokenStorage.layout().burnAddress = newBurnAddress;
    }

    function getBurnAddress() public view returns (address) {
        return ERC721TokenStorage.layout().burnAddress;
    }

    /**
     * @notice Destroys `tokenId`, only callable by the set burn address.
     *
     * @param tokenId The token id to burn.
     */
    function burn(uint256 tokenId) external {
        if (msg.sender != ERC721TokenStorage.layout().burnAddress) {
            revert BurnIncorrectSender();
        }
        _burn(tokenId);
    }
}
