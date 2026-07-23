package class

import (
	"context"
	"main/internal/domain"
	"main/internal/infrastructure/database"
	"main/internal/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IRepository interface {
	GetStudents(ctx context.Context, classID uuid.UUID) ([]*domain.User, error)
	GetMyClasses(ctx context.Context, userID uuid.UUID) ([]*Class, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (repo Repository) GetStudents(ctx context.Context, classID uuid.UUID) ([]*domain.User, error) {
	rows, err := gorm.G[database.Takes](repo.db).
		Preload("User", nil).
		Where("class_id = ?", classID).
		Find(ctx)

	if err != nil {
		return nil, err
	}

	var students []*domain.User

	for _, take := range rows {
		students = append(students, mapper.ToDomainUser(&take.User))
	}

	return students, nil
}

func (repo Repository) GetMyClasses(ctx context.Context, userID uuid.UUID) ([]*Class, error) {
	rows, err := gorm.G[database.Takes](repo.db).
		Preload("Class", nil).
		Where("user_id = ?", userID).
		Find(ctx)

	if err != nil {
		return nil, err
	}

	var classes []*Class

	for _, take := range rows {
		classes = append(classes, ToDomainClass(&take.Class))
	}

	return classes, nil
}

func ToDomainClass(c *database.Class) *Class {
	return &Class{
		ID:        c.ID,
		Name:      c.Name,
		CourseID:  c.CourseID,
		CreatedAt: c.CreatedAt,
	}
}
