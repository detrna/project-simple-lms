package factory

import (
	"context"
	"crypto/rand"
	testsuite "main/integration_test"
	"main/internal/domain"
	"main/internal/infrastructure/database"
	"main/internal/infrastructure/repository/mapper"
	"main/internal/pkg"
	"math/big"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func CreateUser(
	t *testing.T,
	db *gorm.DB,
	hasher pkg.Hasher,
	name string,
) *domain.User {
	t.Helper()

	password := "password123"

	hashedPassword, err := hasher.Hash(password)
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

	err = db.
		WithContext(context.Background()).
		Create(user).Error

	require.NoError(t, err)

	return mapper.ToDomainUser(user)
}

func CreateAdmin(
	t *testing.T,
	db *gorm.DB,
	hasher pkg.Hasher,
	name string,
) *domain.User {
	t.Helper()

	password := "password123"

	hashedPassword, err := hasher.Hash(password)
	require.NoError(t, err)

	user := &database.User{
		ID:       uuid.New(),
		SystemID: "admin-test",
		Name:     "admin",
		Email:    "admin@mail.com",
		Password: string(hashedPassword),
		Role:     "admin",
	}

	err = db.
		WithContext(context.Background()).
		Create(user).Error

	require.NoError(t, err)

	return mapper.ToDomainUser(user)
}

func CreateOTP(t *testing.T, db *gorm.DB, suite *testsuite.Suite, user *domain.User) (string, error) {
	t.Helper()

	rng, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	otp := rng.String()

	if err := suite.Infra.RedisClient.Set(
		context.Background(),
		"otp:"+user.Email,
		otp,
		time.Duration(suite.Config.Mail.OTPExpiryMin)*time.Minute,
	); err != nil {
		return "", err
	}

	return otp, nil
}
