package domain

import (
	"time"

	"github.com/google/uuid"
)

type SemesterTranscript struct {
	ID                   uuid.UUID
	AcademicTranscriptID uuid.UUID
	course_enrollments   CourseEnrollment
	Period               string
	Grade                float64
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
