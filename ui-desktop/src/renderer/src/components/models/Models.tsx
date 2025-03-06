import { useState, useEffect } from 'react';
import withModelsState from "../../store/hocs/withModelsState";

import { LayoutHeader } from '../common/LayoutHeader'
import { View } from '../common/View'
import { BtnAccent } from '../dashboard/BalanceBlock.styles';
import ModelsTable from './ModelsTable';
import Tab from 'react-bootstrap/Tab';
import Tabs from 'react-bootstrap/Tabs';
import styled from 'styled-components'
import FileSelectionModal from './FileSelectionModal';
import PinnedFilesTable from './PinnedFilesTable';


const Container = styled.div`
    overflow-y: auto;
    
    .nav-link {
        color: ${p => p.theme.colors.morMain}
    }

    .nav-link.active {
        color: ${p => p.theme.colors.morMain}
        border-color: ${p => p.theme.colors.morMain}
        background-color: rgba(0,0,0,0.4);
    }
`

const IpfsStatus = styled.div`
    color: ${p => p.theme.colors.morMain};
    font-size: 1.2rem;
`

const Models = ({
    setSelectedModel,
    getIpfsVersion,
    getAllModels,
    openSelectDownloadFolder,
    downloadModelFromIpfs,
    addFileToIpfs,
    getPinnedFiles,
    pinFile,
    unpinFile,
    toasts,
}: any) => {

    const [openChangeModal, setOpenChangeModal] = useState(false);
    const [ipfsVersion, setIpfsVersion] = useState(null);
    const [isIpfsConnected, setIsIpfsConnected] = useState(false);
    const [pinnedFiles, setPinnedFiles] = useState([]);
    const [models, setModels] = useState([]);

    const reload = () => {
        getIpfsVersion().then((response) => {
            if (response?.version) {
                setIpfsVersion(response?.version);
                setIsIpfsConnected(true);
            } else {
                setIsIpfsConnected(false);
            }
        }).catch((error) => {
            console.error("Error", error);
            setIsIpfsConnected(false);
        });

        getAllModels().then((response) => {
            setModels(response);
        }).catch((error) => {
            console.error("Error", error);
        });

        getPinnedFiles().then((response) => {
            setPinnedFiles(response.files);
        }).catch((error) => {
            console.error("Error", error);
        });
    }

    useEffect(() => {
        reload();
    }, []);

    const handleUnpinFile = async (hash) => {
        try {
            const response = await unpinFile(hash);
            if (response) {
                toasts.toast("success", "File unpinned successfully");
                setPinnedFiles(pinnedFiles.filter((file: any) => file.hash !== hash));
            } else {
                toasts.toast("error", "Failed to unpin file");
            }
        } catch (error) {
            toasts.toast("error", "Failed to unpin file");
            console.error("Error", error);
        }
    }

    const onPinModel = async (hash) => {
        const response = await pinFile(hash);
        reload();
        return response;
    }

    return (
        <View data-testid="models-container">
                        {isIpfsConnected ? (
                    <IpfsStatus>
                        <span>IPFS Connected. Version: {ipfsVersion}</span>
                    </IpfsStatus>
                ) : (
                    <IpfsStatus>
                        <span>IPFS is not connected</span>
                    </IpfsStatus>
                )}
            <LayoutHeader title="Models">
                <BtnAccent style={{ padding: '1.5rem' }} onClick={() => setOpenChangeModal(true)}>Pin Model</BtnAccent>
            </LayoutHeader>
            <Container>
                <Tabs
                    defaultActiveKey="registry"
                    id="tab-models"
                    className="mb-3"
                >
                    <Tab eventKey="registry" title="Registry">
                        <ModelsTable setSelectedModel={setSelectedModel} models={models} openSelectDownloadFolder={openSelectDownloadFolder} downloadModelFromIpfs={downloadModelFromIpfs} toasts={toasts} />
                    </Tab>
                    <Tab eventKey="pinned" title="Pinned Models">
                        <PinnedFilesTable pinnedFiles={pinnedFiles} unpinFile={handleUnpinFile} toasts={toasts} />
                    </Tab>
                </Tabs>

            </Container>
            <FileSelectionModal
                isActive={openChangeModal}
                addFileToIpfs={addFileToIpfs}
                pinFile={onPinModel}
                toasts={toasts}
                handleClose={() => setOpenChangeModal(false)}
                 />
        </View>)

}

export default withModelsState(Models);