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

func TestDelete_ClassNotFound(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	_ = factory.NewClass(id, "class-A")

	nonexistentID := uuid.New()
	requestData := nonexistentID

	expected := domain.ErrClassNotFound

	mockRepo := mocks.NewMockClassRepositoryI(t)
	mockRepo.EXPECT().GetByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, domain.ErrClassNotFound)

	classUseCase := usecase.NewClassUseCase(mockRepo)

	err := classUseCase.Delete(ctx, requestData)
	require.ErrorIs(t, err, expected)
}

func TestDelete_Success(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingClass := factory.NewClass(id, "class-A")

	requestData := existingClass.ID

	mockRepo := mocks.NewMockClassRepositoryI(t)
	mockRepo.EXPECT().GetByID(ctx, mock.AnythingOfType("uuid.UUID")).Return(existingClass, nil)
	mockRepo.EXPECT().Delete(ctx, mock.AnythingOfType("uuid.UUID")).Return(nil)

	classUseCase := usecase.NewClassUseCase(mockRepo)

	err := classUseCase.Delete(ctx, requestData)
	require.NoError(t, err)
}
