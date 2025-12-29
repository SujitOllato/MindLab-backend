package services

import "auth-service/internal/config"

var (
	jwtSecret      string
	googleClientID string
)

func Init(cfg *config.Config) {
	jwtSecret = cfg.JWTSecret
	googleClientID = cfg.GoogleClientID
}

func GetGoogleClientID() string {
	return googleClientID
}
