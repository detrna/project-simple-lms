package student

import (
	"github.com/google/uuid"
)

type Student struct {
	uuid     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	id       string
	name     string
	password string
}
