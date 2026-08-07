package class

import (
	"context"
	"main/internal/domain"
	classdomain "main/internal/modules/class/domain"

	"github.com/google/uuid"
)

type IRepository interface {
	GetStudents(ctx context.Context, classID uuid.UUID) ([]*domain.User, error)
	GetClassByID(ctx context.Context, classID uuid.UUID) (*classdomain.Class, error)
	CreateClass(ctx context.Context, data *classdomain.Class) (*classdomain.Class, error)
	UpdateClass(ctx context.Context, data *classdomain.Class) (*classdomain.Class, error)
	DeleteClass(ctx context.Context, classID uuid.UUID) error
}
