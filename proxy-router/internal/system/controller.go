package system

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/config"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/gin-gonic/gin"
)

type HealthCheckResponse struct {
	Status  string
	Version string
	Uptime  string
}

type SystemController struct {
	config        *config.Config
	derivedConfig *config.DerivedConfig
	sysConfig     *SystemConfigurator
	appStartTime  time.Time
	log           lib.ILogger
}

type ConfigResponse struct {
	Version       string
	Commit        string
	DerivedConfig interface{}
	Config        interface{}
}

func NewSystemController(r interfaces.Router, config *config.Config, derived *config.DerivedConfig, sysConfig *SystemConfigurator, appStartTime time.Time, log lib.ILogger) *SystemController {
	c := &SystemController{
		config:        config,
		derivedConfig: derived,
		sysConfig:     sysConfig,
		appStartTime:  appStartTime,
		log:           log,
	}

	r.GET("/healthcheck", c.HealthCheck)
	r.GET("/config", c.GetConfig)
	r.GET("/files", c.GetFiles)

	return c
}

// HealthCheck godoc
//
//		@Summary		Healthcheck example
//		@Description	do ping
//	 	@Tags			healthcheck
//		@Produce		json
//		@Success		200	{object}	HealthCheckResponse
//		@Router			/healthcheck [get]
func (s *SystemController) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, HealthCheckResponse{
		Status:  "healthy",
		Version: config.BuildVersion,
		Uptime:  time.Since(s.appStartTime).Round(time.Second).String(),
	})
}

// GetConfig godoc
//
//		@Summary		Get Config
//		@Description	Return the current config of proxy router
//	 	@Tags				healthcheck
//		@Produce		json
//		@Success		200	{object}	ConfigResponse
//		@Router			/config [get]
func (s *SystemController) GetConfig(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &ConfigResponse{
		Version:       config.BuildVersion,
		Commit:        config.Commit,
		Config:        s.config.GetSanitized(),
		DerivedConfig: s.derivedConfig,
	})
}

// GetFiles godoc
//
//		@Summary		Get files
//		@Description	Returns opened files
//	 	@Tags				healthcheck
//		@Produce		json
//		@Success		200	{object}	[]FD
//		@Router			/files [get]
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
