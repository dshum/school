package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/dshum/school/internal/user"
	"github.com/dshum/school/internal/utils"
	"github.com/jackc/pgx/v4"
)

type Storage interface {
	Attempt(login, password string) (user.User, error)
}

type authStorage struct {
	db *pgx.Conn
}

func NewStorage(db *pgx.Conn) Storage {
	return &authStorage{
		db: db,
	}
}

func (s *authStorage) Attempt(login, password string) (user.User, error) {
	user := user.User{}

	if login == "" || password == "" {
		return user, errors.New("empty login or password")
	}

	row := s.db.QueryRow(context.Background(),
		"SELECT id, login, password, email, first_name, last_name, super_user, banned, created_at, updated_at FROM admin_users WHERE login = $1", login)
	err := row.Scan(
		&user.Id,
		&user.Login,
		&user.Password,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.SuperUser,
		&user.Banned,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err != nil {
		fmt.Printf("%t, %v", err, err)
		return user, errors.New("incorrect login or password")
	}

	if !utils.PasswordVerify(password, user.Password) {
		return user, errors.New("incorrect login or password")
	}

	return user, err
}
