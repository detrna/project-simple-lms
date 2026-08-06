package domain

import (
	"time"

	"github.com/google/uuid"
)

type Assignment struct {
	ID          uuid.UUID
	CourseID    uuid.UUID
	Title       string
	Description string
	Deadline    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Teachers []MaskedUser
	Files    []File
}
