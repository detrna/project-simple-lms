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

func TestGetClassBySystemID_ClassNotFound(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	_ = classfactory.NewClass(id, "class-A")

	requestData := "nonexistent-systemID"

	expected := shared.ErrRecordNotFound

	mockRepo := mocks.NewMockIRepository(t)
	mockRepo.EXPECT().GetClassBySystemID(ctx, mock.AnythingOfType("string")).Return(nil, shared.ErrRecordNotFound)

	usecase := class.NewUseCase(mockRepo)

	result, err := usecase.GetClassBySystemID(ctx, requestData)

	require.ErrorIs(t, err, expected)
	require.Nil(t, result)
}

func TestGetClassBySystemID_Success(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := classfactory.NewClass(id, "class-A")

	requestData := existingClass.SystemID

	expected := existingClass

	mockRepo := mocks.NewMockIRepository(t)
	mockRepo.EXPECT().GetClassBySystemID(ctx, mock.AnythingOfType("string")).Return(existingClass, nil)

	usecase := class.NewUseCase(mockRepo)

	result, err := usecase.GetClassBySystemID(ctx, requestData)

	require.NoError(t, err)
	require.Equal(t, expected, result)
}
