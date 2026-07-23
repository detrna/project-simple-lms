package middleware

import (
	"main/internal/domain"
	"main/internal/pkg"
	"main/internal/shared"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtProvider pkg.JWTProvider, logger pkg.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessTokenPayload, err := parseAccessToken(c, jwtProvider)

		if err != nil {
			shared.HandleError(c, logger, err)
			return
		}

		c.Set("user", accessTokenPayload)

		c.Next()
	}
}

func parseAccessToken(c *gin.Context, tokenService pkg.JWTProvider) (*domain.JWTPayload, error) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		return nil, shared.ErrUnauthorized
	}

	accessToken := strings.TrimPrefix(authHeader, "Bearer ")

	if accessToken == authHeader {
		return nil, shared.ErrUnauthorized
	}

	return tokenService.ParseAccessToken(string(accessToken))
}
