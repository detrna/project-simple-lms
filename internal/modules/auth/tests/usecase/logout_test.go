package auth_usecase_test

import (
	"context"
	"main/internal/modules/auth"
	auth_mocks "main/internal/modules/auth/mocks"
	user_mocks "main/internal/modules/user/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// Logout(ctx context.Context, id uuid.UUID) error

func TestLogout_Success(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()

	repo := auth_mocks.NewMockIRepository(t)

	repo.EXPECT().
		DeleteJWT(ctx, id).
		Return(nil)

	usecase := auth.NewUseCase(repo, user_mocks.NewMockIRepository(t), &auth.UseCasePackages{})

	err := usecase.Logout(ctx, id)

	require.NoError(t, err)
}
