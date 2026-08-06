package domain

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID           uuid.UUID
	SystemID     string
	Name         string
	Credits      int
	AcademicYear string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Teachers MaskedUser
}
