package routes

import (
	"github.com/go-chi/chi"
	"github.com/houseofaragon/go_project/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	// add get methods
	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutId)

	// add post methods
	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)
	return r
}
