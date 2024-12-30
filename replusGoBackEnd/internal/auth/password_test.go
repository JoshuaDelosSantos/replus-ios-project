package auth

import (
	"log"
	"testing"
)

var passwordTestLogger *log.Logger

func init() {
	passwordTestLogger = log.New(log.Writer(), "[PASSWORD_TEST] ", log.LstdFlags|log.Lshortfile)
}

func TestPassword(t *testing.T) {
	passwordTestLogger.Println("Hashing password...")

	userPassword := "password"

	hashedPassword, err := HashPassword(userPassword)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	passwordTestLogger.Println("HashPassword method executed successfully.")

	passwordTestLogger.Println("Checking hash password...")

	if match := CheckPasswordHash(userPassword, hashedPassword); !match {
        t.Errorf("Password should match but doesn't")
    } else {
        passwordTestLogger.Println("Password matched successfully")
    }

}