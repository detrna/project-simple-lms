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

func (jwtProvider *JWTProvider) GenerateAccessToken(data *domain.User) (*domain.JWT, error) {
	payload := domain.JWTPayload{
		JTI:      uuid.New(),
		UserID:   data.ID,
		SystemID: data.SystemID,
		Role:     data.Role,
		Name:     data.Name,
	}

	claims := Claims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtProvider.accessExpiryMinutes) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "project-simple-lms",
		},
	}

	signed := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := signed.SignedString(jwtProvider.accessSecret)

	if err != nil {
		return nil, err
	}

	token := domain.JWT{Payload: payload, Value: tokenString}

	return &token, nil
}

func (jwtProvider *JWTProvider) GenerateRefreshToken(data *domain.User) (*domain.JWT, error) {
	payload := domain.JWTPayload{
		JTI:      uuid.New(),
		UserID:   data.ID,
		SystemID: data.SystemID,
		Role:     data.Role,
		Name:     data.Name,
	}

	claims := Claims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtProvider.refreshExpiryDays) * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "project-simple-lms",
		},
	}

	signed := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := signed.SignedString(jwtProvider.accessSecret)

	if err != nil {
		return nil, err
	}

	token := domain.JWT{Payload: payload, Value: tokenString}

	return &token, nil
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
