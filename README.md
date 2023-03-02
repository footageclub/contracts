# Footage

# 代码结构
```bash
｜ - lib # 第三方库
｜ - lib-upgradeable # 第三方库可升级合约库
｜ - scripts # 部署代码
｜ - src-upgradeable # 可升级合约源代码
｜ - src # 主合约代码
｜ ｜ - market # 二级市场合约
｜ ｜ - token # 掉落合约以及不可升级ERC721A合约
｜ - test # 测试代码
｜ ｜ - market # 二级市场合约测试代码
｜ ｜ - token # 掉落合约以及不可升级ERC721A合约测试代码
｜- tools # 工具代码，主要用于生成服务端代码
- .env.sample #配置
- hardhat.config.ts # 配置文件
- package.json
- remappings.txt
- tsconfig.json
```

# 合约
合约主要分三部分:

1.  二级市场交易合约
 
    二级市场交易合约主要通过SeaPort合约来实现，当前采用1.4版本。 合约支持基本订单、高级订单以及匹配订单履行，同样支持荷兰拍卖模式。我们根据实际需求对代码做了轻微调整，让他更能与Footage项目功能进行匹配

    调整后合约地址: https://github.com/footageclub/eth_contracts/blob/main/src/market/Seaport.sol


2. NFT统一下掉落合约

   掉落合约采用了SeaDrop合约来实现，当版本使用1.0版本. 为了匹配Footage项目的功能，我们采用统一掉落合约，将NFT铸造过成标准化，并对OpenSea有好。合约中我们主要用到统一公开铸造、白名单铸造、签名铸造

   调整后合约地址：https://github.com/footageclub/eth_contracts/blob/main/src/token/SeaDrop.sol

3. NFT合约

   NFT合约按照ERC721A标准进行功能实现, 支持批量铸造、版费、白名单以及合约升级等，合约主代码采用SeaDrop标准合约修改而来，安全可靠。以下是详细合约地址：

   Club合约,增加了质押与质押功能： https://github.com/footageclub/eth_contracts/blob/main/src-upgradeable/src/footage/ERC721ClubUpgradeable.sol

   Clublock合约： https://github.com/footageclub/eth_contracts/blob/main/src-upgradeable/src/footage/ERC721ClubLockUpgradeable.sol

   SuperBee合约: https://github.com/footageclub/eth_contracts/blob/main/src-upgradeable/src/footage/ERC721TokenUpgradeable.sol

   Campaign合约：https://github.com/footageclub/eth_contracts/blob/main/src/token/ERC721PartnerSeaDrop.sol

   以上所有合约均支持在SeaPort进行交易，以及SeaDrop中进行统一铸造



# 安装
```bash
git clone --recurse-submodules https://github.com/footageclub/eth_contracts.git && cd eth_contracts
yarn install
yarn build
```

# 测试合约 
```bash
yarn test
```

# 部署合约
```bash
npx hardhat run scripts/deploy.ts
```
