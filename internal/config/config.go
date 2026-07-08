package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Logger   LoggerConfig
	JWT      JWTConfig
	Bcrypt   BcryptConfig
	App      AppConfig
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		Server:   *LoadServerConfig(),
		Database: *LoadDatabaseConfig(),
		Logger:   *LoadLoggerConfig(),
		JWT:      *LoadJWTConfig(),
		Bcrypt:   *LoadBcryptConfig(),
		App:      *LoadAppConfig(),
	}, nil
}
