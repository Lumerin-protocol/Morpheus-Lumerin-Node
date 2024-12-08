'use strict'

import handlers from '../handlers'
import utils from './utils'

const listeners = {
  'validate-password': handlers.validatePassword,
  'change-password': handlers.changePassword,
  'persist-state': handlers.persistState,
  'clear-cache': handlers.clearCache,
  'handle-client-error': handlers.handleClientSideError,
  'get-pool-address': handlers.getPoolAddress,
  logout: handlers.logout,
  'save-proxy-router-settings': handlers.saveProxyRouterSettings,
  'get-proxy-router-settings': handlers.getProxyRouterSettings,
  'get-default-currency-settings': handlers.getDefaultCurrency,
  'set-default-currency-settings': handlers.setDefaultCurrency,
  'get-custom-env-values': handlers.getCustomEnvs,
  'set-custom-env-values': handlers.setCustomEnvs,
  'get-profit-settings': handlers.getProfitSettings,
  'set-profit-settings': handlers.setProfitSettings,
  'get-contract-hashrate': handlers.getContractHashrate,
  'get-auto-adjust-price': handlers.getAutoAdjustPriceData,
  'set-auto-adjust-price': handlers.setAutoAdjustPriceData,
  // Api Gateway
  'get-all-models': handlers.getAllModels,
  'get-transactions': handlers.getTransactions,
  'get-balances': handlers.getBalances,
  'get-rates': handlers.getMorRate,
  "get-todays-budget": handlers.getTodaysBudget,
  "get-supply": handlers.getTokenSupply,
  // Chat history
  "get-chat-history-titles": handlers.getChatHistoryTitles,
  "get-chat-history": handlers.getChatHistory,
  "delete-chat-history": handlers.deleteChatHistory,
  "update-chat-history-title": handlers.updateChatHistoryTitle,
  // Failover
  "get-failover-setting": handlers.isFailoverEnabled,
  "set-failover-setting": handlers.setFailoverSetting,
  "check-provider-connectivity": handlers.checkProviderConnectivity
}

// Subscribe to messages where no core has to react
export const subscribeWithoutCore = () => utils.subscribeTo(listeners, 'none')

export const unsubscribeWithoutCore = () => utils.unsubscribeTo(listeners)

export default { subscribeWithoutCore, unsubscribeWithoutCore }
