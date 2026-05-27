package config

import "os"

const defaultPort = "8080"

type Config struct {
	Port     string
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
	SSLMode  string
}

func Load() Config {
	return Config{
		Port: envOrDefault("PORT", defaultPort),
		Database: DatabaseConfig{
			Name:     envOrDefault("DB_NAME", "walkquest"),
			User:     envOrDefault("DB_USER", "walkquest"),
			Password: envOrDefault("DB_PASSWORD", "walkquest"),
			Host:     envOrDefault("DB_HOST", "localhost"),
			Port:     envOrDefault("DB_PORT", "5432"),
			SSLMode:  envOrDefault("DB_SSLMODE", "disable"),
		},
	}
}

// envOrDefault는 OS 환경변수를 읽고, 값이 없으면 기본값을 반환합니다.
// (이 함수는 .env 파일을 읽는 것이 아닌 PowerShell, 터미널, Docker, 배포 환경 등에서 OS 환경변수로 직접 설정해야 합니다.)
func envOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
