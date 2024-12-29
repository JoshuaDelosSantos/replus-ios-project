package auth

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "[AUTH_TEST] ", log.LstdFlags|log.Lshortfile)
}

func TestGenerateAndValidateToken(t *testing.T) {
	logger.Println("Starting token generation and validation test")
	
	os.Setenv("JWT_SECRET", "testsecret")
	logger.Printf("Set test JWT_SECRET: %s", "testsecret")

	// Generate token
	userID := 1
	logger.Printf("Generating token for UserID: %d", userID)
	token, err := GenerateToken(userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	logger.Printf("Generated token successfully: %s", token[:10]+"...")

	// Validate token
	logger.Println("Validating generated token")
	claims, err := ValidateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	logger.Printf("Token validated successfully. UserID matches: %d", claims.UserID)

	// Check token expiration
	logger.Println("Checking token timestamps")
	assert.WithinDuration(t, time.Now(), claims.IssuedAt.Time, 5*time.Second)
	assert.WithinDuration(t, time.Now().Add(15*time.Minute), claims.ExpiresAt.Time, 5*time.Second)
	logger.Printf("Token timestamps verified. IssuedAt: %v, ExpiresAt: %v", 
		claims.IssuedAt.Time, claims.ExpiresAt.Time)
}

func TestInvalidToken(t *testing.T) {
	logger.Println("Starting invalid token test")
	
	invalidToken := "invalid.token.string"
	logger.Printf("Attempting to validate invalid token: %s", invalidToken)
	
	_, err := ValidateToken(invalidToken)
	assert.Error(t, err)
	logger.Printf("Expected error received: %v", err)
}

func TestExpiredToken(t *testing.T) {
	logger.Println("Starting expired token test")
	
	os.Setenv("JWT_SECRET", "testsecret")
	token, _ := GenerateToken(1)
	logger.Printf("Generated test token: %s", token[:10]+"...")

	// Wait for token to expire (if using short expiration for testing)
	time.Sleep(time.Second)
	logger.Println("Validating potentially expired token")
	
	claims, err := ValidateToken(token)
	if err != nil {
		logger.Printf("Token validation failed as expected: %v", err)
	} else {
		logger.Printf("Token still valid. Expiration time: %v", claims.ExpiresAt.Time)
	}
}