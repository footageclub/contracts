import { ethers } from "hardhat";

async function main() {
  const ConduitController = await ethers.getContractFactory("ConduitController");
  const conduitController = await ConduitController.deploy();
  await conduitController.deployed();
  console.log(`ConduitController deployed to ${conduitController.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
