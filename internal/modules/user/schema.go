package user

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	SystemID  string    `json:"systemId"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateUserSchema struct {
	SystemID string `json:"systemId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type UpdateUserBodySchema struct {
	SystemID *string `json:"systemId"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Role     *string `json:"role"`
	Password *string `json:"password"`
}

type UpdateUserSchema struct {
	ID       *uuid.UUID `json:"id"`
	SystemID *string    `json:"systemId"`
	Name     *string    `json:"name"`
	Email    *string    `json:"email"`
	Role     *string    `json:"role"`
	Password *string    `json:"password"`
}
