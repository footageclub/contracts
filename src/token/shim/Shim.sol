// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/**
 * @dev HardHat doesn't support multiple source folders; so import everything
 * extra that hardhat tests rely on so they get compiled. Allows for faster
 * feedback than running an extra yarn build.
 */
import { TestTokenERC721 } from "../test/TestTokenERC721.sol";
import { MaliciousRecipient } from "../test/MaliciousRecipient.sol";
