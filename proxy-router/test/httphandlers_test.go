package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/aiengine"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/apibus"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/blockchainapi"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/handlers/httphandlers"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/proxyapi"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/wallet"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/storages"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/walletapi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/stretchr/testify/require"
)

var WALLET_PRIVATE_KEY = lib.MustStringToHexString("") // Set this to a valid private key to run the test.

var DIAMOND_CONTRACT_ADDR = lib.MustStringToAddress("0x70768f0ff919e194e11abfc3a2edf43213359dc1")
var MOR_CONTRACT_ADDR = lib.MustStringToAddress("0xc1664f994fd3991f98ae944bc16b9aed673ef5fd")
var EXPLORER_API_URL = "https://api-sepolia.arbiscan.io/api"
var ETH_LEGACY_TX = false
var ETH_NODE_ADDRESS = "wss://arb-sepolia.g.alchemy.com/v2/Ken3T8xkvWUxtpKvb3yDedzF-sNsQDlZ"

var PROVIDER_ADDR = lib.MustStringToAddress("0x65bBb982d9B0AfE9AED13E999B79c56dDF9e04fC")
var PROVIDER_URL = "thehulk1.stg.lumerin.io:3333"
var BID_ID = lib.MustStringToHash("0xa0d6ea9ce7183510e16cbfd207b9e381a91c20ee75d5db483b5758ddf22a27b1")
var SESSION_DURATION = new(big.Int).SetInt64(5 * 60) // 5 minutes in seconds

func TestChat(t *testing.T) {

	walletAddr, err := lib.PrivKeyBytesToAddr(WALLET_PRIVATE_KEY)
	if err != nil {
		t.Fatalf("failed to get wallet address: %s", err)
		return
	}

	log := lib.NewTestLogger()

	// Create a new instance of the HTTPHandler.
	handler := httphandlers.CreateHTTPServer(log, apiBus)

	server := httptest.NewServer(handler)
	defer server.Close()

	// Make a request to get the token supply.
	// Make a request to get today's budget.
	// Make a request to get the provider's bids.
	// Calculate the stake.
	// Make a request to initiate a session.
	initiateSessionResponse := InitiateSession(t, server, walletAddr, stake)
	approval := initiateSessionResponse.Response.Result.Approval
	approvalSig := initiateSessionResponse.Response.Result.ApprovalSig

	// Make a request to open a session.
	server, providerPubKey, sessionId, _, shouldReturn := PrepSessionForChat(t)
	if shouldReturn {
		return
	}

	// Make a request to send a prompt.
	promptRequestBody := map[string]interface{}{
		"model":  "llama2",
		"stream": true,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": "Why sky is blue?",
			},
		},
	}

	promptBody, err := json.Marshal(promptRequestBody)
	require.NoError(t, err)

	sendPromptURL := server.URL + fmt.Sprintf("/v1/chat/completions")
	request := httptest.NewRequest("POST", sendPromptURL, bytes.NewReader(promptBody))
	request.Header.Add("session_id", sessionId)

	sendPromptResp, err := server.Client().Do(request)
	require.NoError(t, err)

	// Make a request to close a session.
	closeSessionURL := server.URL + fmt.Sprintf("/blockchain/sessions/%s/close", sessionId)
	closeSessionResp, err := http.Post(closeSessionURL, "application/json", nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, closeSessionResp.StatusCode)

	require.Equal(t, http.StatusOK, sendPromptResp.StatusCode)
}

func PrepSessionForChat(t *testing.T) (server *httptest.Server, providerPubKey string, sessionId string, stake *big.Int, shouldReturn bool) {
	apiBus := InitializeApiBus(t)

	walletAddr, err := lib.PrivKeyStringToAddr(WALLET_PRIVATE_KEY)
	if err != nil {
		t.Fatalf("failed to get wallet address: %s", err)
		return nil, "", "", nil, true
	}

	handler := httphandlers.NewHTTPHandler(apiBus)

	server = httptest.NewServer(handler)
	defer server.Close()

	supply := GetTokenSupply(t, server)

	budget := GetTodaysBudget(t, server)

	bid := FindBid(t, server)

	pricePerSecond := new(big.Float).SetFloat64(bid["PricePerSecond"].(float64))
	pricePerSecondInt := new(big.Int)
	pricePerSecondInt, _ = pricePerSecondInt.SetString(pricePerSecond.Text('f', 0), 10)

	totalCost := SESSION_DURATION.Mul(pricePerSecondInt, SESSION_DURATION)
	stake = totalCost.Div(totalCost.Mul(supply, totalCost), budget)

	initiateSessionResponse := InitiateSession(t, server, walletAddr, stake)
	approval := initiateSessionResponse.Response.Result.Approval
	approvalSig := initiateSessionResponse.Response.Result.ApprovalSig
	providerPubKey = initiateSessionResponse.Response.Result.Message

	sessionId = OpenSession(t, server, approval, approvalSig, stake)
	return server, providerPubKey, sessionId, stake, shouldReturn
}

func GetTokenSupply(t *testing.T, server *httptest.Server) *big.Int {
	getTokenSupplyURL := server.URL + "/blockchain/token/supply"
	getTokenSupplyResp, err := http.Get(getTokenSupplyURL)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, getTokenSupplyResp.StatusCode)

	bodyBytes, err := io.ReadAll(getTokenSupplyResp.Body)
	require.NoError(t, err)

	var data map[string]string
	err = json.Unmarshal(bodyBytes, &data)
	require.NoError(t, err)

	supply, ok := new(big.Int).SetString(data["supply"], 10)
	require.True(t, ok)
	return supply
}

