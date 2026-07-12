package factory

import (
	"context"
	"main/internal/infrastructure/database"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func (f Factory) CreateUser(
	t *testing.T,
	name string,
) *database.User {

	t.Helper()

	cost := f.Config.Bcrypt.Cost

	password := "password123"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	require.NoError(t, err)

	user := &database.User{
		ID:       uuid.New(),
		Name:     name,
		Email:    name + "@mail.com",
		Password: string(hashedPassword),
	}

	err = f.DB.
		WithContext(context.Background()).
		Create(user).Error

	require.NoError(t, err)

	return user
}
