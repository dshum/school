package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var (
	// Get current file full path from runtime
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	projectRootPath = filepath.Join(filepath.Dir(b), "../../")
)

var (
	DB    DBConfig
	JWT   JWTConfig
	Redis RedisConfig
)

type Config struct {
	DB    DBConfig
	JWT   JWTConfig
	Redis RedisConfig
}

func init() {
	if err := godotenv.Load(projectRootPath + string(os.PathSeparator) + ".env"); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	c := NewConfig()

	DB = c.DB
	JWT = c.JWT
	Redis = c.Redis
}

func NewConfig() *Config {
	return &Config{
		DB:    *NewDBConfig(),
		JWT:   *NewJWTConfig(),
		Redis: *NewRedisConfig(),
	}
}
