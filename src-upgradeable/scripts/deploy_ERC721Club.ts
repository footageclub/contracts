import fs from "fs";
import { ethers, upgrades } from "hardhat";

async function main() {
    const ERC721ClubUpgradeable = await ethers.getContractFactory("ERC721ClubUpgradeable");
    console.log("Deploying...");
    const club = await upgrades.deployProxy(
      ERC721ClubUpgradeable,
      [
        "Footage",
        "Club",
        "0x5d437a732681eE9c649A6e1c1C44cd37C92F65E0",
        ethers.BigNumber.from("524288000").toString(),
        ethers.BigNumber.from("300000").toString()
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

  console.log(`ERC721ClubUpgradeable deployed to:`, addresses); 
  try {
    await (run as any)("verify", { address: addresses.implementation });
  } catch (e) {}

  fs.writeFileSync("deployment-addresses-ERC721ClubUpgradeable.json", JSON.stringify(addresses));
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });
  