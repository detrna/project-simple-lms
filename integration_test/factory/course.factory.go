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

func (f Factory) CreateCourse(t *testing.T) *domain.Course {
	t.Helper()

	course := &database.Course{
		ID:   uuid.New(),
		Name: "Computer Science",
	}

	err := f.DB.
		WithContext(context.Background()).
		Create(course).Error

	require.NoError(t, err)

	return mapper.ToDomainCourse(course)
}
