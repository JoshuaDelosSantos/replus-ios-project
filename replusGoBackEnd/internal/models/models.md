# Models Documentation

## Overview
This package contains the core data models used throughout the Replus application. These models reflect the database schema and are used for data transfer between layers.

## Models

### User
Represents a user in the system.
```
type User struct {
    ID       int    `json:"user_id"`     // Unique identifier
    UserName string `json:"user_name"`    // User's name
}
```

### Session
Represents a workout session belonging to a user.
```
type Session struct {
    ID          int    `json:"session_id"`    // Unique identifier
    UserID      int    `json:"user_id"`       // References User.ID
    SessionName string `json:"session_name"`   // Name of the session
}
```

### Exercise
Represents an exercise within a session.
```
type Exercise struct {
    ID            int    `json:"exercise_id"`    // Unique identifier
    SessionID     int    `json:"session_id"`     // References Session.ID
    ExerciseName  string `json:"exercise_name"`  // Name of the exercise
}
```

### Line
Represents a set within an exercise with weight, reps and date/timing information.
```
type Line struct {
    ID          int       `json:"line_id"`      // Unique identifier
    ExerciseID  int       `json:"exercise_id"`  // References Exercise.ID
    Weight      float64   `json:"weight"`       // Weight used (in kg/lbs)
    Reps        int       `json:"reps"`         // Number of repetitions
    Date        time.Time `json:"date"`         // When the set was performed
}
```

## Relationships
- User (1) -> (*)Session
- Session (1) -> (*)Exercise
- Exercise (1) -> (*)Line

## Usage example
```
// Creating a new user
user := models.User{
    UserName: "John Doe",
}

// Creating a session for the user
session := models.Session{
    UserID: user.ID,
    SessionName: "Monday Workout",
}

// Creating an exercise in the session
exercise := models.Exercise{
    SessionID: session.ID,
    ExerciseName: "Bench Press",
}

// Recording a set
line := models.Line{
    ExerciseID: exercise.ID,
    Weight: 100.5,
    Reps: 8,
    Date: time.Now(),
}
```