package auth

import (
	"context"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract token from the Authorization header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// Validate token
		claims, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add userID to context
		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)

		// Pass request with updated context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}