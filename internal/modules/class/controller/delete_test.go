package controller_test

import (
	"context"
	"encoding/json"
	"main/internal/modules/class/controller"
	"main/internal/modules/class/controller/mocks"
	"main/internal/modules/class/factory"
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

func TestDelete_ClassNotFound(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	_ = factory.NewClass(id, "class-A")

	expected := shared.ResponseError{
		Error: shared.ErrBadRequest.Error(),
	}

	expectedStatusCode := http.StatusNotFound

	mockUseCase := mocks.NewMockClassUseCaseI(t)
	mockUseCase.EXPECT().Delete(ctx, mock.AnythingOfType("uuid.UUID")).Return(shared.ErrRecordNotFound)

	classController := controller.NewClassController(mockUseCase)

	router := gin.New()
	router.DELETE("/classes/:classId", classController.Delete)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "/classes/"+id.String(), nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expectedStatusCode, w.Code)

	var response shared.ResponseError
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, response)
}

func TestDelete_Success(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	_ = factory.NewClass(id, "class-A")

	expected := http.StatusNoContent

	mockUseCase := mocks.NewMockClassUseCaseI(t)
	mockUseCase.EXPECT().Delete(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil)

	classController := controller.NewClassController(mockUseCase)

	router := gin.New()
	router.DELETE("/classes/:classId", classController.Delete)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "/classes/"+id.String(), nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expected, w.Code)
}

