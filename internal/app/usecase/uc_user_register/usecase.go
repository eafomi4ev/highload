package uc_user_register

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"otus_highload/internal/app/usecase"
	"otus_highload/internal/domain"
	"time"
)

type UseCase struct {
	storage Storage
}

func New(storage Storage) *UseCase {
	return &UseCase{
		storage,
	}
}

func (uc *UseCase) RegisterUser(ctx context.Context, in domain.UserIn) (domain.User, error) {
	err := validate(in)
	if err != nil {
		return domain.User{}, err
	}

	birthdate, err := time.Parse("02-01-2006", in.Birthdate)
	if err != nil {
		return domain.User{}, fmt.Errorf("%w: %s", usecase.ErrDataValidation, "birthdate is not valid, the forman should be DD-MM-YYYY")
	}

	newID, _ := uuid.NewRandom()

	city, err := uc.storage.GetCityByName(ctx, in.City)
	if err != nil {
		return domain.User{}, err
	}

	dto := domain.User{
		ID:        newID,
		FirstName: in.FirstName,
		Surname:   in.Surname,
		Birthdate: birthdate,
		Biography: in.Biography,
		City: domain.City{
			ID:   city.ID,
			Name: city.Name,
		},
		PasswordHash: hashPassword(in.Password),
	}

	user, err := uc.storage.AddUser(ctx, dto)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func validate(in domain.UserIn) error {
	if in.Password == "" {
		return fmt.Errorf("%w: %s", usecase.ErrDataValidation, "password is empty")
	}

	if in.FirstName == "" {
		return fmt.Errorf("%w: %s", usecase.ErrDataValidation, "first name is empty")
	}

	if in.Surname == "" {
		return fmt.Errorf("%w: %s", usecase.ErrDataValidation, "surname is empty")
	}

	if in.City == "" {
		return fmt.Errorf("%w: %s", usecase.ErrDataValidation, "city is empty")
	}

	if in.Birthdate == "" {
		return fmt.Errorf("%w: %s", usecase.ErrDataValidation, "city is empty")
	}

	return nil
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
