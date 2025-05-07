package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/houseofaragon/go_project/internal/api"
	"github.com/houseofaragon/go_project/internal/store"
	"github.com/houseofaragon/go_project/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

// Creates a new app
// returning pointer to Application and error
func NewApplication() (*Application, error) {

	// add db
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	// add migrations
	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// add logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// add stores

	// add handlers
	workoutHandler := api.NewWorkoutHandler()

	// create App
	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}

	defer app.DB.Close()

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}
