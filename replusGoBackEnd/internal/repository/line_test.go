package repository

import (
	"log"
	"time"
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
	repo, mock := setupMockLineDB(t)

	// Parse the dates to time.Time
	date1, err := time.Parse("2006-01-02", "2021-09-01")
	assert.NoError(t, err)
	date2, err := time.Parse("2006-01-02", "2021-09-02")
	assert.NoError(t, err)

	// Set up mock expectations
	mock.ExpectQuery(`
		SELECT line_id, exercise_id, weight, reps, date 
		FROM lines
		`).
		WillReturnRows(sqlmock.NewRows([]string{"line_id", "exercise_id", "weight", "reps", "date"}).
			AddRow(1, 1, 100.0, 10, date1).
			AddRow(2, 1, 200.0, 20, date2))

	// Call the method being tested
	lines, err := repo.GetLines()
	assert.NoError(t, err)

	// Verify results
	expected := []models.Line{
		{ID: 1, ExerciseID: 1, Weight: 100.0, Reps: 10, Date: date1},
		{ID: 2, ExerciseID: 1, Weight: 200.0, Reps: 20, Date: date2},
	}
	assert.Equal(t, expected, lines)
}

func TestGetLinesByExerciseID(t *testing.T) {
	log.Println("Starting TestGetLinesByExerciseID...")

	// Initialize mock DB and repository
	repo, mock := setupMockLineDB(t)

	// Parse the dates to time.Time
	date1, err := time.Parse("2006-01-02", "2021-09-01")
	assert.NoError(t, err)
	date2, err := time.Parse("2006-01-02", "2021-09-02")
	assert.NoError(t, err)

	// Set up mock expectations
	mock.ExpectQuery(`
		SELECT line_id, exercise_id, weight, reps, date 
		FROM lines 
		WHERE exercise_id = \$1`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"line_id", "exercise_id", "weight", "reps", "date"}).
			AddRow(1, 1, 100.0, 10, date1).
			AddRow(2, 1, 200.0, 20, date2))

	// Call the method being tested
	lines, err := repo.GetLinesByExerciseID(1)
	assert.NoError(t, err)

	// Verify results
	expected := []models.Line{
		{ID: 1, ExerciseID: 1, Weight: 100.0, Reps: 10, Date: date1},
		{ID: 2, ExerciseID: 1, Weight: 200.0, Reps: 20, Date: date2},
	}
	assert.Equal(t, expected, lines)
}