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

func TestAdminUpdateUser_Success(t *testing.T) {
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

	newName := "user-test-2"
	requestData := user.UpdateUserBodySchema{
		Name: &newName,
	}

	mockResult := user.UserResponse{
		ID:        account.ID,
		SystemID:  account.SystemID,
		Name:      newName,
		Email:     account.Email,
		Role:      account.Role,
		CreatedAt: account.CreatedAt,
	}

	expected := &mockResult

	ctx := context.Background()
	mockUsecase.On("AdminUpdateUser", ctx, mock.AnythingOfType("user.UpdateUserSchema")).Return(&mockResult, nil)

	router := gin.New()
	router.PATCH("/:id/admin", func(c *gin.Context) {
		ctrl.UpdateUser(c)
	})

	w := httptest.NewRecorder()

	path := "/" + (account.ID).String() + "/admin"
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

func TestAdminUpdateUser_RecordNotFound(t *testing.T) {
	mockUsecase := &user_mocks.MockIUseCase{}
	ctrl := user.NewController(mockUsecase, &pkg_mocks.MockLogger{})

	id := uuid.New()

	account := domain.User{
		ID: id,
	}

	newName := "user-test-2"
	requestData := user.UpdateUserBodySchema{
		Name: &newName,
	}

	ctx := context.Background()
	mockUsecase.On("AdminUpdateUser", ctx, mock.AnythingOfType("user.UpdateUserSchema")).Return(nil, shared.ErrRecordNotFound)

	router := gin.New()
	router.PATCH("/:id/admin", func(c *gin.Context) {
		ctrl.UpdateUser(c)
	})

	w := httptest.NewRecorder()

	path := "/" + (account.ID).String() + "/admin"
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
	assert.Equal(t, shared.ErrRecordNotFound.Error(), response.Error)

	mockUsecase.AssertExpectations(t)
}

func TestAdminUpdateUser_EmailTaken(t *testing.T) {
	mockUsecase := &user_mocks.MockIUseCase{}
	ctrl := user.NewController(mockUsecase, &pkg_mocks.MockLogger{})

	newEmail := "user-test-updated@mail.com"
	requestData := user.UpdateUserBodySchema{
		Email: &newEmail,
	}

	ctx := context.Background()
	mockUsecase.On("AdminUpdateUser", ctx, mock.AnythingOfType("user.UpdateUserSchema")).Return(nil, shared.ErrEmailTaken)

	router := gin.New()
	router.PATCH("/:id/admin", func(c *gin.Context) {
		ctrl.UpdateUser(c)
	})

	w := httptest.NewRecorder()

	path := "/" + uuid.NewString() + "/admin"
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

func TestAdminUpdateUser_SystemIDTaken(t *testing.T) {
	mockUsecase := &user_mocks.MockIUseCase{}
	ctrl := user.NewController(mockUsecase, &pkg_mocks.MockLogger{})

	newSystemID := "user-test-2"
	requestData := user.UpdateUserBodySchema{
		SystemID: &newSystemID,
	}

	ctx := context.Background()
	mockUsecase.On("AdminUpdateUser", ctx, mock.AnythingOfType("user.UpdateUserSchema")).Return(nil, shared.ErrSystemIDTaken)

	router := gin.New()
	router.PATCH("/:id/admin", func(c *gin.Context) {
		ctrl.UpdateUser(c)
	})

	w := httptest.NewRecorder()

	path := "/" + uuid.NewString() + "/admin"
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
