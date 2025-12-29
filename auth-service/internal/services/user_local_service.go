package services

import (
	"auth-service/internal/database"
	"auth-service/internal/models"
	"database/sql"
	"errors"
	"time"
)

/* =======================
   CREATE USER
======================= */

func CreateLocalUser(email, password, name string) (*models.User, error) {

	var id uint64
	err := database.DB.QueryRow(
		"SELECT id FROM users WHERE email = ?",
		email,
	).Scan(&id)

	if err == nil {
		return nil, errors.New("email already registered")
	}

	if err != sql.ErrNoRows {
		return nil, err
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	res, err := database.DB.Exec(`
		INSERT INTO users (email, name, provider, password_hash, created_at)
		VALUES (?, ?, 'local', ?, ?)
	`, email, name, hashedPassword, time.Now())

	if err != nil {
		return nil, err
	}

	userID, _ := res.LastInsertId()

	return &models.User{
		ID:        uint64(userID),
		Email:     email,
		Name:      name,
		Provider:  "local",
		CreatedAt: time.Now(),
	}, nil
}

/* =======================
   LOGIN USER
======================= */

func AuthenticateLocalUser(email, password string) (*models.User, error) {
	var user models.User
	var passwordHash string

	err := database.DB.QueryRow(`
		SELECT id, email, name, provider, password_hash, created_at
		FROM users
		WHERE email = ? AND provider = 'local'
	`, email).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Provider,
		&passwordHash,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !CheckPassword(passwordHash, password) {
		return nil, errors.New("invalid email or password")
	}

	return &user, nil
}
