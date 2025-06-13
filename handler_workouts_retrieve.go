package main

import (
	"net/http"

	"github.com/Graypbj/workout_tracker/internal/auth"
)

func (cfg *apiConfig) handlerWorkoutsRetrieve(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Workouts []Workout `json:"workouts"`
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	dbWorkouts, err := cfg.db.ListWorkoutsByUser(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve workouts", err)
		return
	}

	workouts := make([]Workout, len(dbWorkouts))
	for i, dbWk := range dbWorkouts {
		workouts[i] = Workout{
			ID:          dbWk.ID,
			WorkoutDate: dbWk.WorkoutDate,
			WorkoutType: dbWk.WorkoutType,
			Notes:       dbWk.Notes.String,
			CreatedAt:   dbWk.CreatedAt,
			UpdatedAt:   dbWk.UpdatedAt,
		}
	}

	respondWithJSON(w, http.StatusOK, response{
		Workouts: workouts,
	})
}
