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

func TestDeleteClass_ClassNotFound(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	_ = classfactory.NewClass(id, "class-A") //existing class

	nonexistentID := uuid.New()
	requestData := nonexistentID

	expected := shared.ErrRecordNotFound

	mockRepo := mocks.NewMockIRepository(t)
	mockRepo.EXPECT().GetClassByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, shared.ErrRecordNotFound)

	usecase := class.NewUseCase(mockRepo)

	err := usecase.DeleteClass(ctx, requestData)
	require.ErrorIs(t, err, expected)
}

func TestDeleteClass_Success(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := classfactory.NewClass(id, "class-A")

	requestData := existingClass.ID

	mockRepo := mocks.NewMockIRepository(t)
	mockRepo.EXPECT().GetClassByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(existingClass, nil)
	mockRepo.EXPECT().DeleteClass(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil)

	usecase := class.NewUseCase(mockRepo)

	err := usecase.DeleteClass(ctx, requestData)
	require.NoError(t, err)
}
