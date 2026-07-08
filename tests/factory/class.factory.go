package factory

import (
	"context"
	"main/internal/infrastructure/database"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func CreateClass(
	t *testing.T,
	course *database.Course,
) *database.Class {

	t.Helper()

	class := &database.Class{
		ID:       uuid.New(),
		CourseID: course.ID,
		Name:     "Physics",
	}

	err := database.DB.
		WithContext(context.Background()).
		Create(class).Error

	require.NoError(t, err)

	return class
}
