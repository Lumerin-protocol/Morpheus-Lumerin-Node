import { useRef, useState } from 'react';

import withModelsState from '../../store/hocs/withModelsState';
import styled from 'styled-components';

import Card from 'react-bootstrap/Card';
import { abbreviateAddress } from '../../utils';
import { IconPinnedOff, IconCopy } from '@tabler/icons-react';
import Form from 'react-bootstrap/esm/Form';
import client from '../../client';


const CustomCard = styled(Card)`
  background: #244a47 !important;
  color: #21dc8f !important;
  border: 0.5px solid !important;

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

function ModelCard({ model, toasts, unpinFile }) {
  const onUnpinFile = (e) => {
    e.stopPropagation();
    unpinFile(model.hash);
  };

  const copyHash = () => {
    navigator.clipboard.writeText(model.hash);
    toasts.toast("success", "Hash copied to clipboard",{
      autoClose: 700
    });
  }
  return (
    <CustomCard style={{ width: '36rem' }}>
      <Card.Body>
        <Card.Title
          as={'div'}
          style={{ display: 'flex', justifyContent: 'space-between' }}
        >
          <span>
            {abbreviateAddress(model.hash, 6)}
            <IconCopy style={{ cursor: 'pointer', width: '1.5rem', height: '1.5rem', marginLeft: '1rem' }} onClick={() => copyHash()} />
          </span>
          <IconPinnedOff
            style={{ cursor: 'pointer' }}
            onClick={onUnpinFile}
          />
        </Card.Title>
        <Card.Subtitle className="mb-2">
          {model.cid}
        </Card.Subtitle>
        <Card.Text>
        </Card.Text>
        {/* <Card.Footer className="d-flex gap-20">
          {model.Tags.map((t) => (
            <div key={t}>{t}</div>
          ))}
        </Card.Footer> */}
      </Card.Body>
    </CustomCard>
  );
}


function PinnedFilesTable({
  pinnedFiles,
  toasts,
  unpinFile
}: any) {
  return (<Container>
    {
      pinnedFiles?.length ? pinnedFiles.map((x => (<div>{ModelCard({ model: x, toasts, unpinFile })}</div>))) : null
    }
  </Container>)
}

export default PinnedFilesTable;
