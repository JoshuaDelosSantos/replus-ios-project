package main

import (
    "fmt"
    "log"
    "net/http"

	"github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/config"
)

func main() {
    cfg := config.LoadConfig()

    fmt.Printf("Starting app on port %s\n", cfg.AppPort)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Welcome to Replus App!")
    })
	
    // Here, you can pass the DB config to your database connection
    log.Fatal(http.ListenAndServe(":"+cfg.AppPort, nil))
}