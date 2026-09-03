package usecase

import (
	"context"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/dto"

	"github.com/google/uuid"
)

func (uc ClassUseCase) Create(ctx context.Context, data *dto.CreateClassRequest) (*domain.Class, error) {
	existingClass, err := uc.repo.GetBySystemID(ctx, data.SystemID)

	if existingClass != nil {
		return nil, domain.ErrClassSystemIDTaken
	}

	if err != nil && err != domain.ErrClassNotFound {
		return nil, err
	}

	newClass := domain.Class{
		ID:       uuid.New(),
		SystemID: data.SystemID,
		Name:     data.Name,
	}

	result, err := uc.repo.Create(ctx, &newClass)

	if err != nil {
		return nil, err
	}

	return result, nil
}
