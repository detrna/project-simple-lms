package auth_usecase_test

import (
	"context"
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
	dbToken := "hashed-refresh-token"
	newDbToken := "new-hashed-refresh-token"

	ctx := context.Background()
	repo := auth_mocks.NewMockIRepository(t)
	repo.EXPECT().FindJWT(ctx, mock.Anything).Return(&dbToken, nil)
	repo.EXPECT().DeleteJWT(ctx, mock.Anything).Return(nil)
	repo.EXPECT().CreateJWT(ctx, mock.Anything, mock.Anything).Return(&newDbToken, nil)

	bcryptHasher := pkg_mocks.NewMockBcryptHasher(t)
	bcryptHasher.EXPECT().CompareHashAndPassword(mock.Anything, mock.Anything).Return(nil)

	jwt := pkg_mocks.NewMockJWTProvider(t)
	jwt.
		EXPECT().
		GenerateAccessToken(mock.Anything).
		Return(user_factory.NewJWT("new-access-token", jwtPayload), nil)
	jwt.
		EXPECT().
		GenerateRefreshToken(mock.Anything).
		Return(user_factory.NewJWT("new-refresh-token", jwtPayload), nil)

	pkg := auth.UseCasePackages{
		Bcrypt:        bcryptHasher,
		TokenProvider: jwt,
	}

	u := auth.NewUseCase(repo, user_mocks.NewMockIRepository(t), &pkg)

	requestData := *jwtPayload
	expected := auth.Tokens{
		AccessToken:  "new-access-token",
		RefreshToken: "new-refresh-token",
	}

	result, err := u.Refresh(ctx, &requestData)

	require.NoError(t, err)
	assert.Equal(t, expected, *result)
}

func TestRefresh_RevokedToken(t *testing.T) {
	existingUser := user_factory.NewUser(uuid.New())
	jwtPayload := user_factory.NewJWTPayload(existingUser)

	ctx := context.Background()
	repo := auth_mocks.NewMockIRepository(t)
	repo.EXPECT().FindJWT(ctx, mock.Anything).Return(nil, shared.ErrRecordNotFound)

	pkg := auth.UseCasePackages{}

	u := auth.NewUseCase(repo, user_mocks.NewMockIRepository(t), &pkg)

	requestData := *jwtPayload

	result, err := u.Refresh(ctx, &requestData)

	require.ErrorIs(t, shared.ErrRecordNotFound, err)
	assert.Nil(t, result)
}

func TestRefresh_InvalidToken(t *testing.T) {
	existingUser := user_factory.NewUser(uuid.New())
	jwtPayload := user_factory.NewJWTPayload(existingUser)
	dbToken := "hashed-refresh-token"

	ctx := context.Background()
	repo := auth_mocks.NewMockIRepository(t)
	repo.EXPECT().FindJWT(ctx, mock.Anything).Return(&dbToken, nil)

	bcryptHasher := pkg_mocks.NewMockBcryptHasher(t)
	bcryptHasher.EXPECT().CompareHashAndPassword(mock.Anything, mock.Anything).Return(shared.ErrCredentialsIncorrect)

	pkg := auth.UseCasePackages{
		Bcrypt: bcryptHasher,
	}

	u := auth.NewUseCase(repo, user_mocks.NewMockIRepository(t), &pkg)

	requestData := *jwtPayload

	result, err := u.Refresh(ctx, &requestData)

	require.ErrorIs(t, shared.ErrUnauthorized, err)
	assert.Nil(t, result)
}
