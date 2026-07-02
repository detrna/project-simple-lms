package instructor

import (
	"time"

	"github.com/google/uuid"
)

type Instructor struct {
	UUID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ID        string
	Name      string
	CreatedAt time.Time `gorm:"default.CURRENT_TIMESTAMP"`
}
