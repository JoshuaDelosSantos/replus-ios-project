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

