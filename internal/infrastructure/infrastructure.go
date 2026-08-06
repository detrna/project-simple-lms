package infrastructure

import (
	"main/internal/config"
	"main/internal/infrastructure/database"
	hasherpkg "main/internal/infrastructure/hasher"
	loggerpkg "main/internal/infrastructure/logger"
	mailpkg "main/internal/infrastructure/mailer"
	objstoragepkg "main/internal/infrastructure/object_storage"
	redispkg "main/internal/infrastructure/redis"
	"main/internal/infrastructure/repository"
	tokenpkg "main/internal/infrastructure/token_service"
	"main/internal/pkg"

	"gorm.io/gorm"
)

func Initialize(cfg *config.Config) (*pkg.Packages, *gorm.DB, *repository.Repository, error) {
	db := database.Load(cfg.Database)
	logger := loggerpkg.NewLogger(cfg.Logger)
	redis := redispkg.RedisSetup(cfg.Redis)
	gomail, _ := mailpkg.NewGoMailer(cfg.Mail, cfg.App)
	jwtService := tokenpkg.NewTokenService(cfg.JWT)
	bcrypt := hasherpkg.NewBcryptHasher(cfg.Bcrypt)
	objStorage, _ := objstoragepkg.SetupMinIO(cfg.ObjectStorage)
	repository := repository.NewRepository(db, logger)

	return &pkg.Packages{
		Logger:        logger,
		RedisClient:   redis,
		Mailer:        gomail,
		TokenService:  jwtService,
		Hasher:        bcrypt,
		ObjectStorage: objStorage,
	}, db, repository, nil

}
