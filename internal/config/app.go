package config

type AppConfig struct {
	Mode string
}

func LoadAppConfig() *AppConfig {
	mode := GetEnv("APP_MODE", "DEVELOPMENT")

	return &AppConfig{Mode: mode}
}
