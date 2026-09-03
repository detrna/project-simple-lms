package helper

import "net/http"

func NewRefreshTokenCookie(value string) *http.Cookie {
	return &http.Cookie{
		Name:  "refresh_token",
		Value: value,
	}
}
