package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Graypbj/workout_tracker/internal/auth"
	"github.com/Graypbj/workout_tracker/internal/database"
	"github.com/google/uuid"
)

type CardioTrainingSession struct {
	ID         uuid.UUID     `json:"id"`
	UserID     uuid.UUID     `json:"user_id"`
	WorkoutID  uuid.UUID     `json:"workout_id"`
	ExerciseID uuid.UUID     `json:"exercise_id"`
	Distance   float64       `json:"distance"`
	Time       time.Duration `json:"time"`
	Notes      string        `json:"notes"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
}

func (cfg *apiConfig) handlerCardioTrainingSessionsCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		WorkoutID  uuid.UUID `json:"workout_id"`
		ExerciseID uuid.UUID `json:"exercise_id"`
		Distance   float64   `json:"distance"`
		Time       string    `json:"time"`
		Notes      string    `json:"notes"`
	}

	type response struct {
		CardioTrainingSession
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
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters", err)
		return
	}

	parsedDuration, err := time.ParseDuration(params.Time)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid time format (expecting duration like '1h2m3s')", err)
		return
	}

	cardioTrainingSession, err := cfg.db.CreateCardioSession(r.Context(), database.CreateCardioSessionParams{
		UserID:     userID,
		WorkoutID:  params.WorkoutID,
		ExerciseID: params.ExerciseID,
		Distance:   params.Distance,
		Time:       int64(parsedDuration),
		Notes: sql.NullString{
			String: params.Notes,
			Valid:  params.Notes != "",
		},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create cardio training session", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, response{
		CardioTrainingSession: CardioTrainingSession{
			ID:         cardioTrainingSession.ID,
			WorkoutID:  cardioTrainingSession.WorkoutID,
			ExerciseID: cardioTrainingSession.ExerciseID,
			Distance:   cardioTrainingSession.Distance,
			Time:       time.Duration(cardioTrainingSession.Time),
			Notes:      cardioTrainingSession.Notes.String,
			CreatedAt:  cardioTrainingSession.CreatedAt,
			UpdatedAt:  cardioTrainingSession.UpdatedAt,
		},
	})
}
