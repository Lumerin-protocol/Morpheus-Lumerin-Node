package blockchainapi

import (
	"crypto/rand"
	"math/big"
	"net/http"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/blockchainapi/structs"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type BlockchainController struct {
	service *BlockchainService
	log     lib.ILogger
}

func NewBlockchainController(service *BlockchainService, log lib.ILogger) *BlockchainController {
	c := &BlockchainController{
		service: service,
		log:     log,
	}

	return c
}

func (c *BlockchainController) RegisterRoutes(r interfaces.Router) {
	// transactions
	r.GET("/blockchain/balance", c.getBalance)
	r.GET("/blockchain/transactions", c.getTransactions)
	r.GET("/blockchain/allowance", c.getAllowance)
	r.GET("/blockchain/latestBlock", c.getLatestBlock)
	r.POST("/blockchain/approve", c.approve)
	r.POST("/blockchain/send/eth", c.sendETH)
	r.POST("/blockchain/send/mor", c.sendMOR)

	// providers
	r.GET("/blockchain/providers", c.getAllProviders)
	r.POST("/blockchain/providers", c.createProvider)
	r.DELETE("/blockchain/providers/:id", c.deregisterProvider)

	// models
	r.GET("/blockchain/models", c.getAllModels)
	r.POST("/blockchain/models", c.createNewModel)
	r.DELETE("/blockchain/models/:id", c.deregisterModel)

	// bids
	r.POST("/blockchain/bids", c.createNewBid)
	r.GET("/blockchain/bids/:id", c.getBidByID)
	r.DELETE("/blockchain/bids/:id", c.deleteBid)
	r.GET("/blockchain/models/:id/bids", c.getBidsByModelAgent)
	r.GET("/blockchain/models/:id/bids/rated", c.getRatedBids)
	r.GET("/blockchain/models/:id/bids/active", c.getActiveBidsByModel)
	r.GET("/blockchain/providers/:id/bids", c.getBidsByProvider)
	r.GET("/blockchain/providers/:id/bids/active", c.getActiveBidsByProvider)

	// sessions
	r.GET("/proxy/sessions/:id/providerClaimableBalance", c.getProviderClaimableBalance)
	r.POST("/proxy/sessions/:id/providerClaim", c.claimProviderBalance)
	r.GET("/blockchain/sessions/user", c.getSessionsForUser)
	r.GET("/blockchain/sessions/provider", c.getSessionsForProvider)
	r.GET("/blockchain/sessions/:id", c.getSession)
	r.POST("/blockchain/sessions", c.openSession)
	r.POST("/blockchain/bids/:id/session", c.openSessionByBid)
	r.POST("/blockchain/models/:id/session", c.openSessionByModelId)
	r.POST("/blockchain/sessions/:id/close", c.closeSession)
	r.GET("/blockchain/sessions/budget", c.getBudget)
	r.GET("/blockchain/token/supply", c.getSupply)
}

// GetProviderClaimableBalance godoc
//
//	@Summary		Get Provider Claimable Balance
//	@Description	Get provider claimable balance from session
//	@Tags			sessions
//	@Produce		json
//	@Param			id	path		string	true	"Session ID"
//	@Success		200	{object}	structs.BalanceRes
//	@Router			/proxy/sessions/{id}/providerClaimableBalance [get]
func (c *BlockchainController) getProviderClaimableBalance(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	balance, err := c.service.GetProviderClaimableBalance(ctx, params.ID.Hash)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.BalanceRes{Balance: &lib.BigInt{Int: *balance}})
	return
}

// ClaimProviderBalance godoc
//
//	@Summary		Claim Provider Balance
//	@Description	Claim provider balance from session
//	@Tags			sessions
//	@Produce		json
//	@Param			id		path		string				true	"Session ID"
//	@Success		200		{object}	structs.TxRes
//	@Router			/proxy/sessions/{id}/providerClaim [post]
func (c *BlockchainController) claimProviderBalance(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	txHash, err := c.service.ClaimProviderBalance(ctx, params.ID.Hash)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.TxRes{Tx: txHash})
	return
}

