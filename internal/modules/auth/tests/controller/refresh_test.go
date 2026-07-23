package auth_controller_test

import (
	"context"
	"encoding/json"
	"main/internal/modules/auth"
	auth_mocks "main/internal/modules/auth/mocks"
	auth_test_helper "main/internal/modules/auth/tests"
	user_factory "main/internal/modules/user/tests"
	"main/internal/shared"
	shared_testing "main/internal/shared/testing_helper"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRefresh_Success(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingUser := user_factory.NewUser(id)
	jwtPayload := user_factory.NewJWTPayload(existingUser)

	accessToken := user_factory.NewJWT("access-token", jwtPayload)
	refreshToken := user_factory.NewJWT("refresh-token", jwtPayload)
	newRefreshToken := user_factory.NewJWT("new-refresh-token", jwtPayload)

	mockUsecase := auth_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)

	usecaseResult := auth.Tokens{
		AccessToken:  accessToken.Value,
		RefreshToken: refreshToken.Value,
	}

	expectedResult := auth.TokenResponse{
		AccessToken: accessToken.Value,
	}

	expectedCookie := newRefreshToken.Value

	mockUsecase.EXPECT().Refresh(ctx, mock.AnythingOfType("string")).Return(&usecaseResult, nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/refresh",
		nil,
	)

	router := gin.New()
	c := auth.NewController(mockUsecase, mockLogger, false)

	router.POST("/refresh",
		func(ctx *gin.Context) {
			auth_test_helper.SetRefreshTokenCookie(ctx, refreshToken.Value)
		},
		func(ctx *gin.Context) {
			c.Login(ctx)

			cookie, err := ctx.Cookie("refresh_token")
			require.NoError(t, err)

			assert.Equal(t, expectedCookie, cookie)
		})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response shared.ResponseSuccess[auth.TokenResponse]
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expectedResult, response.Data)
}

func TestRefresh_ExpiredToken(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingUser := user_factory.NewUser(id)
	jwtPayload := user_factory.NewJWTPayload(existingUser)
	refreshToken := user_factory.NewJWT("expired-refresh-token", jwtPayload)

	mockUsecase := auth_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)

	mockUsecase.EXPECT().Refresh(ctx, mock.AnythingOfType("string")).Return(nil, shared.ErrUnauthorized)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/refresh",
		nil,
	)

	router := gin.New()
	c := auth.NewController(mockUsecase, mockLogger, false)

	router.POST("/refresh",
		func(ctx *gin.Context) {
			auth_test_helper.SetRefreshTokenCookie(ctx, refreshToken.Value)
		},
		func(ctx *gin.Context) {
			c.Login(ctx)

			_, err := ctx.Cookie("refresh_token")
			require.Error(t, err)
		})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response shared.ResponseError
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, shared.ErrUnauthorized.Error(), response.Error)
}

func TestRefresh_RevokedToken(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingUser := user_factory.NewUser(id)
	jwtPayload := user_factory.NewJWTPayload(existingUser)
	refreshToken := user_factory.NewJWT("revoked-refresh-token", jwtPayload)

	mockUsecase := auth_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)

	mockUsecase.EXPECT().Refresh(ctx, mock.AnythingOfType("string")).Return(nil, shared.ErrUnauthorized)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/refresh",
		nil,
	)

	router := gin.New()
	c := auth.NewController(mockUsecase, mockLogger, false)

	router.POST("/refresh",
		func(ctx *gin.Context) {
			auth_test_helper.SetRefreshTokenCookie(ctx, refreshToken.Value)
		},
		func(ctx *gin.Context) {
			c.Login(ctx)

			_, err := ctx.Cookie("refresh_token")
			require.Error(t, err)
		})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response shared.ResponseError
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, shared.ErrUnauthorized.Error(), response.Error)
}
