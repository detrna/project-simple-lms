package user_controller_test

import (
	"main/internal/domain"
	pkg_mocks "main/internal/pkg/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func NewUserSample(id uuid.UUID) *domain.User {
	user := domain.User{
		ID:        id,
		SystemID:  "user-test-1",
		Name:      "user-test",
		Role:      "default",
		Email:     "user-test@mail.com",
		Password:  "password123",
		CreatedAt: time.Now(),
	}

	return &user
}

func NewMockLogger(t *testing.T) *pkg_mocks.MockLogger {
	mockLogger := pkg_mocks.NewMockLogger(t)
	mockLogger.On("Warn", mock.Anything).Return().Maybe()

	return mockLogger
}
