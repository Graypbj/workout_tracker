package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Graypbj/workout_tracker/internal/auth"
	"github.com/Graypbj/workout_tracker/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerWorkoutsUpdate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID          uuid.UUID `json:"id"`
		WorkoutDate time.Time `json:"workout_date"`
		WorkoutType string    `json:"workout_type"`
		Notes       string    `json:"notes"`
	}

	type response struct {
		Workout
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

	workout, err := cfg.db.UpdateWorkoutByID(r.Context(), database.UpdateWorkoutByIDParams{
		ID:          params.ID,
		UserID:      userID,
		WorkoutDate: params.WorkoutDate,
		WorkoutType: params.WorkoutType,
		Notes: sql.NullString{
			String: params.Notes,
			Valid:  params.Notes != "",
		},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't update workout", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		Workout: Workout{
			ID:          workout.ID,
			WorkoutDate: workout.WorkoutDate,
			WorkoutType: workout.WorkoutType,
			Notes:       workout.Notes.String,
			CreatedAt:   workout.CreatedAt,
			UpdatedAt:   workout.UpdatedAt,
		},
	})
}
