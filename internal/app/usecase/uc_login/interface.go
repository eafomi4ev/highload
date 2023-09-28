package uc_login

import (
	"context"
	"github.com/google/uuid"
	"otus_highload/internal/domain"
)

type Storage interface {
	GetPasswordHashByUserID(context.Context, uuid.UUID) (domain.User, error)
}
