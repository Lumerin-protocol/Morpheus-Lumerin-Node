import { ISessionStorage } from '../../generated-types/ethers/contracts/interfaces/facets/ISessionRouter';

export const payoutStart = 1707393600;

export function getDefaultPools(): ISessionStorage.PoolStruct[] {
  return [
    {
      payoutStart: payoutStart,
      decreaseInterval: 86400n,
      initialReward: 3456000000000000000000n,
      rewardDecrease: 592558728240000000n,
    },
    {
      payoutStart: 1707393600n,
      decreaseInterval: 86400n,
      initialReward: 3456000000000000000000n,
      rewardDecrease: 592558728240000000n,
    },
    {
      payoutStart: payoutStart,
      decreaseInterval: 86400n,
      initialReward: 3456000000000000000000n,
      rewardDecrease: 592558728240000000n,
    },
    {
      payoutStart: payoutStart,
      decreaseInterval: 86400n,
      initialReward: 3456000000000000000000n,
      rewardDecrease: 592558728240000000n,
    },
    {
      payoutStart: payoutStart,
      decreaseInterval: 86400n,
      initialReward: 576000000000000000000n,
      rewardDecrease: 98759788040000000n,
    },
  ];
}
