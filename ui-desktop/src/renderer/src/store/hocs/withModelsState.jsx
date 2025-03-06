import * as validators from '../validators';
import { withClient } from './clientContext';
import * as utils from '../utils';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import React from 'react';
import { ToastsContext } from '../../components/toasts';
import selectors from '../selectors';
import path from 'path';

const withModelsState = WrappedComponent => {
  class Container extends React.Component {
   
    static contextType = ToastsContext;

    static displayName = `withModelsState(${WrappedComponent.displayName ||
      WrappedComponent.name})`;

    getAllModels = async () => {
        const result = await this.props.client.getAllModels();
        return result;
    }

    getAllProviders = async () => {
      try {
        const authHeaders = await this.props.client.getAuthHeaders();
        const path = `${this.props.config.chain.localProxyRouterUrl}/blockchain/providers`
        const response = await fetch(path, {
          headers: authHeaders
        });
        const data = await response.json();
        return data.providers;
      }
      catch(e) {
        console.log("Error", e)
        return [];
      }
    }

    getBitsByModels = async (modelId) => {
        
    }

    getIpfsVersion = async () => {
      const response = await this.props.client.getIpfsVersion();
      return response;
    }

    openSelectDownloadFolder = async () => {
      const response = await this.props.client.openSelectFolderDialog();
      return response;
    }

    downloadModelFromIpfs = async (modelId, folderPath) => {
      const filePath = path.join(folderPath, modelId);
      const response = await this.props.client.getIpfsFile({ cid: modelId, destinationPath: filePath });
      return response;
    }

    addFileToIpfs = async (filePath) => {
      const response = await this.props.client.addFileToIpfs({ filePath });
      return response;
    }

    pinFile = async (cid) => {
      const response = await this.props.client.pinIpfsFile({ cid });
      return response;
    }

    unpinFile = async (cid) => {
      const response = await this.props.client.unpinIpfsFile({ cid });
      return response;
    }

    getPinnedFiles = async () => {
      const response = await this.props.client.getIpfsPinnedFiles();
      return response;
    }
 
    render() {

      return (
        <WrappedComponent
            getAllModels={this.getAllModels}
            getAllProviders={this.getAllProviders}
            getIpfsVersion={this.getIpfsVersion}
            openSelectDownloadFolder={this.openSelectDownloadFolder}
            downloadModelFromIpfs={this.downloadModelFromIpfs}
            addFileToIpfs={this.addFileToIpfs}
            getPinnedFiles={this.getPinnedFiles}
            pinFile={this.pinFile}
            unpinFile={this.unpinFile}
            toasts={this.context}
            {...this.state}
            {...this.props}
        />
      );
    }
  }

  const mapStateToProps = (state, props) => ({
    // selectedCurrency: selectors.getSellerSelectedCurrency(state),
    config: state.config
  });

  const mapDispatchToProps = dispatch => ({
    setSelectedModel: model => dispatch({ type: 'set-model', payload: model })
  });

  return withClient(connect(mapStateToProps, mapDispatchToProps)(Container));
};

export default withModelsState;
