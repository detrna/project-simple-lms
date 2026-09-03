package domain

import (
	"time"

	"github.com/google/uuid"
)

type Enrollment struct {
	ID           uuid.UUID
	CourseID     uuid.UUID
	StudentID    uuid.UUID
	ClassID      uuid.UUID
	Status       string
	AcademicYear string
	Score        float64
	TeacherNote  string
	LeftAt       time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
