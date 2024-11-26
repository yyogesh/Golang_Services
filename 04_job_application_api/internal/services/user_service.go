package services

import (
	"database/sql"
	"fmt"
	"job_portal/internal/models"
	"job_portal/internal/repository"
	"job_portal/pkg/utils"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
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

func DeleteUser(c *gin.Context, db *sql.DB, userID int) error {
	// Start a transaction
	tx, err := db.BeginTx(c.Request.Context(), &sql.TxOptions{Isolation: sql.LevelSerializable})

	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}

	defer tx.Rollback() // Rollback if not committed

	// Delete user and associated data
	profilePicture, err := repository.DeleteUserWithTransaction(tx, userID)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("error deleting user: %v", err)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	// Delete profile picture after successful transaction if it exists
	if profilePicture != "" {
		filePath := filepath.Join(os.Getenv("UPLOAD_DIR"), profilePicture)
		err = utils.DeleteFileIfExists(filePath)

		if err != nil {
			return fmt.Errorf("error deleting profile picture: %v", err)
		}
	}

	return nil
}

func ChangePassword(db *sql.DB, userID int, currentPassword, newPassword string) error {
	return repository.ChangePassword(db, userID, currentPassword, newPassword)
}
