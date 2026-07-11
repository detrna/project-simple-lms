package middleware

import (
	"main/internal/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")[1]

		if string(token) == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token didn't exist"})
		}

		claims, err := shared.ParseToken(string(token), secretKey)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalid"})
		}

		c.Set("user", claims.Payload)

		c.Next()
	}
}
