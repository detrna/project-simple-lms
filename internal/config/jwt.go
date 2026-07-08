package config

import "strconv"

type JWTConfig struct {
	AccessSecret         string
	RefreshSecret        string
	AccessExpiryMinutes  int
	RefreshExpiryMinutes int
}

func LoadJWTConfig() *JWTConfig {
	accessSecret := GetEnv("JWT_ACCESS_SECRET", "your-secret-key")
	refreshSecret := GetEnv("JWT_REFRESH_SECRET", "your-secret-key")
	accessExpiry, _ := strconv.Atoi(GetEnv("JWT_ACCESS_EXPIRY_MINUTES", "60"))
	refreshExpiry, _ := strconv.Atoi(GetEnv("JWT_REFRESH_EXPIRY_DAYS", "7"))

	return &JWTConfig{
		AccessSecret:         accessSecret,
		RefreshSecret:        refreshSecret,
		AccessExpiryMinutes:  accessExpiry,
		RefreshExpiryMinutes: refreshExpiry,
	}
}
