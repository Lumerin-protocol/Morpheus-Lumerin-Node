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

describe('ModelRegistry', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let SECOND: SignerWithAddress;

  let diamond: LumerinDiamond;
  let providerRegistry: ProviderRegistry;
  let modelRegistry: ModelRegistry;
  let marketplace: Marketplace;

  let token: MorpheusToken;

  const modelId = getHex(Buffer.from('1'));
  const ipfsCID = getHex(Buffer.from('ipfs://ipfsaddress'));

  before(async () => {
    [OWNER, SECOND] = await ethers.getSigners();

    [diamond, token] = await Promise.all([deployLumerinDiamond(), deployMORToken()]);

    [providerRegistry, modelRegistry, , marketplace] = await Promise.all([
      deployFacetProviderRegistry(diamond),
      deployFacetModelRegistry(diamond),
      deployFacetSessionRouter(diamond, OWNER),
      deployFacetMarketplace(diamond, token),
    ]);

    await token.transfer(SECOND, wei(1000));
    await token.connect(SECOND).approve(providerRegistry, wei(1000));
    await token.approve(providerRegistry, wei(1000));
    await token.connect(SECOND).approve(modelRegistry, wei(1000));
    await token.approve(modelRegistry, wei(1000));
    await token.connect(SECOND).approve(marketplace, wei(1000));
    await token.approve(marketplace, wei(1000));

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('#__ModelRegistry_init', () => {
    it('should revert if try to call init function twice', async () => {
      await expect(modelRegistry.__ModelRegistry_init()).to.be.rejectedWith(
        'Initializable: contract is already initialized',
      );
    });
  });

  describe('#modelSetMinStake', async () => {
    it('should set min stake', async () => {
      const minStake = wei(100);

      await expect(modelRegistry.modelSetMinStake(minStake))
        .to.emit(modelRegistry, 'ModelMinimumStakeUpdated')
        .withArgs(minStake);

      expect(await modelRegistry.getModelMinimumStake()).eq(minStake);
    });

    it('should throw error when caller is not an owner', async () => {
      await expect(modelRegistry.connect(SECOND).modelSetMinStake(100)).to.be.revertedWithCustomError(
        diamond,
        'OwnableUnauthorizedAccount',
      );
    });
  });

  describe('#getModelIds', async () => {
    it('should set min stake', async () => {
      await setNextTime(300);
      await modelRegistry
        .connect(SECOND)
        .modelRegister(getHex(Buffer.from('1')), ipfsCID, 0, wei(100), 'name1', ['tag_1']);
      await modelRegistry
        .connect(SECOND)
        .modelRegister(getHex(Buffer.from('2')), ipfsCID, 0, wei(100), 'name2', ['tag_1']);

      const modelIds = await modelRegistry.getModelIds(0, 10);
      expect(modelIds.length).to.eq(2);
    });

    it('should throw error when caller is not an owner', async () => {
      await expect(modelRegistry.connect(SECOND).modelSetMinStake(100)).to.be.revertedWithCustomError(
        diamond,
        'OwnableUnauthorizedAccount',
      );
    });
  });

  describe('#modelRegister', async () => {
    it('should register a new model', async () => {
      await setNextTime(300);
      await modelRegistry.connect(SECOND).modelRegister(modelId, ipfsCID, 0, wei(100), 'name', ['tag_1']);

      const data = await modelRegistry.getModel(modelId);
      expect(data.ipfsCID).to.eq(ipfsCID);
      expect(data.fee).to.eq(0);
      expect(data.stake).to.eq(wei(100));
      expect(data.owner).to.eq(SECOND);
      expect(data.name).to.eq('name');
      expect(data.tags).deep.eq(['tag_1']);
      expect(data.createdAt).to.eq(300);
      expect(data.isDeleted).to.eq(false);
      expect(await modelRegistry.getIsModelActive(modelId)).to.eq(true);

      expect(await token.balanceOf(modelRegistry)).to.eq(wei(100));
      expect(await token.balanceOf(SECOND)).to.eq(wei(900));

      expect(await modelRegistry.getActiveModels(0, 10)).to.deep.eq([modelId]);

      await modelRegistry.connect(SECOND).modelRegister(modelId, ipfsCID, 0, wei(0), 'name', ['tag_1']);
    });
    it('should add stake to existed model', async () => {
      const ipfsCID2 = getHex(Buffer.from('ipfs://ipfsaddress/2'));

      await setNextTime(300);
      await modelRegistry.connect(SECOND).modelRegister(modelId, ipfsCID, 0, wei(100), 'name', ['tag_1']);
      await modelRegistry.connect(SECOND).modelRegister(modelId, ipfsCID2, 1, wei(300), 'name2', ['tag_1', 'tag_2']);

      const data = await modelRegistry.getModel(modelId);
      expect(data.ipfsCID).to.eq(ipfsCID2);
      expect(data.fee).to.eq(1);
      expect(data.stake).to.eq(wei(400));
      expect(data.owner).to.eq(SECOND);
      expect(data.name).to.eq('name2');
      expect(data.tags).deep.eq(['tag_1', 'tag_2']);
      expect(data.createdAt).to.eq(300);
      expect(data.isDeleted).to.eq(false);
      expect(await modelRegistry.getIsModelActive(modelId)).to.eq(true);

      expect(await token.balanceOf(modelRegistry)).to.eq(wei(400));
      expect(await token.balanceOf(SECOND)).to.eq(wei(600));
    });
    it('should activate deregistered model', async () => {
      await setNextTime(300);
      await modelRegistry.connect(SECOND).modelRegister(modelId, ipfsCID, 0, wei(100), 'name', ['tag_1']);
      await modelRegistry.connect(SECOND).modelDeregister(modelId);

      let data = await modelRegistry.getModel(modelId);
      expect(data.isDeleted).to.eq(true);
      expect(await modelRegistry.getIsModelActive(modelId)).to.eq(false);

      await modelRegistry.connect(SECOND).modelRegister(modelId, ipfsCID, 4, wei(200), 'name3', ['tag_3']);

      data = await modelRegistry.getModel(modelId);
      expect(data.ipfsCID).to.eq(ipfsCID);
      expect(data.fee).to.eq(4);
      expect(data.stake).to.eq(wei(200));
      expect(data.owner).to.eq(SECOND);
      expect(data.name).to.eq('name3');
      expect(data.tags).deep.eq(['tag_3']);
      expect(data.createdAt).to.eq(300);
      expect(data.isDeleted).to.eq(false);
      expect(await modelRegistry.getIsModelActive(modelId)).to.eq(true);
    });
    it('should throw error when the caller is not a model owner', async () => {
      await modelRegistry.connect(SECOND).modelRegister(modelId, ipfsCID, 0, wei(100), 'name', ['tag_1']);
      await expect(
        modelRegistry.connect(OWNER).modelRegister(modelId, ipfsCID, 0, wei(100), 'name', ['tag_1']),
      ).to.be.revertedWithCustomError(modelRegistry, 'OwnableUnauthorizedAccount');
    });
    it('should throw error when the stake is too low', async () => {
      await modelRegistry.modelSetMinStake(wei(2));
      await expect(
        modelRegistry.connect(SECOND).modelRegister(modelId, ipfsCID, 0, wei(1), 'name', ['tag_1']),
      ).to.be.revertedWithCustomError(modelRegistry, 'ModelStakeTooLow');
    });
  });

  describe('#modelDeregister', async () => {
    it('should deregister the model', async () => {
      await setNextTime(300);
      await modelRegistry.connect(SECOND).modelRegister(modelId, ipfsCID, 0, wei(100), 'name', ['tag_1']);
      await modelRegistry.connect(SECOND).modelDeregister(modelId);

      expect((await modelRegistry.getModel(modelId)).isDeleted).to.equal(true);
      expect(await modelRegistry.getIsModelActive(modelId)).to.eq(false);
      expect(await token.balanceOf(modelRegistry)).to.eq(0);
      expect(await token.balanceOf(SECOND)).to.eq(wei(1000));

      expect(await modelRegistry.getActiveModels(0, 10)).to.deep.eq([]);
    });
    it('should throw error when the caller is not an owner or specified address', async () => {
      await expect(modelRegistry.connect(SECOND).modelDeregister(modelId)).to.be.revertedWithCustomError(
        modelRegistry,
        'OwnableUnauthorizedAccount',
      );
    });
    it('should throw error when model has active bids', async () => {
      await providerRegistry.connect(SECOND).providerRegister(wei(100), 'test');
      await modelRegistry.connect(SECOND).modelRegister(modelId, ipfsCID, 0, wei(100), 'name', ['tag_1']);
      await marketplace.connect(SECOND).postModelBid(modelId, wei(10));
      await expect(modelRegistry.connect(SECOND).modelDeregister(modelId)).to.be.revertedWithCustomError(
        modelRegistry,
        'ModelHasActiveBids',
      );
    });
    it('should throw error when delete model few times', async () => {
      await modelRegistry.connect(SECOND).modelRegister(modelId, ipfsCID, 0, wei(100), 'name', ['tag_1']);
      await modelRegistry.connect(SECOND).modelDeregister(modelId);
      await expect(modelRegistry.connect(SECOND).modelDeregister(modelId)).to.be.revertedWithCustomError(
        modelRegistry,
        'ModelHasAlreadyDeregistered',
      );
    });
  });
});

// npm run generate-types && npx hardhat test "test/diamond/facets/ModelRegistry.test.ts"
// npx hardhat coverage --solcoverjs ./.solcover.ts --testfiles "test/diamond/facets/ModelRegistry.test.ts"
