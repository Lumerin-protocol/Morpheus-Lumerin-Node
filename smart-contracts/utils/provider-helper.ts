import { SignerWithAddress } from '@nomicfoundation/hardhat-ethers/signers';
import { AbiCoder, getBytes, keccak256 } from 'ethers';

import { getChainId, getCurrentBlockTime } from './block-helper';

export const getProviderApproval = async (
  provider: SignerWithAddress,
  user: SignerWithAddress,
  bidId: string,
  chainId = 0n,
) => {
  chainId = chainId || (await getChainId());
  const timestamp = await getCurrentBlockTime();

  const msg = AbiCoder.defaultAbiCoder().encode(
    ['bytes32', 'uint256', 'address', 'uint128'],
    [bidId, chainId, user.address, timestamp],
  );

  const signature = await provider.signMessage(getBytes(keccak256(msg)));

  return {
    msg,
    signature,
  };
};

export const getReceipt = async (
  reporter: SignerWithAddress,
  sessionId: string,
  tps: number,
  ttftMs: number,
  chainId = 0n,
) => {
  chainId = chainId || (await getChainId());
  const timestamp = await getCurrentBlockTime();

  const msg = AbiCoder.defaultAbiCoder().encode(
    ['bytes32', 'uint256', 'uint128', 'uint32', 'uint32'],
    [sessionId, chainId, timestamp, tps * 1000, ttftMs],
  );
  const signature = await reporter.signMessage(getBytes(keccak256(msg)));

  return {
    msg,
    signature,
  };
};
