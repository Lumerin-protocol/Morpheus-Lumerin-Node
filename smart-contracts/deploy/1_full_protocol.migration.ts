import { Deployer, Reporter } from '@solarity/hardhat-migrate';
import { Fragment } from 'ethers';

import { parseConfig } from './helpers/config-parser';

import {
  IBidStorage__factory,
  IMarketplace__factory,
  IModelRegistry__factory,
  IProviderRegistry__factory,
  ISessionRouter__factory,
  IStatsStorage__factory,
  LinearDistributionIntervalDecrease__factory,
  LumerinDiamond__factory,
  Marketplace,
  Marketplace__factory,
  ModelRegistry,
  ModelRegistry__factory,
  ProviderRegistry,
  ProviderRegistry__factory,
  SessionRouter,
  SessionRouter__factory,
} from '@/generated-types/ethers';
import { FacetAction } from '@/test/helpers/deployers';

module.exports = async function (deployer: Deployer) {
  const config = parseConfig();

  const lumerinDiamond = await deployer.deploy(LumerinDiamond__factory);
  await lumerinDiamond.__LumerinDiamond_init();

  const ldid = await deployer.deploy(LinearDistributionIntervalDecrease__factory);

  let providerRegistryFacet = await deployer.deploy(ProviderRegistry__factory);
  let modelRegistryFacet = await deployer.deploy(ModelRegistry__factory);
  let marketplaceFacet = await deployer.deploy(Marketplace__factory);
  let sessionRouterFacet = await deployer.deploy(SessionRouter__factory, {
    libraries: {
      LinearDistributionIntervalDecrease: ldid,
    },
  });

  await lumerinDiamond['diamondCut((address,uint8,bytes4[])[])']([
    {
      facetAddress: providerRegistryFacet,
      action: FacetAction.Add,
      functionSelectors: IProviderRegistry__factory.createInterface()
        .fragments.filter(Fragment.isFunction)
        .map((f) => f.selector),
    },
    {
      facetAddress: modelRegistryFacet,
      action: FacetAction.Add,
      functionSelectors: IModelRegistry__factory.createInterface()
        .fragments.filter(Fragment.isFunction)
        .map((f) => f.selector),
    },
    {
      facetAddress: marketplaceFacet,
      action: FacetAction.Add,
      functionSelectors: IMarketplace__factory.createInterface()
        .fragments.filter(Fragment.isFunction)
        .map((f) => f.selector),
    },
    {
      facetAddress: marketplaceFacet,
      action: FacetAction.Add,
      functionSelectors: IBidStorage__factory.createInterface()
        .fragments.filter(Fragment.isFunction)
        .map((f) => f.selector),
    },
    {
      facetAddress: sessionRouterFacet,
      action: FacetAction.Add,
      functionSelectors: ISessionRouter__factory.createInterface()
        .fragments.filter(Fragment.isFunction)
        .map((f) => f.selector),
    },
    {
      facetAddress: sessionRouterFacet,
      action: FacetAction.Add,
      functionSelectors: IStatsStorage__factory.createInterface()
        .fragments.filter(Fragment.isFunction)
        .map((f) => f.selector),
    },
  ]);

  providerRegistryFacet = providerRegistryFacet.attach(lumerinDiamond.target) as ProviderRegistry;
  await providerRegistryFacet.__ProviderRegistry_init();
  modelRegistryFacet = modelRegistryFacet.attach(lumerinDiamond.target) as ModelRegistry;
  await modelRegistryFacet.__ModelRegistry_init();
  marketplaceFacet = marketplaceFacet.attach(lumerinDiamond.target) as Marketplace;
  await marketplaceFacet.__Marketplace_init(config.MOR);
  sessionRouterFacet = sessionRouterFacet.attach(lumerinDiamond.target) as SessionRouter;
  await sessionRouterFacet.__SessionRouter_init(config.fundingAccount, config.pools);

  await providerRegistryFacet.providerSetMinStake(config.providerMinStake);
  await modelRegistryFacet.modelSetMinStake(config.modelMinStake);
  await marketplaceFacet.setMarketplaceBidFee(config.marketplaceBidFee);

  Reporter.reportContracts(
    ['Lumerin Diamond', await lumerinDiamond.getAddress()],
    ['Linear Distribution Interval Decrease Library', await ldid.getAddress()],
  );
};

// npx hardhat migrate --only 1

// npx hardhat migrate --network arbitrum_sepolia --only 1 --verify
// npx hardhat migrate --network arbitrum_sepolia --only 1 --verify --continue

// npx hardhat migrate --network arbitrum --only 1 --verify
// npx hardhat migrate --network arbitrum --only 1 --verify --continue
