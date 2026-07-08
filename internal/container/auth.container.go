package container

import (
	"main/internal/infrastructure"
	"main/internal/modules/auth"
	"main/internal/modules/user"
)

type AuthContainer struct {
	UseCase    auth.IUseCase
	Controller auth.IController
	Repo       auth.IRepository
	UserRepo   user.IRepository
}

func NewAuthContainer(infra *infrastructure.Infrastructure, userRepo user.IRepository) *AuthContainer {
	repo := auth.NewRepository(infra.DB, infra.Logger)
	usecase := auth.NewUseCase(repo, userRepo, infra.Logger)
	controller := auth.NewController(usecase, infra.Logger)

	return &AuthContainer{
		UseCase:    usecase,
		Controller: controller,
		Repo:       repo,
		UserRepo:   userRepo,
	}
}
