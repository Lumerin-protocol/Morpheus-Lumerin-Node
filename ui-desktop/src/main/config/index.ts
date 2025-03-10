// When adding/removing variables check 'electron.vite.config.ts' to update the list of 'envsToInject'

const config = {
  chain: {
    bypassAuth: process.env.BYPASS_AUTH === 'true',
    chainId: process.env.CHAIN_ID,
    defaultSellerCurrency: process.env.DEFAULT_SELLER_CURRENCY || 'BTC',
    diamondAddress: process.env.DIAMOND_ADDRESS,
    displayName: process.env.DISPLAY_NAME,
    explorerUrl: process.env.EXPLORER_URL,
    localProxyRouterUrl: `http://localhost:${process.env.PROXY_WEB_DEFAULT_PORT || 8082}`,
    mainTokenAddress: process.env.TOKEN_ADDRESS,
    symbol: process.env.SYMBOL_COIN || 'MOR',
    symbolEth: process.env.SYMBOL_ETH || 'ETH'
  },
  dbAutocompactionInterval: 30000,
  debug: process.env.DEBUG === 'true' && process.env.IGNORE_DEBUG_LOGS !== 'true',
  devTools: process.env.DEV_TOOLS === 'true',
  sentryDsn: process.env.SENTRY_DSN,
  statePersistanceDebounce: 2000,
  trackingId: process.env.TRACKING_ID,
  isFailoverEnabled: process.env.FAILOVER_ENABLED ? process.env.FAILOVER_ENABLED === 'true' : true
}
console.log("🚀 ~ config:", config)

export default config
