package apibus

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/internal/aiengine"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/internal/proxyapi"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/internal/rpcproxy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

// TODO: split implementations into separate client layer
type ApiBus struct {
	rpcProxy       *rpcproxy.RpcProxy
	aiEngine       *aiengine.AiEngine
	proxyRouterApi *proxyapi.ProxyRouterApi
}

func NewApiBus(rpcProxy *rpcproxy.RpcProxy, aiEngine *aiengine.AiEngine, proxyRouterApi *proxyapi.ProxyRouterApi) *ApiBus {
	return &ApiBus{
		rpcProxy:       rpcProxy,
		aiEngine:       aiEngine,
		proxyRouterApi: proxyRouterApi,
	}
}

// Proxy Router Api

// GetConfig godoc
//
//		@Summary		Get Config
//		@Description	Return the current config of proxy router
//	 	@Tags			healthcheck
//		@Produce		json
//		@Success		200	{object}	proxyapi.ConfigResponse
//		@Router			/config [get]
func (apiBus *ApiBus) GetConfig(ctx context.Context) interface{} {
	return apiBus.proxyRouterApi.GetConfig(ctx)
}

func (apiBus *ApiBus) GetFiles(ctx *gin.Context) (int, interface{}) {
	return apiBus.proxyRouterApi.GetFiles(ctx)
}

// HealthCheck godoc
//
//		@Summary		Healthcheck example
//		@Description	do ping
//	 	@Tags			healthcheck
//		@Produce		json
//		@Success		200	{object}	proxyapi.HealthCheckResponse
//		@Router			/healthcheck [get]
func (apiBus *ApiBus) HealthCheck(ctx context.Context) interface{} {
	return apiBus.proxyRouterApi.HealthCheck(ctx)
}

// InitiateSession godoc
//
//		@Summary		Initiate Session with Provider
//		@Description	sends a handshake to the provider
//	 	@Tags			sessions
//		@Produce		json
//		@Success		200	{object}	interface{}
//		@Router			/proxy/sessions/initiate [post]
func (apiBus *ApiBus) InitiateSession(ctx *gin.Context) (int, interface{}) {
	return apiBus.proxyRouterApi.InitiateSession(ctx)
}

// SendPrompt godoc
//
//		@Summary		Send prompt to provider
//		@Description	sens a prompt to the provider by opened session
//	 	@Tags			sessions
//		@Produce		json
//		@Param 			id  path string true "Session ID"
//		@Param			prompt	body		proxyapi.RemotePromptRequest 	true	"RemotePrompt"
//		@Success		200	{object}	interface{}
//		@Router			/proxy/sessions/{id}/prompt [post]
func (apiBus *ApiBus) SendPrompt(ctx *gin.Context) (bool, int, interface{}) {
	return apiBus.proxyRouterApi.SendPrompt(ctx)
}

// AiEngine
func (apiBus *ApiBus) Prompt(ctx context.Context, req interface{}) (interface{}, error) {
	return apiBus.aiEngine.Prompt(ctx, req)
}

// AiEngine
func (apiBus *ApiBus) PromptStream(ctx context.Context, req interface{}, flush interface{}) (interface{}, error) {
	return apiBus.aiEngine.PromptStream(ctx, req, flush)
}

// RpcProxy
func (apiBus *ApiBus) GetLatestBlock(ctx context.Context) (uint64, error) {
	return apiBus.rpcProxy.GetLatestBlock(ctx)
}

// GetBalance godoc
//
//		@Summary		Get ETH and MOR balance
//		@Description	Get ETH and MOR balance of the user
//	 	@Tags			wallet
//		@Produce		json
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/balance [get]
func (apiBus *ApiBus) GetBalance(ctx context.Context) (int, gin.H) {
	apiContext, ok :=  ctx.(*gin.Context)

	if !ok {
		return 500, gin.H{"error": "invalid context"}
	}
	
	return apiBus.rpcProxy.GetBalance(apiContext)
}

// GetAllowance godoc
//
//		@Summary		Get Allowance for MOR
//		@Description	Get MOR allowance for spender
//	 	@Tags			wallet
//		@Produce		json
//		@Param 			spender	query	string	true	"Spender address"
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/allowance [get]
func (apiBus *ApiBus) GetAllowance(ctx *gin.Context) (int, gin.H) {
	return apiBus.rpcProxy.GetAllowance(ctx)
}

