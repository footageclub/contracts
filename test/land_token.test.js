const LandToken = artifacts.require("LandToken");
const common = require('./common');

/*
 * uncomment accounts to access the test accounts made available by the
 * Ethereum client
 * See docs: https://www.trufflesuite.com/docs/truffle/testing/writing-tests-in-javascript
 */
contract("LandToken", function ( accounts ) {
  it("should assert true", async function () {
    await LandToken.deployed();
    return assert.isTrue(true);
  });

  it("should setBaseURI", async function() {
    const landToken = await LandToken.deployed();
    await landToken.setBaseURI("https://www.footage.club/v2/");
    return assert.isTrue(true);
  });

  it("should pause", async function() {
    const landToken = await LandToken.deployed();
    await landToken.pause();
    return assert.isTrue(true);
  });

  it("should unpause", async function() {
    const landToken = await LandToken.deployed();
    await landToken.unpause();
    return assert.isTrue(true);
  });

  it("should getFeeRecipients", async function() {
    const landToken = await LandToken.deployed();
    const address = await landToken.getFeeRecipients(1);
    return assert.equal(address, accounts[0], "invaild address for getFeeRecipients")
  });

  it("should getFeeBps", async function() {
    const landToken = await LandToken.deployed();
    const percebtage = await landToken.getFeeBps(1);
    return assert.equal(percebtage, 250, "invaild percebtage for getFeeBps")
  });

  it("should setFeeRecipient", async function() {
    const landToken = await LandToken.deployed();
    await landToken.setFeeRecipient(accounts[1], 260);
    const address = await landToken.getFeeRecipients(1);
    assert.equal(address, accounts[1], "invaild address for getFeeRecipients")
    const percebtage = await landToken.getFeeBps(1);
    return assert.equal(percebtage, 260, "invaild percebtage for getFeeBps")
  });


  //  function mint(uint amount, uint256 tokenId, string memory tokenUrl, bytes calldata sign) public payable {
  it("should mint", async function() {
    const landToken = await LandToken.deployed();
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
   
    // 获取余额
    let balanceEth = await web3.eth.getBalance(landToken.address)
    assert.equal(balanceEth, 680000000,  "failed balanceEth equal 680000000")

    const balance = await landToken.balanceByOwner(accounts[0]);
    assert.equal(balance.length, 1,  "failed balance equal 1")

    // to
    await landToken.withdraw(accounts[1], {from:accounts[0]});
    balanceEth = await web3.eth.getBalance(landToken.address)
    assert.equal(balanceEth, 0,  "failed balanceEth equal 0")
    const token = await landToken.tokenURI(1980);
    return assert.isTrue(token == "https://www.footage.club/v2/1980")
  });
});
