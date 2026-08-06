package class_enrollment

import (
	"github.com/google/uuid"
	"time"
)

type CreateClassEnrollmentRequest struct {
	AcademicYear string    `json:"academicYear" binding:"required"`
	ClassID      uuid.UUID `json:"classId" binding:"required,uuid"`
}
type PatchClassEnrollmentRequest struct {
	Status       *string    `json:"status"`
	AcademicYear *string    `json:"academicYear"`
	ClassID      *uuid.UUID `json:"classId"`
	LeftAt       *time.Time `json:"leftAt"`
}
