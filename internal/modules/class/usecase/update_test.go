package usecase_test

import (
	"context"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/dto"
	"main/internal/modules/class/factory"
	"main/internal/modules/class/usecase"
	"main/internal/modules/class/usecase/mocks"
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

	expected := domain.ErrClassNotFound

	mockRepo := mocks.NewMockClassRepositoryI(t)
	mockRepo.EXPECT().GetByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, domain.ErrClassNotFound)

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

	expected := domain.ErrClassSystemIDTaken

	mockRepo := mocks.NewMockClassRepositoryI(t)
	mockRepo.EXPECT().GetByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(existingClass, nil)
	mockRepo.EXPECT().GetBySystemID(ctx, mock.AnythingOfType("string")).Return(existingClass, nil)

	classUseCase := usecase.NewClassUseCase(mockRepo)

	result, err := classUseCase.Update(ctx, &requestData)

	require.ErrorIs(t, err, expected)
	require.Nil(t, result)
}

func TestUpdate_SuccessWithoutSystemID(t *testing.T) {
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
	mockRepo.EXPECT().Update(ctx, mock.AnythingOfType("*domain.Class")).Return(expected, nil)

	classUseCase := usecase.NewClassUseCase(mockRepo)

	result, err := classUseCase.Update(ctx, &requestData)

	require.NoError(t, err)
	require.Equal(t, expected, result)
}

func TestUpdate_SuccessWithSystemID(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")

	newName := "new-class-name"
	newSystemID := "new-system-ID"

	requestData := dto.UpdateClassRequest{
		ID:       existingClass.ID,
		SystemID: &newSystemID,
		Name:     &newName,
	}

	expected := existingClass
	expected.Name = newName

	mockRepo := mocks.NewMockClassRepositoryI(t)
	mockRepo.EXPECT().GetByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(existingClass, nil)
	mockRepo.EXPECT().GetBySystemID(ctx, mock.AnythingOfType("string")).Return(nil, domain.ErrClassNotFound)
	mockRepo.EXPECT().Update(ctx, mock.AnythingOfType("*domain.Class")).Return(expected, nil)

	classUseCase := usecase.NewClassUseCase(mockRepo)

	result, err := classUseCase.Update(ctx, &requestData)

	require.NoError(t, err)
	require.Equal(t, expected, result)
}
