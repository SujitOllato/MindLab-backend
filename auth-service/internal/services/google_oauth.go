package services

import (
    "context"
    "errors"

    "google.golang.org/api/idtoken"
)

func VerifyGoogleToken(token string, clientID string) (*idtoken.Payload, error) {
    payload, err := idtoken.Validate(context.Background(), token, clientID)
    if err != nil {
        return nil, errors.New("invalid google token")
    }
    return payload, nil
}
