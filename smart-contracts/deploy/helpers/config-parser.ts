import { readFileSync } from 'fs';

import { ISessionStorage } from '@/generated-types/ethers';

export type Config = {
  MOR: string;
  fundingAccount: string;
  pools: ISessionStorage.PoolStruct[];
  providerMinStake: string;
  modelMinStake: string;
  marketplaceBidFee: string;
};

export function parseConfig(): Config {
  const configPath = `deploy/data/config_arbitrum_sepolia.json`;

  return JSON.parse(readFileSync(configPath, 'utf-8')) as Config;
}
