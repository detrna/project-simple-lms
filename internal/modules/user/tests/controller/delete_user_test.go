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

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDeleteUser_Success(t *testing.T) {
	mockUsecase := &user_mocks.MockIUseCase{}
	ctrl := user.NewController(mockUsecase, &pkg_mocks.MockLogger{})

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

	assert.Equal(t, http.StatusOK, w.Code)

	mockUsecase.AssertExpectations(t)
}

func DeleteUser_RecordNotFound(t *testing.T) {
	mockUsecase := &user_mocks.MockIUseCase{}
	ctrl := user.NewController(mockUsecase, &pkg_mocks.MockLogger{})

	id := uuid.New()

	ctx := context.Background()
	mockUsecase.On("DeleteUser", ctx, mock.AnythingOfType("uuid.UUID")).Return(shared.ErrRecordNotFound)

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

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, shared.ErrBadRequest.Error(), response.Error)

	mockUsecase.AssertExpectations(t)
}
