package testutil

import (
    "testing"
    "database/sql"
    "github.com/DATA-DOG/go-sqlmock"
)

func NewMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Error creating mock DB: %v", err)
    }
    return db, mock
}