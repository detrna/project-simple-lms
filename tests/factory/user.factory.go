package factory

import (
	"context"
	"main/internal/infrastructure/database"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func (f Factory) CreateUser(
	t *testing.T,
	name string,
) *database.User {

	t.Helper()

	password := "password123"

	hashedPassword, err := f.Infra.BcryptHasher.Hash(password)
	require.NoError(t, err)

	user := &database.User{
		ID:       uuid.New(),
		SystemID: name + "-test",
		Name:     name,
		Email:    name + "@mail.com",
		Password: string(hashedPassword),
		Role:     "user",
	}

	err = f.DB.
		WithContext(context.Background()).
		Create(user).Error

	require.NoError(t, err)

	return user
}

func (f Factory) CreateAdmin(t *testing.T) *database.User {
	t.Helper()

	password := "password123"

	hashedPassword, err := f.Infra.BcryptHasher.Hash(password)
	require.NoError(t, err)

	user := &database.User{
		ID:       uuid.New(),
		SystemID: "admin-test",
		Name:     "admin",
		Email:    "admin@mail.com",
		Password: string(hashedPassword),
		Role:     "admin",
	}

	err = f.DB.
		WithContext(context.Background()).
		Create(user).Error

	require.NoError(t, err)

	return user
}
