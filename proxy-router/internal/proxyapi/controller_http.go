package proxyapi

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	constants "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/aiengine"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/blockchainapi/structs"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

type ProxyController struct {
	service     *ProxyServiceSender
	aiEngine    *aiengine.AiEngine
	chatStorage ChatStorageInterface
}

func NewProxyController(service *ProxyServiceSender, aiEngine *aiengine.AiEngine, chatStorage ChatStorageInterface) *ProxyController {
	c := &ProxyController{
		service:     service,
		aiEngine:    aiEngine,
		chatStorage: chatStorage,
	}

	return c
}

func (s *ProxyController) RegisterRoutes(r interfaces.Router) {
	r.POST("/proxy/sessions/initiate", s.InitiateSession)
	r.POST("/v1/chat/completions", s.Prompt)
	r.GET("/v1/models", s.Models)
	r.GET("/v1/chats", s.GetChats)
	r.GET("/v1/chats/:id", s.GetChat)
	r.DELETE("/v1/chats/:id", s.DeleteChat)
	r.POST("/v1/chats/:id", s.UpdateChatTitle)
}

// InitiateSession godoc
//
//	@Summary		Initiate Session with Provider
//	@Description	sends a handshake to the provider
//	@Tags			chat
//	@Produce		json
//	@Param			initiateSession	body		proxyapi.InitiateSessionReq	true	"Initiate Session"
//	@Success		200				{object}	morrpcmesssage.SessionRes
//	@Router			/proxy/sessions/initiate [post]
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
//	@Summary		Send Local Or Remote Prompt
//	@Description	Send prompt to a local or remote model based on session id in header
//	@Tags			chat
//	@Produce		text/event-stream
//	@Param			session_id	header		string								false	"Session ID"	format(hex32)
//	@Param			model_id	header		string								false	"Model ID"		format(hex32)
//	@Param			chat_id		header		string								false	"Chat ID"		format(hex32)
//	@Param			prompt		body		proxyapi.OpenAiCompletitionRequest	true	"Prompt"
//	@Success		200			{object}	proxyapi.ChatCompletionResponse
//	@Router			/v1/chat/completions [post]
func (c *ProxyController) Prompt(ctx *gin.Context) {
	var (
		body openai.ChatCompletionRequest
		head PromptHead
	)
	var responses []interface{}

	if err := ctx.ShouldBindHeader(&head); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var chatID lib.Hash
	if (head.ChatID != lib.Hash{}) {
		chatID = head.ChatID
	} else {
		bytes := make([]byte, 32)
		_, err := rand.Read(bytes[:])
		if err != nil {
			err = fmt.Errorf("error generating chat id: %w", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		chatID = lib.Hash{}
		chatID.SetBytes(bytes)
	}

	promptAt := time.Now()
	chatHistory, errHistory := c.chatStorage.LoadChatFromFile(chatID.Hex())
	if errHistory != nil {
		c.service.log.Warn("No chat history found for chat id", chatID.Hex())
	}

	if (head.SessionID == lib.Hash{}) {
		body.Stream = ctx.GetHeader(constants.HEADER_ACCEPT) == constants.CONTENT_TYPE_JSON
		modelId := head.ModelID.Hex()

		apiType := c.GetLocalPromptApiType(modelId)

		responseAt := time.Now()

		if apiType == "openai" {
			prompt := c.GetLocalOpenAiPrompt(modelId, body)
			var newBody openai.ChatCompletionRequest
			if chatHistory != nil {
				newBody = c.AppendChatHistory(chatHistory, prompt)
			} else {
				newBody = prompt
			}

			res, _ := c.aiEngine.PromptCb(ctx, &newBody)
			responses = res.([]interface{})
			if err := c.chatStorage.StorePromptResponseToFile(chatID.Hex(), "", modelId, prompt, responses, promptAt, responseAt); err != nil {
				fmt.Println("Error storing prompt and responses:", err)
			}
		}
		if apiType == "prodia" {
			prompt := c.GetLocalProdiaPrompt(modelId, body)

			var prodiaResponses []interface{}
			c.aiEngine.PromptProdiaImage(ctx, &prompt, func(completion interface{}) error {
				ctx.Writer.Header().Set(constants.HEADER_CONTENT_TYPE, constants.CONTENT_TYPE_EVENT_STREAM)
				marshalledResponse, err := json.Marshal(completion)
				if err != nil {
					return err
				}
				_, err = ctx.Writer.Write([]byte(fmt.Sprintf("data: %s\n\n", marshalledResponse)))
				if err != nil {
					return err
				}
				ctx.Writer.Flush()

				prodiaResponses = append(prodiaResponses, completion)

				body.Model = prompt.Model
				if err := c.chatStorage.StorePromptResponseToFile(chatID.Hex(), "", modelId, body, prodiaResponses, promptAt, responseAt); err != nil {
					fmt.Println("Error storing prompt and responses:", err)
				}
				return nil
			})
		}
		return
	}

	var newBody openai.ChatCompletionRequest
	if chatHistory != nil {
		newBody = c.AppendChatHistory(chatHistory, body)
	} else {
		newBody = body
	}
	res, err := c.service.SendPrompt(ctx, ctx.Writer, &newBody, head.SessionID.Hash)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses = res.([]interface{})
	responseAt := time.Now()
	sessionID := head.SessionID.Hex()

	modelId := ""
	session, ok := c.service.sessionStorage.GetSession(sessionID)
	if ok {
		modelId = session.ModelID
	}

	if err := c.chatStorage.StorePromptResponseToFile(chatID.Hex(), sessionID, modelId, body, responses, promptAt, responseAt); err != nil {
		fmt.Println("Error storing prompt and responses:", err)
	}
	return
}

// GetLocalModels godoc
//
//	@Summary	Get local models
//	@Tags		chat
//	@Produce	json
//	@Success	200	{object}	[]aiengine.LocalModel
//	@Router		/v1/models [get]
func (c *ProxyController) Models(ctx *gin.Context) {
	models, err := c.aiEngine.GetLocalModels()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, models)
}

// GetChats godoc
//
//	@Summary	Get all chats stored in the system
//	@Tags		chat
//	@Produce	json
//	@Success	200	{object}	[]proxyapi.Chat
//	@Router		/v1/chats [get]
func (c *ProxyController) GetChats(ctx *gin.Context) {
	chats := c.chatStorage.GetChats()
	ctx.JSON(http.StatusOK, chats)
}

// GetChat godoc
//
//	@Summary	Get chat by id
//	@Tags		chat
//	@Produce	json
//	@Param		id	path		string	true	"Chat ID"
//	@Success	200	{object}	proxyapi.ChatHistory
//	@Router		/v1/chats/{id} [get]
func (c *ProxyController) GetChat(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	chat, err := c.chatStorage.LoadChatFromFile(params.ID.Hex())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, chat)
}

