package dto

import (
	"github.com/google/uuid"
)

type IDParams struct {
	ID uuid.UUID `uri:"id,parser=encoding.TextUnmarshaler" binding:"required"`
}

type SystemIDParams struct {
	SystemID string `uri:"id" binding:"required"`
}
