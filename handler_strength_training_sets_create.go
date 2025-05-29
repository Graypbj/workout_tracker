package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Graypbj/workout_tracker/internal/auth"
	"github.com/Graypbj/workout_tracker/internal/database"
	"github.com/google/uuid"
)

type StrengthTrainingSet struct {
	ID        uuid.UUID `json:"id"`
	SessionID uuid.UUID `json:"session_id"`
	SetNumber int32     `json:"set_number"`
	Reps      int32     `json:"reps"`
	Weight    string    `json:"weight"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (cfg *apiConfig) handlerStrengthTrainingSetsCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		SessionID uuid.UUID `json:"session_id"`
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

	strengthTrainingSet, err := cfg.db.CreateStrengthTrainingSet(r.Context(), database.CreateStrengthTrainingSetParams{
		SessionID: params.SessionID,
		UserID:    userID,
		SetNumber: params.SetNumber,
		Reps:      params.Reps,
		Weight:    params.Weight,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create strength training set", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, response{
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
