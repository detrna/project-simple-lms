package usecase

import (
	"context"
	"main/internal/modules/class/domain"

	"github.com/google/uuid"
)

func (uc ClassUseCase) Delete(ctx context.Context, classID uuid.UUID) error {
	_, err := uc.repo.GetByID(ctx, classID)

	if (err != nil && err != domain.ErrClassNotFound) || err == domain.ErrClassNotFound {
		return err
	}

	return uc.repo.Delete(ctx, classID)
}
