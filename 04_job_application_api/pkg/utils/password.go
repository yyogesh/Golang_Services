package utils

import (
	"math/rand"
	"strings"
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
