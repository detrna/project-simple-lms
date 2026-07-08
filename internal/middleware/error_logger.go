package middleware

import (
	"log/slog"
	"main/internal/shared"

	"github.com/gin-gonic/gin"
)

func ErrorLogger(logger shared.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		logger.Error(
			"request failed",
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.Int("status", c.Writer.Status()),
			slog.Any("error", err),
		)
	}
}
