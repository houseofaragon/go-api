package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type WorkoutHandler struct {
}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) HandleGetWorkoutId(w http.ResponseWriter, r *http.Request) {
	// get id from url params
	paramsWorkoutId := chi.URLParam(r, "id")

	// if not found return
	if paramsWorkoutId == "" {
		http.NotFound(w, r)
		return
	}

	// parse workout id from url param
	workoutId, err := strconv.ParseInt(paramsWorkoutId, 10, 64)

	// handle if there is an error
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// fmt.Fprintf workout id
	fmt.Fprintf(w, "Workout id is %d \n", workoutId)
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Created a new workout \n")
}
