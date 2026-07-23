package user_controller_test

import (
	"context"
	"encoding/json"
	"main/internal/domain"
	"main/internal/modules/user"
	user_mocks "main/internal/modules/user/mocks"
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

func TestDeleteUser_Success(t *testing.T) {
	mockUsecase := user_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)
	ctrl := user.NewController(mockUsecase, mockLogger)

	id := uuid.New()
	account := domain.User{
		ID: id,
	}

	ctx := context.Background()
	mockUsecase.On("DeleteUser", ctx, mock.AnythingOfType("uuid.UUID")).Return(nil)

	router := gin.New()
	router.DELETE("/:id", func(c *gin.Context) {
		ctrl.DeleteUser(c)
	})

	w := httptest.NewRecorder()

	path := "/" + (account.ID).String()
	req := httptest.NewRequest(
		http.MethodDelete,
		path,
		nil,
	)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	mockUsecase.AssertExpectations(t)
}

func TestDeleteUser_RecordNotFound(t *testing.T) {
	mockUsecase := user_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)
	ctrl := user.NewController(mockUsecase, mockLogger)

	id := uuid.New()

	ctx := context.Background()
	mockUsecase.On("GetUserByID", ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, shared.ErrRecordNotFound)

	router := gin.New()
	router.DELETE("/:id", func(c *gin.Context) {
		ctrl.GetUserByID(c)
	})

	w := httptest.NewRecorder()

	path := "/" + (id).String()
	req := httptest.NewRequest(
		http.MethodDelete,
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
