import { ethers } from "hardhat";

async function main() {
    const SeaDrop = await ethers.getContractFactory("SeaDrop");
    const seaDrop = await SeaDrop.deploy();
    await seaDrop.deployed();
   console.log(`SeaDrop deployed to: ${seaDrop.address}`); 
}


main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
})
