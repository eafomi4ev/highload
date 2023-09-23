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
	ID           uuid.UUID
	FirstName    string
	Surname      string
	Birthdate    time.Time
	Biography    string
	City         City
	PasswordHash string
}

type City struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
