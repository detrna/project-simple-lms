package middleware

import (
	"fmt"
	"main/internal/domain"
	"main/internal/pkg"
	"main/internal/shared"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtProvider pkg.TokenService, logger pkg.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessTokenPayload, err := parseHeader(c, jwtProvider)

		if err != nil {
			shared.HandleError(c, logger, err)
			return
		}

		c.Set("user", accessTokenPayload)

		c.Next()
	}
}

func parseHeader(c *gin.Context, tokenService pkg.TokenService) (*domain.JWTPayload, error) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		return nil, shared.ErrTokenMissing
	}

	accessToken := strings.TrimPrefix(authHeader, "Bearer ")

	fmt.Print("TOKENIZER " + accessToken)

	return tokenService.ParseAccessToken(string(accessToken))
}
