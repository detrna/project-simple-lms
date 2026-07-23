package infrastructure

import (
	"errors"
	"main/internal/config"
	"main/internal/pkg"
	"main/internal/shared"

	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct {
	salt int
	cost int
}

func NewBcryptHasher(cfg config.BcryptConfig) pkg.BcryptHasher {
	return &BcryptHasher{salt: cfg.Salt, cost: cfg.Cost}
}

func (b BcryptHasher) Compare(hashed string, literal string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(literal))

	if errors.Is(bcrypt.ErrMismatchedHashAndPassword, err) {
		return shared.ErrCredentialsIncorrect
	}

	return err
}

func (b BcryptHasher) Hash(value string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(value), b.cost)

	if err != nil {
		return nil, err
	}

	return hashed, nil
}
