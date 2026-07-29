package factory

import (
	"context"
	"main/internal/domain"
	"main/internal/infrastructure/database"
	"testing"

	"github.com/stretchr/testify/require"
)

func (f Factory) CreateJWT(t *testing.T, user *domain.User) *domain.JWT {
	t.Helper()

	jwt, err := f.Infra.TokenService.GenerateRefreshToken(user)
	require.NoError(t, err)

	hashed := f.Infra.TokenService.HashToken(jwt.Value)

	dbJWT := database.JWT{ID: jwt.Payload.JTI, UserID: jwt.Payload.UserID, Token: hashed}

	err = f.DB.WithContext(context.Background()).
		Create(&dbJWT).Error

	require.NoError(t, err)

	return jwt
}
