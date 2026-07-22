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

func NewJWTPayload(data *domain.User) *domain.JWTPayload {
	jwt := domain.JWTPayload{
		JTI:      uuid.New(),
		UserID:   data.ID,
		SystemID: data.SystemID,
		Name:     data.Name,
		Role:     data.Name,
	}

	return &jwt
}
