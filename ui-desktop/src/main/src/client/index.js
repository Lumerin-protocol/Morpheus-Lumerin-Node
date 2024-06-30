'use strict'

const { ipcMain, app } = require('electron')
const createCore = require('@lumerin/wallet-core')

import logger from '../../logger'
import subscriptions from './subscriptions'
import settings from './settings'
import storage from './storage'

export function startCore({ chain, core, config: coreConfig }, webContent) {
  logger.verbose(`Starting core ${chain}`)
  const { emitter, events, api } = core.start(coreConfig)

  // emitter.setMaxListeners(30);
  emitter.setMaxListeners(50)

  events.push(
    'create-wallet',
    'open-wallet',
    'transactions-scan-started',
    'transactions-scan-finished',
    'contracts-scan-started',
    'contracts-scan-finished',
    'contract-updated'
  )

  function send(eventName, data) {
    try {
      if (!webContent) {
        return
      }
      const payload = Object.assign({}, data, { chain })
      webContent.sender.send(eventName, payload)
    } catch (err) {
      logger.error('send error', err)
    }
  }

  events.forEach((event) =>
    emitter.on(event, function (data) {
      send(event, data)
    })
  )

  emitter.on('wallet-error', function (err) {
    logger.warn(err.inner ? `${err.message} - ${err.inner.message}` : err.message)
  })

  return {
    emitter,
    events,
    api
  }
}

export function stopCore({ core, chain }) {
  logger.verbose(`Stopping core ${chain}`)
  core.stop()
}

export function createClient(config) {
  ipcMain.on('log.error', function (_, args) {
    logger.error('ipcMain error ', args.message)
  })

  settings.presetDefaults()

  let core = {
    chain: config.chain.chainId,
    core: createCore(),
    config: Object.assign({}, config.chain, config)
  }

  ipcMain.on('ui-ready', function (webContent, args) {
    const onboardingComplete = !!settings.getPasswordHash()

    storage
      .getState()
      .catch(function (err) {
        logger.warn('Failed to get state', err.message)
        return {}
      })
      .then(function (persistedState) {
        const payload = Object.assign({}, args, {
          data: {
            onboardingComplete,
            persistedState: persistedState || {},
            config
          }
        })
        webContent.sender.send('ui-ready', payload)
        // logger.verbose(`<-- ui-ready ${stringify(payload)}`);
      })
      .catch(function (err) {
        logger.error('Could not send ui-ready message back', err.message)
      })
      .then(function () {
        const { emitter, events, api } = startCore(core, webContent)
        core.emitter = emitter
        core.events = events
        core.api = api
        subscriptions.subscribe(core)
      })
      .catch(function (err) {
        console.log('panic')
        console.log(err)
        console.log('Unknown chain =', err.message)
        logger.error('Could not start core', err.message)
      })
  })

  ipcMain.on('ui-unload', function () {
    stopCore(core)
    subscriptions.unsubscribe(core)
  })
}
