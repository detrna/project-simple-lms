package factory

import (
	"context"
	"testing"

	"main/internal/database"

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

	err := database.DB.
		WithContext(context.Background()).
		Create(&take).Error

	require.NoError(t, err)
}
