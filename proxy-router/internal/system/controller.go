package system

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/config"
	i "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/gin-gonic/gin"
)

type SystemController struct {
	config       *config.Config
	wallet       i.Wallet
	ethRPC       i.RPCEndpoints
	sysConfig    *SystemConfigurator
	appStartTime time.Time
	chainID      *big.Int
	log          lib.ILogger
}

func NewSystemController(config *config.Config, wallet i.Wallet, ethRPC i.RPCEndpoints, sysConfig *SystemConfigurator, appStartTime time.Time, chainID *big.Int, log lib.ILogger) *SystemController {
	c := &SystemController{
		config:       config,
		wallet:       wallet,
		ethRPC:       ethRPC,
		sysConfig:    sysConfig,
		appStartTime: appStartTime,
		chainID:      chainID,
		log:          log,
	}

	return c
}

func (s *SystemController) RegisterRoutes(r i.Router) {
	r.GET("/healthcheck", s.HealthCheck)
	r.GET("/config", s.GetConfig)
	r.GET("/files", s.GetFiles)

	r.POST("/config/ethNode", s.SetEthNode)
}

// HealthCheck godoc
//
//	@Summary		Healthcheck example
//	@Description	do ping
//	@Tags			system
//	@Produce		json
//	@Success		200	{object}	HealthCheckResponse
//	@Router			/healthcheck [get]
func (s *SystemController) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, HealthCheckResponse{
		Status:  "healthy",
		Version: config.BuildVersion,
		Uptime:  time.Since(s.appStartTime).Round(time.Second).String(),
	})
}

// GetConfig godoc
//
//	@Summary		Get Config
//	@Description	Return the current config of proxy router
//	@Tags			system
//	@Produce		json
//	@Success		200	{object}	ConfigResponse
//	@Router			/config [get]
func (s *SystemController) GetConfig(ctx *gin.Context) {
	prkey, err := s.wallet.GetPrivateKey()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	addr, err := lib.PrivKeyBytesToAddr(prkey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &ConfigResponse{
		Version: config.BuildVersion,
		Commit:  config.Commit,
		Config:  s.config.GetSanitized(),
		DerivedConfig: config.DerivedConfig{
			WalletAddress: addr,
			ChainID:       s.chainID,
			EthNodeURLs:   s.ethRPC.GetURLs(),
		},
	})
}

// GetFiles godoc
//
//	@Summary		Get files
//	@Description	Returns opened files
//	@Tags			system
//	@Produce		json
//	@Success		200	{object}	[]FD
//	@Router			/files [get]
func (s *SystemController) GetFiles(ctx *gin.Context) {
	files, err := s.sysConfig.GetFileDescriptors(ctx, os.Getpid())
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	systemCfg, err := s.sysConfig.GetConfig()
	if err != nil {
		fmt.Fprintf(ctx.Writer, "failed to get system config: %s\n", err)
	} else {
		json, err := json.Marshal(systemCfg)
		if err != nil {
			fmt.Fprintf(ctx.Writer, "failed to marshal system config: %s\n", err)
		} else {
			fmt.Fprintf(ctx.Writer, "system config: %s\n", json)
		}
	}
	fmt.Fprintf(ctx.Writer, "\n")

	err = writeFiles(ctx.Writer, files)
	if err != nil {
		s.log.Errorf("failed to write files: %s", err)
		_ = ctx.Error(err)
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

// SetEthNode godoc
//
//	@Summary		Set Eth Node URLs
//	@Description	Set the Eth Node URLs
//	@Tags			system
//	@Accept			json
//	@Produce		json
//	@Param			urls	body		SetEthNodeURLReq	true	"URLs"
//	@Success		200		{object}	ConfigResponse
//	@Router			/config/ethNode [post]
func (s *SystemController) SetEthNode(ctx *gin.Context) {
	var req SetEthNodeURLReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.ethRPC.SetURLs(req.URLs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// DeleteEthNode godoc
//
//	@Summary		Delete Eth Node URLs
//	@Description	Delete the Eth Node URLs
//	@Tags			system
//	@Produce		json
//	@Success		200	{object}	ConfigResponse
//	@Router			/config/ethNode [delete]
func (c *SystemController) RemoveEthNode(ctx *gin.Context) {
	err := c.ethRPC.RemoveURLs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func writeFiles(writer io.Writer, files []FD) error {
	text := fmt.Sprintf("Total: %d\n", len(files))
	text += "\n"
	text += "fd\tpath\n"

	if _, err := fmt.Fprint(writer, text); err != nil {
		return err
	}

	for _, f := range files {
		if _, err := fmt.Fprintf(writer, "%s\t%s\n", f.ID, f.Path); err != nil {
			return err
		}
	}

	return nil
}
