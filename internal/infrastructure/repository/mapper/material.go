package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"

	"github.com/google/uuid"
)

func ToDomainMaterial(dbMaterial *database.Material) *domain.Material {
	if dbMaterial == nil {
		return nil
	}

	return &domain.Material{
		ID:          dbMaterial.ID,
		CourseID:    dbMaterial.CourseID,
		Title:       dbMaterial.Title,
		Description: dbMaterial.Description,
		CreatedAt:   dbMaterial.CreatedAt,
		UpdatedAt:   dbMaterial.UpdatedAt,
	}
}

func ToDatabaseMaterial(material *domain.Material) *database.Material {
	if material == nil {
		return nil
	}

	return &database.Material{
		ID:          material.ID,
		ClassID:     uuid.Nil,
		CourseID:    material.CourseID,
		Title:       material.Title,
		Description: material.Description,
		CreatedAt:   material.CreatedAt,
		UpdatedAt:   material.UpdatedAt,
	}
}
