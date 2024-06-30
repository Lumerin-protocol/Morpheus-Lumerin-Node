package httphandlers

import (
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/gin-gonic/gin"
)

// RequestLogger is a middleware for logging HTTP requests
func RequestLogger(logger lib.ILogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		if raw != "" {
			path = path + "?" + raw
		}

		// Log details
		logger.Infof("[REQ] %s %s [%d] \n [ERROR]: %s",
			c.Request.Method,
			path,
			c.Writer.Status(),
			c.Errors.ByType(gin.ErrorTypePrivate).String(),
		)
	}
}
