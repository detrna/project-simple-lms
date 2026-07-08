package domain

import (
	"time"

	"github.com/google/uuid"
)

type Class struct {
	ID        uuid.UUID
	CourseID  uuid.UUID
	Name      string
	CreatedAt time.Time
}

type Takes struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ClassID   uuid.UUID
	Grade     float64
	CreatedAt time.Time
}

type Teaches struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ClassID   uuid.UUID
	CreatedAt time.Time
}
