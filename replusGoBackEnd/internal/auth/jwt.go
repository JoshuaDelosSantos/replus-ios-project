package auth

import (
	"errors"
	"time"
	"os"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateTokenWithSecret generates a token using the provided secret key.
func GenerateTokenWithSecret(userID int, secret string) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "replus-ios",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateTokenWithSecret validates a token using the provided secret key.
func ValidateTokenWithSecret(tokenString string, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// Existing functions now delegate to the above functions.
func GenerateToken(userID int) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	return GenerateTokenWithSecret(userID, secret)
}

func ValidateToken(tokenString string) (*Claims, error) {
	secret := os.Getenv("JWT_SECRET")
	return ValidateTokenWithSecret(tokenString, secret)
}