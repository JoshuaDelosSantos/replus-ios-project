package repository

import (
	"fmt"
	"github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/models"
)

// LineRepository defines the contract for line database operations.
type LineRepository interface {
	GetLines() ([]models.Line, error)
	CreateLine(line models.Line) (models.Line, error)
	GetLinesByExerciseID(exerciseID int) ([]models.Line, error)
	UpdateLine(line models.Line) error
	DeleteLine(lineID int) error
}

// lineRepo implements LineRepository interface.
type lineRepo struct {
	db DB
}
