package container

import (
	"main/internal/config"
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

func NewAuthContainer(cfg *config.Config, infra *pkg.Packages, repo *repository.Repository) *AuthContainer {
	authRepo := repo.AuthRepository
	userRepo := repo.UserRepository

	useCasePacakges := auth.UseCasePackages{
		Bcrypt:        infra.BcryptHasher,
		Mailer:        infra.ResendClient,
		TokenProvider: infra.JWTProvider,
		Redis:         infra.RedisClient,
		Logger:        infra.Logger,
	}

	usecase := auth.NewUseCase(authRepo, userRepo, &useCasePacakges, cfg.Mail)
	controller := auth.NewController(usecase, infra.Logger, infra.JWTProvider, cfg.App.Mode == "PRODUCTION")
	routes := auth.NewRoutes(controller, infra.JWTProvider, infra.Logger)

	return &AuthContainer{
		UseCase:    usecase,
		Controller: controller,
		Routes:     routes,
		Repo:       authRepo,
		UserRepo:   userRepo,
	}
}
