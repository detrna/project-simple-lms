package controller

import (
	"context"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/dto"
	"main/internal/pkg"
	"main/internal/shared/pagination"

	"github.com/google/uuid"
)

type ClassUseCaseI interface {
	GetAll(ctx context.Context, pagination *pagination.Pagination) (*[]domain.Class, int, error)
	GetByID(ctx context.Context, classID uuid.UUID) (*domain.Class, error)
	GetBySystemID(ctx context.Context, systemID string) (*domain.Class, error)
	Create(ctx context.Context, data *dto.CreateClassRequest) (*domain.Class, error)
	Update(ctx context.Context, data *dto.UpdateClassRequest) (*domain.Class, error)
	Delete(ctx context.Context, classID uuid.UUID) error
}

type ClassController struct {
	uc     ClassUseCaseI
	logger pkg.Logger
}

func NewClassController(uc ClassUseCaseI, logger pkg.Logger) *ClassController {
	return &ClassController{uc: uc, logger: logger}
}
