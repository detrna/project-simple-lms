package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"
)

func ToDomainSemesterTranscript(dbTranscript *database.SemesterTranscript) *domain.SemesterTranscript {
	if dbTranscript == nil {
		return nil
	}

	return &domain.SemesterTranscript{
		ID:                   dbTranscript.ID,
		AcademicTranscriptID: dbTranscript.AcademicTranscriptID,
		Period:               dbTranscript.Period,
		CreatedAt:            dbTranscript.CreatedAt,
		UpdatedAt:            dbTranscript.UpdatedAt,
	}
}

func ToDatabaseSemesterTranscript(transcript *domain.SemesterTranscript) *database.SemesterTranscript {
	if transcript == nil {
		return nil
	}

	return &database.SemesterTranscript{
		ID:                   transcript.ID,
		AcademicTranscriptID: transcript.AcademicTranscriptID,
		Period:               transcript.Period,
		CreatedAt:            transcript.CreatedAt,
		UpdatedAt:            transcript.UpdatedAt,
	}
}
