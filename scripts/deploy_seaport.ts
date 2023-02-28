import { ethers } from "hardhat";

const conduitAddress = "0x350d6519578E52Fdcd4e5a74F92814f832419949";

async function main() {
  const Seaport = await ethers.getContractFactory("Seaport");
  console.log("Deploying...");
  const seaport = await Seaport.deploy(conduitAddress);
  await seaport.deployed();
  console.log(`Seaport deployed to ${seaport.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
