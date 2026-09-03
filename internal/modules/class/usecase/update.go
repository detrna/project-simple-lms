package usecase

import (
	"context"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/dto"
)

func (uc ClassUseCase) Update(ctx context.Context, data *dto.UpdateClassRequest) (*domain.Class, error) {
	existingClass, err := uc.repo.GetByID(ctx, data.ID)

	if (err != nil && err != domain.ErrClassNotFound) || err == domain.ErrClassNotFound {
		return nil, err
	}

	if data.SystemID != nil {
		takenSystemID, err := uc.repo.GetBySystemID(ctx, *data.SystemID)

		if takenSystemID != nil {
			return nil, domain.ErrClassSystemIDTaken
		}

		if err != nil && err != domain.ErrClassNotFound {
			return nil, err
		}

		existingClass.SystemID = *data.SystemID
	}

	if data.Name != nil {
		existingClass.Name = *data.Name
	}

	class, err := uc.repo.Update(ctx, existingClass)

	if err != nil {
		return nil, err
	}

	return class, nil
}
