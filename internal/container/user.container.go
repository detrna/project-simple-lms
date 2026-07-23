// user/container.go
package container

import (
	"main/internal/infrastructure/repository"
	"main/internal/modules/user"
	"main/internal/pkg"
)

type UserContainer struct {
	Repo       user.IRepository
	UseCase    user.IUseCase
	Controller user.IController
	Routes     *user.Routes
}

func NewUserContainer(infra *pkg.Packages, repo *repository.Repository) *UserContainer {
	userRepo := repo.UserRepository
	usecase := user.NewUseCase(userRepo, infra.BcryptHasher, infra.Logger)
	controller := user.NewController(usecase, infra.Logger)
	routes := user.NewRoutes(controller, infra.JWTProvider, infra.Logger)

	return &UserContainer{
		UseCase:    usecase,
		Controller: controller,
		Repo:       userRepo,
		Routes:     routes,
	}
}