// Approve godoc
//
//		@Summary		Approve MOR allowance
//		@Description	Approve MOR allowance for spender
//	 	@Tags			wallet
//		@Produce		json
//		@Param 			spender	query	string	true	"Spender address"
//		@Param 			amount	query	string	true	"Amount"
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/allowance [post]
func (apiBus *ApiBus) Approve(ctx *gin.Context) (int, gin.H) {
	return apiBus.rpcProxy.Approve(ctx)
}

// GetProviders godoc
//
//		@Summary		Get providers list
//		@Description	Get providers list from blokchain
//	 	@Tags			wallet
//		@Produce		json
//		@Success		200	{object}	[]interface{}
//		@Router			/blockchain/providers [get]
func (apiBus *ApiBus) GetAllProviders(ctx context.Context) (int, gin.H) {
	return apiBus.rpcProxy.GetAllProviders(ctx)
}

func (apiBus *ApiBus) CreateNewProvider(ctx context.Context, address string, addStake uint64, endpoint string) (int, gin.H) {
	return apiBus.rpcProxy.CreateNewProvider(ctx, address, addStake, endpoint)
}

// GetModels godoc
//
//		@Summary		Get models list
//		@Description	Get models list from blokchain
//	 	@Tags			wallet
//		@Produce		json
//		@Success		200	{object}	[]interface{}
//		@Router			/blockchain/models [get]
func (apiBus *ApiBus) GetAllModels(ctx context.Context) (int, gin.H) {
	return apiBus.rpcProxy.GetAllModels(ctx)
}

// GetTransactions godoc
//
//		@Summary		Get Transactions
//		@Description	Get MOR and ETH transactions
//	 	@Tags			wallet
//		@Produce		json
//		@Param 			page	query	string	false	"Page"
//		@Param 			limit	query	string	false	"Limit"
//		@Success		200	{object}	[]interface{}
//		@Router			/blockchain/transactions [get]
func (apiBus *ApiBus) GetTransactions(ctx *gin.Context) (int, gin.H) {
	return apiBus.rpcProxy.GetTransactions(ctx)
}

// GetBidsByProvider godoc
//
//		@Summary		Get Bids by Provider
//		@Description	Get bids from blockchain by provider
//	 	@Tags			wallet
//		@Produce		json
//		@Param 			offset	query	string	false	"Offset"
//		@Param 			limit	query	string	false	"Limit"
//		@Param 			id  path string true "Provider ID"
//		@Success		200	{object}	[]interface{}
//		@Router			/blockchain/providers/{id}/bids [get]
func (apiBus *ApiBus) GetBidsByProvider(ctx context.Context, providerAddr string, offset *big.Int, limit uint8) (int, gin.H) {
	addr := common.HexToAddress(providerAddr)
	return apiBus.rpcProxy.GetBidsByProvider(ctx, addr, offset, limit)
}

// GetBidsByModelAgent godoc
//
//		@Summary		Get Bids by	Model Agent
//		@Description	Get bids from blockchain by model agent
//	 	@Tags			wallet
//		@Produce		json
//		@Param 			offset	query	string	false	"Offset"
//		@Param 			limit	query	string	false	"Limit"
//		@Param 			id  path string true "ModelAgent ID"
//		@Success		200	{object}	[]interface{}
//		@Router			/blockchain/models/{id}/bids [get]
func (apiBus *ApiBus) GetBidsByModelAgent(ctx context.Context, modelAgentId [32]byte, offset *big.Int, limit uint8) (int, gin.H) {
	return apiBus.rpcProxy.GetBidsByModelAgent(ctx, modelAgentId, offset, limit)
}

// SendEth godoc
//
//		@Summary		Send Eth
//		@Description	Send Eth to address
//	 	@Tags			wallet
//		@Produce		json
//		@Param			sendeth	body		rpcproxy.SendRequest 	true	"Send Eth"
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/send/eth [post]
func (apiBus *ApiBus) SendEth(ctx *gin.Context) (int, gin.H) {
	return apiBus.rpcProxy.SendEth(ctx)
}

// SendMor godoc
//
//		@Summary		Send Mor
//		@Description	Send Mor to address
//	 	@Tags			wallet
//		@Produce		json
//		@Param			sendmor	body		rpcproxy.SendRequest 	true	"Send Mor"
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/send/mor [post]
func (apiBus *ApiBus) SendMor(ctx *gin.Context) (int, gin.H) {
	return apiBus.rpcProxy.SendMor(ctx)
}

// OpenSession godoc
//
//		@Summary		Open Session with Provider in blockchain
//		@Description	Sends transaction in blockchain to open a session
//	 	@Tags			sessions
//		@Produce		json
//		@Accept			json
//		@Param			opensession	body		rpcproxy.OpenSessionRequest 	true	"Open session"
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/sessions [post]
func (apiBus *ApiBus) OpenSession(ctx *gin.Context) (int, gin.H) {
	return apiBus.rpcProxy.OpenSession(ctx)
}

