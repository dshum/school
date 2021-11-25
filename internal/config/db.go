package config

import "os"

type DBConfig struct {
	Driver   string
	UserName string
	Password string
	Host     string
	Port     string
	DbName   string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Driver:   os.Getenv("DB_DRIVER"),
		UserName: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_DBNAME"),
	}
}
