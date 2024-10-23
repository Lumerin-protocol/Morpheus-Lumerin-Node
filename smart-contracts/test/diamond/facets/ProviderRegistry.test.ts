import { LumerinDiamond, Marketplace, ModelRegistry, MorpheusToken, ProviderRegistry } from '@ethers-v6';
import { SignerWithAddress } from '@nomicfoundation/hardhat-ethers/signers';
import { expect } from 'chai';
import { ethers } from 'hardhat';

import { getHex, wei } from '@/scripts/utils/utils';
import {
  deployFacetMarketplace,
  deployFacetModelRegistry,
  deployFacetProviderRegistry,
  deployFacetSessionRouter,
  deployLumerinDiamond,
  deployMORToken,
} from '@/test/helpers/deployers';
import { Reverter } from '@/test/helpers/reverter';
import { setNextTime } from '@/utils/block-helper';
import { YEAR } from '@/utils/time';

describe('ProviderRegistry', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let PROVIDER: SignerWithAddress;

  let diamond: LumerinDiamond;
  let providerRegistry: ProviderRegistry;
  let modelRegistry: ModelRegistry;
  let marketplace: Marketplace;

  let token: MorpheusToken;

  const modelId = getHex(Buffer.from('1'));
  const ipfsCID = getHex(Buffer.from('ipfs://ipfsaddress'));

  before(async () => {
    [OWNER, PROVIDER] = await ethers.getSigners();

    [diamond, token] = await Promise.all([deployLumerinDiamond(), deployMORToken()]);

    [providerRegistry, modelRegistry, , marketplace] = await Promise.all([
      deployFacetProviderRegistry(diamond),
      deployFacetModelRegistry(diamond),
      deployFacetSessionRouter(diamond, OWNER),
      deployFacetMarketplace(diamond, token),
    ]);

    await token.transfer(PROVIDER, wei(1000));
    await token.connect(PROVIDER).approve(providerRegistry, wei(1000));
    await token.approve(providerRegistry, wei(1000));

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('#__ProviderRegistry_init', () => {
    it('should revert if try to call init function twice', async () => {
      await expect(providerRegistry.__ProviderRegistry_init()).to.be.rejectedWith(
        'Initializable: contract is already initialized',
      );
    });
  });

  describe('#providerSetMinStake', async () => {
    it('should set min stake', async () => {
      const minStake = wei(100);

      await expect(providerRegistry.providerSetMinStake(minStake))
        .to.emit(providerRegistry, 'ProviderMinimumStakeUpdated')
        .withArgs(minStake);

      expect(await providerRegistry.getProviderMinimumStake()).eq(minStake);
    });

    it('should throw error when caller is not an owner', async () => {
      await expect(providerRegistry.connect(PROVIDER).providerSetMinStake(100)).to.be.revertedWithCustomError(
        diamond,
        'OwnableUnauthorizedAccount',
      );
    });
  });

  describe('#providerRegister', async () => {
    it('should register a new provider', async () => {
      await setNextTime(300);
      await providerRegistry.connect(PROVIDER).providerRegister(wei(100), 'test');

      const data = await providerRegistry.getProvider(PROVIDER);

      expect(data.endpoint).to.eq('test');
      expect(data.stake).to.eq(wei(100));
      expect(data.createdAt).to.eq(300);
      expect(data.limitPeriodEnd).to.eq(YEAR + 300);
      expect(data.limitPeriodEarned).to.eq(0);
      expect(data.isDeleted).to.eq(false);
      expect(await providerRegistry.getIsProviderActive(PROVIDER)).to.eq(true);

      expect(await token.balanceOf(providerRegistry)).to.eq(wei(100));
      expect(await token.balanceOf(PROVIDER)).to.eq(wei(900));

      expect(await providerRegistry.getActiveProviders(0, 10)).to.deep.eq([PROVIDER.address]);

      await providerRegistry.connect(PROVIDER).providerRegister(wei(0), 'test');
    });
    it('should add stake to existed provider', async () => {
      await setNextTime(300);
      await providerRegistry.connect(PROVIDER).providerRegister(wei(100), 'test');
      await providerRegistry.connect(PROVIDER).providerRegister(wei(300), 'test2');

      const data = await providerRegistry.getProvider(PROVIDER);

      expect(data.endpoint).to.eq('test2');
      expect(data.stake).to.eq(wei(400));
      expect(data.createdAt).to.eq(300);
      expect(data.limitPeriodEnd).to.eq(YEAR + 300);
      expect(data.limitPeriodEarned).to.eq(0);
      expect(data.isDeleted).to.eq(false);
      expect(await providerRegistry.getIsProviderActive(PROVIDER)).to.eq(true);

      expect(await token.balanceOf(providerRegistry)).to.eq(wei(400));
      expect(await token.balanceOf(PROVIDER)).to.eq(wei(600));
    });
    it('should activate deregistered provider', async () => {
      await setNextTime(300);
      await providerRegistry.connect(PROVIDER).providerRegister(wei(100), 'test');
      await setNextTime(301 + YEAR);
      await providerRegistry.connect(PROVIDER).providerDeregister();

      let data = await providerRegistry.getProvider(PROVIDER);
      expect(data.isDeleted).to.eq(true);
      expect(await providerRegistry.getIsProviderActive(PROVIDER)).to.eq(false);

      await providerRegistry.connect(PROVIDER).providerRegister(wei(1), 'test2');
      data = await providerRegistry.getProvider(PROVIDER);

      expect(data.endpoint).to.eq('test2');
      expect(data.stake).to.eq(wei(1));
      expect(data.createdAt).to.eq(300);
      expect(data.limitPeriodEnd).to.eq(YEAR + 300);
      expect(data.limitPeriodEarned).to.eq(0);
      expect(data.isDeleted).to.eq(false);
      expect(await providerRegistry.getIsProviderActive(PROVIDER)).to.eq(true);
    });
    it('should throw error when the stake is too low', async () => {
      await providerRegistry.providerSetMinStake(wei(2));
      await expect(providerRegistry.connect(PROVIDER).providerRegister(wei(0), '')).to.be.revertedWithCustomError(
        providerRegistry,
        'ProviderStakeTooLow',
      );
    });
  });

  describe('#providerDeregister', async () => {
    it('should deregister the provider', async () => {
      await setNextTime(300);
      await providerRegistry.connect(PROVIDER).providerRegister(wei(100), 'test');
      await setNextTime(301 + YEAR);
      await providerRegistry.connect(PROVIDER).providerDeregister();

      expect((await providerRegistry.getProvider(PROVIDER)).isDeleted).to.equal(true);
      expect(await providerRegistry.getIsProviderActive(PROVIDER)).to.eq(false);
      expect(await token.balanceOf(providerRegistry)).to.eq(0);
      expect(await token.balanceOf(PROVIDER)).to.eq(wei(1000));

      expect(await providerRegistry.getActiveProviders(0, 10)).to.deep.eq([]);
    });
    it('should deregister the provider without transfer', async () => {
      await providerRegistry.providerSetMinStake(0);
      await setNextTime(300);
      await providerRegistry.connect(PROVIDER).providerRegister(wei(0), 'test');
      await setNextTime(301 + YEAR);
      await providerRegistry.connect(PROVIDER).providerDeregister();

      expect((await providerRegistry.getProvider(PROVIDER)).isDeleted).to.equal(true);
      expect(await providerRegistry.getIsProviderActive(PROVIDER)).to.eq(false);
      expect(await token.balanceOf(providerRegistry)).to.eq(0);
      expect(await token.balanceOf(PROVIDER)).to.eq(wei(1000));
    });
    it('should throw error when provider is not found', async () => {
      await expect(providerRegistry.connect(OWNER).providerDeregister()).to.be.revertedWithCustomError(
        providerRegistry,
        'ProviderNotFound',
      );
    });
    it('should throw error when provider has active bids', async () => {
      await providerRegistry.connect(PROVIDER).providerRegister(wei(100), 'test');
      await modelRegistry.connect(PROVIDER).modelRegister(modelId, ipfsCID, 0, wei(100), 'name', ['tag_1']);
      await marketplace.connect(PROVIDER).postModelBid(modelId, wei(10));
      await expect(providerRegistry.connect(PROVIDER).providerDeregister()).to.be.revertedWithCustomError(
        providerRegistry,
        'ProviderHasActiveBids',
      );
    });
    it('should throw error when delete provider few times', async () => {
      await providerRegistry.connect(OWNER).providerRegister(wei(100), 'test');
      await providerRegistry.connect(OWNER).providerDeregister();
      await expect(providerRegistry.connect(OWNER).providerDeregister()).to.be.revertedWithCustomError(
        providerRegistry,
        'ProviderHasAlreadyDeregistered',
      );
    });
  });
});

// npm run generate-types && npx hardhat test "test/diamond/facets/ProviderRegistry.test.ts"
// npx hardhat coverage --solcoverjs ./.solcover.ts --testfiles "test/diamond/facets/ProviderRegistry.test.ts"
