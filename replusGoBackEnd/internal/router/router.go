package router

import (
    "github.com/gorilla/mux"
    "github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/handlers"
)

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