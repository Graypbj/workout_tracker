package main

import (
	"net/http"
	"time"

	"github.com/Graypbj/workout_tracker/internal/auth"
	"github.com/Graypbj/workout_tracker/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCardioTrainingSessionsRetrieve(w http.ResponseWriter, r *http.Request) {
	type response struct {
		CardioTrainingSessions []CardioTrainingSession `json:"cardio_training_sessions"`
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

	workoutIDStr := r.URL.Query().Get("workout_id")
	if workoutIDStr == "" {
		respondWithError(w, http.StatusBadRequest, "Missing workout_id query parameter", nil)
		return
	}

	workoutID, err := uuid.Parse(workoutIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid workout_id format", err)
		return
	}

	dbSessions, err := cfg.db.ListCardioSessionsByWorkout(r.Context(), database.ListCardioSessionsByWorkoutParams{
		WorkoutID: workoutID,
		UserID:    userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve cardio sessions", err)
		return
	}

	sessions := make([]CardioTrainingSession, len(dbSessions))
	for i, session := range dbSessions {
		sessions[i] = CardioTrainingSession{
			ID:         session.ID,
			WorkoutID:  session.WorkoutID,
			ExerciseID: session.ExerciseID,
			Distance:   session.Distance,
			Time:       time.Duration(session.Time),
			Notes:      session.Notes.String,
			CreatedAt:  session.CreatedAt,
			UpdatedAt:  session.UpdatedAt,
		}
	}

	respondWithJSON(w, http.StatusOK, response{
		CardioTrainingSessions: sessions,
	})
}
