package http

import (
	"encoding/json"
	nethttp "net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w nethttp.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(body)
}

func writeError(w nethttp.ResponseWriter, statusCode int, code string) {
	writeJSON(w, statusCode, errorResponse{
		Error: code,
	})
}
