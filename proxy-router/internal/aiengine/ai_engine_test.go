package aiengine

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	api "github.com/sashabaranov/go-openai"
)

func AiEngine_Prompt(t *testing.T) {
	os.Setenv("OPENAI_BASE_URL", "http://localhost:11434/v1")

	aiEngine := NewAiEngine()
	ctx := context.Background()
	req := &api.ChatCompletionRequest{
		Model:     "llama2",
		MaxTokens: 100,
		Messages: []api.ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Hello, I am a test user"},
		}, // This is a test
	}
	resp, err := aiEngine.Prompt(ctx, req)
	if err != nil {
		t.Errorf("Prompt error: %v", err)
	}
	fmt.Printf("Prompt response: %+v\n", resp)
}

func TestAiEngine_PromptStream(t *testing.T) {
	os.Setenv("OPENAI_BASE_URL", "http://localhost:11434/v1")

	ctx := context.Background()
	req := &api.ChatCompletionRequest{
		Model:     "llama2",
		MaxTokens: 100,
		Stream:    true,
		Messages: []api.ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Hello, I am a test user"},
		}, // This is a test
	}

	choices := make([]api.ChatCompletionStreamChoice, 0)

	resp, err := RequestChatCompletionStream(ctx, req, func(response api.ChatCompletionStreamResponse) error {
		choices = append(choices, response.Choices...)

		if response.Choices[0].Delta.Content == "" {
			return errors.New("empty response")
		}

		return nil
	})

	if err != nil {
		t.Errorf("error: %v", err)
		fmt.Println("error: ", err)
	}

	if resp == nil || resp.Choices == nil {
		t.Errorf("invalid nil response")
	}

	if resp.Choices[0].Delta.Content != "[DONE]" {
		t.Errorf("invalid end of stream response: %s", resp.Choices[0].Delta.Content)
	}

	content := concatenateDeltaContent(choices)

	if content == "" {
		t.Errorf("content is empty")
	}

	if content == "" {
		t.Errorf("content is empty")
	}

	if strings.Contains(content, "Hello there! ") {
		t.Errorf("content is invalid")
	}
}

func concatenateDeltaContent(choices []api.ChatCompletionStreamChoice) string {
	var concatenatedContent string
	for _, choice := range choices {
		concatenatedContent += choice.Delta.Content
	}
	return concatenatedContent
}
