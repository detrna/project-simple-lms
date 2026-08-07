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

func TestClassByID_ClassNotFound(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	_ = classfactory.NewClass(id, "class-A")

	nonexistentID := uuid.New()
	requestData := nonexistentID

	expected := shared.ErrRecordNotFound

	mockRepo := mocks.NewMockIRepository(t)
	mockRepo.EXPECT().GetClassByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, shared.ErrRecordNotFound)

	usecase := class.NewUseCase(mockRepo)

	result, err := usecase.GetClassByID(ctx, requestData)

	require.ErrorIs(t, err, expected)
	require.Nil(t, result)
}

func TestClassByID_Success(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := classfactory.NewClass(id, "class-A")

	requestData := existingClass.ID

	expected := existingClass

	mockRepo := mocks.NewMockIRepository(t)
	mockRepo.EXPECT().GetClassByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(existingClass, nil)

	usecase := class.NewUseCase(mockRepo)

	result, err := usecase.GetClassByID(ctx, requestData)

	require.NoError(t, err)
	require.Equal(t, expected, result)
}
