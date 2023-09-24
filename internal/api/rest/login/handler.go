package login

import (
	"context"
	"errors"
	"net/http"
	"otus_highload/internal/api/rest"
	"otus_highload/internal/app/usecase"
	"otus_highload/internal/app/usecase/uc_login"
	"otus_highload/internal/utils"
)

type Handler struct {
	ctx context.Context
	uc  LoginUseCase
}

func New(ctx context.Context, uc LoginUseCase) *Handler {
	return &Handler{
		ctx: ctx, // todo: подумать, как получше прокинуть контекст в хендлер. Через замыкание?
		uc:  uc,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := h.ctx

	var body requestBody
	err := utils.BodyParser(r.Body, &body)
	if err != nil {
		rest.Response(w, rest.BadRequest("unable to parse request body", err))
		return
	}
	defer r.Body.Close()

	dto := uc_login.LoginInfoIn{
		UserID:   body.ID,
		Password: body.Password,
	}

	token, err := h.uc.Login(ctx, dto)
	if errors.Is(err, usecase.ErrIncorrectPassword) {
		rest.Response(w, rest.BadRequest("bad password", err))
		return
	} else if errors.Is(err, usecase.ErrUserNotFound) {
		rest.Response(w, rest.NotFound("user with such id has not been found", err))
		return
	} else if errors.Is(err, usecase.ErrDataValidation) {
		rest.Response(w, rest.BadRequest("", err))
		return
	} else if err != nil {
		rest.Response(w, rest.InternalServerError("", err))
		return
	}

	rest.Response(w, token)
}
