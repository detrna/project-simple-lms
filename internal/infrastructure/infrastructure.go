package infrastructure

import (
	"main/internal/config"
	"main/internal/infrastructure/database"
	"main/internal/infrastructure/repository"
	"main/internal/pkg"

	"gorm.io/gorm"
)

func Initialize(cfg *config.Config) (*pkg.Packages, *gorm.DB, *repository.Repository, error) {
	db := database.Load(cfg.Database)
	logger := NewLogger(cfg.Logger)
	redis := RedisSetup(cfg.Redis)
	gomail, _ := NewGoMailer(cfg.Mail)
	jwtService := NewTokenService(cfg.JWT)
	bcrypt := NewBcryptHasher(cfg.Bcrypt)
	repository := repository.NewRepository(db, logger)

	return &pkg.Packages{
		Logger:       logger,
		RedisClient:  redis,
		Mailer:       gomail,
		TokenService: jwtService,
		Hasher:       bcrypt,
	}, db, repository, nil

}
