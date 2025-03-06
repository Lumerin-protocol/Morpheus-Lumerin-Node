export type AgentUserRes = {
  agents: AgentUser[]
}

export type AgentUser = {
  allowances: Record<'ETH' | string, string>
  isConfirmed: boolean
  password: string
  perms: string[]
  username: string
}

export type AgentTxRes = {
  txHashes: string[]
  nextCursor: string
}

export type AgentAllowanceRequestsRes = {
  requests: AgentAllowanceRequest[]
}

export type AgentAllowanceRequest = {
  username: string
  token: string
  allowance: string
}

export type ChatTitle = {
  chatId: string
  createdAt: number
  isLocal: boolean
  modelId: string
  title: string
}

export type ChatHistory = {
  title: string
  modelId: string
  sessionId: string
  messages: ChatMessage[]
}

export type ChatMessage = {
  response: string
  prompt: {
    messages: {
      role: string
      content: string
    }[]
  }
  promptAt: number
  responseAt: number
  isImageContent: boolean
  isVideoRawContent: boolean
}

export type ResultResponse = {
  result: boolean
}
