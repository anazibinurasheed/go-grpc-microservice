package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Internally calling the bcrypt.GenerateFromPassword()
func HashPassword(password string) (string, bool) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 11)

	return string(bytes), err == nil
}

// Internally calling the bcrypt.ComparePasswordHash()
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
