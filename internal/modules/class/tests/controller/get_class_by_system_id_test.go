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
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetClassBySystemID_ClassNotFound(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	_ = factory.NewClass(id, "class-A")

	expected := shared.ErrRecordNotFound
	expectedStatusCode := http.StatusNotFound

	mockUsecase := mocks.NewMockIUseCase(t)
	mockUsecase.EXPECT().GetClassBySystemID(ctx, mock.AnythingOfType("string")).Return(nil, shared.ErrRecordNotFound)

	controller := class.NewController(mockUsecase)

	router := gin.New()
	router.GET("/classes/system/:systemId", controller.GetClassBySystemID)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/classes/system/nonexistent-system-id", nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expectedStatusCode, w.Code)

	var response shared.ResponseError
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected.Error(), response.Error)
}

func TestGetClassBySystemID_Success(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")

	expected := existingClass
	expectedStatusCode := http.StatusOK

	mockUsecase := mocks.NewMockIUseCase(t)

	mockResult := existingClass
	mockUsecase.EXPECT().GetClassBySystemID(ctx, mock.AnythingOfType("string")).Return(mockResult, nil)

	controller := class.NewController(mockUsecase)

	router := gin.New()
	router.GET("/classes/system/:systemId", controller.GetClassBySystemID)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/classes/system/"+existingClass.SystemID, nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expectedStatusCode, w.Code)

	var response shared.ResponseSuccess[domain.Class]
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, response)
}
