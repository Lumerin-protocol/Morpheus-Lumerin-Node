package test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/aiengine"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	api "github.com/sashabaranov/go-openai"
)

func AiEngine_Prompt(t *testing.T) {
	aiEngine := aiengine.NewAiEngine("http://localhost:11434/v1", "", lib.NewTestLogger())
	req := &api.ChatCompletionRequest{
		Model:     "llama2",
		MaxTokens: 100,
		Messages: []api.ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Hello, I am a test user"},
		}, // This is a test
	}
	resp, err := aiEngine.Prompt(context.Background(), req)
	if err != nil {
		t.Errorf("Prompt error: %v", err)
	}
	fmt.Printf("Prompt response: %+v\n", resp)
}

func TestAiEngine_PromptStream(t *testing.T) {
	aiEngine := aiengine.NewAiEngine("http://localhost:11434/v1", "", lib.NewTestLogger())
	req := &api.ChatCompletionRequest{
		Model: "llama2",
		// MaxTokens: 100,
		Stream: true,
		Messages: []api.ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Hello, I am a test user",
			},
		}, // This is a test
	}

	choicesChannel := make(chan api.ChatCompletionStreamChoice)
	choices := []api.ChatCompletionStreamChoice{}

	resp, err := aiEngine.PromptStream(context.Background(), req, func(response *api.ChatCompletionStreamResponse) error {
		choices = append(choices, response.Choices...)

			if response.Choices[0].Delta.Content == "" {
				return errors.New("empty response")
			}
			// fmt.Printf("chunk - no error")
			return nil
		})
	}()

	timeout := time.After(60 * time.Second)
outerLoop:
	for {
		select {
		case choice := <-choicesChannel:
			choices = append(choices, choice)

			if len(choices) >= 1 {
				break outerLoop
			}
		case <-timeout:
			break outerLoop
		}
	}

	// if err != nil {
	// 	t.Errorf("error: %v", err)
	// 	// fmt.Println("error: ", err)
	// }

	if len(choices) == 0 {
		t.Errorf("invalid response: %v", choices)
	}

	// if choices[len(choices)-1].Delta.Content != "[DONE]" {
	// 	t.Errorf("invalid end of stream response: %s", choices[0].Delta.Content)
	// }

	// content := concatenateDeltaContent(choices)

	// if content == "" {
	// 	t.Errorf("content is empty")
	// }

	// if content == "" {
	// 	t.Errorf("content is empty")
	// }

	// if strings.Contains(content, "Hello ") {
	// 	t.Errorf("content is invalid: %v", content)
	// }
}

func concatenateDeltaContent(choices []api.ChatCompletionStreamChoice) string {
	var concatenatedContent string
	for _, choice := range choices {
		concatenatedContent += choice.Delta.Content
	}
	return concatenatedContent
}
