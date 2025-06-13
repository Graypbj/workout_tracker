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

func (cfg *apiConfig) handlerCardioSessionsUpdate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID       uuid.UUID `json:"id"`
		Distance float64   `json:"distance"`
		Time     string    `json:"time"`
		Notes    string    `json:"notes"`
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

	session, err := cfg.db.UpdateCardioSessionByID(r.Context(), database.UpdateCardioSessionByIDParams{
		ID:       params.ID,
		UserID:   userID,
		Distance: params.Distance,
		Time:     int64(parsedDuration),
		Notes: sql.NullString{
			String: params.Notes,
			Valid:  params.Notes != "",
		},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't update cardio session", err)
		return
	}

	parsedSessionDuration := time.Duration(session.Time)

	respondWithJSON(w, http.StatusOK, response{
		CardioTrainingSession: CardioTrainingSession{
			ID:         session.ID,
			WorkoutID:  session.WorkoutID,
			ExerciseID: session.ExerciseID,
			Distance:   session.Distance,
			Time:       parsedSessionDuration,
			Notes:      session.Notes.String,
			CreatedAt:  session.CreatedAt,
			UpdatedAt:  session.UpdatedAt,
		},
	})
}
