package testutil

import (
    "testing"
    "database/sql"
    "github.com/DATA-DOG/go-sqlmock"
	"log"
)

func NewMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	log.Println("Setting up mock DB and repository...")
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Error creating mock DB: %v", err)
    }
    return db, mock
}