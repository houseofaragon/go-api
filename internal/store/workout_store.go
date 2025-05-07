package store

import "database/sql"

type Workout struct {
	ID              int            `json:"id"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	DurationMinutes int            `json:"durationMinutes"`
	CaloriesBurned  int            `json:"caloriesBurned"`
	Entries         []WorkoutEntry `json:"entries"`
}

type WorkoutEntry struct {
	ID              int      `json:"id"`
	ExerciseName    string   `json:"exercise_name"`
	Sets            string   `json:"sets"`
	Reps            *int     `json:"reps"`
	DurationSeconds *int     `json:"duration_seconds"`
	Weight          *float64 `json:"weight"`
	Notes           string   `json:"notes"`
	OrderIndex      int      `json:"order_index"`
}

type PostgresWorkoutStore struct {
	db *sql.DB
}

func NewPostgresWorkoutStore(db *sql.DB) *PostgresWorkoutStore {
	return &PostgresWorkoutStore{db: db}
}

type WorkoutStore interface {
	CreateWorkout(*Workout) (*Workout, error)
	GetWorkoutByID(id int64) (*Workout, error)
}

// create CreateWorkout method
func (pg *PostgresWorkoutStore) CreateWorkout(workout *Workout) (*Workout, error) {
	// start transaction pg.db.Begin()
	tx, err := pg.db.Begin()

	// handle if error
	if err != nil {
		return nil, err
	}

	// defer rollback
	defer tx.Rollback()

	// create query to insert into workouts table
	// insert title, description, duration_minues, calories_burned
	// return id
	query :=
		`INSERT INTO workouts (title, description, duration_minues, calories_burned) 
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	// query row tx.QueryRow
	err = tx.QueryRow(query, workout.Title, workout.Description, workout.DurationMinutes, workout.CaloriesBurned).Scan(&workout.ID)
	// handle error
	if err != nil {
		return nil, err
	}

	// insert entries
	for _, entry := range workout.Entries {
		query := `
		INSERT INTO workout_entries (workout_id, exercise_name, sets, reps, duration_seconds, weight, notes, order_index)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

		err = tx.QueryRow(query, workout.ID, entry.ExerciseName, entry.Sets, entry.Reps, entry.DurationSeconds, entry.Weight, entry.Notes, entry.OrderIndex).Scan(&entry.ID)

		if err != nil {
			return nil, err
		}
	}

	// execute query tx.QueryRow
	// handle error

	// commit transaction
	// handle error
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return workout, nil
}

func (pg *PostgresWorkoutStore) GetWorkoutByID(id int64) (*Workout, error) {
	// create workout
	workout := &Workout{}
	return workout, nil
}
