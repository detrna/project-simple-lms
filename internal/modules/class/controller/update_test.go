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
	"main/internal/shared/errors"
	"main/internal/testutil/logger"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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
	expected := domain.ErrClassNotFound

	requestBody, err := json.Marshal(&requestData)
	require.NoError(t, err)

	mockUseCase := mocks.NewMockClassUseCaseI(t)
	mockLogger := logger.NewMockLogger(t)
	mockUseCase.EXPECT().Update(ctx, mock.AnythingOfType("*dto.UpdateClassRequest")).Return(nil, domain.ErrClassNotFound)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPatch, "/classes/"+uuid.NewString(), bytes.NewReader(requestBody))

	classController := controller.NewClassController(mockUseCase, mockLogger)

	router := gin.New()
	router.PATCH("/classes/:id", classController.Update)

	router.ServeHTTP(w, req)

	assert.Equal(t, expectedCode, w.Code)

	var response errors.AppError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected.Message, response.Message)
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
	expected := domain.ErrClassSystemIDTaken

	requestBody, err := json.Marshal(&requestData)
	require.NoError(t, err)

	mockUseCase := mocks.NewMockClassUseCaseI(t)
	mockLogger := logger.NewMockLogger(t)
	mockUseCase.EXPECT().Update(ctx, mock.AnythingOfType("*dto.UpdateClassRequest")).Return(nil, domain.ErrClassSystemIDTaken)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPatch, "/classes/"+existingClass.ID.String(), bytes.NewReader(requestBody))

	classController := controller.NewClassController(mockUseCase, mockLogger)

	router := gin.New()
	router.PATCH("/classes/:id", classController.Update)

	router.ServeHTTP(w, req)

	assert.Equal(t, expectedCode, w.Code)

	var response errors.AppError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected.Message, response.Message)
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
	expected := dto.DomainToResponse(*existingClass)
	expected.Name = newName

	requestBody, err := json.Marshal(&requestData)
	require.NoError(t, err)

	mockUseCase := mocks.NewMockClassUseCaseI(t)
	mockLogger := logger.NewMockLogger(t)

	mockResult := *existingClass
	mockResult.Name = newName

	mockUseCase.EXPECT().Update(ctx, mock.AnythingOfType("*dto.UpdateClassRequest")).Return(&mockResult, nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPatch, "/classes/"+existingClass.ID.String(), bytes.NewReader(requestBody))

	classController := controller.NewClassController(mockUseCase, mockLogger)

	router := gin.New()
	router.PATCH("/classes/:id", classController.Update)

	router.ServeHTTP(w, req)

	assert.Equal(t, expectedCode, w.Code)

	var response shared.ResponseDTO[dto.ClassResponse]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	expected.CreatedAt = time.Time{}
	expected.UpdatedAt = time.Time{}

	response.Data.CreatedAt = time.Time{}
	response.Data.UpdatedAt = time.Time{}

	assert.Equal(t, *expected, *response.Data)
}
