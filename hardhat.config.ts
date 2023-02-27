import fs from "fs"
import { HardhatUserConfig } from "hardhat/config";

import "dotenv/config";
import "@nomicfoundation/hardhat-toolbox";
import "hardhat-preprocessor"

// Configure remappings.
// https://book.getfoundry.sh/config/hardhat
// Re-run `forge remappings > remappings.txt`
// every time you modify libraries in Foundry.
function getRemappings() {
  return fs
    .readFileSync("remappings.txt", "utf8")
    .split("\n")
    .filter(Boolean) // remove empty lines
    .map((line: string) => line.trim().split("="));
}

const config: HardhatUserConfig = {
  solidity: {
    compilers: [
      {
        version: "0.8.17",
        settings: {
          viaIR: false,
          optimizer: {
            enabled: true,
            runs: 1_000_000,
          },
        },
      },
    ],
  },
  networks: {
    goerli: {
      url: process.env.RPC_URL_GOERLI,
      accounts: [process.env.PRIVATE_KEY ?? ""],
    },
    hardhat: {
      blockGasLimit: 30_000_000,
      throwOnCallFailures: false,
    },
    
    verificationNetwork: {
      url: process.env.NETWORK_RPC ?? "",
    },
  },
  preprocess: {
    eachLine: (hre) => ({
      transform: (line: string) => {
        if (line.match(/ from "/i)) {
          getRemappings().forEach(([find, replace]: string[]) => {
            if (line.match(find)) {
              line = line.replace(find, replace);
            }
          });
        }
        return line;
      },
    }),
  },
};

export default config;
