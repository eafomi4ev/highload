package uc_user_search_by_name

import (
	"context"

	"otus_highload/internal/domain"
)

type UseCase struct {
	storage Storage
}

func New(storage Storage) *UseCase {
	return &UseCase{storage: storage}
}

func (uc *UseCase) SearchUsersByName(ctx context.Context, in UserSearchIn) ([]domain.User, error) {
	return uc.storage.SearchUsersByName(ctx, in.FirstNamePrefix, in.SurnamePrefix)
}
