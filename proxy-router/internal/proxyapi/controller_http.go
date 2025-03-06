package proxyapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	constants "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/aiengine"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/blockchainapi/structs"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/chatstorage/genericchatstorage"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/system"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

type AIEngine interface {
	GetLocalModels() ([]aiengine.LocalModel, error)
	GetAdapter(ctx context.Context, chatID, modelID, sessionID common.Hash, storeContext, forwardContext bool) (aiengine.AIEngineStream, error)
}

type ProxyController struct {
	service            *ProxyServiceSender
	aiEngine           AIEngine
	chatStorage        genericchatstorage.ChatStorageInterface
	storeChatContext   bool
	forwardChatContext bool
	log                lib.ILogger
	authConfig         system.HTTPAuthConfig
	ipfsManager        *IpfsManager
}

func NewProxyController(service *ProxyServiceSender, aiEngine AIEngine, chatStorage genericchatstorage.ChatStorageInterface, storeChatContext, forwardChatContext bool, authConfig system.HTTPAuthConfig, log lib.ILogger) *ProxyController {
	ipfsManager := NewIpfsManager()
	c := &ProxyController{
		service:            service,
		aiEngine:           aiEngine,
		chatStorage:        chatStorage,
		storeChatContext:   storeChatContext,
		forwardChatContext: forwardChatContext,
		log:                log,
		authConfig:         authConfig,
		ipfsManager:        ipfsManager,
	}

	return c
}

func (s *ProxyController) RegisterRoutes(r interfaces.Router) {
	r.POST("/proxy/provider/ping", s.Ping)
	r.POST("/proxy/sessions/initiate", s.authConfig.CheckAuth("initiate_session"), s.InitiateSession)
	r.POST("/v1/chat/completions", s.authConfig.CheckAuth("chat"), s.Prompt)
	r.GET("/v1/models", s.authConfig.CheckAuth("get_local_models"), s.Models)
	r.GET("/v1/chats", s.authConfig.CheckAuth("get_chat_history"), s.GetChats)
	r.GET("/v1/chats/:id", s.authConfig.CheckAuth("get_chat_history"), s.GetChat)
	r.DELETE("/v1/chats/:id", s.authConfig.CheckAuth("edit_chat_history"), s.DeleteChat)
	r.POST("/v1/chats/:id", s.authConfig.CheckAuth("edit_chat_history"), s.UpdateChatTitle)

	r.POST("/ipfs/pin", s.authConfig.CheckAuth("ipfs_pin"), s.Pin)
	r.POST("/ipfs/unpin", s.authConfig.CheckAuth("ipfs_unpin"), s.Unpin)
	r.POST("/ipfs/download/:cid", s.authConfig.CheckAuth("ipfs_get"), s.DownloadFile)
	r.POST("/ipfs/add", s.authConfig.CheckAuth("ipfs_add"), s.AddFile)
	r.GET("/ipfs/version", s.authConfig.CheckAuth("ipfs_version"), s.GetIpfsVersion)
	r.GET("/ipfs/pin", s.authConfig.CheckAuth("ipfs_pinned"), s.GetPinnedFiles)
}

// Ping godoc
//
//	@Summary		Ping Provider
//	@Description	sends a ping to the provider on the RPC level
//	@Tags			chat
//	@Produce		json
//	@Param			pingReq	body		proxyapi.PingReq	true	"Ping Request"
//	@Success		200		{object}	proxyapi.PingRes
//	@Router			/proxy/provider/ping [post]
func (s *ProxyController) Ping(ctx *gin.Context) {
	var req *PingReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ping, err := s.service.Ping(ctx, req.ProviderURL, req.ProviderAddr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &PingRes{PingMs: ping.Milliseconds()})
}

