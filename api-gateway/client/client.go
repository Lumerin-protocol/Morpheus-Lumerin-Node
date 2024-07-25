package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

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
	// fmt.Printf("body: %+v", body)
	reqBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(reqBody)

	// fmt.Println("post path: ", c.BaseURL+endpoint)
	req, err := http.NewRequestWithContext(ctx, "POST", c.BaseURL+endpoint, reader)
	if err != nil {
		return err
	}
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	// fmt.Printf("%+v", resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if resp == nil {
		return nil
	}

	if result != nil {
		err = json.NewDecoder(resp.Body).Decode(result)
	}

	// fmt.Println("result: ", result)

	return err
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
		return nil, fmt.Errorf("internal error: %v", err)
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
		return nil, fmt.Errorf("internal error: %v", err)
	}

	return result, nil
}

func (c *ApiGatewayClient) SessionPrompt(ctx context.Context, prompt string, sessionId string) (interface{}, error) {
	var result map[string]interface{}
	err := c.postRequest(ctx, fmt.Sprintf("/proxy/sessions/%s/prompt", sessionId), nil, &result)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return result, nil
}

func (c *ApiGatewayClient) Prompt(ctx context.Context, message string, history []ChatCompletionMessage) (interface{}, error) {

	request := &openai.ChatCompletionRequest{
		// Messages: append(history, ChatCompletionMessage{
		// 	Role:    "user",
		// 	Content: message,
		// }),
		Messages: []ChatCompletionMessage{{
			Role:    "user",
			Content: message,
		}},
		Stream: false,
		Model:  "llama2",
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeText,
		},
	}

	return c.OpenAiClient.CreateChatCompletion(ctx, *request)
}

func (c *ApiGatewayClient) PromptStream(ctx context.Context, message string, sessionId string, flush CompletionCallback) (interface{}, error) {

	request := &openai.ChatCompletionRequest{
		Messages: []ChatCompletionMessage{{
			Role:    "user",
			Content: message,
		},
		},
		Model: "llama2",
	}

	return RequestChatCompletionStream(ctx, request, flush, sessionId)
}

func (c *ApiGatewayClient) GetLatestBlock(ctx context.Context) (result uint64, err error) {

	err = c.getRequest(ctx, "/blockchain/latestBlock", &result)
	return result, err
}

func (c *ApiGatewayClient) GetAllProviders(ctx context.Context) (result map[string]interface{}, err error) {

	err = c.getRequest(ctx, "/blockchain/providers", &result)

	if err != nil {
		return nil, fmt.Errorf("internal error: %v", err)
	}

	return result, nil
}

func (c *ApiGatewayClient) CreateNewProvider(ctx context.Context, addStake uint64, endpoint string) (result interface{}, err error) {
	request := struct {
		AddStake uint64 `json:"addStake"`
		Endpoint string `json:"endpoint"`
	}{addStake, endpoint}

	err = c.postRequest(ctx, "/blockchain/providers", &request, &result)

	if err != nil {
		return nil, fmt.Errorf("internal error: %v", err)
	}

	return result, nil
}

func (c *ApiGatewayClient) CreateNewProviderBid(ctx context.Context, model string, pricePerSecond uint64) (result interface{}, err error) {
	request := struct {
		Model          string `json:"model"`
		PricePerSecond uint64 `json:"pricePerSecond"`
	}{model, pricePerSecond}

	err = c.postRequest(ctx, "/blockchain/providers/bids", &request, &result)

	if err != nil {
		return nil, fmt.Errorf("internal error: %v", err)
	}

	return result, nil
}

func (c *ApiGatewayClient) GetAllModels(ctx context.Context) (result interface{}, err error) {

	err = c.getRequest(ctx, "/blockchain/models", &result)

	if err != nil {
		return nil, fmt.Errorf("internal error: %v", err)
	}

	return result, nil
}

func (c *ApiGatewayClient) GetBidsByProvider(ctx context.Context, providerAddr string, offset *big.Int, limit uint8) (bids interface{}, err error) {
	endpoint := fmt.Sprintf("/blockchain/providers/%s/bids?offset=%s&limit=%d", providerAddr, offset.String(), limit)
	err = c.getRequest(ctx, endpoint, &bids)
	if err != nil {
		return nil, fmt.Errorf("internal error: %v", err)
	}
	return bids, err
}

