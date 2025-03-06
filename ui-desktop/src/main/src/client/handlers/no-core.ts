'use strict'

import { dialog } from 'electron';

import logger from '../../../logger'
import restart from '../electron-restart'
import dbManager from '../database'
import storage from '../storage'
import auth from '../auth'
import wallet from '../wallet'
import {
  setProxyRouterConfig,
  getProxyRouterConfig,
  getDefaultCurrencySetting,
  setDefaultCurrencySetting,
  getKey,
  setKey,
  getFailoverSetting,
  setFailoverSetting
} from '../settings'
import apiGateway from '../apiGateway'
import config from '../../../config'

const validatePassword = (data) => auth.isValidPassword(data)

function clearCache() {
  logger.verbose('Clearing database cache')
  return dbManager.getDb().dropDatabase().then(restart)
}

export const persistState = (data) => storage.persistState(data).then(() => true)

function changePassword({ oldPassword, newPassword }) {
  return validatePassword(oldPassword).then(function (isValid) {
    if (!isValid) {
      return isValid
    }
    return auth.setPassword(newPassword).then(function () {
      const seed = wallet.getSeed(oldPassword)
      wallet.setSeed(seed, newPassword)

      return true
    })
  })
}

const saveProxyRouterSettings = (data) => Promise.resolve(setProxyRouterConfig(data))

const getProxyRouterSettings = async () => {
  return getProxyRouterConfig()
}

const handleClientSideError = (data) => {
  logger.error('client-side error', data.message, data.stack)
}

const getDefaultCurrency = async () => getDefaultCurrencySetting()
const setDefaultCurrency = async (curr) => setDefaultCurrencySetting(curr)

const getCustomEnvs = async () => getKey('customEnvs')
const setCustomEnvs = async (value) => setKey('customEnvs', value)

const getProfitSettings = async () =>
  getKey('profitSettings') || {
    deviation: 2,
    target: 10,
    adaptExisting: false
  }
const setProfitSettings = async (value) => setKey('profitSettings', value)

const getAutoAdjustPriceData = async () => getKey('autoAdjustPriceData')
const setAutoAdjustPriceData = async (value) => {
  const oldData = await getAutoAdjustPriceData()
  setKey('autoAdjustPriceData', {
    ...oldData,
    ...value
  })
}

const getContractHashrate = async (params: { contractId: string; fromDate: Date }) => {
  const { contractId, fromDate } = params
  const collection = await dbManager.getDb().collection('hashrate').findAsync({ id: contractId })
  return collection
    .filter((x) => x.timestamp > fromDate.getTime())
    .sort((a, b) => a.timestamp - b.timestamp)
}

const isFailoverEnabled = async () => {
  const settings = await getFailoverSetting()
  if (!settings) {
    return { isEnabled: config.isFailoverEnabled }
  }
  return settings
}

const restartWallet = () => restart(1);

const openSelectFolderDialog = () => {
  return dialog.showOpenDialog({
    properties: ['openDirectory']
  });
}

const handlers = {
  validatePassword,
  changePassword,
  persistState,
  clearCache,
  saveProxyRouterSettings,
  getProxyRouterSettings,
  handleClientSideError,
  getDefaultCurrency,
  setDefaultCurrency,
  getCustomEnvs,
  setCustomEnvs,
  restartWallet,
  getContractHashrate,
  getProfitSettings,
  setProfitSettings,
  getAutoAdjustPriceData,
  setAutoAdjustPriceData,
  isFailoverEnabled,
  setFailoverSetting,
  openSelectFolderDialog,
  ...apiGateway,
}

export default handlers

export type NoCoreHandlers = typeof handlers
