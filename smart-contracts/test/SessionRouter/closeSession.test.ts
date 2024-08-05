import {
  loadFixture,
  time,
} from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import { expect } from "chai";
import hre from "hardhat";
import { deploySingleBid, getProviderApproval, getReport } from "../fixtures";
import { catchError, getSessionId, getTxTimestamp, nowChain } from "../utils";
import { DAY, SECOND } from "../../utils/time";
import { expectAlmostEqual, expectAlmostEqualDelta } from "../../utils/compare";
import { getAddress, maxUint8, parseUnits } from "viem";

describe("session closeout", function () {
  it("should open short (<1D) session and close after expiration", async function () {
    const {
      sessionRouter,
      provider,
      expectedSession: exp,
      user,
      publicClient,
      tokenMOR,
    } = await loadFixture(deploySingleBid);

    // open session
    const { msg, signature } = await getProviderApproval(
      provider,
      user.account.address,
      exp.bidID,
    );
    const openTx = await sessionRouter.write.openSession(
      [exp.stake, msg, signature],
      { account: user.account.address },
    );
    const sessionId = await getSessionId(publicClient, hre, openTx);

    await time.increase(exp.durationSeconds * 2n);

    const userBalanceBefore = await tokenMOR.read.balanceOf([
      user.account.address,
    ]);
    const providerBalanceBefore = await tokenMOR.read.balanceOf([
      provider.account.address,
    ]);

    // close session
    const report = await getReport(provider, sessionId, 10, 1000);
    await sessionRouter.write.closeSession([report.msg, report.sig], {
      account: user.account,
    });

    // verify session is closed without dispute
    const session = await sessionRouter.read.getSession([sessionId]);
    expect(session.closeoutType).to.equal(0n);

    // verify balances
    const userBalanceAfter = await tokenMOR.read.balanceOf([
      user.account.address,
    ]);
    const providerBalanceAfter = await tokenMOR.read.balanceOf([
      provider.account.address,
    ]);

    const userStakeReturned = userBalanceAfter - userBalanceBefore;
    const providerEarned = providerBalanceAfter - providerBalanceBefore;

    const totalPrice =
      (session.endsAt - session.openedAt) * session.pricePerSecond;

    expectAlmostEqualDelta(
      0,
      userStakeReturned,
      Number(session.pricePerSecond) * 5,
    );
    expectAlmostEqual(providerEarned, totalPrice, 0.0001);
  });

  it("should open short (<1D) session and close early", async function () {
    const {
      sessionRouter,
      provider,
      expectedSession: exp,
      user,
      publicClient,
      tokenMOR,
    } = await loadFixture(deploySingleBid);

    // open session
    const { msg, signature } = await getProviderApproval(
      provider,
      user.account.address,
      exp.bidID,
    );
    const openTx = await sessionRouter.write.openSession(
      [exp.stake, msg, signature],
      { account: user.account.address },
    );
    const sessionId = await getSessionId(publicClient, hre, openTx);

    await time.increase(exp.durationSeconds / 2n - 1n);

    const userBalanceBefore = await tokenMOR.read.balanceOf([
      user.account.address,
    ]);
    const providerBalanceBefore = await tokenMOR.read.balanceOf([
      provider.account.address,
    ]);

    // close session
    const report = await getReport(provider, sessionId, 10, 1000);
    await sessionRouter.write.closeSession([report.msg, report.sig], {
      account: user.account,
    });

    // verify session is closed without dispute
    const session = await sessionRouter.read.getSession([sessionId]);
    expect(session.closeoutType).to.equal(0n);

    // verify balances
    const userBalanceAfter = await tokenMOR.read.balanceOf([
      user.account.address,
    ]);
    const providerBalanceAfter = await tokenMOR.read.balanceOf([
      provider.account.address,
    ]);

    const userStakeReturned = userBalanceAfter - userBalanceBefore;
    const providerEarned = providerBalanceAfter - providerBalanceBefore;

    expectAlmostEqual(userStakeReturned, exp.stake / 2n, 0.001);
    expectAlmostEqual(providerEarned, exp.totalCost / 2n, 0.0001);
  });

  it("should open and close early with user report - dispute", async function () {
    const {
      sessionRouter,
      provider,
      expectedSession: exp,
      user,
      publicClient,
      tokenMOR,
    } = await loadFixture(deploySingleBid);

    // open session
    const { msg, signature } = await getProviderApproval(
      provider,
      user.account.address,
      exp.bidID,
    );
    const openTx = await sessionRouter.write.openSession(
      [exp.stake, msg, signature],
      { account: user.account.address },
    );
    const sessionId = await getSessionId(publicClient, hre, openTx);

    // wait half of the session
    await time.increase(exp.durationSeconds / 2n - 1n);

    const userBalanceBefore = await tokenMOR.read.balanceOf([
      user.account.address,
    ]);
    const providerBalanceBefore = await tokenMOR.read.balanceOf([
      provider.account.address,
    ]);

    // close session with user signature
    const report = await getReport(user, sessionId, 10, 1000);
    await sessionRouter.write.closeSession([report.msg, report.sig], {
      account: user.account,
    });

    // verify session is closed with dispute
    const session = await sessionRouter.read.getSession([sessionId]);
    const totalCost =
      session.pricePerSecond * (session.closedAt - session.openedAt);

    // verify balances
    const userBalanceAfter = await tokenMOR.read.balanceOf([
      user.account.address,
    ]);
    const providerBalanceAfter = await tokenMOR.read.balanceOf([
      provider.account.address,
    ]);

    const claimableProvider =
      await sessionRouter.read.getProviderClaimableBalance([session.id]);

    const [userAvail, userHold] =
      await sessionRouter.read.withdrawableUserStake([
        session.user,
        0n,
        Number(maxUint8),
      ]);

    expect(session.closeoutType).to.equal(1n);
    expect(providerBalanceAfter - providerBalanceBefore).to.equal(0n);
    expect(claimableProvider).to.equal(0n);
    expectAlmostEqual(
      exp.stake / 2n,
      userBalanceAfter - userBalanceBefore,
      0.05,
    );
    expect(userAvail).to.equal(0n);
    expectAlmostEqual(userHold, userBalanceAfter - userBalanceBefore, 0.05);

    // verify provider balance after dispute is released
    await time.increase((1 * DAY) / SECOND);
    const claimableProvider2 =
      await sessionRouter.read.getProviderClaimableBalance([sessionId]);
    expect(claimableProvider2).to.equal(totalCost);

    // claim provider balance
    await sessionRouter.write.claimProviderBalance([
      sessionId,
      claimableProvider2,
    ]);

    // verify provider balance after claim
    const providerBalanceAfterClaim = await tokenMOR.read.balanceOf([
      provider.account.address,
    ]);
    const providerClaimed = providerBalanceAfterClaim - providerBalanceAfter;
    expect(providerClaimed).to.equal(totalCost);
  });

  it("should error when not a user trying to close", async function () {
    const {
      sessionRouter,
      provider,
      expectedSession: exp,
      user,
      publicClient,
    } = await loadFixture(deploySingleBid);

    // open session
    const { msg, signature } = await getProviderApproval(
      provider,
      user.account.address,
      exp.bidID,
    );
    const openTx = await sessionRouter.write.openSession(
      [exp.stake, msg, signature],
      { account: user.account.address },
    );
    const sessionId = await getSessionId(publicClient, hre, openTx);

    // wait half of the session
    await time.increase(exp.durationSeconds / 2n - 1n);

    // close session with user signature
    const report = await getReport(user, sessionId, 10, 10);

    await catchError(sessionRouter.abi, "NotSenderOrOwner", async () => {
      await sessionRouter.write.closeSession([report.msg, report.sig], {
        account: provider.account,
      });
    });
  });

  it("should limit reward by stake amount", async function () {
    const {
      sessionRouter,
      marketplace,
      expectedProvider,
      expectedModel,
      provider,
      decimalsMOR,
      user,
      modelRegistry,
      providerRegistry,
      publicClient,
      tokenMOR,
    } = await loadFixture(deploySingleBid);

    // expected bid
    const expectedBid = {
      id: "" as `0x${string}`,
      providerAddr: getAddress(expectedProvider.address),
      modelId: expectedModel.modelId,
      pricePerSecond: parseUnits("0.1", decimalsMOR),
      nonce: 0n,
      createdAt: 0n,
      deletedAt: 0n,
    };

    // add single bid
    const postBidtx = await marketplace.simulate.postModelBid(
      [
        expectedBid.providerAddr,
        expectedBid.modelId,
        expectedBid.pricePerSecond,
      ],
      { account: provider.account.address },
    );
    const txHash = await provider.writeContract(postBidtx.request);

    expectedBid.id = postBidtx.result;
    expectedBid.createdAt = await getTxTimestamp(publicClient, txHash);

    // calculate data for session opening
    const totalCost = expectedProvider.stake * 2n;
    const durationSeconds = totalCost / expectedBid.pricePerSecond;
    const totalSupply = await sessionRouter.read.totalMORSupply([
      await nowChain(),
    ]);
    const todaysBudget = await sessionRouter.read.getTodaysBudget([
      await nowChain(),
    ]);

    const expectedSession = {
      durationSeconds,
      totalCost,
      pricePerSecond: expectedBid.pricePerSecond,
      user: getAddress(user.account.address),
      provider: expectedBid.providerAddr,
      modelAgentId: expectedBid.modelId,
      bidID: expectedBid.id,
      stake: (totalCost * totalSupply) / todaysBudget,
    };

    // set user balance and approve funds
    await tokenMOR.write.transfer([
      user.account.address,
      expectedSession.stake,
    ]);
    await tokenMOR.write.approve(
      [modelRegistry.address, expectedSession.stake],
      {
        account: user.account,
      },
    );

    // open session
    const { msg, signature } = await getProviderApproval(
      provider,
      user.account.address,
      expectedSession.bidID,
    );
    const openTx = await sessionRouter.write.openSession(
      [expectedSession.stake, msg, signature],
      { account: user.account.address },
    );
    const sessionId = await getSessionId(publicClient, hre, openTx);

    // wait till session ends
    await time.increase(expectedSession.durationSeconds);

    const providerBalanceBefore = await tokenMOR.read.balanceOf([
      provider.account.address,
    ]);
    // close session without dispute
    const report = await getReport(provider, sessionId, 10, 1000);
    await sessionRouter.write.closeSession([report.msg, report.sig], {
      account: user.account,
    });

    const providerBalanceAfter = await tokenMOR.read.balanceOf([
      provider.account.address,
    ]);

    const providerEarned = providerBalanceAfter - providerBalanceBefore;

    expect(providerEarned).to.equal(expectedProvider.stake);

    // check provider record if earning was updated
    const providerRecord = await providerRegistry.read.providerMap([
      provider.account.address,
    ]);
    expect(providerRecord.limitPeriodEarned).to.equal(expectedProvider.stake);
  });

  it("should reset provider limitPeriodEarned after period", async function () {});

  it("should error with WithdrawableBalanceLimitByStakeReached() if claiming more that stake for a period", async function () {});
});
