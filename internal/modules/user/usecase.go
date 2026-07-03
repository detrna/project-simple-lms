package user

import (
	"context"

	"github.com/google/uuid"
)

type UseCase struct {
	repo IRepository
}

func NewUseCase(repo IRepository) *UseCase {
	return (&UseCase{repo: repo})
}

type IUseCase interface {
	GetUserById(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByUserId(ctx context.Context, id uuid.UUID) (*User, error)
	CreateUser(ctx context.Context, data User) error
}

func (usecase UseCase) GetUserById(ctx context.Context, id uuid.UUID) (*User, error) {
	result, err := usecase.repo.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (usecase UseCase) GetUserByUserId(ctx context.Context, id uuid.UUID) (*User, error) {
	result, err := usecase.repo.FindByUserId(ctx, id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (usecase UseCase) CreateUser(ctx context.Context, data CreateUserSchema) error {

	user := User{
		ID:       uuid.New(),
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}

	return usecase.repo.Create(ctx, user)
}
