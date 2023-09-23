package login

import (
	"net/http"
	"otus_highload/internal/utils"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	var body requestBody
	err := utils.BodyParser(r.Body, &body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("OK"))
	return

}
