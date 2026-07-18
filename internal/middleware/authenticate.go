package middleware

import (
	"main/internal/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtProvider pkg.JWTProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "Token didn't exist"},
			)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token == authHeader {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "Invalid authorization format"},
			)
			return
		}
	}
}
