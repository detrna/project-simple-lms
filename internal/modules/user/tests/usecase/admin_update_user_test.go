package user_usecase_test

import (
	"context"
	"main/internal/domain"
	"main/internal/modules/user"
	user_mocks "main/internal/modules/user/mocks"
	pkg_mocks "main/internal/pkg/mocks"
	"main/internal/shared"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAdminUpdateUser_Success(t *testing.T) {
	mockRepo := &user_mocks.MockIRepository{}
	mockBcrypt := &pkg_mocks.MockBcryptHasher{}
	u := user.NewUseCase(mockRepo, mockBcrypt, &pkg_mocks.MockLogger{})

	id := uuid.New()
	systemID := "user-test-1"
	email := "user-test@mail.com"
	createdAt := time.Now()

	newName := "user-test-updated"

	request := user.AdminUpdateUserSchema{
		ID:       &id,
		SystemID: &systemID,
		Email:    &email,
		Name:     &newName,
	}

	existingAccount := domain.User{
		ID:        id,
		SystemID:  systemID,
		Name:      "user-test",
		Email:     email,
		Role:      "default",
		CreatedAt: createdAt,
	}

	expected := domain.User{
		ID:       existingAccount.ID,
		SystemID: existingAccount.SystemID,
		Name:     *request.Name,
		Email:    existingAccount.Email,
		Role:     existingAccount.Role,
	}

	ctx := context.Background()
	mockRepo.On("Update", ctx, mock.AnythingOfType("*domain.User")).
		Return(&expected, nil)

	mockRepo.On("FindByID", ctx, mock.AnythingOfType("uuid.UUID")).Return(&existingAccount, nil)
	mockRepo.On("FindByEmail", ctx, mock.AnythingOfType("string")).Return(nil, shared.ErrRecordNotFound)
	mockRepo.On("FindBySystemID", ctx, mock.AnythingOfType("string")).Return(nil, shared.ErrRecordNotFound)

	result, err := u.AdminUpdateUser(ctx, &request)
	require.NoError(t, err)

	assert.IsType(t, &user.UserResponse{}, result)
	assert.Equal(t, expected.SystemID, result.SystemID)
	assert.Equal(t, expected.Name, result.Name)
	assert.Equal(t, expected.Email, result.Email)
	assert.Equal(t, expected.Role, result.Role)
	assert.Equal(t, expected.ID, result.ID)

	mockRepo.AssertExpectations(t)
}

func TestAdminUpdateUser_RecordNotFound(t *testing.T) {
	mockRepo := &user_mocks.MockIRepository{}
	mockBcrypt := &pkg_mocks.MockBcryptHasher{}
	u := user.NewUseCase(mockRepo, mockBcrypt, &pkg_mocks.MockLogger{})

	id := uuid.New()
	newName := "user-test-updated"

	request := user.AdminUpdateUserSchema{
		ID:   &id,
		Name: &newName,
	}

	ctx := context.Background()
	mockRepo.On("FindByID", ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, shared.ErrRecordNotFound)

	result, err := u.AdminUpdateUser(ctx, &request)
	require.Error(t, err, shared.ErrRecordNotFound)

	assert.Nil(t, result)

	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_EmailTaken(t *testing.T) {
	mockRepo := &user_mocks.MockIRepository{}
	mockBcrypt := &pkg_mocks.MockBcryptHasher{}
	u := user.NewUseCase(mockRepo, mockBcrypt, &pkg_mocks.MockLogger{})

	id := uuid.New()
	systemID := "user-test-1"
	email := "user-test@mail.com"
	newEmail := "user-test-updated@mail.com"
	createdAt := time.Now()

	request := user.AdminUpdateUserSchema{
		ID:    &id,
		Email: &newEmail,
	}

	existingAccount := domain.User{
		ID:        id,
		SystemID:  systemID,
		Name:      "user-test",
		Email:     email,
		Role:      "default",
		CreatedAt: createdAt,
	}

	ctx := context.Background()
	mockRepo.On("FindByID", ctx, mock.AnythingOfType("uuid.UUID")).Return(&existingAccount, nil)
	mockRepo.On("FindByEmail", ctx, mock.AnythingOfType("string")).Return(&existingAccount, nil)

	result, err := u.AdminUpdateUser(ctx, &request)
	require.Error(t, err, shared.ErrEmailTaken)

	assert.Nil(t, result)

	mockRepo.AssertExpectations(t)
}

func TestAdminUpdateUser_SystemIDTaken(t *testing.T) {

	mockRepo := &user_mocks.MockIRepository{}
	mockBcrypt := &pkg_mocks.MockBcryptHasher{}
	u := user.NewUseCase(mockRepo, mockBcrypt, &pkg_mocks.MockLogger{})

	id := uuid.New()
	systemID := "user-test-1"
	email := "user-test@mail.com"
	newSystemID := "user-test-1-updated"
	createdAt := time.Now()

	request := user.AdminUpdateUserSchema{
		ID:       &id,
		SystemID: &newSystemID,
	}

	existingAccount := domain.User{
		ID:        id,
		SystemID:  systemID,
		Name:      "user-test",
		Email:     email,
		Role:      "default",
		CreatedAt: createdAt,
	}

	ctx := context.Background()
	mockRepo.On("FindByID", ctx, mock.AnythingOfType("uuid.UUID")).Return(&existingAccount, nil)
	mockRepo.On("FindBySystemID", ctx, mock.AnythingOfType("string")).Return(&existingAccount, nil)

	result, err := u.AdminUpdateUser(ctx, &request)
	require.Error(t, err, shared.ErrSystemIDTaken)

	assert.Nil(t, result)

	mockRepo.AssertExpectations(t)

}
