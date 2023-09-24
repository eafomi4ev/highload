package uc_login

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"otus_highload/internal/app/usecase"
)

type UseCase struct {
	storage Storage
}

func New(storage Storage) *UseCase {
	return &UseCase{
		storage,
	}
}

func (uc UseCase) Login(ctx context.Context, in LoginInfoIn) (LoginInfoOut, error) {
	userID, err := uuid.Parse(in.UserID) // todo: добавить валидацию на корректность uuid, чтобы не возвращать пятисотку
	if err != nil {
		return LoginInfoOut{}, err
	}

	user, err := uc.storage.GetPasswordHashByUserID(ctx, userID)
	if err != nil {
		return LoginInfoOut{}, err
	}
	if !checkPasswordHash(in.Password, user.PasswordHash) {
		return LoginInfoOut{}, usecase.ErrIncorrectPassword
	}

	token, err := uuid.NewRandom() // todo: сохранять токен куда-нибудь в хранилище
	if err != nil {
		return LoginInfoOut{}, fmt.Errorf("login failed: %w", err)
	}

	res := LoginInfoOut{
		Token: token.String(),
	}

	return res, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
