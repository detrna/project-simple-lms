package class_usecase_test

import (
	"context"
	"main/internal/modules/class"
	"main/internal/modules/class/domain"
	classfactory "main/internal/modules/class/tests/factory"
	"main/internal/modules/class/tests/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetClasses_Success(t *testing.T) {
	ctx := context.Background()

	existingClasses := []domain.Class{
		*classfactory.NewClass(uuid.New(), "class-A"),
		*classfactory.NewClass(uuid.New(), "class-B"),
		*classfactory.NewClass(uuid.New(), "class-C"),
		*classfactory.NewClass(uuid.New(), "class-D"),
		*classfactory.NewClass(uuid.New(), "class-E"),
	}

	expected := existingClasses

	mockRepo := mocks.NewMockIRepository(t)
	mockRepo.EXPECT().GetClasses(ctx).Return(&existingClasses, nil)

	usecase := class.NewUseCase(mockRepo)

	result, err := usecase.GetClasses(ctx)

	require.NoError(t, err)
	require.Equal(t, expected, result)
}
