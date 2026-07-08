// user/container.go
package container

import (
	"main/internal/infrastructure"
	"main/internal/modules/user"
)

type UserContainer struct {
	Repo       user.IRepository
	UseCase    user.IUseCase
	Controller user.IController
}

func NewUserContainer(infra *infrastructure.Infrastructure) *UserContainer {
	repo := user.NewRepository(infra.DB)
	usecase := user.NewUseCase(repo)
	controller := user.NewController(usecase)

	return &UserContainer{
		UseCase:    usecase,
		Controller: controller,
		Repo:       repo,
	}
}