// GetProviders godoc
//
//	@Summary		Get providers list
//	@Description	Get providers list from blokchain
//	@Tags			providers
//	@Produce		json
//	@Success		200	{object}	structs.ProvidersRes
//	@Router			/blockchain/providers [get]
func (c *BlockchainController) getAllProviders(ctx *gin.Context) {
	providers, err := c.service.GetAllProviders(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.ProvidersRes{Providers: providers})
	return
}

// SendEth godoc
//
//	@Summary		Send Eth
//	@Description	Send Eth to address
//	@Tags			transactions
//	@Produce		json
//	@Param			sendeth	body		structs.SendRequest	true	"Send Eth"
//	@Success		200		{object}	structs.TxRes
//	@Router			/blockchain/send/eth [post]
func (c *BlockchainController) sendETH(ctx *gin.Context) {
	to, amount, err := c.getSendParams(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	txHash, err := c.service.SendETH(ctx, to, amount)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.TxRes{Tx: txHash})
	return
}

// SendMor godoc
//
//	@Summary		Send Mor
//	@Description	Send Mor to address
//	@Tags			transactions
//	@Produce		json
//	@Param			sendmor	body		structs.SendRequest	true	"Send Mor"
//	@Success		200		{object}	structs.TxRes
//	@Router			/blockchain/send/mor [post]
func (c *BlockchainController) sendMOR(ctx *gin.Context) {
	to, amount, err := c.getSendParams(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}
	txhash, err := c.service.SendMOR(ctx, to, amount)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.TxRes{Tx: txhash})
	return
}

// GetBidsByProvider godoc
//
//	@Summary		Get Bids by Provider
//	@Description	Get bids from blockchain by provider
//	@Tags			bids
//	@Produce		json
//	@Param			offset	query		string	false	"Offset"
//	@Param			limit	query		string	false	"Limit"
//	@Param			id		path		string	true	"Provider ID"
//	@Success		200		{object}	structs.BidsRes
//	@Router			/blockchain/providers/{id}/bids [get]
func (c *BlockchainController) getBidsByProvider(ctx *gin.Context) {
	var params structs.PathEthAddrID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	offset, limit, err := getOffsetLimit(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	bids, err := c.service.GetBidsByProvider(ctx, params.ID.Address, offset, limit)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.BidsRes{Bids: bids})
	return
}

// GetActiveBidsByProvider godoc
//
//	@Summary		Get Bids by Provider
//	@Description	Get bids from blockchain by provider
//	@Tags			bids
//	@Produce		json
//	@Param			id	path		string	true	"Provider ID"
//	@Success		200	{object}	structs.BidsRes
//	@Router			/blockchain/providers/{id}/bids/active [get]
func (c *BlockchainController) getActiveBidsByProvider(ctx *gin.Context) {
	var params structs.PathEthAddrID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	offset, limit, err := getOffsetLimit(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	bids, err := c.service.GetActiveBidsByProvider(ctx, params.ID.Address, offset, limit)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.BidsRes{Bids: bids})
	return
}

// GetModels godoc
//
//	@Summary		Get models list
//	@Description	Get models list from blokchain
//	@Tags			models
//	@Produce		json
//	@Success		200	{object}	structs.ModelsRes
//	@Router			/blockchain/models [get]
func (c *BlockchainController) getAllModels(ctx *gin.Context) {
	models, err := c.service.GetAllModels(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.ModelsRes{Models: models})
	return
}

// GetBidsByModelAgent godoc
//
//	@Summary		Get Bids by	Model Agent
//	@Description	Get bids from blockchain by model agent
//	@Tags			bids
//	@Produce		json
//	@Param			offset	query		string	false	"Offset"
//	@Param			limit	query		string	false	"Limit"
//	@Param			id		path		string	true	"ModelAgent ID"
//	@Success		200		{object}	structs.BidsRes
//	@Router			/blockchain/models/{id}/bids [get]
func (c *BlockchainController) getBidsByModelAgent(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	offset, limit, err := getOffsetLimit(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	bids, err := c.service.GetBidsByModelAgent(ctx, params.ID.Hash, offset, limit)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.BidsRes{Bids: bids})
	return
}

