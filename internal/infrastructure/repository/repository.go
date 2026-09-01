package repository

import (
	"main/internal/modules/auth"
	"main/internal/modules/user"
	"main/internal/pkg"

	"gorm.io/gorm"

	classrepository "main/internal/modules/class/repository"
	classusecase "main/internal/modules/class/usecase"
)

type Repository struct {
	AuthRepository  auth.IRepository
	UserRepository  user.IRepository
	ClassRepository classusecase.ClassRepositoryI
}

func NewRepository(db *gorm.DB, logger pkg.Logger) *Repository {
	return &Repository{
		AuthRepository:  NewAuthRepository(db, logger),
		UserRepository:  NewUserRepository(db),
		ClassRepository: classrepository.NewClassRepository(db),
	}
}
