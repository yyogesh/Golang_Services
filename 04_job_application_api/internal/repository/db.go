package repository

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./jobportal.db")

	if err != nil {
		return nil, err
	}

	err = createTable(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
	   id INTEGER PRIMARY KEY AUTOINCREMENT,
	   username TEXT NOT NULL UNIQUE,
	   password TEXT NOT NULL,
	   email TEXT NOT NULL UNIQUE,
	   created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	   updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	   is_admin BOOLEAN DEFAULT 0
	   profile_picture TEXT
   )`)

	return err
}
