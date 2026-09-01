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
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetBySystemID_ClassNotFound(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	_ = factory.NewClass(id, "class-A")

	expected := shared.ErrRecordNotFound
	expectedStatusCode := http.StatusNotFound

	mockUseCase := mocks.NewMockClassUseCaseI(t)
	mockUseCase.EXPECT().GetBySystemID(ctx, mock.AnythingOfType("string")).Return(nil, shared.ErrRecordNotFound)

	classController := controller.NewClassController(mockUseCase)

	router := gin.New()
	router.GET("/classes/system/:systemId", classController.GetBySystemID)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/classes/system/nonexistent-system-id", nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expectedStatusCode, w.Code)

	var response shared.ResponseError
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected.Error(), response.Error)
}

func TestGetBySystemID_Success(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")

	expected := existingClass
	expectedStatusCode := http.StatusOK

	mockUseCase := mocks.NewMockClassUseCaseI(t)

	mockResult := existingClass
	mockUseCase.EXPECT().GetBySystemID(ctx, mock.AnythingOfType("string")).Return(mockResult, nil)

	classController := controller.NewClassController(mockUseCase)

	router := gin.New()
	router.GET("/classes/system/:systemId", classController.GetBySystemID)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/classes/system/"+existingClass.SystemID, nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expectedStatusCode, w.Code)

	var response shared.ResponseSuccess[domain.Class]
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, response)
}

