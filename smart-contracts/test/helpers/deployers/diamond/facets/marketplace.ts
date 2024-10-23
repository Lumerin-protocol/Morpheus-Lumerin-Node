import { Fragment } from 'ethers';
import { ethers } from 'hardhat';

import {
  IBidStorage__factory,
  IMarketplace__factory,
  LumerinDiamond,
  Marketplace,
  MorpheusToken,
} from '@/generated-types/ethers';
import { FacetAction } from '@/test/helpers/deployers/diamond/lumerin-diamond';

export const deployFacetMarketplace = async (diamond: LumerinDiamond, token: MorpheusToken): Promise<Marketplace> => {
  let facet: Marketplace;

  const factory = await ethers.getContractFactory('Marketplace');
  facet = await factory.deploy();

  await diamond['diamondCut((address,uint8,bytes4[])[])']([
    {
      facetAddress: facet,
      action: FacetAction.Add,
      functionSelectors: IMarketplace__factory.createInterface()
        .fragments.filter(Fragment.isFunction)
        .map((f) => f.selector),
    },
    {
      facetAddress: facet,
      action: FacetAction.Add,
      functionSelectors: IBidStorage__factory.createInterface()
        .fragments.filter(Fragment.isFunction)
        .map((f) => f.selector),
    },
  ]);

  facet = facet.attach(diamond.target) as Marketplace;
  await facet.__Marketplace_init(token);

  return facet;
};
