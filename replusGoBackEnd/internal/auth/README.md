# Authentication Layer Documentation

## Overview
JWT-based authentication for the Replius API. It includes middleware for protecting routes and validation of authentication tokens.

## Core Components
**1. Token Validation Interface (token_validator.go)**
Defines the contrat for token validation:
```
type TokenValidator interface {
    ValidateToken(token string) (*Claims, error)
}
```

**2. JWT Claims (jwt.go)**
Claims struct that defines the token payload:
```
type Claims struct {
    UserID int `json:"user_id"`
    jwt.RegisteredClaims
}
```