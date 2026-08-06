package domain

import (
	"time"

	"github.com/google/uuid"
)

type Submission struct {
	ID           uuid.UUID
	AssignmentID uuid.UUID
	StudentID    uuid.UUID
	Score        float64
	TeacherNote  string
	StudentNote  string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Student MaskedUser
	Files   []File
}
