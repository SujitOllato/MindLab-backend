package database

import (
	"testing"
)

func TestDBConnection(t *testing.T) {
	db, err := GetDB()
	if err != nil {
		t.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("DB ping failed: %v", err)
	}
}
