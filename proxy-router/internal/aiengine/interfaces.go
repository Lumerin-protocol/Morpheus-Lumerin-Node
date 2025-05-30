package aiengine

import (
	"context"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/chatstorage/genericchatstorage"
	"github.com/sashabaranov/go-openai"
)

type AIEngineStream interface {
	Prompt(ctx context.Context, prompt *openai.ChatCompletionRequest, cb genericchatstorage.CompletionCallback) error
	AudioTranscription(ctx context.Context, prompt *genericchatstorage.AudioTranscriptionRequest, base64Audio string, cb genericchatstorage.CompletionCallback) error
	ApiType() string
}

type ModelParameters map[string]string
