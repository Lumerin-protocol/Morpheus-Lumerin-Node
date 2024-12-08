import { useEffect, useRef, useState } from 'react'
// import component 👇
import Drawer from 'react-modern-drawer'
import { IconHistory, IconArrowUp, IconMessagePlus } from '@tabler/icons-react';
import {
    View,
    ContainerTitle,
    ChatTitleContainer,
    ChatAvatar,
    Avatar,
    TitleRow,
    AvatarHeader,
    MessageBody,
    Container,
    ChatBlock,
    CustomTextArrea,
    Control,
    SendBtn,
    LoadingCover,
    ImageContainer,
    SubPriceLabel,
    VideoContainer
} from './Chat.styles';
import { BtnAccent } from '../dashboard/BalanceBlock.styles';
import { withRouter } from 'react-router-dom';
import withChatState from '../../store/hocs/withChatState';
import { abbreviateAddress } from '../../utils'
import Markdown from 'react-markdown'

import 'react-modern-drawer/dist/index.css'
import './Chat.css'
import { ChatHistory } from './ChatHistory';
import Spinner from 'react-bootstrap/Spinner';
import ModelSelectionModal from './modals/ModelSelectionModal';
import { parseDataChunk, makeId, getColor, isClosed, generateHashId } from './utils';
import { Cooldown } from './Cooldown';
import ImageViewer from "react-simple-image-viewer";
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter'
import { coldarkDark } from 'react-syntax-highlighter/dist/esm/styles/prism'
import { ChatData, ChatHistoryInterface, ChatTitle, HistoryMessage } from './interfaces';

let abort = false;
let cancelScroll = false;
const userMessage = { user: 'Me', role: "user", icon: "M", color: "#20dc8e" };

