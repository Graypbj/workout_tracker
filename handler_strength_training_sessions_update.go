package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Graypbj/workout_tracker/internal/auth"
	"github.com/Graypbj/workout_tracker/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerStrengthTrainingSessionsUpdate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID         uuid.UUID `json:"id"`
		WorkoutID  uuid.UUID `json:"workout_id"`
		ExerciseID uuid.UUID `json:"exercise_id"`
		Notes      string    `json:"notes"`
	}

	type response struct {
		StrengthTrainingSession
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

	strengthTrainingSession, err := cfg.db.UpdateStrengthTrainingSessionByID(r.Context(), database.UpdateStrengthTrainingSessionByIDParams{
		ID:         params.ID,
		WorkoutID:  params.WorkoutID,
		UserID:     userID,
		ExerciseID: params.ExerciseID,
		Notes: sql.NullString{
			String: params.Notes,
			Valid:  params.Notes != "",
		},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't update strength training session", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		StrengthTrainingSession: StrengthTrainingSession{
			ID:         strengthTrainingSession.ID,
			WorkoutID:  strengthTrainingSession.WorkoutID,
			ExerciseID: strengthTrainingSession.ExerciseID,
			Notes:      strengthTrainingSession.Notes.String,
			CreatedAt:  strengthTrainingSession.CreatedAt,
			UpdatedAt:  strengthTrainingSession.UpdatedAt,
		},
	})
}
