package templates

import (
	"bytes"
	"embed"
	"text/template"
)

type ResetPasswordDTO struct {
	Name   string
	OTP    string
	Expiry string
}

//go:embed reset_password.html
var templateFS embed.FS

func VerifyOTP(data ResetPasswordDTO) (string, error) {
	tmpl, err := template.ParseFS(templateFS, "verify_otp.html")
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
