import { Fragment } from 'ethers';
import { ethers } from 'hardhat';

import { IProviderRegistry__factory, LumerinDiamond, ProviderRegistry } from '@/generated-types/ethers';
import { FacetAction } from '@/test/helpers/deployers/diamond/lumerin-diamond';

export const deployFacetProviderRegistry = async (diamond: LumerinDiamond): Promise<ProviderRegistry> => {
  let facet: ProviderRegistry;

  const factory = await ethers.getContractFactory('ProviderRegistry');
  facet = await factory.deploy();

  await diamond['diamondCut((address,uint8,bytes4[])[])']([
    {
      facetAddress: facet,
      action: FacetAction.Add,
      functionSelectors: IProviderRegistry__factory.createInterface()
        .fragments.filter(Fragment.isFunction)
        .map((f) => f.selector),
    },
  ]);

  facet = facet.attach(diamond.target) as ProviderRegistry;
  await facet.__ProviderRegistry_init();

  return facet;
};
