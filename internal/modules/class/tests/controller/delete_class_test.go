package class_controller_test

import (
	"context"
	"encoding/json"
	"main/internal/modules/class"
	"main/internal/modules/class/tests/factory"
	"main/internal/modules/class/tests/mocks"
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

func TestDeleteClass_ClassNotFound(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	_ = factory.NewClass(id, "class-A")

	expected := shared.ResponseError{
		Error: shared.ErrBadRequest.Error(),
	}

	expectedStatusCode := http.StatusNotFound

	mockUsecase := mocks.NewMockIUseCase(t)
	mockUsecase.EXPECT().DeleteClass(ctx, mock.AnythingOfType("uuid.UUID")).Return(shared.ErrRecordNotFound)

	controller := class.NewController(mockUsecase)

	router := gin.New()
	router.DELETE("/classes/:classId", controller.DeleteClass)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "/classes/"+id.String(), nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expectedStatusCode, w.Code)

	var response shared.ResponseError
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, response)
}

func TestDeleteClass_Success(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	_ = factory.NewClass(id, "class-A")

	expected := http.StatusNoContent

	mockUsecase := mocks.NewMockIUseCase(t)
	mockUsecase.EXPECT().DeleteClass(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil)

	controller := class.NewController(mockUsecase)

	router := gin.New()
	router.DELETE("/classes/:classId", controller.DeleteClass)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "/classes/"+id.String(), nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expected, w.Code)
}
