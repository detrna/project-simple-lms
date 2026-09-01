package usecase

import (
	"context"

	"github.com/google/uuid"
)

func (useCase ClassUseCase) Delete(ctx context.Context, classID uuid.UUID) error {
	return nil
}

