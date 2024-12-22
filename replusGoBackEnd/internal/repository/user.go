package repository

import (
    "database/sql"
    "github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/models"
    "fmt"
)

type UserRepository interface {
    GetUsers() ([]models.User, error)
    CreateUser(user models.User) (models.User, error)
}

type userRepo struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
    return &userRepo{db: db}
}

func (r *userRepo) GetUsers() ([]models.User, error) {
    rows, err := r.db.Query("SELECT user_id, user_name FROM users")
    if (err != nil) {
        return nil, fmt.Errorf("error querying users: %v", err)
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.UserName); err != nil {
            return nil, fmt.Errorf("error scanning user: %v", err)
        }
        users = append(users, user)
    }
    return users, nil
}

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