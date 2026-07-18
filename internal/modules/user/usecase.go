package user

import (
	"context"
	"main/internal/domain"
	"main/internal/pkg"
	"main/internal/shared"

	"github.com/google/uuid"
)

type UseCase struct {
	repo   IRepository
	bcrypt pkg.BcryptHasher
	logger pkg.Logger
}

func NewUseCase(repo IRepository, bcrypt pkg.BcryptHasher, logger pkg.Logger) *UseCase {
	return (&UseCase{repo: repo, bcrypt: bcrypt, logger: logger})
}

type IUseCase interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*UserResponse, error)
	GetUserByUserID(ctx context.Context, id uuid.UUID) (*UserResponse, error)
	CreateUser(ctx context.Context, data CreateUserSchema) (*UserResponse, error)
	UpdateUser(ctx context.Context, data UpdateUserSchema) (*UserResponse, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

func OmitPassword(u *domain.User) UserResponse {
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
	dbAccount, err := usecase.repo.FindByEmail(ctx, data.Email)

	if err != shared.ErrRecordNotFound {
		return nil, err
	}

	if dbAccount != nil {
		return nil, shared.ErrEmailTaken
	}

	hashedPassword, err := usecase.bcrypt.Hash(data.Password)

	if err != nil {
		return nil, err
	}

	user := domain.User{
		ID:       uuid.New(),
		SystemID: data.SystemID,
		Name:     data.Name,
		Email:    data.Email,
		Role:     data.Role,
		Password: string(hashedPassword),
	}

	result, err := usecase.repo.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	dto := OmitPassword(result)

	return &dto, nil
}

func (usecase UseCase) UpdateUser(ctx context.Context, data UpdateUserSchema) (*UserResponse, error) {
	var user domain.User

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
