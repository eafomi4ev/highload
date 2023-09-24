package login

import (
	"context"
	"otus_highload/internal/app/usecase/uc_login"
)

type LoginUseCase interface {
	Login(ctx context.Context, in uc_login.LoginInfoIn) (uc_login.LoginInfoOut, error)
}
