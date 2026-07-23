package auth

import (
	"context"
	"main/internal/domain"

	"github.com/google/uuid"
)

type IRepository interface {
	CreateJWT(ctx context.Context, JWTPayload *domain.JWTPayload, token string) (*string, error)
	DeleteJWT(ctx context.Context, ID uuid.UUID) error
	FindJWT(ctx context.Context, JTI uuid.UUID) (*string, error)
}
