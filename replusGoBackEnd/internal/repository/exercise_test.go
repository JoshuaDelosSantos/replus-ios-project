package repository

import (
	"log"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/models"
	"github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/testutil"
)

func setupMockExerciseDB(t *testing.T) (ExerciseRepository, sqlmock.Sqlmock) {
	db, mock := testutil.NewMockDB(t)
	// Use the constructor for consistency
	log.Println("Mock DB and repository setup complete.")
	repo := NewExerciseRepository(db)
	return repo, mock
}

func TestGetExercises(t *testing.T) {
	log.Println("Starting TestGetExercises...")

	// Initialize mock DB and repository
	log.Println("Initializing mock DB and repository...")
	repo, mock := setupMockExerciseDB(t)

	// Set up mock expectations
	log.Println("Setting up mock expectations...")
	mock.ExpectQuery(`
		SELECT exercise_id, session_id, exercise_name 
		FROM exercises
		`).
		WillReturnRows(sqlmock.NewRows([]string{"exercise_id", "session_id", "exercise_name"}).
			AddRow(1, 1, "Exercise 1").
			AddRow(2, 1, "Exercise 2"))

	log.Println("Mock expectations set up successfully.")

	// Call the method being tested
	log.Println("Calling the GetExercises method...")
	exercises, err := repo.GetExercises()
	assert.NoError(t, err)

	log.Println("GetExercises method executed successfully.")

	// Verify results
	log.Println("Verifying results...")
	expected := []models.Exercise{
		{ID: 1, SessionID: 1, ExerciseName: "Exercise 1"},
		{ID: 2, SessionID: 1, ExerciseName: "Exercise 2"},
	}
	assert.Equal(t, expected, exercises)
}

func TestGetExercisesBySessionID(t *testing.T) {
	log.Println("Starting TestGetExercisesBySessionID...")

	// Initialize mock DB and repository
	log.Println("Initializing mock DB and repository...")
	repo, mock := setupMockExerciseDB(t)

	// Set up mock expectations
	log.Println("Setting up mock expectations...")
	mock.ExpectQuery(`
		SELECT exercise_id, session_id, exercise_name 
		FROM exercises 
		WHERE session_id = \$1`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"exercise_id", "session_id", "exercise_name"}).
			AddRow(1, 1, "Exercise 1").
			AddRow(2, 1, "Exercise 2"))

	log.Println("Mock expectations set up successfully.")

	// Call the method being tested
	log.Println("Calling the GetExercisesBySessionID method...")
	exercises, err := repo.GetExercisesBySessionID(1)
	assert.NoError(t, err)

	log.Println("GetExercisesBySessionID method executed successfully.")

	// Verify results
	log.Println("Verifying results...")
	expected := []models.Exercise{
		{ID: 1, SessionID: 1, ExerciseName: "Exercise 1"},
		{ID: 2, SessionID: 1, ExerciseName: "Exercise 2"},
	}
	assert.Equal(t, expected, exercises)
}

func TestUpdateExercise(t *testing.T) {
	log.Println("Starting TestUpdateExercise...")

	// Initialize mock DB and repository
	log.Println("Initializing mock DB and repository...")
	repo, mock := setupMockExerciseDB(t)

	// Set up mock expectations
	log.Println("Setting up mock expectations...")
	mock.ExpectExec(`
		UPDATE exercises
		SET exercise_name = \$1
		WHERE exercise_id = \$2`).
		WithArgs("Exercise 2", 1).
		WillReturnResult(sqlmock.NewResult(0, 1)) // 0 for last insert ID, 1 for rows affected

	log.Println("Mock expectations set up successfully.")

	// Call the method being tested
	log.Println("Calling the UpdateExercise method...")
	err := repo.UpdateExercise(models.Exercise{ID: 1, ExerciseName: "Exercise 2"})
	assert.NoError(t, err)

	log.Println("UpdateExercise method executed successfully.")

	// Ensure all mock expectations were met
	log.Println("Ensuring all mock expectations were met...")
	assert.NoError(t, mock.ExpectationsWereMet())

	log.Println("TestUpdateExercise completed successfully.")
}