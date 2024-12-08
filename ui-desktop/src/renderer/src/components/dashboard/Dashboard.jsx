import React, { useState, useEffect } from 'react'
import styled from 'styled-components'

import withDashboardState from '../../store/hocs/withDashboardState'

import { ChainHeader } from '../common/ChainHeader'
import BalanceBlock from './BalanceBlock'
import TransactionModal from './tx-modal'
import TxList from './tx-list/TxList'
import { View } from '../common/View'
import { toUSD } from '../../store/utils/syncAmounts';
import {
  BtnAccent,
} from './BalanceBlock.styles';

const CustomBtn = styled(BtnAccent)`
  margin-left: 0;
  padding: 1.5rem 1rem;
`
const WidjetsContainer = styled.div`
    display: flex;
    align-items: center;
    justify-content: left;
    gap: 1.6rem;
`

const WidjetItem = styled.div`
    margin: 1.6rem 0 1.6rem;
    padding: 1.6rem 3.2rem;
    border-radius: 0.375rem;
    color: white;
    max-width: 720px;

    color: white;
`

const StakingWidjet = styled(WidjetItem)`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background: rgba(255,255,255,0.04);
  border-width: 1px;
  border: 1px solid rgba(255,255,255,0.04);
`

const Dashboard = ({
  sendDisabled,
  sendDisabledReason,
  syncStatus,
  address,
  hasTransactions,
  copyToClipboard,
  onWalletRefresh,
  getBalances,
  ethCoinPrice,
  loadTransactions,
  getStakedFunds,
  explorerUrl,
  ...props
}) => {
  const [activeModal, setActiveModal] = useState(null)

  const onCloseModal = () => setActiveModal(null)
  const onTabSwitch = (modal) => setActiveModal(modal)

  const [balanceData, setBalanceData] = useState({
    eth: {
      value: 0, rate: 0, usd: 0, symbol: "ETH"
    },
    mor: {
      value: 0, rate: 0, usd: 0, symbol: "MOR"
    }
  });
  const [transactions, setTransactions] = useState([]);
  const [pagging, setPagging] = useState({ page: 1, pageSize: 50, hasNextPage: true })
  const [staked, setStaked] = useState(0);

  const loadBalances = async () => {
    const data = await getBalances();
    const eth = data.balances.eth / 10 ** 18;
    const mor = data.balances.mor / 10 ** 18;
    const ethUsd = toUSD(eth, ethCoinPrice);
    const morUsd = toUSD(mor, +data.rate);

    const balances = {
      eth: {
        value: eth, rate: ethCoinPrice, usd: ethUsd, symbol: props.symbolEth
      },
      mor: {
        value: mor, rate: +data.rate, usd: morUsd, symbol: props.symbol
      }
    }
    setBalanceData(balances);
  }

  const getTransactions = async () => {
    console.log("LOAD NEXT PAGE", pagging, transactions.length);
    let pageTransactions = await loadTransactions(pagging.page, pagging.pageSize);
    const hasNextPage = !!pageTransactions.length;
    const trx = pageTransactions.filter(t => +t.value > 0).map(mapTransaction);
    setPagging({ ...pagging, page: pagging.page + 1, hasNextPage });
    setTransactions([...transactions, ...trx]);
  }

  const mapTransaction = (transaction) => {
    function isSendTransaction(transaction, myAddress) {
      return transaction.from.toLowerCase() === myAddress.toLowerCase();
    }

    function isReceiveTransaction(transaction, myAddress) {
      return transaction.to.toLowerCase() === myAddress.toLowerCase();
    }

    function getTxType(transaction, myAddress) {
      if (isSendTransaction(transaction, myAddress)) {
        return 'sent';
      }
      if (isReceiveTransaction(transaction, myAddress)) {
        return 'received';
      }
      return 'unknown';
    }

    const isMor = !!transaction.contractAddress;

    return {
      hash: transaction.hash,
      from: transaction.from,
      to: transaction.to,
      txType: getTxType(transaction, address),
      isMor: isMor,
      symbol: isMor ? props.symbol : props.symbolEth,
      value: transaction.value / 10 ** 18
    }
  }

  useEffect(() => {
    loadBalances();
    getTransactions();
    getStakedFunds(address).then((data) => {
      setStaked(data);
    })

    const interval = setInterval(() => {
      console.log("Update balances...")
      loadBalances()
    }, 30000);

    return () => clearInterval(interval);
  }, []);


  useEffect(() => {
    loadBalances();
  }, [ethCoinPrice]);

  return (
    <View data-testid="dashboard-container">
      <ChainHeader title="My Wallet" chain={props.config.chain} address={address} copyToClipboard={copyToClipboard} />

      <BalanceBlock
        {...balanceData}
        sendDisabled={sendDisabled}
        sendDisabledReason={sendDisabledReason}
        onTabSwitch={onTabSwitch}
      />

      <WidjetsContainer>
          <StakingWidjet className='staking'>
            <div>
              Staked Balance
            </div>
            <div>{staked} {props.symbol}</div>
          </StakingWidjet>
        <WidjetItem>
        <CustomBtn
             onClick={() => window.openLink(explorerUrl)}
              block
            >
              Transaction Explorer
            </CustomBtn>
        </WidjetItem>
        <WidjetItem>
        <CustomBtn
              onClick={() => window.openLink("https://staking.mor.lumerin.io")}
              block
            >
              Staking Dashboard
            </CustomBtn>
        </WidjetItem>
      </WidjetsContainer>

      <TxList
        {...pagging}
        hasNextPage={pagging.hasNextPage}
        loadNextTransactions={() => {}}
        hasTransactions={!!transactions.length}
        syncStatus={syncStatus}
        transactions={transactions}
      />

      <TransactionModal
        {...balanceData}
        onRequestClose={onCloseModal}
        onTabSwitch={onTabSwitch}
        activeTab={activeModal}
      />
    </View>
  )
}

export default withDashboardState(Dashboard)
