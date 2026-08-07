package domain

import (
	"time"

	"github.com/google/uuid"
)

type Class struct {
	ID        uuid.UUID
	SystemID  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Takes struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ClassID   uuid.UUID
	Grade     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Teaches struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ClassID   uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
