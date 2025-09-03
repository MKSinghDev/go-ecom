// Package config
package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string

	DBUser       string
	DBPassword   string
	DBHost       string
	DBPort       string
	DBName       string
	DBConnString string

	JWTExpirationInSeconds int64
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8000"),

		DBUser:     getEnv("DB_USERNAME", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "12345"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "go-ecom"),
		DBConnString: fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			getEnv("DB_USERNAME", "postgres"),
			getEnv("DB_PASSWORD", "12345"),
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_PORT", "5432"),
			getEnv("DB_NAME", "go-ecom"),
			getEnv("DB_SSL_MODE", "disable"),
		),

		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "not-secret-secret-anymore?"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if val := os.Getenv(key); val != "" {
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
