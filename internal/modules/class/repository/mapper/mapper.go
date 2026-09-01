package mapper

import (
	"main/internal/infrastructure/database"
	"main/internal/modules/class/domain"
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

func ToDatabaseClass(domainClass *domain.Class) *database.Class {
	if domainClass == nil {
		return nil
	}

	return &database.Class{
		ID:        domainClass.ID,
		SystemID:  domainClass.SystemID,
		Name:      domainClass.Name,
		CreatedAt: domainClass.CreatedAt,
		UpdatedAt: domainClass.UpdatedAt,
	}
}
