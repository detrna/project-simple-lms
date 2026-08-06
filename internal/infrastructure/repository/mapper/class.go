package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"
)

func ToDomainClass(dbClass *database.Class) *domain.Class {
	if dbClass == nil {
		return nil
	}

	return &domain.Class{
		ID:        dbClass.ID,
		SystemID:  dbClass.SystemID,
		Name:      dbClass.Name,
		CreatedAt: dbClass.CreatedAt,
		UpdatedAt: dbClass.UpdatedAt,
	}
}

func ToDatabaseClass(classDomain *domain.Class) *database.Class {
	if classDomain == nil {
		return nil
	}

	return &database.Class{
		ID:        classDomain.ID,
		SystemID:  classDomain.SystemID,
		Name:      classDomain.Name,
		CreatedAt: classDomain.CreatedAt,
		UpdatedAt: classDomain.UpdatedAt,
	}
}
