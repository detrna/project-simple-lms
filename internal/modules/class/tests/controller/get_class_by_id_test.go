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

func TestClassByID_ClassNotFound(t *testing.T) {
	ctx := context.Background()

	expected := shared.ErrRecordNotFound
	expectedStatusCode := http.StatusNotFound

	mockUsecase := mocks.NewMockIUseCase(t)
	mockUsecase.EXPECT().GetClassByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, shared.ErrRecordNotFound)

	controller := class.NewController(mockUsecase)

	router := gin.New()
	router.GET("/classes/:id", controller.GetClassByID)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/classes/nonexistent-id", nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expectedStatusCode, w.Code)

	var response shared.ResponseError
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, response.Error)
}

func TestClassByID_Success(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")

	expected := existingClass
	expectedStatusCode := http.StatusOK

	mockUsecase := mocks.NewMockIUseCase(t)

	mockResult := existingClass
	mockUsecase.EXPECT().GetClassByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(mockResult, nil)

	controller := class.NewController(mockUsecase)

	router := gin.New()
	router.GET("/classes/:id", controller.GetClassByID)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/classes/"+existingClass.ID.String(), nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, expectedStatusCode, w.Code)

	var response shared.ResponseSuccess[domain.Class]
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, response)
}
