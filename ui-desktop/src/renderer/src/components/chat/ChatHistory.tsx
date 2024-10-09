import { IconPencil, IconTrash, IconSearch, IconCheck, IconX } from "@tabler/icons-react";
import { abbreviateAddress } from '../../utils';
import { isClosed } from './utils';
import Tab from 'react-bootstrap/Tab';
import Tabs from 'react-bootstrap/Tabs';
import Form from 'react-bootstrap/Form';
import InputGroup from 'react-bootstrap/InputGroup';
import Badge from 'react-bootstrap/Badge';
import { SearchContainer } from '../contracts/modals/CreateContractModal.styles';
import * as components from './ChatHistory.styles';
import { useEffect, useState } from "react";
import { ChatData } from "./interfaces";

interface ChatHistoryProps {
    open: boolean,
    onCloseSession: (id: string) => void;
    onSelectChat: (chat: ChatData) => void;
    refreshSessions: () => void;
    deleteHistory: (chatId: string) => void
    onChangeTitle: (data: { id: string, title: string }) => Promise<void>;
    sessions: any[];
    models: any[];
    activeChat: any,
    chatData: ChatData[]
}

const HisotryEntry = ({ entry, deleteHistory, onSelectChat, isActive, onChangeTitle }) => {
    const [isEdit, setIsEdit] = useState<boolean>(false);
    const [title, setTitle] = useState<string>(entry?.title || "");

    const wrapDelete = (e, id) => {
        e.stopPropagation();
        deleteHistory(id);
    }

    const changetTitle = (e, chatId) => {
        setIsEdit(!isEdit);
        e.stopPropagation();
        onChangeTitle({ id: chatId, title });
    }

    return (
        <components.HistoryEntryTitle data-active={isActive ? true : undefined} onClick={() => onSelectChat(entry)}>
            {
                !isEdit ?
                    (
                        <>
                            <span className="title" style={{ width: isActive ? "75%" : undefined }}>{title}</span>
                            {
                                isActive && (
                                    <components.IconsContainer>
                                        <IconPencil onClick={(e) => {
                                            e.stopPropagation();
                                            setIsEdit(!isEdit);
                                        }} style={{ marginRight: '1.5rem' }} size={22} />
                                        <IconTrash size={22} onClick={(e) => wrapDelete(e, entry.id)} />
                                    </components.IconsContainer>
                                )
                            }

                        </>
                    ) :
                    (
                        <components.ChangeTitleContainer>
                            <InputGroup onClick={(e) => e.stopPropagation()}>
                                <Form.Control
                                    type="text"
                                    value={title}
                                    onChange={(e) => setTitle(e.target.value)}
                                />
                            </InputGroup>
                            <components.IconsContainer>
                                <IconCheck onClick={(e) => {
                                    changetTitle(e, entry.id)
                                }} style={{ margin: '0 1.5rem' }} size={22} />
                                <IconX size={22} onClick={(e) => {
                                    e.stopPropagation();
                                    setTitle(entry?.title || "");
                                    setIsEdit(false);
                                }} />
                            </components.IconsContainer>
                        </components.ChangeTitleContainer>
                    )
            }
        </components.HistoryEntryTitle>)
}

export const ChatHistory = (props: ChatHistoryProps) => {
    const sessions = props.sessions;
    const [search, setSearch] = useState<string | undefined>();

    useEffect(() => {
        setSearch("");
    }, [props.open])

    const renderTitlesGroup = (items: ChatData[]) => {
        return items.map(t => {
            return (<HisotryEntry
                onChangeTitle={props.onChangeTitle}
                isActive={props.activeChat?.id == t.id}
                key={t.id}
                entry={t}
                deleteHistory={props.deleteHistory}
                onSelectChat={props.onSelectChat} />)
        })
    }

    const getGroupHistory = <T extends { createdAt: Date },>(items: T[]) => {
        const getPreviousDate = (shift) => {
            const d = new Date();
            d.setDate(d.getDate() - shift);
            return d;
        }
        const source = (items || []).filter(i => i.createdAt).sort((a, b) => b.createdAt.getTime() - a.createdAt.getTime());
        const result: {
            today: T[],
            last7: T[],
            last30: T[],
            older: T[],
        } = {
            today: [],
            last7: [],
            last30: [],
            older: [],
        };
        const yesterday = getPreviousDate(1);
        const last7 = getPreviousDate(7);
        const last30 = getPreviousDate(30);
        for (const item of source) {
            if (item.createdAt > yesterday) {
                result.today.push(item);
            }
            else if (item.createdAt > last7) {
                result.last7.push(item);
            }
            else if (item.createdAt > last30) {
                result.last30.push(item);
            }
            else {
                result.older.push(item);
            }
        }
        return result;
    }

    const filterdModels = props.chatData && search ? props.chatData.filter(m => m.title?.includes(search)) : (props.chatData || []);
    const groupedItems = getGroupHistory<ChatData>(filterdModels);

    return (
        <components.Container>
            <div className='history-scroll-block'>

                <Tabs
                    defaultActiveKey="history"
                    id="uncontrolled-tab-example"
                    className="mb-3"
                >
                    <Tab eventKey="history" title="History">
                        <SearchContainer>
                            <InputGroup style={{ marginBottom: '15px' }}>
                                <InputGroup.Text>
                                    <IconSearch />
                                </InputGroup.Text>
                                <Form.Control
                                    type="text"
                                    placeholder="Search..."
                                    value={search}
                                    onChange={(e) => setSearch(e.target.value)}
                                />
                            </InputGroup>
                        </SearchContainer>
                        <div className='history-block'>
                            {
                                groupedItems.today.length ? (
                                    <>
                                        <span>Today</span>
                                        {renderTitlesGroup(groupedItems.today)}
                                    </>
                                ) : null
                            }
                            {
                                groupedItems.last7.length ? (
                                    <>
                                        <span>Previous 7 Days</span>
                                        {renderTitlesGroup(groupedItems.last7)}
                                    </>
                                ) : null
                            }
                            {
                                groupedItems.last30.length ? (
                                    <>
                                        <span>Previous 30 Days</span>
                                        {renderTitlesGroup(groupedItems.last30)}
                                    </>
                                ) : null
                            }
                            {
                                groupedItems.older.length ? (
                                    <>
                                        <span>Older</span>
                                        {renderTitlesGroup(groupedItems.older)}
                                    </>
                                ) : null
                            }
                        </div>
                    </Tab>
                    <Tab eventKey="sessions" title="Sessions">
                        {
                            sessions?.length ? (
                                sessions.map(a => {
                                    const model = props.models.find(x => x.Id == a.ModelAgentId);

                                    return (
                                        <components.HistoryEntryContainer key={a.Id}>
                                            <div>
                                                {!isClosed(a) ?
                                                    <components.FlexSpaceBetween>
                                                        <Badge bg="success">Active</Badge>
                                                        <components.CloseBtn onClick={() => props.onCloseSession(a.Id)}>Close</components.CloseBtn>
                                                    </components.FlexSpaceBetween> : null}
                                            </div>
                                            <components.HistoryItem>
                                                <components.ModelName data-rh={abbreviateAddress(a.Id, 3)} data-rh-negative>{model?.Name}</components.ModelName>
                                                <components.Duration>{((a.EndsAt - a.OpenedAt) / 60).toFixed(0)} min</components.Duration>
                                            </components.HistoryItem>
                                        </components.HistoryEntryContainer>
                                    )
                                }
                                ))
                                : <div>You have not any sessions</div>
                        }
                    </Tab>
                </Tabs>
            </div>
        </components.Container>
    )
}