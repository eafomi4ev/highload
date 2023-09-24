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
	ID           uuid.UUID `json:"id" db:"id"`
	FirstName    string    `json:"first_name" db:"first_name"`
	Surname      string    `json:"surname" db:"surname"`
	Birthdate    time.Time `json:"birthdate" db:"birthdate"`
	Biography    string    `json:"biography,omitempty" db:"biography"`
	City         City      `json:"city" db:"city"`
	PasswordHash string    `json:"password_hash,omitempty" db:"password_hash"`
}

type City struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
