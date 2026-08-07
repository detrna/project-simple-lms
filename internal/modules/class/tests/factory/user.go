package factory

import (
	classdomain "main/internal/modules/class/domain"
	"time"

	"github.com/google/uuid"
)

func NewClass(id uuid.UUID, name string) *classdomain.Class {
	return &classdomain.Class{
		ID:        id,
		SystemID:  name,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
