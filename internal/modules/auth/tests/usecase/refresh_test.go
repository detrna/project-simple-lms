package auth_usecase_test

import (
	"context"
	"main/internal/config"
	"main/internal/modules/auth"
	auth_mocks "main/internal/modules/auth/mocks"
	user_mocks "main/internal/modules/user/mocks"
	user_factory "main/internal/modules/user/tests"
	pkg_mocks "main/internal/pkg/mocks"
	"main/internal/shared"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// (usecase *UseCase) Refresh(ctx context.Context, JWTPayload domain.JWTPayload) (*Tokens, error) {

func TestRefresh_Success(t *testing.T) {
	existingUser := user_factory.NewUser(uuid.New())
	jwtPayload := user_factory.NewJWTPayload(existingUser)
	refreshToken := user_factory.NewJWT("refresh-token", jwtPayload)

	dbToken := "hashed-refresh-token"
	newDbToken := "new-hashed-refresh-token"

	ctx := context.Background()
	repo := auth_mocks.NewMockIRepository(t)
	bcryptHasher := pkg_mocks.NewMockHasher(t)
	jwt := pkg_mocks.NewMockTokenService(t)

	repo.EXPECT().FindJWT(ctx, mock.Anything).Return(&dbToken, nil)
	repo.EXPECT().DeleteJWT(ctx, mock.Anything).Return(nil)
	repo.EXPECT().CreateJWT(ctx, mock.Anything, mock.Anything).Return(&newDbToken, nil)

	bcryptHasher.EXPECT().Compare(mock.Anything, mock.Anything).Return(nil)

	jwt.
		EXPECT().
		GenerateAccessToken(mock.Anything).
		Return(user_factory.NewJWT("new-access-token", jwtPayload), nil)
	jwt.
		EXPECT().
		GenerateRefreshToken(mock.Anything).
		Return(user_factory.NewJWT("new-refresh-token", jwtPayload), nil)
	jwt.EXPECT().
		ParseRefreshToken(mock.AnythingOfType("string")).
		Return(jwtPayload, nil)
	jwt.EXPECT().HashToken(mock.AnythingOfType("string")).Return("hashed-refresh-token")

	pkg := auth.UseCasePackages{
		Hasher:       bcryptHasher,
		TokenService: jwt,
	}

	u := auth.NewUseCase(repo, user_mocks.NewMockIRepository(t), &pkg, &config.MailConfig{})

	expected := auth.Tokens{
		AccessToken:  "new-access-token",
		RefreshToken: "new-refresh-token",
	}

	result, err := u.Refresh(ctx, refreshToken.Value)

	require.NoError(t, err)
	assert.Equal(t, expected, *result)
}

func TestRefresh_RevokedToken(t *testing.T) {
	existingUser := user_factory.NewUser(uuid.New())
	jwtPayload := user_factory.NewJWTPayload(existingUser)
	refreshToken := user_factory.NewJWT("refresh_token", jwtPayload)

	ctx := context.Background()
	repo := auth_mocks.NewMockIRepository(t)
	repo.EXPECT().FindJWT(ctx, mock.Anything).Return(nil, shared.ErrRecordNotFound)

	tokenService := pkg_mocks.NewMockTokenService(t)
	tokenService.
		EXPECT().
		ParseRefreshToken(mock.AnythingOfType("string")).
		Return(jwtPayload, nil)

	pkg := auth.UseCasePackages{
		TokenService: tokenService,
	}

	u := auth.NewUseCase(repo, user_mocks.NewMockIRepository(t), &pkg, &config.MailConfig{})

	result, err := u.Refresh(ctx, refreshToken.Value)

	require.ErrorIs(t, shared.ErrRecordNotFound, err)
	assert.Nil(t, result)
}

func TestRefresh_InvalidToken(t *testing.T) {
	existingUser := user_factory.NewUser(uuid.New())
	jwtPayload := user_factory.NewJWTPayload(existingUser)
	refreshToken := user_factory.NewJWT("refresh-token", jwtPayload)

	dbToken := "hashed-refresh-token"

	ctx := context.Background()
	repo := auth_mocks.NewMockIRepository(t)
	repo.EXPECT().FindJWT(ctx, mock.Anything).Return(&dbToken, nil)

	bcryptHasher := pkg_mocks.NewMockHasher(t)
	bcryptHasher.EXPECT().Compare(mock.Anything, mock.Anything).Return(shared.ErrCredentialsIncorrect)

	tokenService := pkg_mocks.NewMockTokenService(t)
	tokenService.
		EXPECT().
		ParseRefreshToken(mock.AnythingOfType("string")).
		Return(jwtPayload, nil)

	pkg := auth.UseCasePackages{
		Hasher:       bcryptHasher,
		TokenService: tokenService,
	}

	u := auth.NewUseCase(repo, user_mocks.NewMockIRepository(t), &pkg, &config.MailConfig{})

	result, err := u.Refresh(ctx, refreshToken.Value)

	require.ErrorIs(t, err, shared.ErrUnauthorized)
	assert.Nil(t, result)
}
