package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"
	"time"
)

func ToDomainEnrollment(dbEnrollment *database.Enrollment) *domain.Enrollment {
	if dbEnrollment == nil {
		return nil
	}

	var leftAtValue time.Time
	if dbEnrollment.LeftAt != nil {
		leftAtValue = *dbEnrollment.LeftAt
	}

	teacherNote := ""
	if dbEnrollment.TeacherNote != nil {
		teacherNote = *dbEnrollment.TeacherNote
	}

	score := 0.0
	if dbEnrollment.Score != nil {
		score = float64(*dbEnrollment.Score)
	}

	return &domain.Enrollment{
		ID:           dbEnrollment.ID,
		CourseID:     dbEnrollment.CourseID,
		StudentID:    dbEnrollment.StudentID,
		ClassID:      dbEnrollment.ClassID,
		Status:       dbEnrollment.Status,
		AcademicYear: dbEnrollment.AcademicYear,
		Score:        score,
		TeacherNote:  teacherNote,
		LeftAt:       leftAtValue,
		CreatedAt:    dbEnrollment.CreatedAt,
		UpdatedAt:    dbEnrollment.UpdatedAt,
	}
}

func ToDatabaseEnrollment(enrollment *domain.Enrollment) *database.Enrollment {
	if enrollment == nil {
		return nil
	}

	var teacherNote *string
	if enrollment.TeacherNote != "" {
		teacherNote = &enrollment.TeacherNote
	}

	var score *int
	if enrollment.Score != 0 {
		scoreValue := int(enrollment.Score)
		score = &scoreValue
	}

	var leftAtValue *time.Time
	if !enrollment.LeftAt.IsZero() {
		leftAtValue = &enrollment.LeftAt
	}

	return &database.Enrollment{
		ID:           enrollment.ID,
		CourseID:     enrollment.CourseID,
		StudentID:    enrollment.StudentID,
		ClassID:      enrollment.ClassID,
		Status:       enrollment.Status,
		AcademicYear: enrollment.AcademicYear,
		Score:        score,
		TeacherNote:  teacherNote,
		LeftAt:       leftAtValue,
		CreatedAt:    enrollment.CreatedAt,
		UpdatedAt:    enrollment.UpdatedAt,
	}
}

func ToDomainCourseEnrollment(dbEnrollment *database.Enrollment) *domain.Enrollment {
	return ToDomainEnrollment(dbEnrollment)
}

func ToDatabaseCourseEnrollment(enrollment *domain.Enrollment) *database.Enrollment {
	return ToDatabaseEnrollment(enrollment)
}
