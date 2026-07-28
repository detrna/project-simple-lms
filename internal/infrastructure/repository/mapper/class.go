package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"
)

func ToDomainClass(dbClass *database.Class) *domain.Class {
	return &domain.Class{
		ID:        dbClass.ID,
		Course:    domain.Course(dbClass.Course),
		Name:      dbClass.Name,
		CreatedAt: dbClass.CreatedAt,
	}
}
