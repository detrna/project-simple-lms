package pkg

import "main/internal/domain"

type JWTProvider interface {
	GenerateAccessToken(data *domain.User) (*domain.JWT, error)
	GenerateRefreshToken(data *domain.User) (*domain.JWT, error)
	ParseAccessToken(tokenString string) (*domain.JWTPayload, error)
	ParseRefreshToken(tokenString string) (*domain.JWTPayload, error)
	HashToken(tokenString string) string
	Compare(hashed string, literal string) bool
}
