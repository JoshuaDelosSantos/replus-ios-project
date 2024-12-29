package main

import (
	"log"      // Package for logging
	"net/http" // Package for HTTP client and server implementations

	"github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/config" // Custom package for configuration
	"github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/pkg/db"          // Custom package for database operations
    "github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/router" // Custom package for routing
)

func main() {
    // Load configuration settings
    cfg := config.LoadConfig()
    log.Printf("Config loaded successfully")

    // Initialize database
    database, err := db.NewDB(cfg)
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    defer database.Close()

    // Test database connection
    if err := database.Ping(); err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }
    log.Printf("Successfully connected to database at %s:%s", cfg.DBHost, cfg.DBPort)

    r := router.NewRouter(cfg)
    // Start the HTTP server on the specified port and log any errors
    log.Printf("Starting server on port %s", cfg.AppPort)
    log.Fatal(http.ListenAndServe(":"+cfg.AppPort, r))
}