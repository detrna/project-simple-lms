package usecase

import (
	"context"
	"main/internal/modules/class/domain"

	"github.com/google/uuid"
)

func (uc ClassUseCase) GetByID(ctx context.Context, classID uuid.UUID) (*domain.Class, error) {
	class, err := uc.repo.GetByID(ctx, classID)

	if err != nil {
		return nil, err
	}

	return class, nil
}
