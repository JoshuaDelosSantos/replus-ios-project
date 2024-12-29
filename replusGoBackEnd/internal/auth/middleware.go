package auth

import (
	"context"
	"log"
	"net/http"
	"os"
)

type TokenValidator interface {
	ValidateToken(token string) (*Claims, error)
}

var middlewareLogger *log.Logger

func init() {
	middlewareLogger = log.New(os.Stdout, "[AUTH_MIDDLEWARE] ", log.LstdFlags|log.Lshortfile)
}

func AuthMiddleware(validator TokenValidator, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middlewareLogger.Printf("Processing request: %s %s", r.Method, r.URL.Path)

		// Extract token from the Authorization header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			middlewareLogger.Printf("Missing token in request from IP: %s", r.RemoteAddr)
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}
		middlewareLogger.Println("Token found in request header")

		// Validate token using the injected validator
		claims, err := validator.ValidateToken(tokenString)
		if err != nil {
			middlewareLogger.Printf("Token validation failed: %v", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		middlewareLogger.Printf("Token validated successfully for UserID: %d", claims.UserID)

		// Add userID to context
		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		middlewareLogger.Printf("Added UserID %d to request context", claims.UserID)

		// Pass request with updated context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
		middlewareLogger.Printf("Request completed for UserID: %d", claims.UserID)
	})
}