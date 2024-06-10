package httphandlers

import (
	"fmt"
	"math/big"
	"net/http"
	"net/http/pprof"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/internal/apibus"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	// gin-swagger middleware
	swaggerFiles "github.com/swaggo/files"

	_ "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/docs"
)

const (
	SUCCESS_STATUS = 200
	ERROR_STATUS   = 500
)

type HTTPHandler struct{}

// @title           ApiBus Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8082
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func NewHTTPHandler(apiBus *apibus.ApiBus) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/healthcheck", (func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, apiBus.HealthCheck(ctx))
	}))
	r.GET("/config", (func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, apiBus.GetConfig(ctx))
	}))
	r.GET("/files", (func(ctx *gin.Context) {
		status, files := apiBus.GetFiles(ctx)
		
		ctx.JSON(status, files)
	}))
	r.POST("/v1/chat/completions", (func(ctx *gin.Context) {
		apiBus.PromptLocal(ctx)
	}))

	r.POST("/proxy/sessions/initiate", (func(ctx *gin.Context) {
		status, response := apiBus.InitiateSession(ctx)
		ctx.JSON(status, response)
	}))

	r.POST("/proxy/sessions/:id/prompt", (func(ctx *gin.Context) {
		if ok, status, response := apiBus.SendPrompt(ctx); !ok {
			ctx.JSON(status, response)
		}
	}))

	r.GET("/proxy/sessions/:id/providerClaimableBalance", (func(ctx *gin.Context) {
		status, response := apiBus.GetProviderClaimableBalance(ctx)
		ctx.JSON(status, response)
	}))

	r.POST("/proxy/sessions/:id/providerClaim", (func(ctx *gin.Context) {
		status, response := apiBus.ClaimProviderBalance(ctx)
		ctx.JSON(status, response)
	}))

	r.GET("/blockchain/providers", (func(ctx *gin.Context) {
		status, providers := apiBus.GetAllProviders(ctx)
		ctx.JSON(status, providers)
	}))

	r.POST("/blockchain/providers", (func(ctx *gin.Context) {
		address := ctx.GetString("address")
		addStake := ctx.GetUint64("addStake")
		endpoint := ctx.GetString("endpoint")

		status, response := apiBus.CreateNewProvider(ctx, address, addStake, endpoint)

		ctx.JSON(status, response)
	}))

	r.POST("/blockchain/send/eth", (func(ctx *gin.Context) {
		status, response := apiBus.SendEth(ctx)
		ctx.JSON(status, response)
	}))

	r.POST("/blockchain/send/mor", (func(ctx *gin.Context) {
		status, response := apiBus.SendMor(ctx)
		ctx.JSON(status, response)
	}))

	r.GET("/blockchain/providers/:id/bids", (func(ctx *gin.Context) {
		providerId := ctx.Param("id")
		offset, limit := getOffsetLimit(ctx)

		if offset == nil {
			return
		}

		status, bids := apiBus.GetBidsByProvider(ctx, providerId, offset, limit)
		ctx.JSON(status, bids)
	}))

	r.GET("/blockchain/models", (func(ctx *gin.Context) {
		status, models := apiBus.GetAllModels(ctx)
		ctx.JSON(status, models)
	}))

	r.GET("/blockchain/models/:id/bids", (func(ctx *gin.Context) {
		modelAgentId := ctx.Param("id")

		offset, limit := getOffsetLimit(ctx)
		if offset == nil {
			return
		}

		id := common.FromHex(modelAgentId)

		status, models := apiBus.GetBidsByModelAgent(ctx, ([32]byte)(id), offset, limit)
		ctx.JSON(status, models)
	}))

	r.GET("/blockchain/balance", (func(ctx *gin.Context) {
		status, balance := apiBus.GetBalance(ctx)
		ctx.JSON(status, balance)
	}))

	r.GET("/blockchain/transactions", (func(ctx *gin.Context) {
		status, transactions := apiBus.GetTransactions(ctx)
		ctx.JSON(status, transactions)
	}))

	r.GET("/blockchain/allowance", (func(ctx *gin.Context) {
		status, balance := apiBus.GetAllowance(ctx)
		ctx.JSON(status, balance)
	}))

	r.POST("/blockchain/approve", (func(ctx *gin.Context) {
		status, response := apiBus.Approve(ctx)
		ctx.JSON(status, response)
	}))

	r.POST("/blockchain/sessions", (func(ctx *gin.Context) {
		status, response := apiBus.OpenSession(ctx)
		ctx.JSON(status, response)
	}))

	r.GET("/blockchain/sessions", (func(ctx *gin.Context) {
		offset, limit := getOffsetLimit(ctx)
		if offset == nil {
			return
		}
		status, response := apiBus.GetSessions(ctx, offset, limit)
		ctx.JSON(status, response)
	}))

	r.GET("/blockchain/sessions/budget", (func(ctx *gin.Context) {
		status, response := apiBus.GetTodaysBudget(ctx)
		ctx.JSON(status, response)
	}))

	r.GET("/blockchain/token/supply", (func(ctx *gin.Context) {
		status, response := apiBus.GetTokenSupply(ctx)
		ctx.JSON(status, response)
	}))

	r.POST("/blockchain/sessions/:id/close", (func(ctx *gin.Context) {
		status, response := apiBus.CloseSession(ctx)
		ctx.JSON(status, response)
	}))

	r.Any("/debug/pprof/*action", gin.WrapF(pprof.Index))

	if err := r.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	return r
}

func getOffsetLimit(ctx *gin.Context) (*big.Int, uint8) {
	offsetStr := ctx.Query("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}
	limitStr := ctx.Query("limit")
	if limitStr == "" {
		limitStr = "100"
	}

	offset, ok := new(big.Int).SetString(offsetStr, 10)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset"})
		return nil, 0
	}

	var limit uint8
	_, err := fmt.Sscanf(limitStr, "%d", &limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
		return nil, 0
	}
	return offset, limit
}
