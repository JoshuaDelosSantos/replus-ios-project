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

func TestCreateLine(t *testing.T) {
	log.Println("Starting TestCreateLine...")

	// Initialize mock DB and repository
	log.Println("Initializing mock DB and repository...")
	repo, mock := setupMockLineDB(t)

	// Parse the date to time.Time
	date, err := time.Parse("2006-01-02", "2021-09-01")
	assert.NoError(t, err)

	// Set up mock expectations
	log.Println("Setting up mock expectations...")
	mock.ExpectQuery(`
		INSERT INTO lines \(exercise_id, weight, reps, date\)
		VALUES \(\$1, \$2, \$3, \$4\)
		RETURNING line_id
		`).
		WithArgs(1, 100.0, 10, date).
		WillReturnRows(sqlmock.NewRows([]string{"line_id"}).AddRow(1))

	// Call the method being tested
	log.Println("Calling the CreateLine method...")
	line, err := repo.CreateLine(models.Line{ExerciseID: 1, Weight: 100.0, Reps: 10, Date: date})
	assert.NoError(t, err)

	// Verify results
	log.Println("Verifying results...")
	expected := models.Line{ID: 1, ExerciseID: 1, Weight: 100.0, Reps: 10, Date: date}
	assert.Equal(t, expected, line)

	log.Println("TestCreateLine executed successfully.")
}

func TestUpdateLine(t *testing.T) {
	log.Println("Starting TestUpdateLine...")

	// Initialize mock DB and repository
	log.Println("Initializing mock DB and repository...")
	repo, mock := setupMockLineDB(t)

	// Parse the date to time.Time
	date, err := time.Parse("2006-01-02", "2021-09-05")
	assert.NoError(t, err)

	// Set up mock expectations
	log.Println("Setting up mock expectations...")
	mock.ExpectExec(`
		UPDATE lines
		SET weight = \$1, reps = \$2, date = \$3
		WHERE line_id = \$4`).
		WithArgs(150.0, 12, date, 1).
		WillReturnResult(sqlmock.NewResult(0, 1)) // Simulate one row updated

	// Call the method being tested
	log.Println("Calling the UpdateLine method...")
	err = repo.UpdateLine(models.Line{ID: 1, Weight: 150.0, Reps: 12, Date: date})
	assert.NoError(t, err)

	// Verify expectations
	log.Println("Verifying expectations...")
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)

	log.Println("TestUpdateLine executed successfully.")
}

func TestDeleteLine(t *testing.T) {
	log.Println("Starting TestDeleteLine...")

	// Initialize mock DB and repository
	log.Println("Initializing mock DB and repository...")
	repo, mock := setupMockLineDB(t)

	// Set up mock expectations
	log.Println("Setting up mock expectations...")
	mock.ExpectExec(`
		DELETE FROM lines
		WHERE line_id = \$1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1)) // Simulate one row deleted

	// Call the method being tested
	log.Println("Calling the DeleteLine method...")
	err := repo.DeleteLine(1)
	assert.NoError(t, err)

	// Verify expectations
	log.Println("Verifying expectations...")
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)

	log.Println("TestDeleteLine executed successfully.")
}