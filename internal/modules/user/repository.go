package user

import (
	"context"
	"main/internal/domain"

	"github.com/google/uuid"
)

type IRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	FindBySystemID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Create(ctx context.Context, data domain.User) (*domain.User, error)
	Update(ctx context.Context, data domain.User) (*domain.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
