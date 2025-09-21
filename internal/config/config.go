package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Values *Config

type Config struct {
	AppPort         string
	DBURL           string
	MaxConnections  int
	MinConnections  int
	MaxConnLifetime string
	MaxConnIdleTime string
	AuthSecret      string
}

func Load() {
	_ = godotenv.Load()

	Values = &Config{
		AppPort:         getEnv("APP_PORT", "8080"),
		DBURL:           mustEnv("DB_URL"),
		MaxConnections:  getEnvAsInt("DB_MAX_CONNECTIONS", 25),
		MinConnections:  getEnvAsInt("DB_MIN_CONNECTIONS", 5),
		MaxConnLifetime: getEnv("DB_MAX_CONN_LIFETIME", "1h"),
		MaxConnIdleTime: getEnv("DB_MAX_CONN_IDLE_TIME", "30m"),
		AuthSecret:      getEnv("AUTH_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func mustEnv(key string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	log.Fatalf("env %s not set", key)
	return ""
}

func getEnvAsInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if intVal, err := strconv.Atoi(v); err == nil {
			return intVal
		}
		log.Printf("Invalid integer value for %s: %s, using fallback: %d", key, v, fallback)
	}
	return fallback
}
