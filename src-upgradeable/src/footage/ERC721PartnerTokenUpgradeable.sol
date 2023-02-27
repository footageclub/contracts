// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import { ERC721PartnerSeaDropUpgradeable } from "../ERC721PartnerSeaDropUpgradeable.sol";
import { ERC721PartnerTokenStorage } from "./ERC721PartnerTokenStorage.sol";

contract ERC721PartnerTokenUpgradeable is ERC721PartnerSeaDropUpgradeable {
    using ERC721PartnerTokenStorage for ERC721PartnerTokenStorage.Layout;

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
        address administrator,
        address[] memory allowedSeaDrop
    ) external initializer initializerERC721A {
        ERC721PartnerSeaDropUpgradeable.__ERC721PartnerSeaDrop_init(
            name,
            symbol,
            administrator,
            allowedSeaDrop
        );
    }

    function setBurnAddress(address newBurnAddress) external onlyOwner {
        ERC721PartnerTokenStorage.layout().burnAddress = newBurnAddress;
    }

    function getBurnAddress() public view returns (address) {
        return ERC721PartnerTokenStorage.layout().burnAddress;
    }

    /**
     * @notice Destroys `tokenId`, only callable by the set burn address.
     *
     * @param tokenId The token id to burn.
     */
    function burn(uint256 tokenId) external {
        if (msg.sender != ERC721PartnerTokenStorage.layout().burnAddress) {
            revert BurnIncorrectSender();
        }
        _burn(tokenId);
    }
}