// GetActiveBidsByModel godoc
//
//	@Summary		Get Active Bids by Model
//	@Description	Get bids from blockchain by model agent
//	@Tags			bids
//	@Produce		json
//	@Param			id	path		string	true	"ModelAgent ID"
//	@Success		200	{object}	structs.BidsRes
//	@Router			/blockchain/models/{id}/bids [get]
func (c *BlockchainController) getActiveBidsByModel(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	offset, limit, err := getOffsetLimit(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	bids, err := c.service.GetActiveBidsByModel(ctx, params.ID.Hash, offset, limit)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.BidsRes{Bids: bids})
	return
}

// GetBalance godoc
//
//	@Summary		Get ETH and MOR balance
//	@Description	Get ETH and MOR balance of the user
//	@Tags			transactions
//	@Produce		json
//	@Success		200	{object}	structs.TokenBalanceRes
//	@Router			/blockchain/balance [get]
func (s *BlockchainController) getBalance(ctx *gin.Context) {
	ethBalance, morBalance, err := s.service.GetBalance(ctx)
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.TokenBalanceRes{
		ETH: &lib.BigInt{Int: *ethBalance},
		MOR: &lib.BigInt{Int: *morBalance},
	})
	return
}

// GetTransactions godoc
//
//	@Summary		Get Transactions
//	@Description	Get MOR and ETH transactions
//	@Tags			transactions
//	@Produce		json
//	@Param			page	query		string	false	"Page"
//	@Param			limit	query		string	false	"Limit"
//	@Success		200		{object}	structs.TransactionsRes
//	@Router			/blockchain/transactions [get]
func (c *BlockchainController) getTransactions(ctx *gin.Context) {
	page, limit, err := getPageLimit(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	txs, err := c.service.GetTransactions(ctx, page, limit)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.TransactionsRes{Transactions: txs})
}

// GetAllowance godoc
//
//	@Summary		Get Allowance for MOR
//	@Description	Get MOR allowance for spender
//	@Tags			transactions
//	@Produce		json
//	@Param			spender	query		string	true	"Spender address"
//	@Success		200		{object}	structs.AllowanceRes
//	@Router			/blockchain/allowance [get]
func (c *BlockchainController) getAllowance(ctx *gin.Context) {
	var query structs.QuerySpender
	err := ctx.ShouldBindQuery(&query)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	allowance, err := c.service.GetAllowance(ctx, query.Spender.Address)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.AllowanceRes{Allowance: &lib.BigInt{Int: *allowance}})
	return
}

// Approve godoc
//
//	@Summary		Approve MOR allowance
//	@Description	Approve MOR allowance for spender
//	@Tags			transactions
//	@Produce		json
//	@Param			spender	query		string	true	"Spender address"
//	@Param			amount	query		string	true	"Amount"
//	@Success		200		{object}	structs.TxRes
//	@Router			/blockchain/approve [post]
func (c *BlockchainController) approve(ctx *gin.Context) {
	var query structs.QueryApprove
	err := ctx.ShouldBindQuery(&query)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	tx, err := c.service.Approve(ctx, query.Spender.Address, query.Amount.Unpack())
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.TxRes{Tx: tx})
	return
}

// OpenSession godoc
//
//	@Summary		Open Session with Provider in blockchain
//	@Description	Sends transaction in blockchain to open a session
//	@Tags			sessions
//	@Produce		json
//	@Accept			json
//	@Param			opensession	body		structs.OpenSessionRequest	true	"Open session"
//	@Success		200			{object}	structs.OpenSessionRes
//	@Router			/blockchain/sessions [post]
func (c *BlockchainController) openSession(ctx *gin.Context) {
	var reqPayload structs.OpenSessionRequest
	if err := ctx.ShouldBindJSON(&reqPayload); err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	sessionId, err := c.service.OpenSession(ctx, reqPayload.Approval, reqPayload.ApprovalSig, reqPayload.Stake.Unpack())
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.OpenSessionRes{SessionID: sessionId})
	return
}

