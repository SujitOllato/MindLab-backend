package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(dbURL string) error {
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return err
	}

	DB = db
	return DB.Ping()
}
