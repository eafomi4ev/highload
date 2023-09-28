package uc_user_get_by_id

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"otus_highload/internal/app/usecase"
	"otus_highload/internal/domain"
	"otus_highload/internal/storage"
)

type UseCase struct {
	storage Storage
}

func New(storage Storage) *UseCase {
	return &UseCase{storage: storage}
}

func (uc *UseCase) GetUserByID(ctx context.Context, userID uuid.UUID) (domain.User, error) {
	user, err := uc.storage.GetUserByID(ctx, userID)
	if errors.Is(err, storage.ErrNotFound) {
		return domain.User{}, usecase.ErrUserNotFound
	} else if err != nil {
		return domain.User{}, err
	}

	user.PasswordHash = "" // на всякий случай зануляем, чтоб не отдавать наружу

	return user, err
}
