import { PublicClient } from "@nomicfoundation/hardhat-viem/types";
import {
  Abi,
  BaseError,
  ContractFunctionRevertedError,
  UnknownRpcError,
  parseEventLogs,
} from "viem";
import {
  DecodeErrorResultReturnType,
  decodeErrorResult,
  padHex,
} from "viem/utils";
import crypto from "crypto";
import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DAY, SECOND } from "../utils/time";
import { time } from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";

export async function getTxTimestamp(
  client: PublicClient,
  txHash: `0x${string}`,
): Promise<bigint> {
  const receipt = await client.waitForTransactionReceipt({
    hash: txHash,
    timeout: 1000,
  });
  const block = await client.getBlock({ blockNumber: receipt.blockNumber });
  return block.timestamp;
}

export async function getSessionId(
  publicClient: PublicClient,
  hre: HardhatRuntimeEnvironment,
  txHash: `0x${string}`,
): Promise<`0x${string}`> {
  const receipt = await publicClient.waitForTransactionReceipt({
    hash: txHash,
  });
  const artifact = await hre.artifacts.readArtifact("SessionRouter");
  const events = parseEventLogs({
    abi: artifact.abi,
    logs: receipt.logs,
    eventName: "SessionOpened",
  });
  if (events.length === 0) {
    throw new Error("SessionOpened event not found");
  }
  if (events.length > 1) {
    throw new Error("Multiple SessionOpened events found");
  }
  return events[0].args.sessionId;
}

/** helper function to catch errors and check if the error is the expected one
 * @example
 * await catchError(abi, "ErrorName", async () => {
 *   await contract.method();
 * });
 **/
export async function catchError<const TAbi extends Abi | readonly unknown[]>(
  abi: TAbi | undefined,
  error:
    | DecodeErrorResultReturnType<TAbi>["errorName"]
    | DecodeErrorResultReturnType<TAbi>["errorName"][],
  cb: () => Promise<unknown>,
) {
  try {
    await cb();
    throw new Error(`No error was thrown, expected error "${error}"`);
  } catch (err) {
    if (Array.isArray(error)) {
      return expectError(err, abi, error);
    } else {
      return expectError(err, abi, [error]);
    }
  }
}

export function expectError<const TAbi extends Abi | readonly unknown[]>(
  err: any,
  abi: TAbi | undefined,
  errors: DecodeErrorResultReturnType<TAbi>["errorName"][],
) {
  for (const error of errors) {
    if (isErr(err, abi, error)) {
      return;
    }
  }

  console.error(err);
  throw new Error(
    `Expected one of blockchain custom errors "${errors.join(" | ")}" was not thrown\n\n${err}`,
    { cause: err },
  );
}

export function isErr<const TAbi extends Abi | readonly unknown[]>(
  err: any,
  abi: TAbi | undefined,
  error: DecodeErrorResultReturnType<TAbi>["errorName"],
): boolean {
  if (err instanceof BaseError) {
    const revertError = err.walk(
      (err) =>
        err instanceof ContractFunctionRevertedError ||
        err instanceof UnknownRpcError,
    );

    // support for regular provider
    if (revertError instanceof ContractFunctionRevertedError) {
      const errorName = revertError.data?.errorName ?? "";
      if (errorName === error) {
        return true;
      }
    }

    // support for hardhat node
    if (revertError instanceof UnknownRpcError) {
      const cause = revertError.cause as any;
      if (cause.data) {
        try {
          const decodedError = decodeErrorResult({ abi, data: cause.data });
          if (decodedError.errorName === error) {
            return true;
          }
        } catch (e) {
          console.error(e);
          return false;
        }
      }
    }
  }
  return false;
}

export const getHex = (buffer: Buffer, padding = 32): `0x${string}` => {
  return padHex(`0x${buffer.toString("hex")}`, { size: padding });
};

export const randomBytes32 = (): `0x${string}` => {
  return getHex(crypto.randomBytes(32));
};

export const randomBytes = (nBytes: number): `0x${string}` => {
  return getHex(crypto.randomBytes(nBytes), nBytes);
};

export const randomAddress = (): `0x${string}` => {
  return getHex(crypto.randomBytes(20), 20);
};

export const now = (): bigint => {
  return BigInt(Math.floor(Date.now() / 1000));
};

export const nowChain = async (): Promise<bigint> => {
  return BigInt(await time.latest());
};

export const startOfTheDay = (timestamp: bigint): bigint => {
  return timestamp - (timestamp % BigInt(DAY / SECOND));
};

export const NewDate = (timestamp: bigint): Date => {
  return new Date(Number(timestamp) * 1000);
};

export const PanicOutOfBoundsRegexp =
  /.*reverted with panic code 0x32 (Array accessed at an out-of-bounds or negative index)*/;
