package repository

import (
	"log"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/models"
	"github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/testutil"
)

func setupMockLineDB(t *testing.T) (LineRepository, sqlmock.Sqlmock) {
	db, mock := testutil.NewMockDB(t)
	// Use the constructor for consistency
	log.Println("Mock DB and repository setup complete.")
	repo := NewLineRepository(db)
	return repo, mock
}

func TestGetLines(t *testing.T) {
	log.Println("Starting TestGetLines...")

	// Initialize mock DB and repository
	log.Println("Initializing mock DB and repository...")
	repo, mock := setupMockLineDB(t)

	// Set up mock expectations
	log.Println("Setting up mock expectations...")
	mock.ExpectQuery(`
		SELECT line_id, exercise_id, weight, reps, date 
		FROM lines
		`).
		WillReturnRows(sqlmock.NewRows([]string{"line_id", "exercise_id", "weight", "reps", "date"}).
			AddRow(1, 1, 100, 10, "2021-09-01").
			AddRow(2, 1, 200, 20, "2021-09-02"))

	log.Println("Mock expectations set up successfully.")

	// Call the method being tested
	log.Println("Calling the GetLines method...")
	lines, err := repo.GetLines()
	assert.NoError(t, err)

	log.Println("GetLines method executed successfully.")

	// Verify results
	log.Println("Verifying results...")
	expected := []models.Line{
		{ID: 1, ExerciseID: 1, Weight: 100, Reps: 10, Date: "2021-09-01"},
		{ID: 2, ExerciseID: 1, Weight: 200, Reps: 20, Date: "2021-09-02"},
	}
	assert.Equal(t, expected, lines)
}