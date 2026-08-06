package infrastructure

import (
	"context"
	"main/internal/domain"
	"main/internal/pkg"

	"github.com/resend/resend-go/v3"
)

type ResendClient struct {
	redis pkg.RedisClient
}

func (r ResendClient) SendRecoveryOTP(ctx context.Context, account *domain.User, otp string) error {
	client := resend.NewClient("apiKey")

	otpHtml := "<strong>Your OTP code: " + otp + "</strong>"

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

	return nil
}
