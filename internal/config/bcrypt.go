package config

import "strconv"

type BcryptConfig struct {
	Cost int
	Salt int
}

func LoadBcryptConfig() *BcryptConfig {
	cost, _ := strconv.Atoi(GetEnv("BCRYPT_COST", "10"))
	salt, _ := strconv.Atoi(GetEnv("SALT", "10"))

	return &BcryptConfig{Cost: cost, Salt: salt}
}
