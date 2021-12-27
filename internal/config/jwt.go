package config

import (
	"os"
	"strconv"
)

type JWTConfig struct {
	Secret string
	TTL    int
}

func NewJWTConfig() *JWTConfig {
	ttl, err := strconv.Atoi(os.Getenv("JWT_TTL"))
	if err != nil || ttl <= 0 {
		ttl = 60
	}

	return &JWTConfig{
		Secret: os.Getenv("JWT_SECRET"),
		TTL:    ttl,
	}
}
