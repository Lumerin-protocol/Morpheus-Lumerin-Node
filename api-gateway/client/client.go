package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/sashabaranov/go-openai"
)

type ChatCompletionMessage = openai.ChatCompletionMessage

func NewApiGatewayClient(baseURL string, httpClient *http.Client) *ApiGatewayClient {
	return &ApiGatewayClient{
		BaseURL:    baseURL,
		HttpClient: httpClient,
		OpenAiClient: openai.NewClientWithConfig(openai.ClientConfig{
			BaseURL:    baseURL + "/v1",
			APIType:    openai.APITypeOpenAI,
			HTTPClient: httpClient,
		}),
	}
}

type ApiGatewayClient struct {
	BaseURL      string
	HttpClient   *http.Client
	OpenAiClient *openai.Client
}

// Helper function to make GET requests
func (c *ApiGatewayClient) getRequest(ctx context.Context, endpoint string, result interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+endpoint, nil)
	if err != nil {
		return err
	}
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		fmt.Println("http client error: ", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("unexpected status code: %d", resp.StatusCode)
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	// line, _, _ := bufio.NewReader(resp.Body).ReadLine()

	// fmt.Println("response body line 1: ", string(line))
	return json.NewDecoder(resp.Body).Decode(result)
}

// Helper function to make POST requests
func (c *ApiGatewayClient) postRequest(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(reqBody)

	fmt.Println("post path: ", c.BaseURL+endpoint)
	req, err := http.NewRequestWithContext(ctx, "POST", c.BaseURL+endpoint, reader)
	if err != nil {
		return err
	}
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("unexpected status code: %d", resp.StatusCode))
	}

	if result == nil {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

func (c *ApiGatewayClient) streamChat(ctx context.Context, path string, req *ChatCompletionRequest, result *ChatCompletionStreamResponse, flush CompletionCallback) error {

	RequestChatCompletionStream(ctx, req, flush)

	return nil
}

func (c *ApiGatewayClient) GetAllowance(ctx context.Context, spender string) (interface{}, error) {
	var result map[string]interface{}
	endpoint := fmt.Sprintf("/blockchain/allowance?spender=%s", spender)
	err := c.getRequest(ctx, endpoint, &result)
	return result, err
}

func (c *ApiGatewayClient) ApproveAllowance(ctx context.Context, spender string, amount uint64) (interface{}, error) {
	var result map[string]interface{}
	endpoint := fmt.Sprintf("/blockchain/allowance?spender=%s&amount=%d", spender, amount)
	err := c.postRequest(ctx, endpoint, nil, &result)
	return result, err
}

func (c *ApiGatewayClient) GetProxyRouterConfig(ctx context.Context) (interface{}, error) {
	var result map[string]interface{}
	err := c.getRequest(ctx, "/config", &result)
	return result, err
}

func (c *ApiGatewayClient) GetProxyRouterFiles(ctx context.Context) (interface{}, error) {
	var result map[string]interface{}
	err := c.getRequest(ctx, "/files", &result)
	if err != nil {
		return nil, fmt.Errorf("internal error: %v; http status: %v", err, http.StatusInternalServerError)
	}
	return result, nil
}

func (c *ApiGatewayClient) HealthCheck(ctx context.Context) (interface{}, error) {
	var result map[string]interface{}
	err := c.getRequest(ctx, "/healthcheck", &result)

	return result, err
}

func (c *ApiGatewayClient) InitiateSession(ctx context.Context) (interface{}, error) {
	var result map[string]interface{}
	err := c.postRequest(ctx, "/proxy/sessions/initiate", nil, &result)

	if err != nil {
		return nil, fmt.Errorf("internal error: %v; http status: %v", err, http.StatusInternalServerError)
	}

	return result, nil
}

func (c *ApiGatewayClient) SessionPrompt(ctx context.Context, prompt string) (interface{}, error) {
	var result map[string]interface{}
	err := c.postRequest(ctx, "/proxy/sessions/:id/prompt", nil, &result)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return result, nil
}

func (c *ApiGatewayClient) Prompt(ctx context.Context, message string, history []ChatCompletionMessage) (interface{}, error) {

	request := &openai.ChatCompletionRequest{
		Messages: append(history, ChatCompletionMessage{
			Role:    "user",
			Content: message,
		}),
		Stream: false,
		Model:  "llama2",
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeText,
		},
	}

	return c.OpenAiClient.CreateChatCompletion(ctx, *request)
}

