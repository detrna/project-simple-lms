package class

import (
	"time"

	"github.com/google/uuid"
)

type Class struct {
	ID        uuid.UUID
	CourseID  uuid.UUID
	Name      string
	CreatedAt time.Time
}
