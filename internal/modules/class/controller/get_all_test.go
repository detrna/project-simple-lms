package controller_test

import (
	"context"
	"encoding/json"
	"main/internal/http/response"
	"main/internal/middleware"
	"main/internal/modules/class/controller"
	"main/internal/modules/class/controller/mocks"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/dto"
	"main/internal/modules/class/factory"
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

func TestGetAll_Success(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	existingClasses := []domain.Class{
		*factory.NewClass(id, "class-A"),
		*factory.NewClass(id, "class-B"),
		*factory.NewClass(id, "class-C"),
		*factory.NewClass(id, "class-D"),
		*factory.NewClass(id, "class-E"),
	}

	expected := *dto.DomainToResponseBatch(existingClasses)
	expectedStatusCode := http.StatusOK

	mockUseCase := mocks.NewMockClassUseCaseI(t)
	mockLogger := logger.NewMockLogger(t)

	mockResult := existingClasses
	mockUseCase.EXPECT().GetAll(ctx, mock.AnythingOfType("*pagination.Pagination")).Return(&mockResult, len(existingClasses), nil)

	classController := controller.NewClassController(mockUseCase, mockLogger)

	router := gin.New()

	router.GET("/classes", middleware.HandlePagination(10, 50, mockLogger), classController.GetAll)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/classes", nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expectedStatusCode, w.Code)

	var response response.ResponseDTO[[]dto.ClassResponse]
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected[0].Name, response.Data[0].Name)
	assert.Equal(t, len(expected), len(response.Data))
}
