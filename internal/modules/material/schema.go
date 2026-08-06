package material

import "github.com/google/uuid"

type CreateMaterialRequest struct {
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description" binding:"required"`
	FileIDs     []uuid.UUID `json:"fileIds" binding:"required"`
}

type PatchMaterialRequest struct {
	Title       *string      `json:"title"`
	Description *string      `json:"description"`
	FileIDs     *[]uuid.UUID `json:"fileIds"`
}
