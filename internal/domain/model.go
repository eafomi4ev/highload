package domain

import (
	"github.com/google/uuid"
	"time"
)

type UserIn struct {
	FirstName string
	Surname   string
	Birthdate string
	Biography string
	City      string
	Password  string
}

type User struct {
	ID           uuid.UUID `db:"id"`
	FirstName    string    `db:"first_name"`
	Surname      string    `db:"surname"`
	Birthdate    time.Time `db:"birthdate"`
	Biography    string    `db:"biography,omitempty"`
	City         City      `db:"city"`
	PasswordHash string    `db:"password_hash"`
}

type City struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
