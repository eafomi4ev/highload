package user_register

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"otus_highload/internal/domain"
	"otus_highload/internal/utils"
)

type Handler struct {
	ctx context.Context
	uc  RegisterUserUseCase
}

func New(ctx context.Context, uc RegisterUserUseCase) *Handler { // todo: погуглить, как сюда прокинуть контекст получше, чем явная передача параметром
	return &Handler{
		ctx: ctx,
		uc:  uc,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := h.ctx

	var body requestBody
	err := utils.BodyParser(r.Body, &body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(fmt.Errorf("unable to parse request body: %w", err).Error()))
		return
	}

	dto := domain.UserIn{
		FirstName: body.FirstName,
		Surname:   body.Surname,
		Birthdate: body.Birthdate,
		Biography: body.Biography,
		City:      body.City,
		Password:  body.Password,
	}

	user, err := h.uc.RegisterUser(ctx, dto)
	if errors.Is(err, domain.DataValidationError) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	res := responseBody{
		UserID: user.ID.String(),
	}

	b, _ := json.Marshal(res)

	_, _ = w.Write(b)
	return

}
