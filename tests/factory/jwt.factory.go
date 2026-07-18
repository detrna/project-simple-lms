package factory

import (
	"context"
	"main/internal/domain"
	"main/internal/infrastructure/database"
	"testing"

	"github.com/stretchr/testify/require"
)

func (f Factory) CreateJWT(t *testing.T, user *domain.User) domain.JWT {
	t.Helper()

	jwt, err := f.Infra.JWTProvider.GenerateAccessToken(user)

	require.NoError(t, err)

	dbJWT := database.JWT{ID: jwt.Payload.JTI, UserID: jwt.Payload.UserID, Token: jwt.Value}

	err = f.DB.WithContext(context.Background()).
		Create(&dbJWT).Error

	require.NoError(t, err)

	return *jwt
}
