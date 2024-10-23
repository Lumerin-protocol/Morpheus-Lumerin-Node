import { Fragment } from 'ethers';
import { ethers } from 'hardhat';

import { IModelRegistry__factory, LumerinDiamond, ModelRegistry } from '@/generated-types/ethers';
import { FacetAction } from '@/test/helpers/deployers/diamond/lumerin-diamond';

export const deployFacetModelRegistry = async (diamond: LumerinDiamond): Promise<ModelRegistry> => {
  let facet: ModelRegistry;

  const factory = await ethers.getContractFactory('ModelRegistry');
  facet = await factory.deploy();

  await diamond['diamondCut((address,uint8,bytes4[])[])']([
    {
      facetAddress: facet,
      action: FacetAction.Add,
      functionSelectors: IModelRegistry__factory.createInterface()
        .fragments.filter(Fragment.isFunction)
        .map((f) => f.selector),
    },
  ]);

  facet = facet.attach(diamond.target) as ModelRegistry;
  await facet.__ModelRegistry_init();

  return facet;
};
