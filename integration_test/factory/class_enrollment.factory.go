package factory

import (
	"context"
	"main/internal/domain"
	"main/internal/infrastructure/database"
	classdomain "main/internal/modules/class/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func (f Factory) EnrollClass(
	t *testing.T,
	class *classdomain.Class,
	user *domain.User,
) {
	t.Helper()

	enrollment := database.ClassEnrollment{
		ClassID:   class.ID,
		StudentID: user.ID,
	}

	err := f.DB.
		WithContext(context.Background()).
		Create(&enrollment).Error

	require.NoError(t, err)
}
