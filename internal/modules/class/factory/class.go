package factory

import (
	"main/internal/modules/class/domain"
	"time"

	"github.com/google/uuid"
)

func NewClass(id uuid.UUID, name string) *domain.Class {
	return &domain.Class{
		ID:        id,
		SystemID:  name,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
