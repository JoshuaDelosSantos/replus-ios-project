package repository

import (
    "testing"

    "github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/models"
    "github.com/stretchr/testify/assert"
    "github.com/DATA-DOG/go-sqlmock"
)

// setupMockDB creates a new sqlmock instance and returns the repository and mock.
func setupMockDB(t *testing.T) (UserRepository, sqlmock.Sqlmock) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("error creating mock DB: %v", err)
    }

    // Use the constructor for consistency
    repo := NewUserRepository(db)
    return repo, mock
}

func TestGetUsers(t *testing.T) {
    // Step 1: Initialize mock DB and repository
    repo, mock := setupMockDB(t)

    // Step 2: Set up mock expectations
    mock.ExpectQuery(`SELECT user_id, user_name FROM users ORDER BY user_id`).
        WillReturnRows(sqlmock.NewRows([]string{"user_id", "user_name"}).
            AddRow(1, "John Doe").
            AddRow(2, "Jane Smith"))

    // Step 3: Call the method being tested
    users, err := repo.GetUsers()
    assert.NoError(t, err)

    // Step 4: Verify results
    expected := []models.User{
        {ID: 1, UserName: "John Doe"},
        {ID: 2, UserName: "Jane Smith"},
    }
    assert.Equal(t, expected, users)

    // Step 5: Ensure all expectations were met
    assert.NoError(t, mock.ExpectationsWereMet())
}