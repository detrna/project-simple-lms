package auth

type RegisterSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Confirm  string `json:"confirm"`
}

type LoginSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type RecoverSchema struct {
	Email string `json:"email"`
}

type VerifyRecoverySchema struct {
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
	OTP         string `json:"otp"`
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
