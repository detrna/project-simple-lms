package class

import (
	"context"
	"main/internal/domain"
	classdomain "main/internal/modules/class/domain"

	"github.com/google/uuid"
)

type UseCase struct {
	repo IRepository
}

type IUseCase interface {
	GetStudents(ctx context.Context, classID uuid.UUID) ([]*domain.User, error)
	GetClassByID(ctx context.Context, classID uuid.UUID) (*classdomain.Class, error)
	CreateClass(ctx context.Context, data *CreateClassRequest) (*classdomain.Class, error)
	UpdateClass(ctx context.Context, data *UpdateClassRequest) (*classdomain.Class, error)
	DeleteClass(ctx context.Context, classID uuid.UUID) error
}

func NewUseCase(repo IRepository) *UseCase {
	return (&UseCase{repo: repo})
}

func (usecase UseCase) GetStudents(ctx context.Context, classID uuid.UUID) ([]*domain.User, error) {
	result, err := usecase.repo.GetStudents(ctx, classID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (usecase UseCase) GetClassByID(ctx context.Context, classID uuid.UUID) (*classdomain.Class, error)
func (usecase UseCase) CreateClass(ctx context.Context, data *CreateClassRequest) (*classdomain.Class, error)
func (usecase UseCase) UpdateClass(ctx context.Context, data *UpdateClassRequest) (*classdomain.Class, error)
func (usecase UseCase) DeleteClass(ctx context.Context, classID uuid.UUID) error
