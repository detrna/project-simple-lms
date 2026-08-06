package infrastructure

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"main/internal/config"
	"main/internal/domain"
	"main/internal/pkg"
	"main/internal/shared"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService struct {
	accessSecret        []byte
	refreshSecret       []byte
	accessExpiryMinutes int
	refreshExpiryDays   int
}

type Claims struct {
	Payload domain.JWTPayload

	jwt.RegisteredClaims
}

func NewTokenService(cfg *config.JWTConfig) pkg.TokenService {
	return &JWTService{
		accessSecret:        []byte(cfg.AccessSecret),
		refreshSecret:       []byte(cfg.RefreshSecret),
		accessExpiryMinutes: cfg.AccessExpiryMinutes,
		refreshExpiryDays:   cfg.RefreshExpiryDays,
	}
}

func (js *JWTService) GenerateAccessToken(data *domain.User) (*domain.JWT, error) {
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(js.accessExpiryMinutes) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "project-simple-lms",
		},
	}

	signed := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := signed.SignedString(js.accessSecret)

	if err != nil {
		return nil, err
	}

	token := domain.JWT{Payload: payload, Value: tokenString}

	return &token, nil
}

func (js *JWTService) GenerateRefreshToken(data *domain.User) (*domain.JWT, error) {
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(js.refreshExpiryDays) * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "project-simple-lms",
		},
	}

	signed := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := signed.SignedString(js.accessSecret)

	if err != nil {
		return nil, err
	}

	token := domain.JWT{Payload: payload, Value: tokenString}

	return &token, nil
}

func (js *JWTService) ParseAccessToken(tokenString string) (*domain.JWTPayload, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return js.accessSecret, nil
		},
	)

	if errors.Is(err, jwt.ErrInvalidKey) {
		return nil, shared.ErrInvalidToken
	}

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*Claims)

	payload := claims.Payload

	return &payload, nil
}

func (js *JWTService) ParseRefreshToken(tokenString string) (*domain.JWTPayload, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return js.refreshSecret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*Claims)

	payload := claims.Payload

	return &payload, nil
}

func (js *JWTService) HashToken(tokenString string) string {
	sum := sha256.Sum256([]byte(tokenString))

	return hex.EncodeToString(sum[:])
}

func (js *JWTService) Compare(hashed string, literal string) bool {
	hashedLiteral := js.HashToken(literal)
	return hashed == hashedLiteral
}
