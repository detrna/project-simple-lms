package container

import (
	"main/internal/modules/class"
	"main/internal/pkg"
)

type ClassContainer struct {
	Repo       class.IRepository
	UseCase    class.IUseCase
	Controller class.IController
	Routes     *class.Routes
}

func NewClassContainer(infra *pkg.Packages, repo class.IRepository) *ClassContainer {
	usecase := class.NewUseCase(repo)
	controller := class.NewController(usecase)
	routes := class.NewRoutes(controller, infra.Logger)

	return &ClassContainer{
		Repo:       repo,
		UseCase:    usecase,
		Controller: controller,
		Routes:     routes,
	}
}
