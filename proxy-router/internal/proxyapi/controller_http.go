package proxyapi

import (
	"net/http"

	constants "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/aiengine"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

type ProxyController struct {
	service  *ProxyServiceSender
	aiEngine *aiengine.AiEngine
}

func NewProxyController(service *ProxyServiceSender, aiEngine *aiengine.AiEngine) *ProxyController {
	c := &ProxyController{
		service:  service,
		aiEngine: aiEngine,
	}

	return c
}

func (s *ProxyController) RegisterRoutes(r interfaces.Router) {
	r.POST("/proxy/sessions/initiate", s.InitiateSession)
	r.POST("/v1/chat/completions", s.Prompt)
}

// InitiateSession godoc
//
//		@Summary		Initiate Session with Provider
//		@Description	sends a handshake to the provider
//	 	@Tags			sessions
//		@Produce		json
//		@Success		200	{object}	interface{}
//		@Router			/proxy/sessions/initiate [post]
func (s *ProxyController) InitiateSession(ctx *gin.Context) {
	var req *InitiateSessionReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.service.InitiateSession(ctx, req.User, req.Provider, req.Spend.Unpack(), req.BidID, req.ProviderUrl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// SendPrompt godoc
//
//		@Summary		Send Local Or Remote Prompt
//		@Description	Send prompt to a local or remote model based on session id in header
//	 	@Tags			wallet
//		@Produce		json
//		@Param			prompt	body		proxyapi.OpenAiCompletitionRequest 	true	"Prompt"
//		@Param 			session_id header string false "Session ID"
//		@Success		200	{object}	interface{}
//		@Router			/v1/chat/completions [post]
func (c *ProxyController) Prompt(ctx *gin.Context) {
	var (
		body openai.ChatCompletionRequest
		head PromptHead
	)
	if err := ctx.ShouldBindHeader(&head); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if (head.SessionID == common.Hash{}) {
		body.Stream = ctx.GetHeader(constants.HEADER_ACCEPT) == constants.CONTENT_TYPE_JSON
		c.aiEngine.PromptCb(ctx, &body)
		return
	}

	err := c.service.SendPrompt(ctx, ctx.Writer, &body, head.SessionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	return
}
