package httphandlers

import (
	"net/http/pprof"

	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/internal/apibus"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	ginSwagger "github.com/swaggo/gin-swagger"

	// gin-swagger middleware
	swaggerFiles "github.com/swaggo/files"

	_ "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/docs"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/config"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
)

type Registrable interface {
	RegisterRoutes(r interfaces.Router)
}

// @title           ApiBus Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func CreateHTTPServer(log lib.ILogger, controllers ...Registrable) *gin.Engine {
	ginValidatorInstance := binding.Validator.Engine().(*validator.Validate)
	err := config.RegisterHex32(ginValidatorInstance)
	if err != nil {
		panic(err)
	}
	err = config.RegisterDuration(ginValidatorInstance)
	if err != nil {
		panic(err)
	}
	err = config.RegisterEthAddr(ginValidatorInstance)
	if err != nil {
		panic(err)
	}
	err = config.RegisterHexadecimal(ginValidatorInstance)
	if err != nil {
		panic(err)
	}

	// gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"session_id"},
	}))

	// r.Use(RequestLogger(log))

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
		shouldSendResponse, status, response := apiBus.RemoteOrLocalPrompt(ctx)
		if !shouldSendResponse {
			return
		}
		ctx.JSON(status, response)
	}))

	r.POST("/proxy/sessions/initiate", (func(ctx *gin.Context) {
		status, response := apiBus.InitiateSession(ctx)
		ctx.JSON(status, response)
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
		status, response := apiBus.CreateNewProvider(ctx)
		ctx.JSON(status, response)
	}))

	r.POST("/blockchain/providers/bids", (func(ctx *gin.Context) {
		status, response := apiBus.CreateNewBid(ctx)
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
		fmt.Printf("POST /blockchain/sessions\n")
		fmt.Printf("body: %+v\n", ctx.Request.Body)
		fmt.Println("approval: ", ctx.GetString("approval"))
		status, response := apiBus.OpenSession(ctx)
		ctx.JSON(status, response)
	}))

	r.POST("/blockchain/sessions/v2", (func(ctx *gin.Context) {
		status, response := apiBus.OpenSessionV2(ctx)
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

	r.POST("/wallet", (func(ctx *gin.Context) {

		var req SetupWalletReqBody
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = apiBus.SetupWallet(ctx, req.PrivateKey)
		if err != nil {
			fmt.Println("wallet error: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}))

	r.GET("/wallet", (func(ctx *gin.Context) {

		address, err := apiBus.GetAddress(ctx)
		if err != nil {
			fmt.Println("wallet error: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"address": address})
	}))

	r.Any("/debug/pprof/*action", gin.WrapF(pprof.Index))

	for _, c := range controllers {
		c.RegisterRoutes(r)
	}

	if err := r.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	return r
}
