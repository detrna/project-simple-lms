package class_controller_test

import (
	"bytes"
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

func TestCreateClass_SystemIDTaken(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")

	requestData := class.CreateClassRequest{
		SystemID: existingClass.SystemID,
		Name:     existingClass.Name,
	}

	expected := shared.ErrSystemIDTaken.Error()

	expectedStatusCode := http.StatusConflict

	mockUsecase := mocks.NewMockIUseCase(t)
	mockUsecase.EXPECT().CreateClass(ctx, mock.AnythingOfType("*domain.CreateClassRequest")).Return(nil, shared.ErrSystemIDTaken)

	controller := class.NewController(mockUsecase)

	router := gin.New()
	router.POST("/classes", controller.CreateClass)

	requestBody, err := json.Marshal(&requestData)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/classes", bytes.NewReader(requestBody))

	router.ServeHTTP(w, request)

	assert.Equal(t, expectedStatusCode, w.Code)

	var response shared.ResponseError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, response.Error)
}

func TestCreateClass_Success(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	classSample := factory.NewClass(id, "class-A")

	requestData := class.CreateClassRequest{
		SystemID: classSample.SystemID,
		Name:     classSample.Name,
	}

	expected := domain.Class{
		SystemID: requestData.SystemID,
		Name:     requestData.Name,
	}

	expectedStatusCode := http.StatusCreated

	mockUsecase := mocks.NewMockIUseCase(t)

	mockResult := classSample
	mockResult.ID = uuid.New()

	mockUsecase.EXPECT().CreateClass(ctx, mock.AnythingOfType("*domain.CreateClassRequest")).Return(mockResult, nil)

	controller := class.NewController(mockUsecase)

	router := gin.New()
	router.POST("/classes", controller.CreateClass)

	requestBody, err := json.Marshal(&requestData)

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/classes", bytes.NewReader(requestBody))

	router.ServeHTTP(w, request)

	assert.Equal(t, expectedStatusCode, w.Code)

	var response shared.ResponseSuccess[domain.Class]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, response)
}
