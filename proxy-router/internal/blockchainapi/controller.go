package blockchainapi

import (
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
	r.GET("/proxy/sessions/:id/providerClaimableBalance", c.getProviderClaimableBalance)
	r.POST("/proxy/sessions/:id/providerClaim", c.claimProviderBalance)

	r.GET("/blockchain/balance", c.getBalance)
	r.POST("/blockchain/send/eth", c.sendETH)
	r.POST("/blockchain/send/mor", c.sendMOR)
	r.GET("/blockchain/transactions", c.getTransactions)
	r.GET("/blockchain/allowance", c.getAllowance)
	r.POST("/blockchain/approve", c.approve)
	r.GET("/blockchain/latestBlock", c.getLatestBlock)

	r.GET("/blockchain/providers", c.getAllProviders)
	r.GET("/blockchain/providers/:id/bids", c.getBidsByProvider)
	r.GET("/blockchain/models", c.getAllModels)
	r.GET("/blockchain/models/:id/bids", c.getBidsByModelAgent)
	r.GET("/blockchain/bids/:id", c.getBidByID)
	r.GET("/blockchain/sessions", c.getSessions)
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
//		@Summary		Get Provider Claimable Balance
//		@Description	Get provider claimable balance from session
//	 	@Tags			sessions
//		@Produce		json
//		@Param 			id  path string true "Session ID"
//		@Success		200	{object}	interface{}
//		@Router			/proxy/sessions/${id}/providerClaimableBalance [get]
func (c *BlockchainController) getProviderClaimableBalance(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	balance, err := c.service.GetProviderClaimableBalance(ctx, params.ID.Hash)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"balance": balance})
	return
}

// ClaimProviderBalance godoc
//
//		@Summary		Claim Provider Balance
//		@Description	Claim provider balance from session
//	 	@Tags			sessions
//		@Produce		json
//		@Param			claim	body		structs.SendRequest 	true	"Claim"
//		@Param 			id  path string true "Session ID"
//		@Success		200	{object}	interface{}
//		@Router			/proxy/sessions/${id}/providerClaim [post]
func (c *BlockchainController) claimProviderBalance(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	to, amount, err := c.getSendParams(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txHash, err := c.service.ClaimProviderBalance(ctx, params.ID.Hash, to, amount)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"txHash": txHash.String()})
	return
}

// GetProviders godoc
//
//		@Summary		Get providers list
//		@Description	Get providers list from blokchain
//	 	@Tags			wallet
//		@Produce		json
//		@Success		200	{object}	[]interface{}
//		@Router			/blockchain/providers [get]
func (c *BlockchainController) getAllProviders(ctx *gin.Context) {
	providers, err := c.service.GetAllProviders(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"providers": providers})
	return
}

// SendEth godoc
//
//		@Summary		Send Eth
//		@Description	Send Eth to address
//	 	@Tags			wallet
//		@Produce		json
//		@Param			sendeth	body		structs.SendRequest 	true	"Send Eth"
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/send/eth [post]
func (c *BlockchainController) sendETH(ctx *gin.Context) {
	to, amount, err := c.getSendParams(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txHash, err := c.service.SendETH(ctx, to, amount)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"txHash": txHash.String()})
	return
}

// SendMor godoc
//
//		@Summary		Send Mor
//		@Description	Send Mor to address
//	 	@Tags			wallet
//		@Produce		json
//		@Param			sendmor	body		structs.SendRequest 	true	"Send Mor"
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/send/mor [post]
func (c *BlockchainController) sendMOR(ctx *gin.Context) {
	to, amount, err := c.getSendParams(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	txhash, err := c.service.SendMOR(ctx, to, amount)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"txHash": txhash.String()})
	return
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
func (c *BlockchainController) getBidsByProvider(ctx *gin.Context) {
	var params structs.PathEthAddrID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offset, limit, err := getOffsetLimit(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bids, err := c.service.GetBidsByProvider(ctx, params.ID, offset, limit)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"bids": bids})
	return
}

// GetModels godoc
//
//		@Summary		Get models list
//		@Description	Get models list from blokchain
//	 	@Tags			wallet
//		@Produce		json
//		@Success		200	{object}	[]interface{}
//		@Router			/blockchain/models [get]
func (c *BlockchainController) getAllModels(ctx *gin.Context) {
	providers, err := c.service.GetAllModels(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"models": providers})
	return
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
func (c *BlockchainController) getBidsByModelAgent(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offset, limit, err := getOffsetLimit(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bids, err := c.service.GetBidsByModelAgent(ctx, params.ID.Hash, offset, limit)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"bids": bids})
	return
}

