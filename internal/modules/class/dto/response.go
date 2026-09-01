package dto

import (
	"time"

	"github.com/google/uuid"
)

type ClassResponse struct {
	ID        uuid.UUID `json:"id"`
	SystemID  string    `json:"systemId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

