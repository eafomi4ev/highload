package user_register

import (
	"context"
	"errors"
	"net/http"
	"otus_highload/internal/api/rest"
	"otus_highload/internal/domain"
	"otus_highload/internal/utils"
)

type Handler struct {
	ctx context.Context
	uc  RegisterUserUseCase
}

func New(ctx context.Context, uc RegisterUserUseCase) *Handler {
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

	dto := domain.UserIn{
		FirstName: body.FirstName,
		Surname:   body.Surname,
		Birthdate: body.Birthdate,
		Biography: body.Biography,
		City:      body.City,
		Password:  body.Password,
	}

	user, err := h.uc.RegisterUser(ctx, dto)
	if errors.Is(err, rest.DataValidationError) {
		rest.Response(w, rest.BadRequest("", err))
		return
	} else if err != nil {
		rest.Response(w, rest.InternalServerError("", err))
		return
	}

	res := responseBody{
		UserID: user.ID.String(),
	}

	rest.Response(w, res)
}
