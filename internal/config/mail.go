package config

import "strconv"

type MailConfig struct {
	OTPExpiryMin int
}

func LoadMailConfig() *MailConfig {
	otpExpiry, _ := strconv.Atoi(GetEnv("MAIL_OTP_EXPIRY_MIN", "15"))

	return &MailConfig{OTPExpiryMin: otpExpiry}
}
