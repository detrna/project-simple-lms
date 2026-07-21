package user_controller_test

import (
	"context"
	"encoding/json"
	"main/internal/domain"
	"main/internal/modules/user"
	user_mocks "main/internal/modules/user/mocks"
	pkg_mocks "main/internal/pkg/mocks"
	"main/internal/shared"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetUserByID_Success(t *testing.T) {
	mockUsecase := &user_mocks.MockIUseCase{}
	ctrl := user.NewController(mockUsecase, &pkg_mocks.MockLogger{})

	id := uuid.New()
	createdAt := time.Now()

	account := domain.User{
		ID:        id,
		SystemID:  "user-test-1",
		Name:      "user-test",
		Role:      "default",
		Email:     "user-test@mail.com",
		Password:  "password123",
		CreatedAt: createdAt,
	}

	mockResult := user.UserResponse{
		ID:       account.ID,
		SystemID: account.SystemID,
		Name:     account.Name,
		Email:    account.Email,
		Role:     account.Role,
	}

	expected := &mockResult

	ctx := context.Background()
	mockUsecase.On("GetUserByID", ctx, mock.AnythingOfType("uuid.UUID")).Return(&mockResult, nil)

	router := gin.New()
	router.GET("/:id", func(c *gin.Context) {
		ctrl.GetUserByID(c)
	})

	w := httptest.NewRecorder()

	path := "/" + (account.ID).String()
	req := httptest.NewRequest(
		http.MethodGet,
		path,
		nil,
	)

	router.ServeHTTP(w, req)

	var response shared.ResponseSuccess[user.UserResponse]
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, *expected, response.Data)

	mockUsecase.AssertExpectations(t)
}

func TestGetUserByID_RecordNotFound(t *testing.T) {
	mockUsecase := &user_mocks.MockIUseCase{}
	ctrl := user.NewController(mockUsecase, &pkg_mocks.MockLogger{})

	id := uuid.New()

	ctx := context.Background()
	mockUsecase.On("GetUserByID", ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, shared.ErrRecordNotFound)

	router := gin.New()
	router.GET("/:id", func(c *gin.Context) {
		ctrl.GetUserByID(c)
	})

	w := httptest.NewRecorder()

	path := "/" + (id).String()
	req := httptest.NewRequest(
		http.MethodGet,
		path,
		nil,
	)

	router.ServeHTTP(w, req)

	var response shared.ResponseError
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, shared.ErrBadRequest.Error(), response.Error)

	mockUsecase.AssertExpectations(t)
}
