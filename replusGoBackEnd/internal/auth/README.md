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

