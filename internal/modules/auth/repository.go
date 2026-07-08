package auth

import (
	"context"
	"main/internal/infrastructure/database"
	"main/internal/shared"

	"gorm.io/gorm"
)

type IRepository interface {
	CreateJWT(ctx context.Context, JWT JWT) (*string, error)
}

type Repository struct {
	db     *gorm.DB
	logger shared.Logger
}

func ToDatabaseJWT(JWT JWT) *database.JWT {
	return &database.JWT{
		ID:        JWT.ID,
		UserID:    JWT.UserID,
		Value:     JWT.Value,
		CreatedAt: JWT.CreatedAt,
	}
}

func NewRepository(db *gorm.DB, logger shared.Logger) *Repository {
	return &Repository{db: db, logger: logger}
}

func (repo Repository) CreateJWT(ctx context.Context, JWT JWT) (*string, error) {
	jwt := ToDatabaseJWT(JWT)
	err := gorm.G[database.JWT](repo.db).Create(ctx, jwt)

	if err != nil {
		return nil, err
	}

	return &jwt.Value, nil
}
