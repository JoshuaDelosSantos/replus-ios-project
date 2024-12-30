package auth

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPassword hashes the user's password.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash checks if the provided password matches the hashed password.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}