/*
Package router handles all routing logic for the Replus API.
It uses gorilla/mux for routing and defines the following endpoints:

Root endpoints:
  - GET /           : Home page
  - GET /health     : Health check endpoint

API endpoints (v1):
  - GET    /api/v1/users     : Retrieve all users
  - POST   /api/v1/users     : Create a new user
  - GET    /api/v1/sessions  : Retrieve all sessions
  - POST   /api/v1/sessions  : Create a new session
*/
package router

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/handlers"
    "github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/auth"
    "github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/config"
)

// NewRouter initializes and returns a new router instance with all API routes configured.
// It separates routes into different sections:
// - Root routes for basic functionality
// - API routes under /api/v1 prefix for versioning
// Returns a configured *mux.Router ready to handle HTTP requests
func NewRouter(cfg config.Config) *mux.Router {
    r := mux.NewRouter()
    
    // Create JWT validator
    validator := auth.NewJWTValidator(cfg.JWT_SECRET)
    
    // Public routes
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to Replus API"))
    }).Methods("GET")
    
    r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    }).Methods("GET")

    // Protected routes
    api := r.PathPrefix("/api/v1").Subrouter()
    
    // Apply middleware with validator
    api.Use(func(next http.Handler) http.Handler {
        return auth.AuthMiddleware(validator, next)
    })

    // User routes
    api.HandleFunc("/users", handlers.GetUsers).Methods("GET")
    api.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    
    // Session routes
    api.HandleFunc("/sessions", handlers.GetSessions).Methods("GET")
    api.HandleFunc("/sessions", handlers.CreateSession).Methods("POST")

    // Exercise routes
    api.HandleFunc("/exercises", handlers.GetExercises).Methods("GET")
    api.HandleFunc("/exercises", handlers.CreateExercise).Methods("POST")
    
    return r
}