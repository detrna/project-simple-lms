package factory

import (
	"context"
	"main/internal/domain"
	"main/internal/infrastructure/database"
	"main/internal/infrastructure/repository/mapper"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func CreateCourse(t *testing.T, db *gorm.DB, name string) *domain.Course {
	t.Helper()

	course := &database.Course{
		ID:           uuid.New(),
		SystemID:     uuid.NewString(),
		Name:         name,
		Credits:      2,
		AcademicYear: "2026",
	}

	err := db.
		WithContext(context.Background()).
		Create(course).Error

	require.NoError(t, err)

	return mapper.ToDomainCourse(course)
}
