package uc_user_get_by_id

import (
	"context"
	"github.com/google/uuid"
	"otus_highload/internal/domain"
)

type Storage interface {
	GetUserByID(ctx context.Context, userID uuid.UUID) (domain.User, error)
}
