package controller

import (
	"context"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/dto"

	"github.com/google/uuid"
)

type ClassUseCaseI interface {
	GetAll(ctx context.Context) ([]*domain.Class, error)
	GetByID(ctx context.Context, classID uuid.UUID) (*domain.Class, error)
	GetBySystemID(ctx context.Context, systemID string) (*domain.Class, error)
	Create(ctx context.Context, data *dto.CreateClassRequest) (*domain.Class, error)
	Update(ctx context.Context, data *dto.UpdateClassRequest) (*domain.Class, error)
	Delete(ctx context.Context, classID uuid.UUID) error
}

type ClassController struct {
	useCase ClassUseCaseI
}

func NewClassController(useCase ClassUseCaseI) *ClassController {
	return &ClassController{useCase: useCase}
}

