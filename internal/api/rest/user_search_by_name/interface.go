package user_search_by_name

import (
	"context"

	"otus_highload/internal/app/usecase/uc_user_search_by_name"
	"otus_highload/internal/domain"
)

type UseCase interface {
	SearchUsersByName(ctx context.Context, in uc_user_search_by_name.UserSearchIn) ([]domain.User, error)
}
