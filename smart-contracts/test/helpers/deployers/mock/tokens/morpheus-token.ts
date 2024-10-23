import { ethers } from 'hardhat';

import { MorpheusToken } from '@/generated-types/ethers';

export const deployMORToken = async (): Promise<MorpheusToken> => {
  const factory = await ethers.getContractFactory('MorpheusToken');
  const contract = await factory.deploy();

  return contract;
};
