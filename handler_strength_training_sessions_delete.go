package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerStrengthTrainingSessionsDelete(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID uuid.UUID `json:"id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	err = cfg.db.DeleteStrengthTrainingSessionByID(r.Context(), params.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete strength training session", err)
		return
	}

	var a any
	respondWithJSON(w, http.StatusOK, a)
}
