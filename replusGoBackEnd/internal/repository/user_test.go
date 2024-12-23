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

    // Initialize mock DB and repository
    log.Println("Initializing mock DB and repository...")
    repo, mock := setupMockDB(t)

    // Set up mock expectations
    log.Println("Setting up mock expectations...")
    mock.ExpectQuery(`
		SELECT user_id, user_name 
		FROM users 
		ORDER BY user_id
		`).
        WillReturnRows(sqlmock.NewRows([]string{"user_id", "user_name"}).
            AddRow(1, "John Doe").
            AddRow(2, "Jane Smith"))

    log.Println("Mock expectations set up successfully.")

    // Call the method being tested
    log.Println("Calling the GetUsers method...")
    users, err := repo.GetUsers()
    assert.NoError(t, err)

    log.Println("GetUsers method executed successfully.")

    // Verify results
    log.Println("Verifying results...")
	log.Println(users)
    expected := []models.User{
        {ID: 1, UserName: "John Doe"},
        {ID: 2, UserName: "Jane Smith"},
    }
    assert.Equal(t, expected, users)

    log.Println("Results verified successfully.")

    // Ensure all expectations were met
    log.Println("Ensuring all mock expectations were met...")
    assert.NoError(t, mock.ExpectationsWereMet())

    log.Println("TestGetUsers completed successfully.")
}

func TestCreateUser(t *testing.T) {
	log.Println("Starting TestCreateUser...")

	// Initialize mock DB and repository
	log.Println("Initializing mock DB and repository...")
	repo, mock := setupMockDB(t)

	// Set up mock expectations
	log.Println("Setting up mock expectations...")
	mock.ExpectQuery(`
		INSERT INTO users \(user_name\) 
		VALUES \(\$1\) 
		RETURNING user_id
		`).
		WithArgs("John Doe").
		WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(1))

	log.Println("Mock expectations set up successfully.")

	// Call the method being tested
	log.Println("Calling the CreateUser method...")
	user := models.User{UserName: "John Doe"}
	createdUser, err := repo.CreateUser(user)
	assert.NoError(t, err)

	log.Println("CreateUser method executed successfully.")

	// Verify results
	log.Println("Verifying results...")
	expected := models.User{ID: 1, UserName: "John Doe"}
	assert.Equal(t, expected, createdUser)

	log.Println("Results verified successfully.")

	// Ensure all expectations were met
	log.Println("Ensuring all mock expectations were met...")
	assert.NoError(t, mock.ExpectationsWereMet())

	log.Println("TestCreateUser completed successfully.")
}

func TestUpdateUser(t *testing.T) {
	log.Println("Starting TestUpdateUser...")

	// Initialize mock DB and repository
	log.Println("Initializing mock DB and repository...")
	repo, mock := setupMockDB(t)

	// Set up mock expectations
	log.Println("Setting up mock expectations...")
	mock.ExpectExec(`
		UPDATE users 
		SET user_name = \$1 
		WHERE user_id = \$2
		`).
		WithArgs("Jane Smith", 2).
		WillReturnResult(sqlmock.NewResult(0, 1))

	log.Println("Mock expectations set up successfully.")

	// Call the method being tested
	log.Println("Calling the UpdateUser method...")
	user := models.User{ID: 2, UserName: "Jane Smith"}
	err := repo.UpdateUser(user)
	assert.NoError(t, err)

	log.Println("UpdateUser method executed successfully.")

	// Ensure all expectations were met
	log.Println("Ensuring all mock expectations were met...")
	assert.NoError(t, mock.ExpectationsWereMet())

	log.Println("TestUpdateUser completed successfully.")
}

func TestDeleteUser(t *testing.T) {
	log.Println("Starting TestDeleteUser...")

	// Initialize mock DB and repository
	log.Println("Initializing mock DB and repository...")
	repo, mock := setupMockDB(t)

	// Set up mock expectations
	log.Println("Setting up mock expectations...")
	mock.ExpectExec(`
		DELETE FROM users 
		WHERE user_id = \$1
		`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	log.Println("Mock expectations set up successfully.")

	// Call the method being tested
	log.Println("Calling the DeleteUser method...")
	err := repo.DeleteUser(1)
	assert.NoError(t, err)

	log.Println("DeleteUser method executed successfully.")

	// Ensure all expectations were met
	log.Println("Ensuring all mock expectations were met...")
	assert.NoError(t, mock.ExpectationsWereMet())

	log.Println("TestDeleteUser completed successfully.")
}