import {
  loadFixture,
  time,
} from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import { expect } from "chai";
import hre from "hardhat";
import {
  deploySingleBid,
  getProviderApproval,
  openEarlyCloseSession,
} from "../fixtures";
import {
  NewDate,
  mine,
  catchError,
  getHex,
  getSessionId,
  getTxTimestamp,
  nowChain,
  randomBytes,
  randomBytes32,
  setAutomine,
  startOfTheDay,
  randomAddress,
} from "../utils";
import { DAY, HOUR, SECOND } from "../../utils/time";
import { UnknownRpcError } from "viem";

describe("session actions", function () {
  describe("positive cases", function () {
    this.afterAll(async function () {
      await setAutomine(hre, true);
    });

    it("should open session without error", async function () {
      const {
        sessionRouter,
        expectedSession: exp,
        user,
        publicClient,
        provider,
      } = await loadFixture(deploySingleBid);

      const { msg, signature } = await getProviderApproval(
        provider,
        user.account.address,
        exp.bidID,
      );
      const openTx = await sessionRouter.write.openSession(
        [exp.stake, msg, signature],
        { account: user.account },
      );

      const sessionId = await getSessionId(publicClient, hre, openTx);
      expect(sessionId).to.be.a("string");
    });

    it("should emit SessionOpened event", async function () {
      const {
        sessionRouter,
        expectedSession: exp,
        user,
        provider,
        publicClient,
      } = await loadFixture(deploySingleBid);

      const { msg, signature } = await getProviderApproval(
        provider,
        user.account.address,
        exp.bidID,
      );
      const hash = await sessionRouter.write.openSession(
        [exp.stake, msg, signature],
        { account: user.account },
      );

      await publicClient.waitForTransactionReceipt({ hash });
      const ev = await sessionRouter.getEvents.SessionOpened({
        providerId: exp.provider,
        userAddress: exp.user,
      });
      expect(ev.length).to.eq(1);
    });

    it("should verify session fields after opening", async function () {
      const {
        sessionRouter,
        expectedSession: exp,
        user,
        publicClient,
        provider,
      } = await loadFixture(deploySingleBid);

      const { msg, signature } = await getProviderApproval(
        provider,
        user.account.address,
        exp.bidID,
      );
      const openTx = await sessionRouter.write.openSession(
        [exp.stake, msg, signature],
        { account: user.account },
      );

      const sessionId = await getSessionId(publicClient, hre, openTx);
      const session = await sessionRouter.read.getSession([sessionId]);
      const createdAt = await getTxTimestamp(publicClient, openTx);

      expect(session).to.deep.include({
        id: sessionId,
        user: exp.user,
        provider: exp.provider,
        modelAgentId: exp.modelAgentId,
        bidID: exp.bidID,
        stake: exp.stake,
        pricePerSecond: exp.pricePerSecond,
        closeoutReceipt: getHex(Buffer.from(""), 0),
        closeoutType: 0n,
        providerWithdrawnAmount: 0n,
        openedAt: createdAt,
        // endsAt: verified in session end time tests
        closedAt: 0n,
      });
    });

    it("should verify balances after opening", async function () {
      const {
        sessionRouter,
        expectedSession: exp,
        publicClient,
        user,
        tokenMOR,
        provider,
      } = await loadFixture(deploySingleBid);

      const srBefore = await tokenMOR.read.balanceOf([sessionRouter.address]);
      const userBefore = await tokenMOR.read.balanceOf([user.account.address]);

      const { msg, signature } = await getProviderApproval(
        provider,
        user.account.address,
        exp.bidID,
      );
      const txHash = await sessionRouter.write.openSession(
        [exp.stake, msg, signature],
        { account: user.account },
      );
      await publicClient.waitForTransactionReceipt({ hash: txHash });

      const srAfter = await tokenMOR.read.balanceOf([sessionRouter.address]);
      const userAfter = await tokenMOR.read.balanceOf([user.account.address]);

      expect(srAfter - srBefore).to.equal(exp.stake);
      expect(userBefore - userAfter).to.equal(exp.stake);
    });

    it("should allow opening two sessions in the same block", async function () {
      const {
        sessionRouter,
        expectedSession: exp,
        user,
        owner,
        publicClient,
        provider,
        tokenMOR,
      } = await loadFixture(deploySingleBid);

      await tokenMOR.write.transfer([user.account.address, exp.stake * 2n], {
        account: owner.account.address,
      });
      await tokenMOR.write.approve([sessionRouter.address, exp.stake * 2n], {
        account: user.account.address,
      });

      const apprv1 = await getProviderApproval(
        provider,
        user.account.address,
        exp.bidID,
      );
      await time.increase(1);
      const apprv2 = await getProviderApproval(
        provider,
        user.account.address,
        exp.bidID,
      );

      await setAutomine(hre, false);

      const openSession1 = await sessionRouter.simulate.openSession(
        [exp.stake, apprv1.msg, apprv1.signature],
        { account: user.account.address },
      );
      const openTx1 = await user.writeContract(openSession1.request);

      const openSession2 = await sessionRouter.simulate.openSession(
        [exp.stake, apprv2.msg, apprv2.signature],
        { account: user.account.address },
      );
      const openTx2 = await user.writeContract(openSession2.request);

      await mine(hre);
      await setAutomine(hre, true);

      const sessionId1 = await getSessionId(publicClient, hre, openTx1);
      const sessionId2 = await getSessionId(publicClient, hre, openTx2);

      expect(sessionId1).not.to.equal(sessionId2);

      const session1 = await sessionRouter.read.getSession([sessionId1]);
      const session2 = await sessionRouter.read.getSession([sessionId2]);

      expect(session1.stake).to.equal(exp.stake);
      expect(session2.stake).to.equal(exp.stake);
    });

    it("should partially use remaining staked tokens for the opening session", async function () {
      const { sessionRouter, user, provider, expectedSession, tokenMOR } =
        await loadFixture(openEarlyCloseSession);
      await time.increaseTo(
        startOfTheDay(await nowChain()) + BigInt(DAY / SECOND),
      );

      const [avail] = await sessionRouter.read.withdrawableUserStake([
        user.account.address,
        0n,
        255,
      ]);
      expect(avail > 0).to.be.true;

      // reset allowance
      await tokenMOR.write.approve([sessionRouter.address, 0n], {
        account: user.account.address,
      });

      const stake = avail / 2n;

      const approval = await getProviderApproval(
        provider,
        user.account.address,
        expectedSession.bidID,
      );
      const openSession = await sessionRouter.simulate.openSession(
        [stake, approval.msg, approval.signature],
        { account: user.account.address },
      );
      await user.writeContract(openSession.request);

      const [avail2] = await sessionRouter.read.withdrawableUserStake([
        user.account.address,
        0n,
        255,
      ]);
      expect(avail2).to.be.equal(stake);
    });

    it("should use all remaining staked tokens for the opening session", async function () {
      const { sessionRouter, user, provider, expectedSession, tokenMOR } =
        await loadFixture(openEarlyCloseSession);
      await time.increaseTo(
        startOfTheDay(await nowChain()) + BigInt(DAY / SECOND),
      );

      const [avail] = await sessionRouter.read.withdrawableUserStake([
        user.account.address,
        0n,
        255,
      ]);
      expect(avail > 0).to.be.true;

      // reset allowance
      await tokenMOR.write.approve([sessionRouter.address, 0n], {
        account: user.account.address,
      });

      const approval = await getProviderApproval(
        provider,
        user.account.address,
        expectedSession.bidID,
      );
      const openSession = await sessionRouter.simulate.openSession(
        [avail, approval.msg, approval.signature],
        { account: user.account.address },
      );
      await user.writeContract(openSession.request);

      const [avail2] = await sessionRouter.read.withdrawableUserStake([
        user.account.address,
        0n,
        255,
      ]);
      expect(avail2).to.be.equal(0n);
    });

    it("should use remaining staked tokens and allowance for opening session", async function () {
      const { sessionRouter, user, provider, expectedSession, tokenMOR } =
        await loadFixture(openEarlyCloseSession);
      await time.increaseTo(
        startOfTheDay(await nowChain()) + BigInt(DAY / SECOND),
      );

      const [avail] = await sessionRouter.read.withdrawableUserStake([
        user.account.address,
        0n,
        255,
      ]);
      expect(avail > 0).to.be.true;

      const allowancePart = 1000n;
      const balanceBefore = await tokenMOR.read.balanceOf([
        user.account.address,
      ]);

      // reset allowance
      await tokenMOR.write.approve([sessionRouter.address, allowancePart], {
        account: user.account.address,
      });

      const approval = await getProviderApproval(
        provider,
        user.account.address,
        expectedSession.bidID,
      );
      const openSession = await sessionRouter.simulate.openSession(
        [avail + allowancePart, approval.msg, approval.signature],
        { account: user.account.address },
      );
      await user.writeContract(openSession.request);

      // check all onHold used
      const [avail2] = await sessionRouter.read.withdrawableUserStake([
        user.account.address,
        0n,
        255,
      ]);
      expect(avail2).to.be.equal(0n);

      // check allowance used
      const balanceAfter = await tokenMOR.read.balanceOf([
        user.account.address,
      ]);
      expect(balanceBefore - balanceAfter).to.be.equal(allowancePart);
    });
  });

  describe("negative cases", function () {
    it("should error when approval generated for a different user", async function () {
      const {
        sessionRouter,
        expectedSession: exp,
        user,
        provider,
      } = await loadFixture(deploySingleBid);

      const { msg, signature } = await getProviderApproval(
        provider,
        randomAddress(),
        exp.bidID,
      );
      await catchError(
        sessionRouter.abi,
        "ApprovedForAnotherUser",
        async () => {
          await sessionRouter.write.openSession([exp.stake, msg, signature], {
            account: user.account,
          });
        },
      );
    });

    it("should error when approval bytes is invalid abi data", async function () {
      const {
        sessionRouter,
        expectedSession: exp,
        user,
      } = await loadFixture(deploySingleBid);
      try {
        await sessionRouter.write.openSession([exp.stake, "0x0", "0x0"], {
          account: user.account,
        });
      } catch (e) {
        expect((e as Error).cause).to.be.an.instanceOf(UnknownRpcError);
      }
    });

    it("should error when approval expired", async function () {
      const {
        sessionRouter,
        expectedSession: exp,
        user,
        provider,
      } = await loadFixture(deploySingleBid);

      const { msg, signature } = await getProviderApproval(
        provider,
        user.account.address,
        exp.bidID,
      );
      const ttl = await sessionRouter.read.SIGNATURE_TTL();
      await time.increase(ttl + 1);

      await catchError(sessionRouter.abi, "SignatureExpired", async () => {
        await sessionRouter.write.openSession([exp.stake, msg, signature], {
          account: user.account,
        });
      });
    });

    it("should error when bid not exist", async function () {
      const {
        sessionRouter,
        user,
        expectedSession: exp,
        provider,
      } = await loadFixture(deploySingleBid);

      const { msg, signature } = await getProviderApproval(
        provider,
        user.account.address,
        randomBytes32(),
      );
      await catchError(sessionRouter.abi, "BidNotFound", async () => {
        await sessionRouter.write.openSession([exp.stake, msg, signature], {
          account: user.account,
        });
      });
    });

    it("should error when bid is deleted", async function () {
      const {
        sessionRouter,
        user,
        expectedSession: exp,
        marketplace,
        provider,
      } = await loadFixture(deploySingleBid);

      await marketplace.write.deleteModelAgentBid([exp.bidID], {
        account: provider.account,
      });

      const { msg, signature } = await getProviderApproval(
        provider,
        user.account.address,
        exp.bidID,
      );
      await catchError(sessionRouter.abi, "BidNotFound", async () => {
        await sessionRouter.write.openSession([exp.stake, msg, signature], {
          account: user.account,
        });
      });
    });

    it("should error when signature has invalid length", async function () {
      const {
        sessionRouter,
        user,
        expectedSession: exp,
        provider,
      } = await loadFixture(deploySingleBid);

      const { msg } = await getProviderApproval(
        provider,
        user.account.address,
        exp.bidID,
      );
      await catchError(
        sessionRouter.abi,
        "ECDSAInvalidSignatureLength",
        async () => {
          await sessionRouter.write.openSession([exp.stake, msg, "0x0"], {
            account: user.account,
          });
        },
      );
    });

    it("should error when signature is invalid", async function () {
      const {
        sessionRouter,
        user,
        expectedSession: exp,
        provider,
      } = await loadFixture(deploySingleBid);

      const { msg } = await getProviderApproval(
        provider,
        user.account.address,
        exp.bidID,
      );
      const sig = randomBytes(65);

      await catchError(
        sessionRouter.abi,
        ["ECDSAInvalidSignatureS", "ECDSAInvalidSignature"],
        async () => {
          await sessionRouter.write.openSession([exp.stake, msg, sig], {
            account: user.account,
          });
        },
      );
    });

    it("should error when opening two bids with same signature", async function () {
      const {
        sessionRouter,
        expectedSession: exp,
        user,
        provider,
        approveUserFunds,
      } = await loadFixture(deploySingleBid);

      const { msg, signature } = await getProviderApproval(
        provider,
        user.account.address,
        exp.bidID,
      );
      await sessionRouter.write.openSession([exp.stake, msg, signature], {
        account: user.account.address,
      });

      await approveUserFunds(exp.stake);

      await catchError(sessionRouter.abi, "DuplicateApproval", async () => {
        await sessionRouter.write.openSession([exp.stake, msg, signature], {
          account: user.account.address,
        });
      });
    });

    it("should not error when opening two bids same time", async function () {
      const {
        sessionRouter,
        expectedSession: exp,
        user,
        provider,
        approveUserFunds,
      } = await loadFixture(deploySingleBid);

      const appr1 = await getProviderApproval(
        provider,
        user.account.address,
        exp.bidID,
      );
      await sessionRouter.write.openSession(
        [exp.stake, appr1.msg, appr1.signature],
        {
          account: user.account.address,
        },
      );

      await approveUserFunds(exp.stake);
      const appr2 = await getProviderApproval(
        provider,
        user.account.address,
        exp.bidID,
      );
      await sessionRouter.write.openSession(
        [exp.stake, appr2.msg, appr2.signature],
        {
          account: user.account.address,
        },
      );
    });

    it("should error with insufficient allowance", async function () {
      const {
        sessionRouter,
        expectedSession: exp,
        user,
        provider,
        tokenMOR,
      } = await loadFixture(deploySingleBid);

      const { msg, signature } = await getProviderApproval(
        provider,
        user.account.address,
        exp.bidID,
      );
      await catchError(tokenMOR.abi, "ERC20InsufficientAllowance", async () => {
        await sessionRouter.write.openSession(
          [exp.stake * 2n, msg, signature],
          {
            account: user.account,
          },
        );
      });
    });

    it("should error with insufficient allowance", async function () {
      const {
        sessionRouter,
        expectedSession: exp,
        user,
        provider,
        tokenMOR,
      } = await loadFixture(deploySingleBid);

      const stake =
        (await tokenMOR.read.balanceOf([user.account.address])) + 1n;
      await tokenMOR.write.approve([sessionRouter.address, stake], {
        account: user.account,
      });

      const { msg, signature } = await getProviderApproval(
        provider,
        user.account.address,
        exp.bidID,
      );
      await catchError(tokenMOR.abi, "ERC20InsufficientBalance", async () => {
        await sessionRouter.write.openSession([stake, msg, signature], {
          account: user.account,
        });
      });
    });
  });
});