// DeleteChat godoc
//
//	@Summary	Delete chat by id from storage
//	@Tags		chat
//	@Produce	json
//	@Param		id	path		string	true	"Chat ID"
//	@Success	200	{object}	proxyapi.ResultResponse
//	@Router		/v1/chats/{id} [delete]
func (c *ProxyController) DeleteChat(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	err = c.chatStorage.DeleteChat(params.ID.Hex())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"result": true})
}

// UpdateChatTitle godoc
//
//	@Summary	Update chat title by id
//	@Tags		chat
//	@Produce	json
//	@Param		id		path		string						true	"Chat ID"
//	@Param		title	body		proxyapi.UpdateChatTitleReq	true	"Chat Title"
//	@Success	200		{object}	proxyapi.ResultResponse
//	@Router		/v1/chats/{id} [post]
func (c *ProxyController) UpdateChatTitle(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	var req UpdateChatTitleReq
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		c.service.log.Errorf("error binding json: %s", err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	err = c.chatStorage.UpdateChatTitle(params.ID.Hex(), req.Title)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"result": true})
}

func (c *ProxyController) GetLocalOpenAiPrompt(modelId string, req openai.ChatCompletionRequest) openai.ChatCompletionRequest {
	if modelId == "" {
		return req
	}

	ids, models := c.aiEngine.GetModelsConfig()

	for i, model := range models {
		if ids[i] == modelId {
			req.Model = model.ModelName
			return req
		}
	}

	req.Model = "llama2"
	return req
}

func (c *ProxyController) GetLocalProdiaPrompt(modelId string, req openai.ChatCompletionRequest) aiengine.ProdiaGenerationRequest {
	ids, models := c.aiEngine.GetModelsConfig()

	for i, model := range models {
		if ids[i] == modelId {
			prompt := aiengine.ProdiaGenerationRequest{
				Model:  model.ModelName,
				Prompt: req.Messages[0].Content,
				ApiUrl: model.ApiURL,
				ApiKey: model.ApiKey,
			}
			return prompt
		}
	}
	return aiengine.ProdiaGenerationRequest{}
}

func (c *ProxyController) GetLocalPromptApiType(modelId string) string {
	if modelId == "" {
		return "openai"
	}

	ids, models := c.aiEngine.GetModelsConfig()

	for i, model := range models {
		if ids[i] == modelId {
			return model.ApiType
		}
	}

	return "openai"
}

func (c *ProxyController) AppendChatHistory(chatHistory *ChatHistory, req openai.ChatCompletionRequest) openai.ChatCompletionRequest {
	messagesWithHistory := make([]openai.ChatCompletionMessage, 0)
	for _, chat := range chatHistory.Messages {
		message := openai.ChatCompletionMessage{
			Role:    chat.Prompt.Messages[0].Role,
			Content: chat.Prompt.Messages[0].Content,
		}
		messagesWithHistory = append(messagesWithHistory, message)
		messagesWithHistory = append(messagesWithHistory, openai.ChatCompletionMessage{
			Role:    "assistant",
			Content: chat.Response,
		})
	}

	messagesWithHistory = append(messagesWithHistory, req.Messages...)
	req.Messages = messagesWithHistory
	return req
}
