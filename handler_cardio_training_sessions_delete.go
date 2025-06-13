package main

import (
	"encoding/json"
	"net/http"

	"github.com/Graypbj/workout_tracker/internal/auth"
	"github.com/Graypbj/workout_tracker/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCardioTrainingSessionsDelete(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID uuid.UUID `json:"id"`
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate jwt", err)
		return
	}

	var params parameters
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters", err)
		return
	}

	err = cfg.db.DeleteCardioSessionByID(r.Context(), database.DeleteCardioSessionByIDParams{
		ID:     params.ID,
		UserID: userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete cardio training session", err)
		return
	}

	var a any
	respondWithJSON(w, http.StatusOK, a)
}
