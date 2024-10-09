import config from '../../config'

const getAllModels = async () => {
    try {
        const path = `${config.chain.localProxyRouterUrl}/blockchain/models`
        const response = await fetch(path);
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
        const response = await fetch(path);
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
            })
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
            })
        });
        const data = await response.json();
        return data.txHash;
    }
    catch (e) {
        console.log("Error", e)
        return undefined;
    }
}

const getTransactions = async (page = 1, size = 15) => {
    try {
        const path = `${config.chain.localProxyRouterUrl}/blockchain/transactions?page=${page}&limit=${size}`
        const response = await fetch(path);
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
        const response = await fetch(path);
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
        const response = await fetch(path);
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
        const response = await fetch(path);
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
       const response = await fetch(path);
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
        });
        const body = await response.json();
        return body.result;
    }
    catch (e) {
        console.log("Error", e)
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
}