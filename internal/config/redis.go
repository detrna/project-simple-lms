package config

import "strconv"

type RedisConfig struct {
	Address  string
	Password string
	DB       int
	Protocol int
}

func LoadRedisConfig() *RedisConfig {
	address := GetEnv("REDIS_ADDRESS", "localhost:6379")
	password := GetEnv("REDIS_PASSWORD", "")
	db, _ := strconv.Atoi(GetEnv("REDIS_DB", "0"))
	protocol, _ := strconv.Atoi(GetEnv("REDIS_PROTOCOL", "2"))

	return &RedisConfig{Address: address, Password: password, DB: db, Protocol: protocol}
}
