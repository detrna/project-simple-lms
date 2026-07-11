package auth

import (
	"context"
	"main/internal/domain"
	"main/internal/infrastructure/database"
	"main/internal/pkg"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IRepository interface {
	CreateJWT(ctx context.Context, JWTPayload domain.JWTPayload, token string) (*string, error)
	DeleteJWT(ctx context.Context, ID uuid.UUID) error
	CheckJWT(ctx context.Context, ID uuid.UUID) error
}

type Repository struct {
	db     *gorm.DB
	logger pkg.Logger
}

func ToDatabaseJWT(JWT domain.JWTPayload, token string) *database.JWT {
	return &database.JWT{
		ID:     JWT.JTI,
		UserID: JWT.UserID,
		Token:  token,
	}
}

func NewRepository(db *gorm.DB, logger pkg.Logger) *Repository {
	return &Repository{db: db, logger: logger}
}

func (repo Repository) CreateJWT(ctx context.Context, JWTPayload domain.JWTPayload, token string) (*string, error) {
	jwt := ToDatabaseJWT(JWTPayload, token)
	err := gorm.G[database.JWT](repo.db).Create(ctx, jwt)

	if err != nil {
		return nil, err
	}

	return &jwt.Token, nil
}

func (repo Repository) DeleteJWT(ctx context.Context, ID uuid.UUID) error {
	_, err := gorm.G[database.JWT](repo.db).Where("jti = ?", ID).Delete(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (repo Repository) CheckJWT(ctx context.Context, ID uuid.UUID) error {
	_, err := gorm.G[database.JWT](repo.db).Where("jti = ?", ID).First(ctx)

	if err != nil {
		return err
	}

	return nil
}
