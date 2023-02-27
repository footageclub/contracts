
import "dotenv/config";
import "@nomiclabs/hardhat-waffle";
import "@openzeppelin/hardhat-upgrades";
import "@nomiclabs/hardhat-etherscan";

module.exports = {
  solidity: {
    compilers: [
      {
        version: "0.8.17",
        settings: {
          viaIR: true,
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
    ]
  },
  networks: {
    goerli: {
      url: process.env.RPC_URL_GOERLI,
      accounts: [process.env.PRIVATE_KEY],
    },

    local: {
      url: process.env.RPC_URL_LOCAL,
      accounts: [process.env.PRIVATE_KEY_LOCAL],
      blockGasLimit: 30_000_000
    },

    hardhat: {
      blockGasLimit: 30_000_000,
      throwOnCallFailures: false,
    },
  },
  etherscan: {
    // Your API key for Etherscan
    // Obtain one at https://etherscan.io/
    apiKey: process.env.ETHERSCAN_KEY,
  },
  paths: { sources: "./src" },
};