func GetTodaysBudget(t *testing.T, server *httptest.Server) *big.Int {
	getTodaysBudgetURL := server.URL + "/blockchain/sessions/budget"
	getTodaysBudgetResp, err := http.Get(getTodaysBudgetURL)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, getTodaysBudgetResp.StatusCode)

	bodyBytes, err := io.ReadAll(getTodaysBudgetResp.Body)
	require.NoError(t, err)

	var data map[string]string
	err = json.Unmarshal(bodyBytes, &data)
	require.NoError(t, err)

	budget, ok := new(big.Int).SetString(data["budget"], 10)
	require.True(t, ok)
	return budget
}

func FindBid(t *testing.T, server *httptest.Server) map[string]interface{} {
	getBidsUrl := server.URL + fmt.Sprintf("/blockchain/providers/%s/bids", PROVIDER_ADDR)
	getBidResp, err := http.Get(getBidsUrl)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, getBidResp.StatusCode)

	bodyBytes, err := io.ReadAll(getBidResp.Body)
	require.NoError(t, err)
	var data map[string][]interface{}

	err = json.Unmarshal(bodyBytes, &data)
	require.NoError(t, err)

	var bid map[string]interface{}
	for _, v := range data["bids"] {
		b := v.(map[string]interface{})
		if b["Id"] == BID_ID {
			bid = b
			break
		}
	}
	require.NotNil(t, bid)
	return bid
}

type InitiateSessionData struct {
	Approval    string `json:"approval"`
	ApprovalSig string `json:"approvalSig"`
	Message     string `json:"message"`
}

type InitiateSessionResult struct {
	Result InitiateSessionData `json:"result"`
}

type InitiateSessionResponse struct {
	Response InitiateSessionResult `json:"response"`
}

func InitiateSession(t *testing.T, server *httptest.Server, walletAddr common.Address, stake *big.Int) InitiateSessionResponse {
	initiateSessionURL := server.URL + "/proxy/sessions/initiate"

	var body map[string]interface{} = make(map[string]interface{})
	body["user"] = walletAddr
	body["provider"] = PROVIDER_ADDR
	body["spend"] = stake
	body["bidId"] = BID_ID
	body["providerUrl"] = PROVIDER_URL
	initiateSessionBody, err := json.Marshal(body)
	require.NoError(t, err)

	initiateSessionResp, err := http.Post(initiateSessionURL, "application/json", bytes.NewReader(initiateSessionBody))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, initiateSessionResp.StatusCode)

	bodyBytes, err := io.ReadAll(initiateSessionResp.Body)
	require.NoError(t, err)

	fmt.Println(string(bodyBytes))

	var response InitiateSessionResponse
	err = json.Unmarshal(bodyBytes, &response)
	require.NoError(t, err)

	return response
}

func OpenSession(t *testing.T, server *httptest.Server, approval string, approvalSig string, stake *big.Int) string {
	var openBody map[string]interface{} = make(map[string]interface{})
	openBody["approval"] = approval
	openBody["approvalSig"] = approvalSig
	openBody["stake"] = stake.String()
	openSessionBody, err := json.Marshal(openBody)
	require.NoError(t, err)

	fmt.Printf("open session body: %s\n", string(openSessionBody))

	openSessionURL := server.URL + "/blockchain/sessions"
	openSessionResp, err := http.Post(openSessionURL, "application/json", bytes.NewReader(openSessionBody))

	fmt.Println("open session error: ", err)
	fmt.Println("open session response: ", openSessionResp)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, openSessionResp.StatusCode)

	bodyBytes, err := io.ReadAll(openSessionResp.Body)
	fmt.Println(string(bodyBytes))
	require.NoError(t, err)

	var data map[string]interface{}
	err = json.Unmarshal(bodyBytes, &data)
	require.NoError(t, err)

	sessionId := data["sessionId"]
	return sessionId.(string)
}

func InitializeApiBus(t *testing.T) *apibus.APIBus {
	log, err := lib.NewLogger("debug", true, false, false, "")
	if err != nil {
		t.Fatalf("failed to create logger: %s", err)
		return nil
	}

	ethClient, err := ethclient.DialContext(context.Background(), ETH_NODE_ADDRESS)
	if err != nil {
		t.Fatalf("failed to connect to the Ethereum node: %s", err)
		return nil
	}

	contractLogStorage := lib.NewCollection[*interfaces.LogStorage]()

	storage := storages.NewStorage(log, "data/test/")
	sessionStorage := storages.NewSessionStorage(storage)

	wlt := wallet.NewEnvWallet(WALLET_PRIVATE_KEY)
	proxyRouterApi := proxyapi.NewProxySender(&url.URL{}, wlt, contractLogStorage, sessionStorage, log)
	blockchainApi := blockchainapi.NewBlockchainService(ethClient, DIAMOND_CONTRACT_ADDR.Address, MOR_CONTRACT_ADDR.Address, EXPLORER_API_URL, wlt, sessionStorage, proxyRouterApi, log, ETH_LEGACY_TX)
	aiEngine := aiengine.NewAiEngine("http://localhost:11434/v1", "", log)

	apiBus := apibus.NewApiBus(
		blockchainapi.NewBlockchainController(blockchainApi, log),
		proxyapi.NewProxyController(proxyRouterApi, aiEngine),
		walletapi.NewWalletController(wlt),
	)
	return apiBus
}
