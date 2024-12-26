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

// NewLineRepository creates a new LineRepository instance.
func NewLineRepository(db DB) LineRepository {
	return &lineRepo{db: db}
}


func (r *lineRepo) GetLines() ([]models.Line, error) {
	rows, err := r.db.Query(`
		SELECT line_id, exercise_id, weight, reps, date 
		FROM lines`)
	if err != nil {
		return nil, fmt.Errorf("error querying lines: %v", err)
	}
	defer rows.Close()

	var lines []models.Line
	for rows.Next() {
		var line models.Line
		if err := rows.Scan(&line.ID, &line.ExerciseID, &line.Weight, &line.Reps, &line.Date); err != nil {
			return nil, fmt.Errorf("error scanning line: %v", err)
		}
		lines = append(lines, line)
	}
	return lines, nil
}

