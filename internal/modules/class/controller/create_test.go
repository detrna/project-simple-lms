package controller_test

import (
	"bytes"
	"context"
	"encoding/json"

	"main/internal/modules/class/controller"
	"main/internal/modules/class/controller/mocks"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/dto"
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

func TestCreate_SystemIDTaken(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")

	requestData := dto.CreateClassRequest{
		SystemID: existingClass.SystemID,
		Name:     existingClass.Name,
	}

	expected := shared.ErrSystemIDTaken.Error()

	expectedStatusCode := http.StatusConflict

	mockUseCase := mocks.NewMockClassUseCaseI(t)
	mockUseCase.EXPECT().Create(ctx, mock.AnythingOfType("*dto.CreateClassRequest")).Return(nil, shared.ErrSystemIDTaken)

	classController := controller.NewClassController(mockUseCase)

	router := gin.New()
	router.POST("/classes", classController.Create)

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

func TestCreate_Success(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	classSample := factory.NewClass(id, "class-A")

	requestData := dto.CreateClassRequest{
		SystemID: classSample.SystemID,
		Name:     classSample.Name,
	}

	expected := domain.Class{
		SystemID: requestData.SystemID,
		Name:     requestData.Name,
	}

	expectedStatusCode := http.StatusCreated

	mockUseCase := mocks.NewMockClassUseCaseI(t)

	mockResult := classSample
	mockResult.ID = uuid.New()

	mockUseCase.EXPECT().Create(ctx, mock.AnythingOfType("*dto.CreateClassRequest")).Return(mockResult, nil)

	classController := controller.NewClassController(mockUseCase)

	router := gin.New()
	router.POST("/classes", classController.Create)

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

