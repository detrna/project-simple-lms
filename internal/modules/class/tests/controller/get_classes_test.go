package class_controller_test

import (
	"context"
	"encoding/json"
	"main/internal/modules/class"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/tests/factory"
	"main/internal/modules/class/tests/mocks"
	"main/internal/shared"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetClasses_Success(t *testing.T) {
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

	mockUsecase := mocks.NewMockIUseCase(t)

	mockResult := existingClasses
	mockUsecase.EXPECT().GetClasses(ctx).Return(mockResult, nil)

	controller := class.NewController(mockUsecase)

	router := gin.New()
	router.GET("/classes/system/:systemId", controller.GetClassBySystemID)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/classes", nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expectedStatusCode, w.Code)

	var response shared.ResponseSuccess[[]domain.Class]
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, response)
}
