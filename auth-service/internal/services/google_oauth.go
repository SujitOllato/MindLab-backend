package services

import (
	"auth-service/internal/database"
	"auth-service/internal/models"
	"database/sql"
	"time"
)

func FindOrCreateGoogleUser(email, name string) (*models.User, error) {
	var user models.User

	err := database.DB.QueryRow(`
		SELECT id, email, name, provider, created_at
		FROM users WHERE email = ?
	`, email).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Provider,
		&user.CreatedAt,
	)

	if err == sql.ErrNoRows {
		res, err := database.DB.Exec(`
			INSERT INTO users (email, name, provider, created_at)
			VALUES (?, ?, 'google', ?)
		`, email, name, time.Now())

		if err != nil {
			return nil, err
		}

		id, _ := res.LastInsertId()
		user = models.User{
			ID:        uint64(id),
			Email:     email,
			Name:      name,
			Provider:  "google",
			CreatedAt: time.Now(),
		}
		return &user, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
