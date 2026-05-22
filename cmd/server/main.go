package main

import (
	"log"
	nethttp "net/http"

	"github.com/asasia1935/walkquest-service/internal/config"
	platformhttp "github.com/asasia1935/walkquest-service/internal/platform/http"
)

func main() {
	cfg := config.Load()
	handler := platformhttp.NewHandler()

	addr := ":" + cfg.Port
	log.Printf("walkquest-service listening on %s", addr)

	if err := nethttp.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
