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

func NewClassContainer(repo usecase.ClassRepositoryI, logger pkg.Logger) *ClassContainer {
	uc := usecase.NewClassUseCase(repo)
	controller := controller.NewClassController(uc, logger)
	routes := routes.NewClassRoutes(controller, logger)

	return &ClassContainer{
		Repo:       repo,
		UseCase:    uc,
		Controller: controller,
		Routes:     routes,
	}
}
