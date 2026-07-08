package domain

import (
	"time"

	"github.com/google/uuid"
)

type Assignment struct {
	ID          uuid.UUID
	ClassID     uuid.UUID
	Title       string
	Description string
	Deadline    time.Time
	CreatedAt   time.Time
}

type AssignmentFile struct {
	ID           uuid.UUID
	AssignmentID uuid.UUID
	URL          string
	CreatedAt    time.Time
}

type SubmissionFile struct {
	ID           uuid.UUID
	AssignmentID uuid.UUID
	UserID       uuid.UUID
	URL          string
	CreatedAt    time.Time
}

type SubmissionGrades struct {
	ID           uuid.UUID
	AssignmentID uuid.UUID
	UserID       uuid.UUID
	Grade        float64
	CreatedAt    time.Time
}
