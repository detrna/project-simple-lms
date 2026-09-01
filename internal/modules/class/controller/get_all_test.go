package controller_test

import (
	"context"
	"encoding/json"
	"main/internal/modules/class/controller"
	"main/internal/modules/class/controller/mocks"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/factory"
	"main/internal/shared"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAll_Success(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	existingClasses := []*domain.Class{
		factory.NewClass(id, "class-A"),
		factory.NewClass(id, "class-B"),
		factory.NewClass(id, "class-C"),
		factory.NewClass(id, "class-D"),
		factory.NewClass(id, "class-E"),
	}

	expected := existingClasses
	expectedStatusCode := http.StatusOK

	mockUseCase := mocks.NewMockClassUseCaseI(t)

	mockResult := existingClasses
	mockUseCase.EXPECT().GetAll(ctx).Return(mockResult, nil)

	classController := controller.NewClassController(mockUseCase)

	router := gin.New()
	router.GET("/classes", classController.GetAll)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/classes", nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expectedStatusCode, w.Code)

	var response shared.ResponseSuccess[[]domain.Class]
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, response)
}

