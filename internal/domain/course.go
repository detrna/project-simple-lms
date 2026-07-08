package domain

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID        uuid.UUID
	Name      string
	Credits   int
	CreatedAt time.Time
}
