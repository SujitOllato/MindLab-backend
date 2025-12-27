package services

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint64, uuid string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "uuid":    uuid,
        "exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(JWTSecret))
}
