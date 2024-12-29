package auth

import (
    "testing"
    "time"
	"github.com/golang-jwt/jwt/v4"
    "github.com/stretchr/testify/assert"
)

func TestJWTValidator(t *testing.T) {
    secret := "test-secret"
    validator := NewJWTValidator(secret)

    t.Run("Valid Token", func(t *testing.T) {
        // Generate a token
        token, err := GenerateTokenWithSecret(1, secret)
        assert.NoError(t, err, "Token generation should succeed")

        // Validate the token
        claims, err := validator.ValidateToken(token)
        assert.NoError(t, err, "Token validation should succeed")
        assert.Equal(t, 1, claims.UserID, "UserID should match")
        assert.WithinDuration(t, time.Now(), claims.IssuedAt.Time, time.Second, "IssuedAt should be near current time")
    })

    t.Run("Invalid Token", func(t *testing.T) {
        // Validate a random invalid token
        _, err := validator.ValidateToken("invalid.token.string")
        assert.Error(t, err, "Invalid token should result in error")
    })

    t.Run("Empty Token", func(t *testing.T) {
		// Validate an empty token
		_, err := validator.ValidateToken("")
		assert.Error(t, err, "Empty token should result in error")
		assert.Equal(t, "token is empty", err.Error(), "Error message should match")
	})

    t.Run("Expired Token", func(t *testing.T) {
        // Generate a token with a past expiration time
        claims := Claims{
            UserID: 1,
            RegisteredClaims: jwt.RegisteredClaims{
                ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Minute)),
                IssuedAt:  jwt.NewNumericDate(time.Now()),
                Issuer:    "replus-ios",
            },
        }
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
        signedToken, _ := token.SignedString([]byte(secret))

        // Validate the expired token
        _, err := validator.ValidateToken(signedToken)
        assert.Error(t, err, "Expired token should result in error")
        assert.Contains(t, err.Error(), "token is expired", "Error should indicate token expiration")
    })
}