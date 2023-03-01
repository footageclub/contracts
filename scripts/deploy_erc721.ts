// ERC721SeaDrop
import { ethers } from "hardhat";


async function main() {
  const ERC721SeaDrop = await ethers.getContractFactory("ERC721SeaDrop");
  console.log("Deploying...");
  const token = await ERC721SeaDrop.deploy("Footage", "Super Bee", ["0x8Ca5014AaC34ABa3Cd050C9112Fe5CD9f4Bb0D43"]);
  await token.deployed();
  console.log(`ERC721SeaDrop deployed to ${token.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
