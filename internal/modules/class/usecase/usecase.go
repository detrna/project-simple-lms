package usecase

import (
	"context"
	"main/internal/modules/class/domain"
	"main/internal/shared/pagination"

	"github.com/google/uuid"
)

type ClassRepositoryI interface {
	GetAll(ctx context.Context, pagination pagination.Pagination) (*[]domain.Class, int, error)
	GetByID(ctx context.Context, classID uuid.UUID) (*domain.Class, error)
	GetBySystemID(ctx context.Context, systemID string) (*domain.Class, error)
	Create(ctx context.Context, data *domain.Class) (*domain.Class, error)
	Update(ctx context.Context, data *domain.Class) (*domain.Class, error)
	Delete(ctx context.Context, classID uuid.UUID) error
}

type ClassUseCase struct {
	repo ClassRepositoryI
}

func NewClassUseCase(repo ClassRepositoryI) *ClassUseCase {
	return (&ClassUseCase{repo: repo})
}
