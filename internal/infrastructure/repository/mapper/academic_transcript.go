package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"
)

func ToDomainAcademicTranscript(dbTranscript *database.AcademicTranscript) *domain.AcademicTranscript {
	if dbTranscript == nil {
		return nil
	}

	return &domain.AcademicTranscript{
		ID:                 dbTranscript.ID,
		StudentID:          dbTranscript.StudentID,
		AcademicYear:       "",
		SemesterTranscript: domain.SemesterTranscript{},
		CreatedAt:          dbTranscript.CreatedAt,
		UpdatedAt:          dbTranscript.UpdatedAt,
	}
}

func ToDatabaseAcademicTranscript(transcript *domain.AcademicTranscript) *database.AcademicTranscript {
	if transcript == nil {
		return nil
	}

	return &database.AcademicTranscript{
		ID:        transcript.ID,
		StudentID: transcript.StudentID,
		CreatedAt: transcript.CreatedAt,
		UpdatedAt: transcript.UpdatedAt,
	}
}
