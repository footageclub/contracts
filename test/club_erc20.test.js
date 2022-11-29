const ClubERC20 = artifacts.require("ClubERC20");

/*
 * uncomment accounts to access the test accounts made available by the
 * Ethereum client
 * See docs: https://www.trufflesuite.com/docs/truffle/testing/writing-tests-in-javascript
 */
contract("ClubERC20", function ( accounts ) {
  it("should assert true", async function () {
    await ClubERC20.deployed();
    return assert.isTrue(true);
  });

  it("should pause", async function() {
    const club = await ClubERC20.deployed();
    await club.pause();
    return assert.isTrue(true);
  });

  it("should unpause", async function() {
    const club = await ClubERC20.deployed();
    await club.unpause();
    return assert.isTrue(true);
  });

  it("should mint", async function() {
    const club = await ClubERC20.deployed();
    await club.mint(accounts[0], 1, {value:68000});
    return assert.isTrue(true);
  });

  it("should withdraw", async function() {
    const club = await ClubERC20.deployed();
    await club.withdraw();
    return assert.isTrue(true);
  });
})