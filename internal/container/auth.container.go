package container

import (
	"main/internal/infrastructure"
	"main/internal/modules/auth"
	"main/internal/modules/user"
)

type AuthContainer struct {
	UseCase    auth.IUseCase
	Controller auth.IController
	Routes     *auth.Routes
	Repo       auth.IRepository
	UserRepo   user.IRepository
}

func NewAuthContainer(infra *infrastructure.Infrastructure, userRepo user.IRepository) *AuthContainer {
	repo := auth.NewRepository(infra.DB, infra.Logger)
	usecase := auth.NewUseCase(repo, userRepo, infra.Logger)
	controller := auth.NewController(usecase, infra.Logger)
	routes := auth.NewRoutes(controller, *infra.Config)

	return &AuthContainer{
		UseCase:    usecase,
		Controller: controller,
		Routes:     routes,
		Repo:       repo,
		UserRepo:   userRepo,
	}
}
