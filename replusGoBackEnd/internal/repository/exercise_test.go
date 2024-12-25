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

func TestCreateExercise(t *testing.T) {
	log.Println("Starting TestCreateExercise...")

	// Initialize mock DB and repository
	log.Println("Initializing mock DB and repository...")
	repo, mock := setupMockExerciseDB(t)

	// Set up mock expectations
	log.Println("Setting up mock expectations...")
	mock.ExpectQuery(`
		INSERT INTO exercises \(session_id, exercise_name\)
		VALUES \(\$1, \$2\)
		RETURNING exercise_id`).
		WithArgs(1, "Exercise 1").
		WillReturnRows(sqlmock.NewRows([]string{"exercise_id"}).
			AddRow(1))

	log.Println("Mock expectations set up successfully.")

	// Call the method being tested
	log.Println("Calling the CreateExercise method...")
	exercise, err := repo.CreateExercise(models.Exercise{SessionID: 1, ExerciseName: "Exercise 1"})
	assert.NoError(t, err)

	log.Println("CreateExercise method executed successfully.")

	// Verify results
	log.Println("Verifying results...")
	expected := models.Exercise{ID: 1, SessionID: 1, ExerciseName: "Exercise 1"}
	assert.Equal(t, expected, exercise)
}