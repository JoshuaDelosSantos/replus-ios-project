package models

type Exercise struct {
    ID          	int    `json:"exercise_id"`
    SessionID   	int    `json:"session_id"`
    ExerciseName	string `json:"exercise_name"`
}