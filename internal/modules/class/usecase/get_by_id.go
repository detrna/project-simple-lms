package usecase

import (
	"context"
	"main/internal/modules/class/domain"

	"github.com/google/uuid"
)

func (useCase ClassUseCase) GetByID(ctx context.Context, classID uuid.UUID) (*domain.Class, error) {
	return nil, nil
}

