package repository

import (
	"github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/models"
	"github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/testutil"
	"testing"
	"log"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
)

func setupMockSessionDB(t *testing.T) (SessionRepository, sqlmock.Sqlmock) {
	db, mock := testutil.NewMockDB(t)
	// Use the constructor for consistency
	log.Println("Mock DB and repository setup complete.")
	repo := NewSessionRepository(db)
	return repo, mock
}

func TestGetSessions(t *testing.T) {
	log.Println("Starting TestGetSessions...")

	// Initialize mock DB and repository
	log.Println("Initializing mock DB and repository...")
	repo, mock := setupMockSessionDB(t)

	// Set up mock expectations
	log.Println("Setting up mock expectations...")
	mock.ExpectQuery(`
		SELECT session_id, user_id, session_name 
		FROM sessions
		`).
		WillReturnRows(sqlmock.NewRows([]string{"session_id", "user_id", "session_name"}).
			AddRow(1, 1, "Session 1").
			AddRow(2, 1, "Session 2"))

	log.Println("Mock expectations set up successfully.")

	// Call the method being tested
	log.Println("Calling the GetSessions method...")
	sessions, err := repo.GetSessions()
	assert.NoError(t, err)

	log.Println("GetSessions method executed successfully.")

	// Verify results
	log.Println("Verifying results...")
	expected := []models.Session{
		{ID: 1, UserID: 1, SessionName: "Session 1"},
		{ID: 2, UserID: 1, SessionName: "Session 2"},
	}
	assert.Equal(t, expected, sessions)
}

func TestGetSessionsByUserID(t *testing.T) {
	log.Println("Starting TestGetSessionsByUserID...")

	// Initialize mock DB and repository
	log.Println("Initializing mock DB and repository...")
	repo, mock := setupMockSessionDB(t)

	// Set up mock expectations for a valid query
	log.Println("Setting up mock expectations...")
	mock.ExpectQuery(`
		SELECT session_id, user_id, session_name 
		FROM sessions 
		WHERE user_id = \$1`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"session_id", "user_id", "session_name"}).
			AddRow(1, 1, "Session 1").
			AddRow(2, 1, "Session 2"))

	log.Println("Mock expectations set up successfully.")

	// Call the method being tested
	log.Println("Calling the GetSessionsByUserID method for a valid user...")
	sessions, err := repo.GetSessionsByUserID(1)
	assert.NoError(t, err)

	log.Println("GetSessionsByUserID method executed successfully.")

	// Verify results
	log.Println("Verifying results...")
	expected := []models.Session{
		{ID: 1, UserID: 1, SessionName: "Session 1"},
		{ID: 2, UserID: 1, SessionName: "Session 2"},
	}
	assert.Equal(t, expected, sessions)

	// Add test for a query error
	log.Println("Testing query error handling...")
	mock.ExpectQuery(`
		SELECT session_id, user_id, session_name 
		FROM sessions 
		WHERE user_id = \$1`).
		WithArgs(2).
		WillReturnError(fmt.Errorf("query error"))

	_, err = repo.GetSessionsByUserID(2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "query error")

	log.Println("Query error handling test passed.")

	// Add test for Scan error
	log.Println("Testing scan error handling...")
	mock.ExpectQuery(`
		SELECT session_id, user_id, session_name 
		FROM sessions 
		WHERE user_id = \$1`).
		WithArgs(3).
		WillReturnRows(sqlmock.NewRows([]string{"session_id", "user_id", "session_name"}).
			AddRow("invalid", 1, "Session 3")) // Simulate a scan error

	_, err = repo.GetSessionsByUserID(3)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error scanning session")

	log.Println("Scan error handling test passed.")

	// Ensure all expectations were met
	log.Println("Ensuring all mock expectations were met...")
	assert.NoError(t, mock.ExpectationsWereMet())

	log.Println("TestGetSessionsByUserID completed successfully.")
}

func TestCreateSession(t *testing.T) {
	log.Println("Starting TestCreateSession...")

	// Initialize mock DB and repository
	log.Println("Initializing mock DB and repository...")
	repo, mock := setupMockSessionDB(t)

	// Set up mock expectations
	log.Println("Setting up mock expectations...")
	mock.ExpectQuery(`
		INSERT INTO sessions \(user_id, session_name\) 
		VALUES \(\$1, \$2\) 
		RETURNING session_id
		`).
		WithArgs(1, "Session 1").
		WillReturnRows(sqlmock.NewRows([]string{"session_id"}).
			AddRow(1))

	log.Println("Mock expectations set up successfully.")

	// Call the method being tested
	log.Println("Calling the CreateSession method...")
	session := models.Session{UserID: 1, SessionName: "Session 1"}
	created, err := repo.CreateSession(session)
	assert.NoError(t, err)

	log.Println("CreateSession method executed successfully.")

	// Verify results
	log.Println("Verifying results...")
	assert.Equal(t, 1, created.ID)

	// Add test for a query error
	log.Println("Testing query error handling...")
	mock.ExpectQuery(`
		INSERT INTO sessions \(user_id, session_name\) 
		VALUES \(\$1, \$2\) 
		RETURNING session_id
		`).
		WithArgs(2, "Session 2").
		WillReturnError(fmt.Errorf("query error"))

	_, err = repo.CreateSession(models.Session{UserID: 2, SessionName: "Session 2"})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "query error")

	log.Println("Query error handling test passed.")

	// Ensure all expectations were met
	log.Println("Ensuring all mock expectations were met...")
	assert.NoError(t, mock.ExpectationsWereMet())

	log.Println("TestCreateSession completed successfully.")
}