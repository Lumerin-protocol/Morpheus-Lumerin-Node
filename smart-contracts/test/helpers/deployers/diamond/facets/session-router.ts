import { SignerWithAddress } from '@nomicfoundation/hardhat-ethers/signers';
import { Fragment } from 'ethers';
import { ethers } from 'hardhat';

import {
  ISessionRouter__factory,
  IStatsStorage__factory,
  LumerinDiamond,
  SessionRouter,
} from '@/generated-types/ethers';
import { FacetAction } from '@/test/helpers/deployers/diamond/lumerin-diamond';
import { getDefaultPools } from '@/test/helpers/pool-helper';

export const deployFacetSessionRouter = async (
  diamond: LumerinDiamond,
  fundingAccount: SignerWithAddress,
): Promise<SessionRouter> => {
  let facet: SessionRouter;

  const LDIDFactory = await ethers.getContractFactory('LinearDistributionIntervalDecrease');
  const LDID = await LDIDFactory.deploy();

  const factory = await ethers.getContractFactory('SessionRouter', {
    libraries: {
      LinearDistributionIntervalDecrease: LDID,
    },
  });
  facet = await factory.deploy();

  await diamond['diamondCut((address,uint8,bytes4[])[])']([
    {
      facetAddress: facet,
      action: FacetAction.Add,
      functionSelectors: ISessionRouter__factory.createInterface()
        .fragments.filter(Fragment.isFunction)
        .map((f) => f.selector),
    },
    {
      facetAddress: facet,
      action: FacetAction.Add,
      functionSelectors: IStatsStorage__factory.createInterface()
        .fragments.filter(Fragment.isFunction)
        .map((f) => f.selector),
    },
  ]);

  facet = facet.attach(diamond.target) as SessionRouter;
  await facet.__SessionRouter_init(fundingAccount, getDefaultPools());

  return facet;
};
