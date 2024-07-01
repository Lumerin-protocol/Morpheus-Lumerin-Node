import hre from "hardhat";
import {
  encodeAbiParameters,
  encodeFunctionData,
  encodePacked,
  getAddress,
  keccak256,
  parseUnits,
} from "viem/utils";
import { getHex, getTxTimestamp, now, nowChain, randomBytes32 } from "./utils";
import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { FacetCutAction, getSelectors } from "../libraries/diamond";
import { DAY, HOUR, MINUTE, SECOND } from "../utils/time";
import {
  GetContractReturnType,
  WalletClient,
} from "@nomicfoundation/hardhat-viem/types";

export async function deployMORtoken() {
  const [owner] = await hre.viem.getWalletClients();

  // Contracts are deployed using the first signer/account by default
  const tokenMOR = await hre.viem.deployContract("MorpheusToken", []);
  const decimalsMOR = await tokenMOR.read.decimals();

  return {
    tokenMOR,
    decimalsMOR,
    owner,
  };
}

export async function deployDiamond() {
  // deploy provider registry and deps
  const { tokenMOR, owner, decimalsMOR } = await loadFixture(deployMORtoken);
  const [_, provider, user] = await hre.viem.getWalletClients();
  const publicClient = await hre.viem.getPublicClient();

  const { diamond } = await onlyDeployDiamond(tokenMOR.address, owner);

  const modelRegistry = await hre.viem.getContractAt(
    "contracts/facets/ModelRegistry.sol:ModelRegistry",
    diamond.address,
  );
  const providerRegistry = await hre.viem.getContractAt(
    "contracts/facets/ProviderRegistry.sol:ProviderRegistry",
    diamond.address,
  );
  const marketplace = await hre.viem.getContractAt(
    "contracts/facets/Marketplace.sol:Marketplace",
    diamond.address,
  );
  const sessionRouter = await hre.viem.getContractAt(
    "contracts/facets/SessionRouter.sol:SessionRouter",
    diamond.address,
  );

  return {
    tokenMOR,
    decimalsMOR,
    diamond,
    owner,
    user,
    provider,
    publicClient,
    modelRegistry,
    providerRegistry,
    marketplace,
    sessionRouter,
  };
}

export async function onlyDeployDiamond(
  morAddress: string,
  owner: WalletClient,
) {
  // 1. deploy diamont init
  const diamondInit = await hre.viem.deployContract("DiamondInit", [], {});
  console.log("diamond init deployed at address", diamondInit.address);

  // 2. deploy faucets
  const FacetNames = [
    "DiamondCutFacet",
    "DiamondLoupeFacet",
    "OwnershipFacet",
    "ModelRegistry",
    "ProviderRegistry",
    "Marketplace",
    "SessionRouter",
  ] as const;

  const facetContracts: GetContractReturnType[] = [];
  for (const name of FacetNames) {
    try {
      const data = await hre.viem.deployContract(name as any, [], {});
      console.log("faucet ", name, " deployed at address", data.address);
      facetContracts.push(data);
    } catch (e) {
      console.log(`error deploying ${name}`);
      throw e;
    }
  }

  // 3. deploy diamond
  const facetCuts = facetContracts.map((facetContract) => ({
    facetAddress: facetContract.address,
    action: FacetCutAction.Add,
    functionSelectors: getSelectors(facetContract.abi),
  }));

  const diamondArgs = {
    owner: owner.account.address,
    init: diamondInit.address,
    initCalldata: encodeFunctionData({
      abi: hre.artifacts.readArtifactSync("DiamondInit").abi,
      functionName: "init",
      args: [morAddress as any, owner.account.address],
    }),
  };

  const diamond = await hre.viem.deployContract("Diamond", [
    facetCuts,
    diamondArgs,
  ]);

  return {
    diamond,
    facets: facetContracts.map((f) => ({
      name: f.constructor.name,
      address: f.address,
    })),
    constructorArgs: [facetCuts, diamondArgs],
    owner,
  };
}

export async function deploySingleProvider() {
  const {
    sessionRouter,
    providerRegistry,
    owner,
    provider,
    publicClient,
    decimalsMOR,
    tokenMOR,
    modelRegistry,
    user,
    marketplace,
  } = await loadFixture(deployDiamond);

  const expectedProvider = {
    address: getAddress(provider.account.address),
    stake: parseUnits("100", decimalsMOR),
    endpoint: "localhost:3334",
    createdAt: 0n,
    limitPeriodEarned: 0n,
    limitPeriodEnd: 0n,
    isDeleted: false,
  };

  await tokenMOR.write.transfer(
    [provider.account.address, expectedProvider.stake * 100n],
    {
      account: owner.account,
    },
  );

  await tokenMOR.write.approve(
    [sessionRouter.address, expectedProvider.stake],
    {
      account: provider.account,
    },
  );

  const addProviderHash = await providerRegistry.write.providerRegister(
    [
      expectedProvider.address,
      expectedProvider.stake,
      expectedProvider.endpoint,
    ],
    { account: provider.account },
  );
  expectedProvider.createdAt = await getTxTimestamp(
    publicClient,
    addProviderHash,
  );
  expectedProvider.limitPeriodEnd =
    expectedProvider.createdAt + BigInt(365 * (DAY / SECOND));

  return {
    expectedProvider,
    providerRegistry,
    modelRegistry,
    sessionRouter,
    marketplace,
    owner,
    provider,
    user,
    publicClient,
    tokenMOR,
    decimalsMOR,
  };
}

