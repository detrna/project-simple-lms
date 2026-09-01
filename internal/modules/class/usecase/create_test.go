package usecase_test

import (
	"context"

	"main/internal/modules/class/dto"
	"main/internal/modules/class/factory"
	"main/internal/modules/class/usecase/mocks"

	"main/internal/modules/class/usecase"
	"main/internal/shared"
	"testing"

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
		Name:     "new-class",
	}

	expected := shared.ErrSystemIDTaken

	mockRepo := mocks.NewMockClassRepositoryI(t)
	mockRepo.EXPECT().GetByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(existingClass, nil)

	classUseCase := usecase.NewClassUseCase(mockRepo)

	result, err := classUseCase.Create(ctx, &requestData)

	require.ErrorIs(t, err, expected)
	assert.Nil(t, result)
}

func TestCreate_Success(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	classSample := factory.NewClass(id, "class-A")

	requestData := dto.CreateClassRequest{
		SystemID: classSample.SystemID,
		Name:     classSample.Name,
	}

	expected := classSample

	mockRepo := mocks.NewMockClassRepositoryI(t)
	mockRepo.EXPECT().GetByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, shared.ErrRecordNotFound)
	mockRepo.EXPECT().Create(ctx, mock.AnythingOfType("*dto.CreateClassRequest")).Return(classSample, nil)

	classUseCase := usecase.NewClassUseCase(mockRepo)

	result, err := classUseCase.Create(ctx, &requestData)

	require.NoError(t, err)
	assert.Equal(t, expected, result)
}

