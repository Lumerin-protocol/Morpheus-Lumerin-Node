import { LumerinDiamond } from '@ethers-v6';
import { SignerWithAddress } from '@nomicfoundation/hardhat-ethers/signers';
import { expect } from 'chai';
import { ethers } from 'hardhat';

import { deployLumerinDiamond } from '../helpers/deployers';
import { Reverter } from '../helpers/reverter';

describe('LumerinDiamond', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;

  let diamond: LumerinDiamond;

  before(async () => {
    [OWNER] = await ethers.getSigners();

    diamond = await deployLumerinDiamond();

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('#__LumerinDiamond_init', () => {
    it('should set correct data after creation', async () => {
      expect(await diamond.owner()).to.eq(await OWNER.getAddress());
    });
    it('should revert if try to call init function twice', async () => {
      await expect(diamond.__LumerinDiamond_init()).to.be.rejectedWith(
        'Initializable: contract is already initialized',
      );
    });
  });
});

// npx hardhat test "test/diamond/LumerinDiamond.test.ts"
// npx hardhat coverage --solcoverjs ./.solcover.ts --testfiles "test/diamond/LumerinDiamond.test.ts"
