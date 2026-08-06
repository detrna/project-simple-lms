package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"
)

func ToDomainCourseEnrollment(dbEnrollment *database.CourseEnrollment) *domain.CourseEnrollment {
	if dbEnrollment == nil {
		return nil
	}

	return &domain.CourseEnrollment{
		ID:          dbEnrollment.ID,
		CourseID:    dbEnrollment.CourseID,
		StudentID:   dbEnrollment.StudentID,
		Score:       dbEnrollment.Score,
		TeacherNote: dbEnrollment.TeacherNote,
		CreatedAt:   dbEnrollment.CreatedAt,
		UpdatedAt:   dbEnrollment.UpdatedAt,
	}
}

func ToDatabaseCourseEnrollment(enrollment *domain.CourseEnrollment) *database.CourseEnrollment {
	if enrollment == nil {
		return nil
	}

	return &database.CourseEnrollment{
		ID:          enrollment.ID,
		CourseID:    enrollment.CourseID,
		StudentID:   enrollment.StudentID,
		Score:       enrollment.Score,
		TeacherNote: enrollment.TeacherNote,
		CreatedAt:   enrollment.CreatedAt,
		UpdatedAt:   enrollment.UpdatedAt,
	}
}