func (c *ApiGatewayClient) PromptStream(ctx context.Context, message string, history []*ChatCompletionMessage, flush CompletionCallback) (interface{}, error) {
	var messages []ChatCompletionMessage

	for i := 0; i < len(history); i++ {
		historyItem := history[i]

		messages = append(messages, ChatCompletionMessage{
			Role:    historyItem.Role,
			Content: historyItem.Content,
		})
	}

	request := &openai.ChatCompletionRequest{

		Messages: messages,
		Stream:   true,
	}

	return RequestChatCompletionStream(ctx, request, flush)
}

func (c *ApiGatewayClient) GetLatestBlock(ctx context.Context) (result uint64, err error) {

	err = c.getRequest(ctx, "/blockchain/latestBlock", &result)
	return result, err
}

func (c *ApiGatewayClient) GetAllProviders(ctx context.Context) (result interface{}, err error) {

	err = c.getRequest(ctx, "/blockchain/providers", &result)

	if err != nil {
		return nil, fmt.Errorf("internal error: %v; http status: %v", err, http.StatusInternalServerError)
	}

	return result, nil
}

func (c *ApiGatewayClient) CreateNewProvider(ctx context.Context, address string, addStake uint64, endpoint string) (result []string, err error) {

	request := struct {
		address  string
		addStake uint64
		endpoint string
	}{address, addStake, endpoint}

	err = c.postRequest(ctx, "/blockchain/providers", &request, &result)

	if err != nil {
		return nil, fmt.Errorf("internal error: %v; http status: %v", err, http.StatusInternalServerError)
	}

	return result, nil
}

func (c *ApiGatewayClient) GetAllModels(ctx context.Context) (result interface{}, err error) {

	err = c.getRequest(ctx, "/blockchain/models", &result)

	if err != nil {
		return nil, fmt.Errorf("internal error: %v; http status: %v", err, http.StatusInternalServerError)
	}

	return result, nil
}

func (c *ApiGatewayClient) GetBidsByProvider(ctx context.Context, providerAddr string, offset *big.Int, limit uint8) (bids []string, err error) {

	endpoint := fmt.Sprintf("/blockchain/providers/%s/bids?offset=%s&limit=%d", providerAddr, offset.String(), limit)
	err = c.getRequest(ctx, endpoint, &bids)
	if err != nil {
		return nil, fmt.Errorf("internal error: %v; http status: %v", err, http.StatusInternalServerError)
	}
	return bids, err
}

func (c *ApiGatewayClient) GetBidsByModelAgent(ctx context.Context, modelAgentId [32]byte, offset *big.Int, limit uint8) (result []string, err error) {

	endpoint := fmt.Sprintf("/blockchain/models/%x/bids?offset=%s&limit=%d", modelAgentId, offset.String(), limit)
	err = c.getRequest(ctx, endpoint, &result)

	if err != nil {
		return nil, fmt.Errorf("internal error: %v; http status: %v", err, http.StatusInternalServerError)
	}

	return result, err
}

func (c *ApiGatewayClient) OpenSession(ctx context.Context) (err error) {

	err = c.postRequest(ctx, "/blockchain/sessions", nil, nil)

	if err != nil {
		return fmt.Errorf("internal error: %v; http status: %v", err, http.StatusInternalServerError)
	}

	return nil
}

func (c *ApiGatewayClient) CloseSession(ctx context.Context) error {
	err := c.postRequest(ctx, "/blockchain/sessions/:id/close", nil, nil)

	if err != nil {
		return fmt.Errorf("internal error: %v; http status: %v", err, http.StatusInternalServerError)
	}

	return nil
}
