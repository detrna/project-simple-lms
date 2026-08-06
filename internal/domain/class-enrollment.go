package domain

import (
	"time"

	"github.com/google/uuid"
)

type ClassEnrollment struct {
	ID           uuid.UUID
	StudentID    uuid.UUID
	ClassID      uuid.UUID
	Status       string
	AcademicYear string
	LeftAt       time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
