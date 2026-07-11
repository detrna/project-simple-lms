package infrastructure

import (
	"main/internal/config"
	"main/internal/domain"
	"main/internal/pkg"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTProvider struct {
	accessSecret        []byte
	refreshSecret       []byte
	accessExpiryMinutes int
	refreshExpiryDays   int
}

type Claims struct {
	Payload domain.JWTPayload

	jwt.RegisteredClaims
}

func NewTokenProvider(cfg config.JWTConfig) pkg.JWTProvider {
	return &JWTProvider{
		accessSecret:        []byte(cfg.AccessSecret),
		refreshSecret:       []byte(cfg.RefreshSecret),
		accessExpiryMinutes: cfg.AccessExpiryMinutes,
		refreshExpiryDays:   cfg.RefreshExpiryDays,
	}
}

func (provider *JWTProvider) GenerateAccessToken(data domain.JWTPayload) (string, error) {
	claims := Claims{
		Payload: domain.JWTPayload{
			JTI:      uuid.New(),
			UserID:   data.UserID,
			SystemID: data.SystemID,
			Role:     data.Role,
			Name:     data.Name,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(provider.accessExpiryMinutes) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "project-simple-lms",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(provider.accessSecret)
}

func (provider *JWTProvider) GenerateRefreshToken(data domain.JWTPayload) (string, error) {
	claims := Claims{
		Payload: domain.JWTPayload{
			JTI:      uuid.New(),
			UserID:   data.UserID,
			SystemID: data.SystemID,
			Role:     data.Role,
			Name:     data.Name,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(provider.refreshExpiryDays) * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "project-simple-lms",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(provider.refreshSecret)
}

func (provider *JWTProvider) ParseAccessToken(tokenString string) (*domain.JWTPayload, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return provider.accessSecret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*Claims)

	payload := claims.Payload

	return &payload, nil
}

func (provider *JWTProvider) ParseRefreshToken(tokenString string) (*domain.JWTPayload, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return provider.refreshSecret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*Claims)

	payload := claims.Payload

	return &payload, nil
}
