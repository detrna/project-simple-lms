package usecase_test

import (
	"context"
	"main/internal/modules/class/dto"
	"main/internal/modules/class/factory"
	"main/internal/modules/class/usecase"
	"main/internal/modules/class/usecase/mocks"
	"main/internal/shared"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdate_ClassNotFound(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	_ = factory.NewClass(id, "class-A")

	nonexistentID := uuid.New()
	newName := "new-class-name"

	requestData := dto.UpdateClassRequest{
		ID:   nonexistentID,
		Name: &newName,
	}

	expected := shared.ErrRecordNotFound

	mockRepo := mocks.NewMockClassRepositoryI(t)
	mockRepo.EXPECT().GetByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, shared.ErrRecordNotFound)

	classUseCase := usecase.NewClassUseCase(mockRepo)

	result, err := classUseCase.Update(ctx, &requestData)

	require.ErrorIs(t, err, expected)
	require.Nil(t, result)
}

func TestUpdate_SystemIDTaken(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")

	newName := "new-class-name"
	newSystemID := "taken-systemID"

	requestData := dto.UpdateClassRequest{
		ID:       existingClass.ID,
		Name:     &newName,
		SystemID: &newSystemID,
	}

	expected := shared.ErrSystemIDTaken

	mockRepo := mocks.NewMockClassRepositoryI(t)
	mockRepo.EXPECT().GetByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(existingClass, nil)
	mockRepo.EXPECT().GetBySystemID(ctx, mock.AnythingOfType("string")).Return(existingClass, nil)

	classUseCase := usecase.NewClassUseCase(mockRepo)

	result, err := classUseCase.Update(ctx, &requestData)

	require.ErrorIs(t, err, expected)
	require.Nil(t, result)
}

func TestUpdate_Success(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")

	newName := "new-class-name"
	requestData := dto.UpdateClassRequest{
		ID:   existingClass.ID,
		Name: &newName,
	}

	expected := existingClass
	expected.Name = newName

	mockRepo := mocks.NewMockClassRepositoryI(t)
	mockRepo.EXPECT().GetByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(existingClass, nil)
	mockRepo.EXPECT().GetBySystemID(ctx, mock.AnythingOfType("string")).Return(nil, shared.ErrRecordNotFound)
	mockRepo.EXPECT().Update(ctx, mock.AnythingOfType("*dto.UpdateClassRequest")).Return(expected, nil)

	classUseCase := usecase.NewClassUseCase(mockRepo)

	result, err := classUseCase.Update(ctx, &requestData)

	require.NoError(t, err)
	require.Equal(t, expected, result)
}

