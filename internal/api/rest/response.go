package rest

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, data interface{}) {
	switch v := data.(type) {
	case ErrorResponse:
		w.WriteHeader(v.Status)
	default:
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")

	b, _ := json.Marshal(data)
	_, _ = w.Write(b)
}
