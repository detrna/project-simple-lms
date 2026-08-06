package transcript

import (
	"github.com/google/uuid"
	"time"
)

type CreateAcademicTranscriptRequest struct {
	StudentID uuid.UUID `json:"studentId" binding:"required,uuid"`
}
type CreateSemesterTranscriptRequest struct {
	Period               string    `json:"period" binding:"required"`
	AcademicTranscriptID uuid.UUID `json:"academicTranscriptId" binding:"required,uuid"`
	CreatedAt            time.Time `json:"createdAt" binding:"required"`
	UpdatedAt            time.Time `json:"updatedAt" binding:"required"`
}
