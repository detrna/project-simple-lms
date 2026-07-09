package shared

import (
	"main/internal/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTPayload = domain.JWT

type Claims struct {
	Payload JWTPayload

	jwt.RegisteredClaims
}

var accessSecret = os.Getenv("JWT_ACCESS_SECRET")
var refreshSecret = os.Getenv("JWT_REFRESH_SECRET")

func GenerateAccessToken(data JWTPayload) (string, error) {
	claims := Claims{
		Payload: JWTPayload{
			JTI:      uuid.New(),
			UserID:   data.UserID,
			SystemID: data.SystemID,
			Role:     data.Role,
			Name:     data.Name,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "project-simple-lms",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(accessSecret))
}

func GenerateRefreshToken(data JWTPayload) (string, error) {
	claims := Claims{
		Payload: JWTPayload{
			JTI:      uuid.New(),
			UserID:   data.UserID,
			SystemID: data.SystemID,
			Role:     data.Role,
			Name:     data.Name,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "project-simple-lms",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(refreshSecret))
}

func ParseToken(tokenString string, key string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*Claims)

	return claims, nil
}
