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

type GetUserByIDSchema struct {
	ID uuid.UUID `uri:"id,parser=encoding.TextUnmarshaler" binding:"required"`
}

type GetUserBySystemIDSchema struct {
	SystemID string `uri:"systemId" binding:"required"`
}

type DeleteUserSchema struct {
	ID uuid.UUID `uri:"id,parser=encoding.TextUnmarshaler" binding:"required"`
}

type CreateUserSchema struct {
	ID       uuid.UUID `uri:"id,parser=encoding.TextUnmarshaler" binding:"required"`
	SystemID string    `json:"systemId" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required,min=8,max=32"`
	Role     string    `json:"role" binding:"required"`
}

type AdminUpdateUserSchema struct {
	ID       uuid.UUID `uri:"id,parser=encoding.TextUnmarshaler" binding:"required"`
	SystemID *string   `json:"systemId"`
	Name     *string   `json:"name"`
	Email    *string   `json:"email" binding:"email"`
	Password *string   `json:"password" binding:"min=8,max=32"`
	Role     *string   `json:"role"`
}

type UpdateUserSchema struct {
	Password string `json:"password" binding:"required,min=8,max=32"`
}

type UpdateUserDTO struct {
	User     *domain.JWTPayload
	Password string
}
