package user_register

import (
	"context"
	"otus_highload/internal/domain"
)

type RegisterUserUseCase interface {
	RegisterUser(ctx context.Context, in domain.UserIn) (domain.User, error)
}
