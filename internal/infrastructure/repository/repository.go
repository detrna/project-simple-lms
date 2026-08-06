package repository

import (
	"main/internal/modules/auth"
	"main/internal/modules/class"
	"main/internal/modules/user"
	"main/internal/pkg"

	"gorm.io/gorm"

	class_repository "main/internal/modules/class/repository"
)

type Repository struct {
	AuthRepository  auth.IRepository
	UserRepository  user.IRepository
	ClassRepository class.IRepository
}

func NewRepository(db *gorm.DB, logger pkg.Logger) *Repository {
	return &Repository{
		AuthRepository:  NewAuthRepository(db, logger),
		UserRepository:  NewUserRepository(db),
		ClassRepository: class_repository.NewClassRepository(db),
	}
}
