package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	DBHost                   string
	DBPort                   string
	DBUser                   string
	DBPassword               string
	DBName                   string
	DBSSLMode                string
	AutoCompleteDelayMinutes int
}

func Load() *Config {
	cfg := &Config{
		DBHost:                   getEnv("DB_HOST", "localhost"),
		DBPort:                   getEnv("DB_PORT", "5432"),
		DBUser:                   getEnv("DB_USER", "postgres"),
		DBPassword:               getEnv("DB_PASSWORD", "postgres123"),
		DBName:                   getEnv("DB_NAME", "postgres"),
		DBSSLMode:                getEnv("DB_SSLMODE", "disable"),
		AutoCompleteDelayMinutes: getEnvInt("AUTO_COMPLETE_DELAY", 10),
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("env %s not set, using default", key)
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("invalid int for %s, using default", key)
		return defaultValue
	}
	return i
}
