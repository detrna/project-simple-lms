package factory

import (
	"context"
	"main/internal/domain"
	"main/internal/infrastructure/database"
	"main/internal/infrastructure/repository/mapper"
	classdomain "main/internal/modules/class/domain"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func Enroll(
	t *testing.T,
	db *gorm.DB,
	class *classdomain.Class,
	user *domain.User,
	course *domain.Course,
) domain.Enrollment {
	t.Helper()

	enrollment := database.Enrollment{
		ClassID:      class.ID,
		StudentID:    user.ID,
		CourseID:     course.ID,
		Status:       "active",
		AcademicYear: "2026/2027",
	}

	err := db.
		WithContext(context.Background()).
		Create(&enrollment).Error

	require.NoError(t, err)

	return *mapper.ToDomainEnrollment(&enrollment)
}
