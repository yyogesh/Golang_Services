package services

import (
	"database/sql"
	"job_portal/internal/models"
	"job_portal/internal/repository"
)

func GetUserByID(db *sql.DB, id int) (*models.User, error) {
	return repository.GetUserByID(db, id)
}

func UpdateUserProfile(db *sql.DB, id int, username, emailId string) (*models.User, error) {
	user := &models.User{ID: id, Username: username, Email: emailId}

	return repository.UpdateUserProfile(db, user)
}

func UpdateProfilePicture(db *sql.DB, id int, profilePicture string) error {
	return repository.UpdateProfilePicture(db, id, profilePicture)
}

func GetAllUsers(db *sql.DB) ([]*models.User, error) {
	return repository.GetAllUsers(db)
}
