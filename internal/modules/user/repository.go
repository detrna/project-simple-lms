package user

import (
	"context"
	"errors"
	"main/internal/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
	FindBySystemID(ctx context.Context, id uuid.UUID) (*User, error)
	Create(ctx context.Context, data User) (*User, error)
	Update(ctx context.Context, data User) (*User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func ToDomainUser(u database.User) User {
	return User{
		ID:        u.ID,
		SystemID:  u.SystemID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}

func ToDatabaseUser(u User) database.User {
	return database.User{
		ID:        u.ID,
		SystemID:  u.SystemID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}

func (repo Repository) FindByID(ctx context.Context, id uuid.UUID) (*User, error) {
	rows, err := gorm.G[User](repo.db).
		Where("id = ?", id).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return &rows, nil
}

func (repo Repository) FindBySystemID(ctx context.Context, id uuid.UUID) (*User, error) {
	rows, err := gorm.G[User](repo.db).
		Where("User_id = ?", id).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return &rows, nil
}

func (repo Repository) Create(ctx context.Context, data User) (*User, error) {
	dbUser := ToDatabaseUser(data)

	err := gorm.G[database.User](repo.db).Create(ctx, &dbUser)
	if err != nil {
		return nil, err
	}

	createdUser := ToDomainUser(dbUser)

	return &createdUser, nil
}

func (repo Repository) Update(ctx context.Context, data User) (*User, error) {
	dbUser := ToDatabaseUser(data)

	var result database.User

	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		_, err := gorm.G[database.User](tx).
			Where("id = ?", data.ID).
			Updates(ctx, dbUser)
		if err != nil {
			return err
		}

		rows, err := gorm.G[database.User](tx).
			Where("id = ?", data.ID).
			First(ctx)
		if err != nil {
			return err
		}

		result = rows

		return nil
	}); err != nil {
		return nil, err
	}

	updatedUser := ToDomainUser(result)

	return &updatedUser, nil
}

func (repo Repository) Delete(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[database.User](repo.db).Where("id = ?", id).Delete(ctx)

	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("User not found")
	}

	return nil
}
