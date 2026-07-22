package middleware

import (
	"main/internal/domain"
	"main/internal/pkg"
	"main/internal/shared"

	"github.com/gin-gonic/gin"
)

func RequiredRole(role string, logger pkg.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exist := c.Get("user")

		if !exist {
			shared.HandleError(c, logger, shared.ErrUnauthorized)
			return
		}

		user, ok := value.(domain.JWTPayload)

		if !ok {
			shared.HandleError(c, logger, shared.ErrUnauthorized)
			return
		}

		if user.Role != role {
			shared.HandleError(c, logger, shared.ErrForbidden)
			return
		}

		c.Next()
	}
}