// OpenSessionByBidId godoc
//
//	@Summary		Open Session by bidId in blockchain
//	@Description	Full flow to open a session by bidId
//	@Tags			sessions
//	@Produce		json
//	@Accept			json
//	@Param			opensession	body		structs.OpenSessionWithDurationRequest	true	"Open session"
//	@Param			id			path		string									true	"Bid ID"
//	@Success		200			{object}	structs.OpenSessionRes
//	@Router			/blockchain/bids/{id}/session [post]
func (s *BlockchainController) openSessionByBid(ctx *gin.Context) {
	var reqPayload structs.OpenSessionWithDurationRequest
	if err := ctx.ShouldBindJSON(&reqPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	sessionId, err := s.service.openSessionByBid(ctx, params.ID.Hash, reqPayload.SessionDuration.Unpack())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.OpenSessionRes{SessionID: sessionId})
	return
}

// OpenSessionByModelId godoc
//
//	@Summary		Open Session by ModelID in blockchain
//	@Description	Full flow to open a session by modelId
//	@Tags			sessions
//	@Produce		json
//	@Accept			json
//	@Param			opensession	body		structs.OpenSessionWithDurationRequest	true	"Open session"
//	@Param			id			path		string									true	"Model ID"
//	@Success		200			{object}	structs.OpenSessionRes
//	@Router			/blockchain/models/{id}/session [post]
func (s *BlockchainController) openSessionByModelId(ctx *gin.Context) {
	var reqPayload structs.OpenSessionWithDurationRequest
	if err := ctx.ShouldBindJSON(&reqPayload); err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sessionId, err := s.service.OpenSessionByModelId(ctx, params.ID.Hash, reqPayload.SessionDuration.Unpack())
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.OpenSessionRes{SessionID: sessionId})
	return
}

// CloseSession godoc
//
//	@Summary		Close Session with Provider
//	@Description	Sends transaction in blockchain to close a session
//	@Tags			sessions
//	@Produce		json
//	@Param			id	path		string	true	"Session ID"
//	@Success		200	{object}	structs.TxRes
//	@Router			/blockchain/sessions/{id}/close [post]
func (c *BlockchainController) closeSession(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	txHash, err := c.service.CloseSession(ctx, params.ID.Hash)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.TxRes{Tx: txHash})
	return
}

// GetSession godoc
//
//	@Summary		Get session
//	@Description	Returns session by ID
//	@Tags			sessions
//	@Produce		json
//	@Param			id	path		string	true	"Session ID"
//	@Success		200	{object}	structs.SessionRes
//	@Router			/blockchain/sessions/{id} [get]
func (c *BlockchainController) getSession(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	session, err := c.service.GetSession(ctx, params.ID.Hash)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.SessionRes{Session: session})
	return
}

// GetSessions godoc
//
//	@Summary		Get Sessions for User
//	@Description	Get sessions from blockchain by user
//	@Tags			sessions
//	@Produce		json
//	@Param			offset	query		string	false	"Offset"
//	@Param			limit	query		string	false	"Limit"
//	@Param			user	query		string	true	"User address"
//	@Success		200		{object}	structs.SessionsRes
//	@Router			/blockchain/sessions/user [get]
func (c *BlockchainController) getSessionsForUser(ctx *gin.Context) {
	offset, limit, err := getOffsetLimit(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	var req structs.QueryUser
	err = ctx.ShouldBindQuery(&req)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	sessions, err := c.service.GetSessions(ctx, req.User.Address, common.Address{}, offset, limit)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.SessionsRes{Sessions: sessions})
	return
}

