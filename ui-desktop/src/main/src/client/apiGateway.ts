import config from '../../config'
import fs from 'fs'
import {
  AgentAllowanceRequestsRes,
  AgentTxRes,
  AgentUserRes,
  ChatHistory,
  ChatTitle,
  ResultResponse
} from './api.types'

import os from 'os'

let auth: Record<string, string> | null = null

const getAuthHeaders = async () => {
  if (auth) {
    return auth
  }

  try {
    const path = `${config.chain.localProxyRouterUrl}/auth/cookie/path`
    const response = await fetch(path)
    const body = await response.json()
    let cookieFilePath = body.path

    const isWindows = os.platform() === 'win32'
    cookieFilePath = isWindows ? cookieFilePath.replace(/\//g, '\\') : cookieFilePath

    const cookieFile = fs.readFileSync(cookieFilePath, 'utf8').trim()
    const [username, password] = cookieFile.split(':')
    auth = {
      Authorization: `Basic ${Buffer.from(`${username}:${password}`, 'utf-8').toString('base64')}`
    }
    return auth
  } catch (e) {
    console.log('Error', e)
    throw e
  }
}

const getAllModels = async (): Promise<unknown[]> => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/blockchain/models`
    const response = await fetch(path, {
      headers: await getAuthHeaders(),
      method: 'GET'
    })
    const data = await response.json()
    return data.models
  } catch (e) {
    console.log('Error', e)
    return []
  }
}

const getBalances = async (): Promise<unknown[]> => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/blockchain/balance`
    const response = await fetch(path, {
      headers: await getAuthHeaders()
    })
    const data = await response.json()
    return data
  } catch (e) {
    console.log('Error', e)
    return []
  }
}

const sendEth = async (to: string, amount: string): Promise<string | undefined> => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/blockchain/send/eth`
    const response = await fetch(path, {
      method: 'POST',
      body: JSON.stringify({
        to,
        amount
      }),
      headers: await getAuthHeaders()
    })
    const data = await response.json()
    return data.txHash
  } catch (e) {
    console.log('Error', e)
    return undefined
  }
}

const sendMor = async (to: string, amount: string): Promise<string | undefined> => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/blockchain/send/mor`
    const response = await fetch(path, {
      method: 'POST',
      body: JSON.stringify({
        to,
        amount
      }),
      headers: await getAuthHeaders()
    })
    const data = await response.json()
    return data.txHash
  } catch (e) {
    console.log('Error', e)
    return undefined
  }
}

const getTransactions = async (payload: { page: number; pageSize: number }): Promise<unknown[]> => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/blockchain/transactions?page=${payload.page}&limit=${payload.pageSize}`
    const response = await fetch(path, {
      headers: await getAuthHeaders()
    })
    const data = await response.json()
    return data.transactions
  } catch (e) {
    console.log('Error', e)
    return []
  }
}

const getMorRate = async (
  tokenAddress = '0x092baadb7def4c3981454dd9c0a0d7ff07bcfc86',
  network = 'arbitrum'
) => {
  try {
    const path = `https://api.geckoterminal.com/api/v2/simple/networks/${network}/token_price/${tokenAddress}`
    const response = await fetch(path)
    const body = await response.json()
    return body.data.attributes.token_prices[tokenAddress]
  } catch (e) {
    console.log('Error', e)
    return null
  }
}

const getTodaysBudget = async () => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/blockchain/sessions/budget`
    const response = await fetch(path, {
      headers: await getAuthHeaders()
    })
    const body = await response.json()
    return body.budget
  } catch (e) {
    console.log('Error', e)
    return null
  }
}

const getTokenSupply = async () => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/blockchain/token/supply`
    const response = await fetch(path, {
      headers: await getAuthHeaders()
    })
    const body = await response.json()
    return body.supply
  } catch (e) {
    console.log('Error', e)
    return null
  }
}

const getChatHistoryTitles = async (): Promise<ChatTitle[] | null> => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/v1/chats`
    const response = await fetch(path, {
      headers: await getAuthHeaders()
    })
    const body = await response.json()
    return body
  } catch (e) {
    console.log('Error', e)
    return null
  }
}

const getChatHistory = async (chatId: string): Promise<ChatHistory | null> => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/v1/chats/${chatId}`
    const response = await fetch(path, {
      headers: await getAuthHeaders()
    })
    const body = await response.json()
    return body
  } catch (e) {
    console.log('Error', e)
    return null
  }
}

