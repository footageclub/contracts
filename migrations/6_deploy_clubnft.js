const ClubNFT = artifacts.require("ClubNFT");

module.exports = async function(deployer) {    
    deployer.deploy(ClubNFT, "0x4cb561273a2757d33f0d044d5489236d5bbd9e36", "0x4cb561273a2757d33f0d044d5489236d5bbd9e36", 250, 250, "https://www.footage.club/{id}.json").then(function(instance){
        console.log('ClubNFT Deployed', instance.address);
    });
};