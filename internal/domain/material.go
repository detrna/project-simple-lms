package domain

import (
	"time"

	"github.com/google/uuid"
)

type Material struct {
	ID          uuid.UUID
	CourseID    uuid.UUID
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Files []File
}
