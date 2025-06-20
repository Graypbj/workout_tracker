package main

import (
	"net/http"

	"github.com/Graypbj/workout_tracker/internal/auth"
	"github.com/Graypbj/workout_tracker/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerStrengthTrainingSetsRetrieve(w http.ResponseWriter, r *http.Request) {
	type response struct {
		StrengthTrainingSets []StrengthTrainingSet `json:"strength_training_sets"`
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

	sessionIDStr := r.URL.Query().Get("session_id")
	if sessionIDStr == "" {
		respondWithError(w, http.StatusBadRequest, "Missing session_id query parameter", err)
		return
	}

	sessionID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid session_id format", err)
		return
	}

	dbSets, err := cfg.db.ListStrengthTrainingSetsBySession(r.Context(), database.ListStrengthTrainingSetsBySessionParams{
		SessionID: sessionID,
		UserID:    userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve sets", err)
		return
	}

	sets := make([]StrengthTrainingSet, len(dbSets))
	for i, dbSet := range dbSets {
		sets[i] = StrengthTrainingSet{
			ID:        dbSet.ID,
			SessionID: dbSet.SessionID,
			SetNumber: dbSet.SetNumber,
			Reps:      dbSet.Reps,
			Weight:    dbSet.Weight,
			CreatedAt: dbSet.CreatedAt,
			UpdatedAt: dbSet.UpdatedAt,
		}
	}

	respondWithJSON(w, http.StatusOK, response{
		StrengthTrainingSets: sets,
	})
}
