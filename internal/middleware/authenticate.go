package middleware

import (
	"main/internal/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtProvider pkg.JWTProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")[1]

		if string(token) == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token didn't exist"})
		}

		jwtPayload, err := jwtProvider.ParseAccessToken(string(token))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalid"})
		}

		c.Set("user", jwtPayload)

		c.Next()
	}
}
