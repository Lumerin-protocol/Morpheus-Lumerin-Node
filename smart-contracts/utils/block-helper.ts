import { ethers } from 'hardhat';

export async function getCurrentBlockTime() {
  return BigInt(await ethers.provider.getBlock('latest').then((block) => block!.timestamp));
}

export async function setNextTime(time: number) {
  await ethers.provider.send('evm_setNextBlockTimestamp', [time]);
}

export async function setTime(time: number) {
  await setNextTime(time);
  await ethers.provider.send('evm_mine', []);
}

export async function getChainId() {
  return (await ethers.provider.getNetwork()).chainId;
}
