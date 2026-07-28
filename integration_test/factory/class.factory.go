package factory

import (
	"context"
	"main/internal/domain"
	"main/internal/infrastructure/database"
	"main/internal/infrastructure/repository/mapper"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func (f Factory) CreateClass(
	t *testing.T,
	course *domain.Course,
) *domain.Class {

	t.Helper()

	class := &database.Class{
		ID:       uuid.New(),
		CourseID: course.ID,
		Name:     "Physics",
	}

	err := f.DB.
		WithContext(context.Background()).
		Create(class).Error

	require.NoError(t, err)

	return mapper.ToDomainClass(class)
}
