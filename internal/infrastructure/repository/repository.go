package repository

import (
	"main/internal/modules/auth"
	"main/internal/modules/user"
	"main/internal/pkg"

	"gorm.io/gorm"
)

type Repository struct {
	AuthRepository auth.IRepository
	UserRepository user.IRepository
}

func NewRepository(db *gorm.DB, logger pkg.Logger) *Repository {
	return &Repository{
		AuthRepository: NewAuthRepository(db, logger),
		UserRepository: NewUserRepository(db),
	}
}
