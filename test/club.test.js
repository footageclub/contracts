const Club = artifacts.require("Club");
const LandToken = artifacts.require("LandToken");
const common = require('./common');

/*
 * uncomment accounts to access the test accounts made available by the
 * Ethereum client
 * See docs: https://www.trufflesuite.com/docs/truffle/testing/writing-tests-in-javascript
 */
contract("Club", function ( accounts ) {
  it("should assert true", async function () {
    await Club.deployed();
    return assert.isTrue(true);
  });


  it("should pause", async function() {
    const club = await Club.deployed();
    await club.pause();
    return assert.isTrue(true);
  });

  it("should unpause", async function() {
    const club = await Club.deployed();
    await club.unpause();
    return assert.isTrue(true);
  });

  it("should setBaseURI", async function() {
    const club = await Club.deployed();
    await club.setBaseURI("https://www.footage.club/v2/");
    return assert.isTrue(true);
  });

  it("should pledge", async function(){
    const landToken = await LandToken.deployed();
    const club = await Club.deployed();
    const tx0 = web3.utils.encodePacked(
      {"type":"address", value:landToken.address},
      {"type":"address", value:accounts[0]},
      {"type":"uint", value:680000000},
      {"type":"uint256", value:1980},
      {"type":"string", value:"1980"}
    );

    const tx1 = web3.utils.keccak256(tx0);
    const key = await common.makePrivateKey(process.env.MNEMONIC, process.env.HDPATH);
    const sign = await web3.eth.accounts.sign(tx1, key.key)
    await landToken.mint(680000000,1980, "1980", sign.signature, {from:accounts[0], value: 680000000});
    await landToken.setApprovalForAll(club.address, true);

    const tx2 = web3.utils.encodePacked(
        {"type":"address", value:club.address},
        {"type":"address", value:accounts[0]},
        {"type":"uint256", value:1},
        {"type":"uint256", value:1980}
      );
  
      const tx3 = web3.utils.keccak256(tx2);
      const sign1 = await web3.eth.accounts.sign(tx3, key.key)
      await club.pledge(1, 1980, "1980", sign1.signature, {from:accounts[0], value: 0});

      const total = await club.countClubLand(1);
      return assert.equal(total, 1, "failed to pledge")
  });

  it("should unpledge", async function(){
    const club = await Club.deployed();
    const tx2 = web3.utils.encodePacked(
        {"type":"address", value:club.address},
        {"type":"address", value:accounts[0]},
        {"type":"uint256", value:1},
        {"type":"uint256", value:1980},
      );
  
      const key = await common.makePrivateKey(process.env.MNEMONIC, process.env.HDPATH);
      const tx3 = web3.utils.keccak256(tx2);
      const sign1 = await web3.eth.accounts.sign(tx3, key.key)

      await club.unpledge(1,1980, sign1.signature);
      total = await club.countClubLand(1);
      return assert.equal(total, 0, "failed to unpledge");
  });
})