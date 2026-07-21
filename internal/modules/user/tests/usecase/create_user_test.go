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

func TestCreateUser_Success(t *testing.T) {
	mockRepo := &user_mocks.MockIRepository{}
	mockBcrypt := &pkg_mocks.MockBcryptHasher{}
	u := user.NewUseCase(mockRepo, mockBcrypt, &pkg_mocks.MockLogger{})

	request := user.CreateUserSchema{
		SystemID: "user-test-1",
		Name:     "user-test",
		Email:    "user-test@mail.com",
		Password: "password123",
		Role:     "default",
	}

	expected := user.UserResponse{
		ID:        uuid.New(),
		SystemID:  request.SystemID,
		Name:      request.Name,
		Email:     request.Email,
		Role:      request.Role,
		CreatedAt: time.Now(),
	}

	ctx := context.Background()
	mockRepo.On("Create", ctx, mock.AnythingOfType("*domain.User")).
		Return(&domain.User{
			ID:        expected.ID,
			Name:      expected.Name,
			Email:     expected.Email,
			SystemID:  expected.SystemID,
			Role:      expected.Role,
			CreatedAt: expected.CreatedAt,
		}, nil)

	mockRepo.On("FindByEmail", ctx, mock.AnythingOfType("string")).Return(nil, shared.ErrRecordNotFound)

	mockBcrypt.On("Hash", mock.AnythingOfType("string")).Return([]byte("hashed-password"), nil)

	result, err := u.CreateUser(ctx, &request)

	require.NoError(t, err)

	assert.IsType(t, &user.UserResponse{}, result)
	assert.Equal(t, expected.SystemID, result.SystemID)
	assert.Equal(t, expected.Name, result.Name)
	assert.Equal(t, expected.Email, result.Email)
	assert.Equal(t, expected.Role, result.Role)
	assert.NotEqual(t, uuid.Nil, result.ID)
	assert.False(t, result.CreatedAt.IsZero())

	mockRepo.AssertExpectations(t)
}

func TestCreateUser_EmailTaken(t *testing.T) {
	mockRepo := &user_mocks.MockIRepository{}
	mockBcrypt := &pkg_mocks.MockBcryptHasher{}
	u := user.NewUseCase(mockRepo, mockBcrypt, &pkg_mocks.MockLogger{})

	request := user.CreateUserSchema{
		SystemID: "user-test-1",
		Name:     "user-test",
		Email:    "user-test-taken@mail.com",
		Password: "password123",
		Role:     "default",
	}

	ctx := context.Background()
	mockRepo.On("FindByEmail", ctx, mock.AnythingOfType("string")).Return(&domain.User{Email: request.Email}, nil)

	result, err := u.CreateUser(ctx, &request)

	require.Nil(t, result)
	require.ErrorIs(t, err, shared.ErrEmailTaken)

	mockRepo.AssertExpectations(t)
}
