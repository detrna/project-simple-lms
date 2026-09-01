package usecase

import (
	"context"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/dto"
)

func (useCase ClassUseCase) Update(ctx context.Context, data *dto.UpdateClassRequest) (*domain.Class, error) {
	return nil, nil
}

