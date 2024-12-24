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
    rows, err := r.db.Query(`
        SELECT session_id, user_id, session_name 
        FROM sessions 
        WHERE user_id = $1`, userID)
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

// UpdateSession updates an existing session in the database
func (r *sessionRepo) UpdateSession(session models.Session) error {
    query := `
        UPDATE sessions 
        SET session_name = $1
        WHERE session_id = $2`
    
    result, err := r.db.Exec(query, session.SessionName, session.ID)
    if err != nil {
        return fmt.Errorf("error updating session: %v", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking update result: %v", err)
    }
    if rows == 0 {
        return fmt.Errorf("session with ID %d not found", session.ID)
    }
    return nil
}

// DeleteSession removes a session from the database by ID
func (r *sessionRepo) DeleteSession(sessionID int) error {
    query := `DELETE FROM sessions WHERE session_id = $1`
    
    result, err := r.db.Exec(query, sessionID)
    if err != nil {
        return fmt.Errorf("error deleting session: %v", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking delete result: %v", err)
    }
    if rows == 0 {
        return fmt.Errorf("session with ID %d not found", sessionID)
    }
    return nil
}