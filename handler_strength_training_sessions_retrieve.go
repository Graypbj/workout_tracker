package main

import (
	"encoding/json"
	"net/http"

	"github.com/Graypbj/workout_tracker/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerStrengthTrainingSessionsRetrieve(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		WorkoutID uuid.UUID `json:"workout_id"`
	}

	type response struct {
		StrengthTrainingSessions []StrengthTrainingSession `json:"strength_training_sessions"`
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}

	_, err = auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	var params parameters
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters", err)
		return
	}

	dbSessions, err := cfg.db.ListStrengthTrainingSessionsByWorkout(r.Context(), params.WorkoutID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve strength training sessions", err)
		return
	}

	sessions := make([]StrengthTrainingSession, len(dbSessions))
	for i, dbSess := range dbSessions {
		sessions[i] = StrengthTrainingSession{
			ID:         dbSess.ID,
			WorkoutID:  dbSess.WorkoutID,
			ExerciseID: dbSess.ExerciseID,
			Notes:      dbSess.Notes.String,
			CreatedAt:  dbSess.CreatedAt,
			UpdatedAt:  dbSess.UpdatedAt,
		}
	}

	respondWithJSON(w, http.StatusOK, response{
		StrengthTrainingSessions: sessions,
	})
}
