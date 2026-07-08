package domain

import (
	"time"

	"github.com/google/uuid"
)

type Material struct {
	ID          uuid.UUID
	ClassID     uuid.UUID
	Title       string
	Description string
	CreatedAt   time.Time
}

type MaterialFile struct {
	ID         uuid.UUID
	MaterialID uuid.UUID
	URL        string
	CreatedAt  time.Time
}
