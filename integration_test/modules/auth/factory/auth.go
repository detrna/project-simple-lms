package factory

import (
	"context"
	userfactory "main/integration_test/modules/user/factory"
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

func CreateUserJWT(t *testing.T, db *gorm.DB, tokenService pkg.TokenService, hasher pkg.Hasher) *domain.JWT {
	t.Helper()

	user := userfactory.CreateUser(t, db, hasher, "name")
	jwt, err := tokenService.GenerateAccessToken(user)
	require.NoError(t, err)

	hashed := tokenService.HashToken(jwt.Value)

	dbJWT := database.JWT{ID: jwt.Payload.JTI, UserID: jwt.Payload.UserID, Token: hashed}

	err = db.WithContext(context.Background()).
		Create(&dbJWT).Error

	require.NoError(t, err)

	return jwt
}

func CreateAdminJWT(t *testing.T, db *gorm.DB, tokenService pkg.TokenService, hasher pkg.Hasher) *domain.JWT {
	t.Helper()

	user := userfactory.CreateAdmin(t, db, hasher, "name")
	jwt, err := tokenService.GenerateAccessToken(user)
	require.NoError(t, err)

	hashed := tokenService.HashToken(jwt.Value)

	dbJWT := database.JWT{ID: jwt.Payload.JTI, UserID: jwt.Payload.UserID, Token: hashed}

	err = db.WithContext(context.Background()).
		Create(&dbJWT).Error

	require.NoError(t, err)

	return jwt
}
