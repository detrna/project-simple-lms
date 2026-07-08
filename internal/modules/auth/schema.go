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

type RecoverSchema struct {
	Email string `json:"email"`
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
