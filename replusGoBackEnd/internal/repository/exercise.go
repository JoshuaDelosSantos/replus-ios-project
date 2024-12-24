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

