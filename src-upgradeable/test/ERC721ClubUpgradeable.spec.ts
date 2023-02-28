import { expect } from "chai";
import { ethers, network, upgrades } from "hardhat";
import { keccak256 } from "@ethersproject/keccak256";

import { randomHex } from "../../test/token/utils/encoding";
import { faucet } from "../../test/token/utils/faucet";
import { VERSION } from "../../test/token/utils/helpers";
import type { Wallet } from "ethers";

describe(`ERC721ClubUpgradeable(v${VERSION})`, function () {
  const { provider } = ethers;
  let token;
  let club;
  let owner: Wallet;
  let admin: Wallet;
  let creator: Wallet;
  let payer: Wallet;
  let minter: Wallet;
  let feeRecipient: Wallet;
  let eip712Domain: { [key: string]: string | number };
  let eip712Types: Record<string, Array<{ name: string; type: string }>>;
  let eip712TypesUnplege: Record<string, Array<{ name: string; type: string }>>;
  let salt: ethers.BigNumber;
  
  after(async () => {
    await network.provider.request({
      method: "hardhat_reset",
    });
  });

  before(async () => {
    // Set the wallets
    owner = new ethers.Wallet(randomHex(32), provider);
    admin = new ethers.Wallet(randomHex(32), provider);
    creator = new ethers.Wallet(randomHex(32), provider);
    payer = new ethers.Wallet(randomHex(32), provider);
    minter = new ethers.Wallet(randomHex(32), provider);
    feeRecipient = new ethers.Wallet(randomHex(32), provider);

    // Add eth to wallets
    for (const wallet of [owner, admin, payer, minter]) {
      await faucet(wallet.address, provider);
    }

    // Deploy SeaDrop
    const ERC721TokenUpgradeable = await ethers.getContractFactory("ERC721PartnerTokenUpgradeable", owner);
    token = await upgrades.deployProxy(
        ERC721TokenUpgradeable,
        [
          "Footage",
          "Super Bee",
          owner.address,
          [owner.address]
        ],
        { initializer: "initialize" }
      );
      await token.deployed();
      await token.setMaxSupply(10);
  });

  beforeEach(async () => {
    const ERC721ClubUpgradeable = await ethers.getContractFactory("ERC721ClubUpgradeable", owner);
    club = await upgrades.deployProxy(
        ERC721ClubUpgradeable,
        [
            "Footage",
            "Club",
            admin.address,
            ethers.BigNumber.from("524288000").toString(),
            ethers.BigNumber.from("300000").toString()
        ],
        { initializer: "initialize" }
      );
      await club.deployed();

       // Configure EIP-712 params
    eip712Domain = {
        name: "Footage",
        version: "1.0",
        chainId: (await provider.getNetwork()).chainId,
        verifyingContract: club.address
      };

    eip712Types = {
        Plege: [
          { name: "nftContract", type: "address" },
          { name: "sender", type: "address" },
          { name: "clubId", type: "uint256" },
          { name: "tokenId", type: "uint256" },
          { name: "plegedToken", type: "uint256[]" },
          { name: "salt", type: "uint256" }
        ]
      };

      eip712TypesUnplege = {
        Unplege: [
          { name: "nftContract", type: "address" },
          { name: "sender", type: "address" },
          { name: "clubId", type: "uint256" },
          { name: "tokenId", type: "uint256" },
          { name: "plegedToken", type: "uint256[]" },
          { name: "salt", type: "uint256" }
        ]
      };

      await club.updateAllowedPledgeToken(token.address, true);
  });

  const signPlege = async (
    nftContract: string,
    sender: Wallet,
    clubId:ethers.BigNumber,
    tokenId:ethers.BigNumber,
    plegedToken:ethers.BigNumber[],
    salt:ethers.BigNumber,
    signer: Wallet
  ) => {
    const plege = {
      nftContract:nftContract,
      sender: sender.address,
      clubId: clubId,
      tokenId: tokenId,
      plegedToken:plegedToken,
      salt
    };

    const signature = await signer._signTypedData(
      eip712Domain,
      eip712Types,
      plege
    );

    // Verify recovered address matchers signer address
    const verifiedAddress = ethers.utils.verifyTypedData(
      eip712Domain,
      eip712Types,
      plege,
      signature
    );

    expect(verifiedAddress).to.eq(signer.address);
    return signature;
  };

  const signUnplege = async (
    nftContract: string,
    sender: Wallet,
    clubId:ethers.BigNumber,
    tokenId:ethers.BigNumber,
    plegedToken:ethers.BigNumber[],
    salt:ethers.BigNumber,
    signer: Wallet
  ) => {
    const unplege = {
      nftContract:nftContract,
      sender: sender.address,
      clubId:clubId,
      tokenId: tokenId,
      plegedToken:plegedToken,
      salt
    };

    const signature = await signer._signTypedData(
      eip712Domain,
      eip712TypesUnplege,
      unplege
    );

    // Verify recovered address matchers signer address
    const verifiedAddress = ethers.utils.verifyTypedData(
      eip712Domain,
      eip712TypesUnplege,
      unplege,
      signature
    );

    expect(verifiedAddress).to.eq(signer.address);
    return signature;
  };


  it("Should plege/unplege token to club", async () => {
    await token.mintSeaDrop(owner.address, 1);
    await token.setApprovalForAll(club.address, true);
    expect(await token.ownerOf(1)).to.be.eq(owner.address);

    

   const salt = ethers.BigNumber.from(randomHex(32));
   const pk:ethers.BigNumber[] = [1];
   const id:ethers.BigNumber = 101;
   const tokenId :ethers.BigNumber = 0;

   //plege(uint256 clubId, uint256 tokenId, address nftContract, uint256[] calldata plegeToken, uint256 salt, bytes calldata signature)
   const sign = await signPlege(token.address, owner, id, tokenId, pk, salt, admin)
   await expect(club.plege(id ,tokenId,token.address, pk,salt, sign)).to.emit(club, "PlegedUpdated");
   expect(await token.ownerOf(1)).to.be.eq(club.address);
   expect(await club.ownerOf(1)).to.be.eq(owner.address);

   const sign3 = await signUnplege(token.address, owner,id,1, pk, salt, admin)
   await expect(club.connect(owner).unplege(id ,1,token.address, pk,salt, sign3)).to.be.revertedWithCustomError(club, "OnlyPleged");

   //expect(await token.ownerOf(1)).to.be.eq(owner.address);
  })
});
