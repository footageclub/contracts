import { expect } from "chai";
import { ethers, network, upgrades} from "hardhat";

import { randomHex } from "../../test/token/utils/encoding";
import { faucet } from "../../test/token/utils/faucet";
import { VERSION } from "../../test/token/utils/helpers";
import type { Wallet } from "ethers";
import type { ERC721PartnerSeaDrop, ISeaDrop } from "../../typechain-types";
import { keccak256 } from "@ethersproject/keccak256";

describe(`ERC721ClubLockUpgradeable(v${VERSION})`, function () {
    const { provider } = ethers;
    
    let clublock;
    let owner:Wallet;
    let signer:Wallet;
    let minter:Wallet;
    let player:Wallet;
    let eip712Domain: { [key: string]: string | number };
    let eip712Types: Record<string, Array<{ name: string; type: string }>>;
    let salt: ethers.BigNumber;

    const mintParams = {
        mintPrice: "100000000000000000", // 0.1 ether
        maxTotalMintableByWallet: 10,
        startTime: Math.round(Date.now() / 1000) - 100,
        endTime: Math.round(Date.now() / 1000) + 100,
        dropStageIndex:1,
        feeBps: 1000,
        maxTokenSupplyForStage:5,
        restrictFeeRecipients: true
    }

    after(async () => {
        await network.provider.request({
          method: "hardhat_reset",
        });
    });

    before(async () => {
        owner = new ethers.Wallet(randomHex(32), provider);
        signer = new ethers.Wallet(randomHex(32), provider);
        minter = new ethers.Wallet(randomHex(32), provider);
        player = new ethers.Wallet(randomHex(32), provider);

        

        // Add eth to wallets
        for (const wallet of [owner, signer, player, minter]) {
            await faucet(wallet.address, provider);
        }

        // Deploy SeaDrop
        const ERC721ClubLockUpgradeable = await ethers.getContractFactory("ERC721ClubLockUpgradeable", owner);
        clublock = await upgrades.deployProxy(
            ERC721ClubLockUpgradeable,
            [
            "Footage",
            "Clublock",
            signer.address
            ],
            { initializer: "initialize" }
        );
        await clublock.deployed();
    });

    beforeEach(async () => {
        await clublock.setMaxSupply(10);
        await clublock.updateMintParams(mintParams);
        await clublock.updateCreatorPayoutAddress(owner.address);

        eip712Domain = {
            name: "Footage",
            version: "1.0",
            chainId: (await provider.getNetwork()).chainId,
            verifyingContract: clublock.address
          };
    
        eip712Types = {
            MintClublock: [
              { name: "minter", type: "address" },
              { name: "to", type: "address" },
              { name: "assets", type: "bytes32" },
              { name: "salt", type: "uint256" }
            ]
          };
    });

    it("ERC721ClubLockUpgradeable set config", async () => {
        await expect(clublock.updateMintParams(mintParams)).to.be.emit(clublock, "MintParamsUpdated");
        await expect(clublock.connect(signer).config( ethers.constants.AddressZero,10, mintParams,
            {
                royaltyAddress:minter.address,
                royaltyBps:1000
            }
        )).to.be.revertedWithCustomError(clublock, "OnlyOwner");


        await expect(clublock.connect(owner).config(ethers.constants.AddressZero, 10, mintParams,
            {
                royaltyAddress:minter.address,
                royaltyBps:1000
            }
        )).to.be.emit(clublock, "MintParamsUpdated");
    })



    const signMintClubLock = async (
        minter:Wallet,
        to: Wallet,
        assets:string[],
        salt:ethers.BigNumber,
        signer: Wallet
      ) => {
        const mintClubLock  = {
          minter:minter.address,
          to:to.address,
          assets: keccak256(ethers.utils.toUtf8Bytes(assets.join(''))) ,
          salt
        };
        
        const signature = await signer._signTypedData(
          eip712Domain,
          eip712Types,
          mintClubLock
        );
    
        // Verify recovered address matchers signer address
        const verifiedAddress = ethers.utils.verifyTypedData(
          eip712Domain,
          eip712Types,
          mintClubLock,
          signature
        );
    
        expect(verifiedAddress).to.eq(signer.address);
        return signature;
      };

    it("ERC721ClubLockUpgradeable mint to clublock", async () => {
        await expect(clublock.updateMintParams(mintParams)).to.be.emit(clublock, "MintParamsUpdated");

        salt = ethers.BigNumber.from(randomHex(32))
        const to = player;
        const assets:string[] = ["1", "A"];
        const value = ethers.BigNumber.from(mintParams.mintPrice).mul(assets.length);

        const sign = await signMintClubLock(to,to, assets, salt, signer)
        await expect(clublock.connect(minter).mintClublock(
            to.address, 
            assets,
            salt,sign,{
                value
            }
        )).to.be.emit(clublock, "mintClublockUpdated");
        
    })
})
