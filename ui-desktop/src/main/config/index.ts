import { parseJSONArray } from './utils'

let httpApiUrls, explorerApiURLs


try {
  httpApiUrls = parseJSONArray(process.env.ETH_NODE_ADDRESS_HTTP)
} catch (err) {
  throw new Error(`Invalid ETH_NODE_ADDRESS_HTTP: ${(err as Error)?.message}`)
}

try {
  explorerApiURLs = parseJSONArray(process.env.EXPLORER_API_URLS)
} catch (err) {
  throw new Error(`Invalid EXPLORER_API_URLS: ${(err as Error)?.message}`)
}

const chain = {
  displayName: process.env.DISPLAY_NAME,
  chainId: process.env.CHAIN_ID,
  symbol: process.env.SYMBOL_LMR || 'LMR',
  symbolEth: process.env.SYMBOL_ETH || 'ETH',

  mainTokenAddress: process.env.TOKEN_ADDRESS,
  cloneFactoryAddress: process.env.CLONE_FACTORY_ADDRESS,

  proxyRouterUrl: process.env.PROXY_ROUTER_URL,
  explorerUrl: process.env.EXPLORER_URL,
  explorerApiURLs: explorerApiURLs,

  wsApiUrl: process.env.ETH_NODE_ADDRESS,
  httpApiUrls: httpApiUrls,
  ipLookupUrl: process.env.IP_LOOKUP_URL,

  coinDefaultGasLimit: process.env.COIN_DEFAULT_GAS_LIMIT,
  lmrDefaultGasLimit: process.env.LMR_DEFAULT_GAS_LIMIT,
  defaultGasPrice: process.env.DEFAULT_GAS_PRICE,
  maxGasPrice: process.env.MAX_GAS_PRICE,

  proxyPort: process.env.PROXY_DEFAULT_PORT || 3333,
  proxyWebPort: process.env.PROXY_WEB_DEFAULT_PORT || 8081,

  portCheckerUrl: process.env.PORT_CHECKER_URL || 'https://portchecker.io/api/v1/query',
  portCheckErrorLink:
    process.env.PORT_CHECK_ERROR_LINK ||
    'https://gitbook.lumerin.io/lumerin-hashpower-marketplace/buyer/2.-network-changes-for-receiving-hashrate',

  localProxyRouterUrl: `http://localhost:${process.env.PROXY_WEB_DEFAULT_PORT || 8082}`,

  faucetUrl: process.env.FAUCET_URL,
  showFaucet: process.env.SHOW_FAUCET === 'true',

  titanLightningPool: process.env.TITAN_LIGHTNING_POOL,
  titanLightningDashboard: process.env.TITAN_LIGHTNING_DASHBOARD || 'https://lightning.titan.io',
  defaultSellerCurrency: process.env.DEFAULT_SELLER_CURRENCY || 'BTC',

  bypassAuth: process.env.BYPASS_AUTH === 'true',
  sellerWhitelistUrl: process.env.SELLER_WHITELIST_URL || 'https://forms.gle/wEcAgppfK2p9YZ3g7'
}

const config = {
  chain,
  dbAutocompactionInterval: 30000,
  debug: process.env.DEBUG === 'true' && process.env.IGNORE_DEBUG_LOGS !== 'true',
  devTools: process.env.DEV_TOOLS === 'true',
  explorerDebounce: 2000,
  ratesUpdateMs: 30000,
  scanTransactionTimeout: 240000,
  sentryDsn: process.env.SENTRY_DSN,
  statePersistanceDebounce: 2000,
  trackingId: process.env.TRACKING_ID,
  web3Timeout: 120000,
  autoAdjustPriceInterval: Number(process.env.AUTO_ADJUST_PRICE_INTERVAL) || 15 * 60 * 1000,
  autoAdjustContractPriceTimeout:
    Number(process.env.AUTO_ADJUST_CONTRACT_PRICE_TIMEOUT) || 24 * 60 * 60 * 1000,
  recaptchaSiteKey: process.env.RECAPTCHA_SITE_KEY
}

export default config