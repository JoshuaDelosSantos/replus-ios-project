package auth

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndValidateToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret") // Set a test secret

	// Generate token
	token, err := GenerateToken(1) // Generate token for UserID = 1
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Validate token
	claims, err := ValidateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, 1, claims.UserID) // Ensure the UserID is correct

	// Check token expiration
	assert.WithinDuration(t, time.Now(), claims.IssuedAt.Time, 5*time.Second) // IssuedAt should be recent
	assert.WithinDuration(t, time.Now().Add(15*time.Minute), claims.ExpiresAt.Time, 5*time.Second) // ExpiresAt should be 15 mins from now
}