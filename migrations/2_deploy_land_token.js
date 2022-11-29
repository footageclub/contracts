const {deployProxy } = require('@openzeppelin/truffle-upgrades');
const LandToken = artifacts.require("LandToken");

module.exports = async function(deployer) {
    const instance = await deployProxy(LandToken, ["Footage", "Land", "https://www.footage.club/"],{
        deployer:deployer,
        initializer:'initialize',
        kind:'uups'
    } )

    console.log('LandToken Deployed', instance.address);
};