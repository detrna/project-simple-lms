package factory

import (
	"context"
	"main/internal/domain"
	"main/internal/infrastructure/database"
	"main/internal/pkg"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func CreateJWT(t *testing.T, db *gorm.DB, tokenService pkg.TokenService, user *domain.User) *domain.JWT {
	t.Helper()

	jwt, err := tokenService.GenerateRefreshToken(user)
	require.NoError(t, err)

	hashed := tokenService.HashToken(jwt.Value)

	dbJWT := database.JWT{ID: jwt.Payload.JTI, UserID: jwt.Payload.UserID, Token: hashed}

	err = db.WithContext(context.Background()).
		Create(&dbJWT).Error

	require.NoError(t, err)

	return jwt
}