func (c *ApiGatewayClient) GetBidsByModelAgent(ctx context.Context, modelAgentId [32]byte, offset *big.Int, limit uint8) (result []string, err error) {

	endpoint := fmt.Sprintf("/blockchain/models/%x/bids?offset=%s&limit=%d", modelAgentId, offset.String(), limit)
	err = c.getRequest(ctx, endpoint, &result)

	if err != nil {
		return nil, fmt.Errorf("internal error: %v", err)
	}

	return result, err
}

type SessionRequest struct {
	ModelId          string `json:"modelId" validate:"required"`
}

type SessionStakeRequest struct {
	Approval    string `json:"approval" validate:"required"`
	ApprovalSig string `json:"approvalSig" validate:"required"`
	Stake       uint64 `json:"stake" validate:"required,number"`
}

type Session struct {
	SessionId string `json:"sessionId"`
}

type CloseSessionRequest struct {
	SessionId string `json:"id" validate:"required"`
}

type SessionListItem struct {
	Bid             string `json:"BidId"`
	Sesssion        string `json:"Id"`
	ModelORAgent    string `json:"ModelAgentId"`
	PricePerSecond  uint64 `json:"PricePerSecond"`
	Provider        string `json:"Provider"`
	User            string `json:"User"`
	ClosedAt        uint64 `json:"ClosedAt"`
	CloseoutReceipt string `json:"CloseoutReceipt"`
	CloseoutType    uint   `json:"CloseoutType"`
	EndsAt          uint64 `json:"EndsAt"`
}

type WalletRequest struct {
	PrivateKey string `json:"privateKey" validate:"required"`
}

func (c *ApiGatewayClient) ListUserSessions(ctx context.Context, user string) (result []SessionListItem, err error) {

	response := map[string][]SessionListItem{}

	err = c.getRequest(ctx, fmt.Sprintf("/blockchain/sessions?user=%s", user), &response)
	// fmt.Println("response: ", response)
	if err != nil {
		return nil, fmt.Errorf("internal error: %v", err)
	}

	return response["sessions"], nil
}

func (c *ApiGatewayClient) OpenStakeSession(ctx context.Context, req *SessionStakeRequest) (session *Session, err error) {

	err = c.postRequest(ctx, "/blockchain/sessions", req, session)

	if err != nil {
		return nil, fmt.Errorf("internal error: %v", err)
	}

	return session, nil
}

func (c *ApiGatewayClient) OpenSession(ctx context.Context, req *SessionRequest) (session *Session, err error) {

	session = &Session{}
	err = c.postRequest(ctx, "/blockchain/sessions/v2", req, session)

	if err != nil {
		return nil, fmt.Errorf("internal error: %v", err)
	}

	return session, nil
}

func (c *ApiGatewayClient) CloseSession(ctx context.Context, sessionId string) error {

	err := c.postRequest(ctx, fmt.Sprintf("/blockchain/sessions/%s/close", sessionId), nil, nil)

	if err != nil {
		return fmt.Errorf("internal error: %v", err)
	}

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

func (c *ApiGatewayClient) CreateWallet(ctx context.Context, privateKey string) error {
	return c.postRequest(ctx, "/wallet", &WalletRequest{PrivateKey: privateKey}, nil)
}

type WalletResponse struct {
	Address string `json:"address"`
}

func (c *ApiGatewayClient) GetWallet(ctx context.Context) (*WalletResponse, error) {
	response := &WalletResponse{}

	err := c.getRequest(ctx, "/wallet", response)

	if err != nil {
		return nil, fmt.Errorf("internal error: %v", err)
	}

	return response, nil
}

func (c *ApiGatewayClient) GetBalance(ctx context.Context) (eth uint64, mor uint64, err error) {
	response := map[string]interface{}{}

	err = c.getRequest(ctx, "/blockchain/balance", &response)

	if err != nil {
		return 0, 0, fmt.Errorf("internal error: %v", err)
	}

	fmt.Println(len(response["eth"].(string)))

	eth, err = strconv.ParseUint(response["eth"].(string), 10, 64)

	if err != nil {
		return 0, 0, fmt.Errorf("error converting eth balance from string to number: %v", err)
	}

	mor, err = strconv.ParseUint(response["mor"].(string), 10, 64)

	if err != nil {
		return 0, 0, fmt.Errorf("error converting mor balance from string to number: %v", err)
	}

	return eth, mor, nil
}
