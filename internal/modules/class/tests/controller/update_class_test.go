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

func TestUpdateClass_ClassNotFound(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")

	newName := "class-B"
	requestData := class.UpdateClassRequest{
		SystemID: &existingClass.SystemID,
		Name:     &newName,
	}

	expectedCode := http.StatusNotFound
	expected := shared.ErrRecordNotFound.Error()

	requestBody, err := json.Marshal(&requestData)
	require.NoError(t, err)

	mockUsecase := mocks.NewMockIUseCase(t)
	mockUsecase.EXPECT().GetClassByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, shared.ErrRecordNotFound)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPatch, "/classes/"+existingClass.ID.String(), bytes.NewReader(requestBody))

	controller := class.NewController(mockUsecase)

	router := gin.New()
	router.PATCH("/classes/:id", controller.UpdateClass)

	router.ServeHTTP(w, req)

	assert.Equal(t, expectedCode, w.Code)

	var response shared.ResponseError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, response.Error)
}

func TestUpdateClass_SystemIDTaken(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")
	otherClass := factory.NewClass(id, "class-C")

	newName := "class-B"
	requestData := class.UpdateClassRequest{
		SystemID: &otherClass.SystemID,
		Name:     &newName,
	}

	expectedCode := http.StatusConflict
	expected := shared.ErrSystemIDTaken.Error()

	requestBody, err := json.Marshal(&requestData)
	require.NoError(t, err)

	mockUsecase := mocks.NewMockIUseCase(t)
	mockUsecase.EXPECT().GetClassByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(existingClass, nil)
	mockUsecase.EXPECT().GetClassBySystemID(ctx, mock.AnythingOfType("string")).Return(otherClass, nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPatch, "/classes/"+existingClass.ID.String(), bytes.NewReader(requestBody))

	controller := class.NewController(mockUsecase)

	router := gin.New()
	router.PATCH("/classes/:id", controller.UpdateClass)

	router.ServeHTTP(w, req)

	assert.Equal(t, expectedCode, w.Code)

	var response shared.ResponseError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, response.Error)
}

func TestUpdateClass_Success(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")

	newName := "class-B"
	requestData := class.UpdateClassRequest{
		Name: &newName,
	}

	expectedCode := http.StatusOK
	expected := *existingClass
	expected.Name = newName

	requestBody, err := json.Marshal(&requestData)
	require.NoError(t, err)

	mockUsecase := mocks.NewMockIUseCase(t)

	mockResult := *existingClass
	mockResult.Name = newName

	mockUsecase.EXPECT().GetClassByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(existingClass, nil)
	mockUsecase.EXPECT().GetClassBySystemID(ctx, mock.AnythingOfType("string")).Return(nil, shared.ErrRecordNotFound)
	mockUsecase.EXPECT().UpdateClass(ctx, mock.AnythingOfType("*class.UpdateClassDTO")).Return(&mockResult, nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPatch, "/classes/"+existingClass.ID.String(), bytes.NewReader(requestBody))

	controller := class.NewController(mockUsecase)

	router := gin.New()
	router.PATCH("/classes/:id", controller.UpdateClass)

	router.ServeHTTP(w, req)

	assert.Equal(t, expectedCode, w.Code)

	var response shared.ResponseSuccess[domain.Class]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, response.Data)
}
