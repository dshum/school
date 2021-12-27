package user

import "time"

type User struct {
	Id        int       `json:"id"`
	Login     string    `json:"login"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	SuperUser bool      `json:"super_user"`
	Banned    bool      `json:"banned"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
