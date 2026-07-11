package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	SystemID  string
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
}

type JWTPayload struct {
	JTI      uuid.UUID
	UserID   uuid.UUID
	SystemID string
	Name     string
	Role     string
}
