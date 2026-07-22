package user_factory

import (
	"main/internal/domain"
	"time"

	"github.com/google/uuid"
)

func NewUser(id uuid.UUID) *domain.User {
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
