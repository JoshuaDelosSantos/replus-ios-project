package models

type Session struct {
    ID          int    `json:"session_id"`
    UserID      int    `json:"user_id"`
    SessionName string `json:"session_name"`
}