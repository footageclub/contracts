const ClubNFTPublic = artifacts.require("ClubNFTPublic");
const common = require('./common');
/*
 * uncomment accounts to access the test accounts made available by the
 * Ethereum client
 * See docs: https://www.trufflesuite.com/docs/truffle/testing/writing-tests-in-javascript
 */
contract("ClubNFTPublic", function ( accounts ) {
  it("should assert true", async function () {
    await ClubNFTPublic.deployed();
    return assert.isTrue(true);
  });

  it("should pause", async function() {
    const club = await ClubNFTPublic.deployed();
    await club.pause();
    return assert.isTrue(true);
  });

  it("should unpause", async function() {
    const club = await ClubNFTPublic.deployed();
    await club.unpause();
    return assert.isTrue(true);
  });

  it("should setURI", async function() {
    const club = await ClubNFTPublic.deployed();
    await club.setURI("https://www.footage.club/v2/{ID}");
    return assert.isTrue(true);
  });

  it("should setRecipientAddress", async function() {
    const club = await ClubNFTPublic.deployed();
    await club.setRecipientAddress(1, accounts[0]);
    return assert.isTrue(true);
  });

  it("should setRRecipientAddress", async function() {
    const club = await ClubNFTPublic.deployed();
    await club.setRRecipientAddress(1, accounts[0]);
    return assert.isTrue(true);
  });

  it("should setRPercentage", async function() {
    const club = await ClubNFTPublic.deployed();
    await club.setRPercentage(1, 500);
    return assert.isTrue(true);
  });

  it("should setClubNFT", async function() {
    const club = await ClubNFTPublic.deployed();
    const tx0 = web3.utils.encodePacked(
      {"type":"address", value:club.address},
      {"type":"address", value:accounts[0]},
      {"type":"uint256", value:1},
      {"type":"uint[]", value:[1,2,3]},
      {"type":"uint[]", value:[100,100,100]},
      {"type":"uint[]", value:[5,5,6]},
      {"type":"uint[]", value:[1,5,7]}
    );

    const tx1 = web3.utils.keccak256(tx0);
    const key = await common.makePrivateKey(process.env.MNEMONIC, process.env.HDPATH);
    const sign = await web3.eth.accounts.sign(tx1, key.key);
    await club.setClubNFT(1, [1,2,3], [100,100,100],[5,5,6],[1,5,7],sign.signature);
    return assert.isTrue(true);
  });

  it("should mint", async function() {
    const club = await ClubNFTPublic.deployed();
    let tx = web3.utils.encodePacked(
      {"type":"address", value:club.address},
      {"type":"address", value:accounts[0]},
      {"type":"uint256", value:1},
      {"type":"uint256", value:1},
      {"type":"uint256", value:5}
    );

    let tx1 = web3.utils.keccak256(tx);
    let key = await common.makePrivateKey(process.env.MNEMONIC, process.env.HDPATH);
    let sign = await web3.eth.accounts.sign(tx1, key.key)
    await club.mint(1, 1, 5,  sign.signature,  {from:accounts[0], value: 5});
    // 获取余额
    let balanceEth = await web3.eth.getBalance(club.address)
    assert.equal(balanceEth, 5,  "failed balanceEth equal 680000000")

    const tx2 = web3.utils.encodePacked(
      {"type":"address", value:club.address},
      {"type":"address", value:accounts[0]},
      {"type":"uint256", value:1},
      {"type":"string", value:"withdraw"}
    );

    const tx4 = web3.utils.keccak256(tx2);
    const sign1 = await web3.eth.accounts.sign(tx4, key.key)

    // to
    await club.withdraw(1, sign1.signature);
    balanceEth = await web3.eth.getBalance(club.address)
    assert.equal(balanceEth, 0,  "failed balanceEth equal 0")

    tx = web3.utils.encodePacked(
      {"type":"address", value:club.address},
      {"type":"address", value:accounts[0]},
      {"type":"uint256", value:1},
      {"type":"uint256", value:2},
      {"type":"uint256", value:10}
    );

    tx1 = web3.utils.keccak256(tx);
    key = await common.makePrivateKey(process.env.MNEMONIC, process.env.HDPATH);
    sign = await web3.eth.accounts.sign(tx1, key.key)
    try {
       await club.mint(1, 2, 10,  sign.signature,  {from:accounts[0], value: 50});
    }catch(e) {
      assert.equal(e.message, "Returned error: VM Exception while processing transaction: revert peer total limit", "revert peer total limit")
    }

    const uri = await club.uri(1);
    return assert.isTrue(uri == "https://www.footage.club/v2/{ID}")
  });
})