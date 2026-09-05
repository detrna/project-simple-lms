package dto

import (
	"main/internal/modules/class/domain"
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

func DomainToResponse(c domain.Class) *ClassResponse {
	return &ClassResponse{
		ID:        c.ID,
		SystemID:  c.SystemID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func DomainToResponseBatch(classes []domain.Class) *[]ClassResponse {
	var response []ClassResponse

	for _, c := range classes {
		response = append(response, ClassResponse{
			ID:        c.ID,
			SystemID:  c.SystemID,
			Name:      c.Name,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}

	return &response
}
