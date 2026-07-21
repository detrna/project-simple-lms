package user_usecase_test

import (
	"context"
	"main/internal/domain"
	"main/internal/modules/user"
	user_mocks "main/internal/modules/user/mocks"
	pkg_mocks "main/internal/pkg/mocks"
	"main/internal/shared"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserByID_Success(t *testing.T) {
	mockRepo := &user_mocks.MockIRepository{}

	uuid := uuid.New()
	expectedPayload := &domain.User{ID: uuid, Name: "test"}

	ctx := context.Background()
	mockRepo.On(
		"FindByID",
		ctx,
		uuid,
	).Return(expectedPayload, nil)

	useCase := user.NewUseCase(mockRepo, &pkg_mocks.MockBcryptHasher{}, &pkg_mocks.MockLogger{})

	result, err := useCase.GetUserByID(ctx, uuid)
	require.NoError(t, err)

	assert.Equal(t, result.ID, expectedPayload.ID)

	mockRepo.AssertExpectations(t)
}

func TestGetUserByID_RecordNotFound(t *testing.T) {
	mockRepo := &user_mocks.MockIRepository{}

	uuid := uuid.New()

	ctx := context.Background()
	mockRepo.On(
		"FindByID",
		ctx,
		uuid,
	).Return(&domain.User{}, shared.ErrRecordNotFound)

	useCase := user.NewUseCase(mockRepo, &pkg_mocks.MockBcryptHasher{}, &pkg_mocks.MockLogger{})

	result, err := useCase.GetUserByID(ctx, uuid)

	require.ErrorIs(t, err, shared.ErrRecordNotFound)
	assert.Nil(t, result)

	mockRepo.AssertExpectations(t)
}