// CloseSession godoc
//
//		@Summary		Close Session with Provider
//		@Description	Sends transaction in blockchain to close a session
//	 	@Tags			sessions
//		@Produce		json
//		@Param 			id  path string true "Session ID"
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/sessions/{id}/close [post]
func (apiBus *ApiBus) CloseSession(ctx *gin.Context) (int, gin.H) {
	return apiBus.rpcProxy.CloseSession(ctx)
}

// ClaimProviderBalance godoc
//
//		@Summary		Claim Provider Balance
//		@Description	Claim provider balance from session
//	 	@Tags			sessions
//		@Produce		json
//		@Param			claim	body		rpcproxy.SendRequest 	true	"Claim"
//		@Param 			id  path string true "Session ID"
//		@Success		200	{object}	interface{}
//		@Router			/proxy/sessions/${id}/providerClaim [post]
func (apiBus *ApiBus) ClaimProviderBalance(ctx *gin.Context) (int, gin.H) {
	return apiBus.rpcProxy.ClaimProviderBalance(ctx)
}

// GetProviderClaimableBalance godoc
//
//		@Summary		Get Provider Claimable Balance
//		@Description	Get provider claimable balance from session
//	 	@Tags			sessions
//		@Produce		json
//		@Param 			id  path string true "Session ID"
//		@Success		200	{object}	interface{}
//		@Router			/proxy/sessions/${id}/providerClaimableBalance [get]
func (apiBus *ApiBus) GetProviderClaimableBalance(ctx *gin.Context) (int, gin.H) {
	return apiBus.rpcProxy.GetProviderClaimableBalance(ctx)
}

// GetTodaysBudget godoc
//
//		@Summary		Get Todays Budget
//		@Description	Get todays budget from blockchain
//	 	@Tags			wallet
//		@Produce		json
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/sessions/budget [get]
func (apiBus *ApiBus) GetTodaysBudget(ctx *gin.Context) (int, gin.H) {
	return apiBus.rpcProxy.GetTodaysBudget(ctx)
}

// GetTokenSupply godoc
//
//		@Summary		Get Token Supply
//		@Description	Get MOR token supply from blockchain
//	 	@Tags			wallet
//		@Produce		json
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/token/supply [get]
func (apiBus *ApiBus) GetTokenSupply(ctx *gin.Context) (int, gin.H) {
	return apiBus.rpcProxy.GetTokenSupply(ctx)
}

// GetSessions godoc
//
//		@Summary		Get Sessions
//		@Description	Get sessions from blockchain by user or provider
//	 	@Tags			sessions
//		@Produce		json
//		@Param 			offset	query	string	false	"Offset"
//		@Param 			limit	query	string	false	"Limit"
//		@Param 			provider	query	string	false	"Provider address"
//		@Param 			user	query	string	false	"User address"
//		@Success		200	{object}	[]interface{}
//		@Router			/blockchain/sessions [get]
func (apiBus *ApiBus) GetSessions(ctx *gin.Context, offset *big.Int, limit uint8) (int, gin.H) {
	return apiBus.rpcProxy.GetSessions(ctx, offset, limit)
}

// SendLocalPrompt godoc
//
//		@Summary		Send prompt to a local model
//		@Description	Send prompt to a local model
//	 	@Tags			wallet
//		@Produce		json
//		@Param			prompt	body		proxyapi.PromptRequest 	true	"LocalPrompt"
//		@Success		200	{object}	interface{}
//		@Router			/v1/chat/completions [post]
func (apiBus *ApiBus) PromptLocal(ctx *gin.Context) {
	var req *openai.ChatCompletionRequest

	err := ctx.ShouldBindJSON(&req)
	switch {
	case errors.Is(err, io.EOF):
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing request body"})
		return
	case err != nil:
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response interface{}

	if req.Stream {
		response, err = apiBus.PromptStream(ctx, req, func(response *openai.ChatCompletionStreamResponse) error {

			marshalledResponse, err := json.Marshal(response)

			if err != nil {
				return err
			}

			ctx.Writer.Header().Set("Content-Type", "text/event-stream")
			_, err = ctx.Writer.Write([]byte(fmt.Sprintf("data: %s\n\n", marshalledResponse)))
			ctx.Writer.Flush()

			if err != nil {
				return err
			}

			return nil
		})
	} else {
		response, err = apiBus.Prompt(ctx, req)
	}

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
