import config from '../../config'
import fs from 'fs'

const cookieFile = fs.readFileSync(config.chain.proxyRouterCookieFilePath, 'utf8').trim();
const [username, password] = cookieFile.split(':');
const auth = {
    Authorization: `Basic ${Buffer.from(`${username}:${password}`, 'utf-8').toString('base64')}`
}

const getAllModels = async () => {
    try {
        const path = `${config.chain.localProxyRouterUrl}/blockchain/models`
        const response = await fetch(path, {
            headers: auth,
            method: "GET"
        });
        const data = await response.json();
        return data.models;
    }
    catch (e) {
        console.log("Error", e)
        return [];
    }
}

const getBalances = async () => {
    try {
        const path = `${config.chain.localProxyRouterUrl}/blockchain/balance`
        const response = await fetch(path, {
            headers: auth
        });
        const data = await response.json();
        return data;
    }
    catch (e) {
        console.log("Error", e)
        return [];
    }
}

const sendEth = async (to, amount) => {
    try {
        const path = `${config.chain.localProxyRouterUrl}/blockchain/send/eth`
        const response = await fetch(path, {
            method: "POST",
            body: JSON.stringify({
                to, amount
            }),
            headers: auth
        });
        const data = await response.json();
        return data.txHash;
    }
    catch (e) {
        console.log("Error", e)
        return undefined;
    }
}

const sendMor = async (to, amount) => {
    try {
        const path = `${config.chain.localProxyRouterUrl}/blockchain/send/mor`
        const response = await fetch(path, {
            method: "POST",
            body: JSON.stringify({
                to, amount
            }),
            headers: auth
        });
        const data = await response.json();
        return data.txHash;
    }
    catch (e) {
        console.log("Error", e)
        return undefined;
    }
}

const getTransactions = async (payload) => {
    try {
        const path = `${config.chain.localProxyRouterUrl}/blockchain/transactions?page=${payload.page}&limit=${payload.pageSize}`
        const response = await fetch(path, {
            headers: auth
        });
        const data = await response.json();
        return data.transactions;
    }
    catch (e) {
        console.log("Error", e)
        return [];
    }
}

const getMorRate = async (tokenAddress = "0x092baadb7def4c3981454dd9c0a0d7ff07bcfc86", network = "arbitrum") => {
    try {
        const path = `https://api.geckoterminal.com/api/v2/simple/networks/${network}/token_price/${tokenAddress}`;
        const response = await fetch(path);
        const body = await response.json();
        return body.data.attributes.token_prices[tokenAddress];
    }
    catch (e) {
        console.log("Error", e)
        return null;
    }
}

const getTodaysBudget = async () => {
    try {
        const path = `${config.chain.localProxyRouterUrl}/blockchain/sessions/budget`;
        const response = await fetch(path, {
            headers: auth
        });
        const body = await response.json();
        return body.budget;
    }
    catch (e) {
        console.log("Error", e)
        return null;
    }
}

const getTokenSupply = async () => {
    try {
        const path = `${config.chain.localProxyRouterUrl}/blockchain/token/supply`;
        const response = await fetch(path, {
            headers: auth
        });
        const body = await response.json();
        return body.supply;
    }
    catch (e) {
        console.log("Error", e)
        return null;
    }
}

/**
 * 
 * @returns {{
 *   title: string,
 *   chatId: string[]
 * }}
 */
const getChatHistoryTitles = async () => {
    try {
        const path = `${config.chain.localProxyRouterUrl}/v1/chats`;
        const response = await fetch(path, {
            headers: auth
        });
        const body = await response.json();
        return body;
    }
    catch (e) {
        console.log("Error", e)
        return null;
    }
}

/**
 * @typedef ChatHistory
 * @property {string} title
 * @property {string} modelId
 * @property {string} sessionId
 * @property {ChatMessage[]} messages
 */

/**
 * @typedef ChatMessage
 * @property {string} response
 * @property {string} prompt
 * @property {number} promptAt
 * @property {number} responseAt
 */ 

/**
 * @typedef ChatPrompt
 * @property {string} model
 * @property {{
 *  role: string,
 *  content: string
 * }[]} messages
 */

/**
 * @param {string} chatId
 * @returns {Promise<ChatHistory>}
*/
const getChatHistory = async (chatId) => {
   try {
       const path = `${config.chain.localProxyRouterUrl}/v1/chats/${chatId}`;
       const response = await fetch(path, {
              headers: auth
       });
       const body = await response.json();
       return body;
   }
   catch (e) {
       console.log("Error", e)
       return null;
   }
}

/**
 * @param {string} chatId
 * @returns {Promise<boolean>}
*/
const deleteChatHistory = async (chatId) => {
    try {
        const path = `${config.chain.localProxyRouterUrl}/v1/chats/${chatId}`;
        const response = await fetch(path, {
            method: "DELETE",
            headers: auth
        });
        const body = await response.json();
        return body.result;
    }
    catch (e) {
        console.log("Error", e)
        return false;
    }
 }

 /**
 * @param {string} chatId
 * @param {string} title
 * @returns {Promise<boolean>}
*/
const updateChatHistoryTitle = async ({ id, title}) => {
    try {
        const path = `${config.chain.localProxyRouterUrl}/v1/chats/${id}`;
        const response = await fetch(path, {
            method: "POST",
            body: JSON.stringify({ title }),
            headers: auth
        });
        const body = await response.json();
        return body.result;
    }
    catch (e) {
        console.log("Error", e)
        return false;
    }
 }

  /**
 * @param {string} address
 * @param {string} endpoint
 * @returns {Promise<boolean>}
*/
const checkProviderConnectivity = async ({ address, endpoint}) => {
    try {
        const path = `${config.chain.localProxyRouterUrl}/proxy/provider/ping`;
        const response = await fetch(path, {
            method: "POST",
            body: JSON.stringify({ 
                providerAddr: address,
                providerUrl: endpoint
             }),
            headers: auth
        });

        if(!response.ok) {
            return false;
        }

        const body = await response.json();
        return !!body.ping;
    }
    catch (e) {
        console.log("checkProviderConnectivity: Error", e)
        return false;
    }
 }

const getAuthHeaders = async () => auth;

const clearEthNodeEnv = async () => {
    try {
        const path = `${config.chain.localProxyRouterUrl}/config/ethNode`
        const response = await fetch(path, { method: "DELETE", headers: auth });
        const data = await response.json();
        return data.status;
    }
    catch (e) {
        console.log("CLEAR ETH NODE ERROR", e)
        return false;
    }
}

const clearWallet = async () => {
    try {
        const path = `${config.chain.localProxyRouterUrl}/wallet`
        const response = await fetch(path, { method: "DELETE", headers: auth });
        const data = await response.json();
        return data.status;
    }
    catch (e) {
        console.log("CLEAR WALLET ERROR", e)
        return false;
    }
}

export default {
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
    clearEthNodeEnv
}