package domain

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID          uuid.UUID
	Name        string
	FileURL     string
	Size        float64
	ContentType string
	ParentType  string
	ParentID    uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
