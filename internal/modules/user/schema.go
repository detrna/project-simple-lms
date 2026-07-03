package user

import "github.com/google/uuid"

type GetUserResponse struct {
	ID       uuid.UUID `json:"id"`
	UserID   string    `json:"userId"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type CreateUserSchema struct {
	UserID   string `json:"userId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
