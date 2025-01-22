package config

import "os"

type Config struct {
	Port  string
	Env   string
	DBURL string
}

func NewConfig() (*Config, error) {
	return &Config{
		Port:  getEnv("PORT", "8080"),
		Env:   getEnv("ENV", "development"),
		DBURL: getEnv("DB_URL", "postgres://localhost:5432/mydb"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
