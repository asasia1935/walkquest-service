package http

import (
	"encoding/json"
	nethttp "net/http"
)

type healthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func NewHandler(explorerProfileHandler *ExplorerProfileHandler) nethttp.Handler {
	mux := nethttp.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("POST /explorer-profiles", explorerProfileHandler.CreateProfile)
	mux.HandleFunc("GET /explorer-profiles/me", explorerProfileHandler.GetMyProfile)

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
