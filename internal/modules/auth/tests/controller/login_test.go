package auth_controller_test

import (
	"bytes"
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

func TestLogin_Success(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingUser := user_factory.NewUser(id)
	jwtPayload := user_factory.NewJWTPayload(existingUser)
	accessToken := user_factory.NewJWT("access-token", jwtPayload)
	refreshToken := user_factory.NewJWT("refresh-token", jwtPayload)

	mockUsecase := auth_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)

	usecaseResult := auth.Tokens{
		AccessToken:  accessToken.Value,
		RefreshToken: refreshToken.Value,
	}

	expectedResult := auth.TokenResponse{
		AccessToken: usecaseResult.AccessToken,
	}

	expectedCookie := usecaseResult.RefreshToken

	mockUsecase.EXPECT().Login(ctx, mock.AnythingOfType("*auth.LoginSchema")).Return(&usecaseResult, nil)

	requestData := auth.LoginSchema{
		Email:    existingUser.Email,
		Password: existingUser.Password,
	}

	requestBody, err := json.Marshal(requestData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/login",
		bytes.NewReader(requestBody),
	)

	req.Header.Set("Content-Type", "application/json")

	router := gin.New()
	c := auth.NewController(mockUsecase, mockLogger, false)

	router.POST("/login", func(ctx *gin.Context) {
		c.Login(ctx)
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	cookies := w.Result().Cookies()

	require.Len(t, cookies, 1)

	assert.Equal(t, "refresh_token", cookies[0].Name)
	assert.Equal(t, expectedCookie, cookies[0].Value)
	assert.Equal(t, "/", cookies[0].Path)
	assert.True(t, cookies[0].HttpOnly)
	assert.False(t, cookies[0].Secure)

	var response shared.ResponseSuccess[auth.TokenResponse]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, &expectedResult, response.Data)
}

func TestLogin_IncorrectEmail(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingUser := user_factory.NewUser(id)

	mockUsecase := auth_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)

	mockUsecase.
		EXPECT().
		Login(ctx, mock.AnythingOfType("*auth.LoginSchema")).
		Return(nil, shared.ErrCredentialsIncorrect)

	requestData := auth.LoginSchema{
		Email:    "incorrect email",
		Password: existingUser.Password,
	}

	requestBody, err := json.Marshal(requestData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/login",
		bytes.NewReader(requestBody),
	)

	req.Header.Set("Content-Type", "application/json")

	router := gin.New()
	c := auth.NewController(mockUsecase, mockLogger, false)

	router.POST("/login", func(ctx *gin.Context) {
		c.Login(ctx)
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response shared.ResponseError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, shared.ErrCredentialsIncorrect.Error(), response.Error)
}

func TestLogin_IncorrectPassword(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingUser := user_factory.NewUser(id)

	mockUsecase := auth_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)

	mockUsecase.
		EXPECT().
		Login(ctx, mock.AnythingOfType("*auth.LoginSchema")).
		Return(nil, shared.ErrCredentialsIncorrect)

	requestData := auth.LoginSchema{
		Email:    existingUser.Email,
		Password: "incorrect password",
	}

	requestBody, err := json.Marshal(requestData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/login",
		bytes.NewReader(requestBody),
	)

	req.Header.Set("Content-Type", "application/json")

	router := gin.New()
	c := auth.NewController(mockUsecase, mockLogger, false)

	router.POST("/login", func(ctx *gin.Context) {
		c.Login(ctx)

		_, err := ctx.Cookie("refresh_token")
		require.Error(t, err)
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response shared.ResponseError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, shared.ErrCredentialsIncorrect.Error(), response.Error)
}
