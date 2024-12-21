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
    "github.com/gorilla/mux"
    "github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/handlers"
)

// NewRouter initializes and returns a new router instance with all API routes configured.
// It separates routes into different sections:
// - Root routes for basic functionality
// - API routes under /api/v1 prefix for versioning
// Returns a configured *mux.Router ready to handle HTTP requests
func NewRouter() *mux.Router {
    r := mux.NewRouter()
    
    // Root route
    r.HandleFunc("/", handlers.Home).Methods("GET")
    
    // Health check endpoint
    r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
    
    // API routes
    api := r.PathPrefix("/api/v1").Subrouter()
    
    // User routes
    api.HandleFunc("/users", handlers.GetUsers).Methods("GET")
    api.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    
    // Session routes
    api.HandleFunc("/sessions", handlers.GetSessions).Methods("GET")
    api.HandleFunc("/sessions", handlers.CreateSession).Methods("POST")
    
    return r
}