package mapper

import (
	"main/internal/infrastructure/database"
	classdomain "main/internal/modules/class/domain"
)

func ToDomainClass(dbClass *database.Class) *classdomain.Class {
	if dbClass == nil {
		return nil
	}

	return &classdomain.Class{
		ID:        dbClass.ID,
		SystemID:  dbClass.SystemID,
		Name:      dbClass.Name,
		CreatedAt: dbClass.CreatedAt,
		UpdatedAt: dbClass.UpdatedAt,
	}
}

func ToDatabaseClass(classDomain *classdomain.Class) *database.Class {
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
