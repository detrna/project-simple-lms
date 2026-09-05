package controller_test

import (
	"context"
	"encoding/json"
	"main/internal/modules/class/controller"
	"main/internal/modules/class/controller/mocks"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/factory"
	"main/internal/shared/errors"
	"main/internal/testutil/logger"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDelete_ClassNotFound(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	_ = factory.NewClass(id, "class-A")

	expected := domain.ErrClassNotFound
	expectedStatusCode := http.StatusNotFound

	mockUseCase := mocks.NewMockClassUseCaseI(t)
	mockLogger := logger.NewMockLogger(t)
	mockUseCase.EXPECT().Delete(ctx, mock.AnythingOfType("uuid.UUID")).Return(domain.ErrClassNotFound)

	classController := controller.NewClassController(mockUseCase, mockLogger)

	router := gin.New()
	router.DELETE("/classes/:id", classController.Delete)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "/classes/"+id.String(), nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expectedStatusCode, w.Code)

	var response errors.AppError
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected.Message, response.Message)
}

func TestDelete_Success(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	_ = factory.NewClass(id, "class-A")

	expected := http.StatusNoContent

	mockUseCase := mocks.NewMockClassUseCaseI(t)
	mockLogger := logger.NewMockLogger(t)
	mockUseCase.EXPECT().Delete(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil)

	classController := controller.NewClassController(mockUseCase, mockLogger)

	router := gin.New()
	router.DELETE("/classes/:id", classController.Delete)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "/classes/"+id.String(), nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expected, w.Code)
}
