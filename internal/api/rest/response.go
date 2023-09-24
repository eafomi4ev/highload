package rest

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, data interface{}) {
	// сетить Content-Type надо до вызова w.WriteHeader(), иначе .Set() не сработает и Content-Type останется plain/text
	// https://github.com/golang/go/issues/17083#issuecomment-246557291
	w.Header().Set("Content-Type", "application/json")

	switch v := data.(type) {
	case ErrorResponse:
		w.WriteHeader(v.Status)
	default:
		w.WriteHeader(http.StatusOK)
	}

	b, _ := json.Marshal(data)
	_, _ = w.Write(b)
}