const Chat = (props) => {
    const chatBlockRef = useRef<null | HTMLDivElement>(null);
    const bidsSpinWaitClosed = useRef(false);

    const [value, setValue] = useState("");
    const [isLoading, setIsLoading] = useState(true);
    const [messages, setMessages] = useState<any>([]);
    const [isOpen, setIsOpen] = useState(false);
    const [sessions, setSessions] = useState<any>();
    const [providersAvailability, setProvidersAvailability] = useState<any[]>([]);

    const [isSpinning, setIsSpinning] = useState(false);
    const [meta, setMeta] = useState({ budget: 0, supply: 0 });

    const [imagePreview, setImagePreview] = useState<string>();
    const [activeSession, setActiveSession] = useState<any>(undefined);

    const [chainData, setChainData] = useState<any>(null);
    const [chatData, setChatsData] = useState<ChatData[]>([]);

    const [openChangeModal, setOpenChangeModal] = useState(false);
    const [isReadonly, setIsReadonly] = useState(false);

    const [selectedBid, setSelectedBid] = useState<any>(null);
    const [selectedModel, setSelectedModel] = useState<any>(undefined);
    const [requiredStake, setRequiredStake] = useState<{ min: Number, max: number }>({ min: 0, max: 0 })
    const [balances, setBalances] = useState<{ eth: Number, mor: number }>({ eth: 0, mor: 0 });

    const [chat, setChat] = useState<ChatData | undefined>(undefined);

    const modelName = selectedModel?.Name || "Model";
    const isLocal = chat?.isLocal;

    const providerAddress = isLocal ? "(local)" : selectedBid?.Provider ? abbreviateAddress(selectedBid?.Provider, 6) : 'Unknown';
    const isDisabled = (!activeSession && !isLocal) || isReadonly;
    const isEnoughFunds = Number(balances.mor) > Number(requiredStake.min);
    const stakedFunds = activeSession ? (((activeSession.EndsAt - activeSession.OpenedAt) * activeSession.PricePerSecond) / 10 ** 18).toFixed(2) : 0;

    useEffect(() => {
        (async () => {
            console.time("LOAD")
            const [chainData, userSessions, chats] = await Promise.all([
                props.getModelsData(),
                props.getSessionsByUser(props.address),
                props.client.getChatHistoryTitles() as Promise<ChatTitle[]>]);

            setBalances(chainData.userBalances)
            setMeta(chainData.meta);
            setChainData(chainData)

            const mappedChatData = chats.reduce((res, item) => {
                const chatModel = chainData.models.find(x => x.Id == item.modelId);
                if (chatModel) {
                    res.push({
                        id: item.chatId,
                        title: item.title,
                        createdAt: new Date(item.createdAt * 1000),
                        modelId: item.modelId,
                        isLocal: item.isLocal,
                    })
                }
                return res;
            }, [] as ChatData[])
            setChatsData(mappedChatData);

            const sessions = userSessions.reduce((res, item) => {
                const sessionModel = chainData.models.find(x => x.Id == item.ModelAgentId);
                if (sessionModel) {
                    item.ModelName = sessionModel.Name;
                    res.push(item);
                }
                return res;
            }, []);
            setSessions(sessions);

            const openSessions = sessions.filter(s => !isClosed(s));

            const useLocalModelChat = () => {
                const localModel = (chainData?.models?.find((m: any) => m.isLocal));
                if (localModel) {
                    setSelectedModel(localModel);
                    setChat({ id: generateHashId(), createdAt: new Date(), modelId: localModel.Id, isLocal: true });
                }
            }

            if (!openSessions.length) {
                useLocalModelChat();
                console.timeEnd("LOAD")
                return;
            }

            const latestSession = openSessions[0];
            const latestSessionModel = (chainData.models.find((m: any) => m.Id == latestSession.ModelAgentId));

            if (!latestSessionModel) {
                useLocalModelChat();
                console.timeEnd("LOAD")
                return;
            }

            const openBid = await props.getBidInfo(latestSession.BidID)

            if (!openBid) {
                useLocalModelChat();
            }

            setSelectedModel(latestSessionModel);
            setSelectedBid(openBid);
            setActiveSession(latestSession);
            setChat({ id: generateHashId(), createdAt: new Date(), modelId: latestSessionModel.ModelAgentId });
            console.timeEnd("LOAD")
        })()
        .then(() => {
            setIsLoading(false);
        })
    }, [])

    useEffect(() => {
        if(!chainData)
            return;

        (async () => {
            const providersMap = chainData.providers.reduce((a, b) => ({ ...a, [b.Address.toLowerCase()]: b }), {});
            const modelsWithBids= (await Promise.all(
                chainData.models.map(async m => {
                    const id = m.Id;
                    if(m.isLocal){
                        return { id }
                    }
                    const bids = (await props.getBidsByModelId(id))
                        .map(b => ({ ...b, ProviderData: providersMap[b.Provider.toLowerCase()], Model: m }))
                        .filter(b => b.ProviderData);

                    if(!bids.length){
                        return null;
                    }

                    return { id, bids }
                })
            )).reduce((acc, next) => {
                if(!next) {
                    return acc;
                }
                const model = chainData.models.find(m => m.Id == next.id);
                return [...acc, { ...model, bids: next.bids}]
            }, []);
            
            setChainData({...chainData, models: modelsWithBids})
            bidsSpinWaitClosed.current = true;
        })();

        (async () => {
            const availabilityResults = await props.getProvidersAvailability(chainData.providers);
            setProvidersAvailability(availabilityResults);            
        })();

    }, chainData)

    const spinWaitForBids = async () => {
        if(bidsSpinWaitClosed.current)
            return;
        setIsLoading(true);
        while(!bidsSpinWaitClosed.current) {
            await new Promise(resolve => setTimeout(resolve, 300));
        }
        setIsLoading(true);
    }

    const toggleDrawer = async () => {
        spinWaitForBids();
        setIsOpen((prevState) => !prevState)
    }

    const scrollToBottom = (behavior: ScrollBehavior = "instant") => {
        if (!cancelScroll) {
            chatBlockRef.current?.scroll({ top: chatBlockRef.current.scrollHeight, behavior: behavior })
        }
    }

    const calculateAcceptableDuration = (pricePerSecond: number, balance: number, stakingInfo) => {
        const delta = 60; // 1 minute

        if (balance > requiredStake.max) {
            return 24 * 60 * 60; // 1 day in seconds
        }

        const targetDuration = Math.round((balance * Number(stakingInfo.budget)) / (Number(stakingInfo.supply) * pricePerSecond))

        if (targetDuration - delta < 5 * 60) {
            return 5 * 60;
        }

        return (targetDuration - (targetDuration % 60)) - delta;
    }

    const setSessionData = async (sessionId) => {
        const allSessions = await refreshSessions();
        const targetSessionData = allSessions.find(x => x.Id == sessionId);
        setActiveSession({ ...targetSessionData, sessionId });
        const targetModel = chainData.models.find(x => x.Id == targetSessionData.ModelAgentId)
        const targetBid = targetModel.bids.find(x => x.Id == targetSessionData.BidID);
        setSelectedBid(targetBid);
    }

    const onOpenSession = async (isReopen) => {
        setIsLoading(true);
        if (!isReopen) {
            setChat({ id: generateHashId(), createdAt: new Date(), modelId: selectedModel.Id });
        }

        const prices = selectedModel.bids.map(x => Number(x.PricePerSecond));
        const maxPrice = Math.max(...prices);
        const duration = calculateAcceptableDuration(maxPrice, Number(balances.mor), meta);

        console.log("open-session", duration);

        try {
            const openedSession = await props.onOpenSession({ modelId: selectedModel.Id, duration });
            if (!openedSession) {
                return;
            }
            await setSessionData(openedSession);
            return openedSession;
        }
        finally {
            setIsLoading(false);
        }
    }

    const loadChatHistory = async (chatId: string) => {
        try {
            const history: ChatHistoryInterface = await props.client.getChatHistory(chatId);
            const messages: HistoryMessage[] = [];

            const model = chainData.models.find((m) => m.Id == history.modelId);
            history.messages.forEach((m) => {
                const modelName = model.Name || "Model";

                const aiIcon = modelName.toUpperCase()[0];
                const aiColor = getColor(aiIcon);

                messages.push({ id: makeId(16), text: m.prompt.messages[0].content, user: userMessage.user, role: userMessage.role, icon: userMessage.icon, color: userMessage.color });
                messages.push({ id: makeId(16), text: m.response, user: modelName, role: "assistant", icon: aiIcon, color: aiColor, isImageContent: m.isImageContent, isVideoRawContent: m.isVideoRawContent });
            });
            setMessages(messages);
        }
        catch (e) {
            props.toasts.toast('error', 'Failed to load chat history');
        }
    }

    const refreshSessions = async () => {
        const sessions = (await props.getSessionsByUser(props.address)).reduce((res, item) => {
            const sessionModel = chainData.models.find(x => x.Id == item.ModelAgentId);
            if (sessionModel) {
                item.ModelName = sessionModel.Name;
                res.push(item);
            }
            return res;
        }, []);

        setSessions(sessions);

        return sessions;
    }

    const closeSession = async (sessionId: string) => {
        setIsLoading(true);
        await props.closeSession(sessionId);
        await refreshSessions();
        setIsLoading(false);

        if (activeSession.Id == sessionId) {
            const localModel = (chainData?.models?.find((m: any) => m.isLocal));
            if (localModel) {
                setSelectedModel(localModel);
                setChat({ id: generateHashId(), createdAt: new Date(), modelId: localModel.Id, isLocal: true });
            }
            setMessages([]);
        }
    }

    const selectChat = async (chatData: ChatData) => {
        console.log("select-session", chatData)

        const modelId = chatData.modelId;
        if (!modelId) {
            console.warn("Model ID is missed");
            return;
        }

        const selectedModel = chainData.isLocal ? chainData.models.find((m: any) => m.Id == modelId) : chainData.models.find((m: any) => m.Id == modelId && m.bids);
        setSelectedModel(selectedModel);
        setIsReadonly(false);

        setChat({ ...chatData })

        if (chatData.isLocal) {
            await loadChatHistory(chatData.id);
            return;
        }

        const openSessions = sessions.filter(s => !isClosed(s));
        // search open session by model ID
        const openSession = openSessions.find(s => s.ModelAgentId == modelId);
        setIsReadonly(!openSession);

        if (openSession) {
            setActiveSession(openSession);
            const activeBid = selectedModel.bids.find((b) => b.Id == openSession.BidID);
            setSelectedBid(activeBid);
        }
        else {
            setActiveSession(undefined);
            setSelectedBid(undefined);
        }

        await loadChatHistory(chatData.id);
        setTimeout(() => scrollToBottom("smooth"), 400);
    }

    const handleReopen = async () => {
        spinWaitForBids();
        setIsLoading(true);
        const newSessionId = await onOpenSession(true);
        setIsReadonly(false);
        console.log("Reopened session id: ", newSessionId)
    }

    const registerScrollEvent = (register) => {
        cancelScroll = false;
        const handler = (event: any) => {
            const isUp = event.wheelDelta ? event.wheelDelta > 0 : event.deltaY < 0;
            if (isUp) {
                cancelScroll = true;
            }
            else {
                if (!chatBlockRef?.current || !cancelScroll) {
                    return;
                }
                // Return scrolling if scrolled to div end 
                if ((chatBlockRef.current.offsetHeight + chatBlockRef.current.scrollTop) >= chatBlockRef.current.scrollHeight) {
                    cancelScroll = false;
                }
            }
        };

        if (register) {
            chatBlockRef?.current?.addEventListener('wheel', handler);
        }
        else {
            chatBlockRef?.current?.removeEventListener('wheel', handler);
        }
    }

    const call = async (message) => {
        let memoState = [...messages, { id: makeId(16), text: value, ...userMessage }];
        setMessages(memoState);
        scrollToBottom();

        const headers = {
            "Accept": "application/json"
        };
        if (isLocal) {
            headers["model_id"] = selectedModel.Id;
        } else {
            headers["session_id"] = activeSession.Id;
        }
        headers["chat_id"] = chat?.id;

        const incommingMessage = { role: "user", content: message };
        const payload = {
            stream: true,
            messages: [incommingMessage]
        };

        // If image take only last message
        const response = await fetch(`${props.config.chain.localProxyRouterUrl}/v1/chat/completions`, {
            method: 'POST',
            headers,
            body: JSON.stringify(payload)
        }).catch((e) => {
            console.log("Failed to send request", e)
            return null;
        });

        if (!response) {
            return;
        }

        if (!response.ok) {
            console.log("Failed", await response.json())
            props.toasts.toast('error', 'Failed to send prompt');
            return;
        }

        const textDecoder = new TextDecoder();

        if (!response.body) {
            console.error("Body is missed");
            return;
        }

        const reader = response.body.getReader()
        registerScrollEvent(true);

        const icon = modelName.toUpperCase()[0];
        const iconProps = { icon, color: getColor(icon) };
        try {
            while (true) {
                if (abort) {
                    await reader.cancel();
                    abort = false;
                }

                const { value, done } = await reader.read();
                if (done) {
                    setIsSpinning(false);
                    break;
                }

                const decodedString = textDecoder.decode(value, { stream: true });
                const parts = parseDataChunk(decodedString);
                parts.forEach(part => {
                    if (!part) {
                        return;
                    }

                    if (part.error) {
                        console.warn(part.error);
                        return;
                    }

                    if (typeof part === 'string') {
                        handleSystemMessage(part);
                        return;
                    }

                    const imageContent = part.imageUrl;
                    const videoRawContent = part.videoRawContent;

                    if (!part?.id && !imageContent && !videoRawContent) {
                        return;
                    }

                    let result: any[] = [];
                    const message = memoState.find(m => m.id == part.id);
                    const otherMessages = memoState.filter(m => m.id != part.id);
                    if (imageContent) {
                        result = [...otherMessages, { id: part.job, user: modelName, role: "assistant", text: imageContent, isImageContent: true, ...iconProps }];
                    }
                    if (videoRawContent) {
                        result = [...otherMessages, { id: part.job, user: modelName, role: "assistant", text: videoRawContent, isVideoRawContent: true, ...iconProps }];
                    }
                    else {
                        const text = `${message?.text || ''}${part?.choices[0]?.delta?.content || ''}`.replace("<|im_start|>", "").replace("<|im_end|>", "");
                        result = [...otherMessages, { id: part.id, user: modelName, role: "assistant", text: text, ...iconProps }];
                    }
                    memoState = result;
                    setMessages(result);
                    scrollToBottom();
                })
            }
        }
        catch (e) {
            props.toasts.toast('error', 'Something goes wrong. Try later.');
            console.error(e);
        }

        registerScrollEvent(false);
        return memoState;
    }

    const handleSystemMessage = (message) => {
        const openSessionEventMessage = "new session opened";
        const failoverTurnOnMessage = "provider failed, failover enabled"

        const renderMessage = (value) => {
            props.toasts.toast('info', value, {
                autoClose: 1500
            });
        }

        if (message.includes(openSessionEventMessage)) {
            const sessionId = message.split(":")[1].trim(); // new session opened: 0x123456
            setSessionData(sessionId).catch((err) => renderMessage(`Failed to load session data: ${err.message}`));
            renderMessage("Opening session with available provider...");
            return;
        }
        if (message.includes(failoverTurnOnMessage)) {
            renderMessage("Target provider unavailable. Applying failover policy...");
            return;
        }
        renderMessage(message);
        return;
    }

    const handleSubmit = () => {
        if (abort) {
            abort = false;
        }

        if (isSpinning) {
            abort = true;
            setIsSpinning(false);
            return;
        }

        if (!value) {
            return;
        }

        if (messages.length === 0 && chat) {
            const title = { ...chat, title: value };
            setChatsData([...chatData, title]);
        }

        setIsSpinning(true);
        call(value).finally(() => setIsSpinning(false));
        setValue("");
    }

    const deleteChatEntry = (id: string) => {
        props.client.deleteChatHistory(id).then(() => {
            const newChats = chatData.filter(x => x.id != id);
            setChatsData(newChats);
        }).catch(console.error);
    }

    const calculateStake = (pricePerSecond, durationInMin) => {
        const totalCost = pricePerSecond * durationInMin * 60;
        const stake = totalCost * Number(meta.supply) / Number(meta.budget);
        return stake;
    }

    const onCreateNewChat = ({ modelId, isLocal }) => {
        abort = true;
        setMessages([]);
        setActiveSession(undefined);
        setSelectedBid(undefined);
        setIsReadonly(false);
        setChat({ id: generateHashId(), createdAt: new Date(), modelId, isLocal });

        const selectedModel = isLocal 
            ? chainData.models.find((m: any) => m.Id == modelId) 
            : chainData.models.find((m: any) => m.Id == modelId && m.bids);

        setSelectedModel(selectedModel);

        if (isLocal) {
            setActiveSession(undefined);
            setSelectedBid(undefined);
            return;
        }

        const openSessions = sessions.filter(s => !isClosed(s));
        const openModelSession = openSessions.find(s => s.ModelAgentId == modelId);

        if (openModelSession) {
            const selectedBid = selectedModel.bids.find(b => b.Id == openModelSession.BidID && b.bids);
            setSelectedBid(selectedBid);
            setActiveSession(openModelSession)
            return;
        }

        const prices = selectedModel.bids.map(x => Number(x.PricePerSecond));
        const maxPrice = Math.max(...prices);

        setRequiredStake({ min: calculateStake(maxPrice, 5), max: calculateStake(maxPrice, 24 * 60) })
    }

    const wrapChangeTitle = async (data: { id, title }) => {
        await props.client.updateChatHistoryTitle(data);
    }

    return (
        <>
            {
                isLoading &&
                <LoadingCover>
                    <Spinner style={{ width: '5rem', height: '5rem' }} animation="border" variant="success" />
                </LoadingCover>
            }
            <Drawer
                open={isOpen}
                onClose={toggleDrawer}
                direction='right'
                className='history-drawer'
            >
                <ChatHistory
                    activeChat={chat}
                    open={isOpen}
                    chatData={chatData}
                    sessions={sessions}
                    deleteHistory={deleteChatEntry}
                    models={chainData?.models || []}
                    onSelectChat={selectChat}
                    refreshSessions={async () => {
                        setIsLoading(true);
                        await refreshSessions()
                        setIsLoading(false);
                    }}
                    onChangeTitle={wrapChangeTitle}
                    onCloseSession={closeSession} />
            </Drawer>
            <View>
                <ContainerTitle>
                    <TitleRow>
                        {/* <Title>Chat</Title> */}
                        <div className='d-flex' style={{ alignItems: 'center' }}>
                            <div className='d-flex model-selector'>
                                <div className='model-selector__info'>
                                    <h3>{isLocal ? "(local)" : providerAddress}</h3>
                                    {
                                        isLocal ?
                                            (
                                                <>
                                                    <span>0 MOR</span>
                                                </>
                                            )
                                            : (
                                                <>
                                                    <SubPriceLabel>{stakedFunds} MOR</SubPriceLabel>
                                                </>
                                            )
                                    }
                                </div>
                                {

                                    !isLocal && activeSession?.EndsAt && (
                                        <div className='model-selector__icons'>
                                            <Cooldown endDate={activeSession?.EndsAt} />
                                        </div>
                                    )
                                }

                            </div>
                            <BtnAccent
                                className='change-modal'
                                onClick={async () => {
                                    await spinWaitForBids();
                                    setOpenChangeModal(true);
                                } }>
                                <IconMessagePlus></IconMessagePlus> New chat
                            </BtnAccent>
                        </div>
                    </TitleRow>
                </ContainerTitle>
                <ChatTitleContainer>
                    <ChatAvatar>
                        <Avatar style={{ color: 'white' }} color={getColor(modelName.toUpperCase()[0])}>
                            {modelName.toUpperCase()[0]}
                        </Avatar>
                        <div style={{ marginLeft: '10px' }}>{modelName}</div>
                    </ChatAvatar>
                    {/* {
                        (selectedBid || isLocal) && <div>
                            <span style={{ color: 'white' }}>Provider:</span> {isLocal ? "(local)" : providerAddress}
                        </div>
                    } */}
                    <div>
                        <div onClick={toggleDrawer}>
                            <IconHistory size={"2.4rem"}></IconHistory>
                        </div>
                    </div>
                </ChatTitleContainer>

                {imagePreview && (
                    <ImageViewer
                        src={[imagePreview]}
                        onClose={() => setImagePreview("")}
                        disableScroll={false}
                        backgroundStyle={{
                            backgroundColor: "rgba(0,0,0,0.9)",
                            zIndex: 1000
                        }}
                        closeOnClickOutside={true}
                    />
                )}

                <Container>
                    <ChatBlock ref={chatBlockRef} className={!messages?.length ? 'createSessionMode' : null}>
                        {
                            messages?.length ? messages.map(x => (
                                <Message key={makeId(16)} message={x} onOpenImage={setImagePreview}></Message>
                            ))
                                : (!isLocal && !activeSession && !isLoading &&
                                    <div className='session-container' style={{ width: '400px' }}>
                                        {
                                            isEnoughFunds ?
                                                <>
                                                    <div className='session-title'>Staked funds will be reserved to start session</div>
                                                    <div className='session-title'>Session may last from 5 mins to 24 hours depending on available balance (min: {(Number(requiredStake.min) / 10 ** 18).toFixed(2)}, max: {(Number(requiredStake.max) / 10 ** 18).toFixed(2)} {props.symbol})</div>
                                                </> :
                                                <div className='session-title'>To start session required balance should be at least {(Number(requiredStake.min) / 10 ** 18).toFixed(2)} {props.symbol}</div>
                                        }
                                        <div>
                                            <BtnAccent
                                                data-modal="receive"
                                                data-testid="receive-btn"
                                                style={{ marginLeft: '0px' }}
                                                block={requiredStake.min}
                                                onClick={onOpenSession}
                                                disabled={!isEnoughFunds}
                                            >
                                                Start
                                            </BtnAccent>
                                        </div>
                                    </div>)
                        }
                    </ChatBlock>
                    <Control>
                        <CustomTextArrea
                            disabled={isDisabled}
                            onKeyPress={(e) => {
                                if (e.key === 'Enter') {
                                    e.preventDefault();
                                    handleSubmit();
                                }
                            }}
                            value={value}
                            onChange={ev => setValue(ev.target.value)}
                            placeholder={isReadonly ? "Session is closed. Chat in ReadOnly Mode" : "Ask me anything..."}
                            minRows={1}
                            maxRows={6} />
                        {
                            isReadonly
                                ? (<SendBtn onClick={handleReopen}>
                                    {isSpinning ? <Spinner animation="border" /> : <span>Reopen</span>}
                                </SendBtn>)
                                : (
                                    <SendBtn disabled={isDisabled} onClick={handleSubmit}>{
                                        isSpinning ? <Spinner animation="border" /> : <IconArrowUp size={"26px"}></IconArrowUp>
                                    }</SendBtn>
                                )
                        }
                    </Control>
                </Container>
            </View>
            <ModelSelectionModal
                models={(chainData as any)?.models}
                isActive={openChangeModal}
                symbol={props.symbol}
                providersAvailability={providersAvailability}
                onChangeModel={(eventData) => {
                    onCreateNewChat(eventData);
                }}
                handleClose={() => setOpenChangeModal(false)} />
        </>
    )
}

