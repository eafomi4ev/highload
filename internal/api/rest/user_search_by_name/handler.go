package user_search_by_name

import (
	"context"
	"net/http"

	"otus_highload/internal/api/rest"
	"otus_highload/internal/app/usecase/uc_user_search_by_name"
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
	firstName := r.URL.Query().Get("first_name")
	surname := r.URL.Query().Get("last_name")

	dto := uc_user_search_by_name.UserSearchIn{
		FirstNamePrefix: firstName,
		SurnamePrefix:   surname,
	}

	users, err := h.uc.SearchUsersByName(h.ctx, dto)
	if err != nil {
		rest.Response(w, rest.InternalServerError("", err))
		return
	}

	rest.Response(w, users)
}