export async function deploySingleModel() {
  const { owner, provider, publicClient, tokenMOR, modelRegistry } =
    await loadFixture(deployDiamond);

  const expectedModel = {
    modelId: randomBytes32(),
    ipfsCID: getHex(Buffer.from("ipfs://ipfsaddress")),
    fee: 100n,
    stake: 100n,
    owner: owner.account.address,
    name: "Llama 2.0",
    createdAt: 0n,
    tags: ["llama", "animal", "cute"],
    isDeleted: false,
  };

  await tokenMOR.write.approve([modelRegistry.address, expectedModel.stake]);

  const addProviderHash = await modelRegistry.write.modelRegister([
    expectedModel.modelId,
    expectedModel.ipfsCID,
    expectedModel.fee,
    expectedModel.stake,
    expectedModel.owner,
    expectedModel.name,
    expectedModel.tags,
  ]);

  expectedModel.createdAt = await getTxTimestamp(publicClient, addProviderHash);

  return {
    expectedModel,
    modelRegistry,
    owner,
    provider,
    publicClient,
    tokenMOR,
  };
}
export async function deploySingleBid() {
  const {
    owner,
    provider,
    publicClient,
    tokenMOR,
    modelRegistry,
    providerRegistry,
    user,
    decimalsMOR,
    marketplace,
    sessionRouter,
    expectedProvider,
  } = await loadFixture(deploySingleProvider);

  // add single model
  const expectedModel = {
    modelId: randomBytes32(),
    ipfsCID: getHex(Buffer.from("ipfs://ipfsaddress")),
    fee: 100n,
    stake: 100n,
    owner: owner.account.address,
    name: "Llama 2.0",
    timestamp: 0n,
    tags: ["llama", "animal", "cute"],
    isDeleted: false,
  };

  await tokenMOR.write.approve([modelRegistry.address, expectedModel.stake]);
  const addProviderHash = await modelRegistry.write.modelRegister([
    expectedModel.modelId,
    expectedModel.ipfsCID,
    expectedModel.fee,
    expectedModel.stake,
    expectedModel.owner,
    expectedModel.name,
    expectedModel.tags,
  ]);

  expectedModel.timestamp = await getTxTimestamp(publicClient, addProviderHash);

  // expected bid
  const expectedBid = {
    id: "" as `0x${string}`,
    providerAddr: getAddress(expectedProvider.address),
    modelId: expectedModel.modelId,
    pricePerSecond: parseUnits("0.0001", decimalsMOR),
    nonce: 0n,
    createdAt: 0n,
    deletedAt: 0n,
  };

  await tokenMOR.write.approve([modelRegistry.address, 10000n * 10n ** 18n]);

  // add single bid
  const postBidtx = await marketplace.simulate.postModelBid(
    [expectedBid.providerAddr, expectedBid.modelId, expectedBid.pricePerSecond],
    { account: provider.account.address },
  );
  const txHash = await provider.writeContract(postBidtx.request);

  expectedBid.id = postBidtx.result;
  expectedBid.createdAt = await getTxTimestamp(publicClient, txHash);

  // generating data for sample session
  const durationSeconds = BigInt(HOUR / SECOND);
  const totalCost = expectedBid.pricePerSecond * durationSeconds;
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

  async function getStake(
    durationSeconds: bigint,
    pricePerSecond: bigint,
  ): Promise<bigint> {
    const totalCost = pricePerSecond * durationSeconds;
    const totalSupply = await sessionRouter.read.totalMORSupply([
      await nowChain(),
    ]);
    const todaysBudget = await sessionRouter.read.getTodaysBudget([
      await nowChain(),
    ]);
    return (totalCost * totalSupply) / todaysBudget;
  }

  async function getDuration(
    stake: bigint,
    pricePerSecond: bigint,
  ): Promise<bigint> {
    const totalSupply = await sessionRouter.read.totalMORSupply([
      await nowChain(),
    ]);
    const todaysBudget = await sessionRouter.read.getTodaysBudget([
      await nowChain(),
    ]);
    const totalCost = (stake * todaysBudget) / totalSupply;
    return totalCost / pricePerSecond;
  }

  async function approveUserFunds(amount: bigint) {
    await tokenMOR.write.transfer([user.account.address, amount]);
    await tokenMOR.write.approve([modelRegistry.address, amount], {
      account: user.account,
    });
  }

  await approveUserFunds(expectedSession.stake);

  return {
    expectedBid,
    expectedSession,
    expectedProvider,
    expectedModel,
    marketplace,
    modelRegistry,
    providerRegistry,
    owner,
    provider,
    publicClient,
    tokenMOR,
    decimalsMOR,
    sessionRouter,
    getStake,
    getDuration,
    approveUserFunds,
    user,
  };
}

export const providerReport = {
  ips: 128,
  timestamp: 10000,
};
export const reportAbi = [
  { type: "bytes32" },
  { type: "uint128" },
  { type: "uint32" },
];

export const approvalAbi = [{ type: "bytes32" }, { type: "uint128" }];

export const getProviderApproval = async (
  provider: WalletClient,
  bidId: `0x${string}`,
) => {
  const timestampMs = (await time.latest()) * 1000;
  const msg = encodeAbiParameters(approvalAbi, [bidId, BigInt(timestampMs)]);
  const signature = await provider.signMessage({
    message: { raw: keccak256(msg) },
  });

  return {
    msg,
    signature,
  };
};

export const getReport = async (
  reporter: WalletClient,
  sessionId: `0x${string}`,
  ips: number,
) => {
  const timestampMs = (await time.latest()) * 1000;
  const msg = encodeAbiParameters(reportAbi, [sessionId, timestampMs, ips]);
  const sig = await reporter.signMessage({
    message: { raw: keccak256(msg) },
  });

  return {
    msg,
    sig,
  };
};
