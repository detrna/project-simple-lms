package usecase

import (
	"context"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/dto"
)

func (useCase ClassUseCase) Create(ctx context.Context, data *dto.CreateClassRequest) (*domain.Class, error) {
	return nil, nil
}

