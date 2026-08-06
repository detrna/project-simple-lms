package domain

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	ParentType string
	ParentID   uuid.UUID
	Content    string
	Replies    []Comment
	CreatedAt  time.Time
}
