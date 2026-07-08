package config

import "strconv"

type BcryptConfig struct {
	Cost int
}

func LoadBcryptConfig() *BcryptConfig {
	cost, _ := strconv.Atoi(GetEnv("BCRYPT_COST", "10"))

	return &BcryptConfig{Cost: cost}
}