// GetBalance godoc
//
//		@Summary		Get ETH and MOR balance
//		@Description	Get ETH and MOR balance of the user
//	 	@Tags			wallet
//		@Produce		json
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/balance [get]
func (s *BlockchainController) getBalance(ctx *gin.Context) {
	ethBalance, morBalance, err := s.service.GetBalance(ctx)
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"eth": ethBalance.String(), "mor": morBalance.String()})
	return
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
func (c *BlockchainController) getTransactions(ctx *gin.Context) {
	page, limit, err := getPageLimit(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txs, err := c.service.GetTransactions(ctx, page, limit)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"transactions": txs})
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
func (c *BlockchainController) getAllowance(ctx *gin.Context) {
	var query structs.QuerySpender
	err := ctx.ShouldBindQuery(&query)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	allowance, err := c.service.GetAllowance(ctx, query.Spender)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"allowance": allowance.String()})
	return
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
func (c *BlockchainController) approve(ctx *gin.Context) {
	var query structs.QueryApprove
	err := ctx.ShouldBindQuery(&query)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := c.service.Approve(ctx, query.Spender, query.Amount.Unpack())
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tx": tx.Hex()})
	return
}

// OpenSession godoc
//
//		@Summary		Open Session with Provider in blockchain
//		@Description	Sends transaction in blockchain to open a session
//	 	@Tags			sessions
//		@Produce		json
//		@Accept			json
//		@Param			opensession	body		structs.OpenSessionRequest 	true	"Open session"
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/sessions [post]
func (c *BlockchainController) openSession(ctx *gin.Context) {
	var reqPayload structs.OpenSessionRequest
	if err := ctx.ShouldBindJSON(&reqPayload); err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txHash, err := c.service.OpenSession(ctx, reqPayload.Approval, reqPayload.ApprovalSig, reqPayload.Stake.Unpack())
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tx": txHash.Hex()})
	return
}

// OpenSessionByBidId godoc
//
//		@Summary		Open Session by bidId in blockchain
//		@Description	Full flow to open a session by bidId
//	 	@Tags			sessions
//		@Produce		json
//		@Accept			json
//		@Param			opensession	body		structs.OpenSessionWithDurationRequest 	true	"Open session"
//		@Param 			id  path string true "Bid ID"
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/bids/:id/session [post]
func (s *BlockchainController) openSessionByBid(ctx *gin.Context) {
	var reqPayload structs.OpenSessionWithDurationRequest
	if err := ctx.ShouldBindJSON(&reqPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txHash, err := s.service.openSessionByBid(ctx, params.ID.Hash, reqPayload.SessionDuration.Unpack())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tx": txHash.Hex()})
	return
}

// OpenSessionByModelId godoc
//
//		@Summary		Open Session by ModelID in blockchain
//		@Description	Full flow to open a session by modelId
//	 	@Tags			sessions
//		@Produce		json
//		@Accept			json
//		@Param			opensession	body		structs.OpenSessionWithDurationRequest 	true	"Open session"
//		@Param 			id  path string true "Model ID"
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/models/:id/session [post]
func (s *BlockchainController) openSessionByModelId(ctx *gin.Context) {
	var reqPayload structs.OpenSessionWithDurationRequest
	if err := ctx.ShouldBindJSON(&reqPayload); err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txHash, err := s.service.OpenSessionByModelId(ctx, params.ID.Hash, reqPayload.SessionDuration.Unpack())
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tx": txHash.Hex()})
	return
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
func (c *BlockchainController) closeSession(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txHash, err := c.service.CloseSession(ctx, params.ID.Hash)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tx": txHash.Hex()})
	return
}

func (c *BlockchainController) getSession(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := c.service.GetSession(ctx, params.ID.Hash)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"session": session})
	return
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
func (c *BlockchainController) getSessions(ctx *gin.Context) {
	offset, limit, err := getOffsetLimit(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req structs.QueryUserOrProvider
	err = ctx.ShouldBindQuery(&req)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hasUser := req.User != lib.Address{}
	hasProvider := req.Provider != lib.Address{}

	if !hasUser && !hasProvider {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user or provider is required"})
		return
	}
	if hasUser && hasProvider {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "only one of user or provider is allowed"})
		return
	}

	sessions, err := c.service.GetSessions(ctx, req.User.Address, req.Provider.Address, offset, limit)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"sessions": sessions})
	return
}

// GetTodaysBudget godoc
//
//		@Summary		Get Todays Budget
//		@Description	Get todays budget from blockchain
//	 	@Tags			wallet
//		@Produce		json
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/sessions/budget [get]
func (s *BlockchainController) getBudget(ctx *gin.Context) {
	budget, err := s.service.GetTodaysBudget(ctx)
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"budget": budget.String()})
	return
}

// GetTokenSupply godoc
//
//		@Summary		Get Token Supply
//		@Description	Get MOR token supply from blockchain
//	 	@Tags			wallet
//		@Produce		json
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/token/supply [get]
func (s *BlockchainController) getSupply(ctx *gin.Context) {
	supply, err := s.service.GetTokenSupply(ctx)
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"supply": supply.String()})
	return
}

// GetLatestBlock godoc
//
//		@Summary		Get Latest Block
//		@Description	Get latest block number from blockchain
//	 	@Tags			wallet
//		@Produce		json
//		@Success		200	{object}	interface{}
//		@Router			/blockchain/latestBlock [get]
func (c *BlockchainController) getLatestBlock(ctx *gin.Context) {
	block, err := c.service.GetLatestBlock(ctx)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"block": block})
	return
}

func (c *BlockchainController) getBidByID(ctx *gin.Context) {
	var params structs.PathHex32ID
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bid, err := c.service.GetBidByID(ctx, params.ID.Hash)
	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"bid": bid})
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
