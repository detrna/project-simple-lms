package template

import "gorm.io/gorm"

type IRepository interface {
	Ping() (string, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (repo Repository) Ping() (string, error) {
	return "pong", nil
}
