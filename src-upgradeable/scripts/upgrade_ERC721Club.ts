import fs from "fs";
import { ethers, upgrades } from "hardhat";

async function main() {
  const Club = await ethers.getContractFactory("ERC721ClubUpgradeable");
  console.log("Upgrading...");
  let addresses = JSON.parse(
    fs.readFileSync("deployment-addresses-ERC721ClubUpgradeable.json").toString()
  );
  const result = await upgrades.upgradeProxy(addresses.proxy, Club);
  await result.deployTransaction.wait();
  console.log("Upgraded");

  addresses = {
    proxy: addresses.proxy,
    admin: await upgrades.erc1967.getAdminAddress(addresses.proxy),
    implementation: await upgrades.erc1967.getImplementationAddress(
      addresses.proxy
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