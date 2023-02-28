import { ethers } from "hardhat";

const ownerAddress = "0x350d6519578E52Fdcd4e5a74F92814f832419949";

async function main() {
  const PausableZoneController = await ethers.getContractFactory("PausableZoneController");
  console.log("Deploying...");
  const pausableZoneController = await PausableZoneController.deploy(ownerAddress);
  await pausableZoneController.deployed();
  console.log(`Zone deployed to ${pausableZoneController.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
