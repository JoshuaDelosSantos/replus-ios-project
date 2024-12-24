package repository


import (
	"fmt"
	"github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/models"
)

// ExerciseRepository defines the contract for exercise database operations.
type ExerciseRepository interface {
	GetExercises() ([]models.Exercise, error)
	CreateExercise(exercise models.Exercise) (models.Exercise, error)
	GetExercisesBySessionID(sessionID int) ([]models.Exercise, error)
	UpdateExercise(exercise models.Exercise) error
	DeleteExercise(exerciseID int) error
}

// exerciseRepo implements ExerciseRepository interface.
type exerciseRepo struct {
	db DB
}

// NewExerciseRepository creates a new ExerciseRepository instance.
func NewExerciseRepository(db DB) ExerciseRepository {
	return &exerciseRepo{db: db}
}

// GetExercises retrieves all exercises from the database
func (r *exerciseRepo) GetExercises() ([]models.Exercise, error) {
	rows, err := r.db.Query(`
		SELECT exercise_id, session_id, exercise_name 
		FROM exercises`)
	if err != nil {
		return nil, fmt.Errorf("error querying exercises: %v", err)
	}
	defer rows.Close()

	var exercises []models.Exercise
	for rows.Next() {
		var exercise models.Exercise
		if err := rows.Scan(&exercise.ID, &exercise.SessionID, &exercise.ExerciseName); err != nil {
			return nil, fmt.Errorf("error scanning exercise: %v", err)
		}
		exercises = append(exercises, exercise)
	}
	return exercises, nil
}

// GetExercisesBySessionID retrieves all exercises for a specific session
func (r *exerciseRepo) GetExercisesBySessionID(sessionID int) ([]models.Exercise, error) {
	rows, err := r.db.Query(`
		SELECT exercise_id, session_id, exercise_name 
		FROM exercises 
		WHERE session_id = $1`, sessionID)
	if err != nil {
		return nil, fmt.Errorf("error querying exercises for session %d: %v", sessionID, err)
	}
	defer rows.Close()

	var exercises []models.Exercise
	for rows.Next() {
		var exercise models.Exercise
		if err := rows.Scan(&exercise.ID, &exercise.SessionID, &exercise.ExerciseName); err != nil {
			return nil, fmt.Errorf("error scanning exercise: %v", err)
		}
		exercises = append(exercises, exercise)
	}
	return exercises, nil
}

// CreateExercise inserts a new exercise into the database.
func (r *exerciseRepo) CreateExercise(exercise models.Exercise) (models.Exercise, error) {
	query := `
		INSERT INTO exercises (session_id, exercise_name)
		VALUES ($1, $2)
		RETURNING exercise_id`

	err := r.db.QueryRow(query, exercise.SessionID, exercise.ExerciseName).Scan(&exercise.ID)
	if err != nil {
		return models.Exercise{}, fmt.Errorf("error creating exercise: %v", err)
	}

	return exercise, nil
}

// UpdateExercise updates an existing exercise in the database.
func (r *exerciseRepo) UpdateExercise(exercise models.Exercise) error {
	query := `
		UPDATE exercises 
		SET exercise_name = $1
		WHERE exercise_id = $2`
	
	result, err := r.db.Exec(query, exercise.ExerciseName, exercise.ID)
	if err != nil {
		return fmt.Errorf("error updating exercise: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking update result: %v", err)
	}
	if rows == 0 {
		return fmt.Errorf("exercise with ID %d not found", exercise.ID)
	}
	return nil
}

// DeleteExercise removes an existing exercise from the database.
func (r *exerciseRepo) DeleteExercise(exerciseID int) error {
	query := `DELETE FROM exercises WHERE exercise_id = $1`
	
	result, err := r.db.Exec(query, exerciseID)
	if err != nil {
		return fmt.Errorf("error deleting exercise: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking delete result: %v", err)
	}
	if rows == 0 {
		return fmt.Errorf("exercise with ID %d not found", exerciseID)
	}
	return nil
}