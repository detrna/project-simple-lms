package container

import (
	"main/internal/modules/class/controller"
	"main/internal/modules/class/routes"
	"main/internal/modules/class/usecase"
	"main/internal/pkg"
)

type ClassContainer struct {
	Repo       usecase.ClassRepositoryI
	UseCase    controller.ClassUseCaseI
	Controller routes.IController
	Routes     *routes.ClassRoutes
}

func NewClassContainer(infra *pkg.Packages, repo usecase.ClassRepositoryI) *ClassContainer {
	uc := usecase.NewClassUseCase(repo)
	controller := controller.NewClassController(uc)
	routes := routes.NewClassRoutes(controller, infra.Logger)

	return &ClassContainer{
		Repo:       repo,
		UseCase:    uc,
		Controller: controller,
		Routes:     routes,
	}
}
