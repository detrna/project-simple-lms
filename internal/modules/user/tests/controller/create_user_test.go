package user_controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"main/internal/modules/user"
	user_mocks "main/internal/modules/user/mocks"
	"main/internal/shared"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateUser_Success(t *testing.T) {
	mockUsecase := user_mocks.NewMockIUseCase(t)
	mockLogger := NewMockLogger(t)
	ctrl := user.NewController(mockUsecase, mockLogger)

	id := uuid.New()
	userSample := NewUserSample(id)

	requestData := user.CreateUserSchema{
		SystemID: userSample.SystemID,
		Name:     userSample.Name,
		Role:     userSample.Role,
		Email:    userSample.Email,
		Password: userSample.Password,
	}

	mockResult := user.UserResponse{
		ID:       userSample.ID,
		SystemID: userSample.SystemID,
		Name:     userSample.Name,
		Email:    userSample.Email,
		Role:     userSample.Role,
	}

	expected := &mockResult

	ctx := context.Background()
	mockUsecase.On("CreateUser", ctx, mock.AnythingOfType("*user.CreateUserSchema")).Return(&mockResult, nil).Once()

	router := gin.New()
	router.POST("/", func(c *gin.Context) {
		ctrl.CreateUser(c)
	})

	w := httptest.NewRecorder()

	path := "/"
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
	assert.Equal(t, expected, response.Data)

	mockUsecase.AssertExpectations(t)
}

func TestCreateUser_EmailTaken(t *testing.T) {
	mockUsecase := user_mocks.NewMockIUseCase(t)
	mockLogger := NewMockLogger(t)
	ctrl := user.NewController(mockUsecase, mockLogger)

	id := uuid.New()
	userSample := NewUserSample(id)

	requestData := user.CreateUserSchema{
		SystemID: userSample.SystemID,
		Name:     userSample.Name,
		Role:     userSample.Role,
		Email:    userSample.Email,
		Password: userSample.Password,
	}

	ctx := context.Background()
	mockUsecase.On("CreateUser", ctx, mock.AnythingOfType("*user.CreateUserSchema")).Return(nil, shared.ErrEmailTaken)

	router := gin.New()
	router.POST("/", func(c *gin.Context) {
		ctrl.CreateUser(c)
	})

	w := httptest.NewRecorder()

	path := "/"
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

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, shared.ErrEmailTaken.Error(), response.Error)

	mockUsecase.AssertExpectations(t)
}

func TestCreateUser_SystemIDTaken(t *testing.T) {
	mockUsecase := user_mocks.NewMockIUseCase(t)
	mockLogger := NewMockLogger(t)
	ctrl := user.NewController(mockUsecase, mockLogger)

	id := uuid.New()
	userSample := NewUserSample(id)

	requestData := user.CreateUserSchema{
		SystemID: userSample.SystemID,
		Name:     userSample.Name,
		Role:     userSample.Role,
		Email:    userSample.Email,
		Password: userSample.Password,
	}

	ctx := context.Background()
	mockUsecase.On("CreateUser", ctx, mock.AnythingOfType("*user.CreateUserSchema")).Return(nil, shared.ErrSystemIDTaken)

	router := gin.New()
	router.POST("/", func(c *gin.Context) {
		ctrl.CreateUser(c)
	})

	w := httptest.NewRecorder()

	path := "/"
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

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, shared.ErrSystemIDTaken.Error(), response.Error)

	mockUsecase.AssertExpectations(t)
}