const deleteChatHistory = async (chatId: string): Promise<boolean> => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/v1/chats/${chatId}`
    const response = await fetch(path, {
      method: 'DELETE',
      headers: await getAuthHeaders()
    })
    const body = await response.json()
    return body.result
  } catch (e) {
    console.log('Error', e)
    return false
  }
}

const updateChatHistoryTitle = async (params: { id: string; title: string }): Promise<boolean> => {
  const { id, title } = params
  try {
    const path = `${config.chain.localProxyRouterUrl}/v1/chats/${id}`
    const response = await fetch(path, {
      method: 'POST',
      body: JSON.stringify({ title }),
      headers: await getAuthHeaders()
    })
    const body = await response.json()
    return body.result
  } catch (e) {
    console.log('Error', e)
    return false
  }
}

const checkProviderConnectivity = async (params: {
  address: string
  endpoint: string
}): Promise<boolean> => {
  const { address, endpoint } = params
  try {
    const path = `${config.chain.localProxyRouterUrl}/proxy/provider/ping`
    const response = await fetch(path, {
      method: 'POST',
      body: JSON.stringify({
        providerAddr: address,
        providerUrl: endpoint
      }),
      headers: await getAuthHeaders()
    })

    if (!response.ok) {
      return false
    }

    const body = await response.json()
    return !!body.ping
  } catch (e) {
    console.log('checkProviderConnectivity: Error', e)
    return false
  }
}

const clearEthNodeEnv = async () => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/config/ethNode`
    const response = await fetch(path, { method: 'DELETE', headers: await getAuthHeaders() })
    const data = await response.json()
    return data.status
  } catch (e) {
    console.log('CLEAR ETH NODE ERROR', e)
    return false
  }
}

const clearWallet = async () => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/wallet`
    const response = await fetch(path, { method: 'DELETE', headers: await getAuthHeaders() })
    const data = await response.json()
    return data.status
  } catch (e) {
    console.log('CLEAR WALLET ERROR', e)
    return false
  }
}

const getAgentUsers = async (): Promise<AgentUserRes | null> => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/auth/users`
    const response = await fetch(path, { method: 'GET', headers: await getAuthHeaders() })
    return await response.json()
  } catch (e) {
    console.log('Error', e)
    return null
  }
}

const confirmDeclineAgentUser = async (params: {
  username: string
  confirm: boolean
}): Promise<boolean> => {
  const { username, confirm } = params
  try {
    const path = `${config.chain.localProxyRouterUrl}/auth/users/confirm`

    const res = await fetch(path, {
      method: 'POST',
      body: JSON.stringify({ username, confirm }),
      headers: await getAuthHeaders()
    })
    await res.json()
    return true
  } catch (e) {
    console.log('Error', e)
    return false
  }
}

const removeAgentUser = async (params: { username: string }): Promise<boolean> => {
  const { username } = params
  try {
    const path = `${config.chain.localProxyRouterUrl}/auth/users`
    await fetch(path, {
      method: 'DELETE',
      body: JSON.stringify({ username }),
      headers: await getAuthHeaders()
    })
    return true
  } catch (e) {
    console.log('Error', e)
    return false
  }
}

const getAgentTxs = async (params: {
  username: string
  cursor: string
  limit: number
}): Promise<AgentTxRes | null> => {
  try {
    const query = new URLSearchParams()
    query.set('cursor', params.cursor)
    query.set('limit', params.limit.toString())

    const path = `${config.chain.localProxyRouterUrl}/auth/users/${encodeURIComponent(params.username)}/txs?${query.toString()}`
    const response = await fetch(path, {
      headers: await getAuthHeaders()
    })
    if (response.ok) {
      return await response.json()
    }
    throw new Error(await response.text())
  } catch (e) {
    console.log('Error', e)
    return null
  }
}

