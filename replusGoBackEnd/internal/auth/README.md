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

**3. Authentication Middleware (middleware.go)**
- Extracts JWT tokens from request headers.
- Validates tokens.
- Adds user information to request context.
```
// Create a validator
validator := auth.NewJWTValidator(secretKey)

// Apply middleware to routes
router.Use(auth.AuthMiddleware(validator, next))
```

**4. JWT Operations (jwt.go)**
Provdes core JWT functionality:
```
// Generate a token
token, err := GenerateToken(userID)

// Validate a token
claims, err := ValidateToken(tokenString)
```

## Workflow

1. Token Generation
When a user logs in:
```
// Generate a JWT token for the user
userID := 123 // From your login logic
token, err := auth.GenerateToken(userID)
if err != nil {
    // Handle error
}

// Return token to client
response := map[string]string{"token": token}
```

2. Client Makes Requests
Client includes token in requests

3. Middleware Validates Request
```
router := mux.NewRouter()

// Create validator with secret
validator := auth.NewJWTValidator(os.Getenv("JWT_SECRET"))

// Apply middleware to protected routes
protected := router.PathPrefix("/api").Subrouter()
protected.Use(func(next http.Handler) http.Handler {
    return auth.AuthMiddleware(validator, next)
})
```

4. Protected Handler Accesses User
```
func protectedHandler(w http.ResponseWriter, r *http.Request) {
    // Get userID from context (set by middleware)
    userID := r.Context().Value("user_id").(int)
    
    fmt.Printf("Request from user: %d\n", userID)
    // ... handle request
}
```