package client

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func TestGetProviders(t *testing.T) {
	client := NewApiGatewayClient("http://localhost:8080", http.DefaultClient)
	res, err := client.GetAllProviders(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("providers: ", res)

}

func TestCreateProvider(t *testing.T) {

	master, err := hdwallet.NewKey()

	if err != nil {
		t.Fatal(err)
	}

	wallet, err := master.GetWallet()

	if err != nil {
		t.Fatal(err)
	}

	address, err := wallet.GetAddress()

	if err != nil {
		t.Fatal(err)
	}
	client := NewApiGatewayClient("http://localhost:8080", http.DefaultClient)

	res, err := client.CreateNewProvider(context.Background(), address, 1, "localhost:8080")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("providers: ", res)
}
 
func TestHealthCheck(t *testing.T) {
	client := NewApiGatewayClient("http://localhost:8080", &http.Client{})
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
	client := NewApiGatewayClient("http://localhost:8080", http.DefaultClient)
	_, err := client.GetProxyRouterFiles(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}

func SkipTestCreateAndStreamChatCompletionMessage(t *testing.T) {
	client := NewApiGatewayClient("http://localhost:8080", http.DefaultClient)

	prompt := "What is the meaning of life?"
	messages := []*ChatCompletionMessage{
		{
			Content: "The meaning of life is 42.",
		},
	}

	res, err := client.PromptStream(context.Background(), prompt, messages, func(msg ChatCompletionStreamResponse) error {
		fmt.Println("chat completion message: ", msg)
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("chat completion result: ", res)

	mapResult := res.(map[string]interface{})

	if mapResult["choices"] == nil {
		t.Fatal("invalid chat completion result")
	}

	if mapResult["choices"].([]interface{})[0].(map[string]interface{})["text"] == "" {
		t.Fatal("invalid chat completion text result")
	}
}

func SkipTestCreateChatCompletionMessage(t *testing.T) {

	os.Setenv("OPENAI_BASE_URL", "http://localhost:8080/v1")

	client := NewApiGatewayClient("http://localhost:8080", http.DefaultClient)

	prompt := "What is the meaning of life?"
	messages := []ChatCompletionMessage{
		{
			Content: "The meaning of life is 42.",
		},
	}

	res, err := client.Prompt(context.Background(), prompt, messages)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("chat completion result: ", res)

	mapResult := res.(map[string]interface{})

	if mapResult["choices"] == nil {
		t.Fatal("invalid chat completion result")
	}

	if mapResult["choices"].([]interface{})[0].(map[string]interface{})["text"] == "" {
		t.Fatal("invalid chat completion text result")
	}
}
