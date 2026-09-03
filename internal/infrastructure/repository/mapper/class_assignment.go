package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"
	"time"
)

func ToDomainClassEnrollment(dbEnrollment *database.Enrollment) *domain.Enrollment {
	return ToDomainEnrollment(dbEnrollment)
}

func ToDatabaseClassEnrollment(enrollment *domain.Enrollment) *database.Enrollment {
	return ToDatabaseEnrollment(enrollment)
}

func ToDomainLegacyClassEnrollment(dbEnrollment *database.Enrollment) *domain.Enrollment {
	return ToDomainEnrollment(dbEnrollment)
}

func ToDatabaseLegacyClassEnrollment(enrollment *domain.Enrollment) *database.Enrollment {
	return ToDatabaseEnrollment(enrollment)
}

func ToDomainClassAssignment(dbEnrollment *database.Enrollment) *domain.Enrollment {
	return ToDomainEnrollment(dbEnrollment)
}

func ToDatabaseClassAssignment(enrollment *domain.Enrollment) *database.Enrollment {
	return ToDatabaseEnrollment(enrollment)
}

func ToDomainClassEnrollmentValue(dbEnrollment *database.Enrollment) *domain.Enrollment {
	var leftAtValue time.Time
	if dbEnrollment != nil && dbEnrollment.LeftAt != nil {
		leftAtValue = *dbEnrollment.LeftAt
	}

	if dbEnrollment == nil {
		return nil
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
