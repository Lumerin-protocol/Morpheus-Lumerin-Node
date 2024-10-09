export interface ChatTitle {
    chatId: string;
    title: string;
    createdAt: number; // timestamp in seconds
    modelId: string;
}

export interface ChatData {
    id: string;
    title?: string;
    createdAt: Date;
    modelId: string;
    isLocal?: boolean;
}

export interface HistoryMessage {
    id: string;
    text: string;
    user: string;
    role: string;
    icon: string;
    color: string;
    isImageContent?: boolean;
}

export interface ChatHistoryInterface {
    title: string;
    modelId: string;
    messages: ChatMessage[];
}

export interface ChatMessage {
    response: string;
    prompt: ChatPrompt;
    promptAt: number;
    responseAt: number;
    isImageContent?: boolean;
}

export interface ChatPrompt {
    model: string;
    messages: {
        role: string;
        content: string;
    }[];
}

