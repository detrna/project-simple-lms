package config

type ServerConfig struct {
	Port string
}

func LoadServerConfig() *ServerConfig {
	serverPort := GetEnv("APP_PORT", "8080")

	return &ServerConfig{Port: serverPort}
}
