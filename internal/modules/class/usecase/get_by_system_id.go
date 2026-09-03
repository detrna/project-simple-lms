package usecase

import (
	"context"
	"main/internal/modules/class/domain"
)

func (uc ClassUseCase) GetBySystemID(ctx context.Context, systemID string) (*domain.Class, error) {
	class, err := uc.repo.GetBySystemID(ctx, systemID)

	if err != nil {
		return nil, err
	}

	return class, nil
}
