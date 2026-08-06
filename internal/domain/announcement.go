package domain

import (
	"time"

	"github.com/google/uuid"
)

type Announcement struct {
	ID        uuid.UUID
	Title     string
	Content   string
	CourseID  uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time

	Teacher MaskedUser
	Files   []File
}
