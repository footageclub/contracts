const { deployProxy } = require('@openzeppelin/truffle-upgrades');
const ClubNFTPublic = artifacts.require("ClubNFTPublic");

module.exports = async function(deployer) {    
    const instance = await deployProxy(ClubNFTPublic, ["https://www.footage.club/{id}.json"], {
        deployer:deployer,
        initializer:'initialize',
        kind:'uups'
    });
    console.log('ClubNFTPublic Deployed', instance.address);

};