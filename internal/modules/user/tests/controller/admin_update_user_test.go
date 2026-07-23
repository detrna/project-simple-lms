package user_controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"main/internal/domain"
	"main/internal/modules/user"
	user_mocks "main/internal/modules/user/mocks"
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

func TestAdminUpdateUser_Success(t *testing.T) {
	mockUsecase := user_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)
	ctrl := user.NewController(mockUsecase, mockLogger)

	id := uuid.New()
	existingUser := user_factory.NewUser(id)

	newName := "user-test-2"
	requestData := user.AdminUpdateUserSchema{
		Name: &newName,
	}

	mockResult := user.UserResponse{
		ID:       existingUser.ID,
		SystemID: existingUser.SystemID,
		Name:     newName,
		Email:    existingUser.Email,
		Role:     existingUser.Role,
	}

	expected := &mockResult

	ctx := context.Background()
	mockUsecase.On("AdminUpdateUser", ctx, mock.AnythingOfType("*user.AdminUpdateUserDTO")).Return(&mockResult, nil)

	router := gin.New()
	router.PATCH("/:id/admin", func(c *gin.Context) {
		ctrl.AdminUpdateUser(c)
	})

	w := httptest.NewRecorder()

	path := "/" + (existingUser.ID).String() + "/admin"
	body, err := json.Marshal(requestData)
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPatch,
		path,
		bytes.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var response shared.ResponseSuccess[user.UserResponse]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, response.Data)

	mockUsecase.AssertExpectations(t)
}

func TestAdminUpdateUser_RecordNotFound(t *testing.T) {
	mockUsecase := user_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)
	ctrl := user.NewController(mockUsecase, mockLogger)

	id := uuid.New()
	account := domain.User{
		ID: id,
	}

	newName := "user-test-2"
	requestData := user.AdminUpdateUserSchema{
		Name: &newName,
	}

	ctx := context.Background()
	mockUsecase.On("AdminUpdateUser", ctx, mock.AnythingOfType("*user.AdminUpdateUserDTO")).Return(nil, shared.ErrRecordNotFound)

	router := gin.New()
	router.PATCH("/:id/admin", func(c *gin.Context) {
		ctrl.AdminUpdateUser(c)
	})

	w := httptest.NewRecorder()

	path := "/" + (account.ID).String() + "/admin"
	body, err := json.Marshal(requestData)
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPatch,
		path,
		bytes.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	var response shared.ResponseError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, shared.ErrRecordNotFound.Error(), response.Error)

	mockUsecase.AssertExpectations(t)
}

func TestAdminUpdateUser_EmailTaken(t *testing.T) {
	mockUsecase := user_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)
	ctrl := user.NewController(mockUsecase, mockLogger)

	newEmail := "user-test-updated@mail.com"
	requestData := user.AdminUpdateUserSchema{
		Email: &newEmail,
	}

	ctx := context.Background()
	mockUsecase.On("AdminUpdateUser", ctx, mock.AnythingOfType("*user.AdminUpdateUserDTO")).Return(nil, shared.ErrEmailTaken)

	router := gin.New()
	router.PATCH("/:id/admin", func(c *gin.Context) {
		ctrl.AdminUpdateUser(c)
	})

	w := httptest.NewRecorder()

	path := "/" + uuid.NewString() + "/admin"
	body, err := json.Marshal(requestData)
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPatch,
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

func TestAdminUpdateUser_SystemIDTaken(t *testing.T) {
	mockUsecase := user_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)
	ctrl := user.NewController(mockUsecase, mockLogger)

	newSystemID := "user-test-2"
	requestData := user.AdminUpdateUserSchema{
		SystemID: &newSystemID,
	}

	ctx := context.Background()
	mockUsecase.On("AdminUpdateUser", ctx, mock.AnythingOfType("*user.AdminUpdateUserDTO")).Return(nil, shared.ErrSystemIDTaken)

	router := gin.New()
	router.PATCH("/:id/admin", func(c *gin.Context) {
		ctrl.AdminUpdateUser(c)
	})

	w := httptest.NewRecorder()

	path := "/" + uuid.NewString() + "/admin"
	body, err := json.Marshal(requestData)
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPatch,
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
