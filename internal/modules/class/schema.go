package class

import "github.com/google/uuid"

type CreateClassRequest struct {
	SystemID string `json:"systemId" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type UpdateClassRequest struct {
	SystemID *string `json:"systemId"`
	Name     *string `json:"name"`
}

type UpdateClassDTO struct {
	ID       uuid.UUID
	SystemID *string
	Name     *string
}
