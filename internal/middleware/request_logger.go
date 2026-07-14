package middleware

import (
	"main/internal/pkg"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestLogger(logger pkg.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := uuid.NewString()

		c.Set("requestID", requestID)

		c.Next()

		logger.RequestLog(
			requestID,
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			start,
		)
	}
}
