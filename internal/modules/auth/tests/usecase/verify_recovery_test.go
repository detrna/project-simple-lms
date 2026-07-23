package auth_usecase_test

import (
	"context"
	"main/internal/config"
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

// (usecase UseCase) VerifyRecovery(ctx context.Context, data *VerifyRecoverSchema) (*domain.User, error)

func TestVerifyRecovery_Success(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	existingAccount := user_factory.NewUser(id)

	newPassword := "new-password"
	requestOTP := "correct-otp-code"
	dbOTP := requestOTP

	updatedAccount := *existingAccount
	updatedAccount.Password = newPassword

	repo := auth_mocks.NewMockIRepository(t)

	userRepo := user_mocks.NewMockIRepository(t)
	userRepo.EXPECT().FindByEmail(ctx, mock.AnythingOfType("string")).Return(existingAccount, nil)
	userRepo.EXPECT().Update(ctx, mock.AnythingOfType("*domain.User")).Return(&updatedAccount, nil)

	redis := pkg_mocks.NewMockRedisClient(t)
	redis.EXPECT().Get(ctx, mock.Anything).Return(dbOTP, nil)

	pkg := auth.UseCasePackages{
		Redis: redis,
	}

	requestData := auth.VerifyRecoverSchema{
		Email:       "previous-valid-email",
		NewPassword: newPassword,
		OTP:         requestOTP,
	}

	u := auth.NewUseCase(repo, userRepo, &pkg, &config.MailConfig{})

	err := u.VerifyRecovery(ctx, &requestData)

	require.NoError(t, err)
}

func TestVerifyRecovery_EmailNotFound(t *testing.T) {
	ctx := context.Background()

	newPassword := "new-password"
	requestOTP := "correct-otp-code"

	repo := auth_mocks.NewMockIRepository(t)

	userRepo := user_mocks.NewMockIRepository(t)
	userRepo.EXPECT().FindByEmail(ctx, mock.AnythingOfType("string")).Return(nil, shared.ErrRecordNotFound)

	pkg := auth.UseCasePackages{}

	requestData := auth.VerifyRecoverSchema{
		Email:       "invalid-email",
		NewPassword: newPassword,
		OTP:         requestOTP,
	}

	u := auth.NewUseCase(repo, userRepo, &pkg, &config.MailConfig{})

	err := u.VerifyRecovery(ctx, &requestData)

	require.ErrorIs(t, shared.ErrRecordNotFound, err)
}

func TestVerifyRecovery_IncorrectOTP(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	existingAccount := user_factory.NewUser(id)

	newPassword := "new-password"
	requestOTP := "incorrect-otp-code"
	dbOTP := "correct-otp-code"

	repo := auth_mocks.NewMockIRepository(t)

	userRepo := user_mocks.NewMockIRepository(t)
	userRepo.EXPECT().FindByEmail(ctx, mock.AnythingOfType("string")).Return(existingAccount, nil)

	redis := pkg_mocks.NewMockRedisClient(t)
	redis.EXPECT().Get(ctx, mock.Anything).Return(dbOTP, nil)

	pkg := auth.UseCasePackages{
		Redis: redis,
	}

	requestData := auth.VerifyRecoverSchema{
		Email:       "previous-valid-email",
		NewPassword: newPassword,
		OTP:         requestOTP,
	}

	u := auth.NewUseCase(repo, userRepo, &pkg, &config.MailConfig{})

	err := u.VerifyRecovery(ctx, &requestData)
	require.ErrorIs(t, shared.ErrIncorrectOTP, err)
}

func TestVerifyRecovery_OTPNotFound(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	existingAccount := user_factory.NewUser(id)

	newPassword := "new-password"
	requestOTP := "correct-otp-code"

	repo := auth_mocks.NewMockIRepository(t)

	userRepo := user_mocks.NewMockIRepository(t)
	userRepo.EXPECT().FindByEmail(ctx, mock.AnythingOfType("string")).Return(existingAccount, nil)

	redis := pkg_mocks.NewMockRedisClient(t)
	redis.EXPECT().Get(ctx, mock.Anything).Return("", shared.ErrRedisRecordNotFound)

	pkg := auth.UseCasePackages{
		Redis: redis,
	}

	requestData := auth.VerifyRecoverSchema{
		Email:       "previous-valid-email",
		NewPassword: newPassword,
		OTP:         requestOTP,
	}

	u := auth.NewUseCase(repo, userRepo, &pkg, &config.MailConfig{})

	err := u.VerifyRecovery(ctx, &requestData)

	require.ErrorIs(t, shared.ErrRedisRecordNotFound, err)
}
