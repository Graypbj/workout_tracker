package main

import (
	"encoding/json"
	"net/http"

	"github.com/Graypbj/workout_tracker/internal/database"
	"github.com/google/uuid"
)

type Exercise struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	ExerciseType string    `json:"exercise_type"`
	Notes        string    `json:"notes"`
}

func (cfg *apiConfig) handlerExerciseCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name         string `json:"name"`
		ExerciseType string `json:"exercise_type"`
		Notes        string `json:"notes"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	exercise, err := cfg.db.CreateExercise(r.Context(), database.CreateExerciseParams{
		Name:         params.Name,
		ExerciseType: params.ExerciseType,
		Notes:        params.Notes,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create exercise", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, Exercise{
		ID:           exercise.ID,
		Name:         exercise.Name,
		ExerciseType: exercise.ExerciseType,
		Notes:        exercise.Notes,
	})
}
