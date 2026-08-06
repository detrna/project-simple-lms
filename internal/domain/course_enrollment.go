package domain

import (
	"time"

	"github.com/google/uuid"
)

type CourseEnrollment struct {
	ID          uuid.UUID
	CourseID    uuid.UUID
	StudentID   uuid.UUID
	Score       float64
	TeacherNote string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
