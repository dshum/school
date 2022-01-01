package auth

import (
	"context"
	"errors"
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
	u := user.User{}

	if login == "" || password == "" {
		return u, errors.New("empty login or password")
	}

	row := s.db.QueryRow(context.Background(),
		"SELECT id, login, password, email, first_name, last_name, super_user, banned, created_at, updated_at FROM admin_users WHERE login = $1", login)
	err := row.Scan(
		&u.Id,
		&u.Login,
		&u.Password,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.SuperUser,
		&u.Banned,
		&u.CreatedAt,
		&u.UpdatedAt)

	if err != nil {
		return u, errors.New("incorrect login")
	}

	if !utils.PasswordVerify(password, u.Password) {
		return u, errors.New("incorrect login or password")
	}

	return u, err
}