const renderMessage = (message, onOpenImage) => {
    if (message.isImageContent) {
        return (<MessageBody>{<ImageContainer src={message.text} onClick={() => onOpenImage(message.text)} />}</MessageBody>)
    }

    if (message.isVideoRawContent) {
        return (<MessageBody><VideoContainer><video controls src={`${message.text}`}/></VideoContainer></MessageBody>)
    }

    return (
        <MessageBody>
            <Markdown
                children={message.text}
                components={{
                    code(props) {
                        const { children, className, node, ...rest } = props
                        const match = /language-(\w+)/.exec(className || '')
                        return match ? (
                            <SyntaxHighlighter
                                {...rest}
                                PreTag="div"
                                children={String(children).replace(/\n$/, '')}
                                language={match[1]}
                                style={coldarkDark}
                            />
                        ) : (
                            <code {...rest} className={className}>
                                {children}
                            </code>
                        )
                    }
                }}
            />
        </MessageBody>)
};

const Message = ({ message, onOpenImage }) => {
    return (
        <div style={{ display: 'flex', margin: '12px 0 28px 0' }}>
            <Avatar color={message.color}>
                {message.icon}
            </Avatar>
            <div>
                <AvatarHeader>{message.user}</AvatarHeader>
                {
                    renderMessage(message, onOpenImage)
                }
            </div>
        </div>)
}

export default withRouter(withChatState(Chat));
