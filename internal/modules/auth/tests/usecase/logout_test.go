package auth_usecase_test

import (
	"context"
	"main/internal/modules/auth"
	auth_mocks "main/internal/modules/auth/mocks"
	user_mocks "main/internal/modules/user/mocks"
	user_factory "main/internal/modules/user/tests"
	pkg_mocks "main/internal/pkg/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLogout_Success(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	existingUser := user_factory.NewUser(id)
	jwtPayload := user_factory.NewJWTPayload(existingUser)
	refreshToken := user_factory.NewJWT("refresh-token", jwtPayload)

	repo := auth_mocks.NewMockIRepository(t)

	dbToken := "hashed-refresh-token"
	repo.EXPECT().
		FindJWT(ctx, mock.AnythingOfType("uuid.UUID")).
		Return(&dbToken, nil)
	repo.EXPECT().
		DeleteJWT(ctx, mock.AnythingOfType("uuid.UUID")).
		Return(nil)

	bcryptHasher := pkg_mocks.NewMockBcryptHasher(t)
	bcryptHasher.
		EXPECT().
		Compare(mock.Anything, mock.Anything).
		Return(nil)

	tokenProvider := pkg_mocks.NewMockJWTProvider(t)
	tokenProvider.
		EXPECT().
		ParseRefreshToken(mock.AnythingOfType("string")).
		Return(jwtPayload, nil)

	pkg := auth.UseCasePackages{
		Bcrypt:        bcryptHasher,
		TokenProvider: tokenProvider,
	}

	u := auth.NewUseCase(repo, user_mocks.NewMockIRepository(t), &pkg, nil)

	err := u.Logout(ctx, refreshToken.Value)

	require.NoError(t, err)
}
