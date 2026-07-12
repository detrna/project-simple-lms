package factory

import (
	"context"
	"main/internal/infrastructure/database"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func (f Factory) CreateCourse(t *testing.T) *database.Course {
	t.Helper()

	course := &database.Course{
		ID:   uuid.New(),
		Name: "Computer Science",
	}

	err := f.DB.
		WithContext(context.Background()).
		Create(course).Error

	require.NoError(t, err)

	return course
}
