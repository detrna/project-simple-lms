package mapper

import (
	"main/internal/domain"
	"main/internal/infrastructure/database"
)

func ToDomainUser(u *database.User) *domain.User {
	if u == nil {
		return nil
	}

	return &domain.User{
		ID:        u.ID,
		SystemID:  u.SystemID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToDomainMaskedUser(u *database.User) domain.MaskedUser {
	if u == nil {
		return domain.MaskedUser{}
	}

	return domain.MaskedUser{
		ID:        u.ID,
		SystemID:  u.SystemID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToDatabaseUser(u *domain.User) *database.User {
	if u == nil {
		return nil
	}

	return &database.User{
		ID:        u.ID,
		SystemID:  u.SystemID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToDatabaseMaskedUser(u *domain.MaskedUser) *database.User {
	if u == nil {
		return nil
	}

	return &database.User{
		ID:        u.ID,
		SystemID:  u.SystemID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToDatabaseJWT(JWT *domain.JWTPayload, token string) *database.JWT {
	return &database.JWT{
		ID:     JWT.JTI,
		UserID: JWT.UserID,
		Token:  token,
	}
}
