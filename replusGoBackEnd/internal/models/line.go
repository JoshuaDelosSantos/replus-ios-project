package models

import "time"

type Line struct {
    ID          	int    `json:"line_id"`
    ExerciseID   	int    `json:"exercise_id"`
    Weight			float64 `json:"weight"`
    Reps			int `json:"reps"`
    Date			time.Time `json:"date"`
}