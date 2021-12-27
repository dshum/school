package redis

import (
	"context"
	"github.com/dshum/school/internal/config"
	"github.com/go-redis/redis/v8"
	"strconv"
)

func NewConnection() (*redis.Client, error) {
	db, _ := strconv.Atoi(config.Redis.DB)

	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.DSN,
		Password: "",
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
