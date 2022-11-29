const { deployProxy } = require('@openzeppelin/truffle-upgrades');
const LandToken = artifacts.require("LandToken");
const Club = artifacts.require("Club");

module.exports = async function(deployer) {
    const land = await LandToken.deployed();
    const instance = await deployProxy(Club, ["Footage", "Club", land.address, "https://www.footage.club/"], {
        deployer:deployer,
        initializer:'initialize',
        kind: 'uups'
    });

    console.log('Club Deployed', instance.address);
};