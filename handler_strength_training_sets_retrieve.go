package main

import (
	"encoding/json"
	"net/http"

	"github.com/Graypbj/workout_tracker/internal/auth"
	"github.com/Graypbj/workout_tracker/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerStrengthTrainingSetsRetrieve(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		SessionID uuid.UUID `json:"session_id"`
	}

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

	var params parameters
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	dbSets, err := cfg.db.ListStrengthTrainingSetsBySession(r.Context(), database.ListStrengthTrainingSetsBySessionParams{
		SessionID: params.SessionID,
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
