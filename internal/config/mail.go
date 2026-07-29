package config

import "strconv"

type MailConfig struct {
	OTPExpiryMin int
	Port         int
	Host         string
	Username     string
	Password     string
	From         string
}

func LoadMailConfig() *MailConfig {
	otpExpiry, _ := strconv.Atoi(GetEnv("MAIL_OTP_EXPIRY_MIN", "15"))
	port, _ := strconv.Atoi(GetEnv("SMTP_PORT", "1025"))
	host := GetEnv("SMTP_HOST", "mailpit")
	username := GetEnv("SMTP_USERNAME", "")
	password := GetEnv("SMTP_PASSWORD", "")
	from := GetEnv("SMTP_FROM", "noreply@simple-lms.local")

	return &MailConfig{
		OTPExpiryMin: otpExpiry,
		Port:         port,
		Host:         host,
		Username:     username,
		Password:     password,
		From:         from,
	}
}
