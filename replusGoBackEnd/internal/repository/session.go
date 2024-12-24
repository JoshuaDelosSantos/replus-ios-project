package repository

import (
	"fmt"
	"github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/models"
)

// SessionRepository defines the contract for session database operations.
type SessionRepository interface {
	GetSessions() ([]models.Session, error)
	CreateSession(session models.Session) (models.Session, error)
	GetSessionsByUserID(userID int) ([]models.Session, error)
}

// sessionRepo implements SessionRepository interface.
type sessionRepo struct {
	db DB
}

// NewSessionRepository creates a new SessionRepository instance.
func NewSessionRepository(db DB) SessionRepository {
	return &sessionRepo{db: db}
}

// GetSessions retrieves all sessions from the database
func (r *sessionRepo) GetSessions() ([]models.Session, error) {
    rows, err := r.db.Query(`
        SELECT session_id, user_id, session_name 
        FROM sessions`)
    if err != nil {
        return nil, fmt.Errorf("error querying sessions: %v", err)
    }
    defer rows.Close()

    var sessions []models.Session
    for rows.Next() {
        var session models.Session
        if err := rows.Scan(&session.ID, &session.UserID, &session.SessionName); err != nil {
            return nil, fmt.Errorf("error scanning session: %v", err)
        }
        sessions = append(sessions, session)
    }
    return sessions, nil
}

// GetSessionsByUserID retrieves all sessions for a specific user
func (r *sessionRepo) GetSessionsByUserID(userID int) ([]models.Session, error) {
    rows, err := r.db.Query("SELECT session_id, user_id, session_name FROM sessions WHERE user_id = $1", userID)
    if err != nil {
        return nil, fmt.Errorf("error querying sessions for user %d: %v", userID, err)
    }
    defer rows.Close()

    var sessions []models.Session
    for rows.Next() {
        var session models.Session
        if err := rows.Scan(&session.ID, &session.UserID, &session.SessionName); err != nil {
            return nil, fmt.Errorf("error scanning session: %v", err)
        }
        sessions = append(sessions, session)
    }
    return sessions, nil
}

// CreateSession inserts a new session into the database
func (r *sessionRepo) CreateSession(session models.Session) (models.Session, error) {
    query := `
        INSERT INTO sessions (user_id, session_name)
        VALUES ($1, $2)
        RETURNING session_id`
    
    err := r.db.QueryRow(query, session.UserID, session.SessionName).Scan(&session.ID)
    if err != nil {
        return models.Session{}, fmt.Errorf("error creating session: %v", err)
    }
    
    return session, nil
}