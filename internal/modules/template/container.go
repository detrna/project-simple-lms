package template

import (
	"gorm.io/gorm"
)

type Container struct {
	Repo       IRepository
	UseCase    IUseCase
	Controller IController
}

func NewContainer(db *gorm.DB) *Container {
	repo := NewRepository(db)
	usecase := NewUseCase(repo)
	controller := NewController(usecase)

	return &Container{
		Repo:       repo,
		UseCase:    usecase,
		Controller: controller,
	}
}
