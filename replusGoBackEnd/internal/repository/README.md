# Repository Layer Documentation

## Package OVervier
Package 'repository' implements the data access layer for the Replus application.

### Common Interface

The DB interface defines a common contract for interacting with the database in the application. It abstracts core database operations, including:
- Query: Executes a query that returns multiple rows of data.
- QueryRow: Executes a query that returns a single row of data.
- Exec: Executes a query that modifies data or performs other operations without returning rows.
This interface facilitates mocking and testing by allowing database operations to be decoupled from the concrete implementation.
```
type DB interface {
    Query(query string, args ...interface{}) (*sql.Rows, error)
    QueryRow(query string, args ...interface{}) *sql.Row
    Exec(query string, args ...interface{}) (sql.Result, error)
}
```

### Repository Interfaces

The **UserRepository** interface provides methods for managing user-related data. It includes operations for:
- GetUsers: Fetching all users from the database.
- CreateUser: Adding a new user to the database.
- UpdateUser: Updating an existing user’s details.
- DeleteUser: Removing a user from the database by their I- 
```
type UserRepository interface {
    GetUsers() ([]models.User, error)
    CreateUser(user models.User) (models.User, error)
    UpdateUser(user models.User) error
    DeleteUser(userID int) error
}
```

The **SessionRepository** interface defines operations for handling session data. It includes methods for:
- GetSessions: Retrieving all sessions.
- CreateSession: Adding a new session to the database.
- GetSessionsByUserID: Fetching sessions associated with a specific user.
- UpdateSession: Modifying details of an existing session.
- DeleteSession: Removing a session from the database by its ID.
```
type SessionRepository interface {
    GetSessions() ([]models.Session, error)
    CreateSession(session models.Session) (models.Session, error)
    GetSessionsByUserID(userID int) ([]models.Session, error)
    UpdateSession(session models.Session) error
    DeleteSession(sessionID int) error
}
```

The **ExerciseRepository** interface manages exercise data operations. It includes:
- GetExercises: Retrieving all exercises.
- CreateExercise: Adding a new exercise to the database.
- GetExercisesBySessionID: Fetching exercises associated with a specific session.
- UpdateExercise: Updating details of an existing exercise.
- DeleteExercise: Removing an exercise from the database by its ID.
```
type ExerciseRepository interface {
    GetExercises() ([]models.Exercise, error)
    CreateExercise(exercise models.Exercise) (models.Exercise, error)
    GetExercisesBySessionID(sessionID int) ([]models.Exercise, error)
    UpdateExercise(exercise models.Exercise) error
    DeleteExercise(exerciseID int) error
}
```

The **LineRepository** interface provides methods for managing line-related data. It includes:
- GetLines: Retrieving all lines from the database.
- CreateLine: Adding a new line to the database.
- GetLinesByExerciseID: Fetching lines associated with a specific exercise.
- UpdateLine: Updating details of an existing line.
- DeleteLine: Removing a line from the database by its ID.
```
type LineRepository interface {
    GetLines() ([]models.Line, error)
    CreateLine(line models.Line) (models.Line, error)
    GetLinesByExerciseID(exerciseID int) ([]models.Line, error)
    UpdateLine(line models.Line) error
    DeleteLine(lineID int) error
}
```