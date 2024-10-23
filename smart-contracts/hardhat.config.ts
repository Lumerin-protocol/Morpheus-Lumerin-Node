import '@nomicfoundation/hardhat-chai-matchers';
import '@nomicfoundation/hardhat-ethers';
import '@solarity/hardhat-gobind';
import '@solarity/hardhat-markup';
import '@solarity/hardhat-migrate';
import '@typechain/hardhat';
import * as dotenv from 'dotenv';
import 'hardhat-gas-reporter';
import { HardhatUserConfig } from 'hardhat/types';
import 'solidity-coverage';
import 'tsconfig-paths/register';

dotenv.config();

function privateKey() {
  return process.env.PRIVATE_KEY !== undefined ? [process.env.PRIVATE_KEY] : [];
}

function typechainTarget() {
  const target = process.env.TYPECHAIN_TARGET;

  return target === '' || target === undefined ? 'ethers-v6' : target;
}

function forceTypechain() {
  return process.env.TYPECHAIN_FORCE === 'false';
}

const config: HardhatUserConfig = {
  networks: {
    hardhat: {
      initialDate: '1970-01-01T00:00:00Z',
      gas: 'auto', // required for tests where two transactions should be mined in the same block
      // loggingEnabled: true,
      // mining: {
      //   auto: true,
      //   interval: 10_000,
      // },
      // forking: {
      //   url: `https://arbitrum-sepolia.infura.io/v3/${process.env.INFURA_KEY}`,
      // },
    },
    localhost: {
      url: 'http://127.0.0.1:8545',
      initialDate: '1970-01-01T00:00:00Z',
      gasMultiplier: 1.2,
      timeout: 1000000000000000,
    },
    arbitrum_sepolia: {
      url: `https://arbitrum-sepolia.infura.io/v3/${process.env.INFURA_KEY}`,
      accounts: privateKey(),
      gasMultiplier: 1.1,
    },
  },
  solidity: {
    version: '0.8.24',
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
      evmVersion: 'paris',
    },
  },
  gobind: {
    outdir: './bindings/go',
    onlyFiles: ['./contracts'],
    skipFiles: [
      'contracts/AppStorage.sol',
      'contracts/libraries',
      'contracts/diamond/libraries',
      'contracts/diamond/interfaces',
    ],
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS ? true : false,
    outputJSON: true,
    outputJSONFile: 'gas.json',
    coinmarketcap: process.env.COINMARKETCAP_API_KEY,
    reportPureAndViewMethods: true,
    darkMode: true,
    currency: 'USD',
    L1Etherscan: process.env.ETHERSCAN_API_KEY,
    L1: 'ethereum',
  },
  typechain: {
    outDir: `generated-types/${typechainTarget().split('-')[0]}`,
    target: typechainTarget(),
    alwaysGenerateOverloads: true,
    discriminateTypes: true,
    dontOverrideCompile: forceTypechain(),
  },
  etherscan: {
    apiKey: {
      mainnet: `${process.env.ETHERSCAN_API_KEY}`,
      arbitrumSepolia: `${process.env.ARBITRUM_API_KEY}`,
    },
  },
};

export default config;
