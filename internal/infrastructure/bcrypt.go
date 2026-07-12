package infrastructure

import (
	"main/internal/config"
	"main/internal/pkg"

	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct {
	salt int
	cost int
}

func NewBcryptHasher(cfg config.BcryptConfig) pkg.BcryptHasher {
	return &BcryptHasher{salt: cfg.Salt, cost: cfg.Cost}
}

func (b BcryptHasher) CompareHashAndPassword(hashed string, literal string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(literal))

	return err
}
