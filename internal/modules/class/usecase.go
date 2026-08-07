package class

import (
	"context"
	classdomain "main/internal/modules/class/domain"

	"github.com/google/uuid"
)

type UseCase struct {
	repo IRepository
}

type IUseCase interface {
	GetClasses(ctx context.Context) ([]*classdomain.Class, error)
	GetClassByID(ctx context.Context, classID uuid.UUID) (*classdomain.Class, error)
	GetClassBySystemID(ctx context.Context, systemID string) (*classdomain.Class, error)
	CreateClass(ctx context.Context, data *CreateClassRequest) (*classdomain.Class, error)
	UpdateClass(ctx context.Context, data *UpdateClassDTO) (*classdomain.Class, error)
	DeleteClass(ctx context.Context, classID uuid.UUID) error
}

func NewUseCase(repo IRepository) *UseCase {
	return (&UseCase{repo: repo})
}

func (usecase UseCase) GetClasses(ctx context.Context) ([]*classdomain.Class, error) {
	return nil, nil
}
func (usecase UseCase) GetClassByID(ctx context.Context, classID uuid.UUID) (*classdomain.Class, error) {
	return nil, nil
}
func (usecase UseCase) GetClassBySystemID(ctx context.Context, systemID string) (*classdomain.Class, error) {
	return nil, nil
}
func (usecase UseCase) CreateClass(ctx context.Context, data *CreateClassRequest) (*classdomain.Class, error) {
	return nil, nil
}
func (usecase UseCase) UpdateClass(ctx context.Context, data *UpdateClassDTO) (*classdomain.Class, error) {
	return nil, nil
}
func (usecase UseCase) DeleteClass(ctx context.Context, classID uuid.UUID) error {
	return nil
}
