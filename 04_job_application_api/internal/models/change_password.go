package models

type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
}

type PasswordValidation struct {
	MinLength  int
	HasUpper   bool
	HasLower   bool
	HasNumber  bool
	HasSpecial bool
}
