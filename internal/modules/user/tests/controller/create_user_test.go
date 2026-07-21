package user_controller_test

import (
	"bytes"
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

func TestCreateUser_Success(t *testing.T) {
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

	requestData := user.CreateUserSchema{
		SystemID: account.SystemID,
		Name:     account.Name,
		Role:     account.Role,
		Email:    account.Email,
		Password: account.Password,
	}

	mockResult := user.UserResponse{
		ID:        account.ID,
		SystemID:  account.SystemID,
		Name:      account.Name,
		Email:     account.Email,
		Role:      account.Role,
		CreatedAt: account.CreatedAt,
	}

	expected := &mockResult

	ctx := context.Background()
	mockUsecase.On("CreateUser", ctx, mock.AnythingOfType("user.CreateUserSchema")).Return(&mockResult, nil)

	router := gin.New()
	router.POST("", func(c *gin.Context) {
		ctrl.CreateUser(c)
	})

	w := httptest.NewRecorder()

	path := ""
	body, err := json.Marshal(requestData)
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPost,
		path,
		bytes.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	var response shared.ResponseSuccess[user.UserResponse]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, *expected, response.Data)

	mockUsecase.AssertExpectations(t)
}

func TestCreateUser_EmailTaken(t *testing.T) {
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

	requestData := user.CreateUserSchema{
		SystemID: account.SystemID,
		Name:     account.Name,
		Role:     account.Role,
		Email:    account.Email,
		Password: account.Password,
	}

	ctx := context.Background()
	mockUsecase.On("CreateUser", ctx, mock.AnythingOfType("user.CreateUserSchema")).Return(nil, shared.ErrEmailTaken)

	router := gin.New()
	router.POST("", func(c *gin.Context) {
		ctrl.CreateUser(c)
	})

	w := httptest.NewRecorder()

	path := ""
	body, err := json.Marshal(requestData)
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPost,
		path,
		bytes.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	var response shared.ResponseError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, shared.ErrEmailTaken.Error(), response.Error)

	mockUsecase.AssertExpectations(t)
}

func TestCreateUser_SystemIDTaken(t *testing.T) {
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

	requestData := user.CreateUserSchema{
		SystemID: account.SystemID,
		Name:     account.Name,
		Role:     account.Role,
		Email:    account.Email,
		Password: account.Password,
	}

	ctx := context.Background()
	mockUsecase.On("CreateUser", ctx, mock.AnythingOfType("user.CreateUserSchema")).Return(nil, shared.ErrSystemIDTaken)

	router := gin.New()
	router.POST("", func(c *gin.Context) {
		ctrl.CreateUser(c)
	})

	w := httptest.NewRecorder()

	path := ""
	body, err := json.Marshal(requestData)
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPost,
		path,
		bytes.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	var response shared.ResponseError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, shared.ErrSystemIDTaken.Error(), response.Error)

	mockUsecase.AssertExpectations(t)
}
