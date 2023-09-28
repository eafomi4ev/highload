package uc_user_register

import (
	"context"
	"otus_highload/internal/domain"
)

type Storage interface {
	AddUser(context.Context, domain.User) (domain.User, error)
	GetCityByName(ctx context.Context, name string) (domain.City, error)
}
