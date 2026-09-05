package repository

import (
	"context"
	"errors"
	"main/internal/infrastructure/database"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/repository/mapper"
	"main/internal/shared/pagination"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) *ClassRepository {
	return &ClassRepository{db: db}
}

func (repo ClassRepository) GetAll(ctx context.Context, pagination pagination.Pagination) (*[]domain.Class, int, error) {
	rows, err := gorm.G[database.Class](repo.db).Limit(pagination.Limit).Offset(pagination.Offset).Order("created_at DESC").Find(ctx)

	if err != nil {
		return nil, 0, err
	}

	var classes []domain.Class
	for _, class := range rows {
		domainClass := mapper.ToDomainClass(&class)
		classes = append(classes, *domainClass)
	}

	count, err := gorm.G[database.Class](repo.db).Count(ctx, "*")
	if err != nil {
		return nil, 0, err
	}

	return &classes, int(count), nil
}

func (repo ClassRepository) GetByID(ctx context.Context, classID uuid.UUID) (*domain.Class, error) {
	rows, err := gorm.G[database.Class](repo.db).Where("id = ?", classID).First(ctx)

	if err != nil {
		return nil, domain.ErrClassNotFound
	}

	class := mapper.ToDomainClass(&rows)

	return class, nil
}

func (repo ClassRepository) GetBySystemID(ctx context.Context, systemID string) (*domain.Class, error) {
	rows, err := gorm.G[database.Class](repo.db).Where("system_id = ?", systemID).First(ctx)

	if err != nil {
		return nil, domain.ErrClassNotFound
	}

	class := mapper.ToDomainClass(&rows)

	return class, nil
}

func (repo ClassRepository) Create(ctx context.Context, data *domain.Class) (*domain.Class, error) {
	dbClass := mapper.ToDatabaseClass(data)

	err := gorm.G[database.Class](repo.db).Create(ctx, dbClass)
	if err != nil {
		return nil, err
	}

	createdClass := mapper.ToDomainClass(dbClass)

	return createdClass, nil
}

func (repo ClassRepository) Update(ctx context.Context, data *domain.Class) (*domain.Class, error) {
	dbClass := mapper.ToDatabaseClass(data)

	_, err := gorm.G[database.Class](repo.db).Where("id = ?", data.ID).Updates(ctx, *dbClass)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrClassNotFound
	}

	if err != nil {
		return nil, err
	}

	result, err := gorm.G[database.Class](repo.db).Where("id = ?", data.ID).First(ctx)

	if err != nil {
		return nil, domain.ErrClassNotFound
	}

	class := mapper.ToDomainClass(&result)

	return class, nil
}

func (repo ClassRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := gorm.G[database.Class](repo.db).Where("id = ?", id).Delete(ctx)

	return err
}
