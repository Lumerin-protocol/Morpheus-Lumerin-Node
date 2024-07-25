package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/sashabaranov/go-openai"
)

func SkipTestGetProviders(t *testing.T) {
	client := NewApiGatewayClient("http://localhost:8082", http.DefaultClient)
	res, err := client.GetAllProviders(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("providers: ", res)

}

func SkipTestCreateProvider(t *testing.T) {

	mnemonic, err := hdwallet.NewMnemonic(128)

	if err != nil {
		t.Fatal(err)
	}

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)

	if err != nil {
		t.Fatal(err)
	}

	account := wallet.Accounts()[0]

	address := account.Address.Hex()

	if err != nil {
		t.Fatal(err)
	}
	client := NewApiGatewayClient("http://localhost:8082", http.DefaultClient)

	res, err := client.CreateNewProvider(context.Background(), address, 1, "localhost:8082")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("providers: ", res)
}

func SkipTestHealthCheck(t *testing.T) {
	client := NewApiGatewayClient("http://localhost:8082", &http.Client{})
	res, err := client.HealthCheck(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("healthcheck result: ", res)

	mapResult := res.(map[string]interface{})

	if mapResult["status"] != "healthy" {
		t.Fatal("healthcheck failed")
	}

	if mapResult["version"] == "" {
		t.Fatal("invalid version result")
	}

	if mapResult["uptime"] == "" {
		t.Fatal("invalid uptime result")
	}
}

func SkipTestFiles(t *testing.T) {
	client := NewApiGatewayClient("http://localhost:8082", http.DefaultClient)
	_, err := client.GetProxyRouterFiles(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}

func SkipTestCreateAndStreamChatCompletionMessage(t *testing.T) {
	client := NewApiGatewayClient("http://localhost:8082", http.DefaultClient)

	prompt := "What is the meaning of life?"
	messages := []*ChatCompletionMessage{
		{
			Content: "The meaning of life is 42.",
			Role:    "user",
		},
	}

	results := make([]*ChatCompletionStreamResponse, 0)

	_, err := client.PromptStream(context.Background(), prompt, messages, func(msg *ChatCompletionStreamResponse) error {
		results = append(results, msg)
		fmt.Println("msg: ", msg)
		return nil
	})

	if err != nil && err != io.EOF {
		t.Fatal(err)
	}

	fmt.Println("chat completion result: ", results)

	if len(results) == 0 {
		t.Fatal("invalid chat completion result: ", results)
	}

	if results[0].Choices[0].Delta.Content == "" {
		t.Fatal("invalid chat completion text result")
	}
}

func SkipTestCreateChatCompletionMessage(t *testing.T) {

	os.Setenv("OPENAI_BASE_URL", "http://localhost:8082/v1")

	client := NewApiGatewayClient("http://localhost:8082", http.DefaultClient)

	prompt := "What is the meaning of life?"
	messages := []ChatCompletionMessage{
		{
			Role:    "user",
			Content: "The meaning of life is 42.",
		},
	}

	res, err := client.Prompt(context.Background(), prompt, messages)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("chat completion result: ", res)

	result := res.(openai.ChatCompletionResponse)

	if result.Choices == nil || len(result.Choices) == 0 {
		t.Fatal("invalid chat completion result")
	}

	if result.Choices[0].Message.Content == "" {
		t.Fatal("invalid chat completion text result")
	}
}

func TestRequestChatCompletionStream(t *testing.T) {

	os.Setenv("OPENAI_BASE_URL", "http://localhost:8082/v1")

	// client := NewApiGatewayClient("http://localhost:11434", http.DefaultClient)

	request := &openai.ChatCompletionRequest{
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Hello, I am a test user",
			},
		},
		Model:     "llama2",
		Stream:    true,
	}

	errChan := make(chan error)
	choicesChannel := make(chan openai.ChatCompletionStreamChoice)
	choices := []openai.ChatCompletionStreamChoice{}

	go func(choicesChannel chan openai.ChatCompletionStreamChoice, errChan chan error) {
		_, err := RequestChatCompletionStream(context.Background(), request, func(response *ChatCompletionStreamResponse) error {
			
			// fmt.Printf("chunk: %+v", response)
			choicesChannel <- response.Choices[0]

			if response.Choices[0].Delta.Content == "" {
				return errors.New("empty response")
			}
			// fmt.Printf("chunk - no error")
			return nil
		})

		if err != nil {
			errChan <- err
			return
		}
	}(choicesChannel, errChan)

	// if err != nil {
	// 	t.Fatal(err)
	// }

	timeout := time.After(60 * time.Second)
outerLoop:
	for {
		select {
		case err := <-errChan:
			t.Fatal(err)
			return
		case choice := <-choicesChannel:
			choices = append(choices, choice)

			if len(choices) >= 1 {
				break outerLoop
			}
		case <-timeout:
			break outerLoop
		}
	}

	if len(choices) == 0 {
		t.Errorf("invalid response: %v", choices)
	}
}