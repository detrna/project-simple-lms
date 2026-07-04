package class

import (
	"context"
	"main/internal/modules/user"

	"github.com/google/uuid"
)

type UseCase struct {
	repo IRepository
}

type IUseCase interface {
	GetStudents(ctx context.Context, classID uuid.UUID) ([]user.User, error)
	GetMyClasses(ctx context.Context, userID uuid.UUID) ([]Class, error)
}

func NewUseCase(repo IRepository) *UseCase {
	return (&UseCase{repo: repo})
}

func (usecase UseCase) GetStudents(ctx context.Context, classID uuid.UUID) ([]user.User, error) {
	result, err := usecase.repo.GetStudents(ctx, classID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (usecase UseCase) GetMyClasses(ctx context.Context, userID uuid.UUID) ([]Class, error) {
	result, err := usecase.repo.GetMyClasses(ctx, userID)

	if err != nil {
		return nil, err
	}

	return result, nil
}
