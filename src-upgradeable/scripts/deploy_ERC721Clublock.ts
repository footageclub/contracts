import fs from "fs";
import { ethers, upgrades } from "hardhat";

async function main() {
    const ERC721ClubLockUpgradeable = await ethers.getContractFactory("ERC721ClubLockUpgradeable");
    console.log("Deploying...");
    const club = await upgrades.deployProxy(
      ERC721ClubLockUpgradeable,
      [
        "Footage",
        "Clublock",
        "0x5d437a732681eE9c649A6e1c1C44cd37C92F65E0"
      ],
      { initializer: "initialize" }
    );
    await club.deployed();

    const addresses = {
      proxy: club.address,
      admin: await upgrades.erc1967.getAdminAddress(club.address),
      implementation: await upgrades.erc1967.getImplementationAddress(
          club.address
      ),
    };
  console.log(`ERC721ClubLockUpgradeable deployed to:`, addresses); 
  try {
    await (run as any)("verify", { address: addresses.implementation });
  } catch (e) {}

  fs.writeFileSync("deployment-addresses-ERC721ClubLockUpgradeable.json", JSON.stringify(addresses));
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });
  