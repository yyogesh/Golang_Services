package repository

import (
	"database/sql"
	"job_portal/internal/models"
)

func CreateUser(db *sql.DB, user *models.User) error {
	_, err := db.Exec(`INSERT INTO users (username, password, email) VALUES (?, ?, ?)`, user.Username, user.Password, user.Email)
	return err
}

func GetUserByID(db *sql.DB, id int) (*models.User, error) {
	var user models.User
	var profilePicture sql.NullString // Use sql.NullString to handle NULL values
	err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin, &profilePicture)
	if err != nil {
		return nil, err
	}
	if profilePicture.Valid {
		user.ProfilePicture = &profilePicture.String
	} else {
		user.ProfilePicture = nil
	}
	return &user, nil
}

func GetUserByUserName(db *sql.DB, username string) (*models.User, error) {
	user := &models.User{}
	err := db.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin, &user.ProfilePicture)

	if err != nil {
		return nil, err
	}
	return user, nil
}
