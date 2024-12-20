package main

import (
	"fmt"      // Package for formatted I/O
	"log"      // Package for logging
	"net/http" // Package for HTTP client and server implementations

	"github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/config" // Custom package for configuration
	"github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/pkg/db"          // Custom package for database operations
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

    // Print a message indicating the app is starting and on which port
    fmt.Printf("Starting app on port %s\n", cfg.AppPort)
	
    // Define a handler function for the root URL path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Welcome to Replus App!")  // Send a welcome message as the HTTP response
    })
	
    // Start the HTTP server on the specified port and log any errors
    log.Fatal(http.ListenAndServe(":"+cfg.AppPort, nil))
}