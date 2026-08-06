package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"

	"github.com/google/uuid"
)

func ToDomainAssignment(dbAssignment *database.Assignment) *domain.Assignment {
	if dbAssignment == nil {
		return nil
	}

	teachers := make([]domain.MaskedUser, 0, len(dbAssignment.Teachers))
	for _, teacher := range dbAssignment.Teachers {
		teachers = append(teachers, ToDomainMaskedUser(&teacher))
	}

	return &domain.Assignment{
		ID:          dbAssignment.ID,
		CourseID:    dbAssignment.CourseID,
		Title:       dbAssignment.Title,
		Description: dbAssignment.Description,
		Deadline:    dbAssignment.Deadline,
		CreatedAt:   dbAssignment.CreatedAt,
		UpdatedAt:   dbAssignment.UpdatedAt,
		Teachers:    teachers,
	}
}

func ToDatabaseAssignment(assignment *domain.Assignment) *database.Assignment {
	if assignment == nil {
		return nil
	}

	teachers := make([]database.User, 0, len(assignment.Teachers))
	for _, teacher := range assignment.Teachers {
		teachers = append(teachers, *ToDatabaseMaskedUser(&teacher))
	}

	return &database.Assignment{
		ID:          assignment.ID,
		ClassID:     uuid.Nil,
		CourseID:    assignment.CourseID,
		Title:       assignment.Title,
		Description: assignment.Description,
		Deadline:    assignment.Deadline,
		CreatedAt:   assignment.CreatedAt,
		UpdatedAt:   assignment.UpdatedAt,
		Teachers:    teachers,
	}
}
