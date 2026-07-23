package auth_test_helper

import (
	"main/internal/modules/auth"
	pkg_mocks "main/internal/pkg/mocks"
	shared_testing "main/internal/shared/testing_helper"
	"testing"
)

func NewUseCasePkgs(t *testing.T) *auth.UseCasePackages {
	pkg := auth.UseCasePackages{
		Bcrypt:        pkg_mocks.NewMockBcryptHasher(t),
		Mailer:        pkg_mocks.NewMockResendClient(t),
		TokenProvider: pkg_mocks.NewMockJWTProvider(t),
		Redis:         pkg_mocks.NewMockRedisClient(t),
		Logger:        shared_testing.NewMockLogger(t),
	}

	return &pkg
}
