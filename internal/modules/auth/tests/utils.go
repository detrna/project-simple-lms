package auth_test_helper

import (
	"main/internal/modules/auth"
	pkg_mocks "main/internal/pkg/mocks"
	shared_testing "main/internal/shared/testing_helper"
	"testing"

	"github.com/gin-gonic/gin"
)

func NewUseCasePkgs(t *testing.T) *auth.UseCasePackages {
	return &auth.UseCasePackages{
		Bcrypt: pkg_mocks.NewMockBcryptHasher(t),
		// Mailer:        pkg_mocks.NewMockResendClient(t),
		TokenProvider: pkg_mocks.NewMockJWTProvider(t),
		Redis:         pkg_mocks.NewMockRedisClient(t),
		Logger:        shared_testing.NewMockLogger(t),
	}
}

func SetRefreshTokenCookie(ctx *gin.Context, token string) {
	ctx.SetCookie(
		"refresh_token", // name
		token,           // value
		3600,            // maxAge (seconds)
		"/",             // path
		"",              // domain
		false,           // secure
		true,            // httpOnly
	)
}
