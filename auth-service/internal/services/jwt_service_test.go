package services

import (
	"testing"
)

func TestGenerateAndValidateJWT(t *testing.T) {
	userID := uint64(1)
	secret := "testsecret"

	// Generate token
	token, err := GenerateJWT(userID, secret)
	if err != nil {
		t.Fatalf("Failed to generate JWT: %v", err)
	}

	// Validate token
	claims, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("Failed to validate JWT: %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("Expected userID %d, got %d", userID, claims.UserID)
	}
}
