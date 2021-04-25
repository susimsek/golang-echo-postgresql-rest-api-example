package config

import "os"

var (
	ServerPort    = GetEnv("SERVER_PORT", "9000")
	PostgresqlUrl = GetEnv("POSTGRESQL_URL", "host=localhost port=5432 user=postgres password=root dbname=demo sslmode=disable")
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
