import { expect } from 'chai';
import { ethers } from 'hardhat';

import { Reverter } from '../helpers/reverter';

import { LumerinToken } from '@/generated-types/ethers/contracts/tokens/LumerinToken';

describe('LumerinToken', () => {
  const reverter = new Reverter();

  let LMR: LumerinToken;

  before(async () => {
    const LMRFactory = await ethers.getContractFactory('LumerinToken');

    LMR = await LMRFactory.deploy('Lumerin Token', 'LMR');

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('#decimals', () => {
    it('should return the correct number of decimals', async () => {
      const decimals = await LMR.decimals();
      expect(decimals).to.equal(8);
    });
  });
});