// GetSessions godoc
//
//	@Summary		Get Sessions for Provider
//	@Description	Get sessions from blockchain by provider
//	@Tags			sessions
//	@Produce		json
//	@Param			offset		query		string	false	"Offset"
//	@Param			limit		query		string	false	"Limit"
//	@Param			provider	query		string	true	"Provider address"
//	@Success		200			{object}	structs.SessionsRes
//	@Router			/blockchain/sessions/provider [get]
func (c *BlockchainController) getSessionsForProvider(ctx *gin.Context) {
	offset, limit, err := getOffsetLimit(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	var req structs.QueryProvider
	err = ctx.ShouldBindQuery(&req)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	sessions, err := c.service.GetSessions(ctx, common.Address{}, req.Provider.Address, offset, limit)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.SessionsRes{Sessions: sessions})
	return
}

// GetTodaysBudget godoc
//
//	@Summary		Get Todays Budget
//	@Description	Get todays budget from blockchain
//	@Tags			sessions
//	@Produce		json
//	@Success		200	{object}	structs.BudgetRes
//	@Router			/blockchain/sessions/budget [get]
func (s *BlockchainController) getBudget(ctx *gin.Context) {
	budget, err := s.service.GetTodaysBudget(ctx)
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.BudgetRes{Budget: &lib.BigInt{Int: *budget}})
	return
}

// GetTokenSupply godoc
//
//	@Summary		Get Token Supply
//	@Description	Get MOR token supply from blockchain
//	@Tags			sessions
//	@Produce		json
//	@Success		200	{object}	structs.SupplyRes
//	@Router			/blockchain/token/supply [get]
func (s *BlockchainController) getSupply(ctx *gin.Context) {
	supply, err := s.service.GetTokenSupply(ctx)
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.SupplyRes{Supply: &lib.BigInt{Int: *supply}})
	return
}

// GetLatestBlock godoc
//
//	@Summary		Get Latest Block
//	@Description	Get latest block number from blockchain
//	@Tags			transactions
//	@Produce		json
//	@Success		200	{object}	structs.BlockRes
//	@Router			/blockchain/latestBlock [get]
func (c *BlockchainController) getLatestBlock(ctx *gin.Context) {
	block, err := c.service.GetLatestBlock(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, structs.BlockRes{Block: block})
	return
}

// GetBidByID godoc
//
//	@Summary		Get Bid by ID
//	@Description	Get bid from blockchain by ID
//	@Tags			bids
//	@Produce		json
//	@Param			id	path		string	true	"Bid ID"
//	@Success		200	{object}	structs.BidRes
//	@Router			/blockchain/bids/{id} [get]
func (c *BlockchainController) getBidByID(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	bid, err := c.service.GetBidByID(ctx, params.ID.Hash)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.BidRes{Bid: bid})
	return
}

// GetRatedBids godoc
//
//	@Summary		Get Rated Bids
//	@Description	Get rated bids from blockchain by model
//	@Tags			bids
//	@Produce		json
//	@Param			id	path		string	true	"Model ID"
//	@Success		200	{object}	structs.ScoredBidsRes
//	@Router			/blockchain/models/{id}/bids/rated [get]
func (c *BlockchainController) getRatedBids(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	bids, err := c.service.GetRatedBids(ctx, params.ID.Hash)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.ScoredBidsRes{Bids: bids})
	return
}

// СreateNewProvider godoc
//
//	@Summary	Creates or updates provider in blockchain
//	@Tags		providers
//	@Produce	json
//	@Accept		json
//	@Param		provider	body		structs.CreateProviderRequest	true	"Provider"
//	@Success	200			{object}	structs.ProviderRes
//	@Router		/blockchain/providers [post]
func (c *BlockchainController) createProvider(ctx *gin.Context) {
	var provider structs.CreateProviderRequest
	if err := ctx.ShouldBindJSON(&provider); err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	result, err := c.service.CreateNewProvider(ctx, provider.Stake, provider.Endpoint)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.ProviderRes{Provider: result})
	return
}

