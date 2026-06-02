package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/asasia1935/walkquest-service/internal/config"
)

func Open(cfg config.DatabaseConfig) (*sql.DB, error) {
	// sql.Open은 실제 DB 연결을 즉시 수행하지 않고,
	// DB 연결 풀을 사용할 준비만 합니다.
	database, err := sql.Open("postgres", dataSourceName(cfg))
	if err != nil {
		return nil, err
	}

	return database, nil
}

func Ping(ctx context.Context, database *sql.DB) error {
	// PingContext로 실제 DB 연결 가능 여부를 확인합니다.
	return database.PingContext(ctx)
}

func dataSourceName(cfg config.DatabaseConfig) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.SSLMode,
	)
}
