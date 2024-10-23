import crypto from 'crypto';
import { hexlify, zeroPadBytes } from 'ethers';
import { ethers } from 'hardhat';

import { DAY, SECOND } from '@/utils/time';

export function wei(value: string | number | bigint, decimal: number = 18): bigint {
  if (typeof value == 'number' || typeof value == 'bigint') {
    value = value.toString();
  }

  return ethers.parseUnits(value as string, decimal);
}

export function fromWei(value: string | number | bigint, decimal: number = 18): string {
  return (BigInt(value) / 10n ** BigInt(decimal)).toString();
}

export const getHex = (buffer: Buffer, padding = 32): string => {
  return zeroPadBytes(`0x${buffer.toString('hex')}`, padding);
};

export const randomBytes32 = (): string => {
  return getHex(crypto.randomBytes(32));
};

export const randomBytes = (nBytes: number): string => {
  return getHex(crypto.randomBytes(nBytes), nBytes);
};

export const startOfTheDay = (timestamp: bigint): bigint => {
  return timestamp - (timestamp % BigInt(DAY / SECOND));
};
