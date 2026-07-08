package factory

import (
	"context"
	"main/internal/infrastructure/database"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func CreateCourse(t *testing.T) *database.Course {
	t.Helper()

	course := &database.Course{
		ID:   uuid.New(),
		Name: "Computer Science",
	}

	err := database.DB.
		WithContext(context.Background()).
		Create(course).Error

	require.NoError(t, err)

	return course
}
