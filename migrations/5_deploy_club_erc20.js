const ClubERC20 = artifacts.require("ClubERC20");

module.exports = async function(deployer) {    
    deployer.deploy(ClubERC20, "Club", "ERC20", 68000).then(function(instance){
        console.log('ClubERC20 Deployed', instance.address);
    });
};