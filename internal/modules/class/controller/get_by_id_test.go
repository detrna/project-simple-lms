package controller_test

import (
	"context"
	"encoding/json"
	"time"

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

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestClassByID_ClassNotFound(t *testing.T) {
	ctx := context.Background()

	expected := domain.ErrClassNotFound
	expectedStatusCode := http.StatusNotFound

	mockUseCase := mocks.NewMockClassUseCaseI(t)
	mockLogger := logger.NewMockLogger(t)
	mockUseCase.EXPECT().GetByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, domain.ErrClassNotFound)

	classController := controller.NewClassController(mockUseCase, mockLogger)

	router := gin.New()
	router.GET("/classes/:id", classController.GetByID)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/classes/"+uuid.NewString(), nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expectedStatusCode, w.Code)

	var response errors.AppError
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected.Message, response.Message)
}

func TestClassByID_Success(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")

	expected := dto.DomainToResponse(*existingClass)
	expectedStatusCode := http.StatusOK

	mockUseCase := mocks.NewMockClassUseCaseI(t)
	mockLogger := logger.NewMockLogger(t)

	mockUseCase.EXPECT().GetByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(existingClass, nil)

	classController := controller.NewClassController(mockUseCase, mockLogger)

	router := gin.New()
	router.GET("/classes/:id", classController.GetByID)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/classes/"+existingClass.ID.String(), nil)

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
