package repository

import (
	"context"
	"main/internal/domain"
	"main/internal/infrastructure/database"
	"main/internal/infrastructure/repository/mapper"
	"main/internal/modules/auth"
	"main/internal/pkg"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db     *gorm.DB
	logger pkg.Logger
}

func NewAuthRepository(db *gorm.DB, logger pkg.Logger) auth.IRepository {
	return &AuthRepository{db: db, logger: logger}
}

func (repo AuthRepository) CreateJWT(ctx context.Context, JWTPayload domain.JWTPayload, token string) (*string, error) {
	jwt := mapper.ToDatabaseJWT(JWTPayload, token)
	err := gorm.G[database.JWT](repo.db).Create(ctx, jwt)

	if err != nil {
		return nil, err
	}

	return &jwt.Token, nil
}

func (repo AuthRepository) DeleteJWT(ctx context.Context, ID uuid.UUID) error {
	_, err := gorm.G[database.JWT](repo.db).Where("jti = ?", ID).Delete(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (repo AuthRepository) CheckJWT(ctx context.Context, ID uuid.UUID) error {
	_, err := gorm.G[database.JWT](repo.db).Where("jti = ?", ID).First(ctx)

	if err != nil {
		return err
	}

	return nil
}
