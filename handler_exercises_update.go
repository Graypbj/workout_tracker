package main

import (
	"encoding/json"
	"net/http"

	"github.com/Graypbj/workout_tracker/internal/auth"
	"github.com/Graypbj/workout_tracker/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerExercisesUpdate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID           uuid.UUID `json:"id"`
		Name         string    `json:"name"`
		ExerciseType string    `json:"exercise_type"`
	}

	type response struct {
		Exercise
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

	exercise, err := cfg.db.UpdateExerciseByID(r.Context(), database.UpdateExerciseByIDParams{
		ID:           params.ID,
		UserID:       userID,
		Name:         params.Name,
		ExerciseType: params.ExerciseType,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't update exercise", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		Exercise: Exercise{
			ID:           exercise.ID,
			Name:         exercise.Name,
			ExerciseType: exercise.ExerciseType,
			CreatedAt:    exercise.CreatedAt,
			UpdatedAt:    exercise.UpdatedAt,
		},
	})
}
