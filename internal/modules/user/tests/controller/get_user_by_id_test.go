package user_controller_test

import (
	"context"
	"encoding/json"
	"main/internal/modules/user"
	user_mocks "main/internal/modules/user/mocks"
	user_factory "main/internal/modules/user/tests"
	pkg_mocks "main/internal/pkg/mocks"
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

func TestGetUserByID_Success(t *testing.T) {
	mockUsecase := user_mocks.NewMockIUseCase(t)
	ctrl := user.NewController(mockUsecase, &pkg_mocks.MockLogger{})

	id := uuid.New()
	userSample := *user_factory.NewUser(id)

	mockResult := user.UserResponse{
		ID:       userSample.ID,
		SystemID: userSample.SystemID,
		Name:     userSample.Name,
		Email:    userSample.Email,
		Role:     userSample.Role,
	}

	expected := &mockResult

	ctx := context.Background()
	mockUsecase.On("GetUserByID", ctx, mock.AnythingOfType("uuid.UUID")).Return(&mockResult, nil)

	router := gin.New()
	router.GET("/:id", func(c *gin.Context) {
		ctrl.GetUserByID(c)
	})

	w := httptest.NewRecorder()

	path := "/" + (userSample.ID).String()
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
	assert.Equal(t, expected, response.Data)

	mockUsecase.AssertExpectations(t)
}

func TestGetUserByID_RecordNotFound(t *testing.T) {
	mockUsecase := user_mocks.NewMockIUseCase(t)
	MockLogger := shared_testing.NewMockLogger(t)
	ctrl := user.NewController(mockUsecase, MockLogger)

	id := uuid.New()

	ctx := context.Background()
	mockUsecase.On("GetUserByID", ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, shared.ErrRecordNotFound)
	MockLogger.On("Warn", mock.Anything).Return()

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

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, shared.ErrRecordNotFound.Error(), response.Error)

	mockUsecase.AssertExpectations(t)
}
