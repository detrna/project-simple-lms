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

func TestUpdate_ClassNotFound(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")

	newName := "class-B"
	requestData := dto.UpdateClassRequest{
		SystemID: &existingClass.SystemID,
		Name:     &newName,
	}

	expectedCode := http.StatusNotFound
	expected := shared.ErrRecordNotFound.Error()

	requestBody, err := json.Marshal(&requestData)
	require.NoError(t, err)

	mockUseCase := mocks.NewMockClassUseCaseI(t)
	mockUseCase.EXPECT().GetByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, shared.ErrRecordNotFound)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPatch, "/classes/"+existingClass.ID.String(), bytes.NewReader(requestBody))

	classController := controller.NewClassController(mockUseCase)

	router := gin.New()
	router.PATCH("/classes/:id", classController.Update)

	router.ServeHTTP(w, req)

	assert.Equal(t, expectedCode, w.Code)

	var response shared.ResponseError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, response.Error)
}

func TestUpdate_SystemIDTaken(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")
	otherClass := factory.NewClass(id, "class-C")

	newName := "class-B"
	requestData := dto.UpdateClassRequest{
		SystemID: &otherClass.SystemID,
		Name:     &newName,
	}

	expectedCode := http.StatusConflict
	expected := shared.ErrSystemIDTaken.Error()

	requestBody, err := json.Marshal(&requestData)
	require.NoError(t, err)

	mockUseCase := mocks.NewMockClassUseCaseI(t)
	mockUseCase.EXPECT().GetByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(existingClass, nil)
	mockUseCase.EXPECT().GetBySystemID(ctx, mock.AnythingOfType("string")).Return(otherClass, nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPatch, "/classes/"+existingClass.ID.String(), bytes.NewReader(requestBody))

	classController := controller.NewClassController(mockUseCase)

	router := gin.New()
	router.PATCH("/classes/:id", classController.Update)

	router.ServeHTTP(w, req)

	assert.Equal(t, expectedCode, w.Code)

	var response shared.ResponseError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, response.Error)
}

func TestUpdate_Success(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")

	newName := "class-B"
	requestData := dto.UpdateClassRequest{
		Name: &newName,
	}

	expectedCode := http.StatusOK
	expected := *existingClass
	expected.Name = newName

	requestBody, err := json.Marshal(&requestData)
	require.NoError(t, err)

	mockUseCase := mocks.NewMockClassUseCaseI(t)

	mockResult := *existingClass
	mockResult.Name = newName

	mockUseCase.EXPECT().GetByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(existingClass, nil)
	mockUseCase.EXPECT().GetBySystemID(ctx, mock.AnythingOfType("string")).Return(nil, shared.ErrRecordNotFound)
	mockUseCase.EXPECT().Update(ctx, mock.AnythingOfType("*dto.UpdateClassRequest")).Return(&mockResult, nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPatch, "/classes/"+existingClass.ID.String(), bytes.NewReader(requestBody))

	classController := controller.NewClassController(mockUseCase)

	router := gin.New()
	router.PATCH("/classes/:id", classController.Update)

	router.ServeHTTP(w, req)

	assert.Equal(t, expectedCode, w.Code)

	var response shared.ResponseSuccess[domain.Class]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, response.Data)
}

