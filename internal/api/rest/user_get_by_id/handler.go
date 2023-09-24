package user_get_by_id

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"otus_highload/internal/api/rest"
	"otus_highload/internal/app/usecase"
)

type Handler struct {
	ctx context.Context
	uc  UseCase
}

func New(ctx context.Context, uc UseCase) *Handler {
	return &Handler{
		ctx: ctx,
		uc:  uc,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userIDRaw, ok := params["id"]
	if !ok {
		rest.Response(w, rest.BadRequest("incorrect user id", errors.New("no user id")))
		return
	}

	userID, err := uuid.Parse(userIDRaw)
	if err != nil {
		rest.Response(w, rest.BadRequest("incorrect user id", errors.New("invalid uuid")))
		return
	}

	user, err := h.uc.GetUserByID(h.ctx, userID)
	if errors.Is(err, usecase.ErrUserNotFound) {
		rest.Response(w, rest.NotFound("user not found", err))
		return
	} else if err != nil {
		rest.Response(w, rest.InternalServerError("", errors.New("something went wrong")))
		return
	}

	res := responseBody{
		FirstName:  user.FirstName,
		SecondName: user.Surname,
		Birthdate:  user.Birthdate.Format("02-01-2006"),
		Biography:  user.Biography,
		City:       user.City.Name,
	}

	rest.Response(w, res)
}
