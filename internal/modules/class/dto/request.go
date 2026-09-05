package dto

import "github.com/google/uuid"

type CreateClassRequest struct {
	SystemID string `json:"systemId" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type UpdateClassRequest struct {
	ID       uuid.UUID `json:"-" uri:"id,parser=encoding.TextUnmarshaler"`
	SystemID *string   `json:"systemId"`
	Name     *string   `json:"name"`
}
