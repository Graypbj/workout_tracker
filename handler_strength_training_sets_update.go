package main

import (
	"encoding/json"
	"net/http"

	"github.com/Graypbj/workout_tracker/internal/auth"
	"github.com/Graypbj/workout_tracker/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerStrengthTrainingSetsUpdate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID        uuid.UUID `json:"id"`
		SetNumber int32     `json:"set_number"`
		Reps      int32     `json:"reps"`
		Weight    string    `json:"weight"`
	}

	type response struct {
		StrengthTrainingSet
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

	strengthTrainingSet, err := cfg.db.UpdateStrengthTrainingSet(r.Context(), database.UpdateStrengthTrainingSetParams{
		ID:        params.ID,
		UserID:    userID,
		SetNumber: params.SetNumber,
		Reps:      params.Reps,
		Weight:    params.Weight,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't update strength training set", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		StrengthTrainingSet: StrengthTrainingSet{
			ID:        strengthTrainingSet.ID,
			SessionID: strengthTrainingSet.SessionID,
			SetNumber: strengthTrainingSet.SetNumber,
			Reps:      strengthTrainingSet.Reps,
			Weight:    strengthTrainingSet.Weight,
			CreatedAt: strengthTrainingSet.CreatedAt,
			UpdatedAt: strengthTrainingSet.UpdatedAt,
		},
	})
}
