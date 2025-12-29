package config

import "os"

type Config struct {
	DBUrl          string
	GoogleClientID string
	JWTSecret      string
}

func Load() *Config {
	return &Config{
		DBUrl:          os.Getenv("DB_URL"),
		GoogleClientID: os.Getenv("GOOGLE_CLIENT_ID"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
	}
}
