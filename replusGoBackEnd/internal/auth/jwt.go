package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID int `json:"userID"`
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT for a given user
func GenerateToken(userID int) (string, error) {
	secret := os.Getenv("JWT_SECRET") // Get secret from env

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), // Token valid for 15 mins
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "replus-ios",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateToken verifies and parses a JWT
func ValidateToken(tokenString string) (*Claims, error) {
	secret := os.Getenv("JWT_SECRET")

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