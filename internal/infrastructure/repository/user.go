package repository

import (
	"context"
	"errors"
	"main/internal/domain"
	"main/internal/infrastructure/database"
	"main/internal/infrastructure/repository/mapper"
	"main/internal/modules/user"
	"main/internal/shared"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.IRepository {
	return &UserRepository{db: db}
}

func (repo UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	rows, err := gorm.G[database.User](repo.db).
		Where("id = ?", id).
		First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, shared.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	user := mapper.ToDomainUser(rows)

	return &user, nil
}

func (repo UserRepository) FindBySystemID(ctx context.Context, id string) (*domain.User, error) {
	rows, err := gorm.G[database.User](repo.db).
		Where("system_id = ?", id).
		First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, shared.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	user := mapper.ToDomainUser(rows)

	return &user, nil
}

func (repo UserRepository) Create(ctx context.Context, data *domain.User) (*domain.User, error) {
	dbUser := mapper.ToDatabaseUser(*data)

	err := gorm.G[database.User](repo.db).Create(ctx, &dbUser)
	if err != nil {
		return nil, err
	}

	createdUser := mapper.ToDomainUser(dbUser)

	return &createdUser, nil
}

func (repo UserRepository) Update(ctx context.Context, data *domain.User) (*domain.User, error) {
	dbUser := mapper.ToDatabaseUser(*data)

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

	updatedUser := mapper.ToDomainUser(result)

	return &updatedUser, nil
}

func (repo UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[database.User](repo.db).
		Where("id = ?", id).
		Delete(ctx)

	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("User not found")
	}

	return nil
}

func (repo UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	rows, err := gorm.G[database.User](repo.db).
		Where("email = ?", email).
		First(ctx)

	if err == gorm.ErrRecordNotFound {
		return nil, shared.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	user := mapper.ToDomainUser(rows)

	return &user, nil
}
