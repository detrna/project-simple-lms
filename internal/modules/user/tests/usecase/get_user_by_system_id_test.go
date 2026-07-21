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

func TestGetUserBySystemID_Success(t *testing.T) {
	mockRepo := &user_mocks.MockIRepository{}

	uuid := uuid.New()
	systemID := "user-test"
	expectedPayload := &domain.User{ID: uuid, SystemID: systemID, Name: "test"}

	ctx := context.Background()
	mockRepo.On(
		"FindBySystemID",
		ctx,
		systemID,
	).Return(expectedPayload, nil)

	useCase := user.NewUseCase(mockRepo, &pkg_mocks.MockBcryptHasher{}, &pkg_mocks.MockLogger{})

	result, err := useCase.GetUserBySystemID(ctx, systemID)
	require.NoError(t, err)

	assert.Equal(t, result.SystemID, expectedPayload.SystemID)

	mockRepo.AssertExpectations(t)
}

func TestGetUserBySystemID_RecordNotFound(t *testing.T) {
	mockRepo := &user_mocks.MockIRepository{}

	systemID := "user-test"

	ctx := context.Background()
	mockRepo.On(
		"FindBySystemID",
		ctx,
		systemID,
	).Return(&domain.User{}, shared.ErrRecordNotFound)

	useCase := user.NewUseCase(mockRepo, &pkg_mocks.MockBcryptHasher{}, &pkg_mocks.MockLogger{})

	result, err := useCase.GetUserBySystemID(ctx, systemID)

	require.ErrorIs(t, err, shared.ErrRecordNotFound)
	assert.Nil(t, result)

	mockRepo.AssertExpectations(t)
}
