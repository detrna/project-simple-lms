package class_usecase_test

import (
	"context"
	"main/internal/modules/class"
	classfactory "main/internal/modules/class/tests/factory"
	"main/internal/modules/class/tests/mocks"

	"main/internal/shared"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateClass_SystemIDTaken(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := classfactory.NewClass(id, "class-A")

	requestData := class.CreateClassRequest{
		SystemID: existingClass.SystemID,
		Name:     "new-class",
	}

	expected := shared.ErrSystemIDTaken

	mockRepo := mocks.NewMockIRepository(t)
	mockRepo.EXPECT().GetClassByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(existingClass, nil)

	usecase := class.NewUseCase(mockRepo)

	result, err := usecase.CreateClass(ctx, &requestData)

	require.ErrorIs(t, err, expected)
	assert.Nil(t, result)
}

func TestCreateClass_Success(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	classSample := classfactory.NewClass(id, "class-A")

	requestData := class.CreateClassRequest{
		SystemID: classSample.SystemID,
		Name:     classSample.Name,
	}

	expected := classSample

	mockRepo := mocks.NewMockIRepository(t)
	mockRepo.EXPECT().GetClassByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, shared.ErrRecordNotFound)
	mockRepo.EXPECT().CreateClass(ctx, mock.AnythingOfType("*class.CreateClassRequest")).Return(classSample, nil)

	usecase := class.NewUseCase(mockRepo)

	result, err := usecase.CreateClass(ctx, &requestData)

	require.NoError(t, err)
	assert.Equal(t, expected, result)
}
