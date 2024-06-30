package walletapi

import (
	"net/http"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/gin-gonic/gin"
)

type WalletController struct {
	service interfaces.Wallet
}

func NewWalletController(service interfaces.Wallet) *WalletController {
	c := &WalletController{
		service: service,
	}

	return c
}

func (s *WalletController) RegisterRoutes(r interfaces.Router) {
	r.GET("/wallet", s.GetWallet)
	r.POST("/wallet", s.SetupWallet)
}

// GetWallet godoc
//
//		@Summary		Get Wallet
//		@Description	Get wallet address
//	 	@Tags			wallet
//		@Produce		json
//		@Success		200	{object}	interface{}
//		@Router			/wallet [get]
func (s *WalletController) GetWallet(ctx *gin.Context) {
	prKey, err := s.service.GetPrivateKey()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	addr, err := lib.PrivKeyBytesToAddr(prKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"address": addr})
}

// SetupWallet godoc
//
//		@Summary		Set Wallet
//		@Description	Set wallet private key
//	 	@Tags			wallet
//		@Produce		json
//		@Param			privatekey	body	walletapi.SetupWalletReqBody true	"Private key"
//		@Success		200	{object}	interface{}
//		@Router			/wallet [post]
func (s *WalletController) SetupWallet(ctx *gin.Context) {
	var req SetupWalletReqBody
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = s.service.SetPrivateKey(req.PrivateKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
