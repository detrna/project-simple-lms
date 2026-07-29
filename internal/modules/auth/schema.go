package auth

type RegisterSchema struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=32"`
	Confirm  string `json:"confirm" binding:"required,min=8,max=32"`
}

type LoginSchema struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=32"`
}

type TokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type RecoverSchema struct {
	Email string `json:"email" binding:"required,email"`
}

type VerifyRecoverySchema struct {
	Email       string `json:"email" binding:"required,email"`
	NewPassword string `json:"newPassword" binding:"required,min=8,max=32"`
	OTP         string `json:"otp" binding:"required,len=8,numeric"`
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
