package db

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/config"
)

func NewDB(cfg config.Config) (*sql.DB, error) {
    connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DBHost, 
        cfg.DBPort, 
        cfg.DBUser, 
        cfg.DBPassword, 
        cfg.DBName,
    )
    
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("error opening database: %v", err)
    }

    if err = db.Ping(); err != nil {
        return nil, fmt.Errorf("error connecting to database: %v", err)
    }

    return db, nil
}