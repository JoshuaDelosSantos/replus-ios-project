package auth

import (
	"fmt"
)

// MockValidator implements the TokenValidator interface for testing purposes.
type MockValidator struct{}

// ValidateToken is a mock implementation of token validation.
func (m *MockValidator) ValidateToken(token string) (*Claims, error) {
	// Return valid claims for "validtoken" and an error for other tokens
	if token == "Bearer validtoken" {
		return &Claims{UserID: 1}, nil
	}
	return nil, fmt.Errorf("invalid token")
}