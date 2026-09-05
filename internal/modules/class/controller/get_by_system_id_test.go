package controller_test

import (
	"context"
	"encoding/json"
	"main/internal/http/response"
	"main/internal/modules/class/controller"
	"main/internal/modules/class/controller/mocks"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/dto"
	"main/internal/modules/class/factory"
	"main/internal/shared/errors"
	"main/internal/testutil/logger"
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

func TestGetBySystemID_ClassNotFound(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	_ = factory.NewClass(id, "class-A")

	expected := domain.ErrClassNotFound
	expectedStatusCode := http.StatusNotFound

	mockUseCase := mocks.NewMockClassUseCaseI(t)
	mockLogger := logger.NewMockLogger(t)
	mockUseCase.EXPECT().GetBySystemID(ctx, mock.AnythingOfType("string")).Return(nil, domain.ErrClassNotFound)

	classController := controller.NewClassController(mockUseCase, mockLogger)

	router := gin.New()
	router.GET("/classes/system/:id", classController.GetBySystemID)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/classes/system/nonexistent-class-systemID", nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expectedStatusCode, w.Code)

	var response errors.AppError
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected.Message, response.Message)
}

func TestGetBySystemID_Success(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")

	expected := dto.DomainToResponse(*existingClass)
	expectedStatusCode := http.StatusOK

	mockUseCase := mocks.NewMockClassUseCaseI(t)
	mockLogger := logger.NewMockLogger(t)

	mockUseCase.EXPECT().GetBySystemID(ctx, mock.AnythingOfType("string")).Return(existingClass, nil)

	classController := controller.NewClassController(mockUseCase, mockLogger)

	router := gin.New()
	router.GET("/classes/system/:id", classController.GetBySystemID)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/classes/system/"+existingClass.SystemID, nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expectedStatusCode, w.Code)

	var response response.ResponseDTO[dto.ClassResponse]
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	expected.CreatedAt = time.Time{}
	expected.UpdatedAt = time.Time{}

	response.Data.CreatedAt = time.Time{}
	response.Data.UpdatedAt = time.Time{}

	assert.Equal(t, *expected, response.Data)
}
