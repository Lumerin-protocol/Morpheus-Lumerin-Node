import React from 'react';
import { withClient } from './clientContext';
import selectors from '../selectors';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import { ToastsContext } from '../../components/toasts';
import { getSessionsByUser } from '../utils/apiCallsHelper';

const withDashboardState = WrappedComponent => {
  class Container extends React.Component {
    static propTypes = {
      sendLmrFeatureStatus: PropTypes.oneOf(['offline', 'no-funds', 'ok'])
        .isRequired,
      syncStatus: PropTypes.oneOf(['up-to-date', 'syncing', 'failed'])
        .isRequired,
      address: PropTypes.string.isRequired,
      client: PropTypes.shape({
        refreshAllTransactions: PropTypes.func.isRequired,
        copyToClipboard: PropTypes.func.isRequired
      }).isRequired
    };

    static contextType = ToastsContext;

    static displayName = `withDashboardState(${WrappedComponent.displayName ||
      WrappedComponent.name})`;

    state = {
      refreshStatus: 'init',
      refreshError: null
    };

    onWalletRefresh = () => {
      this.setState({ refreshStatus: 'pending', refreshError: null });
      this.props.client.getTransactions()
        .then(() => this.setState({ refreshStatus: 'success' }))
        .catch(() => {
          this.context.toast(
            'error',
            'We’re experiencing connection problems.  Please wait a few minutes and try again'
          );
          this.setState({
            refreshStatus: 'failure',
            refreshError: 'Could not refresh'
          });
        });
    };

    loadTransactions = async (page = 1, pageSize = 15) => {
      this.setState({ refreshStatus: 'pending', refreshError: null });
      const transactions = await this.props.client.getTransactions({ page, pageSize });
      this.setState({ refreshStatus: 'success' })
  
      // if (page && pageSize) {
      //   const hasNextPage = transactions.length;
      //   this.props.nextPage({
      //     hasNextPage: Boolean(hasNextPage),
      //     page: page + 1,
      //   })
      // }
      return transactions;
    }

    getStakedFunds = async (user) => {
      const isClosed = (item) => item.ClosedAt || (new Date().getTime() > item.EndsAt * 1000);

      if(!user) {
        return;
      }

      const sessions = await getSessionsByUser(this.props.config.chain.localProxyRouterUrl, user);
      
      try {
        const openSessions = sessions.filter(s => !isClosed(s));
        const sum = openSessions.reduce((curr, next) => curr + next.Stake, 0);
        return (sum / 10 ** 18).toFixed(2);
      }
      catch (e) {
        console.log("Error", e)
        return 0;
      }
    }

    getBalances = async () => {
      const balances = await this.props.client.getBalances();
      const rate = await this.props.client.getRates();
      return { balances, rate };
    }

    render() {
      const { sendLmrFeatureStatus } = this.props;

      const sendDisabledReason =
        sendLmrFeatureStatus === 'offline'
          ? "Can't send while offline"
          : sendLmrFeatureStatus === 'no-funds'
          ? 'You need some funds to send'
          : null;

      return (
        <WrappedComponent
          sendDisabledReason={sendDisabledReason}
          copyToClipboard={this.props.client.copyToClipboard}
          onWalletRefresh={this.onWalletRefresh}
          getBalances={this.getBalances}
          sendDisabled={sendLmrFeatureStatus !== 'ok'}
          loadTransactions={this.loadTransactions}
          getStakedFunds={this.getStakedFunds}
          {...this.props}
          {...this.state}
        />
      );
    }
  }

  const mapStateToProps = state => ({
    config: state.config,
    syncStatus: selectors.getTxSyncStatus(state),
    sendLmrFeatureStatus: selectors.sendLmrFeatureStatus(state),
    hasTransactions: selectors.hasTransactions(state),
    address: selectors.getWalletAddress(state),
    ethCoinPrice: selectors.getRateEth(state),
    symbol: selectors.getCoinSymbol(state),
    symbolEth: selectors.getSymbolEth(state),
    page: selectors.getTransactionPage(state),
    pageSize: selectors.getTransactionPageSize(state),
    hasNextPage: selectors.getHasNextPage(state),
    explorerUrl: selectors.getContractExplorerUrl(state, { 
      hash: selectors.getWalletAddress(state)
    })
  });

  const mapDispatchToProps = dispatch => ({
    nextPage: data => dispatch({ type: 'transactions-next-page', payload: data }),
});


  return withClient(connect(mapStateToProps, mapDispatchToProps)(Container));
};

export default withDashboardState;

function mapApiResponseToTrxReceipt(trx){
  const transaction = {
    from: trx.from,
    to: trx.to,
    value: trx.value,
    input: trx.input,
    gas: trx.gas,
    gasPrice: trx.gasPrice,
    hash: trx.hash,
    nonce: trx.nonce,
    logIndex: trx.logIndex, // emitted only in events, used to differentiate between LMR transfers within one transaction 
    // maxFeePerGas: params.maxFeePerGas,
    // maxPriorityFeePerGas: params.maxPriorityFeePerGas,
  }

  if (trx.returnValues){
    transaction.from = trx.returnValues.from;
    transaction.to = trx.returnValues.to;
    transaction.value = trx.returnValues.value;
    transaction.hash = trx.transactionHash;
  }

  const receipt = {
    transactionHash: trx.hash,
    transactionIndex: trx.transactionIndex,
    blockHash: trx.blockHash,
    blockNumber: trx.blockNumber,
    from: trx.from,
    to: trx.to,
    value: trx.value,
    contractAddress: trx.contractAddress,
    cumulativeGasUsed: trx.cumulativeGasUsed,
    gasUsed: trx.gasUsed,
    tokenSymbol: trx.tokenSymbol,
  }

  if (trx.returnValues){
    receipt.from = trx.returnValues.from;
    receipt.to = trx.returnValues.to;
    receipt.value = trx.returnValues.value;
    receipt.transactionHash = trx.transactionHash;
    receipt.tokenSymbol = trx.address === config.chain.mainTokenAddress ? 'MOR' : undefined;
  }

  return {transaction, receipt}
}