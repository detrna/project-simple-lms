package auth_controller_test

import (
	"context"
	"encoding/json"
	"main/internal/modules/auth"
	auth_mocks "main/internal/modules/auth/mocks"
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
		RefreshToken: newRefreshToken.Value,
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

	req.AddCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: refreshToken.Value,
	})

	router := gin.New()
	c := auth.NewController(mockUsecase, mockLogger, false)

	router.POST("/refresh", func(ctx *gin.Context) {
		c.Refresh(ctx)
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	cookies := w.Result().Cookies()

	assert.Equal(t, 1, len(cookies))
	assert.Equal(t, expectedCookie, cookies[0].Value)

	var response shared.ResponseSuccess[auth.TokenResponse]
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expectedResult, *response.Data)
}

func TestRefresh_ExpiredToken(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingUser := user_factory.NewUser(id)
	jwtPayload := user_factory.NewJWTPayload(existingUser)
	refreshToken := user_factory.NewJWT("expired-refresh-token", jwtPayload)

	mockUsecase := auth_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)

	mockUsecase.
		EXPECT().
		Refresh(ctx, mock.AnythingOfType("string")).
		Return(nil, shared.ErrUnauthorized)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/refresh",
		nil,
	)

	req.AddCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: refreshToken.Value,
	})

	router := gin.New()
	c := auth.NewController(mockUsecase, mockLogger, false)

	router.POST("/refresh",
		func(ctx *gin.Context) {
			c.Refresh(ctx)
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

	req.AddCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: refreshToken.Value,
	})

	router := gin.New()
	c := auth.NewController(mockUsecase, mockLogger, false)

	router.POST("/refresh", func(ctx *gin.Context) {
		c.Refresh(ctx)
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response shared.ResponseError
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, shared.ErrUnauthorized.Error(), response.Error)
}