// DeregisterProvider godoc
//
//	@Summary	Deregister Provider
//	@Tags		providers
//	@Produce	json
//	@Success	200	{object}	structs.TxRes
//	@Router		/blockchain/providers/{id} [delete]
func (c *BlockchainController) deregisterProvider(ctx *gin.Context) {
	txHash, err := c.service.DeregisterProdiver(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.TxRes{Tx: txHash})
	return
}

// CreateNewModel godoc
//
//	@Summary	Creates model in blockchain
//	@Tags		models
//	@Produce	json
//	@Accept		json
//	@Param		model	body		structs.CreateModelRequest	true	"Model"
//	@Success	200		{object}	structs.ModelRes
//	@Router		/blockchain/models [post]
func (c *BlockchainController) createNewModel(ctx *gin.Context) {
	var model structs.CreateModelRequest
	if err := ctx.ShouldBindJSON(&model); err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	var modelId common.Hash
	if model.ID == "" {
		var hash common.Hash
		_, err := rand.Read(hash[:])
		if err != nil {
			c.log.Error(err)
			ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
			return
		}
		modelId = hash
	} else {
		modelId = common.HexToHash(model.ID)
	}
	ipsfHash := common.HexToHash(model.IpfsID)

	result, err := c.service.CreateNewModel(ctx, modelId, ipsfHash, model.Fee, model.Stake, model.Name, model.Tags)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.ModelRes{Model: result})
	return

}

// DeregisterModel godoc
//
//	@Summary	Deregister Model
//	@Tags		models
//	@Produce	json
//	@Param		id	path		string	true	"Model ID"
//	@Success	200	{object}	structs.TxRes
//	@Router		/blockchain/models/{id} [delete]
func (c *BlockchainController) deregisterModel(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	txHash, err := c.service.DeregisterModel(ctx, params.ID.Hash)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.TxRes{Tx: txHash})
	return
}

// CreateBidRequest godoc
//
//	@Summary	Creates bid in blockchain
//	@Tags		bids
//	@Produce	json
//	@Accept		json
//	@Param		bid	body		structs.CreateBidRequest	true	"Bid"
//	@Success	200	{object}	structs.BidRes
//	@Router		/blockchain/bids [post]
func (c *BlockchainController) createNewBid(ctx *gin.Context) {
	var bid structs.CreateBidRequest
	if err := ctx.ShouldBindJSON(&bid); err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	modelId := common.HexToHash(bid.ModelID)
	result, err := c.service.CreateNewBid(ctx, modelId, bid.PricePerSecond)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.BidRes{Bid: result})
	return
}

// DeleteBid godoc
//
//	@Summary	Delete Bid
//	@Tags		bids
//	@Produce	json
//	@Param		id	path		string	true	"Bid ID"
//	@Success	200	{object}	structs.TxRes
//	@Router		/blockchain/bids/{id} [delete]
func (c *BlockchainController) deleteBid(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, structs.ErrRes{Error: err.Error()})
		return
	}

	txHash, err := c.service.DeleteBid(ctx, params.ID.Hash)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, structs.ErrRes{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, structs.TxRes{Tx: txHash})
	return
}

// helpers

func (s *BlockchainController) getSendParams(ctx *gin.Context) (to common.Address, amount *big.Int, err error) {
	var body structs.SendRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		return common.Address{}, nil, err
	}
	return body.To, body.Amount.Unpack(), nil
}

func getOffsetLimit(ctx *gin.Context) (offset *big.Int, limit uint8, err error) {
	var paging structs.QueryOffsetLimit

	err = ctx.ShouldBindQuery(&paging)
	if err != nil {
		return nil, 0, err
	}

	return paging.Offset.Unpack(), paging.Limit, nil
}

func getPageLimit(ctx *gin.Context) (page uint64, limit uint8, err error) {
	var paging structs.QueryPageLimit

	err = ctx.ShouldBindQuery(&paging)
	if err != nil {
		return 0, 0, err
	}

	return paging.Page, paging.Limit, nil
}
