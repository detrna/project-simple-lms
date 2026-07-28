package factory

import (
	"context"
	"main/internal/domain"
	"main/internal/infrastructure/database"
	"testing"

	"github.com/stretchr/testify/require"
)

func (f Factory) EnrollStudent(
	t *testing.T,
	class *domain.Class,
	user *domain.User,
) {
	t.Helper()

	take := database.Takes{
		ClassID: class.ID,
		UserID:  user.ID,
	}

	err := f.DB.
		WithContext(context.Background()).
		Create(&take).Error

	require.NoError(t, err)
}
