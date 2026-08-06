package repository

import (
	"context"
	"errors"
	"main/internal/domain"
	"main/internal/infrastructure/database"
	"main/internal/infrastructure/repository/mapper"
	"main/internal/shared"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) *ClassRepository {
	return &ClassRepository{db: db}
}

func (repo ClassRepository) GetStudents(ctx context.Context, classID uuid.UUID) ([]*domain.User, error) {
	rows, err := gorm.G[database.ClassEnrollment](repo.db).
		Preload("User", nil).
		Where("class_id = ?", classID).
		Find(ctx)

	if err != nil {
		return nil, shared.ErrRecordNotFound
	}

	var students []*domain.User

	for _, take := range rows {
		students = append(students, mapper.ToDomainUser(&take.Student))
	}

	return students, nil
}

func (repo ClassRepository) GetClassByID(ctx context.Context, classID uuid.UUID) (*domain.Class, error) {
	rows, err := gorm.G[database.Class](repo.db).Where("id = ?", classID).First(ctx)

	if err != nil {
		return nil, shared.ErrRecordNotFound
	}

	class := mapper.ToDomainClass(&rows)

	return class, nil
}

func (repo ClassRepository) CreateClass(ctx context.Context, data *domain.Class) (*domain.Class, error) {
	dbClass := mapper.ToDatabaseClass(data)

	err := gorm.G[database.Class](repo.db).Create(ctx, dbClass)
	if err != nil {
		return nil, err
	}

	createdClass := mapper.ToDomainClass(dbClass)

	return createdClass, nil
}

func (repo ClassRepository) UpdateClass(ctx context.Context, data *domain.Class) (*domain.Class, error) {
	dbClass := mapper.ToDatabaseClass(data)

	_, err := gorm.G[database.Class](repo.db).Where("id = ?", data.ID).Updates(ctx, *dbClass)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, shared.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	result, err := gorm.G[database.Class](repo.db).Where("id = ?", data.ID).First(ctx)

	if err != nil {
		return nil, shared.ErrRecordNotFound
	}

	class := mapper.ToDomainClass(&result)

	return class, nil
}

func (repo ClassRepository) DeleteClass(ctx context.Context, id uuid.UUID) error {
	_, err := gorm.G[database.Class](repo.db).Where("id = ?", id).Delete(ctx)

	return err
}
