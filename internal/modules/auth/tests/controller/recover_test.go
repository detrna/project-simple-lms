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

func TestRecover_Success(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingUser := user_factory.NewUser(id)

	mockUsecase := auth_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)

	mockUsecase.EXPECT().Recover(ctx, mock.AnythingOfType("*auth.RecoverSchema")).Return(nil)

	requestData := auth.RecoverSchema{
		Email: existingUser.Email,
	}

	requestBody, err := json.Marshal(requestData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/recover",
		bytes.NewReader(requestBody),
	)

	req.Header.Set("Content-Type", "application/json")

	router := gin.New()
	c := auth.NewController(mockUsecase, mockLogger, false)

	router.POST("/recover", func(ctx *gin.Context) {
		c.Recover(ctx)
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestRecover_EmailNotFound(t *testing.T) {
	ctx := context.Background()

	mockUsecase := auth_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)

	mockUsecase.
		EXPECT().
		Recover(ctx, mock.AnythingOfType("*auth.RecoverSchema")).
		Return(shared.ErrRecordNotFound)

	requestData := auth.RecoverSchema{
		Email: "incorrect-email",
	}

	requestBody, err := json.Marshal(requestData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/recover",
		bytes.NewReader(requestBody),
	)

	req.Header.Set("Content-Type", "application/json")

	router := gin.New()
	c := auth.NewController(mockUsecase, mockLogger, false)

	router.POST("/recover", func(ctx *gin.Context) {
		c.Recover(ctx)
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response shared.ResponseError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, shared.ErrRecordNotFound.Error(), response.Error)
}
