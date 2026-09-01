package usecase_test

import (
	"context"
	"main/internal/modules/class/factory"
	"main/internal/modules/class/usecase"
	"main/internal/modules/class/usecase/mocks"
	"main/internal/shared"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetBySystemID_ClassNotFound(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	_ = factory.NewClass(id, "class-A")

	requestData := "nonexistent-systemID"

	expected := shared.ErrRecordNotFound

	mockRepo := mocks.NewMockClassRepositoryI(t)
	mockRepo.EXPECT().GetBySystemID(ctx, mock.AnythingOfType("string")).Return(nil, shared.ErrRecordNotFound)

	classUseCase := usecase.NewClassUseCase(mockRepo)

	result, err := classUseCase.GetBySystemID(ctx, requestData)

	require.ErrorIs(t, err, expected)
	require.Nil(t, result)
}

func TestGetBySystemID_Success(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")

	requestData := existingClass.SystemID

	expected := existingClass

	mockRepo := mocks.NewMockClassRepositoryI(t)
	mockRepo.EXPECT().GetBySystemID(ctx, mock.AnythingOfType("string")).Return(existingClass, nil)

	classUseCase := usecase.NewClassUseCase(mockRepo)

	result, err := classUseCase.GetBySystemID(ctx, requestData)

	require.NoError(t, err)
	require.Equal(t, expected, result)
}

