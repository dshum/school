package config

import "os"

type RedisConfig struct {
	DSN string
	DB  string
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		DSN: os.Getenv("REDIS_DSN"),
		DB:  os.Getenv("REDIS_DB"),
	}
}
