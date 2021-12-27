package db

import (
	"context"
	"fmt"

	"github.com/dshum/school/internal/config"
	"github.com/jackc/pgx/v4"
)

func NewConnection() (*pgx.Conn, error) {
	url := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		config.DB.Driver,
		config.DB.UserName,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.DbName)

	return pgx.Connect(context.Background(), url)
}
