package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func Initialize() error {
	ex, err := os.Executable()
	if err != nil {
		return err
	}
	exPath := filepath.Dir(ex)

	err = godotenv.Load(exPath + string(os.PathSeparator) + "/.env")
	if err != nil {
		return err
	}

	return nil
}
