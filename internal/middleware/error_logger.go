package middleware

import (
	"main/internal/pkg"

	"github.com/gin-gonic/gin"
)

func ErrorLogger(logger pkg.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		logger.ErrorLog(
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			err,
		)
	}
}
