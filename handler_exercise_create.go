package main

import (
	"encoding/json"
	"net/http"

	"github.com/Graypbj/workout_tracker/internal/auth"
	"github.com/Graypbj/workout_tracker/internal/database"
	"github.com/google/uuid"
)

type Exercise struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	Name         string    `json:"name"`
	ExerciseType string    `json:"exercise_type"`
}

func (cfg *apiConfig) handlerExerciseCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name         string `json:"name"`
		ExerciseType string `json:"exercise_type"`
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

	exercise, err := cfg.db.CreateExercise(r.Context(), database.CreateExerciseParams{
		UserID:       userID,
		Name:         params.Name,
		ExerciseType: params.ExerciseType,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create exercise", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, Exercise{
		ID:           exercise.ID,
		UserID:       exercise.UserID,
		Name:         exercise.Name,
		ExerciseType: exercise.ExerciseType,
	})
}
