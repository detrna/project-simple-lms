package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"
)

func ToDomainUser(u database.User) domain.User {
	return domain.User{
		ID:        u.ID,
		SystemID:  u.SystemID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}

func ToDatabaseUser(u domain.User) database.User {
	return database.User{
		ID:        u.ID,
		SystemID:  u.SystemID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}

func ToDatabaseJWT(JWT domain.JWTPayload, token string) *database.JWT {
	return &database.JWT{
		ID:     JWT.JTI,
		UserID: JWT.UserID,
		Token:  token,
	}
}
