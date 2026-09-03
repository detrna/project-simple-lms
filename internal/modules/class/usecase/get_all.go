package usecase

import (
	"context"
	"main/internal/modules/class/domain"
	"main/internal/shared/pagination"
)

func (uc ClassUseCase) GetAll(ctx context.Context, pagination *pagination.PaginationInput) (*[]domain.Class, int, error) {
	classes, total, err := uc.repo.GetAll(ctx, *pagination)

	if err != nil {
		return nil, 0, err
	}

	return classes, total, nil
}
