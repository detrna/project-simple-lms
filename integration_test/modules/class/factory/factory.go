package factory

import (
	"context"
	"main/internal/infrastructure/database"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/repository/mapper"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func CreateClass(
	t *testing.T,
	db *gorm.DB,
	name string,
) *domain.Class {
	t.Helper()

	class := &database.Class{
		ID:   uuid.New(),
		Name: name,
	}

	err := db.
		WithContext(context.Background()).
		Create(class).Error

	require.NoError(t, err)

	return mapper.ToDomainClass(class)
}
