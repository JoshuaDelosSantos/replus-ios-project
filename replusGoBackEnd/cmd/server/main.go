package main

import (
    "fmt"  // Package for formatted I/O
    "log"  // Package for logging
    "net/http"  // Package for HTTP client and server implementations

	"github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/config"  // Custom package for configuration
)

func main() {
    // Load configuration settings
    cfg := config.LoadConfig()

    // Print a message indicating the app is starting and on which port
    fmt.Printf("Starting app on port %s\n", cfg.AppPort)
	
    // Define a handler function for the root URL path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Welcome to Replus App!")  // Send a welcome message as the HTTP response
    })
	
    // Start the HTTP server on the specified port and log any errors
    log.Fatal(http.ListenAndServe(":"+cfg.AppPort, nil))
}