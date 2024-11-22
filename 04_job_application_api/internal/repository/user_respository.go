package repository

import (
	"database/sql"
	"fmt"
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

func UpdateUserProfile(db *sql.DB, user *models.User) (*models.User, error) {
	_, err := db.Exec("UPDATE users SET username = ?, email = ? WHERE id = ?", user.Username, user.Email, user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateProfilePicture(db *sql.DB, id int, profilePicture string) error {
	_, err := db.Exec("UPDATE users SET profile_picture = ? WHERE id = ?", profilePicture, id)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers(db *sql.DB) ([]*models.User, error) {
	var users []*models.User
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		var profilePicture sql.NullString
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin, &profilePicture)
		if err != nil {
			return nil, err
		}
		if profilePicture.Valid {
			user.ProfilePicture = &profilePicture.String
		} else {
			user.ProfilePicture = nil
		}
		users = append(users, &user)
	}

	return users, nil
}

func UpdateUserPassword(db *sql.DB, user *models.User) error {
	_, err := db.Exec("UPDATE users SET password = ? WHERE id = ?", user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserWithTransaction(tx *sql.Tx, userID int) (string, error) {
	// Delete associated jobs first
	_, err := tx.Exec("DELETE FROM jobs WHERE user_id = ?", userID)

	if err != nil {
		return "", fmt.Errorf("error deleting user's jobs: %v", err)
	}

	// Get user's profile picture before deleting user

	var profilePicture sql.NullString
	err = tx.QueryRow("SELECT profile_picture FROM users WHERE id = ?", userID).Scan(&profilePicture)
	if err != nil {
		return "", fmt.Errorf("error fetching user's profile picture: %v", err)
	}

	// Delete the user
	result, err := tx.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		return "", fmt.Errorf("error deleting user: %v", err)
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return "", fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return "", sql.ErrNoRows
	}

	return profilePicture.String, nil

}
