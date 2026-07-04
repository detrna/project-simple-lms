package factory

import (
	"context"
	"testing"

	"main/internal/database"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func CreateUser(
	t *testing.T,
	name string,
) *database.User {

	t.Helper()

	user := &database.User{
		ID:    uuid.New(),
		Name:  name,
		Email: name + "@gmail.com",
	}

	err := database.DB.
		WithContext(context.Background()).
		Create(user).Error

	require.NoError(t, err)

	return user
}
