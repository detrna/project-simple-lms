package usecase_test

import (
	"context"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/factory"
	"main/internal/modules/class/usecase"
	"main/internal/modules/class/usecase/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetAll_Success(t *testing.T) {
	ctx := context.Background()

	existingClasses := []domain.Class{
		*factory.NewClass(uuid.New(), "class-A"),
		*factory.NewClass(uuid.New(), "class-B"),
		*factory.NewClass(uuid.New(), "class-C"),
		*factory.NewClass(uuid.New(), "class-D"),
		*factory.NewClass(uuid.New(), "class-E"),
	}

	expected := existingClasses

	mockRepo := mocks.NewMockClassRepositoryI(t)
	mockRepo.EXPECT().GetAll(ctx).Return(&existingClasses, nil)

	classUseCase := usecase.NewClassUseCase(mockRepo)

	result, err := classUseCase.GetAll(ctx)

	require.NoError(t, err)
	require.Equal(t, expected, result)
}

