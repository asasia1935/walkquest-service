package main

import (
	"context"
	"log"
	nethttp "net/http"
	"time"

	"github.com/asasia1935/walkquest-service/internal/config"
	"github.com/asasia1935/walkquest-service/internal/db"
	platformhttp "github.com/asasia1935/walkquest-service/internal/platform/http"
)

func main() {
	cfg := config.Load()

	database, err := db.Open(cfg.Database)
	if err != nil {
		log.Fatalf("database open failed: %v", err)
	}
	defer database.Close()

	// timeout을 걸어 DB가 응답하지 않을 때 서버 시작이 오래 멈추지 않도록 합니다.
	pingCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.Ping(pingCtx, database); err != nil {
		log.Fatalf("database ping failed: %v", err)
	}
	log.Print("database ping succeeded")

	handler := platformhttp.NewHandler()

	addr := ":" + cfg.Port
	log.Printf("walkquest-service listening on %s", addr)

	if err := nethttp.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
