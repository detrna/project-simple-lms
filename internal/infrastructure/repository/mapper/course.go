package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"

	"github.com/google/uuid"
)

func ToDomainCourse(dbCourse *database.Course) *domain.Course {
	if dbCourse == nil {
		return nil
	}

	var teacher domain.MaskedUser
	if len(dbCourse.Teachers) > 0 {
		teacher = ToDomainMaskedUser(&dbCourse.Teachers[0])
	}

	return &domain.Course{
		ID:           dbCourse.ID,
		SystemID:     dbCourse.SystemID,
		Name:         dbCourse.Name,
		Credits:      dbCourse.Credits,
		AcademicYear: dbCourse.AcademicYear,
		CreatedAt:    dbCourse.CreatedAt,
		UpdatedAt:    dbCourse.UpdatedAt,
		Teachers:     teacher,
	}
}

func ToDatabaseCourse(course *domain.Course) *database.Course {
	if course == nil {
		return nil
	}

	teachers := []database.User{}
	if course.Teachers.ID != uuid.Nil {
		teachers = append(teachers, *ToDatabaseMaskedUser(&course.Teachers))
	}

	return &database.Course{
		ID:           course.ID,
		SystemID:     course.SystemID,
		Name:         course.Name,
		Credits:      course.Credits,
		AcademicYear: course.AcademicYear,
		CreatedAt:    course.CreatedAt,
		UpdatedAt:    course.UpdatedAt,
		Teachers:     teachers,
	}
}
