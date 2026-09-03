package usecase_test

import (
	"context"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/factory"
	"main/internal/modules/class/usecase"
	"main/internal/modules/class/usecase/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestClassByID_ClassNotFound(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	_ = factory.NewClass(id, "class-A")

	nonexistentID := uuid.New()
	requestData := nonexistentID

	expected := domain.ErrClassNotFound

	mockRepo := mocks.NewMockClassRepositoryI(t)
	mockRepo.EXPECT().GetByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, domain.ErrClassNotFound)

	classUseCase := usecase.NewClassUseCase(mockRepo)

	result, err := classUseCase.GetByID(ctx, requestData)

	require.ErrorIs(t, err, expected)
	require.Nil(t, result)
}

func TestClassByID_Success(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")

	requestData := existingClass.ID

	expected := existingClass

	mockRepo := mocks.NewMockClassRepositoryI(t)
	mockRepo.EXPECT().GetByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(existingClass, nil)

	classUseCase := usecase.NewClassUseCase(mockRepo)

	result, err := classUseCase.GetByID(ctx, requestData)

	require.NoError(t, err)
	require.Equal(t, expected, result)
}
