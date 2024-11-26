package utils

import (
	"fmt"
	"job_portal/internal/models"
	"math/rand"
	"strings"
	"unicode"
)

func GenerateFromPassword(charCount int) string {
	const digit = "0123456789abcdef"

	var password strings.Builder
	password.Grow(charCount)

	for i := 0; i < charCount; i++ {
		password.WriteByte(digit[rand.Intn(len(digit))])
	}
	return password.String()
}

func ValidatePasswordStrength(password string) (bool, []string) {
	validation := models.PasswordValidation{
		MinLength:  8,
		HasUpper:   true,
		HasLower:   true,
		HasNumber:  true,
		HasSpecial: true,
	}

	var validationErrors []string

	// Check minimum length
	if len(password) < validation.MinLength {
		validationErrors = append(validationErrors,
			fmt.Sprintf("Password must be at least %d characters long", validation.MinLength))
	}

	// Check for uppercase letters
	if validation.HasUpper && !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		validationErrors = append(validationErrors,
			"Password must contain at least one uppercase letter")
	}

	// Check for lowercase letters
	if validation.HasLower && !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		validationErrors = append(validationErrors,
			"Password must contain at least one lowercase letter")
	}

	// Check for numbers
	if validation.HasNumber && !strings.ContainsAny(password, "0123456789") {
		validationErrors = append(validationErrors,
			"Password must contain at least one number")
	}

	// Check for special characters
	if validation.HasSpecial {
		hasSpecial := false
		for _, char := range password {
			if unicode.IsPunct(char) || unicode.IsSymbol(char) {
				hasSpecial = true
				break
			}
		}
		if !hasSpecial {
			validationErrors = append(validationErrors,
				"Password must contain at least one special character")
		}
	}

	return len(validationErrors) == 0, validationErrors

}
