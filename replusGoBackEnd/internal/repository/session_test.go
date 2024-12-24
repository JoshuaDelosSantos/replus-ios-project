package repository

import (
	"github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/models"
	"github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/testutil"
	"testing"
	"log"
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