describe("verify session end time", function () {
  it("session that doesn't span across midnight (1h)", async function () {
    const {
      sessionRouter,
      expectedSession: exp,
      user,
      publicClient,
      getStake,
      provider,
    } = await loadFixture(deploySingleBid);

    const durationSeconds = BigInt(HOUR / SECOND);
    const stake = await getStake(durationSeconds, exp.pricePerSecond);

    const { msg, signature } = await getProviderApproval(
      provider,
      user.account.address,
      exp.bidID,
    );
    const txHash = await sessionRouter.write.openSession(
      [stake, msg, signature],
      { account: user.account },
    );

    const sessionId = await getSessionId(publicClient, hre, txHash);
    const session = await sessionRouter.read.getSession([sessionId]);
    const createdAt = await getTxTimestamp(publicClient, txHash);

    expect(session.endsAt).to.equal(createdAt + exp.durationSeconds);
  });

  it("session that spans across midnight (6h) should last 6h", async function () {
    const {
      sessionRouter,
      expectedSession: exp,
      user,
      publicClient,
      getStake,
      approveUserFunds,
      provider,
    } = await loadFixture(deploySingleBid);

    const tomorrow9pm =
      startOfTheDay(await nowChain()) +
      BigInt(DAY / SECOND) +
      21n * BigInt(HOUR / SECOND);
    await time.increaseTo(tomorrow9pm); // 9pm

    // the stake is enough to cover the first day (3h till midnight) and the next day (< 6h)
    const durationSeconds = 6n * BigInt(HOUR / SECOND);
    const stake = await getStake(durationSeconds, exp.pricePerSecond);
    await approveUserFunds(stake);

    const { msg, signature } = await getProviderApproval(
      provider,
      user.account.address,
      exp.bidID,
    );
    const txHash = await sessionRouter.write.openSession(
      [stake, msg, signature],
      { account: user.account },
    );

    await time.increase((3 * HOUR) / SECOND + 1);

    const sessionId = await getSessionId(publicClient, hre, txHash);
    const session = await sessionRouter.read.getSession([sessionId]);
    const createdAt = await getTxTimestamp(publicClient, txHash);

    const endsAt = NewDate(session.endsAt);
    const expEndsAt = NewDate(createdAt + durationSeconds);

    expect(endsAt.getTime()).approximately(expEndsAt.getTime(), 10 * SECOND);
  });

  it("session that lasts multiple days", async function () {
    const {
      sessionRouter,
      expectedSession: exp,
      user,
      publicClient,
      approveUserFunds,
      provider,
    } = await loadFixture(deploySingleBid);

    const midnight = startOfTheDay(await nowChain()) + BigInt(DAY / SECOND);
    await time.increaseTo(midnight);

    // the stake is enough to cover the whole day + extra 1h
    const durationSeconds = 25n * BigInt(HOUR / SECOND);
    const stake = await sessionRouter.read.stipendToStake([
      durationSeconds * exp.pricePerSecond,
      await nowChain(),
    ]);

    await approveUserFunds(stake);

    const { msg, signature } = await getProviderApproval(
      provider,
      user.account.address,
      exp.bidID,
    );
    const txHash = await sessionRouter.write.openSession(
      [stake, msg, signature],
      { account: user.account },
    );

    const sessionId = await getSessionId(publicClient, hre, txHash);
    const session = await sessionRouter.read.getSession([sessionId]);
    const durSeconds = Number(session.endsAt - session.openedAt);

    expect(durSeconds).to.equal(1 * (DAY / SECOND));
  });
});
