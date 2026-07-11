package infrastructure

import (
	"context"
	"crypto/rand"
	"main/internal/domain"
	"main/internal/pkg"
	"math/big"
	"time"

	"github.com/resend/resend-go/v3"
)

type ResendClient struct {
	redis pkg.RedisClient
}

func NewResendClient(redis pkg.RedisClient) pkg.ResendClient {
	return &ResendClient{redis: redis}
}

func (r ResendClient) SendRecoveryOTP(ctx context.Context, account domain.User) error {
	client := resend.NewClient("apiKey")

	otp, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	otpHtml := "<strong>Your OTP code: " + otp.String() + "</strong>"

	params := &resend.SendEmailRequest{
		From:    "Acme <project-simple-lms>",
		To:      []string{account.Email},
		Subject: "Recover your account " + account.Name,
		Html:    otpHtml,
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		return err
	}

	r.redis.Set(ctx, "otp:"+account.Email, otp.String(), 15*time.Minute)

	return nil
}
