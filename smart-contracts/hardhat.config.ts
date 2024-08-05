import "dotenv/config";
import "@nomicfoundation/hardhat-verify";
import "@nomicfoundation/hardhat-toolbox-viem";
import "@nomicfoundation/hardhat-ignition-viem";
import "@solarity/hardhat-gobind";
import "./tasks/upgrade";
import { HardhatUserConfig } from "hardhat/config";

const config: HardhatUserConfig = {
  networks: {
    hardhat: {
      initialDate: "2024-07-16T01:00:00.000Z",
      gas: "auto", // required for tests where two transactions should be mined in the same block
      // loggingEnabled: true,
    },
  },
  solidity: {
    version: "0.8.24",
    settings: {
      optimizer: {
        enabled: true,
        runs: 1000,
        details: {
          yul: true,
          constantOptimizer: true,
        },
      },
      viaIR: true,
    },
  },
  mocha: {
    // reporter: "",
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS ? true : false,
    outputJSON: true,
    outputJSONFile: "gas.json",
    coinmarketcap: process.env.COINMARKETCAP_API_KEY,
    reportPureAndViewMethods: true,
    darkMode: true,
    currency: "USD",
    // offline: true,
    // L2Etherscan: process.env.ETHERSCAN_API_KEY,
    // L2: "arbitrum",
    L1Etherscan: process.env.ETHERSCAN_API_KEY,
    L1: "ethereum",
    reportPureAndViewMethods: true,
  },
  gobind: {
    outdir: "./bindings/go",
    onlyFiles: ["./contracts"],
    skipFiles: [
      "contracts/AppStorage.sol",
      "contracts/libraries",
      "contracts/diamond/libraries",
      "contracts/diamond/interfaces",
    ],
  },
};

export default config;
