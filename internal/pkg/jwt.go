package pkg

import "main/internal/domain"

type JWTProvider interface {
	GenerateAccessToken(data domain.JWTPayload) (string, error)
	GenerateRefreshToken(data domain.JWTPayload) (string, error)
	ParseAccessToken(tokenString string) (*domain.JWTPayload, error)
	ParseRefreshToken(tokenString string) (*domain.JWTPayload, error)
}
