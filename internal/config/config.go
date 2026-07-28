package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   *ServerConfig
	Database *DatabaseConfig
	Logger   *LoggerConfig
	JWT      *JWTConfig
	Bcrypt   *BcryptConfig
	Redis    *RedisConfig
	Mail     *MailConfig
	App      *AppConfig
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func Load() (*Config, error) {
	if err := LoadEnv(); err != nil {
		return nil, err
	}

	return &Config{
		Server:   LoadServerConfig(),
		Database: LoadDatabaseConfig(),
		Logger:   LoadLoggerConfig(),
		JWT:      LoadJWTConfig(),
		Bcrypt:   LoadBcryptConfig(),
		Redis:    LoadRedisConfig(),
		App:      LoadAppConfig(),
		Mail:     LoadMailConfig(),
	}, nil
}

func LoadEnv() error {
	root, err := findProjectRoot()
	if err != nil {
		return err
	}

	envPath := filepath.Join(root, ".env")
	return godotenv.Load(envPath)
}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found")
		}
		dir = parent
	}
}
