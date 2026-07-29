package factory

import (
	"context"
	"main/internal/domain"
	"main/internal/infrastructure/database"
	"main/internal/infrastructure/repository/mapper"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func (f Factory) CreateUser(
	t *testing.T,
	name string,
) *domain.User {

	t.Helper()

	password := "password123"

	hashedPassword, err := f.Infra.Hasher.Hash(password)
	require.NoError(t, err)

	user := &database.User{
		ID:       uuid.New(),
		SystemID: name + "-test",
		Name:     name,
		Email:    name + "@mail.com",
		Password: string(hashedPassword),
		Role:     "user",
	}

	user.Email = "dzakiy1801@student.ub.ac.id"

	err = f.DB.
		WithContext(context.Background()).
		Create(user).Error

	require.NoError(t, err)

	return mapper.ToDomainUser(user)
}

func (f Factory) CreateAdmin(t *testing.T) *domain.User {
	t.Helper()

	password := "password123"

	hashedPassword, err := f.Infra.Hasher.Hash(password)
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

	return mapper.ToDomainUser(user)
}
