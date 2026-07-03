package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	SystemID string
	Name     string
	Email    string
	Password string
	Role     string
}
