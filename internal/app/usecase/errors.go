package usecase

import "errors"

var (
	ErrDataValidation    = errors.New("data validation error")
	ErrIncorrectPassword = errors.New("password validation error")
	ErrUserNotFound      = errors.New("user with such id has not been found")
)
