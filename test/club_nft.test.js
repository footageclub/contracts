const ClubNFT = artifacts.require("ClubNFT");

/*
 * uncomment accounts to access the test accounts made available by the
 * Ethereum client
 * See docs: https://www.trufflesuite.com/docs/truffle/testing/writing-tests-in-javascript
 */
contract("ClubNFT", function ( accounts ) {
  it("should assert true", async function () {
    await ClubNFT.deployed();
    return assert.isTrue(true);
  });

  it("should pause", async function() {
    const club = await ClubNFT.deployed();
    await club.pause();
    return assert.isTrue(true);
  });

  it("should unpause", async function() {
    const club = await ClubNFT.deployed();
    await club.unpause();
    return assert.isTrue(true);
  });

  it("should setURI", async function() {
    const club = await ClubNFT.deployed();
    await club.setURI("https://www.footage.club/v2/{ID}");
    return assert.isTrue(true);
  });

  it("should setRecipientAddress", async function() {
    const club = await ClubNFT.deployed();
    await club.setRecipientAddress(accounts[0]);
    return assert.isTrue(true);
  });

  it("should setRRecipientAddress", async function() {
    const club = await ClubNFT.deployed();
    await club.setRRecipientAddress(accounts[0]);
    return assert.isTrue(true);
  });

  it("should setRPercentage", async function() {
    const club = await ClubNFT.deployed();
    await club.setRPercentage(500);
    return assert.isTrue(true);
  });

  it("should setClubNFT", async function() {
    const club = await ClubNFT.deployed();
    await club.setClubNFT([1,2,3], [100,100,100],[5,5,6],[1,5,7]);
    return assert.isTrue(true);
  });

  it("should mint", async function() {
    const club = await ClubNFT.deployed();
    await club.mint(1, 5, [], {from:accounts[0], value: 5});
    // 获取余额
    let balanceEth = await web3.eth.getBalance(club.address)
    assert.equal(balanceEth, 5,  "failed balanceEth equal 5");

    // to
    await club.withdraw();
    balanceEth = await web3.eth.getBalance(club.address)
    assert.equal(balanceEth, 0,  "failed balanceEth equal 0");

    try {
       await club.mint(2, 10, [],  {from:accounts[0], value: 50});
    }catch(e) {
      assert.equal(e.message, "Returned error: VM Exception while processing transaction: revert peer total limit", "revert peer total limit")
    }

    const uri = await club.uri(1);
    return assert.isTrue(uri == "https://www.footage.club/v2/{ID}")
  });
})