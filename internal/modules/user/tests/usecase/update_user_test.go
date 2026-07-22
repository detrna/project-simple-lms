package user_usecase_test

import (
	"context"
	"main/internal/domain"
	"main/internal/modules/user"
	user_mocks "main/internal/modules/user/mocks"
	user_factory "main/internal/modules/user/tests"
	pkg_mocks "main/internal/pkg/mocks"
	"main/internal/shared"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdateUser_Success(t *testing.T) {
	mockRepo := user_mocks.NewMockIRepository(t)

	id := uuid.New()
	existingAccount := user_factory.NewUser(id)

	newPassword := "password321"

	requestData := user.UpdateUserSchema{
		ID:       existingAccount.ID,
		Password: &newPassword,
	}

	repoRequest := domain.User{
		ID:       existingAccount.ID,
		Password: newPassword,
		SystemID: existingAccount.SystemID,
		Email:    existingAccount.Email,
		Role:     existingAccount.Role,
		Name:     existingAccount.Name,
	}

	repoResult := repoRequest

	expected := user.UserResponse{
		ID:       repoResult.ID,
		SystemID: repoResult.SystemID,
		Name:     repoResult.Name,
		Email:    repoResult.Email,
		Role:     repoResult.Role,
	}

	ctx := context.Background()
	mockRepo.On("FindByID", ctx, mock.AnythingOfType("uuid.UUID")).Return(&existingAccount, nil)
	mockRepo.On("Update", ctx, mock.AnythingOfType("domain.User")).Return(&repoResult, nil)

	u := user.NewUseCase(mockRepo, pkg_mocks.NewMockBcryptHasher(t), pkg_mocks.NewMockLogger(t))

	result, err := u.UpdateUser(ctx, requestData)
	require.NoError(t, err)

	assert.Equal(t, expected, result)

	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_RecordNotFound(t *testing.T) {
	mockRepo := user_mocks.NewMockIRepository(t)

	id := uuid.New()
	existingAccount := user_factory.NewUser(id)

	newPassword := "password321"

	requestData := user.UpdateUserSchema{
		ID:       existingAccount.ID,
		Password: &newPassword,
	}

	ctx := context.Background()
	mockRepo.On("FindByID", ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, shared.ErrRecordNotFound)

	u := user.NewUseCase(mockRepo, pkg_mocks.NewMockBcryptHasher(t), pkg_mocks.NewMockLogger(t))

	result, err := u.UpdateUser(ctx, requestData)
	require.ErrorIs(t, shared.ErrRecordNotFound, err)

	assert.Nil(t, result)

	mockRepo.AssertExpectations(t)
}
