package auth_usecase_test

import (
	"context"
	"main/internal/modules/auth"
	pkg_mocks "main/internal/pkg/mocks"
	"main/internal/shared"

	auth_mocks "main/internal/modules/auth/mocks"
	user_mocks "main/internal/modules/user/mocks"
	user_factory "main/internal/modules/user/tests"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLogin_Success(t *testing.T) {
	id := uuid.New()
	existingUser := user_factory.NewUser(id)

	dbUser := *existingUser
	dbUser.Password = "hashed-password"

	requestData := auth.LoginSchema{
		Email:    existingUser.Email,
		Password: existingUser.Password,
	}

	expected := auth.Tokens{
		AccessToken:  "access-token",
		RefreshToken: "refresh-token",
	}

	ctx := context.Background()
	hashedRefreshToken := "hashed-refresh-token"

	repo := auth_mocks.NewMockIRepository(t)
	repo.
		EXPECT().
		CreateJWT(ctx, mock.AnythingOfType("*domain.JWTPayload"), mock.AnythingOfType("string")).
		Return(&hashedRefreshToken, nil)

	userRepo := user_mocks.NewMockIRepository(t)
	userRepo.
		EXPECT().
		FindByEmail(ctx, mock.AnythingOfType("string")).
		Return(&dbUser, nil)

	jwtPayload := user_factory.NewJWTPayload(existingUser)
	accessToken := user_factory.NewJWT("access-token", jwtPayload)
	refreshToken := user_factory.NewJWT("refresh-token", jwtPayload)

	bcrypt := pkg_mocks.NewMockBcryptHasher(t)
	bcrypt.
		EXPECT().
		CompareHashAndPassword(mock.Anything, mock.Anything).
		Return(nil)

	jwt := pkg_mocks.NewMockJWTProvider(t)
	jwt.
		EXPECT().
		GenerateAccessToken(mock.Anything).
		Return(accessToken, nil)
	jwt.
		EXPECT().
		GenerateRefreshToken(mock.Anything).
		Return(refreshToken, nil)

	pkg := auth.UseCasePackages{
		Bcrypt:        bcrypt,
		TokenProvider: jwt,
	}

	u := auth.NewUseCase(repo, userRepo, &pkg)

	result, err := u.Login(ctx, &requestData)

	require.NoError(t, err)
	assert.Equal(t, expected, *result)
}

func TestLogin_IncorrectEmail(t *testing.T) {
	id := uuid.New()
	existingUser := user_factory.NewUser(id)

	requestData := auth.LoginSchema{
		Email:    "incorrect-email",
		Password: existingUser.Password,
	}

	ctx := context.Background()

	userRepo := user_mocks.NewMockIRepository(t)
	userRepo.
		EXPECT().
		FindByEmail(ctx, mock.AnythingOfType("string")).
		Return(nil, shared.ErrRecordNotFound)

	u := auth.NewUseCase(auth_mocks.NewMockIRepository(t), userRepo, &auth.UseCasePackages{})

	result, err := u.Login(ctx, &requestData)

	require.ErrorIs(t, shared.ErrRecordNotFound, err)
	assert.Nil(t, result)
}

func TestLogin_IncorrectPassword(t *testing.T) {
	id := uuid.New()
	existingUser := user_factory.NewUser(id)

	dbUser := *existingUser
	dbUser.Password = "hashed-password"

	requestData := auth.LoginSchema{
		Email:    existingUser.Email,
		Password: "incorrect-password",
	}

	ctx := context.Background()
	repo := auth_mocks.NewMockIRepository(t)

	userRepo := user_mocks.NewMockIRepository(t)
	userRepo.
		EXPECT().
		FindByEmail(ctx, mock.AnythingOfType("string")).
		Return(&dbUser, nil)

	bcryptHasher := pkg_mocks.NewMockBcryptHasher(t)
	bcryptHasher.
		EXPECT().
		CompareHashAndPassword(mock.Anything, mock.Anything).
		Return(shared.ErrCredentialsIncorrect)

	pkg := auth.UseCasePackages{
		Bcrypt: bcryptHasher,
	}

	u := auth.NewUseCase(repo, userRepo, &pkg)

	result, err := u.Login(ctx, &requestData)

	require.ErrorIs(t, shared.ErrCredentialsIncorrect, err)
	assert.Nil(t, result)
}
