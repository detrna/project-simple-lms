package auth_usecase_test

import (
	"context"
	"main/internal/modules/auth"
	auth_mocks "main/internal/modules/auth/mocks"
	user_mocks "main/internal/modules/user/mocks"
	user_factory "main/internal/modules/user/tests"
	pkg_mocks "main/internal/pkg/mocks"
	"main/internal/shared"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// (usecase UseCase) Recover(ctx context.Context, data RecoverSchema) error

func TestRecover_Success(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	existingAccount := user_factory.NewUser(id)

	repo := auth_mocks.NewMockIRepository(t)

	userRepo := user_mocks.NewMockIRepository(t)
	userRepo.EXPECT().FindByEmail(ctx, mock.AnythingOfType("string")).Return(existingAccount, nil)

	redis := pkg_mocks.NewMockRedisClient(t)
	redis.EXPECT().Set(ctx, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	mailer := pkg_mocks.NewMockResendClient(t)
	mailer.EXPECT().SendRecoveryOTP(ctx, mock.Anything).Return(nil)

	pkg := auth.UseCasePackages{
		Redis:  redis,
		Mailer: mailer,
	}

	u := auth.NewUseCase(repo, userRepo, &pkg)

	requestData := auth.RecoverSchema{
		Email: "valid email",
	}

	err := u.Recover(ctx, &requestData)
	require.NoError(t, err)
}

func TestRecover_EmailNotFound(t *testing.T) {
	ctx := context.Background()

	repo := auth_mocks.NewMockIRepository(t)

	userRepo := user_mocks.NewMockIRepository(t)
	userRepo.EXPECT().FindByEmail(ctx, mock.AnythingOfType("string")).Return(nil, shared.ErrRecordNotFound)

	pkg := auth.UseCasePackages{}

	u := auth.NewUseCase(repo, userRepo, &pkg)

	requestData := auth.RecoverSchema{
		Email: "invalid email",
	}

	err := u.Recover(ctx, &requestData)
	require.ErrorIs(t, shared.ErrRecordNotFound, err)
}
