package services

import (
    "os"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint64, uuid string) (string, error) {
    jwtSecret := os.Getenv("JWT_SECRET")

    claims := jwt.MapClaims{
        "user_id": userID,
        "uuid":    uuid,
        "exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(jwtSecret))
}
