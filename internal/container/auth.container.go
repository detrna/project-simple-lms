package container

import (
	"main/internal/infrastructure/repository"
	"main/internal/modules/auth"
	"main/internal/modules/user"
	"main/internal/pkg"
)

type AuthContainer struct {
	UseCase    auth.IUseCase
	Controller auth.IController
	Routes     *auth.Routes
	Repo       auth.IRepository
	UserRepo   user.IRepository
}

func NewAuthContainer(infra pkg.Packages, repo repository.Repository) *AuthContainer {
	authRepo := repo.AuthRepository
	userRepo := repo.UserRepository

	usecase := auth.NewUseCase(authRepo, userRepo, infra)
	controller := auth.NewController(usecase, infra.Logger)
	routes := auth.NewRoutes(controller, infra.JWTProvider)

	return &AuthContainer{
		UseCase:    usecase,
		Controller: controller,
		Routes:     routes,
		Repo:       authRepo,
		UserRepo:   userRepo,
	}
}
