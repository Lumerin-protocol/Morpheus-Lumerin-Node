package chat

import (
	"context"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/api-gateway/client"
	"github.com/sashabaranov/go-openai"
)

func (m model) sendChat(prompt string, streamResponse client.CompletionCallback) error {
	ctx := context.Background()

	m.openaiRequest.Messages = append(m.openaiRequest.Messages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: prompt,
	})

	_, err := client.RequestChatCompletionStream(ctx, &m.openaiRequest, streamResponse, m.config.SessionId)

	if err != nil {
		return err
	}

	return nil
}
