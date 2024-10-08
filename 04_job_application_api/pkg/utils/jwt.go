package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Clamis struct {
	Username string `json:"username"`
	UserID   int    `json:"userID"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.StandardClaims
}

func GenerateToken(username string, userID int, isAdmin bool) (string, error) {
	exiprationTime := time.Now().Add(10 * time.Hour)

	clamis := &Clamis{
		Username: username,
		UserID:   userID,
		IsAdmin:  isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exiprationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clamis)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
