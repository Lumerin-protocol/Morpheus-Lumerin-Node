import * as validators from '../validators';
import { withClient } from './clientContext';
import * as utils from '../utils';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import React from 'react';
import { ToastsContext } from '../../components/toasts';
import selectors from '../selectors';

const withChatState = WrappedComponent => {
  class Container extends React.Component {
    
    static contextType = ToastsContext;

    static displayName = `withChatState(${WrappedComponent.displayName ||
      WrappedComponent.name})`;

    getBitsByModels = async (modelId) => {
      try {
        const path = `${this.props.config.chain.localProxyRouterUrl}/blockchain/models/${modelId}/bids`
        const response = await fetch(path);
        const data = await response.json();
        return data.bids;
      }
      catch (e) {
        console.log("Error", e)
        return [];
      }
    }

    getProviders = async () => {
      try {
        const path = `${this.props.config.chain.localProxyRouterUrl}/blockchain/providers`
        const response = await fetch(path);
        const data = await response.json();
        return data.providers;
      }
      catch (e) {
        console.log("Error", e)
        return [];
      }
    }

    closeSession = async (sessionId) => {
      try {
        const path = `${this.props.config.chain.localProxyRouterUrl}/blockchain/sessions/${sessionId}/close`;
        const response = await fetch(path, {
          method: "POST"
        });
        const data = await response.json();
        return data.success;
      }
      catch (e) {
        console.log("Error", e)
        return [];
      }
    }

    getAllModels = async () => {
      const result = await this.props.client.getAllModels();
      return result;
    }

    getModelsData = async () => {
      const models = (await this.getAllModels()).filter(m => !m.IsDeleted);
      const providers = (await this.getProviders()).filter(m => !m.IsDeleted);
      const providersMap = providers.reduce((a, b) => ({ ...a, [b.Address.toLowerCase()]: b }), {});
      let result = [];

      for (const model of models) {
        const id = model.Id;
        const bids = (await this.getBitsByModels(id)).filter(b => !b.DeletedAt);
        if (!bids.length) {
          continue;
        }

        const bidsWithProviders = bids.map(b => ({ ...b, ProviderData: providersMap[b.Provider.toLowerCase()], Model: model }))

        result.push({ ...model, bids: model.Name == "Llama 2.0" ? [...bidsWithProviders, { Provider: "Local", Model: model }] : bidsWithProviders })
      }

      return { models: result, providers }
    }


    getMetaInfo = async () => {
      var budget = await this.props.client.getTodaysBudget();
      var supply = await this.props.client.getTokenSupply();
      return { budget, supply };
    }

    getSessionsByUser = async (user) => {
      try {
        const path = `${this.props.config.chain.localProxyRouterUrl}/blockchain/sessions?user=${user}`;
        const response = await fetch(path);
        const data = await response.json();
        return data.sessions;
      }
      catch (e) {
        console.log("Error", e)
        return [];
      }
    }

    onOpenSession = async ({ selectedBid, stake }) => {
      console.log("open-session", stake);
      this.context.toast('info', 'Processing...');
      let signature = {};
      let sessionId = '';

      const bidId = selectedBid.Id;

      try {
        const path = `${this.props.config.chain.localProxyRouterUrl}/proxy/sessions/initiate`;
        const body = {
          user: this.props.address,
          provider: selectedBid.Provider,
          spend: Number(stake),
          bidId,
          providerUrl: selectedBid.ProviderData.Endpoint.replace("http://", "")
        };
        const response = await fetch(path, {
          method: "POST",
          body: JSON.stringify(body)
        });
        const dataResponse = await response.json();
        if (!response.ok) {
          this.context.toast('error', 'Failed to initiate session');
          console.log("Failed initiate session", dataResponse);
          return;
        }
        signature = dataResponse.response.result;
      }
      catch (e) {
        console.error(e);
        this.context.toast('error', 'Failed to initiate session');
        return;
      }

      try {
        const path = `${this.props.config.chain.localProxyRouterUrl}/blockchain/approve?amount=${BigInt(stake).toString()}&spender=${this.props.config.chain.diamondAddress}`;
        const response = await fetch(path, {
          method: "POST",
        });
        const dataResponse = await response.json();
        if (!response.ok) {
          this.context.toast('error', 'Failed to increase allowance');
          console.log("Failed to increase allowance", dataResponse);
          return;
        }
      }
      catch (e) {
        console.error(e);
        this.context.toast('error', 'Failed to increase allowance');
        return;
      }

      try {
        const path = `${this.props.config.chain.localProxyRouterUrl}/blockchain/sessions`;
        const body = {
          approval: signature.approval,
          approvalSig: signature.approvalSig,
          stake: BigInt(stake).toString(),
        };
        const response = await fetch(path, {
          method: "POST",
          body: JSON.stringify(body)
        });
        const dataResponse = await response.json();
        if (!response.ok) {
          this.context.toast('error', 'Failed to set session');
          console.log("Failed to set session", dataResponse);
          return;
        }
        sessionId = dataResponse.sessionId;
      }
      catch (e) {
        console.error(e);
        this.context.toast('error', 'Failed to set session');
        return;
      }

      this.context.toast('success', 'Session successfully created');
      return { sessionId, signature };
    }

    render() {

      return (
        <WrappedComponent
          getProviders={this.getProviders}
          getBitsByModels={this.getBitsByModels}
          getMetaInfo={this.getMetaInfo}
          getModelsData={this.getModelsData}
          getSessionsByUser={this.getSessionsByUser}
          closeSession={this.closeSession}
          onOpenSession={this.onOpenSession}
          toasts={this.context}
          {...this.state}
          {...this.props}
        />
      );
    }
  }

  const mapStateToProps = (state, props) => ({
    // selectedCurrency: selectors.getSellerSelectedCurrency(state),
    // isLocalProxyRouter: selectors.getIsLocalProxyRouter(state),
    // titanLightningPool: state.config.chain.titanLightningPool,
    // titanLightningDashboard: state.config.chain.titanLightningDashboard,
    config: state.config,
    selectedBid: state.models.selectedBid,
    model: state.models.selectedModel,
    provider: state.models.selectedProvider,
    activeSession: state.models.activeSession,
    address: selectors.getWalletAddress(state),
  });

  const mapDispatchToProps = dispatch => ({
    setBid: model => dispatch({ type: 'set-bid', payload: model })
  });

  return withClient(connect(mapStateToProps, mapDispatchToProps)(Container));
};

export default withChatState;
