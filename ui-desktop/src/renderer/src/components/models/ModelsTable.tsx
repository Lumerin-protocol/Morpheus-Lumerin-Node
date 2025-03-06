import { useRef, useState } from 'react';

import withModelsState from '../../store/hocs/withModelsState';
import styled from 'styled-components';

import Card from 'react-bootstrap/Card';
import { abbreviateAddress } from '../../utils';
import { IconDownload } from '@tabler/icons-react';
import Form from 'react-bootstrap/esm/Form';


const CustomCard = styled(Card)`
  background: #244a47 !important;
  color: #21dc8f !important;
  border: 0.5px solid !important;
  cursor: pointer !important;

  p {
    color: white !important;
  }

  .gap-20 {
    gap: 20px !important;
  }
`;

const Container = styled.div`
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  gap: 24px;
  max-height: 75vh;
  overflow-y: auto;
`;

function ModelCard({ onSelect, model, openSelectDownloadFolder, downloadModelFromIpfs, toasts }) {
  const handleFolderSelect = async (e) => {
    e.stopPropagation();
    try {
      const result = await openSelectDownloadFolder();
      const { canceled, filePaths } = result;
      if (canceled) {
        return;
      }
      const folderPath = filePaths[0];
      const response = await downloadModelFromIpfs(model.IpfsCID, folderPath);
      if (response) {
        toasts.toast("success", "Model downloaded successfully");
      } else {
        toasts.toast("error", "Failed to download model");
      }
    } catch (error) {
      if (error?.includes("invalid CID")) {
        toasts.toast("error", "Invalid CID specified in the model.");
      } else if (error?.includes("failed to find file")) {
        toasts.toast("error", "Model is not found in IPFS.");
      } else {
        toasts.toast("error", "Failed to download model");
      }
    }
  };

  return (
    <CustomCard style={{ width: '36rem' }} onClick={() => onSelect(model.Id)}>
      <Card.Body>
        <Card.Title
          as={'div'}
          style={{ display: 'flex', justifyContent: 'space-between' }}
        >
          {model.Name}
          <IconDownload 
            style={{ cursor: 'pointer' }} 
            onClick={handleFolderSelect}
          />
        </Card.Title>
        <Card.Subtitle className="mb-2">
          {abbreviateAddress(model.Id, 6)}
        </Card.Subtitle>
        <Card.Text>
          {/* <div>
          Fee: {model.Fee}
          </div>
          <div>
          Stake: {model.Stake}
          </div> */}
        </Card.Text>
        <Card.Footer className="d-flex gap-20">
          {model.Tags.map((t) => (
            <div key={t}>{t}</div>
          ))}
        </Card.Footer>
      </Card.Body>
    </CustomCard>
  );
}


function  ModelsTable({
  setSelectedModel,
  models,
  client,
  openSelectDownloadFolder,
  downloadModelFromIpfs,
  toasts,
} : any) {
  const onSelect = (id) => {
    setSelectedModel(models.find((x) => x.Id == id));
  };

  return (<Container>
     {
      models.length ? models.map((x => (<div>{ModelCard({ onSelect, model: x, openSelectDownloadFolder, downloadModelFromIpfs, toasts })}</div>))) : null
     }
    </Container>)
}

export default ModelsTable;
