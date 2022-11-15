const LandToken = artifacts.require("LandToken");

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

  //  function mint(uint amount, uint256 tokenId, string memory tokenUrl, bytes calldata sign) public payable {

  it("should mint", async function() {
    const landToken = await LandToken.deployed();
    console.log(accounts);
    const tx0 = web3.utils.encodePacked(
      {"type":"address", value:landToken.address},
      {"type":"address", value:accounts[0]},
      {"type":"uint", value:680000000},
      {"type":"uint256", value:1980},
      {"type":"string", value:"test/1980"}
    );

    const tx1 = web3.utils.keccak256(tx0);
    const sign = await web3.eth.accounts.sign(tx1, "c9e49601c35ec25013aa37548d9ac73f9ac794d80e33f784801a0678a191948a")
    await landToken.mint(680000000,1980, "test/1980", sign.signature, {from:accounts[0], value: 680000000});
  

    // 获取余额
    const balanceEth = await landToken.balance({from:"0x5d437a732681eE9c649A6e1c1C44cd37C92F65E0"});
    console.log(balanceEth);

    // to
    await landToken.withdraw("0x3273131c4B86Efc85DA0c75bea8A0D6F239cFf72", {from:"0x5d437a732681eE9c649A6e1c1C44cd37C92F65E0"});
    const balanceEth2 = await landToken.balance({from:"0x5d437a732681eE9c649A6e1c1C44cd37C92F65E0"});
    console.log(balanceEth2);
    
    const balance = await landToken.balanceByOwner(accounts[0]);
    console.log(balance);

    const token = await landToken.tokenURI(1980);
    console.log(token);
    return assert.isTrue(token == "http://3.27.11.54/test/1980")
  });
});
