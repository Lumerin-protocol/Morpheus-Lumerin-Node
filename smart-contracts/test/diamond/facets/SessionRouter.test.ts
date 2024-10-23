import { LumerinDiamond, Marketplace, ModelRegistry, MorpheusToken, ProviderRegistry, SessionRouter } from '@ethers-v6';
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
import { payoutStart } from '@/test/helpers/pool-helper';
import { Reverter } from '@/test/helpers/reverter';
import { setTime } from '@/utils/block-helper';
import { getProviderApproval, getReceipt } from '@/utils/provider-helper';
import { DAY } from '@/utils/time';

describe('SessionRouter', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let SECOND: SignerWithAddress;
  let FUNDING: SignerWithAddress;
  let PROVIDER: SignerWithAddress;

  let diamond: LumerinDiamond;
  let marketplace: Marketplace;
  let modelRegistry: ModelRegistry;
  let providerRegistry: ProviderRegistry;
  let sessionRouter: SessionRouter;

  let token: MorpheusToken;

  let bidId = '';
  const modelId = getHex(Buffer.from('1'));
  const bidPricePerSecond = wei(0.0001);

  before(async () => {
    [OWNER, SECOND, FUNDING, PROVIDER] = await ethers.getSigners();

    [diamond, token] = await Promise.all([deployLumerinDiamond(), deployMORToken()]);

    [providerRegistry, modelRegistry, sessionRouter, marketplace] = await Promise.all([
      deployFacetProviderRegistry(diamond),
      deployFacetModelRegistry(diamond),
      deployFacetSessionRouter(diamond, FUNDING),
      deployFacetMarketplace(diamond, token),
    ]);

    await token.transfer(SECOND, wei(10000));
    await token.transfer(PROVIDER, wei(10000));
    await token.transfer(FUNDING, wei(10000));
    await token.connect(PROVIDER).approve(providerRegistry, wei(10000));
    await token.connect(PROVIDER).approve(modelRegistry, wei(10000));
    await token.connect(PROVIDER).approve(marketplace, wei(10000));
    await token.connect(SECOND).approve(sessionRouter, wei(10000));
    await token.connect(FUNDING).approve(sessionRouter, wei(10000));

    const ipfsCID = getHex(Buffer.from('ipfs://ipfsaddress'));
    await providerRegistry.connect(PROVIDER).providerRegister(wei(0.2), 'test');
    await modelRegistry.connect(PROVIDER).modelRegister(modelId, ipfsCID, 0, wei(100), 'name', ['tag_1']);

    await marketplace.connect(PROVIDER).postModelBid(modelId, bidPricePerSecond);
    bidId = await marketplace.getBidId(PROVIDER, modelId, 0);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('#__SessionRouter_init', () => {
    it('should set correct data after creation', async () => {
      expect(await sessionRouter.getFundingAccount()).to.eq(FUNDING);
      expect((await sessionRouter.getPools()).length).to.eq(5);
    });
    it('should revert if try to call init function twice', async () => {
      await expect(sessionRouter.__SessionRouter_init(FUNDING, [])).to.be.rejectedWith(
        'Initializable: contract is already initialized',
      );
    });
  });

  describe('#getTodaysBudget', () => {
    it('should return not zero amount', async () => {
      expect(await sessionRouter.getTodaysBudget(payoutStart)).to.eq(0);
      expect(await sessionRouter.getTodaysBudget(payoutStart + 10 * DAY)).to.greaterThan(0);
    });
  });

  describe('#setPoolConfig', () => {
    it('should reset pool', async () => {
      await sessionRouter.setPoolConfig(0, {
        payoutStart: 0,
        decreaseInterval: DAY,
        initialReward: wei(1000),
        rewardDecrease: wei(10),
      });

      const pool = await sessionRouter.getPool(0);
      expect(pool.payoutStart).to.eq(0);
      expect(pool.decreaseInterval).to.eq(DAY);
      expect(pool.initialReward).to.eq(wei(1000));
      expect(pool.rewardDecrease).to.eq(wei(10));
    });
    it('should throw error when the pool index is invalid', async () => {
      await expect(
        sessionRouter.setPoolConfig(100, {
          payoutStart: 0,
          decreaseInterval: DAY,
          initialReward: wei(1000),
          rewardDecrease: wei(10),
        }),
      ).to.be.revertedWithCustomError(sessionRouter, 'SessionPoolIndexOutOfBounds');
    });
    it('should throw error when the caller is invalid', async () => {
      await expect(
        sessionRouter.connect(SECOND).setPoolConfig(0, {
          payoutStart: 0,
          decreaseInterval: DAY,
          initialReward: wei(1000),
          rewardDecrease: wei(10),
        }),
      ).to.be.revertedWithCustomError(sessionRouter, 'OwnableUnauthorizedAccount');
    });
  });

  describe('#openSession', () => {
    let tokenBalBefore = 0n;
    let secondBalBefore = 0n;

    beforeEach(async () => {
      tokenBalBefore = await token.balanceOf(sessionRouter);
      secondBalBefore = await token.balanceOf(SECOND);
    });
    it('should open session', async () => {
      await setTime(payoutStart + 10 * DAY);
      const { msg, signature } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await sessionRouter.connect(SECOND).openSession(wei(50), msg, signature);

      const sessionId = await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 0);
      const data = await sessionRouter.getSession(sessionId);
      expect(data.user).to.eq(SECOND);
      expect(data.bidId).to.eq(bidId);
      expect(data.stake).to.eq(wei(50));
      expect(data.closeoutReceipt).to.eq('0x');
      expect(data.closeoutType).to.eq(0);
      expect(data.providerWithdrawnAmount).to.eq(0);
      expect(data.openedAt).to.eq(payoutStart + 10 * DAY + 1);
      expect(data.endsAt).to.greaterThan(data.openedAt);
      expect(data.closedAt).to.eq(0);
      expect(data.isActive).to.eq(true);

      const tokenBalAfter = await token.balanceOf(sessionRouter);
      expect(tokenBalAfter - tokenBalBefore).to.eq(wei(50));
      const secondBalAfter = await token.balanceOf(SECOND);
      expect(secondBalBefore - secondBalAfter).to.eq(wei(50));

      expect(await sessionRouter.getIsProviderApprovalUsed(msg)).to.eq(true);
      expect(await sessionRouter.getUserSessions(SECOND, 0, 10)).to.deep.eq([sessionId]);
      expect(await sessionRouter.getProviderSessions(PROVIDER, 0, 10)).to.deep.eq([sessionId]);
      expect(await sessionRouter.getModelSessions(modelId, 0, 10)).to.deep.eq([sessionId]);
    });
    it('should open two different session wit the same input params', async () => {
      await setTime(payoutStart + 10 * DAY);
      const { msg: msg1, signature: signature1 } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await setTime(payoutStart + 10 * DAY + 1);
      const { msg: msg2, signature: signature2 } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await sessionRouter.connect(SECOND).openSession(wei(50), msg1, signature1);
      await sessionRouter.connect(SECOND).openSession(wei(50), msg2, signature2);

      const sessionId1 = await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 0);
      const sessionId2 = await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 1);

      const tokenBalAfter = await token.balanceOf(sessionRouter);
      expect(tokenBalAfter - tokenBalBefore).to.eq(wei(100));
      const secondBalAfter = await token.balanceOf(SECOND);
      expect(secondBalBefore - secondBalAfter).to.eq(wei(100));

      expect(await sessionRouter.getIsProviderApprovalUsed(msg1)).to.eq(true);
      expect(await sessionRouter.getIsProviderApprovalUsed(msg2)).to.eq(true);
      expect(await sessionRouter.getUserSessions(SECOND, 0, 10)).to.deep.eq([sessionId1, sessionId2]);
      expect(await sessionRouter.getProviderSessions(PROVIDER, 0, 10)).to.deep.eq([sessionId1, sessionId2]);
      expect(await sessionRouter.getModelSessions(modelId, 0, 10)).to.deep.eq([sessionId1, sessionId2]);

      expect(await sessionRouter.getTotalSessions(PROVIDER)).to.eq(2);
    });
    it('should open session with max duration', async () => {
      await setTime(payoutStart + 10 * DAY);
      const { msg, signature } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await sessionRouter.connect(SECOND).openSession(wei(10000), msg, signature);

      const sessionId = await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 0);
      const data = await sessionRouter.getSession(sessionId);
      expect(data.endsAt).to.eq(Number(data.openedAt.toString()) + DAY);
    });
    it('should throw error when the approval is for an another user', async () => {
      const { msg, signature } = await getProviderApproval(PROVIDER, OWNER, bidId);
      await expect(sessionRouter.connect(SECOND).openSession(wei(50), msg, signature)).to.be.revertedWithCustomError(
        sessionRouter,
        'SessionApprovedForAnotherUser',
      );
    });
    it('should throw error when the approval is for an another chain', async () => {
      const { msg, signature } = await getProviderApproval(PROVIDER, SECOND, bidId, 1n);
      await expect(sessionRouter.connect(SECOND).openSession(wei(50), msg, signature)).to.be.revertedWithCustomError(
        sessionRouter,
        'SesssionApprovedForAnotherChainId',
      );
    });
    it('should throw error when an aprrove expired', async () => {
      await setTime(payoutStart);
      const { msg, signature } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await setTime(payoutStart + 600);
      await expect(sessionRouter.connect(SECOND).openSession(wei(50), msg, signature)).to.be.revertedWithCustomError(
        sessionRouter,
        'SesssionApproveExpired',
      );
    });
    it('should throw error when the bid is not found', async () => {
      const { msg, signature } = await getProviderApproval(PROVIDER, SECOND, getHex(Buffer.from('1')));
      await expect(sessionRouter.connect(SECOND).openSession(wei(50), msg, signature)).to.be.revertedWithCustomError(
        sessionRouter,
        'SessionBidNotFound',
      );
    });
    it('should throw error when the signature mismatch', async () => {
      const { msg, signature } = await getProviderApproval(OWNER, SECOND, bidId);
      await expect(sessionRouter.connect(SECOND).openSession(wei(50), msg, signature)).to.be.revertedWithCustomError(
        sessionRouter,
        'SessionProviderSignatureMismatch',
      );
    });
    it('should throw error when an approval duplicated', async () => {
      await setTime(payoutStart + 10 * DAY);
      const { msg, signature } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await sessionRouter.connect(SECOND).openSession(wei(50), msg, signature);
      await expect(sessionRouter.connect(SECOND).openSession(wei(50), msg, signature)).to.be.revertedWithCustomError(
        sessionRouter,
        'SessionDuplicateApproval',
      );
    });
    it('should throw error when session duration too short', async () => {
      const { msg, signature } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await expect(sessionRouter.connect(SECOND).openSession(wei(50), msg, signature)).to.be.revertedWithCustomError(
        sessionRouter,
        'SessionTooShort',
      );
    });
  });

  describe('#closeSession', () => {
    it('should close session and send rewards for the provider, with dispute, late closure', async () => {
      const { sessionId, openedAt } = await _createSession();

      const providerBalBefore = await token.balanceOf(PROVIDER);
      const fundingBalBefore = await token.balanceOf(FUNDING);

      await setTime(openedAt + 5 * DAY);
      const { msg: receiptMsg } = await getReceipt(PROVIDER, sessionId, 0, 0);
      const { signature: receiptSig } = await getReceipt(OWNER, sessionId, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg, receiptSig);

      const session = await sessionRouter.getSession(sessionId);
      const duration = session.endsAt - session.openedAt;

      expect(session.closedAt).to.eq(openedAt + 5 * DAY + 1);
      expect(session.isActive).to.eq(false);
      expect(session.closeoutReceipt).to.eq(receiptMsg);
      expect(session.providerWithdrawnAmount).to.eq(bidPricePerSecond * duration);

      expect((await sessionRouter.getProvider(PROVIDER)).limitPeriodEarned).to.eq(bidPricePerSecond * duration);

      const providerBalAfter = await token.balanceOf(PROVIDER);
      expect(providerBalAfter - providerBalBefore).to.eq(bidPricePerSecond * duration);
      const fundingBalAfter = await token.balanceOf(FUNDING);
      expect(fundingBalBefore - fundingBalAfter).to.eq(bidPricePerSecond * duration);
    });
    it('should close session and send rewards for the provider, no dispute, early closure', async () => {
      const { sessionId, openedAt } = await _createSession();
      const providerBalBefore = await token.balanceOf(PROVIDER);

      const fundingBalBefore = await token.balanceOf(FUNDING);
      await setTime(openedAt + 200);
      const { msg: receiptMsg, signature: receiptSig } = await getReceipt(PROVIDER, sessionId, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg, receiptSig);

      const session = await sessionRouter.getSession(sessionId);
      const duration = session.closedAt - session.openedAt;

      expect(session.closedAt).to.eq(openedAt + 201);
      expect(session.isActive).to.eq(false);
      expect(session.closeoutReceipt).to.eq(receiptMsg);
      expect(session.providerWithdrawnAmount).to.eq(bidPricePerSecond * duration);

      expect((await sessionRouter.getProvider(PROVIDER)).limitPeriodEarned).to.eq(bidPricePerSecond * duration);

      const providerBalAfter = await token.balanceOf(PROVIDER);
      expect(providerBalAfter - providerBalBefore).to.eq(bidPricePerSecond * duration);
      const fundingBalAfter = await token.balanceOf(FUNDING);
      expect(fundingBalBefore - fundingBalAfter).to.eq(bidPricePerSecond * duration);
    });
    it('should close session and send rewards for the provider, with dispute, late closure before end', async () => {
      const { sessionId, secondsToDayEnd, openedAt } = await _createSession();

      const providerBalBefore = await token.balanceOf(PROVIDER);
      const fundingBalBefore = await token.balanceOf(FUNDING);

      await setTime(openedAt + secondsToDayEnd + 1);
      const { msg: receiptMsg } = await getReceipt(PROVIDER, sessionId, 0, 0);
      const { signature: receiptSig } = await getReceipt(OWNER, sessionId, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg, receiptSig);

      const session = await sessionRouter.getSession(sessionId);
      const duration = BigInt(secondsToDayEnd);

      expect(session.closedAt).to.eq(openedAt + secondsToDayEnd + 2);
      expect(session.isActive).to.eq(false);
      expect(session.closeoutReceipt).to.eq(receiptMsg);
      expect(session.providerWithdrawnAmount).to.eq(bidPricePerSecond * duration);

      expect((await sessionRouter.getProvider(PROVIDER)).limitPeriodEarned).to.eq(bidPricePerSecond * duration);

      const providerBalAfter = await token.balanceOf(PROVIDER);
      expect(providerBalAfter - providerBalBefore).to.eq(bidPricePerSecond * duration);
      const fundingBalAfter = await token.balanceOf(FUNDING);
      expect(fundingBalBefore - fundingBalAfter).to.eq(bidPricePerSecond * duration);
    });
    it('should close session and do not send rewards for the provider, with dispute, same day closure', async () => {
      const { sessionId, secondsToDayEnd, openedAt } = await _createSession();

      const providerBalBefore = await token.balanceOf(PROVIDER);
      const fundingBalBefore = await token.balanceOf(FUNDING);

      await setTime(openedAt + secondsToDayEnd - 50);
      const { msg: receiptMsg } = await getReceipt(PROVIDER, sessionId, 0, 0);
      const { signature: receiptSig } = await getReceipt(OWNER, sessionId, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg, receiptSig);

      const session = await sessionRouter.getSession(sessionId);

      const duration = 0n;
      expect(session.closedAt).to.eq(openedAt + secondsToDayEnd - 49);
      expect(session.isActive).to.eq(false);
      expect(session.closeoutReceipt).to.eq(receiptMsg);

      expect(session.providerWithdrawnAmount).to.eq(bidPricePerSecond * duration);

      expect((await sessionRouter.getProvider(PROVIDER)).limitPeriodEarned).to.eq(bidPricePerSecond * duration);
      const providerBalAfter = await token.balanceOf(PROVIDER);
      expect(providerBalAfter - providerBalBefore).to.eq(bidPricePerSecond * duration);
      const fundingBalAfter = await token.balanceOf(FUNDING);
      expect(fundingBalBefore - fundingBalAfter).to.eq(bidPricePerSecond * duration);
    });
    it('should close session and send rewards for the user, late closure', async () => {
      const { sessionId, openedAt } = await _createSession();

      const userBalBefore = await token.balanceOf(SECOND);
      const contractBalBefore = await token.balanceOf(sessionRouter);

      await setTime(openedAt + 5 * DAY);
      const { msg: receiptMsg, signature: receiptSig } = await getReceipt(PROVIDER, sessionId, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg, receiptSig);

      const stakesOnHold = await sessionRouter.getUserStakesOnHold(SECOND, 20);
      expect(stakesOnHold[0]).to.eq(0);
      expect(stakesOnHold[1]).to.eq(0);

      const userBalAfter = await token.balanceOf(SECOND);
      expect(userBalAfter - userBalBefore).to.eq(wei(50));
      const contractBalAfter = await token.balanceOf(sessionRouter);
      expect(contractBalBefore - contractBalAfter).to.eq(wei(50));
    });
    it('should close session and send rewards for the user, early closure', async () => {
      const { sessionId, openedAt, secondsToDayEnd } = await _createSession();

      const userBalBefore = await token.balanceOf(SECOND);
      const contractBalBefore = await token.balanceOf(sessionRouter);

      await setTime(openedAt + secondsToDayEnd + 1);
      const { msg: receiptMsg, signature: receiptSig } = await getReceipt(PROVIDER, sessionId, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg, receiptSig);

      const stakesOnHold = await sessionRouter.getUserStakesOnHold(SECOND, 1);
      expect(stakesOnHold[0]).to.eq(0);
      expect(stakesOnHold[1]).to.greaterThan(0);

      const userBalAfter = await token.balanceOf(SECOND);
      expect(userBalAfter - userBalBefore).to.lessThan(wei(50));
      const contractBalAfter = await token.balanceOf(sessionRouter);
      expect(contractBalBefore - contractBalAfter).to.lessThan(wei(50));

      await sessionRouter.getProviderModelStats(modelId, PROVIDER);
      await sessionRouter.getModelStats(modelId);
    });
    it('should throw error when the caller is invalid', async () => {
      const { msg: receiptMsg, signature: receiptSig } = await getReceipt(PROVIDER, getHex(Buffer.from('1')), 0, 0);
      await expect(sessionRouter.connect(SECOND).closeSession(receiptMsg, receiptSig)).to.be.revertedWithCustomError(
        sessionRouter,
        'OwnableUnauthorizedAccount',
      );
    });
    it('should throw error when the session already closed', async () => {
      const { sessionId, openedAt, secondsToDayEnd } = await _createSession();

      await setTime(openedAt + secondsToDayEnd + 1);
      const { msg: receiptMsg, signature: receiptSig } = await getReceipt(PROVIDER, sessionId, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg, receiptSig);

      await expect(sessionRouter.connect(SECOND).closeSession(receiptMsg, receiptSig)).to.be.revertedWithCustomError(
        sessionRouter,
        'SessionAlreadyClosed',
      );
    });
    it('should throw error when the provider receipt for another chain', async () => {
      const { sessionId } = await _createSession();

      const { msg: receiptMsg, signature: receiptSig } = await getReceipt(PROVIDER, sessionId, 0, 0, 1n);
      await expect(sessionRouter.connect(SECOND).closeSession(receiptMsg, receiptSig)).to.be.revertedWithCustomError(
        sessionRouter,
        'SesssionReceiptForAnotherChainId',
      );
    });
    it('should throw error when the provider receipt expired', async () => {
      const { sessionId, openedAt } = await _createSession();

      await setTime(openedAt + 100);
      const { msg: receiptMsg, signature: receiptSig } = await getReceipt(PROVIDER, sessionId, 0, 0);

      await setTime(openedAt + 10000);
      await expect(sessionRouter.connect(SECOND).closeSession(receiptMsg, receiptSig)).to.be.revertedWithCustomError(
        sessionRouter,
        'SesssionReceiptExpired',
      );
    });
  });

  describe('#claimForProvider', () => {
    it('should claim provider rewards, remainder, session closed with dispute', async () => {
      const { sessionId, secondsToDayEnd, openedAt } = await _createSession();

      await setTime(openedAt + secondsToDayEnd + 1);
      const { msg: receiptMsg } = await getReceipt(PROVIDER, sessionId, 0, 0);
      const { signature: receiptSig } = await getReceipt(OWNER, sessionId, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg, receiptSig);

      let session = await sessionRouter.getSession(sessionId);
      const fullDuration = BigInt(secondsToDayEnd + 1);
      const duration = 1n;

      const providerBalBefore = await token.balanceOf(PROVIDER);
      const fundingBalBefore = await token.balanceOf(FUNDING);

      await sessionRouter.connect(PROVIDER).claimForProvider(sessionId);
      session = await sessionRouter.getSession(sessionId);

      expect(session.providerWithdrawnAmount).to.eq(bidPricePerSecond * fullDuration);
      expect(await sessionRouter.getProvidersTotalClaimed()).to.eq(bidPricePerSecond * fullDuration);
      expect((await sessionRouter.getProvider(PROVIDER)).limitPeriodEarned).to.eq(bidPricePerSecond * fullDuration);

      const providerBalAfter = await token.balanceOf(PROVIDER);
      expect(providerBalAfter - providerBalBefore).to.eq(bidPricePerSecond * duration);
      const fundingBalAfter = await token.balanceOf(FUNDING);
      expect(fundingBalBefore - fundingBalAfter).to.eq(bidPricePerSecond * duration);
    });
    it('should claim provider rewards, full', async () => {
      const { sessionId, openedAt } = await _createSession();

      let session = await sessionRouter.getSession(sessionId);
      const duration = session.endsAt - session.openedAt;

      const providerBalBefore = await token.balanceOf(PROVIDER);
      const fundingBalBefore = await token.balanceOf(FUNDING);

      await setTime(openedAt + 5 * DAY + 1);
      await sessionRouter.connect(PROVIDER).claimForProvider(sessionId);
      session = await sessionRouter.getSession(sessionId);

      expect(session.providerWithdrawnAmount).to.eq(bidPricePerSecond * duration);
      expect(await sessionRouter.getProvidersTotalClaimed()).to.eq(bidPricePerSecond * duration);
      expect((await sessionRouter.getProvider(PROVIDER)).limitPeriodEarned).to.eq(bidPricePerSecond * duration);

      const providerBalAfter = await token.balanceOf(PROVIDER);
      expect(providerBalAfter - providerBalBefore).to.eq(bidPricePerSecond * duration);
      const fundingBalAfter = await token.balanceOf(FUNDING);
      expect(fundingBalBefore - fundingBalAfter).to.eq(bidPricePerSecond * duration);
    });
    it('should claim provider rewards with reward limiter amount for the period', async () => {
      const providerBalBefore = await token.balanceOf(PROVIDER);
      const fundingBalBefore = await token.balanceOf(FUNDING);

      await setTime(payoutStart + 10 * DAY);
      const { msg: msg1, signature: sig1 } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await sessionRouter.connect(SECOND).openSession(wei(50), msg1, sig1);

      const sessionId1 = await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 0);
      await setTime(payoutStart + 20 * DAY);
      await sessionRouter.connect(PROVIDER).claimForProvider(sessionId1);

      await setTime(payoutStart + 30 * DAY);
      const { msg: msg2, signature: sig2 } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await sessionRouter.connect(SECOND).openSession(wei(50), msg2, sig2);

      const sessionId2 = await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 1);
      await setTime(payoutStart + 40 * DAY);
      await sessionRouter.connect(PROVIDER).claimForProvider(sessionId2);

      expect(await sessionRouter.getProvidersTotalClaimed()).to.eq(wei(0.2));
      expect((await sessionRouter.getProvider(PROVIDER)).limitPeriodEarned).to.eq(wei(0.2));

      const providerBalAfter = await token.balanceOf(PROVIDER);
      expect(providerBalAfter - providerBalBefore).to.eq(wei(0.2));
      const fundingBalAfter = await token.balanceOf(FUNDING);
      expect(fundingBalBefore - fundingBalAfter).to.eq(wei(0.2));
    });
    it('should throw error when caller is not the session provider', async () => {
      const { sessionId } = await _createSession();

      await expect(sessionRouter.connect(SECOND).claimForProvider(sessionId)).to.be.revertedWithCustomError(
        sessionRouter,
        'OwnableUnauthorizedAccount',
      );
    });
    it('should throw error when session is not end', async () => {
      const { sessionId, openedAt } = await _createSession();

      await setTime(openedAt + 10);
      await expect(sessionRouter.connect(PROVIDER).claimForProvider(sessionId)).to.be.revertedWithCustomError(
        sessionRouter,
        'SessionNotEndedOrNotExist',
      );
    });
  });

  describe('#claimForProvider', () => {
    it('should claim provider rewards, remainder, session closed with dispute', async () => {
      const { sessionId, secondsToDayEnd, openedAt } = await _createSession();

      await setTime(openedAt + secondsToDayEnd + 1);
      const { msg: receiptMsg } = await getReceipt(PROVIDER, sessionId, 0, 0);
      const { signature: receiptSig } = await getReceipt(OWNER, sessionId, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg, receiptSig);

      let session = await sessionRouter.getSession(sessionId);
      const fullDuration = BigInt(secondsToDayEnd + 1);
      const duration = 1n;

      const providerBalBefore = await token.balanceOf(PROVIDER);
      const fundingBalBefore = await token.balanceOf(FUNDING);

      await sessionRouter.connect(PROVIDER).claimForProvider(sessionId);
      session = await sessionRouter.getSession(sessionId);

      expect(session.providerWithdrawnAmount).to.eq(bidPricePerSecond * fullDuration);
      expect(await sessionRouter.getProvidersTotalClaimed()).to.eq(bidPricePerSecond * fullDuration);
      expect((await sessionRouter.getProvider(PROVIDER)).limitPeriodEarned).to.eq(bidPricePerSecond * fullDuration);

      const providerBalAfter = await token.balanceOf(PROVIDER);
      expect(providerBalAfter - providerBalBefore).to.eq(bidPricePerSecond * duration);
      const fundingBalAfter = await token.balanceOf(FUNDING);
      expect(fundingBalBefore - fundingBalAfter).to.eq(bidPricePerSecond * duration);
    });
    it('should claim provider rewards, full', async () => {
      const { sessionId, openedAt } = await _createSession();

      let session = await sessionRouter.getSession(sessionId);
      const duration = session.endsAt - session.openedAt;

      const providerBalBefore = await token.balanceOf(PROVIDER);
      const fundingBalBefore = await token.balanceOf(FUNDING);

      await setTime(openedAt + 5 * DAY + 1);
      await sessionRouter.connect(PROVIDER).claimForProvider(sessionId);
      session = await sessionRouter.getSession(sessionId);

      expect(session.providerWithdrawnAmount).to.eq(bidPricePerSecond * duration);
      expect(await sessionRouter.getProvidersTotalClaimed()).to.eq(bidPricePerSecond * duration);
      expect((await sessionRouter.getProvider(PROVIDER)).limitPeriodEarned).to.eq(bidPricePerSecond * duration);

      const providerBalAfter = await token.balanceOf(PROVIDER);
      expect(providerBalAfter - providerBalBefore).to.eq(bidPricePerSecond * duration);
      const fundingBalAfter = await token.balanceOf(FUNDING);
      expect(fundingBalBefore - fundingBalAfter).to.eq(bidPricePerSecond * duration);
    });
    it('should claim provider rewards with reward limiter amount for the period', async () => {
      const providerBalBefore = await token.balanceOf(PROVIDER);
      const fundingBalBefore = await token.balanceOf(FUNDING);

      await setTime(payoutStart + 10 * DAY);
      const { msg: msg1, signature: sig1 } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await sessionRouter.connect(SECOND).openSession(wei(50), msg1, sig1);

      const sessionId1 = await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 0);
      await setTime(payoutStart + 20 * DAY);
      await sessionRouter.connect(PROVIDER).claimForProvider(sessionId1);

      await setTime(payoutStart + 30 * DAY);
      const { msg: msg2, signature: sig2 } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await sessionRouter.connect(SECOND).openSession(wei(50), msg2, sig2);

      const sessionId2 = await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 1);
      await setTime(payoutStart + 40 * DAY);
      await sessionRouter.connect(PROVIDER).claimForProvider(sessionId2);

      expect(await sessionRouter.getProvidersTotalClaimed()).to.eq(wei(0.2));
      expect((await sessionRouter.getProvider(PROVIDER)).limitPeriodEarned).to.eq(wei(0.2));

      const providerBalAfter = await token.balanceOf(PROVIDER);
      expect(providerBalAfter - providerBalBefore).to.eq(wei(0.2));
      const fundingBalAfter = await token.balanceOf(FUNDING);
      expect(fundingBalBefore - fundingBalAfter).to.eq(wei(0.2));
    });
    it('should throw error when caller is not the session provider', async () => {
      const { sessionId } = await _createSession();

      await expect(sessionRouter.connect(SECOND).claimForProvider(sessionId)).to.be.revertedWithCustomError(
        sessionRouter,
        'OwnableUnauthorizedAccount',
      );
    });
    it('should throw error when session is not end', async () => {
      const { sessionId, openedAt } = await _createSession();

      await setTime(openedAt + 10);
      await expect(sessionRouter.connect(PROVIDER).claimForProvider(sessionId)).to.be.revertedWithCustomError(
        sessionRouter,
        'SessionNotEndedOrNotExist',
      );
    });
  });

  describe('#withdrawUserStakes', () => {
    it('should withdraw the user stake on hold, one entity', async () => {
      const openedAt = payoutStart + (payoutStart % DAY) + 10 * DAY - 201;

      await setTime(openedAt + 1 * DAY);
      const { msg, signature } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await sessionRouter.connect(SECOND).openSession(wei(50), msg, signature);
      const sessionId1 = await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 0);

      await setTime(openedAt + 2 * DAY + 1);
      const { msg: receiptMsg, signature: receiptSig } = await getReceipt(PROVIDER, sessionId1, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg, receiptSig);

      const userBalBefore = await token.balanceOf(SECOND);
      const contractBalBefore = await token.balanceOf(sessionRouter);

      await setTime(openedAt + 3 * DAY + 2);
      await sessionRouter.connect(SECOND).withdrawUserStakes(1);

      const stakesOnHold = await sessionRouter.getUserStakesOnHold(SECOND, 1);
      expect(stakesOnHold[0]).to.eq(0);
      expect(stakesOnHold[1]).to.eq(0);

      const userBalAfter = await token.balanceOf(SECOND);
      expect(userBalAfter - userBalBefore).to.greaterThan(0);
      const contractBalAfter = await token.balanceOf(sessionRouter);
      expect(contractBalBefore - contractBalAfter).to.greaterThan(0);
    });
    it('should withdraw the user stake on hold, few entities, with on hold', async () => {
      const openedAt = payoutStart + (payoutStart % DAY) + 10 * DAY - 201;

      await setTime(openedAt + 1 * DAY);
      const { msg: msg1, signature: sig1 } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await sessionRouter.connect(SECOND).openSession(wei(50), msg1, sig1);
      const sessionId1 = await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 0);

      await setTime(openedAt + 1 * DAY + 500);
      const { msg: receiptMsg1, signature: receiptSig1 } = await getReceipt(PROVIDER, sessionId1, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg1, receiptSig1);

      await setTime(openedAt + 3 * DAY);
      const { msg: msg2, signature: sig2 } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await sessionRouter.connect(SECOND).openSession(wei(50), msg2, sig2);
      const sessionId2 = await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 1);

      await setTime(openedAt + 3 * DAY + 500);
      const { msg: receiptMsg2, signature: receiptSig2 } = await getReceipt(PROVIDER, sessionId2, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg2, receiptSig2);

      const stakesOnHold = await sessionRouter.getUserStakesOnHold(SECOND, 20);
      expect(stakesOnHold[0]).to.greaterThan(0);
      expect(stakesOnHold[1]).to.greaterThan(0);

      const userBalBefore = await token.balanceOf(SECOND);
      const contractBalBefore = await token.balanceOf(sessionRouter);

      await setTime(openedAt + 4 * DAY);
      await sessionRouter.connect(SECOND).withdrawUserStakes(20);

      const userBalAfter = await token.balanceOf(SECOND);
      expect(userBalAfter - userBalBefore).to.eq(stakesOnHold[0]);
      const contractBalAfter = await token.balanceOf(sessionRouter);
      expect(contractBalBefore - contractBalAfter).to.eq(stakesOnHold[0]);
    });
    it('should withdraw the user stake on hold, few entities, without on hold', async () => {
      const openedAt = payoutStart + (payoutStart % DAY) + 10 * DAY - 201;

      await setTime(openedAt + 1 * DAY);
      const { msg: msg1, signature: sig1 } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await sessionRouter.connect(SECOND).openSession(wei(50), msg1, sig1);
      const sessionId1 = await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 0);

      await setTime(openedAt + 1 * DAY + 500);
      const { msg: receiptMsg1, signature: receiptSig1 } = await getReceipt(PROVIDER, sessionId1, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg1, receiptSig1);

      await setTime(openedAt + 3 * DAY);
      const { msg: msg2, signature: sig2 } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await sessionRouter.connect(SECOND).openSession(wei(50), msg2, sig2);
      const sessionId2 = await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 1);

      await setTime(openedAt + 3 * DAY + 500);
      const { msg: receiptMsg2, signature: receiptSig2 } = await getReceipt(PROVIDER, sessionId2, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg2, receiptSig2);

      await setTime(openedAt + 10 * DAY);
      const stakesOnHold = await sessionRouter.getUserStakesOnHold(SECOND, 20);
      expect(stakesOnHold[0]).to.greaterThan(0);
      expect(stakesOnHold[1]).to.eq(0);

      const userBalBefore = await token.balanceOf(SECOND);
      const contractBalBefore = await token.balanceOf(sessionRouter);

      await sessionRouter.connect(SECOND).withdrawUserStakes(20);

      const userBalAfter = await token.balanceOf(SECOND);
      expect(userBalAfter - userBalBefore).to.eq(stakesOnHold[0]);
      const contractBalAfter = await token.balanceOf(sessionRouter);
      expect(contractBalBefore - contractBalAfter).to.eq(stakesOnHold[0]);
    });
    it('should withdraw the user stake on hold, few entities, without on hold, partial withdraw', async () => {
      const openedAt = payoutStart + (payoutStart % DAY) + 10 * DAY - 201;

      await setTime(openedAt + 1 * DAY);
      const { msg: msg1, signature: sig1 } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await sessionRouter.connect(SECOND).openSession(wei(50), msg1, sig1);
      const sessionId1 = await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 0);

      await setTime(openedAt + 1 * DAY + 500);
      const { msg: receiptMsg1, signature: receiptSig1 } = await getReceipt(PROVIDER, sessionId1, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg1, receiptSig1);

      await setTime(openedAt + 3 * DAY);
      const { msg: msg2, signature: sig2 } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await sessionRouter.connect(SECOND).openSession(wei(50), msg2, sig2);
      const sessionId2 = await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 1);

      await setTime(openedAt + 3 * DAY + 500);
      const { msg: receiptMsg2, signature: receiptSig2 } = await getReceipt(PROVIDER, sessionId2, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg2, receiptSig2);

      await setTime(openedAt + 10 * DAY);
      // First withdraw
      let stakesOnHold = await sessionRouter.getUserStakesOnHold(SECOND, 1);
      expect(stakesOnHold[0]).to.greaterThan(0);
      expect(stakesOnHold[1]).to.eq(0);

      let userBalBefore = await token.balanceOf(SECOND);
      let contractBalBefore = await token.balanceOf(sessionRouter);

      await sessionRouter.connect(SECOND).withdrawUserStakes(1);

      let userBalAfter = await token.balanceOf(SECOND);
      expect(userBalAfter - userBalBefore).to.lessThan(stakesOnHold[0]);
      let contractBalAfter = await token.balanceOf(sessionRouter);
      expect(contractBalBefore - contractBalAfter).to.lessThan(stakesOnHold[0]);

      // Second withdraw
      stakesOnHold = await sessionRouter.getUserStakesOnHold(SECOND, 1);
      expect(stakesOnHold[0]).to.greaterThan(0);
      expect(stakesOnHold[1]).to.eq(0);

      userBalBefore = await token.balanceOf(SECOND);
      contractBalBefore = await token.balanceOf(sessionRouter);

      await sessionRouter.connect(SECOND).withdrawUserStakes(1);

      userBalAfter = await token.balanceOf(SECOND);
      expect(userBalAfter - userBalBefore).to.eq(stakesOnHold[0]);
      contractBalAfter = await token.balanceOf(sessionRouter);
      expect(contractBalBefore - contractBalAfter).to.eq(stakesOnHold[0]);
    });
    it('should throw error when withdraw amount is zero', async () => {
      const openedAt = payoutStart + (payoutStart % DAY) + 10 * DAY - 201;

      await setTime(openedAt + 1 * DAY);
      const { msg: msg1, signature: sig1 } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await sessionRouter.connect(SECOND).openSession(wei(50), msg1, sig1);
      const sessionId1 = await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 0);

      const { msg: msg2, signature: sig2 } = await getProviderApproval(PROVIDER, SECOND, bidId);
      await sessionRouter.connect(SECOND).openSession(wei(50), msg2, sig2);
      const sessionId2 = await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 1);

      await setTime(openedAt + 1 * DAY + 500);
      const { msg: receiptMsg1, signature: receiptSig1 } = await getReceipt(PROVIDER, sessionId1, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg1, receiptSig1);

      await setTime(openedAt + 1 * DAY + 550);
      const { msg: receiptMsg2, signature: receiptSig2 } = await getReceipt(PROVIDER, sessionId2, 0, 0);
      await sessionRouter.connect(SECOND).closeSession(receiptMsg2, receiptSig2);

      await expect(sessionRouter.connect(SECOND).withdrawUserStakes(20)).to.be.revertedWithCustomError(
        sessionRouter,
        'SessionUserAmountToWithdrawIsZero',
      );
    });
  });

  const _createSession = async () => {
    const secondsToDayEnd = 600n;
    const openedAt = payoutStart + (payoutStart % DAY) + 10 * DAY - Number(secondsToDayEnd) - 1;

    await setTime(openedAt);
    const { msg, signature } = await getProviderApproval(PROVIDER, SECOND, bidId);
    await sessionRouter.connect(SECOND).openSession(wei(50), msg, signature);

    return {
      sessionId: await sessionRouter.getSessionId(SECOND, PROVIDER, bidId, 0),
      secondsToDayEnd: Number(secondsToDayEnd),
      openedAt,
    };
  };
});

// npm run generate-types && npx hardhat test "test/diamond/facets/SessionRouter.test.ts"
// npx hardhat coverage --solcoverjs ./.solcover.ts --testfiles "test/diamond/facets/SessionRouter.test.ts"
