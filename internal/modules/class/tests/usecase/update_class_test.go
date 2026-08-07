package class_usecase_test

import (
	"context"
	"main/internal/modules/class"
	classfactory "main/internal/modules/class/tests/factory"
	"main/internal/modules/class/tests/mocks"
	"main/internal/shared"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdateClass_ClassNotFound(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	_ = classfactory.NewClass(id, "class-A")

	nonexistentID := uuid.New()
	newName := "new-class-name"

	requestData := class.UpdateClassDTO{
		ID:   nonexistentID,
		Name: &newName,
	}

	expected := shared.ErrRecordNotFound

	mockRepo := mocks.NewMockIRepository(t)
	mockRepo.EXPECT().GetClassByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, shared.ErrRecordNotFound)

	usecase := class.NewUseCase(mockRepo)

	result, err := usecase.UpdateClass(ctx, &requestData)

	require.ErrorIs(t, err, expected)
	require.Nil(t, result)
}

func TestUpdateClass_SystemIDTaken(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := classfactory.NewClass(id, "class-A")

	newName := "new-class-name"
	newSystemID := "taken-systemID"

	requestData := class.UpdateClassDTO{
		ID:       existingClass.ID,
		Name:     &newName,
		SystemID: &newSystemID,
	}

	expected := shared.ErrSystemIDTaken

	mockRepo := mocks.NewMockIRepository(t)
	mockRepo.EXPECT().GetClassByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(existingClass, nil)
	mockRepo.EXPECT().GetClassBySystemID(ctx, mock.AnythingOfType("string")).Return(existingClass, nil)

	usecase := class.NewUseCase(mockRepo)

	result, err := usecase.UpdateClass(ctx, &requestData)

	require.ErrorIs(t, err, expected)
	require.Nil(t, result)
}

func TestUpdateClass_Success(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := classfactory.NewClass(id, "class-A")

	newName := "new-class-name"
	requestData := class.UpdateClassDTO{
		ID:   existingClass.ID,
		Name: &newName,
	}

	expected := existingClass
	expected.Name = newName

	mockRepo := mocks.NewMockIRepository(t)
	mockRepo.EXPECT().GetClassByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(existingClass, nil)
	mockRepo.EXPECT().GetClassBySystemID(ctx, mock.AnythingOfType("string")).Return(nil, shared.ErrRecordNotFound)
	mockRepo.EXPECT().UpdateClass(ctx, mock.AnythingOfType("*class.UpdateClassDTO")).Return(expected, nil)

	usecase := class.NewUseCase(mockRepo)

	result, err := usecase.UpdateClass(ctx, &requestData)

	require.NoError(t, err)
	require.Equal(t, expected, result)
}
