package repository

import (
    "database/sql"
    "testing"
    "github.com/JoshuaDelosSantos/replus-ios-project/replus-backend/internal/models"
)

// Mock database for testing
type mockDB struct {
    *sql.DB
    err    error
    users  []models.User
}

func TestGetUsers(t *testing.T) {
    tests := []struct {
        name    string
        db      mockDB
        want    []models.User
        wantErr bool
    }{
        {
            name: "successful retrieval",
            db: mockDB{
                users: []models.User{
                    {ID: 1, UserName: "Test User 1"},
                    {ID: 2, UserName: "Test User 2"},
                },
            },
            want: []models.User{
                {ID: 1, UserName: "Test User 1"},
                {ID: 2, UserName: "Test User 2"},
            },
            wantErr: false,
        },
        {
            name: "database error",
            db: mockDB{
                err: fmt.Errorf("database error"),
            },
            want:    nil,
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            repo := &userRepo{db: tt.db}
            got, err := repo.GetUsers()
            
            if (err != nil) != tt.wantErr {
                t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetUsers() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestCreateUser(t *testing.T) {
    tests := []struct {
        name    string
        input   models.User
        want    models.User
        wantErr bool
    }{
        {
            name:    "successful creation",
            input:   models.User{UserName: "New User"},
            want:    models.User{ID: 1, UserName: "New User"},
            wantErr: false,
        },
        {
            name:    "empty username",
            input:   models.User{UserName: ""},
            want:    models.User{},
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            repo := &userRepo{db: mockDB{}}
            got, err := repo.CreateUser(tt.input)
            
            if (err != nil) != tt.wantErr {
                t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("CreateUser() = %v, want %v", got, tt.want)
            }
        })
    }
}

