package uc_user_search_by_name

import (
	"context"

	"otus_highload/internal/domain"
)

type Storage interface {
	SearchUsersByName(ctx context.Context, firstName, surname string) ([]domain.User, error)
}
