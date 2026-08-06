package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"
)

func ToDomainSubmission(dbSubmission *database.Submission) *domain.Submission {
	if dbSubmission == nil {
		return nil
	}

	return &domain.Submission{
		ID:           dbSubmission.ID,
		AssignmentID: dbSubmission.AssignmentID,
		StudentID:    dbSubmission.StudentID,
		Score:        dbSubmission.Score,
		TeacherNote:  dbSubmission.TeacherNote,
		StudentNote:  dbSubmission.StudentNote,
		CreatedAt:    dbSubmission.CreatedAt,
		UpdatedAt:    dbSubmission.UpdatedAt,
		Student:      ToDomainMaskedUser(&dbSubmission.Student),
	}
}

func ToDatabaseSubmission(submission *domain.Submission) *database.Submission {
	if submission == nil {
		return nil
	}

	return &database.Submission{
		ID:           submission.ID,
		AssignmentID: submission.AssignmentID,
		StudentID:    submission.StudentID,
		Score:        submission.Score,
		TeacherNote:  submission.TeacherNote,
		StudentNote:  submission.StudentNote,
		CreatedAt:    submission.CreatedAt,
		UpdatedAt:    submission.UpdatedAt,
	}
}
