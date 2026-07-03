package user

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*User, error)
	FindByUserId(ctx context.Context, id uuid.UUID) (*User, error)
	Create(ctx context.Context, data User) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (repo Repository) FindById(ctx context.Context, id uuid.UUID) (*User, error) {
	rows, err := gorm.G[User](repo.db).
		Where("id = ?", id).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return &rows, nil
}

func (repo Repository) FindByUserId(ctx context.Context, id uuid.UUID) (*User, error) {
	rows, err := gorm.G[User](repo.db).
		Where("User_id = ?", id).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return &rows, nil
}

func (repo Repository) Create(ctx context.Context, data User) error {
	return gorm.G[User](repo.db).Create(ctx, &data)
}
