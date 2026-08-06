package domain

import (
	"time"

	"github.com/google/uuid"
)

type AcademicTranscript struct {
	ID                 uuid.UUID
	StudentID          uuid.UUID
	AcademicYear       string
	SemesterTranscript SemesterTranscript
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
