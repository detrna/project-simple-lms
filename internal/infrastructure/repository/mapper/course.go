package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"
)

func ToDomainCourse(dbClass *database.Course) *domain.Course {
	return &domain.Course{
		ID:        dbClass.ID,
		Credits:   dbClass.Credits,
		Name:      dbClass.Name,
		CreatedAt: dbClass.CreatedAt,
	}
}
