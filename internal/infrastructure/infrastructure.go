package infrastructure

import (
	"main/internal/config"
	"main/internal/infrastructure/database"
	"main/internal/infrastructure/repository"
	"main/internal/pkg"

	"gorm.io/gorm"
)

func Initialize(cfg config.Config) (*pkg.Packages, *gorm.DB, *repository.Repository, error) {
	db := database.Load(cfg.Database)
	logger := NewLogger(cfg.Logger)
	redis := RedisSetup(cfg.Redis)
	resend := NewResendClient(redis)
	jwtProvider := NewTokenProvider(cfg.JWT)
	bcrypt := NewBcryptHasher(cfg.Bcrypt)
	repository := repository.NewRepository(db, logger)

	return &pkg.Packages{
		Logger:       logger,
		RedisClient:  redis,
		ResendClient: resend,
		JWTProvider:  jwtProvider,
		BcryptHasher: bcrypt,
	}, db, repository, nil

}
