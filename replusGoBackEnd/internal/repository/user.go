// Package repository implements the data access layer for the application.
// It handles all database operations and abstracts them from the business logic.
package repository

import (
    "database/sql"
    "github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/models"
    "fmt"
)

// UserRepository defines the contract for user database operations.
// Interface segregation principle: clients depend only on methods they use.
type UserRepository interface {
    GetUsers() ([]models.User, error)
    CreateUser(user models.User) (models.User, error)
}

// userRepo implements UserRepository interface.
// It holds a database connection to perform operations.
type userRepo struct {
    db *sql.DB // Dependency injection: db connection is passed from outside
}

// NewUserRepository creates a new UserRepository instance.
// This is a constructor function following factory pattern in Go.
func NewUserRepository(db *sql.DB) UserRepository {
    return &userRepo{db: db} // Returns interface, hides concrete implementation
}

// GetUsers retrieves all users from the database.
// Returns a slice of User models and error if any.
func (r *userRepo) GetUsers() ([]models.User, error) {
    // Execute SQL query and handle potential errors
    rows, err := r.db.Query("SELECT user_id, user_name FROM users")
    if (err != nil) {
        return nil, fmt.Errorf("error querying users: %v", err)
    }
    // Ensure rows are closed after function returns
    defer rows.Close()

    // Initialize slice to store users
    var users []models.User
    // Iterate through result rows
    for rows.Next() {
        var user models.User
        // Scan row values into user struct fields
        if err := rows.Scan(&user.ID, &user.UserName); err != nil {
            return nil, fmt.Errorf("error scanning user: %v", err)
        }
        // Append user to slice using built-in append function
        users = append(users, user)
    }
    return users, nil
}

// CreateUser inserts a new user into the database.
// Returns the created User model and error if any.
func (r *userRepo) CreateUser(user models.User) (models.User, error) {
    query := `
        INSERT INTO users (user_name)
        VALUES ($1)
        RETURNING user_id`
    
    err := r.db.QueryRow(query, user.UserName).Scan(&user.ID)
    if err != nil {
        return models.User{}, fmt.Errorf("error creating user: %v", err)
    }
    
    return user, nil
}