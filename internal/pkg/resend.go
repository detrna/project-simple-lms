package pkg

import (
	"context"
	"main/internal/domain"
)

type ResendClient interface {
	SendRecoveryOTP(ctx context.Context, account *domain.User, otp string) error
}
