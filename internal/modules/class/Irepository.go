package class

import (
	"context"
	"main/internal/domain"

	"github.com/google/uuid"
)

type IRepository interface {
	GetStudents(ctx context.Context, classID uuid.UUID) ([]*domain.User, error)
	GetClassByID(ctx context.Context, classID uuid.UUID) (*domain.Class, error)
	CreateClass(ctx context.Context, data *domain.Class) (*domain.Class, error)
	UpdateClass(ctx context.Context, data *domain.Class) (*domain.Class, error)
	DeleteClass(ctx context.Context, classID uuid.UUID) error
}
