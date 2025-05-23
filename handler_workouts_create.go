package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Graypbj/workout_tracker/internal/auth"
	"github.com/Graypbj/workout_tracker/internal/database"
)

type Workout struct {
	WorkoutDate time.Time `json:"workout_date"`
	WorkoutType string    `json:"workout_type"`
	Notes       string    `json:"notes"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (cfg *apiConfig) handlerWorkoutsCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		WorkoutType string `json:"workout_type"`
		Notes       string `json:"notes"`
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

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	workout, err := cfg.db.CreateWorkout(r.Context(), database.CreateWorkoutParams{
		UserID:      userID,
		WorkoutType: params.WorkoutType,
		Notes: sql.NullString{
			String: params.Notes,
			Valid:  params.Notes != "",
		},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create workout", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, Workout{
		WorkoutDate: workout.WorkoutDate,
		WorkoutType: workout.WorkoutType,
		Notes:       workout.Notes.String,
		CreatedAt:   workout.CreatedAt,
		UpdatedAt:   workout.UpdatedAt,
	})
}
