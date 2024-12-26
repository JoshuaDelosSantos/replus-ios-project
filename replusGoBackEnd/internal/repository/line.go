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

func (r *lineRepo) GetLinesByExerciseID(exerciseID int) ([]models.Line, error) {
	rows, err := r.db.Query(`
		SELECT line_id, exercise_id, weight, reps, date 
		FROM lines 
		WHERE exercise_id = $1`, exerciseID)
	if err != nil {
		return nil, fmt.Errorf("error querying lines for exercise %d: %v", exerciseID, err)
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

func (r *lineRepo) CreateLine(line models.Line) (models.Line, error) {
	err := r.db.QueryRow(`
		INSERT INTO lines (exercise_id, weight, reps, date) 
		VALUES ($1, $2, $3, $4) 
		RETURNING line_id
		`, line.ExerciseID, line.Weight, line.Reps, line.Date).
		Scan(&line.ID)
	if err != nil {
		return models.Line{}, fmt.Errorf("error creating line: %v", err)
	}
	return line, nil
}

func (r *lineRepo) UpdateLine(line models.Line) error {
	query := `
		UPDATE lines
		SET weight = $1, reps = $2, date = $3
		WHERE line_id = $4`
	
	result, err := r.db.Exec(query, line.ExerciseID, line.Weight, line.Reps, line.Date, line.ID)
	if err != nil {
		return fmt.Errorf("failed to execute update query: %w", err)
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %w", err)
	}
	
	if rowsAffected == 0 {
		return fmt.Errorf("no line found with ID %d", line.ID)
	}
	
	return nil
}
func (r *lineRepo) DeleteLine(lineID int) error {
	query := `
		DELETE FROM lines
		WHERE line_id = $1`
	
	result, err := r.db.Exec(query, lineID)
	if err != nil {
		return fmt.Errorf("failed to execute delete query: %w", err)
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %w", err)
	}
	
	if rowsAffected == 0 {
		return fmt.Errorf("no line found with ID %d", lineID)
	}
	
	return nil
}