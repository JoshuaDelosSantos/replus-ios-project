package repository

import (
    "log"
    "testing"

    "github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/models"
    "github.com/stretchr/testify/assert"
    "github.com/DATA-DOG/go-sqlmock"
)

// setupMockDB creates a new sqlmock instance and returns the repository and mock.
func setupMockDB(t *testing.T) (UserRepository, sqlmock.Sqlmock) {
    log.Println("Setting up mock DB and repository...")
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Error creating mock DB: %v", err)
    }

    // Use the constructor for consistency
    log.Println("Mock DB and repository setup complete.")
    repo := NewUserRepository(db)
    return repo, mock
}

func TestGetUsers(t *testing.T) {
    log.Println("Starting TestGetUsers...")

    // Step 1: Initialize mock DB and repository
    log.Println("Initializing mock DB and repository...")
    repo, mock := setupMockDB(t)

    // Step 2: Set up mock expectations
    log.Println("Setting up mock expectations...")
    mock.ExpectQuery(`SELECT user_id, user_name FROM users ORDER BY user_id`).
        WillReturnRows(sqlmock.NewRows([]string{"user_id", "user_name"}).
            AddRow(1, "John Doe").
            AddRow(2, "Jane Smith"))

    log.Println("Mock expectations set up successfully.")

    // Step 3: Call the method being tested
    log.Println("Calling the GetUsers method...")
    users, err := repo.GetUsers()
    assert.NoError(t, err)

    log.Println("GetUsers method executed successfully.")

    // Step 4: Verify results
    log.Println("Verifying results...")
    expected := []models.User{
        {ID: 1, UserName: "John Doe"},
        {ID: 2, UserName: "Jane Smith"},
    }
    assert.Equal(t, expected, users)

    log.Println("Results verified successfully.")

    // Step 5: Ensure all expectations were met
    log.Println("Ensuring all mock expectations were met...")
    assert.NoError(t, mock.ExpectationsWereMet())

    log.Println("TestGetUsers completed successfully.")
}