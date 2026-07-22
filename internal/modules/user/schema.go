package user

import (
	"main/internal/domain"
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

type AdminUpdateUserSchema struct {
	SystemID *string `json:"systemId"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Role     *string `json:"role"`
	Password *string `json:"password"`
}

type AdminUpdateUserDTO struct {
	ID       uuid.UUID
	SystemID *string
	Name     *string
	Email    *string
	Role     *string
	Password *string
}

type UpdateUserSchema struct {
	Password *string `json:"password"`
}

type UpdateUserDTO struct {
	User     *domain.JWTPayload
	Password *string
}
