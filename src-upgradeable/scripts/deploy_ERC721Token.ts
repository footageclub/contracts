import fs from "fs";
import { ethers, upgrades } from "hardhat";

async function main() {
  const FootageToken = await ethers.getContractFactory("ERC721TokenUpgradeable");
  console.log("Deploying...");
  const footageToken = await upgrades.deployProxy(
    FootageToken,
    [
      "Footage",
      "Super Bee",
      "0x1fA90a9A760d21105aa25FA52e473d2536803F0b",
      ["0x8Ca5014AaC34ABa3Cd050C9112Fe5CD9f4Bb0D43"],
    ],
    { initializer: "initialize" }
  );
  await footageToken.deployed();
  const addresses = {
    proxy: footageToken.address,
    admin: await upgrades.erc1967.getAdminAddress(footageToken.address),
    implementation: await upgrades.erc1967.getImplementationAddress(
      footageToken.address
    ),
  };

  console.log(`ERC721TokenUpgradeable deployed to:`, addresses); 
  try {
    await (run as any)("verify", { address: addresses.implementation });
  } catch (e) {}

  fs.writeFileSync("deployment-addresses-ERC721TokenUpgradeable.json", JSON.stringify(addresses));
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });
  