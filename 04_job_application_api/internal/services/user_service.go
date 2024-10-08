package services

import (
	"database/sql"
	"job_portal/internal/models"
	"job_portal/internal/repository"
)

func GetUserByID(db *sql.DB, id int) (*models.User, error) {
	return repository.GetUserByID(db, id)
}