const revokeAgentAllowance = async (params: {
  username: string
  token: string
}): Promise<boolean> => {
  const { username, token } = params
  try {
    const path = `${config.chain.localProxyRouterUrl}/auth/allowance/revoke`
    await fetch(path, {
      method: 'POST',
      body: JSON.stringify({ username, token }),
      headers: await getAuthHeaders()
    })
    return true
  } catch (e) {
    console.log('Error', e)
    return false
  }
}

const getAgentAllowanceRequests = async (): Promise<AgentAllowanceRequestsRes | null> => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/auth/allowance/requests`
    const response = await fetch(path, { headers: await getAuthHeaders() })
    const data = await response.json()
    return data
  } catch (e) {
    console.log('Error', e)
    return null
  }
}

const confirmDeclineAgentAllowanceRequest = async (params: {
  username: string
  token: string
  confirm: boolean
}): Promise<boolean> => {
  const { username, token, confirm } = params
  try {
    const path = `${config.chain.localProxyRouterUrl}/auth/allowance/confirm`
    await fetch(path, {
      method: 'POST',
      body: JSON.stringify({ username, token, confirm }),
      headers: await getAuthHeaders()
    })
    return true
  } catch (e) {
    console.log('Error', e)
    return false
  }
}


const getIpfsVersion = async (): Promise<{ version: string } | null> => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/ipfs/version`;
    const response = await fetch(path, { headers: await getAuthHeaders() });
    const body = await response.json();
    return body;
  }
  catch (e) {
    console.log("Error", e)
    return null;
  }
}

const getIpfsFile = async ({ cid, destinationPath }: { cid: string, destinationPath: string }): Promise<ResultResponse | null> => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/ipfs/download/${cid}`;
    const response = await fetch(path, { headers: await getAuthHeaders(), method: "POST", body: JSON.stringify({ destinationPath }) });
    const body = await response.json();
    return body;
  }
  catch (e) {
    console.log("Error", e)
    return null;
  }
}

const pinIpfsFile = async ({ cid }: { cid: string }): Promise<ResultResponse | null> => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/ipfs/pin`;
    const response = await fetch(path, { method: "POST", headers: await getAuthHeaders(), body: JSON.stringify({ cid }) });
    const body = await response.json();
    return body;
  }
  catch (e) {
    console.log("Error", e)
    return null;
  }
}

const unpinIpfsFile = async ({ cid }: { cid: string }): Promise<ResultResponse | null> => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/ipfs/unpin    `;
    const response = await fetch(path, { method: "POST", headers: await getAuthHeaders(), body: JSON.stringify({ cid }) });
    const body = await response.json();
    return body;
  }
  catch (e) {
    console.log("Error", e)
    return null;
  }
}

const addFileToIpfs = async ({ filePath }: { filePath: string }): Promise<{ hash: string, cid: string } | null> => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/ipfs/add`;
    const response = await fetch(path, { method: "POST", headers: await getAuthHeaders(), body: JSON.stringify({ filePath }) });
    const body = await response.json();
    return body;
  }
  catch (e) {
    console.log("Error", e)
    return null;
  }
}

const getIpfsPinnedFiles = async (): Promise<{ files: { cid: string, hash: string }[] } | null> => {
  try {
    const path = `${config.chain.localProxyRouterUrl}/ipfs/pin`;
    const response = await fetch(path, { headers: await getAuthHeaders() });
    const body = await response.json();
    return body;
  }
  catch (e) {
    console.log("Error", e)
    return null;
  }
}

const apiGateway = {
  getAllModels,
  getBalances,
  sendEth,
  sendMor,
  getTransactions,
  getMorRate,
  getTodaysBudget,
  getTokenSupply,
  getChatHistoryTitles,
  getChatHistory,
  updateChatHistoryTitle,
  deleteChatHistory,
  checkProviderConnectivity,
  getAuthHeaders,
  clearWallet,
  clearEthNodeEnv,
  getAgentUsers,
  confirmDeclineAgentUser,
  removeAgentUser,
  getAgentTxs,
  revokeAgentAllowance,
  getAgentAllowanceRequests,
  confirmDeclineAgentAllowanceRequest,
  getIpfsVersion,
  getIpfsFile,
  pinIpfsFile,
  unpinIpfsFile,
  addFileToIpfs,
  getIpfsPinnedFiles,
}

export default apiGateway
export type ApiGateway = typeof apiGateway
