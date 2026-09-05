package usecase_test

import (
	"context"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/factory"
	"main/internal/modules/class/usecase"
	"main/internal/modules/class/usecase/mocks"
	"main/internal/shared/pagination"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
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

	pagination := pagination.Pagination{Limit: 10, Offset: 0}

	expected := existingClasses

	mockRepo := mocks.NewMockClassRepositoryI(t)
	mockRepo.EXPECT().GetAll(ctx, mock.AnythingOfType("pagination.Pagination")).Return(&existingClasses, len(existingClasses), nil)

	classUseCase := usecase.NewClassUseCase(mockRepo)

	result, total, err := classUseCase.GetAll(ctx, &pagination)

	require.NoError(t, err)
	require.Equal(t, expected, *result)
	require.Equal(t, total, len(existingClasses))
}
