package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var DB *pgx.Conn

func InitializeDB() (*pgx.Conn, error) {
	driver := os.Getenv("DB_DRIVER")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DBNAME")
	url := fmt.Sprintf("%s://%s:%s@%s:%s/%s", driver, username, password, host, port, dbname)
	var err error
	DB, err = pgx.Connect(context.Background(), url)
	return DB, err
}
