package http

import (
	"encoding/json"
	nethttp "net/http"
)

type healthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func NewHandler() nethttp.Handler {
	mux := nethttp.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)

	return mux
}

func healthHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(nethttp.StatusOK)

	_ = json.NewEncoder(w).Encode(healthResponse{
		Status:  "ok",
		Service: "walkquest-service",
	})
}
