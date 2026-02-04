package config

import "os"

type Config struct {
	Port  string
	DBDSN string
}

func Load() Config {
	return Config{
		Port:  getEnv("APP_PORT", "8080"),
		DBDSN: getEnv("DB_DSN", ""),
	}
}

func getEnv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
