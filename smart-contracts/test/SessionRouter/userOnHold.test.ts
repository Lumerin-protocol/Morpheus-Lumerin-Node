import { expect } from "chai";
import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { openEarlyCloseSession } from "../fixtures";
import { DAY, HOUR, SECOND } from "../../utils/time";
import { catchError, randomAddress } from "../utils";
import { expectAlmostEqual } from "../../utils/compare";
import { maxUint8 } from "viem";

describe("User on hold tests", () => {
  it("user stake should be locked right after closeout", async () => {
    const { sessionRouter, user, expectedOnHold } = await loadFixture(
      openEarlyCloseSession,
    );

    // right after closeout
    const [available, onHold] = await sessionRouter.read.withdrawableUserStake([
      user.account.address,
      0n,
      Number(maxUint8),
    ]);
    expect(available).to.equal(0n);
    expectAlmostEqual(onHold, expectedOnHold, 0.01);
  });

  it("user stake should be locked before the next day", async () => {
    const { sessionRouter, user, expectedOnHold } = await loadFixture(
      openEarlyCloseSession,
    );

    // before next day
    await time.increaseTo(startOfTomorrow(await time.latest()) - HOUR / SECOND);
    const [available3, onHold3] =
      await sessionRouter.read.withdrawableUserStake([
        user.account.address,
        0n,
        Number(maxUint8),
      ]);
    expect(available3).to.equal(0n);
    expectAlmostEqual(onHold3, expectedOnHold, 0.01);
  });

  it("user stake should be available on the next day and withdrawable", async () => {
    const { sessionRouter, user, expectedOnHold, tokenMOR } = await loadFixture(
      openEarlyCloseSession,
    );

    await time.increaseTo(startOfTomorrow(await time.latest()));
    const [available2, onHold2] =
      await sessionRouter.read.withdrawableUserStake([
        user.account.address,
        0n,
        Number(maxUint8),
      ]);
    expectAlmostEqual(available2, expectedOnHold, 0.01);
    expect(onHold2).to.equal(0n);

    const balanceBefore = await tokenMOR.read.balanceOf([user.account.address]);
    await sessionRouter.write.withdrawUserStake(
      [available2, Number(maxUint8)],
      { account: user.account.address },
    );
    const balanceAfter = await tokenMOR.read.balanceOf([user.account.address]);
    const balanceDelta = balanceAfter - balanceBefore;
    expectAlmostEqual(balanceDelta, expectedOnHold, 0.01);
  });

  it("user shouldn't be able to withdraw more than there is available stake", async () => {
    const { sessionRouter, user } = await loadFixture(openEarlyCloseSession);

    await time.increaseTo(startOfTomorrow(await time.latest()));
    const [available2] = await sessionRouter.read.withdrawableUserStake([
      user.account.address,
      0n,
      Number(maxUint8),
    ]);

    // check that user can't withdraw twice
    await catchError(
      sessionRouter.abi,
      "NotEnoughWithdrawableBalance",
      async () => {
        await sessionRouter.write.withdrawUserStake(
          [available2 * 2n, Number(maxUint8)],
          { account: user.account.address },
        );
      },
    );
  });
});

function startOfTomorrow(epochSeconds: number): number {
  const startOfToday = epochSeconds - (epochSeconds % (DAY / SECOND));
  return startOfToday + DAY / SECOND;
}
