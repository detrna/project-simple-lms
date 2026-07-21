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
	"github.com/stretchr/testify/mock"
)

func TestDeleteUser_Success(t *testing.T) {
	mockRepo := &user_mocks.MockIRepository{}
	u := user.NewUseCase(mockRepo, &pkg_mocks.MockBcryptHasher{}, &pkg_mocks.MockLogger{})

	id := uuid.New()
	existingAccount := domain.User{
		ID: id,
	}

	ctx := context.Background()
	mockRepo.On("FindByID", ctx, mock.AnythingOfType("uuid.UUID")).Return(&existingAccount, nil)
	mockRepo.On("Delete", ctx, mock.AnythingOfType("uuid.UUID")).Return(nil)

	err := u.DeleteUser(ctx, id)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteUser_RecordNotFound(t *testing.T) {
	mockRepo := &user_mocks.MockIRepository{}
	u := user.NewUseCase(mockRepo, &pkg_mocks.MockBcryptHasher{}, &pkg_mocks.MockLogger{})

	id := uuid.New()
	ctx := context.Background()
	mockRepo.On("FindByID", ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, shared.ErrRecordNotFound)

	err := u.DeleteUser(ctx, id)
	assert.ErrorIs(t, err, shared.ErrRecordNotFound)

	mockRepo.AssertExpectations(t)
}
