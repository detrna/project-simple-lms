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

type JWT struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Value     string
	CreatedAt time.Time
}
