package config

import (
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

type Config struct {
	DB DBConfig
}

func LoadEnv() error {
	if err := godotenv.Load(projectRootPath + string(os.PathSeparator) + ".env"); err != nil {
		return err
	}

	return nil
}

func NewConfig() *Config {
	return &Config{
		DB: *NewDBConfig(),
	}
}
