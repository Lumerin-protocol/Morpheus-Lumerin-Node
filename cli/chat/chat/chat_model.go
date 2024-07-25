package chat

import (
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/cli/chat/common"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/sashabaranov/go-openai"
)

type model struct {
	config   common.Config
	err      error
	messages []string
	messageChunks []string
	textarea textarea.Model
	viewport viewport.Model

	openaiClient  *openai.Client
	openaiRequest openai.ChatCompletionRequest

	completionChunkSub       chan string // receive chat stream chunks
}