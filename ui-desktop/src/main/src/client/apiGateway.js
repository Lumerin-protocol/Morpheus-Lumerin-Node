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

export default {
    getAllModels,
    getBalances,
    sendEth,
    sendMor,
    getTransactions,
    getMorRate,
    getTodaysBudget,
    getTokenSupply
}