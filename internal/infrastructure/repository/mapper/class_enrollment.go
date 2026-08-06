package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"
)

func ToDomainClassEnrollment(dbEnrollment *database.ClassEnrollment) *domain.ClassEnrollment {
	if dbEnrollment == nil {
		return nil
	}

	return &domain.ClassEnrollment{
		ID:           dbEnrollment.ID,
		StudentID:    dbEnrollment.StudentID,
		ClassID:      dbEnrollment.ClassID,
		Status:       dbEnrollment.Status,
		AcademicYear: dbEnrollment.AcademicYear,
		LeftAt:       *dbEnrollment.LeftAt,
		CreatedAt:    dbEnrollment.CreatedAt,
		UpdatedAt:    dbEnrollment.UpdatedAt,
	}
}

func ToDatabaseClassEnrollment(enrollment *domain.ClassEnrollment) *database.ClassEnrollment {
	if enrollment == nil {
		return nil
	}

	return &database.ClassEnrollment{
		ID:           enrollment.ID,
		StudentID:    enrollment.StudentID,
		ClassID:      enrollment.ClassID,
		Status:       enrollment.Status,
		AcademicYear: enrollment.AcademicYear,
		LeftAt:       &enrollment.LeftAt,
		CreatedAt:    enrollment.CreatedAt,
		UpdatedAt:    enrollment.UpdatedAt,
	}
}
