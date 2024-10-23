import { ethers } from 'hardhat';

import { LumerinDiamond } from '@/generated-types/ethers';

export enum FacetAction {
  Add = 0,
  Replace = 1,
  Remove = 2,
}

export const deployLumerinDiamond = async (): Promise<LumerinDiamond> => {
  const factory = await ethers.getContractFactory('LumerinDiamond');
  const contract = await factory.deploy();
  await contract.__LumerinDiamond_init();

  return contract;
};
