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
	GetUserByID(ctx context.Context, id uuid.UUID) (*UserResponse, error)
	GetUserByUserID(ctx context.Context, id uuid.UUID) (*UserResponse, error)
	CreateUser(ctx context.Context, data CreateUserSchema) (*UserResponse, error)
	UpdateUser(ctx context.Context, data User) (*UserResponse, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

func OmitPassword(u *User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		SystemID:  u.SystemID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}

func (usecase UseCase) GetUserByID(ctx context.Context, id uuid.UUID) (*UserResponse, error) {
	result, err := usecase.repo.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	dto := OmitPassword(result)

	return &dto, nil
}

func (usecase UseCase) GetUserByUserID(ctx context.Context, id uuid.UUID) (*UserResponse, error) {
	result, err := usecase.repo.FindBySystemID(ctx, id)

	if err != nil {
		return nil, err
	}

	dto := OmitPassword(result)

	return &dto, nil
}

func (usecase UseCase) CreateUser(ctx context.Context, data CreateUserSchema) (*UserResponse, error) {
	user := User{
		ID:       uuid.New(),
		SystemID: data.SystemID,
		Name:     data.Name,
		Email:    data.Email,
		Role:     data.Role,
		Password: data.Password,
	}

	result, err := usecase.repo.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	dto := OmitPassword(result)

	return &dto, nil
}

func (usecase UseCase) UpdateUser(ctx context.Context, data UpdateUserSchema) (*UserResponse, error) {
	var user User

	user.ID = *data.ID

	if data.Email != nil {
		user.Email = *data.Email
	}

	if data.Name != nil {
		user.Name = *data.Name
	}

	if data.Password != nil {
		user.Password = *data.Password
	}

	if data.SystemID != nil {
		user.SystemID = *data.SystemID
	}

	if data.Role != nil {
		user.Role = *data.Role
	}

	result, err := usecase.repo.Update(ctx, user)

	if err != nil {
		return nil, err
	}

	dto := OmitPassword(result)

	return &dto, nil
}

func (usecase UseCase) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return usecase.repo.Delete(ctx, id)
}
