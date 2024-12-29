package auth

import (
    "fmt"
    "log"
    "os"
)

var validatorLogger *log.Logger

func init() {
    validatorLogger = log.New(os.Stdout, "[JWT_VALIDATOR] ", log.LstdFlags|log.Lshortfile)
}

type JWTValidator struct {
    secretKey string
}

func NewJWTValidator(secretKey string) TokenValidator {
    validatorLogger.Printf("Creating new JWT validator with secret key length: %d", len(secretKey))
    return &JWTValidator{
        secretKey: secretKey,
    }
}

func (v *JWTValidator) ValidateToken(tokenString string) (*Claims, error) {
    validatorLogger.Printf("Validating token: %s...", tokenString[:10])
    
    if tokenString == "" {
        return nil, fmt.Errorf("empty token")
    }

    // Remove 'Bearer ' prefix if present
    if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
        tokenString = tokenString[7:]
    }

    // Use the existing ValidateToken function from jwt.go
    claims, err := ValidateToken(tokenString)
    if err != nil {
        validatorLogger.Printf("Token validation failed: %v", err)
        return nil, fmt.Errorf("invalid token: %w", err)
    }

    validatorLogger.Printf("Token validated successfully for UserID: %d", claims.UserID)
    return claims, nil
}