// InitiateSession godoc
//
//	@Summary		Initiate Session with Provider
//	@Description	sends a handshake to the provider
//	@Tags			chat
//	@Produce		json
//	@Param			initiateSession	body		proxyapi.InitiateSessionReq	true	"Initiate Session"
//	@Success		200				{object}	morrpcmesssage.SessionRes
//	@Security		BasicAuth
//
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
//	@Param			session_id	header		string											false	"Session ID"	format(hex32)
//	@Param			model_id	header		string											false	"Model ID"		format(hex32)
//	@Param			chat_id		header		string											false	"Chat ID"		format(hex32)
//	@Param			prompt		body		proxyapi.ChatCompletionRequestSwaggerExample	true	"Prompt"
//	@Success		200			{object}	string
//	@Security		BasicAuth
//
//	@Router			/v1/chat/completions [post]
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

	chatID := head.ChatID
	if chatID == (lib.Hash{}) {
		var err error
		chatID, err = lib.GetRandomHash()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	adapter, err := c.aiEngine.GetAdapter(ctx, chatID.Hash, head.ModelID.Hash, head.SessionID.Hash, c.storeChatContext, c.forwardChatContext)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var contentType string
	if body.Stream {
		contentType = constants.CONTENT_TYPE_EVENT_STREAM
	} else {
		contentType = constants.CONTENT_TYPE_JSON
	}

	ctx.Writer.Header().Set(constants.HEADER_CONTENT_TYPE, contentType)

	err = adapter.Prompt(ctx, &body, func(cbctx context.Context, completion genericchatstorage.Chunk) error {
		marshalledResponse, err := json.Marshal(completion.Data())
		if err != nil {
			return err
		}

		_, err = ctx.Writer.Write([]byte(fmt.Sprintf("data: %s\n\n", marshalledResponse)))
		if err != nil {
			return err
		}

		ctx.Writer.Flush()
		return nil
	})

	if err != nil {
		c.log.Errorf("error sending prompt: %s", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// GetLocalModels godoc
//
//	@Summary	Get local models
//	@Tags		system
//	@Produce	json
//	@Success	200	{object}	[]aiengine.LocalModel
//	@Security	BasicAuth
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
//	@Success	200	{object}	[]genericchatstorage.Chat
//	@Security	BasicAuth
//	@Router		/v1/chats [get]
func (c *ProxyController) GetChats(ctx *gin.Context) {
	chats := c.chatStorage.GetChats()

	if chats == nil {
		ctx.JSON(http.StatusOK, make([]struct{}, 0))
		return
	}

	ctx.JSON(http.StatusOK, chats)
}

// GetChat godoc
//
//	@Summary	Get chat by id
//	@Tags		chat
//	@Produce	json
//	@Param		id	path		string	true	"Chat ID"
//	@Success	200	{object}	genericchatstorage.ChatHistory
//	@Security	BasicAuth
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
//	@Security	BasicAuth
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
//	@Security	BasicAuth
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

// Pin godoc
//
//	@Summary	Pin a file to IPFS
//	@Tags		ipfs
//	@Produce	json
//	@Param		cid	body		proxyapi.CIDReq	true	"CID"
//	@Success	200	{object}	proxyapi.ResultResponse
//	@Security	BasicAuth
//	@Router		/ipfs/pin [post]
func (c *ProxyController) Pin(ctx *gin.Context) {
	var req struct {
		CID lib.Hash `json:"cid"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	CID, err := lib.ManualBytes32ToCID(req.CID.Bytes())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.ipfsManager.Pin(ctx, CID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"result": true})
}

// Unpin godoc
//
//	@Summary	Unpin a file from IPFS
//	@Tags		ipfs
//	@Produce	json
//	@Param		cid	body		proxyapi.CIDReq	true	"CID"
//	@Success	200	{object}	proxyapi.ResultResponse
//	@Security	BasicAuth
//	@Router		/ipfs/unpin [post]
func (c *ProxyController) Unpin(ctx *gin.Context) {
	var req struct {
		CID lib.Hash `json:"cid"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	CID, err := lib.ManualBytes32ToCID(req.CID.Bytes())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.ipfsManager.Unpin(ctx, CID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"result": true})
}

// DownloadFile godoc
//
//	@Summary	Download a file from IPFS
//	@Tags		ipfs
//	@Produce	json
//	@Param		cid	uri			proxyapi.CIDReq	true	"CID"
//	@Success	200	{object}	proxyapi.ResultResponse
//	@Security	BasicAuth
//	@Router		/ipfs/download/{cid} [post]
func (c *ProxyController) DownloadFile(ctx *gin.Context) {
	var params struct {
		CID lib.Hash `uri:"cid" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req struct {
		DestinationPath string `json:"destinationPath"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	CID, err := lib.ManualBytes32ToCID(params.CID.Bytes())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.ipfsManager.GetFile(ctx, CID, req.DestinationPath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"result": true})
}

// AddFile godoc
//
//	@Summary	Add a file to IPFS
//	@Tags		ipfs
//	@Produce	json
//	@Param		filePath	body		proxyapi.AddFileReq	true	"File Path"
//	@Success	200			{object}	proxyapi.AddIpfsFileRes
//	@Security	BasicAuth
//	@Router		/ipfs/add [post]
func (c *ProxyController) AddFile(ctx *gin.Context) {
	var req struct {
		FilePath string `json:"filePath"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cid, err := c.ipfsManager.AddFile(ctx, req.FilePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cidHash, err := lib.CIDToBytes32(cid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	hash := lib.HexString(cidHash)
	ctx.JSON(http.StatusOK, AddIpfsFileRes{CID: cid, Hash: hash})
}

// GetIpfsVersion godoc
//
//	@Summary	Get IPFS Version
//	@Tags		ipfs
//	@Produce	json
//	@Success	200	{object}	proxyapi.IpfsVersionRes
//	@Security	BasicAuth
//	@Router		/ipfs/version [get]
func (c *ProxyController) GetIpfsVersion(ctx *gin.Context) {
	version, err := c.ipfsManager.GetVersion(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, IpfsVersionRes{Version: version})
}

// GetPinnedFiles godoc
//
//	@Summary	Get all pinned files
//	@Tags		ipfs
//	@Produce	json
//	@Success	200	{object}	proxyapi.IpfsPinnedFilesRes
//	@Security	BasicAuth
//	@Router		/ipfs/pin [get]
func (c *ProxyController) GetPinnedFiles(ctx *gin.Context) {
	files, err := c.ipfsManager.GetPinnedFiles(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	filesWithHashes := make([]IpfsPinnedFile, len(files))
	for i, file := range files {
		cidBytes, err := lib.CIDToBytes32(file)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		hash := lib.HexString(cidBytes)
		filesWithHashes[i] = IpfsPinnedFile{CID: file, Hash: hash}
	}
	ctx.JSON(http.StatusOK, IpfsPinnedFilesRes{Files: filesWithHashes})
}
