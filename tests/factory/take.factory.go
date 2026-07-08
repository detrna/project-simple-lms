package factory

import (
	"context"
	"main/internal/infrastructure/database"
	"testing"

	"github.com/stretchr/testify/require"
)

func EnrollStudent(
	t *testing.T,
	class *database.Class,
	user *database.User,
) {

	t.Helper()

	take := database.Takes{
		ClassID: class.ID,
		UserID:  user.ID,
	}

	err := DB.
		WithContext(context.Background()).
		Create(&take).Error

	require.NoError(t, err)
}
